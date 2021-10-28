package gorm_mptt

import (
	"reflect"
)

func (db *Tree) GetNode(o interface{}) interface{} {
	rv := reflect.ValueOf(o).Elem()
	id := rv.FieldByName("ID").String()

	db.Statement.First(&o, map[string]interface{}{"id": id})
	return o

}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
