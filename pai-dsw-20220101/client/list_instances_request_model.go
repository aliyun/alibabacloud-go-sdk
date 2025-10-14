// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *ListInstancesRequest
	GetAcceleratorType() *string
	SetAccessibility(v string) *ListInstancesRequest
	GetAccessibility() *string
	SetCreateUserId(v string) *ListInstancesRequest
	GetCreateUserId() *string
	SetGpuType(v string) *ListInstancesRequest
	GetGpuType() *string
	SetImageName(v string) *ListInstancesRequest
	GetImageName() *string
	SetInstanceId(v string) *ListInstancesRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListInstancesRequest
	GetInstanceName() *string
	SetLabels(v map[string]interface{}) *ListInstancesRequest
	GetLabels() map[string]interface{}
	SetMaxCpu(v string) *ListInstancesRequest
	GetMaxCpu() *string
	SetMaxGpu(v string) *ListInstancesRequest
	GetMaxGpu() *string
	SetMaxGpuMemory(v string) *ListInstancesRequest
	GetMaxGpuMemory() *string
	SetMaxMemory(v string) *ListInstancesRequest
	GetMaxMemory() *string
	SetMinCpu(v string) *ListInstancesRequest
	GetMinCpu() *string
	SetMinGpu(v string) *ListInstancesRequest
	GetMinGpu() *string
	SetMinGpuMemory(v string) *ListInstancesRequest
	GetMinGpuMemory() *string
	SetMinMemory(v string) *ListInstancesRequest
	GetMinMemory() *string
	SetOrder(v string) *ListInstancesRequest
	GetOrder() *string
	SetOversoldInfo(v string) *ListInstancesRequest
	GetOversoldInfo() *string
	SetOversoldType(v string) *ListInstancesRequest
	GetOversoldType() *string
	SetPageNumber(v int64) *ListInstancesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListInstancesRequest
	GetPageSize() *int64
	SetPaymentType(v string) *ListInstancesRequest
	GetPaymentType() *string
	SetResourceId(v string) *ListInstancesRequest
	GetResourceId() *string
	SetSortBy(v string) *ListInstancesRequest
	GetSortBy() *string
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
	SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest
	GetTag() []*ListInstancesRequestTag
	SetWorkspaceId(v string) *ListInstancesRequest
	GetWorkspaceId() *string
}

type ListInstancesRequest struct {
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
	Labels map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
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
	Tag []*ListInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 40823
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListInstancesRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListInstancesRequest) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListInstancesRequest) GetGpuType() *string {
	return s.GpuType
}

func (s *ListInstancesRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ListInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesRequest) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *ListInstancesRequest) GetMaxCpu() *string {
	return s.MaxCpu
}

func (s *ListInstancesRequest) GetMaxGpu() *string {
	return s.MaxGpu
}

func (s *ListInstancesRequest) GetMaxGpuMemory() *string {
	return s.MaxGpuMemory
}

func (s *ListInstancesRequest) GetMaxMemory() *string {
	return s.MaxMemory
}

func (s *ListInstancesRequest) GetMinCpu() *string {
	return s.MinCpu
}

func (s *ListInstancesRequest) GetMinGpu() *string {
	return s.MinGpu
}

func (s *ListInstancesRequest) GetMinGpuMemory() *string {
	return s.MinGpuMemory
}

func (s *ListInstancesRequest) GetMinMemory() *string {
	return s.MinMemory
}

func (s *ListInstancesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListInstancesRequest) GetOversoldInfo() *string {
	return s.OversoldInfo
}

func (s *ListInstancesRequest) GetOversoldType() *string {
	return s.OversoldType
}

func (s *ListInstancesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstancesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) GetTag() []*ListInstancesRequestTag {
	return s.Tag
}

func (s *ListInstancesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListInstancesRequest) SetAcceleratorType(v string) *ListInstancesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListInstancesRequest) SetAccessibility(v string) *ListInstancesRequest {
	s.Accessibility = &v
	return s
}

func (s *ListInstancesRequest) SetCreateUserId(v string) *ListInstancesRequest {
	s.CreateUserId = &v
	return s
}

func (s *ListInstancesRequest) SetGpuType(v string) *ListInstancesRequest {
	s.GpuType = &v
	return s
}

func (s *ListInstancesRequest) SetImageName(v string) *ListInstancesRequest {
	s.ImageName = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceId(v string) *ListInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesRequest) SetLabels(v map[string]interface{}) *ListInstancesRequest {
	s.Labels = v
	return s
}

func (s *ListInstancesRequest) SetMaxCpu(v string) *ListInstancesRequest {
	s.MaxCpu = &v
	return s
}

func (s *ListInstancesRequest) SetMaxGpu(v string) *ListInstancesRequest {
	s.MaxGpu = &v
	return s
}

func (s *ListInstancesRequest) SetMaxGpuMemory(v string) *ListInstancesRequest {
	s.MaxGpuMemory = &v
	return s
}

func (s *ListInstancesRequest) SetMaxMemory(v string) *ListInstancesRequest {
	s.MaxMemory = &v
	return s
}

func (s *ListInstancesRequest) SetMinCpu(v string) *ListInstancesRequest {
	s.MinCpu = &v
	return s
}

func (s *ListInstancesRequest) SetMinGpu(v string) *ListInstancesRequest {
	s.MinGpu = &v
	return s
}

func (s *ListInstancesRequest) SetMinGpuMemory(v string) *ListInstancesRequest {
	s.MinGpuMemory = &v
	return s
}

func (s *ListInstancesRequest) SetMinMemory(v string) *ListInstancesRequest {
	s.MinMemory = &v
	return s
}

func (s *ListInstancesRequest) SetOrder(v string) *ListInstancesRequest {
	s.Order = &v
	return s
}

func (s *ListInstancesRequest) SetOversoldInfo(v string) *ListInstancesRequest {
	s.OversoldInfo = &v
	return s
}

func (s *ListInstancesRequest) SetOversoldType(v string) *ListInstancesRequest {
	s.OversoldType = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int64) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int64) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetPaymentType(v string) *ListInstancesRequest {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesRequest) SetResourceId(v string) *ListInstancesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListInstancesRequest) SetSortBy(v string) *ListInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

func (s *ListInstancesRequest) SetWorkspaceId(v string) *ListInstancesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
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

type ListInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *ListInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
