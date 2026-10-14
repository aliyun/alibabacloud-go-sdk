// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseId(v string) *CreateWebRuleRequest
	GetDefenseId() *string
	SetDomain(v string) *CreateWebRuleRequest
	GetDomain() *string
	SetHttpsExt(v string) *CreateWebRuleRequest
	GetHttpsExt() *string
	SetInstanceIds(v []*string) *CreateWebRuleRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *CreateWebRuleRequest
	GetResourceGroupId() *string
	SetRsType(v int32) *CreateWebRuleRequest
	GetRsType() *int32
	SetRules(v string) *CreateWebRuleRequest
	GetRules() *string
}

type CreateWebRuleRequest struct {
	// The ID of the protection rule to associate. This parameter is used when an Alibaba Cloud service, such as Object Storage Service (OSS), is integrated with an Anti-DDoS instance.
	//
	// > This parameter is for internal use. Do not use this parameter.
	//
	// For example, if your OSS is integrated with an Anti-DDoS Proxy instance, the instance allocates a resource pool of IP addresses to the OSS tenant account. Each IP address has a unique protection ID. The protection ID is a CNAME that resolves to the IP address of the Anti-DDoS instance by default. You can resolve the CNAME to an IP address. This lets you reuse the IP address for flexible scheduling.
	//
	// > You cannot specify both **InstanceIds*	- and **DefenseId**.
	//
	// example:
	//
	// testid
	DefenseId *string `json:"DefenseId,omitempty" xml:"DefenseId,omitempty"`
	// The domain name of the website to add to the Anti-DDoS Proxy instance for protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The advanced HTTPS settings. This parameter takes effect only when the value of **ProxyType*	- includes **https**. This parameter is a JSON string. The value is a JSON object that contains the following fields:
	//
	// - **Http2https**: This field is an integer. Optional. Specifies whether to forcibly redirect HTTP requests to HTTPS requests. Valid values: **0*	- (disabled) and **1*	- (enabled). The default value is 0.
	//
	//   This setting is available when your website supports both HTTP and HTTPS. If you enable this setting, all HTTP requests are redirected to HTTPS requests over port 443.
	//
	// - **Https2http**: This field is an integer. Optional. Specifies whether to enable HTTP back-to-origin. Valid values: **0*	- (disabled) and **1*	- (enabled). The default value is 0.
	//
	//   This setting applies to websites that do not support HTTPS back-to-origin. If you enable this setting, all HTTPS requests are sent to the origin server over HTTP on port 80 by default. All WebSockets requests are sent to the origin server over WebSocket.
	//
	// - **Http2**: This field is an integer. Optional. Specifies whether to enable HTTP/2. Valid values: **0*	- (disabled) and **1*	- (enabled). The default value is 0.
	//
	//   If you enable this setting, HTTP/2 is used.
	//
	// example:
	//
	// {"Http2":1,"Http2https":1,"Https2http":1}
	HttpsExt *string `json:"HttpsExt,omitempty" xml:"HttpsExt,omitempty"`
	// The IDs of the Anti-DDoS Proxy instances to associate.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the Anti-DDoS Proxy instance belongs in Resource Management. If you leave this parameter empty, the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The address type of the origin server. Valid values:
	//
	// - **0**: IP address.
	//
	// - **1**: domain name. This value is typically used when a proxy service, such as Web Application Firewall (WAF), is deployed between the origin server and the Anti-DDoS Proxy instance. The domain name is the back-to-origin address of the proxy service, such as the canonical name (CNAME) of WAF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
	// The details of the forwarding rule, in a JSON string format. The value is an array of objects. Each object contains the following fields:
	//
	// - **ProxyRules**: This field is a JSON array. Required. The information about the origin server, including the port number and server address. Each object in the array contains the following fields:
	//
	//   - **ProxyPort**: This field is an integer. Required. The port number that is used by the protocol.
	//
	//   - **RealServers**: This field is a string array. Required. The addresses of the origin servers.
	//
	// - **ProxyType**: This field is a string. Required. The protocol of the website. Valid values: **http**, **https**, **websocket**, and **websockets**.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ProxyRules":[{"ProxyPort":443,"RealServers":["192.1.XX.XX"]}],"ProxyType":"https"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreateWebRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWebRuleRequest) GetDefenseId() *string {
	return s.DefenseId
}

func (s *CreateWebRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateWebRuleRequest) GetHttpsExt() *string {
	return s.HttpsExt
}

func (s *CreateWebRuleRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateWebRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateWebRuleRequest) GetRsType() *int32 {
	return s.RsType
}

func (s *CreateWebRuleRequest) GetRules() *string {
	return s.Rules
}

func (s *CreateWebRuleRequest) SetDefenseId(v string) *CreateWebRuleRequest {
	s.DefenseId = &v
	return s
}

func (s *CreateWebRuleRequest) SetDomain(v string) *CreateWebRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateWebRuleRequest) SetHttpsExt(v string) *CreateWebRuleRequest {
	s.HttpsExt = &v
	return s
}

func (s *CreateWebRuleRequest) SetInstanceIds(v []*string) *CreateWebRuleRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateWebRuleRequest) SetResourceGroupId(v string) *CreateWebRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWebRuleRequest) SetRsType(v int32) *CreateWebRuleRequest {
	s.RsType = &v
	return s
}

func (s *CreateWebRuleRequest) SetRules(v string) *CreateWebRuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateWebRuleRequest) Validate() error {
	return dara.Validate(s)
}
