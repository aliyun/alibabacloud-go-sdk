// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunWorkflowNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RerunWorkflowNodeRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *RerunWorkflowNodeRequest
	GetCollaborationGroupId() *string
	SetNodeKey(v string) *RerunWorkflowNodeRequest
	GetNodeKey() *string
	SetOpsUserCode(v string) *RerunWorkflowNodeRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RerunWorkflowNodeRequest
	GetOpsUserId() *string
	SetPlanId(v string) *RerunWorkflowNodeRequest
	GetPlanId() *string
	SetTenantId(v string) *RerunWorkflowNodeRequest
	GetTenantId() *string
}

type RerunWorkflowNodeRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The key of the starting node for the rerun, which includes fixed planning nodes and business execution nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// query.amount
	NodeKey *string `json:"nodeKey,omitempty" xml:"nodeKey,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The current physical plan ID after the rerun. Rerunning a fixed planning node may generate a new plan. Use this value for subsequent queries.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan_v3
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RerunWorkflowNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunWorkflowNodeRequest) GoString() string {
	return s.String()
}

func (s *RerunWorkflowNodeRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RerunWorkflowNodeRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *RerunWorkflowNodeRequest) GetNodeKey() *string {
	return s.NodeKey
}

func (s *RerunWorkflowNodeRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RerunWorkflowNodeRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RerunWorkflowNodeRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *RerunWorkflowNodeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RerunWorkflowNodeRequest) SetAppToken(v string) *RerunWorkflowNodeRequest {
	s.AppToken = &v
	return s
}

func (s *RerunWorkflowNodeRequest) SetCollaborationGroupId(v string) *RerunWorkflowNodeRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *RerunWorkflowNodeRequest) SetNodeKey(v string) *RerunWorkflowNodeRequest {
	s.NodeKey = &v
	return s
}

func (s *RerunWorkflowNodeRequest) SetOpsUserCode(v string) *RerunWorkflowNodeRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RerunWorkflowNodeRequest) SetOpsUserId(v string) *RerunWorkflowNodeRequest {
	s.OpsUserId = &v
	return s
}

func (s *RerunWorkflowNodeRequest) SetPlanId(v string) *RerunWorkflowNodeRequest {
	s.PlanId = &v
	return s
}

func (s *RerunWorkflowNodeRequest) SetTenantId(v string) *RerunWorkflowNodeRequest {
	s.TenantId = &v
	return s
}

func (s *RerunWorkflowNodeRequest) Validate() error {
	return dara.Validate(s)
}
