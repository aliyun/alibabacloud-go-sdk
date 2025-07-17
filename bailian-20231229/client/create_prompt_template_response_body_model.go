// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPromptTemplateId(v string) *CreatePromptTemplateResponseBody
	GetPromptTemplateId() *string
	SetRequestId(v string) *CreatePromptTemplateResponseBody
	GetRequestId() *string
}

type CreatePromptTemplateResponseBody struct {
	// example:
	//
	// 6e49109bfeb94a39bb268f4e483ccxxx
	PromptTemplateId *string `json:"promptTemplateId,omitempty" xml:"promptTemplateId,omitempty"`
	// example:
	//
	// FE9B6CBF-47E6-5D76-9C5D-B814DD5ABxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePromptTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePromptTemplateResponseBody) GetPromptTemplateId() *string {
	return s.PromptTemplateId
}

func (s *CreatePromptTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePromptTemplateResponseBody) SetPromptTemplateId(v string) *CreatePromptTemplateResponseBody {
	s.PromptTemplateId = &v
	return s
}

func (s *CreatePromptTemplateResponseBody) SetRequestId(v string) *CreatePromptTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePromptTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
