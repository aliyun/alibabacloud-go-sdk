// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v *GetPhysicalDatabaseResponseBodyDatabase) *GetPhysicalDatabaseResponseBody
	GetDatabase() *GetPhysicalDatabaseResponseBodyDatabase
	SetErrorCode(v string) *GetPhysicalDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPhysicalDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetPhysicalDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPhysicalDatabaseResponseBody
	GetSuccess() *bool
}

type GetPhysicalDatabaseResponseBody struct {
	// The information about the physical database.
	Database *GetPhysicalDatabaseResponseBodyDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	// The error code returned if the request failed.
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
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponseBody) GetDatabase() *GetPhysicalDatabaseResponseBodyDatabase {
	return s.Database
}

func (s *GetPhysicalDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPhysicalDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPhysicalDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhysicalDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPhysicalDatabaseResponseBody) SetDatabase(v *GetPhysicalDatabaseResponseBodyDatabase) *GetPhysicalDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) SetErrorCode(v string) *GetPhysicalDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) SetErrorMessage(v string) *GetPhysicalDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) SetRequestId(v string) *GetPhysicalDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) SetSuccess(v bool) *GetPhysicalDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) Validate() error {
	if s.Database != nil {
		if err := s.Database.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhysicalDatabaseResponseBodyDatabase struct {
	// The name of the catalog to which the database belongs.
	//
	// > : If the database is a PostgreSQL database, the name of the database is displayed.
	//
	// example:
	//
	// def
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The ID of the physical database.
	//
	// example:
	//
	// 43125312
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The user ID of the DBA in the destination database.
	//
	// example:
	//
	// 43253
	DbaId *string `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// The nickname of the database administrator (DBA) in the destination database.
	//
	// example:
	//
	// dmstest
	DbaName *string `json:"DbaName,omitempty" xml:"DbaName,omitempty"`
	// The encoding format of the database.
	//
	// example:
	//
	// utf8mb4
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The type of the environment to which the database belongs. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The endpoint that is used to connect to the database.
	//
	// example:
	//
	// rm-xxxab3r272.mysql.rds.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The alias of the database instance.
	//
	// example:
	//
	// test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The instance ID of the destination database.
	//
	// example:
	//
	// 43215325
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The user IDs of the database owners.
	OwnerIdList *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The nicknames of the database owners.
	OwnerNameList *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The port that is used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the database.
	//
	// > : If the database is a PostgreSQL database, the name of the mode is displayed.
	//
	// example:
	//
	// dmstest
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used for searching the database.
	//
	// example:
	//
	// dmstest@rm-xxxab3r272.mysql.rds.aliyuncs.com:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The system ID (SID) of the database.
	//
	// > : The value of the parameter is returned only for Oracle databases.
	//
	// example:
	//
	// def
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The state of the database. Valid values:
	//
	// 	- **NORMAL**: The database is normal.
	//
	// 	- **DISABLE**: The database is disabled.
	//
	// 	- **OFFLINE**: The database is unpublished.
	//
	// 	- **NOT_EXIST**: The database does not exist.
	//
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetPhysicalDatabaseResponseBodyDatabase) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalDatabaseResponseBodyDatabase) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetDbType() *string {
	return s.DbType
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetDbaId() *string {
	return s.DbaId
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetDbaName() *string {
	return s.DbaName
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetEncoding() *string {
	return s.Encoding
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetEnvType() *string {
	return s.EnvType
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetHost() *string {
	return s.Host
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetOwnerIdList() *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList {
	return s.OwnerIdList
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetOwnerNameList() *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList {
	return s.OwnerNameList
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetPort() *int32 {
	return s.Port
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetSid() *string {
	return s.Sid
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) GetState() *string {
	return s.State
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetCatalogName(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.CatalogName = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetDatabaseId(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.DatabaseId = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetDbType(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.DbType = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetDbaId(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.DbaId = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetDbaName(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.DbaName = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetEncoding(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.Encoding = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetEnvType(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.EnvType = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetHost(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.Host = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetInstanceAlias(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.InstanceAlias = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetInstanceId(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.InstanceId = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetOwnerIdList(v *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) *GetPhysicalDatabaseResponseBodyDatabase {
	s.OwnerIdList = v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetOwnerNameList(v *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) *GetPhysicalDatabaseResponseBodyDatabase {
	s.OwnerNameList = v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetPort(v int32) *GetPhysicalDatabaseResponseBodyDatabase {
	s.Port = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetSchemaName(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.SchemaName = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetSearchName(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.SearchName = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetSid(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.Sid = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetState(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.State = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) Validate() error {
	if s.OwnerIdList != nil {
		if err := s.OwnerIdList.Validate(); err != nil {
			return err
		}
	}
	if s.OwnerNameList != nil {
		if err := s.OwnerNameList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) SetOwnerIds(v []*string) *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) SetOwnerNames(v []*string) *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) Validate() error {
	return dara.Validate(s)
}
