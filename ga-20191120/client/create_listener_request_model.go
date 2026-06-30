// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateListenerRequest
	GetAcceleratorId() *string
	SetCertificates(v []*CreateListenerRequestCertificates) *CreateListenerRequest
	GetCertificates() []*CreateListenerRequestCertificates
	SetClientAffinity(v string) *CreateListenerRequest
	GetClientAffinity() *string
	SetClientToken(v string) *CreateListenerRequest
	GetClientToken() *string
	SetCustomRoutingEndpointGroupConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurations) *CreateListenerRequest
	GetCustomRoutingEndpointGroupConfigurations() []*CreateListenerRequestCustomRoutingEndpointGroupConfigurations
	SetDescription(v string) *CreateListenerRequest
	GetDescription() *string
	SetEndpointGroupConfigurations(v []*CreateListenerRequestEndpointGroupConfigurations) *CreateListenerRequest
	GetEndpointGroupConfigurations() []*CreateListenerRequestEndpointGroupConfigurations
	SetHttpVersion(v string) *CreateListenerRequest
	GetHttpVersion() *string
	SetIdleTimeout(v int32) *CreateListenerRequest
	GetIdleTimeout() *int32
	SetName(v string) *CreateListenerRequest
	GetName() *string
	SetPortRanges(v []*CreateListenerRequestPortRanges) *CreateListenerRequest
	GetPortRanges() []*CreateListenerRequestPortRanges
	SetProtocol(v string) *CreateListenerRequest
	GetProtocol() *string
	SetRegionId(v string) *CreateListenerRequest
	GetRegionId() *string
	SetRequestTimeout(v int32) *CreateListenerRequest
	GetRequestTimeout() *int32
	SetSecurityPolicyId(v string) *CreateListenerRequest
	GetSecurityPolicyId() *string
	SetType(v string) *CreateListenerRequest
	GetType() *string
	SetXForwardedForConfig(v *CreateListenerRequestXForwardedForConfig) *CreateListenerRequest
	GetXForwardedForConfig() *CreateListenerRequestXForwardedForConfig
}

type CreateListenerRequest struct {
	// The ID of the Global Accelerator instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The SSL certificates for an HTTPS listener.
	Certificates []*CreateListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The client affinity for the listener.
	//
	// - By default, client affinity is disabled, and requests from the same client may be routed to different endpoints.
	//
	// - Set to **SOURCE_IP*	- to enable client affinity. This setting directs all requests from the same client to the same endpoint, regardless of the source port or protocol.
	//
	// example:
	//
	// SOURCE_IP
	ClientAffinity *string `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	// A client token that ensures the idempotence of the request.
	//
	// Generate a unique token on your client for each request. The token must contain only ASCII characters.
	//
	// > If you omit this parameter, the system uses the request\\"s **RequestId*	- as the **ClientToken**.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configurations of the endpoint groups for a custom routing listener.
	//
	// You can specify up to five endpoint groups.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	CustomRoutingEndpointGroupConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurations `json:"CustomRoutingEndpointGroupConfigurations,omitempty" xml:"CustomRoutingEndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The description of the listener.
	//
	// The description can be up to 200 characters long and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Listener
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configurations of the endpoint groups for a standard listener.
	//
	// You can specify up to 10 endpoint groups.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	EndpointGroupConfigurations []*CreateListenerRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The maximum HTTP version. Valid values:
	//
	// - **http3**: HTTP/3
	//
	// - **http2*	- (default): HTTP/2
	//
	// - **http1.1**: HTTP/1.1
	//
	// > This parameter applies only to HTTPS listeners.
	//
	// example:
	//
	// http2
	HttpVersion *string `json:"HttpVersion,omitempty" xml:"HttpVersion,omitempty"`
	// The connection idle timeout, in seconds.
	//
	// - TCP: 10–900 seconds. Default: 900 seconds.
	//
	// - UDP: 10–20 seconds. Default: 20 seconds.
	//
	// - HTTP/HTTPS: 1–60 seconds. Default: 15 seconds.
	//
	// example:
	//
	// 900
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 1 to 128 characters long, start with a letter or a Chinese character, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Listener
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The listener port range. The port numbers must be within the range of **1*	- to **65499**. The maximum number of allowed ports depends on the listener\\"s routing type and protocol. For more information, see [Listener ports](https://help.aliyun.com/document_detail/153216.html).
	//
	// This parameter is required.
	PortRanges []*CreateListenerRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The listener\\"s network protocol. Valid values:
	//
	// - **tcp**: TCP.
	//
	// - **udp**: UDP.
	//
	// - **http**: HTTP.
	//
	// - **https**: HTTPS.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID of the Global Accelerator instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request timeout for HTTP/HTTPS connections, in seconds.
	//
	// Valid values: 1–180 seconds. Default: 60 seconds.
	//
	// > This parameter applies only to HTTP or HTTPS listeners. If the backend server does not respond within the timeout period, Global Accelerator returns an HTTP 504 error to the client.
	//
	// example:
	//
	// 15
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The ID of the security policy. Valid values:
	//
	// - **tls_cipher_policy_1_0**
	//
	//   - Supported TLS versions: TLS 1.0, TLS 1.1, and TLS 1.2.
	//
	//   - Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// - **tls_cipher_policy_1_1**
	//
	//   - Supported TLS versions: TLS 1.1 and TLS 1.2.
	//
	//   - Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// - **tls_cipher_policy_1_2**
	//
	//   - Supported TLS version: TLS 1.2.
	//
	//   - Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// - **tls_cipher_policy_1_2_strict**
	//
	//   - Supported TLS version: TLS 1.2.
	//
	//   - Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA.
	//
	// - **tls_cipher_policy_1_2_strict_with_1_3**
	//
	//   - Supported TLS versions: TLS 1.2 and TLS 1.3.
	//
	//   - Supported cipher suites: TLS_AES_128_GCM_SHA256, TLS_AES_256_GCM_SHA384, TLS_CHACHA20_POLY1305_SHA256, TLS_AES_128_CCM_SHA256, TLS_AES_128_CCM_8_SHA256, ECDHE-ECDSA-AES128-GCM-SHA256, ECDHE-ECDSA-AES256-GCM-SHA384, ECDHE-ECDSA-AES128-SHA256, ECDHE-ECDSA-AES256-SHA384, ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-ECDSA-AES128-SHA, ECDHE-ECDSA-AES256-SHA, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA.
	//
	// > This parameter applies only to HTTPS listeners.
	//
	// example:
	//
	// tls_cipher_policy_1_0
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The routing type of the listener. Valid values:
	//
	// - **Standard*	- (default): standard routing.
	//
	// - **CustomRouting**: custom routing.
	//
	// > 	- Custom routing is in invitation-only preview. To use this feature, contact your Alibaba Cloud account manager.
	//
	// >
	//
	// > 	- A standard Global Accelerator instance supports only one routing type for all of its listeners. The routing type cannot be changed after the listener is created. For more information, see [Listener overview](https://help.aliyun.com/document_detail/153216.html).
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Settings for `X-Forwarded-For` related headers.
	XForwardedForConfig *CreateListenerRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s CreateListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateListenerRequest) GetCertificates() []*CreateListenerRequestCertificates {
	return s.Certificates
}

func (s *CreateListenerRequest) GetClientAffinity() *string {
	return s.ClientAffinity
}

func (s *CreateListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateListenerRequest) GetCustomRoutingEndpointGroupConfigurations() []*CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	return s.CustomRoutingEndpointGroupConfigurations
}

func (s *CreateListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateListenerRequest) GetEndpointGroupConfigurations() []*CreateListenerRequestEndpointGroupConfigurations {
	return s.EndpointGroupConfigurations
}

func (s *CreateListenerRequest) GetHttpVersion() *string {
	return s.HttpVersion
}

func (s *CreateListenerRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *CreateListenerRequest) GetName() *string {
	return s.Name
}

func (s *CreateListenerRequest) GetPortRanges() []*CreateListenerRequestPortRanges {
	return s.PortRanges
}

func (s *CreateListenerRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateListenerRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *CreateListenerRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *CreateListenerRequest) GetType() *string {
	return s.Type
}

func (s *CreateListenerRequest) GetXForwardedForConfig() *CreateListenerRequestXForwardedForConfig {
	return s.XForwardedForConfig
}

func (s *CreateListenerRequest) SetAcceleratorId(v string) *CreateListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateListenerRequest) SetCertificates(v []*CreateListenerRequestCertificates) *CreateListenerRequest {
	s.Certificates = v
	return s
}

func (s *CreateListenerRequest) SetClientAffinity(v string) *CreateListenerRequest {
	s.ClientAffinity = &v
	return s
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetCustomRoutingEndpointGroupConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurations) *CreateListenerRequest {
	s.CustomRoutingEndpointGroupConfigurations = v
	return s
}

func (s *CreateListenerRequest) SetDescription(v string) *CreateListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateListenerRequest) SetEndpointGroupConfigurations(v []*CreateListenerRequestEndpointGroupConfigurations) *CreateListenerRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *CreateListenerRequest) SetHttpVersion(v string) *CreateListenerRequest {
	s.HttpVersion = &v
	return s
}

func (s *CreateListenerRequest) SetIdleTimeout(v int32) *CreateListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetName(v string) *CreateListenerRequest {
	s.Name = &v
	return s
}

func (s *CreateListenerRequest) SetPortRanges(v []*CreateListenerRequestPortRanges) *CreateListenerRequest {
	s.PortRanges = v
	return s
}

func (s *CreateListenerRequest) SetProtocol(v string) *CreateListenerRequest {
	s.Protocol = &v
	return s
}

func (s *CreateListenerRequest) SetRegionId(v string) *CreateListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateListenerRequest) SetRequestTimeout(v int32) *CreateListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetSecurityPolicyId(v string) *CreateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateListenerRequest) SetType(v string) *CreateListenerRequest {
	s.Type = &v
	return s
}

func (s *CreateListenerRequest) SetXForwardedForConfig(v *CreateListenerRequestXForwardedForConfig) *CreateListenerRequest {
	s.XForwardedForConfig = v
	return s
}

func (s *CreateListenerRequest) Validate() error {
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomRoutingEndpointGroupConfigurations != nil {
		for _, item := range s.CustomRoutingEndpointGroupConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EndpointGroupConfigurations != nil {
		for _, item := range s.EndpointGroupConfigurations {
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

type CreateListenerRequestCertificates struct {
	// The ID of the SSL certificate.
	//
	// > This parameter is required only for HTTPS listeners.
	//
	// example:
	//
	// 449****-cn-hangzhou
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateListenerRequestCertificates) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCertificates) GetId() *string {
	return s.Id
}

func (s *CreateListenerRequestCertificates) SetId(v string) *CreateListenerRequestCertificates {
	s.Id = &v
	return s
}

func (s *CreateListenerRequestCertificates) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurations struct {
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters long and cannot contain `http://` or `https://`.
	//
	// You can enter up to 5 endpoint group descriptions.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mapping configurations for the endpoint group.
	//
	// You must specify the port ranges and protocols for the backend service. The settings are mapped to the associated listener port ranges.
	//
	// You can specify up to 20 mapping configurations for each endpoint group.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	DestinationConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	// The endpoint configurations.
	//
	// You can specify up to 10 endpoints for each endpoint group.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	EndpointConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the region where the endpoint group is created.
	//
	// You can enter up to 5 endpoint group region IDs.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters long, start with a letter or a Chinese character, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// You can enter up to 5 endpoint group names.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) GetDescription() *string {
	return s.Description
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) GetDestinationConfigurations() []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations {
	return s.DestinationConfigurations
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) GetEndpointConfigurations() []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) GetName() *string {
	return s.Name
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetDescription(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.Description = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetDestinationConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.DestinationConfigurations = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetName(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.Name = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) Validate() error {
	if s.DestinationConfigurations != nil {
		for _, item := range s.DestinationConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EndpointConfigurations != nil {
		for _, item := range s.EndpointConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations struct {
	// The first port of the backend service.
	//
	// The valid port range is **1*	- to **65499**. The value of **FromPort*	- must be less than or equal to the value of **ToPort**.
	//
	// In each endpoint group for a custom routing type listener, you can enter up to 20 backend service starting ports.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The protocols of the backend service.
	//
	// You can specify up to four backend service protocols for each mapping configuration.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The last port of the backend service.
	//
	// The valid port range is **1*	- to **65499**. The value of **FromPort*	- must be less than or equal to the value of **ToPort**.
	//
	// In each endpoint group of a listener of the custom routing type, you can enter a maximum of 20 backend service ports.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) GetFromPort() *int32 {
	return s.FromPort
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) GetProtocols() []*string {
	return s.Protocols
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) GetToPort() *int32 {
	return s.ToPort
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) SetFromPort(v int32) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations {
	s.FromPort = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) SetProtocols(v []*string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations {
	s.Protocols = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) SetToPort(v int32) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations {
	s.ToPort = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations struct {
	// The vSwitch of the custom routing listener.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The destination configurations for a custom routing listener.
	//
	// You can specify up to 20 destinations for each endpoint.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	PolicyConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The traffic policy for the backend service of a custom routing listener. Valid values:
	//
	// - **DenyAll*	- (default): Denies all traffic to the specified backend service.
	//
	// - **AllowAll**: Allows all traffic to the specified backend service.
	//
	// - **AllowCustom**: Allows traffic to specific destinations.
	//
	//   You must specify the IP addresses and port ranges of the allowed destinations. If no port range is specified, all ports of the destination are allowed.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// DenyAll
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	// The type of the backend service for a custom routing listener. Valid value:
	//
	// **PrivateSubNet*	- (default): a private CIDR block.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// PrivateSubNet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) GetPolicyConfigurations() []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	return s.PolicyConfigurations
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) GetTrafficToEndpointPolicy() *string {
	return s.TrafficToEndpointPolicy
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) SetPolicyConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	s.PolicyConfigurations = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) SetTrafficToEndpointPolicy(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) Validate() error {
	if s.PolicyConfigurations != nil {
		for _, item := range s.PolicyConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations struct {
	// The IP address of the destination that is allowed to receive traffic.
	//
	// This parameter is required only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 destination IP addresses for each endpoint.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The port range of the destination that is allowed to receive traffic. The port range must be within the port range of the backend service.
	//
	// If you leave this parameter empty, all ports of the destination are allowed.
	//
	// This parameter is required only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 port ranges for each endpoint, and up to 5 port ranges for each destination.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	PortRanges []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) GetAddress() *string {
	return s.Address
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) GetPortRanges() []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	return s.PortRanges
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) SetAddress(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) SetPortRanges(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	s.PortRanges = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) Validate() error {
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges struct {
	// The first port of the destination that is allowed to receive traffic. The port must be within the port range of the backend service.
	//
	// This parameter is required only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 port ranges for each endpoint, and up to 5 first ports for each destination.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port of the destination that is allowed to receive traffic. The port must be within the port range of the backend service.
	//
	// This parameter is required only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 port ranges for each endpoint, and up to 5 last ports for each destination.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **CustomRouting**.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) SetFromPort(v int32) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) SetToPort(v int32) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestEndpointGroupConfigurations struct {
	// The endpoint configurations.
	EndpointConfigurations []*CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters long and cannot contain `http://` or `https://`.
	//
	// You can enter up to 10 endpoint group descriptions.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// test
	EndpointGroupDescription *string `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters long, start with a letter or a Chinese character, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// You can enter up to 10 endpoint group names.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// test
	EndpointGroupName *string `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	// The ID of the region where the endpoint group is created.
	//
	// You can enter up to 10 endpoint group region IDs.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group. Valid values:
	//
	// - **default*	- (default): a default endpoint group.
	//
	// - **virtual**: a virtual endpoint group.
	//
	// You can enter up to 10 endpoint group types.
	//
	// > - This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// >
	//
	// > - You can create virtual endpoint groups only for HTTP or HTTPS listeners.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	// The IP version used by the backend service. Valid values:
	//
	// - **IPv4*	- (default): GA uses only IPv4 addresses to communicate with backend services.
	//
	// - **IPv6**: GA uses only IPv6 addresses to communicate with backend services.
	//
	// - **ProtocolAffinity**: GA uses the same IP version as the client request to communicate with backend services.
	//
	// example:
	//
	// IPv4
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The protocol version of the backend service. Valid values:
	//
	// - **HTTP1.1*	- (default): HTTP/1.1
	//
	// - **HTTP2**: HTTP/2
	//
	// > This parameter is available only when EndpointRequestProtocol is set to HTTPS.
	//
	// example:
	//
	// HTTP1.1
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The protocol used by the backend service. Valid values:
	//
	// - **HTTP*	- (default)
	//
	// - **HTTPS**
	//
	// You can enter up to 10 backend service protocols.
	//
	// > - This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// >
	//
	// > - You can configure this parameter only for endpoint groups of HTTP or HTTPS listeners.
	//
	// >
	//
	// > - For an HTTP listener, the backend service protocol must be **HTTP**.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable health checks for the endpoint group. Valid values:
	//
	// - **true**: Enables health checks.
	//
	// - **false*	- (Default): Disables health checks.
	//
	// You can enable health checks for up to 10 endpoint groups.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// false
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The domain name that is used for health checks.
	//
	// example:
	//
	// www.taobao.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The health check interval, in seconds.
	//
	// You can enter up to 10 health check intervals.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int64 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path to which health check requests are sent.
	//
	// You can enter up to 10 health check paths.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**.
	//
	// You can enter a maximum of 10 ports for health checks.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// 20
	HealthCheckPort *int64 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol over which health check requests are sent. Valid values:
	//
	// - **tcp*	- or **TCP**: TCP
	//
	// - **http*	- or **HTTP**: HTTP
	//
	// - **https*	- or **HTTPS**: HTTPS
	//
	// You can enter up to 10 health check protocols.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// tcp
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The port mapping. You can specify up to five port mappings.
	PortOverrides []*CreateListenerRequestEndpointGroupConfigurationsPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The number of consecutive successful health checks required to mark an endpoint as healthy, or consecutive failed health checks to mark an endpoint as unhealthy.
	//
	// Valid values: **2*	- to **10**. Default value: **3**.
	//
	// You can enter up to 10 values for the number of consecutive health checks required to trigger a health status change.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// 3
	ThresholdCount *int64 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The traffic distribution ratio. If a standard listener is associated with multiple endpoint groups, this parameter specifies the percentage of traffic that is distributed to each endpoint group.
	//
	// Valid values: **1*	- to **100**. Default value: **100**.
	//
	// You can enter traffic distribution values for up to 10 endpoint groups.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// 100
	TrafficPercentage *int64 `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s CreateListenerRequestEndpointGroupConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetEndpointConfigurations() []*CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetEndpointGroupDescription() *string {
	return s.EndpointGroupDescription
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetEndpointGroupName() *string {
	return s.EndpointGroupName
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetEndpointGroupType() *string {
	return s.EndpointGroupType
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetEndpointIpVersion() *string {
	return s.EndpointIpVersion
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetEndpointProtocolVersion() *string {
	return s.EndpointProtocolVersion
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetEndpointRequestProtocol() *string {
	return s.EndpointRequestProtocol
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetHealthCheckIntervalSeconds() *int64 {
	return s.HealthCheckIntervalSeconds
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetHealthCheckPort() *int64 {
	return s.HealthCheckPort
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetPortOverrides() []*CreateListenerRequestEndpointGroupConfigurationsPortOverrides {
	return s.PortOverrides
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetThresholdCount() *int64 {
	return s.ThresholdCount
}

func (s *CreateListenerRequestEndpointGroupConfigurations) GetTrafficPercentage() *int64 {
	return s.TrafficPercentage
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointGroupDescription(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointGroupDescription = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointGroupName(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointGroupName = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointGroupType(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointGroupType = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointIpVersion(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointIpVersion = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointProtocolVersion(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointProtocolVersion = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointRequestProtocol(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckEnabled(v bool) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckHost(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckIntervalSeconds(v int64) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckPath(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckPort(v int64) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckProtocol(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetPortOverrides(v []*CreateListenerRequestEndpointGroupConfigurationsPortOverrides) *CreateListenerRequestEndpointGroupConfigurations {
	s.PortOverrides = v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetThresholdCount(v int64) *CreateListenerRequestEndpointGroupConfigurations {
	s.ThresholdCount = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetTrafficPercentage(v int64) *CreateListenerRequestEndpointGroupConfigurations {
	s.TrafficPercentage = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) Validate() error {
	if s.EndpointConfigurations != nil {
		for _, item := range s.EndpointConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PortOverrides != nil {
		for _, item := range s.PortOverrides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	ApiKeys []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	// Specifies whether to preserve client source IP addresses. Valid values:
	//
	// - **true**: enables the feature.
	//
	// - **false*	- (default): disables the feature.
	//
	// > 	- This feature is disabled by default for endpoint groups of TCP or UDP listeners. You can enable it as needed.
	//
	// >
	//
	// > 	- This feature is enabled by default for endpoint groups of HTTP or HTTPS listeners. Client source IP addresses are retrieved from the `X-Forwarded-For` header. You cannot disable this feature.
	//
	// >
	//
	// > 	- You cannot set both `EnableClientIPPreservation` and `EnableProxyProtocol` to `true`.
	//
	// >
	//
	// > 	- For more information, see [Preserve client source IP addresses](https://help.aliyun.com/document_detail/158080.html).
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Specifies whether to use the proxy protocol to preserve client source IP addresses. Valid values:
	//
	// - **true**: enables the feature.
	//
	// - **false*	- (default): disables the feature.
	//
	// > 	- You can configure this parameter only for endpoint groups of TCP listeners.
	//
	// >
	//
	// > 	- You cannot set both `EnableClientIPPreservation` and `EnableProxyProtocol` to `true`.
	//
	// >
	//
	// > 	- For more information, see [Preserve client source IP addresses](https://help.aliyun.com/document_detail/158080.html).
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address or domain name of the endpoint.
	//
	// In an endpoint group of an intelligent routing listener, you can enter up to 100 IP addresses or domain names of endpoints.
	//
	// > This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// example:
	//
	// 47.0.XX.XX
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// BAILIAN
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The private IP address of the ENI.
	//
	// > If the endpoint type is **ENI**, you can specify this parameter. If you do not specify this parameter, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// - **Domain**: a custom domain name.
	//
	// - **Ip**: a custom IP address.
	//
	// - **PublicIp**: a public IP address of an Alibaba Cloud service.
	//
	// - **ECS**: an Elastic Compute Service (ECS) instance.
	//
	// - **SLB**: a Server Load Balancer (SLB) instance.
	//
	// - **ALB**: an Application Load Balancer (ALB) instance.
	//
	// - **OSS**: an Object Storage Service (OSS) bucket.
	//
	// - **ENI**: an elastic network interface (ENI).
	//
	// - **NLB**: a Network Load Balancer (NLB) instance.
	//
	// - **IpTarget**: a custom private IP address.
	//
	// You can specify up to 100 endpoints in an endpoint group.
	//
	// > - This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// >
	//
	// > - When you add endpoints, Global Accelerator may create service-linked roles to access your resources. The role created depends on the endpoint type:
	//
	// >
	//
	// > -
	//
	// >
	//
	// > -
	//
	// >
	//
	// > -
	//
	// >
	//
	// > > For more information, see [Service-linked roles](https://help.aliyun.com/document_detail/178360.html).
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The list of vSwitches in the VPC. You can specify up to two vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the Virtual Private Cloud (VPC).
	//
	// In an endpoint group of an intelligent routing listener, you can enter a maximum of 1 VPC ID.
	//
	// > This parameter is required only for **IpTarget*	- endpoints.
	//
	// example:
	//
	// vpc-bp13r1kpr2lel****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// Valid values: **0*	- to **255**.
	//
	// In an endpoint group for an intelligent routing type listener, you can enter weights for up to 100 endpoints.
	//
	// > - This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// >
	//
	// > - If an endpoint\\"s weight is set to 0, Global Accelerator stops sending traffic to it. Use this setting with caution.
	//
	// example:
	//
	// 20
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetProvider() *string {
	return s.Provider
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetSubAddress() *string {
	return s.SubAddress
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetWeight() *int64 {
	return s.Weight
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetApiKeys(v []*string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.ApiKeys = v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetEnableClientIPPreservation(v bool) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetEnableProxyProtocol(v bool) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.EnableProxyProtocol = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetProvider(v string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Provider = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetSubAddress(v string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.SubAddress = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetVSwitchIds(v []*string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.VSwitchIds = v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetVpcId(v string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.VpcId = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetWeight(v int64) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestEndpointGroupConfigurationsPortOverrides struct {
	// The endpoint port that is specified in the port mapping.
	//
	// You can enter a maximum of 5 endpoint ports for port mapping.
	//
	// > - This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// >
	//
	// > - For TCP listeners, you cannot configure a port mapping for a virtual endpoint group. If a virtual endpoint group already exists for the listener, you cannot configure a port mapping for the default endpoint group. If a port mapping is configured for the default endpoint group, you cannot add a virtual endpoint group to the listener.
	//
	// >
	//
	// > - After you configure a port mapping, you cannot modify the listener protocol, except for switching between HTTP and HTTPS.
	//
	// >
	//
	// > - When you modify the listener port range, make sure that the new port range includes all listener ports that are specified in the port mapping. For example, if the listener port range is 80-82 and the listener ports are mapped to the endpoint ports 100-102, you cannot change the listener port range to 80-81.
	//
	// example:
	//
	// 80
	EndpointPort *int64 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The listener port that is specified in the port mapping.
	//
	// You can enter up to 5 listener ports for port mappings.
	//
	// > - This parameter applies only when the listener\\"s routing type (**Type**) is **Standard**.
	//
	// >
	//
	// > - For TCP listeners, you cannot configure a port mapping for a virtual endpoint group. If a virtual endpoint group already exists for the listener, you cannot configure a port mapping for the default endpoint group. If a port mapping is configured for the default endpoint group, you cannot add a virtual endpoint group to the listener.
	//
	// >
	//
	// > - After you configure a port mapping, you cannot modify the listener protocol, except for switching between HTTP and HTTPS.
	//
	// >
	//
	// > - When you modify the listener port range, make sure that the new port range includes all listener ports that are specified in the port mapping. For example, if the listener port range is 80-82 and the listener ports are mapped to the endpoint ports 100-102, you cannot change the listener port range to 80-81.
	//
	// example:
	//
	// 443
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s CreateListenerRequestEndpointGroupConfigurationsPortOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestEndpointGroupConfigurationsPortOverrides) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestEndpointGroupConfigurationsPortOverrides) GetEndpointPort() *int64 {
	return s.EndpointPort
}

func (s *CreateListenerRequestEndpointGroupConfigurationsPortOverrides) GetListenerPort() *int64 {
	return s.ListenerPort
}

func (s *CreateListenerRequestEndpointGroupConfigurationsPortOverrides) SetEndpointPort(v int64) *CreateListenerRequestEndpointGroupConfigurationsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsPortOverrides) SetListenerPort(v int64) *CreateListenerRequestEndpointGroupConfigurationsPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsPortOverrides) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestPortRanges struct {
	// The first port in the listener range used to receive and forward requests to endpoints.
	//
	// The port number must be in the range of **1*	- to **65499**, and the value of **FromPort*	- must be less than or equal to the value of **ToPort**.
	//
	// > For HTTP or HTTPS listeners, you can specify only one listener port. In this case, the value of **FromPort*	- must be the same as the value of **ToPort**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port in the listener range used to receive and forward requests to endpoints.
	//
	// The port number must be in the range of **1*	- to **65499**, and the value of **FromPort*	- must be less than or equal to the value of **ToPort**.
	//
	// > For HTTP or HTTPS listeners, you can specify only one listener port. In this case, the value of **FromPort*	- must be the same as the value of **ToPort**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateListenerRequestPortRanges) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestPortRanges) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *CreateListenerRequestPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *CreateListenerRequestPortRanges) SetFromPort(v int32) *CreateListenerRequestPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateListenerRequestPortRanges) SetToPort(v int32) *CreateListenerRequestPortRanges {
	s.ToPort = &v
	return s
}

func (s *CreateListenerRequestPortRanges) Validate() error {
	return dara.Validate(s)
}

type CreateListenerRequestXForwardedForConfig struct {
	// Specifies whether to use the `GA-AP` header to pass information about the acceleration region to the backend server. Valid values:
	//
	// - **true**
	//
	// - **false*	- (Default)
	//
	// > This parameter applies only to HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForGaApEnabled *bool `json:"XForwardedForGaApEnabled,omitempty" xml:"XForwardedForGaApEnabled,omitempty"`
	// Specifies whether to use the `GA-ID` header to pass the Global Accelerator instance ID to the backend server. Valid values:
	//
	// - **true**
	//
	// - **false*	- (Default)
	//
	// > This parameter applies only to HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForGaIdEnabled *bool `json:"XForwardedForGaIdEnabled,omitempty" xml:"XForwardedForGaIdEnabled,omitempty"`
	// Specifies whether to use the `GA-X-Forward-Port` header to pass the listener port of the Global Accelerator instance to the backend server. Valid values:
	//
	// - **true**
	//
	// - **false*	- (Default)
	//
	// > This parameter applies only to HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForPortEnabled *bool `json:"XForwardedForPortEnabled,omitempty" xml:"XForwardedForPortEnabled,omitempty"`
	// Specifies whether to use the `GA-X-Forward-Proto` header to pass the listener protocol of the Global Accelerator instance to the backend server. Valid values:
	//
	// - **true**
	//
	// - **false*	- (Default)
	//
	// > This parameter applies only to HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Specifies whether to use the `X-Real-IP` header to pass the client\\"s real IP address to the backend server. Valid values:
	//
	// - **true**
	//
	// - **false*	- (Default)
	//
	// > This parameter applies only to HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XRealIpEnabled *bool `json:"XRealIpEnabled,omitempty" xml:"XRealIpEnabled,omitempty"`
}

func (s CreateListenerRequestXForwardedForConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForGaApEnabled() *bool {
	return s.XForwardedForGaApEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForGaIdEnabled() *bool {
	return s.XForwardedForGaIdEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForPortEnabled() *bool {
	return s.XForwardedForPortEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXForwardedForProtoEnabled() *bool {
	return s.XForwardedForProtoEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) GetXRealIpEnabled() *bool {
	return s.XRealIpEnabled
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForGaApEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForGaApEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForGaIdEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForGaIdEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForPortEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForPortEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXRealIpEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XRealIpEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) Validate() error {
	return dara.Validate(s)
}
