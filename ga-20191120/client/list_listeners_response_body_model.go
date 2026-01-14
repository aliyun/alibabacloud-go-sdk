// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody
	GetListeners() []*ListListenersResponseBodyListeners
	SetPageNumber(v int32) *ListListenersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListListenersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListListenersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListListenersResponseBody
	GetTotalCount() *int32
}

type ListListenersResponseBody struct {
	// The information about the listeners.
	Listeners []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) GetListeners() []*ListListenersResponseBodyListeners {
	return s.Listeners
}

func (s *ListListenersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListListenersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListListenersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListListenersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetPageNumber(v int32) *ListListenersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListListenersResponseBody) SetPageSize(v int32) *ListListenersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListenersResponseBody) Validate() error {
	if s.Listeners != nil {
		for _, item := range s.Listeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListListenersResponseBodyListeners struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The range of ports that are used by backend servers.
	BackendPorts []*ListListenersResponseBodyListenersBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The information about the SSL certificates.
	Certificates []*ListListenersResponseBodyListenersCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// Indicates whether client affinity is enabled for the listener.
	//
	// 	- If **NONE*	- is returned, client affinity is disabled. When client affinity is disabled, requests from the same client may be forwarded to different endpoints.
	//
	// 	- If **SOURCE_IP*	- is returned, client affinity is enabled. When a client accesses stateful applications, requests from the same client are forwarded to the same endpoint regardless of the source port or protocol.
	//
	// example:
	//
	// SOURCE_IP
	ClientAffinity *string `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	// The timestamp that indicates when the listener was created. Unit: milliseconds.
	//
	// example:
	//
	// 1577786252000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	PortRanges []*ListListenersResponseBodyListenersPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The network transmission protocol that is used by the listener. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Deprecated
	//
	// Indicates whether client IP address preservation is enabled. Valid values:
	//
	// 	- **true**: Client IP address preservation is enabled. This feature allows you to view client IP addresses on backend servers.
	//
	// 	- **false**: Client IP address preservation is disabled.
	//
	// example:
	//
	// true
	ProxyProtocol *bool `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	// The timeout period of HTTP or HTTPS requests. Unit: seconds.
	//
	// >  This parameter is returned only for HTTP and HTTPS listeners. If no responses are received from the backend server within the timeout period, GA returns an HTTP 504 error code to the client.
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
	// > 	- This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// > 	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*ListListenersResponseBodyListenersServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the listener. Valid values:
	//
	// 	- **active**
	//
	// 	- **init**
	//
	// 	- **updating**
	//
	// 	- **deleting**
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
	XForwardedForConfig *ListListenersResponseBodyListenersXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListListenersResponseBodyListeners) GetBackendPorts() []*ListListenersResponseBodyListenersBackendPorts {
	return s.BackendPorts
}

func (s *ListListenersResponseBodyListeners) GetCertificates() []*ListListenersResponseBodyListenersCertificates {
	return s.Certificates
}

func (s *ListListenersResponseBodyListeners) GetClientAffinity() *string {
	return s.ClientAffinity
}

func (s *ListListenersResponseBodyListeners) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListListenersResponseBodyListeners) GetDescription() *string {
	return s.Description
}

func (s *ListListenersResponseBodyListeners) GetHttpVersion() *string {
	return s.HttpVersion
}

func (s *ListListenersResponseBodyListeners) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *ListListenersResponseBodyListeners) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListListenersResponseBodyListeners) GetName() *string {
	return s.Name
}

func (s *ListListenersResponseBodyListeners) GetPortRanges() []*ListListenersResponseBodyListenersPortRanges {
	return s.PortRanges
}

func (s *ListListenersResponseBodyListeners) GetProtocol() *string {
	return s.Protocol
}

func (s *ListListenersResponseBodyListeners) GetProxyProtocol() *bool {
	return s.ProxyProtocol
}

func (s *ListListenersResponseBodyListeners) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *ListListenersResponseBodyListeners) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *ListListenersResponseBodyListeners) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListListenersResponseBodyListeners) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListListenersResponseBodyListeners) GetServiceManagedInfos() []*ListListenersResponseBodyListenersServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListListenersResponseBodyListeners) GetState() *string {
	return s.State
}

func (s *ListListenersResponseBodyListeners) GetType() *string {
	return s.Type
}

func (s *ListListenersResponseBodyListeners) GetXForwardedForConfig() *ListListenersResponseBodyListenersXForwardedForConfig {
	return s.XForwardedForConfig
}

func (s *ListListenersResponseBodyListeners) SetAcceleratorId(v string) *ListListenersResponseBodyListeners {
	s.AcceleratorId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetBackendPorts(v []*ListListenersResponseBodyListenersBackendPorts) *ListListenersResponseBodyListeners {
	s.BackendPorts = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCertificates(v []*ListListenersResponseBodyListenersCertificates) *ListListenersResponseBodyListeners {
	s.Certificates = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetClientAffinity(v string) *ListListenersResponseBodyListeners {
	s.ClientAffinity = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCreateTime(v int64) *ListListenersResponseBodyListeners {
	s.CreateTime = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetDescription(v string) *ListListenersResponseBodyListeners {
	s.Description = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetHttpVersion(v string) *ListListenersResponseBodyListeners {
	s.HttpVersion = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetIdleTimeout(v int32) *ListListenersResponseBodyListeners {
	s.IdleTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetName(v string) *ListListenersResponseBodyListeners {
	s.Name = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetPortRanges(v []*ListListenersResponseBodyListenersPortRanges) *ListListenersResponseBodyListeners {
	s.PortRanges = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetProtocol(v string) *ListListenersResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetProxyProtocol(v bool) *ListListenersResponseBodyListeners {
	s.ProxyProtocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetRequestTimeout(v int32) *ListListenersResponseBodyListeners {
	s.RequestTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetSecurityPolicyId(v string) *ListListenersResponseBodyListeners {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetServiceId(v string) *ListListenersResponseBodyListeners {
	s.ServiceId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetServiceManaged(v bool) *ListListenersResponseBodyListeners {
	s.ServiceManaged = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetServiceManagedInfos(v []*ListListenersResponseBodyListenersServiceManagedInfos) *ListListenersResponseBodyListeners {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetState(v string) *ListListenersResponseBodyListeners {
	s.State = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetType(v string) *ListListenersResponseBodyListeners {
	s.Type = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetXForwardedForConfig(v *ListListenersResponseBodyListenersXForwardedForConfig) *ListListenersResponseBodyListeners {
	s.XForwardedForConfig = v
	return s
}

func (s *ListListenersResponseBodyListeners) Validate() error {
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

type ListListenersResponseBodyListenersBackendPorts struct {
	// The first port in the range of ports that are used by backend servers.
	//
	// example:
	//
	// 80
	FromPort *string `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port in the range of ports that are used by backend servers.
	//
	// example:
	//
	// 80
	ToPort *string `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListListenersResponseBodyListenersBackendPorts) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersBackendPorts) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersBackendPorts) GetFromPort() *string {
	return s.FromPort
}

func (s *ListListenersResponseBodyListenersBackendPorts) GetToPort() *string {
	return s.ToPort
}

func (s *ListListenersResponseBodyListenersBackendPorts) SetFromPort(v string) *ListListenersResponseBodyListenersBackendPorts {
	s.FromPort = &v
	return s
}

func (s *ListListenersResponseBodyListenersBackendPorts) SetToPort(v string) *ListListenersResponseBodyListenersBackendPorts {
	s.ToPort = &v
	return s
}

func (s *ListListenersResponseBodyListenersBackendPorts) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersCertificates struct {
	// The ID of the SSL certificate.
	//
	// example:
	//
	// 44983xxxx-cn-hangzhou
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

func (s ListListenersResponseBodyListenersCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersCertificates) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersCertificates) GetId() *string {
	return s.Id
}

func (s *ListListenersResponseBodyListenersCertificates) GetType() *string {
	return s.Type
}

func (s *ListListenersResponseBodyListenersCertificates) SetId(v string) *ListListenersResponseBodyListenersCertificates {
	s.Id = &v
	return s
}

func (s *ListListenersResponseBodyListenersCertificates) SetType(v string) *ListListenersResponseBodyListenersCertificates {
	s.Type = &v
	return s
}

func (s *ListListenersResponseBodyListenersCertificates) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersPortRanges struct {
	// The first port in the listener port range that is used to receive and forward requests to endpoints.
	//
	// example:
	//
	// 20
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port in the listener port range that is used to receive and forward requests to endpoints.
	//
	// example:
	//
	// 20
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListListenersResponseBodyListenersPortRanges) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersPortRanges) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *ListListenersResponseBodyListenersPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *ListListenersResponseBodyListenersPortRanges) SetFromPort(v int32) *ListListenersResponseBodyListenersPortRanges {
	s.FromPort = &v
	return s
}

func (s *ListListenersResponseBodyListenersPortRanges) SetToPort(v int32) *ListListenersResponseBodyListenersPortRanges {
	s.ToPort = &v
	return s
}

func (s *ListListenersResponseBodyListenersPortRanges) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersServiceManagedInfos struct {
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
	// 	- **Listener**: listener.
	//
	// 	- **IpSet**: acceleration region.
	//
	// 	- **EndpointGroup**: endpoint group.
	//
	// 	- **ForwardingRule**: forwarding rule.
	//
	// 	- **Endpoint**: endpoint.
	//
	// 	- **EndpointGroupDestination**: protocol mapping of an endpoint group associated with a custom routing listener.
	//
	// 	- **EndpointPolicy**: traffic policy of an endpoint associated with a custom routing listener.
	//
	// >  This parameter takes effect only if the value of **Action*	- is **CreateChild**.
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

func (s ListListenersResponseBodyListenersServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListListenersResponseBodyListenersServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListListenersResponseBodyListenersServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListListenersResponseBodyListenersServiceManagedInfos) SetAction(v string) *ListListenersResponseBodyListenersServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListListenersResponseBodyListenersServiceManagedInfos) SetChildType(v string) *ListListenersResponseBodyListenersServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListListenersResponseBodyListenersServiceManagedInfos) SetIsManaged(v bool) *ListListenersResponseBodyListenersServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListListenersResponseBodyListenersServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}

type ListListenersResponseBodyListenersXForwardedForConfig struct {
	// Indicates whether the `GA-AP` header is used to retrieve the information about acceleration regions. Valid values:
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

func (s ListListenersResponseBodyListenersXForwardedForConfig) String() string {
	return dara.Prettify(s)
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForGaApEnabled() *bool {
	return s.XForwardedForGaApEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForGaIdEnabled() *bool {
	return s.XForwardedForGaIdEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForPortEnabled() *bool {
	return s.XForwardedForPortEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXForwardedForProtoEnabled() *bool {
	return s.XForwardedForProtoEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) GetXRealIpEnabled() *bool {
	return s.XRealIpEnabled
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForGaApEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForGaApEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForGaIdEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForGaIdEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForPortEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForPortEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXRealIpEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XRealIpEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) Validate() error {
	return dara.Validate(s)
}
