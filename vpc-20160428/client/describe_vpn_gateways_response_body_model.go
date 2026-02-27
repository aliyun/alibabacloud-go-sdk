// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpnGatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnGatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnGatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpnGatewaysResponseBody
	GetTotalCount() *int32
	SetVpnGateways(v *DescribeVpnGatewaysResponseBodyVpnGateways) *DescribeVpnGatewaysResponseBody
	GetVpnGateways() *DescribeVpnGatewaysResponseBodyVpnGateways
}

type DescribeVpnGatewaysResponseBody struct {
	// The number of the returned page.
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
	// DF11D6F6-E35A-41C3-9B20-6FC8A901FE65
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpnGateways *DescribeVpnGatewaysResponseBodyVpnGateways `json:"VpnGateways,omitempty" xml:"VpnGateways,omitempty" type:"Struct"`
}

func (s DescribeVpnGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnGatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpnGatewaysResponseBody) GetVpnGateways() *DescribeVpnGatewaysResponseBodyVpnGateways {
	return s.VpnGateways
}

func (s *DescribeVpnGatewaysResponseBody) SetPageNumber(v int32) *DescribeVpnGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) SetPageSize(v int32) *DescribeVpnGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) SetRequestId(v string) *DescribeVpnGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) SetTotalCount(v int32) *DescribeVpnGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) SetVpnGateways(v *DescribeVpnGatewaysResponseBodyVpnGateways) *DescribeVpnGatewaysResponseBody {
	s.VpnGateways = v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) Validate() error {
	if s.VpnGateways != nil {
		if err := s.VpnGateways.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnGatewaysResponseBodyVpnGateways struct {
	VpnGateway []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway `json:"VpnGateway,omitempty" xml:"VpnGateway,omitempty" type:"Repeated"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGateways) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGateways) GetVpnGateway() []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	return s.VpnGateway
}

func (s *DescribeVpnGatewaysResponseBodyVpnGateways) SetVpnGateway(v []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) *DescribeVpnGatewaysResponseBodyVpnGateways {
	s.VpnGateway = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGateways) Validate() error {
	if s.VpnGateway != nil {
		for _, item := range s.VpnGateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway struct {
	// 是否已开启VPN网关的路由自动传播功能。
	//
	// - **true**：已开启。
	//
	// - **false**：未开启。
	//
	// example:
	//
	// true
	AutoPropagate *bool `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	// VPN网关的付费状态。
	//
	// - **Normal**：正常。
	//
	// - **FinancialLocked**：欠费锁定。
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// VPN网关的付费类型。
	//
	// <props="china">仅取值：**Prepay**，包年包月。
	//
	// <props="intl">仅取值：**POSTPAY**，按量计费。
	//
	// <props="partner">仅取值： **POSTPAY**，按量计费。
	//
	// example:
	//
	// 中国站示例值：Prepay，国际站示例值：POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 创建VPN网关的时间戳。单位：毫秒。
	//
	// 时间戳的格式采用Unix时间戳，表示从格林威治时间1970年01月01日00时00分00秒至创建VPN网关实例时的总时长。
	//
	// example:
	//
	// 1515383700000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// VPN网关的描述信息。
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 系统为VPN网关实例分配的用于创建IPsec-VPN连接的第二个IP地址。
	//
	// 仅支持创建双隧道模式IPsec-VPN连接的VPN网关实例会返回当前参数。
	//
	// example:
	//
	// 47.91.XX.XX
	DisasterRecoveryInternetIp *string `json:"DisasterRecoveryInternetIp,omitempty" xml:"DisasterRecoveryInternetIp,omitempty"`
	// VPN网关实例关联的第二个交换机ID。
	//
	// 仅支持创建双隧道模式IPsec-VPN连接的VPN网关实例会返回当前参数。
	//
	// example:
	//
	// vsw-p0w95ql6tmr2ludkt****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// VPN网关BGP功能的开启状态。
	//
	// - **true**：已开启。
	//
	// - **false**：未开启。
	//
	// example:
	//
	// true
	EnableBgp *bool `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	// VPN网关到期时间戳。单位：毫秒。
	//
	// 时间戳的格式采用Unix时间戳，表示从格林威治时间1970年01月01日00时00分00秒至VPN网关实例到期时的总时长。
	//
	// example:
	//
	// 1518105600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 系统为VPN网关实例创建的弹性网卡ENI（Elastic Network Interfaces）列表。
	EniInstanceIds *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds `json:"EniInstanceIds,omitempty" xml:"EniInstanceIds,omitempty" type:"Struct"`
	// VPN 网关类型，取值：
	//
	// Traditional：传统型VPN网关，覆盖IPsec功能和SSL功能
	//
	// Enhance.SiteToSite：增强型站点入云VPN，只覆盖IPsec功能
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// - 在VPN网关实例支持创建单隧道模式IPsec-VPN连接的场景下，该地址为VPN网关实例的IP地址，可用于创建IPsec-VPN连接或SSL-VPN连接。
	//
	// - 在VPN网关实例支持创建双隧道模式IPsec-VPN连接的场景下，该地址为用于创建IPsec-VPN连接的第一个IP地址，不能用于创建SSL-VPN连接。
	//
	//     在VPN网关实例支持创建双隧道模式IPsec-VPN连接的场景下，系统会为VPN网关实例分配两个IPsec地址，用于创建双隧道模式的IPsec-VPN连接。
	//
	// example:
	//
	// 47.12.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// VPN网关是否开启了IPsec-VPN功能。
	//
	// - **enable**：已开启。
	//
	// - **disable**：未开启。
	//
	// example:
	//
	// enable
	IpsecVpn *string `json:"IpsecVpn,omitempty" xml:"IpsecVpn,omitempty"`
	// VPN网关的名称。
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// VPN网关的网络类型。
	//
	// - **public**：公网VPN网关。
	//
	// - **private**：私网VPN网关。
	//
	// example:
	//
	// public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// 未生效的订购数据。
	//
	// >仅**IncludeReservationData**传入**true**才会返回该组参数。
	ReservationData *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData `json:"ReservationData,omitempty" xml:"ReservationData,omitempty" type:"Struct"`
	// VPN网关所属的资源组ID。
	//
	// 您可以调用[ListResourceGroups](https://help.aliyun.com/document_detail/158855.html)接口查询资源组信息。
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// VPN网关的带宽峰值。**M**表示单位Mbps。
	//
	// example:
	//
	// 5M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// VPN网关SSL连接数的规格。
	//
	// example:
	//
	// 5
	SslMaxConnections *int64 `json:"SslMaxConnections,omitempty" xml:"SslMaxConnections,omitempty"`
	// VPN网关是否开启了SSL-VPN功能。
	//
	// - **enable**：已开启。
	//
	// - **disable**：未开启。
	//
	// example:
	//
	// enable
	SslVpn *string `json:"SslVpn,omitempty" xml:"SslVpn,omitempty"`
	// SSL-VPN连接的IP地址。
	//
	// 仅支持创建双隧道模式IPsec-VPN连接的公网网络类型的VPN网关实例开启SSL-VPN功能后，才会返回当前参数。
	//
	// example:
	//
	// 47.74.XX.XX
	SslVpnInternetIp *string `json:"SslVpnInternetIp,omitempty" xml:"SslVpnInternetIp,omitempty"`
	// VPN网关的状态。
	//
	// - **init*	- ：初始化。
	//
	// - **provisioning*	- ：准备中。
	//
	// - **active*	- ：正常。
	//
	// - **updating*	- ：更新中。
	//
	// - **deleting*	- ：删除中。
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 系统自动生成的VPN网关标签。
	//
	// - **VpnEnableBgp**：表示VPN网关是否支持BGP功能。
	//
	//     - **true**：支持。
	//
	//     - **false**：不支持。
	//
	// - **VisuallySsl**：表示VPN网关是否支持查看SSL客户端的连接信息。
	//
	//     - **true**：支持。
	//
	//     - **false**：不支持。
	//
	// - **PbrPriority**：表示VPN网关是否支持为策略路由配置策略优先级。
	//
	//     - **true**：支持。
	//
	//     - **false**：不支持。
	//
	// - **VpnNewImage**：表示VPN网关是否为新型VPN网关。
	//
	//     - **true**：是。
	//
	//     - **false**：否。
	//
	// - **description**：表示VPN网关的描述信息，仅供系统内部使用。
	//
	// - **VpnVersion**：表示VPN网关的版本号。
	//
	// - **IDaaSNewVersion**：表示VPN网关是否支持绑定EIAM 2.0实例。
	//
	//     - **true**：支持。
	//
	//     - **false**：不支持。
	//
	// example:
	//
	// {\\"VpnEnableBgp\\":\\"true\\",\\"VisuallySsl\\":\\"true\\",\\"PbrPriority\\":\\"true\\",\\"VpnNewImage\\":\\"true\\",\\"description\\":\\"转发1.3.24\\",\\"VpnVersion\\":\\"v1.2.4\\",\\"IDaaSNewVersion\\":\\"true\\"}
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// VPN网关绑定的标签列表。
	Tags *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// VPN网关所属交换机的ID。
	//
	// example:
	//
	// vsw-bp15lbk8sgtr6r5b0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPN网关所属VPC的ID。
	//
	// example:
	//
	// vpc-bp1m3i0kn1nd4wiw9****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// VPN网关的ID。
	//
	// example:
	//
	// vpn-bp17lofy9fd0dnvzv****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// VPN网关类型。
	//
	//
	//
	// <props="china">
	//
	// - **Normal**：普通型。
	//
	// - **NationalStandard**：国密型。
	//
	//
	//
	// <props="intl">取值：**Normal**，表示普通型。
	//
	// example:
	//
	// Normal
	VpnType *string `json:"VpnType,omitempty" xml:"VpnType,omitempty"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetAutoPropagate() *bool {
	return s.AutoPropagate
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetDisasterRecoveryInternetIp() *string {
	return s.DisasterRecoveryInternetIp
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetDisasterRecoveryVSwitchId() *string {
	return s.DisasterRecoveryVSwitchId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetEnableBgp() *bool {
	return s.EnableBgp
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetEniInstanceIds() *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds {
	return s.EniInstanceIds
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetGatewayType() *string {
	return s.GatewayType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetIpsecVpn() *string {
	return s.IpsecVpn
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetName() *string {
	return s.Name
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetReservationData() *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	return s.ReservationData
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetSpec() *string {
	return s.Spec
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetSslMaxConnections() *int64 {
	return s.SslMaxConnections
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetSslVpn() *string {
	return s.SslVpn
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetSslVpnInternetIp() *string {
	return s.SslVpnInternetIp
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetTag() *string {
	return s.Tag
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetTags() *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags {
	return s.Tags
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetVpnType() *string {
	return s.VpnType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetAutoPropagate(v bool) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.AutoPropagate = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetBusinessStatus(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetChargeType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.ChargeType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetCreateTime(v int64) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetDescription(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Description = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetDisasterRecoveryInternetIp(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.DisasterRecoveryInternetIp = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetDisasterRecoveryVSwitchId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.DisasterRecoveryVSwitchId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetEnableBgp(v bool) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.EnableBgp = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetEndTime(v int64) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.EndTime = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetEniInstanceIds(v *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.EniInstanceIds = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetGatewayType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.GatewayType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetInternetIp(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.InternetIp = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetIpsecVpn(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.IpsecVpn = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetName(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Name = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetNetworkType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.NetworkType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetReservationData(v *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.ReservationData = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetResourceGroupId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetSpec(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Spec = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetSslMaxConnections(v int64) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.SslMaxConnections = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetSslVpn(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.SslVpn = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetSslVpnInternetIp(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.SslVpnInternetIp = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetStatus(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Status = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetTag(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Tag = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetTags(v *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Tags = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetVSwitchId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetVpcId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.VpcId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetVpnGatewayId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetVpnType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.VpnType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) Validate() error {
	if s.EniInstanceIds != nil {
		if err := s.EniInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.ReservationData != nil {
		if err := s.ReservationData.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds struct {
	EniInstanceId []*string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) GetEniInstanceId() []*string {
	return s.EniInstanceId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) SetEniInstanceId(v []*string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds {
	s.EniInstanceId = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData struct {
	// 如果未生效订单类型为**TEMP_UPGRADE**（临时升配）时，该参数表示为临时升配的还原时间。
	//
	// 如果未生效订单类型为**RENEWCHANGE**（续费变配）或**RENEW**（续费）时，该参数表示为续费或续费变配开始生效时间。
	//
	// example:
	//
	// 2021-07-20T16:00:00Z
	ReservationEndTime *string `json:"ReservationEndTime,omitempty" xml:"ReservationEndTime,omitempty"`
	// 未生效订单IPsec-VPN功能开启状态。
	//
	// - **enable**：已开启。
	//
	// - **disable**：未开启。
	//
	// example:
	//
	// enable
	ReservationIpsec *string `json:"ReservationIpsec,omitempty" xml:"ReservationIpsec,omitempty"`
	// 未生效订单SSL-VPN并发连接用户数的规格。
	//
	// example:
	//
	// 5
	ReservationMaxConnections *int32 `json:"ReservationMaxConnections,omitempty" xml:"ReservationMaxConnections,omitempty"`
	// 未生效订单类型。
	//
	// - **RENEWCHANGE**：续费变配。
	//
	// - **TEMP_UPGRADE**：临时升配。
	//
	// - **RENEW**：续费。
	//
	// example:
	//
	// TEMP_UPGRADE
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// 未生效订单的带宽规格。单位：Mbps。
	//
	// example:
	//
	// 5
	ReservationSpec *string `json:"ReservationSpec,omitempty" xml:"ReservationSpec,omitempty"`
	// 未生效订单SSL-VPN功能开启状态。
	//
	// - **enable**：已开启。
	//
	// - **disable**：未开启。
	//
	// example:
	//
	// enable
	ReservationSsl *string `json:"ReservationSsl,omitempty" xml:"ReservationSsl,omitempty"`
	// 未生效订单状态。
	//
	// - **1**：表示续费或续费变配的订单未生效。
	//
	// - **2**：表示临时升配的订单已生效。在到达还原时间后，系统会将VPN网关规格恢复到临时升配前的规格。此时**ReservationIpsec**、**ReservationMaxConnections**、**ReservationSpec**、**ReservationSsl**表示为VPN网关临时升配前的规格。
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationEndTime() *string {
	return s.ReservationEndTime
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationIpsec() *string {
	return s.ReservationIpsec
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationMaxConnections() *int32 {
	return s.ReservationMaxConnections
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationSpec() *string {
	return s.ReservationSpec
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationSsl() *string {
	return s.ReservationSsl
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationEndTime(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationEndTime = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationIpsec(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationIpsec = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationMaxConnections(v int32) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationMaxConnections = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationOrderType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationSpec(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationSpec = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationSsl(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationSsl = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetStatus(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.Status = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags struct {
	Tag []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) GetTag() []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag {
	return s.Tag
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) SetTag(v []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags {
	s.Tag = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) Validate() error {
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

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag struct {
	// 标签键。
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) SetKey(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) SetValue(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) Validate() error {
	return dara.Validate(s)
}
