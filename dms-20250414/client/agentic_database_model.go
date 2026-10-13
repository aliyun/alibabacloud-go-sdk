// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticDatabase interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *AgenticDatabase
	GetCatalogName() *string
	SetCatalogType(v string) *AgenticDatabase
	GetCatalogType() *string
	SetCatalogUuid(v string) *AgenticDatabase
	GetCatalogUuid() *string
	SetDatabaseBizAttrs(v map[string]interface{}) *AgenticDatabase
	GetDatabaseBizAttrs() map[string]interface{}
	SetDatabaseUuid(v string) *AgenticDatabase
	GetDatabaseUuid() *string
	SetDbType(v string) *AgenticDatabase
	GetDbType() *string
	SetDescription(v string) *AgenticDatabase
	GetDescription() *string
	SetEngineMeta(v *AgenticDatabaseEngineMeta) *AgenticDatabase
	GetEngineMeta() *AgenticDatabaseEngineMeta
	SetName(v string) *AgenticDatabase
	GetName() *string
	SetProperties(v map[string]interface{}) *AgenticDatabase
	GetProperties() map[string]interface{}
	SetQualifiedName(v string) *AgenticDatabase
	GetQualifiedName() *string
	SetRegionId(v string) *AgenticDatabase
	GetRegionId() *string
	SetSearchName(v string) *AgenticDatabase
	GetSearchName() *string
	SetState(v int32) *AgenticDatabase
	GetState() *int32
	SetStorageLocation(v string) *AgenticDatabase
	GetStorageLocation() *string
}

type AgenticDatabase struct {
	// The name of the catalog.
	//
	// example:
	//
	// prod_mysql_order
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
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
	// The UUID of the catalog.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The business attributes of the database, returned as key-value pairs. Currently, this may include **ENV_TYPE**, which indicates the environment type of the database.
	DatabaseBizAttrs map[string]interface{} `json:"DatabaseBizAttrs,omitempty" xml:"DatabaseBizAttrs,omitempty"`
	// The unique identifier of the database.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The database type inherited from the catalog. The value is the uppercase name of the database type, such as MYSQL.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description of the database.
	//
	// example:
	//
	// Order database
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata of the database engine. This field is returned only for databases under a JDBC catalog.
	EngineMeta *AgenticDatabaseEngineMeta `json:"EngineMeta,omitempty" xml:"EngineMeta,omitempty"`
	// The name of the database. For databases that distinguish schemas, such as PostgreSQL, this is the schema name.
	//
	// example:
	//
	// order_db
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended properties of the database, returned as key-value pairs. This field is returned only for databases under a DataLake catalog.
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The fully qualified name of the database. For databases that distinguish schemas, such as PostgreSQL, the format is database name.schema name.
	//
	// example:
	//
	// order_db
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The region ID of the database.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The search name of the database. SearchAgenticDatabase matches search keywords against this name.
	//
	// example:
	//
	// order_db
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The status of the database. Valid values:
	//
	// - **0**: normal
	//
	// - **-1**: deleted
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// The storage location. This field is returned only for databases under a DataLake catalog.
	//
	// example:
	//
	// oss://examplebucket/datalake/order_db/
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s AgenticDatabase) String() string {
	return dara.Prettify(s)
}

func (s AgenticDatabase) GoString() string {
	return s.String()
}

func (s *AgenticDatabase) GetCatalogName() *string {
	return s.CatalogName
}

func (s *AgenticDatabase) GetCatalogType() *string {
	return s.CatalogType
}

func (s *AgenticDatabase) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *AgenticDatabase) GetDatabaseBizAttrs() map[string]interface{} {
	return s.DatabaseBizAttrs
}

func (s *AgenticDatabase) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *AgenticDatabase) GetDbType() *string {
	return s.DbType
}

func (s *AgenticDatabase) GetDescription() *string {
	return s.Description
}

func (s *AgenticDatabase) GetEngineMeta() *AgenticDatabaseEngineMeta {
	return s.EngineMeta
}

func (s *AgenticDatabase) GetName() *string {
	return s.Name
}

func (s *AgenticDatabase) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *AgenticDatabase) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *AgenticDatabase) GetRegionId() *string {
	return s.RegionId
}

func (s *AgenticDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *AgenticDatabase) GetState() *int32 {
	return s.State
}

func (s *AgenticDatabase) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *AgenticDatabase) SetCatalogName(v string) *AgenticDatabase {
	s.CatalogName = &v
	return s
}

func (s *AgenticDatabase) SetCatalogType(v string) *AgenticDatabase {
	s.CatalogType = &v
	return s
}

func (s *AgenticDatabase) SetCatalogUuid(v string) *AgenticDatabase {
	s.CatalogUuid = &v
	return s
}

func (s *AgenticDatabase) SetDatabaseBizAttrs(v map[string]interface{}) *AgenticDatabase {
	s.DatabaseBizAttrs = v
	return s
}

func (s *AgenticDatabase) SetDatabaseUuid(v string) *AgenticDatabase {
	s.DatabaseUuid = &v
	return s
}

func (s *AgenticDatabase) SetDbType(v string) *AgenticDatabase {
	s.DbType = &v
	return s
}

func (s *AgenticDatabase) SetDescription(v string) *AgenticDatabase {
	s.Description = &v
	return s
}

func (s *AgenticDatabase) SetEngineMeta(v *AgenticDatabaseEngineMeta) *AgenticDatabase {
	s.EngineMeta = v
	return s
}

func (s *AgenticDatabase) SetName(v string) *AgenticDatabase {
	s.Name = &v
	return s
}

func (s *AgenticDatabase) SetProperties(v map[string]interface{}) *AgenticDatabase {
	s.Properties = v
	return s
}

func (s *AgenticDatabase) SetQualifiedName(v string) *AgenticDatabase {
	s.QualifiedName = &v
	return s
}

func (s *AgenticDatabase) SetRegionId(v string) *AgenticDatabase {
	s.RegionId = &v
	return s
}

func (s *AgenticDatabase) SetSearchName(v string) *AgenticDatabase {
	s.SearchName = &v
	return s
}

func (s *AgenticDatabase) SetState(v int32) *AgenticDatabase {
	s.State = &v
	return s
}

func (s *AgenticDatabase) SetStorageLocation(v string) *AgenticDatabase {
	s.StorageLocation = &v
	return s
}

func (s *AgenticDatabase) Validate() error {
	if s.EngineMeta != nil {
		if err := s.EngineMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}
