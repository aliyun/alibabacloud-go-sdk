// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAITemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddAITemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *AddAITemplateResponseBody
	GetTemplateId() *string
}

type AddAITemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the AI template.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddAITemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddAITemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAITemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddAITemplateResponseBody) SetRequestId(v string) *AddAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAITemplateResponseBody) SetTemplateId(v string) *AddAITemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *AddAITemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
