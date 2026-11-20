// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimeSeriesPointDTO interface {
	dara.Model
	String() string
	GoString() string
	SetLabel(v string) *TimeSeriesPointDTO
	GetLabel() *string
	SetTimestamp(v string) *TimeSeriesPointDTO
	GetTimestamp() *string
	SetValue(v interface{}) *TimeSeriesPointDTO
	GetValue() interface{}
}

type TimeSeriesPointDTO struct {
	// The label.
	//
	// example:
	//
	// series1
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The timestamp (RFC3339).
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// The value.
	//
	// example:
	//
	// 42.76
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TimeSeriesPointDTO) String() string {
	return dara.Prettify(s)
}

func (s TimeSeriesPointDTO) GoString() string {
	return s.String()
}

func (s *TimeSeriesPointDTO) GetLabel() *string {
	return s.Label
}

func (s *TimeSeriesPointDTO) GetTimestamp() *string {
	return s.Timestamp
}

func (s *TimeSeriesPointDTO) GetValue() interface{} {
	return s.Value
}

func (s *TimeSeriesPointDTO) SetLabel(v string) *TimeSeriesPointDTO {
	s.Label = &v
	return s
}

func (s *TimeSeriesPointDTO) SetTimestamp(v string) *TimeSeriesPointDTO {
	s.Timestamp = &v
	return s
}

func (s *TimeSeriesPointDTO) SetValue(v interface{}) *TimeSeriesPointDTO {
	s.Value = v
	return s
}

func (s *TimeSeriesPointDTO) Validate() error {
	return dara.Validate(s)
}
