// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableMeta interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *TableMeta
	GetCatalogName() *string
	SetDatabaseName(v string) *TableMeta
	GetDatabaseName() *string
	SetTableName(v string) *TableMeta
	GetTableName() *string
}

type TableMeta struct {
	CatalogName  *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	TableName    *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s TableMeta) String() string {
	return dara.Prettify(s)
}

func (s TableMeta) GoString() string {
	return s.String()
}

func (s *TableMeta) GetCatalogName() *string {
	return s.CatalogName
}

func (s *TableMeta) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *TableMeta) GetTableName() *string {
	return s.TableName
}

func (s *TableMeta) SetCatalogName(v string) *TableMeta {
	s.CatalogName = &v
	return s
}

func (s *TableMeta) SetDatabaseName(v string) *TableMeta {
	s.DatabaseName = &v
	return s
}

func (s *TableMeta) SetTableName(v string) *TableMeta {
	s.TableName = &v
	return s
}

func (s *TableMeta) Validate() error {
	return dara.Validate(s)
}
