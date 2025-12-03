// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *DescribeLoadBalancersRequest
	GetAddress() *string
	SetAddressIPVersion(v string) *DescribeLoadBalancersRequest
	GetAddressIPVersion() *string
	SetAddressType(v string) *DescribeLoadBalancersRequest
	GetAddressType() *string
	SetInternetChargeType(v string) *DescribeLoadBalancersRequest
	GetInternetChargeType() *string
	SetLoadBalancerId(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerName() *string
	SetLoadBalancerStatus(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerStatus() *string
	SetMasterZoneId(v string) *DescribeLoadBalancersRequest
	GetMasterZoneId() *string
	SetNetworkType(v string) *DescribeLoadBalancersRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeLoadBalancersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLoadBalancersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadBalancersRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeLoadBalancersRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeLoadBalancersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeLoadBalancersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadBalancersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancersRequest
	GetResourceOwnerId() *int64
	SetServerId(v string) *DescribeLoadBalancersRequest
	GetServerId() *string
	SetServerIntranetAddress(v string) *DescribeLoadBalancersRequest
	GetServerIntranetAddress() *string
	SetSlaveZoneId(v string) *DescribeLoadBalancersRequest
	GetSlaveZoneId() *string
	SetTag(v []*DescribeLoadBalancersRequestTag) *DescribeLoadBalancersRequest
	GetTag() []*DescribeLoadBalancersRequestTag
	SetTags(v string) *DescribeLoadBalancersRequest
	GetTags() *string
	SetVSwitchId(v string) *DescribeLoadBalancersRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeLoadBalancersRequest
	GetVpcId() *string
}

type DescribeLoadBalancersRequest struct {
	// The IP address that the CLB instance uses to provide services.
	//
	// example:
	//
	// 192.168.XX.XX
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
	// The metering method of Internet data transfer. Valid values:
	//
	// 	- **paybybandwidth:*	- pay-by-bandwidth.
	//
	// 	- **paybytraffic:*	- pay-by-data-transfer.
	//
	// example:
	//
	// paybytraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the CLB instance.
	//
	// You can specify up to 10 IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// lb-bp1b6c719dfa****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the CLB instance.
	//
	// The name must be 1 to 80 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// You can specify up to 10 names. Separate multiple names with commas (,).
	//
	// example:
	//
	// test
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The status of the CLB instance. Valid values:
	//
	// 	- **inactive:*	- The CLB instance is disabled. CLB instances in the inactive state do not forward traffic.
	//
	// 	- **active:*	- The CLB instance runs as expected. By default, newly created CLB instances are in the **active*	- state.
	//
	// 	- **locked:*	- The CLB instance is locked. After a CLB instance expires, it is locked for seven days. A locked CLB instance cannot forward traffic and you cannot perform operations on the locked CLB instance. However, other settings such as the IP address are retained.
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
	// The network type of the internal-facing CLB instance. Valid values:
	//
	// 	- **vpc**: VPC
	//
	// 	- **Classic**: classic network
	//
	// example:
	//
	// vpc
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: **1*	- to **100**.
	//
	// >  If you specify the **PageSize*	- parameter, you must also specify the **PageNumber*	- parameter.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method of the CLB instance. Valid values:
	//
	// 	- Set the value to **PayOnDemand**.
	//
	// example:
	//
	// PayOnDemand
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the region where the CLB instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2401682.html) operation to query the most recent region list.
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
	// rg-acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the backend server that is added to the CLB instance.
	//
	// example:
	//
	// vm-server-23****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The private IP address of the backend server that is added to the CLB instance.
	//
	// You can specify multiple IP addresses. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 10.XX.XX.102
	ServerIntranetAddress *string `json:"ServerIntranetAddress,omitempty" xml:"ServerIntranetAddress,omitempty"`
	// The ID of the secondary zone to which the CLB instance belongs.
	//
	// CLB instances on Alibaba Finance Cloud do not support cross-zone deployment.
	//
	// example:
	//
	// cn-hangzhou-d
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	// The tags.
	Tag []*DescribeLoadBalancersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The tags that are added to the CLB instance. The tags must be key-value pairs that are contained in a JSON dictionary.
	//
	// You can specify up to 10 tags in each call.
	//
	// example:
	//
	// [{"tagKey":"Key1","tagValue":"Value1"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the vSwitch to which the CLB instance belongs.
	//
	// example:
	//
	// vsw-bp12mw1f8k3****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the CLB instance belongs.
	//
	// example:
	//
	// vpc-bp1aevy8sof****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeLoadBalancersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersRequest) GetAddress() *string {
	return s.Address
}

func (s *DescribeLoadBalancersRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeLoadBalancersRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeLoadBalancersRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *DescribeLoadBalancersRequest) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *DescribeLoadBalancersRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeLoadBalancersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadBalancersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadBalancersRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeLoadBalancersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLoadBalancersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancersRequest) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeLoadBalancersRequest) GetServerIntranetAddress() *string {
	return s.ServerIntranetAddress
}

func (s *DescribeLoadBalancersRequest) GetSlaveZoneId() *string {
	return s.SlaveZoneId
}

func (s *DescribeLoadBalancersRequest) GetTag() []*DescribeLoadBalancersRequestTag {
	return s.Tag
}

func (s *DescribeLoadBalancersRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeLoadBalancersRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancersRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLoadBalancersRequest) SetAddress(v string) *DescribeLoadBalancersRequest {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetAddressIPVersion(v string) *DescribeLoadBalancersRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetAddressType(v string) *DescribeLoadBalancersRequest {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetInternetChargeType(v string) *DescribeLoadBalancersRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerId(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerName(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerStatus(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetMasterZoneId(v string) *DescribeLoadBalancersRequest {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetNetworkType(v string) *DescribeLoadBalancersRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetOwnerAccount(v string) *DescribeLoadBalancersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetOwnerId(v int64) *DescribeLoadBalancersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPageNumber(v int32) *DescribeLoadBalancersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPageSize(v int32) *DescribeLoadBalancersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPayType(v string) *DescribeLoadBalancersRequest {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetRegionId(v string) *DescribeLoadBalancersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceGroupId(v string) *DescribeLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetServerId(v string) *DescribeLoadBalancersRequest {
	s.ServerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetServerIntranetAddress(v string) *DescribeLoadBalancersRequest {
	s.ServerIntranetAddress = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetSlaveZoneId(v string) *DescribeLoadBalancersRequest {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetTag(v []*DescribeLoadBalancersRequestTag) *DescribeLoadBalancersRequest {
	s.Tag = v
	return s
}

func (s *DescribeLoadBalancersRequest) SetTags(v string) *DescribeLoadBalancersRequest {
	s.Tags = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetVSwitchId(v string) *DescribeLoadBalancersRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetVpcId(v string) *DescribeLoadBalancersRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) Validate() error {
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

type DescribeLoadBalancersRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be at most 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLoadBalancersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeLoadBalancersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeLoadBalancersRequestTag) SetKey(v string) *DescribeLoadBalancersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeLoadBalancersRequestTag) SetValue(v string) *DescribeLoadBalancersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeLoadBalancersRequestTag) Validate() error {
	return dara.Validate(s)
}
