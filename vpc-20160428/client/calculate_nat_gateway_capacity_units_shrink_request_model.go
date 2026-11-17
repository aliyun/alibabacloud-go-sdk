// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCalculateNatGatewayCapacityUnitsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest
	GetClientToken() *string
	SetDryRun(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest
	GetDryRun() *string
	SetNatGatewayCapacityShrink(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest
	GetNatGatewayCapacityShrink() *string
	SetNatGatewayId(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest
	GetNatGatewayId() *string
	SetNatGatewayPerformanceValuesShrink(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest
	GetNatGatewayPerformanceValuesShrink() *string
	SetRegionId(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest
	GetRegionId() *string
}

type CalculateNatGatewayCapacityUnitsShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. Note: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: sends the request without upgrading the Internet NAT gateway. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the check, an error message is returned. If the request passes the check, the DryRunOperation error code is returned.
	//
	// - false (default value): sends a normal request. If the request passes the check, a 2xx HTTP status code is returned and the Internet NAT gateway is directly upgraded.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The details of the reserved capacity units (CUs) to be calculated.
	NatGatewayCapacityShrink *string `json:"NatGatewayCapacity,omitempty" xml:"NatGatewayCapacity,omitempty"`
	// The instance ID of the NAT gateway.
	//
	// example:
	//
	// ngw-2ze0dcn4mq31qx2jc****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The details of the resource reservation metrics to be calculated.
	NatGatewayPerformanceValuesShrink *string `json:"NatGatewayPerformanceValues,omitempty" xml:"NatGatewayPerformanceValues,omitempty"`
	// The region ID of the NAT gateway.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CalculateNatGatewayCapacityUnitsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CalculateNatGatewayCapacityUnitsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) GetNatGatewayCapacityShrink() *string {
	return s.NatGatewayCapacityShrink
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) GetNatGatewayPerformanceValuesShrink() *string {
	return s.NatGatewayPerformanceValuesShrink
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) SetClientToken(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) SetDryRun(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) SetNatGatewayCapacityShrink(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest {
	s.NatGatewayCapacityShrink = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) SetNatGatewayId(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) SetNatGatewayPerformanceValuesShrink(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest {
	s.NatGatewayPerformanceValuesShrink = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) SetRegionId(v string) *CalculateNatGatewayCapacityUnitsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
