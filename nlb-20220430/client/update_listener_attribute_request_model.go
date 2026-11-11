// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlpnEnabled(v bool) *UpdateListenerAttributeRequest
	GetAlpnEnabled() *bool
	SetAlpnPolicy(v string) *UpdateListenerAttributeRequest
	GetAlpnPolicy() *string
	SetCaCertificateIds(v []*string) *UpdateListenerAttributeRequest
	GetCaCertificateIds() []*string
	SetCaEnabled(v bool) *UpdateListenerAttributeRequest
	GetCaEnabled() *bool
	SetCertificateIds(v []*string) *UpdateListenerAttributeRequest
	GetCertificateIds() []*string
	SetClientToken(v string) *UpdateListenerAttributeRequest
	GetClientToken() *string
	SetCps(v int32) *UpdateListenerAttributeRequest
	GetCps() *int32
	SetDryRun(v bool) *UpdateListenerAttributeRequest
	GetDryRun() *bool
	SetIdleTimeout(v int32) *UpdateListenerAttributeRequest
	GetIdleTimeout() *int32
	SetListenerDescription(v string) *UpdateListenerAttributeRequest
	GetListenerDescription() *string
	SetListenerId(v string) *UpdateListenerAttributeRequest
	GetListenerId() *string
	SetMss(v int32) *UpdateListenerAttributeRequest
	GetMss() *int32
	SetProxyProtocolEnabled(v bool) *UpdateListenerAttributeRequest
	GetProxyProtocolEnabled() *bool
	SetProxyProtocolV2Config(v *UpdateListenerAttributeRequestProxyProtocolV2Config) *UpdateListenerAttributeRequest
	GetProxyProtocolV2Config() *UpdateListenerAttributeRequestProxyProtocolV2Config
	SetRegionId(v string) *UpdateListenerAttributeRequest
	GetRegionId() *string
	SetSecSensorEnabled(v bool) *UpdateListenerAttributeRequest
	GetSecSensorEnabled() *bool
	SetSecurityPolicyId(v string) *UpdateListenerAttributeRequest
	GetSecurityPolicyId() *string
	SetServerGroupId(v string) *UpdateListenerAttributeRequest
	GetServerGroupId() *string
}

type UpdateListenerAttributeRequest struct {
	// Specifies whether to enable Application-Layer Protocol Negotiation (ALPN). Valid values:
	//
	// - **true**: enables ALPN.
	//
	// - **false**: disables ALPN.
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
	// > You must configure this parameter after ALPN is enabled.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// HTTP1Only
	AlpnPolicy *string `json:"AlpnPolicy,omitempty" xml:"AlpnPolicy,omitempty"`
	// The CA certificates. You can add only one CA certificate.
	//
	// > This parameter takes effect only for TCPSSL listeners.
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
	// The server certificates. You can add only one server certificate.
	//
	// > This parameter takes effect only for TCPSSL listeners.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of new connections per second that the listener can process in each zone (VIP). Valid values: **0*	- to **1000000**. A value of **0*	- indicates that no limit is applied.
	//
	// example:
	//
	// 10000
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without updating the listener. The system checks whether required parameters are specified, the request format is valid, and the service limits are met. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The idle connection timeout period. Unit: seconds.
	//
	// - If the listener is a **TCP*	- or **TCPSSL*	- listener, valid values are **10*	- to **900**. Default value: **900**.
	//
	// - If the listener is a **UDP*	- listener, valid values are **10*	- to **90**. Default value: **90**.
	//
	// example:
	//
	// 900
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The custom name of the listener.
	//
	// The name must be 2 to 256 characters in length, and can contain Chinese characters, English letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// tcpssl_443
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-bp1bpn0kn908w4nbw****@443
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The maximum segment size (MSS) of a TCP packet. Unit: bytes. Valid values: **0*	- to **1500**. A value of **0*	- indicates that the MSS value of user TCP packets is not modified. This parameter is supported only by TCP and TCPSSL listeners.
	//
	// example:
	//
	// 344
	Mss *int32 `json:"Mss,omitempty" xml:"Mss,omitempty"`
	// Specifies whether to use the Proxy Protocol to pass the client source IP address to backend servers. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	ProxyProtocolEnabled *bool `json:"ProxyProtocolEnabled,omitempty" xml:"ProxyProtocolEnabled,omitempty"`
	// The configuration for passing the VpcId, PrivateLinkEpId, and PrivateLinkEpsId information to backend servers by using the Proxy Protocol.
	ProxyProtocolV2Config *UpdateListenerAttributeRequestProxyProtocolV2Config `json:"ProxyProtocolV2Config,omitempty" xml:"ProxyProtocolV2Config,omitempty" type:"Struct"`
	// The ID of the region where the Network Load Balancer (NLB) instance is deployed.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to retrieve the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable second-level monitoring. Valid values:
	//
	// - **true**: enables second-level monitoring.
	//
	// - **false**: disables second-level monitoring.
	//
	// example:
	//
	// false
	SecSensorEnabled *bool `json:"SecSensorEnabled,omitempty" xml:"SecSensorEnabled,omitempty"`
	// The ID of the security policy.
	//
	// > This parameter takes effect only for TCPSSL listeners.
	//
	// example:
	//
	// tls_cipher_policy_1_1
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of the server group.
	//
	// > - If the listener is a **TCP*	- listener, you can select a server group whose backend protocol is **TCP*	- or **TCP_UDP**. You cannot select a server group whose backend protocol is **UDP**.
	//
	// > - If the listener is a **UDP*	- listener, you can select a server group whose backend protocol is **UDP*	- or **TCP_UDP**. You cannot select a server group whose backend protocol is **TCP**.
	//
	// > - If the listener is a **TCPSSL*	- listener, you can select a server group whose backend protocol is **TCP*	- and for which **client address preservation is disabled**. You cannot select a server group whose backend protocol is **TCP*	- and for which **client address preservation is enabled**, or a server group whose backend protocol is **UDP*	- or **TCP_UDP**.
	//
	// example:
	//
	// sgp-ppdpc14gdm3x4o****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequest) GetAlpnEnabled() *bool {
	return s.AlpnEnabled
}

func (s *UpdateListenerAttributeRequest) GetAlpnPolicy() *string {
	return s.AlpnPolicy
}

func (s *UpdateListenerAttributeRequest) GetCaCertificateIds() []*string {
	return s.CaCertificateIds
}

func (s *UpdateListenerAttributeRequest) GetCaEnabled() *bool {
	return s.CaEnabled
}

func (s *UpdateListenerAttributeRequest) GetCertificateIds() []*string {
	return s.CertificateIds
}

func (s *UpdateListenerAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateListenerAttributeRequest) GetCps() *int32 {
	return s.Cps
}

func (s *UpdateListenerAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateListenerAttributeRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *UpdateListenerAttributeRequest) GetListenerDescription() *string {
	return s.ListenerDescription
}

func (s *UpdateListenerAttributeRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateListenerAttributeRequest) GetMss() *int32 {
	return s.Mss
}

func (s *UpdateListenerAttributeRequest) GetProxyProtocolEnabled() *bool {
	return s.ProxyProtocolEnabled
}

func (s *UpdateListenerAttributeRequest) GetProxyProtocolV2Config() *UpdateListenerAttributeRequestProxyProtocolV2Config {
	return s.ProxyProtocolV2Config
}

func (s *UpdateListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateListenerAttributeRequest) GetSecSensorEnabled() *bool {
	return s.SecSensorEnabled
}

func (s *UpdateListenerAttributeRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *UpdateListenerAttributeRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateListenerAttributeRequest) SetAlpnEnabled(v bool) *UpdateListenerAttributeRequest {
	s.AlpnEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetAlpnPolicy(v string) *UpdateListenerAttributeRequest {
	s.AlpnPolicy = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCaCertificateIds(v []*string) *UpdateListenerAttributeRequest {
	s.CaCertificateIds = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCaEnabled(v bool) *UpdateListenerAttributeRequest {
	s.CaEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCertificateIds(v []*string) *UpdateListenerAttributeRequest {
	s.CertificateIds = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetClientToken(v string) *UpdateListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCps(v int32) *UpdateListenerAttributeRequest {
	s.Cps = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDryRun(v bool) *UpdateListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetIdleTimeout(v int32) *UpdateListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerDescription(v string) *UpdateListenerAttributeRequest {
	s.ListenerDescription = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerId(v string) *UpdateListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetMss(v int32) *UpdateListenerAttributeRequest {
	s.Mss = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetProxyProtocolEnabled(v bool) *UpdateListenerAttributeRequest {
	s.ProxyProtocolEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetProxyProtocolV2Config(v *UpdateListenerAttributeRequestProxyProtocolV2Config) *UpdateListenerAttributeRequest {
	s.ProxyProtocolV2Config = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetRegionId(v string) *UpdateListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetSecSensorEnabled(v bool) *UpdateListenerAttributeRequest {
	s.SecSensorEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetSecurityPolicyId(v string) *UpdateListenerAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetServerGroupId(v string) *UpdateListenerAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) Validate() error {
	if s.ProxyProtocolV2Config != nil {
		if err := s.ProxyProtocolV2Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateListenerAttributeRequestProxyProtocolV2Config struct {
	// Specifies whether to use the Proxy Protocol to pass the CN-ID to backend servers. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	Ppv2CnIdEnabled *bool `json:"Ppv2CnIdEnabled,omitempty" xml:"Ppv2CnIdEnabled,omitempty"`
	// Specifies whether to use the Proxy Protocol to pass the PrivateLinkEpId to backend servers. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	Ppv2PrivateLinkEpIdEnabled *bool `json:"Ppv2PrivateLinkEpIdEnabled,omitempty" xml:"Ppv2PrivateLinkEpIdEnabled,omitempty"`
	// Specifies whether to use the Proxy Protocol to pass the PrivateLinkEpsId to backend servers. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	Ppv2PrivateLinkEpsIdEnabled *bool `json:"Ppv2PrivateLinkEpsIdEnabled,omitempty" xml:"Ppv2PrivateLinkEpsIdEnabled,omitempty"`
	// Specifies whether to use the Proxy Protocol to pass the VpcId to backend servers. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	Ppv2VpcIdEnabled *bool `json:"Ppv2VpcIdEnabled,omitempty" xml:"Ppv2VpcIdEnabled,omitempty"`
}

func (s UpdateListenerAttributeRequestProxyProtocolV2Config) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerAttributeRequestProxyProtocolV2Config) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) GetPpv2CnIdEnabled() *bool {
	return s.Ppv2CnIdEnabled
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) GetPpv2PrivateLinkEpIdEnabled() *bool {
	return s.Ppv2PrivateLinkEpIdEnabled
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) GetPpv2PrivateLinkEpsIdEnabled() *bool {
	return s.Ppv2PrivateLinkEpsIdEnabled
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) GetPpv2VpcIdEnabled() *bool {
	return s.Ppv2VpcIdEnabled
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) SetPpv2CnIdEnabled(v bool) *UpdateListenerAttributeRequestProxyProtocolV2Config {
	s.Ppv2CnIdEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) SetPpv2PrivateLinkEpIdEnabled(v bool) *UpdateListenerAttributeRequestProxyProtocolV2Config {
	s.Ppv2PrivateLinkEpIdEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) SetPpv2PrivateLinkEpsIdEnabled(v bool) *UpdateListenerAttributeRequestProxyProtocolV2Config {
	s.Ppv2PrivateLinkEpsIdEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) SetPpv2VpcIdEnabled(v bool) *UpdateListenerAttributeRequestProxyProtocolV2Config {
	s.Ppv2VpcIdEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestProxyProtocolV2Config) Validate() error {
	return dara.Validate(s)
}
