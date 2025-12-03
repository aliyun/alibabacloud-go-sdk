// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceMetricsResponseBody
	GetInstanceId() *string
	SetPodMetrics(v []*GetInstanceMetricsResponseBodyPodMetrics) *GetInstanceMetricsResponseBody
	GetPodMetrics() []*GetInstanceMetricsResponseBodyPodMetrics
	SetRequestId(v string) *GetInstanceMetricsResponseBody
	GetRequestId() *string
}

type GetInstanceMetricsResponseBody struct {
	// example:
	//
	// eais-bj8b53it29hfhj******
	InstanceId *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PodMetrics []*GetInstanceMetricsResponseBodyPodMetrics `json:"PodMetrics,omitempty" xml:"PodMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceMetricsResponseBody) GetPodMetrics() []*GetInstanceMetricsResponseBodyPodMetrics {
	return s.PodMetrics
}

func (s *GetInstanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceMetricsResponseBody) SetInstanceId(v string) *GetInstanceMetricsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetPodMetrics(v []*GetInstanceMetricsResponseBodyPodMetrics) *GetInstanceMetricsResponseBody {
	s.PodMetrics = v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetRequestId(v string) *GetInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) Validate() error {
	if s.PodMetrics != nil {
		for _, item := range s.PodMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceMetricsResponseBodyPodMetrics struct {
	Metrics []*GetInstanceMetricsResponseBodyPodMetricsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// Pod ID。
	//
	// example:
	//
	// eais-hznzre6ffmz9num4****-579b587ddf-9txr6
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s GetInstanceMetricsResponseBodyPodMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMetricsResponseBodyPodMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) GetMetrics() []*GetInstanceMetricsResponseBodyPodMetricsMetrics {
	return s.Metrics
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) GetPodId() *string {
	return s.PodId
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) SetMetrics(v []*GetInstanceMetricsResponseBodyPodMetricsMetrics) *GetInstanceMetricsResponseBodyPodMetrics {
	s.Metrics = v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) SetPodId(v string) *GetInstanceMetricsResponseBodyPodMetrics {
	s.PodId = &v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) Validate() error {
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

type GetInstanceMetricsResponseBodyPodMetricsMetrics struct {
	// example:
	//
	// 1669107528450
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 4.536552540058814
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceMetricsResponseBodyPodMetricsMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMetricsResponseBodyPodMetricsMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) GetTime() *string {
	return s.Time
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) GetValue() *string {
	return s.Value
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) SetTime(v string) *GetInstanceMetricsResponseBodyPodMetricsMetrics {
	s.Time = &v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) SetValue(v string) *GetInstanceMetricsResponseBodyPodMetricsMetrics {
	s.Value = &v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) Validate() error {
	return dara.Validate(s)
}
