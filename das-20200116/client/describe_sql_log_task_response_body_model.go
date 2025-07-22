// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSqlLogTaskResponseBody
	GetCode() *string
	SetData(v *DescribeSqlLogTaskResponseBodyData) *DescribeSqlLogTaskResponseBody
	GetData() *DescribeSqlLogTaskResponseBodyData
	SetMessage(v string) *DescribeSqlLogTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSqlLogTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSqlLogTaskResponseBody
	GetSuccess() *string
}

type DescribeSqlLogTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeSqlLogTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSqlLogTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSqlLogTaskResponseBody) GetData() *DescribeSqlLogTaskResponseBodyData {
	return s.Data
}

func (s *DescribeSqlLogTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSqlLogTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlLogTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSqlLogTaskResponseBody) SetCode(v string) *DescribeSqlLogTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBody) SetData(v *DescribeSqlLogTaskResponseBodyData) *DescribeSqlLogTaskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlLogTaskResponseBody) SetMessage(v string) *DescribeSqlLogTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBody) SetRequestId(v string) *DescribeSqlLogTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBody) SetSuccess(v string) *DescribeSqlLogTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogTaskResponseBodyData struct {
	// The time when the task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1681363254423
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// Indicates whether the task has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Expire *bool `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The download URL of the export task.
	//
	// example:
	//
	// "https://das-sqllog-download-cn-hongkong.oss-cn-hongkong.aliyuncs.com/****"
	Export *string `json:"Export,omitempty" xml:"Export,omitempty"`
	// The filter parameters.
	Filters []*DescribeSqlLogTaskResponseBodyDataFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The task name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The results of the offline querying task.
	Queries []*DescribeSqlLogTaskResponseBodyDataQueries `json:"Queries,omitempty" xml:"Queries,omitempty" type:"Repeated"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1596177993000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The task state. Valid values:
	//
	// 	- **INIT**: The task is to be scheduled.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **FAILED**: The task failed.
	//
	// 	- **CANCELED**: The task is canceled.
	//
	// 	- **COMPLETED**: The task is complete.
	//
	// >  If a task is in the **COMPLETED*	- state, you can view the results of the task.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 9a4f5c4494dbd6713185d87a97aa53e8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// 	- **Export**
	//
	// 	- **Query**
	//
	// example:
	//
	// Query
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeSqlLogTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTaskResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSqlLogTaskResponseBodyData) GetEnd() *int64 {
	return s.End
}

func (s *DescribeSqlLogTaskResponseBodyData) GetExpire() *bool {
	return s.Expire
}

func (s *DescribeSqlLogTaskResponseBodyData) GetExport() *string {
	return s.Export
}

func (s *DescribeSqlLogTaskResponseBodyData) GetFilters() []*DescribeSqlLogTaskResponseBodyDataFilters {
	return s.Filters
}

func (s *DescribeSqlLogTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeSqlLogTaskResponseBodyData) GetQueries() []*DescribeSqlLogTaskResponseBodyDataQueries {
	return s.Queries
}

func (s *DescribeSqlLogTaskResponseBodyData) GetStart() *int64 {
	return s.Start
}

func (s *DescribeSqlLogTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeSqlLogTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSqlLogTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeSqlLogTaskResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeSqlLogTaskResponseBodyData) SetCreateTime(v int64) *DescribeSqlLogTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetEnd(v int64) *DescribeSqlLogTaskResponseBodyData {
	s.End = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetExpire(v bool) *DescribeSqlLogTaskResponseBodyData {
	s.Expire = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetExport(v string) *DescribeSqlLogTaskResponseBodyData {
	s.Export = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetFilters(v []*DescribeSqlLogTaskResponseBodyDataFilters) *DescribeSqlLogTaskResponseBodyData {
	s.Filters = v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetName(v string) *DescribeSqlLogTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetQueries(v []*DescribeSqlLogTaskResponseBodyDataQueries) *DescribeSqlLogTaskResponseBodyData {
	s.Queries = v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetStart(v int64) *DescribeSqlLogTaskResponseBodyData {
	s.Start = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetStatus(v string) *DescribeSqlLogTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetTaskId(v string) *DescribeSqlLogTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetTaskType(v string) *DescribeSqlLogTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) SetTotal(v int64) *DescribeSqlLogTaskResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogTaskResponseBodyDataFilters struct {
	// The name of the filter parameter.
	//
	// >  For more information about the filter parameters, see the **Valid values of Key*	- section of this topic.
	//
	// example:
	//
	// keyWords
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter parameter.
	//
	// example:
	//
	// select
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSqlLogTaskResponseBodyDataFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTaskResponseBodyDataFilters) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTaskResponseBodyDataFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeSqlLogTaskResponseBodyDataFilters) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSqlLogTaskResponseBodyDataFilters) SetKey(v string) *DescribeSqlLogTaskResponseBodyDataFilters {
	s.Key = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataFilters) SetValue(v interface{}) *DescribeSqlLogTaskResponseBodyDataFilters {
	s.Value = v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataFilters) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogTaskResponseBodyDataQueries struct {
	// The database account.
	//
	// example:
	//
	// testname
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Collection  *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The execution duration. Unit: millisecond.
	//
	// example:
	//
	// 58
	Consume *int64 `json:"Consume,omitempty" xml:"Consume,omitempty"`
	// The CPU execution time. Unit: microsecond.
	//
	// example:
	//
	// 100
	CpuTime *int64 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	// The database name.
	//
	// example:
	//
	// testdb01
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The execution time. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-12-07T02:15:32Z
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The extended information. This parameter is a reserved parameter.
	//
	// example:
	//
	// None
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The number of rows pulled by the CNs of the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 10
	Frows *int64 `json:"Frows,omitempty" xml:"Frows,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 11.197.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The lock wait time. Unit: millisecond.
	//
	// example:
	//
	// 0
	LockTime *int64 `json:"LockTime,omitempty" xml:"LockTime,omitempty"`
	// The number of logical reads.
	//
	// example:
	//
	// 0
	LogicRead *int64 `json:"LogicRead,omitempty" xml:"LogicRead,omitempty"`
	// The ID of the child node.
	//
	// example:
	//
	// pi-bp1o58x3ib7e6****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The execution timestamp. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1701886532000
	OriginTime *int64 `json:"OriginTime,omitempty" xml:"OriginTime,omitempty"`
	// The wait time of parallel queries in the queue in the PolarDB for MySQL instance. Unit: millisecond.
	//
	// example:
	//
	// 10
	ParallelDegree *string `json:"ParallelDegree,omitempty" xml:"ParallelDegree,omitempty"`
	// The degree of parallelism (DOP) value of the PolarDB for MySQL instance.
	//
	// example:
	//
	// 2
	ParallelQueueTime *string `json:"ParallelQueueTime,omitempty" xml:"ParallelQueueTime,omitempty"`
	// The number of physical asynchronous reads.
	//
	// example:
	//
	// 0
	PhysicAsyncRead *int64 `json:"PhysicAsyncRead,omitempty" xml:"PhysicAsyncRead,omitempty"`
	// The total number of physical reads.
	//
	// example:
	//
	// 0
	PhysicRead *int64 `json:"PhysicRead,omitempty" xml:"PhysicRead,omitempty"`
	// The number of physical synchronous reads.
	//
	// example:
	//
	// 0
	PhysicSyncRead *int64 `json:"PhysicSyncRead,omitempty" xml:"PhysicSyncRead,omitempty"`
	// The number of rows returned.
	//
	// example:
	//
	// 0
	ReturnRows *int64 `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	// The total number of rows updated or returned by the CNs of the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 10
	Rows *int64 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The number of rows scanned.
	//
	// example:
	//
	// 0
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The number of requests from the compute nodes (CNs) to the data nodes (DNs) in the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 10
	Scnt       *int64 `json:"Scnt,omitempty" xml:"Scnt,omitempty"`
	SqlCommand *int64 `json:"SqlCommand,omitempty" xml:"SqlCommand,omitempty"`
	// The ID of the SQL statement.
	//
	// example:
	//
	// a4111670e80596c5bf42cf5154438a91
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The queried SQL statement.
	//
	// example:
	//
	// SELECT @@session.transaction_read_only
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
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
	// The execution result of the SQL statement. Valid values:
	//
	// 	- **0**: The execution was successful.
	//
	// 	- **1**: The execution failed.
	//
	// example:
	//
	// 0
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The thread ID.
	//
	// example:
	//
	// None
	ThreadId *int64 `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// The trace ID of the PolarDB-X 2.0 instance, which is the execution ID of the SQL statement on the DN.
	//
	// example:
	//
	// 14c93b7c7bf00000
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The transaction ID.
	//
	// example:
	//
	// 200000
	TrxId *string `json:"TrxId,omitempty" xml:"TrxId,omitempty"`
	// The number of rows updated.
	//
	// example:
	//
	// 0
	UpdateRows *int64 `json:"UpdateRows,omitempty" xml:"UpdateRows,omitempty"`
	// Indicates whether the PolarDB for MySQL instance uses In-Memory Column Indexes (IMCIs). Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	UseImciEngine *string `json:"UseImciEngine,omitempty" xml:"UseImciEngine,omitempty"`
	// The IP address to which the endpoint used for query is resolved.
	//
	// example:
	//
	// 10.146.XX.XX
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
	// The number of writes to the ApsaraDB RDS for SQL Server instance.
	//
	// example:
	//
	// 10
	Writes *int64 `json:"Writes,omitempty" xml:"Writes,omitempty"`
}

func (s DescribeSqlLogTaskResponseBodyDataQueries) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTaskResponseBodyDataQueries) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetCollection() *string {
	return s.Collection
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetConsume() *int64 {
	return s.Consume
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetCpuTime() *int64 {
	return s.CpuTime
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetExt() *string {
	return s.Ext
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetFrows() *int64 {
	return s.Frows
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetLockTime() *int64 {
	return s.LockTime
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetLogicRead() *int64 {
	return s.LogicRead
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetOriginTime() *int64 {
	return s.OriginTime
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetParallelDegree() *string {
	return s.ParallelDegree
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetParallelQueueTime() *string {
	return s.ParallelQueueTime
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetPhysicAsyncRead() *int64 {
	return s.PhysicAsyncRead
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetPhysicRead() *int64 {
	return s.PhysicRead
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetPhysicSyncRead() *int64 {
	return s.PhysicSyncRead
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetReturnRows() *int64 {
	return s.ReturnRows
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetRows() *int64 {
	return s.Rows
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetScnt() *int64 {
	return s.Scnt
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetSqlCommand() *int64 {
	return s.SqlCommand
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetSqlId() *string {
	return s.SqlId
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetSqlText() *string {
	return s.SqlText
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetState() *string {
	return s.State
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetThreadId() *int64 {
	return s.ThreadId
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetTrxId() *string {
	return s.TrxId
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetUpdateRows() *int64 {
	return s.UpdateRows
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetUseImciEngine() *string {
	return s.UseImciEngine
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetVip() *string {
	return s.Vip
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) GetWrites() *int64 {
	return s.Writes
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetAccountName(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.AccountName = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetCollection(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.Collection = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetConsume(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.Consume = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetCpuTime(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.CpuTime = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetDBName(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.DBName = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetExecuteTime(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetExt(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.Ext = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetFrows(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.Frows = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetHostAddress(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.HostAddress = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetLockTime(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.LockTime = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetLogicRead(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.LogicRead = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetNodeId(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.NodeId = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetOriginTime(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.OriginTime = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetParallelDegree(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.ParallelDegree = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetParallelQueueTime(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.ParallelQueueTime = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetPhysicAsyncRead(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.PhysicAsyncRead = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetPhysicRead(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.PhysicRead = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetPhysicSyncRead(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.PhysicSyncRead = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetReturnRows(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetRows(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.Rows = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetScanRows(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.ScanRows = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetScnt(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.Scnt = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetSqlCommand(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.SqlCommand = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetSqlId(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.SqlId = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetSqlText(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.SqlText = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetSqlType(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.SqlType = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetState(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.State = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetThreadId(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.ThreadId = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetTraceId(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.TraceId = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetTrxId(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.TrxId = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetUpdateRows(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.UpdateRows = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetUseImciEngine(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.UseImciEngine = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetVip(v string) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.Vip = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) SetWrites(v int64) *DescribeSqlLogTaskResponseBodyDataQueries {
	s.Writes = &v
	return s
}

func (s *DescribeSqlLogTaskResponseBodyDataQueries) Validate() error {
	return dara.Validate(s)
}
