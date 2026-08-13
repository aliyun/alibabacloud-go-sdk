// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySemanticKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *QuerySemanticKnowledgeRequest
	GetAgentName() *string
	SetGraphName(v string) *QuerySemanticKnowledgeRequest
	GetGraphName() *string
	SetQuery(v string) *QuerySemanticKnowledgeRequest
	GetQuery() *string
	SetTenantId(v string) *QuerySemanticKnowledgeRequest
	GetTenantId() *string
}

type QuerySemanticKnowledgeRequest struct {
	// 数字员工名称，可先调用 listAuthorizedAgents 获取 USE 权限列表
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例数字员工
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// 图谱名称，可先调用 listGraphs 获取
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 自然语言查询问题
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// This parameter is required.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QuerySemanticKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySemanticKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *QuerySemanticKnowledgeRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *QuerySemanticKnowledgeRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *QuerySemanticKnowledgeRequest) GetQuery() *string {
	return s.Query
}

func (s *QuerySemanticKnowledgeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySemanticKnowledgeRequest) SetAgentName(v string) *QuerySemanticKnowledgeRequest {
	s.AgentName = &v
	return s
}

func (s *QuerySemanticKnowledgeRequest) SetGraphName(v string) *QuerySemanticKnowledgeRequest {
	s.GraphName = &v
	return s
}

func (s *QuerySemanticKnowledgeRequest) SetQuery(v string) *QuerySemanticKnowledgeRequest {
	s.Query = &v
	return s
}

func (s *QuerySemanticKnowledgeRequest) SetTenantId(v string) *QuerySemanticKnowledgeRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySemanticKnowledgeRequest) Validate() error {
	return dara.Validate(s)
}
