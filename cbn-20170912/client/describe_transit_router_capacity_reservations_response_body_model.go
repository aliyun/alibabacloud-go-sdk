// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterCapacityReservationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSets(v []*DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) *DescribeTransitRouterCapacityReservationsResponseBody
	GetDeploymentSets() []*DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets
	SetLastModifiedTime(v string) *DescribeTransitRouterCapacityReservationsResponseBody
	GetLastModifiedTime() *string
	SetMinimumTransitRouterCapacities(v []*DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) *DescribeTransitRouterCapacityReservationsResponseBody
	GetMinimumTransitRouterCapacities() []*DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities
	SetRequestId(v string) *DescribeTransitRouterCapacityReservationsResponseBody
	GetRequestId() *string
}

type DescribeTransitRouterCapacityReservationsResponseBody struct {
	// The deployment set information specified for the resource reservation.
	DeploymentSets []*DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets `json:"DeploymentSets,omitempty" xml:"DeploymentSets,omitempty" type:"Repeated"`
	// The time when the resource reservation was last modified.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-10-19T06:31:49Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The list of reserved capacity information for each zone in the resource reservation.
	MinimumTransitRouterCapacities []*DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities `json:"MinimumTransitRouterCapacities,omitempty" xml:"MinimumTransitRouterCapacities,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTransitRouterCapacityReservationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterCapacityReservationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) GetDeploymentSets() []*DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets {
	return s.DeploymentSets
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) GetMinimumTransitRouterCapacities() []*DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities {
	return s.MinimumTransitRouterCapacities
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) SetDeploymentSets(v []*DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) *DescribeTransitRouterCapacityReservationsResponseBody {
	s.DeploymentSets = v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) SetLastModifiedTime(v string) *DescribeTransitRouterCapacityReservationsResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) SetMinimumTransitRouterCapacities(v []*DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) *DescribeTransitRouterCapacityReservationsResponseBody {
	s.MinimumTransitRouterCapacities = v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) SetRequestId(v string) *DescribeTransitRouterCapacityReservationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBody) Validate() error {
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
	return nil
}

type DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets struct {
	// The instance ID of the deployment set configured in the specified active zone of the transit router, which handles forward and routing tasks.
	//
	// example:
	//
	// ds-badxcs*****
	DeploymentSetID *string `json:"DeploymentSetID,omitempty" xml:"DeploymentSetID,omitempty"`
	// The zone for which you want to modify the deployment set of the resource reservation.
	//
	// - To cancel the deployment set, specify the zone and leave the deployment set instance ID empty.
	//
	// example:
	//
	// cn-shanghai-b
	ZoneID *string `json:"ZoneID,omitempty" xml:"ZoneID,omitempty"`
}

func (s DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) GetDeploymentSetID() *string {
	return s.DeploymentSetID
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) GetZoneID() *string {
	return s.ZoneID
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) SetDeploymentSetID(v string) *DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets {
	s.DeploymentSetID = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) SetZoneID(v string) *DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets {
	s.ZoneID = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyDeploymentSets) Validate() error {
	return dara.Validate(s)
}

type DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities struct {
	// The reserved capacity information for the specified zone in the resource reservation.
	DataProcessGbps *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps `json:"DataProcessGbps,omitempty" xml:"DataProcessGbps,omitempty" type:"Struct"`
	// The status information of the specified zone in the resource reservation.
	State *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState `json:"State,omitempty" xml:"State,omitempty" type:"Struct"`
	// The ID of the specified zone for the resource reservation.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) GetDataProcessGbps() *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps {
	return s.DataProcessGbps
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) GetState() *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState {
	return s.State
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) SetDataProcessGbps(v *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps) *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities {
	s.DataProcessGbps = v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) SetState(v *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState) *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities {
	s.State = v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) SetZoneId(v string) *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities {
	s.ZoneId = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacities) Validate() error {
	if s.DataProcessGbps != nil {
		if err := s.DataProcessGbps.Validate(); err != nil {
			return err
		}
	}
	if s.State != nil {
		if err := s.State.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps struct {
	// The effective reserved value for the specified zone in the resource reservation.
	//
	// - The effective reserved value refers to the current billed value for the resource reservation.
	//
	// example:
	//
	// 100
	Effective *int64 `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// The target reserved value for the specified active zone in the resource reservation.
	//
	// - In scenarios where you decrease the quota, the change takes effect in the next month, and the effective reserved value is updated in synchronization.
	//
	// - Upgrade and initial reservation scenarios require spooling by Alibaba Cloud and take effect immediately after the spooling is complete.
	//
	// example:
	//
	// 200
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps) GetEffective() *int64 {
	return s.Effective
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps) GetTarget() *int64 {
	return s.Target
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps) SetEffective(v int64) *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps {
	s.Effective = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps) SetTarget(v int64) *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps {
	s.Target = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesDataProcessGbps) Validate() error {
	return dara.Validate(s)
}

type DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState struct {
	// The reason for failure when the reservation status is ProvisionFailed.
	//
	// - For example, Reservation.InsufficientCapacity indicates that resources are insufficient to meet the target reserved value.
	//
	// example:
	//
	// Reservation.InsufficientCapacity
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The resource reservation status. Valid values:
	//
	// - Pending: The initial reservation is initiated and is pending processing by Alibaba Cloud.
	//
	// - Provisioning: Alibaba Cloud is processing the reservation request.
	//
	// - Provisioned: The reservation is successful.
	//
	// - ProvisionFailed: The reservation failed. You can refer to the failure reason.
	//
	// - Updating: The reservation is being updated. Operations such as modifying the deployment set or increasing the reserved value require spooling by Alibaba Cloud.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState) GetReason() *string {
	return s.Reason
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState) GetStatus() *string {
	return s.Status
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState) SetReason(v string) *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState {
	s.Reason = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState) SetStatus(v string) *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState {
	s.Status = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsResponseBodyMinimumTransitRouterCapacitiesState) Validate() error {
	return dara.Validate(s)
}
