// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLoadBalancerAccessLogRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableLoadBalancerAccessLogRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableLoadBalancerAccessLogRequest
  GetDryRun() *bool 
  SetLoadBalancerId(v string) *EnableLoadBalancerAccessLogRequest
  GetLoadBalancerId() *string 
  SetLogProject(v string) *EnableLoadBalancerAccessLogRequest
  GetLogProject() *string 
  SetLogStore(v string) *EnableLoadBalancerAccessLogRequest
  GetLogStore() *string 
}

type EnableLoadBalancerAccessLogRequest struct {
  // The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
  // 
  // >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
  // 
  // example:
  // 
  // 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // Specifies whether to perform only a dry run, without performing the actual request. Valid values:
  // 
  // 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
  // 
  // 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
  // 
  // example:
  // 
  // true
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // The ALB instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // alb-bd6oylbckp6k9x****
  LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
  // The project to which the access log is shipped.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // sls-setter
  LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
  // The Logstore to which the access log is shipped.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test
  LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s EnableLoadBalancerAccessLogRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableLoadBalancerAccessLogRequest) GoString() string {
  return s.String()
}

func (s *EnableLoadBalancerAccessLogRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableLoadBalancerAccessLogRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableLoadBalancerAccessLogRequest) GetLoadBalancerId() *string  {
  return s.LoadBalancerId
}

func (s *EnableLoadBalancerAccessLogRequest) GetLogProject() *string  {
  return s.LogProject
}

func (s *EnableLoadBalancerAccessLogRequest) GetLogStore() *string  {
  return s.LogStore
}

func (s *EnableLoadBalancerAccessLogRequest) SetClientToken(v string) *EnableLoadBalancerAccessLogRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetDryRun(v bool) *EnableLoadBalancerAccessLogRequest {
  s.DryRun = &v
  return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLoadBalancerId(v string) *EnableLoadBalancerAccessLogRequest {
  s.LoadBalancerId = &v
  return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLogProject(v string) *EnableLoadBalancerAccessLogRequest {
  s.LogProject = &v
  return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLogStore(v string) *EnableLoadBalancerAccessLogRequest {
  s.LogStore = &v
  return s
}

func (s *EnableLoadBalancerAccessLogRequest) Validate() error {
  return dara.Validate(s)
}

