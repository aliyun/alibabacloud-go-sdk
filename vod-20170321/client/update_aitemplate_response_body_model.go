// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAITemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAITemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *UpdateAITemplateResponseBody
	GetTemplateId() *string
}

type UpdateAITemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateAITemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAITemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAITemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateAITemplateResponseBody) SetRequestId(v string) *UpdateAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAITemplateResponseBody) SetTemplateId(v string) *UpdateAITemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UpdateAITemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
