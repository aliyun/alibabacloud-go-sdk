// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWorkspaceQueueRequest
	GetDescription() *string
	SetGpuSpec(v []*string) *CreateWorkspaceQueueRequest
	GetGpuSpec() []*string
	SetInstanceId(v string) *CreateWorkspaceQueueRequest
	GetInstanceId() *string
	SetPaymentType(v string) *CreateWorkspaceQueueRequest
	GetPaymentType() *string
	SetPreheat(v bool) *CreateWorkspaceQueueRequest
	GetPreheat() *bool
	SetQueueCategory(v string) *CreateWorkspaceQueueRequest
	GetQueueCategory() *string
	SetResourceSpec(v *CreateWorkspaceQueueRequestResourceSpec) *CreateWorkspaceQueueRequest
	GetResourceSpec() *CreateWorkspaceQueueRequestResourceSpec
	SetWorkspaceId(v string) *CreateWorkspaceQueueRequest
	GetWorkspaceId() *string
	SetWorkspaceQueueName(v string) *CreateWorkspaceQueueRequest
	GetWorkspaceQueueName() *string
	SetRegionId(v string) *CreateWorkspaceQueueRequest
	GetRegionId() *string
}

type CreateWorkspaceQueueRequest struct {
	// The description.
	//
	// example:
	//
	// Ray Cluster for dev.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of GPU models.
	GpuSpec []*string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty" type:"Repeated"`
	// The Ray cluster instance ID.
	//
	// example:
	//
	// ray-k7nm8ahl5te4tg91-ey7blpbg
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - PayAsYouGo: pay-as-you-go
	//
	// - Pre: subscription
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// Indicates whether resource prefetch is enabled.
	Preheat *bool `json:"preheat,omitempty" xml:"preheat,omitempty"`
	// The queue type. Valid values: CPU and GPU.
	//
	// example:
	//
	// CPU
	QueueCategory *string `json:"queueCategory,omitempty" xml:"queueCategory,omitempty"`
	// The resource specifications.
	ResourceSpec *CreateWorkspaceQueueRequestResourceSpec `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// w-975bcfda9625****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The workspace queue name.
	//
	// example:
	//
	// dev_queue
	WorkspaceQueueName *string `json:"workspaceQueueName,omitempty" xml:"workspaceQueueName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateWorkspaceQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceQueueRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkspaceQueueRequest) GetGpuSpec() []*string {
	return s.GpuSpec
}

func (s *CreateWorkspaceQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWorkspaceQueueRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateWorkspaceQueueRequest) GetPreheat() *bool {
	return s.Preheat
}

func (s *CreateWorkspaceQueueRequest) GetQueueCategory() *string {
	return s.QueueCategory
}

func (s *CreateWorkspaceQueueRequest) GetResourceSpec() *CreateWorkspaceQueueRequestResourceSpec {
	return s.ResourceSpec
}

func (s *CreateWorkspaceQueueRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceQueueRequest) GetWorkspaceQueueName() *string {
	return s.WorkspaceQueueName
}

func (s *CreateWorkspaceQueueRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWorkspaceQueueRequest) SetDescription(v string) *CreateWorkspaceQueueRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetGpuSpec(v []*string) *CreateWorkspaceQueueRequest {
	s.GpuSpec = v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetInstanceId(v string) *CreateWorkspaceQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetPaymentType(v string) *CreateWorkspaceQueueRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetPreheat(v bool) *CreateWorkspaceQueueRequest {
	s.Preheat = &v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetQueueCategory(v string) *CreateWorkspaceQueueRequest {
	s.QueueCategory = &v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetResourceSpec(v *CreateWorkspaceQueueRequestResourceSpec) *CreateWorkspaceQueueRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetWorkspaceId(v string) *CreateWorkspaceQueueRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetWorkspaceQueueName(v string) *CreateWorkspaceQueueRequest {
	s.WorkspaceQueueName = &v
	return s
}

func (s *CreateWorkspaceQueueRequest) SetRegionId(v string) *CreateWorkspaceQueueRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWorkspaceQueueRequest) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceQueueRequestResourceSpec struct {
	// The maximum workspace resource capacity.
	//
	// example:
	//
	// 1000
	Cu *int64 `json:"cu,omitempty" xml:"cu,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 100
	Gpu *int32 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	// The number of GPU machines. This parameter is valid only for subscription instances.
	//
	// example:
	//
	// 8
	GpuMachineNum *int32 `json:"gpuMachineNum,omitempty" xml:"gpuMachineNum,omitempty"`
	// The maximum number of CUs.
	//
	// example:
	//
	// 0.5
	MaxCu *int64 `json:"maxCu,omitempty" xml:"maxCu,omitempty"`
}

func (s CreateWorkspaceQueueRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceQueueRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceQueueRequestResourceSpec) GetCu() *int64 {
	return s.Cu
}

func (s *CreateWorkspaceQueueRequestResourceSpec) GetGpu() *int32 {
	return s.Gpu
}

func (s *CreateWorkspaceQueueRequestResourceSpec) GetGpuMachineNum() *int32 {
	return s.GpuMachineNum
}

func (s *CreateWorkspaceQueueRequestResourceSpec) GetMaxCu() *int64 {
	return s.MaxCu
}

func (s *CreateWorkspaceQueueRequestResourceSpec) SetCu(v int64) *CreateWorkspaceQueueRequestResourceSpec {
	s.Cu = &v
	return s
}

func (s *CreateWorkspaceQueueRequestResourceSpec) SetGpu(v int32) *CreateWorkspaceQueueRequestResourceSpec {
	s.Gpu = &v
	return s
}

func (s *CreateWorkspaceQueueRequestResourceSpec) SetGpuMachineNum(v int32) *CreateWorkspaceQueueRequestResourceSpec {
	s.GpuMachineNum = &v
	return s
}

func (s *CreateWorkspaceQueueRequestResourceSpec) SetMaxCu(v int64) *CreateWorkspaceQueueRequestResourceSpec {
	s.MaxCu = &v
	return s
}

func (s *CreateWorkspaceQueueRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
