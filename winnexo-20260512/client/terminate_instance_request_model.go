// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *TerminateInstanceRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *TerminateInstanceRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *TerminateInstanceRequest
	GetInstanceId() *string
	SetOpsUserCode(v string) *TerminateInstanceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *TerminateInstanceRequest
	GetOpsUserId() *string
	SetReason(v string) *TerminateInstanceRequest
	GetReason() *string
	SetTenantId(v string) *TerminateInstanceRequest
	GetTenantId() *string
}

type TerminateInstanceRequest struct {
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
	// The reason for termination.
	//
	// This parameter is required.
	//
	// example:
	//
	// Terminated by user
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s TerminateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceRequest) GoString() string {
	return s.String()
}

func (s *TerminateInstanceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *TerminateInstanceRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *TerminateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TerminateInstanceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *TerminateInstanceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *TerminateInstanceRequest) GetReason() *string {
	return s.Reason
}

func (s *TerminateInstanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TerminateInstanceRequest) SetAppToken(v string) *TerminateInstanceRequest {
	s.AppToken = &v
	return s
}

func (s *TerminateInstanceRequest) SetCollaborationGroupId(v string) *TerminateInstanceRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *TerminateInstanceRequest) SetInstanceId(v string) *TerminateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *TerminateInstanceRequest) SetOpsUserCode(v string) *TerminateInstanceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *TerminateInstanceRequest) SetOpsUserId(v string) *TerminateInstanceRequest {
	s.OpsUserId = &v
	return s
}

func (s *TerminateInstanceRequest) SetReason(v string) *TerminateInstanceRequest {
	s.Reason = &v
	return s
}

func (s *TerminateInstanceRequest) SetTenantId(v string) *TerminateInstanceRequest {
	s.TenantId = &v
	return s
}

func (s *TerminateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
