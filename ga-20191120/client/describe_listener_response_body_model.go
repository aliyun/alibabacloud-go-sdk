// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeListenerResponseBody
	GetAcceleratorId() *string
	SetAclType(v string) *DescribeListenerResponseBody
	GetAclType() *string
	SetBackendPorts(v []*DescribeListenerResponseBodyBackendPorts) *DescribeListenerResponseBody
	GetBackendPorts() []*DescribeListenerResponseBodyBackendPorts
	SetCertificates(v []*DescribeListenerResponseBodyCertificates) *DescribeListenerResponseBody
	GetCertificates() []*DescribeListenerResponseBodyCertificates
	SetClientAffinity(v string) *DescribeListenerResponseBody
	GetClientAffinity() *string
	SetCreateTime(v string) *DescribeListenerResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeListenerResponseBody
	GetDescription() *string
	SetHttpVersion(v string) *DescribeListenerResponseBody
	GetHttpVersion() *string
	SetIdleTimeout(v int32) *DescribeListenerResponseBody
	GetIdleTimeout() *int32
	SetListenerId(v string) *DescribeListenerResponseBody
	GetListenerId() *string
	SetName(v string) *DescribeListenerResponseBody
	GetName() *string
	SetPortRanges(v []*DescribeListenerResponseBodyPortRanges) *DescribeListenerResponseBody
	GetPortRanges() []*DescribeListenerResponseBodyPortRanges
	SetProtocol(v string) *DescribeListenerResponseBody
	GetProtocol() *string
	SetProxyProtocol(v bool) *DescribeListenerResponseBody
	GetProxyProtocol() *bool
	SetRelatedAcls(v []*DescribeListenerResponseBodyRelatedAcls) *DescribeListenerResponseBody
	GetRelatedAcls() []*DescribeListenerResponseBodyRelatedAcls
	SetRequestId(v string) *DescribeListenerResponseBody
	GetRequestId() *string
	SetRequestTimeout(v int32) *DescribeListenerResponseBody
	GetRequestTimeout() *int32
	SetSecurityPolicyId(v string) *DescribeListenerResponseBody
	GetSecurityPolicyId() *string
	SetServiceId(v string) *DescribeListenerResponseBody
	GetServiceId() *string
	SetServiceManaged(v bool) *DescribeListenerResponseBody
	GetServiceManaged() *bool
	SetServiceManagedInfos(v []*DescribeListenerResponseBodyServiceManagedInfos) *DescribeListenerResponseBody
	GetServiceManagedInfos() []*DescribeListenerResponseBodyServiceManagedInfos
	SetState(v string) *DescribeListenerResponseBody
	GetState() *string
	SetType(v string) *DescribeListenerResponseBody
	GetType() *string
	SetXForwardedForConfig(v *DescribeListenerResponseBodyXForwardedForConfig) *DescribeListenerResponseBody
	GetXForwardedForConfig() *DescribeListenerResponseBodyXForwardedForConfig
}

type DescribeListenerResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The type of the ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. Whitelists are suitable for scenarios in which you want to allow only specific IP addresses to access an application. If a whitelist is improperly configured, risks may arise. After a whitelist is configured for a listener, only requests from the IP addresses that are added to the whitelist are distributed by the listener. If the whitelist is enabled but no IP addresses are added to the ACL, the listener does not forward requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are blocked. Blacklists are suitable for scenarios in which you want to deny access from specific IP addresses to an application. If the blacklist is enabled but no IP addresses are added to the ACL, the listener forwards all requests.
	//
	// This parameter is returned only if the value of **Status*	- is **on**.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The information about the backend ports.
	BackendPorts []*DescribeListenerResponseBodyBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The SSL certificates.
	Certificates []*DescribeListenerResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// Indicates whether client affinity is enabled for the listener.
	//
	// 	- If **NONE*	- is returned, client affinity is disabled. Requests from the same client may be forwarded to different endpoints.
	//
	// 	- If **SOURCE_IP*	- is returned, client affinity is enabled. When a client accesses stateful applications, requests from the same client are forwarded to the same endpoint regardless of the source port or protocol.
	//
	// example:
	//
	// SOURCE_IP
	ClientAffinity *string `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	// The time when the listener was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1577786252000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the listener.
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
	// >  This parameter is returned only for HTTPS listeners.
	//
	// example:
	//
	// http2
	HttpVersion *string `json:"HttpVersion,omitempty" xml:"HttpVersion,omitempty"`
	// The timeout period of idle connections. Unit: seconds.
	//
	// example:
	//
	// 900
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The name of the listener.
	//
	// example:
	//
	// Listener
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the listener ports.
	PortRanges []*DescribeListenerResponseBodyPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The network transmission protocol that is used by the listener. Valid values:
	//
	// 	- **tcp**: TCP.
	//
	// 	- **udp**: UDP.
	//
	// 	- **http**: HTTP.
	//
	// 	- **https**: HTTPS.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Deprecated
	//
	// Indicates whether the client IP address preservation feature is enabled. Valid values:
	//
	// 	- **true*	- You can view the source IP addresses of clients over the backend service.
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ProxyProtocol *bool `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	// The information about the access control list (ACL) that is associated with the listener.
	RelatedAcls []*DescribeListenerResponseBodyRelatedAcls `json:"RelatedAcls,omitempty" xml:"RelatedAcls,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timeout period of HTTP or HTTPS requests. Unit: seconds.
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners. If no responses are received from the backend server within the specified timeout period, GA returns the HTTP 504 error code to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The ID of the security policy.
	//
	// 	- **tls_cipher_policy_1_0**
	//
	//     	- Supported Transport Layer Security (TLS) versions: TLS 1.0, TLS 1.1, and TLS 1.2.
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// 	- **tls_cipher_policy_1_1**
	//
	//     	- Supported TLS versions: TLS 1.1 and TLS 1.2.
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// 	- **tls_cipher_policy_1_2**
	//
	//     	- Supported TLS version: TLS 1.2.
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, AES128-GCM-SHA256, AES256-GCM-SHA384, AES128-SHA256, AES256-SHA256, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, AES128-SHA, AES256-SHA, and DES-CBC3-SHA.
	//
	// 	- **tls_cipher_policy_1_2_strict**
	//
	//     	- Supported TLS version: TLS 1.2.
	//
	//     	- Supported cipher suites: ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA.
	//
	// 	- **tls_cipher_policy_1_2_strict_with_1_3**
	//
	//     	- Supported TLS versions: TLS 1.2 and TLS 1.3.
	//
	//     	- Supported cipher suites: TLS_AES_128_GCM_SHA256, TLS_AES_256_GCM_SHA384, TLS_CHACHA20_POLY1305_SHA256, TLS_AES_128_CCM_SHA256, TLS_AES_128_CCM_8_SHA256, ECDHE-ECDSA-AES128-GCM-SHA256, ECDHE-ECDSA-AES256-GCM-SHA384, ECDHE-ECDSA-AES128-SHA256, ECDHE-ECDSA-AES256-SHA384, ECDHE-RSA-AES128-GCM-SHA256, ECDHE-RSA-AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-ECDSA-AES128-SHA, ECDHE-ECDSA-AES256-SHA, ECDHE-RSA-AES128-SHA, and ECDHE-RSA-AES256-SHA.
	//
	// >  This parameter is returned only for HTTPS listeners.
	//
	// example:
	//
	// tls_cipher_policy_1_0
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of the service that manages the instance.
	//
	// >  This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the instance is managed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The actions that users can perform on the managed instance.
	//
	// >	- This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// >	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*DescribeListenerResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the listener. Valid values:
	//
	// 	- **configuring**: The listener is being configured.
	//
	// 	- **init**: The listener is being initialized.
	//
	// 	- **updating**: The listener is being updated.
	//
	// 	- **deleting:*	- The listener is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The routing type of the listener. Valid values:
	//
	// 	- **Standard**: intelligent routing.
	//
	// 	- **CustomRouting**: custom routing.
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The configurations of the `XForward` headers.
	XForwardedForConfig *DescribeListenerResponseBodyXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s DescribeListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeListenerResponseBody) GetAclType() *string {
	return s.AclType
}

func (s *DescribeListenerResponseBody) GetBackendPorts() []*DescribeListenerResponseBodyBackendPorts {
	return s.BackendPorts
}

func (s *DescribeListenerResponseBody) GetCertificates() []*DescribeListenerResponseBodyCertificates {
	return s.Certificates
}

func (s *DescribeListenerResponseBody) GetClientAffinity() *string {
	return s.ClientAffinity
}

func (s *DescribeListenerResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeListenerResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeListenerResponseBody) GetHttpVersion() *string {
	return s.HttpVersion
}

func (s *DescribeListenerResponseBody) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeListenerResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeListenerResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeListenerResponseBody) GetPortRanges() []*DescribeListenerResponseBodyPortRanges {
	return s.PortRanges
}

func (s *DescribeListenerResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeListenerResponseBody) GetProxyProtocol() *bool {
	return s.ProxyProtocol
}

func (s *DescribeListenerResponseBody) GetRelatedAcls() []*DescribeListenerResponseBodyRelatedAcls {
	return s.RelatedAcls
}

func (s *DescribeListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeListenerResponseBody) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeListenerResponseBody) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *DescribeListenerResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeListenerResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeListenerResponseBody) GetServiceManagedInfos() []*DescribeListenerResponseBodyServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *DescribeListenerResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeListenerResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeListenerResponseBody) GetXForwardedForConfig() *DescribeListenerResponseBodyXForwardedForConfig {
	return s.XForwardedForConfig
}

func (s *DescribeListenerResponseBody) SetAcceleratorId(v string) *DescribeListenerResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetAclType(v string) *DescribeListenerResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeListenerResponseBody) SetBackendPorts(v []*DescribeListenerResponseBodyBackendPorts) *DescribeListenerResponseBody {
	s.BackendPorts = v
	return s
}

func (s *DescribeListenerResponseBody) SetCertificates(v []*DescribeListenerResponseBodyCertificates) *DescribeListenerResponseBody {
	s.Certificates = v
	return s
}

func (s *DescribeListenerResponseBody) SetClientAffinity(v string) *DescribeListenerResponseBody {
	s.ClientAffinity = &v
	return s
}

func (s *DescribeListenerResponseBody) SetCreateTime(v string) *DescribeListenerResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeListenerResponseBody) SetDescription(v string) *DescribeListenerResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeListenerResponseBody) SetHttpVersion(v string) *DescribeListenerResponseBody {
	s.HttpVersion = &v
	return s
}

func (s *DescribeListenerResponseBody) SetIdleTimeout(v int32) *DescribeListenerResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeListenerResponseBody) SetListenerId(v string) *DescribeListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetName(v string) *DescribeListenerResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeListenerResponseBody) SetPortRanges(v []*DescribeListenerResponseBodyPortRanges) *DescribeListenerResponseBody {
	s.PortRanges = v
	return s
}

func (s *DescribeListenerResponseBody) SetProtocol(v string) *DescribeListenerResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeListenerResponseBody) SetProxyProtocol(v bool) *DescribeListenerResponseBody {
	s.ProxyProtocol = &v
	return s
}

func (s *DescribeListenerResponseBody) SetRelatedAcls(v []*DescribeListenerResponseBodyRelatedAcls) *DescribeListenerResponseBody {
	s.RelatedAcls = v
	return s
}

func (s *DescribeListenerResponseBody) SetRequestId(v string) *DescribeListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetRequestTimeout(v int32) *DescribeListenerResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeListenerResponseBody) SetSecurityPolicyId(v string) *DescribeListenerResponseBody {
	s.SecurityPolicyId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetServiceId(v string) *DescribeListenerResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetServiceManaged(v bool) *DescribeListenerResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeListenerResponseBody) SetServiceManagedInfos(v []*DescribeListenerResponseBodyServiceManagedInfos) *DescribeListenerResponseBody {
	s.ServiceManagedInfos = v
	return s
}

func (s *DescribeListenerResponseBody) SetState(v string) *DescribeListenerResponseBody {
	s.State = &v
	return s
}

func (s *DescribeListenerResponseBody) SetType(v string) *DescribeListenerResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeListenerResponseBody) SetXForwardedForConfig(v *DescribeListenerResponseBodyXForwardedForConfig) *DescribeListenerResponseBody {
	s.XForwardedForConfig = v
	return s
}

func (s *DescribeListenerResponseBody) Validate() error {
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
	if s.RelatedAcls != nil {
		for _, item := range s.RelatedAcls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceManagedInfos != nil {
		for _, item := range s.ServiceManagedInfos {
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

type DescribeListenerResponseBodyBackendPorts struct {
	// The first port in the range of ports that are used by the backend server to receive requests.
	//
	// This parameter is returned only if an HTTPS listener is configured and the listener port is the same as the service port of the backend server.
	//
	// example:
	//
	// 80
	FromPort *string `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port in the range of ports that are used by the backend server to receive requests.
	//
	// example:
	//
	// 80
	ToPort *string `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s DescribeListenerResponseBodyBackendPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerResponseBodyBackendPorts) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyBackendPorts) GetFromPort() *string {
	return s.FromPort
}

func (s *DescribeListenerResponseBodyBackendPorts) GetToPort() *string {
	return s.ToPort
}

func (s *DescribeListenerResponseBodyBackendPorts) SetFromPort(v string) *DescribeListenerResponseBodyBackendPorts {
	s.FromPort = &v
	return s
}

func (s *DescribeListenerResponseBodyBackendPorts) SetToPort(v string) *DescribeListenerResponseBodyBackendPorts {
	s.ToPort = &v
	return s
}

func (s *DescribeListenerResponseBodyBackendPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeListenerResponseBodyCertificates struct {
	// The ID of the SSL certificate.
	//
	// example:
	//
	// 449****-cn-hangzhou
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the SSL certificate.
	//
	// Only **Server*	- may be returned, which indicates a server certificate.
	//
	// example:
	//
	// Server
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeListenerResponseBodyCertificates) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyCertificates) GetId() *string {
	return s.Id
}

func (s *DescribeListenerResponseBodyCertificates) GetType() *string {
	return s.Type
}

func (s *DescribeListenerResponseBodyCertificates) SetId(v string) *DescribeListenerResponseBodyCertificates {
	s.Id = &v
	return s
}

func (s *DescribeListenerResponseBodyCertificates) SetType(v string) *DescribeListenerResponseBodyCertificates {
	s.Type = &v
	return s
}

func (s *DescribeListenerResponseBodyCertificates) Validate() error {
	return dara.Validate(s)
}

type DescribeListenerResponseBodyPortRanges struct {
	// The first port in the range of listener ports that are used to receive and forward requests to endpoints.
	//
	// example:
	//
	// 20
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port in the range of listener ports that are used to receive and forward requests to endpoints.
	//
	// example:
	//
	// 20
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s DescribeListenerResponseBodyPortRanges) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerResponseBodyPortRanges) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *DescribeListenerResponseBodyPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *DescribeListenerResponseBodyPortRanges) SetFromPort(v int32) *DescribeListenerResponseBodyPortRanges {
	s.FromPort = &v
	return s
}

func (s *DescribeListenerResponseBodyPortRanges) SetToPort(v int32) *DescribeListenerResponseBodyPortRanges {
	s.ToPort = &v
	return s
}

func (s *DescribeListenerResponseBodyPortRanges) Validate() error {
	return dara.Validate(s)
}

type DescribeListenerResponseBodyRelatedAcls struct {
	// The ID of the ACL that is associated with the listener.
	//
	// example:
	//
	// 123
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Indicates whether the access control feature is enabled. Valid values:
	//
	// 	- **on**: enabled.
	//
	// 	- **off**: disabled.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeListenerResponseBodyRelatedAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerResponseBodyRelatedAcls) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyRelatedAcls) GetAclId() *string {
	return s.AclId
}

func (s *DescribeListenerResponseBodyRelatedAcls) GetStatus() *string {
	return s.Status
}

func (s *DescribeListenerResponseBodyRelatedAcls) SetAclId(v string) *DescribeListenerResponseBodyRelatedAcls {
	s.AclId = &v
	return s
}

func (s *DescribeListenerResponseBodyRelatedAcls) SetStatus(v string) *DescribeListenerResponseBodyRelatedAcls {
	s.Status = &v
	return s
}

func (s *DescribeListenerResponseBodyRelatedAcls) Validate() error {
	return dara.Validate(s)
}

type DescribeListenerResponseBodyServiceManagedInfos struct {
	// The name of the action on the managed instance. Valid values:
	//
	// 	- **Create**
	//
	// 	- **Update**
	//
	// 	- **Delete**
	//
	// 	- **Associate**
	//
	// 	- **UserUnmanaged**
	//
	// 	- **CreateChild**
	//
	// example:
	//
	// Update
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the child resource. Valid values:
	//
	// 	- **Listener**: a listener.
	//
	// 	- **IpSet**: an acceleration region.
	//
	// 	- **EndpointGroup**: an endpoint group.
	//
	// 	- **ForwardingRule**: a forwarding rule.
	//
	// 	- **Endpoint**: an endpoint.
	//
	// 	- **EndpointGroupDestination**: a protocol mapping of an endpoint group associated with a custom routing listener.
	//
	// 	- **EndpointPolicy**: a traffic policy of an endpoint associated with a custom routing listener.
	//
	// >  This parameter is returned only if the value of **Action*	- is **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed.
	//
	// 	- **true**: The specified actions are managed, and users cannot perform the specified actions on the managed instance.
	//
	// 	- **false**: The specified actions are not managed, and users can perform the specified actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s DescribeListenerResponseBodyServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerResponseBodyServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeListenerResponseBodyServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *DescribeListenerResponseBodyServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *DescribeListenerResponseBodyServiceManagedInfos) SetAction(v string) *DescribeListenerResponseBodyServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *DescribeListenerResponseBodyServiceManagedInfos) SetChildType(v string) *DescribeListenerResponseBodyServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *DescribeListenerResponseBodyServiceManagedInfos) SetIsManaged(v bool) *DescribeListenerResponseBodyServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *DescribeListenerResponseBodyServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeListenerResponseBodyXForwardedForConfig struct {
	// Indicates whether the `GA-AP` header is used to retrieve information about acceleration regions. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForGaApEnabled *bool `json:"XForwardedForGaApEnabled,omitempty" xml:"XForwardedForGaApEnabled,omitempty"`
	// Indicates whether the `GA-ID` header is used to retrieve the ID of the GA instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForGaIdEnabled *bool `json:"XForwardedForGaIdEnabled,omitempty" xml:"XForwardedForGaIdEnabled,omitempty"`
	// Indicates whether the `GA-X-Forward-Port` header is used to retrieve the listener ports of the GA instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForPortEnabled *bool `json:"XForwardedForPortEnabled,omitempty" xml:"XForwardedForPortEnabled,omitempty"`
	// Indicates whether the `GA-X-Forward-Proto` header is used to retrieve the listener protocol of the GA instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// Indicates whether the `X-Real-IP` header is used to retrieve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners.
	//
	// example:
	//
	// false
	XRealIpEnabled *bool `json:"XRealIpEnabled,omitempty" xml:"XRealIpEnabled,omitempty"`
}

func (s DescribeListenerResponseBodyXForwardedForConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerResponseBodyXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) GetXForwardedForGaApEnabled() *bool {
	return s.XForwardedForGaApEnabled
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) GetXForwardedForGaIdEnabled() *bool {
	return s.XForwardedForGaIdEnabled
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) GetXForwardedForPortEnabled() *bool {
	return s.XForwardedForPortEnabled
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) GetXForwardedForProtoEnabled() *bool {
	return s.XForwardedForProtoEnabled
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) GetXRealIpEnabled() *bool {
	return s.XRealIpEnabled
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXForwardedForGaApEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XForwardedForGaApEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXForwardedForGaIdEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XForwardedForGaIdEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXForwardedForPortEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XForwardedForPortEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXRealIpEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XRealIpEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) Validate() error {
	return dara.Validate(s)
}
