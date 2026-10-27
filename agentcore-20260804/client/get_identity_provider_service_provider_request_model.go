// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderServiceProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProtocol(v string) *GetIdentityProviderServiceProviderRequest
	GetProtocol() *string
}

type GetIdentityProviderServiceProviderRequest struct {
	// 查询参数必填，仅支持精确值 SAML；响应返回 SAML。
	//
	// This parameter is required.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s GetIdentityProviderServiceProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderServiceProviderRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderServiceProviderRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetIdentityProviderServiceProviderRequest) SetProtocol(v string) *GetIdentityProviderServiceProviderRequest {
	s.Protocol = &v
	return s
}

func (s *GetIdentityProviderServiceProviderRequest) Validate() error {
	return dara.Validate(s)
}
