// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelatedWorkspacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeRecent(v bool) *GetRelatedWorkspacesShrinkRequest
	GetIncludeRecent() *bool
	SetTenantContextShrink(v string) *GetRelatedWorkspacesShrinkRequest
	GetTenantContextShrink() *string
}

type GetRelatedWorkspacesShrinkRequest struct {
	// example:
	//
	// true
	IncludeRecent       *bool   `json:"IncludeRecent,omitempty" xml:"IncludeRecent,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetRelatedWorkspacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesShrinkRequest) GetIncludeRecent() *bool {
	return s.IncludeRecent
}

func (s *GetRelatedWorkspacesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetRelatedWorkspacesShrinkRequest) SetIncludeRecent(v bool) *GetRelatedWorkspacesShrinkRequest {
	s.IncludeRecent = &v
	return s
}

func (s *GetRelatedWorkspacesShrinkRequest) SetTenantContextShrink(v string) *GetRelatedWorkspacesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetRelatedWorkspacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
