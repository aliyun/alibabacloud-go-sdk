// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultAITemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *SetDefaultAITemplateRequest
	GetTemplateId() *string
}

type SetDefaultAITemplateRequest struct {
	// The AI template ID. You can obtain the ID by using one of the following methods:
	//
	// - When you call the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template, the AI template ID is the value of the TemplateId response parameter.
	//
	// - After the AI template is added, call the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation to query the AI template ID, which is the value of the TemplateId response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SetDefaultAITemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultAITemplateRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultAITemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetDefaultAITemplateRequest) SetTemplateId(v string) *SetDefaultAITemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *SetDefaultAITemplateRequest) Validate() error {
	return dara.Validate(s)
}
