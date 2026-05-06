// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProviderTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *UpdateModelProviderTemplateRequestConfig) *UpdateModelProviderTemplateRequest
	GetConfig() *UpdateModelProviderTemplateRequestConfig
	SetDescription(v string) *UpdateModelProviderTemplateRequest
	GetDescription() *string
	SetEnableWuyingProxy(v bool) *UpdateModelProviderTemplateRequest
	GetEnableWuyingProxy() *bool
	SetName(v string) *UpdateModelProviderTemplateRequest
	GetName() *string
	SetProviderTemplateId(v string) *UpdateModelProviderTemplateRequest
	GetProviderTemplateId() *string
}

type UpdateModelProviderTemplateRequest struct {
	Config      *UpdateModelProviderTemplateRequestConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	Description *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	EnableWuyingProxy *bool   `json:"EnableWuyingProxy,omitempty" xml:"EnableWuyingProxy,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s UpdateModelProviderTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProviderTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelProviderTemplateRequest) GetConfig() *UpdateModelProviderTemplateRequestConfig {
	return s.Config
}

func (s *UpdateModelProviderTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelProviderTemplateRequest) GetEnableWuyingProxy() *bool {
	return s.EnableWuyingProxy
}

func (s *UpdateModelProviderTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateModelProviderTemplateRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *UpdateModelProviderTemplateRequest) SetConfig(v *UpdateModelProviderTemplateRequestConfig) *UpdateModelProviderTemplateRequest {
	s.Config = v
	return s
}

func (s *UpdateModelProviderTemplateRequest) SetDescription(v string) *UpdateModelProviderTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelProviderTemplateRequest) SetEnableWuyingProxy(v bool) *UpdateModelProviderTemplateRequest {
	s.EnableWuyingProxy = &v
	return s
}

func (s *UpdateModelProviderTemplateRequest) SetName(v string) *UpdateModelProviderTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateModelProviderTemplateRequest) SetProviderTemplateId(v string) *UpdateModelProviderTemplateRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *UpdateModelProviderTemplateRequest) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelProviderTemplateRequestConfig struct {
	// example:
	//
	// sk-xxxxxxxxxxxxxxxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// openai-completions
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
}

func (s UpdateModelProviderTemplateRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProviderTemplateRequestConfig) GoString() string {
	return s.String()
}

func (s *UpdateModelProviderTemplateRequestConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateModelProviderTemplateRequestConfig) GetApiType() *string {
	return s.ApiType
}

func (s *UpdateModelProviderTemplateRequestConfig) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *UpdateModelProviderTemplateRequestConfig) SetApiKey(v string) *UpdateModelProviderTemplateRequestConfig {
	s.ApiKey = &v
	return s
}

func (s *UpdateModelProviderTemplateRequestConfig) SetApiType(v string) *UpdateModelProviderTemplateRequestConfig {
	s.ApiType = &v
	return s
}

func (s *UpdateModelProviderTemplateRequestConfig) SetBaseUrl(v string) *UpdateModelProviderTemplateRequestConfig {
	s.BaseUrl = &v
	return s
}

func (s *UpdateModelProviderTemplateRequestConfig) Validate() error {
	return dara.Validate(s)
}
