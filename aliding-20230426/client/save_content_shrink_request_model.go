// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentsShrink(v string) *SaveContentShrinkRequest
	GetContentsShrink() *string
	SetDdFrom(v string) *SaveContentShrinkRequest
	GetDdFrom() *string
	SetTemplateId(v string) *SaveContentShrinkRequest
	GetTemplateId() *string
	SetTenantContextShrink(v string) *SaveContentShrinkRequest
	GetTenantContextShrink() *string
}

type SaveContentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// []
	ContentsShrink *string `json:"Contents,omitempty" xml:"Contents,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client
	DdFrom *string `json:"DdFrom,omitempty" xml:"DdFrom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sdfafdsfsafdfsaf
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s SaveContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveContentShrinkRequest) GetContentsShrink() *string {
	return s.ContentsShrink
}

func (s *SaveContentShrinkRequest) GetDdFrom() *string {
	return s.DdFrom
}

func (s *SaveContentShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SaveContentShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SaveContentShrinkRequest) SetContentsShrink(v string) *SaveContentShrinkRequest {
	s.ContentsShrink = &v
	return s
}

func (s *SaveContentShrinkRequest) SetDdFrom(v string) *SaveContentShrinkRequest {
	s.DdFrom = &v
	return s
}

func (s *SaveContentShrinkRequest) SetTemplateId(v string) *SaveContentShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SaveContentShrinkRequest) SetTenantContextShrink(v string) *SaveContentShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SaveContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
