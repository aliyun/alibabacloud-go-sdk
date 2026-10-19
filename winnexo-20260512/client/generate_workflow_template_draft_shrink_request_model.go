// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWorkflowTemplateDraftShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GenerateWorkflowTemplateDraftShrinkRequest
	GetCollaborationGroupId() *string
	SetPreviousDraftShrink(v string) *GenerateWorkflowTemplateDraftShrinkRequest
	GetPreviousDraftShrink() *string
	SetRequirement(v string) *GenerateWorkflowTemplateDraftShrinkRequest
	GetRequirement() *string
	SetTenantId(v string) *GenerateWorkflowTemplateDraftShrinkRequest
	GetTenantId() *string
}

type GenerateWorkflowTemplateDraftShrinkRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 上一轮草稿；提供时基于它迭代修改
	//
	// example:
	//
	// example
	PreviousDraftShrink *string `json:"previousDraft,omitempty" xml:"previousDraft,omitempty"`
	// 模板需求描述
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Requirement *string `json:"requirement,omitempty" xml:"requirement,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GenerateWorkflowTemplateDraftShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateWorkflowTemplateDraftShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) GetPreviousDraftShrink() *string {
	return s.PreviousDraftShrink
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) GetRequirement() *string {
	return s.Requirement
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) SetCollaborationGroupId(v string) *GenerateWorkflowTemplateDraftShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) SetPreviousDraftShrink(v string) *GenerateWorkflowTemplateDraftShrinkRequest {
	s.PreviousDraftShrink = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) SetRequirement(v string) *GenerateWorkflowTemplateDraftShrinkRequest {
	s.Requirement = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) SetTenantId(v string) *GenerateWorkflowTemplateDraftShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftShrinkRequest) Validate() error {
	return dara.Validate(s)
}
