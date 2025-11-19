// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialPublicConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v map[string]*string) *CredentialPublicConfig
	GetAuthConfig() map[string]*string
	SetAuthType(v string) *CredentialPublicConfig
	GetAuthType() *string
	SetHeaderKey(v string) *CredentialPublicConfig
	GetHeaderKey() *string
	SetProvider(v string) *CredentialPublicConfig
	GetProvider() *string
	SetRemoteConfig(v *CredentialPublicConfigRemoteConfig) *CredentialPublicConfig
	GetRemoteConfig() *CredentialPublicConfigRemoteConfig
	SetUsers(v []*CredentialPublicConfigUsers) *CredentialPublicConfig
	GetUsers() []*CredentialPublicConfigUsers
}

type CredentialPublicConfig struct {
	AuthConfig   map[string]*string                  `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType     *string                             `json:"authType,omitempty" xml:"authType,omitempty"`
	HeaderKey    *string                             `json:"headerKey,omitempty" xml:"headerKey,omitempty"`
	Provider     *string                             `json:"provider,omitempty" xml:"provider,omitempty"`
	RemoteConfig *CredentialPublicConfigRemoteConfig `json:"remoteConfig,omitempty" xml:"remoteConfig,omitempty" type:"Struct"`
	Users        []*CredentialPublicConfigUsers      `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s CredentialPublicConfig) String() string {
	return dara.Prettify(s)
}

func (s CredentialPublicConfig) GoString() string {
	return s.String()
}

func (s *CredentialPublicConfig) GetAuthConfig() map[string]*string {
	return s.AuthConfig
}

func (s *CredentialPublicConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *CredentialPublicConfig) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *CredentialPublicConfig) GetProvider() *string {
	return s.Provider
}

func (s *CredentialPublicConfig) GetRemoteConfig() *CredentialPublicConfigRemoteConfig {
	return s.RemoteConfig
}

func (s *CredentialPublicConfig) GetUsers() []*CredentialPublicConfigUsers {
	return s.Users
}

func (s *CredentialPublicConfig) SetAuthConfig(v map[string]*string) *CredentialPublicConfig {
	s.AuthConfig = v
	return s
}

func (s *CredentialPublicConfig) SetAuthType(v string) *CredentialPublicConfig {
	s.AuthType = &v
	return s
}

func (s *CredentialPublicConfig) SetHeaderKey(v string) *CredentialPublicConfig {
	s.HeaderKey = &v
	return s
}

func (s *CredentialPublicConfig) SetProvider(v string) *CredentialPublicConfig {
	s.Provider = &v
	return s
}

func (s *CredentialPublicConfig) SetRemoteConfig(v *CredentialPublicConfigRemoteConfig) *CredentialPublicConfig {
	s.RemoteConfig = v
	return s
}

func (s *CredentialPublicConfig) SetUsers(v []*CredentialPublicConfigUsers) *CredentialPublicConfig {
	s.Users = v
	return s
}

func (s *CredentialPublicConfig) Validate() error {
	if s.RemoteConfig != nil {
		if err := s.RemoteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CredentialPublicConfigRemoteConfig struct {
	Timeout *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Ttl     *int32  `json:"ttl,omitempty" xml:"ttl,omitempty"`
	Uri     *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s CredentialPublicConfigRemoteConfig) String() string {
	return dara.Prettify(s)
}

func (s CredentialPublicConfigRemoteConfig) GoString() string {
	return s.String()
}

func (s *CredentialPublicConfigRemoteConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CredentialPublicConfigRemoteConfig) GetTtl() *int32 {
	return s.Ttl
}

func (s *CredentialPublicConfigRemoteConfig) GetUri() *string {
	return s.Uri
}

func (s *CredentialPublicConfigRemoteConfig) SetTimeout(v int32) *CredentialPublicConfigRemoteConfig {
	s.Timeout = &v
	return s
}

func (s *CredentialPublicConfigRemoteConfig) SetTtl(v int32) *CredentialPublicConfigRemoteConfig {
	s.Ttl = &v
	return s
}

func (s *CredentialPublicConfigRemoteConfig) SetUri(v string) *CredentialPublicConfigRemoteConfig {
	s.Uri = &v
	return s
}

func (s *CredentialPublicConfigRemoteConfig) Validate() error {
	return dara.Validate(s)
}

type CredentialPublicConfigUsers struct {
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CredentialPublicConfigUsers) String() string {
	return dara.Prettify(s)
}

func (s CredentialPublicConfigUsers) GoString() string {
	return s.String()
}

func (s *CredentialPublicConfigUsers) GetPassword() *string {
	return s.Password
}

func (s *CredentialPublicConfigUsers) GetUsername() *string {
	return s.Username
}

func (s *CredentialPublicConfigUsers) SetPassword(v string) *CredentialPublicConfigUsers {
	s.Password = &v
	return s
}

func (s *CredentialPublicConfigUsers) SetUsername(v string) *CredentialPublicConfigUsers {
	s.Username = &v
	return s
}

func (s *CredentialPublicConfigUsers) Validate() error {
	return dara.Validate(s)
}
