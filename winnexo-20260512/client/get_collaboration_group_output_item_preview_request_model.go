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
