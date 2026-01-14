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
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The SSL certificates.
	Certificates []*CreateListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// Specifies whether to enable client affinity for the listener.
	//
	// 	- If this parameter is left empty, client affinity is disabled. After client affinity is disabled, requests from a specific client IP address may be forwarded to different endpoints.
	//
	// 	- To enable client affinity, set this parameter to **SOURCE_IP**. In this case, when a client accesses stateful applications, requests from the same client are forwarded to the same endpoint regardless of the source port or protocol.
	//
	// example:
	//
	// SOURCE_IP
	ClientAffinity *string `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- is different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint group that is associated with the custom routing listener.
	//
	// The endpoint groups that are associated with the custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	CustomRoutingEndpointGroupConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurations `json:"CustomRoutingEndpointGroupConfigurations,omitempty" xml:"CustomRoutingEndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The description of the listener.
	//
	// The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Listener
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint groups that are associated with the intelligent routing listener.
	//
	// You can configure up to 10 endpoint groups for an intelligent routing listener.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	EndpointGroupConfigurations []*CreateListenerRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The maximum version of the HTTP protocol. Valid values:
	//
	// 	- **http3**
	//
	// 	- **http2*	- (default)
	//
	// 	- **http1.1**
	//
	// >  Only HTTPS listeners support this parameter.
	//
	// example:
	//
	// http2
	HttpVersion *string `json:"HttpVersion,omitempty" xml:"HttpVersion,omitempty"`
	// The timeout period of idle connections. Unit: seconds.
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
	// The name of the listener.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// Listener
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The listener ports. Valid values: **1*	- to **65499**. The maximum number of ports that can be configured depends on the routing type and protocol of the listener. For more information, see [Listener overview](https://help.aliyun.com/document_detail/153216.html).
	//
	// This parameter is required.
	PortRanges []*CreateListenerRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The network transmission protocol that you want to use for the listener. Valid values:
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
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period for HTTP or HTTPS requests. Unit: seconds.
	//
	// Valid values: 1 to 180. Default value: 60. Unit: seconds.
	//
	// >  This parameter takes effect only for HTTP or HTTPS listeners. If the backend server does not respond within the timeout period, GA returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 15
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The ID of the security policy. Valid values:
	//
	// 	- **tls_cipher_policy_1_0**
	//
	//     	- Supported Transport Layer Security (TLS) versions: TLS 1.0, TLS 1.1, and TLS 1.2
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// 	- **tls_cipher_policy_1_1**
	//
	//     	- Supported TLS versions: TLS 1.1 and TLS 1.2
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// 	- **tls_cipher_policy_1_2**
	//
	//     	- Supported TLS version: TLS 1.2
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
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
	// The routing type of the listener. Valid values:
	//
	// 	- **Standard*	- (default): intelligent routing
	//
	// 	- **CustomRouting**: custom routing
	//
	// > 	- Custom routing listeners are in invitational preview. To use custom routing listeners, contact your account manager.
	//
	// > 	- You can create only listeners of the same routing type for a standard GA instance. You cannot change the routing types of listeners. For more information, see [Listener overview](https://help.aliyun.com/document_detail/153216.html).
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The `XForward` headers.
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
	// > This parameter is required only when you create an HTTPS listener.
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
	// The description of the endpoint group that is associated with the custom routing listener.
	//
	// The description can be up to 200 characters in length and cannot contain `http://` or `https://`.
	//
	// You can specify up to five endpoint group descriptions.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mapping configurations of the endpoint group that is associated with the custom routing listener.
	//
	// You need to specify the port ranges and protocols used by the endpoint group. The ports are mapped to listener ports.
	//
	// You can specify up to 20 mapping configurations for an endpoint group of a custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	DestinationConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	// The endpoints that are associated with the custom routing listener.
	//
	// You can configure up to 10 endpoints for an endpoint group of a custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	EndpointConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The region ID of the endpoint group that is associated with the custom routing listener.
	//
	// You can enter the region IDs of up to five endpoint groups.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The name of the endpoint group that is associated with the custom routing listener.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// You can specify up to five endpoint group names.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
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
	// The start port used by the endpoint group that is associated with the custom routing listener.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be equal to or smaller than the value of **ToPort**.
	//
	// You can specify up to 20 start ports for an endpoint group of a custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The protocol used by the endpoint group that is associated with the custom routing listener.
	//
	// You can specify up to four protocol types for an endpoint group of a custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The end port used by the endpoint group that is associated with the custom routing listener.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be equal to or smaller than the value of **ToPort**.
	//
	// You can specify up to 20 end ports for an endpoint group of a custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
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
	// The name of the vSwitch attached to the endpoint of the custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The destinations in the endpoint that is associated with the custom routing listener.
	//
	// You can specify up to 20 traffic destinations for each endpoint of a custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	PolicyConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The traffic policy for the endpoint that is associated with the custom routing listener. Default value: DenyAll. Valid values:
	//
	// 	- **DenyAll*	- (default): denies all traffic to the specified backend service.
	//
	// 	- **AllowAll**: allows all traffic to the specified backend service.
	//
	// 	- **AllowCustom**: allows traffic only to specified destinations in the endpoint. If you set this parameter to AllowCustom, you must specify IP addresses and port ranges as the destinations to which you want to distribute traffic. If you specify only IP addresses and do not specify port ranges, GA can forward traffic to the specified IP addresses over all destination ports.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	//
	// example:
	//
	// DenyAll
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	// The service type of the endpoint that is associated with the custom routing listener. Default value: PrivateSubNet. Set the value to
	//
	// **PrivateSubNet**, which specifies a private CIDR block.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
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
	// The IP address of the destination.
	//
	// This parameter takes effect only if **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 destination IP addresses for each endpoint of a custom routing listener.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The port ranges of the destination to which traffic is forwarded. The value of this parameter must fall within the port range of the endpoint group.
	//
	// If you do not specify this parameter, traffic is forwarded over all ports.
	//
	// This parameter takes effect only if **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for up to 20 destinations in each endpoint of a custom routing listener. You can specify up to five port ranges for each destination.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
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
	// The start port of the port range. The value of this parameter must fall within the port range of the backend service.
	//
	// This parameter takes effect only if **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for up to 20 destinations in each endpoint of a custom routing listener. You can specify up to five start ports for each destination.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The end port of the destination port range. The value of this parameter must fall within the port range of the backend service.
	//
	// This parameter takes effect only if **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for up to 20 destinations in each endpoint of a custom routing listener. You can specify up to five end ports for each destination.
	//
	// >  You can configure endpoint groups and endpoints for a custom routing listener only if **Type*	- is set to **CustomRouting**.
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
	// The endpoints that are associated with the intelligent routing listener.
	EndpointConfigurations []*CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The description of the endpoint group that is associated with the intelligent routing listener.
	//
	// The description can be up to 200 characters in length and cannot contain `http://` or `https://`.
	//
	// You can enter the descriptions of up to 10 endpoint groups.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// test
	EndpointGroupDescription *string `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	// The name of the endpoint group that is associated with the intelligent routing listener.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// You can enter the names of up to 10 endpoint groups.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// test
	EndpointGroupName *string `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	// The region ID of the endpoint group that is associated with the intelligent routing listener.
	//
	// You can enter the IDs of up to 10 regions.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group associated with the intelligent routing listener. Valid values:
	//
	// 	- **default*	- (default): a default endpoint group.
	//
	// 	- **virtual**: a virtual endpoint group.
	//
	// You can specify up to 10 endpoint group types.
	//
	// > 	- You can configure endpoint groups and endpoints for an intelligent routing listener only if you set **Type*	- to **Standard**.
	//
	// >	- Only HTTP and HTTPS intelligent routing listeners support virtual endpoint groups.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The backend service protocol version of the endpoint that is associated with the intelligent routing listener. Valid values:
	//
	// 	- **HTTP1.1*	- (default)
	//
	// 	- **HTTP2**
	//
	// >  You can specify this parameter only if EndpointRequestProtocol is set to HTTPS.
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The backend service protocol of the endpoint that is associated with the intelligent routing listener. Valid values:
	//
	// 	- **HTTP*	- (default)
	//
	// 	- **HTTPS**
	//
	// You can specify up to 10 backend service protocols.
	//
	// > 	- You can configure endpoint groups and endpoints for an intelligent routing listener only if you set **Type*	- to **Standard**.
	//
	// >	- You can specify this parameter only for HTTP and HTTPS intelligent routing listeners.
	//
	// >	- For an HTTP listener, the backend service protocol must be **HTTP**.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable health checks for the endpoint group. Valid values:
	//
	// 	- **true**: enables the health check feature.
	//
	// 	- **false*	- (default): disables the health check feature.
	//
	// You can enable the health check feature for up to 10 endpoint groups.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// false
	HealthCheckEnabled *bool   `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHost    *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// You can specify up to 10 health check intervals.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int64 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The health check path.
	//
	// You can specify up to 10 health check paths.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port that you want to use for health checks. Valid values: **1*	- to **65535**.
	//
	// You can specify up to 10 ports for health checks.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// 20
	HealthCheckPort *int64 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol over which health check requests are sent. Valid values:
	//
	// 	- **tcp*	- or **TCP**
	//
	// 	- **http*	- or **HTTP**
	//
	// 	- **https*	- or **HTTPS**
	//
	// You can specify up to 10 health check protocols.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// tcp
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The port mappings.
	PortOverrides []*CreateListenerRequestEndpointGroupConfigurationsPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The number of failed consecutive health checks that must occur before a healthy endpoint group is considered unhealthy or the number of successful consecutive health checks that must occur before an unhealthy endpoint group is considered healthy. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// You can specify up to 10 values (the number of consecutive health check successes or consecutive health check failures).
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
	//
	// example:
	//
	// 3
	ThresholdCount *int64 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The traffic distribution ratio. If an intelligent routing listener is associated with multiple endpoint groups, you can configure this parameter to specify the ratio of traffic distributed to each endpoint group.
	//
	// Valid values: **1*	- to **100**. Default value: **100**.
	//
	// You can specify traffic distribution ratios for up to 10 endpoint groups.
	//
	// >  You can configure endpoint groups and endpoints only if you set **Type*	- to **Standard**.
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
	// Specifies whether to automatically preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > 	- By default, client IP address preservation is disabled for an endpoint group that is associated with a UDP or TCP listener. You can configure this parameter based on your business requirements.
	//
	// >	- By default, client IP address preservation is enabled for an endpoint group that is associated with a HTTP or HTTPS listener. Client IP addresses are obtained by using the X-Forwarded-For header. You cannot disable the feature.
	//
	// >	- EnableClientIPPreservation and EnableProxyProtocol cannot be set to true at the same time.
	//
	// >>For more information, see [Preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Specifies whether to use the proxy protocol to preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >	- This parameter is supported only by endpoint groups associated with TCP listeners.
	//
	// >	- EnableClientIPPreservation and EnableProxyProtocol cannot be set to true at the same time.
	//
	// >>For more information, see [Preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address or domain name of the endpoint that is associated with the intelligent routing listener.
	//
	// You can enter the IP addresses or domain names of up to 100 endpoints in an endpoint group that is associated with the intelligent routing listener.
	//
	// >  If you set **Type*	- to **Standard**, you can configure endpoint groups and endpoints, and this parameter is required.
	//
	// example:
	//
	// 47.0.XX.XX
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The private IP address of the ENI.
	//
	// >  This parameter is available only when you set the endpoint type to **ENI**. If you leave this parameter empty, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of the endpoint that is associated with the intelligent routing listener. Valid values:
	//
	// 	- **Domain**: a custom domain name.
	//
	// 	- **Ip**: a custom IP address.
	//
	// 	- **PublicIp**: a public IP address provided by Alibaba Cloud.
	//
	// 	- **ECS**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **SLB**: a Server Load Balancer (SLB) instance.
	//
	// 	- **ALB**: an Application Load Balancer (ALB) instance.
	//
	// 	- **OSS**: an Object Storage Service (OSS) bucket.
	//
	// 	- **ENI**: an elastic network interface (ENI).
	//
	// 	- **NLB**: a Network Load Balancer (NLB) instance.
	//
	// 	- **IpTarget**: a custom private IP address.
	//
	// You can specify up to 100 endpoint types in the endpoint group that is associated with the intelligent routing listener.
	//
	// > 	- If you set **Type*	- to **Standard**, you can configure the endpoint group and endpoint that are associated with the intelligent routing listener. In addition, this parameter is required.
	//
	// >	- If you set this parameter to **ECS**, **ENI**, **SLB**, **ALB**, **NLB**, or **IpTarget*	- and the AliyunServiceRoleForGaVpcEndpoint service-linked role does not exist, the system automatically creates the role.
	//
	// >	- If you set this parameter to **ALB*	- and the AliyunServiceRoleForGaAlb service-linked role does not exist, the system automatically creates the role.
	//
	// >	- If you set this parameter to **OSS*	- and the AliyunServiceRoleForGaOss service-linked role does not exist, the system automatically creates the role.
	//
	// >	- If you set this parameter to **NLB*	- and the AliyunServiceRoleForGaNlb service-linked role does not exist, the system automatically creates the role.
	//
	// >>For more information, see [Service-linked roles](https://help.aliyun.com/document_detail/178360.html).
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IDs of vSwitches that are deployed in the VPC.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID.
	//
	// You can specify one VPC ID for an endpoint group of an intelligent routing listener.
	//
	// >  This parameter is valid and required only if Type is set to **IpTarget**.
	//
	// example:
	//
	// vpc-bp13r1kpr2lel****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint that is associated with the intelligent routing listener.
	//
	// Valid values: **0*	- to **255**.
	//
	// You can specify the weights of up to 100 endpoints for an endpoint group of an intelligent routing listener.
	//
	// > 	- If you set **Type*	- to **Standard**, you can configure the endpoint group and endpoint that are associated with the intelligent routing listener. In addition, this parameter is required.
	//
	// >	- If the weight of an endpoint is set to 0, GA stops distributing network traffic to the endpoint. Proceed with caution.
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

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
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
	// The endpoint port that is mapped to the listener port.
	//
	// You can specify endpoint ports in up to five port mappings.
	//
	// > 	- You can configure endpoint groups and endpoints for an intelligent routing listener only if you set **Type*	- to **Standard**.
	//
	// >	- You cannot configure port mappings for virtual endpoint groups of TCP listeners. If a virtual endpoint group already exists on the listener, you cannot configure port mappings for the default endpoint group. If port mappings are configured for the default endpoint group, you cannot add a virtual endpoint group.
	//
	// >	- If you configure port mappings for a listener, you cannot modify the listener protocol. You can only switch between HTTP and HTTPS.
	//
	// >	- Listener port: When you modify the listener port range, make sure that the port range includes the ports configured in port mappings. For example, if you set the listener port range to 80 to 82 and map the listener ports to endpoint ports 100 to 102, you cannot change the listener port range to 80 to 81.
	//
	// example:
	//
	// 80
	EndpointPort *int64 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The listener port that is mapped to the endpoint port.
	//
	// You can specify listener ports in up to five port mappings.
	//
	// > 	- You can configure endpoint groups and endpoints for an intelligent routing listener only if you set **Type*	- to **Standard**.
	//
	// >	- You cannot configure port mappings for virtual endpoint groups of TCP listeners. If a virtual endpoint group already exists on the listener, you cannot configure port mappings for the default endpoint group. If port mappings are configured for the default endpoint group, you cannot add a virtual endpoint group.
	//
	// >	- If you configure port mappings for a listener, you cannot modify the listener protocol. You can only switch between HTTP and HTTPS.
	//
	// >	- Listener port: When you modify the listener port range, make sure that the port range includes the ports configured in port mappings. For example, if you set the listener port range to 80 to 82 and map the listener ports to endpoint ports 100 to 102, you cannot change the listener port range to 80 to 81.
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
	// The first port of the listener port range that you want to use to receive and forward requests to endpoints.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be smaller than or equal to the value of **ToPort**.
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
	// The last port of the listener port range that you want to use to receive and forward requests to endpoints.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be smaller than or equal to the value of **ToPort**.
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
	// Specifies whether to use the `GA-AP` header to retrieve the information about acceleration regions. Valid values:
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
