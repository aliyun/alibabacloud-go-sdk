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
	SetGatewayType(v string) *DescribeVpnGatewayResponseBody
	GetGatewayType() *string
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
	// Indicates whether automatic propagation is enabled for the VPN gateway. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	AutoPropagate *bool `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	// The payment status of the VPN gateway. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **FinancialLocked**: locked due to overdue payment.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method. Value:
	//
	// <props="intl">**POSTPAY**: pay-as-you-go billing method.
	//
	// <props="partner">**POSTPAY**: pay-as-you-go billing method.
	//
	// <props="china">**Prepay**: subscription.
	//
	// example:
	//
	// China site example: Prepay, International site example: POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The timestamp when the VPN gateway was created. Unit: milliseconds.
	//
	// The timestamp follows the UNIX time format, which represents the total number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
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
	// The second IP address assigned by the system to the VPN gateway instance for creating IPsec-VPN connections.
	//
	// This parameter is returned only for VPN gateway instances that support creating dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// 47.91.XX.XX
	DisasterRecoveryInternetIp *string `json:"DisasterRecoveryInternetIp,omitempty" xml:"DisasterRecoveryInternetIp,omitempty"`
	// The ID of the second vSwitch associated with the VPN gateway instance.
	//
	// This parameter is returned only for VPN gateway instances that support creating dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// vsw-p0w95ql6tmr2ludkt****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// The enabling status of the BGP feature for the VPN gateway. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	EnableBgp *bool `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	// The timestamp when the VPN gateway expires. Unit: milliseconds.
	//
	// The timestamp follows the UNIX time format, which represents the total number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1544666102000
	EndTime        *int64                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EniInstanceIds *DescribeVpnGatewayResponseBodyEniInstanceIds `json:"EniInstanceIds,omitempty" xml:"EniInstanceIds,omitempty" type:"Struct"`
	// The type of the VPN gateway. Valid values:
	//
	// - **Traditional**: traditional VPN gateway that supports both IPsec and SSL features.
	//
	// - **Enhanced.SiteToSite**: enhanced site-to-cloud VPN gateway that supports only the IPsec feature.
	//
	// example:
	//
	// Enhanced.SiteToSite
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// - If the VPN gateway instance supports creating single-tunnel IPsec-VPN connections, this address is the IP address of the VPN gateway instance and can be used to create IPsec-VPN connections or SSL-VPN connections.
	//
	// - If the VPN gateway instance supports creating dual-tunnel IPsec-VPN connections, this address is the first IP address used to create IPsec-VPN connections and cannot be used to create SSL-VPN connections.
	//
	//     If the VPN gateway instance supports creating dual-tunnel IPsec-VPN connections, the system assigns two IPsec IP addresses to the VPN gateway instance for creating dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// 47.22.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Indicates whether IPsec-VPN is enabled. Valid values:
	//
	// - **enable**: enabled.
	//
	// - **disable**: disabled.
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
	// - **public**: public VPN gateway.
	//
	// - **private**: private VPN gateway.
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
	// The pending order data.
	//
	// >This parameter is returned only when **IncludeReservationData*	- is set to **true**.
	ReservationData *DescribeVpnGatewayResponseBodyReservationData `json:"ReservationData,omitempty" xml:"ReservationData,omitempty" type:"Struct"`
	// The ID of the resource group to which the VPN gateway belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource group information.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The bandwidth specification of the VPN gateway. Unit: Mbit/s.
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
	// The enabling status of the SSL-VPN feature. Valid values:
	//
	// - **enable**: enabled.
	//
	// - **disable**: disabled.
	//
	// example:
	//
	// enable
	SslVpn *string `json:"SslVpn,omitempty" xml:"SslVpn,omitempty"`
	// The IP address for SSL-VPN connections.
	//
	// This parameter is returned only when the SSL-VPN feature is enabled on a VPN gateway instance that has a public network type and supports creating dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// 47.74.XX.XX
	SslVpnInternetIp *string `json:"SslVpnInternetIp,omitempty" xml:"SslVpnInternetIp,omitempty"`
	// The status of the VPN gateway. Valid values:
	//
	// - **init**: initializing.
	//
	// - **provisioning**: preparing.
	//
	// - **active**: Normal.
	//
	// - **updating**: updating.
	//
	// - **deleting**: deleting.
	//
	// example:
	//
	// init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The system-generated labels of the VPN gateway.
	//
	// - **VpnEnableBgp**: indicates whether the VPN gateway supports the BGP feature.
	//
	//     - **true**: supported.
	//
	//     - **false**: not supported.
	//
	// - **VisuallySsl**: indicates whether the VPN gateway supports viewing SSL client connection information.
	//
	//     - **true**: supported.
	//
	//     - **false**: not supported.
	//
	// - **PbrPriority**: indicates whether the VPN gateway supports configuring policy priority for policy-based routing.
	//
	//     - **true**: supported.
	//
	//     - **false**: not supported.
	//
	// - **VpnNewImage**: indicates whether the VPN gateway is a new-generation VPN gateway.
	//
	//     - **true**: yes.
	//
	//     - **false**: no.
	//
	// - **description**: the description of the VPN gateway, which is used only for internal system purposes.
	//
	// - **VpnVersion**: the version number of the VPN gateway.
	//
	// - **IDaaSNewVersion**: indicates whether the VPN gateway supports attaching to an EIAM 2.0 instance.
	//
	//     - **true**: supported.
	//
	//     - **false**: not supported.
	//
	// example:
	//
	// {\\"VpnEnableBgp\\":\\"true\\",\\"VisuallySsl\\":\\"true\\",\\"PbrPriority\\":\\"true\\",\\"VpnNewImage\\":\\"true\\",\\"description\\":\\"Forwarding 1.3.24\\",\\"VpnVersion\\":\\"v1.2.4\\",\\"IDaaSNewVersion\\":\\"true\\"}
	Tag  *string                             `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
	// The instance ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1r3v1xqkl0w519g****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// <props="intl">The type of the VPN gateway. Value: **Normal**, which indicates a standard VPN gateway.
	//
	// <props="china">
	//
	// The type of the VPN gateway. Valid values:
	//
	// - **Normal**: standard.
	//
	// - **NationalStandard**: Chinese SM-based.
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

func (s *DescribeVpnGatewayResponseBody) GetGatewayType() *string {
	return s.GatewayType
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

func (s *DescribeVpnGatewayResponseBody) SetGatewayType(v string) *DescribeVpnGatewayResponseBody {
	s.GatewayType = &v
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
	// If the pending order type is **TEMP_UPGRADE*	- (temporary upgrade), this parameter indicates the revert time for the temporary upgrade.
	//
	// If the pending order type is **RENEWCHANGE*	- (renewal with specification change) or **RENEW*	- (renewal), this parameter indicates the effective period when the renewal or renewal with specification change takes effect.
	//
	// example:
	//
	// 2020-07-20T16:00:00Z
	ReservationEndTime *string `json:"ReservationEndTime,omitempty" xml:"ReservationEndTime,omitempty"`
	// The enabling status of the IPsec-VPN feature for the pending order. Valid values:
	//
	// - **enable**: enabled.
	//
	// - **disable**: disabled.
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
	// - **RENEWCHANGE**: renewal with specification change.
	//
	// - **TEMP_UPGRADE**: temporary upgrade.
	//
	// - **RENEW**: renewal.
	//
	// example:
	//
	// TEMP_UPGRADE
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// The bandwidth specification of the pending order. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	ReservationSpec *string `json:"ReservationSpec,omitempty" xml:"ReservationSpec,omitempty"`
	// The enabling status of the SSL-VPN feature for the pending order. Valid values:
	//
	// - **enable**: enabled.
	//
	// - **disable**: disabled.
	//
	// example:
	//
	// enable
	ReservationSsl *string `json:"ReservationSsl,omitempty" xml:"ReservationSsl,omitempty"`
	// The status of the pending order. Valid values:
	//
	// - **1**: the renewal or renewal with specification change order has not taken effect.
	//
	// - **2**: the temporary upgrade order has taken effect. After the restoration time is reached, the system restores the VPN gateway to the specification before the temporary upgrade. In this case, **ReservationIpsec**, **ReservationMaxConnections**, **ReservationSpec**, and **ReservationSsl*	- indicate the specifications before the temporary upgrade.
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

type DescribeVpnGatewayResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
