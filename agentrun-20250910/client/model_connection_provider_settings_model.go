// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelConnectionProviderSettings interface {
	dara.Model
	String() string
	GoString() string
	SetBaseUrl(v string) *ModelConnectionProviderSettings
	GetBaseUrl() *string
	SetModelNames(v []*string) *ModelConnectionProviderSettings
	GetModelNames() []*string
}

type ModelConnectionProviderSettings struct {
	// 模型提供商的默认API基础地址
	//
	// example:
	//
	// https://api.openai.com/v1
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// 该连接支持的模型名称列表
	ModelNames []*string `json:"modelNames" xml:"modelNames" type:"Repeated"`
}

func (s ModelConnectionProviderSettings) String() string {
	return dara.Prettify(s)
}

func (s ModelConnectionProviderSettings) GoString() string {
	return s.String()
}

func (s *ModelConnectionProviderSettings) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelConnectionProviderSettings) GetModelNames() []*string {
	return s.ModelNames
}

func (s *ModelConnectionProviderSettings) SetBaseUrl(v string) *ModelConnectionProviderSettings {
	s.BaseUrl = &v
	return s
}

func (s *ModelConnectionProviderSettings) SetModelNames(v []*string) *ModelConnectionProviderSettings {
	s.ModelNames = v
	return s
}

func (s *ModelConnectionProviderSettings) Validate() error {
	return dara.Validate(s)
}
