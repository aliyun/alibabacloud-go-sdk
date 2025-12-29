// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMessageTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateDetail(v string) *AddMessageTemplateRequest
	GetTemplateDetail() *string
	SetTemplateName(v string) *AddMessageTemplateRequest
	GetTemplateName() *string
}

type AddMessageTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 这是${hotel}的一个测试模板
	TemplateDetail *string `json:"TemplateDetail,omitempty" xml:"TemplateDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s AddMessageTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMessageTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddMessageTemplateRequest) GetTemplateDetail() *string {
	return s.TemplateDetail
}

func (s *AddMessageTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AddMessageTemplateRequest) SetTemplateDetail(v string) *AddMessageTemplateRequest {
	s.TemplateDetail = &v
	return s
}

func (s *AddMessageTemplateRequest) SetTemplateName(v string) *AddMessageTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddMessageTemplateRequest) Validate() error {
	return dara.Validate(s)
}
