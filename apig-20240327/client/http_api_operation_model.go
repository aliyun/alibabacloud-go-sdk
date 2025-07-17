// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiOperation interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *AuthConfig) *HttpApiOperation
	GetAuthConfig() *AuthConfig
	SetDeployConfigs(v []*HttpApiDeployConfig) *HttpApiOperation
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *HttpApiOperation
	GetDescription() *string
	SetEnableAuth(v bool) *HttpApiOperation
	GetEnableAuth() *bool
	SetMethod(v string) *HttpApiOperation
	GetMethod() *string
	SetMock(v *HttpApiMockContract) *HttpApiOperation
	GetMock() *HttpApiMockContract
	SetName(v string) *HttpApiOperation
	GetName() *string
	SetPath(v string) *HttpApiOperation
	GetPath() *string
	SetRequest(v *HttpApiRequestContract) *HttpApiOperation
	GetRequest() *HttpApiRequestContract
	SetResponse(v *HttpApiResponseContract) *HttpApiOperation
	GetResponse() *HttpApiResponseContract
}

type HttpApiOperation struct {
	AuthConfig    *AuthConfig            `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 获取用户信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EnableAuth  *bool   `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GET
	Method *string              `json:"method,omitempty" xml:"method,omitempty"`
	Mock   *HttpApiMockContract `json:"mock,omitempty" xml:"mock,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GetUserInfo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /user
	Path     *string                  `json:"path,omitempty" xml:"path,omitempty"`
	Request  *HttpApiRequestContract  `json:"request,omitempty" xml:"request,omitempty"`
	Response *HttpApiResponseContract `json:"response,omitempty" xml:"response,omitempty"`
}

func (s HttpApiOperation) String() string {
	return dara.Prettify(s)
}

func (s HttpApiOperation) GoString() string {
	return s.String()
}

func (s *HttpApiOperation) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *HttpApiOperation) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *HttpApiOperation) GetDescription() *string {
	return s.Description
}

func (s *HttpApiOperation) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *HttpApiOperation) GetMethod() *string {
	return s.Method
}

func (s *HttpApiOperation) GetMock() *HttpApiMockContract {
	return s.Mock
}

func (s *HttpApiOperation) GetName() *string {
	return s.Name
}

func (s *HttpApiOperation) GetPath() *string {
	return s.Path
}

func (s *HttpApiOperation) GetRequest() *HttpApiRequestContract {
	return s.Request
}

func (s *HttpApiOperation) GetResponse() *HttpApiResponseContract {
	return s.Response
}

func (s *HttpApiOperation) SetAuthConfig(v *AuthConfig) *HttpApiOperation {
	s.AuthConfig = v
	return s
}

func (s *HttpApiOperation) SetDeployConfigs(v []*HttpApiDeployConfig) *HttpApiOperation {
	s.DeployConfigs = v
	return s
}

func (s *HttpApiOperation) SetDescription(v string) *HttpApiOperation {
	s.Description = &v
	return s
}

func (s *HttpApiOperation) SetEnableAuth(v bool) *HttpApiOperation {
	s.EnableAuth = &v
	return s
}

func (s *HttpApiOperation) SetMethod(v string) *HttpApiOperation {
	s.Method = &v
	return s
}

func (s *HttpApiOperation) SetMock(v *HttpApiMockContract) *HttpApiOperation {
	s.Mock = v
	return s
}

func (s *HttpApiOperation) SetName(v string) *HttpApiOperation {
	s.Name = &v
	return s
}

func (s *HttpApiOperation) SetPath(v string) *HttpApiOperation {
	s.Path = &v
	return s
}

func (s *HttpApiOperation) SetRequest(v *HttpApiRequestContract) *HttpApiOperation {
	s.Request = v
	return s
}

func (s *HttpApiOperation) SetResponse(v *HttpApiResponseContract) *HttpApiOperation {
	s.Response = v
	return s
}

func (s *HttpApiOperation) Validate() error {
	return dara.Validate(s)
}
