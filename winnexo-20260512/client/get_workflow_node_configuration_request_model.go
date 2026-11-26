// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowNodeConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetWorkflowNodeConfigurationRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetWorkflowNodeConfigurationRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *GetWorkflowNodeConfigurationRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetWorkflowNodeConfigurationRequest
	GetNodeId() *string
	SetOpsUserCode(v string) *GetWorkflowNodeConfigurationRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetWorkflowNodeConfigurationRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetWorkflowNodeConfigurationRequest
	GetTenantId() *string
}

type GetWorkflowNodeConfigurationRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The ID of the current physical plan (planId).
	//
	// This parameter is required.
	//
	// example:
	//
	// inst_2081
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The key of the business node (nodeKey).
	//
	// This parameter is required.
	//
	// example:
	//
	// node_7f3a2c
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
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
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetWorkflowNodeConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowNodeConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowNodeConfigurationRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetWorkflowNodeConfigurationRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkflowNodeConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkflowNodeConfigurationRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetWorkflowNodeConfigurationRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetWorkflowNodeConfigurationRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetWorkflowNodeConfigurationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkflowNodeConfigurationRequest) SetAppToken(v string) *GetWorkflowNodeConfigurationRequest {
	s.AppToken = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetCollaborationGroupId(v string) *GetWorkflowNodeConfigurationRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetInstanceId(v string) *GetWorkflowNodeConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetNodeId(v string) *GetWorkflowNodeConfigurationRequest {
	s.NodeId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetOpsUserCode(v string) *GetWorkflowNodeConfigurationRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetOpsUserId(v string) *GetWorkflowNodeConfigurationRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) SetTenantId(v string) *GetWorkflowNodeConfigurationRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
