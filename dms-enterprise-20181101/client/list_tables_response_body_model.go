// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListTablesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTablesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTablesResponseBody
	GetSuccess() *bool
	SetTableList(v *ListTablesResponseBodyTableList) *ListTablesResponseBody
	GetTableList() *ListTablesResponseBodyTableList
	SetTotalCount(v int64) *ListTablesResponseBody
	GetTotalCount() *int64
}

type ListTablesResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B16FB618-5E96-4FFD-BB0D-490C890A4030
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the tables.
	TableList *ListTablesResponseBodyTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	// The total number of tables that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTablesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTablesResponseBody) GetTableList() *ListTablesResponseBodyTableList {
	return s.TableList
}

func (s *ListTablesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTablesResponseBody) SetErrorCode(v string) *ListTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTablesResponseBody) SetErrorMessage(v string) *ListTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetSuccess(v bool) *ListTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTablesResponseBody) SetTableList(v *ListTablesResponseBodyTableList) *ListTablesResponseBody {
	s.TableList = v
	return s
}

func (s *ListTablesResponseBody) SetTotalCount(v int64) *ListTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyTableList struct {
	Table []*ListTablesResponseBodyTableListTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyTableList) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyTableList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTableList) GetTable() []*ListTablesResponseBodyTableListTable {
	return s.Table
}

func (s *ListTablesResponseBodyTableList) SetTable(v []*ListTablesResponseBodyTableListTable) *ListTablesResponseBodyTableList {
	s.Table = v
	return s
}

func (s *ListTablesResponseBodyTableList) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyTableListTable struct {
	// The ID of the physical database.
	//
	// example:
	//
	// 1860****
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
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
	// InnoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The number of rows in the table. This is a statistical value and does not indicate the actual number of rows.
	//
	// example:
	//
	// 10085
	NumRows *int64 `json:"NumRows,omitempty" xml:"NumRows,omitempty"`
	// The ID list of the table owners.
	OwnerIdList *ListTablesResponseBodyTableListTableOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The nickname list of the table owners.
	OwnerNameList *ListTablesResponseBodyTableListTableOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The storage space that is occupied by the table. This is a statistical value and does not indicate the accurate storage space. Unit: MB.
	//
	// example:
	//
	// 1024
	StoreCapacity *int64 `json:"StoreCapacity,omitempty" xml:"StoreCapacity,omitempty"`
	// The GUID of the table in DMS.
	//
	// example:
	//
	// IDB_44743****.qntest.consumption_records
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The ID of the table.
	//
	// example:
	//
	// 44743****
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The table name.
	//
	// example:
	//
	// consumption_records
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The database in which the table resides.
	//
	// example:
	//
	// qntest
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
	// The type of the table. Default value: NORMAL.
	//
	// example:
	//
	// NORMAL
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ListTablesResponseBodyTableListTable) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyTableListTable) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTableListTable) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListTablesResponseBodyTableListTable) GetDescription() *string {
	return s.Description
}

func (s *ListTablesResponseBodyTableListTable) GetEncoding() *string {
	return s.Encoding
}

func (s *ListTablesResponseBodyTableListTable) GetEngine() *string {
	return s.Engine
}

func (s *ListTablesResponseBodyTableListTable) GetNumRows() *int64 {
	return s.NumRows
}

func (s *ListTablesResponseBodyTableListTable) GetOwnerIdList() *ListTablesResponseBodyTableListTableOwnerIdList {
	return s.OwnerIdList
}

func (s *ListTablesResponseBodyTableListTable) GetOwnerNameList() *ListTablesResponseBodyTableListTableOwnerNameList {
	return s.OwnerNameList
}

func (s *ListTablesResponseBodyTableListTable) GetStoreCapacity() *int64 {
	return s.StoreCapacity
}

func (s *ListTablesResponseBodyTableListTable) GetTableGuid() *string {
	return s.TableGuid
}

func (s *ListTablesResponseBodyTableListTable) GetTableId() *string {
	return s.TableId
}

func (s *ListTablesResponseBodyTableListTable) GetTableName() *string {
	return s.TableName
}

func (s *ListTablesResponseBodyTableListTable) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *ListTablesResponseBodyTableListTable) GetTableType() *string {
	return s.TableType
}

func (s *ListTablesResponseBodyTableListTable) SetDatabaseId(v string) *ListTablesResponseBodyTableListTable {
	s.DatabaseId = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetDescription(v string) *ListTablesResponseBodyTableListTable {
	s.Description = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetEncoding(v string) *ListTablesResponseBodyTableListTable {
	s.Encoding = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetEngine(v string) *ListTablesResponseBodyTableListTable {
	s.Engine = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetNumRows(v int64) *ListTablesResponseBodyTableListTable {
	s.NumRows = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetOwnerIdList(v *ListTablesResponseBodyTableListTableOwnerIdList) *ListTablesResponseBodyTableListTable {
	s.OwnerIdList = v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetOwnerNameList(v *ListTablesResponseBodyTableListTableOwnerNameList) *ListTablesResponseBodyTableListTable {
	s.OwnerNameList = v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetStoreCapacity(v int64) *ListTablesResponseBodyTableListTable {
	s.StoreCapacity = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableGuid(v string) *ListTablesResponseBodyTableListTable {
	s.TableGuid = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableId(v string) *ListTablesResponseBodyTableListTable {
	s.TableId = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableName(v string) *ListTablesResponseBodyTableListTable {
	s.TableName = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableSchemaName(v string) *ListTablesResponseBodyTableListTable {
	s.TableSchemaName = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableType(v string) *ListTablesResponseBodyTableListTable {
	s.TableType = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyTableListTableOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyTableListTableOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyTableListTableOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTableListTableOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *ListTablesResponseBodyTableListTableOwnerIdList) SetOwnerIds(v []*string) *ListTablesResponseBodyTableListTableOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *ListTablesResponseBodyTableListTableOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyTableListTableOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyTableListTableOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyTableListTableOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTableListTableOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *ListTablesResponseBodyTableListTableOwnerNameList) SetOwnerNames(v []*string) *ListTablesResponseBodyTableListTableOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *ListTablesResponseBodyTableListTableOwnerNameList) Validate() error {
	return dara.Validate(s)
}
