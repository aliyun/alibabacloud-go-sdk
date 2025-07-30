// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodMetric interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*Metric) *PodMetric
	GetMetrics() []*Metric
	SetPodId(v string) *PodMetric
	GetPodId() *string
}

type PodMetric struct {
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// dlc-20210329110128-746bf7cl47pr8-worker-0
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s PodMetric) String() string {
	return dara.Prettify(s)
}

func (s PodMetric) GoString() string {
	return s.String()
}

func (s *PodMetric) GetMetrics() []*Metric {
	return s.Metrics
}

func (s *PodMetric) GetPodId() *string {
	return s.PodId
}

func (s *PodMetric) SetMetrics(v []*Metric) *PodMetric {
	s.Metrics = v
	return s
}

func (s *PodMetric) SetPodId(v string) *PodMetric {
	s.PodId = &v
	return s
}

func (s *PodMetric) Validate() error {
	return dara.Validate(s)
}
