// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSlowLogRecordsResponseBody
	GetCode() *string
	SetData(v *DescribeSlowLogRecordsResponseBodyData) *DescribeSlowLogRecordsResponseBody
	GetData() *DescribeSlowLogRecordsResponseBodyData
	SetMessage(v string) *DescribeSlowLogRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSlowLogRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSlowLogRecordsResponseBody
	GetSuccess() *string
}

type DescribeSlowLogRecordsResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// DBLogRecords<SlowLogItem>
	Data *DescribeSlowLogRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A1C79EE2-D04D-571B-8C60-961FAF8E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSlowLogRecordsResponseBody) GetData() *DescribeSlowLogRecordsResponseBodyData {
	return s.Data
}

func (s *DescribeSlowLogRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSlowLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogRecordsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSlowLogRecordsResponseBody) SetCode(v string) *DescribeSlowLogRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetData(v *DescribeSlowLogRecordsResponseBodyData) *DescribeSlowLogRecordsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetMessage(v string) *DescribeSlowLogRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetSuccess(v string) *DescribeSlowLogRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogRecordsResponseBodyData struct {
	// example:
	//
	// 100
	DbInstanceId *int64 `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// rm-bp157g54vy772****
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 1672617600000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	ItemsNumbers *int64                                        `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	Logs         []*DescribeSlowLogRecordsResponseBodyDataLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxRecordsPerPage *int32 `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	// example:
	//
	// node123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 1
	PageNumbers *int32 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// example:
	//
	// 1672531200000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetDbInstanceId() *int64 {
	return s.DbInstanceId
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetItemsNumbers() *int64 {
	return s.ItemsNumbers
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetLogs() []*DescribeSlowLogRecordsResponseBodyDataLogs {
	return s.Logs
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetMaxRecordsPerPage() *int32 {
	return s.MaxRecordsPerPage
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetPageNumbers() *int32 {
	return s.PageNumbers
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetTotalRecords() *int64 {
	return s.TotalRecords
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetDbInstanceId(v int64) *DescribeSlowLogRecordsResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetDbInstanceName(v string) *DescribeSlowLogRecordsResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetEndTime(v string) *DescribeSlowLogRecordsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetItemsNumbers(v int64) *DescribeSlowLogRecordsResponseBodyData {
	s.ItemsNumbers = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetLogs(v []*DescribeSlowLogRecordsResponseBodyDataLogs) *DescribeSlowLogRecordsResponseBodyData {
	s.Logs = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetMaxRecordsPerPage(v int32) *DescribeSlowLogRecordsResponseBodyData {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetNodeId(v string) *DescribeSlowLogRecordsResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetPageNumbers(v int32) *DescribeSlowLogRecordsResponseBodyData {
	s.PageNumbers = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetStartTime(v string) *DescribeSlowLogRecordsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetTotalRecords(v int64) *DescribeSlowLogRecordsResponseBodyData {
	s.TotalRecords = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogRecordsResponseBodyDataLogs struct {
	// example:
	//
	// user1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// MyApp
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// 50
	CPUTime *float64 `json:"CPUTime,omitempty" xml:"CPUTime,omitempty"`
	// example:
	//
	// 100
	CPUTimeSeconds *float64 `json:"CPUTimeSeconds,omitempty" xml:"CPUTimeSeconds,omitempty"`
	// example:
	//
	// SELECT
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// example:
	//
	// test
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// rm-2zebg30mk056g****
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 100
	DocsExamined *string `json:"DocsExamined,omitempty" xml:"DocsExamined,omitempty"`
	// example:
	//
	// 10
	Frows *int64 `json:"Frows,omitempty" xml:"Frows,omitempty"`
	// example:
	//
	// 192.168.1.1
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// example:
	//
	// 1
	IOWrites *int64 `json:"IOWrites,omitempty" xml:"IOWrites,omitempty"`
	// example:
	//
	// test
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// example:
	//
	// valueA
	KeysExamined *string `json:"KeysExamined,omitempty" xml:"KeysExamined,omitempty"`
	// example:
	//
	// 10
	LastRowsCountAffected *int64 `json:"LastRowsCountAffected,omitempty" xml:"LastRowsCountAffected,omitempty"`
	// example:
	//
	// 100
	LockTime *float64 `json:"LockTime,omitempty" xml:"LockTime,omitempty"`
	// example:
	//
	// 100
	LockTimeSeconds *float64 `json:"LockTimeSeconds,omitempty" xml:"LockTimeSeconds,omitempty"`
	// example:
	//
	// 1
	LogicalIOReads *int64 `json:"LogicalIOReads,omitempty" xml:"LogicalIOReads,omitempty"`
	// example:
	//
	// pro-test
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	PhysicalIOReads *int64 `json:"PhysicalIOReads,omitempty" xml:"PhysicalIOReads,omitempty"`
	// example:
	//
	// SELECT 	- FROM my_table WHERE ROWNUM <= 10
	Psql *string `json:"Psql,omitempty" xml:"Psql,omitempty"`
	// example:
	//
	// sq-1pzcdMwRb
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// example:
	//
	// 2024-04-01 11:00:00
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// example:
	//
	// 121
	QueryTime *int64 `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// example:
	//
	// 100
	QueryTimeSeconds *float64 `json:"QueryTimeSeconds,omitempty" xml:"QueryTimeSeconds,omitempty"`
	// example:
	//
	// test
	ReturnItemNumbers *string `json:"ReturnItemNumbers,omitempty" xml:"ReturnItemNumbers,omitempty"`
	// example:
	//
	// 20
	ReturnNum *string `json:"ReturnNum,omitempty" xml:"ReturnNum,omitempty"`
	// example:
	//
	// 20
	Rows *int64 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 10
	RowsCountAffected *int64 `json:"RowsCountAffected,omitempty" xml:"RowsCountAffected,omitempty"`
	// example:
	//
	// 100
	RowsExamined *int64 `json:"RowsExamined,omitempty" xml:"RowsExamined,omitempty"`
	// example:
	//
	// 10
	RowsSent *int64 `json:"RowsSent,omitempty" xml:"RowsSent,omitempty"`
	// example:
	//
	// SELECT 	- FROM my_table WHERE ROWNUM <= 10
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// example:
	//
	// HTTPS
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	// example:
	//
	// 10
	Scnt *int64 `json:"Scnt,omitempty" xml:"Scnt,omitempty"`
	// example:
	//
	// sqlId
	SqlId  *string                                           `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlTag *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag `json:"SqlTag,omitempty" xml:"SqlTag,omitempty" type:"Struct"`
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// example:
	//
	// r-8vb219d10038****
	SubInstanceId *string `json:"SubInstanceId,omitempty" xml:"SubInstanceId,omitempty"`
	// example:
	//
	// tableNameExample
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 6a63b6ac4572abfaef7d1163f684****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 57472578
	ThreadId *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// example:
	//
	// 1747118812
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 074ce334-5247-40b9-b0c1-158aea5d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyDataLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyDataLogs) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetCPUTime() *float64 {
	return s.CPUTime
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetCPUTimeSeconds() *float64 {
	return s.CPUTimeSeconds
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetCommand() *string {
	return s.Command
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetDocsExamined() *string {
	return s.DocsExamined
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetFrows() *int64 {
	return s.Frows
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetIOWrites() *int64 {
	return s.IOWrites
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetInsName() *string {
	return s.InsName
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetKeysExamined() *string {
	return s.KeysExamined
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetLastRowsCountAffected() *int64 {
	return s.LastRowsCountAffected
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetLockTime() *float64 {
	return s.LockTime
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetLockTimeSeconds() *float64 {
	return s.LockTimeSeconds
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetLogicalIOReads() *int64 {
	return s.LogicalIOReads
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetPhysicalIOReads() *int64 {
	return s.PhysicalIOReads
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetPsql() *string {
	return s.Psql
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetQueryTime() *int64 {
	return s.QueryTime
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetQueryTimeSeconds() *float64 {
	return s.QueryTimeSeconds
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetReturnItemNumbers() *string {
	return s.ReturnItemNumbers
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetReturnNum() *string {
	return s.ReturnNum
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetRows() *int64 {
	return s.Rows
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetRowsCountAffected() *int64 {
	return s.RowsCountAffected
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetRowsExamined() *int64 {
	return s.RowsExamined
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetRowsSent() *int64 {
	return s.RowsSent
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetScheme() *string {
	return s.Scheme
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetScnt() *int64 {
	return s.Scnt
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetSqlId() *string {
	return s.SqlId
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetSqlTag() *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag {
	return s.SqlTag
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetSubInstanceId() *string {
	return s.SubInstanceId
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetThreadId() *string {
	return s.ThreadId
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetAccountName(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.AccountName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetApplicationName(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.ApplicationName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetCPUTime(v float64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.CPUTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetCPUTimeSeconds(v float64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.CPUTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetCommand(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.Command = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetDbInstanceName(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetDocsExamined(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.DocsExamined = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetFrows(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.Frows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetIOWrites(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.IOWrites = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetInsName(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.InsName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetKeysExamined(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.KeysExamined = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetLastRowsCountAffected(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.LastRowsCountAffected = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetLockTime(v float64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.LockTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetLockTimeSeconds(v float64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.LockTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetLogicalIOReads(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.LogicalIOReads = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetNamespace(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.Namespace = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetPhysicalIOReads(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.PhysicalIOReads = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetPsql(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.Psql = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetQueryId(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.QueryId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetQueryStartTime(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetQueryTime(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.QueryTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetQueryTimeSeconds(v float64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.QueryTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetReturnItemNumbers(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.ReturnItemNumbers = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetReturnNum(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.ReturnNum = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetRows(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.Rows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetRowsCountAffected(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.RowsCountAffected = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetRowsExamined(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.RowsExamined = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetRowsSent(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.RowsSent = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetScheme(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.Scheme = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetScnt(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.Scnt = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetSqlId(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.SqlId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetSqlTag(v *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.SqlTag = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetSqlType(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.SqlType = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetSubInstanceId(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.SubInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetTableName(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.TableName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetTemplateId(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.TemplateId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetThreadId(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.ThreadId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetTimestamp(v int64) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.Timestamp = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) SetTraceId(v string) *DescribeSlowLogRecordsResponseBodyDataLogs {
	s.TraceId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogs) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogRecordsResponseBodyDataLogsSqlTag struct {
	// example:
	//
	// test
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// sqlid。
	//
	// example:
	//
	// 8ad7069f236bcdaaa9b3ae4b6299****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// example:
	//
	// DAS_IMPORTANT,DAS_IN_PLAN
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) GetComments() *string {
	return s.Comments
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) GetSqlId() *string {
	return s.SqlId
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) GetTags() *string {
	return s.Tags
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) SetComments(v string) *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag {
	s.Comments = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) SetSqlId(v string) *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag {
	s.SqlId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) SetTags(v string) *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag {
	s.Tags = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataLogsSqlTag) Validate() error {
	return dara.Validate(s)
}
