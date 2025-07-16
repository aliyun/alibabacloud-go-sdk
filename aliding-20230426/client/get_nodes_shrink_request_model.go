// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeIdsShrink(v string) *GetNodesShrinkRequest
	GetNodeIdsShrink() *string
	SetOptionShrink(v string) *GetNodesShrinkRequest
	GetOptionShrink() *string
	SetTenantContextShrink(v string) *GetNodesShrinkRequest
	GetTenantContextShrink() *string
}

type GetNodesShrinkRequest struct {
	// This parameter is required.
	NodeIdsShrink       *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	OptionShrink        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNodesShrinkRequest) GetNodeIdsShrink() *string {
	return s.NodeIdsShrink
}

func (s *GetNodesShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *GetNodesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetNodesShrinkRequest) SetNodeIdsShrink(v string) *GetNodesShrinkRequest {
	s.NodeIdsShrink = &v
	return s
}

func (s *GetNodesShrinkRequest) SetOptionShrink(v string) *GetNodesShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *GetNodesShrinkRequest) SetTenantContextShrink(v string) *GetNodesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
