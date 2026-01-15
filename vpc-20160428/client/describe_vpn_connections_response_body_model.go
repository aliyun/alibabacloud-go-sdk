// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpnConnectionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnConnectionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpnConnectionsResponseBody
	GetTotalCount() *int32
	SetVpnConnections(v *DescribeVpnConnectionsResponseBodyVpnConnections) *DescribeVpnConnectionsResponseBody
	GetVpnConnections() *DescribeVpnConnectionsResponseBodyVpnConnections
}

type DescribeVpnConnectionsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 238752DC-0693-49BE-9C85-711D5691D3E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the IPsec-VPN connections.
	VpnConnections *DescribeVpnConnectionsResponseBodyVpnConnections `json:"VpnConnections,omitempty" xml:"VpnConnections,omitempty" type:"Struct"`
}

func (s DescribeVpnConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnConnectionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnConnectionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpnConnectionsResponseBody) GetVpnConnections() *DescribeVpnConnectionsResponseBodyVpnConnections {
	return s.VpnConnections
}

func (s *DescribeVpnConnectionsResponseBody) SetPageNumber(v int32) *DescribeVpnConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBody) SetPageSize(v int32) *DescribeVpnConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBody) SetRequestId(v string) *DescribeVpnConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBody) SetTotalCount(v int32) *DescribeVpnConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBody) SetVpnConnections(v *DescribeVpnConnectionsResponseBodyVpnConnections) *DescribeVpnConnectionsResponseBody {
	s.VpnConnections = v
	return s
}

func (s *DescribeVpnConnectionsResponseBody) Validate() error {
	if s.VpnConnections != nil {
		if err := s.VpnConnections.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnConnectionsResponseBodyVpnConnections struct {
	VpnConnection []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection `json:"VpnConnection,omitempty" xml:"VpnConnection,omitempty" type:"Repeated"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnections) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnections) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnections) GetVpnConnection() []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	return s.VpnConnection
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnections) SetVpnConnection(v []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) *DescribeVpnConnectionsResponseBodyVpnConnections {
	s.VpnConnection = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnections) Validate() error {
	if s.VpnConnection != nil {
		for _, item := range s.VpnConnection {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection struct {
	// The ID of the CEN instance to which the transit router belongs.
	//
	// example:
	//
	// cen-lxxpbpalc776qz****
	AttachInstanceId *string `json:"AttachInstanceId,omitempty" xml:"AttachInstanceId,omitempty"`
	// The type of resource that is associated with the IPsec-VPN connection. Valid values:
	//
	// 	- **CEN**: indicates that the IPsec-VPN connection is associated with a transit router of a Cloud Enterprise Network (CEN) instance.
	//
	// 	- **NO_ASSOCIATED**: indicates that the IPsec-VPN connection is not associated with any resource.
	//
	// 	- **VPNGW**: indicates that the IPsec-VPN connection is associated with a VPN gateway.
	//
	// example:
	//
	// CEN
	AttachType *string `json:"AttachType,omitempty" xml:"AttachType,omitempty"`
	// The timestamp generated when the IPsec-VPN connection was established. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1492753817000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the IPsec-VPN connection is associated with a transit router that belongs to another Alibaba Cloud account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CrossAccountAuthorized *bool `json:"CrossAccountAuthorized,omitempty" xml:"CrossAccountAuthorized,omitempty"`
	// The ID of the customer gateway associated with the IPsec-VPN connection.
	//
	// example:
	//
	// cgw-bp1mvj4g9kogw****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Indicates whether IPsec negotiations immediately start.
	//
	// 	- **true**: Negotiations are reinitiated after the configuration is changed.
	//
	// 	- **false**: Negotiations are reinitiated after traffic is detected.
	//
	// example:
	//
	// true
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// Indicates whether dead peer detection (DPD) is enabled for the IPsec-VPN connection. Valid values:
	//
	// 	- **true**
	//
	//     The initiator of the IPsec-VPN connection sends DPD packets to check the existence and availability of the peer. If no feedback is received from the peer within a specific period of time, the connection fails. Then, the ISAKMP security association (SA), IPsec SA, and IPsec tunnel are deleted.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Indicates whether NAT traversal is enabled for the IPsec-VPN connection.
	//
	// 	- **true**
	//
	//     After NAT traversal is enabled, the initiator does not check the UDP ports during IKE negotiations and can automatically discover NAT gateway devices along the IPsec tunnel.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// Indicates whether BGP is enabled for the tunnel. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// The configurations of Phase 1 negotiations.
	IkeConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty" type:"Struct"`
	// The gateway IP address of the IPsec-VPN connection.
	//
	// >  This parameter is returned only if the IPsec-VPN connection is associated with a transit router.
	//
	// example:
	//
	// 10.XX.XX.10
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The configurations of Phase 2 negotiations.
	IpsecConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty" type:"Struct"`
	// The CIDR block on the Alibaba Cloud side.
	//
	// Multiple CIDR blocks are separated by commas (,).
	//
	// example:
	//
	// 192.168.0.0/16,172.17.0.0/16
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the IPsec-VPN connection. Valid values:
	//
	// 	- **public**
	//
	// 	- **private**
	//
	// example:
	//
	// public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The certificate authority (CA) certificate of the peer.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The CIDR block of the data center.
	//
	// Multiple CIDR blocks are separated by commas (,).
	//
	// example:
	//
	// 10.0.0.0/8,172.16.0.0/16
	RemoteSubnet *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	// The ID of the resource group to which the IPsec-VPN connection belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the resource group information.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The bandwidth specification of the IPsec-VPN connection. Unit: **Mbit/s**.
	//
	// example:
	//
	// 1000M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The association state of the IPsec-VPN connection. Valid values:
	//
	// 	- **active**: The IPsec-VPN connection is associated with a VPN gateway.
	//
	// 	- **init**: The IPsec-VPN connection is not associated with any resource and is being initialized.
	//
	// 	- **attaching**: The IPsec-VPN connection is being associated with a transit router.
	//
	// 	- **attached**: The IPsec-VPN connection is associated with a transit router.
	//
	// 	- **detaching**: The IPsec-VPN connection is being disassociated from a transit router.
	//
	// 	- **financialLocked**: The IPsec-VPN connection is locked due to overdue payments.
	//
	// 	- **provisioning**: The IPsec-VPN connection is being prepared.
	//
	// 	- **updating**: The IPsec-VPN connection is being updated.
	//
	// 	- **Upgrading**: The IPsec-VPN connection is being upgraded.
	//
	// 	- **deleted**: The IPsec-VPN connection is deleted.
	//
	// example:
	//
	// attached
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the IPsec-VPN connection. Valid values:
	//
	// 	- **ike_sa_not_established**: Phase 1 negotiations failed.
	//
	// 	- **ike_sa_established**: Phase 1 negotiations succeeded.
	//
	// 	- **ipsec_sa_not_established**: Phase 2 negotiations failed.
	//
	// 	- **ipsec_sa_established**: Phase 2 negotiations succeeded.
	//
	// example:
	//
	// ipsec_sa_established
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags to be added to the IPsec-VPN connection.
	Tag *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	// The ID of the transit router with which the IPsec-VPN connection is associated.
	//
	// example:
	//
	// tr-p0we2edef9qr44a85****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The name of the transit router.
	//
	// example:
	//
	// nametest
	TransitRouterName *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
	TunnelBandwidth   *string `json:"TunnelBandwidth,omitempty" xml:"TunnelBandwidth,omitempty"`
	// The tunnel configurations of the IPsec-VPN connection.
	//
	// Parameters in **TunnelOptionsSpecification*	- are returned only if you query an IPsec-VPN connection in dual-tunnel mode.
	TunnelOptionsSpecification *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Struct"`
	// The health check configuration of the IPsec-VPN connection.
	VcoHealthCheck *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck `json:"VcoHealthCheck,omitempty" xml:"VcoHealthCheck,omitempty" type:"Struct"`
	// The BGP configuration of the IPsec-VPN connection.
	VpnBgpConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig `json:"VpnBgpConfig,omitempty" xml:"VpnBgpConfig,omitempty" type:"Struct"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-bp10lz7aejumd****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetAttachInstanceId() *string {
	return s.AttachInstanceId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetAttachType() *string {
	return s.AttachType
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetCrossAccountAuthorized() *bool {
	return s.CrossAccountAuthorized
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetEffectImmediately() *bool {
	return s.EffectImmediately
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetEnableDpd() *bool {
	return s.EnableDpd
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetEnableNatTraversal() *bool {
	return s.EnableNatTraversal
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetEnableTunnelsBgp() *bool {
	return s.EnableTunnelsBgp
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetIkeConfig() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	return s.IkeConfig
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetIpsecConfig() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig {
	return s.IpsecConfig
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetName() *string {
	return s.Name
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetRemoteSubnet() *string {
	return s.RemoteSubnet
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetSpec() *string {
	return s.Spec
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetState() *string {
	return s.State
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetTag() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag {
	return s.Tag
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetTransitRouterName() *string {
	return s.TransitRouterName
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetTunnelBandwidth() *string {
	return s.TunnelBandwidth
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetTunnelOptionsSpecification() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification {
	return s.TunnelOptionsSpecification
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetVcoHealthCheck() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck {
	return s.VcoHealthCheck
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetVpnBgpConfig() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig {
	return s.VpnBgpConfig
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetAttachInstanceId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.AttachInstanceId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetAttachType(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.AttachType = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetCreateTime(v int64) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetCrossAccountAuthorized(v bool) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.CrossAccountAuthorized = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetCustomerGatewayId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetEffectImmediately(v bool) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.EffectImmediately = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetEnableDpd(v bool) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.EnableDpd = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetEnableNatTraversal(v bool) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.EnableNatTraversal = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetEnableTunnelsBgp(v bool) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.EnableTunnelsBgp = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetIkeConfig(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.IkeConfig = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetInternetIp(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.InternetIp = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetIpsecConfig(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.IpsecConfig = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetLocalSubnet(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.LocalSubnet = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetName(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.Name = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetNetworkType(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.NetworkType = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetRemoteCaCertificate(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.RemoteCaCertificate = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetRemoteSubnet(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.RemoteSubnet = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetResourceGroupId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetSpec(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.Spec = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetState(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.State = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetStatus(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.Status = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetTag(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.Tag = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetTransitRouterId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetTransitRouterName(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.TransitRouterName = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetTunnelBandwidth(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.TunnelBandwidth = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetTunnelOptionsSpecification(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.TunnelOptionsSpecification = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetVcoHealthCheck(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.VcoHealthCheck = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetVpnBgpConfig(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.VpnBgpConfig = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetVpnConnectionId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) SetVpnGatewayId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnection) Validate() error {
	if s.IkeConfig != nil {
		if err := s.IkeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IpsecConfig != nil {
		if err := s.IpsecConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		if err := s.Tag.Validate(); err != nil {
			return err
		}
	}
	if s.TunnelOptionsSpecification != nil {
		if err := s.TunnelOptionsSpecification.Validate(); err != nil {
			return err
		}
	}
	if s.VcoHealthCheck != nil {
		if err := s.VcoHealthCheck.Validate(); err != nil {
			return err
		}
	}
	if s.VpnBgpConfig != nil {
		if err := s.VpnBgpConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig struct {
	// The authentication algorithm in the IKE phase.
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm in the IKE phase.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The lifetime in the IKE phase. Unit: seconds.
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The IKE negotiation mode.
	//
	// 	- **main**: This mode offers higher security during negotiations.
	//
	// 	- **aggressive**: This mode is faster and has a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The DH group in the IKE phase.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The version of the IKE protocol.
	//
	// 	- **ikev1**
	//
	// 	- **ikev2**
	//
	// Compared with IKEv1, IKEv2 simplifies the SA negotiation process and is more suitable for scenarios in which multiple CIDR blocks are used.
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the IPsec-VPN connection on the data center side.
	//
	// example:
	//
	// 116.64.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// pgw6dy7****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the IPsec-VPN connection on the Alibaba Cloud side.
	//
	// example:
	//
	// 139.17.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetIkeLifetime() *int64 {
	return s.IkeLifetime
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetIkeAuthAlg(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetIkeEncAlg(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetIkeLifetime(v int64) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetIkeMode(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetIkePfs(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetIkeVersion(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetLocalId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.LocalId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetPsk(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.Psk = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) SetRemoteId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig struct {
	// The authentication algorithm in the IPsec phase.
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm in the IPsec phase.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The lifetime in the IPsec phase. Unit: seconds.
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The DH group in the IPsec phase.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) GetIpsecLifetime() *int64 {
	return s.IpsecLifetime
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) SetIpsecAuthAlg(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) SetIpsecEncAlg(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) SetIpsecLifetime(v int64) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) SetIpsecPfs(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag struct {
	Tag []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag) GetTag() []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag {
	return s.Tag
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag) SetTag(v []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag {
	s.Tag = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag struct {
	// The tag key of the IPsec-VPN connection.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the IPsec-VPN connection.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag) SetKey(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag {
	s.Key = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag) SetValue(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag {
	s.Value = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTagTag) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification struct {
	TunnelOptions []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions `json:"TunnelOptions,omitempty" xml:"TunnelOptions,omitempty" type:"Repeated"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification) GetTunnelOptions() []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	return s.TunnelOptions
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification) SetTunnelOptions(v []*DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification {
	s.TunnelOptions = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification) Validate() error {
	if s.TunnelOptions != nil {
		for _, item := range s.TunnelOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions struct {
	// The ID of the customer gateway associated with the tunnel.
	//
	// example:
	//
	// cgw-p0wy363lucf1uyae8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// Indicates whether the DPD feature is enabled for the tunnel. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	EnableDpd *string `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// Indicates whether NAT traversal is enabled for the tunnel. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	EnableNatTraversal *string `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// The tunnel IP address.
	//
	// example:
	//
	// 47.21.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The CA certificate of the tunnel peer.
	//
	// This parameter is returned only if the VPN gateway is of the SM type.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// The tunnel role. Valid values:
	//
	// 	- **master**: The tunnel is an active tunnel.
	//
	// 	- **slave**: The tunnel is a standby tunnel.
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The tunnel status. Valid values:
	//
	// 	- **active**
	//
	// 	- **updating**
	//
	// 	- **deleting**
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The state of the IPsec-VPN connection. Valid values:
	//
	// 	- **ike_sa_not_established**: Phase 1 negotiations failed.
	//
	// 	- **ike_sa_established**: Phase 1 negotiations succeeded.
	//
	// 	- **ipsec_sa_not_established**: Phase 2 negotiations failed.
	//
	// 	- **ipsec_sa_established**: Phase 2 negotiations succeeded.
	//
	// example:
	//
	// ipsec_sa_established
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The BGP configurations.
	TunnelBgpConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-opsqc4d97wni27****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The configuration of Phase 1 negotiations.
	TunnelIkeConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// The order in which the tunnel is created.
	//
	// 	- **1**: Tunnel 1.
	//
	// 	- **2**: Tunnel 2.
	//
	// >  This parameter is returned only if the IPsec-VPN connection is associated with a transit router.
	//
	// example:
	//
	// 1
	TunnelIndex *int32 `json:"TunnelIndex,omitempty" xml:"TunnelIndex,omitempty"`
	// The configurations of Phase 2 negotiations.
	TunnelIpsecConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
	// The zone of the tunnel.
	//
	// example:
	//
	// ap-southeast-5a
	ZoneNo *string `json:"ZoneNo,omitempty" xml:"ZoneNo,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetEnableDpd() *string {
	return s.EnableDpd
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetEnableNatTraversal() *string {
	return s.EnableNatTraversal
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetRemoteCaCertificate() *string {
	return s.RemoteCaCertificate
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetRole() *string {
	return s.Role
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetState() *string {
	return s.State
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetTunnelBgpConfig() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	return s.TunnelBgpConfig
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetTunnelId() *string {
	return s.TunnelId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetTunnelIkeConfig() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	return s.TunnelIkeConfig
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetTunnelIndex() *int32 {
	return s.TunnelIndex
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetTunnelIpsecConfig() *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	return s.TunnelIpsecConfig
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) GetZoneNo() *string {
	return s.ZoneNo
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetCustomerGatewayId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetEnableDpd(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.EnableDpd = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetEnableNatTraversal(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.EnableNatTraversal = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetInternetIp(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.InternetIp = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetRemoteCaCertificate(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.RemoteCaCertificate = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetRole(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.Role = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetState(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.State = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetStatus(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.Status = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetTunnelBgpConfig(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.TunnelBgpConfig = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetTunnelId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.TunnelId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetTunnelIkeConfig(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.TunnelIkeConfig = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetTunnelIndex(v int32) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.TunnelIndex = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetTunnelIpsecConfig(v *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.TunnelIpsecConfig = v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) SetZoneNo(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions {
	s.ZoneNo = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptions) Validate() error {
	if s.TunnelBgpConfig != nil {
		if err := s.TunnelBgpConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TunnelIkeConfig != nil {
		if err := s.TunnelIkeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TunnelIpsecConfig != nil {
		if err := s.TunnelIpsecConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig struct {
	// The negotiation state of BGP. Valid values:
	//
	// 	- **success**
	//
	// 	- **false**
	//
	// example:
	//
	// success
	BgpStatus *string `json:"BgpStatus,omitempty" xml:"BgpStatus,omitempty"`
	// The ASN on the Alibaba Cloud side.
	//
	// example:
	//
	// 65530
	LocalAsn *string `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP address on the Alibaba Cloud side.
	//
	// example:
	//
	// 169.254.10.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The ASN of the tunnel peer.
	//
	// example:
	//
	// 65531
	PeerAsn *string `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The BGP IP address of the tunnel peer.
	//
	// example:
	//
	// 169.254.10.2
	PeerBgpIp *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	// The BGP CIDR block of the tunnel.
	//
	// example:
	//
	// 169.254.10.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetBgpStatus() *string {
	return s.BgpStatus
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetLocalAsn() *string {
	return s.LocalAsn
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetPeerAsn() *string {
	return s.PeerAsn
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetBgpStatus(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.BgpStatus = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetLocalAsn(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetLocalBgpIp(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetPeerAsn(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetPeerBgpIp(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) SetTunnelCidr(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig struct {
	// The authentication algorithm in the IKE phase.
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// The encryption algorithm in the IKE phase.
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// The lifetime in the IKE phase. Unit: seconds.
	//
	// example:
	//
	// 86400
	IkeLifetime *string `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// The IKE negotiation mode.
	//
	// 	- **main**: This mode offers higher security during negotiations.
	//
	// 	- **aggressive**: This mode is faster and has a higher success rate.
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// The DH group in the IKE phase.
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// The version of the IKE protocol.
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// The identifier of the tunnel on the Alibaba Cloud side.
	//
	// example:
	//
	// 47.21.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// The pre-shared key.
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// The identifier of the tunnel peer.
	//
	// example:
	//
	// 47.42.XX.XX
	RemoteId *string `json:"RemoteId,omitempty" xml:"RemoteId,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeAuthAlg() *string {
	return s.IkeAuthAlg
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeEncAlg() *string {
	return s.IkeEncAlg
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeLifetime() *string {
	return s.IkeLifetime
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeMode() *string {
	return s.IkeMode
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkePfs() *string {
	return s.IkePfs
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetIkeVersion() *string {
	return s.IkeVersion
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetLocalId() *string {
	return s.LocalId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetPsk() *string {
	return s.Psk
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) GetRemoteId() *string {
	return s.RemoteId
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeAuthAlg(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeAuthAlg = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeEncAlg(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeEncAlg = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeLifetime(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeLifetime = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeMode(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeMode = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkePfs(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkePfs = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetIkeVersion(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.IkeVersion = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetLocalId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.LocalId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetPsk(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.Psk = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) SetRemoteId(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig {
	s.RemoteId = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig struct {
	// The authentication algorithm in the IPsec phase.
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// The encryption algorithm in the IPsec phase.
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// The lifetime in the IPsec phase. Unit: seconds.
	//
	// example:
	//
	// 86400
	IpsecLifetime *string `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// The DH group in the IPsec phase.
	//
	// example:
	//
	// group2
	IpsecPfs *string `json:"IpsecPfs,omitempty" xml:"IpsecPfs,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecAuthAlg() *string {
	return s.IpsecAuthAlg
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecEncAlg() *string {
	return s.IpsecEncAlg
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecLifetime() *string {
	return s.IpsecLifetime
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) GetIpsecPfs() *string {
	return s.IpsecPfs
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecAuthAlg(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecAuthAlg = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecEncAlg(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecEncAlg = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecLifetime(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecLifetime = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) SetIpsecPfs(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig {
	s.IpsecPfs = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck struct {
	// The destination IP address.
	//
	// example:
	//
	// 192.168.0.1
	Dip *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	// Indicates whether the health check feature is enabled.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// example:
	//
	// 2
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Indicates whether advertised routes are withdrawn when the health check fails.
	//
	// 	- **revoke_route**: Advertised routes are withdrawn.
	//
	// 	- **reserve_route**: Advertised routes are not withdrawn.
	//
	// example:
	//
	// revoke_route
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The maximum number of health check retries.
	//
	// example:
	//
	// 3
	Retry *int32 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 192.168.0.50
	Sip *string `json:"Sip,omitempty" xml:"Sip,omitempty"`
	// The state of the health check. Valid values:
	//
	// 	- **success**
	//
	// 	- **failed**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) GetDip() *string {
	return s.Dip
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) GetEnable() *string {
	return s.Enable
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) GetRetry() *int32 {
	return s.Retry
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) GetSip() *string {
	return s.Sip
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) SetDip(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck {
	s.Dip = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) SetEnable(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck {
	s.Enable = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) SetInterval(v int32) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck {
	s.Interval = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) SetPolicy(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck {
	s.Policy = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) SetRetry(v int32) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck {
	s.Retry = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) SetSip(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck {
	s.Sip = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) SetStatus(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck {
	s.Status = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig struct {
	// The authentication key of the BGP routing protocol.
	//
	// example:
	//
	// AuthKey****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The ASN on the Alibaba Cloud side.
	//
	// example:
	//
	// 65531
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP IP address on the Alibaba Cloud side.
	//
	// example:
	//
	// 169.254.10.2
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// The ASN of the peer.
	//
	// example:
	//
	// 65530
	PeerAsn *int64 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The BGP IP address of the peer.
	//
	// example:
	//
	// 169.254.10.1
	PeerBgpIp *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	// The negotiation state of the BGP routing protocol. Valid values:
	//
	// 	- **success**
	//
	// 	- **false**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The BGP CIDR block of the IPsec-VPN connection. The CIDR block falls within 169.254.0.0/16. The subnet mask of the CIDR block must be 30 bits in length.
	//
	// example:
	//
	// 169.254.10.0/30
	TunnelCidr *string `json:"TunnelCidr,omitempty" xml:"TunnelCidr,omitempty"`
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) GetLocalBgpIp() *string {
	return s.LocalBgpIp
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) GetPeerAsn() *int64 {
	return s.PeerAsn
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) GetPeerBgpIp() *string {
	return s.PeerBgpIp
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) GetTunnelCidr() *string {
	return s.TunnelCidr
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) SetAuthKey(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig {
	s.AuthKey = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) SetLocalAsn(v int64) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig {
	s.LocalAsn = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) SetLocalBgpIp(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig {
	s.LocalBgpIp = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) SetPeerAsn(v int64) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig {
	s.PeerAsn = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) SetPeerBgpIp(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig {
	s.PeerBgpIp = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) SetStatus(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig {
	s.Status = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) SetTunnelCidr(v string) *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig {
	s.TunnelCidr = &v
	return s
}

func (s *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig) Validate() error {
	return dara.Validate(s)
}
