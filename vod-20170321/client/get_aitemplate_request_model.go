// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetAITemplateRequest
	GetTemplateId() *string
}

type GetAITemplateRequest struct {
	// The AI template ID. You can obtain the ID by using one of the following methods:
	//
	// - When you call the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template, the AI template ID is the value of the TemplateId parameter in the response.
	//
	// - After the AI template is added, you can call the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation to query the AI template ID, which is the value of the TemplateId parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetAITemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAITemplateRequest) GoString() string {
	return s.String()
}

func (s *GetAITemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAITemplateRequest) SetTemplateId(v string) *GetAITemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetAITemplateRequest) Validate() error {
	return dara.Validate(s)
}
