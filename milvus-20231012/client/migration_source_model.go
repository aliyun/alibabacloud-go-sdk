// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrationSource interface {
	dara.Model
	String() string
	GoString() string
	SetAuthInfo(v *MigrationSourceAuthInfo) *MigrationSource
	GetAuthInfo() *MigrationSourceAuthInfo
	SetDatabase(v string) *MigrationSource
	GetDatabase() *string
	SetEndpoint(v *MigrationSourceEndpoint) *MigrationSource
	GetEndpoint() *MigrationSourceEndpoint
}

type MigrationSource struct {
	AuthInfo *MigrationSourceAuthInfo `json:"authInfo,omitempty" xml:"authInfo,omitempty" type:"Struct"`
	Database *string                  `json:"database,omitempty" xml:"database,omitempty"`
	Endpoint *MigrationSourceEndpoint `json:"endpoint,omitempty" xml:"endpoint,omitempty" type:"Struct"`
}

func (s MigrationSource) String() string {
	return dara.Prettify(s)
}

func (s MigrationSource) GoString() string {
	return s.String()
}

func (s *MigrationSource) GetAuthInfo() *MigrationSourceAuthInfo {
	return s.AuthInfo
}

func (s *MigrationSource) GetDatabase() *string {
	return s.Database
}

func (s *MigrationSource) GetEndpoint() *MigrationSourceEndpoint {
	return s.Endpoint
}

func (s *MigrationSource) SetAuthInfo(v *MigrationSourceAuthInfo) *MigrationSource {
	s.AuthInfo = v
	return s
}

func (s *MigrationSource) SetDatabase(v string) *MigrationSource {
	s.Database = &v
	return s
}

func (s *MigrationSource) SetEndpoint(v *MigrationSourceEndpoint) *MigrationSource {
	s.Endpoint = v
	return s
}

func (s *MigrationSource) Validate() error {
	if s.AuthInfo != nil {
		if err := s.AuthInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Endpoint != nil {
		if err := s.Endpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MigrationSourceAuthInfo struct {
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	Token    *string `json:"token,omitempty" xml:"token,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s MigrationSourceAuthInfo) String() string {
	return dara.Prettify(s)
}

func (s MigrationSourceAuthInfo) GoString() string {
	return s.String()
}

func (s *MigrationSourceAuthInfo) GetAuthType() *string {
	return s.AuthType
}

func (s *MigrationSourceAuthInfo) GetPassword() *string {
	return s.Password
}

func (s *MigrationSourceAuthInfo) GetToken() *string {
	return s.Token
}

func (s *MigrationSourceAuthInfo) GetUsername() *string {
	return s.Username
}

func (s *MigrationSourceAuthInfo) SetAuthType(v string) *MigrationSourceAuthInfo {
	s.AuthType = &v
	return s
}

func (s *MigrationSourceAuthInfo) SetPassword(v string) *MigrationSourceAuthInfo {
	s.Password = &v
	return s
}

func (s *MigrationSourceAuthInfo) SetToken(v string) *MigrationSourceAuthInfo {
	s.Token = &v
	return s
}

func (s *MigrationSourceAuthInfo) SetUsername(v string) *MigrationSourceAuthInfo {
	s.Username = &v
	return s
}

func (s *MigrationSourceAuthInfo) Validate() error {
	return dara.Validate(s)
}

type MigrationSourceEndpoint struct {
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Port     *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s MigrationSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s MigrationSourceEndpoint) GoString() string {
	return s.String()
}

func (s *MigrationSourceEndpoint) GetEndpoint() *string {
	return s.Endpoint
}

func (s *MigrationSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *MigrationSourceEndpoint) SetEndpoint(v string) *MigrationSourceEndpoint {
	s.Endpoint = &v
	return s
}

func (s *MigrationSourceEndpoint) SetPort(v string) *MigrationSourceEndpoint {
	s.Port = &v
	return s
}

func (s *MigrationSourceEndpoint) Validate() error {
	return dara.Validate(s)
}
