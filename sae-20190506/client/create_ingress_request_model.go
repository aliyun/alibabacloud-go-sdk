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
	// -
	//
	// example:
	//
	// -
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the **CLB*	- certificate. Valid values:
	//
	// 	- If you set **LoadBalanceType*	- to **clb**, you can use CertId to configure a certificate for the HTTPS listener.
	//
	// For more information about how to use SSL certificate IDs for CLB, see [Manage certificates (CLB)](https://help.aliyun.com/document_detail/90792.html).
	//
	// example:
	//
	// 188077086902****_176993d****_181437****_108724****
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The ID of the multi-certificate **ALB**. Valid values:
	//
	// 	- If the **LoadBalanceType*	- is **alb**, use this field to configure multiple certificates for HTTPS listeners. Separate multiple certificate IDs with commas (,).
	//
	// 	- The ID of the SSL certificate used by ALB must be obtained from the digital certificate product. For example, in the configuration `756***-cn-hangzhou`, the `756***` is the certificate ID obtained from the product page, and the `-cn-hangzhou` is a fixed suffix. For more information, see [Manage certificates](https://help.aliyun.com/document_detail/209076.html).
	//
	// example:
	//
	// 87***35-cn-hangzhou,812***3-cn-hangzhou
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// -
	//
	// example:
	//
	// -
	CorsConfig *string `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty"`
	// The default forwarding rule. Forwards traffic to a specified application through a specified port based on the IP address. The following table describes the parameters.
	//
	// 	- **appId**: the ID of the application.
	//
	// 	- **containerPort**: The port of the application instance.
	//
	// >  All requests that do not match or satisfy **Rules*	- forwarding rules are forwarded to the specified application.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080}
	DefaultRule *string `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	// The name of the routing rule.
	//
	// example:
	//
	// ingress-for-sae-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// -
	//
	// example:
	//
	// -
	EnableGzip *bool `json:"EnableGzip,omitempty" xml:"EnableGzip,omitempty"`
	// -
	//
	// example:
	//
	// -
	EnableXForwardedFor *bool `json:"EnableXForwardedFor,omitempty" xml:"EnableXForwardedFor,omitempty"`
	// -
	//
	// example:
	//
	// -
	EnableXForwardedForClientSrcPort *bool `json:"EnableXForwardedForClientSrcPort,omitempty" xml:"EnableXForwardedForClientSrcPort,omitempty"`
	// -
	//
	// example:
	//
	// -
	EnableXForwardedForProto *bool `json:"EnableXForwardedForProto,omitempty" xml:"EnableXForwardedForProto,omitempty"`
	// -
	//
	// example:
	//
	// -
	EnableXForwardedForSlbId *bool `json:"EnableXForwardedForSlbId,omitempty" xml:"EnableXForwardedForSlbId,omitempty"`
	// -
	//
	// example:
	//
	// -
	EnableXForwardedForSlbPort *bool `json:"EnableXForwardedForSlbPort,omitempty" xml:"EnableXForwardedForSlbPort,omitempty"`
	// Specifies the connection idle timeout period. Unit: seconds. Valid values: 1 to 60. If there is no access request within the timeout period, the SLB will temporarily interrupt the current connection until the next request comes to re-establish a new connection.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The SLB listening port. This port cannot be occupied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The request forwarding protocol. Valid values:
	//
	// 	- **HTTP**: suitable for applications that need to identify data content.
	//
	// 	- **HTTPS**: suitable for applications that require encrypted transmission.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// SLB the type of the SLB instance. It depends on the type that you entered when you created the routing rule and cannot be changed when you update it. Valid values:
	//
	// 	- **clb**: traditional SLB CLB (formerly SLB).
	//
	// 	- **alb**: Applied SLB ALB.
	//
	// example:
	//
	// clb
	LoadBalanceType *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	// -
	//
	// example:
	//
	// -
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The ID of the namespace where the application resides. Currently, cross-namespace applications are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:sae-test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// Specifies the request timeout period. Unit: seconds. Valid values: 1 to 180. If the backend server does not respond within the timeout period, the SLB abandons the wait and returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 3
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The forwarding rule. Forwards traffic to a specified application through a specified port based on the domain name and request path. The following table describes the parameters.
	//
	// 	- **appId**: the ID of the application.
	//
	// 	- **containerPort**: The port of the application instance.
	//
	// 	- **domain**: the domain name.
	//
	// 	- **path**: the request path.
	//
	// 	- **backendProtocol**: The backend service protocol. Valid values: http, https, and grpc. Default value: http.
	//
	// 	- **rewritePath**: Rewrite the path.
	//
	// >  Only ALB allows you to set the RewritePath feature. CLB does not support this feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080,"domain":"www.sae.site","path":"/path1"},{"appId":"666403ce-d25b-47cf-87fe-497565d2****","containerPort":8080,"domain":"sae.site","path":"/path2"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the security policy instance.
	//
	// example:
	//
	// sp-bp1bpn0kn9****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The Server Load Balancer (SLB) instance that is used by the routing rule.
	//
	// >  SLB SLB instances include CLB instances and ALB instances.
	//
	// example:
	//
	// lb-uf6hucc7inlqrtcq5****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// -
	//
	// example:
	//
	// -
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
