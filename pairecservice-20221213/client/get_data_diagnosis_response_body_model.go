// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetDataDiagnosisResponseBody
	GetConfig() *string
	SetCycleTime(v string) *GetDataDiagnosisResponseBody
	GetCycleTime() *string
	SetGmtCreateTime(v string) *GetDataDiagnosisResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetDataDiagnosisResponseBody
	GetGmtModifiedTime() *string
	SetLeftTableMetaId(v string) *GetDataDiagnosisResponseBody
	GetLeftTableMetaId() *string
	SetLeftTablePartitionField(v string) *GetDataDiagnosisResponseBody
	GetLeftTablePartitionField() *string
	SetName(v string) *GetDataDiagnosisResponseBody
	GetName() *string
	SetPartitionField(v string) *GetDataDiagnosisResponseBody
	GetPartitionField() *string
	SetRequestId(v string) *GetDataDiagnosisResponseBody
	GetRequestId() *string
	SetRightTableMetaId(v string) *GetDataDiagnosisResponseBody
	GetRightTableMetaId() *string
	SetRightTablePartitionField(v string) *GetDataDiagnosisResponseBody
	GetRightTablePartitionField() *string
	SetTableMetaId(v string) *GetDataDiagnosisResponseBody
	GetTableMetaId() *string
	SetTableMetaName(v string) *GetDataDiagnosisResponseBody
	GetTableMetaName() *string
	SetTopNQuantity(v int64) *GetDataDiagnosisResponseBody
	GetTopNQuantity() *int64
	SetType(v string) *GetDataDiagnosisResponseBody
	GetType() *string
}

type GetDataDiagnosisResponseBody struct {
	// example:
	//
	// {"AnalysisField":"userid","PartitionFieldFormat":"yyyymmdd"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 08:00
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 4
	LeftTableMetaId *string `json:"LeftTableMetaId,omitempty" xml:"LeftTableMetaId,omitempty"`
	// example:
	//
	// dt
	LeftTablePartitionField *string `json:"LeftTablePartitionField,omitempty" xml:"LeftTablePartitionField,omitempty"`
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
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// table_meta_1
	TableMetaName *string `json:"TableMetaName,omitempty" xml:"TableMetaName,omitempty"`
	// example:
	//
	// 10
	TopNQuantity *int64 `json:"TopNQuantity,omitempty" xml:"TopNQuantity,omitempty"`
	// example:
	//
	// ChangeRate
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataDiagnosisResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetDataDiagnosisResponseBody) GetCycleTime() *string {
	return s.CycleTime
}

func (s *GetDataDiagnosisResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetDataDiagnosisResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetDataDiagnosisResponseBody) GetLeftTableMetaId() *string {
	return s.LeftTableMetaId
}

func (s *GetDataDiagnosisResponseBody) GetLeftTablePartitionField() *string {
	return s.LeftTablePartitionField
}

func (s *GetDataDiagnosisResponseBody) GetName() *string {
	return s.Name
}

func (s *GetDataDiagnosisResponseBody) GetPartitionField() *string {
	return s.PartitionField
}

func (s *GetDataDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataDiagnosisResponseBody) GetRightTableMetaId() *string {
	return s.RightTableMetaId
}

func (s *GetDataDiagnosisResponseBody) GetRightTablePartitionField() *string {
	return s.RightTablePartitionField
}

func (s *GetDataDiagnosisResponseBody) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *GetDataDiagnosisResponseBody) GetTableMetaName() *string {
	return s.TableMetaName
}

func (s *GetDataDiagnosisResponseBody) GetTopNQuantity() *int64 {
	return s.TopNQuantity
}

func (s *GetDataDiagnosisResponseBody) GetType() *string {
	return s.Type
}

func (s *GetDataDiagnosisResponseBody) SetConfig(v string) *GetDataDiagnosisResponseBody {
	s.Config = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetCycleTime(v string) *GetDataDiagnosisResponseBody {
	s.CycleTime = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetGmtCreateTime(v string) *GetDataDiagnosisResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetGmtModifiedTime(v string) *GetDataDiagnosisResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetLeftTableMetaId(v string) *GetDataDiagnosisResponseBody {
	s.LeftTableMetaId = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetLeftTablePartitionField(v string) *GetDataDiagnosisResponseBody {
	s.LeftTablePartitionField = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetName(v string) *GetDataDiagnosisResponseBody {
	s.Name = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetPartitionField(v string) *GetDataDiagnosisResponseBody {
	s.PartitionField = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetRequestId(v string) *GetDataDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetRightTableMetaId(v string) *GetDataDiagnosisResponseBody {
	s.RightTableMetaId = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetRightTablePartitionField(v string) *GetDataDiagnosisResponseBody {
	s.RightTablePartitionField = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetTableMetaId(v string) *GetDataDiagnosisResponseBody {
	s.TableMetaId = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetTableMetaName(v string) *GetDataDiagnosisResponseBody {
	s.TableMetaName = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetTopNQuantity(v int64) *GetDataDiagnosisResponseBody {
	s.TopNQuantity = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) SetType(v string) *GetDataDiagnosisResponseBody {
	s.Type = &v
	return s
}

func (s *GetDataDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
