// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewayCapacityReservationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyNatGatewayCapacityReservationShrinkRequest
	GetClientToken() *string
	SetDryRun(v string) *ModifyNatGatewayCapacityReservationShrinkRequest
	GetDryRun() *string
	SetNatGatewayCapacityShrink(v string) *ModifyNatGatewayCapacityReservationShrinkRequest
	GetNatGatewayCapacityShrink() *string
	SetNatGatewayId(v string) *ModifyNatGatewayCapacityReservationShrinkRequest
	GetNatGatewayId() *string
	SetRegionId(v string) *ModifyNatGatewayCapacityReservationShrinkRequest
	GetRegionId() *string
	SetResetCapacityReservation(v string) *ModifyNatGatewayCapacityReservationShrinkRequest
	GetResetCapacityReservation() *string
}

type ModifyNatGatewayCapacityReservationShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters. Note: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: Performs a dry run without upgrading the Internet NAT gateway. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - false (default): Sends a Normal request. If the request passes the dry run, a 2xx HTTP status code is returned and the Internet NAT gateway is upgraded.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The number of capacity units (CUs) for the NAT resource reservation.
	NatGatewayCapacityShrink *string `json:"NatGatewayCapacity,omitempty" xml:"NatGatewayCapacity,omitempty"`
	// The instance ID of the NAT gateway.
	//
	// example:
	//
	// ngw-112za33e4****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to cancel the resource reservation. Valid values:
	//
	// - **true**: Cancels the resource reservation.
	//
	// - **false*	- (default): Does not cancel the resource reservation. You do not need to specify this parameter when you create or modify a resource reservation.
	//
	// example:
	//
	// true
	ResetCapacityReservation *string `json:"ResetCapacityReservation,omitempty" xml:"ResetCapacityReservation,omitempty"`
}

func (s ModifyNatGatewayCapacityReservationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayCapacityReservationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) GetNatGatewayCapacityShrink() *string {
	return s.NatGatewayCapacityShrink
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) GetResetCapacityReservation() *string {
	return s.ResetCapacityReservation
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) SetClientToken(v string) *ModifyNatGatewayCapacityReservationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) SetDryRun(v string) *ModifyNatGatewayCapacityReservationShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) SetNatGatewayCapacityShrink(v string) *ModifyNatGatewayCapacityReservationShrinkRequest {
	s.NatGatewayCapacityShrink = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) SetNatGatewayId(v string) *ModifyNatGatewayCapacityReservationShrinkRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) SetRegionId(v string) *ModifyNatGatewayCapacityReservationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) SetResetCapacityReservation(v string) *ModifyNatGatewayCapacityReservationShrinkRequest {
	s.ResetCapacityReservation = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
