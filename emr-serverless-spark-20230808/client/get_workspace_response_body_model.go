// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspace(v *GetWorkspaceResponseBodyWorkspace) *GetWorkspaceResponseBody
	GetWorkspace() *GetWorkspaceResponseBodyWorkspace
}

type GetWorkspaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The workspace object.
	Workspace *GetWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s GetWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceResponseBody) GetWorkspace() *GetWorkspaceResponseBodyWorkspace {
	return s.Workspace
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspace(v *GetWorkspaceResponseBodyWorkspace) *GetWorkspaceResponseBody {
	s.Workspace = v
	return s
}

func (s *GetWorkspaceResponseBody) Validate() error {
	if s.Workspace != nil {
		if err := s.Workspace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceResponseBodyWorkspace struct {
	// Specifies whether auto-renewal is enabled. This parameter is required for subscription workspaces.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The auto-renewal period. This parameter is required for subscription workspaces.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The auto-renewal cycle. This parameter is required for subscription workspaces.
	//
	// example:
	//
	// YEAR, MONTH, WEEK, DAY, HOUR, MINUTE
	AutoRenewPeriodUnit *string `json:"autoRenewPeriodUnit,omitempty" xml:"autoRenewPeriodUnit,omitempty"`
	// The time when the workspace was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1776824869000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The DLF catalog information.
	//
	// example:
	//
	// default
	DlfCatalogId *string `json:"dlfCatalogId,omitempty" xml:"dlfCatalogId,omitempty"`
	// The bound DLF type.
	//
	// example:
	//
	// DLF-NEXT
	DlfType *string `json:"dlfType,omitempty" xml:"dlfType,omitempty"`
	// The subscription duration. This parameter is required for subscription workspaces.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The time when the workspace is released.
	//
	// example:
	//
	// 1786586477000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// out of stock
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// The list of GPU instance types.
	GpuSpec []*string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty" type:"Repeated"`
	// The IP whitelist of the workspace.
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
	// The unit of the subscription period. This parameter is required for subscription workspaces.
	//
	// example:
	//
	// YEAR, MONTH, WEEK, DAY, HOUR, MINUTE
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	// The payment status.
	//
	// example:
	//
	// PAID/UNPAID
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PayAsYouGo or Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The information about the subscription resource quota.
	PrePaidQuota *GetWorkspaceResponseBodyWorkspacePrePaidQuota `json:"prePaidQuota,omitempty" xml:"prePaidQuota,omitempty" type:"Struct"`
	// The list of GPU subscription quotas.
	PrePaidQuotaGpu []*GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu `json:"prePaidQuotaGpu,omitempty" xml:"prePaidQuotaGpu,omitempty" type:"Repeated"`
	// The name of the RAM role used to execute Spark jobs.
	//
	// example:
	//
	// Aliyun
	RamRoleName *string `json:"ramRoleName,omitempty" xml:"ramRoleName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The reason why the workspace is released.
	//
	// example:
	//
	// stable
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek4f7uj74wdthi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The UID of the user who created the workspace.
	//
	// example:
	//
	// 13*************
	ResourceOwnerId *string `json:"resourceOwnerId,omitempty" xml:"resourceOwnerId,omitempty"`
	// The resource specifications.
	//
	// example:
	//
	// {"cu": 100}
	ResourceSpec *string `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	// The information about the workspace status change.
	StateChangeReason *GetWorkspaceResponseBodyWorkspaceStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The OSS path.
	//
	// example:
	//
	// spark-result
	Storage *string `json:"storage,omitempty" xml:"storage,omitempty"`
	// The tag keys.
	Tags []*GetWorkspaceResponseBodyWorkspaceTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// w-******
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// Spark batch job workspace-1
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// The workspace status.
	//
	// example:
	//
	// STARTING,RUNNING,TERMINATED
	WorkspaceStatus *string `json:"workspaceStatus,omitempty" xml:"workspaceStatus,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspace) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspace) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *GetWorkspaceResponseBodyWorkspace) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *GetWorkspaceResponseBodyWorkspace) GetAutoRenewPeriodUnit() *string {
	return s.AutoRenewPeriodUnit
}

func (s *GetWorkspaceResponseBodyWorkspace) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBodyWorkspace) GetDlfCatalogId() *string {
	return s.DlfCatalogId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetDlfType() *string {
	return s.DlfType
}

func (s *GetWorkspaceResponseBodyWorkspace) GetDuration() *int32 {
	return s.Duration
}

func (s *GetWorkspaceResponseBodyWorkspace) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetWorkspaceResponseBodyWorkspace) GetFailReason() *string {
	return s.FailReason
}

func (s *GetWorkspaceResponseBodyWorkspace) GetGpuSpec() []*string {
	return s.GpuSpec
}

func (s *GetWorkspaceResponseBodyWorkspace) GetIpWhiteList() []*string {
	return s.IpWhiteList
}

func (s *GetWorkspaceResponseBodyWorkspace) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *GetWorkspaceResponseBodyWorkspace) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *GetWorkspaceResponseBodyWorkspace) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetWorkspaceResponseBodyWorkspace) GetPrePaidQuota() *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	return s.PrePaidQuota
}

func (s *GetWorkspaceResponseBodyWorkspace) GetPrePaidQuotaGpu() []*GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	return s.PrePaidQuotaGpu
}

func (s *GetWorkspaceResponseBodyWorkspace) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *GetWorkspaceResponseBodyWorkspace) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *GetWorkspaceResponseBodyWorkspace) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetResourceSpec() *string {
	return s.ResourceSpec
}

func (s *GetWorkspaceResponseBodyWorkspace) GetStateChangeReason() *GetWorkspaceResponseBodyWorkspaceStateChangeReason {
	return s.StateChangeReason
}

func (s *GetWorkspaceResponseBodyWorkspace) GetStorage() *string {
	return s.Storage
}

func (s *GetWorkspaceResponseBodyWorkspace) GetTags() []*GetWorkspaceResponseBodyWorkspaceTags {
	return s.Tags
}

func (s *GetWorkspaceResponseBodyWorkspace) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetWorkspaceResponseBodyWorkspace) GetWorkspaceStatus() *string {
	return s.WorkspaceStatus
}

func (s *GetWorkspaceResponseBodyWorkspace) SetAutoRenew(v bool) *GetWorkspaceResponseBodyWorkspace {
	s.AutoRenew = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetAutoRenewPeriod(v int32) *GetWorkspaceResponseBodyWorkspace {
	s.AutoRenewPeriod = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetAutoRenewPeriodUnit(v string) *GetWorkspaceResponseBodyWorkspace {
	s.AutoRenewPeriodUnit = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCreateTime(v int64) *GetWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetDlfCatalogId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.DlfCatalogId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetDlfType(v string) *GetWorkspaceResponseBodyWorkspace {
	s.DlfType = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetDuration(v int32) *GetWorkspaceResponseBodyWorkspace {
	s.Duration = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetEndTime(v int64) *GetWorkspaceResponseBodyWorkspace {
	s.EndTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetFailReason(v string) *GetWorkspaceResponseBodyWorkspace {
	s.FailReason = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetGpuSpec(v []*string) *GetWorkspaceResponseBodyWorkspace {
	s.GpuSpec = v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetIpWhiteList(v []*string) *GetWorkspaceResponseBodyWorkspace {
	s.IpWhiteList = v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetPaymentDurationUnit(v string) *GetWorkspaceResponseBodyWorkspace {
	s.PaymentDurationUnit = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetPaymentStatus(v string) *GetWorkspaceResponseBodyWorkspace {
	s.PaymentStatus = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetPaymentType(v string) *GetWorkspaceResponseBodyWorkspace {
	s.PaymentType = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetPrePaidQuota(v *GetWorkspaceResponseBodyWorkspacePrePaidQuota) *GetWorkspaceResponseBodyWorkspace {
	s.PrePaidQuota = v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetPrePaidQuotaGpu(v []*GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) *GetWorkspaceResponseBodyWorkspace {
	s.PrePaidQuotaGpu = v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetRamRoleName(v string) *GetWorkspaceResponseBodyWorkspace {
	s.RamRoleName = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetRegionId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.RegionId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetReleaseType(v string) *GetWorkspaceResponseBodyWorkspace {
	s.ReleaseType = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetResourceGroupId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.ResourceGroupId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetResourceOwnerId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetResourceSpec(v string) *GetWorkspaceResponseBodyWorkspace {
	s.ResourceSpec = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetStateChangeReason(v *GetWorkspaceResponseBodyWorkspaceStateChangeReason) *GetWorkspaceResponseBodyWorkspace {
	s.StateChangeReason = v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetStorage(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Storage = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetTags(v []*GetWorkspaceResponseBodyWorkspaceTags) *GetWorkspaceResponseBodyWorkspace {
	s.Tags = v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetWorkspaceId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetWorkspaceName(v string) *GetWorkspaceResponseBodyWorkspace {
	s.WorkspaceName = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetWorkspaceStatus(v string) *GetWorkspaceResponseBodyWorkspace {
	s.WorkspaceStatus = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) Validate() error {
	if s.PrePaidQuota != nil {
		if err := s.PrePaidQuota.Validate(); err != nil {
			return err
		}
	}
	if s.PrePaidQuotaGpu != nil {
		for _, item := range s.PrePaidQuotaGpu {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StateChangeReason != nil {
		if err := s.StateChangeReason.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkspaceResponseBodyWorkspacePrePaidQuota struct {
	// The amount of resources currently allocated.
	//
	// example:
	//
	// {"cu": 100}
	AllocatedResource *string `json:"allocatedResource,omitempty" xml:"allocatedResource,omitempty"`
	// Specifies whether auto-renewal is enabled for the resource.
	//
	// - true: Auto-renewal is enabled. The resource is automatically renewed upon expiration.
	//
	// - false: Auto-renewal is disabled. The resource is stopped upon expiration.
	//
	// example:
	//
	// False
	AutoRenewal *bool `json:"autoRenewal,omitempty" xml:"autoRenewal,omitempty"`
	// The time when the workspace was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1791802817000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The expiration time of the GPU subscription quota. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1779768979586
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ray-df0tghlhat1776wy-klbifz8l
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The maximum amount of resources that can be used.
	//
	// example:
	//
	// {"cu": 100}
	MaxResource *string `json:"maxResource,omitempty" xml:"maxResource,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 298581757170384
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The order status.
	//
	// example:
	//
	// NORMAL
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
	// The amount of resources currently used.
	//
	// example:
	//
	// {"cu": 50}
	UsedResource *string `json:"usedResource,omitempty" xml:"usedResource,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspacePrePaidQuota) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspacePrePaidQuota) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetAllocatedResource() *string {
	return s.AllocatedResource
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetMaxResource() *string {
	return s.MaxResource
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetOrderId() *string {
	return s.OrderId
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) GetUsedResource() *string {
	return s.UsedResource
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetAllocatedResource(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.AllocatedResource = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetAutoRenewal(v bool) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.AutoRenewal = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetCreateTime(v int64) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetExpireTime(v int64) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.ExpireTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetInstanceId(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.InstanceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetMaxResource(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.MaxResource = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetOrderId(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.OrderId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetPaymentStatus(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.PaymentStatus = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) SetUsedResource(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuota {
	s.UsedResource = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuota) Validate() error {
	return dara.Validate(s)
}

type GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu struct {
	// Specifies whether auto-renewal is enabled.
	//
	// example:
	//
	// False
	AutoRenewal *bool `json:"autoRenewal,omitempty" xml:"autoRenewal,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 1
	CpuCoreCount *string `json:"cpuCoreCount,omitempty" xml:"cpuCoreCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1782292672000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1782292772000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	GpuAmount *int32 `json:"gpuAmount,omitempty" xml:"gpuAmount,omitempty"`
	// The number of GPU instances.
	//
	// example:
	//
	// 1
	GpuMachineNum *int32 `json:"gpuMachineNum,omitempty" xml:"gpuMachineNum,omitempty"`
	// The GPU memory size.
	//
	// example:
	//
	// 4096
	GpuMemorySize *int64 `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	// The number of GPU instances.
	//
	// example:
	//
	// 8
	GpuNum *int32 `json:"gpuNum,omitempty" xml:"gpuNum,omitempty"`
	// The GPU instance type.
	//
	// example:
	//
	// ecs.gn7i-c8g1.2xlarge
	GpuSpec *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// w-xxxxxx-quota-gpu-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance family.
	//
	// example:
	//
	// ecs.g5
	InstanceTypeFamily *string `json:"instanceTypeFamily,omitempty" xml:"instanceTypeFamily,omitempty"`
	// The instance type ID.
	//
	// example:
	//
	// ecs.g5.small
	InstanceTypeId *string `json:"instanceTypeId,omitempty" xml:"instanceTypeId,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 4096
	MemorySize *string `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2345634254
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The order status.
	//
	// example:
	//
	// NORMAL
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetCpuCoreCount() *string {
	return s.CpuCoreCount
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetGpuAmount() *int32 {
	return s.GpuAmount
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetGpuMachineNum() *int32 {
	return s.GpuMachineNum
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetGpuMemorySize() *int64 {
	return s.GpuMemorySize
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetGpuNum() *int32 {
	return s.GpuNum
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetInstanceTypeId() *string {
	return s.InstanceTypeId
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetMemorySize() *string {
	return s.MemorySize
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetOrderId() *string {
	return s.OrderId
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetAutoRenewal(v bool) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.AutoRenewal = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetCpuCoreCount(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.CpuCoreCount = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetCreateTime(v int64) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetExpireTime(v int64) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.ExpireTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetGpuAmount(v int32) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.GpuAmount = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetGpuMachineNum(v int32) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.GpuMachineNum = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetGpuMemorySize(v int64) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.GpuMemorySize = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetGpuNum(v int32) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.GpuNum = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetGpuSpec(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.GpuSpec = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetInstanceId(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.InstanceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetInstanceTypeFamily(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.InstanceTypeFamily = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetInstanceTypeId(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.InstanceTypeId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetMemorySize(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.MemorySize = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetOrderId(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.OrderId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) SetPaymentStatus(v string) *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu {
	s.PaymentStatus = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspacePrePaidQuotaGpu) Validate() error {
	return dara.Validate(s)
}

type GetWorkspaceResponseBodyWorkspaceStateChangeReason struct {
	// The error code.
	//
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status information of the node pod.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspaceStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspaceStateChangeReason) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspaceStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *GetWorkspaceResponseBodyWorkspaceStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspaceResponseBodyWorkspaceStateChangeReason) SetCode(v string) *GetWorkspaceResponseBodyWorkspaceStateChangeReason {
	s.Code = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspaceStateChangeReason) SetMessage(v string) *GetWorkspaceResponseBodyWorkspaceStateChangeReason {
	s.Message = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspaceStateChangeReason) Validate() error {
	return dara.Validate(s)
}

type GetWorkspaceResponseBodyWorkspaceTags struct {
	// The tag keys.
	//
	// example:
	//
	// [\\"COST\\"]
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 3d775731-9ec1-4fe2-a00d-061d51fce677
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspaceTags) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspaceTags) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspaceTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetWorkspaceResponseBodyWorkspaceTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetWorkspaceResponseBodyWorkspaceTags) SetTagKey(v string) *GetWorkspaceResponseBodyWorkspaceTags {
	s.TagKey = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspaceTags) SetTagValue(v string) *GetWorkspaceResponseBodyWorkspaceTags {
	s.TagValue = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspaceTags) Validate() error {
	return dara.Validate(s)
}
