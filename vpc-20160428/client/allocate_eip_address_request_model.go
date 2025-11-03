// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *AllocateEipAddressRequest
	GetActivityId() *int64
	SetAutoPay(v bool) *AllocateEipAddressRequest
	GetAutoPay() *bool
	SetBandwidth(v string) *AllocateEipAddressRequest
	GetBandwidth() *string
	SetClientToken(v string) *AllocateEipAddressRequest
	GetClientToken() *string
	SetDescription(v string) *AllocateEipAddressRequest
	GetDescription() *string
	SetISP(v string) *AllocateEipAddressRequest
	GetISP() *string
	SetInstanceChargeType(v string) *AllocateEipAddressRequest
	GetInstanceChargeType() *string
	SetInstanceId(v string) *AllocateEipAddressRequest
	GetInstanceId() *string
	SetInternetChargeType(v string) *AllocateEipAddressRequest
	GetInternetChargeType() *string
	SetIpAddress(v string) *AllocateEipAddressRequest
	GetIpAddress() *string
	SetName(v string) *AllocateEipAddressRequest
	GetName() *string
	SetNetmode(v string) *AllocateEipAddressRequest
	GetNetmode() *string
	SetOwnerAccount(v string) *AllocateEipAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateEipAddressRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *AllocateEipAddressRequest
	GetPeriod() *int32
	SetPricingCycle(v string) *AllocateEipAddressRequest
	GetPricingCycle() *string
	SetPublicIpAddressPoolId(v string) *AllocateEipAddressRequest
	GetPublicIpAddressPoolId() *string
	SetRegionId(v string) *AllocateEipAddressRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AllocateEipAddressRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AllocateEipAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateEipAddressRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionTypes(v []*string) *AllocateEipAddressRequest
	GetSecurityProtectionTypes() []*string
	SetTag(v []*AllocateEipAddressRequestTag) *AllocateEipAddressRequest
	GetTag() []*AllocateEipAddressRequestTag
	SetZone(v string) *AllocateEipAddressRequest
	GetZone() *string
}

type AllocateEipAddressRequest struct {
	// The promotion code. This parameter is not required.
	//
	// example:
	//
	// 123456
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **false*	- (default): The automatic payment is disabled. If you select this option, you must go to the Order Center to complete the payment after an order is generated.
	//
	// 	- **true**: The automatic payment is enabled. Payments are automatically complete after an order is generated.
	//
	// If **InstanceChargeType*	- is set to **PrePaid**, this parameter is required. If **InstanceChargeType*	- is set to **PostPaid**, this parameter is not required.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// 	- Valid values when **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByBandwidth**: **1*	- to **500**.****
	//
	// 	- Valid values when **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByTraffic**: **1*	- to **200**.****
	//
	// 	- Valid values when **InstanceChargeType*	- is set to **PrePaid**: **1*	- to **1000**.****
	//
	// Default value: **5**. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The **client token*	- can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the **client token**. The value of **RequestId*	- is different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the EIP.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// >  You cannot specify this parameter if you create a subscription EIP.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP*	- (default): BGP (Multi-ISP) All regions support BGP (Multi-ISP) EIPs.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro Only the following regions support BGP (Multi-ISP) Pro lines: China (Hong Kong), Singapore, Japan (Tokyo), Malaysia (Kuala Lumpur), Philippines (Manila), Indonesia (Jakarta), and Thailand (Bangkok).
	//
	// For more information about BGP (Multi-ISP) and BGP (Multi-ISP) Pro, see the "Line types" section of [What is EIP?](https://help.aliyun.com/document_detail/32321.html)
	//
	// 	- If you are allowed to use single-ISP bandwidth, you can also choose one of the following values:
	//
	//     	- **ChinaTelecom**
	//
	//     	- **ChinaUnicom**
	//
	//     	- **ChinaMobile**
	//
	//     	- **ChinaTelecom_L2**
	//
	//     	- **ChinaUnicom_L2**
	//
	//     	- **ChinaMobile_L2**
	//
	// 	- If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The billing method of the EIP. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid*	- (default): pay-as-you-go
	//
	// If **InstanceChargeType*	- is set to **PrePaid**, set **InternetChargeType*	- to **PayByBandwidth**. If **InstanceChargeType*	- is set to **PostPaid**, set **InternetChargeType*	- to **PayByBandwidth*	- or **PayByTraffic**.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The EIP ID.
	//
	// Specify **IpAddress*	- or **InstanceId**. If you leave both parameters empty, the system randomly allocates an EIP.
	//
	// example:
	//
	// eip-25877c70gddh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- **PayByBandwidth*	- (default): pay-by-bandwidth
	//
	// 	- **PayByTraffic**: pay-by-data-transfer
	//
	// When **InstanceChargeType*	- is set to **PrePaid**, set **InternetChargeType*	- to **PayByBandwidth**.
	//
	// When **InstanceChargeType*	- is set to **PostPaid**, set **InternetChargeType*	- to **PayByBandwidth*	- or **PayByTraffic**.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address of the EIP that you want to request.
	//
	// Specify **IpAddress*	- or **InstanceId**. If you leave both parameters empty, the system randomly allocates an EIP.
	//
	// example:
	//
	// 192.0.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The EIP name.
	//
	// The name must be 1 to 128 characters in length and start with a letter, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// >  You cannot specify this parameter if you create a subscription EIP.
	//
	// example:
	//
	// EIP1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type. Default value: **public**.
	//
	// example:
	//
	// public
	Netmode      *string `json:"Netmode,omitempty" xml:"Netmode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the EIP.
	//
	// Valid values when **PricingCycle*	- is set to **Month**: **1*	- to **9**.****
	//
	// Valid values when **PricingCycle*	- is set to **Year**: **1*	- to **5**.****
	//
	// This parameter must be specified when **InstanceChargeType*	- is set to **PrePaid**. This parameter is optional when **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle of the subscription EIP. Valid values:
	//
	// 	- **Month*	- (default)
	//
	// 	- **Year**
	//
	// If **InstanceChargeType*	- is set to **PrePaid**, this parameter is required. If **InstanceChargeType*	- is set to **PostPaid**, this parameter is not required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the IP address pool.
	//
	// The EIP is allocated from the IP address pool.
	//
	// By default, the IP address pool feature is unavailable. To use the IP address pool, apply for the privilege in the Quota Center console. For more information, see the "Request a quota increase in the Quota Center console" section in [Manage EIP quotas](https://help.aliyun.com/document_detail/108213.html).
	//
	// example:
	//
	// pippool-2vc0kxcedhquybdsz****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The ID of the region to which the EIP belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazffggds****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The editions of Anti-DDoS.
	//
	// 	- If you do not specify this parameter, Anti-DDoS Origin Basic is used.
	//
	// 	- If you set the parameter to **AntiDDoS_Enhanced**, Anti-DDoS Pro/Premium is used.
	//
	// You can specify up to 10 editions of Anti-DDoS.
	//
	// example:
	//
	// AntiDDoS_Enhanced
	SecurityProtectionTypes []*string                       `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Repeated"`
	Tag                     []*AllocateEipAddressRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone of the EIP.
	//
	// When the service type of the IP address pool specified by **PublicIpAddressPoolId*	- is CloudBox, the default value is the zone of the IP address pool.
	//
	// For more information, see [ListPublicIpAddressPools](https://help.aliyun.com/document_detail/429433.html).
	//
	// example:
	//
	// cn-hangzhou-a
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s AllocateEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *AllocateEipAddressRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *AllocateEipAddressRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *AllocateEipAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateEipAddressRequest) GetDescription() *string {
	return s.Description
}

func (s *AllocateEipAddressRequest) GetISP() *string {
	return s.ISP
}

func (s *AllocateEipAddressRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *AllocateEipAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AllocateEipAddressRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *AllocateEipAddressRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *AllocateEipAddressRequest) GetName() *string {
	return s.Name
}

func (s *AllocateEipAddressRequest) GetNetmode() *string {
	return s.Netmode
}

func (s *AllocateEipAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateEipAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateEipAddressRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *AllocateEipAddressRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *AllocateEipAddressRequest) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *AllocateEipAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateEipAddressRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateEipAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateEipAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateEipAddressRequest) GetSecurityProtectionTypes() []*string {
	return s.SecurityProtectionTypes
}

func (s *AllocateEipAddressRequest) GetTag() []*AllocateEipAddressRequestTag {
	return s.Tag
}

func (s *AllocateEipAddressRequest) GetZone() *string {
	return s.Zone
}

func (s *AllocateEipAddressRequest) SetActivityId(v int64) *AllocateEipAddressRequest {
	s.ActivityId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetAutoPay(v bool) *AllocateEipAddressRequest {
	s.AutoPay = &v
	return s
}

func (s *AllocateEipAddressRequest) SetBandwidth(v string) *AllocateEipAddressRequest {
	s.Bandwidth = &v
	return s
}

func (s *AllocateEipAddressRequest) SetClientToken(v string) *AllocateEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateEipAddressRequest) SetDescription(v string) *AllocateEipAddressRequest {
	s.Description = &v
	return s
}

func (s *AllocateEipAddressRequest) SetISP(v string) *AllocateEipAddressRequest {
	s.ISP = &v
	return s
}

func (s *AllocateEipAddressRequest) SetInstanceChargeType(v string) *AllocateEipAddressRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *AllocateEipAddressRequest) SetInstanceId(v string) *AllocateEipAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetInternetChargeType(v string) *AllocateEipAddressRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AllocateEipAddressRequest) SetIpAddress(v string) *AllocateEipAddressRequest {
	s.IpAddress = &v
	return s
}

func (s *AllocateEipAddressRequest) SetName(v string) *AllocateEipAddressRequest {
	s.Name = &v
	return s
}

func (s *AllocateEipAddressRequest) SetNetmode(v string) *AllocateEipAddressRequest {
	s.Netmode = &v
	return s
}

func (s *AllocateEipAddressRequest) SetOwnerAccount(v string) *AllocateEipAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateEipAddressRequest) SetOwnerId(v int64) *AllocateEipAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetPeriod(v int32) *AllocateEipAddressRequest {
	s.Period = &v
	return s
}

func (s *AllocateEipAddressRequest) SetPricingCycle(v string) *AllocateEipAddressRequest {
	s.PricingCycle = &v
	return s
}

func (s *AllocateEipAddressRequest) SetPublicIpAddressPoolId(v string) *AllocateEipAddressRequest {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetRegionId(v string) *AllocateEipAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetResourceGroupId(v string) *AllocateEipAddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetResourceOwnerAccount(v string) *AllocateEipAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateEipAddressRequest) SetResourceOwnerId(v int64) *AllocateEipAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetSecurityProtectionTypes(v []*string) *AllocateEipAddressRequest {
	s.SecurityProtectionTypes = v
	return s
}

func (s *AllocateEipAddressRequest) SetTag(v []*AllocateEipAddressRequestTag) *AllocateEipAddressRequest {
	s.Tag = v
	return s
}

func (s *AllocateEipAddressRequest) SetZone(v string) *AllocateEipAddressRequest {
	s.Zone = &v
	return s
}

func (s *AllocateEipAddressRequest) Validate() error {
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

type AllocateEipAddressRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AllocateEipAddressRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressRequestTag) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressRequestTag) GetKey() *string {
	return s.Key
}

func (s *AllocateEipAddressRequestTag) GetValue() *string {
	return s.Value
}

func (s *AllocateEipAddressRequestTag) SetKey(v string) *AllocateEipAddressRequestTag {
	s.Key = &v
	return s
}

func (s *AllocateEipAddressRequestTag) SetValue(v string) *AllocateEipAddressRequestTag {
	s.Value = &v
	return s
}

func (s *AllocateEipAddressRequestTag) Validate() error {
	return dara.Validate(s)
}
