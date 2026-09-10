// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataCheckReportResponseBodyData) *ListDataCheckReportResponseBody
	GetData() []*ListDataCheckReportResponseBodyData
	SetErrCode(v string) *ListDataCheckReportResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListDataCheckReportResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListDataCheckReportResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckReportResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataCheckReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCheckReportResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDataCheckReportResponseBody
	GetTotalCount() *int32
}

type ListDataCheckReportResponseBody struct {
	Data []*ListDataCheckReportResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ListDataCheckReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportResponseBody) GetData() []*ListDataCheckReportResponseBodyData {
	return s.Data
}

func (s *ListDataCheckReportResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListDataCheckReportResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckReportResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckReportResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCheckReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCheckReportResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataCheckReportResponseBody) SetData(v []*ListDataCheckReportResponseBodyData) *ListDataCheckReportResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCheckReportResponseBody) SetErrCode(v string) *ListDataCheckReportResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDataCheckReportResponseBody) SetErrMessage(v string) *ListDataCheckReportResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckReportResponseBody) SetPageIndex(v int32) *ListDataCheckReportResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckReportResponseBody) SetPageSize(v int32) *ListDataCheckReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckReportResponseBody) SetRequestId(v string) *ListDataCheckReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCheckReportResponseBody) SetSuccess(v bool) *ListDataCheckReportResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCheckReportResponseBody) SetTotalCount(v int32) *ListDataCheckReportResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataCheckReportResponseBody) Validate() error {
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

type ListDataCheckReportResponseBodyData struct {
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// example:
	//
	// 10
	CheckColumCount *int64 `json:"checkColumCount,omitempty" xml:"checkColumCount,omitempty"`
	// example:
	//
	// 1
	CheckResult *int32 `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// example:
	//
	// 1000
	CompareRowCount *int64 `json:"compareRowCount,omitempty" xml:"compareRowCount,omitempty"`
	// example:
	//
	// 95.00%
	CompletionRate *string `json:"completionRate,omitempty" xml:"completionRate,omitempty"`
	// example:
	//
	// 0.00%
	DiffRate *string `json:"diffRate,omitempty" xml:"diffRate,omitempty"`
	// example:
	//
	// id
	DstCompareColumn *string `json:"dstCompareColumn,omitempty" xml:"dstCompareColumn,omitempty"`
	DstHint          *string `json:"dstHint,omitempty" xml:"dstHint,omitempty"`
	// example:
	//
	// amount
	DstMetricName *string   `json:"dstMetricName,omitempty" xml:"dstMetricName,omitempty"`
	DstSqlList    []*string `json:"dstSqlList,omitempty" xml:"dstSqlList,omitempty" type:"Repeated"`
	// example:
	//
	// connection timeout
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 2026-01-16 10:00:00
	ExecTime *string `json:"execTime,omitempty" xml:"execTime,omitempty"`
	// example:
	//
	// 0
	ExpDiffCount *string `json:"expDiffCount,omitempty" xml:"expDiffCount,omitempty"`
	// example:
	//
	// 2026-01-16T10:00:00Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// example:
	//
	// 0
	IsSkipped *int32 `json:"isSkipped,omitempty" xml:"isSkipped,omitempty"`
	// example:
	//
	// 10001
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 2
	JobStatus *int32 `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
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
	// 0
	OnlyDstCount *int64 `json:"onlyDstCount,omitempty" xml:"onlyDstCount,omitempty"`
	// example:
	//
	// 0
	OnlySrcCount *int64 `json:"onlySrcCount,omitempty" xml:"onlySrcCount,omitempty"`
	// example:
	//
	// 8
	PassColumCount *int64 `json:"passColumCount,omitempty" xml:"passColumCount,omitempty"`
	// example:
	//
	// 0
	RealDiffCount *int64 `json:"realDiffCount,omitempty" xml:"realDiffCount,omitempty"`
	// example:
	//
	// 1000
	RealSameCount *int64 `json:"realSameCount,omitempty" xml:"realSameCount,omitempty"`
	// example:
	//
	// 30001
	ResultId *string `json:"resultId,omitempty" xml:"resultId,omitempty"`
	// example:
	//
	// amount
	SourceColumn *string `json:"sourceColumn,omitempty" xml:"sourceColumn,omitempty"`
	// example:
	//
	// 1000
	SourceCount *string `json:"sourceCount,omitempty" xml:"sourceCount,omitempty"`
	// example:
	//
	// ds_demo
	SourceDataSource *string `json:"sourceDataSource,omitempty" xml:"sourceDataSource,omitempty"`
	// example:
	//
	// Table \\"src_db.src_table\\" doesn\\"t exist
	SourceError *string `json:"sourceError,omitempty" xml:"sourceError,omitempty"`
	// example:
	//
	// col_a,col_b
	SourceGroupClause *string `json:"sourceGroupClause,omitempty" xml:"sourceGroupClause,omitempty"`
	// example:
	//
	// ds=20260116
	SourcePartition *string `json:"sourcePartition,omitempty" xml:"sourcePartition,omitempty"`
	// example:
	//
	// table_demo
	SourceTable *string `json:"sourceTable,omitempty" xml:"sourceTable,omitempty"`
	// example:
	//
	// Hive
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// col_a > 0 and col_b = \\"x\\"
	SourceWhereClause *string `json:"sourceWhereClause,omitempty" xml:"sourceWhereClause,omitempty"`
	// example:
	//
	// id
	SrcCompareColumn *string `json:"srcCompareColumn,omitempty" xml:"srcCompareColumn,omitempty"`
	SrcHint          *string `json:"srcHint,omitempty" xml:"srcHint,omitempty"`
	// example:
	//
	// amount
	SrcMetricName *string   `json:"srcMetricName,omitempty" xml:"srcMetricName,omitempty"`
	SrcSqlList    []*string `json:"srcSqlList,omitempty" xml:"srcSqlList,omitempty" type:"Repeated"`
	// example:
	//
	// amount
	TargetColumn *string `json:"targetColumn,omitempty" xml:"targetColumn,omitempty"`
	// example:
	//
	// 1000
	TargetCount *string `json:"targetCount,omitempty" xml:"targetCount,omitempty"`
	// example:
	//
	// ds_demo
	TargetDataSource *string `json:"targetDataSource,omitempty" xml:"targetDataSource,omitempty"`
	// example:
	//
	// Table \\"dst_db.dst_table\\" doesn\\"t exist
	TargetError *string `json:"targetError,omitempty" xml:"targetError,omitempty"`
	// example:
	//
	// col_a,col_b
	TargetGroupClause *string `json:"targetGroupClause,omitempty" xml:"targetGroupClause,omitempty"`
	// example:
	//
	// ds=20260116
	TargetPartition *string `json:"targetPartition,omitempty" xml:"targetPartition,omitempty"`
	// example:
	//
	// table_demo
	TargetTable *string `json:"targetTable,omitempty" xml:"targetTable,omitempty"`
	// example:
	//
	// hive
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// example:
	//
	// col_a > 0 and col_b = \\"x\\"
	TargetWhereClause *string `json:"targetWhereClause,omitempty" xml:"targetWhereClause,omitempty"`
	// example:
	//
	// 1001
	TaskConfigId *int64 `json:"taskConfigId,omitempty" xml:"taskConfigId,omitempty"`
	// example:
	//
	// 数据量校验模板
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// example:
	//
	// 0.0
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// example:
	//
	// 0.5
	TotalCountThreshold *string `json:"totalCountThreshold,omitempty" xml:"totalCountThreshold,omitempty"`
}

func (s ListDataCheckReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportResponseBodyData) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ListDataCheckReportResponseBodyData) GetCheckColumCount() *int64 {
	return s.CheckColumCount
}

func (s *ListDataCheckReportResponseBodyData) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *ListDataCheckReportResponseBodyData) GetCompareRowCount() *int64 {
	return s.CompareRowCount
}

func (s *ListDataCheckReportResponseBodyData) GetCompletionRate() *string {
	return s.CompletionRate
}

func (s *ListDataCheckReportResponseBodyData) GetDiffRate() *string {
	return s.DiffRate
}

func (s *ListDataCheckReportResponseBodyData) GetDstCompareColumn() *string {
	return s.DstCompareColumn
}

func (s *ListDataCheckReportResponseBodyData) GetDstHint() *string {
	return s.DstHint
}

func (s *ListDataCheckReportResponseBodyData) GetDstMetricName() *string {
	return s.DstMetricName
}

func (s *ListDataCheckReportResponseBodyData) GetDstSqlList() []*string {
	return s.DstSqlList
}

func (s *ListDataCheckReportResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListDataCheckReportResponseBodyData) GetExecTime() *string {
	return s.ExecTime
}

func (s *ListDataCheckReportResponseBodyData) GetExpDiffCount() *string {
	return s.ExpDiffCount
}

func (s *ListDataCheckReportResponseBodyData) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListDataCheckReportResponseBodyData) GetIsSkipped() *int32 {
	return s.IsSkipped
}

func (s *ListDataCheckReportResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *ListDataCheckReportResponseBodyData) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *ListDataCheckReportResponseBodyData) GetMetricColumCount() *int64 {
	return s.MetricColumCount
}

func (s *ListDataCheckReportResponseBodyData) GetMetricPassColumCount() *int64 {
	return s.MetricPassColumCount
}

func (s *ListDataCheckReportResponseBodyData) GetOnlyDstCount() *int64 {
	return s.OnlyDstCount
}

func (s *ListDataCheckReportResponseBodyData) GetOnlySrcCount() *int64 {
	return s.OnlySrcCount
}

func (s *ListDataCheckReportResponseBodyData) GetPassColumCount() *int64 {
	return s.PassColumCount
}

func (s *ListDataCheckReportResponseBodyData) GetRealDiffCount() *int64 {
	return s.RealDiffCount
}

func (s *ListDataCheckReportResponseBodyData) GetRealSameCount() *int64 {
	return s.RealSameCount
}

func (s *ListDataCheckReportResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *ListDataCheckReportResponseBodyData) GetSourceColumn() *string {
	return s.SourceColumn
}

func (s *ListDataCheckReportResponseBodyData) GetSourceCount() *string {
	return s.SourceCount
}

func (s *ListDataCheckReportResponseBodyData) GetSourceDataSource() *string {
	return s.SourceDataSource
}

func (s *ListDataCheckReportResponseBodyData) GetSourceError() *string {
	return s.SourceError
}

func (s *ListDataCheckReportResponseBodyData) GetSourceGroupClause() *string {
	return s.SourceGroupClause
}

func (s *ListDataCheckReportResponseBodyData) GetSourcePartition() *string {
	return s.SourcePartition
}

func (s *ListDataCheckReportResponseBodyData) GetSourceTable() *string {
	return s.SourceTable
}

func (s *ListDataCheckReportResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *ListDataCheckReportResponseBodyData) GetSourceWhereClause() *string {
	return s.SourceWhereClause
}

func (s *ListDataCheckReportResponseBodyData) GetSrcCompareColumn() *string {
	return s.SrcCompareColumn
}

func (s *ListDataCheckReportResponseBodyData) GetSrcHint() *string {
	return s.SrcHint
}

func (s *ListDataCheckReportResponseBodyData) GetSrcMetricName() *string {
	return s.SrcMetricName
}

func (s *ListDataCheckReportResponseBodyData) GetSrcSqlList() []*string {
	return s.SrcSqlList
}

func (s *ListDataCheckReportResponseBodyData) GetTargetColumn() *string {
	return s.TargetColumn
}

func (s *ListDataCheckReportResponseBodyData) GetTargetCount() *string {
	return s.TargetCount
}

func (s *ListDataCheckReportResponseBodyData) GetTargetDataSource() *string {
	return s.TargetDataSource
}

func (s *ListDataCheckReportResponseBodyData) GetTargetError() *string {
	return s.TargetError
}

func (s *ListDataCheckReportResponseBodyData) GetTargetGroupClause() *string {
	return s.TargetGroupClause
}

func (s *ListDataCheckReportResponseBodyData) GetTargetPartition() *string {
	return s.TargetPartition
}

func (s *ListDataCheckReportResponseBodyData) GetTargetTable() *string {
	return s.TargetTable
}

func (s *ListDataCheckReportResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *ListDataCheckReportResponseBodyData) GetTargetWhereClause() *string {
	return s.TargetWhereClause
}

func (s *ListDataCheckReportResponseBodyData) GetTaskConfigId() *int64 {
	return s.TaskConfigId
}

func (s *ListDataCheckReportResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListDataCheckReportResponseBodyData) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ListDataCheckReportResponseBodyData) GetTotalCountThreshold() *string {
	return s.TotalCountThreshold
}

func (s *ListDataCheckReportResponseBodyData) SetBatchId(v int64) *ListDataCheckReportResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetCheckColumCount(v int64) *ListDataCheckReportResponseBodyData {
	s.CheckColumCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetCheckResult(v int32) *ListDataCheckReportResponseBodyData {
	s.CheckResult = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetCompareRowCount(v int64) *ListDataCheckReportResponseBodyData {
	s.CompareRowCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetCompletionRate(v string) *ListDataCheckReportResponseBodyData {
	s.CompletionRate = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetDiffRate(v string) *ListDataCheckReportResponseBodyData {
	s.DiffRate = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetDstCompareColumn(v string) *ListDataCheckReportResponseBodyData {
	s.DstCompareColumn = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetDstHint(v string) *ListDataCheckReportResponseBodyData {
	s.DstHint = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetDstMetricName(v string) *ListDataCheckReportResponseBodyData {
	s.DstMetricName = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetDstSqlList(v []*string) *ListDataCheckReportResponseBodyData {
	s.DstSqlList = v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetErrorMsg(v string) *ListDataCheckReportResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetExecTime(v string) *ListDataCheckReportResponseBodyData {
	s.ExecTime = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetExpDiffCount(v string) *ListDataCheckReportResponseBodyData {
	s.ExpDiffCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetFinishTime(v string) *ListDataCheckReportResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetIsSkipped(v int32) *ListDataCheckReportResponseBodyData {
	s.IsSkipped = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetJobId(v string) *ListDataCheckReportResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetJobStatus(v int32) *ListDataCheckReportResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetMetricColumCount(v int64) *ListDataCheckReportResponseBodyData {
	s.MetricColumCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetMetricPassColumCount(v int64) *ListDataCheckReportResponseBodyData {
	s.MetricPassColumCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetOnlyDstCount(v int64) *ListDataCheckReportResponseBodyData {
	s.OnlyDstCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetOnlySrcCount(v int64) *ListDataCheckReportResponseBodyData {
	s.OnlySrcCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetPassColumCount(v int64) *ListDataCheckReportResponseBodyData {
	s.PassColumCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetRealDiffCount(v int64) *ListDataCheckReportResponseBodyData {
	s.RealDiffCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetRealSameCount(v int64) *ListDataCheckReportResponseBodyData {
	s.RealSameCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetResultId(v string) *ListDataCheckReportResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourceColumn(v string) *ListDataCheckReportResponseBodyData {
	s.SourceColumn = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourceCount(v string) *ListDataCheckReportResponseBodyData {
	s.SourceCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourceDataSource(v string) *ListDataCheckReportResponseBodyData {
	s.SourceDataSource = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourceError(v string) *ListDataCheckReportResponseBodyData {
	s.SourceError = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourceGroupClause(v string) *ListDataCheckReportResponseBodyData {
	s.SourceGroupClause = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourcePartition(v string) *ListDataCheckReportResponseBodyData {
	s.SourcePartition = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourceTable(v string) *ListDataCheckReportResponseBodyData {
	s.SourceTable = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourceType(v string) *ListDataCheckReportResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSourceWhereClause(v string) *ListDataCheckReportResponseBodyData {
	s.SourceWhereClause = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSrcCompareColumn(v string) *ListDataCheckReportResponseBodyData {
	s.SrcCompareColumn = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSrcHint(v string) *ListDataCheckReportResponseBodyData {
	s.SrcHint = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSrcMetricName(v string) *ListDataCheckReportResponseBodyData {
	s.SrcMetricName = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetSrcSqlList(v []*string) *ListDataCheckReportResponseBodyData {
	s.SrcSqlList = v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetColumn(v string) *ListDataCheckReportResponseBodyData {
	s.TargetColumn = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetCount(v string) *ListDataCheckReportResponseBodyData {
	s.TargetCount = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetDataSource(v string) *ListDataCheckReportResponseBodyData {
	s.TargetDataSource = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetError(v string) *ListDataCheckReportResponseBodyData {
	s.TargetError = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetGroupClause(v string) *ListDataCheckReportResponseBodyData {
	s.TargetGroupClause = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetPartition(v string) *ListDataCheckReportResponseBodyData {
	s.TargetPartition = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetTable(v string) *ListDataCheckReportResponseBodyData {
	s.TargetTable = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetType(v string) *ListDataCheckReportResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTargetWhereClause(v string) *ListDataCheckReportResponseBodyData {
	s.TargetWhereClause = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTaskConfigId(v int64) *ListDataCheckReportResponseBodyData {
	s.TaskConfigId = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTemplateName(v string) *ListDataCheckReportResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetThreshold(v float32) *ListDataCheckReportResponseBodyData {
	s.Threshold = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) SetTotalCountThreshold(v string) *ListDataCheckReportResponseBodyData {
	s.TotalCountThreshold = &v
	return s
}

func (s *ListDataCheckReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
