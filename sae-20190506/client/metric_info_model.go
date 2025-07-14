// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAverage(v float32) *MetricInfo
	GetAverage() *float32
	SetCount(v float32) *MetricInfo
	GetCount() *float32
	SetMaximum(v float32) *MetricInfo
	GetMaximum() *float32
	SetMinimum(v float32) *MetricInfo
	GetMinimum() *float32
	SetSum(v float32) *MetricInfo
	GetSum() *float32
	SetTimestamp(v int64) *MetricInfo
	GetTimestamp() *int64
	SetValue(v float32) *MetricInfo
	GetValue() *float32
}

type MetricInfo struct {
	// example:
	//
	// 1234.5
	Average *float32 `json:"Average,omitempty" xml:"Average,omitempty"`
	// example:
	//
	// 1234.5
	Count *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1234.5
	Maximum *float32 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// example:
	//
	// 1234.5
	Minimum *float32 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// example:
	//
	// 1234.5
	Sum *float32 `json:"Sum,omitempty" xml:"Sum,omitempty"`
	// example:
	//
	// 1686568800000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// 1234.5
	Value *float32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MetricInfo) String() string {
	return dara.Prettify(s)
}

func (s MetricInfo) GoString() string {
	return s.String()
}

func (s *MetricInfo) GetAverage() *float32 {
	return s.Average
}

func (s *MetricInfo) GetCount() *float32 {
	return s.Count
}

func (s *MetricInfo) GetMaximum() *float32 {
	return s.Maximum
}

func (s *MetricInfo) GetMinimum() *float32 {
	return s.Minimum
}

func (s *MetricInfo) GetSum() *float32 {
	return s.Sum
}

func (s *MetricInfo) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *MetricInfo) GetValue() *float32 {
	return s.Value
}

func (s *MetricInfo) SetAverage(v float32) *MetricInfo {
	s.Average = &v
	return s
}

func (s *MetricInfo) SetCount(v float32) *MetricInfo {
	s.Count = &v
	return s
}

func (s *MetricInfo) SetMaximum(v float32) *MetricInfo {
	s.Maximum = &v
	return s
}

func (s *MetricInfo) SetMinimum(v float32) *MetricInfo {
	s.Minimum = &v
	return s
}

func (s *MetricInfo) SetSum(v float32) *MetricInfo {
	s.Sum = &v
	return s
}

func (s *MetricInfo) SetTimestamp(v int64) *MetricInfo {
	s.Timestamp = &v
	return s
}

func (s *MetricInfo) SetValue(v float32) *MetricInfo {
	s.Value = &v
	return s
}

func (s *MetricInfo) Validate() error {
	return dara.Validate(s)
}
