// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoMValues interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentValue(v int64) *MoMValues
	GetCurrentValue() *int64
	SetLastDayValue(v int64) *MoMValues
	GetLastDayValue() *int64
	SetLastMonthValue(v int64) *MoMValues
	GetLastMonthValue() *int64
}

type MoMValues struct {
	// total
	CurrentValue *int64 `json:"currentValue,omitempty" xml:"currentValue,omitempty"`
	// daily addition
	LastDayValue *int64 `json:"lastDayValue,omitempty" xml:"lastDayValue,omitempty"`
	// monthly addition
	LastMonthValue *int64 `json:"lastMonthValue,omitempty" xml:"lastMonthValue,omitempty"`
}

func (s MoMValues) String() string {
	return dara.Prettify(s)
}

func (s MoMValues) GoString() string {
	return s.String()
}

func (s *MoMValues) GetCurrentValue() *int64 {
	return s.CurrentValue
}

func (s *MoMValues) GetLastDayValue() *int64 {
	return s.LastDayValue
}

func (s *MoMValues) GetLastMonthValue() *int64 {
	return s.LastMonthValue
}

func (s *MoMValues) SetCurrentValue(v int64) *MoMValues {
	s.CurrentValue = &v
	return s
}

func (s *MoMValues) SetLastDayValue(v int64) *MoMValues {
	s.LastDayValue = &v
	return s
}

func (s *MoMValues) SetLastMonthValue(v int64) *MoMValues {
	s.LastMonthValue = &v
	return s
}

func (s *MoMValues) Validate() error {
	return dara.Validate(s)
}
