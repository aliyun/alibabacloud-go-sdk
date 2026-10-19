// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *CreateWorkflowInstanceShrinkRequest
	GetAcceptanceCriteria() *string
	SetBusinessId(v string) *CreateWorkflowInstanceShrinkRequest
	GetBusinessId() *string
	SetCollaborationGroupId(v string) *CreateWorkflowInstanceShrinkRequest
	GetCollaborationGroupId() *string
	SetInitiator(v string) *CreateWorkflowInstanceShrinkRequest
	GetInitiator() *string
	SetPriority(v string) *CreateWorkflowInstanceShrinkRequest
	GetPriority() *string
	SetTemplateId(v string) *CreateWorkflowInstanceShrinkRequest
	GetTemplateId() *string
	SetTenantId(v string) *CreateWorkflowInstanceShrinkRequest
	GetTenantId() *string
	SetTitle(v string) *CreateWorkflowInstanceShrinkRequest
	GetTitle() *string
	SetVariablesShrink(v string) *CreateWorkflowInstanceShrinkRequest
	GetVariablesShrink() *string
}

type CreateWorkflowInstanceShrinkRequest struct {
	// Workflow 必填验收标准
	//
	// This parameter is required.
	//
	// example:
	//
	// 完成全部审批节点并产出报告
	AcceptanceCriteria *string `json:"acceptanceCriteria,omitempty" xml:"acceptanceCriteria,omitempty"`
	// 空间内唯一需求编号，重复请求幂等
	//
	// This parameter is required.
	//
	// example:
	//
	// REQ-1001
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_20261016_abc123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 外部发起人展示标签
	//
	// example:
	//
	// 外部发起人张三
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// 工作项优先级
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// 工作流模板 ID，最多 64 字符
	//
	// This parameter is required.
	//
	// example:
	//
	// wft_20261016_abc123
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 需求标题
	//
	// This parameter is required.
	//
	// example:
	//
	// 季度经营分析报告
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 业务变量对象，tags 可为字符串标签列表
	//
	// example:
	//
	// tags=季度报告
	VariablesShrink *string `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s CreateWorkflowInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstanceShrinkRequest) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *CreateWorkflowInstanceShrinkRequest) GetBusinessId() *string {
	return s.BusinessId
}

func (s *CreateWorkflowInstanceShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowInstanceShrinkRequest) GetInitiator() *string {
	return s.Initiator
}

func (s *CreateWorkflowInstanceShrinkRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateWorkflowInstanceShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkflowInstanceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkflowInstanceShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkflowInstanceShrinkRequest) GetVariablesShrink() *string {
	return s.VariablesShrink
}

func (s *CreateWorkflowInstanceShrinkRequest) SetAcceptanceCriteria(v string) *CreateWorkflowInstanceShrinkRequest {
	s.AcceptanceCriteria = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetBusinessId(v string) *CreateWorkflowInstanceShrinkRequest {
	s.BusinessId = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetCollaborationGroupId(v string) *CreateWorkflowInstanceShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetInitiator(v string) *CreateWorkflowInstanceShrinkRequest {
	s.Initiator = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetPriority(v string) *CreateWorkflowInstanceShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetTemplateId(v string) *CreateWorkflowInstanceShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetTenantId(v string) *CreateWorkflowInstanceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetTitle(v string) *CreateWorkflowInstanceShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) SetVariablesShrink(v string) *CreateWorkflowInstanceShrinkRequest {
	s.VariablesShrink = &v
	return s
}

func (s *CreateWorkflowInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
