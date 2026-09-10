// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsFullTableCount(v int32) *AddDataCheckConfigRequest
	GetIsFullTableCount() *int32
	SetSourceColumns(v string) *AddDataCheckConfigRequest
	GetSourceColumns() *string
	SetSourceGroupClause(v string) *AddDataCheckConfigRequest
	GetSourceGroupClause() *string
	SetSourceHint(v string) *AddDataCheckConfigRequest
	GetSourceHint() *string
	SetSourcePartition(v string) *AddDataCheckConfigRequest
	GetSourcePartition() *string
	SetSourceTable(v string) *AddDataCheckConfigRequest
	GetSourceTable() *string
	SetSourceWhereClause(v string) *AddDataCheckConfigRequest
	GetSourceWhereClause() *string
	SetTargetColumns(v string) *AddDataCheckConfigRequest
	GetTargetColumns() *string
	SetTargetGroupClause(v string) *AddDataCheckConfigRequest
	GetTargetGroupClause() *string
	SetTargetHint(v string) *AddDataCheckConfigRequest
	GetTargetHint() *string
	SetTargetPartition(v string) *AddDataCheckConfigRequest
	GetTargetPartition() *string
	SetTargetTable(v string) *AddDataCheckConfigRequest
	GetTargetTable() *string
	SetTargetWhereClause(v string) *AddDataCheckConfigRequest
	GetTargetWhereClause() *string
	SetTaskConfigInfo(v string) *AddDataCheckConfigRequest
	GetTaskConfigInfo() *string
	SetTaskId(v int64) *AddDataCheckConfigRequest
	GetTaskId() *int64
	SetTotalCountThreshold(v float32) *AddDataCheckConfigRequest
	GetTotalCountThreshold() *float32
}

type AddDataCheckConfigRequest struct {
	// Specifies whether to perform full-table validation. Valid values:
	//
	// - 0: partition-level comparison.
	//
	// - 1: full-table comparison.
	//
	// example:
	//
	// 0
	IsFullTableCount *int32 `json:"isFullTableCount,omitempty" xml:"isFullTableCount,omitempty"`
	// The columns of the source table. You can specify multiple columns separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	SourceColumns *string `json:"sourceColumns,omitempty" xml:"sourceColumns,omitempty"`
	// The GROUP condition of the source table.
	//
	// example:
	//
	// col_a,col_b
	SourceGroupClause *string `json:"sourceGroupClause,omitempty" xml:"sourceGroupClause,omitempty"`
	// The hint for the source.
	SourceHint *string `json:"sourceHint,omitempty" xml:"sourceHint,omitempty"`
	// The partition of the source table.
	//
	// example:
	//
	// ds=20260116
	SourcePartition *string `json:"sourcePartition,omitempty" xml:"sourcePartition,omitempty"`
	// The name of the source table.
	//
	// example:
	//
	// table_demo
	SourceTable *string `json:"sourceTable,omitempty" xml:"sourceTable,omitempty"`
	// The WHERE condition of the source table.
	//
	// example:
	//
	// col_a > 0 and col_b = \\"x\\"
	SourceWhereClause *string `json:"sourceWhereClause,omitempty" xml:"sourceWhereClause,omitempty"`
	// The columns of the target table. You can specify multiple columns separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	TargetColumns *string `json:"targetColumns,omitempty" xml:"targetColumns,omitempty"`
	// The GROUP condition of the target table.
	//
	// example:
	//
	// col_a,col_b
	TargetGroupClause *string `json:"targetGroupClause,omitempty" xml:"targetGroupClause,omitempty"`
	// The hint for the target.
	TargetHint *string `json:"targetHint,omitempty" xml:"targetHint,omitempty"`
	// The partition of the target table.
	//
	// example:
	//
	// ds=20260116
	TargetPartition *string `json:"targetPartition,omitempty" xml:"targetPartition,omitempty"`
	// The name of the target table.
	//
	// example:
	//
	// table_demo
	TargetTable *string `json:"targetTable,omitempty" xml:"targetTable,omitempty"`
	// The WHERE condition of the target table.
	//
	// example:
	//
	// col_a > 0 and col_b = \\"x\\"
	TargetWhereClause *string `json:"targetWhereClause,omitempty" xml:"targetWhereClause,omitempty"`
	// The batch table configurations for same-pattern creation (`taskMode=1`). Separate multiple configurations with a line break (`
	//
	// `).
	//
	// example:
	//
	// lhm|lhm_dw|*
	TaskConfigInfo *string `json:"taskConfigInfo,omitempty" xml:"taskConfigInfo,omitempty"`
	// The ID of the validation task.
	//
	// This parameter is required.
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

func (s AddDataCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *AddDataCheckConfigRequest) GetIsFullTableCount() *int32 {
	return s.IsFullTableCount
}

func (s *AddDataCheckConfigRequest) GetSourceColumns() *string {
	return s.SourceColumns
}

func (s *AddDataCheckConfigRequest) GetSourceGroupClause() *string {
	return s.SourceGroupClause
}

func (s *AddDataCheckConfigRequest) GetSourceHint() *string {
	return s.SourceHint
}

func (s *AddDataCheckConfigRequest) GetSourcePartition() *string {
	return s.SourcePartition
}

func (s *AddDataCheckConfigRequest) GetSourceTable() *string {
	return s.SourceTable
}

func (s *AddDataCheckConfigRequest) GetSourceWhereClause() *string {
	return s.SourceWhereClause
}

func (s *AddDataCheckConfigRequest) GetTargetColumns() *string {
	return s.TargetColumns
}

func (s *AddDataCheckConfigRequest) GetTargetGroupClause() *string {
	return s.TargetGroupClause
}

func (s *AddDataCheckConfigRequest) GetTargetHint() *string {
	return s.TargetHint
}

func (s *AddDataCheckConfigRequest) GetTargetPartition() *string {
	return s.TargetPartition
}

func (s *AddDataCheckConfigRequest) GetTargetTable() *string {
	return s.TargetTable
}

func (s *AddDataCheckConfigRequest) GetTargetWhereClause() *string {
	return s.TargetWhereClause
}

func (s *AddDataCheckConfigRequest) GetTaskConfigInfo() *string {
	return s.TaskConfigInfo
}

func (s *AddDataCheckConfigRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *AddDataCheckConfigRequest) GetTotalCountThreshold() *float32 {
	return s.TotalCountThreshold
}

func (s *AddDataCheckConfigRequest) SetIsFullTableCount(v int32) *AddDataCheckConfigRequest {
	s.IsFullTableCount = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetSourceColumns(v string) *AddDataCheckConfigRequest {
	s.SourceColumns = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetSourceGroupClause(v string) *AddDataCheckConfigRequest {
	s.SourceGroupClause = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetSourceHint(v string) *AddDataCheckConfigRequest {
	s.SourceHint = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetSourcePartition(v string) *AddDataCheckConfigRequest {
	s.SourcePartition = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetSourceTable(v string) *AddDataCheckConfigRequest {
	s.SourceTable = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetSourceWhereClause(v string) *AddDataCheckConfigRequest {
	s.SourceWhereClause = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTargetColumns(v string) *AddDataCheckConfigRequest {
	s.TargetColumns = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTargetGroupClause(v string) *AddDataCheckConfigRequest {
	s.TargetGroupClause = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTargetHint(v string) *AddDataCheckConfigRequest {
	s.TargetHint = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTargetPartition(v string) *AddDataCheckConfigRequest {
	s.TargetPartition = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTargetTable(v string) *AddDataCheckConfigRequest {
	s.TargetTable = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTargetWhereClause(v string) *AddDataCheckConfigRequest {
	s.TargetWhereClause = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTaskConfigInfo(v string) *AddDataCheckConfigRequest {
	s.TaskConfigInfo = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTaskId(v int64) *AddDataCheckConfigRequest {
	s.TaskId = &v
	return s
}

func (s *AddDataCheckConfigRequest) SetTotalCountThreshold(v float32) *AddDataCheckConfigRequest {
	s.TotalCountThreshold = &v
	return s
}

func (s *AddDataCheckConfigRequest) Validate() error {
	return dara.Validate(s)
}
