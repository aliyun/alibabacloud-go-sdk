// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCollaborationGroupOutputItemPreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetCollaborationGroupId() *string
	SetItemId(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetItemId() *string
	SetOpsUserCode(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetTenantId() *string
}

type GetCollaborationGroupOutputItemPreviewRequest struct {
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
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The output detail ID. This ID is obtained from the ref.id of a USER_OUTPUT_ITEM node or the space output list.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51469466
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
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
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetCollaborationGroupOutputItemPreviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCollaborationGroupOutputItemPreviewRequest) GoString() string {
	return s.String()
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetItemId() *string {
	return s.ItemId
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetAppToken(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.AppToken = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetCollaborationGroupId(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetItemId(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.ItemId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetOpsUserCode(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetOpsUserId(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetTenantId(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.TenantId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) Validate() error {
	return dara.Validate(s)
}
