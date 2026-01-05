// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*Metric) *QueryInstanceMetricsResponseBody
	GetMetrics() []*Metric
	SetPeriod(v string) *QueryInstanceMetricsResponseBody
	GetPeriod() *string
	SetRequestId(v string) *QueryInstanceMetricsResponseBody
	GetRequestId() *string
}

type QueryInstanceMetricsResponseBody struct {
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

func (s QueryInstanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceMetricsResponseBody) GetMetrics() []*Metric {
	return s.Metrics
}

func (s *QueryInstanceMetricsResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *QueryInstanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstanceMetricsResponseBody) SetMetrics(v []*Metric) *QueryInstanceMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *QueryInstanceMetricsResponseBody) SetPeriod(v string) *QueryInstanceMetricsResponseBody {
	s.Period = &v
	return s
}

func (s *QueryInstanceMetricsResponseBody) SetRequestId(v string) *QueryInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceMetricsResponseBody) Validate() error {
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
