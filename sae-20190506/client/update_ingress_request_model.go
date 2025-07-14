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
	// The ID of the certificate that is associated with the Classic Load Balancer (**CLB**) instance.
	//
	// 	- If you set **LoadBalanceType*	- to **clb**, you can use CertId to configure a certificate for the HTTPS listener.
	//
	// For more information about how to manage the SSL certificate IDs that are used by CLB instances, see [Overview](https://help.aliyun.com/document_detail/90792.html).
	//
	// example:
	//
	// 188077086902****_176993d****_181437****_108724****
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The IDs of the certificates that are associated with the Application Load Balancer (**ALB**) instance.
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
	// The default forwarding rule. You can specify a port and an application in the default forwarding rule to forward traffic based on the IP address. The following list describes the involved parameters:
	//
	// 	- **appId**: the ID of the application.
	//
	// 	- **containerPort**: the container port of the application.
	//
	// >  All requests that do not match the forwarding rules specified for Rules are forwarded over the port to the application.
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
	Description                      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableXForwardedFor              *bool   `json:"EnableXForwardedFor,omitempty" xml:"EnableXForwardedFor,omitempty"`
	EnableXForwardedForClientSrcPort *bool   `json:"EnableXForwardedForClientSrcPort,omitempty" xml:"EnableXForwardedForClientSrcPort,omitempty"`
	EnableXForwardedForProto         *bool   `json:"EnableXForwardedForProto,omitempty" xml:"EnableXForwardedForProto,omitempty"`
	EnableXForwardedForSlbId         *bool   `json:"EnableXForwardedForSlbId,omitempty" xml:"EnableXForwardedForSlbId,omitempty"`
	EnableXForwardedForSlbPort       *bool   `json:"EnableXForwardedForSlbPort,omitempty" xml:"EnableXForwardedForSlbPort,omitempty"`
	// The timeout period of idle connections. Unit: seconds.
	//
	// >  A value of 0 indicates that the default value is used.
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
	// The port specified for the Server Load Balancer (SLB) listener. You must specify a vacant port.
	//
	// example:
	//
	// 443
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The protocol that is used to forward requests. Valid values:
	//
	// 	- **HTTP**: HTTP is suitable for applications that need to identify the transmitted data.
	//
	// 	- **HTTPS**: HTTPS is suitable for applications that require encrypted data transmission.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// This parameter is discontinued.
	//
	// example:
	//
	// clb
	LoadBalanceType *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	// The request timed out. Unit: seconds.
	//
	// example:
	//
	// 60
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
	// example:
	//
	// [{"appId":"395b60e4-0550-458d-9c54-a265d036****","containerPort":8080,"domain":"www.sae.site","path":"/path1"},{"appId":"666403ce-d25b-47cf-87fe-497565d2****","containerPort":8080,"domain":"sae.site","path":"/path2"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of a security policy.
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
