// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDateSummary interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *DateSummary
	GetDate() *string
	SetValue(v int64) *DateSummary
	GetValue() *int64
}

type DateSummary struct {
	// example:
	//
	// 2025-06-01
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// Metric value at corresponding date
	//
	// example:
	//
	// 100
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DateSummary) String() string {
	return dara.Prettify(s)
}

func (s DateSummary) GoString() string {
	return s.String()
}

func (s *DateSummary) GetDate() *string {
	return s.Date
}

func (s *DateSummary) GetValue() *int64 {
	return s.Value
}

func (s *DateSummary) SetDate(v string) *DateSummary {
	s.Date = &v
	return s
}

func (s *DateSummary) SetValue(v int64) *DateSummary {
	s.Value = &v
	return s
}

func (s *DateSummary) Validate() error {
	return dara.Validate(s)
}
