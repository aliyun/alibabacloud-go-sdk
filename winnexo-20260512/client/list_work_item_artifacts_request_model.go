// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemArtifactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListWorkItemArtifactsRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ListWorkItemArtifactsRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *ListWorkItemArtifactsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListWorkItemArtifactsRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListWorkItemArtifactsRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ListWorkItemArtifactsRequest
	GetWorkItemId() *string
}

type ListWorkItemArtifactsRequest struct {
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
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

func (s ListWorkItemArtifactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemArtifactsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListWorkItemArtifactsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkItemArtifactsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListWorkItemArtifactsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListWorkItemArtifactsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkItemArtifactsRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ListWorkItemArtifactsRequest) SetAppToken(v string) *ListWorkItemArtifactsRequest {
	s.AppToken = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) SetCollaborationGroupId(v string) *ListWorkItemArtifactsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) SetOpsUserCode(v string) *ListWorkItemArtifactsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) SetOpsUserId(v string) *ListWorkItemArtifactsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) SetTenantId(v string) *ListWorkItemArtifactsRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) SetWorkItemId(v string) *ListWorkItemArtifactsRequest {
	s.WorkItemId = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) Validate() error {
	return dara.Validate(s)
}
