// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostGroups(v *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) *DescribeDedicatedHostGroupsResponseBody
	GetDedicatedHostGroups() *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups
	SetRequestId(v string) *DescribeDedicatedHostGroupsResponseBody
	GetRequestId() *string
}

type DescribeDedicatedHostGroupsResponseBody struct {
	// The information about dedicated clusters returned.
	DedicatedHostGroups *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups `json:"DedicatedHostGroups,omitempty" xml:"DedicatedHostGroups,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AB44DC0A-7E77-442A-97A9-C6418694CB22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedHostGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBody) GetDedicatedHostGroups() *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups {
	return s.DedicatedHostGroups
}

func (s *DescribeDedicatedHostGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedHostGroupsResponseBody) SetDedicatedHostGroups(v *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) *DescribeDedicatedHostGroupsResponseBody {
	s.DedicatedHostGroups = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBody) SetRequestId(v string) *DescribeDedicatedHostGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBody) Validate() error {
	if s.DedicatedHostGroups != nil {
		if err := s.DedicatedHostGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups struct {
	DedicatedHostGroups []*DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups `json:"DedicatedHostGroups,omitempty" xml:"DedicatedHostGroups,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) GetDedicatedHostGroups() []*DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	return s.DedicatedHostGroups
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) SetDedicatedHostGroups(v []*DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups {
	s.DedicatedHostGroups = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) Validate() error {
	if s.DedicatedHostGroups != nil {
		for _, item := range s.DedicatedHostGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups struct {
	// The policy based on which the system allocates resources in the dedicated cluster. Valid values:
	//
	// 	- **Evenly**: The system evenly allocates the resources to all the hosts in the dedicated cluster.
	//
	// 	- **Intensively**: The system preferentially allocates the resources to the heavily loaded hosts in the dedicated cluster.
	//
	// example:
	//
	// Evenly
	AllocationPolicy *string `json:"AllocationPolicy,omitempty" xml:"AllocationPolicy,omitempty"`
	// The ID of the bastion host.
	//
	// example:
	//
	// bastionhost-cn-m7xxxxxxxx
	BastionInstanceId *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	// The percentage of allocated cores in the dedicated cluster. Unit: %.
	//
	// example:
	//
	// 25
	CpuAllocateRation *float32 `json:"CpuAllocateRation,omitempty" xml:"CpuAllocateRation,omitempty"`
	// The number of allocated cores in the dedicated cluster.
	//
	// example:
	//
	// 8
	CpuAllocatedAmount *float32 `json:"CpuAllocatedAmount,omitempty" xml:"CpuAllocatedAmount,omitempty"`
	// The core overcommitment ratio of the dedicated cluster. Unit: %. For more information about the core overcommitment ratio, see [Manage a dedicated cluster](https://help.aliyun.com/document_detail/182328.html).
	//
	// example:
	//
	// 200
	CpuAllocationRatio *int32 `json:"CpuAllocationRatio,omitempty" xml:"CpuAllocationRatio,omitempty"`
	// The timestamp when the dedicated cluster was created.
	//
	// example:
	//
	// 1571125370000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of storage media that is used for the hosts in the dedicated cluster. Valid values:
	//
	// 	- **dhg_cloud_ssd**: cloud disks
	//
	// 	- **dhg_local_ssd**: local disks
	//
	// example:
	//
	// dhg_cloud_ssd
	DedicatedHostCountGroupByHostType map[string]interface{} `json:"DedicatedHostCountGroupByHostType,omitempty" xml:"DedicatedHostCountGroupByHostType,omitempty"`
	// The name of the dedicated cluster.
	//
	// example:
	//
	// testHostGroup
	DedicatedHostGroupDesc *string `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
	// The ID of the dedicated cluster.
	//
	// example:
	//
	// dhg-7a9xxxxxxxx
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The percentage of allocated disk space in the dedicated cluster. Unit: %.
	//
	// example:
	//
	// 0.49
	DiskAllocateRation *float32 `json:"DiskAllocateRation,omitempty" xml:"DiskAllocateRation,omitempty"`
	// The amount of allocated disk space in the dedicated cluster. Unit: GB.
	//
	// example:
	//
	// 200
	DiskAllocatedAmount *float32 `json:"DiskAllocatedAmount,omitempty" xml:"DiskAllocatedAmount,omitempty"`
	// The disk overcommitment ratio of the dedicated cluster. Unit: %. For more information about the core overcommitment ratio, see [Manage a dedicated cluster](https://help.aliyun.com/document_detail/182328.html).
	//
	// example:
	//
	// 200
	DiskAllocationRatio *int32 `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	// The amount of used disk space in the dedicated cluster. Unit: GB.
	//
	// example:
	//
	// 20
	DiskUsedAmount *float32 `json:"DiskUsedAmount,omitempty" xml:"DiskUsedAmount,omitempty"`
	// The disk usage of the dedicated cluster. Unit: %.
	//
	// example:
	//
	// 0
	DiskUtility *float32 `json:"DiskUtility,omitempty" xml:"DiskUtility,omitempty"`
	// The database engine of the instances in the dedicated cluster.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The total number of hosts in the dedicated cluster.
	//
	// example:
	//
	// 3
	HostNumber *int32 `json:"HostNumber,omitempty" xml:"HostNumber,omitempty"`
	// The policy that is used to handle host failures. Valid values:
	//
	// 	- **Auto**: The system automatically replaces faulty hosts.
	//
	// 	- **Manual**: You must manually replace faulty hosts.
	//
	// example:
	//
	// Auto
	HostReplacePolicy *string `json:"HostReplacePolicy,omitempty" xml:"HostReplacePolicy,omitempty"`
	// The total number of instances in the dedicated cluster.
	//
	// example:
	//
	// 4
	InstanceNumber *int32 `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	// The percentage of allocated memory space in the dedicated cluster. Unit: %.
	//
	// example:
	//
	// 33.7
	MemAllocateRation *float32 `json:"MemAllocateRation,omitempty" xml:"MemAllocateRation,omitempty"`
	// The amount of allocated memory space in the dedicated cluster.
	//
	// example:
	//
	// 16384
	MemAllocatedAmount *float32 `json:"MemAllocatedAmount,omitempty" xml:"MemAllocatedAmount,omitempty"`
	// The memory overcommitment ratio of the dedicated cluster. Unit: %. For more information about the core overcommitment ratio, see [Manage a dedicated cluster](https://help.aliyun.com/document_detail/182328.html).
	//
	// example:
	//
	// 90
	MemAllocationRatio *int32 `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	// The amount of used memory space in the dedicated cluster. Unit: MB.
	//
	// example:
	//
	// 0
	MemUsedAmount *float32 `json:"MemUsedAmount,omitempty" xml:"MemUsedAmount,omitempty"`
	// The memory usage of the dedicated cluster. Unit: %.
	//
	// example:
	//
	// 0
	MemUtility *float32 `json:"MemUtility,omitempty" xml:"MemUtility,omitempty"`
	// Indicates whether the feature that allows you to have the OS permissions on the host is enabled. Valid values:
	//
	// 	- **0*	- or **null**: The permissions cannot be granted.
	//
	// 	- **1**: The permissions can be granted.
	//
	// 	- **3**: The permissions have been granted.
	//
	// example:
	//
	// 3
	OpenPermission *string `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	// The name and ID of the dedicated cluster. The value consists of **DedicatedHostGroupDesc*	- and **DedicatedHostGroupId**. Format: DedicatedHostGroupDesc/DedicatedHostGroupId.
	//
	// example:
	//
	// testHostGroup/dhg-7a9xxxxxxxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the dedicated cluster belongs.
	//
	// example:
	//
	// vpc-bp1oxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The zones to which the hosts of the dedicated cluster belong.
	ZoneIDList *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList `json:"ZoneIDList,omitempty" xml:"ZoneIDList,omitempty" type:"Struct"`
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetAllocationPolicy() *string {
	return s.AllocationPolicy
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetBastionInstanceId() *string {
	return s.BastionInstanceId
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetCpuAllocateRation() *float32 {
	return s.CpuAllocateRation
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetCpuAllocatedAmount() *float32 {
	return s.CpuAllocatedAmount
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetCpuAllocationRatio() *int32 {
	return s.CpuAllocationRatio
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetDedicatedHostCountGroupByHostType() map[string]interface{} {
	return s.DedicatedHostCountGroupByHostType
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetDedicatedHostGroupDesc() *string {
	return s.DedicatedHostGroupDesc
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetDiskAllocateRation() *float32 {
	return s.DiskAllocateRation
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetDiskAllocatedAmount() *float32 {
	return s.DiskAllocatedAmount
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetDiskAllocationRatio() *int32 {
	return s.DiskAllocationRatio
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetDiskUsedAmount() *float32 {
	return s.DiskUsedAmount
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetDiskUtility() *float32 {
	return s.DiskUtility
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetHostNumber() *int32 {
	return s.HostNumber
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetHostReplacePolicy() *string {
	return s.HostReplacePolicy
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetInstanceNumber() *int32 {
	return s.InstanceNumber
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetMemAllocateRation() *float32 {
	return s.MemAllocateRation
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetMemAllocatedAmount() *float32 {
	return s.MemAllocatedAmount
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetMemAllocationRatio() *int32 {
	return s.MemAllocationRatio
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetMemUsedAmount() *float32 {
	return s.MemUsedAmount
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetMemUtility() *float32 {
	return s.MemUtility
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetOpenPermission() *string {
	return s.OpenPermission
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetText() *string {
	return s.Text
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GetZoneIDList() *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList {
	return s.ZoneIDList
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetAllocationPolicy(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.AllocationPolicy = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetBastionInstanceId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.BastionInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCreateTime(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDedicatedHostCountGroupByHostType(v map[string]interface{}) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DedicatedHostCountGroupByHostType = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDedicatedHostGroupDesc(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DedicatedHostGroupDesc = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskUsedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskUsedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskUtility(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskUtility = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetEngine(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetHostNumber(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.HostNumber = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetHostReplacePolicy(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.HostReplacePolicy = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetInstanceNumber(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemUsedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemUsedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemUtility(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemUtility = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetOpenPermission(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.OpenPermission = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetText(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.Text = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetVPCId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.VPCId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetZoneIDList(v *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.ZoneIDList = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) Validate() error {
	if s.ZoneIDList != nil {
		if err := s.ZoneIDList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList struct {
	ZoneIDList []*string `json:"ZoneIDList,omitempty" xml:"ZoneIDList,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) GetZoneIDList() []*string {
	return s.ZoneIDList
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) SetZoneIDList(v []*string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList {
	s.ZoneIDList = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) Validate() error {
	return dara.Validate(s)
}
