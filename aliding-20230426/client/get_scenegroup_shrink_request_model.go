// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenegroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *GetScenegroupShrinkRequest
	GetOpenConversationId() *string
	SetTenantContextShrink(v string) *GetScenegroupShrinkRequest
	GetTenantContextShrink() *string
}

type GetScenegroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId  *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetScenegroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScenegroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetScenegroupShrinkRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetScenegroupShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetScenegroupShrinkRequest) SetOpenConversationId(v string) *GetScenegroupShrinkRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetScenegroupShrinkRequest) SetTenantContextShrink(v string) *GetScenegroupShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetScenegroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
