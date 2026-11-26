// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetNodeDetailRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetNodeDetailRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *GetNodeDetailRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetNodeDetailRequest
	GetNodeId() *string
	SetOpsUserCode(v string) *GetNodeDetailRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetNodeDetailRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetNodeDetailRequest
	GetTenantId() *string
}

type GetNodeDetailRequest struct {
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
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The physical workflow instance ID (planId).
	//
	// This parameter is required.
	//
	// example:
	//
	// plan_example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The ID of the node (nodeKey).
	//
	// This parameter is required.
	//
	// example:
	//
	// step_1
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

func (s GetNodeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailRequest) GoString() string {
	return s.String()
}

func (s *GetNodeDetailRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetNodeDetailRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetNodeDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNodeDetailRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeDetailRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetNodeDetailRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetNodeDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetNodeDetailRequest) SetAppToken(v string) *GetNodeDetailRequest {
	s.AppToken = &v
	return s
}

func (s *GetNodeDetailRequest) SetCollaborationGroupId(v string) *GetNodeDetailRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetNodeDetailRequest) SetInstanceId(v string) *GetNodeDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNodeDetailRequest) SetNodeId(v string) *GetNodeDetailRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeDetailRequest) SetOpsUserCode(v string) *GetNodeDetailRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetNodeDetailRequest) SetOpsUserId(v string) *GetNodeDetailRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetNodeDetailRequest) SetTenantId(v string) *GetNodeDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetNodeDetailRequest) Validate() error {
	return dara.Validate(s)
}
