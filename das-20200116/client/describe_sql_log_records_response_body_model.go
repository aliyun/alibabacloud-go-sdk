// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSqlLogRecordsResponseBody
	GetCode() *string
	SetData(v *DescribeSqlLogRecordsResponseBodyData) *DescribeSqlLogRecordsResponseBody
	GetData() *DescribeSqlLogRecordsResponseBodyData
	SetMessage(v string) *DescribeSqlLogRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSqlLogRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSqlLogRecordsResponseBody
	GetSuccess() *string
}

type DescribeSqlLogRecordsResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *DescribeSqlLogRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// >  If the request is successful, **Successful*	- is returned. If the request fails, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F43E7FB3-CE67-5FFD-A59C-EFD278BCD7BE
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

func (s DescribeSqlLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSqlLogRecordsResponseBody) GetData() *DescribeSqlLogRecordsResponseBodyData {
	return s.Data
}

func (s *DescribeSqlLogRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSqlLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlLogRecordsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSqlLogRecordsResponseBody) SetCode(v string) *DescribeSqlLogRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBody) SetData(v *DescribeSqlLogRecordsResponseBodyData) *DescribeSqlLogRecordsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlLogRecordsResponseBody) SetMessage(v string) *DescribeSqlLogRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBody) SetRequestId(v string) *DescribeSqlLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBody) SetSuccess(v string) *DescribeSqlLogRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogRecordsResponseBodyData struct {
	// The end of the time range to query. This value is a UNIX timestamp. Unit: millisecond.
	//
	// example:
	//
	// 1608888296000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the task was complete. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// >  If the value of **Finish*	- is 0 and the value of **JobId*	- is returned, the request is an asynchronous request and the return result cannot be directly obtained. You must query the return result based on the value of **JobId**. Specify JobId as the key of **Filters*	- and the value of **JobId*	- as the value of Filters. Example: `Filters=[{"Key": "JobId", "Value": "******"}]`.
	//
	// example:
	//
	// 1
	Finish *string `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// The data.
	Items *DescribeSqlLogRecordsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// MzI4NTZfUUlOR0RBT19DTTlfTlUyMF9NWVNRTF9PREJTX0xWU18zMjg1Nl9teXNxbF9XZWQgTWFyIDA2IDE0OjUwOjQ3IENTVCAyMDI0XzBfMzBfRXhlY3V0ZVRpbWVfREVTQ19XZWQgTWFyIDA2IDE0OjM1OjQ3IENTVCAyMDI0Xw==_1709708406465
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: millisecond.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalRecords *int64 `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeSqlLogRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogRecordsResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSqlLogRecordsResponseBodyData) GetFinish() *string {
	return s.Finish
}

func (s *DescribeSqlLogRecordsResponseBodyData) GetItems() *DescribeSqlLogRecordsResponseBodyDataItems {
	return s.Items
}

func (s *DescribeSqlLogRecordsResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *DescribeSqlLogRecordsResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSqlLogRecordsResponseBodyData) GetTotalRecords() *int64 {
	return s.TotalRecords
}

func (s *DescribeSqlLogRecordsResponseBodyData) SetEndTime(v int64) *DescribeSqlLogRecordsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyData) SetFinish(v string) *DescribeSqlLogRecordsResponseBodyData {
	s.Finish = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyData) SetItems(v *DescribeSqlLogRecordsResponseBodyDataItems) *DescribeSqlLogRecordsResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyData) SetJobId(v string) *DescribeSqlLogRecordsResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyData) SetStartTime(v int64) *DescribeSqlLogRecordsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyData) SetTotalRecords(v int64) *DescribeSqlLogRecordsResponseBodyData {
	s.TotalRecords = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogRecordsResponseBodyDataItems struct {
	// The SQL log data.
	SQLLogRecord []*DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord `json:"SQLLogRecord,omitempty" xml:"SQLLogRecord,omitempty" type:"Repeated"`
}

func (s DescribeSqlLogRecordsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogRecordsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogRecordsResponseBodyDataItems) GetSQLLogRecord() []*DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	return s.SQLLogRecord
}

func (s *DescribeSqlLogRecordsResponseBodyDataItems) SetSQLLogRecord(v []*DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) *DescribeSqlLogRecordsResponseBodyDataItems {
	s.SQLLogRecord = v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord struct {
	// The account of the database.
	//
	// example:
	//
	// testname
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The amount of time that is consumed to execute the SQL statement. Unit: millisecond.
	//
	// example:
	//
	// 58
	Consume *int64 `json:"Consume,omitempty" xml:"Consume,omitempty"`
	// The CPU execution duration. Unit: microsecond.
	//
	// example:
	//
	// 100
	CpuTime *int64 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	// The database name.
	//
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The time when the SQL statement was executed. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
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
	// The number of rows that are pulled by the compute nodes of the PolarDB-X 2.0 instance.
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
	// The lock wait duration. Unit: millisecond.
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
	// The node ID.
	//
	// example:
	//
	// pi-uf6k5f6g3912i****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The timestamp generated when the SQL statement was executed. The value of this parameter is a UNIX timestamp. Unit: millisecond.
	//
	// example:
	//
	// 1701886532000
	OriginTime *int64 `json:"OriginTime,omitempty" xml:"OriginTime,omitempty"`
	// The parallel queue time of the PolarDB for MySQL instance. Unit: millisecond.
	//
	// example:
	//
	// 10
	ParallelDegree *string `json:"ParallelDegree,omitempty" xml:"ParallelDegree,omitempty"`
	// The parallelism of the PolarDB for MySQL cluster.
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
	// The number of rows returned by the SQL statement.
	//
	// example:
	//
	// 0
	ReturnRows *int64 `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	// The total number of rows that are updated or returned by the compute nodes of the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 10
	Rows *int64 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The number of scanned rows.
	//
	// example:
	//
	// 0
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The number of requests that are sent from the compute nodes to the data nodes of the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 10
	Scnt *int64 `json:"Scnt,omitempty" xml:"Scnt,omitempty"`
	// The SQL statement ID.
	//
	// example:
	//
	// c67649d4a7fb62c4f8c7a447c52b5b17
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// select resource_id as cluster_id, tpl_name \\n\\tfrom dbfree_alert_resource_tpl_ref\\n\\twhere user_id=? and type=\\"cluster\\" group by resource_id, tpl_name
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// The type of the SQL statement.
	//
	// example:
	//
	// select
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The execution status of the SQL statement.
	//
	// 	- **0**: The execution was successful.
	//
	// 	- **1**: The execution failed.
	//
	// example:
	//
	// 0
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The thread ID.
	//
	// example:
	//
	// None
	ThreadId *int64 `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// The trace ID of the PolarDB-X 2.0 instance. The value is the execution ID of the SQL statement on the data node.
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
	// The number of rows that are updated.
	//
	// example:
	//
	// 0
	UpdateRows *int64 `json:"UpdateRows,omitempty" xml:"UpdateRows,omitempty"`
	// Indicates whether the In-Memory Column Index (IMCI) feature is enabled for the PolarDB for MySQL cluster. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	UseImciEngine *string `json:"UseImciEngine,omitempty" xml:"UseImciEngine,omitempty"`
	// The IP address that is resolved from the endpoint of the query link.
	//
	// example:
	//
	// 100.115.XX.XX
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
	// The number of writes to the ApsaraDB RDS for SQL Server instance.
	//
	// example:
	//
	// 10
	Writes *int64 `json:"Writes,omitempty" xml:"Writes,omitempty"`
}

func (s DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetCollection() *string {
	return s.Collection
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetConsume() *int64 {
	return s.Consume
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetCpuTime() *int64 {
	return s.CpuTime
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetExt() *string {
	return s.Ext
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetFrows() *int64 {
	return s.Frows
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetLockTime() *int64 {
	return s.LockTime
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetLogicRead() *int64 {
	return s.LogicRead
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetOriginTime() *int64 {
	return s.OriginTime
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetParallelDegree() *string {
	return s.ParallelDegree
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetParallelQueueTime() *string {
	return s.ParallelQueueTime
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetPhysicAsyncRead() *int64 {
	return s.PhysicAsyncRead
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetPhysicRead() *int64 {
	return s.PhysicRead
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetPhysicSyncRead() *int64 {
	return s.PhysicSyncRead
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetReturnRows() *int64 {
	return s.ReturnRows
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetRows() *int64 {
	return s.Rows
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetScnt() *int64 {
	return s.Scnt
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetSqlId() *string {
	return s.SqlId
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetSqlText() *string {
	return s.SqlText
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetState() *string {
	return s.State
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetThreadId() *int64 {
	return s.ThreadId
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetTrxId() *string {
	return s.TrxId
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetUpdateRows() *int64 {
	return s.UpdateRows
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetUseImciEngine() *string {
	return s.UseImciEngine
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetVip() *string {
	return s.Vip
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetWrites() *int64 {
	return s.Writes
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetAccountName(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.AccountName = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetCollection(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Collection = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetConsume(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Consume = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetCpuTime(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.CpuTime = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetDBName(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetExecuteTime(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetExt(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Ext = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetFrows(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Frows = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetHostAddress(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetLockTime(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.LockTime = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetLogicRead(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.LogicRead = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetNodeId(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.NodeId = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetOriginTime(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.OriginTime = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetParallelDegree(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ParallelDegree = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetParallelQueueTime(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ParallelQueueTime = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetPhysicAsyncRead(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.PhysicAsyncRead = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetPhysicRead(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.PhysicRead = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetPhysicSyncRead(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.PhysicSyncRead = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetReturnRows(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetRows(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Rows = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetScanRows(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ScanRows = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetScnt(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Scnt = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetSqlId(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.SqlId = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetSqlText(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.SqlText = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetSqlType(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.SqlType = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetState(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.State = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetTableName(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.TableName = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetThreadId(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ThreadId = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetTraceId(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.TraceId = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetTrxId(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.TrxId = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetUpdateRows(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.UpdateRows = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetUseImciEngine(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.UseImciEngine = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetVip(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Vip = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetWrites(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Writes = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) Validate() error {
	return dara.Validate(s)
}
