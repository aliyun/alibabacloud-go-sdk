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
	// The ID of the AI template. You can use one of the following methods to obtain the ID:
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
