// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogHistogramAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSlowLogHistogramAsyncResponseBody
	GetCode() *string
	SetData(v *DescribeSlowLogHistogramAsyncResponseBodyData) *DescribeSlowLogHistogramAsyncResponseBody
	GetData() *DescribeSlowLogHistogramAsyncResponseBodyData
	SetMessage(v string) *DescribeSlowLogHistogramAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSlowLogHistogramAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSlowLogHistogramAsyncResponseBody
	GetSuccess() *string
}

type DescribeSlowLogHistogramAsyncResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// AsyncResult<Histogram>。
	Data *DescribeSlowLogHistogramAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSlowLogHistogramAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogHistogramAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) GetData() *DescribeSlowLogHistogramAsyncResponseBodyData {
	return s.Data
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) SetCode(v string) *DescribeSlowLogHistogramAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) SetData(v *DescribeSlowLogHistogramAsyncResponseBodyData) *DescribeSlowLogHistogramAsyncResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) SetMessage(v string) *DescribeSlowLogHistogramAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) SetRequestId(v string) *DescribeSlowLogHistogramAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) SetSuccess(v string) *DescribeSlowLogHistogramAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogHistogramAsyncResponseBodyData struct {
	Data *DescribeSlowLogHistogramAsyncResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestKey *string `json:"RequestKey,omitempty" xml:"RequestKey,omitempty"`
	// example:
	//
	// async__20ee808e72257f16a4fe024057ca****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	// example:
	//
	// SUCCESS
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1645668213000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeSlowLogHistogramAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogHistogramAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) GetData() *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	return s.Data
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) GetRequestKey() *string {
	return s.RequestKey
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) GetState() *string {
	return s.State
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) SetData(v *DescribeSlowLogHistogramAsyncResponseBodyDataData) *DescribeSlowLogHistogramAsyncResponseBodyData {
	s.Data = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) SetErrorCode(v int32) *DescribeSlowLogHistogramAsyncResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) SetIsFinish(v bool) *DescribeSlowLogHistogramAsyncResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) SetMessage(v string) *DescribeSlowLogHistogramAsyncResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) SetRequestKey(v string) *DescribeSlowLogHistogramAsyncResponseBodyData {
	s.RequestKey = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) SetResultId(v string) *DescribeSlowLogHistogramAsyncResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) SetState(v string) *DescribeSlowLogHistogramAsyncResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) SetTimestamp(v int64) *DescribeSlowLogHistogramAsyncResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogHistogramAsyncResponseBodyDataData struct {
	AvgCPUTime               []*float64                                               `json:"AvgCPUTime,omitempty" xml:"AvgCPUTime,omitempty" type:"Repeated"`
	AvgDocExamined           []*float64                                               `json:"AvgDocExamined,omitempty" xml:"AvgDocExamined,omitempty" type:"Repeated"`
	AvgFrows                 []*float64                                               `json:"AvgFrows,omitempty" xml:"AvgFrows,omitempty" type:"Repeated"`
	AvgIOWrites              []*float64                                               `json:"AvgIOWrites,omitempty" xml:"AvgIOWrites,omitempty" type:"Repeated"`
	AvgKeysExamined          []*float64                                               `json:"AvgKeysExamined,omitempty" xml:"AvgKeysExamined,omitempty" type:"Repeated"`
	AvgLastRowsCountAffected []*float64                                               `json:"AvgLastRowsCountAffected,omitempty" xml:"AvgLastRowsCountAffected,omitempty" type:"Repeated"`
	AvgLockTime              []*float64                                               `json:"AvgLockTime,omitempty" xml:"AvgLockTime,omitempty" type:"Repeated"`
	AvgLogicalIOReads        []*float64                                               `json:"AvgLogicalIOReads,omitempty" xml:"AvgLogicalIOReads,omitempty" type:"Repeated"`
	AvgPhysicalIOReads       []*float64                                               `json:"AvgPhysicalIOReads,omitempty" xml:"AvgPhysicalIOReads,omitempty" type:"Repeated"`
	AvgReturnNum             []*float64                                               `json:"AvgReturnNum,omitempty" xml:"AvgReturnNum,omitempty" type:"Repeated"`
	AvgRows                  []*float64                                               `json:"AvgRows,omitempty" xml:"AvgRows,omitempty" type:"Repeated"`
	AvgRowsCountAffected     []*float64                                               `json:"AvgRowsCountAffected,omitempty" xml:"AvgRowsCountAffected,omitempty" type:"Repeated"`
	AvgRowsExamined          []*float64                                               `json:"AvgRowsExamined,omitempty" xml:"AvgRowsExamined,omitempty" type:"Repeated"`
	AvgRowsSent              []*float64                                               `json:"AvgRowsSent,omitempty" xml:"AvgRowsSent,omitempty" type:"Repeated"`
	AvgRt                    []*float64                                               `json:"AvgRt,omitempty" xml:"AvgRt,omitempty" type:"Repeated"`
	AvgScnt                  []*float64                                               `json:"AvgScnt,omitempty" xml:"AvgScnt,omitempty" type:"Repeated"`
	CPUTime                  []*float64                                               `json:"CPUTime,omitempty" xml:"CPUTime,omitempty" type:"Repeated"`
	Count                    []*int64                                                 `json:"Count,omitempty" xml:"Count,omitempty" type:"Repeated"`
	DocExamined              []*int64                                                 `json:"DocExamined,omitempty" xml:"DocExamined,omitempty" type:"Repeated"`
	Frows                    []*int64                                                 `json:"Frows,omitempty" xml:"Frows,omitempty" type:"Repeated"`
	IOWrites                 []*int64                                                 `json:"IOWrites,omitempty" xml:"IOWrites,omitempty" type:"Repeated"`
	Item                     []*DescribeSlowLogHistogramAsyncResponseBodyDataDataItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	KeysExamined             []*int64                                                 `json:"KeysExamined,omitempty" xml:"KeysExamined,omitempty" type:"Repeated"`
	LastRowsCountAffected    []*int64                                                 `json:"LastRowsCountAffected,omitempty" xml:"LastRowsCountAffected,omitempty" type:"Repeated"`
	LockTime                 []*float64                                               `json:"LockTime,omitempty" xml:"LockTime,omitempty" type:"Repeated"`
	LogicalIOReads           []*int64                                                 `json:"LogicalIOReads,omitempty" xml:"LogicalIOReads,omitempty" type:"Repeated"`
	MaxCPUTime               []*float64                                               `json:"MaxCPUTime,omitempty" xml:"MaxCPUTime,omitempty" type:"Repeated"`
	MaxDocExamined           []*int64                                                 `json:"MaxDocExamined,omitempty" xml:"MaxDocExamined,omitempty" type:"Repeated"`
	MaxFrows                 []*int64                                                 `json:"MaxFrows,omitempty" xml:"MaxFrows,omitempty" type:"Repeated"`
	MaxIOWrites              []*int64                                                 `json:"MaxIOWrites,omitempty" xml:"MaxIOWrites,omitempty" type:"Repeated"`
	MaxKeysExamined          []*int64                                                 `json:"MaxKeysExamined,omitempty" xml:"MaxKeysExamined,omitempty" type:"Repeated"`
	MaxLastRowsCountAffected []*int64                                                 `json:"MaxLastRowsCountAffected,omitempty" xml:"MaxLastRowsCountAffected,omitempty" type:"Repeated"`
	MaxLockTime              []*float64                                               `json:"MaxLockTime,omitempty" xml:"MaxLockTime,omitempty" type:"Repeated"`
	MaxLogicalIOReads        []*int64                                                 `json:"MaxLogicalIOReads,omitempty" xml:"MaxLogicalIOReads,omitempty" type:"Repeated"`
	MaxPhysicalIOReads       []*int64                                                 `json:"MaxPhysicalIOReads,omitempty" xml:"MaxPhysicalIOReads,omitempty" type:"Repeated"`
	MaxReturnNum             []*int64                                                 `json:"MaxReturnNum,omitempty" xml:"MaxReturnNum,omitempty" type:"Repeated"`
	MaxRows                  []*int64                                                 `json:"MaxRows,omitempty" xml:"MaxRows,omitempty" type:"Repeated"`
	MaxRowsCountAffected     []*int64                                                 `json:"MaxRowsCountAffected,omitempty" xml:"MaxRowsCountAffected,omitempty" type:"Repeated"`
	MaxRowsExamined          []*int64                                                 `json:"MaxRowsExamined,omitempty" xml:"MaxRowsExamined,omitempty" type:"Repeated"`
	MaxRowsSent              []*int64                                                 `json:"MaxRowsSent,omitempty" xml:"MaxRowsSent,omitempty" type:"Repeated"`
	MaxRt                    []*float64                                               `json:"MaxRt,omitempty" xml:"MaxRt,omitempty" type:"Repeated"`
	MaxScnt                  []*int64                                                 `json:"MaxScnt,omitempty" xml:"MaxScnt,omitempty" type:"Repeated"`
	PhysicalIOReads          []*int64                                                 `json:"PhysicalIOReads,omitempty" xml:"PhysicalIOReads,omitempty" type:"Repeated"`
	ReturnNum                []*int64                                                 `json:"ReturnNum,omitempty" xml:"ReturnNum,omitempty" type:"Repeated"`
	Rows                     []*int64                                                 `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	RowsCountAffected        []*int64                                                 `json:"RowsCountAffected,omitempty" xml:"RowsCountAffected,omitempty" type:"Repeated"`
	RowsExamined             []*int64                                                 `json:"RowsExamined,omitempty" xml:"RowsExamined,omitempty" type:"Repeated"`
	RowsSent                 []*int64                                                 `json:"RowsSent,omitempty" xml:"RowsSent,omitempty" type:"Repeated"`
	Rt                       []*float64                                               `json:"Rt,omitempty" xml:"Rt,omitempty" type:"Repeated"`
	Scnt                     []*int64                                                 `json:"Scnt,omitempty" xml:"Scnt,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Total *int64   `json:"Total,omitempty" xml:"Total,omitempty"`
	Ts    []*int64 `json:"Ts,omitempty" xml:"Ts,omitempty" type:"Repeated"`
	TsEnd []*int64 `json:"TsEnd,omitempty" xml:"TsEnd,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogHistogramAsyncResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogHistogramAsyncResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgCPUTime() []*float64 {
	return s.AvgCPUTime
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgDocExamined() []*float64 {
	return s.AvgDocExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgFrows() []*float64 {
	return s.AvgFrows
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgIOWrites() []*float64 {
	return s.AvgIOWrites
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgKeysExamined() []*float64 {
	return s.AvgKeysExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgLastRowsCountAffected() []*float64 {
	return s.AvgLastRowsCountAffected
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgLockTime() []*float64 {
	return s.AvgLockTime
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgLogicalIOReads() []*float64 {
	return s.AvgLogicalIOReads
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgPhysicalIOReads() []*float64 {
	return s.AvgPhysicalIOReads
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgReturnNum() []*float64 {
	return s.AvgReturnNum
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgRows() []*float64 {
	return s.AvgRows
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgRowsCountAffected() []*float64 {
	return s.AvgRowsCountAffected
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgRowsExamined() []*float64 {
	return s.AvgRowsExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgRowsSent() []*float64 {
	return s.AvgRowsSent
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgRt() []*float64 {
	return s.AvgRt
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetAvgScnt() []*float64 {
	return s.AvgScnt
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetCPUTime() []*float64 {
	return s.CPUTime
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetCount() []*int64 {
	return s.Count
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetDocExamined() []*int64 {
	return s.DocExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetFrows() []*int64 {
	return s.Frows
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetIOWrites() []*int64 {
	return s.IOWrites
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetItem() []*DescribeSlowLogHistogramAsyncResponseBodyDataDataItem {
	return s.Item
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetKeysExamined() []*int64 {
	return s.KeysExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetLastRowsCountAffected() []*int64 {
	return s.LastRowsCountAffected
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetLockTime() []*float64 {
	return s.LockTime
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetLogicalIOReads() []*int64 {
	return s.LogicalIOReads
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxCPUTime() []*float64 {
	return s.MaxCPUTime
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxDocExamined() []*int64 {
	return s.MaxDocExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxFrows() []*int64 {
	return s.MaxFrows
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxIOWrites() []*int64 {
	return s.MaxIOWrites
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxKeysExamined() []*int64 {
	return s.MaxKeysExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxLastRowsCountAffected() []*int64 {
	return s.MaxLastRowsCountAffected
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxLockTime() []*float64 {
	return s.MaxLockTime
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxLogicalIOReads() []*int64 {
	return s.MaxLogicalIOReads
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxPhysicalIOReads() []*int64 {
	return s.MaxPhysicalIOReads
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxReturnNum() []*int64 {
	return s.MaxReturnNum
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxRows() []*int64 {
	return s.MaxRows
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxRowsCountAffected() []*int64 {
	return s.MaxRowsCountAffected
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxRowsExamined() []*int64 {
	return s.MaxRowsExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxRowsSent() []*int64 {
	return s.MaxRowsSent
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxRt() []*float64 {
	return s.MaxRt
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetMaxScnt() []*int64 {
	return s.MaxScnt
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetPhysicalIOReads() []*int64 {
	return s.PhysicalIOReads
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetReturnNum() []*int64 {
	return s.ReturnNum
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetRows() []*int64 {
	return s.Rows
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetRowsCountAffected() []*int64 {
	return s.RowsCountAffected
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetRowsExamined() []*int64 {
	return s.RowsExamined
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetRowsSent() []*int64 {
	return s.RowsSent
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetRt() []*float64 {
	return s.Rt
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetScnt() []*int64 {
	return s.Scnt
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetTs() []*int64 {
	return s.Ts
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) GetTsEnd() []*int64 {
	return s.TsEnd
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgCPUTime(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgCPUTime = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgDocExamined(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgDocExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgFrows(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgFrows = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgIOWrites(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgIOWrites = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgKeysExamined(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgKeysExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgLastRowsCountAffected(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgLastRowsCountAffected = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgLockTime(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgLockTime = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgLogicalIOReads(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgLogicalIOReads = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgPhysicalIOReads(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgPhysicalIOReads = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgReturnNum(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgReturnNum = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgRows(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgRows = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgRowsCountAffected(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgRowsCountAffected = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgRowsExamined(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgRowsExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgRowsSent(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgRowsSent = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgRt(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgRt = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetAvgScnt(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.AvgScnt = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetCPUTime(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.CPUTime = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetCount(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.Count = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetDocExamined(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.DocExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetFrows(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.Frows = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetIOWrites(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.IOWrites = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetItem(v []*DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.Item = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetKeysExamined(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.KeysExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetLastRowsCountAffected(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.LastRowsCountAffected = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetLockTime(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.LockTime = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetLogicalIOReads(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.LogicalIOReads = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxCPUTime(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxCPUTime = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxDocExamined(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxDocExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxFrows(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxFrows = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxIOWrites(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxIOWrites = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxKeysExamined(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxKeysExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxLastRowsCountAffected(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxLastRowsCountAffected = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxLockTime(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxLockTime = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxLogicalIOReads(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxLogicalIOReads = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxPhysicalIOReads(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxPhysicalIOReads = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxReturnNum(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxReturnNum = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxRows(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxRows = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxRowsCountAffected(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxRowsCountAffected = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxRowsExamined(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxRowsExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxRowsSent(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxRowsSent = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxRt(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxRt = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetMaxScnt(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.MaxScnt = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetPhysicalIOReads(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.PhysicalIOReads = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetReturnNum(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.ReturnNum = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetRows(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.Rows = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetRowsCountAffected(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.RowsCountAffected = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetRowsExamined(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.RowsExamined = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetRowsSent(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.RowsSent = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetRt(v []*float64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.Rt = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetScnt(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.Scnt = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetTotal(v int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.Total = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetTs(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.Ts = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) SetTsEnd(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataData {
	s.TsEnd = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogHistogramAsyncResponseBodyDataDataItem struct {
	Count    []*int64                                                         `json:"Count,omitempty" xml:"Count,omitempty" type:"Repeated"`
	InsItems []*DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems `json:"InsItems,omitempty" xml:"InsItems,omitempty" type:"Repeated"`
	// example:
	//
	// r-bp1hi0wg57s3n0i3n8-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) GetCount() []*int64 {
	return s.Count
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) GetInsItems() []*DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems {
	return s.InsItems
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) SetCount(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem {
	s.Count = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) SetInsItems(v []*DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem {
	s.InsItems = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) SetNodeId(v string) *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItem) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems struct {
	Count []*int64 `json:"Count,omitempty" xml:"Count,omitempty" type:"Repeated"`
	// example:
	//
	// 2492
	InsId *string `json:"InsId,omitempty" xml:"InsId,omitempty"`
	// example:
	//
	// userAdmin
	InsRole *string `json:"InsRole,omitempty" xml:"InsRole,omitempty"`
}

func (s DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) GetCount() []*int64 {
	return s.Count
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) GetInsId() *string {
	return s.InsId
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) GetInsRole() *string {
	return s.InsRole
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) SetCount(v []*int64) *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems {
	s.Count = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) SetInsId(v string) *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems {
	s.InsId = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) SetInsRole(v string) *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems {
	s.InsRole = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncResponseBodyDataDataItemInsItems) Validate() error {
	return dara.Validate(s)
}
