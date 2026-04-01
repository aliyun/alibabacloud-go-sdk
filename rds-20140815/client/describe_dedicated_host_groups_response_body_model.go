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
	AllocationPolicy                  *string                                                                                  `json:"AllocationPolicy,omitempty" xml:"AllocationPolicy,omitempty"`
	BastionInstanceId                 *string                                                                                  `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	CpuAllocateRation                 *float32                                                                                 `json:"CpuAllocateRation,omitempty" xml:"CpuAllocateRation,omitempty"`
	CpuAllocatedAmount                *float32                                                                                 `json:"CpuAllocatedAmount,omitempty" xml:"CpuAllocatedAmount,omitempty"`
	CpuAllocationRatio                *int32                                                                                   `json:"CpuAllocationRatio,omitempty" xml:"CpuAllocationRatio,omitempty"`
	CreateTime                        *string                                                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DedicatedHostCountGroupByHostType map[string]interface{}                                                                   `json:"DedicatedHostCountGroupByHostType,omitempty" xml:"DedicatedHostCountGroupByHostType,omitempty"`
	DedicatedHostGroupDesc            *string                                                                                  `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
	DedicatedHostGroupId              *string                                                                                  `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DiskAllocateRation                *float32                                                                                 `json:"DiskAllocateRation,omitempty" xml:"DiskAllocateRation,omitempty"`
	DiskAllocatedAmount               *float32                                                                                 `json:"DiskAllocatedAmount,omitempty" xml:"DiskAllocatedAmount,omitempty"`
	DiskAllocationRatio               *int32                                                                                   `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	DiskUsedAmount                    *float32                                                                                 `json:"DiskUsedAmount,omitempty" xml:"DiskUsedAmount,omitempty"`
	DiskUtility                       *float32                                                                                 `json:"DiskUtility,omitempty" xml:"DiskUtility,omitempty"`
	Engine                            *string                                                                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
	HostNumber                        *int32                                                                                   `json:"HostNumber,omitempty" xml:"HostNumber,omitempty"`
	HostReplacePolicy                 *string                                                                                  `json:"HostReplacePolicy,omitempty" xml:"HostReplacePolicy,omitempty"`
	InstanceNumber                    *int32                                                                                   `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	MemAllocateRation                 *float32                                                                                 `json:"MemAllocateRation,omitempty" xml:"MemAllocateRation,omitempty"`
	MemAllocatedAmount                *float32                                                                                 `json:"MemAllocatedAmount,omitempty" xml:"MemAllocatedAmount,omitempty"`
	MemAllocationRatio                *int32                                                                                   `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	MemUsedAmount                     *float32                                                                                 `json:"MemUsedAmount,omitempty" xml:"MemUsedAmount,omitempty"`
	MemUtility                        *float32                                                                                 `json:"MemUtility,omitempty" xml:"MemUtility,omitempty"`
	OpenPermission                    *string                                                                                  `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	Text                              *string                                                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
	VPCId                             *string                                                                                  `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	ZoneIDList                        *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList `json:"ZoneIDList,omitempty" xml:"ZoneIDList,omitempty" type:"Struct"`
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
