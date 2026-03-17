// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagStaticRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidr(v string) *ModifySagStaticRouteRequest
	GetDestinationCidr() *string
	SetNextHop(v string) *ModifySagStaticRouteRequest
	GetNextHop() *string
	SetOwnerAccount(v string) *ModifySagStaticRouteRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagStaticRouteRequest
	GetOwnerId() *int64
	SetPortName(v string) *ModifySagStaticRouteRequest
	GetPortName() *string
	SetRegionId(v string) *ModifySagStaticRouteRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagStaticRouteRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagStaticRouteRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySagStaticRouteRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagStaticRouteRequest
	GetSmartAGSn() *string
	SetVlan(v string) *ModifySagStaticRouteRequest
	GetVlan() *string
}

type ModifySagStaticRouteRequest struct {
	// The destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.0/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The IP address of the next hop.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.1
	NextHop      *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The ID of the region where the Smart Access Gateway (SAG) instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
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
	// Valid values: **0*	- indicates a physical port. **1 to 4094*	- indicates that VLANs are supported.
	//
	// example:
	//
	// 1
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s ModifySagStaticRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagStaticRouteRequest) GoString() string {
	return s.String()
}

func (s *ModifySagStaticRouteRequest) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *ModifySagStaticRouteRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *ModifySagStaticRouteRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagStaticRouteRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagStaticRouteRequest) GetPortName() *string {
	return s.PortName
}

func (s *ModifySagStaticRouteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagStaticRouteRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagStaticRouteRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagStaticRouteRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagStaticRouteRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagStaticRouteRequest) GetVlan() *string {
	return s.Vlan
}

func (s *ModifySagStaticRouteRequest) SetDestinationCidr(v string) *ModifySagStaticRouteRequest {
	s.DestinationCidr = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetNextHop(v string) *ModifySagStaticRouteRequest {
	s.NextHop = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetOwnerAccount(v string) *ModifySagStaticRouteRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetOwnerId(v int64) *ModifySagStaticRouteRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetPortName(v string) *ModifySagStaticRouteRequest {
	s.PortName = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetRegionId(v string) *ModifySagStaticRouteRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetResourceOwnerAccount(v string) *ModifySagStaticRouteRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetResourceOwnerId(v int64) *ModifySagStaticRouteRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetSmartAGId(v string) *ModifySagStaticRouteRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetSmartAGSn(v string) *ModifySagStaticRouteRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagStaticRouteRequest) SetVlan(v string) *ModifySagStaticRouteRequest {
	s.Vlan = &v
	return s
}

func (s *ModifySagStaticRouteRequest) Validate() error {
	return dara.Validate(s)
}
