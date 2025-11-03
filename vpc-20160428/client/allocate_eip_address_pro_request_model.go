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
	// Specifies whether to enable automatic payment. Default value: true. Valid values:
	//
	// 	- **false**: Automatic payment is disabled. After an order is generated, you must go to the Order Center to complete the payment.
	//
	// 	- **true**: Automatic payment is enabled. After an order is generated, the payment is automatically completed.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. This parameter is optional if **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the specified EIP. Unit: Mbit/s.
	//
	// 	- When **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByBandwidth**, valid values for **Bandwidth*	- are **1*	- to **500**.
	//
	// 	- When **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByTraffic**, valid values for **Bandwidth*	- are **1*	- to **200**.
	//
	// 	- When **InstanceChargeType*	- is set to **PrePaid**, valid values for **Bandwidth*	- are **1*	- to **1000**.
	//
	// Default value: **5*	- Mbit /s.
	//
	// example:
	//
	// 5
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP*	- (default): BGP (Multi-ISP) line The BGP (Multi-ISP) line is supported in all regions.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro line The BGP (Multi-ISP) Pro line is supported in the China (Hong Kong), Singapore, Malaysia (Kuala Lumpur), Philippines (Manila), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// For more information about the BGP (Multi-ISP) line and BGP (Multi-ISP) Pro line, see the "Line types" section of [What is EIP?](https://help.aliyun.com/document_detail/32321.html)
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
	// 	- If your services are deployed in China East 1 Finance, this parameter is required and you must set the parameter to **BGP_FinanceCloud**.
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
	// Set the value of **InternetChargeType*	- to **PayByBandwidth*	- if **InstanceChargeType*	- is set to **PrePaid**.
	//
	// Valid values when **InstanceChargeType*	- is set to **PostPaid**: **PayByBandwidth*	- or **PayByTraffic**.
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
	// 	- **PayByBandwidth*	- (default): pay-by-bandwidth.
	//
	// 	- **PayByTraffic**: pay-by-data-transfer.
	//
	// When **InstanceChargeType*	- is set to **PrePaid**, you must set **InternetChargeType*	- to **PayByBandwidth**.
	//
	// When **InstanceChargeType*	- is set to **PostPaid**, set **InternetChargeType*	- to **PayByBandwidth*	- or **PayByTraffic**.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address of the EIP.
	//
	// Specify **IpAddress*	- or **InstanceId**. If you leave both parameters empty, the system randomly allocates an EIP.
	//
	// example:
	//
	// 192.0.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. By default, this value is set to **public**, which specifies the public network type.
	//
	// example:
	//
	// public
	Netmode      *string `json:"Netmode,omitempty" xml:"Netmode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration.
	//
	// 	- Valid values when **PricingCycle*	- is set to **Month**: **1 to 9**.****
	//
	// 	- Valid values when **PricingCycle*	- is set to **Year**: **1 to 3**.****
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**.
	//
	// Leave this parameter empty if **InstanceChargeType*	- is set to **PostPaid**.
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
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. This parameter is optional if **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the IP address pool.
	//
	// The EIP is allocated from the IP address pool.
	//
	// By default, you cannot use the IP address pool. To use this feature, apply for the privilege in the Quota Center console. For more information, see the "Request a quota increase in the Quota Center console" section of [Manage EIP quotas](https://help.aliyun.com/document_detail/108213.html).
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
	// The ID of the resource group to which the EIP belongs.
	//
	// example:
	//
	// rg-resourcegroup****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The editions of Anti-DDoS.
	//
	// 	- If you do not specify this parameter, Anti-DDoS Origin Basic is used.
	//
	// 	- If you set the parameter to **AntiDDoS_Enhanced**, Anti-DDoS Pro/Premium is used.
	//
	// You can configure Anti-DDoS editions for up to 10 EIPs.
	SecurityProtectionTypes []*string                          `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Repeated"`
	Tag                     []*AllocateEipAddressProRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
