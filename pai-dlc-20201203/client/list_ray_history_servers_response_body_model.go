// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayHistoryServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRayHistoryServers(v []*ListRayHistoryServersResponseBodyRayHistoryServers) *ListRayHistoryServersResponseBody
	GetRayHistoryServers() []*ListRayHistoryServersResponseBodyRayHistoryServers
	SetRequestId(v string) *ListRayHistoryServersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRayHistoryServersResponseBody
	GetTotalCount() *int32
}

type ListRayHistoryServersResponseBody struct {
	// The list of RayHistoryServer entries.
	RayHistoryServers []*ListRayHistoryServersResponseBodyRayHistoryServers `json:"RayHistoryServers,omitempty" xml:"RayHistoryServers,omitempty" type:"Repeated"`
	// The request ID, which is used for diagnostics and troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that match the filter conditions.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRayHistoryServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRayHistoryServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRayHistoryServersResponseBody) GetRayHistoryServers() []*ListRayHistoryServersResponseBodyRayHistoryServers {
	return s.RayHistoryServers
}

func (s *ListRayHistoryServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRayHistoryServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRayHistoryServersResponseBody) SetRayHistoryServers(v []*ListRayHistoryServersResponseBodyRayHistoryServers) *ListRayHistoryServersResponseBody {
	s.RayHistoryServers = v
	return s
}

func (s *ListRayHistoryServersResponseBody) SetRequestId(v string) *ListRayHistoryServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRayHistoryServersResponseBody) SetTotalCount(v int32) *ListRayHistoryServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRayHistoryServersResponseBody) Validate() error {
	if s.RayHistoryServers != nil {
		for _, item := range s.RayHistoryServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRayHistoryServersResponseBodyRayHistoryServers struct {
	// The visibility of the job. Valid values:
	//
	// - PUBLIC: visible to all users in the workspace.
	//
	// - PRIVATE (default): visible only to you and administrators in the workspace.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The display name of the job.
	//
	// example:
	//
	// AEB-RECHARGE-TASK-14478-1778466397-main-4-
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The hardware specifications of the public resource group. Visit [PAI-DLC billing](https://help.aliyun.com/document_detail/171758.html) for a detailed list of specifications.	Notice: Prices vary depending on the specifications..
	//
	// example:
	//
	// ecs.g6.xlarge
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// The time when the job was created, in UTC.
	//
	// example:
	//
	// 2025-12-30T02:43:52Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the job ended, in UTC.
	//
	// example:
	//
	// 2026-01-27T09:17:11Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// The time when the job was last modified, in UTC.
	//
	// example:
	//
	// 2026-05-19T04:05:46Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// The maximum runtime in minutes.
	//
	// example:
	//
	// 1000
	MaxRuntimeMinutes *int32 `json:"MaxRuntimeMinutes,omitempty" xml:"MaxRuntimeMinutes,omitempty"`
	// The ID of the created RayHistoryServer.
	//
	// example:
	//
	// rhsxxx
	RayHistoryServerId *string `json:"RayHistoryServerId,omitempty" xml:"RayHistoryServerId,omitempty"`
	// Ray Dashboard URL。
	//
	// example:
	//
	// https://rhsxxx-dashboard.dsw-gateway-cn-wulanchabu.data.aliyun.com/
	RayHistoryServerUrl *string `json:"RayHistoryServerUrl,omitempty" xml:"RayHistoryServerUrl,omitempty"`
	// The status detail code.
	//
	// example:
	//
	// InvalidParameter
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The status details.
	//
	// example:
	//
	// PyTorchJob dlc1tx4b9lw3ntb9 is running.
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The resource group ID. For information about how to query the ID of a dedicated resource group, see [Manage resource quotas](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// quotazoqd53w0q75
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource on which the job runs.
	//
	// example:
	//
	// OWNER_REPO
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The RayHistoryServer status. Valid values:
	//
	// - Creating: being created.
	//
	// - Running: running.
	//
	// - Stopped: stopped.
	//
	// - Succeeded: succeeded.
	//
	// - Failed: failed.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage path of Ray logs.
	//
	// example:
	//
	// oss://bucket-test-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/tmp
	StoragePath *string `json:"StoragePath,omitempty" xml:"StoragePath,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1335237941080704
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// myusername
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The workspace ID. <props="china">For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html)..
	//
	// example:
	//
	// 153466
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListRayHistoryServersResponseBodyRayHistoryServers) String() string {
	return dara.Prettify(s)
}

func (s ListRayHistoryServersResponseBodyRayHistoryServers) GoString() string {
	return s.String()
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetMaxRuntimeMinutes() *int32 {
	return s.MaxRuntimeMinutes
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetRayHistoryServerId() *string {
	return s.RayHistoryServerId
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetRayHistoryServerUrl() *string {
	return s.RayHistoryServerUrl
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetStatus() *string {
	return s.Status
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetStoragePath() *string {
	return s.StoragePath
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetTenantId() *string {
	return s.TenantId
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetUserId() *string {
	return s.UserId
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetUsername() *string {
	return s.Username
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetAccessibility(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.Accessibility = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetDisplayName(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.DisplayName = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetEcsSpec(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.EcsSpec = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetGmtCreateTime(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.GmtCreateTime = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetGmtFinishTime(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.GmtFinishTime = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetGmtModifyTime(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.GmtModifyTime = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetMaxRuntimeMinutes(v int32) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.MaxRuntimeMinutes = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetRayHistoryServerId(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.RayHistoryServerId = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetRayHistoryServerUrl(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.RayHistoryServerUrl = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetReasonCode(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.ReasonCode = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetReasonMessage(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.ReasonMessage = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetResourceId(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.ResourceId = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetResourceName(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.ResourceName = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetStatus(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.Status = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetStoragePath(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.StoragePath = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetTenantId(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.TenantId = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetUserId(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.UserId = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetUsername(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.Username = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) SetWorkspaceId(v string) *ListRayHistoryServersResponseBodyRayHistoryServers {
	s.WorkspaceId = &v
	return s
}

func (s *ListRayHistoryServersResponseBodyRayHistoryServers) Validate() error {
	return dara.Validate(s)
}
