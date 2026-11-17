// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipAddressProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *AllocateEipAddressProRequest
	GetAutoPay() *bool
	SetBandwidth(v string) *AllocateEipAddressProRequest
	GetBandwidth() *string
	SetClientToken(v string) *AllocateEipAddressProRequest
	GetClientToken() *string
	SetISP(v string) *AllocateEipAddressProRequest
	GetISP() *string
	SetInstanceChargeType(v string) *AllocateEipAddressProRequest
	GetInstanceChargeType() *string
	SetInstanceId(v string) *AllocateEipAddressProRequest
	GetInstanceId() *string
	SetInternetChargeType(v string) *AllocateEipAddressProRequest
	GetInternetChargeType() *string
	SetIpAddress(v string) *AllocateEipAddressProRequest
	GetIpAddress() *string
	SetNetmode(v string) *AllocateEipAddressProRequest
	GetNetmode() *string
	SetOwnerAccount(v string) *AllocateEipAddressProRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateEipAddressProRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *AllocateEipAddressProRequest
	GetPeriod() *int32
	SetPricingCycle(v string) *AllocateEipAddressProRequest
	GetPricingCycle() *string
	SetPublicIpAddressPoolId(v string) *AllocateEipAddressProRequest
	GetPublicIpAddressPoolId() *string
	SetRegionId(v string) *AllocateEipAddressProRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AllocateEipAddressProRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AllocateEipAddressProRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateEipAddressProRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionTypes(v []*string) *AllocateEipAddressProRequest
	GetSecurityProtectionTypes() []*string
	SetTag(v []*AllocateEipAddressProRequestTag) *AllocateEipAddressProRequest
	GetTag() []*AllocateEipAddressProRequestTag
}

type AllocateEipAddressProRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **false**: Automatic payment is disabled. After an order is generated, you must go to the User Center to complete the payment.
	//
	// - **true**: Automatic payment is enabled, and the order is automatically paid.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. If **InstanceChargeType*	- is set to **PostPaid**, you can leave this parameter empty.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the EIP to be allocated. Unit: Mbit/s.
	//
	// - If **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByBandwidth**, valid values for **Bandwidth*	- are **1*	- to **500**.
	//
	// - If **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByTraffic**, valid values for **Bandwidth*	- are **1*	- to **200**.
	//
	// - If **InstanceChargeType*	- is set to **PrePaid**, valid values for **Bandwidth*	- are **1*	- to **1000**.
	//
	// Default value: **5**.
	//
	// example:
	//
	// 5
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The line type. Valid values:
	//
	// - **BGP*	- (default): BGP (multi-ISP) line. BGP (multi-ISP) EIPs are supported in all regions.
	//
	// - **BGP_PRO**: BGP (multi-ISP) Pro line. BGP (multi-ISP) Pro EIPs are supported only in the Hong Kong (China), Singapore, Malaysia (Kuala Lumpur), Philippines (Manila), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// For more information about BGP (multi-ISP) and BGP (multi-ISP) Pro lines, see [EIP line types](https://help.aliyun.com/document_detail/32321.html).
	//
	// - If you are a single-ISP bandwidth whitelist user, you can also select the following types:
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
	// The billing method of the EIP to be allocated. Valid values:
	//
	//
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid*	- (default): pay-as-you-go.
	//
	// If **InstanceChargeType*	- is set to **PrePaid**, **InternetChargeType*	- must be set to **PayByBandwidth**.
	//
	// If **InstanceChargeType*	- is set to **PostPaid**, **InternetChargeType*	- can be set to **PayByBandwidth*	- or **PayByTraffic**.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID of the EIP to be allocated.
	//
	// You must specify either **IpAddress*	- or **InstanceId**. If neither is specified, the system randomly allocates an EIP.
	//
	// example:
	//
	// eip-25877c70gddh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method of the EIP to be allocated. Valid values:
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
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address of the EIP to be allocated.
	//
	// You must specify either **IpAddress*	- or **InstanceId**. If neither is specified, the system randomly allocates an EIP.
	//
	// example:
	//
	// 192.0.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
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
	// - If **PricingCycle*	- is set to **Month**, valid values for **Period*	- are **1*	- to **9**.
	//
	// - If **PricingCycle*	- is set to **Year**, valid values for **Period*	- are **1*	- to **3**.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**.
	//
	// Leave this parameter empty if **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle of the subscription. Valid values:
	//
	// - **Month*	- (default): billed by month.
	//
	// - **Year**: billed by year.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. If **InstanceChargeType*	- is set to **PostPaid**, you can leave this parameter empty.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the IP address pool.
	//
	// The EIP is allocated from this IP address pool.
	//
	// By default, the IP address pool feature is not available. To use this feature, contact your account manager.
	//
	// example:
	//
	// pippool-2vc0kxcedhquybdsz****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the EIP to be allocated.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the EIP to be allocated belongs.
	//
	// example:
	//
	// rg-resourcegroup****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security protection level.
	//
	// - If this parameter is left empty, Anti-DDoS Basic is used by default.
	//
	// - If this parameter is set to **AntiDDoS_Enhanced**, Anti-DDoS (Enhanced) is used.
	SecurityProtectionTypes []*string `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Repeated"`
	// The tag information.
	//
	// You can specify up to 20 tags at a time.
	Tag []*AllocateEipAddressProRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AllocateEipAddressProRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressProRequest) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressProRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *AllocateEipAddressProRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *AllocateEipAddressProRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateEipAddressProRequest) GetISP() *string {
	return s.ISP
}

func (s *AllocateEipAddressProRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *AllocateEipAddressProRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AllocateEipAddressProRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *AllocateEipAddressProRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *AllocateEipAddressProRequest) GetNetmode() *string {
	return s.Netmode
}

func (s *AllocateEipAddressProRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateEipAddressProRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateEipAddressProRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *AllocateEipAddressProRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *AllocateEipAddressProRequest) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *AllocateEipAddressProRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateEipAddressProRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateEipAddressProRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateEipAddressProRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateEipAddressProRequest) GetSecurityProtectionTypes() []*string {
	return s.SecurityProtectionTypes
}

func (s *AllocateEipAddressProRequest) GetTag() []*AllocateEipAddressProRequestTag {
	return s.Tag
}

func (s *AllocateEipAddressProRequest) SetAutoPay(v bool) *AllocateEipAddressProRequest {
	s.AutoPay = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetBandwidth(v string) *AllocateEipAddressProRequest {
	s.Bandwidth = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetClientToken(v string) *AllocateEipAddressProRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetISP(v string) *AllocateEipAddressProRequest {
	s.ISP = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetInstanceChargeType(v string) *AllocateEipAddressProRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetInstanceId(v string) *AllocateEipAddressProRequest {
	s.InstanceId = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetInternetChargeType(v string) *AllocateEipAddressProRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetIpAddress(v string) *AllocateEipAddressProRequest {
	s.IpAddress = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetNetmode(v string) *AllocateEipAddressProRequest {
	s.Netmode = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetOwnerAccount(v string) *AllocateEipAddressProRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetOwnerId(v int64) *AllocateEipAddressProRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetPeriod(v int32) *AllocateEipAddressProRequest {
	s.Period = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetPricingCycle(v string) *AllocateEipAddressProRequest {
	s.PricingCycle = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetPublicIpAddressPoolId(v string) *AllocateEipAddressProRequest {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetRegionId(v string) *AllocateEipAddressProRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetResourceGroupId(v string) *AllocateEipAddressProRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetResourceOwnerAccount(v string) *AllocateEipAddressProRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetResourceOwnerId(v int64) *AllocateEipAddressProRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateEipAddressProRequest) SetSecurityProtectionTypes(v []*string) *AllocateEipAddressProRequest {
	s.SecurityProtectionTypes = v
	return s
}

func (s *AllocateEipAddressProRequest) SetTag(v []*AllocateEipAddressProRequestTag) *AllocateEipAddressProRequest {
	s.Tag = v
	return s
}

func (s *AllocateEipAddressProRequest) Validate() error {
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

type AllocateEipAddressProRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. If you specify this parameter, you cannot enter an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. If you specify this parameter, you can enter an empty string.
	//
	// A tag value can be up to 128 characters in length. It must start with a letter or a Chinese character, and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AllocateEipAddressProRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressProRequestTag) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressProRequestTag) GetKey() *string {
	return s.Key
}

func (s *AllocateEipAddressProRequestTag) GetValue() *string {
	return s.Value
}

func (s *AllocateEipAddressProRequestTag) SetKey(v string) *AllocateEipAddressProRequestTag {
	s.Key = &v
	return s
}

func (s *AllocateEipAddressProRequestTag) SetValue(v string) *AllocateEipAddressProRequestTag {
	s.Value = &v
	return s
}

func (s *AllocateEipAddressProRequestTag) Validate() error {
	return dara.Validate(s)
}
