// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyWebRuleRequest
	GetDomain() *string
	SetHttpsExt(v string) *ModifyWebRuleRequest
	GetHttpsExt() *string
	SetInstanceIds(v []*string) *ModifyWebRuleRequest
	GetInstanceIds() []*string
	SetProxyTypes(v string) *ModifyWebRuleRequest
	GetProxyTypes() *string
	SetRealServers(v []*string) *ModifyWebRuleRequest
	GetRealServers() []*string
	SetResourceGroupId(v string) *ModifyWebRuleRequest
	GetResourceGroupId() *string
	SetRsType(v int32) *ModifyWebRuleRequest
	GetRsType() *int32
}

type ModifyWebRuleRequest struct {
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query the domain names for which forwarding rules are configured.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The advanced HTTPS settings. This parameter takes effect only when the value of **ProxyType*	- includes **https**. The value is a string that consists of a JSON struct. The JSON struct contains the following fields:
	//
	// 	- **Http2https**: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: **0*	- and **1**. The value 0 indicates that Enforce HTTPS Routing is turned off. The value 1 indicates that Enforce HTTPS Routing is turned on. The default value is 0.
	//
	//     If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
	//
	// 	- **Https2http**: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: **0*	- and **1**. The value 0 indicates that Enable HTTP is turned off. The value 1 indicates that Enable HTTP is turned on. The default value is 0.
	//
	//     If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
	//
	// 	- **Http2**: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: **0*	- and **1**. The value 0 indicates that Enable HTTP/2 is turned off. The value 1 indicates that Enable HTTP/2 is turned on. The default value is 0.
	//
	//     After you turn on the switch, the protocol type is HTTP/2.
	//
	// example:
	//
	// {"Http2":1,"Http2https":1,"Https2http":1}
	HttpsExt *string `json:"HttpsExt,omitempty" xml:"HttpsExt,omitempty"`
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The protocol of the forwarding rule. The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// 	- **ProxyType**: the protocol type. This field is required and must be a string. Valid values: **http**, **https**, **websocket**, and **websockets**.
	//
	// 	- **ProxyPort**: the port number. This field is required and must be an array.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ProxyType":"https","ProxyPorts":[443]}]
	ProxyTypes *string `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.xxx.xxx.1
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The address type of the origin server. Valid values:
	//
	// 	- **0**: IP address.
	//
	// 	- **1**: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s ModifyWebRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebRuleRequest) GetHttpsExt() *string {
	return s.HttpsExt
}

func (s *ModifyWebRuleRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyWebRuleRequest) GetProxyTypes() *string {
	return s.ProxyTypes
}

func (s *ModifyWebRuleRequest) GetRealServers() []*string {
	return s.RealServers
}

func (s *ModifyWebRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebRuleRequest) GetRsType() *int32 {
	return s.RsType
}

func (s *ModifyWebRuleRequest) SetDomain(v string) *ModifyWebRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebRuleRequest) SetHttpsExt(v string) *ModifyWebRuleRequest {
	s.HttpsExt = &v
	return s
}

func (s *ModifyWebRuleRequest) SetInstanceIds(v []*string) *ModifyWebRuleRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyWebRuleRequest) SetProxyTypes(v string) *ModifyWebRuleRequest {
	s.ProxyTypes = &v
	return s
}

func (s *ModifyWebRuleRequest) SetRealServers(v []*string) *ModifyWebRuleRequest {
	s.RealServers = v
	return s
}

func (s *ModifyWebRuleRequest) SetResourceGroupId(v string) *ModifyWebRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebRuleRequest) SetRsType(v int32) *ModifyWebRuleRequest {
	s.RsType = &v
	return s
}

func (s *ModifyWebRuleRequest) Validate() error {
	return dara.Validate(s)
}
