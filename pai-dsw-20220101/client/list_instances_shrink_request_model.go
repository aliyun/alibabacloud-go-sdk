// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *ListInstancesShrinkRequest
	GetAcceleratorType() *string
	SetAccessibility(v string) *ListInstancesShrinkRequest
	GetAccessibility() *string
	SetCreateUserId(v string) *ListInstancesShrinkRequest
	GetCreateUserId() *string
	SetGpuType(v string) *ListInstancesShrinkRequest
	GetGpuType() *string
	SetImageName(v string) *ListInstancesShrinkRequest
	GetImageName() *string
	SetInstanceId(v string) *ListInstancesShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListInstancesShrinkRequest
	GetInstanceName() *string
	SetLabelsShrink(v string) *ListInstancesShrinkRequest
	GetLabelsShrink() *string
	SetMaxCpu(v string) *ListInstancesShrinkRequest
	GetMaxCpu() *string
	SetMaxGpu(v string) *ListInstancesShrinkRequest
	GetMaxGpu() *string
	SetMaxGpuMemory(v string) *ListInstancesShrinkRequest
	GetMaxGpuMemory() *string
	SetMaxMemory(v string) *ListInstancesShrinkRequest
	GetMaxMemory() *string
	SetMinCpu(v string) *ListInstancesShrinkRequest
	GetMinCpu() *string
	SetMinGpu(v string) *ListInstancesShrinkRequest
	GetMinGpu() *string
	SetMinGpuMemory(v string) *ListInstancesShrinkRequest
	GetMinGpuMemory() *string
	SetMinMemory(v string) *ListInstancesShrinkRequest
	GetMinMemory() *string
	SetOrder(v string) *ListInstancesShrinkRequest
	GetOrder() *string
	SetOversoldInfo(v string) *ListInstancesShrinkRequest
	GetOversoldInfo() *string
	SetOversoldType(v string) *ListInstancesShrinkRequest
	GetOversoldType() *string
	SetPageNumber(v int64) *ListInstancesShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListInstancesShrinkRequest
	GetPageSize() *int64
	SetPaymentType(v string) *ListInstancesShrinkRequest
	GetPaymentType() *string
	SetResourceId(v string) *ListInstancesShrinkRequest
	GetResourceId() *string
	SetSortBy(v string) *ListInstancesShrinkRequest
	GetSortBy() *string
	SetStatus(v string) *ListInstancesShrinkRequest
	GetStatus() *string
	SetTagShrink(v string) *ListInstancesShrinkRequest
	GetTagShrink() *string
	SetWorkspaceId(v string) *ListInstancesShrinkRequest
	GetWorkspaceId() *string
}

type ListInstancesShrinkRequest struct {
	// The accelerator type.
	//
	// 	- CPU: Only CPU computing is used.
	//
	// 	- GPU: GPUs are used to accelerate computing.
	//
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The accessibility. Valid values:
	//
	// 	- PRIVATE (default): The instances are accessible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: The instances are accessible only to all members in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The UID of the creator.
	//
	// example:
	//
	// 12345*****67890
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// The GPU type.
	//
	// example:
	//
	// NVIDIA A10
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// The image name.
	//
	// example:
	//
	// modelscope:1.9.4-pytorch2.0.1tensorflow2.13.0-cpu-py38-ubuntu20.04
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The instance ID. You can call [ListInstances](https://help.aliyun.com/document_detail/470439.html) to obtain the instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// training_data
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The labels. A maximum of four labels are supported.
	//
	// example:
	//
	// {
	//
	//   "key1": "value1",
	//
	//   "key2": "value2",
	//
	//   "key3": "value3"
	//
	// }
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of CPUs. Unit: 0.001 CPU. The value 1000 indicates one CPU.
	//
	// example:
	//
	// 30000
	MaxCpu *string `json:"MaxCpu,omitempty" xml:"MaxCpu,omitempty"`
	// The maximum number of GPUs. Unit: 0.001 GPU. The value 1000 indicates one GPU.
	//
	// example:
	//
	// 8000
	MaxGpu *string `json:"MaxGpu,omitempty" xml:"MaxGpu,omitempty"`
	// The maximum memory size per GPU card. Unit: GB.
	//
	// example:
	//
	// 16
	MaxGpuMemory *string `json:"MaxGpuMemory,omitempty" xml:"MaxGpuMemory,omitempty"`
	// The maximum memory size. Unit: GB.
	//
	// example:
	//
	// 48
	MaxMemory *string `json:"MaxMemory,omitempty" xml:"MaxMemory,omitempty"`
	// The minimum number of CPUs. Unit: 0.001 CPU. The value 1000 indicates one CPU.
	//
	// example:
	//
	// 2000
	MinCpu *string `json:"MinCpu,omitempty" xml:"MinCpu,omitempty"`
	// The minimum number of GPUs. Unit: 0.001 GPU. The value 1000 indicates one GPU.
	//
	// example:
	//
	// 100
	MinGpu *string `json:"MinGpu,omitempty" xml:"MinGpu,omitempty"`
	// The minimum memory size per GPU card. Unit: GB.
	//
	// example:
	//
	// 8
	MinGpuMemory *string `json:"MinGpuMemory,omitempty" xml:"MinGpuMemory,omitempty"`
	// The minimum memory size. Unit: GB.
	//
	// example:
	//
	// 4
	MinMemory *string `json:"MinMemory,omitempty" xml:"MinMemory,omitempty"`
	// The order that you use to sort the query results.
	//
	// Valid values:
	//
	// 	- ASC
	//
	// 	- DESC
	//
	// example:
	//
	// DESC
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OversoldInfo *string `json:"OversoldInfo,omitempty" xml:"OversoldInfo,omitempty"`
	OversoldType *string `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Subscription
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The resource group ID. If you leave this parameter empty, the instances in the pay-as-you-go resource group are queried. If you set this parameter to ALL, all instances are queried.
	//
	// example:
	//
	// rg-123456789
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The field that you use to sort the query results.
	//
	// Valid values:
	//
	// 	- Priority
	//
	// 	- GmtCreateTime
	//
	// 	- GmtModifiedTime
	//
	// example:
	//
	// gmtCreate
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The instance status.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- SaveFailed
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- ResourceAllocating
	//
	// 	- Stopping
	//
	// 	- Updating
	//
	// 	- Saving
	//
	// 	- Queuing
	//
	// 	- Recovering
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Saved
	//
	// 	- Deleting
	//
	// 	- EnvPreparing
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 40823
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListInstancesShrinkRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListInstancesShrinkRequest) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListInstancesShrinkRequest) GetGpuType() *string {
	return s.GpuType
}

func (s *ListInstancesShrinkRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ListInstancesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *ListInstancesShrinkRequest) GetMaxCpu() *string {
	return s.MaxCpu
}

func (s *ListInstancesShrinkRequest) GetMaxGpu() *string {
	return s.MaxGpu
}

func (s *ListInstancesShrinkRequest) GetMaxGpuMemory() *string {
	return s.MaxGpuMemory
}

func (s *ListInstancesShrinkRequest) GetMaxMemory() *string {
	return s.MaxMemory
}

func (s *ListInstancesShrinkRequest) GetMinCpu() *string {
	return s.MinCpu
}

func (s *ListInstancesShrinkRequest) GetMinGpu() *string {
	return s.MinGpu
}

func (s *ListInstancesShrinkRequest) GetMinGpuMemory() *string {
	return s.MinGpuMemory
}

func (s *ListInstancesShrinkRequest) GetMinMemory() *string {
	return s.MinMemory
}

func (s *ListInstancesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListInstancesShrinkRequest) GetOversoldInfo() *string {
	return s.OversoldInfo
}

func (s *ListInstancesShrinkRequest) GetOversoldType() *string {
	return s.OversoldType
}

func (s *ListInstancesShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstancesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstancesShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstancesShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListInstancesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListInstancesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListInstancesShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListInstancesShrinkRequest) SetAcceleratorType(v string) *ListInstancesShrinkRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetAccessibility(v string) *ListInstancesShrinkRequest {
	s.Accessibility = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetCreateUserId(v string) *ListInstancesShrinkRequest {
	s.CreateUserId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetGpuType(v string) *ListInstancesShrinkRequest {
	s.GpuType = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetImageName(v string) *ListInstancesShrinkRequest {
	s.ImageName = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceId(v string) *ListInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceName(v string) *ListInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetLabelsShrink(v string) *ListInstancesShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMaxCpu(v string) *ListInstancesShrinkRequest {
	s.MaxCpu = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMaxGpu(v string) *ListInstancesShrinkRequest {
	s.MaxGpu = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMaxGpuMemory(v string) *ListInstancesShrinkRequest {
	s.MaxGpuMemory = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMaxMemory(v string) *ListInstancesShrinkRequest {
	s.MaxMemory = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMinCpu(v string) *ListInstancesShrinkRequest {
	s.MinCpu = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMinGpu(v string) *ListInstancesShrinkRequest {
	s.MinGpu = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMinGpuMemory(v string) *ListInstancesShrinkRequest {
	s.MinGpuMemory = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMinMemory(v string) *ListInstancesShrinkRequest {
	s.MinMemory = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetOrder(v string) *ListInstancesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetOversoldInfo(v string) *ListInstancesShrinkRequest {
	s.OversoldInfo = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetOversoldType(v string) *ListInstancesShrinkRequest {
	s.OversoldType = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageNumber(v int64) *ListInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageSize(v int64) *ListInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPaymentType(v string) *ListInstancesShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceId(v string) *ListInstancesShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetSortBy(v string) *ListInstancesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetStatus(v string) *ListInstancesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTagShrink(v string) *ListInstancesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetWorkspaceId(v string) *ListInstancesShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
