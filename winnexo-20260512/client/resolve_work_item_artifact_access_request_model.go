// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveWorkItemArtifactAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ResolveWorkItemArtifactAccessRequest
	GetAppToken() *string
	SetArtifactId(v string) *ResolveWorkItemArtifactAccessRequest
	GetArtifactId() *string
	SetCollaborationGroupId(v string) *ResolveWorkItemArtifactAccessRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *ResolveWorkItemArtifactAccessRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ResolveWorkItemArtifactAccessRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ResolveWorkItemArtifactAccessRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ResolveWorkItemArtifactAccessRequest
	GetWorkItemId() *string
}

type ResolveWorkItemArtifactAccessRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the artifact.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// The ID of the collaboration workspace.
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
	// The ID of the work item.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s ResolveWorkItemArtifactAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s ResolveWorkItemArtifactAccessRequest) GoString() string {
	return s.String()
}

func (s *ResolveWorkItemArtifactAccessRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ResolveWorkItemArtifactAccessRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ResolveWorkItemArtifactAccessRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ResolveWorkItemArtifactAccessRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ResolveWorkItemArtifactAccessRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ResolveWorkItemArtifactAccessRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResolveWorkItemArtifactAccessRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ResolveWorkItemArtifactAccessRequest) SetAppToken(v string) *ResolveWorkItemArtifactAccessRequest {
	s.AppToken = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetArtifactId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.ArtifactId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetCollaborationGroupId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetOpsUserCode(v string) *ResolveWorkItemArtifactAccessRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetOpsUserId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.OpsUserId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetTenantId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.TenantId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetWorkItemId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.WorkItemId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) Validate() error {
	return dara.Validate(s)
}
