// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlotMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*Metric) *QuerySlotMetricsResponseBody
	GetMetrics() []*Metric
	SetPeriod(v string) *QuerySlotMetricsResponseBody
	GetPeriod() *string
	SetRequestId(v string) *QuerySlotMetricsResponseBody
	GetRequestId() *string
}

type QuerySlotMetricsResponseBody struct {
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// 60s
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySlotMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySlotMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySlotMetricsResponseBody) GetMetrics() []*Metric {
	return s.Metrics
}

func (s *QuerySlotMetricsResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *QuerySlotMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySlotMetricsResponseBody) SetMetrics(v []*Metric) *QuerySlotMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *QuerySlotMetricsResponseBody) SetPeriod(v string) *QuerySlotMetricsResponseBody {
	s.Period = &v
	return s
}

func (s *QuerySlotMetricsResponseBody) SetRequestId(v string) *QuerySlotMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySlotMetricsResponseBody) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
