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
	// The address type. Valid values:
	//
	// - Internet: public network address.
	//
	// - Intranet: private network address.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the **CLB*	- certificate. Valid values:
	//
	// - If **LoadBalanceType*	- is set to **clb**, use this field to configure the certificate for the HTTPS listener.
	//
	// For more information about how to use SSL certificate IDs in CLB, see [Manage certificates (CLB)](https://help.aliyun.com/document_detail/90792.html).
	//
	// example:
	//
	// 188077086902****_176993d****_181437****_108724****
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The IDs of the **ALB*	- certificates. Valid values:
	//
	// - If **LoadBalanceType*	- is set to **alb**, use this field to configure multiple certificates for the HTTPS listener. Separate multiple certificate IDs with commas (,).
	//
	// - You must obtain the SSL certificate IDs used by ALB from the Certificate Management Service. For example, if you configure `756***-cn-hangzhou`, `756***` is the certificate ID obtained from the product page, and `-cn-hangzhou` is a fixed suffix. For more information, see [Manage certificates (ALB)](https://help.aliyun.com/document_detail/209076.html).
	//
	// example:
	//
	// 87***35-cn-hangzhou,812***3-cn-hangzhou
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// The cross-origin resource sharing (CORS) configuration.
	//
	// example:
	//
	// {"Enable":"true"}
	CorsConfig *string `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty"`
	// The default forwarding rule. Traffic is forwarded to the specified application based on the IP address through the specified port. The parameters are as follows:
	//
	// - **appId**: the ID of the application.
	//
	// - **containerPort**: the port of the application instance.
	//
	// > All requests that do not match or meet the forwarding rules in **Rules*	- are forwarded to the specified application.
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
	// Specifies whether to enable data compression.
	//
	// example:
	//
	// true
	EnableGzip *bool `json:"EnableGzip,omitempty" xml:"EnableGzip,omitempty"`
	// Specifies whether to enable the feature to obtain the client IP address of the visitor through the X-Forwarded-For header field.
	//
	// example:
	//
	// true
	EnableXForwardedFor *bool `json:"EnableXForwardedFor,omitempty" xml:"EnableXForwardedFor,omitempty"`
	// Specifies whether to obtain the listening port of the load balancing instance through the X-Forwarded-Port header field.
	//
	// example:
	//
	// true
	EnableXForwardedForClientSrcPort *bool `json:"EnableXForwardedForClientSrcPort,omitempty" xml:"EnableXForwardedForClientSrcPort,omitempty"`
	// Specifies whether to obtain the listening protocol of the load balancing instance through the X-Forwarded-Proto header field.
	//
	// example:
	//
	// true
	EnableXForwardedForProto *bool `json:"EnableXForwardedForProto,omitempty" xml:"EnableXForwardedForProto,omitempty"`
	// Specifies whether to obtain the load balancing instance ID through the SLB-ID header field.
	//
	// example:
	//
	// true
	EnableXForwardedForSlbId *bool `json:"EnableXForwardedForSlbId,omitempty" xml:"EnableXForwardedForSlbId,omitempty"`
	// Specifies whether to obtain the listening port of the load balancing instance through the X-Forwarded-Port header field.
	//
	// example:
	//
	// true
	EnableXForwardedForSlbPort *bool `json:"EnableXForwardedForSlbPort,omitempty" xml:"EnableXForwardedForSlbPort,omitempty"`
	// Specifies the connection idle timeout period. Unit: seconds. Valid values: 1 to 60. If no access request is received within the timeout period, load balancing temporarily breaks the current connection until a new connection is established when the next request arrives.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The listening port of the SLB instance. The port must not be occupied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The request forwarding protocol. Valid values:
	//
	// - **HTTP**: applicable to applications that need to identify data content.
	//
	// - **HTTPS**: applicable to applications that require encrypted transmission.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The type of the SLB instance. The type depends on the value specified when the routing rule is created and cannot be changed during updates. Valid values:
	//
	// - **clb**: Classic Load Balancer (CLB).
	//
	// - **alb**: Application Load Balancer (ALB).
	//
	// example:
	//
	// clb
	LoadBalanceType *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	// The edition of the Application Load Balancer (ALB). Different editions have different feature limits and billing policy. Valid values:
	//
	// - **Standard**: Standard Edition.
	//
	// - **StandardWithWaf**: WAF-enhanced Edition.
	//
	// example:
	//
	// Standard
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The ID of the namespace where the application resides. Cross-namespace applications are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:sae-test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// Specifies the request timeout period. Unit: seconds. Valid values: 1 to 180. If the backend server does not send a response within the timeout period, load balancing stops waiting and returns an HTTP 504 error code to the client.
	//
	// example:
	//
	// 3
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The forwarding rules. Traffic is forwarded to the specified application based on the domain name and the request path through the specified port. The parameters are as follows:
	//
	// - **appId**: the ID of the application.
	//
	// - **containerPort**: the port of the application instance.
	//
	// - **domain**: the domain name.
	//
	// - **path**: the URI of the request.
	//
	// - **backendProtocol**: the backend service Terms of Service. Valid values: http, https, and grpc. Default value: http.
	//
	// - **rewritePath**: the rewrite path.
	//
	// > Currently, only ALB supports the URL path rewrite (RewritePath) feature. CLB does not support this feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080,"domain":"www.sae.site","path":"/path1"},{"appId":"666403ce-d25b-47cf-87fe-497565d2****","containerPort":8080,"domain":"sae.site","path":"/sys/(.*)/(.*)/aaa","backendProtocol":"http"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the security policy instance.
	//
	// example:
	//
	// sp-bp1bpn0kn9****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The SLB instance ID used by the routing rule.
	//
	// > Server Load Balancer (SLB) includes CLB and ALB instances.
	//
	// example:
	//
	// lb-uf6hucc7inlqrtcq5****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The JSON string of the zone and vSwitch mapping list. If the current region supports two or more zones, you must add at least two zones. A ZoneMapping structure includes:
	//
	// - VSwitchId: the vSwitch corresponding to the zone. Data type: String. Only one vSwitch and one subnet can be used in each zone.
	//
	// - ZoneId: the zone ID of the Network Load Balancer (NLB) instance. Data type: String.
	//
	// example:
	//
	// [{"VSwitchId":"vsw-wz9klui6icc08p6******","ZoneId":"cn-shenzhen-c"},{"VSwitchId":"vsw-wz9frrmoeuki2wp******","ZoneId":"cn-shenzhen-e"}]
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
