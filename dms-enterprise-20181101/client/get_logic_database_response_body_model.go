// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogicDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetLogicDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetLogicDatabaseResponseBody
	GetErrorMessage() *string
	SetLogicDatabase(v *GetLogicDatabaseResponseBodyLogicDatabase) *GetLogicDatabaseResponseBody
	GetLogicDatabase() *GetLogicDatabaseResponseBodyLogicDatabase
	SetRequestId(v string) *GetLogicDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLogicDatabaseResponseBody
	GetSuccess() *bool
}

type GetLogicDatabaseResponseBody struct {
	// The status code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details of the logical database.
	LogicDatabase *GetLogicDatabaseResponseBodyLogicDatabase `json:"LogicDatabase,omitempty" xml:"LogicDatabase,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 90260530-565C-42B9-A6E8-893481FE6AB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLogicDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogicDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetLogicDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetLogicDatabaseResponseBody) GetLogicDatabase() *GetLogicDatabaseResponseBodyLogicDatabase {
	return s.LogicDatabase
}

func (s *GetLogicDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogicDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLogicDatabaseResponseBody) SetErrorCode(v string) *GetLogicDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLogicDatabaseResponseBody) SetErrorMessage(v string) *GetLogicDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLogicDatabaseResponseBody) SetLogicDatabase(v *GetLogicDatabaseResponseBodyLogicDatabase) *GetLogicDatabaseResponseBody {
	s.LogicDatabase = v
	return s
}

func (s *GetLogicDatabaseResponseBody) SetRequestId(v string) *GetLogicDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogicDatabaseResponseBody) SetSuccess(v bool) *GetLogicDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetLogicDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLogicDatabaseResponseBodyLogicDatabase struct {
	// The alias of the logical database.
	//
	// example:
	//
	// test_logic_alias
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The ID of the logical database.
	//
	// example:
	//
	// 1***
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The IDs of database shards of the logical database.
	DatabaseIds *GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Struct"`
	// The database engine. For more information about the valid values of the DbType parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// POLARDB
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database belongs. Valid values:
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
	// Indicates whether the database is a logical database. The return value is true.
	//
	// example:
	//
	// true
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The IDs of the owners of the logical database.
	OwnerIdList *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The names of the owners of the logical database.
	OwnerNameList *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The name of the logical database.
	//
	// example:
	//
	// test_logic_db
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used to search for the logical database.
	//
	// example:
	//
	// test_logic_db[test_logic_alias]
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetLogicDatabaseResponseBodyLogicDatabase) String() string {
	return dara.Prettify(s)
}

func (s GetLogicDatabaseResponseBodyLogicDatabase) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetAlias() *string {
	return s.Alias
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetDatabaseIds() *GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds {
	return s.DatabaseIds
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetDbType() *string {
	return s.DbType
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetEnvType() *string {
	return s.EnvType
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetLogic() *bool {
	return s.Logic
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetOwnerIdList() *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList {
	return s.OwnerIdList
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetOwnerNameList() *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList {
	return s.OwnerNameList
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetAlias(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.Alias = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetDatabaseId(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.DatabaseId = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetDatabaseIds(v *GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.DatabaseIds = v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetDbType(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.DbType = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetEnvType(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.EnvType = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetLogic(v bool) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.Logic = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetOwnerIdList(v *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.OwnerIdList = v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetOwnerNameList(v *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.OwnerNameList = v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetSchemaName(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.SchemaName = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetSearchName(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.SearchName = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) Validate() error {
	return dara.Validate(s)
}

type GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds struct {
	DatabaseIds []*int64 `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Repeated"`
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds) String() string {
	return dara.Prettify(s)
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds) GetDatabaseIds() []*int64 {
	return s.DatabaseIds
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds) SetDatabaseIds(v []*int64) *GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds {
	s.DatabaseIds = v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseDatabaseIds) Validate() error {
	return dara.Validate(s)
}

type GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) SetOwnerIds(v []*string) *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) SetOwnerNames(v []*string) *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) Validate() error {
	return dara.Validate(s)
}
