// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListInstanceNodesRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ListInstanceNodesRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *ListInstanceNodesRequest
	GetInstanceId() *string
	SetOpsUserCode(v string) *ListInstanceNodesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListInstanceNodesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListInstanceNodesRequest
	GetTenantId() *string
}

type ListInstanceNodesRequest struct {
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

func (s ListInstanceNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceNodesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceNodesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListInstanceNodesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListInstanceNodesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceNodesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListInstanceNodesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListInstanceNodesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListInstanceNodesRequest) SetAppToken(v string) *ListInstanceNodesRequest {
	s.AppToken = &v
	return s
}

func (s *ListInstanceNodesRequest) SetCollaborationGroupId(v string) *ListInstanceNodesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListInstanceNodesRequest) SetInstanceId(v string) *ListInstanceNodesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceNodesRequest) SetOpsUserCode(v string) *ListInstanceNodesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListInstanceNodesRequest) SetOpsUserId(v string) *ListInstanceNodesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListInstanceNodesRequest) SetTenantId(v string) *ListInstanceNodesRequest {
	s.TenantId = &v
	return s
}

func (s *ListInstanceNodesRequest) Validate() error {
	return dara.Validate(s)
}
