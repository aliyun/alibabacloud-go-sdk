// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetricDropRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateMetricDropRequest
	GetClusterId() *string
	SetMetricDrop(v string) *UpdateMetricDropRequest
	GetMetricDrop() *string
	SetRegionId(v string) *UpdateMetricDropRequest
	GetRegionId() *string
}

type UpdateMetricDropRequest struct {
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// c3ca36c8e2693403d85c0d9f8bb1d7b6c
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of discarded metrics. Specify one metric name in each line.
	//
	// example:
	//
	// apiserver_request_duration_seconds_bucket
	//
	// etcd_request_duration_seconds_bucket
	//
	// apiserver_request_total
	//
	// container_tasks_state
	MetricDrop *string `json:"MetricDrop,omitempty" xml:"MetricDrop,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateMetricDropRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetricDropRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetricDropRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateMetricDropRequest) GetMetricDrop() *string {
	return s.MetricDrop
}

func (s *UpdateMetricDropRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateMetricDropRequest) SetClusterId(v string) *UpdateMetricDropRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateMetricDropRequest) SetMetricDrop(v string) *UpdateMetricDropRequest {
	s.MetricDrop = &v
	return s
}

func (s *UpdateMetricDropRequest) SetRegionId(v string) *UpdateMetricDropRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateMetricDropRequest) Validate() error {
	return dara.Validate(s)
}
