// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultAITemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultAITemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *SetDefaultAITemplateResponseBody
	GetTemplateId() *string
}

type SetDefaultAITemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8E70E3F8-E2EE-47BC-4677-379D6F28****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the AI template.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SetDefaultAITemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultAITemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultAITemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetDefaultAITemplateResponseBody) SetRequestId(v string) *SetDefaultAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultAITemplateResponseBody) SetTemplateId(v string) *SetDefaultAITemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *SetDefaultAITemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
