// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcEndpointConnectionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBandwidth(v int32) *EnableVpcEndpointConnectionRequest
  GetBandwidth() *int32 
  SetClientToken(v string) *EnableVpcEndpointConnectionRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableVpcEndpointConnectionRequest
  GetDryRun() *bool 
  SetEndpointId(v string) *EnableVpcEndpointConnectionRequest
  GetEndpointId() *string 
  SetRegionId(v string) *EnableVpcEndpointConnectionRequest
  GetRegionId() *string 
  SetServiceId(v string) *EnableVpcEndpointConnectionRequest
  GetServiceId() *string 
  SetTrafficControlMode(v string) *EnableVpcEndpointConnectionRequest
  GetTrafficControlMode() *string 
}

type EnableVpcEndpointConnectionRequest struct {
  // The bandwidth of the endpoint connection. Unit: Mbit/s. Valid values: **3072 to 10240**.
  // 
  // >  The bandwidth of an endpoint connection is in the range of **100 to 10,240*	- Mbit/s. The default bandwidth is **3,072*	- Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is **3,072 to 10,240*	- Mbit/s. If Classic Load Balancer (CLB) instances or Application Load Balancer (ALB) instances are specified as service resources, you can modify the endpoint connection bandwidth based on your business requirements. This parameter is invalid if Network Load Balancer (NLB) instances are specified as service resources.
  // 
  // example:
  // 
  // 1024
  Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
  // 
  // example:
  // 
  // 0c593ea1-3bea-11e9-b96b-88e9fe637760
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // Specifies whether to perform only a dry run, without performing the actual request. Valid values:
  // 
  // 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
  // 
  // 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
  // 
  // example:
  // 
  // false
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // The endpoint ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ep-hp33b2e43fays7s8****
  EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
  // The ID of the region where the connection request is accepted.
  // 
  // You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-huhehaote
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The endpoint service ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // epsrv-hp3vpx8yqxblby3i****
  ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
  TrafficControlMode *string `json:"TrafficControlMode,omitempty" xml:"TrafficControlMode,omitempty"`
}

func (s EnableVpcEndpointConnectionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcEndpointConnectionRequest) GoString() string {
  return s.String()
}

func (s *EnableVpcEndpointConnectionRequest) GetBandwidth() *int32  {
  return s.Bandwidth
}

func (s *EnableVpcEndpointConnectionRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableVpcEndpointConnectionRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableVpcEndpointConnectionRequest) GetEndpointId() *string  {
  return s.EndpointId
}

func (s *EnableVpcEndpointConnectionRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableVpcEndpointConnectionRequest) GetServiceId() *string  {
  return s.ServiceId
}

func (s *EnableVpcEndpointConnectionRequest) GetTrafficControlMode() *string  {
  return s.TrafficControlMode
}

func (s *EnableVpcEndpointConnectionRequest) SetBandwidth(v int32) *EnableVpcEndpointConnectionRequest {
  s.Bandwidth = &v
  return s
}

func (s *EnableVpcEndpointConnectionRequest) SetClientToken(v string) *EnableVpcEndpointConnectionRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableVpcEndpointConnectionRequest) SetDryRun(v bool) *EnableVpcEndpointConnectionRequest {
  s.DryRun = &v
  return s
}

func (s *EnableVpcEndpointConnectionRequest) SetEndpointId(v string) *EnableVpcEndpointConnectionRequest {
  s.EndpointId = &v
  return s
}

func (s *EnableVpcEndpointConnectionRequest) SetRegionId(v string) *EnableVpcEndpointConnectionRequest {
  s.RegionId = &v
  return s
}

func (s *EnableVpcEndpointConnectionRequest) SetServiceId(v string) *EnableVpcEndpointConnectionRequest {
  s.ServiceId = &v
  return s
}

func (s *EnableVpcEndpointConnectionRequest) SetTrafficControlMode(v string) *EnableVpcEndpointConnectionRequest {
  s.TrafficControlMode = &v
  return s
}

func (s *EnableVpcEndpointConnectionRequest) Validate() error {
  return dara.Validate(s)
}

