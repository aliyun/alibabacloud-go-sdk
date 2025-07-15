// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPropagate(v bool) *DescribeVpnGatewayResponseBody
	GetAutoPropagate() *bool
	SetBusinessStatus(v string) *DescribeVpnGatewayResponseBody
	GetBusinessStatus() *string
	SetChargeType(v string) *DescribeVpnGatewayResponseBody
	GetChargeType() *string
	SetCreateTime(v int64) *DescribeVpnGatewayResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *DescribeVpnGatewayResponseBody
	GetDescription() *string
	SetDisasterRecoveryInternetIp(v string) *DescribeVpnGatewayResponseBody
	GetDisasterRecoveryInternetIp() *string
	SetDisasterRecoveryVSwitchId(v string) *DescribeVpnGatewayResponseBody
	GetDisasterRecoveryVSwitchId() *string
	SetEnableBgp(v bool) *DescribeVpnGatewayResponseBody
	GetEnableBgp() *bool
	SetEndTime(v int64) *DescribeVpnGatewayResponseBody
	GetEndTime() *int64
	SetEniInstanceIds(v *DescribeVpnGatewayResponseBodyEniInstanceIds) *DescribeVpnGatewayResponseBody
	GetEniInstanceIds() *DescribeVpnGatewayResponseBodyEniInstanceIds
	SetInternetIp(v string) *DescribeVpnGatewayResponseBody
	GetInternetIp() *string
	SetIpsecVpn(v string) *DescribeVpnGatewayResponseBody
	GetIpsecVpn() *string
	SetName(v string) *DescribeVpnGatewayResponseBody
	GetName() *string
	SetNetworkType(v string) *DescribeVpnGatewayResponseBody
	GetNetworkType() *string
	SetRequestId(v string) *DescribeVpnGatewayResponseBody
	GetRequestId() *string
	SetReservationData(v *DescribeVpnGatewayResponseBodyReservationData) *DescribeVpnGatewayResponseBody
	GetReservationData() *DescribeVpnGatewayResponseBodyReservationData
	SetResourceGroupId(v string) *DescribeVpnGatewayResponseBody
	GetResourceGroupId() *string
	SetSpec(v string) *DescribeVpnGatewayResponseBody
	GetSpec() *string
	SetSslMaxConnections(v int64) *DescribeVpnGatewayResponseBody
	GetSslMaxConnections() *int64
	SetSslVpn(v string) *DescribeVpnGatewayResponseBody
	GetSslVpn() *string
	SetSslVpnInternetIp(v string) *DescribeVpnGatewayResponseBody
	GetSslVpnInternetIp() *string
	SetStatus(v string) *DescribeVpnGatewayResponseBody
	GetStatus() *string
	SetTag(v string) *DescribeVpnGatewayResponseBody
	GetTag() *string
	SetTags(v *DescribeVpnGatewayResponseBodyTags) *DescribeVpnGatewayResponseBody
	GetTags() *DescribeVpnGatewayResponseBodyTags
	SetVSwitchId(v string) *DescribeVpnGatewayResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeVpnGatewayResponseBody
	GetVpcId() *string
	SetVpnGatewayId(v string) *DescribeVpnGatewayResponseBody
	GetVpnGatewayId() *string
	SetVpnType(v string) *DescribeVpnGatewayResponseBody
	GetVpnType() *string
}

type DescribeVpnGatewayResponseBody struct {
	// Indicates whether BGP routes are automatically advertised to the VPC. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoPropagate *bool `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	// The payment status of the VPN gateway. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method. Valid value:
	//
	// **POSTPAY**: pay-as-you-go
	//
	// example:
	//
	// China site (aliyun.com): Prepay. International site (alibabacloud.com): POSTPAY.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The timestamp when the VPN gateway was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1495382400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the VPN gateway.
	//
	// example:
	//
	// vpngatewaydescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The second IP address assigned by the system to create an IPsec-VPN connection.
	//
	// This parameter is returned only when the VPN gateway supports the dual-tunnel mode.
	//
	// example:
	//
	// 47.91.XX.XX
	DisasterRecoveryInternetIp *string `json:"DisasterRecoveryInternetIp,omitempty" xml:"DisasterRecoveryInternetIp,omitempty"`
	// The ID of the second vSwitch associated with the VPN gateway.
	//
	// This parameter is returned only when the VPN gateway supports the dual-tunnel mode.
	//
	// example:
	//
	// vsw-p0w95ql6tmr2ludkt****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// Indicates whether BGP is enabled for the VPN gateway. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableBgp *bool `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	// The timestamp when the VPN gateway expires. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1544666102000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ENIs created by the system for the VPN gateway.
	EniInstanceIds *DescribeVpnGatewayResponseBodyEniInstanceIds `json:"EniInstanceIds,omitempty" xml:"EniInstanceIds,omitempty" type:"Struct"`
	// 	- If the VPN gateway supports IPsec-VPN connections in single-tunnel mode, the address is the IP address of the VPN gateway and can be used to create an IPsec-VPN connection or an SSL-VPN connection.
	//
	// 	- If the VPN gateway supports IPsec-VPN connections in dual-tunnel mode, the address is the first IP address used to create an IPsec-VPN connection. The address cannot be used to create an SSL-VPN connection.
	//
	//     If the VPN gateway supports IPsec-VPN connections in dual-tunnel mode, the system assigns two IP addresses to the VPN gateway to create two encrypted tunnels.
	//
	// example:
	//
	// 47.22.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Indicates whether the IPsec-VPN feature is enabled. Valid values:
	//
	// 	- **enable**
	//
	// 	- **disable**
	//
	// example:
	//
	// enable
	IpsecVpn *string `json:"IpsecVpn,omitempty" xml:"IpsecVpn,omitempty"`
	// The name of the VPN gateway.
	//
	// example:
	//
	// vpngatewayname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the VPN gateway.
	//
	// 	- **public**
	//
	// 	- **private**
	//
	// example:
	//
	// public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 27E4E088-8DE0-4672-BF5C-0A412389DB9E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about pending orders.
	//
	// > This set of parameters is returned only when **IncludeReservationData*	- is set to **true**.
	ReservationData *DescribeVpnGatewayResponseBodyReservationData `json:"ReservationData,omitempty" xml:"ReservationData,omitempty" type:"Struct"`
	// The ID of the resource group to which the VPN gateway belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The maximum bandwidth of the VPN gateway. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The maximum number of concurrent SSL-VPN connections.
	//
	// example:
	//
	// 5
	SslMaxConnections *int64 `json:"SslMaxConnections,omitempty" xml:"SslMaxConnections,omitempty"`
	// The status of the SSL-VPN feature. Valid values:
	//
	// 	- **enable**
	//
	// 	- **disable**
	//
	// example:
	//
	// enable
	SslVpn *string `json:"SslVpn,omitempty" xml:"SslVpn,omitempty"`
	// The IP address of the SSL-VPN connection.
	//
	// This parameter is returned only when the VPN gateway is a public VPN gateway and supports only the single-tunnel mode. In addition, the VPN gateway must have the SSL-VPN feature enabled.
	//
	// example:
	//
	// 47.74.XX.XX
	SslVpnInternetIp *string `json:"SslVpnInternetIp,omitempty" xml:"SslVpnInternetIp,omitempty"`
	// The status of the VPN gateway. Valid values:
	//
	// 	- **init**
	//
	// 	- **provisioning**
	//
	// 	- **active**
	//
	// 	- **updating**
	//
	// 	- **deleting**
	//
	// example:
	//
	// init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is automatically generated for the VPN gateway. The tag consists of the following parameters:
	//
	// 	- **VpnEnableBgp**: indicates whether the VPN gateway supports BGP. Valid values:
	//
	//     	- **true**
	//
	//     	- **false**
	//
	// 	- **VisuallySsl**: indicates whether the VPN gateway allows you to view the connection information of SSL clients. Valid values:
	//
	//     	- **true**
	//
	//     	- **false**
	//
	// 	- **PbrPriority**: indicates whether the VPN gateway allows you to configure priorities for policy-based routes. Valid values:
	//
	//     	- **true**
	//
	//     	- **false**
	//
	// 	- **VpnNewImage**: indicates whether the VPN gateway is upgraded. Valid values:
	//
	//     	- **true**
	//
	//     	- **false**
	//
	// 	- **description**: the description of the VPN gateway. This parameter is only for internal use.
	//
	// 	- **VpnVersion**: the version of the VPN gateway.
	//
	// 	- **IDaaSNewVersion**: indicates whether the VPN gateway can be associated with an EIAM 2.0 instance.
	//
	//     	- **true**
	//
	//     	- **false**
	//
	// example:
	//
	// {\\"VpnEnableBgp\\":\\"true\\",\\"VisuallySsl\\":\\"true\\",\\"PbrPriority\\":\\"true\\",\\"VpnNewImage\\":\\"true\\",\\"description\\":\\"forwarding1.3.7\\",\\"VpnVersion\\":\\"v1.2.4\\"}
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The tags that are added to the VPN gateway.
	Tags *DescribeVpnGatewayResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vSwitch to which the VPN gateway belongs.
	//
	// example:
	//
	// vsw-bp1dmzugdikc6hdgx****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the VPN gateway belongs.
	//
	// example:
	//
	// vpc-bp19m2yx1m5q0avyq****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1r3v1xqkl0w519g****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// The type of VPN gateway. Only **Normal*	- may be returned, which indicates a standard VPN gateway.
	//
	// example:
	//
	// Normal
	VpnType *string `json:"VpnType,omitempty" xml:"VpnType,omitempty"`
}

func (s DescribeVpnGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayResponseBody) GetAutoPropagate() *bool {
	return s.AutoPropagate
}

func (s *DescribeVpnGatewayResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeVpnGatewayResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeVpnGatewayResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpnGatewayResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpnGatewayResponseBody) GetDisasterRecoveryInternetIp() *string {
	return s.DisasterRecoveryInternetIp
}

func (s *DescribeVpnGatewayResponseBody) GetDisasterRecoveryVSwitchId() *string {
	return s.DisasterRecoveryVSwitchId
}

func (s *DescribeVpnGatewayResponseBody) GetEnableBgp() *bool {
	return s.EnableBgp
}

func (s *DescribeVpnGatewayResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeVpnGatewayResponseBody) GetEniInstanceIds() *DescribeVpnGatewayResponseBodyEniInstanceIds {
	return s.EniInstanceIds
}

func (s *DescribeVpnGatewayResponseBody) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVpnGatewayResponseBody) GetIpsecVpn() *string {
	return s.IpsecVpn
}

func (s *DescribeVpnGatewayResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeVpnGatewayResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeVpnGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnGatewayResponseBody) GetReservationData() *DescribeVpnGatewayResponseBodyReservationData {
	return s.ReservationData
}

func (s *DescribeVpnGatewayResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpnGatewayResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *DescribeVpnGatewayResponseBody) GetSslMaxConnections() *int64 {
	return s.SslMaxConnections
}

func (s *DescribeVpnGatewayResponseBody) GetSslVpn() *string {
	return s.SslVpn
}

func (s *DescribeVpnGatewayResponseBody) GetSslVpnInternetIp() *string {
	return s.SslVpnInternetIp
}

func (s *DescribeVpnGatewayResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnGatewayResponseBody) GetTag() *string {
	return s.Tag
}

func (s *DescribeVpnGatewayResponseBody) GetTags() *DescribeVpnGatewayResponseBodyTags {
	return s.Tags
}

func (s *DescribeVpnGatewayResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVpnGatewayResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpnGatewayResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnGatewayResponseBody) GetVpnType() *string {
	return s.VpnType
}

func (s *DescribeVpnGatewayResponseBody) SetAutoPropagate(v bool) *DescribeVpnGatewayResponseBody {
	s.AutoPropagate = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetBusinessStatus(v string) *DescribeVpnGatewayResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetChargeType(v string) *DescribeVpnGatewayResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetCreateTime(v int64) *DescribeVpnGatewayResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetDescription(v string) *DescribeVpnGatewayResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetDisasterRecoveryInternetIp(v string) *DescribeVpnGatewayResponseBody {
	s.DisasterRecoveryInternetIp = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetDisasterRecoveryVSwitchId(v string) *DescribeVpnGatewayResponseBody {
	s.DisasterRecoveryVSwitchId = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetEnableBgp(v bool) *DescribeVpnGatewayResponseBody {
	s.EnableBgp = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetEndTime(v int64) *DescribeVpnGatewayResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetEniInstanceIds(v *DescribeVpnGatewayResponseBodyEniInstanceIds) *DescribeVpnGatewayResponseBody {
	s.EniInstanceIds = v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetInternetIp(v string) *DescribeVpnGatewayResponseBody {
	s.InternetIp = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetIpsecVpn(v string) *DescribeVpnGatewayResponseBody {
	s.IpsecVpn = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetName(v string) *DescribeVpnGatewayResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetNetworkType(v string) *DescribeVpnGatewayResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetRequestId(v string) *DescribeVpnGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetReservationData(v *DescribeVpnGatewayResponseBodyReservationData) *DescribeVpnGatewayResponseBody {
	s.ReservationData = v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetResourceGroupId(v string) *DescribeVpnGatewayResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetSpec(v string) *DescribeVpnGatewayResponseBody {
	s.Spec = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetSslMaxConnections(v int64) *DescribeVpnGatewayResponseBody {
	s.SslMaxConnections = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetSslVpn(v string) *DescribeVpnGatewayResponseBody {
	s.SslVpn = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetSslVpnInternetIp(v string) *DescribeVpnGatewayResponseBody {
	s.SslVpnInternetIp = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetStatus(v string) *DescribeVpnGatewayResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetTag(v string) *DescribeVpnGatewayResponseBody {
	s.Tag = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetTags(v *DescribeVpnGatewayResponseBodyTags) *DescribeVpnGatewayResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetVSwitchId(v string) *DescribeVpnGatewayResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetVpcId(v string) *DescribeVpnGatewayResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetVpnGatewayId(v string) *DescribeVpnGatewayResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) SetVpnType(v string) *DescribeVpnGatewayResponseBody {
	s.VpnType = &v
	return s
}

func (s *DescribeVpnGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewayResponseBodyEniInstanceIds struct {
	EniInstanceId []*string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeVpnGatewayResponseBodyEniInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayResponseBodyEniInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayResponseBodyEniInstanceIds) GetEniInstanceId() []*string {
	return s.EniInstanceId
}

func (s *DescribeVpnGatewayResponseBodyEniInstanceIds) SetEniInstanceId(v []*string) *DescribeVpnGatewayResponseBodyEniInstanceIds {
	s.EniInstanceId = v
	return s
}

func (s *DescribeVpnGatewayResponseBodyEniInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewayResponseBodyReservationData struct {
	// If the order type is **TEMP_UPGRADE*	- (temporary upgrade), this parameter specifies the time when the temporary upgrade expires.
	//
	// If the order type is **RENEWCHANGE*	- (renewal with a specification change) or **RENEW*	- (renewal), this parameter indicates the time when the renewal or renewal with a specification change takes effect.
	//
	// example:
	//
	// 2020-07-20T16:00:00Z
	ReservationEndTime *string `json:"ReservationEndTime,omitempty" xml:"ReservationEndTime,omitempty"`
	// The IPsec-VPN status of the pending order. Valid values:
	//
	// 	- **enable**
	//
	// 	- **disable**
	//
	// example:
	//
	// enable
	ReservationIpsec *string `json:"ReservationIpsec,omitempty" xml:"ReservationIpsec,omitempty"`
	// The maximum number of concurrent SSL-VPN connections of the pending order.
	//
	// example:
	//
	// 5
	ReservationMaxConnections *int32 `json:"ReservationMaxConnections,omitempty" xml:"ReservationMaxConnections,omitempty"`
	// The type of the pending order. Valid values:
	//
	// 	- **RENEWCHANGE**: renewal with upgrade or downgrade
	//
	// 	- **TEMP_UPGRADE**: temporary upgrade
	//
	// 	- **RENEW**: renewal
	//
	// example:
	//
	// TEMP_UPGRADE
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// The bandwidth of the pending order. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	ReservationSpec *string `json:"ReservationSpec,omitempty" xml:"ReservationSpec,omitempty"`
	// The SSL-VPN status of the pending order. Valid values:
	//
	// 	- **enable**
	//
	// 	- **disable**
	//
	// example:
	//
	// enable
	ReservationSsl *string `json:"ReservationSsl,omitempty" xml:"ReservationSsl,omitempty"`
	// The status of the pending order. Valid values:
	//
	// 	- **1**: indicates that the order of the renewal or specification change has not taken effect.
	//
	// 	- **2**: indicates that the order is an order for temporary upgrade and the order has taken effect. After the temporary upgrade expires, the system restores the VPN gateway to its previous specifications. In this case, **ReservationIpsec**, **ReservationMaxConnections**, **ReservationSpec**, and **ReservationSsl*	- indicate the previous specification.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVpnGatewayResponseBodyReservationData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayResponseBodyReservationData) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayResponseBodyReservationData) GetReservationEndTime() *string {
	return s.ReservationEndTime
}

func (s *DescribeVpnGatewayResponseBodyReservationData) GetReservationIpsec() *string {
	return s.ReservationIpsec
}

func (s *DescribeVpnGatewayResponseBodyReservationData) GetReservationMaxConnections() *int32 {
	return s.ReservationMaxConnections
}

func (s *DescribeVpnGatewayResponseBodyReservationData) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeVpnGatewayResponseBodyReservationData) GetReservationSpec() *string {
	return s.ReservationSpec
}

func (s *DescribeVpnGatewayResponseBodyReservationData) GetReservationSsl() *string {
	return s.ReservationSsl
}

func (s *DescribeVpnGatewayResponseBodyReservationData) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnGatewayResponseBodyReservationData) SetReservationEndTime(v string) *DescribeVpnGatewayResponseBodyReservationData {
	s.ReservationEndTime = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyReservationData) SetReservationIpsec(v string) *DescribeVpnGatewayResponseBodyReservationData {
	s.ReservationIpsec = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyReservationData) SetReservationMaxConnections(v int32) *DescribeVpnGatewayResponseBodyReservationData {
	s.ReservationMaxConnections = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyReservationData) SetReservationOrderType(v string) *DescribeVpnGatewayResponseBodyReservationData {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyReservationData) SetReservationSpec(v string) *DescribeVpnGatewayResponseBodyReservationData {
	s.ReservationSpec = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyReservationData) SetReservationSsl(v string) *DescribeVpnGatewayResponseBodyReservationData {
	s.ReservationSsl = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyReservationData) SetStatus(v string) *DescribeVpnGatewayResponseBodyReservationData {
	s.Status = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyReservationData) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewayResponseBodyTags struct {
	Tag []*DescribeVpnGatewayResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVpnGatewayResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayResponseBodyTags) GetTag() []*DescribeVpnGatewayResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeVpnGatewayResponseBodyTags) SetTag(v []*DescribeVpnGatewayResponseBodyTagsTag) *DescribeVpnGatewayResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeVpnGatewayResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewayResponseBodyTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// aaa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// bbb
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpnGatewayResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpnGatewayResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpnGatewayResponseBodyTagsTag) SetKey(v string) *DescribeVpnGatewayResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyTagsTag) SetValue(v string) *DescribeVpnGatewayResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVpnGatewayResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
