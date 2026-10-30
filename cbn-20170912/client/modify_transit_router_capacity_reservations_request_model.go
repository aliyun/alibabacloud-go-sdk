// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterCapacityReservationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyTransitRouterCapacityReservationsRequest
	GetClientToken() *string
	SetDeploymentSets(v []*ModifyTransitRouterCapacityReservationsRequestDeploymentSets) *ModifyTransitRouterCapacityReservationsRequest
	GetDeploymentSets() []*ModifyTransitRouterCapacityReservationsRequestDeploymentSets
	SetMinimumTransitRouterCapacities(v []*ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) *ModifyTransitRouterCapacityReservationsRequest
	GetMinimumTransitRouterCapacities() []*ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities
	SetResetCapacityReservations(v []*ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations) *ModifyTransitRouterCapacityReservationsRequest
	GetResetCapacityReservations() []*ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations
	SetTransitRouterId(v string) *ModifyTransitRouterCapacityReservationsRequest
	GetTransitRouterId() *string
	SetVersion(v string) *ModifyTransitRouterCapacityReservationsRequest
	GetVersion() *string
}

type ModifyTransitRouterCapacityReservationsRequest struct {
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
	DeploymentSets []*ModifyTransitRouterCapacityReservationsRequestDeploymentSets `json:"DeploymentSets,omitempty" xml:"DeploymentSets,omitempty" type:"Repeated"`
	// The reserved capacity for each zone in the capacity reservation.
	MinimumTransitRouterCapacities []*ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities `json:"MinimumTransitRouterCapacities,omitempty" xml:"MinimumTransitRouterCapacities,omitempty" type:"Repeated"`
	// Cancels capacity reservations.
	ResetCapacityReservations []*ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations `json:"ResetCapacityReservations,omitempty" xml:"ResetCapacityReservations,omitempty" type:"Repeated"`
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

func (s ModifyTransitRouterCapacityReservationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCapacityReservationsRequest) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCapacityReservationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyTransitRouterCapacityReservationsRequest) GetDeploymentSets() []*ModifyTransitRouterCapacityReservationsRequestDeploymentSets {
	return s.DeploymentSets
}

func (s *ModifyTransitRouterCapacityReservationsRequest) GetMinimumTransitRouterCapacities() []*ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities {
	return s.MinimumTransitRouterCapacities
}

func (s *ModifyTransitRouterCapacityReservationsRequest) GetResetCapacityReservations() []*ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations {
	return s.ResetCapacityReservations
}

func (s *ModifyTransitRouterCapacityReservationsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ModifyTransitRouterCapacityReservationsRequest) GetVersion() *string {
	return s.Version
}

func (s *ModifyTransitRouterCapacityReservationsRequest) SetClientToken(v string) *ModifyTransitRouterCapacityReservationsRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequest) SetDeploymentSets(v []*ModifyTransitRouterCapacityReservationsRequestDeploymentSets) *ModifyTransitRouterCapacityReservationsRequest {
	s.DeploymentSets = v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequest) SetMinimumTransitRouterCapacities(v []*ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) *ModifyTransitRouterCapacityReservationsRequest {
	s.MinimumTransitRouterCapacities = v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequest) SetResetCapacityReservations(v []*ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations) *ModifyTransitRouterCapacityReservationsRequest {
	s.ResetCapacityReservations = v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequest) SetTransitRouterId(v string) *ModifyTransitRouterCapacityReservationsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequest) SetVersion(v string) *ModifyTransitRouterCapacityReservationsRequest {
	s.Version = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequest) Validate() error {
	if s.DeploymentSets != nil {
		for _, item := range s.DeploymentSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MinimumTransitRouterCapacities != nil {
		for _, item := range s.MinimumTransitRouterCapacities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResetCapacityReservations != nil {
		for _, item := range s.ResetCapacityReservations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyTransitRouterCapacityReservationsRequestDeploymentSets struct {
	// The ID of the deployment set instance used.
	//
	// example:
	//
	// ds-badxcs*****
	DeploymentSetID *string `json:"DeploymentSetID,omitempty" xml:"DeploymentSetID,omitempty"`
	// The zone for which to modify the deployment set of the capacity reservation.
	//
	// - To cancel the deployment set, specify the zone and leave the deployment set ID empty.
	//
	// example:
	//
	// cn-shanghai-b
	ZoneID *string `json:"ZoneID,omitempty" xml:"ZoneID,omitempty"`
}

func (s ModifyTransitRouterCapacityReservationsRequestDeploymentSets) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCapacityReservationsRequestDeploymentSets) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCapacityReservationsRequestDeploymentSets) GetDeploymentSetID() *string {
	return s.DeploymentSetID
}

func (s *ModifyTransitRouterCapacityReservationsRequestDeploymentSets) GetZoneID() *string {
	return s.ZoneID
}

func (s *ModifyTransitRouterCapacityReservationsRequestDeploymentSets) SetDeploymentSetID(v string) *ModifyTransitRouterCapacityReservationsRequestDeploymentSets {
	s.DeploymentSetID = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequestDeploymentSets) SetZoneID(v string) *ModifyTransitRouterCapacityReservationsRequestDeploymentSets {
	s.ZoneID = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequestDeploymentSets) Validate() error {
	return dara.Validate(s)
}

type ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities struct {
	// The reserved capacity for the specified zone.
	DataProcessGbps *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps `json:"DataProcessGbps,omitempty" xml:"DataProcessGbps,omitempty" type:"Struct"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-shanghai-l
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) GetDataProcessGbps() *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps {
	return s.DataProcessGbps
}

func (s *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) SetDataProcessGbps(v *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps) *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities {
	s.DataProcessGbps = v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) SetZoneId(v string) *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities {
	s.ZoneId = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacities) Validate() error {
	if s.DataProcessGbps != nil {
		if err := s.DataProcessGbps.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps struct {
	// The target reserved capacity for the specified zone.
	//
	// example:
	//
	// 100
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps) GetTarget() *int64 {
	return s.Target
}

func (s *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps) SetTarget(v int64) *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps {
	s.Target = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequestMinimumTransitRouterCapacitiesDataProcessGbps) Validate() error {
	return dara.Validate(s)
}

type ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations struct {
	// The ID of the zone for which to cancel the capacity reservation.
	//
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations) SetZoneId(v string) *ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations {
	s.ZoneId = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsRequestResetCapacityReservations) Validate() error {
	return dara.Validate(s)
}
