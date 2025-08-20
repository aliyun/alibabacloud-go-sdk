// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSQLInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosisSQLInfo(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetDiagnosisSQLInfo() *string
	SetRequestId(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetRequestId() *string
	SetStageInfos(v []*DescribeDiagnosisSQLInfoResponseBodyStageInfos) *DescribeDiagnosisSQLInfoResponseBody
	GetStageInfos() []*DescribeDiagnosisSQLInfoResponseBodyStageInfos
}

type DescribeDiagnosisSQLInfoResponseBody struct {
	// The queried execution information, including the SQL statement, statistics, execution plan, and operator information.
	DiagnosisSQLInfo *string `json:"DiagnosisSQLInfo,omitempty" xml:"DiagnosisSQLInfo,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried execution information by stage.
	StageInfos []*DescribeDiagnosisSQLInfoResponseBodyStageInfos `json:"StageInfos,omitempty" xml:"StageInfos,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisSQLInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetDiagnosisSQLInfo() *string {
	return s.DiagnosisSQLInfo
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetStageInfos() []*DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	return s.StageInfos
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDiagnosisSQLInfo(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.DiagnosisSQLInfo = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetRequestId(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStageInfos(v []*DescribeDiagnosisSQLInfoResponseBodyStageInfos) *DescribeDiagnosisSQLInfoResponseBody {
	s.StageInfos = v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDiagnosisSQLInfoResponseBodyStageInfos struct {
	ExecutionType *string `json:"ExecutionType,omitempty" xml:"ExecutionType,omitempty"`
	// The total amount of input data in the stage. Unit: bytes.
	//
	// example:
	//
	// 2341
	InputDataSize *int64 `json:"InputDataSize,omitempty" xml:"InputDataSize,omitempty"`
	// The total number of input rows in the stage.
	//
	// example:
	//
	// 2341
	InputRows *int64 `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	// The total amount of time consumed by all operators in the stage. Unit: milliseconds.
	//
	// example:
	//
	// 2341
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The total amount of output data in the stage. Unit: bytes.
	//
	// example:
	//
	// 2341
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The total number of output rows in the stage.
	//
	// example:
	//
	// 2341
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The total peak memory of the stage. Unit: bytes.
	//
	// example:
	//
	// 2341
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The execution progress of the stage.
	//
	// example:
	//
	// 0.3
	Progress *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The stage ID.
	//
	// example:
	//
	// Stage[26]
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The state of the stage.
	//
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDiagnosisSQLInfoResponseBodyStageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBodyStageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetInputDataSize() *int64 {
	return s.InputDataSize
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetInputRows() *int64 {
	return s.InputRows
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetOperatorCost() *int64 {
	return s.OperatorCost
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetOutputDataSize() *int64 {
	return s.OutputDataSize
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetOutputRows() *int64 {
	return s.OutputRows
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetProgress() *float64 {
	return s.Progress
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetStageId() *string {
	return s.StageId
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) GetState() *string {
	return s.State
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetExecutionType(v string) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.ExecutionType = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetInputDataSize(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.InputDataSize = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetInputRows(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.InputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOperatorCost(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OperatorCost = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOutputDataSize(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOutputRows(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetPeakMemory(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetProgress(v float64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.Progress = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetStageId(v string) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.StageId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetState(v string) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.State = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) Validate() error {
	return dara.Validate(s)
}
