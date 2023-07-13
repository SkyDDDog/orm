package session

import (
	"fmt"
	"learning-orm/log"
	"learning-orm/schema"
	"reflect"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	// nil or different model, update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}

	desc := strings.Join(columns, ",")
	rawSql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s);", table.Name, desc)
	_, err := s.Raw(rawSql).Exec()
	return err
}

func (s *Session) DropTable() error {
	rawSql := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", s.RefTable().Name)
	_, err := s.Raw(rawSql).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
