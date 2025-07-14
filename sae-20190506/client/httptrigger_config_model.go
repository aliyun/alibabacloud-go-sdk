// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHTTPTriggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAclConfig(v *AclConfig) *HTTPTriggerConfig
	GetAclConfig() *AclConfig
	SetAuthConfig(v interface{}) *HTTPTriggerConfig
	GetAuthConfig() interface{}
	SetAuthType(v string) *HTTPTriggerConfig
	GetAuthType() *string
	SetDisableURLInternet(v bool) *HTTPTriggerConfig
	GetDisableURLInternet() *bool
	SetSafeMode(v bool) *HTTPTriggerConfig
	GetSafeMode() *bool
}

type HTTPTriggerConfig struct {
	AclConfig          *AclConfig  `json:"aclConfig,omitempty" xml:"aclConfig,omitempty"`
	AuthConfig         interface{} `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType           *string     `json:"authType,omitempty" xml:"authType,omitempty"`
	DisableURLInternet *bool       `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	SafeMode           *bool       `json:"safeMode,omitempty" xml:"safeMode,omitempty"`
}

func (s HTTPTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s HTTPTriggerConfig) GoString() string {
	return s.String()
}

func (s *HTTPTriggerConfig) GetAclConfig() *AclConfig {
	return s.AclConfig
}

func (s *HTTPTriggerConfig) GetAuthConfig() interface{} {
	return s.AuthConfig
}

func (s *HTTPTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *HTTPTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *HTTPTriggerConfig) GetSafeMode() *bool {
	return s.SafeMode
}

func (s *HTTPTriggerConfig) SetAclConfig(v *AclConfig) *HTTPTriggerConfig {
	s.AclConfig = v
	return s
}

func (s *HTTPTriggerConfig) SetAuthConfig(v interface{}) *HTTPTriggerConfig {
	s.AuthConfig = v
	return s
}

func (s *HTTPTriggerConfig) SetAuthType(v string) *HTTPTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *HTTPTriggerConfig) SetDisableURLInternet(v bool) *HTTPTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *HTTPTriggerConfig) SetSafeMode(v bool) *HTTPTriggerConfig {
	s.SafeMode = &v
	return s
}

func (s *HTTPTriggerConfig) Validate() error {
	return dara.Validate(s)
}
