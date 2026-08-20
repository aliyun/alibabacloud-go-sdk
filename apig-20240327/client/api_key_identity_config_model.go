// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiKeyIdentityConfig interface {
	dara.Model
	String() string
	GoString() string
	SetApikeySource(v *ApiKeyIdentityConfigApikeySource) *ApiKeyIdentityConfig
	GetApikeySource() *ApiKeyIdentityConfigApikeySource
	SetApikeySources(v []*ApiKeyIdentityConfigApikeySources) *ApiKeyIdentityConfig
	GetApikeySources() []*ApiKeyIdentityConfigApikeySources
	SetCredentials(v []*ApiKeyIdentityConfigCredentials) *ApiKeyIdentityConfig
	GetCredentials() []*ApiKeyIdentityConfigCredentials
	SetType(v string) *ApiKeyIdentityConfig
	GetType() *string
}

type ApiKeyIdentityConfig struct {
	// The API key source configuration.
	ApikeySource *ApiKeyIdentityConfigApikeySource `json:"apikeySource,omitempty" xml:"apikeySource,omitempty" type:"Struct"`
	// The complete set of API key credential sources. The set contains one to three items. Multiple sources are applicable only to the AI gateway Header mode. Query String and non-AI gateway allow only a single source. If submitted together with apikeySource, the latter must be consistent with the compatible projection.
	ApikeySources []*ApiKeyIdentityConfigApikeySources `json:"apikeySources,omitempty" xml:"apikeySources,omitempty" type:"Repeated"`
	// The list of credentials.
	Credentials []*ApiKeyIdentityConfigCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	// The type.
	//
	// example:
	//
	// Apikey
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApiKeyIdentityConfig) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyIdentityConfig) GoString() string {
	return s.String()
}

func (s *ApiKeyIdentityConfig) GetApikeySource() *ApiKeyIdentityConfigApikeySource {
	return s.ApikeySource
}

func (s *ApiKeyIdentityConfig) GetApikeySources() []*ApiKeyIdentityConfigApikeySources {
	return s.ApikeySources
}

func (s *ApiKeyIdentityConfig) GetCredentials() []*ApiKeyIdentityConfigCredentials {
	return s.Credentials
}

func (s *ApiKeyIdentityConfig) GetType() *string {
	return s.Type
}

func (s *ApiKeyIdentityConfig) SetApikeySource(v *ApiKeyIdentityConfigApikeySource) *ApiKeyIdentityConfig {
	s.ApikeySource = v
	return s
}

func (s *ApiKeyIdentityConfig) SetApikeySources(v []*ApiKeyIdentityConfigApikeySources) *ApiKeyIdentityConfig {
	s.ApikeySources = v
	return s
}

func (s *ApiKeyIdentityConfig) SetCredentials(v []*ApiKeyIdentityConfigCredentials) *ApiKeyIdentityConfig {
	s.Credentials = v
	return s
}

func (s *ApiKeyIdentityConfig) SetType(v string) *ApiKeyIdentityConfig {
	s.Type = &v
	return s
}

func (s *ApiKeyIdentityConfig) Validate() error {
	if s.ApikeySource != nil {
		if err := s.ApikeySource.Validate(); err != nil {
			return err
		}
	}
	if s.ApikeySources != nil {
		for _, item := range s.ApikeySources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Credentials != nil {
		for _, item := range s.Credentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApiKeyIdentityConfigApikeySource struct {
	// The API key source.
	//
	// example:
	//
	// Default
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The API key value.
	//
	// example:
	//
	// xxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ApiKeyIdentityConfigApikeySource) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyIdentityConfigApikeySource) GoString() string {
	return s.String()
}

func (s *ApiKeyIdentityConfigApikeySource) GetSource() *string {
	return s.Source
}

func (s *ApiKeyIdentityConfigApikeySource) GetValue() *string {
	return s.Value
}

func (s *ApiKeyIdentityConfigApikeySource) SetSource(v string) *ApiKeyIdentityConfigApikeySource {
	s.Source = &v
	return s
}

func (s *ApiKeyIdentityConfigApikeySource) SetValue(v string) *ApiKeyIdentityConfigApikeySource {
	s.Value = &v
	return s
}

func (s *ApiKeyIdentityConfigApikeySource) Validate() error {
	return dara.Validate(s)
}

type ApiKeyIdentityConfigApikeySources struct {
	// The credential source type.
	//
	// example:
	//
	// Default
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The field name of the HTTP header or query string.
	//
	// example:
	//
	// Authorization
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ApiKeyIdentityConfigApikeySources) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyIdentityConfigApikeySources) GoString() string {
	return s.String()
}

func (s *ApiKeyIdentityConfigApikeySources) GetSource() *string {
	return s.Source
}

func (s *ApiKeyIdentityConfigApikeySources) GetValue() *string {
	return s.Value
}

func (s *ApiKeyIdentityConfigApikeySources) SetSource(v string) *ApiKeyIdentityConfigApikeySources {
	s.Source = &v
	return s
}

func (s *ApiKeyIdentityConfigApikeySources) SetValue(v string) *ApiKeyIdentityConfigApikeySources {
	s.Value = &v
	return s
}

func (s *ApiKeyIdentityConfigApikeySources) Validate() error {
	return dara.Validate(s)
}

type ApiKeyIdentityConfigCredentials struct {
	// The API key configuration.
	//
	// example:
	//
	// xxxxx
	Apikey *string `json:"apikey,omitempty" xml:"apikey,omitempty"`
	// The generation mode.
	//
	// example:
	//
	// System
	GenerateMode *string `json:"generateMode,omitempty" xml:"generateMode,omitempty"`
}

func (s ApiKeyIdentityConfigCredentials) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyIdentityConfigCredentials) GoString() string {
	return s.String()
}

func (s *ApiKeyIdentityConfigCredentials) GetApikey() *string {
	return s.Apikey
}

func (s *ApiKeyIdentityConfigCredentials) GetGenerateMode() *string {
	return s.GenerateMode
}

func (s *ApiKeyIdentityConfigCredentials) SetApikey(v string) *ApiKeyIdentityConfigCredentials {
	s.Apikey = &v
	return s
}

func (s *ApiKeyIdentityConfigCredentials) SetGenerateMode(v string) *ApiKeyIdentityConfigCredentials {
	s.GenerateMode = &v
	return s
}

func (s *ApiKeyIdentityConfigCredentials) Validate() error {
	return dara.Validate(s)
}
