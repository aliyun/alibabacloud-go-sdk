// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLoadBalancerIpv6InternetResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetJobId(v string) *EnableLoadBalancerIpv6InternetResponseBody
  GetJobId() *string 
  SetRequestId(v string) *EnableLoadBalancerIpv6InternetResponseBody
  GetRequestId() *string 
}

type EnableLoadBalancerIpv6InternetResponseBody struct {
  // The asynchronous task ID.
  // 
  // example:
  // 
  // 4a6e3ad4-ef08-4ab1-b332-fa621cfe****
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // BB920797-D70E-567F-8098-55A861DD7912
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLoadBalancerIpv6InternetResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableLoadBalancerIpv6InternetResponseBody) GoString() string {
  return s.String()
}

func (s *EnableLoadBalancerIpv6InternetResponseBody) GetJobId() *string  {
  return s.JobId
}

func (s *EnableLoadBalancerIpv6InternetResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableLoadBalancerIpv6InternetResponseBody) SetJobId(v string) *EnableLoadBalancerIpv6InternetResponseBody {
  s.JobId = &v
  return s
}

func (s *EnableLoadBalancerIpv6InternetResponseBody) SetRequestId(v string) *EnableLoadBalancerIpv6InternetResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableLoadBalancerIpv6InternetResponseBody) Validate() error {
  return dara.Validate(s)
}

