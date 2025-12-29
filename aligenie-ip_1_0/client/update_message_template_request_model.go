// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateDetail(v string) *UpdateMessageTemplateRequest
	GetTemplateDetail() *string
	SetTemplateId(v int64) *UpdateMessageTemplateRequest
	GetTemplateId() *int64
	SetTemplateName(v string) *UpdateMessageTemplateRequest
	GetTemplateName() *string
}

type UpdateMessageTemplateRequest struct {
	TemplateDetail *string `json:"TemplateDetail,omitempty" xml:"TemplateDetail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123123
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateMessageTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageTemplateRequest) GetTemplateDetail() *string {
	return s.TemplateDetail
}

func (s *UpdateMessageTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpdateMessageTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateMessageTemplateRequest) SetTemplateDetail(v string) *UpdateMessageTemplateRequest {
	s.TemplateDetail = &v
	return s
}

func (s *UpdateMessageTemplateRequest) SetTemplateId(v int64) *UpdateMessageTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateMessageTemplateRequest) SetTemplateName(v string) *UpdateMessageTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateMessageTemplateRequest) Validate() error {
	return dara.Validate(s)
}
