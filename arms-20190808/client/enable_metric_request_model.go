// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClusterId(v string) *EnableMetricRequest
  GetClusterId() *string 
  SetDropMetric(v string) *EnableMetricRequest
  GetDropMetric() *string 
  SetRegionId(v string) *EnableMetricRequest
  GetRegionId() *string 
}

type EnableMetricRequest struct {
  // The cluster ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ccfa5e34a5c1f4ce6b916a40a12151d88
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // The metric name.
  // 
  // example:
  // 
  // kube_pod_container_status_ready
  DropMetric *string `json:"DropMetric,omitempty" xml:"DropMetric,omitempty"`
  // The region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableMetricRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableMetricRequest) GoString() string {
  return s.String()
}

func (s *EnableMetricRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *EnableMetricRequest) GetDropMetric() *string  {
  return s.DropMetric
}

func (s *EnableMetricRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableMetricRequest) SetClusterId(v string) *EnableMetricRequest {
  s.ClusterId = &v
  return s
}

func (s *EnableMetricRequest) SetDropMetric(v string) *EnableMetricRequest {
  s.DropMetric = &v
  return s
}

func (s *EnableMetricRequest) SetRegionId(v string) *EnableMetricRequest {
  s.RegionId = &v
  return s
}

func (s *EnableMetricRequest) Validate() error {
  return dara.Validate(s)
}

