// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedClusterMonitorRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuAlarmThreshold(v int64) *CreateDedicatedClusterMonitorRuleRequest
	GetCpuAlarmThreshold() *int64
	SetDedicatedClusterId(v string) *CreateDedicatedClusterMonitorRuleRequest
	GetDedicatedClusterId() *string
	SetDiskAlarmThreshold(v int64) *CreateDedicatedClusterMonitorRuleRequest
	GetDiskAlarmThreshold() *int64
	SetDuAlarmThreshold(v int64) *CreateDedicatedClusterMonitorRuleRequest
	GetDuAlarmThreshold() *int64
	SetInstanceId(v string) *CreateDedicatedClusterMonitorRuleRequest
	GetInstanceId() *string
	SetMemAlarmThreshold(v int64) *CreateDedicatedClusterMonitorRuleRequest
	GetMemAlarmThreshold() *int64
	SetNoticeSwitch(v int64) *CreateDedicatedClusterMonitorRuleRequest
	GetNoticeSwitch() *int64
	SetOwnerId(v string) *CreateDedicatedClusterMonitorRuleRequest
	GetOwnerId() *string
	SetPhones(v string) *CreateDedicatedClusterMonitorRuleRequest
	GetPhones() *string
	SetRegionId(v string) *CreateDedicatedClusterMonitorRuleRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDedicatedClusterMonitorRuleRequest
	GetResourceGroupId() *string
}

type CreateDedicatedClusterMonitorRuleRequest struct {
	// The alert threshold for CPU utilization. Unit: percentage.
	//
	// example:
	//
	// 30
	CpuAlarmThreshold *int64 `json:"CpuAlarmThreshold,omitempty" xml:"CpuAlarmThreshold,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// dts-dasd22******
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The alert threshold for disk usage. Unit: percentage.
	//
	// example:
	//
	// 100
	DiskAlarmThreshold *int64 `json:"DiskAlarmThreshold,omitempty" xml:"DiskAlarmThreshold,omitempty"`
	// The alert threshold for DTS Unit (DU) usage. Unit: percentage.
	//
	// example:
	//
	// 20
	DuAlarmThreshold *int64 `json:"DuAlarmThreshold,omitempty" xml:"DuAlarmThreshold,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-bp1162kryivb8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The alert threshold for memory usage. Unit: percentage.
	//
	// example:
	//
	// 40
	MemAlarmThreshold *int64 `json:"MemAlarmThreshold,omitempty" xml:"MemAlarmThreshold,omitempty"`
	// Specifies whether to enable the alert feature. Valid values:
	//
	// 	- **1**: enables the alert feature.
	//
	// 	- **0**: disables the alert feature.
	//
	// example:
	//
	// 1
	NoticeSwitch *int64  `json:"NoticeSwitch,omitempty" xml:"NoticeSwitch,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The mobile phone number to which alerts are sent. Separate multiple mobile phone numbers with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 186****7654
	Phones *string `json:"Phones,omitempty" xml:"Phones,omitempty"`
	// The ID of the region in which the Data Transmission Service (DTS) instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This parameter is a global parameter and not required.
	//
	// example:
	//
	// The resource group ID. This parameter is a global parameter and not required.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDedicatedClusterMonitorRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedClusterMonitorRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetCpuAlarmThreshold() *int64 {
	return s.CpuAlarmThreshold
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetDiskAlarmThreshold() *int64 {
	return s.DiskAlarmThreshold
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetDuAlarmThreshold() *int64 {
	return s.DuAlarmThreshold
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetMemAlarmThreshold() *int64 {
	return s.MemAlarmThreshold
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetNoticeSwitch() *int64 {
	return s.NoticeSwitch
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetPhones() *string {
	return s.Phones
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDedicatedClusterMonitorRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetCpuAlarmThreshold(v int64) *CreateDedicatedClusterMonitorRuleRequest {
	s.CpuAlarmThreshold = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetDedicatedClusterId(v string) *CreateDedicatedClusterMonitorRuleRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetDiskAlarmThreshold(v int64) *CreateDedicatedClusterMonitorRuleRequest {
	s.DiskAlarmThreshold = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetDuAlarmThreshold(v int64) *CreateDedicatedClusterMonitorRuleRequest {
	s.DuAlarmThreshold = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetInstanceId(v string) *CreateDedicatedClusterMonitorRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetMemAlarmThreshold(v int64) *CreateDedicatedClusterMonitorRuleRequest {
	s.MemAlarmThreshold = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetNoticeSwitch(v int64) *CreateDedicatedClusterMonitorRuleRequest {
	s.NoticeSwitch = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetOwnerId(v string) *CreateDedicatedClusterMonitorRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetPhones(v string) *CreateDedicatedClusterMonitorRuleRequest {
	s.Phones = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetRegionId(v string) *CreateDedicatedClusterMonitorRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) SetResourceGroupId(v string) *CreateDedicatedClusterMonitorRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleRequest) Validate() error {
	return dara.Validate(s)
}
