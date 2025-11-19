// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterServiceInput interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialName(v string) *RegisterServiceInput
	GetCredentialName() *string
	SetProtocol(v string) *RegisterServiceInput
	GetProtocol() *string
	SetResourceName(v string) *RegisterServiceInput
	GetResourceName() *string
	SetServiceBackendEndpoint(v string) *RegisterServiceInput
	GetServiceBackendEndpoint() *string
	SetServiceName(v string) *RegisterServiceInput
	GetServiceName() *string
	SetServiceType(v string) *RegisterServiceInput
	GetServiceType() *string
	SetTenantId(v string) *RegisterServiceInput
	GetTenantId() *string
}

type RegisterServiceInput struct {
	// 关联的凭证ID，用于服务认证
	//
	// example:
	//
	// my-credential
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// 服务的协议类型
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// 关联的资源名称
	//
	// example:
	//
	// my-resource
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// 转发的下游服务端点URL，必须是有效的HTTP/HTTPS地址（这里是 FC trigger endpoint）
	//
	// This parameter is required.
	//
	// example:
	//
	// https://123456789.cn-hangzhou.fc.aliyuncs.com/2016-08-15/proxy/my-service/my-function/
	ServiceBackendEndpoint *string `json:"serviceBackendEndpoint,omitempty" xml:"serviceBackendEndpoint,omitempty"`
	// 服务名称，在租户内唯一
	//
	// This parameter is required.
	//
	// example:
	//
	// my-service
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// 服务类型
	//
	// This parameter is required.
	//
	// example:
	//
	// FC
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// 租户ID，用于多租户隔离
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123456
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RegisterServiceInput) String() string {
	return dara.Prettify(s)
}

func (s RegisterServiceInput) GoString() string {
	return s.String()
}

func (s *RegisterServiceInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *RegisterServiceInput) GetProtocol() *string {
	return s.Protocol
}

func (s *RegisterServiceInput) GetResourceName() *string {
	return s.ResourceName
}

func (s *RegisterServiceInput) GetServiceBackendEndpoint() *string {
	return s.ServiceBackendEndpoint
}

func (s *RegisterServiceInput) GetServiceName() *string {
	return s.ServiceName
}

func (s *RegisterServiceInput) GetServiceType() *string {
	return s.ServiceType
}

func (s *RegisterServiceInput) GetTenantId() *string {
	return s.TenantId
}

func (s *RegisterServiceInput) SetCredentialName(v string) *RegisterServiceInput {
	s.CredentialName = &v
	return s
}

func (s *RegisterServiceInput) SetProtocol(v string) *RegisterServiceInput {
	s.Protocol = &v
	return s
}

func (s *RegisterServiceInput) SetResourceName(v string) *RegisterServiceInput {
	s.ResourceName = &v
	return s
}

func (s *RegisterServiceInput) SetServiceBackendEndpoint(v string) *RegisterServiceInput {
	s.ServiceBackendEndpoint = &v
	return s
}

func (s *RegisterServiceInput) SetServiceName(v string) *RegisterServiceInput {
	s.ServiceName = &v
	return s
}

func (s *RegisterServiceInput) SetServiceType(v string) *RegisterServiceInput {
	s.ServiceType = &v
	return s
}

func (s *RegisterServiceInput) SetTenantId(v string) *RegisterServiceInput {
	s.TenantId = &v
	return s
}

func (s *RegisterServiceInput) Validate() error {
	return dara.Validate(s)
}
