// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SearchTableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchTableResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SearchTableResponseBody
	GetRequestId() *string
	SetSearchTableList(v *SearchTableResponseBodySearchTableList) *SearchTableResponseBody
	GetSearchTableList() *SearchTableResponseBodySearchTableList
	SetSuccess(v bool) *SearchTableResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *SearchTableResponseBody
	GetTotalCount() *int64
}

type SearchTableResponseBody struct {
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
	// 1489257F-1B5D-4B5B-89EF-923C12CEEBD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the tables.
	SearchTableList *SearchTableResponseBodySearchTableList `json:"SearchTableList,omitempty" xml:"SearchTableList,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTableResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchTableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTableResponseBody) GetSearchTableList() *SearchTableResponseBodySearchTableList {
	return s.SearchTableList
}

func (s *SearchTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchTableResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchTableResponseBody) SetErrorCode(v string) *SearchTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchTableResponseBody) SetErrorMessage(v string) *SearchTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchTableResponseBody) SetRequestId(v string) *SearchTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTableResponseBody) SetSearchTableList(v *SearchTableResponseBodySearchTableList) *SearchTableResponseBody {
	s.SearchTableList = v
	return s
}

func (s *SearchTableResponseBody) SetSuccess(v bool) *SearchTableResponseBody {
	s.Success = &v
	return s
}

func (s *SearchTableResponseBody) SetTotalCount(v int64) *SearchTableResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchTableResponseBodySearchTableList struct {
	SearchTable []*SearchTableResponseBodySearchTableListSearchTable `json:"SearchTable,omitempty" xml:"SearchTable,omitempty" type:"Repeated"`
}

func (s SearchTableResponseBodySearchTableList) String() string {
	return dara.Prettify(s)
}

func (s SearchTableResponseBodySearchTableList) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBodySearchTableList) GetSearchTable() []*SearchTableResponseBodySearchTableListSearchTable {
	return s.SearchTable
}

func (s *SearchTableResponseBodySearchTableList) SetSearchTable(v []*SearchTableResponseBodySearchTableListSearchTable) *SearchTableResponseBodySearchTableList {
	s.SearchTable = v
	return s
}

func (s *SearchTableResponseBodySearchTableList) Validate() error {
	return dara.Validate(s)
}

type SearchTableResponseBodySearchTableListSearchTable struct {
	// The name that is used to search for the database to which the table belongs.
	//
	// example:
	//
	// test
	DBSearchName *string `json:"DBSearchName,omitempty" xml:"DBSearchName,omitempty"`
	// The ID of the database to which the table belongs.
	//
	// example:
	//
	// 1
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// 	- **DRDS**
	//
	// 	- **OceanBase**
	//
	// 	- **Mongo**
	//
	// 	- **Redis**
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description of the table.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encoding format of the table.
	//
	// example:
	//
	// utf8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The engine of the table.
	//
	// example:
	//
	// innodb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The type of the environment to which the database belongs.
	//
	// example:
	//
	// test
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the table is a logical table. Valid values:
	//
	// 	- **true**: The table is a logical table.
	//
	// 	- **false**: The table is not a logical table.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The IDs of the table owners.
	OwnerIdList *SearchTableResponseBodySearchTableListSearchTableOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The nicknames of the table owners.
	OwnerNameList *SearchTableResponseBodySearchTableListSearchTableOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The GUID of the table.
	//
	// example:
	//
	// IDB_L_9032.db-test.yuyang_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The ID of the table.
	//
	// example:
	//
	// 1
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The name of the database to which the table belongs.
	//
	// example:
	//
	// test@xxx.xxx.xxx.xxx:3306
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
}

func (s SearchTableResponseBodySearchTableListSearchTable) String() string {
	return dara.Prettify(s)
}

func (s SearchTableResponseBodySearchTableListSearchTable) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetDBSearchName() *string {
	return s.DBSearchName
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetDbName() *string {
	return s.DbName
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetDbType() *string {
	return s.DbType
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetDescription() *string {
	return s.Description
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetEncoding() *string {
	return s.Encoding
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetEngine() *string {
	return s.Engine
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetEnvType() *string {
	return s.EnvType
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetLogic() *bool {
	return s.Logic
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetOwnerIdList() *SearchTableResponseBodySearchTableListSearchTableOwnerIdList {
	return s.OwnerIdList
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetOwnerNameList() *SearchTableResponseBodySearchTableListSearchTableOwnerNameList {
	return s.OwnerNameList
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetTableGuid() *string {
	return s.TableGuid
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetTableId() *string {
	return s.TableId
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetTableName() *string {
	return s.TableName
}

func (s *SearchTableResponseBodySearchTableListSearchTable) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDBSearchName(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.DBSearchName = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDatabaseId(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.DatabaseId = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDbName(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.DbName = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDbType(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.DbType = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDescription(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.Description = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetEncoding(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.Encoding = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetEngine(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.Engine = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetEnvType(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.EnvType = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetLogic(v bool) *SearchTableResponseBodySearchTableListSearchTable {
	s.Logic = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetOwnerIdList(v *SearchTableResponseBodySearchTableListSearchTableOwnerIdList) *SearchTableResponseBodySearchTableListSearchTable {
	s.OwnerIdList = v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetOwnerNameList(v *SearchTableResponseBodySearchTableListSearchTableOwnerNameList) *SearchTableResponseBodySearchTableListSearchTable {
	s.OwnerNameList = v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetTableGuid(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.TableGuid = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetTableId(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.TableId = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetTableName(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.TableName = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetTableSchemaName(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.TableSchemaName = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) Validate() error {
	return dara.Validate(s)
}

type SearchTableResponseBodySearchTableListSearchTableOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s SearchTableResponseBodySearchTableListSearchTableOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s SearchTableResponseBodySearchTableListSearchTableOwnerIdList) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBodySearchTableListSearchTableOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *SearchTableResponseBodySearchTableListSearchTableOwnerIdList) SetOwnerIds(v []*string) *SearchTableResponseBodySearchTableListSearchTableOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTableOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type SearchTableResponseBodySearchTableListSearchTableOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s SearchTableResponseBodySearchTableListSearchTableOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s SearchTableResponseBodySearchTableListSearchTableOwnerNameList) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBodySearchTableListSearchTableOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *SearchTableResponseBodySearchTableListSearchTableOwnerNameList) SetOwnerNames(v []*string) *SearchTableResponseBodySearchTableListSearchTableOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTableOwnerNameList) Validate() error {
	return dara.Validate(s)
}
