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
	// - **false**: Automatic payment is disabled. After an order is generated, go to the Order Center to complete the payment.
	//
	// - **true**: Automatic payment is enabled. The order is automatically paid.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. This parameter is optional if **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the EIP to allocate. Unit: Mbit/s.
	//
	// - If **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByBandwidth**, valid values of **Bandwidth*	- are **1*	- to **500**.
	//
	// - If **InstanceChargeType*	- is set to **PostPaid*	- and **InternetChargeType*	- is set to **PayByTraffic**, valid values of **Bandwidth*	- are **1*	- to **200**.
	//
	// - If **InstanceChargeType*	- is set to **PrePaid**, valid values of **Bandwidth*	- are **1*	- to **1000**.
	//
	// Default value: **5*	- Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The line type. Valid values:
	//
	// - **BGP*	- (default): BGP (multi-ISP) line. All regions support BGP (multi-ISP) EIPs.
	//
	// - **BGP_PRO**: BGP (multi-ISP) Pro line. Only the following regions support BGP (multi-ISP) Pro EIPs: Hong Kong (China), Singapore, Malaysia (Kuala Lumpur), Philippines (Manila), Indonesia (Jakarta), and Thailand (Bangkok).
	//
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
	// - If you are a China (Hangzhou) Finance Cloud user, this parameter is required. Set the value to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The billing method of the EIP to allocate. Valid values:
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
	// The instance ID of the EIP to allocate.
	//
	// You need to specify only one of **IpAddress*	- and **InstanceId**. If neither is specified, the system randomly allocates an EIP.
	//
	// example:
	//
	// eip-25877c70gddh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method of the EIP to allocate. Valid values:
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
	// The IP address of the EIP to allocate.
	//
	// You need to specify only one of **IpAddress*	- and **InstanceId**. If neither is specified, the system randomly allocates an EIP.
	//
	// example:
	//
	// 192.0.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. The value can only be **public*	- (default), which indicates the public network.
	//
	// example:
	//
	// public
	Netmode      *string `json:"Netmode,omitempty" xml:"Netmode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration.
	//
	// - If **PricingCycle*	- is set to **Month**, valid values of **Period*	- are **1*	- to **9**.
	//
	// - If **PricingCycle*	- is set to **Year**, valid values of **Period*	- are **1*	- to **3**.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**.
	//
	// Do not set this parameter if **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle of the subscription. Valid values:
	//
	// - **Month*	- (default): billed on a monthly basis.
	//
	// - **Year**: billed on a yearly basis.
	//
	// This parameter is required if **InstanceChargeType*	- is set to **PrePaid**. This parameter is optional if **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the IP address pool.
	//
	// The EIP is allocated from the specified IP address pool.
	//
	// The IP address pool feature is not available by default. To use this feature, apply for the IP address pool privilege quota in Quota Center. For more information, see [Increase a quota in Quota Center](https://help.aliyun.com/document_detail/108213.html).
	//
	// example:
	//
	// pippool-2vc0kxcedhquybdsz****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the EIP to allocate.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
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
	// The security protection level.
	//
	// - If this parameter is left empty, the default value is Anti-DDoS Basic.
	//
	// - If this parameter is set to **AntiDDoS_Enhanced**, Anti-DDoS (Enhanced) is enabled.
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
