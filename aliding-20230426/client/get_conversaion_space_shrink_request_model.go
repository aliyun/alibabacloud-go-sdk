// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversaionSpaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *GetConversaionSpaceShrinkRequest
	GetOpenConversationId() *string
	SetTenantContextShrink(v string) *GetConversaionSpaceShrinkRequest
	GetTenantContextShrink() *string
}

type GetConversaionSpaceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidB8Pzg****FIWPv2PMA==
	OpenConversationId  *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetConversaionSpaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceShrinkRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetConversaionSpaceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetConversaionSpaceShrinkRequest) SetOpenConversationId(v string) *GetConversaionSpaceShrinkRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetConversaionSpaceShrinkRequest) SetTenantContextShrink(v string) *GetConversaionSpaceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetConversaionSpaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
