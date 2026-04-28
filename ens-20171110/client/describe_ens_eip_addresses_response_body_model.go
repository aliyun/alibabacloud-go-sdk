// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsEipAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipAddresses(v *DescribeEnsEipAddressesResponseBodyEipAddresses) *DescribeEnsEipAddressesResponseBody
	GetEipAddresses() *DescribeEnsEipAddressesResponseBodyEipAddresses
	SetPageNumber(v int32) *DescribeEnsEipAddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnsEipAddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEnsEipAddressesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEnsEipAddressesResponseBody
	GetTotalCount() *int32
}

type DescribeEnsEipAddressesResponseBody struct {
	// Details of the EIPs.
	//
	// example:
	//
	// [\\"106.14.194.193\\"]
	EipAddresses *DescribeEnsEipAddressesResponseBodyEipAddresses `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Struct"`
	// The page number. Valid values: an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: **10*	- to **100**.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8629F679-B51D-4194-A1CC-5D8F504C362B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEnsEipAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBody) GetEipAddresses() *DescribeEnsEipAddressesResponseBodyEipAddresses {
	return s.EipAddresses
}

func (s *DescribeEnsEipAddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnsEipAddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnsEipAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsEipAddressesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEnsEipAddressesResponseBody) SetEipAddresses(v *DescribeEnsEipAddressesResponseBodyEipAddresses) *DescribeEnsEipAddressesResponseBody {
	s.EipAddresses = v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) SetPageNumber(v int32) *DescribeEnsEipAddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) SetPageSize(v int32) *DescribeEnsEipAddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) SetRequestId(v string) *DescribeEnsEipAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) SetTotalCount(v int32) *DescribeEnsEipAddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBody) Validate() error {
	if s.EipAddresses != nil {
		if err := s.EipAddresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnsEipAddressesResponseBodyEipAddresses struct {
	EipAddress []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Repeated"`
}

func (s DescribeEnsEipAddressesResponseBodyEipAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBodyEipAddresses) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddresses) GetEipAddress() []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	return s.EipAddress
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddresses) SetEipAddress(v []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) *DescribeEnsEipAddressesResponseBodyEipAddresses {
	s.EipAddress = v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddresses) Validate() error {
	if s.EipAddress != nil {
		for _, item := range s.EipAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress struct {
	// EIP的实例ID。
	//
	// example:
	//
	// eip-5sainglpw7qfem3icir4s****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// EIP的创建时间
	//
	// example:
	//
	// 1624885274000
	AllocationTime *string `json:"AllocationTime,omitempty" xml:"AllocationTime,omitempty"`
	// EIP的带宽峰值，默认值为5。取值范围：**5**~**10000**，单位：Mbps。
	//
	// example:
	//
	// 50
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// example:
	//
	// 100
	BandwidthPackageBandwidth *int32 `json:"BandwidthPackageBandwidth,omitempty" xml:"BandwidthPackageBandwidth,omitempty"`
	// example:
	//
	// cbwp-5***
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// EIP的计费模式。
	//
	// - **PrePaid**：包年包月。
	//
	// - **PostPaid**：按量计费。
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// EIP的描述信息。
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ENS节点ID。
	//
	// example:
	//
	// cn-xian-telecom
	EnsRegionId      *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	IcmpReplyEnabled *bool   `json:"IcmpReplyEnabled,omitempty" xml:"IcmpReplyEnabled,omitempty"`
	// 当前绑定的实例的ID。
	//
	// example:
	//
	// lb-5sc3kum2e0sz34wbqrws9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 要绑定的云产品实例的类型，取值：
	//
	// - **EnsInstance**：VPC类型的ENS实例。
	//
	// - **SlbInstance**：负载均衡实例。
	//
	// example:
	//
	// SlbInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// EIP的计费方式。
	//
	// - **95BandwidthByMonth**：月95峰值带宽。
	//
	// - **PayByBandwidth**：固定带宽计费。
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// EIP的IP地址。
	//
	// example:
	//
	// 120.XXX.XXX.4
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// 如果该弹性IP是备用EIP时，表示当前EIP的运行状态。
	//
	// - **Stopped**：已停止。
	//
	// - **Running**：运行中。
	//
	// - **Starting**：启动中。
	//
	// - **Stopping**：停止中。
	//
	// example:
	//
	// Stopped
	IpStatus *string `json:"IpStatus,omitempty" xml:"IpStatus,omitempty"`
	// 运营商，取值：
	//
	// - **cmcc**：中国移动。
	//
	// - **unicom**：中国联通。
	//
	// - **telecom**：中国电信。
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// EIP实例名称。
	//
	// example:
	//
	// test
	Name           *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationLocks *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// 该EIP是否是备用。
	//
	// example:
	//
	// true
	Standby *bool `json:"Standby,omitempty" xml:"Standby,omitempty"`
	// EIP的状态。
	//
	// - **Associating**：绑定中。
	//
	// - **Unassociating**：解绑中。
	//
	// - **InUse**：已分配。
	//
	// - **Available**：可用。
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 标签集合。
	Tags *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetAllocationTime() *string {
	return s.AllocationTime
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetBandwidthPackageBandwidth() *int32 {
	return s.BandwidthPackageBandwidth
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetDescription() *string {
	return s.Description
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetIcmpReplyEnabled() *bool {
	return s.IcmpReplyEnabled
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetIpStatus() *string {
	return s.IpStatus
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetIsp() *string {
	return s.Isp
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetName() *string {
	return s.Name
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetOperationLocks() *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks {
	return s.OperationLocks
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetStandby() *bool {
	return s.Standby
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) GetTags() *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags {
	return s.Tags
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetAllocationId(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.AllocationId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetAllocationTime(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.AllocationTime = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidth(v int32) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Bandwidth = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidthPackageBandwidth(v int32) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.BandwidthPackageBandwidth = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidthPackageId(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetChargeType(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.ChargeType = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetDescription(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Description = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetEnsRegionId(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetIcmpReplyEnabled(v bool) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.IcmpReplyEnabled = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceId(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceId = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceType(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceType = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetInternetChargeType(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetIpAddress(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.IpAddress = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetIpStatus(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.IpStatus = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetIsp(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Isp = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetName(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Name = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetOperationLocks(v *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.OperationLocks = v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetStandby(v bool) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Standby = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetStatus(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Status = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) SetTags(v *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress {
	s.Tags = v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddress) Validate() error {
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
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

type DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks struct {
	Lock []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock `json:"Lock,omitempty" xml:"Lock,omitempty" type:"Repeated"`
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) GetLock() []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock {
	return s.Lock
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) SetLock(v []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks {
	s.Lock = v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) Validate() error {
	if s.Lock != nil {
		for _, item := range s.Lock {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock struct {
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock) SetLockReason(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock {
	s.LockReason = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLock) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags struct {
	Tag []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags) GetTag() []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag {
	return s.Tag
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags) SetTag(v []*DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags {
	s.Tag = v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTags) Validate() error {
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

type DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag struct {
	// 标签键
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Deprecated
	//
	// 标签键。
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// Deprecated
	//
	// 标签值。
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// 标签值。
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) SetKey(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) SetTagKey(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) SetTagValue(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) SetValue(v string) *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeEnsEipAddressesResponseBodyEipAddressesEipAddressTagsTag) Validate() error {
	return dara.Validate(s)
}
