// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBTopology(v *GetDBTopologyResponseBodyDBTopology) *GetDBTopologyResponseBody
	GetDBTopology() *GetDBTopologyResponseBodyDBTopology
	SetErrorCode(v string) *GetDBTopologyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDBTopologyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDBTopologyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDBTopologyResponseBody
	GetSuccess() *bool
}

type GetDBTopologyResponseBody struct {
	// The topology of the data table.
	DBTopology *GetDBTopologyResponseBodyDBTopology `json:"DBTopology,omitempty" xml:"DBTopology,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C5B8E84B-42B6-4374-AD5A-6264E1753378
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDBTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDBTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBTopologyResponseBody) GetDBTopology() *GetDBTopologyResponseBodyDBTopology {
	return s.DBTopology
}

func (s *GetDBTopologyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDBTopologyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDBTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDBTopologyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDBTopologyResponseBody) SetDBTopology(v *GetDBTopologyResponseBodyDBTopology) *GetDBTopologyResponseBody {
	s.DBTopology = v
	return s
}

func (s *GetDBTopologyResponseBody) SetErrorCode(v string) *GetDBTopologyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDBTopologyResponseBody) SetErrorMessage(v string) *GetDBTopologyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDBTopologyResponseBody) SetRequestId(v string) *GetDBTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBTopologyResponseBody) SetSuccess(v bool) *GetDBTopologyResponseBody {
	s.Success = &v
	return s
}

func (s *GetDBTopologyResponseBody) Validate() error {
	if s.DBTopology != nil {
		if err := s.DBTopology.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDBTopologyResponseBodyDBTopology struct {
	// The alias of the access point.
	//
	// example:
	//
	// logic_db_test
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The list of database splitting topology information.
	DBTopologyInfoList []*GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList `json:"DBTopologyInfoList,omitempty" xml:"DBTopologyInfoList,omitempty" type:"Repeated"`
	// The type of the database engine.
	//
	// example:
	//
	// polardb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment in which the database instance is deployed. Valid values:
	//
	// 	- product: production environment
	//
	// 	- dev: development environment
	//
	// 	- pre: pre-release environment
	//
	// 	- test: test environment
	//
	// 	- sit: system integration testing (SIT) environment
	//
	// 	- uat: user acceptance testing (UAT) environment
	//
	// 	- pet: stress testing environment
	//
	// 	- stag: staging environment
	//
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the logical database.
	//
	// example:
	//
	// 1234
	LogicDbId *int64 `json:"LogicDbId,omitempty" xml:"LogicDbId,omitempty"`
	// Logical database name.
	//
	// example:
	//
	// logic_db_test
	LogicDbName *string `json:"LogicDbName,omitempty" xml:"LogicDbName,omitempty"`
	// The name of the saved search.
	//
	// example:
	//
	// logic_db_test
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetDBTopologyResponseBodyDBTopology) String() string {
	return dara.Prettify(s)
}

func (s GetDBTopologyResponseBodyDBTopology) GoString() string {
	return s.String()
}

func (s *GetDBTopologyResponseBodyDBTopology) GetAlias() *string {
	return s.Alias
}

func (s *GetDBTopologyResponseBodyDBTopology) GetDBTopologyInfoList() []*GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	return s.DBTopologyInfoList
}

func (s *GetDBTopologyResponseBodyDBTopology) GetDbType() *string {
	return s.DbType
}

func (s *GetDBTopologyResponseBodyDBTopology) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDBTopologyResponseBodyDBTopology) GetLogicDbId() *int64 {
	return s.LogicDbId
}

func (s *GetDBTopologyResponseBodyDBTopology) GetLogicDbName() *string {
	return s.LogicDbName
}

func (s *GetDBTopologyResponseBodyDBTopology) GetSearchName() *string {
	return s.SearchName
}

func (s *GetDBTopologyResponseBodyDBTopology) SetAlias(v string) *GetDBTopologyResponseBodyDBTopology {
	s.Alias = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetDBTopologyInfoList(v []*GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) *GetDBTopologyResponseBodyDBTopology {
	s.DBTopologyInfoList = v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetDbType(v string) *GetDBTopologyResponseBodyDBTopology {
	s.DbType = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetEnvType(v string) *GetDBTopologyResponseBodyDBTopology {
	s.EnvType = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetLogicDbId(v int64) *GetDBTopologyResponseBodyDBTopology {
	s.LogicDbId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetLogicDbName(v string) *GetDBTopologyResponseBodyDBTopology {
	s.LogicDbName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetSearchName(v string) *GetDBTopologyResponseBodyDBTopology {
	s.SearchName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) Validate() error {
	if s.DBTopologyInfoList != nil {
		for _, item := range s.DBTopologyInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList struct {
	// The name of the catalog to which the database belongs.
	//
	// > If the database is a PostgreSQL database, the value of this parameter is the name of the database.
	//
	// example:
	//
	// def
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The ID of the database for which the schema design is executed.
	//
	// example:
	//
	// 423532
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// polardb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database belongs. Valid values:
	//
	// 	- product: production environment
	//
	// 	- dev: development environment
	//
	// 	- pre: staging environment
	//
	// 	- test: test environment
	//
	// 	- sit: SIT environment
	//
	// 	- uat: user acceptance testing (UAT) environment
	//
	// 	- pet: stress testing environment
	//
	// 	- stag: STAG environment
	//
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the instance. The valid value is returned if you call the ListInstances operation. The instance ID is not the ID of the RDS instance.
	//
	// example:
	//
	// 4325325
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance resource ID.
	//
	// example:
	//
	// rm-xxx
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	// The source of the database instance. Valid values:
	//
	// 	- **PUBLIC_OWN:*	- a self-managed database instance that is deployed on the Internet
	//
	// 	- **RDS:*	- an ApsaraDB RDS instance
	//
	// 	- **ECS_OWN:*	- a self-managed database that is deployed on an Elastic Compute Service (ECS) instance
	//
	// 	- **VPC_IDC:*	- a self-managed database instance that is deployed in a data center connected over a virtual private cloud (VPC)
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the logical database.
	//
	// > If the database is a PostgreSQL database, the value of this parameter is the name of the database schema.
	//
	// example:
	//
	// db_test@rm-xxx:3306
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the saved search.
	//
	// example:
	//
	// db_test
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GoString() string {
	return s.String()
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetDbId() *int64 {
	return s.DbId
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetDbType() *string {
	return s.DbType
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetInstanceResourceId() *string {
	return s.InstanceResourceId
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GetSearchName() *string {
	return s.SearchName
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetCatalogName(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.CatalogName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetDbId(v int64) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.DbId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetDbType(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.DbType = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetEnvType(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.EnvType = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetInstanceId(v int64) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetInstanceResourceId(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.InstanceResourceId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetInstanceSource(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.InstanceSource = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetRegionId(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.RegionId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetSchemaName(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.SchemaName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetSearchName(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.SearchName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) Validate() error {
	return dara.Validate(s)
}
