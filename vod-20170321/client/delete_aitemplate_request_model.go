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
	// The ID of the AI template. You can use one of the following methods to obtain the ID of the AI template:
	//
	// 	- Call the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template if no AI template exists. The value of TemplateId in the response is the ID of the AI template.
	//
	// 	- Call the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation if the template already exists. The value of TemplateId in the response is the ID of the AI template.
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
