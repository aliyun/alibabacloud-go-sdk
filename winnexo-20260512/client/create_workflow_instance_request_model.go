// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *CreateWorkflowInstanceRequest
	GetAcceptanceCriteria() *string
	SetBusinessId(v string) *CreateWorkflowInstanceRequest
	GetBusinessId() *string
	SetCollaborationGroupId(v string) *CreateWorkflowInstanceRequest
	GetCollaborationGroupId() *string
	SetInitiator(v string) *CreateWorkflowInstanceRequest
	GetInitiator() *string
	SetPriority(v string) *CreateWorkflowInstanceRequest
	GetPriority() *string
	SetTemplateId(v string) *CreateWorkflowInstanceRequest
	GetTemplateId() *string
	SetTenantId(v string) *CreateWorkflowInstanceRequest
	GetTenantId() *string
	SetTitle(v string) *CreateWorkflowInstanceRequest
	GetTitle() *string
	SetVariables(v map[string]interface{}) *CreateWorkflowInstanceRequest
	GetVariables() map[string]interface{}
}

type CreateWorkflowInstanceRequest struct {
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
	Variables map[string]interface{} `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s CreateWorkflowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstanceRequest) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *CreateWorkflowInstanceRequest) GetBusinessId() *string {
	return s.BusinessId
}

func (s *CreateWorkflowInstanceRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowInstanceRequest) GetInitiator() *string {
	return s.Initiator
}

func (s *CreateWorkflowInstanceRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateWorkflowInstanceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkflowInstanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkflowInstanceRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkflowInstanceRequest) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *CreateWorkflowInstanceRequest) SetAcceptanceCriteria(v string) *CreateWorkflowInstanceRequest {
	s.AcceptanceCriteria = &v
	return s
}

func (s *CreateWorkflowInstanceRequest) SetBusinessId(v string) *CreateWorkflowInstanceRequest {
	s.BusinessId = &v
	return s
}

func (s *CreateWorkflowInstanceRequest) SetCollaborationGroupId(v string) *CreateWorkflowInstanceRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowInstanceRequest) SetInitiator(v string) *CreateWorkflowInstanceRequest {
	s.Initiator = &v
	return s
}

func (s *CreateWorkflowInstanceRequest) SetPriority(v string) *CreateWorkflowInstanceRequest {
	s.Priority = &v
	return s
}

func (s *CreateWorkflowInstanceRequest) SetTemplateId(v string) *CreateWorkflowInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkflowInstanceRequest) SetTenantId(v string) *CreateWorkflowInstanceRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkflowInstanceRequest) SetTitle(v string) *CreateWorkflowInstanceRequest {
	s.Title = &v
	return s
}

func (s *CreateWorkflowInstanceRequest) SetVariables(v map[string]interface{}) *CreateWorkflowInstanceRequest {
	s.Variables = v
	return s
}

func (s *CreateWorkflowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
