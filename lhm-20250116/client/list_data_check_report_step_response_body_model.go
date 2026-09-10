// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportStepResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataCheckReportStepResponseBodyData) *ListDataCheckReportStepResponseBody
	GetData() []*ListDataCheckReportStepResponseBodyData
	SetErrCode(v string) *ListDataCheckReportStepResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListDataCheckReportStepResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListDataCheckReportStepResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckReportStepResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataCheckReportStepResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCheckReportStepResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDataCheckReportStepResponseBody
	GetTotalCount() *int32
}

type ListDataCheckReportStepResponseBody struct {
	// The data list returned by the operation. For the structure of each element, see the child field descriptions.
	Data []*ListDataCheckReportStepResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The page size, which is the number of records returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues with this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that meet the query conditions. This value is used for pagination.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataCheckReportStepResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportStepResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportStepResponseBody) GetData() []*ListDataCheckReportStepResponseBodyData {
	return s.Data
}

func (s *ListDataCheckReportStepResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListDataCheckReportStepResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckReportStepResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckReportStepResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckReportStepResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCheckReportStepResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCheckReportStepResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataCheckReportStepResponseBody) SetData(v []*ListDataCheckReportStepResponseBodyData) *ListDataCheckReportStepResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCheckReportStepResponseBody) SetErrCode(v string) *ListDataCheckReportStepResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDataCheckReportStepResponseBody) SetErrMessage(v string) *ListDataCheckReportStepResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckReportStepResponseBody) SetPageIndex(v int32) *ListDataCheckReportStepResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckReportStepResponseBody) SetPageSize(v int32) *ListDataCheckReportStepResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckReportStepResponseBody) SetRequestId(v string) *ListDataCheckReportStepResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCheckReportStepResponseBody) SetSuccess(v bool) *ListDataCheckReportStepResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCheckReportStepResponseBody) SetTotalCount(v int32) *ListDataCheckReportStepResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataCheckReportStepResponseBody) Validate() error {
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

type ListDataCheckReportStepResponseBodyData struct {
	// The shard boundary information.
	//
	// example:
	//
	// R[1->1000)
	Boundary *string `json:"boundary,omitempty" xml:"boundary,omitempty"`
	// The number of verified fields.
	//
	// example:
	//
	// 10
	CheckColumCount *int64 `json:"checkColumCount,omitempty" xml:"checkColumCount,omitempty"`
	// The destination data volume. For single-SQL verification, this indicates the data volume of the destination result set. For data volume verification and full verification, this indicates the total data volume on the destination.
	//
	// example:
	//
	// 1000
	DstCount *string `json:"dstCount,omitempty" xml:"dstCount,omitempty"`
	// The SQL statement executed on the destination.
	//
	// example:
	//
	// SELECT 	- FROM dst_table;
	DstSql *string `json:"dstSql,omitempty" xml:"dstSql,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The reserved field.
	//
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-16T10:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The job end time.
	//
	// example:
	//
	// 2026-01-16T10:00:00Z
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-01-16T10:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The job start time.
	//
	// example:
	//
	// 2026-01-16T10:00:00Z
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// The primary key ID that uniquely identifies a record.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the data is consistent. Valid values:
	//
	// - 0: inconsistent.
	//
	// - 1: consistent.
	//
	// example:
	//
	// 1
	IsConsistent *int32 `json:"isConsistent,omitempty" xml:"isConsistent,omitempty"`
	// The unique job ID.
	//
	// example:
	//
	// 10001
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The number of verified metrics.
	//
	// example:
	//
	// 10
	MetricColumCount *int64 `json:"metricColumCount,omitempty" xml:"metricColumCount,omitempty"`
	// The number of metrics that passed verification.
	//
	// example:
	//
	// 8
	MetricPassColumCount *int64 `json:"metricPassColumCount,omitempty" xml:"metricPassColumCount,omitempty"`
	// The number of fields that passed verification.
	//
	// example:
	//
	// 8
	PassColumCount *int64 `json:"passColumCount,omitempty" xml:"passColumCount,omitempty"`
	// The unique ID of the verification result.
	//
	// example:
	//
	// 30001
	ResultId *string `json:"resultId,omitempty" xml:"resultId,omitempty"`
	// The list of label names.
	SignNameList []*string `json:"signNameList,omitempty" xml:"signNameList,omitempty" type:"Repeated"`
	// The source partition name.
	//
	// example:
	//
	// ds=20260116
	SourcePtName *string `json:"sourcePtName,omitempty" xml:"sourcePtName,omitempty"`
	// The source data volume. For single-SQL verification, this indicates the data volume of the source result set. For data volume verification and full verification, this indicates the total data volume on the source.
	//
	// example:
	//
	// 1000
	SrcCount *string `json:"srcCount,omitempty" xml:"srcCount,omitempty"`
	// The SQL statement executed on the source.
	//
	// example:
	//
	// SELECT 	- FROM src_table;
	SrcSql *string `json:"srcSql,omitempty" xml:"srcSql,omitempty"`
	// The task status. Valid values:
	//
	// - 0: created.
	//
	// - 1: running.
	//
	// - 2: completed.
	//
	// - 3: stopped.
	//
	// - 4: canceled.
	//
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The step ID that uniquely identifies an execution step within the job.
	//
	// example:
	//
	// 1
	StepId *string `json:"stepId,omitempty" xml:"stepId,omitempty"`
	// The destination partition name.
	//
	// example:
	//
	// ds=20260116
	TargetPtName *string `json:"targetPtName,omitempty" xml:"targetPtName,omitempty"`
}

func (s ListDataCheckReportStepResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportStepResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportStepResponseBodyData) GetBoundary() *string {
	return s.Boundary
}

func (s *ListDataCheckReportStepResponseBodyData) GetCheckColumCount() *int64 {
	return s.CheckColumCount
}

func (s *ListDataCheckReportStepResponseBodyData) GetDstCount() *string {
	return s.DstCount
}

func (s *ListDataCheckReportStepResponseBodyData) GetDstSql() *string {
	return s.DstSql
}

func (s *ListDataCheckReportStepResponseBodyData) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckReportStepResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *ListDataCheckReportStepResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDataCheckReportStepResponseBodyData) GetGmtEnd() *string {
	return s.GmtEnd
}

func (s *ListDataCheckReportStepResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataCheckReportStepResponseBodyData) GetGmtStart() *string {
	return s.GmtStart
}

func (s *ListDataCheckReportStepResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListDataCheckReportStepResponseBodyData) GetIsConsistent() *int32 {
	return s.IsConsistent
}

func (s *ListDataCheckReportStepResponseBodyData) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDataCheckReportStepResponseBodyData) GetMetricColumCount() *int64 {
	return s.MetricColumCount
}

func (s *ListDataCheckReportStepResponseBodyData) GetMetricPassColumCount() *int64 {
	return s.MetricPassColumCount
}

func (s *ListDataCheckReportStepResponseBodyData) GetPassColumCount() *int64 {
	return s.PassColumCount
}

func (s *ListDataCheckReportStepResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *ListDataCheckReportStepResponseBodyData) GetSignNameList() []*string {
	return s.SignNameList
}

func (s *ListDataCheckReportStepResponseBodyData) GetSourcePtName() *string {
	return s.SourcePtName
}

func (s *ListDataCheckReportStepResponseBodyData) GetSrcCount() *string {
	return s.SrcCount
}

func (s *ListDataCheckReportStepResponseBodyData) GetSrcSql() *string {
	return s.SrcSql
}

func (s *ListDataCheckReportStepResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListDataCheckReportStepResponseBodyData) GetStepId() *string {
	return s.StepId
}

func (s *ListDataCheckReportStepResponseBodyData) GetTargetPtName() *string {
	return s.TargetPtName
}

func (s *ListDataCheckReportStepResponseBodyData) SetBoundary(v string) *ListDataCheckReportStepResponseBodyData {
	s.Boundary = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetCheckColumCount(v int64) *ListDataCheckReportStepResponseBodyData {
	s.CheckColumCount = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetDstCount(v string) *ListDataCheckReportStepResponseBodyData {
	s.DstCount = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetDstSql(v string) *ListDataCheckReportStepResponseBodyData {
	s.DstSql = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetErrMessage(v string) *ListDataCheckReportStepResponseBodyData {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetExtra(v string) *ListDataCheckReportStepResponseBodyData {
	s.Extra = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetGmtCreate(v string) *ListDataCheckReportStepResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetGmtEnd(v string) *ListDataCheckReportStepResponseBodyData {
	s.GmtEnd = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetGmtModified(v string) *ListDataCheckReportStepResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetGmtStart(v string) *ListDataCheckReportStepResponseBodyData {
	s.GmtStart = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetId(v int64) *ListDataCheckReportStepResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetIsConsistent(v int32) *ListDataCheckReportStepResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetJobId(v int64) *ListDataCheckReportStepResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetMetricColumCount(v int64) *ListDataCheckReportStepResponseBodyData {
	s.MetricColumCount = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetMetricPassColumCount(v int64) *ListDataCheckReportStepResponseBodyData {
	s.MetricPassColumCount = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetPassColumCount(v int64) *ListDataCheckReportStepResponseBodyData {
	s.PassColumCount = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetResultId(v string) *ListDataCheckReportStepResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetSignNameList(v []*string) *ListDataCheckReportStepResponseBodyData {
	s.SignNameList = v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetSourcePtName(v string) *ListDataCheckReportStepResponseBodyData {
	s.SourcePtName = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetSrcCount(v string) *ListDataCheckReportStepResponseBodyData {
	s.SrcCount = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetSrcSql(v string) *ListDataCheckReportStepResponseBodyData {
	s.SrcSql = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetStatus(v int32) *ListDataCheckReportStepResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetStepId(v string) *ListDataCheckReportStepResponseBodyData {
	s.StepId = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) SetTargetPtName(v string) *ListDataCheckReportStepResponseBodyData {
	s.TargetPtName = &v
	return s
}

func (s *ListDataCheckReportStepResponseBodyData) Validate() error {
	return dara.Validate(s)
}
