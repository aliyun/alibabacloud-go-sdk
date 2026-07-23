// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiagnoses(v []*ListDataDiagnosesResponseBodyDataDiagnoses) *ListDataDiagnosesResponseBody
	GetDataDiagnoses() []*ListDataDiagnosesResponseBodyDataDiagnoses
	SetRequestId(v string) *ListDataDiagnosesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDataDiagnosesResponseBody
	GetTotalCount() *int64
}

type ListDataDiagnosesResponseBody struct {
	// The list of data diagnoses.
	DataDiagnoses []*ListDataDiagnosesResponseBodyDataDiagnoses `json:"DataDiagnoses,omitempty" xml:"DataDiagnoses,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataDiagnosesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosesResponseBody) GetDataDiagnoses() []*ListDataDiagnosesResponseBodyDataDiagnoses {
	return s.DataDiagnoses
}

func (s *ListDataDiagnosesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataDiagnosesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataDiagnosesResponseBody) SetDataDiagnoses(v []*ListDataDiagnosesResponseBodyDataDiagnoses) *ListDataDiagnosesResponseBody {
	s.DataDiagnoses = v
	return s
}

func (s *ListDataDiagnosesResponseBody) SetRequestId(v string) *ListDataDiagnosesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataDiagnosesResponseBody) SetTotalCount(v int64) *ListDataDiagnosesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataDiagnosesResponseBody) Validate() error {
	if s.DataDiagnoses != nil {
		for _, item := range s.DataDiagnoses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataDiagnosesResponseBodyDataDiagnoses struct {
	// The configuration for the data diagnosis task, in JSON format. The required fields depend on the `Type` value:<br>
	//
	// example:
	//
	// {"AnalysisField": "userid","PartitionFieldFormat": "yyyymmdd"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The time at which the task is scheduled to run periodically. If this parameter is empty, the task runs only once.
	//
	// example:
	//
	// 08:00
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The data diagnosis ID.
	//
	// example:
	//
	// 3
	DataDiagnosisId *string `json:"DataDiagnosisId,omitempty" xml:"DataDiagnosisId,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the task was last updated.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
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
	// The name of the data diagnosis.
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
	// The data table ID.
	//
	// example:
	//
	// 3
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// The name of the data table.
	//
	// example:
	//
	// table_meta_1
	TableMetaName *string `json:"TableMetaName,omitempty" xml:"TableMetaName,omitempty"`
	// The number of top results to return.
	//
	// example:
	//
	// 10
	TopNQuantity *int64 `json:"TopNQuantity,omitempty" xml:"TopNQuantity,omitempty"`
	// The type of data diagnosis. Valid values:
	//
	// - `ChangeRate`: Change Rate Analysis.
	//
	// - `PreferenceStatisticsCycle`: Preference Statistics Cycle Analysis.
	//
	// - `JoinTables`: Join Tables Analysis.
	//
	// - `BaseStatistics`: Base Statistics Analysis.
	//
	// - `AbnormalBehavior`: Abnormal Behavior Analysis.
	//
	// example:
	//
	// ChangeRate
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataDiagnosesResponseBodyDataDiagnoses) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosesResponseBodyDataDiagnoses) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetConfig() *string {
	return s.Config
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetCycleTime() *string {
	return s.CycleTime
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetDataDiagnosisId() *string {
	return s.DataDiagnosisId
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetLeftTableMetaId() *string {
	return s.LeftTableMetaId
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetLeftTablePartitionField() *string {
	return s.LeftTablePartitionField
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetName() *string {
	return s.Name
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetPartitionField() *string {
	return s.PartitionField
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetRightTableMetaId() *string {
	return s.RightTableMetaId
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetRightTablePartitionField() *string {
	return s.RightTablePartitionField
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetTableMetaName() *string {
	return s.TableMetaName
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetTopNQuantity() *int64 {
	return s.TopNQuantity
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) GetType() *string {
	return s.Type
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetConfig(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.Config = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetCycleTime(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.CycleTime = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetDataDiagnosisId(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.DataDiagnosisId = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetGmtCreateTime(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.GmtCreateTime = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetGmtModifiedTime(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetLeftTableMetaId(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.LeftTableMetaId = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetLeftTablePartitionField(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.LeftTablePartitionField = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetName(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.Name = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetPartitionField(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.PartitionField = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetRightTableMetaId(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.RightTableMetaId = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetRightTablePartitionField(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.RightTablePartitionField = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetTableMetaId(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.TableMetaId = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetTableMetaName(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.TableMetaName = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetTopNQuantity(v int64) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.TopNQuantity = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) SetType(v string) *ListDataDiagnosesResponseBodyDataDiagnoses {
	s.Type = &v
	return s
}

func (s *ListDataDiagnosesResponseBodyDataDiagnoses) Validate() error {
	return dara.Validate(s)
}
