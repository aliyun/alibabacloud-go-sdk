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
	SetDataSourceUuid(v string) *AgenticCatalog
	GetDataSourceUuid() *string
	SetDbType(v string) *AgenticCatalog
	GetDbType() *string
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
	// The business attributes of the catalog, returned as key-value pairs. The following keys may be included:
	//
	// - **ALIAS**: the alias of the data source
	//
	// - **HOST**: the endpoint of the data source
	//
	// - **PORT**: the connection port of the data source
	//
	// - **ENV_TYPE**: the environment type of the data source
	//
	// - **VERSION**: the database version of the data source
	//
	// - **SYNC_ALIAS**: whether automatic synchronization of instance aliases is enabled
	CatalogBizAttrs map[string]interface{} `json:"CatalogBizAttrs,omitempty" xml:"CatalogBizAttrs,omitempty"`
	// The type of the catalog. Valid values:
	//
	// - **JDBC**: relational database catalog
	//
	// - **DataLake**: data lake catalog
	//
	// - **Dataset**: knowledge base catalog
	//
	// example:
	//
	// JDBC
	CatalogType *string `json:"CatalogType,omitempty" xml:"CatalogType,omitempty"`
	// The unique identifier of the catalog.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The unique identifier of the data source associated with the catalog.
	//
	// example:
	//
	// ds-HZ-Xk2Pq9Lm***
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
	// The database type of the data source associated with the catalog. The value is the uppercase name of the database type, such as MYSQL.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description of the catalog.
	//
	// example:
	//
	// Order production database
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the catalog.
	//
	// example:
	//
	// prod_mysql_order
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended properties of the catalog, returned as key-value pairs.
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The region ID of the data source associated with the catalog.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the catalog. Valid values:
	//
	// - **0**: Normal. This value is set when the data source is created.
	//
	// - **-1**: Deleted. This value is set when the data source is deleted.
	//
	// Disabling a data source does not change the status of the catalog.
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// The storage location. This property is returned only for catalogs of the DataLake type.
	//
	// example:
	//
	// oss://examplebucket/datalake/
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

func (s *AgenticCatalog) GetDataSourceUuid() *string {
	return s.DataSourceUuid
}

func (s *AgenticCatalog) GetDbType() *string {
	return s.DbType
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

func (s *AgenticCatalog) SetDataSourceUuid(v string) *AgenticCatalog {
	s.DataSourceUuid = &v
	return s
}

func (s *AgenticCatalog) SetDbType(v string) *AgenticCatalog {
	s.DbType = &v
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
