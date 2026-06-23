// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnGatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPropagate(v bool) *ModifyVpnGatewayAttributeResponseBody
	GetAutoPropagate() *bool
	SetBusinessStatus(v string) *ModifyVpnGatewayAttributeResponseBody
	GetBusinessStatus() *string
	SetCreateTime(v int64) *ModifyVpnGatewayAttributeResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *ModifyVpnGatewayAttributeResponseBody
	GetDescription() *string
	SetDisasterRecoveryInternetIp(v string) *ModifyVpnGatewayAttributeResponseBody
	GetDisasterRecoveryInternetIp() *string
	SetDisasterRecoveryVSwitchId(v string) *ModifyVpnGatewayAttributeResponseBody
	GetDisasterRecoveryVSwitchId() *string
	SetEnableBgp(v bool) *ModifyVpnGatewayAttributeResponseBody
	GetEnableBgp() *bool
	SetEndTime(v int64) *ModifyVpnGatewayAttributeResponseBody
	GetEndTime() *int64
	SetInternetIp(v string) *ModifyVpnGatewayAttributeResponseBody
	GetInternetIp() *string
	SetIntranetIp(v string) *ModifyVpnGatewayAttributeResponseBody
	GetIntranetIp() *string
	SetName(v string) *ModifyVpnGatewayAttributeResponseBody
	GetName() *string
	SetRequestId(v string) *ModifyVpnGatewayAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *ModifyVpnGatewayAttributeResponseBody
	GetResourceGroupId() *string
	SetSpec(v string) *ModifyVpnGatewayAttributeResponseBody
	GetSpec() *string
	SetSslVpnInternetIp(v string) *ModifyVpnGatewayAttributeResponseBody
	GetSslVpnInternetIp() *string
	SetStatus(v string) *ModifyVpnGatewayAttributeResponseBody
	GetStatus() *string
	SetVSwitchId(v string) *ModifyVpnGatewayAttributeResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *ModifyVpnGatewayAttributeResponseBody
	GetVpcId() *string
	SetVpnGatewayId(v string) *ModifyVpnGatewayAttributeResponseBody
	GetVpnGatewayId() *string
}

type ModifyVpnGatewayAttributeResponseBody struct {
	// Indicates whether BGP routes are automatically propagated to the VPC. Valid values:
	//
	// - **true**: automatic propagation is enabled.
	//
	// - **false**: automatic propagation is not enabled.
	//
	// example:
	//
	// true
	AutoPropagate *bool `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	// The billing status of the VPN gateway instance. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **FinancialLocked**: locked due to overdue payment.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The timestamp when the VPN gateway instance was created. Unit: milliseconds.
	//
	// The timestamp follows the UNIX timestamp format, which represents the total number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC to the time when the VPN gateway instance was created.
	//
	// example:
	//
	// 1492753580000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the VPN gateway instance.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The second IP address assigned by the system to the VPN gateway instance for creating IPsec-VPN connections.
	//
	// This parameter is returned only for VPN gateway instances that support dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// 116.11.XX.XX
	DisasterRecoveryInternetIp *string `json:"DisasterRecoveryInternetIp,omitempty" xml:"DisasterRecoveryInternetIp,omitempty"`
	// The ID of the second vSwitch associated with the VPN gateway instance.
	//
	// This parameter is returned only for VPN gateway instances that support dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// vsw-p0w95ql6tmr2ludkt****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// The enabling status of the BGP feature for the VPN gateway. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: not enabled.
	//
	// example:
	//
	// true
	EnableBgp *bool `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	// The timestamp when the VPN gateway instance expires. Unit: milliseconds.
	//
	// The timestamp follows the UNIX timestamp format, which represents the total number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC to the time when the VPN gateway instance expires.
	//
	// example:
	//
	// 1495382400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// - If the VPN gateway instance supports single-tunnel IPsec-VPN connections, this address is the IP address of the VPN gateway instance and can be used to create IPsec-VPN connections or SSL-VPN connections.
	//
	// - If the VPN gateway instance supports dual-tunnel IPsec-VPN connections, this address is the first IP address used to create IPsec-VPN connections and cannot be used to create SSL-VPN connections.
	//
	//     If the VPN gateway instance supports dual-tunnel IPsec-VPN connections, the system assigns two IPsec IP addresses to the VPN gateway instance for creating dual-tunnel IPsec-VPN connections.
	//
	// example:
	//
	// 116.62.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the vSwitch occupied by the system when the VPN gateway instance was deployed.
	//
	// This parameter is returned only for VPN gateway instances that support single-tunnel IPsec-VPN connections and have the IPsec-VPN feature enabled.
	//
	// example:
	//
	// 172.16.10.46
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The name of the VPN gateway instance.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the VPN gateway instance belongs.
	//
	// You can call [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) to query resource group information.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The bandwidth specification of the VPN gateway instance. Unit: Mbit/s.
	//
	// example:
	//
	// 5M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The IP address of the SSL-VPN connection.
	//
	// This parameter is returned only when the SSL-VPN feature is enabled for a VPN gateway instance that supports dual-tunnel IPsec-VPN connections and uses the public network type.
	//
	// example:
	//
	// 116.33.XX.XX
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
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch associated with the VPN gateway instance.
	//
	// example:
	//
	// vsw-bp1y9ovl1cu9ou4tv****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the VPN gateway instance belongs.
	//
	// example:
	//
	// vpc-bp1ub1yt9cvakoel****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VPN gateway instance.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s ModifyVpnGatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnGatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetAutoPropagate() *bool {
	return s.AutoPropagate
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetDisasterRecoveryInternetIp() *string {
	return s.DisasterRecoveryInternetIp
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetDisasterRecoveryVSwitchId() *string {
	return s.DisasterRecoveryVSwitchId
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetEnableBgp() *bool {
	return s.EnableBgp
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetSslVpnInternetIp() *string {
	return s.SslVpnInternetIp
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyVpnGatewayAttributeResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetAutoPropagate(v bool) *ModifyVpnGatewayAttributeResponseBody {
	s.AutoPropagate = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetBusinessStatus(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetCreateTime(v int64) *ModifyVpnGatewayAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetDescription(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetDisasterRecoveryInternetIp(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.DisasterRecoveryInternetIp = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetDisasterRecoveryVSwitchId(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.DisasterRecoveryVSwitchId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetEnableBgp(v bool) *ModifyVpnGatewayAttributeResponseBody {
	s.EnableBgp = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetEndTime(v int64) *ModifyVpnGatewayAttributeResponseBody {
	s.EndTime = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetInternetIp(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.InternetIp = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetIntranetIp(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.IntranetIp = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetName(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetRequestId(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetResourceGroupId(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetSpec(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.Spec = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetSslVpnInternetIp(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.SslVpnInternetIp = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetStatus(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetVSwitchId(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetVpcId(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) SetVpnGatewayId(v string) *ModifyVpnGatewayAttributeResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *ModifyVpnGatewayAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
