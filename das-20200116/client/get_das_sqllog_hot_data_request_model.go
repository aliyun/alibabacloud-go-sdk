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
	// The account of the database.
	//
	// >  You can specify multiple database accounts that are separated by spaces. Example: `user1 user2 user3`.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The node ID.
	//
	// >  This parameter must be specified if the database instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-bp179lg03445l****
	ChildDBInstanceIDs *string `json:"ChildDBInstanceIDs,omitempty" xml:"ChildDBInstanceIDs,omitempty"`
	// The name of the database.
	//
	// >  You can specify multiple database names that are separated by spaces. Example: `DB1 DB2 DB3`.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. The interval between the start time and the end time cannot exceed 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684820697000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The error code of SQL execution. You can call the [GetAsyncErrorRequestStatByCode](https://help.aliyun.com/document_detail/409804.html) operation to query MySQL error codes in SQL Explorer data.
	//
	// example:
	//
	// 1064
	Fail *string `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// The IP address of the client.
	//
	// >  You can specify multiple IP addresses that are separated by spaces. Example: `IP1 IP2 IP3`.
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
	// The logical relationship among multiple keywords.
	//
	// 	- **or**
	//
	// 	- **and**
	//
	// example:
	//
	// or
	LogicalOperator *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	// The maximum execution duration. Unit: microseconds. You can specify this parameter to query the SQL statements whose execution duration is smaller than the value of this parameter.
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
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	MaxRows *int64 `json:"MaxRows,omitempty" xml:"MaxRows,omitempty"`
	// The maximum number of scanned rows. You can specify this parameter to query the SQL statements that scan a smaller number of rows than the value of this parameter.
	//
	// example:
	//
	// 10000
	MaxScanRows *int64 `json:"MaxScanRows,omitempty" xml:"MaxScanRows,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	MaxSpillCnt *int64 `json:"MaxSpillCnt,omitempty" xml:"MaxSpillCnt,omitempty"`
	// The minimum execution duration. Unit: microseconds. You can specify this parameter to query the SQL statements whose execution duration is greater than or equal to the value of this parameter.
	//
	// example:
	//
	// 10
	MinLatancy *int64 `json:"MinLatancy,omitempty" xml:"MinLatancy,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	MinRows *int64 `json:"MinRows,omitempty" xml:"MinRows,omitempty"`
	// The minimum number of scanned rows. You can specify this parameter to query the SQL statements that scan a larger or an equal number of rows than the value of this parameter.
	//
	// example:
	//
	// 10
	MinScanRows *int64 `json:"MinScanRows,omitempty" xml:"MinScanRows,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	MinSpillCnt *int64 `json:"MinSpillCnt,omitempty" xml:"MinSpillCnt,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 2
	PageNumbers *int64 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// The keyword that is used for the query.
	//
	// >  Fuzzy search is not supported. You can query data by using multiple keywords. Separate keywords with spaces.
	//
	// example:
	//
	// test
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The basis on which you want to sort the query results.
	//
	// 	- **SCAN_ROWS**: the number of scanned rows.
	//
	// 	- **UPDATE_ROWS**: the number of updated rows.
	//
	// 	- **CONSUME**: the time consumed.
	//
	// 	- **ORIGIN_TIME**: the execution duration.
	//
	// example:
	//
	// SCAN_ROWS
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	// The order in which you want to sort the query results.
	//
	// 	- **ase**: ascending order.
	//
	// 	- **desc**: descending order.
	//
	// example:
	//
	// ase
	SortMethod *string `json:"SortMethod,omitempty" xml:"SortMethod,omitempty"`
	// The type of the SQL statement. Valid values:
	//
	// 	- **SELECT**
	//
	// 	- **UPDATE**
	//
	// 	- **DELETE**
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The beginning of the time range to query must be later than the time when DAS Enterprise Edition is enabled, and can be up to seven days earlier than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684734297000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The execution results. You can specify **0*	- to query the SQL statements that are successfully executed. You can also specify an error code to query the corresponding SQL statements that fail to be executed.
	//
	// example:
	//
	// 0
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The thread ID.
	//
	// >  You can specify multiple thread IDs that are separated by spaces. Example: `Thread ID1 Thread ID2 Thread ID3`.
	//
	// example:
	//
	// 657
	ThreadID *string `json:"ThreadID,omitempty" xml:"ThreadID,omitempty"`
	// The reserved parameter. This parameter is not supported.
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
