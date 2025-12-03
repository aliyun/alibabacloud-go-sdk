// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateLoadBalancerRequest
	GetAddress() *string
	SetAddressIPVersion(v string) *CreateLoadBalancerRequest
	GetAddressIPVersion() *string
	SetAddressType(v string) *CreateLoadBalancerRequest
	GetAddressType() *string
	SetAutoPay(v bool) *CreateLoadBalancerRequest
	GetAutoPay() *bool
	SetBandwidth(v int32) *CreateLoadBalancerRequest
	GetBandwidth() *int32
	SetClientToken(v string) *CreateLoadBalancerRequest
	GetClientToken() *string
	SetDeleteProtection(v string) *CreateLoadBalancerRequest
	GetDeleteProtection() *string
	SetDuration(v int32) *CreateLoadBalancerRequest
	GetDuration() *int32
	SetInstanceChargeType(v string) *CreateLoadBalancerRequest
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *CreateLoadBalancerRequest
	GetInternetChargeType() *string
	SetLoadBalancerName(v string) *CreateLoadBalancerRequest
	GetLoadBalancerName() *string
	SetLoadBalancerSpec(v string) *CreateLoadBalancerRequest
	GetLoadBalancerSpec() *string
	SetMasterZoneId(v string) *CreateLoadBalancerRequest
	GetMasterZoneId() *string
	SetModificationProtectionReason(v string) *CreateLoadBalancerRequest
	GetModificationProtectionReason() *string
	SetModificationProtectionStatus(v string) *CreateLoadBalancerRequest
	GetModificationProtectionStatus() *string
	SetOwnerAccount(v string) *CreateLoadBalancerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLoadBalancerRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateLoadBalancerRequest
	GetPayType() *string
	SetPricingCycle(v string) *CreateLoadBalancerRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreateLoadBalancerRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateLoadBalancerRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateLoadBalancerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLoadBalancerRequest
	GetResourceOwnerId() *int64
	SetSlaveZoneId(v string) *CreateLoadBalancerRequest
	GetSlaveZoneId() *string
	SetTag(v []*CreateLoadBalancerRequestTag) *CreateLoadBalancerRequest
	GetTag() []*CreateLoadBalancerRequestTag
	SetVSwitchId(v string) *CreateLoadBalancerRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateLoadBalancerRequest
	GetVpcId() *string
}

type CreateLoadBalancerRequest struct {
	// The private IP address of the CLB instance. The private IP address must belong to the destination CIDR block of the vSwitch.
	//
	// example:
	//
	// 192.168.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The IP version of the CLB instance. Valid values: **ipv4*	- and **ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The network type of the CLB instance. Valid values:
	//
	// 	- **internet**: After an Internet-facing CLB instance is created, the system allocates a public IP address to the instance. The CLB instance can forward requests over the Internet.
	//
	// 	- **intranet**: After an internal-facing CLB instance is created, the system allocates a private IP address to the CLB instance. The CLB instance can forward requests only within the VPC.
	//
	// example:
	//
	// internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// Deprecated
	//
	// Specifies whether to automatically pay the subscription fee of the Internet-facing CLB instance. Valid values:
	//
	// 	- **true**: yes. The CLB instance is created after you call the API operation.
	//
	// 	- **false*	- (default): After you call the operation, the order is created but the payment is not completed. You can view pending orders in the console. The CLB instance will not be created until you complete the payment.
	//
	// >  This parameter takes effect only for subscription CLB instances created on the Alibaba Cloud China site.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s.
	//
	// Valid values: **1*	- to **5120**. For a pay-by-bandwidth Internet-facing CLB instance, you can specify a maximum bandwidth for each listener. The sum of the maximum bandwidth of all listeners cannot exceed the maximum bandwidth of the CLB instance.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// >  If you do not specify this parameter, the system uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable deletion protection for the CLB instance. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	DeleteProtection *string `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	// Deprecated
	//
	// The subscription duration of the Internet-facing CLB instance. Valid values:
	//
	// 	- If **PricingCycle*	- is set to **month**, the valid values are **1 to 9**.
	//
	// 	- If **PricingCycle*	- is set to **year**, the valid values are **1 to 5**.
	//
	// >  This parameter is supported only by subscription instances created on the Alibaba Cloud China site.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The metering method of the CLB instance. Valid values:
	//
	// 	- **PayBySpec*	- (default)
	//
	// 	- **PayByCLCU**
	//
	// >  This parameter is supported only by instances created on the Alibaba Cloud China site and only when **PayType*	- is set to **PayOnDemand**.
	//
	// example:
	//
	// PayBySpec
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Internet-facing CLB instance. Valid values:
	//
	// 	- **paybytraffic*	- (default): pay-by-data-transfer
	//
	//     >  If you set InternetChargeType to **paybytraffic**, you do not need to configure the **Bandwidth*	- parameter. The value of **Bandwidth*	- does not take effect even if you specify one.
	//
	// 	- **paybybandwidth**: pay-by-bandwidth
	//
	// >  If you set **PayType*	- to **PayOnDemand*	- and **InstanceChargeType*	- to **PayByCLCU**, the only valid value for InternetChargeType is **paybytraffic**.
	//
	// example:
	//
	// paybytraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The CLB instance name.
	//
	// The name must be 1 to 80 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// If you do not specify this parameter, the system automatically assigns a name to the CLB instance.
	//
	// example:
	//
	// lb-bp1o94dp5i6ea****
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The specification of the CLB instance. Valid values:
	//
	// 	- **slb.s1.small**
	//
	// 	- **slb.s2.small**
	//
	// 	- **slb.s2.medium**
	//
	// 	- **slb.s3.small**
	//
	// 	- **slb.s3.medium**
	//
	// 	- **slb.s3.large**
	//
	//
	//
	//  >   If you do not specify this parameter, a shared-resource CLB instance is created. Shared-resource CLB instances are no longer available for purchase. Therefore, you must specify this parameter.
	//
	// If **InstanceChargeType*	- is set to **PayByCLCU**, this parameter is invalid and you do not need to specify this parameter.
	//
	// example:
	//
	// slb.s1.small
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	// The ID of the primary zone to which the CLB instance belongs.
	//
	// You can call the [DescribeZone](https://help.aliyun.com/document_detail/2401683.html) operation to query the primary and secondary zones in the region where you want to create the CLB instance.
	//
	// example:
	//
	// cn-hangzhou-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The reason for enabling the configuration read-only mode. The reason must be 1 to 80 characters in length. It must start with a letter and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// >  This parameter takes effect only when **ModificationProtectionStatus*	- is set to **ConsoleProtection**.
	//
	// example:
	//
	// Managed instance
	ModificationProtectionReason *string `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	// Specifies whether to enable the configuration read-only mode. Valid values:
	//
	// 	- **NonProtection**: disables the configuration read-only mode. After you disable the configuration read-only mode, the value of **ModificationProtectionReason*	- is cleared.
	//
	// 	- **ConsoleProtection**: enables the configuration read-only mode.
	//
	// >  If you set this parameter to **ConsoleProtection**, you cannot modify instance configurations in the CLB console. However, you can modify instance configurations by calling API operations.
	//
	// example:
	//
	// ConsoleProtection
	ModificationProtectionStatus *string `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	OwnerAccount                 *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Deprecated
	//
	// The billing method of the CLB instance. Valid values:
	//
	// **PayOnDemand**: pay-as-you-go.
	//
	// >  The Alibaba Cloud International site supports only pay-as-you-go CLB instances.
	//
	// example:
	//
	// PayOnDemand
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Deprecated
	//
	// The billing cycle of the subscription Internet-facing CLB instance. Valid values:
	//
	// 	- **month**
	//
	// 	- **year**
	//
	// >  This parameter is supported only by subscription instances created on the Alibaba Cloud China site.
	//
	// example:
	//
	// month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the most recent region list.
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
	// rg-atstuj3rtopt****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the secondary zone to which the CLB instance belongs.
	//
	// You can call the [DescribeZone](https://help.aliyun.com/document_detail/2401683.html) operation to query the primary and secondary zones in the region where you want to create the CLB instance.
	//
	// example:
	//
	// cn-hangzhou-d
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	// The tags.
	Tag []*CreateLoadBalancerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the CLB instance belongs.
	//
	// If you want to deploy the CLB instance in a VPC, this parameter is required. If this parameter is specified, **AddessType*	- is set to **intranet*	- by default.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgy****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the CLB instance belongs.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateLoadBalancerRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *CreateLoadBalancerRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *CreateLoadBalancerRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateLoadBalancerRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateLoadBalancerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLoadBalancerRequest) GetDeleteProtection() *string {
	return s.DeleteProtection
}

func (s *CreateLoadBalancerRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateLoadBalancerRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateLoadBalancerRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *CreateLoadBalancerRequest) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *CreateLoadBalancerRequest) GetModificationProtectionReason() *string {
	return s.ModificationProtectionReason
}

func (s *CreateLoadBalancerRequest) GetModificationProtectionStatus() *string {
	return s.ModificationProtectionStatus
}

func (s *CreateLoadBalancerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLoadBalancerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLoadBalancerRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateLoadBalancerRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateLoadBalancerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLoadBalancerRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLoadBalancerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLoadBalancerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLoadBalancerRequest) GetSlaveZoneId() *string {
	return s.SlaveZoneId
}

func (s *CreateLoadBalancerRequest) GetTag() []*CreateLoadBalancerRequestTag {
	return s.Tag
}

func (s *CreateLoadBalancerRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLoadBalancerRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateLoadBalancerRequest) SetAddress(v string) *CreateLoadBalancerRequest {
	s.Address = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressIPVersion(v string) *CreateLoadBalancerRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressType(v string) *CreateLoadBalancerRequest {
	s.AddressType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAutoPay(v bool) *CreateLoadBalancerRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetBandwidth(v int32) *CreateLoadBalancerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDeleteProtection(v string) *CreateLoadBalancerRequest {
	s.DeleteProtection = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDuration(v int32) *CreateLoadBalancerRequest {
	s.Duration = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetInstanceChargeType(v string) *CreateLoadBalancerRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetInternetChargeType(v string) *CreateLoadBalancerRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerSpec(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetMasterZoneId(v string) *CreateLoadBalancerRequest {
	s.MasterZoneId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionReason(v string) *CreateLoadBalancerRequest {
	s.ModificationProtectionReason = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionStatus(v string) *CreateLoadBalancerRequest {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetOwnerAccount(v string) *CreateLoadBalancerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetOwnerId(v int64) *CreateLoadBalancerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetPayType(v string) *CreateLoadBalancerRequest {
	s.PayType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetPricingCycle(v string) *CreateLoadBalancerRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetRegionId(v string) *CreateLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceGroupId(v string) *CreateLoadBalancerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetSlaveZoneId(v string) *CreateLoadBalancerRequest {
	s.SlaveZoneId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetTag(v []*CreateLoadBalancerRequestTag) *CreateLoadBalancerRequest {
	s.Tag = v
	return s
}

func (s *CreateLoadBalancerRequest) SetVSwitchId(v string) *CreateLoadBalancerRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetVpcId(v string) *CreateLoadBalancerRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLoadBalancerRequest) Validate() error {
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

type CreateLoadBalancerRequestTag struct {
	// The tag key of the bastion host. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// The tag key can be at most 64 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: **1 to 20**. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadBalancerRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLoadBalancerRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLoadBalancerRequestTag) SetKey(v string) *CreateLoadBalancerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLoadBalancerRequestTag) SetValue(v string) *CreateLoadBalancerRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLoadBalancerRequestTag) Validate() error {
	return dara.Validate(s)
}
