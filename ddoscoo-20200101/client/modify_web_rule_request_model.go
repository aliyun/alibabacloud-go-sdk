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
	// The domain name of the website business to operate.
	//
	// > The domain name must have a website forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names that have website forwarding rules configured.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The HTTPS advanced settings, which take effect only when the website protocol type supports HTTPS (**ProxyType*	- includes **https**). This parameter is a string converted from a JSON struct that contains the following fields:
	//
	// - **Http2https**: Integer type | Optional | Specifies whether to enable the HTTPS force redirect feature. Valid values: **0*	- (disabled) | **1*	- (enabled). Default value: 0 (disabled).
	//
	//   This is applicable when your website supports both HTTP and HTTPS protocols. After this setting is enabled, all HTTP requests will be forcibly converted to HTTPS requests and redirected to port 443 by default.
	//
	// - **Https2http**: Integer type | Optional | Specifies whether to enable the HTTP back-to-origin feature. Valid values: **0*	- (disabled) | **1*	- (enabled). Default value: 0 (disabled).
	//
	//   This is applicable when your website does not support HTTPS back-to-origin. After this setting is enabled, all HTTPS protocol requests will be forwarded to the origin through HTTP (Websockets protocol requests will be forwarded through Websocket), and the default back-to-origin port is 80.
	//
	// - **Http2**: Integer type | Optional | Specifies whether to enable the HTTP2.0 protocol type. Valid values: **0*	- (disabled) | **1*	- (enabled). Default value: 0 (disabled).
	//
	//   After this setting is enabled, the protocol version is HTTP2.0.
	//
	// example:
	//
	// {"Http2":1,"Http2https":1,"Https2http":1}
	HttpsExt *string `json:"HttpsExt,omitempty" xml:"HttpsExt,omitempty"`
	// The list of Anti-DDoS Proxy instance IDs to associate. You can associate up to 100 instances.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The protocol information of the website forwarding rule. This parameter is a string converted from a JSON array. Each element in the JSON array is a JSON struct that contains the following fields:
	//
	// - **ProxyType**: String type | Required | The website protocol type. Valid values: **http*	- | **https*	- | **websocket*	- | **websockets**.
	//
	// - **ProxyPorts**: array type | Required | The port numbers corresponding to the protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ProxyType":"https","ProxyPorts":[443]}]
	ProxyTypes *string `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty"`
	// The list of origin server addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.xxx.xxx.1
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The resource group ID of the Anti-DDoS Proxy instance in Resource Management. Default value: empty, which indicates that the instance belongs to the default resource group.
	//
	// For more information, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The address type of the origin server. Valid values:
	//
	// - **0**: indicates the IP address of the origin server.
	//
	// - **1**: indicates the domain name of the origin server. This is typically applicable to scenarios where other proxy services (such as WAF) are deployed between the origin server and Anti-DDoS Proxy, specifically referring to the redirect address of the proxy service (such as the WAF CNAME address).
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
