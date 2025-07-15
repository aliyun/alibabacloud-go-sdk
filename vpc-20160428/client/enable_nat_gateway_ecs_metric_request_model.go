// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNatGatewayEcsMetricRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDryRun(v bool) *EnableNatGatewayEcsMetricRequest
  GetDryRun() *bool 
  SetNatGatewayId(v string) *EnableNatGatewayEcsMetricRequest
  GetNatGatewayId() *string 
  SetRegionId(v string) *EnableNatGatewayEcsMetricRequest
  GetRegionId() *string 
}

type EnableNatGatewayEcsMetricRequest struct {
  // Specifies whether to perform only a dry run, without performing the actual request. Valid values:
  // 
  // **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
  // 
  // **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
  // 
  // example:
  // 
  // false
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // The ID of the NAT gateway for which you want to enable ECS traffic monitoring.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ngw-2vc53wynunp35lw1y****
  NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
  // The region ID of the NAT gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-chengdu
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableNatGatewayEcsMetricRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableNatGatewayEcsMetricRequest) GoString() string {
  return s.String()
}

func (s *EnableNatGatewayEcsMetricRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableNatGatewayEcsMetricRequest) GetNatGatewayId() *string  {
  return s.NatGatewayId
}

func (s *EnableNatGatewayEcsMetricRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableNatGatewayEcsMetricRequest) SetDryRun(v bool) *EnableNatGatewayEcsMetricRequest {
  s.DryRun = &v
  return s
}

func (s *EnableNatGatewayEcsMetricRequest) SetNatGatewayId(v string) *EnableNatGatewayEcsMetricRequest {
  s.NatGatewayId = &v
  return s
}

func (s *EnableNatGatewayEcsMetricRequest) SetRegionId(v string) *EnableNatGatewayEcsMetricRequest {
  s.RegionId = &v
  return s
}

func (s *EnableNatGatewayEcsMetricRequest) Validate() error {
  return dara.Validate(s)
}

