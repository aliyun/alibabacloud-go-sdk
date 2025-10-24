// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIngressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *CreateIngressRequest
	GetAddressType() *string
	SetCertId(v string) *CreateIngressRequest
	GetCertId() *string
	SetCertIds(v string) *CreateIngressRequest
	GetCertIds() *string
	SetCorsConfig(v string) *CreateIngressRequest
	GetCorsConfig() *string
	SetDefaultRule(v string) *CreateIngressRequest
	GetDefaultRule() *string
	SetDescription(v string) *CreateIngressRequest
	GetDescription() *string
	SetEnableGzip(v bool) *CreateIngressRequest
	GetEnableGzip() *bool
	SetEnableXForwardedFor(v bool) *CreateIngressRequest
	GetEnableXForwardedFor() *bool
	SetEnableXForwardedForClientSrcPort(v bool) *CreateIngressRequest
	GetEnableXForwardedForClientSrcPort() *bool
	SetEnableXForwardedForProto(v bool) *CreateIngressRequest
	GetEnableXForwardedForProto() *bool
	SetEnableXForwardedForSlbId(v bool) *CreateIngressRequest
	GetEnableXForwardedForSlbId() *bool
	SetEnableXForwardedForSlbPort(v bool) *CreateIngressRequest
	GetEnableXForwardedForSlbPort() *bool
	SetIdleTimeout(v int32) *CreateIngressRequest
	GetIdleTimeout() *int32
	SetListenerPort(v int32) *CreateIngressRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *CreateIngressRequest
	GetListenerProtocol() *string
	SetLoadBalanceType(v string) *CreateIngressRequest
	GetLoadBalanceType() *string
	SetLoadBalancerEdition(v string) *CreateIngressRequest
	GetLoadBalancerEdition() *string
	SetNamespaceId(v string) *CreateIngressRequest
	GetNamespaceId() *string
	SetRequestTimeout(v int32) *CreateIngressRequest
	GetRequestTimeout() *int32
	SetRules(v string) *CreateIngressRequest
	GetRules() *string
	SetSecurityPolicyId(v string) *CreateIngressRequest
	GetSecurityPolicyId() *string
	SetSlbId(v string) *CreateIngressRequest
	GetSlbId() *string
	SetZoneMappings(v string) *CreateIngressRequest
	GetZoneMappings() *string
}

type CreateIngressRequest struct {
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the certificate that is associated with the **CLB*	- instance.
	//
	// 	- If you set **LoadBalanceType*	- to **clb**, you can use CertId to configure a certificate for the HTTPS listener.
	//
	// For more information about how to manage the SSL certificate IDs that are used by CLB instances, see [Overview](https://help.aliyun.com/document_detail/90792.html).
	//
	// example:
	//
	// 188077086902****_176993d****_181437****_108724****
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The IDs of the certificates that are associated with the **ALB*	- instance.
	//
	// 	- If you set **LoadBalanceType*	- to **alb**, you can use CertIds to configure multiple certificates for the HTTPS listener. Separate multiple certificate IDs with commas (,).
	//
	// 	- The ID of the SSL certificate that is used by an ALB instance can be obtained from Certificate Management Service. For example, if you specify `756***-cn-hangzhou`, `756***` is the certificate ID that is obtained from the service page, and `-cn-hangzhou` is the fixed suffix. For more information, see [Manage certificates](https://help.aliyun.com/document_detail/209076.html).
	//
	// example:
	//
	// 87***35-cn-hangzhou,812***3-cn-hangzhou
	CertIds    *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	CorsConfig *string `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty"`
	// Default forwarding rule. Traffic is forwarded to the specified application through a designated port based on the IP address. Parameter descriptions are as follows:
	//
	// - **appId**: Application ID. - **containerPort**: Application instance port.
	//
	// > All requests that do not match or do not meet the **Rules*	- for forwarding will be directed to this specified application.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080}
	DefaultRule *string `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	// Route rule name.
	//
	// example:
	//
	// ingress-for-sae-test
	Description                      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableGzip                       *bool   `json:"EnableGzip,omitempty" xml:"EnableGzip,omitempty"`
	EnableXForwardedFor              *bool   `json:"EnableXForwardedFor,omitempty" xml:"EnableXForwardedFor,omitempty"`
	EnableXForwardedForClientSrcPort *bool   `json:"EnableXForwardedForClientSrcPort,omitempty" xml:"EnableXForwardedForClientSrcPort,omitempty"`
	EnableXForwardedForProto         *bool   `json:"EnableXForwardedForProto,omitempty" xml:"EnableXForwardedForProto,omitempty"`
	EnableXForwardedForSlbId         *bool   `json:"EnableXForwardedForSlbId,omitempty" xml:"EnableXForwardedForSlbId,omitempty"`
	EnableXForwardedForSlbPort       *bool   `json:"EnableXForwardedForSlbPort,omitempty" xml:"EnableXForwardedForSlbPort,omitempty"`
	// The timeout period of an idle connection. Unit: seconds Valid values: 1 to 60. If no requests are received within the specified timeout period, ALB closes the current connection. When a new request is received, ALB establishes a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// SThe frontend port that is used by the ALB instance.
	//
	// Valid values: 1 to 65535.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// Request forwarding protocol. The value description is as follows:
	//
	// - **HTTP**: Suitable for applications that need to identify data content. - **HTTPS**: Suitable for applications that require encrypted transmission.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The type of the SLB instance. The instance type can be specified only when you create a routing rule. You cannot change the instance type when you update the routing rule. Valid values:
	//
	// 	- **clb**
	//
	// 	- **alb**
	//
	// example:
	//
	// clb
	LoadBalanceType     *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The ID of the namespace where the application is located. Currently, cross-namespace applications are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:sae-test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The timeout period of a request. Unit: seconds. Valid values: 1 to 180. If no response is received from the backend server within the specified timeout period, ALB stops waiting for the response and returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 3
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The forwarding rules. You can specify a port and an application in a forwarding rule to forward traffic based on the specified domain name and request path. The following list describes the involved parameters:
	//
	// 	- **appId**: the ID of the application.
	//
	// 	- **containerPort**: the container port of the application.
	//
	// 	- **domain**: the domain name.
	//
	// 	- **path**: the request path.
	//
	// 	- **backendProtocol**: the backend service protocol. Valid values: http, https, and grpc. Default value: http.
	//
	// 	- **rewritePath**: the rewrite path.
	//
	// >  The path rewrite feature is supported only by ALB instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080,"domain":"www.sae.site","path":"/path1"},{"appId":"666403ce-d25b-47cf-87fe-497565d2****","containerPort":8080,"domain":"sae.site","path":"/path2"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of a security policy.
	//
	// example:
	//
	// sp-bp1bpn0kn9****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The Server Load Balancer (SLB) instance that is used by the routing rule.
	//
	// >  The SLB instance can be a Classic Load Balancer (CLB) instance or an Application Load Balancer (ALB) instance.
	//
	// example:
	//
	// lb-uf6hucc7inlqrtcq5****
	SlbId        *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	ZoneMappings *string `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty"`
}

func (s CreateIngressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIngressRequest) GoString() string {
	return s.String()
}

func (s *CreateIngressRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *CreateIngressRequest) GetCertId() *string {
	return s.CertId
}

func (s *CreateIngressRequest) GetCertIds() *string {
	return s.CertIds
}

func (s *CreateIngressRequest) GetCorsConfig() *string {
	return s.CorsConfig
}

func (s *CreateIngressRequest) GetDefaultRule() *string {
	return s.DefaultRule
}

func (s *CreateIngressRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIngressRequest) GetEnableGzip() *bool {
	return s.EnableGzip
}

func (s *CreateIngressRequest) GetEnableXForwardedFor() *bool {
	return s.EnableXForwardedFor
}

func (s *CreateIngressRequest) GetEnableXForwardedForClientSrcPort() *bool {
	return s.EnableXForwardedForClientSrcPort
}

func (s *CreateIngressRequest) GetEnableXForwardedForProto() *bool {
	return s.EnableXForwardedForProto
}

func (s *CreateIngressRequest) GetEnableXForwardedForSlbId() *bool {
	return s.EnableXForwardedForSlbId
}

func (s *CreateIngressRequest) GetEnableXForwardedForSlbPort() *bool {
	return s.EnableXForwardedForSlbPort
}

func (s *CreateIngressRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *CreateIngressRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateIngressRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *CreateIngressRequest) GetLoadBalanceType() *string {
	return s.LoadBalanceType
}

func (s *CreateIngressRequest) GetLoadBalancerEdition() *string {
	return s.LoadBalancerEdition
}

func (s *CreateIngressRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateIngressRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *CreateIngressRequest) GetRules() *string {
	return s.Rules
}

func (s *CreateIngressRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *CreateIngressRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *CreateIngressRequest) GetZoneMappings() *string {
	return s.ZoneMappings
}

func (s *CreateIngressRequest) SetAddressType(v string) *CreateIngressRequest {
	s.AddressType = &v
	return s
}

func (s *CreateIngressRequest) SetCertId(v string) *CreateIngressRequest {
	s.CertId = &v
	return s
}

func (s *CreateIngressRequest) SetCertIds(v string) *CreateIngressRequest {
	s.CertIds = &v
	return s
}

func (s *CreateIngressRequest) SetCorsConfig(v string) *CreateIngressRequest {
	s.CorsConfig = &v
	return s
}

func (s *CreateIngressRequest) SetDefaultRule(v string) *CreateIngressRequest {
	s.DefaultRule = &v
	return s
}

func (s *CreateIngressRequest) SetDescription(v string) *CreateIngressRequest {
	s.Description = &v
	return s
}

func (s *CreateIngressRequest) SetEnableGzip(v bool) *CreateIngressRequest {
	s.EnableGzip = &v
	return s
}

func (s *CreateIngressRequest) SetEnableXForwardedFor(v bool) *CreateIngressRequest {
	s.EnableXForwardedFor = &v
	return s
}

func (s *CreateIngressRequest) SetEnableXForwardedForClientSrcPort(v bool) *CreateIngressRequest {
	s.EnableXForwardedForClientSrcPort = &v
	return s
}

func (s *CreateIngressRequest) SetEnableXForwardedForProto(v bool) *CreateIngressRequest {
	s.EnableXForwardedForProto = &v
	return s
}

func (s *CreateIngressRequest) SetEnableXForwardedForSlbId(v bool) *CreateIngressRequest {
	s.EnableXForwardedForSlbId = &v
	return s
}

func (s *CreateIngressRequest) SetEnableXForwardedForSlbPort(v bool) *CreateIngressRequest {
	s.EnableXForwardedForSlbPort = &v
	return s
}

func (s *CreateIngressRequest) SetIdleTimeout(v int32) *CreateIngressRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateIngressRequest) SetListenerPort(v int32) *CreateIngressRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateIngressRequest) SetListenerProtocol(v string) *CreateIngressRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateIngressRequest) SetLoadBalanceType(v string) *CreateIngressRequest {
	s.LoadBalanceType = &v
	return s
}

func (s *CreateIngressRequest) SetLoadBalancerEdition(v string) *CreateIngressRequest {
	s.LoadBalancerEdition = &v
	return s
}

func (s *CreateIngressRequest) SetNamespaceId(v string) *CreateIngressRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateIngressRequest) SetRequestTimeout(v int32) *CreateIngressRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateIngressRequest) SetRules(v string) *CreateIngressRequest {
	s.Rules = &v
	return s
}

func (s *CreateIngressRequest) SetSecurityPolicyId(v string) *CreateIngressRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateIngressRequest) SetSlbId(v string) *CreateIngressRequest {
	s.SlbId = &v
	return s
}

func (s *CreateIngressRequest) SetZoneMappings(v string) *CreateIngressRequest {
	s.ZoneMappings = &v
	return s
}

func (s *CreateIngressRequest) Validate() error {
	return dara.Validate(s)
}
