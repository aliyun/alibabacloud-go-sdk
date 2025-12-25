// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() []*DescribeInstancesResponseBodyInstances
	SetPageIndex(v int32) *DescribeInstancesResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeInstancesResponseBody
	GetTotalCount() *int64
	SetTotalPage(v int32) *DescribeInstancesResponseBody
	GetTotalPage() *int32
}

type DescribeInstancesResponseBody struct {
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() []*DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageIndex(v int32) *DescribeInstancesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetSuccess(v bool) *DescribeInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalPage(v int32) *DescribeInstancesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// c20c******404
	AskClusterId *string `json:"AskClusterId,omitempty" xml:"AskClusterId,omitempty"`
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// RUNNING
	ClusterStatus  *string                                               `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	Ha             *bool                                                 `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec *DescribeInstancesResponseBodyInstancesHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	HaVSwitchIds   []*string                                             `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	HaZoneId       *string                                               `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	HostAliases []*DescribeInstancesResponseBodyInstancesHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// example:
	//
	// sc_flinkserverlesspost_public_cn-0ju2bj2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// vvp1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType  *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// NORMAL
	OrderState *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1629879567394
	ResourceCreateTime *int64 `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// example:
	//
	// 1637337600000
	ResourceExpiredTime *int64 `json:"ResourceExpiredTime,omitempty" xml:"ResourceExpiredTime,omitempty"`
	// example:
	//
	// rg-***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// b3690a1655da47
	ResourceId   *string                                             `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceSpec *DescribeInstancesResponseBodyInstancesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage      *DescribeInstancesResponseBodyInstancesStorage      `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Tags         []*DescribeInstancesResponseBodyInstancesTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 183899668736****
	Uid        *string   `json:"Uid,omitempty" xml:"Uid,omitempty"`
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-2ze9*******nxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesResponseBodyInstances) GetAskClusterId() *string {
	return s.AskClusterId
}

func (s *DescribeInstancesResponseBodyInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesResponseBodyInstances) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *DescribeInstancesResponseBodyInstances) GetHa() *bool {
	return s.Ha
}

func (s *DescribeInstancesResponseBodyInstances) GetHaResourceSpec() *DescribeInstancesResponseBodyInstancesHaResourceSpec {
	return s.HaResourceSpec
}

func (s *DescribeInstancesResponseBodyInstances) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *DescribeInstancesResponseBodyInstances) GetHaZoneId() *string {
	return s.HaZoneId
}

func (s *DescribeInstancesResponseBodyInstances) GetHostAliases() []*DescribeInstancesResponseBodyInstancesHostAliases {
	return s.HostAliases
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesResponseBodyInstances) GetMonitorType() *string {
	return s.MonitorType
}

func (s *DescribeInstancesResponseBodyInstances) GetOrderState() *string {
	return s.OrderState
}

func (s *DescribeInstancesResponseBodyInstances) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceCreateTime() *int64 {
	return s.ResourceCreateTime
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceExpiredTime() *int64 {
	return s.ResourceExpiredTime
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceSpec() *DescribeInstancesResponseBodyInstancesResourceSpec {
	return s.ResourceSpec
}

func (s *DescribeInstancesResponseBodyInstances) GetStorage() *DescribeInstancesResponseBodyInstancesStorage {
	return s.Storage
}

func (s *DescribeInstancesResponseBodyInstances) GetTags() []*DescribeInstancesResponseBodyInstancesTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyInstances) GetUid() *string {
	return s.Uid
}

func (s *DescribeInstancesResponseBodyInstances) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeInstancesResponseBodyInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstances) SetArchitectureType(v string) *DescribeInstancesResponseBodyInstances {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetAskClusterId(v string) *DescribeInstancesResponseBodyInstances {
	s.AskClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetChargeType(v string) *DescribeInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.ClusterStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHa(v bool) *DescribeInstancesResponseBodyInstances {
	s.Ha = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaResourceSpec(v *DescribeInstancesResponseBodyInstancesHaResourceSpec) *DescribeInstancesResponseBodyInstances {
	s.HaResourceSpec = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaVSwitchIds(v []*string) *DescribeInstancesResponseBodyInstances {
	s.HaVSwitchIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaZoneId(v string) *DescribeInstancesResponseBodyInstances {
	s.HaZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHostAliases(v []*DescribeInstancesResponseBodyInstancesHostAliases) *DescribeInstancesResponseBodyInstances {
	s.HostAliases = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMonitorType(v string) *DescribeInstancesResponseBodyInstances {
	s.MonitorType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetOrderState(v string) *DescribeInstancesResponseBodyInstances {
	s.OrderState = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegion(v string) *DescribeInstancesResponseBodyInstances {
	s.Region = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceCreateTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ResourceCreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceExpiredTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ResourceExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceSpec(v *DescribeInstancesResponseBodyInstancesResourceSpec) *DescribeInstancesResponseBodyInstances {
	s.ResourceSpec = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStorage(v *DescribeInstancesResponseBodyInstancesStorage) *DescribeInstancesResponseBodyInstances {
	s.Storage = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetTags(v []*DescribeInstancesResponseBodyInstancesTags) *DescribeInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUid(v string) *DescribeInstancesResponseBodyInstances {
	s.Uid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVSwitchIds(v []*string) *DescribeInstancesResponseBodyInstances {
	s.VSwitchIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetZoneId(v string) *DescribeInstancesResponseBodyInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	if s.HaResourceSpec != nil {
		if err := s.HaResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.HostAliases != nil {
		for _, item := range s.HostAliases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Storage != nil {
		if err := s.Storage.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesHaResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) SetCpu(v int32) *DescribeInstancesResponseBodyInstancesHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) SetMemoryGB(v int32) *DescribeInstancesResponseBodyInstancesHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesHostAliases struct {
	// This parameter is required.
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// This parameter is required.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesHostAliases) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) GetHostNames() []*string {
	return s.HostNames
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) SetHostNames(v []*string) *DescribeInstancesResponseBodyInstancesHostAliases {
	s.HostNames = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) SetIp(v string) *DescribeInstancesResponseBodyInstancesHostAliases {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesResourceSpec struct {
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) SetCpu(v int32) *DescribeInstancesResponseBodyInstancesResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) SetMemoryGB(v int32) *DescribeInstancesResponseBodyInstancesResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesStorage struct {
	Oss *DescribeInstancesResponseBodyInstancesStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s DescribeInstancesResponseBodyInstancesStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesStorage) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesStorage) GetOss() *DescribeInstancesResponseBodyInstancesStorageOss {
	return s.Oss
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetOss(v *DescribeInstancesResponseBodyInstancesStorageOss) *DescribeInstancesResponseBodyInstancesStorage {
	s.Oss = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorage) Validate() error {
	if s.Oss != nil {
		if err := s.Oss.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesStorageOss struct {
	// example:
	//
	// oss_flink
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesStorageOss) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesStorageOss) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesStorageOss) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeInstancesResponseBodyInstancesStorageOss) SetBucket(v string) *DescribeInstancesResponseBodyInstancesStorageOss {
	s.Bucket = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorageOss) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesTags struct {
	// example:
	//
	// flink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesTags) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesResponseBodyInstancesTags) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesResponseBodyInstancesTags) SetKey(v string) *DescribeInstancesResponseBodyInstancesTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesTags) SetValue(v string) *DescribeInstancesResponseBodyInstancesTags {
	s.Value = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesTags) Validate() error {
	return dara.Validate(s)
}
