// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObservationChartsDTO interface {
	dara.Model
	String() string
	GoString() string
	SetCallVolume(v []*TimeSeriesPointDTO) *ObservationChartsDTO
	GetCallVolume() []*TimeSeriesPointDTO
	SetConcurrency(v []*TimeSeriesPointDTO) *ObservationChartsDTO
	GetConcurrency() []*TimeSeriesPointDTO
	SetQpm(v []*TimeSeriesPointDTO) *ObservationChartsDTO
	GetQpm() []*TimeSeriesPointDTO
	SetResponseTime(v []*TimeSeriesPointDTO) *ObservationChartsDTO
	GetResponseTime() []*TimeSeriesPointDTO
	SetSuccessRate(v []*ObservationChartsDTOSuccessRate) *ObservationChartsDTO
	GetSuccessRate() []*ObservationChartsDTOSuccessRate
	SetTpm(v []*TimeSeriesPointDTO) *ObservationChartsDTO
	GetTpm() []*TimeSeriesPointDTO
}

type ObservationChartsDTO struct {
	CallVolume   []*TimeSeriesPointDTO              `json:"callVolume,omitempty" xml:"callVolume,omitempty" type:"Repeated"`
	Concurrency  []*TimeSeriesPointDTO              `json:"concurrency,omitempty" xml:"concurrency,omitempty" type:"Repeated"`
	Qpm          []*TimeSeriesPointDTO              `json:"qpm,omitempty" xml:"qpm,omitempty" type:"Repeated"`
	ResponseTime []*TimeSeriesPointDTO              `json:"responseTime,omitempty" xml:"responseTime,omitempty" type:"Repeated"`
	SuccessRate  []*ObservationChartsDTOSuccessRate `json:"successRate,omitempty" xml:"successRate,omitempty" type:"Repeated"`
	Tpm          []*TimeSeriesPointDTO              `json:"tpm,omitempty" xml:"tpm,omitempty" type:"Repeated"`
}

func (s ObservationChartsDTO) String() string {
	return dara.Prettify(s)
}

func (s ObservationChartsDTO) GoString() string {
	return s.String()
}

func (s *ObservationChartsDTO) GetCallVolume() []*TimeSeriesPointDTO {
	return s.CallVolume
}

func (s *ObservationChartsDTO) GetConcurrency() []*TimeSeriesPointDTO {
	return s.Concurrency
}

func (s *ObservationChartsDTO) GetQpm() []*TimeSeriesPointDTO {
	return s.Qpm
}

func (s *ObservationChartsDTO) GetResponseTime() []*TimeSeriesPointDTO {
	return s.ResponseTime
}

func (s *ObservationChartsDTO) GetSuccessRate() []*ObservationChartsDTOSuccessRate {
	return s.SuccessRate
}

func (s *ObservationChartsDTO) GetTpm() []*TimeSeriesPointDTO {
	return s.Tpm
}

func (s *ObservationChartsDTO) SetCallVolume(v []*TimeSeriesPointDTO) *ObservationChartsDTO {
	s.CallVolume = v
	return s
}

func (s *ObservationChartsDTO) SetConcurrency(v []*TimeSeriesPointDTO) *ObservationChartsDTO {
	s.Concurrency = v
	return s
}

func (s *ObservationChartsDTO) SetQpm(v []*TimeSeriesPointDTO) *ObservationChartsDTO {
	s.Qpm = v
	return s
}

func (s *ObservationChartsDTO) SetResponseTime(v []*TimeSeriesPointDTO) *ObservationChartsDTO {
	s.ResponseTime = v
	return s
}

func (s *ObservationChartsDTO) SetSuccessRate(v []*ObservationChartsDTOSuccessRate) *ObservationChartsDTO {
	s.SuccessRate = v
	return s
}

func (s *ObservationChartsDTO) SetTpm(v []*TimeSeriesPointDTO) *ObservationChartsDTO {
	s.Tpm = v
	return s
}

func (s *ObservationChartsDTO) Validate() error {
	if s.CallVolume != nil {
		for _, item := range s.CallVolume {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Concurrency != nil {
		for _, item := range s.Concurrency {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Qpm != nil {
		for _, item := range s.Qpm {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResponseTime != nil {
		for _, item := range s.ResponseTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessRate != nil {
		for _, item := range s.SuccessRate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tpm != nil {
		for _, item := range s.Tpm {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ObservationChartsDTOSuccessRate struct {
	// example:
	//
	// series1
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	Timestamp *string     `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Value     interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ObservationChartsDTOSuccessRate) String() string {
	return dara.Prettify(s)
}

func (s ObservationChartsDTOSuccessRate) GoString() string {
	return s.String()
}

func (s *ObservationChartsDTOSuccessRate) GetLabel() *string {
	return s.Label
}

func (s *ObservationChartsDTOSuccessRate) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ObservationChartsDTOSuccessRate) GetValue() interface{} {
	return s.Value
}

func (s *ObservationChartsDTOSuccessRate) SetLabel(v string) *ObservationChartsDTOSuccessRate {
	s.Label = &v
	return s
}

func (s *ObservationChartsDTOSuccessRate) SetTimestamp(v string) *ObservationChartsDTOSuccessRate {
	s.Timestamp = &v
	return s
}

func (s *ObservationChartsDTOSuccessRate) SetValue(v interface{}) *ObservationChartsDTOSuccessRate {
	s.Value = v
	return s
}

func (s *ObservationChartsDTOSuccessRate) Validate() error {
	return dara.Validate(s)
}
