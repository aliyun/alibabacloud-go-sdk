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
	// The ID of the special activity. You do not need to specify this parameter.
	//
	// example:
	//
	// 123456
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **false*	- (default): disables automatic payment. After an order is generated, you must go to the User Center to complete the payment.
	//
	// - **true**: enables automatic payment. The order is automatically paid.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. If **InstanceChargeType*	- is set to **PostPaid**, you do not need to specify this parameter.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// - If **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByBandwidth**, valid values for **Bandwidth*	- are **1*	- to **500**.
	//
	// - If **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByTraffic**, valid values for **Bandwidth*	- are **1*	- to **200**.
	//
	// - If **InstanceChargeType*	- is set to **PrePaid**, valid values for **Bandwidth*	- are **1*	- to **1000**.
	//
	// Default value: **5*	- Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. **ClientToken*	- can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the EIP.
	//
	// The description must be 0 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// > You cannot specify this parameter when you create a subscription EIP.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The line type. Valid values:
	//
	// - **BGP*	- (default): BGP (Multi-ISP) line. BGP (Multi-ISP) EIPs are supported in all regions.
	//
	// - **BGP_PRO**: BGP (Multi-ISP) Pro line. BGP (Multi-ISP) Pro EIPs are supported only in the Hong Kong (China), Singapore, Japan (Tokyo), Malaysia (Kuala Lumpur), Philippines (Manila), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// For more information about BGP (Multi-ISP) and BGP (Multi-ISP) Pro lines, see [EIP line types](https://help.aliyun.com/document_detail/32321.html).
	//
	// - If you are a whitelisted user of single-ISP bandwidth, you can also select the following types:
	//
	//     - **ChinaTelecom**: China Telecom
	//
	//     - **ChinaUnicom**: China Unicom
	//
	//     - **ChinaMobile**: China Mobile
	//
	//     - **ChinaTelecom_L2**: China Telecom L2
	//
	//     - **ChinaUnicom_L2**: China Unicom L2
	//
	//     - **ChinaMobile_L2**: China Mobile L2
	//
	// - If you are an Alibaba Finance Cloud user in the China (Hangzhou) region, this parameter is required. Set the value to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The billing method of the EIP. Valid values:
	//
	//
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid*	- (default): pay-as-you-go billing method.
	//
	// If **InstanceChargeType*	- is set to **PrePaid**, **InternetChargeType*	- must be set to **PayByBandwidth**. If **InstanceChargeType*	- is set to **PostPaid**, **InternetChargeType*	- can be set to **PayByBandwidth*	- or **PayByTraffic**.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID of the EIP to be allocated.
	//
	// You need to specify only one of the **IpAddress*	- and **InstanceId*	- parameters. If neither is specified, the system randomly allocates an EIP.
	//
	// example:
	//
	// eip-25877c70gddh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// - **PayByBandwidth*	- (default): pay-by-bandwidth.
	//
	// - **PayByTraffic**: pay-by-data-transfer.
	//
	// If **InstanceChargeType*	- is set to **PrePaid**, **InternetChargeType*	- must be set to **PayByBandwidth**.
	//
	// If **InstanceChargeType*	- is set to **PostPaid**, **InternetChargeType*	- can be set to **PayByBandwidth*	- or **PayByTraffic**.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address of the EIP to be allocated.
	//
	// You need to specify only one of the **IpAddress*	- and **InstanceId*	- parameters. If neither is specified, the system randomly allocates an EIP.
	//
	// example:
	//
	// 192.0.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the EIP.
	//
	// The name must be 0 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// > You cannot specify this parameter when you create a subscription EIP.
	//
	// example:
	//
	// EIP1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type. Set the value to **public*	- (default), which indicates the Internet.
	//
	// example:
	//
	// public
	Netmode      *string `json:"Netmode,omitempty" xml:"Netmode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration.
	//
	// If **PricingCycle*	- is set to **Month**, valid values for **Period*	- are **1*	- to **9**.
	//
	// If **PricingCycle*	- is set to **Year**, valid values for **Period*	- are **1*	- to **5**.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. If **InstanceChargeType*	- is set to **PostPaid**, you do not need to specify this parameter.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle of the subscription. Valid values:
	//
	// - **Month*	- (default): monthly subscription.
	//
	// - **Year**: yearly subscription.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. If **InstanceChargeType*	- is set to **PostPaid**, you do not need to specify this parameter.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the IP address pool.
	//
	// The EIP is allocated from the IP address pool.
	//
	// By default, the IP address pool feature is disabled. To use this feature, apply for the IP address pool quota privilege quota Center. For more information, see [Increase quotas privilege quota Center](https://help.aliyun.com/document_detail/108213.html).
	//
	// example:
	//
	// pippool-2vc0kxcedhquybdsz****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the EIP.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
	// The security protection level.
	//
	// - If this parameter is left empty, Anti-DDoS Basic is used by default.
	//
	// - If this parameter is set to **AntiDDoS_Enhanced**, Anti-DDoS Pro/Premium is used.
	//
	// You can specify only one security protection level.
	//
	// example:
	//
	// AntiDDoS_Enhanced
	SecurityProtectionTypes []*string `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Repeated"`
	// The tags of the resource.
	Tag []*AllocateEipAddressRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone of the EIP.
	//
	// If the IP address pool specified by **PublicIpAddressPoolId*	- is of the CloudBox type, this parameter defaults to the zone of the IP address pool.
	//
	// For more information about how to view the service type of an IP address pool, see [ListPublicIpAddressPools](https://api.aliyun.com/document/Vpc/2016-04-28/ListPublicIpAddressPools).
	//
	// example:
	//
	// ap-southeast-1-lzdvn-cb
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
	// The tag key of the resource. You can specify up to 20 tag keys. If you specify this parameter, you cannot specify an empty string.
	//
	// A tag key can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. When you call the API operation, use the Tag.N.Value format. Valid values of N: 1 to 20. If you specify this parameter, you cannot specify an empty string. A tag value can be up to 128 characters in length, and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceJoshua
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
