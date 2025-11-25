// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyDomainResourceRequest
	GetDomain() *string
	SetHttpsExt(v string) *ModifyDomainResourceRequest
	GetHttpsExt() *string
	SetInstanceIds(v []*string) *ModifyDomainResourceRequest
	GetInstanceIds() []*string
	SetProxyTypes(v []*ModifyDomainResourceRequestProxyTypes) *ModifyDomainResourceRequest
	GetProxyTypes() []*ModifyDomainResourceRequestProxyTypes
	SetRealServers(v []*string) *ModifyDomainResourceRequest
	GetRealServers() []*string
	SetRsType(v int32) *ModifyDomainResourceRequest
	GetRsType() *int32
}

type ModifyDomainResourceRequest struct {
	// The domain name that is added to the Anti-DDoS Pro or Anti-DDoS Premium instance.
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
	ProxyTypes []*ModifyDomainResourceRequestProxyTypes `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	// An array that consists of the addresses of origin servers.
	//
	// This parameter is required.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The address type of the origin server. Valid values:
	//
	// 	- **0**: IP address
	//
	// 	- **1**: domain name
	//
	//     If you deploy proxies, such as a Web Application Firewall (WAF) instance, between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance, set the value to 1. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s ModifyDomainResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainResourceRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyDomainResourceRequest) GetHttpsExt() *string {
	return s.HttpsExt
}

func (s *ModifyDomainResourceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyDomainResourceRequest) GetProxyTypes() []*ModifyDomainResourceRequestProxyTypes {
	return s.ProxyTypes
}

func (s *ModifyDomainResourceRequest) GetRealServers() []*string {
	return s.RealServers
}

func (s *ModifyDomainResourceRequest) GetRsType() *int32 {
	return s.RsType
}

func (s *ModifyDomainResourceRequest) SetDomain(v string) *ModifyDomainResourceRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainResourceRequest) SetHttpsExt(v string) *ModifyDomainResourceRequest {
	s.HttpsExt = &v
	return s
}

func (s *ModifyDomainResourceRequest) SetInstanceIds(v []*string) *ModifyDomainResourceRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyDomainResourceRequest) SetProxyTypes(v []*ModifyDomainResourceRequestProxyTypes) *ModifyDomainResourceRequest {
	s.ProxyTypes = v
	return s
}

func (s *ModifyDomainResourceRequest) SetRealServers(v []*string) *ModifyDomainResourceRequest {
	s.RealServers = v
	return s
}

func (s *ModifyDomainResourceRequest) SetRsType(v int32) *ModifyDomainResourceRequest {
	s.RsType = &v
	return s
}

func (s *ModifyDomainResourceRequest) Validate() error {
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

type ModifyDomainResourceRequestProxyTypes struct {
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
	// https
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s ModifyDomainResourceRequestProxyTypes) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainResourceRequestProxyTypes) GoString() string {
	return s.String()
}

func (s *ModifyDomainResourceRequestProxyTypes) GetProxyPorts() []*int32 {
	return s.ProxyPorts
}

func (s *ModifyDomainResourceRequestProxyTypes) GetProxyType() *string {
	return s.ProxyType
}

func (s *ModifyDomainResourceRequestProxyTypes) SetProxyPorts(v []*int32) *ModifyDomainResourceRequestProxyTypes {
	s.ProxyPorts = v
	return s
}

func (s *ModifyDomainResourceRequestProxyTypes) SetProxyType(v string) *ModifyDomainResourceRequestProxyTypes {
	s.ProxyType = &v
	return s
}

func (s *ModifyDomainResourceRequestProxyTypes) Validate() error {
	return dara.Validate(s)
}
