// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticCatalog interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogBizAttrs(v map[string]interface{}) *AgenticCatalog
	GetCatalogBizAttrs() map[string]interface{}
	SetCatalogType(v string) *AgenticCatalog
	GetCatalogType() *string
	SetCatalogUuid(v string) *AgenticCatalog
	GetCatalogUuid() *string
	SetDataSourceType(v string) *AgenticCatalog
	GetDataSourceType() *string
	SetDataSourceUuid(v string) *AgenticCatalog
	GetDataSourceUuid() *string
	SetDescription(v string) *AgenticCatalog
	GetDescription() *string
	SetName(v string) *AgenticCatalog
	GetName() *string
	SetProperties(v map[string]interface{}) *AgenticCatalog
	GetProperties() map[string]interface{}
	SetRegionId(v string) *AgenticCatalog
	GetRegionId() *string
	SetState(v int32) *AgenticCatalog
	GetState() *int32
	SetStorageLocation(v string) *AgenticCatalog
	GetStorageLocation() *string
}

type AgenticCatalog struct {
	// A collection of key-value pairs that represents business attributes for the catalog, such as the data owner or department.
	CatalogBizAttrs map[string]interface{} `json:"CatalogBizAttrs,omitempty" xml:"CatalogBizAttrs,omitempty"`
	// The type of the catalog. For example, `INTERNAL_METADATA` or `THIRD_PARTY`.
	CatalogType *string `json:"CatalogType,omitempty" xml:"CatalogType,omitempty"`
	// The unique identifier (UUID) of the catalog. This parameter is system-generated and output-only.
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The type of the data source associated with the catalog. For example, `MySQL`, `PostgreSQL`, or `OSS`.
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The unique identifier (UUID) of the associated data source.
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
	// The description of the catalog. It can be up to 2,048 characters long.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the catalog. The name can be up to 256 characters long.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A collection of key-value pairs that represents additional technical properties for the catalog.
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The ID of the region where the catalog is located. For example, `cn-hangzhou`.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The current state of the catalog. Valid values are: `0` (Creating), `1` (Active), `2` (Deleting), and `3` (Error).
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// The storage location for the catalog\\"s metadata, such as a database name or a file path.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s AgenticCatalog) String() string {
	return dara.Prettify(s)
}

func (s AgenticCatalog) GoString() string {
	return s.String()
}

func (s *AgenticCatalog) GetCatalogBizAttrs() map[string]interface{} {
	return s.CatalogBizAttrs
}

func (s *AgenticCatalog) GetCatalogType() *string {
	return s.CatalogType
}

func (s *AgenticCatalog) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *AgenticCatalog) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *AgenticCatalog) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *AgenticCatalog) GetDescription() *string {
	return s.Description
}

func (s *AgenticCatalog) GetName() *string {
	return s.Name
}

func (s *AgenticCatalog) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *AgenticCatalog) GetRegionId() *string {
	return s.RegionId
}

func (s *AgenticCatalog) GetState() *int32 {
	return s.State
}

func (s *AgenticCatalog) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *AgenticCatalog) SetCatalogBizAttrs(v map[string]interface{}) *AgenticCatalog {
	s.CatalogBizAttrs = v
	return s
}

func (s *AgenticCatalog) SetCatalogType(v string) *AgenticCatalog {
	s.CatalogType = &v
	return s
}

func (s *AgenticCatalog) SetCatalogUuid(v string) *AgenticCatalog {
	s.CatalogUuid = &v
	return s
}

func (s *AgenticCatalog) SetDataSourceType(v string) *AgenticCatalog {
	s.DataSourceType = &v
	return s
}

func (s *AgenticCatalog) SetDataSourceUuid(v string) *AgenticCatalog {
	s.DataSourceUuid = &v
	return s
}

func (s *AgenticCatalog) SetDescription(v string) *AgenticCatalog {
	s.Description = &v
	return s
}

func (s *AgenticCatalog) SetName(v string) *AgenticCatalog {
	s.Name = &v
	return s
}

func (s *AgenticCatalog) SetProperties(v map[string]interface{}) *AgenticCatalog {
	s.Properties = v
	return s
}

func (s *AgenticCatalog) SetRegionId(v string) *AgenticCatalog {
	s.RegionId = &v
	return s
}

func (s *AgenticCatalog) SetState(v int32) *AgenticCatalog {
	s.State = &v
	return s
}

func (s *AgenticCatalog) SetStorageLocation(v string) *AgenticCatalog {
	s.StorageLocation = &v
	return s
}

func (s *AgenticCatalog) Validate() error {
	return dara.Validate(s)
}
