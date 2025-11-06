// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCodeServerAuth(v string) *ServiceConfig
	GetCodeServerAuth() *string
	SetCodeServerPassword(v string) *ServiceConfig
	GetCodeServerPassword() *string
	SetJupyterServerAuth(v string) *ServiceConfig
	GetJupyterServerAuth() *string
	SetJupyterServerPassword(v string) *ServiceConfig
	GetJupyterServerPassword() *string
}

type ServiceConfig struct {
	CodeServerAuth        *string `json:"CodeServerAuth,omitempty" xml:"CodeServerAuth,omitempty"`
	CodeServerPassword    *string `json:"CodeServerPassword,omitempty" xml:"CodeServerPassword,omitempty"`
	JupyterServerAuth     *string `json:"JupyterServerAuth,omitempty" xml:"JupyterServerAuth,omitempty"`
	JupyterServerPassword *string `json:"JupyterServerPassword,omitempty" xml:"JupyterServerPassword,omitempty"`
}

func (s ServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s ServiceConfig) GoString() string {
	return s.String()
}

func (s *ServiceConfig) GetCodeServerAuth() *string {
	return s.CodeServerAuth
}

func (s *ServiceConfig) GetCodeServerPassword() *string {
	return s.CodeServerPassword
}

func (s *ServiceConfig) GetJupyterServerAuth() *string {
	return s.JupyterServerAuth
}

func (s *ServiceConfig) GetJupyterServerPassword() *string {
	return s.JupyterServerPassword
}

func (s *ServiceConfig) SetCodeServerAuth(v string) *ServiceConfig {
	s.CodeServerAuth = &v
	return s
}

func (s *ServiceConfig) SetCodeServerPassword(v string) *ServiceConfig {
	s.CodeServerPassword = &v
	return s
}

func (s *ServiceConfig) SetJupyterServerAuth(v string) *ServiceConfig {
	s.JupyterServerAuth = &v
	return s
}

func (s *ServiceConfig) SetJupyterServerPassword(v string) *ServiceConfig {
	s.JupyterServerPassword = &v
	return s
}

func (s *ServiceConfig) Validate() error {
	return dara.Validate(s)
}
