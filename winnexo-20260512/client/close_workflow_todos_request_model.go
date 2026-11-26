// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseWorkflowTodosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CloseWorkflowTodosRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *CloseWorkflowTodosRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *CloseWorkflowTodosRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CloseWorkflowTodosRequest
	GetOpsUserId() *string
	SetSourceWorkflowId(v string) *CloseWorkflowTodosRequest
	GetSourceWorkflowId() *string
	SetTenantId(v string) *CloseWorkflowTodosRequest
	GetTenantId() *string
}

type CloseWorkflowTodosRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaborative workspace. The match scope is limited to this workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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
	// The workflow ID at the run level.
	//
	// This parameter is required.
	//
	// example:
	//
	// wf-1
	SourceWorkflowId *string `json:"sourceWorkflowId,omitempty" xml:"sourceWorkflowId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CloseWorkflowTodosRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseWorkflowTodosRequest) GoString() string {
	return s.String()
}

func (s *CloseWorkflowTodosRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CloseWorkflowTodosRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CloseWorkflowTodosRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CloseWorkflowTodosRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CloseWorkflowTodosRequest) GetSourceWorkflowId() *string {
	return s.SourceWorkflowId
}

func (s *CloseWorkflowTodosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CloseWorkflowTodosRequest) SetAppToken(v string) *CloseWorkflowTodosRequest {
	s.AppToken = &v
	return s
}

func (s *CloseWorkflowTodosRequest) SetCollaborationGroupId(v string) *CloseWorkflowTodosRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CloseWorkflowTodosRequest) SetOpsUserCode(v string) *CloseWorkflowTodosRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CloseWorkflowTodosRequest) SetOpsUserId(v string) *CloseWorkflowTodosRequest {
	s.OpsUserId = &v
	return s
}

func (s *CloseWorkflowTodosRequest) SetSourceWorkflowId(v string) *CloseWorkflowTodosRequest {
	s.SourceWorkflowId = &v
	return s
}

func (s *CloseWorkflowTodosRequest) SetTenantId(v string) *CloseWorkflowTodosRequest {
	s.TenantId = &v
	return s
}

func (s *CloseWorkflowTodosRequest) Validate() error {
	return dara.Validate(s)
}
