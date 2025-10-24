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
	SetRegion(v string) *ResourceInstance
	GetRegion() *string
	SetResourceId(v string) *ResourceInstance
	GetResourceId() *string
	SetZone(v string) *ResourceInstance
	GetZone() *string
}

type ResourceInstance struct {
	Arch                   *string                   `json:"Arch,omitempty" xml:"Arch,omitempty"`
	AutoRenewal            *bool                     `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	ChargeType             *string                   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime             *string                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredTime            *string                   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceCpuCount       *int32                    `json:"InstanceCpuCount,omitempty" xml:"InstanceCpuCount,omitempty"`
	InstanceGpuCount       *int32                    `json:"InstanceGpuCount,omitempty" xml:"InstanceGpuCount,omitempty"`
	InstanceGpuMemory      *string                   `json:"InstanceGpuMemory,omitempty" xml:"InstanceGpuMemory,omitempty"`
	InstanceId             *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIp             *string                   `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	InstanceMemory         *string                   `json:"InstanceMemory,omitempty" xml:"InstanceMemory,omitempty"`
	InstanceName           *string                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus         *string                   `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceSystemDiskSize *int32                    `json:"InstanceSystemDiskSize,omitempty" xml:"InstanceSystemDiskSize,omitempty"`
	InstanceTenantIp       *string                   `json:"InstanceTenantIp,omitempty" xml:"InstanceTenantIp,omitempty"`
	InstanceType           *string                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceUsedCpu        *float32                  `json:"InstanceUsedCpu,omitempty" xml:"InstanceUsedCpu,omitempty"`
	InstanceUsedGpu        *float32                  `json:"InstanceUsedGpu,omitempty" xml:"InstanceUsedGpu,omitempty"`
	InstanceUsedGpuMemory  *string                   `json:"InstanceUsedGpuMemory,omitempty" xml:"InstanceUsedGpuMemory,omitempty"`
	InstanceUsedMemory     *string                   `json:"InstanceUsedMemory,omitempty" xml:"InstanceUsedMemory,omitempty"`
	Labels                 []*ResourceInstanceLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Region                 *string                   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId             *string                   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Zone                   *string                   `json:"Zone,omitempty" xml:"Zone,omitempty"`
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
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
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
