// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNatGatewayEcsMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DisableNatGatewayEcsMetricRequest
	GetDryRun() *bool
	SetNatGatewayId(v string) *DisableNatGatewayEcsMetricRequest
	GetNatGatewayId() *string
	SetRegionId(v string) *DisableNatGatewayEcsMetricRequest
	GetRegionId() *string
}

type DisableNatGatewayEcsMetricRequest struct {
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
	// The ID of the NAT gateway for which you want to disable ECS traffic monitoring.
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

func (s DisableNatGatewayEcsMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableNatGatewayEcsMetricRequest) GoString() string {
	return s.String()
}

func (s *DisableNatGatewayEcsMetricRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableNatGatewayEcsMetricRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DisableNatGatewayEcsMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableNatGatewayEcsMetricRequest) SetDryRun(v bool) *DisableNatGatewayEcsMetricRequest {
	s.DryRun = &v
	return s
}

func (s *DisableNatGatewayEcsMetricRequest) SetNatGatewayId(v string) *DisableNatGatewayEcsMetricRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DisableNatGatewayEcsMetricRequest) SetRegionId(v string) *DisableNatGatewayEcsMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DisableNatGatewayEcsMetricRequest) Validate() error {
	return dara.Validate(s)
}
