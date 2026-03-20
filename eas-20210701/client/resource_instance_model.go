// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceInstance interface {
	dara.Model
	String() string
	GoString() string
	SetArch(v string) *ResourceInstance
	GetArch() *string
	SetAutoRenewal(v bool) *ResourceInstance
	GetAutoRenewal() *bool
	SetChargeType(v string) *ResourceInstance
	GetChargeType() *string
	SetCreateTime(v string) *ResourceInstance
	GetCreateTime() *string
	SetExpiredTime(v string) *ResourceInstance
	GetExpiredTime() *string
	SetInstanceCpuCount(v int32) *ResourceInstance
	GetInstanceCpuCount() *int32
	SetInstanceGpuCount(v int32) *ResourceInstance
	GetInstanceGpuCount() *int32
	SetInstanceGpuMemory(v string) *ResourceInstance
	GetInstanceGpuMemory() *string
	SetInstanceId(v string) *ResourceInstance
	GetInstanceId() *string
	SetInstanceIp(v string) *ResourceInstance
	GetInstanceIp() *string
	SetInstanceMemory(v string) *ResourceInstance
	GetInstanceMemory() *string
	SetInstanceName(v string) *ResourceInstance
	GetInstanceName() *string
	SetInstancePhase(v string) *ResourceInstance
	GetInstancePhase() *string
	SetInstanceStatus(v string) *ResourceInstance
	GetInstanceStatus() *string
	SetInstanceSystemDiskSize(v int32) *ResourceInstance
	GetInstanceSystemDiskSize() *int32
	SetInstanceTenantIp(v string) *ResourceInstance
	GetInstanceTenantIp() *string
	SetInstanceType(v string) *ResourceInstance
	GetInstanceType() *string
	SetInstanceUsedCpu(v float32) *ResourceInstance
	GetInstanceUsedCpu() *float32
	SetInstanceUsedGpu(v float32) *ResourceInstance
	GetInstanceUsedGpu() *float32
	SetInstanceUsedGpuMemory(v string) *ResourceInstance
	GetInstanceUsedGpuMemory() *string
	SetInstanceUsedMemory(v string) *ResourceInstance
	GetInstanceUsedMemory() *string
	SetLabels(v []*ResourceInstanceLabels) *ResourceInstance
	GetLabels() []*ResourceInstanceLabels
	SetLastCordonOperator(v string) *ResourceInstance
	GetLastCordonOperator() *string
	SetLastCordonReason(v string) *ResourceInstance
	GetLastCordonReason() *string
	SetRegion(v string) *ResourceInstance
	GetRegion() *string
	SetResourceId(v string) *ResourceInstance
	GetResourceId() *string
	SetZone(v string) *ResourceInstance
	GetZone() *string
}

type ResourceInstance struct {
	// The system architecture of the instance.
	//
	// example:
	//
	// arm64
	Arch *string `json:"Arch,omitempty" xml:"Arch,omitempty"`
	// Indicates whether auto-renewal is enabled for the instance.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The billing method of the instance.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2020-07-05T22:51:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2020-08-05T22:51:32Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The number of CPU cores for the instance.
	//
	// example:
	//
	// 4
	InstanceCpuCount *int32 `json:"InstanceCpuCount,omitempty" xml:"InstanceCpuCount,omitempty"`
	// The number of GPUs for the instance.
	//
	// example:
	//
	// 0
	InstanceGpuCount *int32 `json:"InstanceGpuCount,omitempty" xml:"InstanceGpuCount,omitempty"`
	// The GPU memory size of the instance.
	//
	// example:
	//
	// 0G
	InstanceGpuMemory *string `json:"InstanceGpuMemory,omitempty" xml:"InstanceGpuMemory,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// eas-i-1800z74n30kao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 11.227.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The memory size of the instance.
	//
	// example:
	//
	// 8192M
	InstanceMemory *string `json:"InstanceMemory,omitempty" xml:"InstanceMemory,omitempty"`
	// The instance name.
	//
	// example:
	//
	// eas01122713204*****
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstancePhase *string `json:"InstancePhase,omitempty" xml:"InstancePhase,omitempty"`
	// The instance status.
	//
	// Valid values:
	//
	// 	- Ready-SchedulingDisabled
	//
	// 	- Ready
	//
	// 	- NotReady
	//
	// example:
	//
	// Ready
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The system disk size of the instance.
	//
	// example:
	//
	// 200
	InstanceSystemDiskSize *int32 `json:"InstanceSystemDiskSize,omitempty" xml:"InstanceSystemDiskSize,omitempty"`
	// The IP address of the instance in the VPC.
	//
	// example:
	//
	// 192.168.xx.xx
	InstanceTenantIp *string `json:"InstanceTenantIp,omitempty" xml:"InstanceTenantIp,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.s6-c1m2.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of CPU cores used by the instance.
	//
	// example:
	//
	// 2.4
	InstanceUsedCpu *float32 `json:"InstanceUsedCpu,omitempty" xml:"InstanceUsedCpu,omitempty"`
	// The number of GPUs used by the instance.
	//
	// example:
	//
	// 0
	InstanceUsedGpu *float32 `json:"InstanceUsedGpu,omitempty" xml:"InstanceUsedGpu,omitempty"`
	// The size of the GPU memory used by the instance.
	//
	// example:
	//
	// 470M
	InstanceUsedGpuMemory *string `json:"InstanceUsedGpuMemory,omitempty" xml:"InstanceUsedGpuMemory,omitempty"`
	// The size of the memory used by the instance.
	//
	// example:
	//
	// 1000M
	InstanceUsedMemory *string `json:"InstanceUsedMemory,omitempty" xml:"InstanceUsedMemory,omitempty"`
	// The instance tags.
	Labels             []*ResourceInstanceLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LastCordonOperator *string                   `json:"LastCordonOperator,omitempty" xml:"LastCordonOperator,omitempty"`
	LastCordonReason   *string                   `json:"LastCordonReason,omitempty" xml:"LastCordonReason,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// eas-r-xxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the zone to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou-b
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s ResourceInstance) String() string {
	return dara.Prettify(s)
}

func (s ResourceInstance) GoString() string {
	return s.String()
}

func (s *ResourceInstance) GetArch() *string {
	return s.Arch
}

func (s *ResourceInstance) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *ResourceInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *ResourceInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ResourceInstance) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ResourceInstance) GetInstanceCpuCount() *int32 {
	return s.InstanceCpuCount
}

func (s *ResourceInstance) GetInstanceGpuCount() *int32 {
	return s.InstanceGpuCount
}

func (s *ResourceInstance) GetInstanceGpuMemory() *string {
	return s.InstanceGpuMemory
}

func (s *ResourceInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResourceInstance) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *ResourceInstance) GetInstanceMemory() *string {
	return s.InstanceMemory
}

func (s *ResourceInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ResourceInstance) GetInstancePhase() *string {
	return s.InstancePhase
}

func (s *ResourceInstance) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ResourceInstance) GetInstanceSystemDiskSize() *int32 {
	return s.InstanceSystemDiskSize
}

func (s *ResourceInstance) GetInstanceTenantIp() *string {
	return s.InstanceTenantIp
}

func (s *ResourceInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ResourceInstance) GetInstanceUsedCpu() *float32 {
	return s.InstanceUsedCpu
}

func (s *ResourceInstance) GetInstanceUsedGpu() *float32 {
	return s.InstanceUsedGpu
}

func (s *ResourceInstance) GetInstanceUsedGpuMemory() *string {
	return s.InstanceUsedGpuMemory
}

func (s *ResourceInstance) GetInstanceUsedMemory() *string {
	return s.InstanceUsedMemory
}

func (s *ResourceInstance) GetLabels() []*ResourceInstanceLabels {
	return s.Labels
}

func (s *ResourceInstance) GetLastCordonOperator() *string {
	return s.LastCordonOperator
}

func (s *ResourceInstance) GetLastCordonReason() *string {
	return s.LastCordonReason
}

func (s *ResourceInstance) GetRegion() *string {
	return s.Region
}

func (s *ResourceInstance) GetResourceId() *string {
	return s.ResourceId
}

func (s *ResourceInstance) GetZone() *string {
	return s.Zone
}

func (s *ResourceInstance) SetArch(v string) *ResourceInstance {
	s.Arch = &v
	return s
}

func (s *ResourceInstance) SetAutoRenewal(v bool) *ResourceInstance {
	s.AutoRenewal = &v
	return s
}

func (s *ResourceInstance) SetChargeType(v string) *ResourceInstance {
	s.ChargeType = &v
	return s
}

func (s *ResourceInstance) SetCreateTime(v string) *ResourceInstance {
	s.CreateTime = &v
	return s
}

func (s *ResourceInstance) SetExpiredTime(v string) *ResourceInstance {
	s.ExpiredTime = &v
	return s
}

func (s *ResourceInstance) SetInstanceCpuCount(v int32) *ResourceInstance {
	s.InstanceCpuCount = &v
	return s
}

func (s *ResourceInstance) SetInstanceGpuCount(v int32) *ResourceInstance {
	s.InstanceGpuCount = &v
	return s
}

func (s *ResourceInstance) SetInstanceGpuMemory(v string) *ResourceInstance {
	s.InstanceGpuMemory = &v
	return s
}

func (s *ResourceInstance) SetInstanceId(v string) *ResourceInstance {
	s.InstanceId = &v
	return s
}

func (s *ResourceInstance) SetInstanceIp(v string) *ResourceInstance {
	s.InstanceIp = &v
	return s
}

func (s *ResourceInstance) SetInstanceMemory(v string) *ResourceInstance {
	s.InstanceMemory = &v
	return s
}

func (s *ResourceInstance) SetInstanceName(v string) *ResourceInstance {
	s.InstanceName = &v
	return s
}

func (s *ResourceInstance) SetInstancePhase(v string) *ResourceInstance {
	s.InstancePhase = &v
	return s
}

func (s *ResourceInstance) SetInstanceStatus(v string) *ResourceInstance {
	s.InstanceStatus = &v
	return s
}

func (s *ResourceInstance) SetInstanceSystemDiskSize(v int32) *ResourceInstance {
	s.InstanceSystemDiskSize = &v
	return s
}

func (s *ResourceInstance) SetInstanceTenantIp(v string) *ResourceInstance {
	s.InstanceTenantIp = &v
	return s
}

func (s *ResourceInstance) SetInstanceType(v string) *ResourceInstance {
	s.InstanceType = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedCpu(v float32) *ResourceInstance {
	s.InstanceUsedCpu = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedGpu(v float32) *ResourceInstance {
	s.InstanceUsedGpu = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedGpuMemory(v string) *ResourceInstance {
	s.InstanceUsedGpuMemory = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedMemory(v string) *ResourceInstance {
	s.InstanceUsedMemory = &v
	return s
}

func (s *ResourceInstance) SetLabels(v []*ResourceInstanceLabels) *ResourceInstance {
	s.Labels = v
	return s
}

func (s *ResourceInstance) SetLastCordonOperator(v string) *ResourceInstance {
	s.LastCordonOperator = &v
	return s
}

func (s *ResourceInstance) SetLastCordonReason(v string) *ResourceInstance {
	s.LastCordonReason = &v
	return s
}

func (s *ResourceInstance) SetRegion(v string) *ResourceInstance {
	s.Region = &v
	return s
}

func (s *ResourceInstance) SetResourceId(v string) *ResourceInstance {
	s.ResourceId = &v
	return s
}

func (s *ResourceInstance) SetZone(v string) *ResourceInstance {
	s.Zone = &v
	return s
}

func (s *ResourceInstance) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResourceInstanceLabels struct {
	// The tag key of the instance.
	//
	// example:
	//
	// key
	LabelKey *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	// The tag value of the instance.
	//
	// example:
	//
	// value
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s ResourceInstanceLabels) String() string {
	return dara.Prettify(s)
}

func (s ResourceInstanceLabels) GoString() string {
	return s.String()
}

func (s *ResourceInstanceLabels) GetLabelKey() *string {
	return s.LabelKey
}

func (s *ResourceInstanceLabels) GetLabelValue() *string {
	return s.LabelValue
}

func (s *ResourceInstanceLabels) SetLabelKey(v string) *ResourceInstanceLabels {
	s.LabelKey = &v
	return s
}

func (s *ResourceInstanceLabels) SetLabelValue(v string) *ResourceInstanceLabels {
	s.LabelValue = &v
	return s
}

func (s *ResourceInstanceLabels) Validate() error {
	return dara.Validate(s)
}
