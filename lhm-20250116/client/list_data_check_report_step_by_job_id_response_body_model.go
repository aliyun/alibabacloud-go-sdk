// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportStepByJobIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataCheckReportStepByJobIdResponseBodyData) *ListDataCheckReportStepByJobIdResponseBody
	GetData() []*ListDataCheckReportStepByJobIdResponseBodyData
	SetErrCode(v string) *ListDataCheckReportStepByJobIdResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListDataCheckReportStepByJobIdResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListDataCheckReportStepByJobIdResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckReportStepByJobIdResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataCheckReportStepByJobIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCheckReportStepByJobIdResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDataCheckReportStepByJobIdResponseBody
	GetTotalCount() *int32
}

type ListDataCheckReportStepByJobIdResponseBody struct {
	Data []*ListDataCheckReportStepByJobIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataCheckReportStepByJobIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportStepByJobIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportStepByJobIdResponseBody) GetData() []*ListDataCheckReportStepByJobIdResponseBodyData {
	return s.Data
}

func (s *ListDataCheckReportStepByJobIdResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListDataCheckReportStepByJobIdResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckReportStepByJobIdResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckReportStepByJobIdResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckReportStepByJobIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCheckReportStepByJobIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCheckReportStepByJobIdResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataCheckReportStepByJobIdResponseBody) SetData(v []*ListDataCheckReportStepByJobIdResponseBodyData) *ListDataCheckReportStepByJobIdResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBody) SetErrCode(v string) *ListDataCheckReportStepByJobIdResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBody) SetErrMessage(v string) *ListDataCheckReportStepByJobIdResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBody) SetPageIndex(v int32) *ListDataCheckReportStepByJobIdResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBody) SetPageSize(v int32) *ListDataCheckReportStepByJobIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBody) SetRequestId(v string) *ListDataCheckReportStepByJobIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBody) SetSuccess(v bool) *ListDataCheckReportStepByJobIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBody) SetTotalCount(v int32) *ListDataCheckReportStepByJobIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBody) Validate() error {
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

type ListDataCheckReportStepByJobIdResponseBodyData struct {
	// example:
	//
	// R[1->1000)
	Boundary *string `json:"boundary,omitempty" xml:"boundary,omitempty"`
	// example:
	//
	// 10
	CheckColumCount *int64 `json:"checkColumCount,omitempty" xml:"checkColumCount,omitempty"`
	// example:
	//
	// 1000
	DstCount *string `json:"dstCount,omitempty" xml:"dstCount,omitempty"`
	// example:
	//
	// SELECT 	- FROM dst_table;
	DstSql *string `json:"dstSql,omitempty" xml:"dstSql,omitempty"`
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 2026-01-16T10:00:00Z
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// example:
	//
	// 2026-01-16T10:00:00Z
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// example:
	//
	// 1
	IsConsistent *int32 `json:"isConsistent,omitempty" xml:"isConsistent,omitempty"`
	// example:
	//
	// 10
	MetricColumCount *int64 `json:"metricColumCount,omitempty" xml:"metricColumCount,omitempty"`
	// example:
	//
	// 8
	MetricPassColumCount *int64 `json:"metricPassColumCount,omitempty" xml:"metricPassColumCount,omitempty"`
	// example:
	//
	// 8
	PassColumCount *int64 `json:"passColumCount,omitempty" xml:"passColumCount,omitempty"`
	// example:
	//
	// 30001
	ResultId *string `json:"resultId,omitempty" xml:"resultId,omitempty"`
	// example:
	//
	// ds=20260116
	SourcePtName *string `json:"sourcePtName,omitempty" xml:"sourcePtName,omitempty"`
	// example:
	//
	// 1000
	SrcCount *string `json:"srcCount,omitempty" xml:"srcCount,omitempty"`
	// example:
	//
	// SELECT 	- FROM src_table;
	SrcSql *string `json:"srcSql,omitempty" xml:"srcSql,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	StepId *string `json:"stepId,omitempty" xml:"stepId,omitempty"`
	// example:
	//
	// ds=20260116
	TargetPtName *string `json:"targetPtName,omitempty" xml:"targetPtName,omitempty"`
}

func (s ListDataCheckReportStepByJobIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportStepByJobIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetBoundary() *string {
	return s.Boundary
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetCheckColumCount() *int64 {
	return s.CheckColumCount
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetDstCount() *string {
	return s.DstCount
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetDstSql() *string {
	return s.DstSql
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetGmtEnd() *string {
	return s.GmtEnd
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetGmtStart() *string {
	return s.GmtStart
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetIsConsistent() *int32 {
	return s.IsConsistent
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetMetricColumCount() *int64 {
	return s.MetricColumCount
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetMetricPassColumCount() *int64 {
	return s.MetricPassColumCount
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetPassColumCount() *int64 {
	return s.PassColumCount
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetSourcePtName() *string {
	return s.SourcePtName
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetSrcCount() *string {
	return s.SrcCount
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetSrcSql() *string {
	return s.SrcSql
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetStepId() *string {
	return s.StepId
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) GetTargetPtName() *string {
	return s.TargetPtName
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetBoundary(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.Boundary = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetCheckColumCount(v int64) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.CheckColumCount = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetDstCount(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.DstCount = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetDstSql(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.DstSql = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetErrMessage(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetExtra(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.Extra = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetGmtEnd(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.GmtEnd = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetGmtStart(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.GmtStart = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetIsConsistent(v int32) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetMetricColumCount(v int64) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.MetricColumCount = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetMetricPassColumCount(v int64) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.MetricPassColumCount = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetPassColumCount(v int64) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.PassColumCount = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetResultId(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetSourcePtName(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.SourcePtName = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetSrcCount(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.SrcCount = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetSrcSql(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.SrcSql = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetStatus(v int32) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetStepId(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.StepId = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) SetTargetPtName(v string) *ListDataCheckReportStepByJobIdResponseBodyData {
	s.TargetPtName = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
