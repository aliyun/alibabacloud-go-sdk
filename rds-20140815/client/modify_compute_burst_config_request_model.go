// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComputeBurstConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBurstStatus(v string) *ModifyComputeBurstConfigRequest
	GetBurstStatus() *string
	SetClientToken(v string) *ModifyComputeBurstConfigRequest
	GetClientToken() *string
	SetCpuEnlargeThreshold(v string) *ModifyComputeBurstConfigRequest
	GetCpuEnlargeThreshold() *string
	SetCpuShrinkThreshold(v string) *ModifyComputeBurstConfigRequest
	GetCpuShrinkThreshold() *string
	SetCrontabJobId(v string) *ModifyComputeBurstConfigRequest
	GetCrontabJobId() *string
	SetDBInstanceId(v string) *ModifyComputeBurstConfigRequest
	GetDBInstanceId() *string
	SetMemoryEnlargeThreshold(v string) *ModifyComputeBurstConfigRequest
	GetMemoryEnlargeThreshold() *string
	SetMemoryShrinkThreshold(v string) *ModifyComputeBurstConfigRequest
	GetMemoryShrinkThreshold() *string
	SetOwnerAccount(v string) *ModifyComputeBurstConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyComputeBurstConfigRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyComputeBurstConfigRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyComputeBurstConfigRequest
	GetResourceOwnerAccount() *string
	SetScaleMaxCpus(v string) *ModifyComputeBurstConfigRequest
	GetScaleMaxCpus() *string
	SetScaleMaxMemory(v string) *ModifyComputeBurstConfigRequest
	GetScaleMaxMemory() *string
	SetSwitchTime(v string) *ModifyComputeBurstConfigRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *ModifyComputeBurstConfigRequest
	GetSwitchTimeMode() *string
	SetTaskId(v string) *ModifyComputeBurstConfigRequest
	GetTaskId() *string
}

type ModifyComputeBurstConfigRequest struct {
	// This parameter is set to **disabled*	- if the assured serverless feature is disabled.
	//
	// example:
	//
	// disabled
	BurstStatus *string `json:"BurstStatus,omitempty" xml:"BurstStatus,omitempty"`
	// The client token that is used to ensure the idempotence of requests and prevent repeated requests from being submitted. You can use the client to generate the value, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCziJZNwH****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The CPU utilization threshold for **scale-out**. Valid values: 60 to 90. Unit: %.
	//
	// example:
	//
	// 80
	CpuEnlargeThreshold *string `json:"CpuEnlargeThreshold,omitempty" xml:"CpuEnlargeThreshold,omitempty"`
	// The CPU utilization threshold for **scale-in**. Valid values: 30 to 55. Unit: %.
	//
	// example:
	//
	// 50
	CpuShrinkThreshold *string `json:"CpuShrinkThreshold,omitempty" xml:"CpuShrinkThreshold,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	CrontabJobId *string `json:"CrontabJobId,omitempty" xml:"CrontabJobId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-2ze63v2p3o3k****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The memory usage threshold for **scale-out**. Valid values: 60 to 90. Unit: %.
	//
	// example:
	//
	// 80
	MemoryEnlargeThreshold *string `json:"MemoryEnlargeThreshold,omitempty" xml:"MemoryEnlargeThreshold,omitempty"`
	// The memory usage threshold for **scale-in**. Valid values: 30 to 55. Unit: %.
	//
	// example:
	//
	// 50
	MemoryShrinkThreshold *string `json:"MemoryShrinkThreshold,omitempty" xml:"MemoryShrinkThreshold,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The maximum number of CPU cores for elastic scaling. The maximum value cannot exceed twice the initial CPU configuration.
	//
	// example:
	//
	// 2
	ScaleMaxCpus *string `json:"ScaleMaxCpus,omitempty" xml:"ScaleMaxCpus,omitempty"`
	// The maximum memory for elastic scaling. The value cannot exceed twice the instance\\"s initial memory size. Unit: GB. Step size: 2 GB.
	//
	// example:
	//
	// 4
	ScaleMaxMemory *string `json:"ScaleMaxMemory,omitempty" xml:"ScaleMaxMemory,omitempty"`
	// The time when the specified entry takes effect. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// >  This parameter is required only if **SwitchTimeMode*	- is set to **2**.
	//
	// example:
	//
	// 2025-05-06T09:24:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The effective policy. Valid values:
	//
	// 	- **0**: Immediately takes effect.
	//
	// 	- **1**: Takes effect within the maintenance window. You can call the **ModifyDBInstanceMaintainTime*	- operation to change the maintenance window of an instance.
	//
	// 	- **2**: Takes effect at a specified point in time.
	//
	// example:
	//
	// Immediate
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyComputeBurstConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyComputeBurstConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyComputeBurstConfigRequest) GetBurstStatus() *string {
	return s.BurstStatus
}

func (s *ModifyComputeBurstConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyComputeBurstConfigRequest) GetCpuEnlargeThreshold() *string {
	return s.CpuEnlargeThreshold
}

func (s *ModifyComputeBurstConfigRequest) GetCpuShrinkThreshold() *string {
	return s.CpuShrinkThreshold
}

func (s *ModifyComputeBurstConfigRequest) GetCrontabJobId() *string {
	return s.CrontabJobId
}

func (s *ModifyComputeBurstConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyComputeBurstConfigRequest) GetMemoryEnlargeThreshold() *string {
	return s.MemoryEnlargeThreshold
}

func (s *ModifyComputeBurstConfigRequest) GetMemoryShrinkThreshold() *string {
	return s.MemoryShrinkThreshold
}

func (s *ModifyComputeBurstConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyComputeBurstConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyComputeBurstConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyComputeBurstConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyComputeBurstConfigRequest) GetScaleMaxCpus() *string {
	return s.ScaleMaxCpus
}

func (s *ModifyComputeBurstConfigRequest) GetScaleMaxMemory() *string {
	return s.ScaleMaxMemory
}

func (s *ModifyComputeBurstConfigRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyComputeBurstConfigRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *ModifyComputeBurstConfigRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyComputeBurstConfigRequest) SetBurstStatus(v string) *ModifyComputeBurstConfigRequest {
	s.BurstStatus = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetClientToken(v string) *ModifyComputeBurstConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetCpuEnlargeThreshold(v string) *ModifyComputeBurstConfigRequest {
	s.CpuEnlargeThreshold = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetCpuShrinkThreshold(v string) *ModifyComputeBurstConfigRequest {
	s.CpuShrinkThreshold = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetCrontabJobId(v string) *ModifyComputeBurstConfigRequest {
	s.CrontabJobId = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetDBInstanceId(v string) *ModifyComputeBurstConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetMemoryEnlargeThreshold(v string) *ModifyComputeBurstConfigRequest {
	s.MemoryEnlargeThreshold = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetMemoryShrinkThreshold(v string) *ModifyComputeBurstConfigRequest {
	s.MemoryShrinkThreshold = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetOwnerAccount(v string) *ModifyComputeBurstConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetOwnerId(v int64) *ModifyComputeBurstConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetResourceGroupId(v string) *ModifyComputeBurstConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetResourceOwnerAccount(v string) *ModifyComputeBurstConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetScaleMaxCpus(v string) *ModifyComputeBurstConfigRequest {
	s.ScaleMaxCpus = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetScaleMaxMemory(v string) *ModifyComputeBurstConfigRequest {
	s.ScaleMaxMemory = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetSwitchTime(v string) *ModifyComputeBurstConfigRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetSwitchTimeMode(v string) *ModifyComputeBurstConfigRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) SetTaskId(v string) *ModifyComputeBurstConfigRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyComputeBurstConfigRequest) Validate() error {
	return dara.Validate(s)
}
