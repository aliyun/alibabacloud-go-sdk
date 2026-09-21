// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlInsightStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSqlInsightStatisticResponseBody
	GetCode() *string
	SetData(v *DescribeSqlInsightStatisticResponseBodyData) *DescribeSqlInsightStatisticResponseBody
	GetData() *DescribeSqlInsightStatisticResponseBodyData
	SetMessage(v string) *DescribeSqlInsightStatisticResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSqlInsightStatisticResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSqlInsightStatisticResponseBody
	GetSuccess() *string
}

type DescribeSqlInsightStatisticResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The envelope for asynchronous query results. The first call returns **ResultId*	- and **State**. Poll with the exact same request parameters until **State*	- is **SUCCESS**, then retrieve the statistical details from **List**.
	Data *DescribeSqlInsightStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed. Check the **Code*	- and **Message*	- fields to determine the cause.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSqlInsightStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlInsightStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlInsightStatisticResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSqlInsightStatisticResponseBody) GetData() *DescribeSqlInsightStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeSqlInsightStatisticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSqlInsightStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlInsightStatisticResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSqlInsightStatisticResponseBody) SetCode(v string) *DescribeSqlInsightStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBody) SetData(v *DescribeSqlInsightStatisticResponseBodyData) *DescribeSqlInsightStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBody) SetMessage(v string) *DescribeSqlInsightStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBody) SetRequestId(v string) *DescribeSqlInsightStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBody) SetSuccess(v string) *DescribeSqlInsightStatisticResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSqlInsightStatisticResponseBodyData struct {
	// The SQL Explorer statistical query results.
	//
	// > Returned only when **State*	- is **SUCCESS**.
	Data *DescribeSqlInsightStatisticResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code of the asynchronous query failure.
	//
	// > This field is returned only when the query fails.
	//
	// example:
	//
	// -10200020
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the asynchronous query has completed. Valid values:
	//
	// - **true**: **State*	- is **SUCCESS*	- or **FAIL**.
	//
	// - **false**: **State*	- is **RUNNING**.
	IsFinish *bool `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// The error description of the asynchronous query failure.
	//
	// > This field is returned only when the query fails.
	//
	// example:
	//
	// startTime must be in 30 days and the interval must be within 7 day
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The hash identifier of the request parameters.
	//
	// > This operation does not return this field. Use **ResultId*	- to identify the asynchronous query.
	//
	// example:
	//
	// 507044db6c4eadfa2dab9b084e80****
	RequestKey *string `json:"RequestKey,omitempty" xml:"RequestKey,omitempty"`
	// The asynchronous query result ID, in the format of an async_ prefix followed by a hash value computed from all business parameters of the request.
	//
	// > Repeated calls with the same parameters return the same query result. Therefore, when polling, you must use exactly the same request parameters as the initial call. Any change in parameters generates a different **ResultId*	- and triggers a new query.
	//
	// example:
	//
	// async__507044db6c4eadfa2dab9b084e80****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	// The current status of the asynchronous query. Valid values:
	//
	// - **RUNNING**: The query is in progress. Continue polling.
	//
	// - **SUCCESS**: The query succeeded. The **Data*	- field contains data only in this state.
	//
	// - **FAIL**: The query failed.
	//
	// > When the query fails, the operation directly returns an error code and error message instead of a normal response body with **State*	- set to **FAIL**.
	//
	// example:
	//
	// SUCCESS
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the asynchronous query was submitted. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// > When the query fails, this value indicates the time when the failure occurred.
	//
	// example:
	//
	// 1718600000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeSqlInsightStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlInsightStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlInsightStatisticResponseBodyData) GetData() *DescribeSqlInsightStatisticResponseBodyDataData {
	return s.Data
}

func (s *DescribeSqlInsightStatisticResponseBodyData) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DescribeSqlInsightStatisticResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *DescribeSqlInsightStatisticResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DescribeSqlInsightStatisticResponseBodyData) GetRequestKey() *string {
	return s.RequestKey
}

func (s *DescribeSqlInsightStatisticResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *DescribeSqlInsightStatisticResponseBodyData) GetState() *string {
	return s.State
}

func (s *DescribeSqlInsightStatisticResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSqlInsightStatisticResponseBodyData) SetData(v *DescribeSqlInsightStatisticResponseBodyDataData) *DescribeSqlInsightStatisticResponseBodyData {
	s.Data = v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyData) SetErrorCode(v int32) *DescribeSqlInsightStatisticResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyData) SetIsFinish(v bool) *DescribeSqlInsightStatisticResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyData) SetMessage(v string) *DescribeSqlInsightStatisticResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyData) SetRequestKey(v string) *DescribeSqlInsightStatisticResponseBodyData {
	s.RequestKey = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyData) SetResultId(v string) *DescribeSqlInsightStatisticResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyData) SetState(v string) *DescribeSqlInsightStatisticResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyData) SetTimestamp(v int64) *DescribeSqlInsightStatisticResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyData) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSqlInsightStatisticResponseBodyDataData struct {
	// The extended information.
	//
	// > This field is not returned by this operation.
	//
	// example:
	//
	// {}
	Extra interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The list of SQL Explorer statistical results. Each element is a statistical entry under an aggregation dimension.
	List []*DescribeSqlInsightStatisticResponseBodyDataDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The current page number, corresponding to the request parameter **PageNo**.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page, corresponding to the request parameter **PageSize**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of statistical entries that match the query conditions. You can use this value for pagination calculation.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeSqlInsightStatisticResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlInsightStatisticResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) GetExtra() interface{} {
	return s.Extra
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) GetList() []*DescribeSqlInsightStatisticResponseBodyDataDataList {
	return s.List
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) SetExtra(v interface{}) *DescribeSqlInsightStatisticResponseBodyDataData {
	s.Extra = v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) SetList(v []*DescribeSqlInsightStatisticResponseBodyDataDataList) *DescribeSqlInsightStatisticResponseBodyDataData {
	s.List = v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) SetPageNo(v int64) *DescribeSqlInsightStatisticResponseBodyDataData {
	s.PageNo = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) SetPageSize(v int64) *DescribeSqlInsightStatisticResponseBodyDataData {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) SetTotal(v int64) *DescribeSqlInsightStatisticResponseBodyDataData {
	s.Total = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataData) Validate() error {
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

type DescribeSqlInsightStatisticResponseBodyDataDataList struct {
	// The number of affected rows for a single SQL statement. For **SELECT*	- statements, this indicates the number of scanned rows. For **DML*	- statements, this indicates the number of affected rows.
	//
	// > Returned only for Lindorm instances.
	//
	// example:
	//
	// 30
	AffectRows *int64 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// The value of the aggregation dimension for this statistical entry, which varies based on the **Type*	- request parameter. Valid values:
	//
	// - When aggregated by SQL template: the SQL template ID, which is the same as **SqlId**.
	//
	// - When **Type*	- is set to **FullRequestOrigin**: the access source address.
	//
	// - When **Type*	- is set to **FullRequestUser**: the database username.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	AggKey *string `json:"AggKey,omitempty" xml:"AggKey,omitempty"`
	// The average number of affected rows.
	//
	// > Returned only for Lindorm instances. The value is null for other database engines.
	//
	// example:
	//
	// 30.2
	AvgAffectRows *float64 `json:"AvgAffectRows,omitempty" xml:"AvgAffectRows,omitempty"`
	// The average CPU time consumed by SQL execution, in microseconds.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 12.5
	AvgCpuTime *float64 `json:"AvgCpuTime,omitempty" xml:"AvgCpuTime,omitempty"`
	// The average number of rows fetched by the PolarDB-X compute node from data nodes.
	//
	// > This metric is exclusive to PolarDB-X compute nodes. The value is null or 0 for other database engines.
	//
	// example:
	//
	// 10
	AvgFrows *float64 `json:"AvgFrows,omitempty" xml:"AvgFrows,omitempty"`
	// The average lock wait time per execution, in milliseconds.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 0.00009589874265269765
	AvgLockWaitTime *float64 `json:"AvgLockWaitTime,omitempty" xml:"AvgLockWaitTime,omitempty"`
	// The average number of logical reads per execution.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 654.4470327860251
	AvgLogicalRead *float64 `json:"AvgLogicalRead,omitempty" xml:"AvgLogicalRead,omitempty"`
	// The average number of physical asynchronous reads per execution.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 0
	AvgPhysicalAsyncRead *float64 `json:"AvgPhysicalAsyncRead,omitempty" xml:"AvgPhysicalAsyncRead,omitempty"`
	// The average number of physical reads.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 654.4
	AvgPhysicalRead *float64 `json:"AvgPhysicalRead,omitempty" xml:"AvgPhysicalRead,omitempty"`
	// The average number of physical synchronous reads per execution.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 0
	AvgPhysicalSyncRead *float64 `json:"AvgPhysicalSyncRead,omitempty" xml:"AvgPhysicalSyncRead,omitempty"`
	// The average number of updated rows and returned rows for the PolarDB-X compute node.
	//
	// > This metric is exclusive to PolarDB-X compute nodes. The value is null or 0 for other database engines.
	//
	// example:
	//
	// 10
	AvgRows *float64 `json:"AvgRows,omitempty" xml:"AvgRows,omitempty"`
	// The average number of rows scanned per execution.
	//
	// example:
	//
	// 53421.0
	AvgRowsExamined *float64 `json:"AvgRowsExamined,omitempty" xml:"AvgRowsExamined,omitempty"`
	// The average number of rows returned per execution.
	//
	// example:
	//
	// 14
	AvgRowsReturned *float64 `json:"AvgRowsReturned,omitempty" xml:"AvgRowsReturned,omitempty"`
	// The average number of rows updated per execution.
	//
	// example:
	//
	// 30.2
	AvgRowsUpdated *float64 `json:"AvgRowsUpdated,omitempty" xml:"AvgRowsUpdated,omitempty"`
	// The average execution time per execution, in milliseconds.
	//
	// example:
	//
	// 2.499
	AvgRt *float64 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// The average number of scanned rows.
	//
	// > This field is not returned by this operation. Use **AvgRowsExamined*	- for the average number of scanned rows.
	//
	// example:
	//
	// 53421.0
	AvgScanRows *float64 `json:"AvgScanRows,omitempty" xml:"AvgScanRows,omitempty"`
	// The average number of requests sent by the PolarDB-X compute node to data nodes.
	//
	// > This metric is exclusive to PolarDB-X compute nodes. The value is null or 0 for other database engines.
	//
	// example:
	//
	// 10
	AvgScnt *float64 `json:"AvgScnt,omitempty" xml:"AvgScnt,omitempty"`
	// The average number of logical writes.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 10
	AvgWrites *float64 `json:"AvgWrites,omitempty" xml:"AvgWrites,omitempty"`
	// The total number of executions of the SQL template within the statistical interval.
	//
	// example:
	//
	// 127
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ratio of the number of executions of this statistical entry to the total number of executions of all SQL statements on the instance. The value ranges from 0 to 1.
	//
	// example:
	//
	// 0.0586
	CountRate *float64 `json:"CountRate,omitempty" xml:"CountRate,omitempty"`
	// The total CPU time consumed by SQL execution, in microseconds.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 100
	CpuTime *int64 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	// The name of the database where the SQL statement is executed.
	//
	// example:
	//
	// dbtest01
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The error code returned by SQL execution.
	//
	// > The error code is a detail of a single SQL statement. This operation returns template-level aggregated statistics and does not return this field. Use **ErrorCount*	- for error information.
	//
	// example:
	//
	// 1146
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The number of execution errors for the SQL template within the statistical interval.
	//
	// example:
	//
	// 1
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// The time when the SQL template first appeared.
	//
	// > This field is not returned by this operation.
	//
	// example:
	//
	// 1659308149000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// The total number of rows fetched by the PolarDB-X compute node from data nodes.
	//
	// > This metric is exclusive to PolarDB-X compute nodes. The value is null or 0 for other database engines.
	//
	// example:
	//
	// 10
	Frows *int64 `json:"Frows,omitempty" xml:"Frows,omitempty"`
	// The hash value of the SQL template, returned together with the SQL template.
	//
	// > This value is generated by the PolarDB-X compute node kernel. The value is empty for non-PolarDB-X compute node instances.
	//
	// example:
	//
	// 2e8147b5ca2dfc640dfd5e43d96a****
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The endpoint of the instance to which the statistical data belongs.
	//
	// > Whether this field is returned depends on the aggregated storage link of the instance. The value is null for some links.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****.mysql.rds.aliyuncs.com
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The total lock wait time, in milliseconds.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 1089.4177720290281
	LockWaitTime *float64 `json:"LockWaitTime,omitempty" xml:"LockWaitTime,omitempty"`
	// The total number of logical reads.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 165848
	LogicalRead *float64 `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	// The maximum CPU time in a single execution, in microseconds.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 30
	MaxCpuTime *int64 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	// The maximum number of logical reads in a single execution.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 3186
	MaxLogicalRead *int64 `json:"MaxLogicalRead,omitempty" xml:"MaxLogicalRead,omitempty"`
	// The maximum number of physical reads in a single execution.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 3186
	MaxPhysicalRead *int64 `json:"MaxPhysicalRead,omitempty" xml:"MaxPhysicalRead,omitempty"`
	// The maximum number of rows scanned in a single execution.
	//
	// > This field is not returned by this operation. Use **RowsExamined*	- and **AvgRowsExamined*	- for scanned row counts.
	//
	// example:
	//
	// 318613
	MaxRowsExamined *int64 `json:"MaxRowsExamined,omitempty" xml:"MaxRowsExamined,omitempty"`
	// The maximum number of rows returned in a single execution.
	//
	// example:
	//
	// 20
	MaxRowsReturned *int64 `json:"MaxRowsReturned,omitempty" xml:"MaxRowsReturned,omitempty"`
	// The maximum execution time in a single execution, in milliseconds.
	//
	// example:
	//
	// 12.499
	MaxRt *float64 `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	// The maximum number of logical writes in a single execution.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 12
	MaxWrites *int64 `json:"MaxWrites,omitempty" xml:"MaxWrites,omitempty"`
	// The minimum CPU time in a single execution, in microseconds.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 5
	MinCpuTime *int64 `json:"MinCpuTime,omitempty" xml:"MinCpuTime,omitempty"`
	// The minimum number of logical reads in a single execution.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 12
	MinLogicalRead *int64 `json:"MinLogicalRead,omitempty" xml:"MinLogicalRead,omitempty"`
	// The minimum number of physical reads in a single execution.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 12
	MinPhysicalRead *int64 `json:"MinPhysicalRead,omitempty" xml:"MinPhysicalRead,omitempty"`
	// The minimum number of rows returned in a single execution.
	//
	// example:
	//
	// 1
	MinRowsReturned *int64 `json:"MinRowsReturned,omitempty" xml:"MinRowsReturned,omitempty"`
	// The minimum execution time in a single execution, in milliseconds.
	//
	// example:
	//
	// 0.409789
	MinRt *float64 `json:"MinRt,omitempty" xml:"MinRt,omitempty"`
	// The minimum number of logical writes in a single execution.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 8
	MinWrites *int64 `json:"MinWrites,omitempty" xml:"MinWrites,omitempty"`
	// The display alias configured for the access source address.
	//
	// > Returned only when aggregated by access source (when **Type*	- is set to **FullRequestOrigin**). The value is null in other scenarios.
	//
	// example:
	//
	// order-1
	OriginAlias *string `json:"OriginAlias,omitempty" xml:"OriginAlias,omitempty"`
	// The source address of the client that initiated the SQL statement.
	//
	// > When **Type*	- is set to **FullRequestOrigin**, this field serves as the aggregation dimension for the statistical entry.
	//
	// example:
	//
	// 172.26.XX.XXX
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The parameter content of the SQL sample.
	//
	// > This operation returns template-level aggregated statistics and does not return this field.
	//
	// example:
	//
	// [1, "das"]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The total number of physical asynchronous reads.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 0
	PhysicalAsyncRead *float64 `json:"PhysicalAsyncRead,omitempty" xml:"PhysicalAsyncRead,omitempty"`
	// The total number of physical reads.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other database engines.
	//
	// example:
	//
	// 165848
	PhysicalRead *int64 `json:"PhysicalRead,omitempty" xml:"PhysicalRead,omitempty"`
	// The total number of physical synchronous reads.
	//
	// > Data is available only when **Version*	- is set to **1**.
	//
	// example:
	//
	// 0
	PhysicalSyncRead *float64 `json:"PhysicalSyncRead,omitempty" xml:"PhysicalSyncRead,omitempty"`
	// The port of the instance to which the statistical data belongs.
	//
	// > Whether this field is returned depends on the aggregated storage link of the instance. The value is null for some links.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The parameterized SQL template text, which is the statement with constants in the SQL replaced by placeholders.
	//
	// example:
	//
	// select 	- from t_order where id = ?
	Psql *string `json:"Psql,omitempty" xml:"Psql,omitempty"`
	// The total number of updated rows and returned rows for the PolarDB-X compute node.
	//
	// > This metric is exclusive to PolarDB-X compute nodes. The value is null or 0 for other engines.
	//
	// example:
	//
	// 10
	Rows *int64 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The total number of rows examined by the SQL template within the statistical interval.
	//
	// example:
	//
	// 2048576
	RowsExamined *int64 `json:"RowsExamined,omitempty" xml:"RowsExamined,omitempty"`
	// The total number of rows returned by the SQL template within the statistical interval.
	//
	// example:
	//
	// 14
	RowsReturned *int64 `json:"RowsReturned,omitempty" xml:"RowsReturned,omitempty"`
	// The total execution duration of the SQL template within the statistical interval. Unit: milliseconds.
	//
	// > For PolarDB-X compute nodes (where **Role*	- is **polarx_cn**) with kernel versions earlier than 5.4.13, this value is converted from microseconds to milliseconds.
	//
	// example:
	//
	// 0.409789
	Rt *float64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// The number of times the execution duration exceeds 1 second.
	//
	// > Whether this field is returned depends on the aggregation storage link of the instance. The value is null for certain links.
	//
	// example:
	//
	// 20
	RtGreaterThanOneSecondCount *int64 `json:"RtGreaterThanOneSecondCount,omitempty" xml:"RtGreaterThanOneSecondCount,omitempty"`
	// The ratio of the total execution duration of this entry to the total execution duration of all SQL statements on the instance. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.1384
	RtRate *float64 `json:"RtRate,omitempty" xml:"RtRate,omitempty"`
	// The type identifier of the sample data.
	//
	// > This operation returns aggregated statistics and does not return this field.
	//
	// example:
	//
	// sql
	SampleType *string `json:"SampleType,omitempty" xml:"SampleType,omitempty"`
	// The number of rows scanned by a single SQL statement.
	//
	// > This operation returns template-level aggregated statistics and does not return this field. Use the aggregated metrics **RowsExamined*	- and **AvgRowsExamined*	- instead.
	//
	// example:
	//
	// 29
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The total number of requests sent from the PolarDB-X compute node to data nodes.
	//
	// > This metric is exclusive to PolarDB-X compute nodes. The value is null or 0 for other engines.
	//
	// example:
	//
	// 10
	Scnt *int64 `json:"Scnt,omitempty" xml:"Scnt,omitempty"`
	// The original SQL text.
	//
	// > The statistical results return the SQL template (**Psql**) and do not return this field.
	//
	// example:
	//
	// select 	- from t_order where id = 1
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The SQL template ID that uniquely identifies a type of parameterized SQL statement. Multiple executions under the same template are aggregated into a single statistical entry. You can use this ID to correlate the same type of SQL across multi-dimensional queries.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The SQL text with parameter values uniformly processed, used in sample data scenarios.
	//
	// > This operation does not return this field.
	//
	// example:
	//
	// select 	- from t_order where id = ?
	SqlNew *string `json:"SqlNew,omitempty" xml:"SqlNew,omitempty"`
	// The SQL text feature value, used in SQL analysis scenarios.
	//
	// > This operation does not return this field.
	//
	// example:
	//
	// select_from_t_order
	SqlTextFeature *string `json:"SqlTextFeature,omitempty" xml:"SqlTextFeature,omitempty"`
	// The SQL type. Valid values:
	//
	// - **select**
	//
	// - **insert**
	//
	// - **update**
	//
	// - **delete**
	//
	// - **other**
	//
	// example:
	//
	// select
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The total number of rows updated by the SQL template within the statistical interval.
	//
	// example:
	//
	// 3810
	SumRowsUpdated *float64 `json:"SumRowsUpdated,omitempty" xml:"SumRowsUpdated,omitempty"`
	// The list of table names involved in the SQL statement.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The database thread ID that executed the SQL statement.
	//
	// > The thread ID is a detail of a single SQL statement. This operation returns template-level aggregated statistics and does not return this field.
	//
	// example:
	//
	// 57472578
	ThreadId *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// The execution duration ratio.
	//
	// > This operation returns the execution duration ratio through **RtRate*	- and does not return this field.
	//
	// example:
	//
	// 0.1384
	TimeRate *float64 `json:"TimeRate,omitempty" xml:"TimeRate,omitempty"`
	// The data timestamp. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// > The statistical results are aggregated at the SQL template level and do not return this field.
	//
	// example:
	//
	// 1718600000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The total number of affected rows.
	//
	// > This field is returned only for Lindorm instances. The value is null for other engines.
	//
	// example:
	//
	// 3810
	TotalAffectRows *int64 `json:"TotalAffectRows,omitempty" xml:"TotalAffectRows,omitempty"`
	// The total SQL execution duration.
	//
	// > This operation does not return this field. Use **Rt*	- for the total execution duration.
	//
	// example:
	//
	// 310
	TotalRt *int64 `json:"TotalRt,omitempty" xml:"TotalRt,omitempty"`
	// The total number of rows scanned.
	//
	// > This operation does not return this field. Use **RowsExamined*	- for the total number of rows scanned.
	//
	// example:
	//
	// 2048576
	TotalScanRows *int64 `json:"TotalScanRows,omitempty" xml:"TotalScanRows,omitempty"`
	// The execution count trend sequence of the SQL template, divided into time slices within the query time window.
	//
	// > This field is returned only when the request parameter **DoFillTrend*	- is set to **true*	- and the trend padding capability is enabled for the instance. The time slice interval is automatically determined by the query span. Time slices with no data may be padded with zeros.
	Trend []*DescribeSqlInsightStatisticResponseBodyDataDataListTrend `json:"Trend,omitempty" xml:"Trend,omitempty" type:"Repeated"`
	// The number of rows updated by a single SQL statement.
	//
	// > This operation returns template-level aggregated statistics and does not return this field. Use the aggregated metrics **SumRowsUpdated*	- and **AvgRowsUpdated*	- instead.
	//
	// example:
	//
	// 30
	UpdateRows *int64 `json:"UpdateRows,omitempty" xml:"UpdateRows,omitempty"`
	// The database username that executed the SQL statement.
	//
	// > When **Type*	- is set to **FullRequestUser**, this field serves as the aggregation dimension for the statistical entry.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The SQL Explorer data collection link version. The value **1*	- is returned when the instance collects logical read or lock wait data. Otherwise, the value **0*	- is returned. Valid values:
	//
	// - **0**: V0 basic collection link.
	//
	// - **1**: V1 collection link, which additionally collects four metrics (**LockWaitTime**, **LogicalRead**, **PhysicalSyncRead**, and **PhysicalAsyncRead**) on top of V0.
	//
	// > When the value is **0**, the extended metrics contain no data.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The VPC ID of the instance to which the statistical data belongs.
	//
	// > Whether this field is returned depends on the aggregation storage link of the instance. The value is null for certain links.
	//
	// example:
	//
	// vpc-2zentqj1sk4qmolci****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The total number of logical writes.
	//
	// > This metric is exclusive to SQL Server instances. The value is null for other engines.
	//
	// example:
	//
	// 10
	Writes *int64 `json:"Writes,omitempty" xml:"Writes,omitempty"`
}

func (s DescribeSqlInsightStatisticResponseBodyDataDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlInsightStatisticResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAggKey() *string {
	return s.AggKey
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgAffectRows() *float64 {
	return s.AvgAffectRows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgCpuTime() *float64 {
	return s.AvgCpuTime
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgFrows() *float64 {
	return s.AvgFrows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgLockWaitTime() *float64 {
	return s.AvgLockWaitTime
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgLogicalRead() *float64 {
	return s.AvgLogicalRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgPhysicalAsyncRead() *float64 {
	return s.AvgPhysicalAsyncRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgPhysicalRead() *float64 {
	return s.AvgPhysicalRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgPhysicalSyncRead() *float64 {
	return s.AvgPhysicalSyncRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgRows() *float64 {
	return s.AvgRows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgRowsExamined() *float64 {
	return s.AvgRowsExamined
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgRowsReturned() *float64 {
	return s.AvgRowsReturned
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgRowsUpdated() *float64 {
	return s.AvgRowsUpdated
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgRt() *float64 {
	return s.AvgRt
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgScanRows() *float64 {
	return s.AvgScanRows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgScnt() *float64 {
	return s.AvgScnt
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetAvgWrites() *float64 {
	return s.AvgWrites
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetCountRate() *float64 {
	return s.CountRate
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetCpuTime() *int64 {
	return s.CpuTime
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetDatabase() *string {
	return s.Database
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetFrows() *int64 {
	return s.Frows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetHash() *string {
	return s.Hash
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetIp() *string {
	return s.Ip
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetLockWaitTime() *float64 {
	return s.LockWaitTime
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetLogicalRead() *float64 {
	return s.LogicalRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMaxCpuTime() *int64 {
	return s.MaxCpuTime
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMaxLogicalRead() *int64 {
	return s.MaxLogicalRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMaxPhysicalRead() *int64 {
	return s.MaxPhysicalRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMaxRowsExamined() *int64 {
	return s.MaxRowsExamined
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMaxRowsReturned() *int64 {
	return s.MaxRowsReturned
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMaxRt() *float64 {
	return s.MaxRt
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMaxWrites() *int64 {
	return s.MaxWrites
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMinCpuTime() *int64 {
	return s.MinCpuTime
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMinLogicalRead() *int64 {
	return s.MinLogicalRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMinPhysicalRead() *int64 {
	return s.MinPhysicalRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMinRowsReturned() *int64 {
	return s.MinRowsReturned
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMinRt() *float64 {
	return s.MinRt
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetMinWrites() *int64 {
	return s.MinWrites
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetOriginAlias() *string {
	return s.OriginAlias
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetOriginHost() *string {
	return s.OriginHost
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetParams() *string {
	return s.Params
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetPhysicalAsyncRead() *float64 {
	return s.PhysicalAsyncRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetPhysicalRead() *int64 {
	return s.PhysicalRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetPhysicalSyncRead() *float64 {
	return s.PhysicalSyncRead
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetPsql() *string {
	return s.Psql
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetRows() *int64 {
	return s.Rows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetRowsExamined() *int64 {
	return s.RowsExamined
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetRowsReturned() *int64 {
	return s.RowsReturned
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetRt() *float64 {
	return s.Rt
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetRtGreaterThanOneSecondCount() *int64 {
	return s.RtGreaterThanOneSecondCount
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetRtRate() *float64 {
	return s.RtRate
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetSampleType() *string {
	return s.SampleType
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetScnt() *int64 {
	return s.Scnt
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetSql() *string {
	return s.Sql
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetSqlId() *string {
	return s.SqlId
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetSqlNew() *string {
	return s.SqlNew
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetSqlTextFeature() *string {
	return s.SqlTextFeature
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetSumRowsUpdated() *float64 {
	return s.SumRowsUpdated
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetTables() []*string {
	return s.Tables
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetThreadId() *string {
	return s.ThreadId
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetTimeRate() *float64 {
	return s.TimeRate
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetTotalAffectRows() *int64 {
	return s.TotalAffectRows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetTotalRt() *int64 {
	return s.TotalRt
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetTotalScanRows() *int64 {
	return s.TotalScanRows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetTrend() []*DescribeSqlInsightStatisticResponseBodyDataDataListTrend {
	return s.Trend
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetUpdateRows() *int64 {
	return s.UpdateRows
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetUser() *string {
	return s.User
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) GetWrites() *int64 {
	return s.Writes
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAffectRows(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AffectRows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAggKey(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AggKey = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgAffectRows(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgAffectRows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgCpuTime(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgCpuTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgFrows(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgFrows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgLockWaitTime(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgLockWaitTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgLogicalRead(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgLogicalRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgPhysicalAsyncRead(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgPhysicalAsyncRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgPhysicalRead(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgPhysicalRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgPhysicalSyncRead(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgPhysicalSyncRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgRows(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgRows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgRowsExamined(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgRowsExamined = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgRowsReturned(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgRowsReturned = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgRowsUpdated(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgRowsUpdated = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgRt(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgRt = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgScanRows(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgScanRows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgScnt(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgScnt = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetAvgWrites(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.AvgWrites = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetCount(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Count = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetCountRate(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.CountRate = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetCpuTime(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.CpuTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetDatabase(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Database = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetErrorCode(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetErrorCount(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.ErrorCount = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetFirstTime(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.FirstTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetFrows(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Frows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetHash(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Hash = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetIp(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Ip = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetLockWaitTime(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.LockWaitTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetLogicalRead(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMaxCpuTime(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMaxLogicalRead(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MaxLogicalRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMaxPhysicalRead(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MaxPhysicalRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMaxRowsExamined(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MaxRowsExamined = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMaxRowsReturned(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MaxRowsReturned = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMaxRt(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MaxRt = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMaxWrites(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MaxWrites = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMinCpuTime(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MinCpuTime = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMinLogicalRead(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MinLogicalRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMinPhysicalRead(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MinPhysicalRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMinRowsReturned(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MinRowsReturned = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMinRt(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MinRt = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetMinWrites(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.MinWrites = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetOriginAlias(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.OriginAlias = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetOriginHost(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.OriginHost = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetParams(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Params = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetPhysicalAsyncRead(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.PhysicalAsyncRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetPhysicalRead(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.PhysicalRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetPhysicalSyncRead(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.PhysicalSyncRead = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetPort(v int32) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Port = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetPsql(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Psql = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetRows(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Rows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetRowsExamined(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.RowsExamined = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetRowsReturned(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.RowsReturned = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetRt(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Rt = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetRtGreaterThanOneSecondCount(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.RtGreaterThanOneSecondCount = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetRtRate(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.RtRate = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetSampleType(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.SampleType = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetScanRows(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.ScanRows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetScnt(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Scnt = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetSql(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Sql = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetSqlId(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.SqlId = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetSqlNew(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.SqlNew = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetSqlTextFeature(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.SqlTextFeature = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetSqlType(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.SqlType = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetSumRowsUpdated(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.SumRowsUpdated = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetTables(v []*string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Tables = v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetThreadId(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.ThreadId = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetTimeRate(v float64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.TimeRate = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetTimestamp(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Timestamp = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetTotalAffectRows(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.TotalAffectRows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetTotalRt(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.TotalRt = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetTotalScanRows(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.TotalScanRows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetTrend(v []*DescribeSqlInsightStatisticResponseBodyDataDataListTrend) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Trend = v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetUpdateRows(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.UpdateRows = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetUser(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.User = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetVersion(v int32) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Version = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetVpcId(v string) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.VpcId = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) SetWrites(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataList {
	s.Writes = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataList) Validate() error {
	if s.Trend != nil {
		for _, item := range s.Trend {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSqlInsightStatisticResponseBodyDataDataListTrend struct {
	// The timestamp of the trend data point. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1718000000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The number of SQL executions within the time slice.
	//
	// example:
	//
	// 12
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSqlInsightStatisticResponseBodyDataDataListTrend) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlInsightStatisticResponseBodyDataDataListTrend) GoString() string {
	return s.String()
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataListTrend) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataListTrend) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataListTrend) SetTimestamp(v int64) *DescribeSqlInsightStatisticResponseBodyDataDataListTrend {
	s.Timestamp = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataListTrend) SetValue(v interface{}) *DescribeSqlInsightStatisticResponseBodyDataDataListTrend {
	s.Value = v
	return s
}

func (s *DescribeSqlInsightStatisticResponseBodyDataDataListTrend) Validate() error {
	return dara.Validate(s)
}
