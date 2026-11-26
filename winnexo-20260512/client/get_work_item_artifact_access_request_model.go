// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemArtifactAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetWorkItemArtifactAccessRequest
	GetAppToken() *string
	SetArtifactId(v string) *GetWorkItemArtifactAccessRequest
	GetArtifactId() *string
	SetCollaborationGroupId(v string) *GetWorkItemArtifactAccessRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *GetWorkItemArtifactAccessRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetWorkItemArtifactAccessRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetWorkItemArtifactAccessRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemArtifactAccessRequest
	GetWorkItemId() *string
}

type GetWorkItemArtifactAccessRequest struct {
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
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
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

func (s GetWorkItemArtifactAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemArtifactAccessRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemArtifactAccessRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetWorkItemArtifactAccessRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *GetWorkItemArtifactAccessRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemArtifactAccessRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetWorkItemArtifactAccessRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetWorkItemArtifactAccessRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemArtifactAccessRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemArtifactAccessRequest) SetAppToken(v string) *GetWorkItemArtifactAccessRequest {
	s.AppToken = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetArtifactId(v string) *GetWorkItemArtifactAccessRequest {
	s.ArtifactId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetCollaborationGroupId(v string) *GetWorkItemArtifactAccessRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetOpsUserCode(v string) *GetWorkItemArtifactAccessRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetOpsUserId(v string) *GetWorkItemArtifactAccessRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetTenantId(v string) *GetWorkItemArtifactAccessRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetWorkItemId(v string) *GetWorkItemArtifactAccessRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) Validate() error {
	return dara.Validate(s)
}
