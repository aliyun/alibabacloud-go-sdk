// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWorkflowTemplateDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GenerateWorkflowTemplateDraftRequest
	GetCollaborationGroupId() *string
	SetPreviousDraft(v map[string]interface{}) *GenerateWorkflowTemplateDraftRequest
	GetPreviousDraft() map[string]interface{}
	SetRequirement(v string) *GenerateWorkflowTemplateDraftRequest
	GetRequirement() *string
	SetTenantId(v string) *GenerateWorkflowTemplateDraftRequest
	GetTenantId() *string
}

type GenerateWorkflowTemplateDraftRequest struct {
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
	PreviousDraft map[string]interface{} `json:"previousDraft,omitempty" xml:"previousDraft,omitempty"`
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

func (s GenerateWorkflowTemplateDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateWorkflowTemplateDraftRequest) GoString() string {
	return s.String()
}

func (s *GenerateWorkflowTemplateDraftRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GenerateWorkflowTemplateDraftRequest) GetPreviousDraft() map[string]interface{} {
	return s.PreviousDraft
}

func (s *GenerateWorkflowTemplateDraftRequest) GetRequirement() *string {
	return s.Requirement
}

func (s *GenerateWorkflowTemplateDraftRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GenerateWorkflowTemplateDraftRequest) SetCollaborationGroupId(v string) *GenerateWorkflowTemplateDraftRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftRequest) SetPreviousDraft(v map[string]interface{}) *GenerateWorkflowTemplateDraftRequest {
	s.PreviousDraft = v
	return s
}

func (s *GenerateWorkflowTemplateDraftRequest) SetRequirement(v string) *GenerateWorkflowTemplateDraftRequest {
	s.Requirement = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftRequest) SetTenantId(v string) *GenerateWorkflowTemplateDraftRequest {
	s.TenantId = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftRequest) Validate() error {
	return dara.Validate(s)
}
