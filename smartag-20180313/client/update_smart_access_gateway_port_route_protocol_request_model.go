// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayPortRouteProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetCrossAccount() *bool
	SetPortName(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetPortName() *string
	SetRegionId(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetRegionId() *string
	SetRemoteAs(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetRemoteAs() *string
	SetRemoteIp(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetRemoteIp() *string
	SetResourceUid(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetResourceUid() *string
	SetRouteProtocol(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetRouteProtocol() *string
	SetSagInsId(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetSagInsId() *string
	SetSagSn(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetSagSn() *string
	SetVlan(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest
	GetVlan() *string
}

type UpdateSmartAccessGatewayPortRouteProtocolRequest struct {
	// Specifies whether to query only the SAG instances that belong to another Alibaba Cloud account. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	CrossAccount *bool `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty"`
	// The port name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The region ID of the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The autonomous system number (ASN) of the SAG device.
	//
	// > When you enable BGP, you must set this parameter.
	//
	// example:
	//
	// 65535
	RemoteAs *string `json:"RemoteAs,omitempty" xml:"RemoteAs,omitempty"`
	// The IP address of the peer device.
	//
	// > When you enable BGP, you must set this parameter.
	//
	// example:
	//
	// 192.XX.XX.1
	RemoteIp *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 109790620697****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// The routing protocol. Valid values:
	//
	// 	- **STATIC**
	//
	// 	- **OSPF**
	//
	// 	- **BGP**
	//
	// This parameter is required.
	//
	// example:
	//
	// STATIC
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The ID of the Smart Access Gateway (SAG) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-3manef62evrfr6****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4dk****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
	// The VLAN ID.
	//
	// example:
	//
	// 10
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s UpdateSmartAccessGatewayPortRouteProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayPortRouteProtocolRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetPortName() *string {
	return s.PortName
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetRemoteAs() *string {
	return s.RemoteAs
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetRemoteIp() *string {
	return s.RemoteIp
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) GetVlan() *string {
	return s.Vlan
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetCrossAccount(v bool) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.CrossAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetPortName(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.PortName = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetRegionId(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetRemoteAs(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.RemoteAs = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetRemoteIp(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.RemoteIp = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetResourceUid(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.ResourceUid = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetRouteProtocol(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.RouteProtocol = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetSagInsId(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.SagInsId = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetSagSn(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.SagSn = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) SetVlan(v string) *UpdateSmartAccessGatewayPortRouteProtocolRequest {
	s.Vlan = &v
	return s
}

func (s *UpdateSmartAccessGatewayPortRouteProtocolRequest) Validate() error {
	return dara.Validate(s)
}
