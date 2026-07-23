// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateDataDiagnosisRequest
	GetConfig() *string
	SetCycleTime(v string) *CreateDataDiagnosisRequest
	GetCycleTime() *string
	SetInstanceId(v string) *CreateDataDiagnosisRequest
	GetInstanceId() *string
	SetLeftTableMetaId(v string) *CreateDataDiagnosisRequest
	GetLeftTableMetaId() *string
	SetLeftTablePartitionField(v string) *CreateDataDiagnosisRequest
	GetLeftTablePartitionField() *string
	SetName(v string) *CreateDataDiagnosisRequest
	GetName() *string
	SetPartitionField(v string) *CreateDataDiagnosisRequest
	GetPartitionField() *string
	SetRightTableMetaId(v string) *CreateDataDiagnosisRequest
	GetRightTableMetaId() *string
	SetRightTablePartitionField(v string) *CreateDataDiagnosisRequest
	GetRightTablePartitionField() *string
	SetTableMetaId(v string) *CreateDataDiagnosisRequest
	GetTableMetaId() *string
	SetTopNQuantity(v int64) *CreateDataDiagnosisRequest
	GetTopNQuantity() *int64
	SetType(v string) *CreateDataDiagnosisRequest
	GetType() *string
}

type CreateDataDiagnosisRequest struct {
	// The configuration for the data diagnosis task, specified as a JSON string. The required fields in this object depend on the value of the `Type` parameter.
	//
	// - If `Type` is set to `ChangeRate`, specify the following fields: `AnalysisField` and `PartitionFieldFormat`.
	//
	// - If `Type` is set to `PreferenceStatisticsCycle`, specify the following fields: `UserIdField`, `RemainDays`, `EverAppearedDays`, `RemainRatePeriods`, and `PartitionFieldFormat`.
	//
	// - If `Type` is set to `JoinTables`, specify the following fields: `LeftTableAnalysisField`, `RightTableAnalysisField`, `LeftJoinField`, `RightJoinField`, `SampleQuantity`, `LeftTablePartitionFieldFormat`, and `RightTablePartitionFieldFormat`.
	//
	// - If `Type` is set to `BaseStatistics`, specify the following fields: `TagField`, `TagFieldSeparator`, `KVField`, `KVFieldSeparator`, `KVPairSeparator`, `TextField`, `Quantiles`, `DefaultValueOfString`, `NullStringField`, and `PartitionFieldFormat`.
	//
	// - If `Type` is set to `AbnormalBehavior`, specify the following fields: `UserId`, `ItemId`, `EventField`, `UpStreamBehavior`, `DownstreamBehavior`, `NumericHistogramBins`, and `PartitionFieldFormat`.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "AnalysisField": "userid", "PartitionFieldFormat": "yyyymmdd" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The scheduled time to run the task. If this parameter is omitted, the task runs only once.
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
	// The partition field for the left data table.
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
	// The partition field for the right data table.
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
	// The number of top results to return.
	//
	// example:
	//
	// 10
	TopNQuantity *int64 `json:"TopNQuantity,omitempty" xml:"TopNQuantity,omitempty"`
	// The type of the data diagnosis task. Valid values:
	//
	// - ChangeRate: Item or user change rate analysis.
	//
	// - PreferenceStatisticsCycle: User preference statistics cycle analysis.
	//
	// - JoinTables: Two-table join analysis.
	//
	// - BaseStatistics: Basic statistical analysis.
	//
	// - AbnormalBehavior: Abnormal behavior analysis.
	//
	// This parameter is required.
	//
	// example:
	//
	// ChangeRate
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CreateDataDiagnosisRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateDataDiagnosisRequest) GetCycleTime() *string {
	return s.CycleTime
}

func (s *CreateDataDiagnosisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDataDiagnosisRequest) GetLeftTableMetaId() *string {
	return s.LeftTableMetaId
}

func (s *CreateDataDiagnosisRequest) GetLeftTablePartitionField() *string {
	return s.LeftTablePartitionField
}

func (s *CreateDataDiagnosisRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataDiagnosisRequest) GetPartitionField() *string {
	return s.PartitionField
}

func (s *CreateDataDiagnosisRequest) GetRightTableMetaId() *string {
	return s.RightTableMetaId
}

func (s *CreateDataDiagnosisRequest) GetRightTablePartitionField() *string {
	return s.RightTablePartitionField
}

func (s *CreateDataDiagnosisRequest) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *CreateDataDiagnosisRequest) GetTopNQuantity() *int64 {
	return s.TopNQuantity
}

func (s *CreateDataDiagnosisRequest) GetType() *string {
	return s.Type
}

func (s *CreateDataDiagnosisRequest) SetConfig(v string) *CreateDataDiagnosisRequest {
	s.Config = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetCycleTime(v string) *CreateDataDiagnosisRequest {
	s.CycleTime = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetInstanceId(v string) *CreateDataDiagnosisRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetLeftTableMetaId(v string) *CreateDataDiagnosisRequest {
	s.LeftTableMetaId = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetLeftTablePartitionField(v string) *CreateDataDiagnosisRequest {
	s.LeftTablePartitionField = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetName(v string) *CreateDataDiagnosisRequest {
	s.Name = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetPartitionField(v string) *CreateDataDiagnosisRequest {
	s.PartitionField = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetRightTableMetaId(v string) *CreateDataDiagnosisRequest {
	s.RightTableMetaId = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetRightTablePartitionField(v string) *CreateDataDiagnosisRequest {
	s.RightTablePartitionField = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetTableMetaId(v string) *CreateDataDiagnosisRequest {
	s.TableMetaId = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetTopNQuantity(v int64) *CreateDataDiagnosisRequest {
	s.TopNQuantity = &v
	return s
}

func (s *CreateDataDiagnosisRequest) SetType(v string) *CreateDataDiagnosisRequest {
	s.Type = &v
	return s
}

func (s *CreateDataDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
