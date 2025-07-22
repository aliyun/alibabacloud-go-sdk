// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestOriginStatByInstanceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetFullRequestOriginStatByInstanceIdResponseBody
	GetCode() *int64
	SetData(v *GetFullRequestOriginStatByInstanceIdResponseBodyData) *GetFullRequestOriginStatByInstanceIdResponseBody
	GetData() *GetFullRequestOriginStatByInstanceIdResponseBodyData
	SetMessage(v string) *GetFullRequestOriginStatByInstanceIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFullRequestOriginStatByInstanceIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFullRequestOriginStatByInstanceIdResponseBody
	GetSuccess() *bool
}

type GetFullRequestOriginStatByInstanceIdResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetFullRequestOriginStatByInstanceIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetFullRequestOriginStatByInstanceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) GetData() *GetFullRequestOriginStatByInstanceIdResponseBodyData {
	return s.Data
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetCode(v int64) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetData(v *GetFullRequestOriginStatByInstanceIdResponseBodyData) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.Data = v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetMessage(v string) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetRequestId(v string) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) SetSuccess(v bool) *GetFullRequestOriginStatByInstanceIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFullRequestOriginStatByInstanceIdResponseBodyData struct {
	// The details of the full request data.
	List []*GetFullRequestOriginStatByInstanceIdResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetFullRequestOriginStatByInstanceIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyData) GetList() []*GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	return s.List
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyData) SetList(v []*GetFullRequestOriginStatByInstanceIdResponseBodyDataList) *GetFullRequestOriginStatByInstanceIdResponseBodyData {
	s.List = v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyData) SetTotal(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetFullRequestOriginStatByInstanceIdResponseBodyDataList struct {
	// The average number of scanned rows.
	//
	// > This parameter is returned only for ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and PolarDB for MySQL databases.
	//
	// example:
	//
	// 10000
	AvgExaminedRows *float64 `json:"AvgExaminedRows,omitempty" xml:"AvgExaminedRows,omitempty"`
	// The average number of rows that are fetched from data nodes by compute nodes on the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 0
	AvgFetchRows *int64 `json:"AvgFetchRows,omitempty" xml:"AvgFetchRows,omitempty"`
	// The average lock wait duration. Unit: seconds.
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
	AvgPhysicalSyncRead *float64 `json:"AvgPhysicalSyncRead,omitempty" xml:"AvgPhysicalSyncRead,omitempty"`
	// The average number of returned rows.
	//
	// example:
	//
	// 10000
	AvgReturnedRows *float64 `json:"AvgReturnedRows,omitempty" xml:"AvgReturnedRows,omitempty"`
	// The average number of rows.
	//
	// example:
	//
	// 0
	AvgRows *int64 `json:"AvgRows,omitempty" xml:"AvgRows,omitempty"`
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
	// > This parameter is returned only for ApsaraDB RDS for MySQL and PolarDB-X 2.0 databases.
	//
	// example:
	//
	// 10000
	AvgUpdatedRows *float64 `json:"AvgUpdatedRows,omitempty" xml:"AvgUpdatedRows,omitempty"`
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
	// The number of rows that are fetched from data nodes by compute nodes on the PolarDB-X 2.0 instance.
	//
	// example:
	//
	// 200
	FetchRows *int64 `json:"FetchRows,omitempty" xml:"FetchRows,omitempty"`
	// The network address of the database instance.
	//
	// example:
	//
	// rm-uf6dyi58dm6****.mysql.rds.aliy****.com
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP address of the client that executes the SQL statement.
	//
	// example:
	//
	// 172.26.6****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The lock wait duration. Unit: seconds.
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
	// The IP address of the client that executes the SQL statement.
	//
	// example:
	//
	// 172.26.6****
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
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
	// 2
	RtGreaterThanOneSecondCount *int64 `json:"RtGreaterThanOneSecondCount,omitempty" xml:"RtGreaterThanOneSecondCount,omitempty"`
	// The execution duration percentage.
	//
	// example:
	//
	// 0.1384
	RtRate *float64 `json:"RtRate,omitempty" xml:"RtRate,omitempty"`
	// The number of SQL statements.
	//
	// example:
	//
	// 200
	SqlCount *int64 `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	// The total number of updated rows.
	//
	// example:
	//
	// 200
	SumUpdatedRows *int64 `json:"SumUpdatedRows,omitempty" xml:"SumUpdatedRows,omitempty"`
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

func (s GetFullRequestOriginStatByInstanceIdResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgExaminedRows() *float64 {
	return s.AvgExaminedRows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgFetchRows() *int64 {
	return s.AvgFetchRows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgLockWaitTime() *float64 {
	return s.AvgLockWaitTime
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgLogicalRead() *float64 {
	return s.AvgLogicalRead
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgPhysicalAsyncRead() *int64 {
	return s.AvgPhysicalAsyncRead
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgPhysicalSyncRead() *float64 {
	return s.AvgPhysicalSyncRead
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgReturnedRows() *float64 {
	return s.AvgReturnedRows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgRows() *int64 {
	return s.AvgRows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgRt() *float64 {
	return s.AvgRt
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgSqlCount() *int64 {
	return s.AvgSqlCount
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetAvgUpdatedRows() *float64 {
	return s.AvgUpdatedRows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetCount() *int64 {
	return s.Count
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetCountRate() *float64 {
	return s.CountRate
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetDatabase() *string {
	return s.Database
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetExaminedRows() *int64 {
	return s.ExaminedRows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetFetchRows() *int64 {
	return s.FetchRows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetIp() *string {
	return s.Ip
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetKey() *string {
	return s.Key
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetLockWaitTime() *float64 {
	return s.LockWaitTime
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetLogicalRead() *int64 {
	return s.LogicalRead
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetOriginHost() *string {
	return s.OriginHost
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetPhysicalAsyncRead() *int64 {
	return s.PhysicalAsyncRead
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetPhysicalSyncRead() *int64 {
	return s.PhysicalSyncRead
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetPort() *int64 {
	return s.Port
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetRows() *int64 {
	return s.Rows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetRtGreaterThanOneSecondCount() *int64 {
	return s.RtGreaterThanOneSecondCount
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetRtRate() *float64 {
	return s.RtRate
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetSqlCount() *int64 {
	return s.SqlCount
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetSumUpdatedRows() *int64 {
	return s.SumUpdatedRows
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetVersion() *int64 {
	return s.Version
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) GetVpcId() *string {
	return s.VpcId
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgExaminedRows(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgExaminedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgFetchRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgFetchRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgLockWaitTime(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgLockWaitTime = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgLogicalRead(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgLogicalRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgPhysicalAsyncRead(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgPhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgPhysicalSyncRead(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgPhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgReturnedRows(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgReturnedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgRt(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgRt = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgSqlCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgSqlCount = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetAvgUpdatedRows(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.AvgUpdatedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetCountRate(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.CountRate = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetDatabase(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Database = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetErrorCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.ErrorCount = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetExaminedRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.ExaminedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetFetchRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.FetchRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetIp(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetKey(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Key = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetLockWaitTime(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.LockWaitTime = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetLogicalRead(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.LogicalRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetOriginHost(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.OriginHost = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetPhysicalAsyncRead(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.PhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetPhysicalSyncRead(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.PhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetPort(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Port = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Rows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetRtGreaterThanOneSecondCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.RtGreaterThanOneSecondCount = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetRtRate(v float64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.RtRate = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetSqlCount(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.SqlCount = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetSumUpdatedRows(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.SumUpdatedRows = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetVersion(v int64) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.Version = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) SetVpcId(v string) *GetFullRequestOriginStatByInstanceIdResponseBodyDataList {
	s.VpcId = &v
	return s
}

func (s *GetFullRequestOriginStatByInstanceIdResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
