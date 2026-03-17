// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSagStaticRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidr(v string) *CreateSagStaticRouteRequest
	GetDestinationCidr() *string
	SetNextHop(v string) *CreateSagStaticRouteRequest
	GetNextHop() *string
	SetOwnerAccount(v string) *CreateSagStaticRouteRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSagStaticRouteRequest
	GetOwnerId() *int64
	SetPortName(v string) *CreateSagStaticRouteRequest
	GetPortName() *string
	SetRegionId(v string) *CreateSagStaticRouteRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSagStaticRouteRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSagStaticRouteRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *CreateSagStaticRouteRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *CreateSagStaticRouteRequest
	GetSmartAGSn() *string
	SetVlan(v string) *CreateSagStaticRouteRequest
	GetVlan() *string
}

type CreateSagStaticRouteRequest struct {
	// The destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.0/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// Enter the IP address of the next hop.
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
	// The ID of the region where the SAG instance is deployed.
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
	// 20
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s CreateSagStaticRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSagStaticRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateSagStaticRouteRequest) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *CreateSagStaticRouteRequest) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateSagStaticRouteRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSagStaticRouteRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSagStaticRouteRequest) GetPortName() *string {
	return s.PortName
}

func (s *CreateSagStaticRouteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSagStaticRouteRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSagStaticRouteRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSagStaticRouteRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *CreateSagStaticRouteRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *CreateSagStaticRouteRequest) GetVlan() *string {
	return s.Vlan
}

func (s *CreateSagStaticRouteRequest) SetDestinationCidr(v string) *CreateSagStaticRouteRequest {
	s.DestinationCidr = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetNextHop(v string) *CreateSagStaticRouteRequest {
	s.NextHop = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetOwnerAccount(v string) *CreateSagStaticRouteRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetOwnerId(v int64) *CreateSagStaticRouteRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetPortName(v string) *CreateSagStaticRouteRequest {
	s.PortName = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetRegionId(v string) *CreateSagStaticRouteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetResourceOwnerAccount(v string) *CreateSagStaticRouteRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetResourceOwnerId(v int64) *CreateSagStaticRouteRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetSmartAGId(v string) *CreateSagStaticRouteRequest {
	s.SmartAGId = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetSmartAGSn(v string) *CreateSagStaticRouteRequest {
	s.SmartAGSn = &v
	return s
}

func (s *CreateSagStaticRouteRequest) SetVlan(v string) *CreateSagStaticRouteRequest {
	s.Vlan = &v
	return s
}

func (s *CreateSagStaticRouteRequest) Validate() error {
	return dara.Validate(s)
}
