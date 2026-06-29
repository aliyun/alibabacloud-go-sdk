// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskStatistic interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptItemCount(v float32) *TaskStatistic
	GetAcceptItemCount() *float32
	SetCheckAbandon(v float32) *TaskStatistic
	GetCheckAbandon() *float32
	SetCheckAccuracy(v float32) *TaskStatistic
	GetCheckAccuracy() *float32
	SetCheckEfficiency(v float32) *TaskStatistic
	GetCheckEfficiency() *float32
	SetCheckedAccuracy(v float32) *TaskStatistic
	GetCheckedAccuracy() *float32
	SetCheckedError(v float32) *TaskStatistic
	GetCheckedError() *float32
	SetCheckedRejectCount(v float32) *TaskStatistic
	GetCheckedRejectCount() *float32
	SetFinalAbandonCount(v float32) *TaskStatistic
	GetFinalAbandonCount() *float32
	SetFinishedItemCount(v int64) *TaskStatistic
	GetFinishedItemCount() *int64
	SetFinishedSubtaskCount(v int64) *TaskStatistic
	GetFinishedSubtaskCount() *int64
	SetMarkEfficiency(v float32) *TaskStatistic
	GetMarkEfficiency() *float32
	SetPreMarkFixedCount(v float32) *TaskStatistic
	GetPreMarkFixedCount() *float32
	SetSampledAccuracy(v float32) *TaskStatistic
	GetSampledAccuracy() *float32
	SetSampledErrorCount(v float32) *TaskStatistic
	GetSampledErrorCount() *float32
	SetSampledRejectCount(v float32) *TaskStatistic
	GetSampledRejectCount() *float32
	SetSamplingAccuracy(v float32) *TaskStatistic
	GetSamplingAccuracy() *float32
	SetTotalCheckCount(v float32) *TaskStatistic
	GetTotalCheckCount() *float32
	SetTotalCheckTime(v float32) *TaskStatistic
	GetTotalCheckTime() *float32
	SetTotalCheckedCount(v float32) *TaskStatistic
	GetTotalCheckedCount() *float32
	SetTotalItemCount(v int64) *TaskStatistic
	GetTotalItemCount() *int64
	SetTotalMarkTime(v float32) *TaskStatistic
	GetTotalMarkTime() *float32
	SetTotalSampledCount(v float32) *TaskStatistic
	GetTotalSampledCount() *float32
	SetTotalSamplingCount(v float32) *TaskStatistic
	GetTotalSamplingCount() *float32
	SetTotalSubtaskCount(v int64) *TaskStatistic
	GetTotalSubtaskCount() *int64
	SetTotalWorkTime(v float32) *TaskStatistic
	GetTotalWorkTime() *float32
}

type TaskStatistic struct {
	// Data items that passed
	//
	// example:
	//
	// 0
	AcceptItemCount *float32 `json:"AcceptItemCount,omitempty" xml:"AcceptItemCount,omitempty"`
	// Quantity abandoned in the check flow
	//
	// example:
	//
	// 0
	CheckAbandon *float32 `json:"CheckAbandon,omitempty" xml:"CheckAbandon,omitempty"`
	// Inspection accuracy
	//
	// example:
	//
	// 0
	CheckAccuracy *float32 `json:"CheckAccuracy,omitempty" xml:"CheckAccuracy,omitempty"`
	// Inspection efficiency (items/hour)
	//
	// example:
	//
	// 0
	CheckEfficiency *float32 `json:"CheckEfficiency,omitempty" xml:"CheckEfficiency,omitempty"`
	// Check accuracy
	//
	// example:
	//
	// 0
	CheckedAccuracy *float32 `json:"CheckedAccuracy,omitempty" xml:"CheckedAccuracy,omitempty"`
	// Number of errors found in the inspection flow
	//
	// example:
	//
	// 0
	CheckedError *float32 `json:"CheckedError,omitempty" xml:"CheckedError,omitempty"`
	// Number of checks
	//
	// example:
	//
	// 0
	CheckedRejectCount *float32 `json:"CheckedRejectCount,omitempty" xml:"CheckedRejectCount,omitempty"`
	// Discarded data items
	//
	// example:
	//
	// 0
	FinalAbandonCount *float32 `json:"FinalAbandonCount,omitempty" xml:"FinalAbandonCount,omitempty"`
	// Completed data items
	//
	// example:
	//
	// 3
	FinishedItemCount *int64 `json:"FinishedItemCount,omitempty" xml:"FinishedItemCount,omitempty"`
	// Quantity of completed subtasks
	//
	// example:
	//
	// 3
	FinishedSubtaskCount *int64 `json:"FinishedSubtaskCount,omitempty" xml:"FinishedSubtaskCount,omitempty"`
	// Annotation efficiency (items/hour)
	//
	// example:
	//
	// 0
	MarkEfficiency *float32 `json:"MarkEfficiency,omitempty" xml:"MarkEfficiency,omitempty"`
	// Quantity of corrections made during pre-annotation
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0
	PreMarkFixedCount *float32 `json:"PreMarkFixedCount,omitempty" xml:"PreMarkFixedCount,omitempty"`
	// Sampling accuracy
	//
	// example:
	//
	// 0
	SampledAccuracy *float32 `json:"SampledAccuracy,omitempty" xml:"SampledAccuracy,omitempty"`
	// Number of sampled fault samples
	//
	// example:
	//
	// 0
	SampledErrorCount *float32 `json:"SampledErrorCount,omitempty" xml:"SampledErrorCount,omitempty"`
	// Number of samples denied
	//
	// example:
	//
	// 0
	SampledRejectCount *float32 `json:"SampledRejectCount,omitempty" xml:"SampledRejectCount,omitempty"`
	// Sampling accuracy
	//
	// example:
	//
	// 0
	SamplingAccuracy *float32 `json:"SamplingAccuracy,omitempty" xml:"SamplingAccuracy,omitempty"`
	// Total number of check flow steps
	//
	// example:
	//
	// 0
	TotalCheckCount *float32 `json:"TotalCheckCount,omitempty" xml:"TotalCheckCount,omitempty"`
	// Total check time (hours)
	//
	// example:
	//
	// 0
	TotalCheckTime *float32 `json:"TotalCheckTime,omitempty" xml:"TotalCheckTime,omitempty"`
	// Total number of checks
	//
	// example:
	//
	// 0
	TotalCheckedCount *float32 `json:"TotalCheckedCount,omitempty" xml:"TotalCheckedCount,omitempty"`
	// Total number of data items
	//
	// example:
	//
	// 3
	TotalItemCount *int64 `json:"TotalItemCount,omitempty" xml:"TotalItemCount,omitempty"`
	// Total time spent in the annotation phase (hours)
	//
	// example:
	//
	// 0
	TotalMarkTime *float32 `json:"TotalMarkTime,omitempty" xml:"TotalMarkTime,omitempty"`
	// Total sampling quantity
	//
	// example:
	//
	// 0
	TotalSampledCount *float32 `json:"TotalSampledCount,omitempty" xml:"TotalSampledCount,omitempty"`
	// Total number of sampled validations
	//
	// example:
	//
	// 0
	TotalSamplingCount *float32 `json:"TotalSamplingCount,omitempty" xml:"TotalSamplingCount,omitempty"`
	// Total number of subtasks
	//
	// example:
	//
	// 3
	TotalSubtaskCount *int64 `json:"TotalSubtaskCount,omitempty" xml:"TotalSubtaskCount,omitempty"`
	// Total work time (hours)
	//
	// example:
	//
	// 0
	TotalWorkTime *float32 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s TaskStatistic) String() string {
	return dara.Prettify(s)
}

func (s TaskStatistic) GoString() string {
	return s.String()
}

func (s *TaskStatistic) GetAcceptItemCount() *float32 {
	return s.AcceptItemCount
}

func (s *TaskStatistic) GetCheckAbandon() *float32 {
	return s.CheckAbandon
}

func (s *TaskStatistic) GetCheckAccuracy() *float32 {
	return s.CheckAccuracy
}

func (s *TaskStatistic) GetCheckEfficiency() *float32 {
	return s.CheckEfficiency
}

func (s *TaskStatistic) GetCheckedAccuracy() *float32 {
	return s.CheckedAccuracy
}

func (s *TaskStatistic) GetCheckedError() *float32 {
	return s.CheckedError
}

func (s *TaskStatistic) GetCheckedRejectCount() *float32 {
	return s.CheckedRejectCount
}

func (s *TaskStatistic) GetFinalAbandonCount() *float32 {
	return s.FinalAbandonCount
}

func (s *TaskStatistic) GetFinishedItemCount() *int64 {
	return s.FinishedItemCount
}

func (s *TaskStatistic) GetFinishedSubtaskCount() *int64 {
	return s.FinishedSubtaskCount
}

func (s *TaskStatistic) GetMarkEfficiency() *float32 {
	return s.MarkEfficiency
}

func (s *TaskStatistic) GetPreMarkFixedCount() *float32 {
	return s.PreMarkFixedCount
}

func (s *TaskStatistic) GetSampledAccuracy() *float32 {
	return s.SampledAccuracy
}

func (s *TaskStatistic) GetSampledErrorCount() *float32 {
	return s.SampledErrorCount
}

func (s *TaskStatistic) GetSampledRejectCount() *float32 {
	return s.SampledRejectCount
}

func (s *TaskStatistic) GetSamplingAccuracy() *float32 {
	return s.SamplingAccuracy
}

func (s *TaskStatistic) GetTotalCheckCount() *float32 {
	return s.TotalCheckCount
}

func (s *TaskStatistic) GetTotalCheckTime() *float32 {
	return s.TotalCheckTime
}

func (s *TaskStatistic) GetTotalCheckedCount() *float32 {
	return s.TotalCheckedCount
}

func (s *TaskStatistic) GetTotalItemCount() *int64 {
	return s.TotalItemCount
}

func (s *TaskStatistic) GetTotalMarkTime() *float32 {
	return s.TotalMarkTime
}

func (s *TaskStatistic) GetTotalSampledCount() *float32 {
	return s.TotalSampledCount
}

func (s *TaskStatistic) GetTotalSamplingCount() *float32 {
	return s.TotalSamplingCount
}

func (s *TaskStatistic) GetTotalSubtaskCount() *int64 {
	return s.TotalSubtaskCount
}

func (s *TaskStatistic) GetTotalWorkTime() *float32 {
	return s.TotalWorkTime
}

func (s *TaskStatistic) SetAcceptItemCount(v float32) *TaskStatistic {
	s.AcceptItemCount = &v
	return s
}

func (s *TaskStatistic) SetCheckAbandon(v float32) *TaskStatistic {
	s.CheckAbandon = &v
	return s
}

func (s *TaskStatistic) SetCheckAccuracy(v float32) *TaskStatistic {
	s.CheckAccuracy = &v
	return s
}

func (s *TaskStatistic) SetCheckEfficiency(v float32) *TaskStatistic {
	s.CheckEfficiency = &v
	return s
}

func (s *TaskStatistic) SetCheckedAccuracy(v float32) *TaskStatistic {
	s.CheckedAccuracy = &v
	return s
}

func (s *TaskStatistic) SetCheckedError(v float32) *TaskStatistic {
	s.CheckedError = &v
	return s
}

func (s *TaskStatistic) SetCheckedRejectCount(v float32) *TaskStatistic {
	s.CheckedRejectCount = &v
	return s
}

func (s *TaskStatistic) SetFinalAbandonCount(v float32) *TaskStatistic {
	s.FinalAbandonCount = &v
	return s
}

func (s *TaskStatistic) SetFinishedItemCount(v int64) *TaskStatistic {
	s.FinishedItemCount = &v
	return s
}

func (s *TaskStatistic) SetFinishedSubtaskCount(v int64) *TaskStatistic {
	s.FinishedSubtaskCount = &v
	return s
}

func (s *TaskStatistic) SetMarkEfficiency(v float32) *TaskStatistic {
	s.MarkEfficiency = &v
	return s
}

func (s *TaskStatistic) SetPreMarkFixedCount(v float32) *TaskStatistic {
	s.PreMarkFixedCount = &v
	return s
}

func (s *TaskStatistic) SetSampledAccuracy(v float32) *TaskStatistic {
	s.SampledAccuracy = &v
	return s
}

func (s *TaskStatistic) SetSampledErrorCount(v float32) *TaskStatistic {
	s.SampledErrorCount = &v
	return s
}

func (s *TaskStatistic) SetSampledRejectCount(v float32) *TaskStatistic {
	s.SampledRejectCount = &v
	return s
}

func (s *TaskStatistic) SetSamplingAccuracy(v float32) *TaskStatistic {
	s.SamplingAccuracy = &v
	return s
}

func (s *TaskStatistic) SetTotalCheckCount(v float32) *TaskStatistic {
	s.TotalCheckCount = &v
	return s
}

func (s *TaskStatistic) SetTotalCheckTime(v float32) *TaskStatistic {
	s.TotalCheckTime = &v
	return s
}

func (s *TaskStatistic) SetTotalCheckedCount(v float32) *TaskStatistic {
	s.TotalCheckedCount = &v
	return s
}

func (s *TaskStatistic) SetTotalItemCount(v int64) *TaskStatistic {
	s.TotalItemCount = &v
	return s
}

func (s *TaskStatistic) SetTotalMarkTime(v float32) *TaskStatistic {
	s.TotalMarkTime = &v
	return s
}

func (s *TaskStatistic) SetTotalSampledCount(v float32) *TaskStatistic {
	s.TotalSampledCount = &v
	return s
}

func (s *TaskStatistic) SetTotalSamplingCount(v float32) *TaskStatistic {
	s.TotalSamplingCount = &v
	return s
}

func (s *TaskStatistic) SetTotalSubtaskCount(v int64) *TaskStatistic {
	s.TotalSubtaskCount = &v
	return s
}

func (s *TaskStatistic) SetTotalWorkTime(v float32) *TaskStatistic {
	s.TotalWorkTime = &v
	return s
}

func (s *TaskStatistic) Validate() error {
	return dara.Validate(s)
}
