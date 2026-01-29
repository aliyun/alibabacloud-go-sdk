// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHTTPTriggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v string) *HTTPTriggerConfig
	GetAuthConfig() *string
	SetAuthType(v string) *HTTPTriggerConfig
	GetAuthType() *string
	SetCorsConfig(v *CORSConfig) *HTTPTriggerConfig
	GetCorsConfig() *CORSConfig
	SetDisableURLInternet(v bool) *HTTPTriggerConfig
	GetDisableURLInternet() *bool
	SetMethods(v []*string) *HTTPTriggerConfig
	GetMethods() []*string
}

type HTTPTriggerConfig struct {
	// example:
	//
	// {"JWKS":{"foo":"bar"},"TokenLookup":"header:Authorization:Bearer,cookie:AuthorizationCookie","ClaimPassBy":"query:uid:uid,header:name:name"}
	AuthConfig *string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// anonymous
	AuthType   *string     `json:"authType,omitempty" xml:"authType,omitempty"`
	CorsConfig *CORSConfig `json:"corsConfig,omitempty" xml:"corsConfig,omitempty"`
	// example:
	//
	// true
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	Methods            []*string `json:"methods" xml:"methods" type:"Repeated"`
}

func (s HTTPTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s HTTPTriggerConfig) GoString() string {
	return s.String()
}

func (s *HTTPTriggerConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *HTTPTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *HTTPTriggerConfig) GetCorsConfig() *CORSConfig {
	return s.CorsConfig
}

func (s *HTTPTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *HTTPTriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *HTTPTriggerConfig) SetAuthConfig(v string) *HTTPTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *HTTPTriggerConfig) SetAuthType(v string) *HTTPTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *HTTPTriggerConfig) SetCorsConfig(v *CORSConfig) *HTTPTriggerConfig {
	s.CorsConfig = v
	return s
}

func (s *HTTPTriggerConfig) SetDisableURLInternet(v bool) *HTTPTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *HTTPTriggerConfig) SetMethods(v []*string) *HTTPTriggerConfig {
	s.Methods = v
	return s
}

func (s *HTTPTriggerConfig) Validate() error {
	if s.CorsConfig != nil {
		if err := s.CorsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
