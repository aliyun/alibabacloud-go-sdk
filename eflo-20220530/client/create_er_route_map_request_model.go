// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateErRouteMapRequest
	GetDescription() *string
	SetDestinationCidrBlock(v string) *CreateErRouteMapRequest
	GetDestinationCidrBlock() *string
	SetErId(v string) *CreateErRouteMapRequest
	GetErId() *string
	SetReceptionInstanceId(v string) *CreateErRouteMapRequest
	GetReceptionInstanceId() *string
	SetReceptionInstanceOwner(v string) *CreateErRouteMapRequest
	GetReceptionInstanceOwner() *string
	SetReceptionInstanceType(v string) *CreateErRouteMapRequest
	GetReceptionInstanceType() *string
	SetRegionId(v string) *CreateErRouteMapRequest
	GetRegionId() *string
	SetRouteMapAction(v string) *CreateErRouteMapRequest
	GetRouteMapAction() *string
	SetRouteMapNum(v int32) *CreateErRouteMapRequest
	GetRouteMapNum() *int32
	SetTransmissionInstanceId(v string) *CreateErRouteMapRequest
	GetTransmissionInstanceId() *string
	SetTransmissionInstanceOwner(v string) *CreateErRouteMapRequest
	GetTransmissionInstanceOwner() *string
	SetTransmissionInstanceType(v string) *CreateErRouteMapRequest
	GetTransmissionInstanceType() *string
}

type CreateErRouteMapRequest struct {
	// Policy description
	//
	// example:
	//
	// terraform-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination CIDR Block
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the destination instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xlhsvdvt
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// The tenant to which the route receiving instance belongs.
	//
	// example:
	//
	// 1620939556166277
	ReceptionInstanceOwner *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	// The type of the destination instance. Valid values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Policy behavior; optional values:
	//
	// 	- **permit**: Allow
	//
	// 	- **deny**: Rejected
	//
	// This parameter is required.
	//
	// example:
	//
	// permit
	RouteMapAction *string `json:"RouteMapAction,omitempty" xml:"RouteMapAction,omitempty"`
	// The ID of the policy.
	//
	// A smaller sequence number indicates a lower priority. When a route is matched, a policy with a higher priority is preferentially matched.
	//
	// **Valid values: 1001 to 2000**
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	RouteMapNum *int32 `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	// The ID of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-xlsjsdvt
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// The tenant to which the route publish instance belongs
	//
	// example:
	//
	// 1620939556166277
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	// The type of the source instance. Valid values:
	//
	// 	- **VPD**: Lingjun network segment.
	//
	// 	- **VCC**: Lingjun Connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPD
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s CreateErRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateErRouteMapRequest) GoString() string {
	return s.String()
}

func (s *CreateErRouteMapRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateErRouteMapRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateErRouteMapRequest) GetErId() *string {
	return s.ErId
}

func (s *CreateErRouteMapRequest) GetReceptionInstanceId() *string {
	return s.ReceptionInstanceId
}

func (s *CreateErRouteMapRequest) GetReceptionInstanceOwner() *string {
	return s.ReceptionInstanceOwner
}

func (s *CreateErRouteMapRequest) GetReceptionInstanceType() *string {
	return s.ReceptionInstanceType
}

func (s *CreateErRouteMapRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateErRouteMapRequest) GetRouteMapAction() *string {
	return s.RouteMapAction
}

func (s *CreateErRouteMapRequest) GetRouteMapNum() *int32 {
	return s.RouteMapNum
}

func (s *CreateErRouteMapRequest) GetTransmissionInstanceId() *string {
	return s.TransmissionInstanceId
}

func (s *CreateErRouteMapRequest) GetTransmissionInstanceOwner() *string {
	return s.TransmissionInstanceOwner
}

func (s *CreateErRouteMapRequest) GetTransmissionInstanceType() *string {
	return s.TransmissionInstanceType
}

func (s *CreateErRouteMapRequest) SetDescription(v string) *CreateErRouteMapRequest {
	s.Description = &v
	return s
}

func (s *CreateErRouteMapRequest) SetDestinationCidrBlock(v string) *CreateErRouteMapRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateErRouteMapRequest) SetErId(v string) *CreateErRouteMapRequest {
	s.ErId = &v
	return s
}

func (s *CreateErRouteMapRequest) SetReceptionInstanceId(v string) *CreateErRouteMapRequest {
	s.ReceptionInstanceId = &v
	return s
}

func (s *CreateErRouteMapRequest) SetReceptionInstanceOwner(v string) *CreateErRouteMapRequest {
	s.ReceptionInstanceOwner = &v
	return s
}

func (s *CreateErRouteMapRequest) SetReceptionInstanceType(v string) *CreateErRouteMapRequest {
	s.ReceptionInstanceType = &v
	return s
}

func (s *CreateErRouteMapRequest) SetRegionId(v string) *CreateErRouteMapRequest {
	s.RegionId = &v
	return s
}

func (s *CreateErRouteMapRequest) SetRouteMapAction(v string) *CreateErRouteMapRequest {
	s.RouteMapAction = &v
	return s
}

func (s *CreateErRouteMapRequest) SetRouteMapNum(v int32) *CreateErRouteMapRequest {
	s.RouteMapNum = &v
	return s
}

func (s *CreateErRouteMapRequest) SetTransmissionInstanceId(v string) *CreateErRouteMapRequest {
	s.TransmissionInstanceId = &v
	return s
}

func (s *CreateErRouteMapRequest) SetTransmissionInstanceOwner(v string) *CreateErRouteMapRequest {
	s.TransmissionInstanceOwner = &v
	return s
}

func (s *CreateErRouteMapRequest) SetTransmissionInstanceType(v string) *CreateErRouteMapRequest {
	s.TransmissionInstanceType = &v
	return s
}

func (s *CreateErRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
