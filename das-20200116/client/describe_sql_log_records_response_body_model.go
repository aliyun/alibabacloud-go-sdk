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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeSqlLogRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request is successful, **Successful*	- is returned. Otherwise, an error message is returned.
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
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSqlLogRecordsResponseBodyData struct {
	// The end time of the query. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1608888296000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the task is complete. Valid values:
	//
	// - **0**: The task is in progress.
	//
	// - **1**: The task is complete.
	//
	// > If this parameter is **0*	- and the **JobId*	- parameter is returned, the current request is an asynchronous request and you cannot obtain the returned results. You must use the value of **JobId*	- to initiate another request. Set the **Filters*	- parameter to the value of **JobId**. Example: `Filters=[{"Key": "JobId", "Value": "******"}]`.
	//
	// example:
	//
	// 1
	Finish *string `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// The details of the SQL logs.
	Items *DescribeSqlLogRecordsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The asynchronous task ID.
	//
	// example:
	//
	// MzI4NTZfUUlOR0RBT19DTTlfTlUyMF9NWVNRTF9PREJTX0xWU18zMjg1Nl9teXNxbF9XZWQgTWFyIDA2IDE0OjUwOjQ3IENTVCAyMDI0XzBfMzBfRXhlY3V0ZVRpbWVfREVTQ19XZWQgTWFyIDA2IDE0OjM1OjQ3IENTVCAyMDI0Xw==_1709708406465
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The start time of the query. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of entries returned.
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
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.SQLLogRecord != nil {
		for _, item := range s.SQLLogRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord struct {
	// The database account.
	//
	// example:
	//
	// testname
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The affected columns.
	//
	// example:
	//
	// ["col1"]
	AffectColumns *string `json:"AffectColumns,omitempty" xml:"AffectColumns,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 10.0.0.1xx
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The client port.
	//
	// example:
	//
	// 3306
	ClientPort *int64 `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// None
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The connection ID.
	//
	// example:
	//
	// ld-******
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The execution duration. Unit: microseconds (μs).
	//
	// example:
	//
	// 58
	Consume *int64 `json:"Consume,omitempty" xml:"Consume,omitempty"`
	// The CPU execution time. Unit: microseconds (μs).
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
	// The execution time. The time is in UTC. Format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// example:
	//
	// 2023-12-07T02:15:32Z
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The extended information. This parameter is reserved.
	//
	// example:
	//
	// None
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The number of rows fetched by the compute node (CN) in a PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 10
	Frows *int64 `json:"Frows,omitempty" xml:"Frows,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 11.197.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The lock wait time. Unit: milliseconds.
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
	// The execution time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1701886532000
	OriginTime *int64 `json:"OriginTime,omitempty" xml:"OriginTime,omitempty"`
	// The degree of parallelism (DOP) for the PolarDB for MySQL instance.
	//
	// example:
	//
	// 10
	ParallelDegree *string `json:"ParallelDegree,omitempty" xml:"ParallelDegree,omitempty"`
	// The parallel queue time for the PolarDB for MySQL instance. Unit: milliseconds.
	//
	// example:
	//
	// 2
	ParallelQueueTime *string `json:"ParallelQueueTime,omitempty" xml:"ParallelQueueTime,omitempty"`
	// The SQL parameters.
	//
	// example:
	//
	// [1, "das"]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The number of asynchronous physical reads.
	//
	// example:
	//
	// 0
	PhysicAsyncRead *int64 `json:"PhysicAsyncRead,omitempty" xml:"PhysicAsyncRead,omitempty"`
	// The number of physical reads.
	//
	// example:
	//
	// 0
	PhysicRead *int64 `json:"PhysicRead,omitempty" xml:"PhysicRead,omitempty"`
	// The number of synchronous physical reads.
	//
	// example:
	//
	// 0
	PhysicSyncRead *int64 `json:"PhysicSyncRead,omitempty" xml:"PhysicSyncRead,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// MySQL
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of returned rows.
	//
	// example:
	//
	// 0
	ReturnRows *int64 `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	// The row key of the SQL log record.
	//
	// example:
	//
	// 23
	RowKey *string `json:"RowKey,omitempty" xml:"RowKey,omitempty"`
	// The total number of rows updated or returned by the compute node (CN) of a PolarDB-X 2.0 instance.
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
	// The number of requests sent from a compute node (CN) to data nodes (DNs) in a PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 10
	Scnt *int64 `json:"Scnt,omitempty" xml:"Scnt,omitempty"`
	// The SQL ID.
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
	// The execution status. Valid values:
	//
	// - **0**: The execution was successful.
	//
	// - **1**: The execution failed.
	//
	// example:
	//
	// 0
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The name of the table that the SQL statement references.
	//
	// example:
	//
	// das
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The thread ID.
	//
	// example:
	//
	// None
	ThreadId *int64 `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// The trace ID for a PolarDB-X 2.0 instance. This is the ID of the SQL statement that was executed on a data node (DN).
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
	// The number of updated rows.
	//
	// example:
	//
	// 0
	UpdateRows *int64 `json:"UpdateRows,omitempty" xml:"UpdateRows,omitempty"`
	// Indicates whether an In-Memory Column Index (IMCI) is used for the PolarDB for MySQL instance.
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	UseImciEngine *string `json:"UseImciEngine,omitempty" xml:"UseImciEngine,omitempty"`
	// The endpoint that is resolved from the query connection string.
	//
	// example:
	//
	// 100.115.XX.XX
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
	// The number of write operations on an ApsaraDB RDS for SQL Server instance.
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

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetAffectColumns() *string {
	return s.AffectColumns
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetClientPort() *int64 {
	return s.ClientPort
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetCollection() *string {
	return s.Collection
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetConnectionId() *string {
	return s.ConnectionId
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

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetParams() *string {
	return s.Params
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

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetReturnRows() *int64 {
	return s.ReturnRows
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) GetRowKey() *string {
	return s.RowKey
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

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetAffectColumns(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.AffectColumns = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetClientIp(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ClientIp = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetClientPort(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ClientPort = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetCollection(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Collection = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetConnectionId(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ConnectionId = &v
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

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetParams(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Params = &v
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

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetProtocol(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.Protocol = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetReturnRows(v int64) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord) SetRowKey(v string) *DescribeSqlLogRecordsResponseBodyDataItemsSQLLogRecord {
	s.RowKey = &v
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
