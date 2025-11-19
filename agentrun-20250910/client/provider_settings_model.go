// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProviderSettings interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ProviderSettings
	GetApiKey() *string
	SetBaseUrl(v string) *ProviderSettings
	GetBaseUrl() *string
	SetModelNames(v []*string) *ProviderSettings
	GetModelNames() []*string
}

type ProviderSettings struct {
	ApiKey     *string   `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	BaseUrl    *string   `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	ModelNames []*string `json:"modelNames,omitempty" xml:"modelNames,omitempty" type:"Repeated"`
}

func (s ProviderSettings) String() string {
	return dara.Prettify(s)
}

func (s ProviderSettings) GoString() string {
	return s.String()
}

func (s *ProviderSettings) GetApiKey() *string {
	return s.ApiKey
}

func (s *ProviderSettings) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ProviderSettings) GetModelNames() []*string {
	return s.ModelNames
}

func (s *ProviderSettings) SetApiKey(v string) *ProviderSettings {
	s.ApiKey = &v
	return s
}

func (s *ProviderSettings) SetBaseUrl(v string) *ProviderSettings {
	s.BaseUrl = &v
	return s
}

func (s *ProviderSettings) SetModelNames(v []*string) *ProviderSettings {
	s.ModelNames = v
	return s
}

func (s *ProviderSettings) Validate() error {
	return dara.Validate(s)
}
