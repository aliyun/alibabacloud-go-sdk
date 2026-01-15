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
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the VPN gateways.
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
	// Indicates whether BGP routes are automatically advertised to the VPC.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoPropagate *bool `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	// The payment status of the VPN gateway.
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method of the VPN gateway.
	//
	// Only **POSTPAY*	- may be returned, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// Example value for the China site (aliyun.com): Prepay. Example value for the International site (alibabacloud.com): POSTPAY.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The timestamp generated when the VPN gateway was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1515383700000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the VPN gateway.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The second IP address that is assigned to the VPN gateway to create IPsec-VPN connections.
	//
	// This parameter is returned only if the VPN gateway supports IPsec-VPN connections in dual-tunnel mode.
	//
	// example:
	//
	// 47.91.XX.XX
	DisasterRecoveryInternetIp *string `json:"DisasterRecoveryInternetIp,omitempty" xml:"DisasterRecoveryInternetIp,omitempty"`
	// The ID of the second vSwitch that is associated with the VPN gateway.
	//
	// This parameter is returned only if the VPN gateway supports IPsec-VPN connections in dual-tunnel mode.
	//
	// example:
	//
	// vsw-p0w95ql6tmr2ludkt****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// The BGP status of the VPN gateway. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableBgp *bool `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	// The timestamp generated when the VPN gateway expires. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1518105600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ENIs created by the system for the VPN gateway.
	EniInstanceIds *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds `json:"EniInstanceIds,omitempty" xml:"EniInstanceIds,omitempty" type:"Struct"`
	GatewayType    *string                                                             `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// 	- If the VPN gateway supports IPsec-VPN connections in single-tunnel mode, the value of this parameter is the IP address of the VPN gateway, which can be used to create IPsec-VPN or SSL-VPN connections.
	//
	// 	- If the VPN gateway supports IPsec-VPN connections in dual-tunnel mode, the value of this parameter is the first IP address that is used to create an IPsec-VPN connection. The IP address cannot be used to create SSL-VPN connections.
	//
	//     If the VPN gateway supports IPsec-VPN connections in dual-tunnel mode, the system assigns two IPsec addresses to the VPN gateway to create IPsec-VPN connections in dual-tunnel mode.
	//
	// example:
	//
	// 47.12.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Indicates whether IPsec-VPN is enabled for the VPN gateway. Valid values:
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
	// test
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
	// The information about pending orders.
	//
	// > This parameter is returned only if **IncludeReservationData*	- is set to **true**.
	ReservationData *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData `json:"ReservationData,omitempty" xml:"ReservationData,omitempty" type:"Struct"`
	// The ID of the resource group to which the VPN gateway belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The maximum bandwidth of the VPN gateway. Unit: **M**, which indicates Mbit/s.
	//
	// example:
	//
	// 5M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The number of SSL-VPN connections supported by the VPN gateway.
	//
	// example:
	//
	// 5
	SslMaxConnections *int64 `json:"SslMaxConnections,omitempty" xml:"SslMaxConnections,omitempty"`
	// Indicates whether SSL-VPN is enabled for the VPN gateway. Valid values:
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
	// This parameter is returned only if the VPN gateway is a public VPN gateway and supports IPsec-VPN connections in dual-tunnel mode. In addition, SSL-VPN must be enabled for the VPN gateway.
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
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is automatically generated for the VPN gateway.
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
	//     	- **true**: queries only SQL templates that need to be optimized.
	//
	//     	- **false**: does not query only SQL statements that need to be optimized.
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
	Tags *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vSwitch to which the VPN gateway belongs.
	//
	// example:
	//
	// vsw-bp15lbk8sgtr6r5b0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the VPN gateway belongs.
	//
	// example:
	//
	// vpc-bp1m3i0kn1nd4wiw9****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp17lofy9fd0dnvzv****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// The type of VPN gateway.
	//
	// Only **Normal*	- may be returned, which indicates a standard VPN gateway.
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
	// If the value of ReservationOrderType is **TEMP_UPGRADE**, this parameter indicates the time when the temporary upgrade expires.
	//
	// If the value of ReservationOrderType is **RENEWCHANGE*	- or **RENEW**, this parameter indicates the time when the renewal or renewal with a specification change takes effect.
	//
	// example:
	//
	// 2021-07-20T16:00:00Z
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
	// The type of the order that has not taken effect. Valid values:
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
	// The status of the pending order.
	//
	// 	- **1**: indicates that the order for renewal or the order for renewal with a specification change has not taken effect.
	//
	// 	- **2**: indicates that the order of the temporary upgrade has taken effect. After the temporary upgrade expires, the system restores the VPN gateway to its previous specifications. In this case, the values of **ReservationIpsec**, **ReservationMaxConnections**, **ReservationSpec**, and **ReservationSsl*	- indicate the previous specifications of the VPN gateway.
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
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
