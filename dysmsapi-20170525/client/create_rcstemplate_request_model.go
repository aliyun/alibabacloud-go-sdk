// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelatedSignNames(v string) *CreateRCSTemplateRequest
	GetRelatedSignNames() *string
	SetTemplateContent(v string) *CreateRCSTemplateRequest
	GetTemplateContent() *string
	SetTemplateFormat(v string) *CreateRCSTemplateRequest
	GetTemplateFormat() *string
	SetTemplateMenu(v string) *CreateRCSTemplateRequest
	GetTemplateMenu() *string
	SetTemplateName(v string) *CreateRCSTemplateRequest
	GetTemplateName() *string
	SetTemplateRule(v string) *CreateRCSTemplateRequest
	GetTemplateRule() *string
	SetTemplateType(v int64) *CreateRCSTemplateRequest
	GetTemplateType() *int64
}

type CreateRCSTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	RelatedSignNames *string `json:"RelatedSignNames,omitempty" xml:"RelatedSignNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	TemplateMenu *string `json:"TemplateMenu,omitempty" xml:"TemplateMenu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 示例值示例值
	TemplateRule *string `json:"TemplateRule,omitempty" xml:"TemplateRule,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateRCSTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateRCSTemplateRequest) GetRelatedSignNames() *string {
	return s.RelatedSignNames
}

func (s *CreateRCSTemplateRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *CreateRCSTemplateRequest) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *CreateRCSTemplateRequest) GetTemplateMenu() *string {
	return s.TemplateMenu
}

func (s *CreateRCSTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateRCSTemplateRequest) GetTemplateRule() *string {
	return s.TemplateRule
}

func (s *CreateRCSTemplateRequest) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *CreateRCSTemplateRequest) SetRelatedSignNames(v string) *CreateRCSTemplateRequest {
	s.RelatedSignNames = &v
	return s
}

func (s *CreateRCSTemplateRequest) SetTemplateContent(v string) *CreateRCSTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateRCSTemplateRequest) SetTemplateFormat(v string) *CreateRCSTemplateRequest {
	s.TemplateFormat = &v
	return s
}

func (s *CreateRCSTemplateRequest) SetTemplateMenu(v string) *CreateRCSTemplateRequest {
	s.TemplateMenu = &v
	return s
}

func (s *CreateRCSTemplateRequest) SetTemplateName(v string) *CreateRCSTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateRCSTemplateRequest) SetTemplateRule(v string) *CreateRCSTemplateRequest {
	s.TemplateRule = &v
	return s
}

func (s *CreateRCSTemplateRequest) SetTemplateType(v int64) *CreateRCSTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateRCSTemplateRequest) Validate() error {
	return dara.Validate(s)
}
