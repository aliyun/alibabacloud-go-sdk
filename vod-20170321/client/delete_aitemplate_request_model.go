// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAITemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteAITemplateRequest
	GetTemplateId() *string
}

type DeleteAITemplateRequest struct {
	// The ID of the AI template. You can obtain the template ID by using one of the following methods:
	//
	// - When you call the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template, the template ID is the value of the TemplateId parameter in the response.
	//
	// - After the AI template is added, call the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation to query the template ID, which is the value of the TemplateId parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteAITemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAITemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAITemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteAITemplateRequest) SetTemplateId(v string) *DeleteAITemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteAITemplateRequest) Validate() error {
	return dara.Validate(s)
}
