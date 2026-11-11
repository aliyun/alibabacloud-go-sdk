// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListenerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlpnEnabled(v bool) *CreateListenerShrinkRequest
	GetAlpnEnabled() *bool
	SetAlpnPolicy(v string) *CreateListenerShrinkRequest
	GetAlpnPolicy() *string
	SetCaCertificateIds(v []*string) *CreateListenerShrinkRequest
	GetCaCertificateIds() []*string
	SetCaEnabled(v bool) *CreateListenerShrinkRequest
	GetCaEnabled() *bool
	SetCertificateIds(v []*string) *CreateListenerShrinkRequest
	GetCertificateIds() []*string
	SetClientToken(v string) *CreateListenerShrinkRequest
	GetClientToken() *string
	SetCps(v int32) *CreateListenerShrinkRequest
	GetCps() *int32
	SetDryRun(v bool) *CreateListenerShrinkRequest
	GetDryRun() *bool
	SetEndPort(v int32) *CreateListenerShrinkRequest
	GetEndPort() *int32
	SetIdleTimeout(v int32) *CreateListenerShrinkRequest
	GetIdleTimeout() *int32
	SetListenerDescription(v string) *CreateListenerShrinkRequest
	GetListenerDescription() *string
	SetListenerPort(v int32) *CreateListenerShrinkRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *CreateListenerShrinkRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *CreateListenerShrinkRequest
	GetLoadBalancerId() *string
	SetMss(v int32) *CreateListenerShrinkRequest
	GetMss() *int32
	SetProxyProtocolEnabled(v bool) *CreateListenerShrinkRequest
	GetProxyProtocolEnabled() *bool
	SetProxyProtocolV2ConfigShrink(v string) *CreateListenerShrinkRequest
	GetProxyProtocolV2ConfigShrink() *string
	SetRegionId(v string) *CreateListenerShrinkRequest
	GetRegionId() *string
	SetSecSensorEnabled(v bool) *CreateListenerShrinkRequest
	GetSecSensorEnabled() *bool
	SetSecurityPolicyId(v string) *CreateListenerShrinkRequest
	GetSecurityPolicyId() *string
	SetServerGroupId(v string) *CreateListenerShrinkRequest
	GetServerGroupId() *string
	SetStartPort(v int32) *CreateListenerShrinkRequest
	GetStartPort() *int32
	SetTag(v []*CreateListenerShrinkRequestTag) *CreateListenerShrinkRequest
	GetTag() []*CreateListenerShrinkRequestTag
}

type CreateListenerShrinkRequest struct {
	// Specifies whether to enable ALPN. Valid values:
	//
	// - **true**: enables ALPN.
	//
	// - **false*	- (default): disables ALPN.
	//
	// example:
	//
	// false
	AlpnEnabled *bool `json:"AlpnEnabled,omitempty" xml:"AlpnEnabled,omitempty"`
	// The ALPN policy. Valid values:
	//
	// - **HTTP1Only**: negotiates only HTTP/1.x. The priority is HTTP/1.1 > HTTP/1.0.
	//
	// - **HTTP2Only**: negotiates only HTTP/2.
	//
	// - **HTTP2Optional**: prefers HTTP/1.x but also accepts HTTP/2. The priority is HTTP/1.1 > HTTP/1.0 > HTTP/2.
	//
	// - **HTTP2Preferred**: prefers HTTP/2 but also accepts HTTP/1.x. The priority is HTTP/2 > HTTP/1.1 > HTTP/1.0.
	//
	// > You must configure this parameter after **AlpnEnabled*	- is enabled.
	//
	// example:
	//
	// HTTP1Only
	AlpnPolicy *string `json:"AlpnPolicy,omitempty" xml:"AlpnPolicy,omitempty"`
	// The CA certificate list. This parameter takes effect only for TCPSSL listeners.
	//
	// > Currently, you can add only one CA certificate.
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	// Specifies whether to enable mutual authentication. Valid values:
	//
	// - **true**: enables mutual authentication.
	//
	// - **false*	- (default): disables mutual authentication.
	//
	// example:
	//
	// false
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// The server certificate list. This parameter takes effect only for TCPSSL listeners.
	//
	// > Currently, you can add only one server certificate.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a parameter value on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The limit on the number of new connections processed per second by the listener in each zone (VIP). Valid values: **0*	- to **1000000**. **0*	- indicates that no limit is applied.
	//
	// example:
	//
	// 100
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: sends a precheck request without creating the listener. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends a regular request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The end port for all-port listening. Valid values: **1*	- to **65535**.
	//
	// The end port must be greater than the start port.
	//
	// > You must configure this parameter when **ListenerPort*	- is set to **0**.
	//
	// example:
	//
	// 566
	EndPort *int32 `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// The idle timeout period for connections. Unit: seconds.
	//
	// - If **ListenerProtocol*	- is set to **TCP*	- or **TCPSSL**, the valid values are **10*	- to **900**. Default value: **900**.
	//
	// - If **ListenerProtocol*	- is set to **UDP**, the valid values are **10*	- to **90**. Default value: **90**.
	//
	// example:
	//
	// 900
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The custom listener name.
	//
	// The name must be 2 to 256 characters in length, and can contain Chinese characters, English letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// tcp_80
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener port. Valid values: **0*	- to **65535**.
	//
	// **0**: indicates that the all-port feature is used. If you set this parameter to **0**, you must configure **StartPort*	- and **EndPort**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol. Valid values: **TCP**, **UDP**, and **TCPSSL**.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The Network Load Balancer (NLB) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// nlb-83ckzc8d4xlp8o****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The maximum segment size (MSS) of TCP packets. Unit: bytes. Valid values: **0*	- to **1500**.
	//
	// **0**: indicates that the MSS value of user TCP packets is not modified.
	//
	// > This parameter is supported only by TCP and TCPSSL listeners.
	//
	// example:
	//
	// 43
	Mss *int32 `json:"Mss,omitempty" xml:"Mss,omitempty"`
	// Specifies whether to use the Proxy Protocol to pass the client source IP address to backend servers. Valid values:
	//
	// - **true**: enables the feature.
	//
	// - **false*	- (default): disables the feature.
	//
	// example:
	//
	// false
	ProxyProtocolEnabled *bool `json:"ProxyProtocolEnabled,omitempty" xml:"ProxyProtocolEnabled,omitempty"`
	// The configuration for passing the VpcId, PrivateLinkEpId, and PrivateLinkEpsId information to backend servers by using the Proxy Protocol.
	ProxyProtocolV2ConfigShrink *string `json:"ProxyProtocolV2Config,omitempty" xml:"ProxyProtocolV2Config,omitempty"`
	// The region ID of the Network Load Balancer (NLB) instance.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable second-level monitoring. Valid values:
	//
	// - **true**: enables second-level monitoring.
	//
	// - **false*	- (default): disables second-level monitoring.
	//
	// example:
	//
	// false
	SecSensorEnabled *bool `json:"SecSensorEnabled,omitempty" xml:"SecSensorEnabled,omitempty"`
	// The ID of the security policy. System security policies and custom security policies are supported.
	//
	// - System policies: **tls_cipher_policy_1_0*	- (default), **tls_cipher_policy_1_1**, **tls_cipher_policy_1_2**, **tls_cipher_policy_1_2_strict**, and **tls_cipher_policy_1_2_strict_with_1_3**.
	//
	// - Custom security policies: Enter the ID of a custom security policy.
	//
	//     - To create a custom policy, refer to [CreateSecurityPolicy](https://help.aliyun.com/document_detail/445901.html).
	//
	//     - To query security policies, refer to [ListSecurityPolicy](https://help.aliyun.com/document_detail/445900.html).
	//
	// > This parameter takes effect only for TCPSSL listeners.
	//
	// example:
	//
	// tls_cipher_policy_1_0
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of the server group.
	//
	// > - If **ListenerProtocol*	- is set to **TCP**, the listener supports server groups whose backend protocol is **TCP*	- or **TCP_UDP**, but does not support server groups whose backend protocol is **UDP**.
	//
	// > - If **ListenerProtocol*	- is set to **UDP**, the listener supports server groups whose backend protocol is **UDP*	- or **TCP_UDP**, but does not support server groups whose backend protocol is **TCP**.
	//
	// > - If **ListenerProtocol*	- is set to **TCPSSL**, the listener supports server groups whose backend protocol is **TCP*	- and for which **client address preservation is disabled**. It does not support server groups whose backend protocol is **TCP*	- and for which **client address preservation is enabled**, or server groups whose backend protocol is **UDP*	- or **TCP_UDP**.
	//
	// example:
	//
	// sgp-ppdpc14gdm3x4o****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The start port for all-port listening. Valid values: **1*	- to **65535**.
	//
	// > You must configure this parameter when **ListenerPort*	- is set to **0**.
	//
	// example:
	//
	// 244
	StartPort *int32 `json:"StartPort,omitempty" xml:"StartPort,omitempty"`
	// The list of tags.
	//
	// if can be null:
	// true
	Tag []*CreateListenerShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateListenerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerShrinkRequest) GetAlpnEnabled() *bool {
	return s.AlpnEnabled
}

func (s *CreateListenerShrinkRequest) GetAlpnPolicy() *string {
	return s.AlpnPolicy
}

func (s *CreateListenerShrinkRequest) GetCaCertificateIds() []*string {
	return s.CaCertificateIds
}

func (s *CreateListenerShrinkRequest) GetCaEnabled() *bool {
	return s.CaEnabled
}

func (s *CreateListenerShrinkRequest) GetCertificateIds() []*string {
	return s.CertificateIds
}

func (s *CreateListenerShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateListenerShrinkRequest) GetCps() *int32 {
	return s.Cps
}

func (s *CreateListenerShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateListenerShrinkRequest) GetEndPort() *int32 {
	return s.EndPort
}

func (s *CreateListenerShrinkRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *CreateListenerShrinkRequest) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *CreateListenerShrinkRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateListenerShrinkRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *CreateListenerShrinkRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateListenerShrinkRequest) GetMss() *int32 {
	return s.Mss
}

func (s *CreateListenerShrinkRequest) GetProxyProtocolEnabled() *bool {
	return s.ProxyProtocolEnabled
}

func (s *CreateListenerShrinkRequest) GetProxyProtocolV2ConfigShrink() *string {
	return s.ProxyProtocolV2ConfigShrink
}

func (s *CreateListenerShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateListenerShrinkRequest) GetSecSensorEnabled() *bool {
	return s.SecSensorEnabled
}

func (s *CreateListenerShrinkRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *CreateListenerShrinkRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *CreateListenerShrinkRequest) GetStartPort() *int32 {
	return s.StartPort
}

func (s *CreateListenerShrinkRequest) GetTag() []*CreateListenerShrinkRequestTag {
	return s.Tag
}

func (s *CreateListenerShrinkRequest) SetAlpnEnabled(v bool) *CreateListenerShrinkRequest {
	s.AlpnEnabled = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetAlpnPolicy(v string) *CreateListenerShrinkRequest {
	s.AlpnPolicy = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetCaCertificateIds(v []*string) *CreateListenerShrinkRequest {
	s.CaCertificateIds = v
	return s
}

func (s *CreateListenerShrinkRequest) SetCaEnabled(v bool) *CreateListenerShrinkRequest {
	s.CaEnabled = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetCertificateIds(v []*string) *CreateListenerShrinkRequest {
	s.CertificateIds = v
	return s
}

func (s *CreateListenerShrinkRequest) SetClientToken(v string) *CreateListenerShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetCps(v int32) *CreateListenerShrinkRequest {
	s.Cps = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetDryRun(v bool) *CreateListenerShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetEndPort(v int32) *CreateListenerShrinkRequest {
	s.EndPort = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetIdleTimeout(v int32) *CreateListenerShrinkRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetListenerDescription(v string) *CreateListenerShrinkRequest {
	s.ListenerDescription = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetListenerPort(v int32) *CreateListenerShrinkRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetListenerProtocol(v string) *CreateListenerShrinkRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetLoadBalancerId(v string) *CreateListenerShrinkRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetMss(v int32) *CreateListenerShrinkRequest {
	s.Mss = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetProxyProtocolEnabled(v bool) *CreateListenerShrinkRequest {
	s.ProxyProtocolEnabled = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetProxyProtocolV2ConfigShrink(v string) *CreateListenerShrinkRequest {
	s.ProxyProtocolV2ConfigShrink = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetRegionId(v string) *CreateListenerShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetSecSensorEnabled(v bool) *CreateListenerShrinkRequest {
	s.SecSensorEnabled = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetSecurityPolicyId(v string) *CreateListenerShrinkRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetServerGroupId(v string) *CreateListenerShrinkRequest {
	s.ServerGroupId = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetStartPort(v int32) *CreateListenerShrinkRequest {
	s.StartPort = &v
	return s
}

func (s *CreateListenerShrinkRequest) SetTag(v []*CreateListenerShrinkRequestTag) *CreateListenerShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateListenerShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateListenerShrinkRequestTag struct {
	// The tag key. The key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`. Supported characters include letters, digits, underscores (_), periods (.), colons (:), forward slashes (/), equal signs (=), plus signs (+), hyphens (-), and at signs (@).
	//
	// You can add up to 20 tags in a single call.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`. Supported characters include letters, digits, underscores (_), periods (.), colons (:), forward slashes (/), equal signs (=), plus signs (+), hyphens (-), and at signs (@).
	//
	// You can add up to 20 tags in a single call.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateListenerShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateListenerShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateListenerShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateListenerShrinkRequestTag) SetKey(v string) *CreateListenerShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateListenerShrinkRequestTag) SetValue(v string) *CreateListenerShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateListenerShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
