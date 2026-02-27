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
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// 转发路由器实例所属的云企业网实例ID。
	//
	// example:
	//
	// cen-lxxpbpalc776qz****
	AttachInstanceId *string `json:"AttachInstanceId,omitempty" xml:"AttachInstanceId,omitempty"`
	// IPsec连接绑定的资源类型。
	//
	// - **CEN**：表示IPsec连接已绑定云企业网实例下的转发路由器实例。
	//
	// - **NO_ASSOCIATED**：表示IPsec连接未绑定任何资源。
	//
	// - **VPNGW**：表示IPsec连接绑定了VPN网关实例。
	//
	// example:
	//
	// CEN
	AttachType *string `json:"AttachType,omitempty" xml:"AttachType,omitempty"`
	// 创建IPsec连接的时间戳。单位：毫秒。
	//
	// 时间戳的格式采用Unix时间戳，表示从格林威治时间1970年01月01日00时00分00秒至创建IPsec连接时的总时长。
	//
	// example:
	//
	// 1492753817000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// IPsec连接是否绑定了跨账号的转发路由器实例。
	//
	// - **true**：是。
	//
	// - **false**：否。
	//
	// example:
	//
	// false
	CrossAccountAuthorized *bool `json:"CrossAccountAuthorized,omitempty" xml:"CrossAccountAuthorized,omitempty"`
	// IPsec连接关联的用户网关的实例ID。
	//
	// example:
	//
	// cgw-bp1mvj4g9kogw****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// IPsec连接的配置是否立即生效。
	//
	// - **true**：是，配置变更完成后触发重连。
	//
	// - **false**：否，有流量时触发重连。
	//
	// example:
	//
	// true
	EffectImmediately *bool `json:"EffectImmediately,omitempty" xml:"EffectImmediately,omitempty"`
	// IPsec连接是否已开启DPD（对等体存活检测）功能。
	//
	// - **true**：开启DPD功能。
	//
	//     IPsec发起端会发送DPD报文用来检测对端的设备是否存活，如果在设定时间内未收到正确回应则认为对端已经断线，IPsec将删除ISAKMP SA和相应的IPsec SA，安全隧道同样也会被删除。
	//
	// - **false**：不开启DPD功能，IPsec发起端不会发送DPD探测报文。
	//
	// example:
	//
	// true
	EnableDpd *bool `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// IPsec连接是否已开启NAT穿越功能。
	//
	// - **true**：开启NAT穿越功能。
	//
	//    开启后，IKE协商过程会删除对UDP端口号的验证过程，同时实现对VPN隧道中NAT网关设备的发现功能。
	//
	// - **false**：不开启NAT穿越功能。
	//
	// example:
	//
	// true
	EnableNatTraversal *bool `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// 隧道BGP的开启状态。
	//
	// - **true**：已开启。
	//
	// - **false**：未开启。
	//
	// example:
	//
	// true
	EnableTunnelsBgp *bool `json:"EnableTunnelsBgp,omitempty" xml:"EnableTunnelsBgp,omitempty"`
	// 第一阶段协商的配置。
	IkeConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIkeConfig `json:"IkeConfig,omitempty" xml:"IkeConfig,omitempty" type:"Struct"`
	// IPsec连接的网关IP地址。
	//
	// > 仅IPsec连接绑定转发路由器实例时会返回当前参数。
	//
	// example:
	//
	// 10.XX.XX.10
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// 第二阶段协商的配置。
	IpsecConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionIpsecConfig `json:"IpsecConfig,omitempty" xml:"IpsecConfig,omitempty" type:"Struct"`
	// IPsec连接阿里云侧的网段。
	//
	// 在多个网段的情况下，网段之间使用半角逗号（,）分隔。
	//
	// example:
	//
	// 192.168.0.0/16,172.17.0.0/16
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// IPsec连接的名称。
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// IPsec连接的网络类型。
	//
	// - **public**：公网，表示IPsec连接通过公网建立加密通信通道。
	//
	// - **private**：私网，表示IPsec连接通过私网建立加密通信通道。
	//
	// example:
	//
	// public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// 对端的CA证书。
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// 本地数据中心侧的网段。
	//
	// 在多个网段的情况下，网段之间使用半角逗号（,）分隔。
	//
	// example:
	//
	// 10.0.0.0/8,172.16.0.0/16
	RemoteSubnet *string `json:"RemoteSubnet,omitempty" xml:"RemoteSubnet,omitempty"`
	// IPsec连接所属的资源组ID。
	//
	// 您可以调用[ListResourceGroups](https://help.aliyun.com/document_detail/158855.html)接口查询资源组信息。
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// IPsec连接的带宽规格。单位：**Mbps**。
	//
	// example:
	//
	// 1000M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// IPsec连接与转发路由器实例的绑定状态。
	//
	// - **active**：IPsec连接已与VPN网关实例绑定，状态正常。
	//
	// - **init**：IPsec连接未绑定任何资源，IPsec连接初始化。
	//
	// - **attaching**：IPsec连接与转发路由器实例绑定中。
	//
	// - **attached**：IPsec连接已与转发路由器实例绑定。
	//
	// - **detaching**：IPsec连接与转发路由器实例解绑中。
	//
	// - **financialLocked**：欠费锁定。
	//
	// - **provisioning**：资源准备中。
	//
	// - **updating**：更新中。
	//
	// - **upgrading**：升级中。
	//
	// - **deleted**：已删除。
	//
	// example:
	//
	// attached
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// IPsec连接的状态。
	//
	// - **ike_sa_not_established**：第一阶段协商失败。
	//
	// - **ike_sa_established**：第一阶段协商成功。
	//
	// - **ipsec_sa_not_established**：第二阶段协商失败。
	//
	// - **ipsec_sa_established**：第二阶段协商成功。
	//
	// example:
	//
	// ipsec_sa_established
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// IPsec连接绑定的标签列表。
	Tag *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	// IPsec连接绑定的转发路由器实例ID。
	//
	// example:
	//
	// tr-p0we2edef9qr44a85****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// 转发路由器实例的名称。
	//
	// example:
	//
	// nametest
	TransitRouterName *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
	// 用于说明VPN单条隧道的带宽规格，取值：
	//
	// Standard（默认值）：标准型，默认带宽1Gbps
	//
	// Large（大型）：大型，默认带宽3Gbps
	//
	// example:
	//
	// Standard
	TunnelBandwidth *string `json:"TunnelBandwidth,omitempty" xml:"TunnelBandwidth,omitempty"`
	// IPsec连接的隧道配置信息。
	//
	// 仅查询双隧道模式的IPsec连接会返回**TunnelOptionsSpecification**数组下的参数。
	TunnelOptionsSpecification *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecification `json:"TunnelOptionsSpecification,omitempty" xml:"TunnelOptionsSpecification,omitempty" type:"Struct"`
	// IPsec连接的健康检查配置。
	VcoHealthCheck *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVcoHealthCheck `json:"VcoHealthCheck,omitempty" xml:"VcoHealthCheck,omitempty" type:"Struct"`
	// IPsec连接BGP路由协议的配置。
	VpnBgpConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionVpnBgpConfig `json:"VpnBgpConfig,omitempty" xml:"VpnBgpConfig,omitempty" type:"Struct"`
	// IPsec连接的ID。
	//
	// example:
	//
	// vco-bp10lz7aejumd****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// VPN网关的实例ID。
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
	// IKE阶段认证算法。
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// IKE阶段加密算法。
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// IKE阶段生存时间。单位：秒。
	//
	// example:
	//
	// 86400
	IkeLifetime *int64 `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// IKE阶段协商模式。
	//
	// - **main**：主模式，协商过程安全性高。
	//
	// - **aggressive**：野蛮模式，协商快速且协商成功率高。
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// IKE阶段DH分组。
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// IKE协议版本。
	//
	// - **ikev1**
	//
	// - **ikev2**
	//
	// 相对于IKEv1版本，IKEv2版本简化了SA的协商过程并且对于多网段的场景提供了更好的支持。
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// IPsec连接对端本地数据中心侧的标识。
	//
	// example:
	//
	// 116.64.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// 预共享密钥。
	//
	// example:
	//
	// pgw6dy7****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// IPsec连接阿里云侧的标识。
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
	// IPsec阶段认证算法。
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// IPsec阶段加密算法。
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// IPsec阶段生存时间。单位：秒。
	//
	// example:
	//
	// 86400
	IpsecLifetime *int64 `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// IPsec阶段DH分组。
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
	// 标签键。
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。
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
	// 隧道关联的用户网关ID。
	//
	// example:
	//
	// cgw-p0wy363lucf1uyae8****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// 隧道是否已开启DPD（对等体存活检测）功能。
	//
	// - **false**：未开启。
	//
	// - **true**：已开启。
	//
	// example:
	//
	// true
	EnableDpd *string `json:"EnableDpd,omitempty" xml:"EnableDpd,omitempty"`
	// 隧道是否已开启NAT穿越功能。
	//
	// - **false**：未开启。
	//
	// - **true**：已开启。
	//
	// example:
	//
	// true
	EnableNatTraversal *string `json:"EnableNatTraversal,omitempty" xml:"EnableNatTraversal,omitempty"`
	// 隧道的IP地址。
	//
	// example:
	//
	// 47.21.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// 隧道对端的CA证书。
	//
	// 仅VPN网关实例的类型为国密型时才会返回当前参数。
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIB7zCCAZW***	- -----END CERTIFICATE-----
	RemoteCaCertificate *string `json:"RemoteCaCertificate,omitempty" xml:"RemoteCaCertificate,omitempty"`
	// 隧道的角色。
	//
	// - **master**：表示当前隧道为主隧道。
	//
	// - **slave**：表示当前隧道为备隧道。
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 隧道的状态。
	//
	// - **active**：状态正常。
	//
	// - **updating**：更新中。
	//
	// - **deleting**：删除中。
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// IPsec连接的状态。
	//
	// - **ike_sa_not_established**：第一阶段协商失败。
	//
	// - **ike_sa_established**：第一阶段协商成功。
	//
	// - **ipsec_sa_not_established**：第二阶段协商失败。
	//
	// - **ipsec_sa_established**：第二阶段协商成功。
	//
	// example:
	//
	// ipsec_sa_established
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 隧道的BGP配置信息。
	TunnelBgpConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelBgpConfig `json:"TunnelBgpConfig,omitempty" xml:"TunnelBgpConfig,omitempty" type:"Struct"`
	// 隧道ID。
	//
	// example:
	//
	// tun-opsqc4d97wni27****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// 第一阶段协商的配置。
	TunnelIkeConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIkeConfig `json:"TunnelIkeConfig,omitempty" xml:"TunnelIkeConfig,omitempty" type:"Struct"`
	// 隧道的创建顺序。
	//
	// - **1**：第一条隧道。
	//
	// - **2**：第二条隧道。
	//
	// > 仅IPsec连接绑定转发路由器时会返回该参数。
	//
	// example:
	//
	// 1
	TunnelIndex *int32 `json:"TunnelIndex,omitempty" xml:"TunnelIndex,omitempty"`
	// 第二阶段协商的配置。
	TunnelIpsecConfig *DescribeVpnConnectionsResponseBodyVpnConnectionsVpnConnectionTunnelOptionsSpecificationTunnelOptionsTunnelIpsecConfig `json:"TunnelIpsecConfig,omitempty" xml:"TunnelIpsecConfig,omitempty" type:"Struct"`
	// 隧道部署的可用区。
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
	// BGP的协商状态。
	//
	// - **success**：正常。
	//
	// - **failed**：异常。
	//
	// example:
	//
	// success
	BgpStatus *string `json:"BgpStatus,omitempty" xml:"BgpStatus,omitempty"`
	// 隧道本端（阿里云侧）的自治系统号。
	//
	// example:
	//
	// 65530
	LocalAsn *string `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// 隧道本端（阿里云侧）的BGP地址。
	//
	// example:
	//
	// 169.254.10.1
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// 隧道对端的自治系统号。
	//
	// example:
	//
	// 65531
	PeerAsn *string `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// 隧道对端的BGP地址。
	//
	// example:
	//
	// 169.254.10.2
	PeerBgpIp *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	// 隧道的BGP网段。
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
	// IKE阶段认证算法。
	//
	// example:
	//
	// sha1
	IkeAuthAlg *string `json:"IkeAuthAlg,omitempty" xml:"IkeAuthAlg,omitempty"`
	// IKE阶段加密算法。
	//
	// example:
	//
	// aes
	IkeEncAlg *string `json:"IkeEncAlg,omitempty" xml:"IkeEncAlg,omitempty"`
	// IKE阶段生存时间。单位：秒。
	//
	// example:
	//
	// 86400
	IkeLifetime *string `json:"IkeLifetime,omitempty" xml:"IkeLifetime,omitempty"`
	// IKE协商模式。
	//
	// - **main**：主模式，协商过程安全性高。
	//
	// - **aggressive**：野蛮模式，协商快速且协商成功率高。
	//
	// example:
	//
	// main
	IkeMode *string `json:"IkeMode,omitempty" xml:"IkeMode,omitempty"`
	// IKE阶段DH分组。
	//
	// example:
	//
	// group2
	IkePfs *string `json:"IkePfs,omitempty" xml:"IkePfs,omitempty"`
	// IKE协议版本。
	//
	// example:
	//
	// ikev1
	IkeVersion *string `json:"IkeVersion,omitempty" xml:"IkeVersion,omitempty"`
	// 隧道本端（阿里云侧）的标识。
	//
	// example:
	//
	// 47.21.XX.XX
	LocalId *string `json:"LocalId,omitempty" xml:"LocalId,omitempty"`
	// 预共享密钥。
	//
	// example:
	//
	// 123456****
	Psk *string `json:"Psk,omitempty" xml:"Psk,omitempty"`
	// 隧道对端的标识。
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
	// IPsec阶段认证算法。
	//
	// example:
	//
	// sha1
	IpsecAuthAlg *string `json:"IpsecAuthAlg,omitempty" xml:"IpsecAuthAlg,omitempty"`
	// IPsec阶段加密算法。
	//
	// example:
	//
	// aes
	IpsecEncAlg *string `json:"IpsecEncAlg,omitempty" xml:"IpsecEncAlg,omitempty"`
	// IPsec阶段生存时间。单位：秒。
	//
	// example:
	//
	// 86400
	IpsecLifetime *string `json:"IpsecLifetime,omitempty" xml:"IpsecLifetime,omitempty"`
	// IPsec阶段DH分组。
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
	// 目的IP地址。
	//
	// example:
	//
	// 192.168.0.1
	Dip *string `json:"Dip,omitempty" xml:"Dip,omitempty"`
	// 健康检查的开启状态。
	//
	// - **true**：已开启。
	//
	// - **false**：未开启。
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// 健康检查的时间间隔。单位：秒。
	//
	// example:
	//
	// 2
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// 健康检查失败时是否撤销已发布的路由。
	//
	// - **revoke_route**：撤销路由。
	//
	// - **reserve_route**：不撤销路由。
	//
	// example:
	//
	// revoke_route
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 健康检查的重试发包次数。
	//
	// example:
	//
	// 3
	Retry *int32 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// 源IP地址。
	//
	// example:
	//
	// 192.168.0.50
	Sip *string `json:"Sip,omitempty" xml:"Sip,omitempty"`
	// 健康检查状态。
	//
	// - **success**：正常。
	//
	// - **failed**：异常。
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
	// BGP路由协议的认证密钥。
	//
	// example:
	//
	// AuthKey****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// 阿里云侧自治系统号。
	//
	// example:
	//
	// 65531
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// 阿里云侧BGP地址。
	//
	// example:
	//
	// 169.254.10.2
	LocalBgpIp *string `json:"LocalBgpIp,omitempty" xml:"LocalBgpIp,omitempty"`
	// 对端自治系统号。
	//
	// example:
	//
	// 65530
	PeerAsn *int64 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// 对端BGP地址。
	//
	// example:
	//
	// 169.254.10.1
	PeerBgpIp *string `json:"PeerBgpIp,omitempty" xml:"PeerBgpIp,omitempty"`
	// BGP路由协议的协商状态。
	//
	// - **success**：正常。
	//
	// - **false**：异常。
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// IPsec连接BGP网段。该网段是一个在169.254.0.0/16内的子网掩码长度为30的网段。
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
