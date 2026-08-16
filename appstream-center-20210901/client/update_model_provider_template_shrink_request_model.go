// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProviderTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigShrink(v string) *UpdateModelProviderTemplateShrinkRequest
	GetConfigShrink() *string
	SetDescription(v string) *UpdateModelProviderTemplateShrinkRequest
	GetDescription() *string
	SetEnableWuyingProxy(v bool) *UpdateModelProviderTemplateShrinkRequest
	GetEnableWuyingProxy() *bool
	SetName(v string) *UpdateModelProviderTemplateShrinkRequest
	GetName() *string
	SetProviderTemplateId(v string) *UpdateModelProviderTemplateShrinkRequest
	GetProviderTemplateId() *string
}

type UpdateModelProviderTemplateShrinkRequest struct {
	// The model provider configuration.
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The description of the model provider template.
	//
	// example:
	//
	// 阿里云百炼服务商
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the WUYING security gateway proxy.
	//
	// example:
	//
	// true
	EnableWuyingProxy *bool `json:"EnableWuyingProxy,omitempty" xml:"EnableWuyingProxy,omitempty"`
	// The name of the model provider template.
	//
	// example:
	//
	// 阿里云百炼
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the model provider template.
	//
	// This parameter is required.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s UpdateModelProviderTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProviderTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelProviderTemplateShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *UpdateModelProviderTemplateShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelProviderTemplateShrinkRequest) GetEnableWuyingProxy() *bool {
	return s.EnableWuyingProxy
}

func (s *UpdateModelProviderTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateModelProviderTemplateShrinkRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *UpdateModelProviderTemplateShrinkRequest) SetConfigShrink(v string) *UpdateModelProviderTemplateShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *UpdateModelProviderTemplateShrinkRequest) SetDescription(v string) *UpdateModelProviderTemplateShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelProviderTemplateShrinkRequest) SetEnableWuyingProxy(v bool) *UpdateModelProviderTemplateShrinkRequest {
	s.EnableWuyingProxy = &v
	return s
}

func (s *UpdateModelProviderTemplateShrinkRequest) SetName(v string) *UpdateModelProviderTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateModelProviderTemplateShrinkRequest) SetProviderTemplateId(v string) *UpdateModelProviderTemplateShrinkRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *UpdateModelProviderTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
