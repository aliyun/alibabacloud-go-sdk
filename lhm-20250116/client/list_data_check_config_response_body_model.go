// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataCheckConfigResponseBodyData) *ListDataCheckConfigResponseBody
	GetData() []*ListDataCheckConfigResponseBodyData
	SetErrCode(v string) *ListDataCheckConfigResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListDataCheckConfigResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListDataCheckConfigResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckConfigResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataCheckConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCheckConfigResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDataCheckConfigResponseBody
	GetTotalCount() *int32
}

type ListDataCheckConfigResponseBody struct {
	Data []*ListDataCheckConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ListDataCheckConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCheckConfigResponseBody) GetData() []*ListDataCheckConfigResponseBodyData {
	return s.Data
}

func (s *ListDataCheckConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListDataCheckConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckConfigResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckConfigResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCheckConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCheckConfigResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataCheckConfigResponseBody) SetData(v []*ListDataCheckConfigResponseBodyData) *ListDataCheckConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCheckConfigResponseBody) SetErrCode(v string) *ListDataCheckConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDataCheckConfigResponseBody) SetErrMessage(v string) *ListDataCheckConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckConfigResponseBody) SetPageIndex(v int32) *ListDataCheckConfigResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckConfigResponseBody) SetPageSize(v int32) *ListDataCheckConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckConfigResponseBody) SetRequestId(v string) *ListDataCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCheckConfigResponseBody) SetSuccess(v bool) *ListDataCheckConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCheckConfigResponseBody) SetTotalCount(v int32) *ListDataCheckConfigResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataCheckConfigResponseBody) Validate() error {
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

type ListDataCheckConfigResponseBodyData struct {
	// example:
	//
	// 1
	Algorithm *int32 `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// example:
	//
	// 1000
	BatchSize *int32 `json:"batchSize,omitempty" xml:"batchSize,omitempty"`
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// example:
	//
	// =
	Comparator *string `json:"comparator,omitempty" xml:"comparator,omitempty"`
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 0.5
	GroupCountThreshold *float32 `json:"groupCountThreshold,omitempty" xml:"groupCountThreshold,omitempty"`
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	IsFullTableCount *int32 `json:"isFullTableCount,omitempty" xml:"isFullTableCount,omitempty"`
	// example:
	//
	// 0
	IsSkipped *int32 `json:"isSkipped,omitempty" xml:"isSkipped,omitempty"`
	// example:
	//
	// CUSTOM_METRIC_MIX
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	// example:
	//
	// 1
	SourceCheckAllColumn *int32 `json:"sourceCheckAllColumn,omitempty" xml:"sourceCheckAllColumn,omitempty"`
	// example:
	//
	// col_a,col_b
	SourceColumns *string `json:"sourceColumns,omitempty" xml:"sourceColumns,omitempty"`
	// example:
	//
	// id
	SourceCompareKey *string `json:"sourceCompareKey,omitempty" xml:"sourceCompareKey,omitempty"`
	// example:
	//
	// ds_demo
	SourceDataSource *string `json:"sourceDataSource,omitempty" xml:"sourceDataSource,omitempty"`
	// example:
	//
	// col_a,col_b
	SourceGroupClause *string `json:"sourceGroupClause,omitempty" xml:"sourceGroupClause,omitempty"`
	SourceHint        *string `json:"sourceHint,omitempty" xml:"sourceHint,omitempty"`
	// example:
	//
	// 1001
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// ds=20260116
	SourcePartition *string `json:"sourcePartition,omitempty" xml:"sourcePartition,omitempty"`
	// example:
	//
	// SELECT 	- FROM t;
	SourceSql *string `json:"sourceSql,omitempty" xml:"sourceSql,omitempty"`
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
	// 1
	TargetCheckAllColumn *int32 `json:"targetCheckAllColumn,omitempty" xml:"targetCheckAllColumn,omitempty"`
	// example:
	//
	// col_a,col_b
	TargetColumns *string `json:"targetColumns,omitempty" xml:"targetColumns,omitempty"`
	// example:
	//
	// id
	TargetCompareKey *string `json:"targetCompareKey,omitempty" xml:"targetCompareKey,omitempty"`
	// example:
	//
	// ds_demo
	TargetDataSource *string `json:"targetDataSource,omitempty" xml:"targetDataSource,omitempty"`
	// example:
	//
	// col_a,col_b
	TargetGroupClause *string `json:"targetGroupClause,omitempty" xml:"targetGroupClause,omitempty"`
	TargetHint        *string `json:"targetHint,omitempty" xml:"targetHint,omitempty"`
	// example:
	//
	// 2001
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// example:
	//
	// ds=20260116
	TargetPartition *string `json:"targetPartition,omitempty" xml:"targetPartition,omitempty"`
	// example:
	//
	// SELECT 	- FROM t;
	TargetSql *string `json:"targetSql,omitempty" xml:"targetSql,omitempty"`
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
	// lhm|lhm_dw|*
	TaskConfigInfo *string `json:"taskConfigInfo,omitempty" xml:"taskConfigInfo,omitempty"`
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 0.5
	TotalCountThreshold *float32 `json:"totalCountThreshold,omitempty" xml:"totalCountThreshold,omitempty"`
}

func (s ListDataCheckConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCheckConfigResponseBodyData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *ListDataCheckConfigResponseBodyData) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *ListDataCheckConfigResponseBodyData) GetCheckType() *int32 {
	return s.CheckType
}

func (s *ListDataCheckConfigResponseBodyData) GetComparator() *string {
	return s.Comparator
}

func (s *ListDataCheckConfigResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *ListDataCheckConfigResponseBodyData) GetGroupCountThreshold() *float32 {
	return s.GroupCountThreshold
}

func (s *ListDataCheckConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListDataCheckConfigResponseBodyData) GetIsFullTableCount() *int32 {
	return s.IsFullTableCount
}

func (s *ListDataCheckConfigResponseBodyData) GetIsSkipped() *int32 {
	return s.IsSkipped
}

func (s *ListDataCheckConfigResponseBodyData) GetMetricType() *string {
	return s.MetricType
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceCheckAllColumn() *int32 {
	return s.SourceCheckAllColumn
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceColumns() *string {
	return s.SourceColumns
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceCompareKey() *string {
	return s.SourceCompareKey
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceDataSource() *string {
	return s.SourceDataSource
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceGroupClause() *string {
	return s.SourceGroupClause
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceHint() *string {
	return s.SourceHint
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *ListDataCheckConfigResponseBodyData) GetSourcePartition() *string {
	return s.SourcePartition
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceSql() *string {
	return s.SourceSql
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceTable() *string {
	return s.SourceTable
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *ListDataCheckConfigResponseBodyData) GetSourceWhereClause() *string {
	return s.SourceWhereClause
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetCheckAllColumn() *int32 {
	return s.TargetCheckAllColumn
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetColumns() *string {
	return s.TargetColumns
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetCompareKey() *string {
	return s.TargetCompareKey
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetDataSource() *string {
	return s.TargetDataSource
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetGroupClause() *string {
	return s.TargetGroupClause
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetHint() *string {
	return s.TargetHint
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetPartition() *string {
	return s.TargetPartition
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetSql() *string {
	return s.TargetSql
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetTable() *string {
	return s.TargetTable
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *ListDataCheckConfigResponseBodyData) GetTargetWhereClause() *string {
	return s.TargetWhereClause
}

func (s *ListDataCheckConfigResponseBodyData) GetTaskConfigInfo() *string {
	return s.TaskConfigInfo
}

func (s *ListDataCheckConfigResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListDataCheckConfigResponseBodyData) GetTotalCountThreshold() *float32 {
	return s.TotalCountThreshold
}

func (s *ListDataCheckConfigResponseBodyData) SetAlgorithm(v int32) *ListDataCheckConfigResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetBatchSize(v int32) *ListDataCheckConfigResponseBodyData {
	s.BatchSize = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetCheckType(v int32) *ListDataCheckConfigResponseBodyData {
	s.CheckType = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetComparator(v string) *ListDataCheckConfigResponseBodyData {
	s.Comparator = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetExtra(v string) *ListDataCheckConfigResponseBodyData {
	s.Extra = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetGroupCountThreshold(v float32) *ListDataCheckConfigResponseBodyData {
	s.GroupCountThreshold = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetId(v int64) *ListDataCheckConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetIsFullTableCount(v int32) *ListDataCheckConfigResponseBodyData {
	s.IsFullTableCount = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetIsSkipped(v int32) *ListDataCheckConfigResponseBodyData {
	s.IsSkipped = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetMetricType(v string) *ListDataCheckConfigResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceCheckAllColumn(v int32) *ListDataCheckConfigResponseBodyData {
	s.SourceCheckAllColumn = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceColumns(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceColumns = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceCompareKey(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceCompareKey = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceDataSource(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceDataSource = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceGroupClause(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceGroupClause = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceHint(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceHint = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceId(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourcePartition(v string) *ListDataCheckConfigResponseBodyData {
	s.SourcePartition = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceSql(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceSql = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceTable(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceTable = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceType(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetSourceWhereClause(v string) *ListDataCheckConfigResponseBodyData {
	s.SourceWhereClause = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetCheckAllColumn(v int32) *ListDataCheckConfigResponseBodyData {
	s.TargetCheckAllColumn = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetColumns(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetColumns = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetCompareKey(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetCompareKey = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetDataSource(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetDataSource = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetGroupClause(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetGroupClause = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetHint(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetHint = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetId(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetPartition(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetPartition = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetSql(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetSql = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetTable(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetTable = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetType(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTargetWhereClause(v string) *ListDataCheckConfigResponseBodyData {
	s.TargetWhereClause = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTaskConfigInfo(v string) *ListDataCheckConfigResponseBodyData {
	s.TaskConfigInfo = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTaskId(v int64) *ListDataCheckConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) SetTotalCountThreshold(v float32) *ListDataCheckConfigResponseBodyData {
	s.TotalCountThreshold = &v
	return s
}

func (s *ListDataCheckConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
