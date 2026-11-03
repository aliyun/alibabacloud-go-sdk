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
	// The ID of the Classic Load Balancer (CLB) certificate. Valid values:
	//
	// - If **LoadBalanceType*	- is set to **clb**, use this field to configure the certificate for the HTTPS listener.
	//
	// For more information about using SSL certificate IDs in CLB, see [Manage certificates (CLB)](https://help.aliyun.com/document_detail/90792.html).
	//
	// example:
	//
	// 188077086902****_176993d****_181437****_108724****
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The IDs of multiple Application Load Balancer (ALB) certificates. Valid values:
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
	// {\\"Enable\\":\\"true\\"}
	CorsConfig *string `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty"`
	// The default forwarding rule. Traffic is forwarded to a specified application through a specified port based on the IP address. The parameters are described as follows:
	//
	// - **appId**: the ID of the application.
	//
	// - **containerPort**: the port of the application instance.
	//
	// > All requests that do not match or meet the forwarding rules in **Rules*	- are forwarded to the specified application.
	//
	// example:
	//
	// {"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080}
	DefaultRule *string `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	// The name of the routing rule.
	//
	// example:
	//
	// ingress-sae-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable data compression.
	//
	// example:
	//
	// true
	EnableGzip *bool `json:"EnableGzip,omitempty" xml:"EnableGzip,omitempty"`
	// Specifies whether to enable the retrieval of the client IP address through the X-Forwarded-For header field.
	//
	// example:
	//
	// true
	EnableXForwardedFor *bool `json:"EnableXForwardedFor,omitempty" xml:"EnableXForwardedFor,omitempty"`
	// Specifies whether to retrieve the listening port of the load balancing instance through the X-Forwarded-Port header field.
	//
	// example:
	//
	// true
	EnableXForwardedForClientSrcPort *bool `json:"EnableXForwardedForClientSrcPort,omitempty" xml:"EnableXForwardedForClientSrcPort,omitempty"`
	// Specifies whether to retrieve the listening protocol of the load balancing instance through the X-Forwarded-Proto header field.
	//
	// example:
	//
	// true
	EnableXForwardedForProto *bool `json:"EnableXForwardedForProto,omitempty" xml:"EnableXForwardedForProto,omitempty"`
	// Specifies whether to retrieve the instance ID of the load balancing instance through the SLB-ID header field.
	//
	// example:
	//
	// true
	EnableXForwardedForSlbId *bool `json:"EnableXForwardedForSlbId,omitempty" xml:"EnableXForwardedForSlbId,omitempty"`
	// Specifies whether to retrieve the listening port of the load balancing instance through the X-Forwarded-Port header field.
	//
	// example:
	//
	// true
	EnableXForwardedForSlbPort *bool `json:"EnableXForwardedForSlbPort,omitempty" xml:"EnableXForwardedForSlbPort,omitempty"`
	// The idle connection timeout period, in seconds.
	//
	// > A value of 0 indicates that the default idle timeout value is used.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The ID of the routing rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 87
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
	// The listening port of the Server Load Balancer (SLB) instance. The port must not be occupied.
	//
	// example:
	//
	// 443
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
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
	// **[Deprecated]*	- This parameter is deprecated and updates are no longer supported.
	//
	// example:
	//
	// clb
	LoadBalanceType *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	// The request timeout period, in seconds.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The forwarding rules. Traffic is forwarded to a specified application through a specified port based on the domain name and request path. The parameters are described as follows:
	//
	// - **appId**: the ID of the application.
	//
	// - **containerPort**: the port of the application instance.
	//
	// - **domain**: the domain name.
	//
	// - **path**: the request path.
	//
	// example:
	//
	// [{"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080,"domain":"www.sae.site","path":"/path1"},{"appId":"666403ce-d25b-47cf-87fe-497565d2****","containerPort":8080,"domain":"sae.site","path":"/path2"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the security policy instance.
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
