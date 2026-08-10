// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKgBySemanticShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SearchKgBySemanticShrinkRequest
	GetOpTenantId() *int64
	SetSearchCommandShrink(v string) *SearchKgBySemanticShrinkRequest
	GetSearchCommandShrink() *string
	SetWorkspaceId(v string) *SearchKgBySemanticShrinkRequest
	GetWorkspaceId() *string
}

type SearchKgBySemanticShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The search command.
	//
	// This parameter is required.
	SearchCommandShrink *string `json:"SearchCommand,omitempty" xml:"SearchCommand,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SearchKgBySemanticShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchKgBySemanticShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchKgBySemanticShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SearchKgBySemanticShrinkRequest) GetSearchCommandShrink() *string {
	return s.SearchCommandShrink
}

func (s *SearchKgBySemanticShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SearchKgBySemanticShrinkRequest) SetOpTenantId(v int64) *SearchKgBySemanticShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SearchKgBySemanticShrinkRequest) SetSearchCommandShrink(v string) *SearchKgBySemanticShrinkRequest {
	s.SearchCommandShrink = &v
	return s
}

func (s *SearchKgBySemanticShrinkRequest) SetWorkspaceId(v string) *SearchKgBySemanticShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SearchKgBySemanticShrinkRequest) Validate() error {
	return dara.Validate(s)
}
