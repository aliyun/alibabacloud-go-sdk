// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetUserShrinkRequest
	GetTenantContextShrink() *string
	SetLanguage(v string) *GetUserShrinkRequest
	GetLanguage() *string
}

type GetUserShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetUserShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetUserShrinkRequest) SetTenantContextShrink(v string) *GetUserShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetUserShrinkRequest) SetLanguage(v string) *GetUserShrinkRequest {
	s.Language = &v
	return s
}

func (s *GetUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
