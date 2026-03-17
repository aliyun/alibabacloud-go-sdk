// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSagStaticRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidr(v string) *DeleteSagStaticRouteRequest
	GetDestinationCidr() *string
	SetOwnerAccount(v string) *DeleteSagStaticRouteRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSagStaticRouteRequest
	GetOwnerId() *int64
	SetPortName(v string) *DeleteSagStaticRouteRequest
	GetPortName() *string
	SetRegionId(v string) *DeleteSagStaticRouteRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSagStaticRouteRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSagStaticRouteRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DeleteSagStaticRouteRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DeleteSagStaticRouteRequest
	GetSmartAGSn() *string
	SetVlan(v string) *DeleteSagStaticRouteRequest
	GetVlan() *string
}

type DeleteSagStaticRouteRequest struct {
	// The destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.0/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
	// The VLAN ID.
	//
	// Valid values: **0*	- (physical port) and **1 to 4094*	- (VLAN).
	//
	// example:
	//
	// 1
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s DeleteSagStaticRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSagStaticRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteSagStaticRouteRequest) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DeleteSagStaticRouteRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSagStaticRouteRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSagStaticRouteRequest) GetPortName() *string {
	return s.PortName
}

func (s *DeleteSagStaticRouteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSagStaticRouteRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSagStaticRouteRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSagStaticRouteRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DeleteSagStaticRouteRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DeleteSagStaticRouteRequest) GetVlan() *string {
	return s.Vlan
}

func (s *DeleteSagStaticRouteRequest) SetDestinationCidr(v string) *DeleteSagStaticRouteRequest {
	s.DestinationCidr = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetOwnerAccount(v string) *DeleteSagStaticRouteRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetOwnerId(v int64) *DeleteSagStaticRouteRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetPortName(v string) *DeleteSagStaticRouteRequest {
	s.PortName = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetRegionId(v string) *DeleteSagStaticRouteRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetResourceOwnerAccount(v string) *DeleteSagStaticRouteRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetResourceOwnerId(v int64) *DeleteSagStaticRouteRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetSmartAGId(v string) *DeleteSagStaticRouteRequest {
	s.SmartAGId = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetSmartAGSn(v string) *DeleteSagStaticRouteRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) SetVlan(v string) *DeleteSagStaticRouteRequest {
	s.Vlan = &v
	return s
}

func (s *DeleteSagStaticRouteRequest) Validate() error {
	return dara.Validate(s)
}
