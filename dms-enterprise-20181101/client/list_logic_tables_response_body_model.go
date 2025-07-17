// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListLogicTablesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListLogicTablesResponseBody
	GetErrorMessage() *string
	SetLogicTableList(v *ListLogicTablesResponseBodyLogicTableList) *ListLogicTablesResponseBody
	GetLogicTableList() *ListLogicTablesResponseBodyLogicTableList
	SetRequestId(v string) *ListLogicTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLogicTablesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListLogicTablesResponseBody
	GetTotalCount() *int64
}

type ListLogicTablesResponseBody struct {
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
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details of the logical tables.
	LogicTableList *ListLogicTablesResponseBodyLogicTableList `json:"LogicTableList,omitempty" xml:"LogicTableList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F1E6484F-9DF1-4406-9BDE-0861C4629B69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of logical tables that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLogicTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListLogicTablesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListLogicTablesResponseBody) GetLogicTableList() *ListLogicTablesResponseBodyLogicTableList {
	return s.LogicTableList
}

func (s *ListLogicTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogicTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLogicTablesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLogicTablesResponseBody) SetErrorCode(v string) *ListLogicTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLogicTablesResponseBody) SetErrorMessage(v string) *ListLogicTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListLogicTablesResponseBody) SetLogicTableList(v *ListLogicTablesResponseBodyLogicTableList) *ListLogicTablesResponseBody {
	s.LogicTableList = v
	return s
}

func (s *ListLogicTablesResponseBody) SetRequestId(v string) *ListLogicTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogicTablesResponseBody) SetSuccess(v bool) *ListLogicTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLogicTablesResponseBody) SetTotalCount(v int64) *ListLogicTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLogicTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLogicTablesResponseBodyLogicTableList struct {
	LogicTable []*ListLogicTablesResponseBodyLogicTableListLogicTable `json:"LogicTable,omitempty" xml:"LogicTable,omitempty" type:"Repeated"`
}

func (s ListLogicTablesResponseBodyLogicTableList) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTablesResponseBodyLogicTableList) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBodyLogicTableList) GetLogicTable() []*ListLogicTablesResponseBodyLogicTableListLogicTable {
	return s.LogicTable
}

func (s *ListLogicTablesResponseBodyLogicTableList) SetLogicTable(v []*ListLogicTablesResponseBodyLogicTableListLogicTable) *ListLogicTablesResponseBodyLogicTableList {
	s.LogicTable = v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableList) Validate() error {
	return dara.Validate(s)
}

type ListLogicTablesResponseBodyLogicTableListLogicTable struct {
	// The ID of the logical database.
	//
	// example:
	//
	// 1
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// Indicates whether the table is a logical table. The value is fixed to true.
	//
	// example:
	//
	// true
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The IDs of the owners of the logical tables.
	OwnerIdList *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The nicknames of the owners of the logical tables.
	OwnerNameList *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The logical database to which the logical table belongs.
	//
	// example:
	//
	// yuyang_test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The number of logical tables.
	//
	// example:
	//
	// 4
	TableCount *string `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	// The expression of the logical table.
	//
	// example:
	//
	// test[1-4]
	TableExpr *string `json:"TableExpr,omitempty" xml:"TableExpr,omitempty"`
	// The GUID of the logical table.
	//
	// example:
	//
	// IDB_L_308302.yuyang_test.test_ch
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The ID of the logical table.
	//
	// example:
	//
	// 1
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the logical table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTable) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTable) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetLogic() *bool {
	return s.Logic
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetOwnerIdList() *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList {
	return s.OwnerIdList
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetOwnerNameList() *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList {
	return s.OwnerNameList
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetTableCount() *string {
	return s.TableCount
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetTableExpr() *string {
	return s.TableExpr
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetTableGuid() *string {
	return s.TableGuid
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetTableId() *string {
	return s.TableId
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) GetTableName() *string {
	return s.TableName
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetDatabaseId(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.DatabaseId = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetLogic(v bool) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.Logic = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetOwnerIdList(v *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.OwnerIdList = v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetOwnerNameList(v *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.OwnerNameList = v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetSchemaName(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.SchemaName = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableCount(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableCount = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableExpr(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableExpr = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableGuid(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableGuid = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableId(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableId = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableName(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableName = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) Validate() error {
	return dara.Validate(s)
}

type ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) SetOwnerIds(v []*string) *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) SetOwnerNames(v []*string) *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) Validate() error {
	return dara.Validate(s)
}
