// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPfsSqlSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetPfsSqlSampleResponseBody
	GetCode() *int64
	SetData(v []*GetPfsSqlSampleResponseBodyData) *GetPfsSqlSampleResponseBody
	GetData() []*GetPfsSqlSampleResponseBodyData
	SetMessage(v string) *GetPfsSqlSampleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPfsSqlSampleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPfsSqlSampleResponseBody
	GetSuccess() *bool
}

type GetPfsSqlSampleResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The SQL sample data.
	Data []*GetPfsSqlSampleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 9CB97BC4-6479-55D0-B9D0-EA925AFE****
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
}

func (s GetPfsSqlSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPfsSqlSampleResponseBody) GoString() string {
	return s.String()
}

func (s *GetPfsSqlSampleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetPfsSqlSampleResponseBody) GetData() []*GetPfsSqlSampleResponseBodyData {
	return s.Data
}

func (s *GetPfsSqlSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPfsSqlSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPfsSqlSampleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPfsSqlSampleResponseBody) SetCode(v int64) *GetPfsSqlSampleResponseBody {
	s.Code = &v
	return s
}

func (s *GetPfsSqlSampleResponseBody) SetData(v []*GetPfsSqlSampleResponseBodyData) *GetPfsSqlSampleResponseBody {
	s.Data = v
	return s
}

func (s *GetPfsSqlSampleResponseBody) SetMessage(v string) *GetPfsSqlSampleResponseBody {
	s.Message = &v
	return s
}

func (s *GetPfsSqlSampleResponseBody) SetRequestId(v string) *GetPfsSqlSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBody) SetSuccess(v bool) *GetPfsSqlSampleResponseBody {
	s.Success = &v
	return s
}

func (s *GetPfsSqlSampleResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPfsSqlSampleResponseBodyData struct {
	// The number of internal on-disk temporary tables that were created when the SQL statement was executed.
	//
	// example:
	//
	// 0
	CreateTmpDiskTables *int32 `json:"CreateTmpDiskTables,omitempty" xml:"CreateTmpDiskTables,omitempty"`
	// The number of internal temporary tables that were created when the SQL statement was executed.
	//
	// example:
	//
	// 0
	CreateTmpTables *int32 `json:"CreateTmpTables,omitempty" xml:"CreateTmpTables,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testDB
	Db *string `json:"Db,omitempty" xml:"Db,omitempty"`
	// The end ID of the event. By default, the value of this parameter is NULL when the event starts and is changed to the event ID when the event ends.
	//
	// example:
	//
	// 0
	EndEventId *int32 `json:"EndEventId,omitempty" xml:"EndEventId,omitempty"`
	// The number of errors returned for the SQL statement.
	//
	// example:
	//
	// 0
	Errors *int32 `json:"Errors,omitempty" xml:"Errors,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 63735293
	EventId *int32 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// statement/sql/select
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution duration. Unit: millisecond.
	//
	// example:
	//
	// 0.199
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The lock wait duration. Unit: millisecond.
	//
	// example:
	//
	// 0.09
	LockLatency *float64 `json:"LockLatency,omitempty" xml:"LockLatency,omitempty"`
	// The ID of the logical database.
	//
	// example:
	//
	// xxxxx
	LogicId *string `json:"LogicId,omitempty" xml:"LogicId,omitempty"`
	// Indicates whether the server failed to find an index that can be used for the SQL statement. Valid values:
	//
	// 	- **1**: yes.
	//
	// 	- **0**: no.
	//
	// example:
	//
	// 1
	NoGoodIndexUsed *int32 `json:"NoGoodIndexUsed,omitempty" xml:"NoGoodIndexUsed,omitempty"`
	// Indicates whether table scans were performed when indexes were not used. Valid values:
	//
	// 	- **1**: yes.
	//
	// 	- **0**: no.
	//
	// example:
	//
	// 1
	NoIndexUsed *int32 `json:"NoIndexUsed,omitempty" xml:"NoIndexUsed,omitempty"`
	// The node ID.
	//
	// >  This parameter is returned only for ApsaraDB RDS for MySQL Cluster Edition instances or PolarDB for MySQL clusters.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The number of rows affected by the SQL statement.
	//
	// example:
	//
	// 0
	RowsAffected *int32 `json:"RowsAffected,omitempty" xml:"RowsAffected,omitempty"`
	// The number of rows scanned by the SQL statement.
	//
	// example:
	//
	// 2048576
	RowsExamined *int32 `json:"RowsExamined,omitempty" xml:"RowsExamined,omitempty"`
	// The number of rows returned by the SQL statement.
	//
	// example:
	//
	// 0
	RowsSent *int32 `json:"RowsSent,omitempty" xml:"RowsSent,omitempty"`
	// The number of joins that are used to perform table scans without using indexes.
	//
	// > : This parameter is used for the scenario in which indexes are not used in a union query. If the returned value is not 0, check the indexes of tables.
	//
	// example:
	//
	// 0
	SelectFullJoin *int32 `json:"SelectFullJoin,omitempty" xml:"SelectFullJoin,omitempty"`
	// The number of joins that used ranges on referenced tables.
	//
	// example:
	//
	// 0
	SelectFullRangeJoin *int32 `json:"SelectFullRangeJoin,omitempty" xml:"SelectFullRangeJoin,omitempty"`
	// The number of joins that used ranges on the first table.
	//
	// example:
	//
	// 0
	SelectRange *int32 `json:"SelectRange,omitempty" xml:"SelectRange,omitempty"`
	// The number of joins that did not have key values. The keys and values were checked for each row of data.
	//
	// > : This parameter is used for the scenario in which indexes are not used in a union query. If the returned value is not 0, check the indexes of tables.
	//
	// example:
	//
	// 0
	SelectRangeCheck *int32 `json:"SelectRangeCheck,omitempty" xml:"SelectRangeCheck,omitempty"`
	// The number of scans.
	//
	// example:
	//
	// 0
	SelectScan *int32 `json:"SelectScan,omitempty" xml:"SelectScan,omitempty"`
	// The number of merges that the sorting algorithm must perform.
	//
	// example:
	//
	// 0
	SortMergePasses *int32 `json:"SortMergePasses,omitempty" xml:"SortMergePasses,omitempty"`
	// The number of times the data was sorted by using ranges.
	//
	// example:
	//
	// 0
	SortRange *int32 `json:"SortRange,omitempty" xml:"SortRange,omitempty"`
	// The number of sorted rows.
	//
	// example:
	//
	// 0
	SortRows *int32 `json:"SortRows,omitempty" xml:"SortRows,omitempty"`
	// The number of sorts that were performed during table scans.
	//
	// example:
	//
	// 1
	SortScan *int32 `json:"SortScan,omitempty" xml:"SortScan,omitempty"`
	// The sample SQL statement.
	//
	// example:
	//
	// select 	- from xxxx where ****
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The SQL statement ID.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The thread ID.
	//
	// example:
	//
	// 81751940
	ThreadId *int32 `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// The time when the SQL statement was executed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1660100753556
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 196278346919****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The number of warnings returned for the SQL statement.
	//
	// example:
	//
	// 0
	Warnings *int32 `json:"Warnings,omitempty" xml:"Warnings,omitempty"`
}

func (s GetPfsSqlSampleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPfsSqlSampleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPfsSqlSampleResponseBodyData) GetCreateTmpDiskTables() *int32 {
	return s.CreateTmpDiskTables
}

func (s *GetPfsSqlSampleResponseBodyData) GetCreateTmpTables() *int32 {
	return s.CreateTmpTables
}

func (s *GetPfsSqlSampleResponseBodyData) GetDb() *string {
	return s.Db
}

func (s *GetPfsSqlSampleResponseBodyData) GetEndEventId() *int32 {
	return s.EndEventId
}

func (s *GetPfsSqlSampleResponseBodyData) GetErrors() *int32 {
	return s.Errors
}

func (s *GetPfsSqlSampleResponseBodyData) GetEventId() *int32 {
	return s.EventId
}

func (s *GetPfsSqlSampleResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *GetPfsSqlSampleResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPfsSqlSampleResponseBodyData) GetLatency() *float64 {
	return s.Latency
}

func (s *GetPfsSqlSampleResponseBodyData) GetLockLatency() *float64 {
	return s.LockLatency
}

func (s *GetPfsSqlSampleResponseBodyData) GetLogicId() *string {
	return s.LogicId
}

func (s *GetPfsSqlSampleResponseBodyData) GetNoGoodIndexUsed() *int32 {
	return s.NoGoodIndexUsed
}

func (s *GetPfsSqlSampleResponseBodyData) GetNoIndexUsed() *int32 {
	return s.NoIndexUsed
}

func (s *GetPfsSqlSampleResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPfsSqlSampleResponseBodyData) GetRowsAffected() *int32 {
	return s.RowsAffected
}

func (s *GetPfsSqlSampleResponseBodyData) GetRowsExamined() *int32 {
	return s.RowsExamined
}

func (s *GetPfsSqlSampleResponseBodyData) GetRowsSent() *int32 {
	return s.RowsSent
}

func (s *GetPfsSqlSampleResponseBodyData) GetSelectFullJoin() *int32 {
	return s.SelectFullJoin
}

func (s *GetPfsSqlSampleResponseBodyData) GetSelectFullRangeJoin() *int32 {
	return s.SelectFullRangeJoin
}

func (s *GetPfsSqlSampleResponseBodyData) GetSelectRange() *int32 {
	return s.SelectRange
}

func (s *GetPfsSqlSampleResponseBodyData) GetSelectRangeCheck() *int32 {
	return s.SelectRangeCheck
}

func (s *GetPfsSqlSampleResponseBodyData) GetSelectScan() *int32 {
	return s.SelectScan
}

func (s *GetPfsSqlSampleResponseBodyData) GetSortMergePasses() *int32 {
	return s.SortMergePasses
}

func (s *GetPfsSqlSampleResponseBodyData) GetSortRange() *int32 {
	return s.SortRange
}

func (s *GetPfsSqlSampleResponseBodyData) GetSortRows() *int32 {
	return s.SortRows
}

func (s *GetPfsSqlSampleResponseBodyData) GetSortScan() *int32 {
	return s.SortScan
}

func (s *GetPfsSqlSampleResponseBodyData) GetSql() *string {
	return s.Sql
}

func (s *GetPfsSqlSampleResponseBodyData) GetSqlId() *string {
	return s.SqlId
}

func (s *GetPfsSqlSampleResponseBodyData) GetThreadId() *int32 {
	return s.ThreadId
}

func (s *GetPfsSqlSampleResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetPfsSqlSampleResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetPfsSqlSampleResponseBodyData) GetWarnings() *int32 {
	return s.Warnings
}

func (s *GetPfsSqlSampleResponseBodyData) SetCreateTmpDiskTables(v int32) *GetPfsSqlSampleResponseBodyData {
	s.CreateTmpDiskTables = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetCreateTmpTables(v int32) *GetPfsSqlSampleResponseBodyData {
	s.CreateTmpTables = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetDb(v string) *GetPfsSqlSampleResponseBodyData {
	s.Db = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetEndEventId(v int32) *GetPfsSqlSampleResponseBodyData {
	s.EndEventId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetErrors(v int32) *GetPfsSqlSampleResponseBodyData {
	s.Errors = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetEventId(v int32) *GetPfsSqlSampleResponseBodyData {
	s.EventId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetEventName(v string) *GetPfsSqlSampleResponseBodyData {
	s.EventName = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetInstanceId(v string) *GetPfsSqlSampleResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetLatency(v float64) *GetPfsSqlSampleResponseBodyData {
	s.Latency = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetLockLatency(v float64) *GetPfsSqlSampleResponseBodyData {
	s.LockLatency = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetLogicId(v string) *GetPfsSqlSampleResponseBodyData {
	s.LogicId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetNoGoodIndexUsed(v int32) *GetPfsSqlSampleResponseBodyData {
	s.NoGoodIndexUsed = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetNoIndexUsed(v int32) *GetPfsSqlSampleResponseBodyData {
	s.NoIndexUsed = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetNodeId(v string) *GetPfsSqlSampleResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetRowsAffected(v int32) *GetPfsSqlSampleResponseBodyData {
	s.RowsAffected = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetRowsExamined(v int32) *GetPfsSqlSampleResponseBodyData {
	s.RowsExamined = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetRowsSent(v int32) *GetPfsSqlSampleResponseBodyData {
	s.RowsSent = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSelectFullJoin(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SelectFullJoin = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSelectFullRangeJoin(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SelectFullRangeJoin = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSelectRange(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SelectRange = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSelectRangeCheck(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SelectRangeCheck = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSelectScan(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SelectScan = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSortMergePasses(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SortMergePasses = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSortRange(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SortRange = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSortRows(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SortRows = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSortScan(v int32) *GetPfsSqlSampleResponseBodyData {
	s.SortScan = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSql(v string) *GetPfsSqlSampleResponseBodyData {
	s.Sql = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetSqlId(v string) *GetPfsSqlSampleResponseBodyData {
	s.SqlId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetThreadId(v int32) *GetPfsSqlSampleResponseBodyData {
	s.ThreadId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetTimestamp(v int64) *GetPfsSqlSampleResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetUserId(v string) *GetPfsSqlSampleResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) SetWarnings(v int32) *GetPfsSqlSampleResponseBodyData {
	s.Warnings = &v
	return s
}

func (s *GetPfsSqlSampleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
