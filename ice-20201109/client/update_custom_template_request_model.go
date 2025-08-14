// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateCustomTemplateRequest
	GetName() *string
	SetTemplateConfig(v string) *UpdateCustomTemplateRequest
	GetTemplateConfig() *string
	SetTemplateId(v string) *UpdateCustomTemplateRequest
	GetTemplateId() *string
}

type UpdateCustomTemplateRequest struct {
	// The template name.
	//
	// example:
	//
	// test-template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The [template parameters](https://help.aliyun.com/document_detail/448291.html).
	//
	// example:
	//
	// {"param": "sample"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCustomTemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *UpdateCustomTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateCustomTemplateRequest) SetName(v string) *UpdateCustomTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomTemplateRequest) SetTemplateConfig(v string) *UpdateCustomTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *UpdateCustomTemplateRequest) SetTemplateId(v string) *UpdateCustomTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
