// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7CCTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ConfigLayer7CCTemplateRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ConfigLayer7CCTemplateRequest
	GetResourceGroupId() *string
	SetTemplate(v string) *ConfigLayer7CCTemplateRequest
	GetTemplate() *string
}

type ConfigLayer7CCTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ConfigLayer7CCTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7CCTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCTemplateRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigLayer7CCTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigLayer7CCTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *ConfigLayer7CCTemplateRequest) SetDomain(v string) *ConfigLayer7CCTemplateRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetResourceGroupId(v string) *ConfigLayer7CCTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetTemplate(v string) *ConfigLayer7CCTemplateRequest {
	s.Template = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) Validate() error {
	return dara.Validate(s)
}
