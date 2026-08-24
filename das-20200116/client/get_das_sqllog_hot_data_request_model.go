// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasSQLLogHotDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetDasSQLLogHotDataRequest
	GetAccountName() *string
	SetChildDBInstanceIDs(v string) *GetDasSQLLogHotDataRequest
	GetChildDBInstanceIDs() *string
	SetDBName(v string) *GetDasSQLLogHotDataRequest
	GetDBName() *string
	SetEnd(v int64) *GetDasSQLLogHotDataRequest
	GetEnd() *int64
	SetFail(v string) *GetDasSQLLogHotDataRequest
	GetFail() *string
	SetHostAddress(v string) *GetDasSQLLogHotDataRequest
	GetHostAddress() *string
	SetInstanceId(v string) *GetDasSQLLogHotDataRequest
	GetInstanceId() *string
	SetLogicalOperator(v string) *GetDasSQLLogHotDataRequest
	GetLogicalOperator() *string
	SetMaxLatancy(v int64) *GetDasSQLLogHotDataRequest
	GetMaxLatancy() *int64
	SetMaxRecordsPerPage(v int64) *GetDasSQLLogHotDataRequest
	GetMaxRecordsPerPage() *int64
	SetMaxRows(v int64) *GetDasSQLLogHotDataRequest
	GetMaxRows() *int64
	SetMaxScanRows(v int64) *GetDasSQLLogHotDataRequest
	GetMaxScanRows() *int64
	SetMaxSpillCnt(v int64) *GetDasSQLLogHotDataRequest
	GetMaxSpillCnt() *int64
	SetMinLatancy(v int64) *GetDasSQLLogHotDataRequest
	GetMinLatancy() *int64
	SetMinRows(v int64) *GetDasSQLLogHotDataRequest
	GetMinRows() *int64
	SetMinScanRows(v int64) *GetDasSQLLogHotDataRequest
	GetMinScanRows() *int64
	SetMinSpillCnt(v int64) *GetDasSQLLogHotDataRequest
	GetMinSpillCnt() *int64
	SetPageNumbers(v int64) *GetDasSQLLogHotDataRequest
	GetPageNumbers() *int64
	SetQueryKeyword(v string) *GetDasSQLLogHotDataRequest
	GetQueryKeyword() *string
	SetRole(v string) *GetDasSQLLogHotDataRequest
	GetRole() *string
	SetSortKey(v string) *GetDasSQLLogHotDataRequest
	GetSortKey() *string
	SetSortMethod(v string) *GetDasSQLLogHotDataRequest
	GetSortMethod() *string
	SetSqlType(v string) *GetDasSQLLogHotDataRequest
	GetSqlType() *string
	SetStart(v int64) *GetDasSQLLogHotDataRequest
	GetStart() *int64
	SetState(v string) *GetDasSQLLogHotDataRequest
	GetState() *string
	SetThreadID(v string) *GetDasSQLLogHotDataRequest
	GetThreadID() *string
	SetTraceId(v string) *GetDasSQLLogHotDataRequest
	GetTraceId() *string
	SetTransactionId(v string) *GetDasSQLLogHotDataRequest
	GetTransactionId() *string
}

type GetDasSQLLogHotDataRequest struct {
	// The database account.
	//
	// > You can specify multiple database accounts. Separate multiple accounts with a space. For example: `user1 user2 user3`.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The node ID.
	//
	// > This parameter is required if the database instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-bp179lg03445l****
	ChildDBInstanceIDs *string `json:"ChildDBInstanceIDs,omitempty" xml:"ChildDBInstanceIDs,omitempty"`
	// The database name.
	//
	// > You can specify multiple database names. Separate multiple names with a space. For example: `DB1 DB2 DB3`.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. This value must be a Unix timestamp in milliseconds.
	//
	// > The end time must be later than the start time. The time range cannot exceed one day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684820697000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The SQL execution error code. You can call the [GetAsyncErrorRequestStatByCode](https://help.aliyun.com/document_detail/409804.html) operation to obtain the error code.
	//
	// example:
	//
	// 1064
	Fail *string `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// The client IP address.
	//
	// > You can specify multiple client IP addresses. Separate multiple IP addresses with a space. For example: `IP1 IP2 IP3`.
	//
	// example:
	//
	// 47.100.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The logical operator to use with multiple keywords. Valid values:
	//
	// - **or**
	//
	// - **and**
	//
	// example:
	//
	// or
	LogicalOperator *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	// The maximum execution time in microseconds. Returns SQL statements that have an execution time less than this value.
	//
	// example:
	//
	// 100
	MaxLatancy *int64 `json:"MaxLatancy,omitempty" xml:"MaxLatancy,omitempty"`
	// The maximum number of entries per page. Valid values: 5 to 100.
	//
	// example:
	//
	// 10
	MaxRecordsPerPage *int64 `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	MaxRows *int64 `json:"MaxRows,omitempty" xml:"MaxRows,omitempty"`
	// The maximum number of scanned rows. Returns SQL statements that scanned fewer than this number of rows.
	//
	// example:
	//
	// 10000
	MaxScanRows *int64 `json:"MaxScanRows,omitempty" xml:"MaxScanRows,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	MaxSpillCnt *int64 `json:"MaxSpillCnt,omitempty" xml:"MaxSpillCnt,omitempty"`
	// The minimum execution time in microseconds. Returns SQL statements with an execution time greater than or equal to this value.
	//
	// example:
	//
	// 10
	MinLatancy *int64 `json:"MinLatancy,omitempty" xml:"MinLatancy,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	MinRows *int64 `json:"MinRows,omitempty" xml:"MinRows,omitempty"`
	// The minimum number of scanned rows. Returns SQL statements that scanned at least this number of rows.
	//
	// example:
	//
	// 10
	MinScanRows *int64 `json:"MinScanRows,omitempty" xml:"MinScanRows,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	MinSpillCnt *int64 `json:"MinSpillCnt,omitempty" xml:"MinSpillCnt,omitempty"`
	// The page number to return. Pages start from 1. The default value is 1.
	//
	// example:
	//
	// 2
	PageNumbers *int64 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// The query keyword.
	//
	// > Fuzzy search is supported. You can specify up to 10 keywords. Separate multiple keywords with a space. For example: a1 b2 c3.
	//
	// example:
	//
	// a1 b2
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The sort key. Valid values:
	//
	// - **ScanRows**: scanned rows.
	//
	// - **UpdateRows**: updated rows.
	//
	// - **Consume**: execution time.
	//
	// - **OriginTime**: The execution start time.
	//
	// - **ReturnRows**: returned rows.
	//
	// example:
	//
	// ScanRows
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	// The sort order. Valid values:
	//
	// - **ASC**: ascending
	//
	// - **DESC**: descending
	//
	// example:
	//
	// ASC
	SortMethod *string `json:"SortMethod,omitempty" xml:"SortMethod,omitempty"`
	// The SQL type.
	//
	// example:
	//
	// select
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start of the time range to query. This value must be a Unix timestamp in milliseconds.
	//
	// > You can query only data that is generated after you enable DAS Enterprise Edition. The start time cannot be earlier than seven days before the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684734297000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The execution state. Set this parameter to **0*	- to query for successfully executed SQL statements. You can also specify an error code to query for the corresponding SQL statements.
	//
	// example:
	//
	// 0
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The thread ID.
	//
	// > You can specify multiple thread IDs. Separate multiple IDs with a space. For example: `657 658 659`.
	//
	// example:
	//
	// 657
	ThreadID *string `json:"ThreadID,omitempty" xml:"ThreadID,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The transaction ID.
	//
	// example:
	//
	// 0
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s GetDasSQLLogHotDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDasSQLLogHotDataRequest) GoString() string {
	return s.String()
}

func (s *GetDasSQLLogHotDataRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetDasSQLLogHotDataRequest) GetChildDBInstanceIDs() *string {
	return s.ChildDBInstanceIDs
}

func (s *GetDasSQLLogHotDataRequest) GetDBName() *string {
	return s.DBName
}

func (s *GetDasSQLLogHotDataRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetDasSQLLogHotDataRequest) GetFail() *string {
	return s.Fail
}

func (s *GetDasSQLLogHotDataRequest) GetHostAddress() *string {
	return s.HostAddress
}

func (s *GetDasSQLLogHotDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDasSQLLogHotDataRequest) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *GetDasSQLLogHotDataRequest) GetMaxLatancy() *int64 {
	return s.MaxLatancy
}

func (s *GetDasSQLLogHotDataRequest) GetMaxRecordsPerPage() *int64 {
	return s.MaxRecordsPerPage
}

func (s *GetDasSQLLogHotDataRequest) GetMaxRows() *int64 {
	return s.MaxRows
}

func (s *GetDasSQLLogHotDataRequest) GetMaxScanRows() *int64 {
	return s.MaxScanRows
}

func (s *GetDasSQLLogHotDataRequest) GetMaxSpillCnt() *int64 {
	return s.MaxSpillCnt
}

func (s *GetDasSQLLogHotDataRequest) GetMinLatancy() *int64 {
	return s.MinLatancy
}

func (s *GetDasSQLLogHotDataRequest) GetMinRows() *int64 {
	return s.MinRows
}

func (s *GetDasSQLLogHotDataRequest) GetMinScanRows() *int64 {
	return s.MinScanRows
}

func (s *GetDasSQLLogHotDataRequest) GetMinSpillCnt() *int64 {
	return s.MinSpillCnt
}

func (s *GetDasSQLLogHotDataRequest) GetPageNumbers() *int64 {
	return s.PageNumbers
}

func (s *GetDasSQLLogHotDataRequest) GetQueryKeyword() *string {
	return s.QueryKeyword
}

func (s *GetDasSQLLogHotDataRequest) GetRole() *string {
	return s.Role
}

func (s *GetDasSQLLogHotDataRequest) GetSortKey() *string {
	return s.SortKey
}

func (s *GetDasSQLLogHotDataRequest) GetSortMethod() *string {
	return s.SortMethod
}

func (s *GetDasSQLLogHotDataRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *GetDasSQLLogHotDataRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetDasSQLLogHotDataRequest) GetState() *string {
	return s.State
}

func (s *GetDasSQLLogHotDataRequest) GetThreadID() *string {
	return s.ThreadID
}

func (s *GetDasSQLLogHotDataRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *GetDasSQLLogHotDataRequest) GetTransactionId() *string {
	return s.TransactionId
}

func (s *GetDasSQLLogHotDataRequest) SetAccountName(v string) *GetDasSQLLogHotDataRequest {
	s.AccountName = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetChildDBInstanceIDs(v string) *GetDasSQLLogHotDataRequest {
	s.ChildDBInstanceIDs = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetDBName(v string) *GetDasSQLLogHotDataRequest {
	s.DBName = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetEnd(v int64) *GetDasSQLLogHotDataRequest {
	s.End = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetFail(v string) *GetDasSQLLogHotDataRequest {
	s.Fail = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetHostAddress(v string) *GetDasSQLLogHotDataRequest {
	s.HostAddress = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetInstanceId(v string) *GetDasSQLLogHotDataRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetLogicalOperator(v string) *GetDasSQLLogHotDataRequest {
	s.LogicalOperator = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMaxLatancy(v int64) *GetDasSQLLogHotDataRequest {
	s.MaxLatancy = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMaxRecordsPerPage(v int64) *GetDasSQLLogHotDataRequest {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMaxRows(v int64) *GetDasSQLLogHotDataRequest {
	s.MaxRows = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMaxScanRows(v int64) *GetDasSQLLogHotDataRequest {
	s.MaxScanRows = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMaxSpillCnt(v int64) *GetDasSQLLogHotDataRequest {
	s.MaxSpillCnt = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMinLatancy(v int64) *GetDasSQLLogHotDataRequest {
	s.MinLatancy = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMinRows(v int64) *GetDasSQLLogHotDataRequest {
	s.MinRows = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMinScanRows(v int64) *GetDasSQLLogHotDataRequest {
	s.MinScanRows = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetMinSpillCnt(v int64) *GetDasSQLLogHotDataRequest {
	s.MinSpillCnt = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetPageNumbers(v int64) *GetDasSQLLogHotDataRequest {
	s.PageNumbers = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetQueryKeyword(v string) *GetDasSQLLogHotDataRequest {
	s.QueryKeyword = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetRole(v string) *GetDasSQLLogHotDataRequest {
	s.Role = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetSortKey(v string) *GetDasSQLLogHotDataRequest {
	s.SortKey = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetSortMethod(v string) *GetDasSQLLogHotDataRequest {
	s.SortMethod = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetSqlType(v string) *GetDasSQLLogHotDataRequest {
	s.SqlType = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetStart(v int64) *GetDasSQLLogHotDataRequest {
	s.Start = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetState(v string) *GetDasSQLLogHotDataRequest {
	s.State = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetThreadID(v string) *GetDasSQLLogHotDataRequest {
	s.ThreadID = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetTraceId(v string) *GetDasSQLLogHotDataRequest {
	s.TraceId = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) SetTransactionId(v string) *GetDasSQLLogHotDataRequest {
	s.TransactionId = &v
	return s
}

func (s *GetDasSQLLogHotDataRequest) Validate() error {
	return dara.Validate(s)
}
