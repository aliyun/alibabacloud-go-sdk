// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelProviderAuthorization interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v map[string]*string) *ModelProviderAuthorization
	GetAuthConfig() map[string]*string
	SetType(v string) *ModelProviderAuthorization
	GetType() *string
}

type ModelProviderAuthorization struct {
	AuthConfig map[string]*string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// apiKey
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelProviderAuthorization) String() string {
	return dara.Prettify(s)
}

func (s ModelProviderAuthorization) GoString() string {
	return s.String()
}

func (s *ModelProviderAuthorization) GetAuthConfig() map[string]*string {
	return s.AuthConfig
}

func (s *ModelProviderAuthorization) GetType() *string {
	return s.Type
}

func (s *ModelProviderAuthorization) SetAuthConfig(v map[string]*string) *ModelProviderAuthorization {
	s.AuthConfig = v
	return s
}

func (s *ModelProviderAuthorization) SetType(v string) *ModelProviderAuthorization {
	s.Type = &v
	return s
}

func (s *ModelProviderAuthorization) Validate() error {
	return dara.Validate(s)
}
