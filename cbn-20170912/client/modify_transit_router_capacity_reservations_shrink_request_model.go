// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterCapacityReservationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest
	GetClientToken() *string
	SetDeploymentSetsShrink(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest
	GetDeploymentSetsShrink() *string
	SetMinimumTransitRouterCapacitiesShrink(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest
	GetMinimumTransitRouterCapacitiesShrink() *string
	SetResetCapacityReservationsShrink(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest
	GetResetCapacityReservationsShrink() *string
	SetTransitRouterId(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest
	GetTransitRouterId() *string
	SetVersion(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest
	GetVersion() *string
}

type ModifyTransitRouterCapacityReservationsShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the token on your client to ensure that the token is unique across different requests. The client token only supports ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID of the API request as the client token. The request ID may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The deployment set information for the capacity reservation.
	DeploymentSetsShrink *string `json:"DeploymentSets,omitempty" xml:"DeploymentSets,omitempty"`
	// The reserved capacity for each zone in the capacity reservation.
	MinimumTransitRouterCapacitiesShrink *string `json:"MinimumTransitRouterCapacities,omitempty" xml:"MinimumTransitRouterCapacities,omitempty"`
	// Cancels capacity reservations.
	ResetCapacityReservationsShrink *string `json:"ResetCapacityReservations,omitempty" xml:"ResetCapacityReservations,omitempty"`
	// The ID of the transit router instance for forward routing.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ModifyTransitRouterCapacityReservationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCapacityReservationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) GetDeploymentSetsShrink() *string {
	return s.DeploymentSetsShrink
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) GetMinimumTransitRouterCapacitiesShrink() *string {
	return s.MinimumTransitRouterCapacitiesShrink
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) GetResetCapacityReservationsShrink() *string {
	return s.ResetCapacityReservationsShrink
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) GetVersion() *string {
	return s.Version
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) SetClientToken(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) SetDeploymentSetsShrink(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest {
	s.DeploymentSetsShrink = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) SetMinimumTransitRouterCapacitiesShrink(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest {
	s.MinimumTransitRouterCapacitiesShrink = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) SetResetCapacityReservationsShrink(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest {
	s.ResetCapacityReservationsShrink = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) SetTransitRouterId(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) SetVersion(v string) *ModifyTransitRouterCapacityReservationsShrinkRequest {
	s.Version = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
