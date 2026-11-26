// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ResumeWorkItemRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ResumeWorkItemRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *ResumeWorkItemRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ResumeWorkItemRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ResumeWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ResumeWorkItemRequest
	GetWorkItemId() *string
}

type ResumeWorkItemRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration space ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
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
	// The tenant ID. This is a common parameter. If not specified, the caller\"s default tenant is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The work item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s ResumeWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeWorkItemRequest) GoString() string {
	return s.String()
}

func (s *ResumeWorkItemRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ResumeWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ResumeWorkItemRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ResumeWorkItemRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ResumeWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResumeWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ResumeWorkItemRequest) SetAppToken(v string) *ResumeWorkItemRequest {
	s.AppToken = &v
	return s
}

func (s *ResumeWorkItemRequest) SetCollaborationGroupId(v string) *ResumeWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ResumeWorkItemRequest) SetOpsUserCode(v string) *ResumeWorkItemRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ResumeWorkItemRequest) SetOpsUserId(v string) *ResumeWorkItemRequest {
	s.OpsUserId = &v
	return s
}

func (s *ResumeWorkItemRequest) SetTenantId(v string) *ResumeWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *ResumeWorkItemRequest) SetWorkItemId(v string) *ResumeWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *ResumeWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
