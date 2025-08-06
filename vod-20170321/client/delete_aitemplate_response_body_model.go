// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAITemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAITemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *DeleteAITemplateResponseBody
	GetTemplateId() *string
}

type DeleteAITemplateResponseBody struct {
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

func (s DeleteAITemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAITemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAITemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteAITemplateResponseBody) SetRequestId(v string) *DeleteAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAITemplateResponseBody) SetTemplateId(v string) *DeleteAITemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DeleteAITemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
