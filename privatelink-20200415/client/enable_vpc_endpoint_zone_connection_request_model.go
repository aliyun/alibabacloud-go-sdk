// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcEndpointZoneConnectionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableVpcEndpointZoneConnectionRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableVpcEndpointZoneConnectionRequest
  GetDryRun() *bool 
  SetEndpointId(v string) *EnableVpcEndpointZoneConnectionRequest
  GetEndpointId() *string 
  SetRegionId(v string) *EnableVpcEndpointZoneConnectionRequest
  GetRegionId() *string 
  SetServiceId(v string) *EnableVpcEndpointZoneConnectionRequest
  GetServiceId() *string 
  SetZoneId(v string) *EnableVpcEndpointZoneConnectionRequest
  GetZoneId() *string 
}

type EnableVpcEndpointZoneConnectionRequest struct {
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
  // 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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
  // The ID of the region where the endpoint connection request is accepted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The endpoint service ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // epsrv-hp3vpx8yqxblby3i****
  ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
  // The ID of the zone that is associated with the endpoint.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-b
  ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EnableVpcEndpointZoneConnectionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcEndpointZoneConnectionRequest) GoString() string {
  return s.String()
}

func (s *EnableVpcEndpointZoneConnectionRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableVpcEndpointZoneConnectionRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableVpcEndpointZoneConnectionRequest) GetEndpointId() *string  {
  return s.EndpointId
}

func (s *EnableVpcEndpointZoneConnectionRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableVpcEndpointZoneConnectionRequest) GetServiceId() *string  {
  return s.ServiceId
}

func (s *EnableVpcEndpointZoneConnectionRequest) GetZoneId() *string  {
  return s.ZoneId
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetClientToken(v string) *EnableVpcEndpointZoneConnectionRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetDryRun(v bool) *EnableVpcEndpointZoneConnectionRequest {
  s.DryRun = &v
  return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetEndpointId(v string) *EnableVpcEndpointZoneConnectionRequest {
  s.EndpointId = &v
  return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetRegionId(v string) *EnableVpcEndpointZoneConnectionRequest {
  s.RegionId = &v
  return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetServiceId(v string) *EnableVpcEndpointZoneConnectionRequest {
  s.ServiceId = &v
  return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) SetZoneId(v string) *EnableVpcEndpointZoneConnectionRequest {
  s.ZoneId = &v
  return s
}

func (s *EnableVpcEndpointZoneConnectionRequest) Validate() error {
  return dara.Validate(s)
}

