// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIngressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeIngressResponseBody
	GetCode() *string
	SetData(v *DescribeIngressResponseBodyData) *DescribeIngressResponseBody
	GetData() *DescribeIngressResponseBodyData
	SetErrorCode(v string) *DescribeIngressResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeIngressResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeIngressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeIngressResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeIngressResponseBody
	GetTraceId() *string
}

type DescribeIngressResponseBody struct {
	// The HTTP status code returned for the request. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A client error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is returned only if the request fails.
	//
	// - For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query the details of a call.
	//
	// example:
	//
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeIngressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIngressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeIngressResponseBody) GetData() *DescribeIngressResponseBodyData {
	return s.Data
}

func (s *DescribeIngressResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeIngressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeIngressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIngressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeIngressResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeIngressResponseBody) SetCode(v string) *DescribeIngressResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeIngressResponseBody) SetData(v *DescribeIngressResponseBodyData) *DescribeIngressResponseBody {
	s.Data = v
	return s
}

func (s *DescribeIngressResponseBody) SetErrorCode(v string) *DescribeIngressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeIngressResponseBody) SetMessage(v string) *DescribeIngressResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeIngressResponseBody) SetRequestId(v string) *DescribeIngressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIngressResponseBody) SetSuccess(v bool) *DescribeIngressResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeIngressResponseBody) SetTraceId(v string) *DescribeIngressResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeIngressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIngressResponseBodyData struct {
	// The ID of the **Classic Load Balancer (CLB)*	- certificate.
	//
	// example:
	//
	// 13623****809_16cad216b32_845_-419427029
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The comma-separated IDs of the **Application Load Balancer (ALB)*	- certificates.
	//
	// example:
	//
	// 87***35-cn-hangzhou,812***3-cn-hangzhou
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// The configurations for Cross-Origin Resource Sharing (CORS). Valid HTTP methods:
	//
	// - **GET**
	//
	// - **POST**
	//
	// - **PUT**
	//
	// - **DELETE**
	//
	// - **HEAD**
	//
	// - **OPTIONS**
	//
	// - **PATCH**
	CorsConfig *DescribeIngressResponseBodyDataCorsConfig `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	// Indicates whether the Application Load Balancer (ALB) instance was provisioned by SAE.
	//
	// example:
	//
	// true
	CreatedBySae *bool `json:"CreatedBySae,omitempty" xml:"CreatedBySae,omitempty"`
	// The default rule.
	DefaultRule *DescribeIngressResponseBodyDataDefaultRule `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty" type:"Struct"`
	// The description of the Ingress.
	//
	// example:
	//
	// ingress-sae-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable Gzip compression.
	EnableGzip *bool `json:"EnableGzip,omitempty" xml:"EnableGzip,omitempty"`
	// Specifies whether to use the `X-Forwarded-For` header to retrieve client IP addresses.
	//
	// example:
	//
	// true
	EnableXForwardedFor *bool `json:"EnableXForwardedFor,omitempty" xml:"EnableXForwardedFor,omitempty"`
	// Specifies whether to use a header to retrieve the source port of the client.
	//
	// example:
	//
	// true
	EnableXForwardedForClientSrcPort *bool `json:"EnableXForwardedForClientSrcPort,omitempty" xml:"EnableXForwardedForClientSrcPort,omitempty"`
	// Specifies whether to use the `X-Forwarded-Proto` header to retrieve the listener protocol of the SLB instance.
	//
	// example:
	//
	// true
	EnableXForwardedForProto *bool `json:"EnableXForwardedForProto,omitempty" xml:"EnableXForwardedForProto,omitempty"`
	// Specifies whether to use the `SLB-ID` header to retrieve the ID of the SLB instance.
	//
	// example:
	//
	// true
	EnableXForwardedForSlbId *bool `json:"EnableXForwardedForSlbId,omitempty" xml:"EnableXForwardedForSlbId,omitempty"`
	// Specifies whether to use the `X-Forwarded-Port` header to retrieve the listener port of the SLB instance.
	//
	// example:
	//
	// true
	EnableXForwardedForSlbPort *bool `json:"EnableXForwardedForSlbPort,omitempty" xml:"EnableXForwardedForSlbPort,omitempty"`
	// The ID of the Ingress.
	//
	// example:
	//
	// 87
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The connection idle timeout, in seconds.
	//
	// Valid values: 1 to 60.
	//
	// Default value: 15.
	//
	// If no request is received within the timeout period, the load balancer closes the connection. A new connection is established when the next request is received.
	//
	// example:
	//
	// 3
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The listener port of the SLB instance.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The request forwarding protocol. Valid values:
	//
	// - **HTTP**: Suitable for applications that need to identify data content.
	//
	// - **HTTPS**: Suitable for applications that require encrypted transmission.
	//
	// This parameter is optional for the `CreateIngress` and `UpdateIngress` operations. It is not returned if it was not specified when the Ingress was created or updated.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The type of the Server Load Balancer (SLB) instance. Valid values:
	//
	// - **clb**: Classic Load Balancer (CLB), formerly known as SLB.
	//
	// - **alb**: Application Load Balancer (ALB).
	//
	// example:
	//
	// clb
	LoadBalanceType *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	// The name of the Ingress.
	//
	// example:
	//
	// lb-uf6jt0nu4z6ior943****-80-f5****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:sae-test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The request timeout, in seconds.
	//
	// Valid values: 1 to 180.
	//
	// Default value: 60.
	//
	// If a backend server does not respond within the specified timeout period, the load balancer terminates the request and returns an HTTP 504 error to the client.
	//
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The forwarding rules.
	Rules []*DescribeIngressResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The ID of the security policy instance.
	//
	// example:
	//
	// sp-n0kn923****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of the Server Load Balancer (SLB) instance.
	//
	// example:
	//
	// lb-uf62****6d13tq2u5
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The type of the SLB instance. Valid values:
	//
	// - **internet**: An internet-facing instance.
	//
	// - **intranet**: An internal-facing instance.
	//
	// example:
	//
	// internet
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s DescribeIngressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeIngressResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBodyData) GetCertId() *string {
	return s.CertId
}

func (s *DescribeIngressResponseBodyData) GetCertIds() *string {
	return s.CertIds
}

func (s *DescribeIngressResponseBodyData) GetCorsConfig() *DescribeIngressResponseBodyDataCorsConfig {
	return s.CorsConfig
}

func (s *DescribeIngressResponseBodyData) GetCreatedBySae() *bool {
	return s.CreatedBySae
}

func (s *DescribeIngressResponseBodyData) GetDefaultRule() *DescribeIngressResponseBodyDataDefaultRule {
	return s.DefaultRule
}

func (s *DescribeIngressResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeIngressResponseBodyData) GetEnableGzip() *bool {
	return s.EnableGzip
}

func (s *DescribeIngressResponseBodyData) GetEnableXForwardedFor() *bool {
	return s.EnableXForwardedFor
}

func (s *DescribeIngressResponseBodyData) GetEnableXForwardedForClientSrcPort() *bool {
	return s.EnableXForwardedForClientSrcPort
}

func (s *DescribeIngressResponseBodyData) GetEnableXForwardedForProto() *bool {
	return s.EnableXForwardedForProto
}

func (s *DescribeIngressResponseBodyData) GetEnableXForwardedForSlbId() *bool {
	return s.EnableXForwardedForSlbId
}

func (s *DescribeIngressResponseBodyData) GetEnableXForwardedForSlbPort() *bool {
	return s.EnableXForwardedForSlbPort
}

func (s *DescribeIngressResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DescribeIngressResponseBodyData) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeIngressResponseBodyData) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeIngressResponseBodyData) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeIngressResponseBodyData) GetLoadBalanceType() *string {
	return s.LoadBalanceType
}

func (s *DescribeIngressResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeIngressResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeIngressResponseBodyData) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *DescribeIngressResponseBodyData) GetRules() []*DescribeIngressResponseBodyDataRules {
	return s.Rules
}

func (s *DescribeIngressResponseBodyData) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *DescribeIngressResponseBodyData) GetSlbId() *string {
	return s.SlbId
}

func (s *DescribeIngressResponseBodyData) GetSlbType() *string {
	return s.SlbType
}

func (s *DescribeIngressResponseBodyData) SetCertId(v string) *DescribeIngressResponseBodyData {
	s.CertId = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetCertIds(v string) *DescribeIngressResponseBodyData {
	s.CertIds = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetCorsConfig(v *DescribeIngressResponseBodyDataCorsConfig) *DescribeIngressResponseBodyData {
	s.CorsConfig = v
	return s
}

func (s *DescribeIngressResponseBodyData) SetCreatedBySae(v bool) *DescribeIngressResponseBodyData {
	s.CreatedBySae = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetDefaultRule(v *DescribeIngressResponseBodyDataDefaultRule) *DescribeIngressResponseBodyData {
	s.DefaultRule = v
	return s
}

func (s *DescribeIngressResponseBodyData) SetDescription(v string) *DescribeIngressResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetEnableGzip(v bool) *DescribeIngressResponseBodyData {
	s.EnableGzip = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetEnableXForwardedFor(v bool) *DescribeIngressResponseBodyData {
	s.EnableXForwardedFor = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetEnableXForwardedForClientSrcPort(v bool) *DescribeIngressResponseBodyData {
	s.EnableXForwardedForClientSrcPort = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetEnableXForwardedForProto(v bool) *DescribeIngressResponseBodyData {
	s.EnableXForwardedForProto = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetEnableXForwardedForSlbId(v bool) *DescribeIngressResponseBodyData {
	s.EnableXForwardedForSlbId = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetEnableXForwardedForSlbPort(v bool) *DescribeIngressResponseBodyData {
	s.EnableXForwardedForSlbPort = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetId(v int64) *DescribeIngressResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetIdleTimeout(v int32) *DescribeIngressResponseBodyData {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetListenerPort(v int32) *DescribeIngressResponseBodyData {
	s.ListenerPort = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetListenerProtocol(v string) *DescribeIngressResponseBodyData {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetLoadBalanceType(v string) *DescribeIngressResponseBodyData {
	s.LoadBalanceType = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetName(v string) *DescribeIngressResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetNamespaceId(v string) *DescribeIngressResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetRequestTimeout(v int32) *DescribeIngressResponseBodyData {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetRules(v []*DescribeIngressResponseBodyDataRules) *DescribeIngressResponseBodyData {
	s.Rules = v
	return s
}

func (s *DescribeIngressResponseBodyData) SetSecurityPolicyId(v string) *DescribeIngressResponseBodyData {
	s.SecurityPolicyId = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetSlbId(v string) *DescribeIngressResponseBodyData {
	s.SlbId = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetSlbType(v string) *DescribeIngressResponseBodyData {
	s.SlbType = &v
	return s
}

func (s *DescribeIngressResponseBodyData) Validate() error {
	if s.CorsConfig != nil {
		if err := s.CorsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultRule != nil {
		if err := s.DefaultRule.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIngressResponseBodyDataCorsConfig struct {
	// Specifies whether to allow credentials in cross-origin requests. Valid values:
	//
	// - **on**: yes
	//
	// - **off**: no
	//
	// example:
	//
	// on
	AllowCredentials *string `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// The headers that are allowed in cross-origin requests.
	//
	// example:
	//
	// test_123
	AllowHeaders *string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty"`
	// The HTTP methods that are allowed for cross-origin requests.
	//
	// example:
	//
	// GET
	AllowMethods *string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty"`
	// The origins that are allowed to access the resource. You can specify a single asterisk (`*`) or one or more specific origins.
	//
	// - A specific origin must start with `http://` or `https://` and be a valid domain name or a first-level wildcard domain name. Example: `http://*.test.abc.example.com`.
	//
	// - You can optionally specify a port. The valid port range is **1*	- to **65535**.
	//
	// example:
	//
	// *
	AllowOrigin *string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty"`
	// Specifies whether to enable CORS.
	//
	// example:
	//
	// false
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The headers that are exposed to clients.
	//
	// example:
	//
	// test_123
	ExposeHeaders *string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty"`
	// The maximum cache duration of preflight requests in the browser, in seconds.
	//
	// Valid values: **-1*	- to **172800**.
	//
	// example:
	//
	// 1000
	MaxAge *string `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s DescribeIngressResponseBodyDataCorsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeIngressResponseBodyDataCorsConfig) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBodyDataCorsConfig) GetAllowCredentials() *string {
	return s.AllowCredentials
}

func (s *DescribeIngressResponseBodyDataCorsConfig) GetAllowHeaders() *string {
	return s.AllowHeaders
}

func (s *DescribeIngressResponseBodyDataCorsConfig) GetAllowMethods() *string {
	return s.AllowMethods
}

func (s *DescribeIngressResponseBodyDataCorsConfig) GetAllowOrigin() *string {
	return s.AllowOrigin
}

func (s *DescribeIngressResponseBodyDataCorsConfig) GetEnable() *string {
	return s.Enable
}

func (s *DescribeIngressResponseBodyDataCorsConfig) GetExposeHeaders() *string {
	return s.ExposeHeaders
}

func (s *DescribeIngressResponseBodyDataCorsConfig) GetMaxAge() *string {
	return s.MaxAge
}

func (s *DescribeIngressResponseBodyDataCorsConfig) SetAllowCredentials(v string) *DescribeIngressResponseBodyDataCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *DescribeIngressResponseBodyDataCorsConfig) SetAllowHeaders(v string) *DescribeIngressResponseBodyDataCorsConfig {
	s.AllowHeaders = &v
	return s
}

func (s *DescribeIngressResponseBodyDataCorsConfig) SetAllowMethods(v string) *DescribeIngressResponseBodyDataCorsConfig {
	s.AllowMethods = &v
	return s
}

func (s *DescribeIngressResponseBodyDataCorsConfig) SetAllowOrigin(v string) *DescribeIngressResponseBodyDataCorsConfig {
	s.AllowOrigin = &v
	return s
}

func (s *DescribeIngressResponseBodyDataCorsConfig) SetEnable(v string) *DescribeIngressResponseBodyDataCorsConfig {
	s.Enable = &v
	return s
}

func (s *DescribeIngressResponseBodyDataCorsConfig) SetExposeHeaders(v string) *DescribeIngressResponseBodyDataCorsConfig {
	s.ExposeHeaders = &v
	return s
}

func (s *DescribeIngressResponseBodyDataCorsConfig) SetMaxAge(v string) *DescribeIngressResponseBodyDataCorsConfig {
	s.MaxAge = &v
	return s
}

func (s *DescribeIngressResponseBodyDataCorsConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeIngressResponseBodyDataDefaultRule struct {
	// The ID of the application for the default rule.
	//
	// example:
	//
	// 395b60e4-0550-458d-9c54-a265d036****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application for the default rule.
	//
	// example:
	//
	// app1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The backend protocol. Valid values:
	//
	// - **http**: Suitable for applications that need to identify data content.
	//
	// - **https**: Suitable for applications that require encrypted transmission.
	//
	// - **grpc**: Suitable for load balancing gRPC services developed in multiple languages, such as .NET.
	//
	// This parameter is valid only when the `LoadBalanceType` parameter is set to `alb` and the `ListenerProtocol` parameter is set to `HTTPS`.
	//
	// example:
	//
	// http
	BackendProtocol *string `json:"BackendProtocol,omitempty" xml:"BackendProtocol,omitempty"`
	// The backend port for the default rule.
	//
	// example:
	//
	// 8080
	ContainerPort *int32 `json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
}

func (s DescribeIngressResponseBodyDataDefaultRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeIngressResponseBodyDataDefaultRule) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBodyDataDefaultRule) GetAppId() *string {
	return s.AppId
}

func (s *DescribeIngressResponseBodyDataDefaultRule) GetAppName() *string {
	return s.AppName
}

func (s *DescribeIngressResponseBodyDataDefaultRule) GetBackendProtocol() *string {
	return s.BackendProtocol
}

func (s *DescribeIngressResponseBodyDataDefaultRule) GetContainerPort() *int32 {
	return s.ContainerPort
}

func (s *DescribeIngressResponseBodyDataDefaultRule) SetAppId(v string) *DescribeIngressResponseBodyDataDefaultRule {
	s.AppId = &v
	return s
}

func (s *DescribeIngressResponseBodyDataDefaultRule) SetAppName(v string) *DescribeIngressResponseBodyDataDefaultRule {
	s.AppName = &v
	return s
}

func (s *DescribeIngressResponseBodyDataDefaultRule) SetBackendProtocol(v string) *DescribeIngressResponseBodyDataDefaultRule {
	s.BackendProtocol = &v
	return s
}

func (s *DescribeIngressResponseBodyDataDefaultRule) SetContainerPort(v int32) *DescribeIngressResponseBodyDataDefaultRule {
	s.ContainerPort = &v
	return s
}

func (s *DescribeIngressResponseBodyDataDefaultRule) Validate() error {
	return dara.Validate(s)
}

type DescribeIngressResponseBodyDataRules struct {
	// The ID of the destination application.
	//
	// example:
	//
	// 395b60e4-0550-458d-9c54-a265d036****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the destination application.
	//
	// example:
	//
	// app1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The backend protocol. Valid values:
	//
	// - **http**: Suitable for applications that need to identify data content.
	//
	// - **https**: Suitable for applications that require encrypted transmission.
	//
	// - **grpc**: Suitable for load balancing gRPC services developed in multiple languages, such as .NET.
	//
	// This parameter is valid only when the `LoadBalanceType` parameter is set to `alb` and the `ListenerProtocol` parameter is set to `HTTPS`.
	//
	// example:
	//
	// http
	BackendProtocol *string `json:"BackendProtocol,omitempty" xml:"BackendProtocol,omitempty"`
	// The backend port of the application.
	//
	// example:
	//
	// 8080
	ContainerPort *int32 `json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
	// The domain name of the application.
	//
	// example:
	//
	// edas.site
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The URL path.
	//
	// example:
	//
	// /path1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The rewritten path.
	//
	// example:
	//
	// /${1}
	RewritePath *string `json:"RewritePath,omitempty" xml:"RewritePath,omitempty"`
	// The actions of the forwarding rule.
	RuleActions []*DescribeIngressResponseBodyDataRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
}

func (s DescribeIngressResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeIngressResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBodyDataRules) GetAppId() *string {
	return s.AppId
}

func (s *DescribeIngressResponseBodyDataRules) GetAppName() *string {
	return s.AppName
}

func (s *DescribeIngressResponseBodyDataRules) GetBackendProtocol() *string {
	return s.BackendProtocol
}

func (s *DescribeIngressResponseBodyDataRules) GetContainerPort() *int32 {
	return s.ContainerPort
}

func (s *DescribeIngressResponseBodyDataRules) GetDomain() *string {
	return s.Domain
}

func (s *DescribeIngressResponseBodyDataRules) GetPath() *string {
	return s.Path
}

func (s *DescribeIngressResponseBodyDataRules) GetRewritePath() *string {
	return s.RewritePath
}

func (s *DescribeIngressResponseBodyDataRules) GetRuleActions() []*DescribeIngressResponseBodyDataRulesRuleActions {
	return s.RuleActions
}

func (s *DescribeIngressResponseBodyDataRules) SetAppId(v string) *DescribeIngressResponseBodyDataRules {
	s.AppId = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetAppName(v string) *DescribeIngressResponseBodyDataRules {
	s.AppName = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetBackendProtocol(v string) *DescribeIngressResponseBodyDataRules {
	s.BackendProtocol = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetContainerPort(v int32) *DescribeIngressResponseBodyDataRules {
	s.ContainerPort = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetDomain(v string) *DescribeIngressResponseBodyDataRules {
	s.Domain = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetPath(v string) *DescribeIngressResponseBodyDataRules {
	s.Path = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetRewritePath(v string) *DescribeIngressResponseBodyDataRules {
	s.RewritePath = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetRuleActions(v []*DescribeIngressResponseBodyDataRulesRuleActions) *DescribeIngressResponseBodyDataRules {
	s.RuleActions = v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) Validate() error {
	if s.RuleActions != nil {
		for _, item := range s.RuleActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIngressResponseBodyDataRulesRuleActions struct {
	// The configuration of the action.
	//
	// example:
	//
	// {\\"host\\":\\"www.example.com\\",\\"path\\":\\"/example/text\\",\\"query\\":\\"x=1\\"}
	ActionConfig *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	// The type of the action. Valid values:
	//
	// - rewrite: a rewrite policy
	//
	// - redirect: a redirection policy
	//
	// example:
	//
	// rewrite
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
}

func (s DescribeIngressResponseBodyDataRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s DescribeIngressResponseBodyDataRulesRuleActions) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBodyDataRulesRuleActions) GetActionConfig() *string {
	return s.ActionConfig
}

func (s *DescribeIngressResponseBodyDataRulesRuleActions) GetActionType() *string {
	return s.ActionType
}

func (s *DescribeIngressResponseBodyDataRulesRuleActions) SetActionConfig(v string) *DescribeIngressResponseBodyDataRulesRuleActions {
	s.ActionConfig = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRulesRuleActions) SetActionType(v string) *DescribeIngressResponseBodyDataRulesRuleActions {
	s.ActionType = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRulesRuleActions) Validate() error {
	return dara.Validate(s)
}
