// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaDatabase interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *OneMetaDatabase
	GetCatalogName() *string
	SetCatalogType(v string) *OneMetaDatabase
	GetCatalogType() *string
	SetCatalogUuid(v string) *OneMetaDatabase
	GetCatalogUuid() *string
	SetDataSourceType(v string) *OneMetaDatabase
	GetDataSourceType() *string
	SetDatabaseBizAttrs(v map[string]interface{}) *OneMetaDatabase
	GetDatabaseBizAttrs() map[string]interface{}
	SetDatabaseUuid(v string) *OneMetaDatabase
	GetDatabaseUuid() *string
	SetDescription(v string) *OneMetaDatabase
	GetDescription() *string
	SetEngineMeta(v *OneMetaDatabaseEngineMeta) *OneMetaDatabase
	GetEngineMeta() *OneMetaDatabaseEngineMeta
	SetName(v string) *OneMetaDatabase
	GetName() *string
	SetProperties(v map[string]interface{}) *OneMetaDatabase
	GetProperties() map[string]interface{}
	SetQualifiedName(v string) *OneMetaDatabase
	GetQualifiedName() *string
	SetRegionId(v string) *OneMetaDatabase
	GetRegionId() *string
	SetSearchName(v string) *OneMetaDatabase
	GetSearchName() *string
	SetState(v int32) *OneMetaDatabase
	GetState() *int32
	SetStorageLocation(v string) *OneMetaDatabase
	GetStorageLocation() *string
}

type OneMetaDatabase struct {
	CatalogName      *string                    `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	CatalogType      *string                    `json:"CatalogType,omitempty" xml:"CatalogType,omitempty"`
	CatalogUuid      *string                    `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	DataSourceType   *string                    `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DatabaseBizAttrs map[string]interface{}     `json:"DatabaseBizAttrs,omitempty" xml:"DatabaseBizAttrs,omitempty"`
	DatabaseUuid     *string                    `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description      *string                    `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineMeta       *OneMetaDatabaseEngineMeta `json:"EngineMeta,omitempty" xml:"EngineMeta,omitempty"`
	Name             *string                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties       map[string]interface{}     `json:"Properties,omitempty" xml:"Properties,omitempty"`
	QualifiedName    *string                    `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	RegionId         *string                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchName       *string                    `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	State            *int32                     `json:"State,omitempty" xml:"State,omitempty"`
	StorageLocation  *string                    `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s OneMetaDatabase) String() string {
	return dara.Prettify(s)
}

func (s OneMetaDatabase) GoString() string {
	return s.String()
}

func (s *OneMetaDatabase) GetCatalogName() *string {
	return s.CatalogName
}

func (s *OneMetaDatabase) GetCatalogType() *string {
	return s.CatalogType
}

func (s *OneMetaDatabase) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *OneMetaDatabase) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *OneMetaDatabase) GetDatabaseBizAttrs() map[string]interface{} {
	return s.DatabaseBizAttrs
}

func (s *OneMetaDatabase) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *OneMetaDatabase) GetDescription() *string {
	return s.Description
}

func (s *OneMetaDatabase) GetEngineMeta() *OneMetaDatabaseEngineMeta {
	return s.EngineMeta
}

func (s *OneMetaDatabase) GetName() *string {
	return s.Name
}

func (s *OneMetaDatabase) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *OneMetaDatabase) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *OneMetaDatabase) GetRegionId() *string {
	return s.RegionId
}

func (s *OneMetaDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *OneMetaDatabase) GetState() *int32 {
	return s.State
}

func (s *OneMetaDatabase) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *OneMetaDatabase) SetCatalogName(v string) *OneMetaDatabase {
	s.CatalogName = &v
	return s
}

func (s *OneMetaDatabase) SetCatalogType(v string) *OneMetaDatabase {
	s.CatalogType = &v
	return s
}

func (s *OneMetaDatabase) SetCatalogUuid(v string) *OneMetaDatabase {
	s.CatalogUuid = &v
	return s
}

func (s *OneMetaDatabase) SetDataSourceType(v string) *OneMetaDatabase {
	s.DataSourceType = &v
	return s
}

func (s *OneMetaDatabase) SetDatabaseBizAttrs(v map[string]interface{}) *OneMetaDatabase {
	s.DatabaseBizAttrs = v
	return s
}

func (s *OneMetaDatabase) SetDatabaseUuid(v string) *OneMetaDatabase {
	s.DatabaseUuid = &v
	return s
}

func (s *OneMetaDatabase) SetDescription(v string) *OneMetaDatabase {
	s.Description = &v
	return s
}

func (s *OneMetaDatabase) SetEngineMeta(v *OneMetaDatabaseEngineMeta) *OneMetaDatabase {
	s.EngineMeta = v
	return s
}

func (s *OneMetaDatabase) SetName(v string) *OneMetaDatabase {
	s.Name = &v
	return s
}

func (s *OneMetaDatabase) SetProperties(v map[string]interface{}) *OneMetaDatabase {
	s.Properties = v
	return s
}

func (s *OneMetaDatabase) SetQualifiedName(v string) *OneMetaDatabase {
	s.QualifiedName = &v
	return s
}

func (s *OneMetaDatabase) SetRegionId(v string) *OneMetaDatabase {
	s.RegionId = &v
	return s
}

func (s *OneMetaDatabase) SetSearchName(v string) *OneMetaDatabase {
	s.SearchName = &v
	return s
}

func (s *OneMetaDatabase) SetState(v int32) *OneMetaDatabase {
	s.State = &v
	return s
}

func (s *OneMetaDatabase) SetStorageLocation(v string) *OneMetaDatabase {
	s.StorageLocation = &v
	return s
}

func (s *OneMetaDatabase) Validate() error {
	if s.EngineMeta != nil {
		if err := s.EngineMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}
