// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetDataCheckConfigResponseBodyData) *GetDataCheckConfigResponseBody
	GetData() []*GetDataCheckConfigResponseBodyData
	SetErrCode(v string) *GetDataCheckConfigResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetDataCheckConfigResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetDataCheckConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCheckConfigResponseBody
	GetSuccess() *bool
}

type GetDataCheckConfigResponseBody struct {
	// The data list returned by the operation. For the structure of each element, see the child field descriptions.
	Data []*GetDataCheckConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check the values of errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDataCheckConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCheckConfigResponseBody) GetData() []*GetDataCheckConfigResponseBodyData {
	return s.Data
}

func (s *GetDataCheckConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDataCheckConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetDataCheckConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCheckConfigResponseBody) SetData(v []*GetDataCheckConfigResponseBodyData) *GetDataCheckConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetDataCheckConfigResponseBody) SetErrCode(v string) *GetDataCheckConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataCheckConfigResponseBody) SetErrMessage(v string) *GetDataCheckConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDataCheckConfigResponseBody) SetRequestId(v string) *GetDataCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckConfigResponseBody) SetSuccess(v bool) *GetDataCheckConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCheckConfigResponseBody) Validate() error {
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

type GetDataCheckConfigResponseBodyData struct {
	// The check algorithm.
	//
	// example:
	//
	// 1
	Algorithm *int32 `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// The batch size.
	//
	// example:
	//
	// 1000
	BatchSize *int32 `json:"batchSize,omitempty" xml:"batchSize,omitempty"`
	// The check type.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The comparison type. Valid values: =, !=, >, <, >=, <=, contains, does not contain, and ==.
	//
	// example:
	//
	// =
	Comparator *string `json:"comparator,omitempty" xml:"comparator,omitempty"`
	// The reserved field.
	//
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The group data volume comparison threshold.
	//
	// example:
	//
	// 0.5
	GroupCountThreshold *float32 `json:"groupCountThreshold,omitempty" xml:"groupCountThreshold,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether a full table count is performed.
	//
	// example:
	//
	// 1
	IsFullTableCount *int32 `json:"isFullTableCount,omitempty" xml:"isFullTableCount,omitempty"`
	// Indicates whether the check is skipped.
	//
	// example:
	//
	// 0
	IsSkipped *int32 `json:"isSkipped,omitempty" xml:"isSkipped,omitempty"`
	// The metric type. Valid values:
	//
	// - CUSTOM_METRIC_NUM: built-in NUM mode.
	//
	// - CUSTOM_METRIC_LEN: built-in LEN mode.
	//
	// - CUSTOM_METRIC_MIX: built-in MIX mode.
	//
	// example:
	//
	// CUSTOM_METRIC_MIX
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	// Indicates whether all columns are checked on the source side. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 1
	SourceCheckAllColumn *int32 `json:"sourceCheckAllColumn,omitempty" xml:"sourceCheckAllColumn,omitempty"`
	// The source table columns. You can specify multiple columns separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	SourceColumns *string `json:"sourceColumns,omitempty" xml:"sourceColumns,omitempty"`
	// The source comparison key.
	//
	// example:
	//
	// id
	SourceCompareKey *string `json:"sourceCompareKey,omitempty" xml:"sourceCompareKey,omitempty"`
	// The name of the source datasource.
	//
	// example:
	//
	// ds_demo
	SourceDataSource *string `json:"sourceDataSource,omitempty" xml:"sourceDataSource,omitempty"`
	// The GROUP BY clause for the source table.
	//
	// example:
	//
	// col_a,col_b
	SourceGroupClause *string `json:"sourceGroupClause,omitempty" xml:"sourceGroupClause,omitempty"`
	// The hint for the source side.
	SourceHint *string `json:"sourceHint,omitempty" xml:"sourceHint,omitempty"`
	// The ID of the source datasource.
	//
	// example:
	//
	// 1001
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The source partition.
	//
	// example:
	//
	// ds=20260116
	SourcePartition *string `json:"sourcePartition,omitempty" xml:"sourcePartition,omitempty"`
	// The SQL statement for the source side.
	//
	// example:
	//
	// SELECT 	- FROM t;
	SourceSql *string `json:"sourceSql,omitempty" xml:"sourceSql,omitempty"`
	// The source table.
	//
	// example:
	//
	// table_demo
	SourceTable *string `json:"sourceTable,omitempty" xml:"sourceTable,omitempty"`
	// The type of the source datasource.
	//
	// example:
	//
	// Hive
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The WHERE clause for the source table.
	//
	// example:
	//
	// col_a > 0 and col_b = \\"x\\"
	SourceWhereClause *string `json:"sourceWhereClause,omitempty" xml:"sourceWhereClause,omitempty"`
	// Indicates whether all columns are checked on the target side. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 1
	TargetCheckAllColumn *int32 `json:"targetCheckAllColumn,omitempty" xml:"targetCheckAllColumn,omitempty"`
	// The target table columns. You can specify multiple columns separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	TargetColumns *string `json:"targetColumns,omitempty" xml:"targetColumns,omitempty"`
	// The target comparison key.
	//
	// example:
	//
	// id
	TargetCompareKey *string `json:"targetCompareKey,omitempty" xml:"targetCompareKey,omitempty"`
	// The target datasource.
	//
	// example:
	//
	// ds_demo
	TargetDataSource *string `json:"targetDataSource,omitempty" xml:"targetDataSource,omitempty"`
	// The GROUP BY clause for the target table.
	//
	// example:
	//
	// col_a,col_b
	TargetGroupClause *string `json:"targetGroupClause,omitempty" xml:"targetGroupClause,omitempty"`
	// The hint for the target side.
	TargetHint *string `json:"targetHint,omitempty" xml:"targetHint,omitempty"`
	// The ID of the target datasource.
	//
	// example:
	//
	// 2001
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// The target partition.
	//
	// example:
	//
	// ds=20260116
	TargetPartition *string `json:"targetPartition,omitempty" xml:"targetPartition,omitempty"`
	// The SQL statement for the target side.
	//
	// example:
	//
	// SELECT 	- FROM t;
	TargetSql *string `json:"targetSql,omitempty" xml:"targetSql,omitempty"`
	// The target table.
	//
	// example:
	//
	// table_demo
	TargetTable *string `json:"targetTable,omitempty" xml:"targetTable,omitempty"`
	// The type of the target datasource.
	//
	// example:
	//
	// hive
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The WHERE clause for the target table.
	//
	// example:
	//
	// col_a > 0 and col_b = \\"x\\"
	TargetWhereClause *string `json:"targetWhereClause,omitempty" xml:"targetWhereClause,omitempty"`
	// The configuration details.
	//
	// example:
	//
	// lhm|lhm_dw|*
	TaskConfigInfo *string `json:"taskConfigInfo,omitempty" xml:"taskConfigInfo,omitempty"`
	// The batch ID.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The total data volume comparison threshold.
	//
	// example:
	//
	// 0.5
	TotalCountThreshold *float32 `json:"totalCountThreshold,omitempty" xml:"totalCountThreshold,omitempty"`
}

func (s GetDataCheckConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataCheckConfigResponseBodyData) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *GetDataCheckConfigResponseBodyData) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *GetDataCheckConfigResponseBodyData) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckConfigResponseBodyData) GetComparator() *string {
	return s.Comparator
}

func (s *GetDataCheckConfigResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetDataCheckConfigResponseBodyData) GetGroupCountThreshold() *float32 {
	return s.GroupCountThreshold
}

func (s *GetDataCheckConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetDataCheckConfigResponseBodyData) GetIsFullTableCount() *int32 {
	return s.IsFullTableCount
}

func (s *GetDataCheckConfigResponseBodyData) GetIsSkipped() *int32 {
	return s.IsSkipped
}

func (s *GetDataCheckConfigResponseBodyData) GetMetricType() *string {
	return s.MetricType
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceCheckAllColumn() *int32 {
	return s.SourceCheckAllColumn
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceColumns() *string {
	return s.SourceColumns
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceCompareKey() *string {
	return s.SourceCompareKey
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceDataSource() *string {
	return s.SourceDataSource
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceGroupClause() *string {
	return s.SourceGroupClause
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceHint() *string {
	return s.SourceHint
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *GetDataCheckConfigResponseBodyData) GetSourcePartition() *string {
	return s.SourcePartition
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceSql() *string {
	return s.SourceSql
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceTable() *string {
	return s.SourceTable
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *GetDataCheckConfigResponseBodyData) GetSourceWhereClause() *string {
	return s.SourceWhereClause
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetCheckAllColumn() *int32 {
	return s.TargetCheckAllColumn
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetColumns() *string {
	return s.TargetColumns
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetCompareKey() *string {
	return s.TargetCompareKey
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetDataSource() *string {
	return s.TargetDataSource
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetGroupClause() *string {
	return s.TargetGroupClause
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetHint() *string {
	return s.TargetHint
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetPartition() *string {
	return s.TargetPartition
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetSql() *string {
	return s.TargetSql
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetTable() *string {
	return s.TargetTable
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *GetDataCheckConfigResponseBodyData) GetTargetWhereClause() *string {
	return s.TargetWhereClause
}

func (s *GetDataCheckConfigResponseBodyData) GetTaskConfigInfo() *string {
	return s.TaskConfigInfo
}

func (s *GetDataCheckConfigResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetDataCheckConfigResponseBodyData) GetTotalCountThreshold() *float32 {
	return s.TotalCountThreshold
}

func (s *GetDataCheckConfigResponseBodyData) SetAlgorithm(v int32) *GetDataCheckConfigResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetBatchSize(v int32) *GetDataCheckConfigResponseBodyData {
	s.BatchSize = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetCheckType(v int32) *GetDataCheckConfigResponseBodyData {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetComparator(v string) *GetDataCheckConfigResponseBodyData {
	s.Comparator = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetExtra(v string) *GetDataCheckConfigResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetGroupCountThreshold(v float32) *GetDataCheckConfigResponseBodyData {
	s.GroupCountThreshold = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetId(v int64) *GetDataCheckConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetIsFullTableCount(v int32) *GetDataCheckConfigResponseBodyData {
	s.IsFullTableCount = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetIsSkipped(v int32) *GetDataCheckConfigResponseBodyData {
	s.IsSkipped = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetMetricType(v string) *GetDataCheckConfigResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceCheckAllColumn(v int32) *GetDataCheckConfigResponseBodyData {
	s.SourceCheckAllColumn = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceColumns(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceColumns = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceCompareKey(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceCompareKey = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceDataSource(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceDataSource = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceGroupClause(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceGroupClause = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceHint(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceHint = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceId(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourcePartition(v string) *GetDataCheckConfigResponseBodyData {
	s.SourcePartition = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceSql(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceSql = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceTable(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceTable = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceType(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetSourceWhereClause(v string) *GetDataCheckConfigResponseBodyData {
	s.SourceWhereClause = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetCheckAllColumn(v int32) *GetDataCheckConfigResponseBodyData {
	s.TargetCheckAllColumn = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetColumns(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetColumns = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetCompareKey(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetCompareKey = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetDataSource(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetDataSource = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetGroupClause(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetGroupClause = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetHint(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetHint = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetId(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetPartition(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetPartition = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetSql(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetSql = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetTable(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetTable = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetType(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTargetWhereClause(v string) *GetDataCheckConfigResponseBodyData {
	s.TargetWhereClause = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTaskConfigInfo(v string) *GetDataCheckConfigResponseBodyData {
	s.TaskConfigInfo = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTaskId(v int64) *GetDataCheckConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) SetTotalCountThreshold(v float32) *GetDataCheckConfigResponseBodyData {
	s.TotalCountThreshold = &v
	return s
}

func (s *GetDataCheckConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
