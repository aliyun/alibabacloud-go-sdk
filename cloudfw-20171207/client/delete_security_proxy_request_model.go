// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteSecurityProxyRequest
	GetLang() *string
	SetProxyId(v string) *DeleteSecurityProxyRequest
	GetProxyId() *string
}

type DeleteSecurityProxyRequest struct {
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// proxy-nat00ab412ef93d4275a6b5
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
}

func (s DeleteSecurityProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityProxyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityProxyRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteSecurityProxyRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DeleteSecurityProxyRequest) SetLang(v string) *DeleteSecurityProxyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteSecurityProxyRequest) SetProxyId(v string) *DeleteSecurityProxyRequest {
	s.ProxyId = &v
	return s
}

func (s *DeleteSecurityProxyRequest) Validate() error {
	return dara.Validate(s)
}
