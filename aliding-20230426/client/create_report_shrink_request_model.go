// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentsShrink(v string) *CreateReportShrinkRequest
	GetContentsShrink() *string
	SetDdFrom(v string) *CreateReportShrinkRequest
	GetDdFrom() *string
	SetTemplateId(v string) *CreateReportShrinkRequest
	GetTemplateId() *string
	SetTenantContextShrink(v string) *CreateReportShrinkRequest
	GetTenantContextShrink() *string
	SetToChat(v bool) *CreateReportShrinkRequest
	GetToChat() *bool
	SetToCidsShrink(v string) *CreateReportShrinkRequest
	GetToCidsShrink() *string
	SetToUseridsShrink(v string) *CreateReportShrinkRequest
	GetToUseridsShrink() *string
}

type CreateReportShrinkRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// true
	ToChat *bool `json:"ToChat,omitempty" xml:"ToChat,omitempty"`
	// example:
	//
	// []
	ToCidsShrink *string `json:"ToCids,omitempty" xml:"ToCids,omitempty"`
	// example:
	//
	// [123,456]
	ToUseridsShrink *string `json:"ToUserids,omitempty" xml:"ToUserids,omitempty"`
}

func (s CreateReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateReportShrinkRequest) GetContentsShrink() *string {
	return s.ContentsShrink
}

func (s *CreateReportShrinkRequest) GetDdFrom() *string {
	return s.DdFrom
}

func (s *CreateReportShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateReportShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateReportShrinkRequest) GetToChat() *bool {
	return s.ToChat
}

func (s *CreateReportShrinkRequest) GetToCidsShrink() *string {
	return s.ToCidsShrink
}

func (s *CreateReportShrinkRequest) GetToUseridsShrink() *string {
	return s.ToUseridsShrink
}

func (s *CreateReportShrinkRequest) SetContentsShrink(v string) *CreateReportShrinkRequest {
	s.ContentsShrink = &v
	return s
}

func (s *CreateReportShrinkRequest) SetDdFrom(v string) *CreateReportShrinkRequest {
	s.DdFrom = &v
	return s
}

func (s *CreateReportShrinkRequest) SetTemplateId(v string) *CreateReportShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateReportShrinkRequest) SetTenantContextShrink(v string) *CreateReportShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateReportShrinkRequest) SetToChat(v bool) *CreateReportShrinkRequest {
	s.ToChat = &v
	return s
}

func (s *CreateReportShrinkRequest) SetToCidsShrink(v string) *CreateReportShrinkRequest {
	s.ToCidsShrink = &v
	return s
}

func (s *CreateReportShrinkRequest) SetToUseridsShrink(v string) *CreateReportShrinkRequest {
	s.ToUseridsShrink = &v
	return s
}

func (s *CreateReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
