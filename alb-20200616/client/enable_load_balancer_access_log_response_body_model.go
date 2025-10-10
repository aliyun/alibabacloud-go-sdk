// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLoadBalancerAccessLogResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetJobId(v string) *EnableLoadBalancerAccessLogResponseBody
  GetJobId() *string 
  SetRequestId(v string) *EnableLoadBalancerAccessLogResponseBody
  GetRequestId() *string 
}

type EnableLoadBalancerAccessLogResponseBody struct {
  // The ID of the asynchronous job.
  // 
  // example:
  // 
  // ff7713ca-5818-4120-85e3-0bf9fr******
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 593B0448-D13E-4C56-AC0D-FDF0FD******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLoadBalancerAccessLogResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableLoadBalancerAccessLogResponseBody) GoString() string {
  return s.String()
}

func (s *EnableLoadBalancerAccessLogResponseBody) GetJobId() *string  {
  return s.JobId
}

func (s *EnableLoadBalancerAccessLogResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableLoadBalancerAccessLogResponseBody) SetJobId(v string) *EnableLoadBalancerAccessLogResponseBody {
  s.JobId = &v
  return s
}

func (s *EnableLoadBalancerAccessLogResponseBody) SetRequestId(v string) *EnableLoadBalancerAccessLogResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableLoadBalancerAccessLogResponseBody) Validate() error {
  return dara.Validate(s)
}

