// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkDomain(v *GetNetworkDomainResponseBodyNetworkDomain) *GetNetworkDomainResponseBody
	GetNetworkDomain() *GetNetworkDomainResponseBodyNetworkDomain
	SetRequestId(v string) *GetNetworkDomainResponseBody
	GetRequestId() *string
}

type GetNetworkDomainResponseBody struct {
	// The detailed information about the network domain.
	NetworkDomain *GetNetworkDomainResponseBodyNetworkDomain `json:"NetworkDomain,omitempty" xml:"NetworkDomain,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05F59944-2E24-595C-B21A-8C9955E60FAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkDomainResponseBody) GetNetworkDomain() *GetNetworkDomainResponseBodyNetworkDomain {
	return s.NetworkDomain
}

func (s *GetNetworkDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkDomainResponseBody) SetNetworkDomain(v *GetNetworkDomainResponseBodyNetworkDomain) *GetNetworkDomainResponseBody {
	s.NetworkDomain = v
	return s
}

func (s *GetNetworkDomainResponseBody) SetRequestId(v string) *GetNetworkDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNetworkDomainResponseBodyNetworkDomain struct {
	// The remarks of the network domain.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Indicates whether the network domain is a built-in network domain.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Default *bool `json:"Default,omitempty" xml:"Default,omitempty"`
	// The network domain ID.
	//
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The name of the network domain.
	//
	// example:
	//
	// SSH Proxy
	NetworkDomainName *string `json:"NetworkDomainName,omitempty" xml:"NetworkDomainName,omitempty"`
	// The connection mode of the network domain. Valid values:
	//
	// 	- Direct
	//
	// 	- Proxy
	//
	// example:
	//
	// Proxy
	NetworkDomainType *string `json:"NetworkDomainType,omitempty" xml:"NetworkDomainType,omitempty"`
	// The information about the proxy servers.
	Proxies []*GetNetworkDomainResponseBodyNetworkDomainProxies `json:"Proxies,omitempty" xml:"Proxies,omitempty" type:"Repeated"`
}

func (s GetNetworkDomainResponseBodyNetworkDomain) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkDomainResponseBodyNetworkDomain) GoString() string {
	return s.String()
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) GetComment() *string {
	return s.Comment
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) GetDefault() *bool {
	return s.Default
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) GetNetworkDomainName() *string {
	return s.NetworkDomainName
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) GetNetworkDomainType() *string {
	return s.NetworkDomainType
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) GetProxies() []*GetNetworkDomainResponseBodyNetworkDomainProxies {
	return s.Proxies
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) SetComment(v string) *GetNetworkDomainResponseBodyNetworkDomain {
	s.Comment = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) SetDefault(v bool) *GetNetworkDomainResponseBodyNetworkDomain {
	s.Default = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) SetNetworkDomainId(v string) *GetNetworkDomainResponseBodyNetworkDomain {
	s.NetworkDomainId = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) SetNetworkDomainName(v string) *GetNetworkDomainResponseBodyNetworkDomain {
	s.NetworkDomainName = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) SetNetworkDomainType(v string) *GetNetworkDomainResponseBodyNetworkDomain {
	s.NetworkDomainType = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) SetProxies(v []*GetNetworkDomainResponseBodyNetworkDomainProxies) *GetNetworkDomainResponseBodyNetworkDomain {
	s.Proxies = v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomain) Validate() error {
	return dara.Validate(s)
}

type GetNetworkDomainResponseBodyNetworkDomainProxies struct {
	// The IP address of the proxy server.
	//
	// example:
	//
	// ``47.102.**.**``
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Indicates whether the proxy server has a password. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	HasPassword *bool `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The node type of the proxy server. Valid values:
	//
	// - **Master**: primary proxy server.
	//
	// - **Slave**: secondary proxy server.
	//
	// example:
	//
	// Master
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The port of the proxy server.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The status of the proxy server.
	//
	// - **Available**
	//
	// - **Unavailable**
	//
	// example:
	//
	// Unavailable
	ProxyState *string `json:"ProxyState,omitempty" xml:"ProxyState,omitempty"`
	// The error code that indicates the status of the proxy server.
	//
	// - **CHECK_PWD_FAILED**: The password is invalid.
	//
	// - **CHECK_PWD_TIMEOUT**: The password verification session timed out.
	//
	// - **CHECK_PWD_NETWORK_ERR**: A network error occurred.
	//
	// - **UNEXPECTED**: An unknown error occurred.
	//
	// example:
	//
	// CHECK_PWD_TIMEOUT
	ProxyStateErrorCode *string `json:"ProxyStateErrorCode,omitempty" xml:"ProxyStateErrorCode,omitempty"`
	// The proxy type. Valid values:
	//
	// - **SSHProxy**
	//
	// - **HTTPProxy**
	//
	// - **Socks5Proxy**
	//
	// example:
	//
	// HTTPProxy
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The username of the proxy server.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetNetworkDomainResponseBodyNetworkDomainProxies) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkDomainResponseBodyNetworkDomainProxies) GoString() string {
	return s.String()
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) GetAddress() *string {
	return s.Address
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) GetHasPassword() *bool {
	return s.HasPassword
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) GetNodeType() *string {
	return s.NodeType
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) GetPort() *int32 {
	return s.Port
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) GetProxyState() *string {
	return s.ProxyState
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) GetProxyStateErrorCode() *string {
	return s.ProxyStateErrorCode
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) GetProxyType() *string {
	return s.ProxyType
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) GetUser() *string {
	return s.User
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) SetAddress(v string) *GetNetworkDomainResponseBodyNetworkDomainProxies {
	s.Address = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) SetHasPassword(v bool) *GetNetworkDomainResponseBodyNetworkDomainProxies {
	s.HasPassword = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) SetNodeType(v string) *GetNetworkDomainResponseBodyNetworkDomainProxies {
	s.NodeType = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) SetPort(v int32) *GetNetworkDomainResponseBodyNetworkDomainProxies {
	s.Port = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) SetProxyState(v string) *GetNetworkDomainResponseBodyNetworkDomainProxies {
	s.ProxyState = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) SetProxyStateErrorCode(v string) *GetNetworkDomainResponseBodyNetworkDomainProxies {
	s.ProxyStateErrorCode = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) SetProxyType(v string) *GetNetworkDomainResponseBodyNetworkDomainProxies {
	s.ProxyType = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) SetUser(v string) *GetNetworkDomainResponseBodyNetworkDomainProxies {
	s.User = &v
	return s
}

func (s *GetNetworkDomainResponseBodyNetworkDomainProxies) Validate() error {
	return dara.Validate(s)
}
