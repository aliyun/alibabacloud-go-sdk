// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaCategoryTableEntity interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *MetaCategoryTableEntity
	GetCatalogName() *string
	SetCategoryId(v int64) *MetaCategoryTableEntity
	GetCategoryId() *int64
	SetDatabaseSearchName(v string) *MetaCategoryTableEntity
	GetDatabaseSearchName() *string
	SetDbId(v int32) *MetaCategoryTableEntity
	GetDbId() *int32
	SetDbType(v string) *MetaCategoryTableEntity
	GetDbType() *string
	SetDescription(v string) *MetaCategoryTableEntity
	GetDescription() *string
	SetInstanceId(v int32) *MetaCategoryTableEntity
	GetInstanceId() *int32
	SetSchemaName(v string) *MetaCategoryTableEntity
	GetSchemaName() *string
	SetTableName(v string) *MetaCategoryTableEntity
	GetTableName() *string
	SetTableSchemaName(v string) *MetaCategoryTableEntity
	GetTableSchemaName() *string
}

type MetaCategoryTableEntity struct {
	CatalogName        *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	CategoryId         *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	DatabaseSearchName *string `json:"DatabaseSearchName,omitempty" xml:"DatabaseSearchName,omitempty"`
	DbId               *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbType             *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId         *int32  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SchemaName         *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName          *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableSchemaName    *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
}

func (s MetaCategoryTableEntity) String() string {
	return dara.Prettify(s)
}

func (s MetaCategoryTableEntity) GoString() string {
	return s.String()
}

func (s *MetaCategoryTableEntity) GetCatalogName() *string {
	return s.CatalogName
}

func (s *MetaCategoryTableEntity) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *MetaCategoryTableEntity) GetDatabaseSearchName() *string {
	return s.DatabaseSearchName
}

func (s *MetaCategoryTableEntity) GetDbId() *int32 {
	return s.DbId
}

func (s *MetaCategoryTableEntity) GetDbType() *string {
	return s.DbType
}

func (s *MetaCategoryTableEntity) GetDescription() *string {
	return s.Description
}

func (s *MetaCategoryTableEntity) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *MetaCategoryTableEntity) GetSchemaName() *string {
	return s.SchemaName
}

func (s *MetaCategoryTableEntity) GetTableName() *string {
	return s.TableName
}

func (s *MetaCategoryTableEntity) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *MetaCategoryTableEntity) SetCatalogName(v string) *MetaCategoryTableEntity {
	s.CatalogName = &v
	return s
}

func (s *MetaCategoryTableEntity) SetCategoryId(v int64) *MetaCategoryTableEntity {
	s.CategoryId = &v
	return s
}

func (s *MetaCategoryTableEntity) SetDatabaseSearchName(v string) *MetaCategoryTableEntity {
	s.DatabaseSearchName = &v
	return s
}

func (s *MetaCategoryTableEntity) SetDbId(v int32) *MetaCategoryTableEntity {
	s.DbId = &v
	return s
}

func (s *MetaCategoryTableEntity) SetDbType(v string) *MetaCategoryTableEntity {
	s.DbType = &v
	return s
}

func (s *MetaCategoryTableEntity) SetDescription(v string) *MetaCategoryTableEntity {
	s.Description = &v
	return s
}

func (s *MetaCategoryTableEntity) SetInstanceId(v int32) *MetaCategoryTableEntity {
	s.InstanceId = &v
	return s
}

func (s *MetaCategoryTableEntity) SetSchemaName(v string) *MetaCategoryTableEntity {
	s.SchemaName = &v
	return s
}

func (s *MetaCategoryTableEntity) SetTableName(v string) *MetaCategoryTableEntity {
	s.TableName = &v
	return s
}

func (s *MetaCategoryTableEntity) SetTableSchemaName(v string) *MetaCategoryTableEntity {
	s.TableSchemaName = &v
	return s
}

func (s *MetaCategoryTableEntity) Validate() error {
	return dara.Validate(s)
}
