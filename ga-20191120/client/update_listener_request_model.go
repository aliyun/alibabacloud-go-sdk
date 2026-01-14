// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendPorts(v []*UpdateListenerRequestBackendPorts) *UpdateListenerRequest
	GetBackendPorts() []*UpdateListenerRequestBackendPorts
	SetCertificates(v []*UpdateListenerRequestCertificates) *UpdateListenerRequest
	GetCertificates() []*UpdateListenerRequestCertificates
	SetClientAffinity(v string) *UpdateListenerRequest
	GetClientAffinity() *string
	SetClientToken(v string) *UpdateListenerRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateListenerRequest
	GetDescription() *string
	SetHttpVersion(v string) *UpdateListenerRequest
	GetHttpVersion() *string
	SetIdleTimeout(v int32) *UpdateListenerRequest
	GetIdleTimeout() *int32
	SetListenerId(v string) *UpdateListenerRequest
	GetListenerId() *string
	SetName(v string) *UpdateListenerRequest
	GetName() *string
	SetPortRanges(v []*UpdateListenerRequestPortRanges) *UpdateListenerRequest
	GetPortRanges() []*UpdateListenerRequestPortRanges
	SetProtocol(v string) *UpdateListenerRequest
	GetProtocol() *string
	SetProxyProtocol(v string) *UpdateListenerRequest
	GetProxyProtocol() *string
	SetRegionId(v string) *UpdateListenerRequest
	GetRegionId() *string
	SetRequestTimeout(v int32) *UpdateListenerRequest
	GetRequestTimeout() *int32
	SetSecurityPolicyId(v string) *UpdateListenerRequest
	GetSecurityPolicyId() *string
	SetXForwardedForConfig(v *UpdateListenerRequestXForwardedForConfig) *UpdateListenerRequest
	GetXForwardedForConfig() *UpdateListenerRequestXForwardedForConfig
}

type UpdateListenerRequest struct {
	// The range of ports that are used by backend servers to receive requests.
	BackendPorts []*UpdateListenerRequestBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The SSL certificate.
	Certificates []*UpdateListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// Indicates whether client affinity is enabled for the listener. Valid values:
	//
	// 	- **NONE**: Client affinity is disabled. Requests from the same client may be forwarded to different endpoints.
	//
	// 	- **SOURCE_IP**: Client affinity is enabled. When a client accesses stateful applications, requests from the same client are forwarded to the same endpoint regardless of the source port or protocol.
	//
	// example:
	//
	// SOURCE_IP
	ClientAffinity *string `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the listener.
	//
	// The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Listener
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum version of the HTTP protocol. Valid values:
	//
	// 	- **http3**
	//
	// 	- **http2**
	//
	// 	- **http1.1**
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// http2
	HttpVersion *string `json:"HttpVersion,omitempty" xml:"HttpVersion,omitempty"`
	// The timeout period for idle connections. Unit: seconds.
	//
	// 	- TCP: 10-900. Default value: 900. Unit: seconds.
	//
	// 	- UDP: 10-20. Default value: 20. Unit: seconds.
	//
	// 	- HTTP/HTTPS: 1-60. Default value: 15. Unit: seconds.
	//
	// example:
	//
	// 900
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The name of the listener.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// Listener
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The listener ports that are used to receive requests and forward the requests to endpoints.
	//
	// Valid values: **1*	- to **65499**.
	//
	// The maximum number of ports that can be configured varies based on the routing type and protocol of the listener. For more information, see [Listener overview](https://help.aliyun.com/document_detail/153216.html).
	PortRanges []*UpdateListenerRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The network transmission protocol that is used by the listener. Valid values:
	//
	// 	- **tcp**: TCP
	//
	// 	- **udp**: UDP
	//
	// 	- **http**: HTTP
	//
	// 	- **https**: HTTPS
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Deprecated
	//
	// Specifies whether to preserve source IP addresses of clients.
	//
	// 	- **true*	- This feature allows you to view client IP addresses on backend servers.
	//
	// 	- **false*	- (default)
	//
	// >  This parameter will be discontinued in the API operations that are used to configure listeners. We recommend that you set this parameter when you call API operations to configure endpoint groups. For more information about the **ProxyProtocol*	- parameter, see [CreateEndpointGroup](https://help.aliyun.com/document_detail/153259.html) and [UpdateEndpointGroup](https://help.aliyun.com/document_detail/153262.html).
	//
	// example:
	//
	// false
	ProxyProtocol *string `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period for HTTP or HTTPS requests.
	//
	// Valid values: 1 to 180. Default value: 60. Unit: seconds.
	//
	// >  This parameter takes effect only for HTTP or HTTPS listeners. If the backend server does not respond within the timeout period, GA returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The ID of the security policy. Valid values:
	//
	// 	- **tls_cipher_policy_1_0**
	//
	//     	- Supported Transport Layer Security (TLS) versions: TLS 1.0, TLS 1.1, and TLS 1.2
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA
	//
	// 	- **tls_cipher_policy_1_1**
	//
	//     	- Supported TLS versions: TLS 1.1 and TLS 1.2
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA
	//
	// 	- **tls_cipher_policy_1_2**
	//
	//     	- Supported TLS version: TLS 1.2
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA
	//
	// 	- **tls_cipher_policy_1_2_strict**
	//
	//     	- Supported TLS version: TLS 1.2
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA
	//
	// 	- **tls_cipher_policy_1_2_strict_with_1_3**
	//
	//     	- Supported TLS versions: TLS 1.2 and TLS 1.3
	//
	//     	- Supported cipher suites: TLS_AES_128_GCM_SHA256, TLS_AES_256_GCM_SHA384, TLS_CHACHA20_POLY1305_SHA256, TLS_AES_128_CCM_SHA256, TLS_AES_128_CCM_8_SHA256, ECDHE-ECDSA-AES128-GCM-SHA256, ECDHE-ECDSA-AES256-GCM-SHA384, ECDHE-ECDSA-AES128-SHA256, ECDHE-ECDSA-AES256-SHA384, ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-ECDSA-AES128-SHA, ECDHE-ECDSA-AES256-SHA, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA
	//
	// > This parameter is available only when you create an HTTPS listener.
	//
	// example:
	//
	// tls_cipher_policy_1_0
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The `XForward` headers.
	XForwardedForConfig *UpdateListenerRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s UpdateListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequest) GetBackendPorts() []*UpdateListenerRequestBackendPorts {
	return s.BackendPorts
}

func (s *UpdateListenerRequest) GetCertificates() []*UpdateListenerRequestCertificates {
	return s.Certificates
}

func (s *UpdateListenerRequest) GetClientAffinity() *string {
	return s.ClientAffinity
}

func (s *UpdateListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateListenerRequest) GetHttpVersion() *string {
	return s.HttpVersion
}

func (s *UpdateListenerRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *UpdateListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateListenerRequest) GetName() *string {
	return s.Name
}

func (s *UpdateListenerRequest) GetPortRanges() []*UpdateListenerRequestPortRanges {
	return s.PortRanges
}

func (s *UpdateListenerRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateListenerRequest) GetProxyProtocol() *string {
	return s.ProxyProtocol
}

func (s *UpdateListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateListenerRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *UpdateListenerRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *UpdateListenerRequest) GetXForwardedForConfig() *UpdateListenerRequestXForwardedForConfig {
	return s.XForwardedForConfig
}

func (s *UpdateListenerRequest) SetBackendPorts(v []*UpdateListenerRequestBackendPorts) *UpdateListenerRequest {
	s.BackendPorts = v
	return s
}

func (s *UpdateListenerRequest) SetCertificates(v []*UpdateListenerRequestCertificates) *UpdateListenerRequest {
	s.Certificates = v
	return s
}

func (s *UpdateListenerRequest) SetClientAffinity(v string) *UpdateListenerRequest {
	s.ClientAffinity = &v
	return s
}

func (s *UpdateListenerRequest) SetClientToken(v string) *UpdateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerRequest) SetDescription(v string) *UpdateListenerRequest {
	s.Description = &v
	return s
}

func (s *UpdateListenerRequest) SetHttpVersion(v string) *UpdateListenerRequest {
	s.HttpVersion = &v
	return s
}

func (s *UpdateListenerRequest) SetIdleTimeout(v int32) *UpdateListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateListenerRequest) SetListenerId(v string) *UpdateListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerRequest) SetName(v string) *UpdateListenerRequest {
	s.Name = &v
	return s
}

func (s *UpdateListenerRequest) SetPortRanges(v []*UpdateListenerRequestPortRanges) *UpdateListenerRequest {
	s.PortRanges = v
	return s
}

func (s *UpdateListenerRequest) SetProtocol(v string) *UpdateListenerRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateListenerRequest) SetProxyProtocol(v string) *UpdateListenerRequest {
	s.ProxyProtocol = &v
	return s
}

func (s *UpdateListenerRequest) SetRegionId(v string) *UpdateListenerRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateListenerRequest) SetRequestTimeout(v int32) *UpdateListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *UpdateListenerRequest) SetSecurityPolicyId(v string) *UpdateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateListenerRequest) SetXForwardedForConfig(v *UpdateListenerRequestXForwardedForConfig) *UpdateListenerRequest {
	s.XForwardedForConfig = v
	return s
}

func (s *UpdateListenerRequest) Validate() error {
	if s.BackendPorts != nil {
		for _, item := range s.BackendPorts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.XForwardedForConfig != nil {
		if err := s.XForwardedForConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateListenerRequestBackendPorts struct {
	// The first port in the range of ports that are used by backend servers to receive requests.
	//
	// > This parameter is required only when you configure an HTTPS or HTTP listener and the listener port is different from the service port of the backend servers. In this case, the first port that is used by the backend servers to receive requests must be the same as the last port.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port in the range of ports that are used by backend servers to receive requests.
	//
	// > This parameter is required only when you configure an HTTPS or HTTP listener and the listener port is different from the service port of the backend servers. In this case, the first port that is used by the backend servers to receive requests must be the same as the last port.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateListenerRequestBackendPorts) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerRequestBackendPorts) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequestBackendPorts) GetFromPort() *int32 {
	return s.FromPort
}

func (s *UpdateListenerRequestBackendPorts) GetToPort() *int32 {
	return s.ToPort
}

func (s *UpdateListenerRequestBackendPorts) SetFromPort(v int32) *UpdateListenerRequestBackendPorts {
	s.FromPort = &v
	return s
}

func (s *UpdateListenerRequestBackendPorts) SetToPort(v int32) *UpdateListenerRequestBackendPorts {
	s.ToPort = &v
	return s
}

func (s *UpdateListenerRequestBackendPorts) Validate() error {
	return dara.Validate(s)
}

type UpdateListenerRequestCertificates struct {
	// The ID of the SSL certificate.
	//
	// > This parameter is required only when you configure an HTTPS listener.
	//
	// example:
	//
	// 449****-cn-hangzhou
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateListenerRequestCertificates) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequestCertificates) GetId() *string {
	return s.Id
}

func (s *UpdateListenerRequestCertificates) SetId(v string) *UpdateListenerRequestCertificates {
	s.Id = &v
	return s
}

func (s *UpdateListenerRequestCertificates) Validate() error {
	return dara.Validate(s)
}

type UpdateListenerRequestPortRanges struct {
	// The first port of the listener port range that is used to receive and forward requests to endpoints.
	//
	// Valid values: **1*	- to **65499**. The **FromPort*	- value must be smaller than or equal to the **ToPort*	- value.
	//
	// The maximum number of ports that can be configured varies based on the routing type and protocol of the listener. For more information, see [Listener overview](https://help.aliyun.com/document_detail/153216.html).
	//
	// > You can configure only one listener port for an HTTP or HTTPS listener. In this case, the first port is the same as the last port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port of the listener port range that is used to receive and forward requests to endpoints.
	//
	// Valid values: **1*	- to **65499**. The **FromPort*	- value must be smaller than or equal to the **ToPort*	- value.
	//
	// The maximum number of ports that can be configured varies based on the routing type and protocol of the listener. For more information, see [Listener overview](https://help.aliyun.com/document_detail/153216.html).
	//
	// > You can configure only one listener port for an HTTP or HTTPS listener. In this case, the first port is the same as the last port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateListenerRequestPortRanges) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerRequestPortRanges) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequestPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *UpdateListenerRequestPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *UpdateListenerRequestPortRanges) SetFromPort(v int32) *UpdateListenerRequestPortRanges {
	s.FromPort = &v
	return s
}

func (s *UpdateListenerRequestPortRanges) SetToPort(v int32) *UpdateListenerRequestPortRanges {
	s.ToPort = &v
	return s
}

func (s *UpdateListenerRequestPortRanges) Validate() error {
	return dara.Validate(s)
}

type UpdateListenerRequestXForwardedForConfig struct {
	// Specifies whether to use the `GA-AP` header to retrieve information about acceleration regions. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// > This parameter is available only when you create an HTTPS or HTTP listener.
	//
	// example:
	//
	// false
	XForwardedForGaApEnabled *bool `json:"XForwardedForGaApEnabled,omitempty" xml:"XForwardedForGaApEnabled,omitempty"`
	// Specifies whether to use the `GA-ID` header to retrieve the ID of the GA instance. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// > This parameter is available only when you create an HTTPS or HTTP listener.
	//
	// example:
	//
	// false
	XForwardedForGaIdEnabled *bool `json:"XForwardedForGaIdEnabled,omitempty" xml:"XForwardedForGaIdEnabled,omitempty"`
	// Specifies whether to use the `GA-X-Forward-Port` header to retrieve the listener ports of the GA instance. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// > This parameter is available only when you create an HTTPS or HTTP listener.
	//
	// example:
	//
	// false
	XForwardedForPortEnabled *bool `json:"XForwardedForPortEnabled,omitempty" xml:"XForwardedForPortEnabled,omitempty"`
	// Specifies whether to use the `GA-X-Forward-Proto` header to retrieve the listener protocol of the GA instance. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// > This parameter is available only when you create an HTTPS or HTTP listener.
	//
	// example:
	//
	// false
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Specifies whether to use the `X-Real-IP` header to retrieve client IP addresses. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// > This parameter is available only when you create an HTTPS or HTTP listener.
	//
	// example:
	//
	// false
	XRealIpEnabled *bool `json:"XRealIpEnabled,omitempty" xml:"XRealIpEnabled,omitempty"`
}

func (s UpdateListenerRequestXForwardedForConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequestXForwardedForConfig) GetXForwardedForGaApEnabled() *bool {
	return s.XForwardedForGaApEnabled
}

func (s *UpdateListenerRequestXForwardedForConfig) GetXForwardedForGaIdEnabled() *bool {
	return s.XForwardedForGaIdEnabled
}

func (s *UpdateListenerRequestXForwardedForConfig) GetXForwardedForPortEnabled() *bool {
	return s.XForwardedForPortEnabled
}

func (s *UpdateListenerRequestXForwardedForConfig) GetXForwardedForProtoEnabled() *bool {
	return s.XForwardedForProtoEnabled
}

func (s *UpdateListenerRequestXForwardedForConfig) GetXRealIpEnabled() *bool {
	return s.XRealIpEnabled
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXForwardedForGaApEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XForwardedForGaApEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXForwardedForGaIdEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XForwardedForGaIdEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXForwardedForPortEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XForwardedForPortEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXRealIpEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XRealIpEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) Validate() error {
	return dara.Validate(s)
}
