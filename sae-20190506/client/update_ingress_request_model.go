// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIngressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *UpdateIngressRequest
	GetCertId() *string
	SetCertIds(v string) *UpdateIngressRequest
	GetCertIds() *string
	SetCorsConfig(v string) *UpdateIngressRequest
	GetCorsConfig() *string
	SetDefaultRule(v string) *UpdateIngressRequest
	GetDefaultRule() *string
	SetDescription(v string) *UpdateIngressRequest
	GetDescription() *string
	SetEnableGzip(v bool) *UpdateIngressRequest
	GetEnableGzip() *bool
	SetEnableXForwardedFor(v bool) *UpdateIngressRequest
	GetEnableXForwardedFor() *bool
	SetEnableXForwardedForClientSrcPort(v bool) *UpdateIngressRequest
	GetEnableXForwardedForClientSrcPort() *bool
	SetEnableXForwardedForProto(v bool) *UpdateIngressRequest
	GetEnableXForwardedForProto() *bool
	SetEnableXForwardedForSlbId(v bool) *UpdateIngressRequest
	GetEnableXForwardedForSlbId() *bool
	SetEnableXForwardedForSlbPort(v bool) *UpdateIngressRequest
	GetEnableXForwardedForSlbPort() *bool
	SetIdleTimeout(v int32) *UpdateIngressRequest
	GetIdleTimeout() *int32
	SetIngressId(v int64) *UpdateIngressRequest
	GetIngressId() *int64
	SetListenerPort(v string) *UpdateIngressRequest
	GetListenerPort() *string
	SetListenerProtocol(v string) *UpdateIngressRequest
	GetListenerProtocol() *string
	SetLoadBalanceType(v string) *UpdateIngressRequest
	GetLoadBalanceType() *string
	SetRequestTimeout(v int32) *UpdateIngressRequest
	GetRequestTimeout() *int32
	SetRules(v string) *UpdateIngressRequest
	GetRules() *string
	SetSecurityPolicyId(v string) *UpdateIngressRequest
	GetSecurityPolicyId() *string
}

type UpdateIngressRequest struct {
	// **CLB*	- certificate ID. Details are as follows:
	//
	// - If **LoadBalanceType*	- is **clb**, use this field to configure the HTTPS listener certificate.
	//
	// For more information about using SSL certificate IDs with CLB, see [Manage Certificates (CLB)]().
	//
	// example:
	//
	// 188077086902****_176993d****_181437****_108724****
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// **ALB*	- multiple certificate IDs. Details are as follows:
	//
	// - If **LoadBalanceType*	- is **alb**, use this field to configure multiple HTTPS listener certificates. Separate multiple certificate IDs with commas.
	//
	// - Obtain the SSL certificate ID used by ALB from the digital certificate product. For example, configure `756***-cn-hangzhou`, where `756***` is the certificate ID obtained from the product page, and `-cn-hangzhou` is a fixed suffix. For more information, see [Manage Certificates (ALB)]().
	//
	// example:
	//
	// 87***35-cn-hangzhou,812***3-cn-hangzhou
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// Cross-domain configuration.
	//
	// example:
	//
	// {\\"Enable\\":\\"true\\"}
	CorsConfig *string `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty"`
	// Default forwarding rule. Forward traffic to the specified application by IP address through the specified port. Parameter description:
	//
	// - **appId**: Application ID.
	//
	// - **containerPort**: Application instance port.
	//
	// > All requests that do not match or satisfy the **Rules*	- forwarding rule are forwarded to this specified application.
	//
	// example:
	//
	// {"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080}
	DefaultRule *string `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	// Routing rule name.
	//
	// example:
	//
	// ingress-sae-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Enable or disable data compression.
	//
	// example:
	//
	// true
	EnableGzip *bool `json:"EnableGzip,omitempty" xml:"EnableGzip,omitempty"`
	// Enable or disable obtaining the client IP address of the visitor through the X-Forwarded-For header field.
	//
	// example:
	//
	// true
	EnableXForwardedFor *bool `json:"EnableXForwardedFor,omitempty" xml:"EnableXForwardedFor,omitempty"`
	// Obtain the listening port of the SLB instance through the X-Forwarded-Port header field.
	//
	// example:
	//
	// true
	EnableXForwardedForClientSrcPort *bool `json:"EnableXForwardedForClientSrcPort,omitempty" xml:"EnableXForwardedForClientSrcPort,omitempty"`
	// Specifies whether to determine the listener protocol of the SLB instance from the X-Forwarded-Proto header field.
	//
	// example:
	//
	// true
	EnableXForwardedForProto *bool `json:"EnableXForwardedForProto,omitempty" xml:"EnableXForwardedForProto,omitempty"`
	// Obtain the SLB instance ID through the SLB-ID header field.
	//
	// example:
	//
	// true
	EnableXForwardedForSlbId *bool `json:"EnableXForwardedForSlbId,omitempty" xml:"EnableXForwardedForSlbId,omitempty"`
	// Whether to obtain the listening port of the Server Load Balancer instance from the X-Forwarded-Port header field.
	//
	// example:
	//
	// true
	EnableXForwardedForSlbPort *bool `json:"EnableXForwardedForSlbPort,omitempty" xml:"EnableXForwardedForSlbPort,omitempty"`
	// Idle connection timeout, in seconds (s).
	//
	// > A value of 0 indicates that the default idle timeout is used.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Routing rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 87
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
	// SLB listening port. This port must not be occupied.
	//
	// example:
	//
	// 443
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// Forwarding Protocol. Details are as follows:
	//
	// - **HTTP**: Applies to applications that need to identify data content.
	//
	// - **HTTPS**: Applies to applications that need encrypted transmission.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// Deprecated parameter. Updates are no longer supported.
	//
	// example:
	//
	// clb
	LoadBalanceType *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	// Request timeout, in seconds (s).
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// Forwarding rules. Forward traffic to the specified application by domain name and URI of the request through the specified port. Parameter description:
	//
	// - **appId**: Application ID.
	//
	// - **containerPort**: Application instance port.
	//
	// - **domain**: Domain name.
	//
	// - **path**: URI of the request.
	//
	// example:
	//
	// [{"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080,"domain":"www.sae.site","path":"/path1"},{"appId":"666403ce-d25b-47cf-87fe-497565d2****","containerPort":8080,"domain":"sae.site","path":"/path2"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// Security policy instance ID.
	//
	// example:
	//
	// tls_cipher_policy_1_2_strict_with_1_3
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s UpdateIngressRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIngressRequest) GoString() string {
	return s.String()
}

func (s *UpdateIngressRequest) GetCertId() *string {
	return s.CertId
}

func (s *UpdateIngressRequest) GetCertIds() *string {
	return s.CertIds
}

func (s *UpdateIngressRequest) GetCorsConfig() *string {
	return s.CorsConfig
}

func (s *UpdateIngressRequest) GetDefaultRule() *string {
	return s.DefaultRule
}

func (s *UpdateIngressRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateIngressRequest) GetEnableGzip() *bool {
	return s.EnableGzip
}

func (s *UpdateIngressRequest) GetEnableXForwardedFor() *bool {
	return s.EnableXForwardedFor
}

func (s *UpdateIngressRequest) GetEnableXForwardedForClientSrcPort() *bool {
	return s.EnableXForwardedForClientSrcPort
}

func (s *UpdateIngressRequest) GetEnableXForwardedForProto() *bool {
	return s.EnableXForwardedForProto
}

func (s *UpdateIngressRequest) GetEnableXForwardedForSlbId() *bool {
	return s.EnableXForwardedForSlbId
}

func (s *UpdateIngressRequest) GetEnableXForwardedForSlbPort() *bool {
	return s.EnableXForwardedForSlbPort
}

func (s *UpdateIngressRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *UpdateIngressRequest) GetIngressId() *int64 {
	return s.IngressId
}

func (s *UpdateIngressRequest) GetListenerPort() *string {
	return s.ListenerPort
}

func (s *UpdateIngressRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *UpdateIngressRequest) GetLoadBalanceType() *string {
	return s.LoadBalanceType
}

func (s *UpdateIngressRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *UpdateIngressRequest) GetRules() *string {
	return s.Rules
}

func (s *UpdateIngressRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *UpdateIngressRequest) SetCertId(v string) *UpdateIngressRequest {
	s.CertId = &v
	return s
}

func (s *UpdateIngressRequest) SetCertIds(v string) *UpdateIngressRequest {
	s.CertIds = &v
	return s
}

func (s *UpdateIngressRequest) SetCorsConfig(v string) *UpdateIngressRequest {
	s.CorsConfig = &v
	return s
}

func (s *UpdateIngressRequest) SetDefaultRule(v string) *UpdateIngressRequest {
	s.DefaultRule = &v
	return s
}

func (s *UpdateIngressRequest) SetDescription(v string) *UpdateIngressRequest {
	s.Description = &v
	return s
}

func (s *UpdateIngressRequest) SetEnableGzip(v bool) *UpdateIngressRequest {
	s.EnableGzip = &v
	return s
}

func (s *UpdateIngressRequest) SetEnableXForwardedFor(v bool) *UpdateIngressRequest {
	s.EnableXForwardedFor = &v
	return s
}

func (s *UpdateIngressRequest) SetEnableXForwardedForClientSrcPort(v bool) *UpdateIngressRequest {
	s.EnableXForwardedForClientSrcPort = &v
	return s
}

func (s *UpdateIngressRequest) SetEnableXForwardedForProto(v bool) *UpdateIngressRequest {
	s.EnableXForwardedForProto = &v
	return s
}

func (s *UpdateIngressRequest) SetEnableXForwardedForSlbId(v bool) *UpdateIngressRequest {
	s.EnableXForwardedForSlbId = &v
	return s
}

func (s *UpdateIngressRequest) SetEnableXForwardedForSlbPort(v bool) *UpdateIngressRequest {
	s.EnableXForwardedForSlbPort = &v
	return s
}

func (s *UpdateIngressRequest) SetIdleTimeout(v int32) *UpdateIngressRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateIngressRequest) SetIngressId(v int64) *UpdateIngressRequest {
	s.IngressId = &v
	return s
}

func (s *UpdateIngressRequest) SetListenerPort(v string) *UpdateIngressRequest {
	s.ListenerPort = &v
	return s
}

func (s *UpdateIngressRequest) SetListenerProtocol(v string) *UpdateIngressRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *UpdateIngressRequest) SetLoadBalanceType(v string) *UpdateIngressRequest {
	s.LoadBalanceType = &v
	return s
}

func (s *UpdateIngressRequest) SetRequestTimeout(v int32) *UpdateIngressRequest {
	s.RequestTimeout = &v
	return s
}

func (s *UpdateIngressRequest) SetRules(v string) *UpdateIngressRequest {
	s.Rules = &v
	return s
}

func (s *UpdateIngressRequest) SetSecurityPolicyId(v string) *UpdateIngressRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateIngressRequest) Validate() error {
	return dara.Validate(s)
}
