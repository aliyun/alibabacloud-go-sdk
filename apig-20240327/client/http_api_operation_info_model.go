// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiOperationInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *AuthConfig) *HttpApiOperationInfo
	GetAuthConfig() *AuthConfig
	SetCreateTimestamp(v int64) *HttpApiOperationInfo
	GetCreateTimestamp() *int64
	SetDeployConfigs(v []*HttpApiDeployConfig) *HttpApiOperationInfo
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *HttpApiOperationInfo
	GetDescription() *string
	SetEnableAuth(v bool) *HttpApiOperationInfo
	GetEnableAuth() *bool
	SetMethod(v string) *HttpApiOperationInfo
	GetMethod() *string
	SetMock(v *HttpApiMockContract) *HttpApiOperationInfo
	GetMock() *HttpApiMockContract
	SetName(v string) *HttpApiOperationInfo
	GetName() *string
	SetOperationId(v string) *HttpApiOperationInfo
	GetOperationId() *string
	SetPath(v string) *HttpApiOperationInfo
	GetPath() *string
	SetRequest(v *HttpApiRequestContract) *HttpApiOperationInfo
	GetRequest() *HttpApiRequestContract
	SetResponse(v *HttpApiResponseContract) *HttpApiOperationInfo
	GetResponse() *HttpApiResponseContract
	SetStatus(v string) *HttpApiOperationInfo
	GetStatus() *string
}

type HttpApiOperationInfo struct {
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64                 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	DeployConfigs   []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 获取用户信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// example:
	//
	// GET
	Method *string              `json:"method,omitempty" xml:"method,omitempty"`
	Mock   *HttpApiMockContract `json:"mock,omitempty" xml:"mock,omitempty"`
	// example:
	//
	// GetUserInfo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// op-xxx
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// example:
	//
	// /user/123
	Path     *string                  `json:"path,omitempty" xml:"path,omitempty"`
	Request  *HttpApiRequestContract  `json:"request,omitempty" xml:"request,omitempty"`
	Response *HttpApiResponseContract `json:"response,omitempty" xml:"response,omitempty"`
	// example:
	//
	// Deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s HttpApiOperationInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiOperationInfo) GoString() string {
	return s.String()
}

func (s *HttpApiOperationInfo) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *HttpApiOperationInfo) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *HttpApiOperationInfo) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *HttpApiOperationInfo) GetDescription() *string {
	return s.Description
}

func (s *HttpApiOperationInfo) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *HttpApiOperationInfo) GetMethod() *string {
	return s.Method
}

func (s *HttpApiOperationInfo) GetMock() *HttpApiMockContract {
	return s.Mock
}

func (s *HttpApiOperationInfo) GetName() *string {
	return s.Name
}

func (s *HttpApiOperationInfo) GetOperationId() *string {
	return s.OperationId
}

func (s *HttpApiOperationInfo) GetPath() *string {
	return s.Path
}

func (s *HttpApiOperationInfo) GetRequest() *HttpApiRequestContract {
	return s.Request
}

func (s *HttpApiOperationInfo) GetResponse() *HttpApiResponseContract {
	return s.Response
}

func (s *HttpApiOperationInfo) GetStatus() *string {
	return s.Status
}

func (s *HttpApiOperationInfo) SetAuthConfig(v *AuthConfig) *HttpApiOperationInfo {
	s.AuthConfig = v
	return s
}

func (s *HttpApiOperationInfo) SetCreateTimestamp(v int64) *HttpApiOperationInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *HttpApiOperationInfo) SetDeployConfigs(v []*HttpApiDeployConfig) *HttpApiOperationInfo {
	s.DeployConfigs = v
	return s
}

func (s *HttpApiOperationInfo) SetDescription(v string) *HttpApiOperationInfo {
	s.Description = &v
	return s
}

func (s *HttpApiOperationInfo) SetEnableAuth(v bool) *HttpApiOperationInfo {
	s.EnableAuth = &v
	return s
}

func (s *HttpApiOperationInfo) SetMethod(v string) *HttpApiOperationInfo {
	s.Method = &v
	return s
}

func (s *HttpApiOperationInfo) SetMock(v *HttpApiMockContract) *HttpApiOperationInfo {
	s.Mock = v
	return s
}

func (s *HttpApiOperationInfo) SetName(v string) *HttpApiOperationInfo {
	s.Name = &v
	return s
}

func (s *HttpApiOperationInfo) SetOperationId(v string) *HttpApiOperationInfo {
	s.OperationId = &v
	return s
}

func (s *HttpApiOperationInfo) SetPath(v string) *HttpApiOperationInfo {
	s.Path = &v
	return s
}

func (s *HttpApiOperationInfo) SetRequest(v *HttpApiRequestContract) *HttpApiOperationInfo {
	s.Request = v
	return s
}

func (s *HttpApiOperationInfo) SetResponse(v *HttpApiResponseContract) *HttpApiOperationInfo {
	s.Response = v
	return s
}

func (s *HttpApiOperationInfo) SetStatus(v string) *HttpApiOperationInfo {
	s.Status = &v
	return s
}

func (s *HttpApiOperationInfo) Validate() error {
	return dara.Validate(s)
}
