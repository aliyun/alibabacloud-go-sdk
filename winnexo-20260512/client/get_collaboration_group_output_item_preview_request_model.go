// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCollaborationGroupOutputItemPreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetCollaborationGroupId() *string
	SetItemId(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetItemId() *string
	SetTenantId(v string) *GetCollaborationGroupOutputItemPreviewRequest
	GetTenantId() *string
}

type GetCollaborationGroupOutputItemPreviewRequest struct {
	// 协作空间ID，必填，长度1至64。调用者必须是该租户下该空间的有效成员。可从ListCollaborationGroups返回的items中获取。
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 对话生成文件的产出明细ID。取GetNodeDetail返回的artifacts中ref.kind为USER_OUTPUT_ITEM的ref.id；不是artifactId。
	//
	// This parameter is required.
	//
	// example:
	//
	// 51469466
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 业务租户ID，可选。必须属于当前调用身份的可访问租户；无权访问时拒绝请求。未传时按现有鉴权规则解析租户，多租户用户无法唯一确定时需显式指定。
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

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetItemId() *string {
	return s.ItemId
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetCollaborationGroupId(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetItemId(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.ItemId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) SetTenantId(v string) *GetCollaborationGroupOutputItemPreviewRequest {
	s.TenantId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewRequest) Validate() error {
	return dara.Validate(s)
}
