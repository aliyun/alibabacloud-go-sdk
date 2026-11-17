// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewayCapacityReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyNatGatewayCapacityReservationRequest
	GetClientToken() *string
	SetDryRun(v string) *ModifyNatGatewayCapacityReservationRequest
	GetDryRun() *string
	SetNatGatewayCapacity(v *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity) *ModifyNatGatewayCapacityReservationRequest
	GetNatGatewayCapacity() *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity
	SetNatGatewayId(v string) *ModifyNatGatewayCapacityReservationRequest
	GetNatGatewayId() *string
	SetRegionId(v string) *ModifyNatGatewayCapacityReservationRequest
	GetRegionId() *string
	SetResetCapacityReservation(v string) *ModifyNatGatewayCapacityReservationRequest
	GetResetCapacityReservation() *string
}

type ModifyNatGatewayCapacityReservationRequest struct {
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
	NatGatewayCapacity *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity `json:"NatGatewayCapacity,omitempty" xml:"NatGatewayCapacity,omitempty" type:"Struct"`
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

func (s ModifyNatGatewayCapacityReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayCapacityReservationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNatGatewayCapacityReservationRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *ModifyNatGatewayCapacityReservationRequest) GetNatGatewayCapacity() *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity {
	return s.NatGatewayCapacity
}

func (s *ModifyNatGatewayCapacityReservationRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatGatewayCapacityReservationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNatGatewayCapacityReservationRequest) GetResetCapacityReservation() *string {
	return s.ResetCapacityReservation
}

func (s *ModifyNatGatewayCapacityReservationRequest) SetClientToken(v string) *ModifyNatGatewayCapacityReservationRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationRequest) SetDryRun(v string) *ModifyNatGatewayCapacityReservationRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationRequest) SetNatGatewayCapacity(v *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity) *ModifyNatGatewayCapacityReservationRequest {
	s.NatGatewayCapacity = v
	return s
}

func (s *ModifyNatGatewayCapacityReservationRequest) SetNatGatewayId(v string) *ModifyNatGatewayCapacityReservationRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationRequest) SetRegionId(v string) *ModifyNatGatewayCapacityReservationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationRequest) SetResetCapacityReservation(v string) *ModifyNatGatewayCapacityReservationRequest {
	s.ResetCapacityReservation = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationRequest) Validate() error {
	if s.NatGatewayCapacity != nil {
		if err := s.NatGatewayCapacity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity struct {
	// The number of LCUs reserved per hour.
	//
	// example:
	//
	// 100
	CapacityUnits *float64 `json:"CapacityUnits,omitempty" xml:"CapacityUnits,omitempty"`
}

func (s ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity) GetCapacityUnits() *float64 {
	return s.CapacityUnits
}

func (s *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity) SetCapacityUnits(v float64) *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity {
	s.CapacityUnits = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationRequestNatGatewayCapacity) Validate() error {
	return dara.Validate(s)
}
