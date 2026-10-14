// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListWorkspacesResponseBody
	GetTotalCount() *int32
	SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody
	GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces
}

type ListWorkspacesResponseBody struct {
	// The maximum number of records to return in a single request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The next token.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The list of workspaces.
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkspacesResponseBody) GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces {
	return s.Workspaces
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int32) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

func (s *ListWorkspacesResponseBody) Validate() error {
	if s.Workspaces != nil {
		for _, item := range s.Workspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesResponseBodyWorkspaces struct {
	// Specifies whether to enable auto-renewal. This parameter is required for subscription resources.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The auto-renewal duration. This parameter is required for subscription resources.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The auto-renewal epoch. This parameter is required for subscription resources.
	//
	// example:
	//
	// YEAR, MONTH, WEEK, DAY, HOUR, MINUTE
	AutoRenewPeriodUnit *string `json:"autoRenewPeriodUnit,omitempty" xml:"autoRenewPeriodUnit,omitempty"`
	// The time when the workspace was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1684115879955
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The DLF catalog information.
	//
	// example:
	//
	// default
	DlfCatalogId *string `json:"dlfCatalogId,omitempty" xml:"dlfCatalogId,omitempty"`
	// The type of the associated DLF.
	//
	// example:
	//
	// 1.0
	DlfType *string `json:"dlfType,omitempty" xml:"dlfType,omitempty"`
	// The duration of the subscription. This parameter is required for subscription resources.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The time when the workspace was released. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1687103999999
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// out of stock
	FailReason *string `json:"failReason,omitempty" xml:"failReason,omitempty"`
	// The list of GPU models.
	GpuSpec []*string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty" type:"Repeated"`
	// The IP whitelist of the workspace.
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
	// The unit of the subscription period. This parameter is required for subscription resources.
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
	PrePaidQuota *ListWorkspacesResponseBodyWorkspacesPrePaidQuota `json:"prePaidQuota,omitempty" xml:"prePaidQuota,omitempty" type:"Struct"`
	// The list of GPU subscription quotas.
	PrePaidQuotaGpu []*ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu `json:"prePaidQuotaGpu,omitempty" xml:"prePaidQuotaGpu,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The reason why the workspace was released.
	//
	// example:
	//
	// SERVICE_RELEASE
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxxxxxxxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The resource specification.
	//
	// example:
	//
	// 100cu
	ResourceSpec *string `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	// The information about the workspace status change.
	StateChangeReason *ListWorkspacesResponseBodyWorkspacesStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The OSS path.
	//
	// example:
	//
	// spark-result
	Storage *string `json:"storage,omitempty" xml:"storage,omitempty"`
	// The list of workspace tags.
	Tags []*ListWorkspacesResponseBodyWorkspacesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// w-******
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// spark-batch-job-workspace-1
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// The status of the workspace.
	//
	// example:
	//
	// STARTING,RUNNING,TERMINATED
	WorkspaceStatus *string `json:"workspaceStatus,omitempty" xml:"workspaceStatus,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetAutoRenewPeriodUnit() *string {
	return s.AutoRenewPeriodUnit
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetDlfCatalogId() *string {
	return s.DlfCatalogId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetDlfType() *string {
	return s.DlfType
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetDuration() *int32 {
	return s.Duration
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetFailReason() *string {
	return s.FailReason
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetGpuSpec() []*string {
	return s.GpuSpec
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetIpWhiteList() []*string {
	return s.IpWhiteList
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetPrePaidQuota() *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	return s.PrePaidQuota
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetPrePaidQuotaGpu() []*ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	return s.PrePaidQuotaGpu
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetResourceSpec() *string {
	return s.ResourceSpec
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetStateChangeReason() *ListWorkspacesResponseBodyWorkspacesStateChangeReason {
	return s.StateChangeReason
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetStorage() *string {
	return s.Storage
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetTags() []*ListWorkspacesResponseBodyWorkspacesTags {
	return s.Tags
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceStatus() *string {
	return s.WorkspaceStatus
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetAutoRenew(v bool) *ListWorkspacesResponseBodyWorkspaces {
	s.AutoRenew = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetAutoRenewPeriod(v int32) *ListWorkspacesResponseBodyWorkspaces {
	s.AutoRenewPeriod = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetAutoRenewPeriodUnit(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.AutoRenewPeriodUnit = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v int64) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDlfCatalogId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.DlfCatalogId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDlfType(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.DlfType = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDuration(v int32) *ListWorkspacesResponseBodyWorkspaces {
	s.Duration = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetEndTime(v int64) *ListWorkspacesResponseBodyWorkspaces {
	s.EndTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetFailReason(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.FailReason = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetGpuSpec(v []*string) *ListWorkspacesResponseBodyWorkspaces {
	s.GpuSpec = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetIpWhiteList(v []*string) *ListWorkspacesResponseBodyWorkspaces {
	s.IpWhiteList = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPaymentDurationUnit(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.PaymentDurationUnit = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPaymentStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.PaymentStatus = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPaymentType(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.PaymentType = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPrePaidQuota(v *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) *ListWorkspacesResponseBodyWorkspaces {
	s.PrePaidQuota = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPrePaidQuotaGpu(v []*ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) *ListWorkspacesResponseBodyWorkspaces {
	s.PrePaidQuotaGpu = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRegionId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetReleaseType(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ReleaseType = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetResourceGroupId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ResourceGroupId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetResourceSpec(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ResourceSpec = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStateChangeReason(v *ListWorkspacesResponseBodyWorkspacesStateChangeReason) *ListWorkspacesResponseBodyWorkspaces {
	s.StateChangeReason = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStorage(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Storage = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetTags(v []*ListWorkspacesResponseBodyWorkspacesTags) *ListWorkspacesResponseBodyWorkspaces {
	s.Tags = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceStatus = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) Validate() error {
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

type ListWorkspacesResponseBodyWorkspacesPrePaidQuota struct {
	// The amount of resources currently allocated.
	//
	// example:
	//
	// {\\"cpu\\":\\"1\\",\\"memory\\":\\"4Gi\\",\\"cu\\":\\"1\\"}
	AllocatedResource *string `json:"allocatedResource,omitempty" xml:"allocatedResource,omitempty"`
	// Specifies whether auto-renewal is enabled for the resource. Valid values:
	//
	// - true: Auto-renewal is enabled. The resource is automatically renewed upon expiration.
	//
	// - false: Auto-renewal is disabled. The resource is stopped upon expiration.
	//
	// example:
	//
	// true
	AutoRenewal *bool `json:"autoRenewal,omitempty" xml:"autoRenewal,omitempty"`
	// The time when the resource quota was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1745683200000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The expiration time of the resource quota. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1740537153000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The instance ID of the resource associated with the quota.
	//
	// example:
	//
	// i-abc12345
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The maximum amount of resources that can be used.
	//
	// example:
	//
	// {\\"cpu\\":\\"1\\",\\"memory\\":\\"4Gi\\",\\"cu\\":\\"1\\"}
	MaxResource *string `json:"maxResource,omitempty" xml:"maxResource,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 23464687565
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The payment status of the current resource. Valid values:
	//
	// - NORMAL: Valid.
	//
	// - WAIT_FOR_EXPIRE: Expiring.
	//
	// - EXPIRED: Expired.
	//
	// example:
	//
	// NORMAL
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
	// The amount of resources currently used.
	//
	// example:
	//
	// {\\"cpu\\":\\"0\\",\\"memory\\":\\"0Gi\\",\\"cu\\":\\"0\\"}
	UsedResource *string `json:"usedResource,omitempty" xml:"usedResource,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesPrePaidQuota) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetAllocatedResource() *string {
	return s.AllocatedResource
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetMaxResource() *string {
	return s.MaxResource
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetOrderId() *string {
	return s.OrderId
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) GetUsedResource() *string {
	return s.UsedResource
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetAllocatedResource(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.AllocatedResource = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetAutoRenewal(v bool) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.AutoRenewal = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetCreateTime(v int64) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetExpireTime(v int64) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.ExpireTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetInstanceId(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.InstanceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetMaxResource(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.MaxResource = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetOrderId(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.OrderId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetPaymentStatus(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.PaymentStatus = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) SetUsedResource(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuota {
	s.UsedResource = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuota) Validate() error {
	return dara.Validate(s)
}

type ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu struct {
	// Specifies whether auto-renewal is enabled.
	//
	// example:
	//
	// true
	AutoRenewal *bool `json:"autoRenewal,omitempty" xml:"autoRenewal,omitempty"`
	// The number of CPU cores.
	CpuCoreCount *string `json:"cpuCoreCount,omitempty" xml:"cpuCoreCount,omitempty"`
	// The time when the GPU subscription quota was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1782292672000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The expiration time of the GPU subscription quota. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1782292772000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 0
	GpuAmount *int32 `json:"gpuAmount,omitempty" xml:"gpuAmount,omitempty"`
	// The number of GPU machines.
	//
	// example:
	//
	// 4
	GpuMachineNum *int32 `json:"gpuMachineNum,omitempty" xml:"gpuMachineNum,omitempty"`
	// The GPU memory size.
	//
	// example:
	//
	// 0
	GpuMemorySize *int64 `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 8
	GpuNum *int32 `json:"gpuNum,omitempty" xml:"gpuNum,omitempty"`
	// The GPU model.
	//
	// example:
	//
	// ecs.gn7i-c8g1.2xlarge
	GpuSpec *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// w-xxxxxxxxx-gpu-quota-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance family.
	InstanceTypeFamily *string `json:"instanceTypeFamily,omitempty" xml:"instanceTypeFamily,omitempty"`
	// The instance type ID.
	InstanceTypeId *string `json:"instanceTypeId,omitempty" xml:"instanceTypeId,omitempty"`
	// The memory size.
	MemorySize *string `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2534863936
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The order status.
	//
	// example:
	//
	// NORMAL
	PaymentStatus *string `json:"paymentStatus,omitempty" xml:"paymentStatus,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetCpuCoreCount() *string {
	return s.CpuCoreCount
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetGpuAmount() *int32 {
	return s.GpuAmount
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetGpuMachineNum() *int32 {
	return s.GpuMachineNum
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetGpuMemorySize() *int64 {
	return s.GpuMemorySize
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetGpuNum() *int32 {
	return s.GpuNum
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetInstanceTypeId() *string {
	return s.InstanceTypeId
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetMemorySize() *string {
	return s.MemorySize
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetOrderId() *string {
	return s.OrderId
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetAutoRenewal(v bool) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.AutoRenewal = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetCpuCoreCount(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.CpuCoreCount = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetCreateTime(v int64) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetExpireTime(v int64) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.ExpireTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetGpuAmount(v int32) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.GpuAmount = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetGpuMachineNum(v int32) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.GpuMachineNum = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetGpuMemorySize(v int64) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.GpuMemorySize = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetGpuNum(v int32) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.GpuNum = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetGpuSpec(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.GpuSpec = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetInstanceId(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.InstanceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetInstanceTypeFamily(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.InstanceTypeFamily = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetInstanceTypeId(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.InstanceTypeId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetMemorySize(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.MemorySize = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetOrderId(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.OrderId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) SetPaymentStatus(v string) *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu {
	s.PaymentStatus = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesPrePaidQuotaGpu) Validate() error {
	return dara.Validate(s)
}

type ListWorkspacesResponseBodyWorkspacesStateChangeReason struct {
	// The error code.
	//
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesStateChangeReason) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *ListWorkspacesResponseBodyWorkspacesStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *ListWorkspacesResponseBodyWorkspacesStateChangeReason) SetCode(v string) *ListWorkspacesResponseBodyWorkspacesStateChangeReason {
	s.Code = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesStateChangeReason) SetMessage(v string) *ListWorkspacesResponseBodyWorkspacesStateChangeReason {
	s.Message = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesStateChangeReason) Validate() error {
	return dara.Validate(s)
}

type ListWorkspacesResponseBodyWorkspacesTags struct {
	// The tag key.
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesTags) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesTags) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListWorkspacesResponseBodyWorkspacesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListWorkspacesResponseBodyWorkspacesTags) SetTagKey(v string) *ListWorkspacesResponseBodyWorkspacesTags {
	s.TagKey = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesTags) SetTagValue(v string) *ListWorkspacesResponseBodyWorkspacesTags {
	s.TagValue = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesTags) Validate() error {
	return dara.Validate(s)
}
