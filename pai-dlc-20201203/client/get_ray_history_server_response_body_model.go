// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayHistoryServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetRayHistoryServerResponseBody
	GetAccessibility() *string
	SetDisplayName(v string) *GetRayHistoryServerResponseBody
	GetDisplayName() *string
	SetEcsSpec(v string) *GetRayHistoryServerResponseBody
	GetEcsSpec() *string
	SetGmtCreateTime(v string) *GetRayHistoryServerResponseBody
	GetGmtCreateTime() *string
	SetGmtFinishTime(v string) *GetRayHistoryServerResponseBody
	GetGmtFinishTime() *string
	SetGmtModifyTime(v string) *GetRayHistoryServerResponseBody
	GetGmtModifyTime() *string
	SetMaxRuntimeMinutes(v int32) *GetRayHistoryServerResponseBody
	GetMaxRuntimeMinutes() *int32
	SetRayHistoryServerId(v string) *GetRayHistoryServerResponseBody
	GetRayHistoryServerId() *string
	SetRayHistoryServerUrl(v string) *GetRayHistoryServerResponseBody
	GetRayHistoryServerUrl() *string
	SetReasonCode(v string) *GetRayHistoryServerResponseBody
	GetReasonCode() *string
	SetReasonMessage(v string) *GetRayHistoryServerResponseBody
	GetReasonMessage() *string
	SetResourceId(v string) *GetRayHistoryServerResponseBody
	GetResourceId() *string
	SetResourceName(v string) *GetRayHistoryServerResponseBody
	GetResourceName() *string
	SetStatus(v string) *GetRayHistoryServerResponseBody
	GetStatus() *string
	SetStoragePath(v string) *GetRayHistoryServerResponseBody
	GetStoragePath() *string
	SetTenantId(v string) *GetRayHistoryServerResponseBody
	GetTenantId() *string
	SetUserId(v string) *GetRayHistoryServerResponseBody
	GetUserId() *string
	SetUsername(v string) *GetRayHistoryServerResponseBody
	GetUsername() *string
	SetWorkspaceId(v string) *GetRayHistoryServerResponseBody
	GetWorkspaceId() *string
}

type GetRayHistoryServerResponseBody struct {
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// my-ray-history-server
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// ecs.g6.large
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// 1000
	MaxRuntimeMinutes *int32 `json:"MaxRuntimeMinutes,omitempty" xml:"MaxRuntimeMinutes,omitempty"`
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
	// example:
	//
	// NotFound
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// ""
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// quotaxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// my-resource-name
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// oss://bucket-test-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/tmp
	StoragePath *string `json:"StoragePath,omitempty" xml:"StoragePath,omitempty"`
	// example:
	//
	// 10**************14
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// myusername
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// 46099
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetRayHistoryServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRayHistoryServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetRayHistoryServerResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetRayHistoryServerResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetRayHistoryServerResponseBody) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *GetRayHistoryServerResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetRayHistoryServerResponseBody) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *GetRayHistoryServerResponseBody) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *GetRayHistoryServerResponseBody) GetMaxRuntimeMinutes() *int32 {
	return s.MaxRuntimeMinutes
}

func (s *GetRayHistoryServerResponseBody) GetRayHistoryServerId() *string {
	return s.RayHistoryServerId
}

func (s *GetRayHistoryServerResponseBody) GetRayHistoryServerUrl() *string {
	return s.RayHistoryServerUrl
}

func (s *GetRayHistoryServerResponseBody) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetRayHistoryServerResponseBody) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetRayHistoryServerResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetRayHistoryServerResponseBody) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetRayHistoryServerResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRayHistoryServerResponseBody) GetStoragePath() *string {
	return s.StoragePath
}

func (s *GetRayHistoryServerResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *GetRayHistoryServerResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetRayHistoryServerResponseBody) GetUsername() *string {
	return s.Username
}

func (s *GetRayHistoryServerResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetRayHistoryServerResponseBody) SetAccessibility(v string) *GetRayHistoryServerResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetDisplayName(v string) *GetRayHistoryServerResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetEcsSpec(v string) *GetRayHistoryServerResponseBody {
	s.EcsSpec = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetGmtCreateTime(v string) *GetRayHistoryServerResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetGmtFinishTime(v string) *GetRayHistoryServerResponseBody {
	s.GmtFinishTime = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetGmtModifyTime(v string) *GetRayHistoryServerResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetMaxRuntimeMinutes(v int32) *GetRayHistoryServerResponseBody {
	s.MaxRuntimeMinutes = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetRayHistoryServerId(v string) *GetRayHistoryServerResponseBody {
	s.RayHistoryServerId = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetRayHistoryServerUrl(v string) *GetRayHistoryServerResponseBody {
	s.RayHistoryServerUrl = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetReasonCode(v string) *GetRayHistoryServerResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetReasonMessage(v string) *GetRayHistoryServerResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetResourceId(v string) *GetRayHistoryServerResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetResourceName(v string) *GetRayHistoryServerResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetStatus(v string) *GetRayHistoryServerResponseBody {
	s.Status = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetStoragePath(v string) *GetRayHistoryServerResponseBody {
	s.StoragePath = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetTenantId(v string) *GetRayHistoryServerResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetUserId(v string) *GetRayHistoryServerResponseBody {
	s.UserId = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetUsername(v string) *GetRayHistoryServerResponseBody {
	s.Username = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) SetWorkspaceId(v string) *GetRayHistoryServerResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetRayHistoryServerResponseBody) Validate() error {
	return dara.Validate(s)
}
