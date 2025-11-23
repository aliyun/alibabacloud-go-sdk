// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v *GetDatabaseResponseBodyDatabase) *GetDatabaseResponseBody
	GetDatabase() *GetDatabaseResponseBodyDatabase
	SetErrorCode(v string) *GetDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatabaseResponseBody
	GetSuccess() *bool
}

type GetDatabaseResponseBody struct {
	// The details of the database.
	Database *GetDatabaseResponseBodyDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// An unknown error occurred.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3CDB8601-AD74-4A47-8114-08E08CD6****
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

func (s GetDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBody) GetDatabase() *GetDatabaseResponseBodyDatabase {
	return s.Database
}

func (s *GetDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatabaseResponseBody) SetDatabase(v *GetDatabaseResponseBodyDatabase) *GetDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *GetDatabaseResponseBody) SetErrorCode(v string) *GetDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDatabaseResponseBody) SetErrorMessage(v string) *GetDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDatabaseResponseBody) SetRequestId(v string) *GetDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseResponseBody) SetSuccess(v bool) *GetDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatabaseResponseBody) Validate() error {
	if s.Database != nil {
		if err := s.Database.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatabaseResponseBodyDatabase struct {
	// The name of the catalog to which the database belongs.
	//
	// example:
	//
	// def
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 984****
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The ID of the database administrator (DBA).
	//
	// example:
	//
	// 27****
	DbaId *string `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// The nickname of the DBA.
	//
	// example:
	//
	// dba_name
	DbaName *string `json:"DbaName,omitempty" xml:"DbaName,omitempty"`
	// The encoding format of the database.
	//
	// example:
	//
	// utf8mb4
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The type of the environment to which the database belongs. Valid values:
	//
	// 	- **product**: production environment
	//
	// 	- **dev**: development environment
	//
	// 	- **pre**: staging environment
	//
	// 	- **test**: test environment
	//
	// 	- **sit**: SIT environment
	//
	// 	- **uat**: user acceptance testing (UAT) environment
	//
	// 	- **pet**: stress testing environment
	//
	// 	- **stag**: STAG environment
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The endpoint that is used to connect to the database.
	//
	// example:
	//
	// 192.168.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The alias of the instance.
	//
	// example:
	//
	// test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 149****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the owners of the database.
	OwnerIdList *GetDatabaseResponseBodyDatabaseOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The names of the owners of the database.
	OwnerNameList *GetDatabaseResponseBodyDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The port that is used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// mysql
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The keyword that is used to search for the database.
	//
	// example:
	//
	// mysql@192.168.XX.XX:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The SID of the database.
	//
	// >  The value of the parameter is returned only for Oracle databases.
	//
	// example:
	//
	// test_sid
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The status of the database. Valid values:
	//
	// 	- **NORMAL**: The database is running as expected.
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

func (s GetDatabaseResponseBodyDatabase) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseResponseBodyDatabase) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBodyDatabase) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDatabaseResponseBodyDatabase) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *GetDatabaseResponseBodyDatabase) GetDbType() *string {
	return s.DbType
}

func (s *GetDatabaseResponseBodyDatabase) GetDbaId() *string {
	return s.DbaId
}

func (s *GetDatabaseResponseBodyDatabase) GetDbaName() *string {
	return s.DbaName
}

func (s *GetDatabaseResponseBodyDatabase) GetEncoding() *string {
	return s.Encoding
}

func (s *GetDatabaseResponseBodyDatabase) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDatabaseResponseBodyDatabase) GetHost() *string {
	return s.Host
}

func (s *GetDatabaseResponseBodyDatabase) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetDatabaseResponseBodyDatabase) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDatabaseResponseBodyDatabase) GetOwnerIdList() *GetDatabaseResponseBodyDatabaseOwnerIdList {
	return s.OwnerIdList
}

func (s *GetDatabaseResponseBodyDatabase) GetOwnerNameList() *GetDatabaseResponseBodyDatabaseOwnerNameList {
	return s.OwnerNameList
}

func (s *GetDatabaseResponseBodyDatabase) GetPort() *int32 {
	return s.Port
}

func (s *GetDatabaseResponseBodyDatabase) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetDatabaseResponseBodyDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *GetDatabaseResponseBodyDatabase) GetSid() *string {
	return s.Sid
}

func (s *GetDatabaseResponseBodyDatabase) GetState() *string {
	return s.State
}

func (s *GetDatabaseResponseBodyDatabase) SetCatalogName(v string) *GetDatabaseResponseBodyDatabase {
	s.CatalogName = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDatabaseId(v string) *GetDatabaseResponseBodyDatabase {
	s.DatabaseId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDbType(v string) *GetDatabaseResponseBodyDatabase {
	s.DbType = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDbaId(v string) *GetDatabaseResponseBodyDatabase {
	s.DbaId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDbaName(v string) *GetDatabaseResponseBodyDatabase {
	s.DbaName = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetEncoding(v string) *GetDatabaseResponseBodyDatabase {
	s.Encoding = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetEnvType(v string) *GetDatabaseResponseBodyDatabase {
	s.EnvType = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetHost(v string) *GetDatabaseResponseBodyDatabase {
	s.Host = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetInstanceAlias(v string) *GetDatabaseResponseBodyDatabase {
	s.InstanceAlias = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetInstanceId(v string) *GetDatabaseResponseBodyDatabase {
	s.InstanceId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetOwnerIdList(v *GetDatabaseResponseBodyDatabaseOwnerIdList) *GetDatabaseResponseBodyDatabase {
	s.OwnerIdList = v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetOwnerNameList(v *GetDatabaseResponseBodyDatabaseOwnerNameList) *GetDatabaseResponseBodyDatabase {
	s.OwnerNameList = v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetPort(v int32) *GetDatabaseResponseBodyDatabase {
	s.Port = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSchemaName(v string) *GetDatabaseResponseBodyDatabase {
	s.SchemaName = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSearchName(v string) *GetDatabaseResponseBodyDatabase {
	s.SearchName = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSid(v string) *GetDatabaseResponseBodyDatabase {
	s.Sid = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetState(v string) *GetDatabaseResponseBodyDatabase {
	s.State = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) Validate() error {
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

type GetDatabaseResponseBodyDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s GetDatabaseResponseBodyDatabaseOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseResponseBodyDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBodyDatabaseOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *GetDatabaseResponseBodyDatabaseOwnerIdList) SetOwnerIds(v []*string) *GetDatabaseResponseBodyDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *GetDatabaseResponseBodyDatabaseOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type GetDatabaseResponseBodyDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s GetDatabaseResponseBodyDatabaseOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseResponseBodyDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBodyDatabaseOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *GetDatabaseResponseBodyDatabaseOwnerNameList) SetOwnerNames(v []*string) *GetDatabaseResponseBodyDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *GetDatabaseResponseBodyDatabaseOwnerNameList) Validate() error {
	return dara.Validate(s)
}
