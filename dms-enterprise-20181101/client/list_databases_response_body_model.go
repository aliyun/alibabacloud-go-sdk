// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseList(v *ListDatabasesResponseBodyDatabaseList) *ListDatabasesResponseBody
	GetDatabaseList() *ListDatabasesResponseBodyDatabaseList
	SetErrorCode(v string) *ListDatabasesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDatabasesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDatabasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatabasesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListDatabasesResponseBody
	GetTotalCount() *int64
}

type ListDatabasesResponseBody struct {
	// The details of the databases.
	DatabaseList *ListDatabasesResponseBodyDatabaseList `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Struct"`
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
	// 5541CDA6-F674-435C-81BD-40C2FB926CE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of databases that belong to an instance.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) GetDatabaseList() *ListDatabasesResponseBodyDatabaseList {
	return s.DatabaseList
}

func (s *ListDatabasesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDatabasesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatabasesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabasesResponseBody) SetDatabaseList(v *ListDatabasesResponseBodyDatabaseList) *ListDatabasesResponseBody {
	s.DatabaseList = v
	return s
}

func (s *ListDatabasesResponseBody) SetErrorCode(v string) *ListDatabasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDatabasesResponseBody) SetErrorMessage(v string) *ListDatabasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDatabasesResponseBody) SetRequestId(v string) *ListDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesResponseBody) SetSuccess(v bool) *ListDatabasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatabasesResponseBody) SetTotalCount(v int64) *ListDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyDatabaseList struct {
	Database []*ListDatabasesResponseBodyDatabaseListDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyDatabaseList) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseList) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseList) GetDatabase() []*ListDatabasesResponseBodyDatabaseListDatabase {
	return s.Database
}

func (s *ListDatabasesResponseBodyDatabaseList) SetDatabase(v []*ListDatabasesResponseBodyDatabaseListDatabase) *ListDatabasesResponseBodyDatabaseList {
	s.Database = v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseList) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyDatabaseListDatabase struct {
	// The name of the catalog to which the database belongs.
	//
	// example:
	//
	// 1
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 1
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The ID of the DBA.
	//
	// example:
	//
	// 1
	DbaId *string `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// The nickname of the Database administrator (DBA) to which the database belongs.
	//
	// example:
	//
	// dba_user
	DbaName *string `json:"DbaName,omitempty" xml:"DbaName,omitempty"`
	// The encoding format of the database.
	//
	// example:
	//
	// utf-8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The type of the environment to which the database belongs.
	//
	// example:
	//
	// dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The endpoint of the instance to which the database belongs.
	//
	// example:
	//
	// xxx.xxx.xxx.xxx
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the instance to which the database belongs.
	//
	// example:
	//
	// 1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the owners of the database.
	OwnerIdList *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The nicknames of the database owners.
	OwnerNameList *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The connection port of the instance to which the database belongs.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used for searching the database.
	//
	// example:
	//
	// test@xxx.xxx.xxx.xxx:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The system ID (SID) of the instance to which the database belongs.
	//
	// example:
	//
	// test
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The state of the database. Valid values:
	//
	// 	- NORMAL: The database is normal.
	//
	// 	- DISABLE: The database is disabled.
	//
	// 	- OFFLINE: The database is unpublished.
	//
	// 	- NOT_EXIST: The database does not exist.
	//
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListDatabasesResponseBodyDatabaseListDatabase) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseListDatabase) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetDbType() *string {
	return s.DbType
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetDbaId() *string {
	return s.DbaId
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetDbaName() *string {
	return s.DbaName
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetEncoding() *string {
	return s.Encoding
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetHost() *string {
	return s.Host
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetOwnerIdList() *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList {
	return s.OwnerIdList
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetOwnerNameList() *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList {
	return s.OwnerNameList
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetPort() *int32 {
	return s.Port
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetSid() *string {
	return s.Sid
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) GetState() *string {
	return s.State
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetCatalogName(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.CatalogName = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetDatabaseId(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetDbType(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.DbType = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetDbaId(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.DbaId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetDbaName(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.DbaName = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetEncoding(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.Encoding = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetEnvType(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.EnvType = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetHost(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.Host = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetInstanceId(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetOwnerIdList(v *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.OwnerIdList = v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetOwnerNameList(v *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.OwnerNameList = v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetPort(v int32) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.Port = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetSchemaName(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.SchemaName = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetSearchName(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.SearchName = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetSid(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.Sid = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetState(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.State = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) SetOwnerIds(v []*string) *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) SetOwnerNames(v []*string) *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) Validate() error {
	return dara.Validate(s)
}
