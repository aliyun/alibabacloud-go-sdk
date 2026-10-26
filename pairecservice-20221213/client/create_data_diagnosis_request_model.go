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
	// The configuration content of the data modeling diagnostic task.
	//
	// - If Type is set to ChangeRate, the configuration must include: analysis field (AnalysisField) and partition field format (PartitionFieldFormat).
	//
	// - If Type is set to PreferenceStatisticsCycle, the configuration must include: user ID field (UserIdField), single-day retention rate statistical period (RemainDays), ever-appeared rate statistical period (EverAppearedDays), periodic retention rate statistical period (RemainRatePeriods), and partition field format (PartitionFieldFormat).
	//
	// - If Type is set to JoinTables, the configuration must include: left table analysis field (LeftTableAnalysisField), right table analysis field (RightTableAnalysisField), left table JOIN field (LeftJoinField), right table JOIN field (RightJoinField), display sample quantity (SampleQuantity), left table partition field format (LeftTablePartitionFieldFormat), and right table partition field format (RightTablePartitionFieldFormat).
	//
	// - If Type is set to BaseStatistics, the configuration must include: tag field (TagField), tag field separator (TagFieldSeparator), KV field (KVField), KV field separator (KVFieldSeparator), KV pair separator (KVPairSeparator), text field (TextField), data percentile distribution (Quantiles), default value for string type (DefaultValueOfString), null string field for string type (NullStringField), and partition field format (PartitionFieldFormat).
	//
	// - If Type is set to AbnormalBehavior, the configuration must include: user ID field (UserId), item ID field (ItemId), behavior field (EventField), upstream behavior value (UpStreamBehavior), downstream behavior value (DownstreamBehavior), number of histogram bins (NumericHistogramBins), and partition field format (PartitionFieldFormat).
	//
	// This parameter is required.
	//
	// example:
	//
	// { "AnalysisField": "userid", "PartitionFieldFormat": "yyyymmdd" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The periodic run time. If this field is empty, the task does not run periodically. The format is HH:mm.
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
	// The name of the data modeling diagnostic task configuration.
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
	// The number of top N results.
	//
	// example:
	//
	// 10
	TopNQuantity *int64 `json:"TopNQuantity,omitempty" xml:"TopNQuantity,omitempty"`
	// The type of the data diagnostic task. Valid values:
	//
	// - ChangeRate: item/user change rate analysis.
	//
	//
	//
	// - PreferenceStatisticsCycle: user preference statistical period analysis.
	//
	//
	//
	// - JoinTables: join table analysis.
	//
	//
	//
	// - BaseStatistics: basic statistics analysis.
	//
	//
	//
	// - AbnormalBehavior: abnormal behavior analysis.
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
