// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateDataDiagnosisRequest
	GetConfig() *string
	SetCycleTime(v string) *UpdateDataDiagnosisRequest
	GetCycleTime() *string
	SetInstanceId(v string) *UpdateDataDiagnosisRequest
	GetInstanceId() *string
	SetLeftTableMetaId(v string) *UpdateDataDiagnosisRequest
	GetLeftTableMetaId() *string
	SetLeftTablePartitionField(v string) *UpdateDataDiagnosisRequest
	GetLeftTablePartitionField() *string
	SetName(v string) *UpdateDataDiagnosisRequest
	GetName() *string
	SetPartitionField(v string) *UpdateDataDiagnosisRequest
	GetPartitionField() *string
	SetRightTableMetaId(v string) *UpdateDataDiagnosisRequest
	GetRightTableMetaId() *string
	SetRightTablePartitionField(v string) *UpdateDataDiagnosisRequest
	GetRightTablePartitionField() *string
	SetTableMetaId(v string) *UpdateDataDiagnosisRequest
	GetTableMetaId() *string
	SetTopNQuantity(v int64) *UpdateDataDiagnosisRequest
	GetTopNQuantity() *int64
	SetType(v string) *UpdateDataDiagnosisRequest
	GetType() *string
}

type UpdateDataDiagnosisRequest struct {
	// The configuration of the data diagnosis task, provided as a JSON string. The required fields vary based on the `Type` parameter. For `ChangeRate`, specify `AnalysisField` and `PartitionFieldFormat`. For `PreferenceStatisticsCycle`, specify `UserIdField`, `RemainDays`, `EverAppearedDays`, `RemainRatePeriods`, and `PartitionFieldFormat`. For `JoinTables`, specify `LeftTableAnalysisField`, `RightTableAnalysisField`, `LeftJoinField`, `RightJoinField`, `SampleQuantity`, `LeftTablePartitionFieldFormat`, and `RightTablePartitionFieldFormat`. For `BaseStatistics`, specify `TagField`, `TagFieldSeparator`, `KVField`, `KVFieldSeparator`, `KVPairSeparator`, `TextField`, `Quantiles`, `DefaultValueOfString`, `NullStringField`, and `PartitionFieldFormat`. For `AbnormalBehavior`, specify `UserId`, `ItemId`, `EventField`, `UpStreamBehavior`, `DownstreamBehavior`, `NumericHistogramBins`, and `PartitionFieldFormat`.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AnalysisField":"userid","PartitionFieldFormat":"yyyymmdd"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The execution time for periodic runs. If this field is omitted, the task does not run periodically.
	//
	// example:
	//
	// 08:00
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the left data table.
	//
	// example:
	//
	// 4
	LeftTableMetaId *string `json:"LeftTableMetaId,omitempty" xml:"LeftTableMetaId,omitempty"`
	// The partition field of the left table.
	//
	// example:
	//
	// dt
	LeftTablePartitionField *string `json:"LeftTablePartitionField,omitempty" xml:"LeftTablePartitionField,omitempty"`
	// The name of the data diagnosis task.
	//
	// This parameter is required.
	//
	// example:
	//
	// data_diagnosis_job1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The partition field.
	//
	// example:
	//
	// dt
	PartitionField *string `json:"PartitionField,omitempty" xml:"PartitionField,omitempty"`
	// The ID of the right data table.
	//
	// example:
	//
	// 5
	RightTableMetaId *string `json:"RightTableMetaId,omitempty" xml:"RightTableMetaId,omitempty"`
	// The partition field of the right table.
	//
	// example:
	//
	// dt
	RightTablePartitionField *string `json:"RightTablePartitionField,omitempty" xml:"RightTablePartitionField,omitempty"`
	// The ID of the data table.
	//
	// example:
	//
	// 3
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// The number of top-ranked results to return.
	//
	// example:
	//
	// 10
	TopNQuantity *int64 `json:"TopNQuantity,omitempty" xml:"TopNQuantity,omitempty"`
	// The type of the data diagnosis task. Valid values:
	//
	// - `ChangeRate`: item/user change rate analysis.
	//
	// - `PreferenceStatisticsCycle`: user preference statistics cycle analysis.
	//
	// - `JoinTables`: two-table join analysis.
	//
	// - `BaseStatistics`: basic statistical analysis.
	//
	// - `AbnormalBehavior`: abnormal behavior analysis.
	//
	// This parameter is required.
	//
	// example:
	//
	// ChangeRate
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataDiagnosisRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateDataDiagnosisRequest) GetCycleTime() *string {
	return s.CycleTime
}

func (s *UpdateDataDiagnosisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDataDiagnosisRequest) GetLeftTableMetaId() *string {
	return s.LeftTableMetaId
}

func (s *UpdateDataDiagnosisRequest) GetLeftTablePartitionField() *string {
	return s.LeftTablePartitionField
}

func (s *UpdateDataDiagnosisRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataDiagnosisRequest) GetPartitionField() *string {
	return s.PartitionField
}

func (s *UpdateDataDiagnosisRequest) GetRightTableMetaId() *string {
	return s.RightTableMetaId
}

func (s *UpdateDataDiagnosisRequest) GetRightTablePartitionField() *string {
	return s.RightTablePartitionField
}

func (s *UpdateDataDiagnosisRequest) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *UpdateDataDiagnosisRequest) GetTopNQuantity() *int64 {
	return s.TopNQuantity
}

func (s *UpdateDataDiagnosisRequest) GetType() *string {
	return s.Type
}

func (s *UpdateDataDiagnosisRequest) SetConfig(v string) *UpdateDataDiagnosisRequest {
	s.Config = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetCycleTime(v string) *UpdateDataDiagnosisRequest {
	s.CycleTime = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetInstanceId(v string) *UpdateDataDiagnosisRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetLeftTableMetaId(v string) *UpdateDataDiagnosisRequest {
	s.LeftTableMetaId = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetLeftTablePartitionField(v string) *UpdateDataDiagnosisRequest {
	s.LeftTablePartitionField = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetName(v string) *UpdateDataDiagnosisRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetPartitionField(v string) *UpdateDataDiagnosisRequest {
	s.PartitionField = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetRightTableMetaId(v string) *UpdateDataDiagnosisRequest {
	s.RightTableMetaId = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetRightTablePartitionField(v string) *UpdateDataDiagnosisRequest {
	s.RightTablePartitionField = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetTableMetaId(v string) *UpdateDataDiagnosisRequest {
	s.TableMetaId = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetTopNQuantity(v int64) *UpdateDataDiagnosisRequest {
	s.TopNQuantity = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) SetType(v string) *UpdateDataDiagnosisRequest {
	s.Type = &v
	return s
}

func (s *UpdateDataDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
