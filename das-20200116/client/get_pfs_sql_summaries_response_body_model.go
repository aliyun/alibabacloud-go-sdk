// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPfsSqlSummariesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetPfsSqlSummariesResponseBody
	GetCode() *int64
	SetData(v *GetPfsSqlSummariesResponseBodyData) *GetPfsSqlSummariesResponseBody
	GetData() *GetPfsSqlSummariesResponseBodyData
	SetMessage(v string) *GetPfsSqlSummariesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPfsSqlSummariesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPfsSqlSummariesResponseBody
	GetSuccess() *bool
}

type GetPfsSqlSummariesResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetPfsSqlSummariesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54F3DBAE-9420-511A-9C29-265E8C04****
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPfsSqlSummariesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPfsSqlSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *GetPfsSqlSummariesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetPfsSqlSummariesResponseBody) GetData() *GetPfsSqlSummariesResponseBodyData {
	return s.Data
}

func (s *GetPfsSqlSummariesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPfsSqlSummariesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPfsSqlSummariesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPfsSqlSummariesResponseBody) SetCode(v int64) *GetPfsSqlSummariesResponseBody {
	s.Code = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBody) SetData(v *GetPfsSqlSummariesResponseBodyData) *GetPfsSqlSummariesResponseBody {
	s.Data = v
	return s
}

func (s *GetPfsSqlSummariesResponseBody) SetMessage(v string) *GetPfsSqlSummariesResponseBody {
	s.Message = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBody) SetRequestId(v string) *GetPfsSqlSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBody) SetSuccess(v bool) *GetPfsSqlSummariesResponseBody {
	s.Success = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPfsSqlSummariesResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The details of the data returned.
	List []*GetPfsSqlSummariesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 264
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetPfsSqlSummariesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPfsSqlSummariesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPfsSqlSummariesResponseBodyData) GetExtra() interface{} {
	return s.Extra
}

func (s *GetPfsSqlSummariesResponseBodyData) GetList() []*GetPfsSqlSummariesResponseBodyDataList {
	return s.List
}

func (s *GetPfsSqlSummariesResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetPfsSqlSummariesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetPfsSqlSummariesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetPfsSqlSummariesResponseBodyData) SetExtra(v interface{}) *GetPfsSqlSummariesResponseBodyData {
	s.Extra = v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyData) SetList(v []*GetPfsSqlSummariesResponseBodyDataList) *GetPfsSqlSummariesResponseBodyData {
	s.List = v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyData) SetPageNo(v int64) *GetPfsSqlSummariesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyData) SetPageSize(v int64) *GetPfsSqlSummariesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyData) SetTotal(v int64) *GetPfsSqlSummariesResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetPfsSqlSummariesResponseBodyDataList struct {
	// The average execution latency. Unit: millisecond.
	//
	// example:
	//
	// 0.1717
	AvgLatency *float64 `json:"AvgLatency,omitempty" xml:"AvgLatency,omitempty"`
	// The total number of executions.
	//
	// example:
	//
	// 100000
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The percentage of the number of executions.
	//
	// example:
	//
	// 0.0586
	CountRate *float64 `json:"CountRate,omitempty" xml:"CountRate,omitempty"`
	// The ratio of the CPU execution duration to the total execution duration of the SQL statement.
	//
	// example:
	//
	// 0
	CpuRate *float64 `json:"CpuRate,omitempty" xml:"CpuRate,omitempty"`
	// The CPU execution duration. Unit: millisecond.
	//
	// example:
	//
	// 0
	CpuTime *float64 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	// The data read duration. Unit: millisecond.
	//
	// example:
	//
	// 0
	DataReadTime *float64 `json:"DataReadTime,omitempty" xml:"DataReadTime,omitempty"`
	// The number of nodes from which data can be read.
	//
	// example:
	//
	// 0
	DataReads *int32 `json:"DataReads,omitempty" xml:"DataReads,omitempty"`
	// The data write duration. Unit: millisecond.
	//
	// example:
	//
	// 0
	DataWriteTime *float64 `json:"DataWriteTime,omitempty" xml:"DataWriteTime,omitempty"`
	// The number of nodes to which data can be written.
	//
	// example:
	//
	// 0
	DataWrites *int32 `json:"DataWrites,omitempty" xml:"DataWrites,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testDB
	Db *string `json:"Db,omitempty" xml:"Db,omitempty"`
	// The execution duration. Unit: millisecond.
	//
	// example:
	//
	// 0
	ElapsedTime *float64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The number of errors.
	//
	// example:
	//
	// 0
	ErrCount *int64 `json:"ErrCount,omitempty" xml:"ErrCount,omitempty"`
	// The time when the SQL statement was executed for the first time. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1659308149000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// Indicates whether full table scan was enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FullScan *bool `json:"FullScan,omitempty" xml:"FullScan,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 26186357
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the SQL statement was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1661306520000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The average lock wait latency. Unit: millisecond.
	//
	// example:
	//
	// 0
	LockLatencyAvg *float64 `json:"LockLatencyAvg,omitempty" xml:"LockLatencyAvg,omitempty"`
	// The logical database ID.
	//
	// example:
	//
	// 58275984
	LogicId *int64 `json:"LogicId,omitempty" xml:"LogicId,omitempty"`
	// The number of logical nodes.
	//
	// example:
	//
	// 0
	LogicReads *int64 `json:"LogicReads,omitempty" xml:"LogicReads,omitempty"`
	// The maximum execution latency. Unit: millisecond.
	//
	// example:
	//
	// 36.233
	MaxLatency *float64 `json:"MaxLatency,omitempty" xml:"MaxLatency,omitempty"`
	// The number of mutex spins.
	//
	// example:
	//
	// 1
	MutexSpins *int32 `json:"MutexSpins,omitempty" xml:"MutexSpins,omitempty"`
	// The number of mutex waits.
	//
	// example:
	//
	// 1
	MutexWaits *int32 `json:"MutexWaits,omitempty" xml:"MutexWaits,omitempty"`
	// The node ID.
	//
	// >  This parameter is returned only if the database instance is an ApsaraDB RDS for MySQL Cluster Edition instance or a PolarDB for MySQL cluster.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The number of physical asynchronous nodes.
	//
	// example:
	//
	// 0
	PhysicalAsyncReads *int64 `json:"PhysicalAsyncReads,omitempty" xml:"PhysicalAsyncReads,omitempty"`
	// The number of physical nodes.
	//
	// example:
	//
	// 0
	PhysicalReads *int64 `json:"PhysicalReads,omitempty" xml:"PhysicalReads,omitempty"`
	// The SQL template.
	//
	// example:
	//
	// select ?
	Psql *string `json:"Psql,omitempty" xml:"Psql,omitempty"`
	// The number of redo nodes.
	//
	// example:
	//
	// 0
	RedoWrites *int64 `json:"RedoWrites,omitempty" xml:"RedoWrites,omitempty"`
	// The number of rows that are affected by the SQL statement.
	//
	// example:
	//
	// 0
	RowsAffected *int64 `json:"RowsAffected,omitempty" xml:"RowsAffected,omitempty"`
	// The average number of rows affected by the SQL statement.
	//
	// example:
	//
	// 0
	RowsAffectedAvg *float64 `json:"RowsAffectedAvg,omitempty" xml:"RowsAffectedAvg,omitempty"`
	// The total number of scanned rows.
	//
	// example:
	//
	// 100
	RowsExamined *int64 `json:"RowsExamined,omitempty" xml:"RowsExamined,omitempty"`
	// The average number of scanned rows.
	//
	// example:
	//
	// 0
	RowsExaminedAvg *float64 `json:"RowsExaminedAvg,omitempty" xml:"RowsExaminedAvg,omitempty"`
	// The average number of returned rows.
	//
	// example:
	//
	// 0
	RowsSendAvg *float64 `json:"RowsSendAvg,omitempty" xml:"RowsSendAvg,omitempty"`
	// The number of rows returned by the SQL statement.
	//
	// example:
	//
	// 0
	RowsSent *int64 `json:"RowsSent,omitempty" xml:"RowsSent,omitempty"`
	// The average number of rows returned for the SQL statement.
	//
	// example:
	//
	// 0.52
	RowsSentAvg *float64 `json:"RowsSentAvg,omitempty" xml:"RowsSentAvg,omitempty"`
	// The number of sorted rows.
	//
	// example:
	//
	// 0
	RowsSorted *int64 `json:"RowsSorted,omitempty" xml:"RowsSorted,omitempty"`
	// The execution duration percentage.
	//
	// example:
	//
	// 0.1384
	RtRate *float64 `json:"RtRate,omitempty" xml:"RtRate,omitempty"`
	// Indicates whether read/write splitting was enabled. Valid values:
	//
	// 	- **0:*	- Read/write splitting was disabled.
	//
	// 	- **1:*	- Read/write splitting was enabled.
	//
	// example:
	//
	// 0
	RwlockOsWaits *int32 `json:"RwlockOsWaits,omitempty" xml:"RwlockOsWaits,omitempty"`
	// The read/write splitting parameters.
	//
	// example:
	//
	// 0
	RwlockSpinRounds *int32 `json:"RwlockSpinRounds,omitempty" xml:"RwlockSpinRounds,omitempty"`
	// Indices whether multi-index scanning was enabled. Valid values:
	//
	// 	- **0:*	- Multi-index scanning was disabled.
	//
	// 	- **1:*	- Multi-index scanning was enabled.
	//
	// example:
	//
	// 0
	RwlockSpinWaits *int32 `json:"RwlockSpinWaits,omitempty" xml:"RwlockSpinWaits,omitempty"`
	// The average number of joins that performed table scans without using indexes.
	//
	// > If the value of this parameter is not 0, check the table indexes.
	//
	// example:
	//
	// 0
	SelectFullJoinAvg *float64 `json:"SelectFullJoinAvg,omitempty" xml:"SelectFullJoinAvg,omitempty"`
	// The average number of joins that selected a range.
	//
	// example:
	//
	// 0
	SelectFullRangeJoinAvg *float64 `json:"SelectFullRangeJoinAvg,omitempty" xml:"SelectFullRangeJoinAvg,omitempty"`
	// The average selected range.
	//
	// example:
	//
	// 0
	SelectRangeAvg *float64 `json:"SelectRangeAvg,omitempty" xml:"SelectRangeAvg,omitempty"`
	// The average number of scanned rows.
	//
	// example:
	//
	// 0
	SelectScanAvg *float64 `json:"SelectScanAvg,omitempty" xml:"SelectScanAvg,omitempty"`
	// The semi-synchronous replication latency. Unit: millisecond.
	//
	// example:
	//
	// 0.12
	SemisyncDelayTime *float64 `json:"SemisyncDelayTime,omitempty" xml:"SemisyncDelayTime,omitempty"`
	// The amount of time consumed for locking the server. Unit: millisecond.
	//
	// example:
	//
	// 0
	ServerLockTime *float64 `json:"ServerLockTime,omitempty" xml:"ServerLockTime,omitempty"`
	// The number of merges that the sorting algorithm must perform.
	//
	// example:
	//
	// 0
	SortMergePasses *int64 `json:"SortMergePasses,omitempty" xml:"SortMergePasses,omitempty"`
	// The average number of sorts that were performed by using a range.
	//
	// example:
	//
	// 0
	SortRangeAvg *float64 `json:"SortRangeAvg,omitempty" xml:"SortRangeAvg,omitempty"`
	// The average number of sorted rows.
	//
	// example:
	//
	// 0
	SortRowsAvg *float64 `json:"SortRowsAvg,omitempty" xml:"SortRowsAvg,omitempty"`
	// The average number of sorts that were performed during table scans.
	//
	// example:
	//
	// 0
	SortScanAvg *float64 `json:"SortScanAvg,omitempty" xml:"SortScanAvg,omitempty"`
	// The SQL template ID.
	//
	// example:
	//
	// 2e8147b5ca2dfc640dfd5e43d96a****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
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
	// The names of tables in the database.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	TimerWaitAvg *float64 `json:"TimerWaitAvg,omitempty" xml:"TimerWaitAvg,omitempty"`
	// The data timestamp. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1643040000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The number of on-disk temporary tables.
	//
	// example:
	//
	// 0
	TmpDiskTables *int64 `json:"TmpDiskTables,omitempty" xml:"TmpDiskTables,omitempty"`
	// The average number of on-disk temporary tables.
	//
	// example:
	//
	// 0
	TmpDiskTablesAvg *float64 `json:"TmpDiskTablesAvg,omitempty" xml:"TmpDiskTablesAvg,omitempty"`
	// The number of temporary tables.
	//
	// example:
	//
	// 0
	TmpTables *int64 `json:"TmpTables,omitempty" xml:"TmpTables,omitempty"`
	// The average number of temporary tables.
	//
	// example:
	//
	// 0
	TmpTablesAvg *float64 `json:"TmpTablesAvg,omitempty" xml:"TmpTablesAvg,omitempty"`
	// The execution latency. Unit: millisecond.
	//
	// example:
	//
	// 60913.256
	TotalLatency *float64 `json:"TotalLatency,omitempty" xml:"TotalLatency,omitempty"`
	// The amount of time consumed for locking the storage transaction. Unit: millisecond.
	//
	// example:
	//
	// 0
	TransactionLockTime *float64 `json:"TransactionLockTime,omitempty" xml:"TransactionLockTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 196278346919****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The number of warnings.
	//
	// example:
	//
	// 0
	WarnCount *int64 `json:"WarnCount,omitempty" xml:"WarnCount,omitempty"`
}

func (s GetPfsSqlSummariesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetPfsSqlSummariesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetAvgLatency() *float64 {
	return s.AvgLatency
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetCount() *int64 {
	return s.Count
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetCountRate() *float64 {
	return s.CountRate
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetCpuRate() *float64 {
	return s.CpuRate
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetCpuTime() *float64 {
	return s.CpuTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetDataReadTime() *float64 {
	return s.DataReadTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetDataReads() *int32 {
	return s.DataReads
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetDataWriteTime() *float64 {
	return s.DataWriteTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetDataWrites() *int32 {
	return s.DataWrites
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetDb() *string {
	return s.Db
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetElapsedTime() *float64 {
	return s.ElapsedTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetErrCount() *int64 {
	return s.ErrCount
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetFullScan() *bool {
	return s.FullScan
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetLastTime() *int64 {
	return s.LastTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetLockLatencyAvg() *float64 {
	return s.LockLatencyAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetLogicId() *int64 {
	return s.LogicId
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetLogicReads() *int64 {
	return s.LogicReads
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetMaxLatency() *float64 {
	return s.MaxLatency
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetMutexSpins() *int32 {
	return s.MutexSpins
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetMutexWaits() *int32 {
	return s.MutexWaits
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetPhysicalAsyncReads() *int64 {
	return s.PhysicalAsyncReads
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetPhysicalReads() *int64 {
	return s.PhysicalReads
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetPsql() *string {
	return s.Psql
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRedoWrites() *int64 {
	return s.RedoWrites
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRowsAffected() *int64 {
	return s.RowsAffected
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRowsAffectedAvg() *float64 {
	return s.RowsAffectedAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRowsExamined() *int64 {
	return s.RowsExamined
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRowsExaminedAvg() *float64 {
	return s.RowsExaminedAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRowsSendAvg() *float64 {
	return s.RowsSendAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRowsSent() *int64 {
	return s.RowsSent
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRowsSentAvg() *float64 {
	return s.RowsSentAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRowsSorted() *int64 {
	return s.RowsSorted
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRtRate() *float64 {
	return s.RtRate
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRwlockOsWaits() *int32 {
	return s.RwlockOsWaits
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRwlockSpinRounds() *int32 {
	return s.RwlockSpinRounds
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetRwlockSpinWaits() *int32 {
	return s.RwlockSpinWaits
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSelectFullJoinAvg() *float64 {
	return s.SelectFullJoinAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSelectFullRangeJoinAvg() *float64 {
	return s.SelectFullRangeJoinAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSelectRangeAvg() *float64 {
	return s.SelectRangeAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSelectScanAvg() *float64 {
	return s.SelectScanAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSemisyncDelayTime() *float64 {
	return s.SemisyncDelayTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetServerLockTime() *float64 {
	return s.ServerLockTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSortMergePasses() *int64 {
	return s.SortMergePasses
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSortRangeAvg() *float64 {
	return s.SortRangeAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSortRowsAvg() *float64 {
	return s.SortRowsAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSortScanAvg() *float64 {
	return s.SortScanAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSqlId() *string {
	return s.SqlId
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetSqlType() *string {
	return s.SqlType
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTables() []*string {
	return s.Tables
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTimerWaitAvg() *float64 {
	return s.TimerWaitAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTmpDiskTables() *int64 {
	return s.TmpDiskTables
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTmpDiskTablesAvg() *float64 {
	return s.TmpDiskTablesAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTmpTables() *int64 {
	return s.TmpTables
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTmpTablesAvg() *float64 {
	return s.TmpTablesAvg
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTotalLatency() *float64 {
	return s.TotalLatency
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetTransactionLockTime() *float64 {
	return s.TransactionLockTime
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetUserId() *string {
	return s.UserId
}

func (s *GetPfsSqlSummariesResponseBodyDataList) GetWarnCount() *int64 {
	return s.WarnCount
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetAvgLatency(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.AvgLatency = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetCount(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetCountRate(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.CountRate = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetCpuRate(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.CpuRate = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetCpuTime(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.CpuTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetDataReadTime(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.DataReadTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetDataReads(v int32) *GetPfsSqlSummariesResponseBodyDataList {
	s.DataReads = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetDataWriteTime(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.DataWriteTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetDataWrites(v int32) *GetPfsSqlSummariesResponseBodyDataList {
	s.DataWrites = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetDb(v string) *GetPfsSqlSummariesResponseBodyDataList {
	s.Db = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetElapsedTime(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.ElapsedTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetErrCount(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.ErrCount = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetFirstTime(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.FirstTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetFullScan(v bool) *GetPfsSqlSummariesResponseBodyDataList {
	s.FullScan = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetId(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetInstanceId(v string) *GetPfsSqlSummariesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetLastTime(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.LastTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetLockLatencyAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.LockLatencyAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetLogicId(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.LogicId = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetLogicReads(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.LogicReads = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetMaxLatency(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.MaxLatency = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetMutexSpins(v int32) *GetPfsSqlSummariesResponseBodyDataList {
	s.MutexSpins = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetMutexWaits(v int32) *GetPfsSqlSummariesResponseBodyDataList {
	s.MutexWaits = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetNodeId(v string) *GetPfsSqlSummariesResponseBodyDataList {
	s.NodeId = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetPhysicalAsyncReads(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.PhysicalAsyncReads = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetPhysicalReads(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.PhysicalReads = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetPsql(v string) *GetPfsSqlSummariesResponseBodyDataList {
	s.Psql = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRedoWrites(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RedoWrites = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRowsAffected(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RowsAffected = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRowsAffectedAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RowsAffectedAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRowsExamined(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RowsExamined = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRowsExaminedAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RowsExaminedAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRowsSendAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RowsSendAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRowsSent(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RowsSent = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRowsSentAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RowsSentAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRowsSorted(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RowsSorted = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRtRate(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.RtRate = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRwlockOsWaits(v int32) *GetPfsSqlSummariesResponseBodyDataList {
	s.RwlockOsWaits = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRwlockSpinRounds(v int32) *GetPfsSqlSummariesResponseBodyDataList {
	s.RwlockSpinRounds = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetRwlockSpinWaits(v int32) *GetPfsSqlSummariesResponseBodyDataList {
	s.RwlockSpinWaits = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSelectFullJoinAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SelectFullJoinAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSelectFullRangeJoinAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SelectFullRangeJoinAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSelectRangeAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SelectRangeAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSelectScanAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SelectScanAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSemisyncDelayTime(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SemisyncDelayTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetServerLockTime(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.ServerLockTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSortMergePasses(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SortMergePasses = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSortRangeAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SortRangeAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSortRowsAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SortRowsAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSortScanAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.SortScanAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSqlId(v string) *GetPfsSqlSummariesResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetSqlType(v string) *GetPfsSqlSummariesResponseBodyDataList {
	s.SqlType = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTables(v []*string) *GetPfsSqlSummariesResponseBodyDataList {
	s.Tables = v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTimerWaitAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.TimerWaitAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTimestamp(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.Timestamp = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTmpDiskTables(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.TmpDiskTables = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTmpDiskTablesAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.TmpDiskTablesAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTmpTables(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.TmpTables = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTmpTablesAvg(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.TmpTablesAvg = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTotalLatency(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.TotalLatency = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetTransactionLockTime(v float64) *GetPfsSqlSummariesResponseBodyDataList {
	s.TransactionLockTime = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetUserId(v string) *GetPfsSqlSummariesResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) SetWarnCount(v int64) *GetPfsSqlSummariesResponseBodyDataList {
	s.WarnCount = &v
	return s
}

func (s *GetPfsSqlSummariesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
