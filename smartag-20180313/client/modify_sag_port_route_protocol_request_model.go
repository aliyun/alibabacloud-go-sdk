// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagPortRouteProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifySagPortRouteProtocolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagPortRouteProtocolRequest
	GetOwnerId() *int64
	SetPortName(v string) *ModifySagPortRouteProtocolRequest
	GetPortName() *string
	SetRegionId(v string) *ModifySagPortRouteProtocolRequest
	GetRegionId() *string
	SetRemoteAs(v string) *ModifySagPortRouteProtocolRequest
	GetRemoteAs() *string
	SetRemoteIp(v string) *ModifySagPortRouteProtocolRequest
	GetRemoteIp() *string
	SetResourceOwnerAccount(v string) *ModifySagPortRouteProtocolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagPortRouteProtocolRequest
	GetResourceOwnerId() *int64
	SetRouteProtocol(v string) *ModifySagPortRouteProtocolRequest
	GetRouteProtocol() *string
	SetSmartAGId(v string) *ModifySagPortRouteProtocolRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagPortRouteProtocolRequest
	GetSmartAGSn() *string
	SetVlan(v string) *ModifySagPortRouteProtocolRequest
	GetVlan() *string
}

type ModifySagPortRouteProtocolRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The ID of the region where the Smart Access Gateway (SAG) instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The BGP autonomous system number (ASN) of the peer device.
	//
	// >  You must set this parameter when you enable BGP.
	//
	// example:
	//
	// 65535
	RemoteAs *string `json:"RemoteAs,omitempty" xml:"RemoteAs,omitempty"`
	// The IP address of the peer device.
	//
	// >  You must set this parameter when you enable BGP.
	//
	// example:
	//
	// 192.XX.XX.1
	RemoteIp             *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The routing protocol. Valid values:
	//
	// 	- **STATIC**: uses a static routing protocol.
	//
	// 	- **OSPF**: uses the Open Shortest Path First protocol (OSPF).
	//
	// 	- **BGP**: uses the Border Gateway Protocol (BGP).
	//
	// This parameter is required.
	//
	// example:
	//
	// STATIC
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
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
	// example:
	//
	// 2
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s ModifySagPortRouteProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagPortRouteProtocolRequest) GoString() string {
	return s.String()
}

func (s *ModifySagPortRouteProtocolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagPortRouteProtocolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagPortRouteProtocolRequest) GetPortName() *string {
	return s.PortName
}

func (s *ModifySagPortRouteProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagPortRouteProtocolRequest) GetRemoteAs() *string {
	return s.RemoteAs
}

func (s *ModifySagPortRouteProtocolRequest) GetRemoteIp() *string {
	return s.RemoteIp
}

func (s *ModifySagPortRouteProtocolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagPortRouteProtocolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagPortRouteProtocolRequest) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *ModifySagPortRouteProtocolRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagPortRouteProtocolRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagPortRouteProtocolRequest) GetVlan() *string {
	return s.Vlan
}

func (s *ModifySagPortRouteProtocolRequest) SetOwnerAccount(v string) *ModifySagPortRouteProtocolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetOwnerId(v int64) *ModifySagPortRouteProtocolRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetPortName(v string) *ModifySagPortRouteProtocolRequest {
	s.PortName = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetRegionId(v string) *ModifySagPortRouteProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetRemoteAs(v string) *ModifySagPortRouteProtocolRequest {
	s.RemoteAs = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetRemoteIp(v string) *ModifySagPortRouteProtocolRequest {
	s.RemoteIp = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetResourceOwnerAccount(v string) *ModifySagPortRouteProtocolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetResourceOwnerId(v int64) *ModifySagPortRouteProtocolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetRouteProtocol(v string) *ModifySagPortRouteProtocolRequest {
	s.RouteProtocol = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetSmartAGId(v string) *ModifySagPortRouteProtocolRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetSmartAGSn(v string) *ModifySagPortRouteProtocolRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) SetVlan(v string) *ModifySagPortRouteProtocolRequest {
	s.Vlan = &v
	return s
}

func (s *ModifySagPortRouteProtocolRequest) Validate() error {
	return dara.Validate(s)
}
