// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateSecurityProxyRequest
	GetLang() *string
	SetProxyId(v string) *UpdateSecurityProxyRequest
	GetProxyId() *string
	SetProxyName(v string) *UpdateSecurityProxyRequest
	GetProxyName() *string
	SetStrictMode(v int32) *UpdateSecurityProxyRequest
	GetStrictMode() *int32
}

type UpdateSecurityProxyRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// proxy-natfdc73073e031****8e0d
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// proxy_auto_heyuan
	ProxyName *string `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	// example:
	//
	// 1
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
}

func (s UpdateSecurityProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityProxyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityProxyRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateSecurityProxyRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *UpdateSecurityProxyRequest) GetProxyName() *string {
	return s.ProxyName
}

func (s *UpdateSecurityProxyRequest) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *UpdateSecurityProxyRequest) SetLang(v string) *UpdateSecurityProxyRequest {
	s.Lang = &v
	return s
}

func (s *UpdateSecurityProxyRequest) SetProxyId(v string) *UpdateSecurityProxyRequest {
	s.ProxyId = &v
	return s
}

func (s *UpdateSecurityProxyRequest) SetProxyName(v string) *UpdateSecurityProxyRequest {
	s.ProxyName = &v
	return s
}

func (s *UpdateSecurityProxyRequest) SetStrictMode(v int32) *UpdateSecurityProxyRequest {
	s.StrictMode = &v
	return s
}

func (s *UpdateSecurityProxyRequest) Validate() error {
	return dara.Validate(s)
}
