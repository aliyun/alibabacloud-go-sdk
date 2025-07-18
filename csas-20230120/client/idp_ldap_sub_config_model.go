// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpLdapSubConfig interface {
	dara.Model
	String() string
	GoString() string
	SetBaseDn(v string) *IdpLdapSubConfig
	GetBaseDn() *string
	SetGroupBaseDn(v string) *IdpLdapSubConfig
	GetGroupBaseDn() *string
	SetGroupFilter(v string) *IdpLdapSubConfig
	GetGroupFilter() *string
	SetGroupMembershipAttr(v string) *IdpLdapSubConfig
	GetGroupMembershipAttr() *string
	SetGroupNameAttr(v string) *IdpLdapSubConfig
	GetGroupNameAttr() *string
	SetLoginNameAttr(v string) *IdpLdapSubConfig
	GetLoginNameAttr() *string
	SetServerAddr(v string) *IdpLdapSubConfig
	GetServerAddr() *string
	SetServerPort(v string) *IdpLdapSubConfig
	GetServerPort() *string
	SetServerType(v string) *IdpLdapSubConfig
	GetServerType() *string
	SetUseSsl(v bool) *IdpLdapSubConfig
	GetUseSsl() *bool
	SetUserDn(v string) *IdpLdapSubConfig
	GetUserDn() *string
	SetUserFilter(v string) *IdpLdapSubConfig
	GetUserFilter() *string
	SetUserPassword(v string) *IdpLdapSubConfig
	GetUserPassword() *string
}

type IdpLdapSubConfig struct {
	BaseDn              *string `json:"BaseDn,omitempty" xml:"BaseDn,omitempty"`
	GroupBaseDn         *string `json:"GroupBaseDn,omitempty" xml:"GroupBaseDn,omitempty"`
	GroupFilter         *string `json:"GroupFilter,omitempty" xml:"GroupFilter,omitempty"`
	GroupMembershipAttr *string `json:"GroupMembershipAttr,omitempty" xml:"GroupMembershipAttr,omitempty"`
	GroupNameAttr       *string `json:"GroupNameAttr,omitempty" xml:"GroupNameAttr,omitempty"`
	LoginNameAttr       *string `json:"LoginNameAttr,omitempty" xml:"LoginNameAttr,omitempty"`
	ServerAddr          *string `json:"ServerAddr,omitempty" xml:"ServerAddr,omitempty"`
	ServerPort          *string `json:"ServerPort,omitempty" xml:"ServerPort,omitempty"`
	ServerType          *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	UseSsl              *bool   `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	UserDn              *string `json:"UserDn,omitempty" xml:"UserDn,omitempty"`
	UserFilter          *string `json:"UserFilter,omitempty" xml:"UserFilter,omitempty"`
	UserPassword        *string `json:"UserPassword,omitempty" xml:"UserPassword,omitempty"`
}

func (s IdpLdapSubConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpLdapSubConfig) GoString() string {
	return s.String()
}

func (s *IdpLdapSubConfig) GetBaseDn() *string {
	return s.BaseDn
}

func (s *IdpLdapSubConfig) GetGroupBaseDn() *string {
	return s.GroupBaseDn
}

func (s *IdpLdapSubConfig) GetGroupFilter() *string {
	return s.GroupFilter
}

func (s *IdpLdapSubConfig) GetGroupMembershipAttr() *string {
	return s.GroupMembershipAttr
}

func (s *IdpLdapSubConfig) GetGroupNameAttr() *string {
	return s.GroupNameAttr
}

func (s *IdpLdapSubConfig) GetLoginNameAttr() *string {
	return s.LoginNameAttr
}

func (s *IdpLdapSubConfig) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *IdpLdapSubConfig) GetServerPort() *string {
	return s.ServerPort
}

func (s *IdpLdapSubConfig) GetServerType() *string {
	return s.ServerType
}

func (s *IdpLdapSubConfig) GetUseSsl() *bool {
	return s.UseSsl
}

func (s *IdpLdapSubConfig) GetUserDn() *string {
	return s.UserDn
}

func (s *IdpLdapSubConfig) GetUserFilter() *string {
	return s.UserFilter
}

func (s *IdpLdapSubConfig) GetUserPassword() *string {
	return s.UserPassword
}

func (s *IdpLdapSubConfig) SetBaseDn(v string) *IdpLdapSubConfig {
	s.BaseDn = &v
	return s
}

func (s *IdpLdapSubConfig) SetGroupBaseDn(v string) *IdpLdapSubConfig {
	s.GroupBaseDn = &v
	return s
}

func (s *IdpLdapSubConfig) SetGroupFilter(v string) *IdpLdapSubConfig {
	s.GroupFilter = &v
	return s
}

func (s *IdpLdapSubConfig) SetGroupMembershipAttr(v string) *IdpLdapSubConfig {
	s.GroupMembershipAttr = &v
	return s
}

func (s *IdpLdapSubConfig) SetGroupNameAttr(v string) *IdpLdapSubConfig {
	s.GroupNameAttr = &v
	return s
}

func (s *IdpLdapSubConfig) SetLoginNameAttr(v string) *IdpLdapSubConfig {
	s.LoginNameAttr = &v
	return s
}

func (s *IdpLdapSubConfig) SetServerAddr(v string) *IdpLdapSubConfig {
	s.ServerAddr = &v
	return s
}

func (s *IdpLdapSubConfig) SetServerPort(v string) *IdpLdapSubConfig {
	s.ServerPort = &v
	return s
}

func (s *IdpLdapSubConfig) SetServerType(v string) *IdpLdapSubConfig {
	s.ServerType = &v
	return s
}

func (s *IdpLdapSubConfig) SetUseSsl(v bool) *IdpLdapSubConfig {
	s.UseSsl = &v
	return s
}

func (s *IdpLdapSubConfig) SetUserDn(v string) *IdpLdapSubConfig {
	s.UserDn = &v
	return s
}

func (s *IdpLdapSubConfig) SetUserFilter(v string) *IdpLdapSubConfig {
	s.UserFilter = &v
	return s
}

func (s *IdpLdapSubConfig) SetUserPassword(v string) *IdpLdapSubConfig {
	s.UserPassword = &v
	return s
}

func (s *IdpLdapSubConfig) Validate() error {
	return dara.Validate(s)
}
