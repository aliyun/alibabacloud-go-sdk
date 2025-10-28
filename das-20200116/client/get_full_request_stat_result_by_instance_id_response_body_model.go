// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestStatResultByInstanceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetFullRequestStatResultByInstanceIdResponseBody
	GetCode() *int64
	SetData(v *GetFullRequestStatResultByInstanceIdResponseBodyData) *GetFullRequestStatResultByInstanceIdResponseBody
	GetData() *GetFullRequestStatResultByInstanceIdResponseBodyData
	SetMessage(v string) *GetFullRequestStatResultByInstanceIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFullRequestStatResultByInstanceIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFullRequestStatResultByInstanceIdResponseBody
	GetSuccess() *bool
}

type GetFullRequestStatResultByInstanceIdResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetFullRequestStatResultByInstanceIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7172BECE-588A-5961-8126-C216E16B****
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

func (s GetFullRequestStatResultByInstanceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) GetData() *GetFullRequestStatResultByInstanceIdResponseBodyData {
	return s.Data
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetCode(v int64) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetData(v *GetFullRequestStatResultByInstanceIdResponseBodyData) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.Data = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetMessage(v string) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetRequestId(v string) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) SetSuccess(v bool) *GetFullRequestStatResultByInstanceIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFullRequestStatResultByInstanceIdResponseBodyData struct {
	// Indicates whether the asynchronous request failed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Fail *bool `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// Indicates whether the asynchronous request was complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsFinish *bool `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// The returned full request data.
	Result *GetFullRequestStatResultByInstanceIdResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9CB97BC4-6479-55D0-B9D0-EA925AFE****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	// The state of the asynchronous request. Valid values:
	//
	// 	- **RUNNING**
	//
	// 	- **SUCCESS**
	//
	// 	- **FAIL**
	//
	// example:
	//
	// SUCCESS
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the asynchronous request was sent. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1645668213000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) GetFail() *bool {
	return s.Fail
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) GetResult() *GetFullRequestStatResultByInstanceIdResponseBodyDataResult {
	return s.Result
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetFail(v bool) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.Fail = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetIsFinish(v bool) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetResult(v *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.Result = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetResultId(v string) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetState(v string) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.State = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) SetTimestamp(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFullRequestStatResultByInstanceIdResponseBodyDataResult struct {
	// The full request data.
	List []*GetFullRequestStatResultByInstanceIdResponseBodyDataResultList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) GetList() []*GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	return s.List
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) GetTotal() *int64 {
	return s.Total
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) SetList(v []*GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) *GetFullRequestStatResultByInstanceIdResponseBodyDataResult {
	s.List = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) SetTotal(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResult {
	s.Total = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResult) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFullRequestStatResultByInstanceIdResponseBodyDataResultList struct {
	// The average number of scanned rows.
	//
	// > This parameter is returned only for ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and PolarDB for MySQL databases.
	//
	// example:
	//
	// 10000
	AvgExaminedRows *float64 `json:"AvgExaminedRows,omitempty" xml:"AvgExaminedRows,omitempty"`
	// The average number of rows that are fetched by compute nodes from data nodes on the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 0
	AvgFetchRows *int64 `json:"AvgFetchRows,omitempty" xml:"AvgFetchRows,omitempty"`
	// The average lock wait latency. Unit: seconds.
	//
	// example:
	//
	// 0.00009589874265269765
	AvgLockWaitTime *float64 `json:"AvgLockWaitTime,omitempty" xml:"AvgLockWaitTime,omitempty"`
	// The average number of logical reads.
	//
	// example:
	//
	// 654.4470327860251
	AvgLogicalRead *float64 `json:"AvgLogicalRead,omitempty" xml:"AvgLogicalRead,omitempty"`
	// The average number of physical asynchronous reads.
	//
	// example:
	//
	// 0
	AvgPhysicalAsyncRead *int64 `json:"AvgPhysicalAsyncRead,omitempty" xml:"AvgPhysicalAsyncRead,omitempty"`
	// The average number of physical synchronous reads.
	//
	// example:
	//
	// 0
	AvgPhysicalSyncRead *int64 `json:"AvgPhysicalSyncRead,omitempty" xml:"AvgPhysicalSyncRead,omitempty"`
	// The average number of returned rows.
	//
	// example:
	//
	// 10000
	AvgReturnedRows *float64 `json:"AvgReturnedRows,omitempty" xml:"AvgReturnedRows,omitempty"`
	// The average execution duration.
	//
	// example:
	//
	// 2.499
	AvgRt *float64 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// The average number of SQL statements.
	//
	// example:
	//
	// 10000
	AvgSqlCount *int64 `json:"AvgSqlCount,omitempty" xml:"AvgSqlCount,omitempty"`
	// The average number of updated rows.
	//
	//  > This parameter is returned only for ApsaraDB RDS for MySQL and PolarDB-X 2.0 databases.
	//
	// example:
	//
	// 10000
	AvgUpdatedRows *int64 `json:"AvgUpdatedRows,omitempty" xml:"AvgUpdatedRows,omitempty"`
	// The total number of executions.
	//
	// example:
	//
	// 100000
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The percentage of the total number of executions.
	//
	// example:
	//
	// 0.0586
	CountRate *float64 `json:"CountRate,omitempty" xml:"CountRate,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// dbtest01
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The number of failed executions.
	//
	// example:
	//
	// 1
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// The total number of scanned rows.
	//
	// > This parameter is returned only for ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and PolarDB for MySQL databases.
	//
	// example:
	//
	// 10000
	ExaminedRows *int64 `json:"ExaminedRows,omitempty" xml:"ExaminedRows,omitempty"`
	// The number of rows that are fetched by compute nodes from data nodes on the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 0
	FetchRows *int64 `json:"FetchRows,omitempty" xml:"FetchRows,omitempty"`
	// The IP address of the database instance.
	//
	// example:
	//
	// rm-uf6dyi58dm6****.mysql.rds.aliy****.com
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The lock wait latency. Unit: seconds.
	//
	// example:
	//
	// 1089.4177720290281
	LockWaitTime *float64 `json:"LockWaitTime,omitempty" xml:"LockWaitTime,omitempty"`
	// The number of logical reads.
	//
	// example:
	//
	// 7.434573266E9
	LogicalRead *int64 `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	// The number of physical asynchronous reads.
	//
	// example:
	//
	// 0
	PhysicalAsyncRead *int64 `json:"PhysicalAsyncRead,omitempty" xml:"PhysicalAsyncRead,omitempty"`
	// The number of physical synchronous reads.
	//
	// example:
	//
	// 0
	PhysicalSyncRead *int64 `json:"PhysicalSyncRead,omitempty" xml:"PhysicalSyncRead,omitempty"`
	// The port number that is used to connect to the database instance.
	//
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The SQL template.
	//
	// example:
	//
	// select 	- from dbtest01 where ****
	Psql *string `json:"Psql,omitempty" xml:"Psql,omitempty"`
	// The total number of rows updated or returned by the compute nodes of the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 0
	Rows *int64 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The number of SQL statements that take longer than 1 second to execute.
	//
	// example:
	//
	// 20
	RtGreaterThanOneSecondCount *int64 `json:"RtGreaterThanOneSecondCount,omitempty" xml:"RtGreaterThanOneSecondCount,omitempty"`
	// The execution duration percentage.
	//
	// example:
	//
	// 2.499
	RtRate *float64 `json:"RtRate,omitempty" xml:"RtRate,omitempty"`
	// The number of SQL statements.
	//
	// example:
	//
	// 200
	SqlCount *int64 `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	// The SQL ID.
	//
	// example:
	//
	// d71f82be1eef72bd105128204d2e****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The total number of updated rows.
	//
	// example:
	//
	// 100
	SumUpdatedRows *int64 `json:"SumUpdatedRows,omitempty" xml:"SumUpdatedRows,omitempty"`
	// The names of tables in the database.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-2zentqj1sk4qmolci****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgExaminedRows() *float64 {
	return s.AvgExaminedRows
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgFetchRows() *int64 {
	return s.AvgFetchRows
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgLockWaitTime() *float64 {
	return s.AvgLockWaitTime
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgLogicalRead() *float64 {
	return s.AvgLogicalRead
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgPhysicalAsyncRead() *int64 {
	return s.AvgPhysicalAsyncRead
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgPhysicalSyncRead() *int64 {
	return s.AvgPhysicalSyncRead
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgReturnedRows() *float64 {
	return s.AvgReturnedRows
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgRt() *float64 {
	return s.AvgRt
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgSqlCount() *int64 {
	return s.AvgSqlCount
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetAvgUpdatedRows() *int64 {
	return s.AvgUpdatedRows
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetCount() *int64 {
	return s.Count
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetCountRate() *float64 {
	return s.CountRate
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetDatabase() *string {
	return s.Database
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetExaminedRows() *int64 {
	return s.ExaminedRows
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetFetchRows() *int64 {
	return s.FetchRows
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetIp() *string {
	return s.Ip
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetLockWaitTime() *float64 {
	return s.LockWaitTime
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetLogicalRead() *int64 {
	return s.LogicalRead
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetPhysicalAsyncRead() *int64 {
	return s.PhysicalAsyncRead
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetPhysicalSyncRead() *int64 {
	return s.PhysicalSyncRead
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetPort() *int64 {
	return s.Port
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetPsql() *string {
	return s.Psql
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetRows() *int64 {
	return s.Rows
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetRtGreaterThanOneSecondCount() *int64 {
	return s.RtGreaterThanOneSecondCount
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetRtRate() *float64 {
	return s.RtRate
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetSqlCount() *int64 {
	return s.SqlCount
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetSqlId() *string {
	return s.SqlId
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetSumUpdatedRows() *int64 {
	return s.SumUpdatedRows
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetTables() []*string {
	return s.Tables
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetVersion() *int64 {
	return s.Version
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) GetVpcId() *string {
	return s.VpcId
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgExaminedRows(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgExaminedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgFetchRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgFetchRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgLockWaitTime(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgLockWaitTime = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgLogicalRead(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgLogicalRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgPhysicalAsyncRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgPhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgPhysicalSyncRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgPhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgReturnedRows(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgReturnedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgRt(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgRt = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgSqlCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgSqlCount = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetAvgUpdatedRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.AvgUpdatedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Count = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetCountRate(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.CountRate = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetDatabase(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Database = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetErrorCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.ErrorCount = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetExaminedRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.ExaminedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetFetchRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.FetchRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetIp(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Ip = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetLockWaitTime(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.LockWaitTime = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetLogicalRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.LogicalRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetPhysicalAsyncRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.PhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetPhysicalSyncRead(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.PhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetPort(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Port = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetPsql(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Psql = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Rows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetRtGreaterThanOneSecondCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.RtGreaterThanOneSecondCount = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetRtRate(v float64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.RtRate = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetSqlCount(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.SqlCount = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetSqlId(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.SqlId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetSumUpdatedRows(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.SumUpdatedRows = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetTables(v []*string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Tables = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetVersion(v int64) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.Version = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) SetVpcId(v string) *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList {
	s.VpcId = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponseBodyDataResultList) Validate() error {
	return dara.Validate(s)
}
