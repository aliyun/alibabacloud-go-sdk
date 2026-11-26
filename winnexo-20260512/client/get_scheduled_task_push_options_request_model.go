// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskPushOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetScheduledTaskPushOptionsRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetScheduledTaskPushOptionsRequest
	GetCollaborationGroupId() *string
	SetDigitalEmployeeName(v string) *GetScheduledTaskPushOptionsRequest
	GetDigitalEmployeeName() *string
	SetOpsUserCode(v string) *GetScheduledTaskPushOptionsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetScheduledTaskPushOptionsRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetScheduledTaskPushOptionsRequest
	GetTenantId() *string
}

type GetScheduledTaskPushOptionsRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration group, such as cg_101. If specified, a group workspace task is created and the caller must be a valid group member. If left empty, a personal task is created.
	//
	// example:
	//
	// cg_401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The name of the currently active digital employee. This value is empty if not configured.
	//
	// example:
	//
	// exampleDigitalEmployeeName
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
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
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScheduledTaskPushOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskPushOptionsRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskPushOptionsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetScheduledTaskPushOptionsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskPushOptionsRequest) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *GetScheduledTaskPushOptionsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetScheduledTaskPushOptionsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetScheduledTaskPushOptionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskPushOptionsRequest) SetAppToken(v string) *GetScheduledTaskPushOptionsRequest {
	s.AppToken = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) SetCollaborationGroupId(v string) *GetScheduledTaskPushOptionsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) SetDigitalEmployeeName(v string) *GetScheduledTaskPushOptionsRequest {
	s.DigitalEmployeeName = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) SetOpsUserCode(v string) *GetScheduledTaskPushOptionsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) SetOpsUserId(v string) *GetScheduledTaskPushOptionsRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) SetTenantId(v string) *GetScheduledTaskPushOptionsRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskPushOptionsRequest) Validate() error {
	return dara.Validate(s)
}
