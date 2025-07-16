// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeResourceResponseBody
	GetClusterId() *string
	SetCpuCount(v int32) *DescribeResourceResponseBody
	GetCpuCount() *int32
	SetCpuUsed(v int32) *DescribeResourceResponseBody
	GetCpuUsed() *int32
	SetCreateTime(v string) *DescribeResourceResponseBody
	GetCreateTime() *string
	SetExtraData(v string) *DescribeResourceResponseBody
	GetExtraData() *string
	SetGpuCount(v int32) *DescribeResourceResponseBody
	GetGpuCount() *int32
	SetGpuUsed(v float32) *DescribeResourceResponseBody
	GetGpuUsed() *float32
	SetInstanceCount(v int32) *DescribeResourceResponseBody
	GetInstanceCount() *int32
	SetMemory(v int32) *DescribeResourceResponseBody
	GetMemory() *int32
	SetMemoryUsed(v int32) *DescribeResourceResponseBody
	GetMemoryUsed() *int32
	SetMessage(v string) *DescribeResourceResponseBody
	GetMessage() *string
	SetOwnerUid(v string) *DescribeResourceResponseBody
	GetOwnerUid() *string
	SetPostPaidInstanceCount(v int32) *DescribeResourceResponseBody
	GetPostPaidInstanceCount() *int32
	SetPrePaidInstanceCount(v int32) *DescribeResourceResponseBody
	GetPrePaidInstanceCount() *int32
	SetRequestId(v string) *DescribeResourceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *DescribeResourceResponseBody
	GetResourceId() *string
	SetResourceName(v string) *DescribeResourceResponseBody
	GetResourceName() *string
	SetResourceType(v string) *DescribeResourceResponseBody
	GetResourceType() *string
	SetStatus(v string) *DescribeResourceResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *DescribeResourceResponseBody
	GetUpdateTime() *string
}

type DescribeResourceResponseBody struct {
	// The ID of the cluster to which the resource group belongs.
	//
	// example:
	//
	// cn-beijing
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The total number of CPU cores.
	//
	// example:
	//
	// 16
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The number of vCPUs that is used.
	//
	// example:
	//
	// 8
	CpuUsed *int32 `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	// The time when the resource group was created.
	//
	// example:
	//
	// 2020-05-19T14:19:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The additional information, such as the connection status of a virtual private cloud (VPC) and the log status of Log Service.
	//
	// example:
	//
	// {"vswitch_id":"vsw-bp17uo6xebcusy****","gpu_share":true,"aux_vswitch_id_list":["vsw-bp13b3pvjap3vxn****","vsw-bp1nls8o5hk8mt8*****"],"security_group_id":"sg-bp1j1z7297hcink*****","vpc_id":"vpc-bp1kjr3rfyhx01*****","destination_cidr":"172.16.0.12/28","role_arn":"acs:ram::1157703270*****:role/AliyunServiceRoleForPaiEas","sls_project":"","sls_logstore":"","sls_status":"ResourceReady","sls_message":"","update_time":""}
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// The total number of GPUs.
	//
	// example:
	//
	// 1
	GpuCount *int32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The number of GPUs that is used.
	//
	// example:
	//
	// 2
	GpuUsed *float32 `json:"GpuUsed,omitempty" xml:"GpuUsed,omitempty"`
	// The total number of instances in the resource group.
	//
	// example:
	//
	// 4
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The total memory size. Unit: MB.
	//
	// example:
	//
	// 8192
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The size of memory that is used. Unit: MB.
	//
	// example:
	//
	// 2048
	MemoryUsed *int32 `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Resource is ready
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the resource group owner.
	//
	// example:
	//
	// 14401087478****
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The total number of pay-as-you-go instances in the resource group.
	//
	// example:
	//
	// 3
	PostPaidInstanceCount *int32 `json:"PostPaidInstanceCount,omitempty" xml:"PostPaidInstanceCount,omitempty"`
	// The total number of subscription instances in the resource group.
	//
	// example:
	//
	// 1
	PrePaidInstanceCount *int32 `json:"PrePaidInstanceCount,omitempty" xml:"PrePaidInstanceCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 902976F2-6FAF-5404-8A4D-6CC223***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the Elastic Algorithm Service (EAS) resource.
	//
	// example:
	//
	// eas-r-glkfpsxuw57x1h*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the EAS resource.
	//
	// example:
	//
	// my-resouce****
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- Dedicated: the dedicated resource group.
	//
	// 	- SelfManaged: the self-managed resource group.
	//
	// example:
	//
	// Dedicated
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The state of the resource group.
	//
	// example:
	//
	// ResourceReady
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the resource group was last updated.
	//
	// example:
	//
	// 2021-02-24T11:52:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeResourceResponseBody) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *DescribeResourceResponseBody) GetCpuUsed() *int32 {
	return s.CpuUsed
}

func (s *DescribeResourceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeResourceResponseBody) GetExtraData() *string {
	return s.ExtraData
}

func (s *DescribeResourceResponseBody) GetGpuCount() *int32 {
	return s.GpuCount
}

func (s *DescribeResourceResponseBody) GetGpuUsed() *float32 {
	return s.GpuUsed
}

func (s *DescribeResourceResponseBody) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeResourceResponseBody) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeResourceResponseBody) GetMemoryUsed() *int32 {
	return s.MemoryUsed
}

func (s *DescribeResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResourceResponseBody) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *DescribeResourceResponseBody) GetPostPaidInstanceCount() *int32 {
	return s.PostPaidInstanceCount
}

func (s *DescribeResourceResponseBody) GetPrePaidInstanceCount() *int32 {
	return s.PrePaidInstanceCount
}

func (s *DescribeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourceResponseBody) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeResourceResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourceResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeResourceResponseBody) SetClusterId(v string) *DescribeResourceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetCpuCount(v int32) *DescribeResourceResponseBody {
	s.CpuCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetCpuUsed(v int32) *DescribeResourceResponseBody {
	s.CpuUsed = &v
	return s
}

func (s *DescribeResourceResponseBody) SetCreateTime(v string) *DescribeResourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourceResponseBody) SetExtraData(v string) *DescribeResourceResponseBody {
	s.ExtraData = &v
	return s
}

func (s *DescribeResourceResponseBody) SetGpuCount(v int32) *DescribeResourceResponseBody {
	s.GpuCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetGpuUsed(v float32) *DescribeResourceResponseBody {
	s.GpuUsed = &v
	return s
}

func (s *DescribeResourceResponseBody) SetInstanceCount(v int32) *DescribeResourceResponseBody {
	s.InstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetMemory(v int32) *DescribeResourceResponseBody {
	s.Memory = &v
	return s
}

func (s *DescribeResourceResponseBody) SetMemoryUsed(v int32) *DescribeResourceResponseBody {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeResourceResponseBody) SetMessage(v string) *DescribeResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceResponseBody) SetOwnerUid(v string) *DescribeResourceResponseBody {
	s.OwnerUid = &v
	return s
}

func (s *DescribeResourceResponseBody) SetPostPaidInstanceCount(v int32) *DescribeResourceResponseBody {
	s.PostPaidInstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetPrePaidInstanceCount(v int32) *DescribeResourceResponseBody {
	s.PrePaidInstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetRequestId(v string) *DescribeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetResourceId(v string) *DescribeResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetResourceName(v string) *DescribeResourceResponseBody {
	s.ResourceName = &v
	return s
}

func (s *DescribeResourceResponseBody) SetResourceType(v string) *DescribeResourceResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceResponseBody) SetStatus(v string) *DescribeResourceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeResourceResponseBody) SetUpdateTime(v string) *DescribeResourceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
