// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAITemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateConfig(v string) *UpdateAITemplateRequest
	GetTemplateConfig() *string
	SetTemplateId(v string) *UpdateAITemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateAITemplateRequest
	GetTemplateName() *string
}

type UpdateAITemplateRequest struct {
	// The detailed configuration of the AI template. The value is a JSON string. For more information, see [AITemplateConfig](~~89863#title-vd3-499-o36~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AuditItem":["terrorism","porn"],"AuditRange":["text-title","video"],"AuditContent":["screen"],"AuditAutoBlock":"yes"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The ID of the AI template. You can obtain the template ID by using one of the following methods:
	//
	// - When you call the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template, the AI template ID is the value of TemplateId in the response.
	//
	// - After the AI template is added, call the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation to query the AI template ID, which is the value of TemplateId in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the AI template. The name can be up to 128 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoAITemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateAITemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAITemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAITemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *UpdateAITemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateAITemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateAITemplateRequest) SetTemplateConfig(v string) *UpdateAITemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *UpdateAITemplateRequest) SetTemplateId(v string) *UpdateAITemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateAITemplateRequest) SetTemplateName(v string) *UpdateAITemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateAITemplateRequest) Validate() error {
	return dara.Validate(s)
}
