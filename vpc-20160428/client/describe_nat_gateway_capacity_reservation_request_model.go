// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewayCapacityReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeNatGatewayCapacityReservationRequest
	GetClientToken() *string
	SetDryRun(v string) *DescribeNatGatewayCapacityReservationRequest
	GetDryRun() *string
	SetNatGatewayId(v string) *DescribeNatGatewayCapacityReservationRequest
	GetNatGatewayId() *string
	SetRegionId(v string) *DescribeNatGatewayCapacityReservationRequest
	GetRegionId() *string
}

type DescribeNatGatewayCapacityReservationRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, and you must ensure that the value is unique among different requests. The client token can contain only ASCII characters. Note: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request is different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for this request. Valid values:
	//
	// - true: sends a dry run request without querying the resource status. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the dry run, the corresponding error is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - false (default): sends a Normal request. If the request passes the dry run, a 2xx HTTP status code is returned and the resource status is queried.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NAT gateway to query.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNatGatewayCapacityReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayCapacityReservationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeNatGatewayCapacityReservationRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *DescribeNatGatewayCapacityReservationRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewayCapacityReservationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNatGatewayCapacityReservationRequest) SetClientToken(v string) *DescribeNatGatewayCapacityReservationRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationRequest) SetDryRun(v string) *DescribeNatGatewayCapacityReservationRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationRequest) SetNatGatewayId(v string) *DescribeNatGatewayCapacityReservationRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationRequest) SetRegionId(v string) *DescribeNatGatewayCapacityReservationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationRequest) Validate() error {
	return dara.Validate(s)
}
