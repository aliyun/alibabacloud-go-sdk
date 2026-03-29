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
	// This parameter is required.
	//
	// example:
	//
	// { "AnalysisField": "userid", "PartitionFieldFormat": "yyyymmdd" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 08:00
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 4
	LeftTableMetaId *string `json:"LeftTableMetaId,omitempty" xml:"LeftTableMetaId,omitempty"`
	// example:
	//
	// dt
	LeftTablePartitionField *string `json:"LeftTablePartitionField,omitempty" xml:"LeftTablePartitionField,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// data_diagnosis_job1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// dt
	PartitionField *string `json:"PartitionField,omitempty" xml:"PartitionField,omitempty"`
	// example:
	//
	// 5
	RightTableMetaId *string `json:"RightTableMetaId,omitempty" xml:"RightTableMetaId,omitempty"`
	// example:
	//
	// dt
	RightTablePartitionField *string `json:"RightTablePartitionField,omitempty" xml:"RightTablePartitionField,omitempty"`
	// example:
	//
	// 3
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// 10
	TopNQuantity *int64 `json:"TopNQuantity,omitempty" xml:"TopNQuantity,omitempty"`
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
