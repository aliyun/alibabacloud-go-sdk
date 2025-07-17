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
	SetCredentials(v []*ApiKeyIdentityConfigCredentials) *ApiKeyIdentityConfig
	GetCredentials() []*ApiKeyIdentityConfigCredentials
	SetType(v string) *ApiKeyIdentityConfig
	GetType() *string
}

type ApiKeyIdentityConfig struct {
	ApikeySource *ApiKeyIdentityConfigApikeySource  `json:"apikeySource,omitempty" xml:"apikeySource,omitempty" type:"Struct"`
	Credentials  []*ApiKeyIdentityConfigCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	Type         *string                            `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *ApiKeyIdentityConfig) SetCredentials(v []*ApiKeyIdentityConfigCredentials) *ApiKeyIdentityConfig {
	s.Credentials = v
	return s
}

func (s *ApiKeyIdentityConfig) SetType(v string) *ApiKeyIdentityConfig {
	s.Type = &v
	return s
}

func (s *ApiKeyIdentityConfig) Validate() error {
	return dara.Validate(s)
}

type ApiKeyIdentityConfigApikeySource struct {
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
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

type ApiKeyIdentityConfigCredentials struct {
	Apikey       *string `json:"apikey,omitempty" xml:"apikey,omitempty"`
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
