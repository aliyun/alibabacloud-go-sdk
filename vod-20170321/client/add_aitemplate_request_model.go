// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAITemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateConfig(v string) *AddAITemplateRequest
	GetTemplateConfig() *string
	SetTemplateName(v string) *AddAITemplateRequest
	GetTemplateName() *string
	SetTemplateType(v string) *AddAITemplateRequest
	GetTemplateType() *string
}

type AddAITemplateRequest struct {
	// The detailed configurations of the AI template. The value must be a JSON string. For more information, see [AITemplateConfig](~~89863#title-vd3-499-o36~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AuditItem":["terrorism","porn"],"AuditRange":["image-cover","text-title","video"],"AuditContent":["screen"],"AuditAutoBlock":"yes"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The name of the AI template. The name can be up to 128 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// AI-media-test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the AI template. Valid values:
	//
	// 	- **AIMediaAudit**: automated review
	//
	// 	- **AIImage**: smart thumbnail
	//
	// This parameter is required.
	//
	// example:
	//
	// AIMediaAudit
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s AddAITemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAITemplateRequest) GoString() string {
	return s.String()
}

func (s *AddAITemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *AddAITemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AddAITemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *AddAITemplateRequest) SetTemplateConfig(v string) *AddAITemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *AddAITemplateRequest) SetTemplateName(v string) *AddAITemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddAITemplateRequest) SetTemplateType(v string) *AddAITemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *AddAITemplateRequest) Validate() error {
	return dara.Validate(s)
}
