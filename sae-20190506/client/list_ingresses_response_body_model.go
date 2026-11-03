// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIngressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIngressesResponseBody
	GetCode() *string
	SetData(v *ListIngressesResponseBodyData) *ListIngressesResponseBody
	GetData() *ListIngressesResponseBodyData
	SetErrorCode(v string) *ListIngressesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListIngressesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIngressesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIngressesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListIngressesResponseBody
	GetTraceId() *string
}

type ListIngressesResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirect.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned results.
	Data *ListIngressesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. The value is described as follows:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, the specific error code is returned.
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
	// Specifies whether the ingress list was successfully retrieved. Valid values:
	//
	// - **true**: The ingress list was successfully retrieved.
	//
	// - **false**: The ingress list failed to be retrieved.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListIngressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIngressesResponseBody) GetData() *ListIngressesResponseBodyData {
	return s.Data
}

func (s *ListIngressesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListIngressesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIngressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIngressesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIngressesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListIngressesResponseBody) SetCode(v string) *ListIngressesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIngressesResponseBody) SetData(v *ListIngressesResponseBodyData) *ListIngressesResponseBody {
	s.Data = v
	return s
}

func (s *ListIngressesResponseBody) SetErrorCode(v string) *ListIngressesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListIngressesResponseBody) SetMessage(v string) *ListIngressesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIngressesResponseBody) SetRequestId(v string) *ListIngressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIngressesResponseBody) SetSuccess(v bool) *ListIngressesResponseBody {
	s.Success = &v
	return s
}

func (s *ListIngressesResponseBody) SetTraceId(v string) *ListIngressesResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListIngressesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIngressesResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of routing rules.
	IngressList []*ListIngressesResponseBodyDataIngressList `json:"IngressList,omitempty" xml:"IngressList,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of queried entries.
	//
	// example:
	//
	// 10
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListIngressesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListIngressesResponseBodyData) GetIngressList() []*ListIngressesResponseBodyDataIngressList {
	return s.IngressList
}

func (s *ListIngressesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIngressesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListIngressesResponseBodyData) SetCurrentPage(v int32) *ListIngressesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListIngressesResponseBodyData) SetIngressList(v []*ListIngressesResponseBodyDataIngressList) *ListIngressesResponseBodyData {
	s.IngressList = v
	return s
}

func (s *ListIngressesResponseBodyData) SetPageSize(v int32) *ListIngressesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIngressesResponseBodyData) SetTotalSize(v int32) *ListIngressesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListIngressesResponseBodyData) Validate() error {
	if s.IngressList != nil {
		for _, item := range s.IngressList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIngressesResponseBodyDataIngressList struct {
	// The **CLB*	- certificate ID.
	//
	// example:
	//
	// 13624*****73809_16f8e549a20_1175189789_12****3210
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The **ALB*	- certificate IDs.
	//
	// example:
	//
	// 87***35-cn-hangzhou,812***3-cn-hangzhou
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// The CORS configuration.
	CorsConfig *ListIngressesResponseBodyDataIngressListCorsConfig `json:"CorsConfig,omitempty" xml:"CorsConfig,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 1593746835111
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default rule.
	DefaultRule *ListIngressesResponseBodyDataIngressListDefaultRule `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty" type:"Struct"`
	// The name of the routing rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The routing rule ID.
	//
	// example:
	//
	// 18
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The idle connection timeout period.
	//
	// example:
	//
	// 3
	IdleTimeout *int64 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The SLB listener port.
	//
	// example:
	//
	// 80
	ListenerPort *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The request forwarding protocol supported by the load balancing SLB. Valid values:
	//
	// - **HTTP**: Applicable to applications that need to identify data content.
	//
	// - **HTTPS**: Applicable to applications that require encrypted transmission.
	//
	// This parameter is optional in **CreateIngress*	- and **UpdateIngress**. If you do not specify this parameter when you create or update a gateway routing rule, this parameter is not returned when you invoke this API.
	//
	// example:
	//
	// HTTP
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The type of the load balancing SLB. Valid values:
	//
	// - **clb**: Classic Load Balancer (CLB), formerly known as SLB.
	//
	// - **alb**: Application Load Balancer (ALB).
	//
	// example:
	//
	// clb
	LoadBalanceType *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	// The MSE cloud-native gateway instance ID.
	//
	// example:
	//
	// gw-d5df01a1bae748f1a7c4e325d2fd****
	MseGatewayId *string `json:"MseGatewayId,omitempty" xml:"MseGatewayId,omitempty"`
	// The port corresponding to the service.
	//
	// example:
	//
	// 80
	MseGatewayPort *string `json:"MseGatewayPort,omitempty" xml:"MseGatewayPort,omitempty"`
	// The request forwarding protocol supported by the MSE cloud-native gateway. Valid values:
	//
	// - **HTTP**: Applicable to applications that need to identify data content.
	//
	// - **HTTPS**: Applicable to applications that require encrypted transmission.
	//
	// example:
	//
	// HTTP
	MseGatewayProtocol *string `json:"MseGatewayProtocol,omitempty" xml:"MseGatewayProtocol,omitempty"`
	// The name of the routing rule.
	//
	// example:
	//
	// lb-uf6jt0nu4z6ior943****-80-f5****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The request timeout period.
	//
	// example:
	//
	// 60
	RequestTimeout *int64 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The rules.
	Rules []*ListIngressesResponseBodyDataIngressListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The SLB ID.
	//
	// example:
	//
	// lb-uf62****6d13tq2u5
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The SLB type. Valid values:
	//
	// - **internet**: Internet.
	//
	// - **intranet**: Internal network.
	//
	// example:
	//
	// internet
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s ListIngressesResponseBodyDataIngressList) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesResponseBodyDataIngressList) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBodyDataIngressList) GetCertId() *string {
	return s.CertId
}

func (s *ListIngressesResponseBodyDataIngressList) GetCertIds() *string {
	return s.CertIds
}

func (s *ListIngressesResponseBodyDataIngressList) GetCorsConfig() *ListIngressesResponseBodyDataIngressListCorsConfig {
	return s.CorsConfig
}

func (s *ListIngressesResponseBodyDataIngressList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListIngressesResponseBodyDataIngressList) GetDefaultRule() *ListIngressesResponseBodyDataIngressListDefaultRule {
	return s.DefaultRule
}

func (s *ListIngressesResponseBodyDataIngressList) GetDescription() *string {
	return s.Description
}

func (s *ListIngressesResponseBodyDataIngressList) GetId() *int64 {
	return s.Id
}

func (s *ListIngressesResponseBodyDataIngressList) GetIdleTimeout() *int64 {
	return s.IdleTimeout
}

func (s *ListIngressesResponseBodyDataIngressList) GetListenerPort() *string {
	return s.ListenerPort
}

func (s *ListIngressesResponseBodyDataIngressList) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *ListIngressesResponseBodyDataIngressList) GetLoadBalanceType() *string {
	return s.LoadBalanceType
}

func (s *ListIngressesResponseBodyDataIngressList) GetMseGatewayId() *string {
	return s.MseGatewayId
}

func (s *ListIngressesResponseBodyDataIngressList) GetMseGatewayPort() *string {
	return s.MseGatewayPort
}

func (s *ListIngressesResponseBodyDataIngressList) GetMseGatewayProtocol() *string {
	return s.MseGatewayProtocol
}

func (s *ListIngressesResponseBodyDataIngressList) GetName() *string {
	return s.Name
}

func (s *ListIngressesResponseBodyDataIngressList) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListIngressesResponseBodyDataIngressList) GetRequestTimeout() *int64 {
	return s.RequestTimeout
}

func (s *ListIngressesResponseBodyDataIngressList) GetRules() []*ListIngressesResponseBodyDataIngressListRules {
	return s.Rules
}

func (s *ListIngressesResponseBodyDataIngressList) GetSlbId() *string {
	return s.SlbId
}

func (s *ListIngressesResponseBodyDataIngressList) GetSlbType() *string {
	return s.SlbType
}

func (s *ListIngressesResponseBodyDataIngressList) SetCertId(v string) *ListIngressesResponseBodyDataIngressList {
	s.CertId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetCertIds(v string) *ListIngressesResponseBodyDataIngressList {
	s.CertIds = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetCorsConfig(v *ListIngressesResponseBodyDataIngressListCorsConfig) *ListIngressesResponseBodyDataIngressList {
	s.CorsConfig = v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetCreateTime(v int64) *ListIngressesResponseBodyDataIngressList {
	s.CreateTime = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetDefaultRule(v *ListIngressesResponseBodyDataIngressListDefaultRule) *ListIngressesResponseBodyDataIngressList {
	s.DefaultRule = v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetDescription(v string) *ListIngressesResponseBodyDataIngressList {
	s.Description = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetId(v int64) *ListIngressesResponseBodyDataIngressList {
	s.Id = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetIdleTimeout(v int64) *ListIngressesResponseBodyDataIngressList {
	s.IdleTimeout = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetListenerPort(v string) *ListIngressesResponseBodyDataIngressList {
	s.ListenerPort = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetListenerProtocol(v string) *ListIngressesResponseBodyDataIngressList {
	s.ListenerProtocol = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetLoadBalanceType(v string) *ListIngressesResponseBodyDataIngressList {
	s.LoadBalanceType = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetMseGatewayId(v string) *ListIngressesResponseBodyDataIngressList {
	s.MseGatewayId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetMseGatewayPort(v string) *ListIngressesResponseBodyDataIngressList {
	s.MseGatewayPort = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetMseGatewayProtocol(v string) *ListIngressesResponseBodyDataIngressList {
	s.MseGatewayProtocol = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetName(v string) *ListIngressesResponseBodyDataIngressList {
	s.Name = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetNamespaceId(v string) *ListIngressesResponseBodyDataIngressList {
	s.NamespaceId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetRequestTimeout(v int64) *ListIngressesResponseBodyDataIngressList {
	s.RequestTimeout = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetRules(v []*ListIngressesResponseBodyDataIngressListRules) *ListIngressesResponseBodyDataIngressList {
	s.Rules = v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetSlbId(v string) *ListIngressesResponseBodyDataIngressList {
	s.SlbId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetSlbType(v string) *ListIngressesResponseBodyDataIngressList {
	s.SlbType = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) Validate() error {
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

type ListIngressesResponseBodyDataIngressListCorsConfig struct {
	// Specifies whether credentials are allowed.
	//
	// example:
	//
	// on
	AllowCredentials *string `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// The allowed headers.
	//
	// example:
	//
	// test_123
	AllowHeaders *string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty"`
	// The allowed methods.
	//
	// example:
	//
	// GET
	AllowMethods *string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty"`
	// The allowed origins.
	//
	// example:
	//
	// *
	AllowOrigin *string `json:"AllowOrigin,omitempty" xml:"AllowOrigin,omitempty"`
	// Specifies whether the feature is enabled.
	//
	// example:
	//
	// false
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The allowed exposed headers.
	//
	// example:
	//
	// test_123
	ExposeHeaders *string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty"`
	// The max age.
	//
	// example:
	//
	// 1000
	MaxAge *string `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
}

func (s ListIngressesResponseBodyDataIngressListCorsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesResponseBodyDataIngressListCorsConfig) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) GetAllowCredentials() *string {
	return s.AllowCredentials
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) GetAllowHeaders() *string {
	return s.AllowHeaders
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) GetAllowMethods() *string {
	return s.AllowMethods
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) GetAllowOrigin() *string {
	return s.AllowOrigin
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) GetEnable() *string {
	return s.Enable
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) GetExposeHeaders() *string {
	return s.ExposeHeaders
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) GetMaxAge() *string {
	return s.MaxAge
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) SetAllowCredentials(v string) *ListIngressesResponseBodyDataIngressListCorsConfig {
	s.AllowCredentials = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) SetAllowHeaders(v string) *ListIngressesResponseBodyDataIngressListCorsConfig {
	s.AllowHeaders = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) SetAllowMethods(v string) *ListIngressesResponseBodyDataIngressListCorsConfig {
	s.AllowMethods = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) SetAllowOrigin(v string) *ListIngressesResponseBodyDataIngressListCorsConfig {
	s.AllowOrigin = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) SetEnable(v string) *ListIngressesResponseBodyDataIngressListCorsConfig {
	s.Enable = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) SetExposeHeaders(v string) *ListIngressesResponseBodyDataIngressListCorsConfig {
	s.ExposeHeaders = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) SetMaxAge(v string) *ListIngressesResponseBodyDataIngressListCorsConfig {
	s.MaxAge = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListCorsConfig) Validate() error {
	return dara.Validate(s)
}

type ListIngressesResponseBodyDataIngressListDefaultRule struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The backend protocol.
	//
	// example:
	//
	// http
	BackendProtocol *string `json:"BackendProtocol,omitempty" xml:"BackendProtocol,omitempty"`
	// The container port.
	//
	// example:
	//
	// 8080
	ContainerPort *int32 `json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
}

func (s ListIngressesResponseBodyDataIngressListDefaultRule) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesResponseBodyDataIngressListDefaultRule) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) GetAppId() *string {
	return s.AppId
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) GetAppName() *string {
	return s.AppName
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) GetBackendProtocol() *string {
	return s.BackendProtocol
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) GetContainerPort() *int32 {
	return s.ContainerPort
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) SetAppId(v string) *ListIngressesResponseBodyDataIngressListDefaultRule {
	s.AppId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) SetAppName(v string) *ListIngressesResponseBodyDataIngressListDefaultRule {
	s.AppName = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) SetBackendProtocol(v string) *ListIngressesResponseBodyDataIngressListDefaultRule {
	s.BackendProtocol = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) SetContainerPort(v int32) *ListIngressesResponseBodyDataIngressListDefaultRule {
	s.ContainerPort = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListDefaultRule) Validate() error {
	return dara.Validate(s)
}

type ListIngressesResponseBodyDataIngressListRules struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The backend protocol.
	//
	// example:
	//
	// http
	BackendProtocol *string `json:"BackendProtocol,omitempty" xml:"BackendProtocol,omitempty"`
	// The container port.
	//
	// example:
	//
	// 8080
	ContainerPort *int32 `json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
	// The domain name.
	//
	// example:
	//
	// edas.site
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The path.
	//
	// example:
	//
	// /Path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The rewrite path.
	//
	// example:
	//
	// /${1}
	RewritePath *string `json:"RewritePath,omitempty" xml:"RewritePath,omitempty"`
	// The list of rule actions.
	RuleActions []*ListIngressesResponseBodyDataIngressListRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
}

func (s ListIngressesResponseBodyDataIngressListRules) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesResponseBodyDataIngressListRules) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBodyDataIngressListRules) GetAppId() *string {
	return s.AppId
}

func (s *ListIngressesResponseBodyDataIngressListRules) GetAppName() *string {
	return s.AppName
}

func (s *ListIngressesResponseBodyDataIngressListRules) GetBackendProtocol() *string {
	return s.BackendProtocol
}

func (s *ListIngressesResponseBodyDataIngressListRules) GetContainerPort() *int32 {
	return s.ContainerPort
}

func (s *ListIngressesResponseBodyDataIngressListRules) GetDomain() *string {
	return s.Domain
}

func (s *ListIngressesResponseBodyDataIngressListRules) GetPath() *string {
	return s.Path
}

func (s *ListIngressesResponseBodyDataIngressListRules) GetRewritePath() *string {
	return s.RewritePath
}

func (s *ListIngressesResponseBodyDataIngressListRules) GetRuleActions() []*ListIngressesResponseBodyDataIngressListRulesRuleActions {
	return s.RuleActions
}

func (s *ListIngressesResponseBodyDataIngressListRules) SetAppId(v string) *ListIngressesResponseBodyDataIngressListRules {
	s.AppId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRules) SetAppName(v string) *ListIngressesResponseBodyDataIngressListRules {
	s.AppName = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRules) SetBackendProtocol(v string) *ListIngressesResponseBodyDataIngressListRules {
	s.BackendProtocol = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRules) SetContainerPort(v int32) *ListIngressesResponseBodyDataIngressListRules {
	s.ContainerPort = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRules) SetDomain(v string) *ListIngressesResponseBodyDataIngressListRules {
	s.Domain = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRules) SetPath(v string) *ListIngressesResponseBodyDataIngressListRules {
	s.Path = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRules) SetRewritePath(v string) *ListIngressesResponseBodyDataIngressListRules {
	s.RewritePath = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRules) SetRuleActions(v []*ListIngressesResponseBodyDataIngressListRulesRuleActions) *ListIngressesResponseBodyDataIngressListRules {
	s.RuleActions = v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRules) Validate() error {
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

type ListIngressesResponseBodyDataIngressListRulesRuleActions struct {
	// The action configuration.
	//
	// example:
	//
	// {\\"host\\":\\"www.example.com\\",\\"path\\":\\"/example/text\\",\\"query\\":\\"x=1\\"}
	ActionConfig *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	// The action type.
	//
	// example:
	//
	// rewrite
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
}

func (s ListIngressesResponseBodyDataIngressListRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesResponseBodyDataIngressListRulesRuleActions) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBodyDataIngressListRulesRuleActions) GetActionConfig() *string {
	return s.ActionConfig
}

func (s *ListIngressesResponseBodyDataIngressListRulesRuleActions) GetActionType() *string {
	return s.ActionType
}

func (s *ListIngressesResponseBodyDataIngressListRulesRuleActions) SetActionConfig(v string) *ListIngressesResponseBodyDataIngressListRulesRuleActions {
	s.ActionConfig = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRulesRuleActions) SetActionType(v string) *ListIngressesResponseBodyDataIngressListRulesRuleActions {
	s.ActionType = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressListRulesRuleActions) Validate() error {
	return dara.Validate(s)
}
