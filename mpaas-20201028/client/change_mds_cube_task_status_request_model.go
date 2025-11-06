// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMdsCubeTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ChangeMdsCubeTaskStatusRequest
	GetAppId() *string
	SetTaskStatus(v int32) *ChangeMdsCubeTaskStatusRequest
	GetTaskStatus() *int32
	SetTemplateResourceId(v int64) *ChangeMdsCubeTaskStatusRequest
	GetTemplateResourceId() *int64
	SetTemplateTaskId(v int64) *ChangeMdsCubeTaskStatusRequest
	GetTemplateTaskId() *int64
	SetTenantId(v string) *ChangeMdsCubeTaskStatusRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ChangeMdsCubeTaskStatusRequest
	GetWorkspaceId() *string
}

type ChangeMdsCubeTaskStatusRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1
	TemplateResourceId *int64 `json:"TemplateResourceId,omitempty" xml:"TemplateResourceId,omitempty"`
	// example:
	//
	// 1
	TemplateTaskId *int64 `json:"TemplateTaskId,omitempty" xml:"TemplateTaskId,omitempty"`
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ChangeMdsCubeTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeMdsCubeTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeMdsCubeTaskStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ChangeMdsCubeTaskStatusRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ChangeMdsCubeTaskStatusRequest) GetTemplateResourceId() *int64 {
	return s.TemplateResourceId
}

func (s *ChangeMdsCubeTaskStatusRequest) GetTemplateTaskId() *int64 {
	return s.TemplateTaskId
}

func (s *ChangeMdsCubeTaskStatusRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ChangeMdsCubeTaskStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ChangeMdsCubeTaskStatusRequest) SetAppId(v string) *ChangeMdsCubeTaskStatusRequest {
	s.AppId = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusRequest) SetTaskStatus(v int32) *ChangeMdsCubeTaskStatusRequest {
	s.TaskStatus = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusRequest) SetTemplateResourceId(v int64) *ChangeMdsCubeTaskStatusRequest {
	s.TemplateResourceId = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusRequest) SetTemplateTaskId(v int64) *ChangeMdsCubeTaskStatusRequest {
	s.TemplateTaskId = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusRequest) SetTenantId(v string) *ChangeMdsCubeTaskStatusRequest {
	s.TenantId = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusRequest) SetWorkspaceId(v string) *ChangeMdsCubeTaskStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
