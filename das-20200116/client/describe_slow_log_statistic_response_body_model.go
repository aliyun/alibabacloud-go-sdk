// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSlowLogStatisticResponseBody
	GetCode() *string
	SetData(v *DescribeSlowLogStatisticResponseBodyData) *DescribeSlowLogStatisticResponseBody
	GetData() *DescribeSlowLogStatisticResponseBodyData
	SetMessage(v string) *DescribeSlowLogStatisticResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSlowLogStatisticResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSlowLogStatisticResponseBody
	GetSuccess() *string
}

type DescribeSlowLogStatisticResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// AsyncResult<DBLogRecords<SlowLogStat>>
	Data *DescribeSlowLogStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 52D540CF-C517-1F57-BB42-9035F96******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSlowLogStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSlowLogStatisticResponseBody) GetData() *DescribeSlowLogStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeSlowLogStatisticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSlowLogStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogStatisticResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSlowLogStatisticResponseBody) SetCode(v string) *DescribeSlowLogStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBody) SetData(v *DescribeSlowLogStatisticResponseBodyData) *DescribeSlowLogStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBody) SetMessage(v string) *DescribeSlowLogStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBody) SetRequestId(v string) *DescribeSlowLogStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBody) SetSuccess(v string) *DescribeSlowLogStatisticResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogStatisticResponseBodyData struct {
	Data *DescribeSlowLogStatisticResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 10910
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// true
	IsFinish *bool `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 123456789
	RequestKey *string `json:"RequestKey,omitempty" xml:"RequestKey,omitempty"`
	// example:
	//
	// async__665ee69612f1627c7fd9f3c85075****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	// example:
	//
	// SUCCESS
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1735104224250
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeSlowLogStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponseBodyData) GetData() *DescribeSlowLogStatisticResponseBodyDataData {
	return s.Data
}

func (s *DescribeSlowLogStatisticResponseBodyData) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DescribeSlowLogStatisticResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *DescribeSlowLogStatisticResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DescribeSlowLogStatisticResponseBodyData) GetRequestKey() *string {
	return s.RequestKey
}

func (s *DescribeSlowLogStatisticResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *DescribeSlowLogStatisticResponseBodyData) GetState() *string {
	return s.State
}

func (s *DescribeSlowLogStatisticResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSlowLogStatisticResponseBodyData) SetData(v *DescribeSlowLogStatisticResponseBodyDataData) *DescribeSlowLogStatisticResponseBodyData {
	s.Data = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyData) SetErrorCode(v int32) *DescribeSlowLogStatisticResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyData) SetIsFinish(v bool) *DescribeSlowLogStatisticResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyData) SetMessage(v string) *DescribeSlowLogStatisticResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyData) SetRequestKey(v string) *DescribeSlowLogStatisticResponseBodyData {
	s.RequestKey = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyData) SetResultId(v string) *DescribeSlowLogStatisticResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyData) SetState(v string) *DescribeSlowLogStatisticResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyData) SetTimestamp(v int64) *DescribeSlowLogStatisticResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyData) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogStatisticResponseBodyDataData struct {
	// example:
	//
	// rm-k2ja51w7cnusg5a1x
	DbInstanceId *int64 `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// 0
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 2024-08-08T02:15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	ItemsNumbers *int64                                              `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	Logs         []*DescribeSlowLogStatisticResponseBodyDataDataLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxRecordsPerPage *int32 `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	// example:
	//
	// pi-wz99g5rn7w1x8h0sf
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 1
	PageNumbers *int32 `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// example:
	//
	// 2024-10-08T02:01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 13
	TotalRecords *int64 `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeSlowLogStatisticResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetDbInstanceId() *int64 {
	return s.DbInstanceId
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetItemsNumbers() *int64 {
	return s.ItemsNumbers
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetLogs() []*DescribeSlowLogStatisticResponseBodyDataDataLogs {
	return s.Logs
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetMaxRecordsPerPage() *int32 {
	return s.MaxRecordsPerPage
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetPageNumbers() *int32 {
	return s.PageNumbers
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) GetTotalRecords() *int64 {
	return s.TotalRecords
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetDbInstanceId(v int64) *DescribeSlowLogStatisticResponseBodyDataData {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetDbInstanceName(v string) *DescribeSlowLogStatisticResponseBodyDataData {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetEndTime(v string) *DescribeSlowLogStatisticResponseBodyDataData {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetItemsNumbers(v int64) *DescribeSlowLogStatisticResponseBodyDataData {
	s.ItemsNumbers = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetLogs(v []*DescribeSlowLogStatisticResponseBodyDataDataLogs) *DescribeSlowLogStatisticResponseBodyDataData {
	s.Logs = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetMaxRecordsPerPage(v int32) *DescribeSlowLogStatisticResponseBodyDataData {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetNodeId(v string) *DescribeSlowLogStatisticResponseBodyDataData {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetPageNumbers(v int32) *DescribeSlowLogStatisticResponseBodyDataData {
	s.PageNumbers = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetStartTime(v string) *DescribeSlowLogStatisticResponseBodyDataData {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) SetTotalRecords(v int64) *DescribeSlowLogStatisticResponseBodyDataData {
	s.TotalRecords = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataData) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogStatisticResponseBodyDataDataLogs struct {
	// example:
	//
	// edu_admin
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// -
	AvgCPUTime        *float64 `json:"AvgCPUTime,omitempty" xml:"AvgCPUTime,omitempty"`
	AvgCPUTimeSeconds *float64 `json:"AvgCPUTimeSeconds,omitempty" xml:"AvgCPUTimeSeconds,omitempty"`
	// example:
	//
	// 10000
	AvgDocExamined *float64 `json:"AvgDocExamined,omitempty" xml:"AvgDocExamined,omitempty"`
	AvgFrows       *float64 `json:"AvgFrows,omitempty" xml:"AvgFrows,omitempty"`
	// example:
	//
	// 1000
	AvgIOWrites *float64 `json:"AvgIOWrites,omitempty" xml:"AvgIOWrites,omitempty"`
	// example:
	//
	// 20000
	AvgKeysExamined *float64 `json:"AvgKeysExamined,omitempty" xml:"AvgKeysExamined,omitempty"`
	// example:
	//
	// 1000
	AvgLastRowsCountAffected *float64 `json:"AvgLastRowsCountAffected,omitempty" xml:"AvgLastRowsCountAffected,omitempty"`
	// example:
	//
	// -
	AvgLockTime *float64 `json:"AvgLockTime,omitempty" xml:"AvgLockTime,omitempty"`
	// example:
	//
	// 0.0
	AvgLockTimeSeconds *float64 `json:"AvgLockTimeSeconds,omitempty" xml:"AvgLockTimeSeconds,omitempty"`
	// example:
	//
	// 1000
	AvgLogicalIOReads *float64 `json:"AvgLogicalIOReads,omitempty" xml:"AvgLogicalIOReads,omitempty"`
	// example:
	//
	// 1000
	AvgPhysicalIOReads *float64 `json:"AvgPhysicalIOReads,omitempty" xml:"AvgPhysicalIOReads,omitempty"`
	// example:
	//
	// -
	AvgQueryTime *float64 `json:"AvgQueryTime,omitempty" xml:"AvgQueryTime,omitempty"`
	// example:
	//
	// 6.211
	AvgQueryTimeSeconds *float64 `json:"AvgQueryTimeSeconds,omitempty" xml:"AvgQueryTimeSeconds,omitempty"`
	AvgRequestSize      *float64 `json:"AvgRequestSize,omitempty" xml:"AvgRequestSize,omitempty"`
	AvgResponseSize     *float64 `json:"AvgResponseSize,omitempty" xml:"AvgResponseSize,omitempty"`
	// example:
	//
	// 1
	AvgReturnNum *float64 `json:"AvgReturnNum,omitempty" xml:"AvgReturnNum,omitempty"`
	AvgRows      *float64 `json:"AvgRows,omitempty" xml:"AvgRows,omitempty"`
	// example:
	//
	// 1000
	AvgRowsCountAffected *float64 `json:"AvgRowsCountAffected,omitempty" xml:"AvgRowsCountAffected,omitempty"`
	// example:
	//
	// 53421.0
	AvgRowsExamined *float64 `json:"AvgRowsExamined,omitempty" xml:"AvgRowsExamined,omitempty"`
	// example:
	//
	// 2.0
	AvgRowsSent *float64 `json:"AvgRowsSent,omitempty" xml:"AvgRowsSent,omitempty"`
	AvgRt       *float64 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	AvgScnt     *float64 `json:"AvgScnt,omitempty" xml:"AvgScnt,omitempty"`
	// example:
	//
	// -
	CPUTime        *float64 `json:"CPUTime,omitempty" xml:"CPUTime,omitempty"`
	CPUTimeSeconds *float64 `json:"CPUTimeSeconds,omitempty" xml:"CPUTimeSeconds,omitempty"`
	ClientIp       *string  `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Cmd            *string  `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	// example:
	//
	// SELECT b?.id,b?.t?,b?.id,b?.t? FROM testtb? b? JOIN testtb? b? ON b?.id=b?.id WHERE b?.t? LIKE ? ORDER BY b?.t? DESC
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// example:
	//
	// 12
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0.2034
	CountRate *float64 `json:"CountRate,omitempty" xml:"CountRate,omitempty"`
	// example:
	//
	// member_score
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// work-wechat-api
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	DbId     *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// 0
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 2000000
	DocExamined *int64 `json:"DocExamined,omitempty" xml:"DocExamined,omitempty"`
	// example:
	//
	// 1
	DocsExamined *int64                                                     `json:"DocsExamined,omitempty" xml:"DocsExamined,omitempty"`
	Frows        *int64                                                     `json:"Frows,omitempty" xml:"Frows,omitempty"`
	Histogram    *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram `json:"Histogram,omitempty" xml:"Histogram,omitempty" type:"Struct"`
	// example:
	//
	// 172.23.142.178
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// example:
	//
	// mongo-abc123456
	HostInsId *string `json:"HostInsId,omitempty" xml:"HostInsId,omitempty"`
	// example:
	//
	// 1000
	IOWrites *int64 `json:"IOWrites,omitempty" xml:"IOWrites,omitempty"`
	// example:
	//
	// rm-uf6zix1z2jh1y6fe5
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// example:
	//
	// __system
	InsRole *string `json:"InsRole,omitempty" xml:"InsRole,omitempty"`
	// example:
	//
	// 20000
	KeysExamined *int64 `json:"KeysExamined,omitempty" xml:"KeysExamined,omitempty"`
	// example:
	//
	// 1000
	LastRowsCountAffected *int64 `json:"LastRowsCountAffected,omitempty" xml:"LastRowsCountAffected,omitempty"`
	// example:
	//
	// -
	LockTime *float64 `json:"LockTime,omitempty" xml:"LockTime,omitempty"`
	// example:
	//
	// 0.0
	LockTimeSeconds *float64 `json:"LockTimeSeconds,omitempty" xml:"LockTimeSeconds,omitempty"`
	// example:
	//
	// 1000
	LogicalIOReads *int64 `json:"LogicalIOReads,omitempty" xml:"LogicalIOReads,omitempty"`
	// example:
	//
	// -
	MaxCPUTime        *float64 `json:"MaxCPUTime,omitempty" xml:"MaxCPUTime,omitempty"`
	MaxCPUTimeSeconds *float64 `json:"MaxCPUTimeSeconds,omitempty" xml:"MaxCPUTimeSeconds,omitempty"`
	// example:
	//
	// 1000000
	MaxDocExamined *int64 `json:"MaxDocExamined,omitempty" xml:"MaxDocExamined,omitempty"`
	MaxFrows       *int64 `json:"MaxFrows,omitempty" xml:"MaxFrows,omitempty"`
	// example:
	//
	// 1000
	MaxIOWrites *int64 `json:"MaxIOWrites,omitempty" xml:"MaxIOWrites,omitempty"`
	// example:
	//
	// 2000000
	MaxKeysExamined *int64 `json:"MaxKeysExamined,omitempty" xml:"MaxKeysExamined,omitempty"`
	// example:
	//
	// 1000
	MaxLastRowsCountAffected *int64 `json:"MaxLastRowsCountAffected,omitempty" xml:"MaxLastRowsCountAffected,omitempty"`
	// example:
	//
	// -
	MaxLockTime *float64 `json:"MaxLockTime,omitempty" xml:"MaxLockTime,omitempty"`
	// example:
	//
	// 0.0
	MaxLockTimeSeconds *float64 `json:"MaxLockTimeSeconds,omitempty" xml:"MaxLockTimeSeconds,omitempty"`
	// example:
	//
	// 1000
	MaxLogicalIOReads *int64 `json:"MaxLogicalIOReads,omitempty" xml:"MaxLogicalIOReads,omitempty"`
	// example:
	//
	// 1000
	MaxPhysicalIOReads *int64 `json:"MaxPhysicalIOReads,omitempty" xml:"MaxPhysicalIOReads,omitempty"`
	// example:
	//
	// -
	MaxQueryTime *float64 `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	// example:
	//
	// 14.402
	MaxQueryTimeSeconds *float64 `json:"MaxQueryTimeSeconds,omitempty" xml:"MaxQueryTimeSeconds,omitempty"`
	MaxRequestSize      *float64 `json:"MaxRequestSize,omitempty" xml:"MaxRequestSize,omitempty"`
	MaxResponseSize     *float64 `json:"MaxResponseSize,omitempty" xml:"MaxResponseSize,omitempty"`
	// example:
	//
	// 1
	MaxReturnNum *int64 `json:"MaxReturnNum,omitempty" xml:"MaxReturnNum,omitempty"`
	MaxRows      *int64 `json:"MaxRows,omitempty" xml:"MaxRows,omitempty"`
	// example:
	//
	// 1000
	MaxRowsCountAffected *int64 `json:"MaxRowsCountAffected,omitempty" xml:"MaxRowsCountAffected,omitempty"`
	// example:
	//
	// 318613
	MaxRowsExamined *int64 `json:"MaxRowsExamined,omitempty" xml:"MaxRowsExamined,omitempty"`
	// example:
	//
	// 256
	MaxRowsSent *int64   `json:"MaxRowsSent,omitempty" xml:"MaxRowsSent,omitempty"`
	MaxRt       *float64 `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	MaxScnt     *int64   `json:"MaxScnt,omitempty" xml:"MaxScnt,omitempty"`
	// example:
	//
	// database.collection
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// DLNode
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// Insert
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// example:
	//
	// order-1
	OriginAlias *string `json:"OriginAlias,omitempty" xml:"OriginAlias,omitempty"`
	// example:
	//
	// 1000
	PhysicalIOReads *int64 `json:"PhysicalIOReads,omitempty" xml:"PhysicalIOReads,omitempty"`
	// example:
	//
	// SELECT b?.id,b?.t?,b?.id,b?.t? FROM testtb? b? JOIN testtb? b? ON b?.id=b?.id WHERE b?.t? LIKE ? ORDER BY b?.id DESC
	Psql *string `json:"Psql,omitempty" xml:"Psql,omitempty"`
	// example:
	//
	// sq-1pzcdMwRb
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// example:
	//
	// 2024-12-25T03:00:00Z
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// example:
	//
	// 272.444
	QueryTime *int64 `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// example:
	//
	// 0.1018
	QueryTimeRate *float64 `json:"QueryTimeRate,omitempty" xml:"QueryTimeRate,omitempty"`
	// example:
	//
	// 25.472
	QueryTimeSeconds  *float64 `json:"QueryTimeSeconds,omitempty" xml:"QueryTimeSeconds,omitempty"`
	ReturnItemNumbers *string  `json:"ReturnItemNumbers,omitempty" xml:"ReturnItemNumbers,omitempty"`
	// example:
	//
	// 1
	ReturnNum *int64 `json:"ReturnNum,omitempty" xml:"ReturnNum,omitempty"`
	Rows      *int64 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 1000
	RowsCountAffected *int64 `json:"RowsCountAffected,omitempty" xml:"RowsCountAffected,omitempty"`
	// example:
	//
	// 2444081
	RowsExamined *int64 `json:"RowsExamined,omitempty" xml:"RowsExamined,omitempty"`
	// example:
	//
	// 772
	RowsSent *int64  `json:"RowsSent,omitempty" xml:"RowsSent,omitempty"`
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// SELECT \\"Hello, World!\\" FROM DUAL
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// example:
	//
	// HTTP
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Scnt   *int64  `json:"Scnt,omitempty" xml:"Scnt,omitempty"`
	// SQL ID。
	//
	// example:
	//
	// 2dca88762ec6b3812504ab8a4b******
	SqlId  *string                                                 `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlTag *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag `json:"SqlTag,omitempty" xml:"SqlTag,omitempty" type:"Struct"`
	// example:
	//
	// LOGIN
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// example:
	//
	// r-8vba51c588ba3a94
	SubInstanceId *string `json:"SubInstanceId,omitempty" xml:"SubInstanceId,omitempty"`
	// example:
	//
	// users\\nifconfig\\n
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 1
	ThreadId *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// example:
	//
	// 1708568930
	Timestamp  *int64                                                   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TotalCount *int64                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Trend      []*DescribeSlowLogStatisticResponseBodyDataDataLogsTrend `json:"Trend,omitempty" xml:"Trend,omitempty" type:"Repeated"`
	// example:
	//
	// user-1
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogs) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgCPUTime() *float64 {
	return s.AvgCPUTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgCPUTimeSeconds() *float64 {
	return s.AvgCPUTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgDocExamined() *float64 {
	return s.AvgDocExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgFrows() *float64 {
	return s.AvgFrows
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgIOWrites() *float64 {
	return s.AvgIOWrites
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgKeysExamined() *float64 {
	return s.AvgKeysExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgLastRowsCountAffected() *float64 {
	return s.AvgLastRowsCountAffected
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgLockTime() *float64 {
	return s.AvgLockTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgLockTimeSeconds() *float64 {
	return s.AvgLockTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgLogicalIOReads() *float64 {
	return s.AvgLogicalIOReads
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgPhysicalIOReads() *float64 {
	return s.AvgPhysicalIOReads
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgQueryTime() *float64 {
	return s.AvgQueryTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgQueryTimeSeconds() *float64 {
	return s.AvgQueryTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgRequestSize() *float64 {
	return s.AvgRequestSize
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgResponseSize() *float64 {
	return s.AvgResponseSize
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgReturnNum() *float64 {
	return s.AvgReturnNum
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgRows() *float64 {
	return s.AvgRows
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgRowsCountAffected() *float64 {
	return s.AvgRowsCountAffected
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgRowsExamined() *float64 {
	return s.AvgRowsExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgRowsSent() *float64 {
	return s.AvgRowsSent
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgRt() *float64 {
	return s.AvgRt
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetAvgScnt() *float64 {
	return s.AvgScnt
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetCPUTime() *float64 {
	return s.CPUTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetCPUTimeSeconds() *float64 {
	return s.CPUTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetCmd() *string {
	return s.Cmd
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetCommand() *string {
	return s.Command
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetCount() *int64 {
	return s.Count
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetCountRate() *float64 {
	return s.CountRate
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetDatabase() *string {
	return s.Database
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetDbId() *string {
	return s.DbId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetDocExamined() *int64 {
	return s.DocExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetDocsExamined() *int64 {
	return s.DocsExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetFrows() *int64 {
	return s.Frows
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetHistogram() *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	return s.Histogram
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetHostInsId() *string {
	return s.HostInsId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetIOWrites() *int64 {
	return s.IOWrites
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetInsName() *string {
	return s.InsName
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetInsRole() *string {
	return s.InsRole
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetKeysExamined() *int64 {
	return s.KeysExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetLastRowsCountAffected() *int64 {
	return s.LastRowsCountAffected
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetLockTime() *float64 {
	return s.LockTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetLockTimeSeconds() *float64 {
	return s.LockTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetLogicalIOReads() *int64 {
	return s.LogicalIOReads
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxCPUTime() *float64 {
	return s.MaxCPUTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxCPUTimeSeconds() *float64 {
	return s.MaxCPUTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxDocExamined() *int64 {
	return s.MaxDocExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxFrows() *int64 {
	return s.MaxFrows
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxIOWrites() *int64 {
	return s.MaxIOWrites
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxKeysExamined() *int64 {
	return s.MaxKeysExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxLastRowsCountAffected() *int64 {
	return s.MaxLastRowsCountAffected
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxLockTime() *float64 {
	return s.MaxLockTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxLockTimeSeconds() *float64 {
	return s.MaxLockTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxLogicalIOReads() *int64 {
	return s.MaxLogicalIOReads
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxPhysicalIOReads() *int64 {
	return s.MaxPhysicalIOReads
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxQueryTime() *float64 {
	return s.MaxQueryTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxQueryTimeSeconds() *float64 {
	return s.MaxQueryTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxRequestSize() *float64 {
	return s.MaxRequestSize
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxResponseSize() *float64 {
	return s.MaxResponseSize
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxReturnNum() *int64 {
	return s.MaxReturnNum
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxRows() *int64 {
	return s.MaxRows
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxRowsCountAffected() *int64 {
	return s.MaxRowsCountAffected
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxRowsExamined() *int64 {
	return s.MaxRowsExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxRowsSent() *int64 {
	return s.MaxRowsSent
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxRt() *float64 {
	return s.MaxRt
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetMaxScnt() *int64 {
	return s.MaxScnt
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetOpType() *string {
	return s.OpType
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetOriginAlias() *string {
	return s.OriginAlias
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetPhysicalIOReads() *int64 {
	return s.PhysicalIOReads
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetPsql() *string {
	return s.Psql
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetQueryTime() *int64 {
	return s.QueryTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetQueryTimeRate() *float64 {
	return s.QueryTimeRate
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetQueryTimeSeconds() *float64 {
	return s.QueryTimeSeconds
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetReturnItemNumbers() *string {
	return s.ReturnItemNumbers
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetReturnNum() *int64 {
	return s.ReturnNum
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetRows() *int64 {
	return s.Rows
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetRowsCountAffected() *int64 {
	return s.RowsCountAffected
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetRowsExamined() *int64 {
	return s.RowsExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetRowsSent() *int64 {
	return s.RowsSent
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetScheme() *string {
	return s.Scheme
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetScnt() *int64 {
	return s.Scnt
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetSqlId() *string {
	return s.SqlId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetSqlTag() *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag {
	return s.SqlTag
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetSubInstanceId() *string {
	return s.SubInstanceId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetThreadId() *string {
	return s.ThreadId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetTrend() []*DescribeSlowLogStatisticResponseBodyDataDataLogsTrend {
	return s.Trend
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) GetUser() *string {
	return s.User
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAccountName(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AccountName = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgCPUTime(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgCPUTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgCPUTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgCPUTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgDocExamined(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgDocExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgFrows(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgFrows = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgIOWrites(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgIOWrites = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgKeysExamined(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgKeysExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgLastRowsCountAffected(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgLastRowsCountAffected = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgLockTime(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgLockTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgLockTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgLockTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgLogicalIOReads(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgLogicalIOReads = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgPhysicalIOReads(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgPhysicalIOReads = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgQueryTime(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgQueryTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgQueryTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgQueryTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgRequestSize(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgRequestSize = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgResponseSize(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgResponseSize = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgReturnNum(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgReturnNum = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgRows(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgRows = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgRowsCountAffected(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgRowsCountAffected = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgRowsExamined(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgRowsExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgRowsSent(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgRowsSent = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgRt(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgRt = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetAvgScnt(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.AvgScnt = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetCPUTime(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.CPUTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetCPUTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.CPUTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetClientIp(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.ClientIp = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetCmd(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Cmd = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetCommand(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Command = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetCount(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Count = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetCountRate(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.CountRate = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetDBName(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetDatabase(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Database = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetDbId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.DbId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetDbInstanceName(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetDocExamined(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.DocExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetDocsExamined(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.DocsExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetFrows(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Frows = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetHistogram(v *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Histogram = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetHostAddress(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetHostInsId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.HostInsId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetIOWrites(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.IOWrites = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetInsName(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.InsName = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetInsRole(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.InsRole = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetKeysExamined(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.KeysExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetLastRowsCountAffected(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.LastRowsCountAffected = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetLockTime(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.LockTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetLockTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.LockTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetLogicalIOReads(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.LogicalIOReads = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxCPUTime(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxCPUTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxCPUTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxCPUTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxDocExamined(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxDocExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxFrows(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxFrows = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxIOWrites(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxIOWrites = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxKeysExamined(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxKeysExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxLastRowsCountAffected(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxLastRowsCountAffected = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxLockTime(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxLockTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxLockTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxLockTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxLogicalIOReads(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxLogicalIOReads = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxPhysicalIOReads(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxPhysicalIOReads = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxQueryTime(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxQueryTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxQueryTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxQueryTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxRequestSize(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxRequestSize = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxResponseSize(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxResponseSize = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxReturnNum(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxReturnNum = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxRows(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxRows = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxRowsCountAffected(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxRowsCountAffected = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxRowsExamined(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxRowsExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxRowsSent(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxRowsSent = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxRt(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxRt = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetMaxScnt(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.MaxScnt = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetNamespace(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Namespace = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetNodeType(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.NodeType = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetOpType(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.OpType = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetOriginAlias(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.OriginAlias = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetPhysicalIOReads(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.PhysicalIOReads = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetPsql(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Psql = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetQueryId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.QueryId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetQueryStartTime(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetQueryTime(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.QueryTime = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetQueryTimeRate(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.QueryTimeRate = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetQueryTimeSeconds(v float64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.QueryTimeSeconds = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetReturnItemNumbers(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.ReturnItemNumbers = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetReturnNum(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.ReturnNum = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetRows(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Rows = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetRowsCountAffected(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.RowsCountAffected = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetRowsExamined(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.RowsExamined = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetRowsSent(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.RowsSent = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetRuleId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.RuleId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetSQLText(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetScheme(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Scheme = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetScnt(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Scnt = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetSqlId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.SqlId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetSqlTag(v *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.SqlTag = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetSqlType(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.SqlType = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetSubInstanceId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.SubInstanceId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetTableName(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.TableName = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetThreadId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.ThreadId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetTimestamp(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Timestamp = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetTotalCount(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.TotalCount = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetTrend(v []*DescribeSlowLogStatisticResponseBodyDataDataLogsTrend) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.Trend = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) SetUser(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogs {
	s.User = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogs) Validate() error {
	if s.Histogram != nil {
		if err := s.Histogram.Validate(); err != nil {
			return err
		}
	}
	if s.SqlTag != nil {
		if err := s.SqlTag.Validate(); err != nil {
			return err
		}
	}
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

type DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram struct {
	AvgLockTime     []*float64                                                       `json:"AvgLockTime,omitempty" xml:"AvgLockTime,omitempty" type:"Repeated"`
	AvgRowsExamined []*float64                                                       `json:"AvgRowsExamined,omitempty" xml:"AvgRowsExamined,omitempty" type:"Repeated"`
	AvgRowsSent     []*float64                                                       `json:"AvgRowsSent,omitempty" xml:"AvgRowsSent,omitempty" type:"Repeated"`
	AvgRt           []*float64                                                       `json:"AvgRt,omitempty" xml:"AvgRt,omitempty" type:"Repeated"`
	Count           []*int64                                                         `json:"Count,omitempty" xml:"Count,omitempty" type:"Repeated"`
	Item            []*DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	LockTime        []*float64                                                       `json:"LockTime,omitempty" xml:"LockTime,omitempty" type:"Repeated"`
	MaxLockTime     []*float64                                                       `json:"MaxLockTime,omitempty" xml:"MaxLockTime,omitempty" type:"Repeated"`
	MaxRowsExamined []*int64                                                         `json:"MaxRowsExamined,omitempty" xml:"MaxRowsExamined,omitempty" type:"Repeated"`
	MaxRowsSent     []*int64                                                         `json:"MaxRowsSent,omitempty" xml:"MaxRowsSent,omitempty" type:"Repeated"`
	MaxRt           []*float64                                                       `json:"MaxRt,omitempty" xml:"MaxRt,omitempty" type:"Repeated"`
	RowsExamined    []*int64                                                         `json:"RowsExamined,omitempty" xml:"RowsExamined,omitempty" type:"Repeated"`
	RowsSent        []*int64                                                         `json:"RowsSent,omitempty" xml:"RowsSent,omitempty" type:"Repeated"`
	Rt              []*float64                                                       `json:"Rt,omitempty" xml:"Rt,omitempty" type:"Repeated"`
	// example:
	//
	// 7
	Total *int64   `json:"Total,omitempty" xml:"Total,omitempty"`
	Ts    []*int64 `json:"Ts,omitempty" xml:"Ts,omitempty" type:"Repeated"`
	TsEnd []*int64 `json:"TsEnd,omitempty" xml:"TsEnd,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetAvgLockTime() []*float64 {
	return s.AvgLockTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetAvgRowsExamined() []*float64 {
	return s.AvgRowsExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetAvgRowsSent() []*float64 {
	return s.AvgRowsSent
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetAvgRt() []*float64 {
	return s.AvgRt
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetCount() []*int64 {
	return s.Count
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetItem() []*DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem {
	return s.Item
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetLockTime() []*float64 {
	return s.LockTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetMaxLockTime() []*float64 {
	return s.MaxLockTime
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetMaxRowsExamined() []*int64 {
	return s.MaxRowsExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetMaxRowsSent() []*int64 {
	return s.MaxRowsSent
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetMaxRt() []*float64 {
	return s.MaxRt
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetRowsExamined() []*int64 {
	return s.RowsExamined
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetRowsSent() []*int64 {
	return s.RowsSent
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetRt() []*float64 {
	return s.Rt
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetTs() []*int64 {
	return s.Ts
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) GetTsEnd() []*int64 {
	return s.TsEnd
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetAvgLockTime(v []*float64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.AvgLockTime = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetAvgRowsExamined(v []*float64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.AvgRowsExamined = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetAvgRowsSent(v []*float64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.AvgRowsSent = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetAvgRt(v []*float64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.AvgRt = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetCount(v []*int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.Count = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetItem(v []*DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.Item = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetLockTime(v []*float64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.LockTime = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetMaxLockTime(v []*float64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.MaxLockTime = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetMaxRowsExamined(v []*int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.MaxRowsExamined = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetMaxRowsSent(v []*int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.MaxRowsSent = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetMaxRt(v []*float64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.MaxRt = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetRowsExamined(v []*int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.RowsExamined = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetRowsSent(v []*int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.RowsSent = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetRt(v []*float64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.Rt = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetTotal(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.Total = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetTs(v []*int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.Ts = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) SetTsEnd(v []*int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram {
	s.TsEnd = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogram) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem struct {
	Count []*int64 `json:"Count,omitempty" xml:"Count,omitempty" type:"Repeated"`
	// example:
	//
	// r-bp1s1m8hwzrm77kfvz-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem) GetCount() []*int64 {
	return s.Count
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem) SetCount(v []*int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem {
	s.Count = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem) SetNodeId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsHistogramItem) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag struct {
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// SQL ID。
	//
	// example:
	//
	// a3931d8c3a9315dd5ed016d71cf*****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// example:
	//
	// DAS_IN_PLAN
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) GetComments() *string {
	return s.Comments
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) GetSqlId() *string {
	return s.SqlId
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) GetTags() *string {
	return s.Tags
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) SetComments(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag {
	s.Comments = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) SetSqlId(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag {
	s.SqlId = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) SetTags(v string) *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag {
	s.Tags = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsSqlTag) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogStatisticResponseBodyDataDataLogsTrend struct {
	// example:
	//
	// 1723775362
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// tf-testacc-oos-parameter
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogsTrend) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponseBodyDataDataLogsTrend) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsTrend) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsTrend) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsTrend) SetTimestamp(v int64) *DescribeSlowLogStatisticResponseBodyDataDataLogsTrend {
	s.Timestamp = &v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsTrend) SetValue(v interface{}) *DescribeSlowLogStatisticResponseBodyDataDataLogsTrend {
	s.Value = v
	return s
}

func (s *DescribeSlowLogStatisticResponseBodyDataDataLogsTrend) Validate() error {
	return dara.Validate(s)
}
