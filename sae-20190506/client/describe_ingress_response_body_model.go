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
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result returned.
	Data *DescribeIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error codes. Valid values:
	//
	// 	- **ErrorCode*	- is not returned if a request is successful.
	//
	// 	- **ErrorCode*	- is returned if a request failed. For more information, see **Error code*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. Valid values:
	//
	// 	- **success*	- is returned when a request is successful.
	//
	// 	- An error code is returned when the request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of a request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the configurations of Ingresses were queried successfully. Valid values:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: The information failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of a trace. The ID is used to query the details of a request.
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
	// The ID of the certificate that is associated with a Classic Load Balancer (**CLB**) instance.
	//
	// example:
	//
	// 13623****809_16cad216b32_845_-419427029
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The ID of the certificate that is associated with an Application Load Balancer **ALB*	- instance.
	//
	// example:
	//
	// 87***35-cn-hangzhou,812***3-cn-hangzhou
	CertIds      *string                                    `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	CorsConfig   *DescribeIngressResponseBodyDataCorsConfig `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	CreatedBySae *bool                                      `json:"CreatedBySae,omitempty" xml:"CreatedBySae,omitempty"`
	// The default rule.
	DefaultRule *DescribeIngressResponseBodyDataDefaultRule `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty" type:"Struct"`
	// The name of a routing rule.
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
	// The ID of a routing rule.
	//
	// example:
	//
	// 87
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 3
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The listener ports for an SLB instance.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The protocol used to forward requests. Valid values:
	//
	// 	- **HTTP**: HTTP is suitable for applications that need to identify the transmitted data.
	//
	// 	- **HTTPS**: HTTPS is suitable for applications that require encrypted data transmission.
	//
	// This parameter is optional in the **CreateIngress*	- and **UpadateIngress*	- operations. If you do not configure this parameter when you call the CreateIngress or UpdateIngress operation to create or update a gateway routing rule, this parameter is not returned for the corresponding response.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The type of SLB instances. Valid values:
	//
	// 	- **clb**: Classic Load Balancer (formerly known as SLB).
	//
	// 	- **alb**: Application Load Balancer.
	//
	// example:
	//
	// clb
	LoadBalanceType *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	// The name of a routing rule.
	//
	// example:
	//
	// lb-uf6jt0nu4z6ior943****-80-f5****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of a namespace.
	//
	// example:
	//
	// cn-beijing:sae-test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// 60
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The forwarding rules.
	Rules []*DescribeIngressResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// sp-n0kn923****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of a Server Load Balancer (SLB) instance.
	//
	// example:
	//
	// lb-uf62****6d13tq2u5
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The type of an SLB instance. Valid values:
	//
	// 	- **internet**: an Internet-facing SLB instance
	//
	// 	- **intranet**: an Intranet-facing SLB instance
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
	AllowCredentials *string `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	AllowHeaders     *string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty"`
	AllowMethods     *string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty"`
	AllowOrigin      *string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty"`
	Enable           *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	ExposeHeaders    *string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty"`
	MaxAge           *string `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
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
	// The ID of the application that is specified in the default rule.
	//
	// example:
	//
	// 395b60e4-0550-458d-9c54-a265d036****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application that is specified in the default rule.
	//
	// example:
	//
	// app1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The backend protocol. Valid values:
	//
	// 	- **http**: HTTP is suitable for applications that need to identify the transmitted data.
	//
	// 	- **https**: HTTP is suitable for applications that require encrypted data transmission.
	//
	// 	- **grpc**: GRPC is suitable for load balancing scenarios in which you want to deploy services in multi-language frameworks, such as the .NET framework.
	//
	// This parameter is returned only if the**LoadBalanceType*	- parameter is set to **ALB*	- and the **ListenerProtocol*	- parameter **is set to HTTPS**.
	//
	// example:
	//
	// HTTP
	BackendProtocol *string `json:"BackendProtocol,omitempty" xml:"BackendProtocol,omitempty"`
	// The container port of the application that is specified in the default rule.
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
	// The ID of the application specified in the forwarding rule.
	//
	// example:
	//
	// 395b60e4-0550-458d-9c54-a265d036****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application specified in the forwarding rules.
	//
	// example:
	//
	// app1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The backend protocol. Valid values:
	//
	// 	- **http**: HTTP is suitable for applications that need to identify the transmitted data.
	//
	// 	- **https**: HTTPS is suitable for applications that require encrypted data transmission.
	//
	// 	- **grpc**: GRPC is suitable for load balancing scenarios in which you want to deploy services in multi-language frameworks, such as the .NET framework.
	//
	// This parameter is returned only if the **LoadBalanceType*	- parameter is set to **ALB*	- and the **ListenerProtocol*	- parameter is set to **HTTPS**.
	//
	// example:
	//
	// HTTP
	BackendProtocol *string `json:"BackendProtocol,omitempty" xml:"BackendProtocol,omitempty"`
	// Tthe container port of the application specified in the forwarding rules.
	//
	// example:
	//
	// 8080
	ContainerPort *int32 `json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
	// The domain name of the application specified in the forwarding rules.
	//
	// example:
	//
	// edas.site
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The path of a URL.
	//
	// example:
	//
	// /path1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path that is used to rewrite the original path.
	//
	// example:
	//
	// /${1}
	RewritePath *string                                            `json:"RewritePath,omitempty" xml:"RewritePath,omitempty"`
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
	ActionConfig *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	ActionType   *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
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
