// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserStatistic interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptedMarkItemsCount(v float32) *UserStatistic
	GetAcceptedMarkItemsCount() *float32
	SetCheckCount(v float32) *UserStatistic
	GetCheckCount() *float32
	SetCheckedAcceptedCount(v float32) *UserStatistic
	GetCheckedAcceptedCount() *float32
	SetCheckedAccuracy(v float32) *UserStatistic
	GetCheckedAccuracy() *float32
	SetMarkEfficiency(v float32) *UserStatistic
	GetMarkEfficiency() *float32
	SetMarkTime(v float32) *UserStatistic
	GetMarkTime() *float32
	SetSamplingAccuracy(v float32) *UserStatistic
	GetSamplingAccuracy() *float32
	SetSamplingCount(v float32) *UserStatistic
	GetSamplingCount() *float32
	SetSamplingErrorCount(v float32) *UserStatistic
	GetSamplingErrorCount() *float32
	SetTotalMarkItemsCount(v float32) *UserStatistic
	GetTotalMarkItemsCount() *float32
	SetUserId(v string) *UserStatistic
	GetUserId() *string
}

type UserStatistic struct {
	// Quantity of Data items passed
	//
	// example:
	//
	// 172
	AcceptedMarkItemsCount *float32 `json:"AcceptedMarkItemsCount,omitempty" xml:"AcceptedMarkItemsCount,omitempty"`
	// Total inspection count
	//
	// example:
	//
	// 140
	CheckCount *float32 `json:"CheckCount,omitempty" xml:"CheckCount,omitempty"`
	// Quantity passed in inspection
	//
	// example:
	//
	// 100
	CheckedAcceptedCount *float32 `json:"CheckedAcceptedCount,omitempty" xml:"CheckedAcceptedCount,omitempty"`
	// Inspection accuracy.
	//
	// Inspection accuracy = Number Of Error inspected / Quantity inspected
	//
	// example:
	//
	// 95.33
	CheckedAccuracy *float32 `json:"CheckedAccuracy,omitempty" xml:"CheckedAccuracy,omitempty"`
	// Annotation efficiency. Unit: items/hour
	//
	// Annotation efficiency = Quantity annotated / Annotation duration (including rejections)
	//
	// example:
	//
	// 0.1
	MarkEfficiency *float32 `json:"MarkEfficiency,omitempty" xml:"MarkEfficiency,omitempty"`
	// Annotation duration. Unit: hours
	//
	// example:
	//
	// 0.1
	MarkTime *float32 `json:"MarkTime,omitempty" xml:"MarkTime,omitempty"`
	// Sampling accuracy.
	//
	// Validated accuracy = Number Of Error validated / Quantity validated
	//
	// example:
	//
	// 84.92
	SamplingAccuracy *float32 `json:"SamplingAccuracy,omitempty" xml:"SamplingAccuracy,omitempty"`
	// Total sampling quantity
	//
	// example:
	//
	// 1
	SamplingCount *float32 `json:"SamplingCount,omitempty" xml:"SamplingCount,omitempty"`
	// Number Of Error in sampling
	//
	// example:
	//
	// 1
	SamplingErrorCount *float32 `json:"SamplingErrorCount,omitempty" xml:"SamplingErrorCount,omitempty"`
	// Total Data items
	//
	// example:
	//
	// 172
	TotalMarkItemsCount *float32 `json:"TotalMarkItemsCount,omitempty" xml:"TotalMarkItemsCount,omitempty"`
	// User ID
	//
	// example:
	//
	// 166***9980757311
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UserStatistic) String() string {
	return dara.Prettify(s)
}

func (s UserStatistic) GoString() string {
	return s.String()
}

func (s *UserStatistic) GetAcceptedMarkItemsCount() *float32 {
	return s.AcceptedMarkItemsCount
}

func (s *UserStatistic) GetCheckCount() *float32 {
	return s.CheckCount
}

func (s *UserStatistic) GetCheckedAcceptedCount() *float32 {
	return s.CheckedAcceptedCount
}

func (s *UserStatistic) GetCheckedAccuracy() *float32 {
	return s.CheckedAccuracy
}

func (s *UserStatistic) GetMarkEfficiency() *float32 {
	return s.MarkEfficiency
}

func (s *UserStatistic) GetMarkTime() *float32 {
	return s.MarkTime
}

func (s *UserStatistic) GetSamplingAccuracy() *float32 {
	return s.SamplingAccuracy
}

func (s *UserStatistic) GetSamplingCount() *float32 {
	return s.SamplingCount
}

func (s *UserStatistic) GetSamplingErrorCount() *float32 {
	return s.SamplingErrorCount
}

func (s *UserStatistic) GetTotalMarkItemsCount() *float32 {
	return s.TotalMarkItemsCount
}

func (s *UserStatistic) GetUserId() *string {
	return s.UserId
}

func (s *UserStatistic) SetAcceptedMarkItemsCount(v float32) *UserStatistic {
	s.AcceptedMarkItemsCount = &v
	return s
}

func (s *UserStatistic) SetCheckCount(v float32) *UserStatistic {
	s.CheckCount = &v
	return s
}

func (s *UserStatistic) SetCheckedAcceptedCount(v float32) *UserStatistic {
	s.CheckedAcceptedCount = &v
	return s
}

func (s *UserStatistic) SetCheckedAccuracy(v float32) *UserStatistic {
	s.CheckedAccuracy = &v
	return s
}

func (s *UserStatistic) SetMarkEfficiency(v float32) *UserStatistic {
	s.MarkEfficiency = &v
	return s
}

func (s *UserStatistic) SetMarkTime(v float32) *UserStatistic {
	s.MarkTime = &v
	return s
}

func (s *UserStatistic) SetSamplingAccuracy(v float32) *UserStatistic {
	s.SamplingAccuracy = &v
	return s
}

func (s *UserStatistic) SetSamplingCount(v float32) *UserStatistic {
	s.SamplingCount = &v
	return s
}

func (s *UserStatistic) SetSamplingErrorCount(v float32) *UserStatistic {
	s.SamplingErrorCount = &v
	return s
}

func (s *UserStatistic) SetTotalMarkItemsCount(v float32) *UserStatistic {
	s.TotalMarkItemsCount = &v
	return s
}

func (s *UserStatistic) SetUserId(v string) *UserStatistic {
	s.UserId = &v
	return s
}

func (s *UserStatistic) Validate() error {
	return dara.Validate(s)
}
