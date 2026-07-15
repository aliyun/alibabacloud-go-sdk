// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRayHistoryServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateRayHistoryServerRequest
	GetAccessibility() *string
	SetDisplayName(v string) *CreateRayHistoryServerRequest
	GetDisplayName() *string
	SetEcsSpec(v string) *CreateRayHistoryServerRequest
	GetEcsSpec() *string
	SetMaxRuntimeMinutes(v int32) *CreateRayHistoryServerRequest
	GetMaxRuntimeMinutes() *int32
	SetResourceId(v string) *CreateRayHistoryServerRequest
	GetResourceId() *string
	SetStoragePath(v string) *CreateRayHistoryServerRequest
	GetStoragePath() *string
	SetWorkspaceId(v string) *CreateRayHistoryServerRequest
	GetWorkspaceId() *string
}

type CreateRayHistoryServerRequest struct {
	// The visibility of the job. Valid values:
	//
	// - PUBLIC: visible to all members in the workspace.
	//
	// - PRIVATE: visible only to you and administrators in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The display name of the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-ray-history-server
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The hardware specifications of the public resource group. To access the detailed list of specifications, see [PAI-DLC billing](https://help.aliyun.com/document_detail/171758.html).	Notice: Prices vary depending on the specifications..
	//
	// example:
	//
	// ecs.c6.large
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// The maximum runtime in minutes.
	//
	// example:
	//
	// 1000
	MaxRuntimeMinutes *int32 `json:"MaxRuntimeMinutes,omitempty" xml:"MaxRuntimeMinutes,omitempty"`
	// The resource group ID. This parameter is optional.
	//
	// - If this parameter is left empty, the job is committed to the public resource group.
	//
	// - If the current workspace is attached to a resource quota, you can specify the corresponding resource quota ID. For details about how to query the resource quota ID, see [Manage resource quotas](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// quotaxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The storage path for Ray logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket-test-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/tmp
	StoragePath *string `json:"StoragePath,omitempty" xml:"StoragePath,omitempty"`
	// The workspace ID. <props="china">For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html)..
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateRayHistoryServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRayHistoryServerRequest) GoString() string {
	return s.String()
}

func (s *CreateRayHistoryServerRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateRayHistoryServerRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateRayHistoryServerRequest) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *CreateRayHistoryServerRequest) GetMaxRuntimeMinutes() *int32 {
	return s.MaxRuntimeMinutes
}

func (s *CreateRayHistoryServerRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateRayHistoryServerRequest) GetStoragePath() *string {
	return s.StoragePath
}

func (s *CreateRayHistoryServerRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateRayHistoryServerRequest) SetAccessibility(v string) *CreateRayHistoryServerRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateRayHistoryServerRequest) SetDisplayName(v string) *CreateRayHistoryServerRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateRayHistoryServerRequest) SetEcsSpec(v string) *CreateRayHistoryServerRequest {
	s.EcsSpec = &v
	return s
}

func (s *CreateRayHistoryServerRequest) SetMaxRuntimeMinutes(v int32) *CreateRayHistoryServerRequest {
	s.MaxRuntimeMinutes = &v
	return s
}

func (s *CreateRayHistoryServerRequest) SetResourceId(v string) *CreateRayHistoryServerRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateRayHistoryServerRequest) SetStoragePath(v string) *CreateRayHistoryServerRequest {
	s.StoragePath = &v
	return s
}

func (s *CreateRayHistoryServerRequest) SetWorkspaceId(v string) *CreateRayHistoryServerRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateRayHistoryServerRequest) Validate() error {
	return dara.Validate(s)
}
