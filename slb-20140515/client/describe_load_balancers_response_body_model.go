// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancers(v *DescribeLoadBalancersResponseBodyLoadBalancers) *DescribeLoadBalancersResponseBody
	GetLoadBalancers() *DescribeLoadBalancersResponseBodyLoadBalancers
	SetPageNumber(v int32) *DescribeLoadBalancersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadBalancersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLoadBalancersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLoadBalancersResponseBody
	GetTotalCount() *int32
}

type DescribeLoadBalancersResponseBody struct {
	// An array of CLB instances.
	LoadBalancers *DescribeLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8B9DB03B-ED39-5DB8-9C9F-1ED5F548D61E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of instances returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadBalancersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBody) GetLoadBalancers() *DescribeLoadBalancersResponseBodyLoadBalancers {
	return s.LoadBalancers
}

func (s *DescribeLoadBalancersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadBalancersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadBalancersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLoadBalancersResponseBody) SetLoadBalancers(v *DescribeLoadBalancersResponseBodyLoadBalancers) *DescribeLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetPageNumber(v int32) *DescribeLoadBalancersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetPageSize(v int32) *DescribeLoadBalancersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetRequestId(v string) *DescribeLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetTotalCount(v int32) *DescribeLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) Validate() error {
	if s.LoadBalancers != nil {
		if err := s.LoadBalancers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancersResponseBodyLoadBalancers struct {
	LoadBalancer []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancers) GetLoadBalancer() []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	return s.LoadBalancer
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancers) SetLoadBalancer(v []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) *DescribeLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancer = v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancers) Validate() error {
	if s.LoadBalancer != nil {
		for _, item := range s.LoadBalancer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer struct {
	// The endpoint of the CLB instance.
	//
	// example:
	//
	// 100.XX.XX.28
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The IP version that is used by the CLB instance. Valid values: **ipv4*	- and **ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The network type of the CLB instance. Valid values:
	//
	// 	- **internet:*	- After an Internet-facing CLB instance is created, the system assigns a public IP address to the CLB instance. Then, the CLB instance can forward requests over the Internet.
	//
	// 	- **intranet:*	- After an internal-facing CLB instance is created, the system assigns a private IP address to the CLB instance. Then, the CLB instance can forward requests only over internal networks.
	//
	// example:
	//
	// intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The time when the CLB instance was created. The time follows the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2021-12-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp when the instance was created.
	//
	// example:
	//
	// 1504147745000
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// Indicates whether deletion protection is enabled for the CLB instance. Valid values:
	//
	// 	- **on:*	- Deletion protection is enabled.
	//
	// 	- **off:*	- Deletion protection is disabled.
	//
	// example:
	//
	// off
	DeleteProtection *string `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	// The metering method of the CLB instance. Valid values:
	//
	// 	- **PayBySpec:*	- pay-by-specification.
	//
	// 	- **PayByCLCU:*	- pay-by-LCU.
	//
	// >  This parameter takes effect only for accounts registered on the China site (aliyun.com) and when the **PayType*	- parameter is set to **PayOnDemand**.
	//
	// example:
	//
	// PayBySpec
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Internet-facing CLB instance. Valid values:
	//
	// 	- **3:*	- pay-by-bandwidth (**paybybandwidth**).
	//
	// 	- **4:*	- pay-by-data-transfer (**paybytraffic**).
	//
	// example:
	//
	// 3
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The metering method of Internet data transfer. Valid values:
	//
	// 	- **paybybandwidth:*	- pay-by-bandwidth.
	//
	// 	- **paybytraffic:*	- pay-by-data-transfer.
	//
	// example:
	//
	// paybybandwidth
	InternetChargeTypeAlias *string `json:"InternetChargeTypeAlias,omitempty" xml:"InternetChargeTypeAlias,omitempty"`
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-bp1b6c719dfa****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the CLB instance.
	//
	// example:
	//
	// lb-bp1o94dp5i6ea****
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The specification of the CLB instance.
	//
	// >  Pay-as-you-go CLB instances are not subject to specifications. **slb.lcu.elastic*	- is returned by default.
	//
	// example:
	//
	// slb.s1.small
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	// The status of the CLB instance. Valid values:
	//
	// 	- **inactive:*	- The CLB instance is disabled. CLB instances in the inactive state do not forward traffic.
	//
	// 	- **active:*	- The CLB instance runs as expected. By default, newly created CLB instances are in the **active*	- state.
	//
	// 	- **locked:*	- The CLB instance is locked.
	//
	// example:
	//
	// active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The ID of the primary zone to which the CLB instance belongs.
	//
	// example:
	//
	// cn-hangzhou-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The reason why the configuration read-only mode was enabled. The reason must be 1 to 80 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The reason must start with a letter.
	//
	// > This parameter takes effect only when you set the `ModificationProtectionStatus` parameter to **ConsoleProtection**.
	//
	// example:
	//
	// ManagedInstance
	ModificationProtectionReason *string `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	// Indicates whether the configuration read-only mode is enabled for the CLB instance. Valid values:
	//
	// 	- **NonProtection:*	- The configuration read-only mode is disabled. In this case, you cannot specify the ModificationProtectionReason parameter. If you specify the `ModificationProtectionReason` parameter, the value is cleared.
	//
	// 	- **ConsoleProtection:*	- The configuration read-only mode is enabled.
	//
	// >  If you set this parameter to **ConsoleProtection**, you cannot modify the configurations of the CLB instance in the CLB console. However, you can call API operations to modify the configurations of the CLB instance.
	//
	// example:
	//
	// ConsoleProtection
	ModificationProtectionStatus *string `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	// The network type of the internal-facing CLB instance. Valid values:
	//
	// 	- **vpc**: VPC
	//
	// 	- **Classic**: classic network
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the CLB instance.
	//
	// 	- **PayOnDemand*	- is returned, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PayOnDemand
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the region where the CLB instance was deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region where the CLB instance was deployed.
	//
	// example:
	//
	// hangzhou
	RegionIdAlias *string `json:"RegionIdAlias,omitempty" xml:"RegionIdAlias,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3r****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the secondary zone to which the CLB instance belongs.
	//
	// example:
	//
	// cn-hangzhou-d
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	// The tags.
	Tags *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vSwitch to which the internal-facing CLB instance belongs.
	//
	// example:
	//
	// vsw-255ecr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC in which the internal-facing CLB instance was deployed.
	//
	// example:
	//
	// vpc-25dvzy9f8****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetAddress() *string {
	return s.Address
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetDeleteProtection() *string {
	return s.DeleteProtection
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetInternetChargeTypeAlias() *string {
	return s.InternetChargeTypeAlias
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetModificationProtectionReason() *string {
	return s.ModificationProtectionReason
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetModificationProtectionStatus() *string {
	return s.ModificationProtectionStatus
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetPayType() *string {
	return s.PayType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetRegionIdAlias() *string {
	return s.RegionIdAlias
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetSlaveZoneId() *string {
	return s.SlaveZoneId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetTags() *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags {
	return s.Tags
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddress(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddressIPVersion(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddressType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetBandwidth(v int32) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetCreateTime(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetCreateTimeStamp(v int64) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetDeleteProtection(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.DeleteProtection = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetInstanceChargeType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetInternetChargeType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetInternetChargeTypeAlias(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.InternetChargeTypeAlias = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerName(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerSpec(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerStatus(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetMasterZoneId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetModificationProtectionReason(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.ModificationProtectionReason = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetModificationProtectionStatus(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetNetworkType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetPayType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetRegionId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetRegionIdAlias(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.RegionIdAlias = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetResourceGroupId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetSlaveZoneId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetTags(v *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.Tags = v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetVSwitchId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetVpcId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.VpcId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags struct {
	Tag []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) GetTag() []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag {
	return s.Tag
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) SetTag(v []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags {
	s.Tag = v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) Validate() error {
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

type DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// FinanceDept
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// FinanceJoshua
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) SetTagKey(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) SetTagValue(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) Validate() error {
	return dara.Validate(s)
}
