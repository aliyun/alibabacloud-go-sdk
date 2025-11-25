// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CreateDomainResourceRequest
	GetDomain() *string
	SetHttpsExt(v string) *CreateDomainResourceRequest
	GetHttpsExt() *string
	SetInstanceIds(v []*string) *CreateDomainResourceRequest
	GetInstanceIds() []*string
	SetProxyTypes(v []*CreateDomainResourceRequestProxyTypes) *CreateDomainResourceRequest
	GetProxyTypes() []*CreateDomainResourceRequestProxyTypes
	SetRealServers(v []*string) *CreateDomainResourceRequest
	GetRealServers() []*string
	SetRsType(v int32) *CreateDomainResourceRequest
	GetRsType() *int32
}

type CreateDomainResourceRequest struct {
	// The domain name of the website that you want to add to the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The advanced HTTPS settings. This parameter takes effect only when the value of the **ProxyType*	- parameter includes **https**. The value is a string that consists of a JSON struct. The JSON struct contains the following fields:
	//
	// 	- **Http2https**: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: **0*	- and **1**. The value 0 indicates that Enforce HTTPS Routing is turned off. The value 1 indicates that Enforce HTTPS Routing is turned on. The default value is 0.
	//
	//     If your website supports both HTTP and HTTPS, this feature meets your business requirements. If you enable this feature, all HTTP requests to access the website are redirected to HTTPS requests on the standard port 443.
	//
	// 	- **Https2http**: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: **0*	- and **1**. The value 0 indicates that Enable HTTP is turned off. The value 1 indicates that Enable HTTP is turned on. The default value is 0.
	//
	//     If your website does not support HTTPS, this feature meets your business requirements If this feature is enabled, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. This feature can redirect WebSockets requests to WebSocket requests. Requests are redirected over the standard port 80.
	//
	// 	- **Http2**: specifies whether to turn on Enable HTTP/2. This field is optional. Data type: integer. Valid values: **0*	- and **1**. The value 0 indicates that Enable HTTP/2 is turned off. The value 1 indicates that Enable HTTP/2 is turned on. The default value is 0.
	//
	//     After you turn on the switch, HTTP/2 is used.
	//
	// example:
	//
	// {"Http2":1,"Http2https":1,"Https2http":1}
	HttpsExt *string `json:"HttpsExt,omitempty" xml:"HttpsExt,omitempty"`
	// An array consisting of the IDs of instances that you want to associate.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The details about the protocol type and port number.
	//
	// This parameter is required.
	ProxyTypes []*CreateDomainResourceRequestProxyTypes `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	// An array that consists of the addresses of origin servers.
	//
	// This parameter is required.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The address type of the origin server. Valid values:
	//
	// 	- **0**: IP address.
	//
	// 	- **1**: domain name.
	//
	//     This parameter is suitable for scenarios in which another proxy service, such as Web Application Firewall (WAF), is deployed between the origin server and Anti-DDoS Proxy. The address is the redirection address of the proxy service, such as the CNAME of WAF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s CreateDomainResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainResourceRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateDomainResourceRequest) GetHttpsExt() *string {
	return s.HttpsExt
}

func (s *CreateDomainResourceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateDomainResourceRequest) GetProxyTypes() []*CreateDomainResourceRequestProxyTypes {
	return s.ProxyTypes
}

func (s *CreateDomainResourceRequest) GetRealServers() []*string {
	return s.RealServers
}

func (s *CreateDomainResourceRequest) GetRsType() *int32 {
	return s.RsType
}

func (s *CreateDomainResourceRequest) SetDomain(v string) *CreateDomainResourceRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainResourceRequest) SetHttpsExt(v string) *CreateDomainResourceRequest {
	s.HttpsExt = &v
	return s
}

func (s *CreateDomainResourceRequest) SetInstanceIds(v []*string) *CreateDomainResourceRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateDomainResourceRequest) SetProxyTypes(v []*CreateDomainResourceRequestProxyTypes) *CreateDomainResourceRequest {
	s.ProxyTypes = v
	return s
}

func (s *CreateDomainResourceRequest) SetRealServers(v []*string) *CreateDomainResourceRequest {
	s.RealServers = v
	return s
}

func (s *CreateDomainResourceRequest) SetRsType(v int32) *CreateDomainResourceRequest {
	s.RsType = &v
	return s
}

func (s *CreateDomainResourceRequest) Validate() error {
	if s.ProxyTypes != nil {
		for _, item := range s.ProxyTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDomainResourceRequestProxyTypes struct {
	// The port numbers.
	//
	// This parameter is required.
	ProxyPorts []*int32 `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	// The type of the protocol. Valid values:
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// 	- **websocket**
	//
	// 	- **websockets**
	//
	// example:
	//
	// http
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s CreateDomainResourceRequestProxyTypes) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResourceRequestProxyTypes) GoString() string {
	return s.String()
}

func (s *CreateDomainResourceRequestProxyTypes) GetProxyPorts() []*int32 {
	return s.ProxyPorts
}

func (s *CreateDomainResourceRequestProxyTypes) GetProxyType() *string {
	return s.ProxyType
}

func (s *CreateDomainResourceRequestProxyTypes) SetProxyPorts(v []*int32) *CreateDomainResourceRequestProxyTypes {
	s.ProxyPorts = v
	return s
}

func (s *CreateDomainResourceRequestProxyTypes) SetProxyType(v string) *CreateDomainResourceRequestProxyTypes {
	s.ProxyType = &v
	return s
}

func (s *CreateDomainResourceRequestProxyTypes) Validate() error {
	return dara.Validate(s)
}
