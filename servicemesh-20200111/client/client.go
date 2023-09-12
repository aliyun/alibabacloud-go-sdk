// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SecretCreateRecordValue struct {
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SecretCreateRecordValue) String() string {
	return tea.Prettify(s)
}

func (s SecretCreateRecordValue) GoString() string {
	return s.String()
}

func (s *SecretCreateRecordValue) SetState(v string) *SecretCreateRecordValue {
	s.State = &v
	return s
}

func (s *SecretCreateRecordValue) SetClusterId(v string) *SecretCreateRecordValue {
	s.ClusterId = &v
	return s
}

func (s *SecretCreateRecordValue) SetMessage(v string) *SecretCreateRecordValue {
	s.Message = &v
	return s
}

type SecretDeleteRecordValue struct {
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SecretDeleteRecordValue) String() string {
	return tea.Prettify(s)
}

func (s SecretDeleteRecordValue) GoString() string {
	return s.String()
}

func (s *SecretDeleteRecordValue) SetState(v string) *SecretDeleteRecordValue {
	s.State = &v
	return s
}

func (s *SecretDeleteRecordValue) SetClusterId(v string) *SecretDeleteRecordValue {
	s.ClusterId = &v
	return s
}

func (s *SecretDeleteRecordValue) SetMessage(v string) *SecretDeleteRecordValue {
	s.Message = &v
	return s
}

type CCMVersionsValue struct {
	QueryState              *string `json:"QueryState,omitempty" xml:"QueryState,omitempty"`
	Version                 *string `json:"Version,omitempty" xml:"Version,omitempty"`
	SLBGracefulDrainSupport *bool   `json:"SLBGracefulDrainSupport,omitempty" xml:"SLBGracefulDrainSupport,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Message                 *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CCMVersionsValue) String() string {
	return tea.Prettify(s)
}

func (s CCMVersionsValue) GoString() string {
	return s.String()
}

func (s *CCMVersionsValue) SetQueryState(v string) *CCMVersionsValue {
	s.QueryState = &v
	return s
}

func (s *CCMVersionsValue) SetVersion(v string) *CCMVersionsValue {
	s.Version = &v
	return s
}

func (s *CCMVersionsValue) SetSLBGracefulDrainSupport(v bool) *CCMVersionsValue {
	s.SLBGracefulDrainSupport = &v
	return s
}

func (s *CCMVersionsValue) SetClusterId(v string) *CCMVersionsValue {
	s.ClusterId = &v
	return s
}

func (s *CCMVersionsValue) SetMessage(v string) *CCMVersionsValue {
	s.Message = &v
	return s
}

type UpgradeDetailGatewayStatusRecordValue struct {
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpgradeDetailGatewayStatusRecordValue) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDetailGatewayStatusRecordValue) GoString() string {
	return s.String()
}

func (s *UpgradeDetailGatewayStatusRecordValue) SetStatus(v string) *UpgradeDetailGatewayStatusRecordValue {
	s.Status = &v
	return s
}

func (s *UpgradeDetailGatewayStatusRecordValue) SetMessage(v string) *UpgradeDetailGatewayStatusRecordValue {
	s.Message = &v
	return s
}

func (s *UpgradeDetailGatewayStatusRecordValue) SetVersion(v string) *UpgradeDetailGatewayStatusRecordValue {
	s.Version = &v
	return s
}

type AddClusterIntoServiceMeshRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 添加集群时不检查目标集群是否存在istio-system namespace，一般用于自建istio 迁移ASM 场景
	IgnoreNamespaceCheck *bool   `json:"IgnoreNamespaceCheck,omitempty" xml:"IgnoreNamespaceCheck,omitempty"`
	ServiceMeshId        *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s AddClusterIntoServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s AddClusterIntoServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *AddClusterIntoServiceMeshRequest) SetClusterId(v string) *AddClusterIntoServiceMeshRequest {
	s.ClusterId = &v
	return s
}

func (s *AddClusterIntoServiceMeshRequest) SetIgnoreNamespaceCheck(v bool) *AddClusterIntoServiceMeshRequest {
	s.IgnoreNamespaceCheck = &v
	return s
}

func (s *AddClusterIntoServiceMeshRequest) SetServiceMeshId(v string) *AddClusterIntoServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

type AddClusterIntoServiceMeshResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddClusterIntoServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddClusterIntoServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *AddClusterIntoServiceMeshResponseBody) SetCode(v string) *AddClusterIntoServiceMeshResponseBody {
	s.Code = &v
	return s
}

func (s *AddClusterIntoServiceMeshResponseBody) SetMessage(v string) *AddClusterIntoServiceMeshResponseBody {
	s.Message = &v
	return s
}

func (s *AddClusterIntoServiceMeshResponseBody) SetRequestId(v string) *AddClusterIntoServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

type AddClusterIntoServiceMeshResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddClusterIntoServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddClusterIntoServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s AddClusterIntoServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *AddClusterIntoServiceMeshResponse) SetHeaders(v map[string]*string) *AddClusterIntoServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *AddClusterIntoServiceMeshResponse) SetStatusCode(v int32) *AddClusterIntoServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *AddClusterIntoServiceMeshResponse) SetBody(v *AddClusterIntoServiceMeshResponseBody) *AddClusterIntoServiceMeshResponse {
	s.Body = v
	return s
}

type AddVMIntoServiceMeshRequest struct {
	// The ID of the ECS instance.
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s AddVMIntoServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVMIntoServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *AddVMIntoServiceMeshRequest) SetEcsId(v string) *AddVMIntoServiceMeshRequest {
	s.EcsId = &v
	return s
}

func (s *AddVMIntoServiceMeshRequest) SetServiceMeshId(v string) *AddVMIntoServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

type AddVMIntoServiceMeshResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVMIntoServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVMIntoServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *AddVMIntoServiceMeshResponseBody) SetRequestId(v string) *AddVMIntoServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

type AddVMIntoServiceMeshResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddVMIntoServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVMIntoServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVMIntoServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *AddVMIntoServiceMeshResponse) SetHeaders(v map[string]*string) *AddVMIntoServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *AddVMIntoServiceMeshResponse) SetStatusCode(v int32) *AddVMIntoServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVMIntoServiceMeshResponse) SetBody(v *AddVMIntoServiceMeshResponseBody) *AddVMIntoServiceMeshResponse {
	s.Body = v
	return s
}

type CreateASMGatewayRequest struct {
	// The YAML content that is used to create the ASM gateway.
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s CreateASMGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateASMGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateASMGatewayRequest) SetBody(v string) *CreateASMGatewayRequest {
	s.Body = &v
	return s
}

func (s *CreateASMGatewayRequest) SetIstioGatewayName(v string) *CreateASMGatewayRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateASMGatewayRequest) SetServiceMeshId(v string) *CreateASMGatewayRequest {
	s.ServiceMeshId = &v
	return s
}

type CreateASMGatewayResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateASMGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateASMGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateASMGatewayResponseBody) SetRequestId(v string) *CreateASMGatewayResponseBody {
	s.RequestId = &v
	return s
}

type CreateASMGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateASMGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateASMGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateASMGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateASMGatewayResponse) SetHeaders(v map[string]*string) *CreateASMGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateASMGatewayResponse) SetStatusCode(v int32) *CreateASMGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateASMGatewayResponse) SetBody(v *CreateASMGatewayResponseBody) *CreateASMGatewayResponse {
	s.Body = v
	return s
}

type CreateGatewaySecretRequest struct {
	// The content of the certificate.
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The private key of the certificate.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s CreateGatewaySecretRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewaySecretRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewaySecretRequest) SetCert(v string) *CreateGatewaySecretRequest {
	s.Cert = &v
	return s
}

func (s *CreateGatewaySecretRequest) SetIstioGatewayName(v string) *CreateGatewaySecretRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateGatewaySecretRequest) SetKey(v string) *CreateGatewaySecretRequest {
	s.Key = &v
	return s
}

func (s *CreateGatewaySecretRequest) SetSecretName(v string) *CreateGatewaySecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateGatewaySecretRequest) SetServiceMeshId(v string) *CreateGatewaySecretRequest {
	s.ServiceMeshId = &v
	return s
}

type CreateGatewaySecretResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The record of creating the secret.
	SecretCreateRecord map[string]*SecretCreateRecordValue `json:"SecretCreateRecord,omitempty" xml:"SecretCreateRecord,omitempty"`
}

func (s CreateGatewaySecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewaySecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewaySecretResponseBody) SetRequestId(v string) *CreateGatewaySecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewaySecretResponseBody) SetSecretCreateRecord(v map[string]*SecretCreateRecordValue) *CreateGatewaySecretResponseBody {
	s.SecretCreateRecord = v
	return s
}

type CreateGatewaySecretResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewaySecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewaySecretResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewaySecretResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewaySecretResponse) SetHeaders(v map[string]*string) *CreateGatewaySecretResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewaySecretResponse) SetStatusCode(v int32) *CreateGatewaySecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewaySecretResponse) SetBody(v *CreateGatewaySecretResponseBody) *CreateGatewaySecretResponse {
	s.Body = v
	return s
}

type CreateIstioGatewayDomainsRequest struct {
	// The name of the secret that contains the Transport Layer Security (TLS) certificate and certificate authority (CA) certificate.
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// Specifies whether to forcibly use TLS to protect connection security.
	//
	// *   `true`: forcibly uses TLS to protect connection security.
	// *   `false`: does not forcibly use TLS to protect connection security.
	ForceHttps *bool `json:"ForceHttps,omitempty" xml:"ForceHttps,omitempty"`
	// The domain names of the one or more hosts that are exposed by the ASM gateway. Separate multiple domain names with commas (,).
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The maximum number of ASM gateways to query.
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The port that is provided by the ASM gateway.
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// The name of the port.
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The type of the protocol. Valid values: `HTTP`, `HTTPS`, `GRPC`, `HTTP2`, `MONGO`, `TCP`, and `TLS`.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s CreateIstioGatewayDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayDomainsRequest) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayDomainsRequest) SetCredential(v string) *CreateIstioGatewayDomainsRequest {
	s.Credential = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetForceHttps(v bool) *CreateIstioGatewayDomainsRequest {
	s.ForceHttps = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetHosts(v string) *CreateIstioGatewayDomainsRequest {
	s.Hosts = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetIstioGatewayName(v string) *CreateIstioGatewayDomainsRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetLimit(v string) *CreateIstioGatewayDomainsRequest {
	s.Limit = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetNamespace(v string) *CreateIstioGatewayDomainsRequest {
	s.Namespace = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetNumber(v int32) *CreateIstioGatewayDomainsRequest {
	s.Number = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetPortName(v string) *CreateIstioGatewayDomainsRequest {
	s.PortName = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetProtocol(v string) *CreateIstioGatewayDomainsRequest {
	s.Protocol = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetServiceMeshId(v string) *CreateIstioGatewayDomainsRequest {
	s.ServiceMeshId = &v
	return s
}

type CreateIstioGatewayDomainsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIstioGatewayDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayDomainsResponseBody) SetRequestId(v string) *CreateIstioGatewayDomainsResponseBody {
	s.RequestId = &v
	return s
}

type CreateIstioGatewayDomainsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIstioGatewayDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIstioGatewayDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayDomainsResponse) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayDomainsResponse) SetHeaders(v map[string]*string) *CreateIstioGatewayDomainsResponse {
	s.Headers = v
	return s
}

func (s *CreateIstioGatewayDomainsResponse) SetStatusCode(v int32) *CreateIstioGatewayDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIstioGatewayDomainsResponse) SetBody(v *CreateIstioGatewayDomainsResponseBody) *CreateIstioGatewayDomainsResponse {
	s.Body = v
	return s
}

type CreateIstioGatewayRoutesRequest struct {
	// The description of the routing rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the routing rule to be created for the ASM gateway.
	GatewayRoute *CreateIstioGatewayRoutesRequestGatewayRoute `json:"GatewayRoute,omitempty" xml:"GatewayRoute,omitempty" type:"Struct"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// *   `0`: The routing rule is valid.
	// *   `1`: The routing rule is invalid.
	// *   `2`: An error occurs during the creation or update of the routing rule.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateIstioGatewayRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequest) SetDescription(v string) *CreateIstioGatewayRoutesRequest {
	s.Description = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetGatewayRoute(v *CreateIstioGatewayRoutesRequestGatewayRoute) *CreateIstioGatewayRoutesRequest {
	s.GatewayRoute = v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetIstioGatewayName(v string) *CreateIstioGatewayRoutesRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetPriority(v int32) *CreateIstioGatewayRoutesRequest {
	s.Priority = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetServiceMeshId(v string) *CreateIstioGatewayRoutesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequest) SetStatus(v int32) *CreateIstioGatewayRoutesRequest {
	s.Status = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRoute struct {
	// The requested domain names.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The advanced settings for routing HTTP traffic.
	HTTPAdvancedOptions *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions `json:"HTTPAdvancedOptions,omitempty" xml:"HTTPAdvancedOptions,omitempty" type:"Struct"`
	// The matching rules for traffic routing.
	MatchRequest *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest `json:"MatchRequest,omitempty" xml:"MatchRequest,omitempty" type:"Struct"`
	// The namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// A JSON string. This parameter corresponds to the three routing types in virtual services and provides configuration entries for advanced features. The value of this parameter overwrites the configurations in RouteName, RouteType, MatchRequest, and HTTPAdvancedOptions.
	RawVSRoute interface{} `json:"RawVSRoute,omitempty" xml:"RawVSRoute,omitempty"`
	// The endpoints of destination services for Layer 4 weighted routing.
	RouteDestinations []*CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations `json:"RouteDestinations,omitempty" xml:"RouteDestinations,omitempty" type:"Repeated"`
	// The name of the routing rule.
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The type of the traffic to be routed. Valid values: `HTTP`, `TLS`, and `TCP`.
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRoute) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRoute) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetDomains(v []*string) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.Domains = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetHTTPAdvancedOptions(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.HTTPAdvancedOptions = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetMatchRequest(v *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.MatchRequest = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetNamespace(v string) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.Namespace = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetRawVSRoute(v interface{}) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.RawVSRoute = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetRouteDestinations(v []*CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteDestinations = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetRouteName(v string) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteName = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRoute) SetRouteType(v string) *CreateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteType = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions struct {
	// The virtual service that defines traffic routing.
	Delegate *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate `json:"Delegate,omitempty" xml:"Delegate,omitempty" type:"Struct"`
	// The configurations of fault injection.
	Fault *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault `json:"Fault,omitempty" xml:"Fault,omitempty" type:"Struct"`
	// The HTTP redirection rule.
	HTTPRedirect *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect `json:"HTTPRedirect,omitempty" xml:"HTTPRedirect,omitempty" type:"Struct"`
	// The configurations for mirroring HTTP traffic to another destination in addition to forwarding requests to the specified destination.
	Mirror *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror `json:"Mirror,omitempty" xml:"Mirror,omitempty" type:"Struct"`
	// The percentage of requests that are mirrored to another destination except for the original destination.
	MirrorPercentage *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage `json:"MirrorPercentage,omitempty" xml:"MirrorPercentage,omitempty" type:"Struct"`
	// The configurations of retries for failed requests.
	Retries *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries `json:"Retries,omitempty" xml:"Retries,omitempty" type:"Struct"`
	// The configurations for rewriting the virtual service.
	Rewrite *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Struct"`
	// The timeout period for requests.
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetDelegate(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Delegate = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetFault(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Fault = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetHTTPRedirect(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.HTTPRedirect = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetMirror(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Mirror = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetMirrorPercentage(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.MirrorPercentage = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetRetries(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Retries = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetRewrite(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Rewrite = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetTimeout(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Timeout = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate struct {
	// The name of the virtual service.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the virtual service belongs.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) SetName(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	s.Name = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) SetNamespace(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	s.Namespace = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault struct {
	// The configurations for aborting requests with specified error codes.
	Abort *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort `json:"Abort,omitempty" xml:"Abort,omitempty" type:"Struct"`
	// The duration to delay a request.
	Delay *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay `json:"Delay,omitempty" xml:"Delay,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) SetAbort(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	s.Abort = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) SetDelay(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	s.Delay = v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort struct {
	// The HTTP status code.
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The percentage of requests that are aborted with the specified error code.
	Percentage *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) SetHttpStatus(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	s.HttpStatus = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) SetPercentage(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	s.Percentage = v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage struct {
	// The percentage of requests that are aborted with the specified error code, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) SetValue(v float32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage {
	s.Value = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay struct {
	// The fixed duration for request delay.
	FixedDelay *string `json:"FixedDelay,omitempty" xml:"FixedDelay,omitempty"`
	// The percentage of requests to which the delay fault is injected.
	Percentage *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) SetFixedDelay(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	s.FixedDelay = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) SetPercentage(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	s.Percentage = v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage struct {
	// The percentage of requests to which the delay fault is injected, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) SetValue(v float32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage {
	s.Value = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect struct {
	// The value to be used to overwrite the value of the Authority or Host header during redirection.``
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The HTTP status code to be used to indicate URL redirection. Default value: 301.
	RedirectCode *int32 `json:"RedirectCode,omitempty" xml:"RedirectCode,omitempty"`
	// The value to be used to overwrite the URL path during redirection.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetAuthority(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.Authority = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetRedirectCode(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.RedirectCode = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetUri(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.Uri = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror struct {
	// The name of the service defined in the service registry.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The name of the service subset.
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) SetHost(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	s.Host = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) SetSubset(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	s.Subset = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage struct {
	// The percentage of requests that are mirrored to another destination except for the original destination, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) SetValue(v float32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage {
	s.Value = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries struct {
	// The number of retries that are allowed for a request.
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// The timeout period for each retry. Example: `5s`.
	PerTryTimeout *string `json:"PerTryTimeout,omitempty" xml:"PerTryTimeout,omitempty"`
	// The condition for retries. Example: `connect-failure,refused-stream,503`.
	RetryOn *string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty"`
	// Specifies whether to allow retries to other localities.
	RetryRemoteLocalities *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities `json:"RetryRemoteLocalities,omitempty" xml:"RetryRemoteLocalities,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetAttempts(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.Attempts = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetPerTryTimeout(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.PerTryTimeout = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetRetryOn(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.RetryOn = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetRetryRemoteLocalities(v *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.RetryRemoteLocalities = v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities struct {
	// Specifies whether to allow retries to other localities. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	Value *bool `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) SetValue(v bool) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	s.Value = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite struct {
	// The value to be used to overwrite the value of the Authority or Host header.
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The value to be used to overwrite the path or prefix of the URI.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) SetAuthority(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	s.Authority = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) SetUri(v string) *CreateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	s.Uri = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest struct {
	// The request headers to be matched.
	Headers []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The ports of destination services for Layer 4 weighted routing.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The matching rule for Transport Layer Security (TLS) traffic.
	TLSMatchAttributes []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes `json:"TLSMatchAttributes,omitempty" xml:"TLSMatchAttributes,omitempty" type:"Repeated"`
	// The matching rule for URIs.
	URI *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI `json:"URI,omitempty" xml:"URI,omitempty" type:"Struct"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetHeaders(v []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.Headers = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetPorts(v []*int32) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.Ports = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetTLSMatchAttributes(v []*CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.TLSMatchAttributes = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetURI(v *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.URI = v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders struct {
	// The header value to be matched.
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode for the header value. Valid values:
	//
	// *   `exact`: exact match
	// *   `prefix`: match by prefix
	// *   `regex`: match by regular expression
	MatchingMode *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
	// The header key to be matched.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetMatchingContent(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.MatchingContent = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetMatchingMode(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.MatchingMode = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetName(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.Name = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes struct {
	// The Server Name Indication (SNI) values to be matched.
	SNIHosts []*string `json:"SNIHosts,omitempty" xml:"SNIHosts,omitempty" type:"Repeated"`
	// The TLS port.
	TLSPort *int32 `json:"TLSPort,omitempty" xml:"TLSPort,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) SetSNIHosts(v []*string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	s.SNIHosts = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) SetTLSPort(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	s.TLSPort = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI struct {
	// The content to be matched.
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode for the routing rule. Valid values:
	//
	// *   `exact`: exact match
	// *   `prefix`: match by prefix
	// *   `regex`: match by regular expression
	MatchingMode *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) SetMatchingContent(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	s.MatchingContent = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) SetMatchingMode(v string) *CreateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	s.MatchingMode = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations struct {
	// The unique endpoint of the destination service to which the specified requests are sent.
	Destination *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Struct"`
	// The weight of the service subset.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) SetDestination(v *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	s.Destination = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) SetWeight(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	s.Weight = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination struct {
	// The name of the service defined in the service registry.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port.
	Port *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Struct"`
	// The name of the service subset.
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetHost(v string) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Host = &v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetPort(v *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Port = v
	return s
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetSubset(v string) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Subset = &v
	return s
}

type CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort struct {
	// The port number.
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) SetNumber(v int32) *CreateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort {
	s.Number = &v
	return s
}

type CreateIstioGatewayRoutesShrinkRequest struct {
	// The description of the routing rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the routing rule to be created for the ASM gateway.
	GatewayRouteShrink *string `json:"GatewayRoute,omitempty" xml:"GatewayRoute,omitempty"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// *   `0`: The routing rule is valid.
	// *   `1`: The routing rule is invalid.
	// *   `2`: An error occurs during the creation or update of the routing rule.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateIstioGatewayRoutesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetDescription(v string) *CreateIstioGatewayRoutesShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetGatewayRouteShrink(v string) *CreateIstioGatewayRoutesShrinkRequest {
	s.GatewayRouteShrink = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetIstioGatewayName(v string) *CreateIstioGatewayRoutesShrinkRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetPriority(v int32) *CreateIstioGatewayRoutesShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetServiceMeshId(v string) *CreateIstioGatewayRoutesShrinkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetStatus(v int32) *CreateIstioGatewayRoutesShrinkRequest {
	s.Status = &v
	return s
}

type CreateIstioGatewayRoutesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIstioGatewayRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesResponseBody) SetRequestId(v string) *CreateIstioGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

type CreateIstioGatewayRoutesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIstioGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIstioGatewayRoutesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIstioGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesResponse) SetHeaders(v map[string]*string) *CreateIstioGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *CreateIstioGatewayRoutesResponse) SetStatusCode(v int32) *CreateIstioGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIstioGatewayRoutesResponse) SetBody(v *CreateIstioGatewayRoutesResponseBody) *CreateIstioGatewayRoutesResponse {
	s.Body = v
	return s
}

type CreateServiceMeshRequest struct {
	// Specifies whether to enable access log collection. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	AccessLogEnabled *bool `json:"AccessLogEnabled,omitempty" xml:"AccessLogEnabled,omitempty"`
	// Specifies whether to enable access log collection. Valid values:
	//
	// *   "": disables access log collection.
	// *   `/dev/stdout`: enables access log collection. Access logs are written to /dev/stdout.
	AccessLogFile *string `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	// Custom fields of access logs. To set this parameter, you must enable access log collection. Otherwise, you cannot set this parameter. The value must be a JSON string. The following key values must be contained: authority_for, bytes_received, bytes_sent, downstream_local_address, downstream_remote_address, duration, istio_policy_status, method, path, protocol, requested_server_name, response_code, response_flags, route_name, start_time, trace_id, upstream_cluster, upstream_host, upstream_local_address, upstream_service_time, upstream_transport_failure_reason, user_agent, and x_forwarded_for.
	AccessLogFormat *string `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	// The SLS project from which access logs are collected.
	AccessLogProject *string `json:"AccessLogProject,omitempty" xml:"AccessLogProject,omitempty"`
	// Specifies whether to enable gRPC Access Log Service (ALS) of Envoy. gRPC is short for Google Remote Procedure Call. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	AccessLogServiceEnabled *bool `json:"AccessLogServiceEnabled,omitempty" xml:"AccessLogServiceEnabled,omitempty"`
	// The endpoint of Envoy\"s gRPC ALS.
	AccessLogServiceHost *string `json:"AccessLogServiceHost,omitempty" xml:"AccessLogServiceHost,omitempty"`
	// The port of Envoy\"s gRPC ALS.
	AccessLogServicePort *int32 `json:"AccessLogServicePort,omitempty" xml:"AccessLogServicePort,omitempty"`
	// The type of the SLB instance that is bound to Istio Pilot. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`, `slb.s3.small`, `slb.s3.medium`, and `slb.s3.large`.
	ApiServerLoadBalancerSpec *string `json:"ApiServerLoadBalancerSpec,omitempty" xml:"ApiServerLoadBalancerSpec,omitempty"`
	// Specifies whether to expose the API server to the Internet. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	//
	// > If you set this parameter to `false`, the API server cannot be accessed over the Internet.
	ApiServerPublicEip *bool `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	// The name of the Log Service project that is used for mesh audit.
	//
	// Default value: mesh-log-{ASM instance ID}.
	AuditProject *string `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	// Specifies whether to enable auto-renewal for the SLB instance if the SLB instance uses the subscription billing method. Valid values:
	//
	// - true
	//
	// - false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period of the SLB instance. This parameter is valid only if the `ChargeType` parameter is set to `PrePay`. If the original subscription period of the SLB instance is less than one year, the value of this parameter indicates the number of months for auto-renewal. If the original subscription period of the SLB instance is more than one year, the value of this parameter indicates the number of years for auto-renewal.
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to allow the Kubernetes API of clusters on the data plane to access Istio resources. The version of the ASM instance must be V1.9.7.93 or later. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	CRAggregationEnabled *bool `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	// The billing method of the SLB instance. Valid values:
	//
	// *   `PayOnDemand`: pay-as-you-go.
	// *   `PrePay`: subscription.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The edition of the ASM instance. Valid values:
	//
	// - `standard`: Standard Edition
	//
	// - `enterprise`: Enterprise Edition
	//
	// - `ultimate`: Ultimate Edition
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// Specifies whether to enable the external service registry. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	ConfigSourceEnabled *bool `json:"ConfigSourceEnabled,omitempty" xml:"ConfigSourceEnabled,omitempty"`
	// The instance ID of the Nacos registry.
	ConfigSourceNacosID *string `json:"ConfigSourceNacosID,omitempty" xml:"ConfigSourceNacosID,omitempty"`
	// Specifies whether to enable the collection of control plane logs. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	ControlPlaneLogEnabled *bool `json:"ControlPlaneLogEnabled,omitempty" xml:"ControlPlaneLogEnabled,omitempty"`
	// The name of the Log Service project that is used to collect the logs of the control plane.
	ControlPlaneLogProject *string `json:"ControlPlaneLogProject,omitempty" xml:"ControlPlaneLogProject,omitempty"`
	// Specifies whether to use a custom Prometheus instance. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	CustomizedPrometheus *bool `json:"CustomizedPrometheus,omitempty" xml:"CustomizedPrometheus,omitempty"`
	// Specifies whether to use a self-managed Zipkin system to collect tracing data. Valid values:
	//
	// *   `true`: uses a self-managed Zipkin system to collect tracing data.
	// *   `false`: uses Alibaba Cloud Tracing Analysis to collect tracing data.
	//
	// Default value: `false`.
	CustomizedZipkin *bool `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	// Specifies whether to enable the DNS proxy feature. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	DNSProxyingEnabled *bool `json:"DNSProxyingEnabled,omitempty" xml:"DNSProxyingEnabled,omitempty"`
	// Specifies whether to enable Dubbo Filter. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	DubboFilterEnabled *bool `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	// The edition of the ASM instance.
	Edition       *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	EnableAmbient *bool   `json:"EnableAmbient,omitempty" xml:"EnableAmbient,omitempty"`
	// Specifies whether to enable the mesh audit feature. To enable this feature, make sure that you have activated [Log Service](https://sls.console.aliyun.com/). Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	EnableAudit *bool `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	// Specifies whether to enable the rollback feature for Istio resources. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	EnableCRHistory *bool `json:"EnableCRHistory,omitempty" xml:"EnableCRHistory,omitempty"`
	// Specifies whether to enable Secret Discovery Service (SDS). Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	EnableSDSServer *bool `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
	// The IP ranges in CIDR form to be excluded from redirection to the sidecar proxy in the ASM instance.
	ExcludeIPRanges *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	// The inbound ports to be excluded from redirection to the sidecar proxy in the ASM instance. Separate multiple port numbers with commas (,).
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The outbound ports to be excluded from redirection to the sidecar proxy in the ASM instance. Separate multiple port numbers with commas (,).
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// The existing CA certificate, which is encoded in Base64. This parameter is used in scenarios where you migrate open source Istio to ASM. It specifies the content of the ca-cert.pem file in the istio-ca-secret secret. The secret is in the istio-system namespace of the Kubernetes cluster where the open source Istio is installed.
	ExistingCaCert *string `json:"ExistingCaCert,omitempty" xml:"ExistingCaCert,omitempty"`
	// The existing CA key, which is encoded in Base64. This parameter is used in scenarios where you migrate open source Istio to ASM. It specifies the content of the ca-key.pem file in the istio-ca-secret secret. The secret is in the istio-system namespace of the Kubernetes cluster where the open source Istio is installed.
	ExistingCaKey *string `json:"ExistingCaKey,omitempty" xml:"ExistingCaKey,omitempty"`
	// The type of the existing CA certificate. Valid values:
	//
	// *   1: Self-signed certificate generated by istiod. The certificate corresponds to the secret named istio-ca-secret in the istio-system namespace. If you use this type of certificate, you must set the `ExistingCaCert` and `ExsitingCaKey` parameters.
	// *   2: Administrator-specified certificate. For more information, see [plugin ca cert](https://istio.io/latest/docs/tasks/security/cert-management/plugin-ca-cert/). In most cases, the certificate corresponds to the secret named cacerts in the istio-system namespace. If you use this type of certificate, you must set the `ExisingRootCaCert` and `ExisingRootCaKey` parameters.
	ExistingCaType *string `json:"ExistingCaType,omitempty" xml:"ExistingCaType,omitempty"`
	// The existing root certificate, which is encoded in Base64.
	ExistingRootCaCert *string `json:"ExistingRootCaCert,omitempty" xml:"ExistingRootCaCert,omitempty"`
	// The private key that corresponds to the root certificate, which is encoded in Base64.
	ExistingRootCaKey *string `json:"ExistingRootCaKey,omitempty" xml:"ExistingRootCaKey,omitempty"`
	// Specifies whether to enable gateway configuration filtering. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	FilterGatewayClusterConfig *bool `json:"FilterGatewayClusterConfig,omitempty" xml:"FilterGatewayClusterConfig,omitempty"`
	// Specifies whether to enable Gateway API. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	GatewayAPIEnabled *bool `json:"GatewayAPIEnabled,omitempty" xml:"GatewayAPIEnabled,omitempty"`
	// After this ASM instance is successfully created, automatically add an ACK cluster to it.
	// Make sure this ASM instance and ACK cluster have same VPC, VSwitch, cluster domain.
	GuestCluster *string `json:"GuestCluster,omitempty" xml:"GuestCluster,omitempty"`
	// The IP ranges in CIDR form for which traffic is to be redirected to the sidecar proxy in the ASM instance.
	IncludeIPRanges *string `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	// The Istio version of the ASM instance.
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	// Specifies whether to enable the mesh topology feature. To enable this feature, make sure that you have enabled Prometheus monitoring. If Prometheus monitoring is disabled, you must set this parameter to `false`.`` Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	KialiEnabled *bool `json:"KialiEnabled,omitempty" xml:"KialiEnabled,omitempty"`
	// The configurations for the access to the nearest instance.
	LocalityLBConf *string `json:"LocalityLBConf,omitempty" xml:"LocalityLBConf,omitempty"`
	// Specifies whether to route traffic to the nearest instance. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	LocalityLoadBalancing *bool `json:"LocalityLoadBalancing,omitempty" xml:"LocalityLoadBalancing,omitempty"`
	// Specifies whether to enable Microservices Engine (MSE). Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	MSEEnabled *bool `json:"MSEEnabled,omitempty" xml:"MSEEnabled,omitempty"`
	// Specifies whether to enable MultiBuffer-based Transport Layer Security (TLS) acceleration. Valid values:
	//
	// - `true`
	//
	// - `false`
	//
	//
	// Default value: `true`
	MultiBufferEnabled *bool `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	// The pull-request latency. Default value: 30. Unit: seconds.
	MultiBufferPollDelay *string `json:"MultiBufferPollDelay,omitempty" xml:"MultiBufferPollDelay,omitempty"`
	// Specifies whether to enable MySQL Filter. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	MysqlFilterEnabled *bool `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	// The name of the ASM instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of CPU cores that are available to the OPA container.
	OPALimitCPU *string `json:"OPALimitCPU,omitempty" xml:"OPALimitCPU,omitempty"`
	// The maximum size of the memory that is available to the OPA container. You can specify the parameter value in the standard quantity representation form used by Kubernetes. 1 Mi equals 1,024 KB.
	OPALimitMemory *string `json:"OPALimitMemory,omitempty" xml:"OPALimitMemory,omitempty"`
	// The log level of the OPA container.
	OPALogLevel *string `json:"OPALogLevel,omitempty" xml:"OPALogLevel,omitempty"`
	// The minimum number of CPU cores that are required by the OPA container. You can specify the parameter value in the standard representation form of CPUs in Kubernetes. For example, if you set the value to 1, one CPU core is required.
	OPARequestCPU *string `json:"OPARequestCPU,omitempty" xml:"OPARequestCPU,omitempty"`
	// The minimum size of the memory that is required by the OPA container. You can specify the parameter value in the standard quantity representation form used by Kubernetes. 1 Mi equals 1,024 KB.
	OPARequestMemory *string `json:"OPARequestMemory,omitempty" xml:"OPARequestMemory,omitempty"`
	// Specifies whether to enable the OPA plug-in. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	OpaEnabled *bool `json:"OpaEnabled,omitempty" xml:"OpaEnabled,omitempty"`
	// Specifies whether to install the Open Policy Agent (OPA) plug-in. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	OpenAgentPolicy *bool `json:"OpenAgentPolicy,omitempty" xml:"OpenAgentPolicy,omitempty"`
	// The auto-renewal period of the SLB instance. This parameter is valid only if `ChargeType` is set to `PrePaid`. The value of this parameter indicates the purchased month of the SLB instance when the subscription billing method is used. For example, if the subscription period is one year, set this parameter to 12.
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The type of the SLB instance that is bound to Istio Pilot. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`, `slb.s3.small`, `slb.s3.medium`, and `slb.s3.large`.
	PilotLoadBalancerSpec *string `json:"PilotLoadBalancerSpec,omitempty" xml:"PilotLoadBalancerSpec,omitempty"`
	// The endpoint of the custom Prometheus instance.
	PrometheusUrl *string `json:"PrometheusUrl,omitempty" xml:"PrometheusUrl,omitempty"`
	// The maximum number of CPU cores that are available to the proxy container.
	ProxyLimitCPU *string `json:"ProxyLimitCPU,omitempty" xml:"ProxyLimitCPU,omitempty"`
	// The maximum size of the memory that is available to the proxy container.
	ProxyLimitMemory *string `json:"ProxyLimitMemory,omitempty" xml:"ProxyLimitMemory,omitempty"`
	// The minimum number of CPU cores that are required by the proxy container.
	ProxyRequestCPU *string `json:"ProxyRequestCPU,omitempty" xml:"ProxyRequestCPU,omitempty"`
	// The minimum size of the memory that is required by the proxy container.
	ProxyRequestMemory *string `json:"ProxyRequestMemory,omitempty" xml:"ProxyRequestMemory,omitempty"`
	// Specifies whether to enable Redis Filter. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	RedisFilterEnabled *bool `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
	// The ID of the region in which the ASM instance resides.
	RegionId *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag      []*CreateServiceMeshRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to enable Prometheus monitoring. We recommend that you use Prometheus Service of [Application Real-Time Monitoring Service (ARMS)](https://arms.console.aliyun.com/). Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	Telemetry *bool `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	// Specifies whether to enable Thrift Filter. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	ThriftFilterEnabled *bool `json:"ThriftFilterEnabled,omitempty" xml:"ThriftFilterEnabled,omitempty"`
	// The sampling percentage of Tracing Analysis.
	TraceSampling *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
	// Specifies whether to enable the Tracing Analysis feature. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	Tracing *bool `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
	// Specifies whether to use an existing CA certificate and private key.
	UseExistingCA *bool `json:"UseExistingCA,omitempty" xml:"UseExistingCA,omitempty"`
	// The ID of the vSwitch to which the ASM instance is connected.
	VSwitches *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the ASM instance resides.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Specifies whether to enable WebAssembly Filter. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	WebAssemblyFilterEnabled *bool `json:"WebAssemblyFilterEnabled,omitempty" xml:"WebAssemblyFilterEnabled,omitempty"`
}

func (s CreateServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshRequest) SetAccessLogEnabled(v bool) *CreateServiceMeshRequest {
	s.AccessLogEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAccessLogFile(v string) *CreateServiceMeshRequest {
	s.AccessLogFile = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAccessLogFormat(v string) *CreateServiceMeshRequest {
	s.AccessLogFormat = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAccessLogProject(v string) *CreateServiceMeshRequest {
	s.AccessLogProject = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAccessLogServiceEnabled(v bool) *CreateServiceMeshRequest {
	s.AccessLogServiceEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAccessLogServiceHost(v string) *CreateServiceMeshRequest {
	s.AccessLogServiceHost = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAccessLogServicePort(v int32) *CreateServiceMeshRequest {
	s.AccessLogServicePort = &v
	return s
}

func (s *CreateServiceMeshRequest) SetApiServerLoadBalancerSpec(v string) *CreateServiceMeshRequest {
	s.ApiServerLoadBalancerSpec = &v
	return s
}

func (s *CreateServiceMeshRequest) SetApiServerPublicEip(v bool) *CreateServiceMeshRequest {
	s.ApiServerPublicEip = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAuditProject(v string) *CreateServiceMeshRequest {
	s.AuditProject = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAutoRenew(v bool) *CreateServiceMeshRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateServiceMeshRequest) SetAutoRenewPeriod(v int32) *CreateServiceMeshRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateServiceMeshRequest) SetCRAggregationEnabled(v bool) *CreateServiceMeshRequest {
	s.CRAggregationEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetChargeType(v string) *CreateServiceMeshRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateServiceMeshRequest) SetClusterSpec(v string) *CreateServiceMeshRequest {
	s.ClusterSpec = &v
	return s
}

func (s *CreateServiceMeshRequest) SetConfigSourceEnabled(v bool) *CreateServiceMeshRequest {
	s.ConfigSourceEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetConfigSourceNacosID(v string) *CreateServiceMeshRequest {
	s.ConfigSourceNacosID = &v
	return s
}

func (s *CreateServiceMeshRequest) SetControlPlaneLogEnabled(v bool) *CreateServiceMeshRequest {
	s.ControlPlaneLogEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetControlPlaneLogProject(v string) *CreateServiceMeshRequest {
	s.ControlPlaneLogProject = &v
	return s
}

func (s *CreateServiceMeshRequest) SetCustomizedPrometheus(v bool) *CreateServiceMeshRequest {
	s.CustomizedPrometheus = &v
	return s
}

func (s *CreateServiceMeshRequest) SetCustomizedZipkin(v bool) *CreateServiceMeshRequest {
	s.CustomizedZipkin = &v
	return s
}

func (s *CreateServiceMeshRequest) SetDNSProxyingEnabled(v bool) *CreateServiceMeshRequest {
	s.DNSProxyingEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetDubboFilterEnabled(v bool) *CreateServiceMeshRequest {
	s.DubboFilterEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetEdition(v string) *CreateServiceMeshRequest {
	s.Edition = &v
	return s
}

func (s *CreateServiceMeshRequest) SetEnableAmbient(v bool) *CreateServiceMeshRequest {
	s.EnableAmbient = &v
	return s
}

func (s *CreateServiceMeshRequest) SetEnableAudit(v bool) *CreateServiceMeshRequest {
	s.EnableAudit = &v
	return s
}

func (s *CreateServiceMeshRequest) SetEnableCRHistory(v bool) *CreateServiceMeshRequest {
	s.EnableCRHistory = &v
	return s
}

func (s *CreateServiceMeshRequest) SetEnableSDSServer(v bool) *CreateServiceMeshRequest {
	s.EnableSDSServer = &v
	return s
}

func (s *CreateServiceMeshRequest) SetExcludeIPRanges(v string) *CreateServiceMeshRequest {
	s.ExcludeIPRanges = &v
	return s
}

func (s *CreateServiceMeshRequest) SetExcludeInboundPorts(v string) *CreateServiceMeshRequest {
	s.ExcludeInboundPorts = &v
	return s
}

func (s *CreateServiceMeshRequest) SetExcludeOutboundPorts(v string) *CreateServiceMeshRequest {
	s.ExcludeOutboundPorts = &v
	return s
}

func (s *CreateServiceMeshRequest) SetExistingCaCert(v string) *CreateServiceMeshRequest {
	s.ExistingCaCert = &v
	return s
}

func (s *CreateServiceMeshRequest) SetExistingCaKey(v string) *CreateServiceMeshRequest {
	s.ExistingCaKey = &v
	return s
}

func (s *CreateServiceMeshRequest) SetExistingCaType(v string) *CreateServiceMeshRequest {
	s.ExistingCaType = &v
	return s
}

func (s *CreateServiceMeshRequest) SetExistingRootCaCert(v string) *CreateServiceMeshRequest {
	s.ExistingRootCaCert = &v
	return s
}

func (s *CreateServiceMeshRequest) SetExistingRootCaKey(v string) *CreateServiceMeshRequest {
	s.ExistingRootCaKey = &v
	return s
}

func (s *CreateServiceMeshRequest) SetFilterGatewayClusterConfig(v bool) *CreateServiceMeshRequest {
	s.FilterGatewayClusterConfig = &v
	return s
}

func (s *CreateServiceMeshRequest) SetGatewayAPIEnabled(v bool) *CreateServiceMeshRequest {
	s.GatewayAPIEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetGuestCluster(v string) *CreateServiceMeshRequest {
	s.GuestCluster = &v
	return s
}

func (s *CreateServiceMeshRequest) SetIncludeIPRanges(v string) *CreateServiceMeshRequest {
	s.IncludeIPRanges = &v
	return s
}

func (s *CreateServiceMeshRequest) SetIstioVersion(v string) *CreateServiceMeshRequest {
	s.IstioVersion = &v
	return s
}

func (s *CreateServiceMeshRequest) SetKialiEnabled(v bool) *CreateServiceMeshRequest {
	s.KialiEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetLocalityLBConf(v string) *CreateServiceMeshRequest {
	s.LocalityLBConf = &v
	return s
}

func (s *CreateServiceMeshRequest) SetLocalityLoadBalancing(v bool) *CreateServiceMeshRequest {
	s.LocalityLoadBalancing = &v
	return s
}

func (s *CreateServiceMeshRequest) SetMSEEnabled(v bool) *CreateServiceMeshRequest {
	s.MSEEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetMultiBufferEnabled(v bool) *CreateServiceMeshRequest {
	s.MultiBufferEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetMultiBufferPollDelay(v string) *CreateServiceMeshRequest {
	s.MultiBufferPollDelay = &v
	return s
}

func (s *CreateServiceMeshRequest) SetMysqlFilterEnabled(v bool) *CreateServiceMeshRequest {
	s.MysqlFilterEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetName(v string) *CreateServiceMeshRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceMeshRequest) SetOPALimitCPU(v string) *CreateServiceMeshRequest {
	s.OPALimitCPU = &v
	return s
}

func (s *CreateServiceMeshRequest) SetOPALimitMemory(v string) *CreateServiceMeshRequest {
	s.OPALimitMemory = &v
	return s
}

func (s *CreateServiceMeshRequest) SetOPALogLevel(v string) *CreateServiceMeshRequest {
	s.OPALogLevel = &v
	return s
}

func (s *CreateServiceMeshRequest) SetOPARequestCPU(v string) *CreateServiceMeshRequest {
	s.OPARequestCPU = &v
	return s
}

func (s *CreateServiceMeshRequest) SetOPARequestMemory(v string) *CreateServiceMeshRequest {
	s.OPARequestMemory = &v
	return s
}

func (s *CreateServiceMeshRequest) SetOpaEnabled(v bool) *CreateServiceMeshRequest {
	s.OpaEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetOpenAgentPolicy(v bool) *CreateServiceMeshRequest {
	s.OpenAgentPolicy = &v
	return s
}

func (s *CreateServiceMeshRequest) SetPeriod(v int32) *CreateServiceMeshRequest {
	s.Period = &v
	return s
}

func (s *CreateServiceMeshRequest) SetPilotLoadBalancerSpec(v string) *CreateServiceMeshRequest {
	s.PilotLoadBalancerSpec = &v
	return s
}

func (s *CreateServiceMeshRequest) SetPrometheusUrl(v string) *CreateServiceMeshRequest {
	s.PrometheusUrl = &v
	return s
}

func (s *CreateServiceMeshRequest) SetProxyLimitCPU(v string) *CreateServiceMeshRequest {
	s.ProxyLimitCPU = &v
	return s
}

func (s *CreateServiceMeshRequest) SetProxyLimitMemory(v string) *CreateServiceMeshRequest {
	s.ProxyLimitMemory = &v
	return s
}

func (s *CreateServiceMeshRequest) SetProxyRequestCPU(v string) *CreateServiceMeshRequest {
	s.ProxyRequestCPU = &v
	return s
}

func (s *CreateServiceMeshRequest) SetProxyRequestMemory(v string) *CreateServiceMeshRequest {
	s.ProxyRequestMemory = &v
	return s
}

func (s *CreateServiceMeshRequest) SetRedisFilterEnabled(v bool) *CreateServiceMeshRequest {
	s.RedisFilterEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetRegionId(v string) *CreateServiceMeshRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceMeshRequest) SetTag(v []*CreateServiceMeshRequestTag) *CreateServiceMeshRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceMeshRequest) SetTelemetry(v bool) *CreateServiceMeshRequest {
	s.Telemetry = &v
	return s
}

func (s *CreateServiceMeshRequest) SetThriftFilterEnabled(v bool) *CreateServiceMeshRequest {
	s.ThriftFilterEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetTraceSampling(v float32) *CreateServiceMeshRequest {
	s.TraceSampling = &v
	return s
}

func (s *CreateServiceMeshRequest) SetTracing(v bool) *CreateServiceMeshRequest {
	s.Tracing = &v
	return s
}

func (s *CreateServiceMeshRequest) SetUseExistingCA(v bool) *CreateServiceMeshRequest {
	s.UseExistingCA = &v
	return s
}

func (s *CreateServiceMeshRequest) SetVSwitches(v string) *CreateServiceMeshRequest {
	s.VSwitches = &v
	return s
}

func (s *CreateServiceMeshRequest) SetVpcId(v string) *CreateServiceMeshRequest {
	s.VpcId = &v
	return s
}

func (s *CreateServiceMeshRequest) SetWebAssemblyFilterEnabled(v bool) *CreateServiceMeshRequest {
	s.WebAssemblyFilterEnabled = &v
	return s
}

type CreateServiceMeshRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceMeshRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMeshRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshRequestTag) SetKey(v string) *CreateServiceMeshRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceMeshRequestTag) SetValue(v string) *CreateServiceMeshRequestTag {
	s.Value = &v
	return s
}

type CreateServiceMeshResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s CreateServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshResponseBody) SetRequestId(v string) *CreateServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceMeshResponseBody) SetServiceMeshId(v string) *CreateServiceMeshResponseBody {
	s.ServiceMeshId = &v
	return s
}

type CreateServiceMeshResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshResponse) SetHeaders(v map[string]*string) *CreateServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceMeshResponse) SetStatusCode(v int32) *CreateServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceMeshResponse) SetBody(v *CreateServiceMeshResponseBody) *CreateServiceMeshResponse {
	s.Body = v
	return s
}

type CreateSwimLaneRequest struct {
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The label key of the associated service workload. Set the value to `ASM_TRAFFIC_TAG`.
	LabelSelectorKey *string `json:"LabelSelectorKey,omitempty" xml:"LabelSelectorKey,omitempty"`
	// The label value of the associated service workload.``
	LabelSelectorValue *string `json:"LabelSelectorValue,omitempty" xml:"LabelSelectorValue,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The list of services associated with the lane. The value is a JSON array. The format of a single service is `$Cluster name /$Cluster ID/$Namespace/$Service name`.
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
	// The name of the lane.
	SwimLaneName *string `json:"SwimLaneName,omitempty" xml:"SwimLaneName,omitempty"`
}

func (s CreateSwimLaneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSwimLaneRequest) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneRequest) SetGroupName(v string) *CreateSwimLaneRequest {
	s.GroupName = &v
	return s
}

func (s *CreateSwimLaneRequest) SetLabelSelectorKey(v string) *CreateSwimLaneRequest {
	s.LabelSelectorKey = &v
	return s
}

func (s *CreateSwimLaneRequest) SetLabelSelectorValue(v string) *CreateSwimLaneRequest {
	s.LabelSelectorValue = &v
	return s
}

func (s *CreateSwimLaneRequest) SetServiceMeshId(v string) *CreateSwimLaneRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateSwimLaneRequest) SetServicesList(v string) *CreateSwimLaneRequest {
	s.ServicesList = &v
	return s
}

func (s *CreateSwimLaneRequest) SetSwimLaneName(v string) *CreateSwimLaneRequest {
	s.SwimLaneName = &v
	return s
}

type CreateSwimLaneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSwimLaneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSwimLaneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneResponseBody) SetRequestId(v string) *CreateSwimLaneResponseBody {
	s.RequestId = &v
	return s
}

type CreateSwimLaneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSwimLaneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSwimLaneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSwimLaneResponse) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneResponse) SetHeaders(v map[string]*string) *CreateSwimLaneResponse {
	s.Headers = v
	return s
}

func (s *CreateSwimLaneResponse) SetStatusCode(v int32) *CreateSwimLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSwimLaneResponse) SetBody(v *CreateSwimLaneResponseBody) *CreateSwimLaneResponse {
	s.Body = v
	return s
}

type CreateSwimLaneGroupRequest struct {
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the ingress gateway.
	IngressGatewayName *string `json:"IngressGatewayName,omitempty" xml:"IngressGatewayName,omitempty"`
	// The type of the gateway for ingress traffic. Only ASM ingress gateways are supported.
	IngressType  *string `json:"IngressType,omitempty" xml:"IngressType,omitempty"`
	IsPermissive *bool   `json:"IsPermissive,omitempty" xml:"IsPermissive,omitempty"`
	RouteHeader  *string `json:"RouteHeader,omitempty" xml:"RouteHeader,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// A list of services associated with the lane group. The value is a JSON array. The format of a service is `$Cluster name/$Cluster ID/$Namespace/$Service name`.
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
	TraceHeader  *string `json:"TraceHeader,omitempty" xml:"TraceHeader,omitempty"`
}

func (s CreateSwimLaneGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSwimLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneGroupRequest) SetGroupName(v string) *CreateSwimLaneGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetIngressGatewayName(v string) *CreateSwimLaneGroupRequest {
	s.IngressGatewayName = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetIngressType(v string) *CreateSwimLaneGroupRequest {
	s.IngressType = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetIsPermissive(v bool) *CreateSwimLaneGroupRequest {
	s.IsPermissive = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetRouteHeader(v string) *CreateSwimLaneGroupRequest {
	s.RouteHeader = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetServiceMeshId(v string) *CreateSwimLaneGroupRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetServicesList(v string) *CreateSwimLaneGroupRequest {
	s.ServicesList = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetTraceHeader(v string) *CreateSwimLaneGroupRequest {
	s.TraceHeader = &v
	return s
}

type CreateSwimLaneGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSwimLaneGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSwimLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneGroupResponseBody) SetRequestId(v string) *CreateSwimLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateSwimLaneGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSwimLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSwimLaneGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSwimLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneGroupResponse) SetHeaders(v map[string]*string) *CreateSwimLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSwimLaneGroupResponse) SetStatusCode(v int32) *CreateSwimLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSwimLaneGroupResponse) SetBody(v *CreateSwimLaneGroupResponseBody) *CreateSwimLaneGroupResponse {
	s.Body = v
	return s
}

type DeleteGatewayRouteRequest struct {
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The name of the routing rule.
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteGatewayRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteRequest) SetIstioGatewayName(v string) *DeleteGatewayRouteRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DeleteGatewayRouteRequest) SetRouteName(v string) *DeleteGatewayRouteRequest {
	s.RouteName = &v
	return s
}

func (s *DeleteGatewayRouteRequest) SetServiceMeshId(v string) *DeleteGatewayRouteRequest {
	s.ServiceMeshId = &v
	return s
}

type DeleteGatewayRouteResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteResponseBody) SetRequestId(v string) *DeleteGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGatewayRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteResponse) SetHeaders(v map[string]*string) *DeleteGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayRouteResponse) SetStatusCode(v int32) *DeleteGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayRouteResponse) SetBody(v *DeleteGatewayRouteResponseBody) *DeleteGatewayRouteResponse {
	s.Body = v
	return s
}

type DeleteGatewaySecretRequest struct {
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteGatewaySecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecretRequest) SetIstioGatewayName(v string) *DeleteGatewaySecretRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DeleteGatewaySecretRequest) SetSecretName(v string) *DeleteGatewaySecretRequest {
	s.SecretName = &v
	return s
}

func (s *DeleteGatewaySecretRequest) SetServiceMeshId(v string) *DeleteGatewaySecretRequest {
	s.ServiceMeshId = &v
	return s
}

type DeleteGatewaySecretResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The record of deleting the secret.
	SecretDeleteRecord map[string]*SecretDeleteRecordValue `json:"SecretDeleteRecord,omitempty" xml:"SecretDeleteRecord,omitempty"`
}

func (s DeleteGatewaySecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecretResponseBody) SetRequestId(v string) *DeleteGatewaySecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewaySecretResponseBody) SetSecretDeleteRecord(v map[string]*SecretDeleteRecordValue) *DeleteGatewaySecretResponseBody {
	s.SecretDeleteRecord = v
	return s
}

type DeleteGatewaySecretResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewaySecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewaySecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewaySecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecretResponse) SetHeaders(v map[string]*string) *DeleteGatewaySecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewaySecretResponse) SetStatusCode(v int32) *DeleteGatewaySecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewaySecretResponse) SetBody(v *DeleteGatewaySecretResponseBody) *DeleteGatewaySecretResponse {
	s.Body = v
	return s
}

type DeleteIstioGatewayDomainsRequest struct {
	// The domain names of the one or more hosts that are exposed by the ASM gateway. Separate multiple domain names with commas (,).
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The maximum number of ASM gateways to query.
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The namespace in which the ASM gateway resides.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the port.
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteIstioGatewayDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIstioGatewayDomainsRequest) GoString() string {
	return s.String()
}

func (s *DeleteIstioGatewayDomainsRequest) SetHosts(v string) *DeleteIstioGatewayDomainsRequest {
	s.Hosts = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetIstioGatewayName(v string) *DeleteIstioGatewayDomainsRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetLimit(v string) *DeleteIstioGatewayDomainsRequest {
	s.Limit = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetNamespace(v string) *DeleteIstioGatewayDomainsRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetPortName(v string) *DeleteIstioGatewayDomainsRequest {
	s.PortName = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetServiceMeshId(v string) *DeleteIstioGatewayDomainsRequest {
	s.ServiceMeshId = &v
	return s
}

type DeleteIstioGatewayDomainsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIstioGatewayDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIstioGatewayDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIstioGatewayDomainsResponseBody) SetRequestId(v string) *DeleteIstioGatewayDomainsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIstioGatewayDomainsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIstioGatewayDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIstioGatewayDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIstioGatewayDomainsResponse) GoString() string {
	return s.String()
}

func (s *DeleteIstioGatewayDomainsResponse) SetHeaders(v map[string]*string) *DeleteIstioGatewayDomainsResponse {
	s.Headers = v
	return s
}

func (s *DeleteIstioGatewayDomainsResponse) SetStatusCode(v int32) *DeleteIstioGatewayDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIstioGatewayDomainsResponse) SetBody(v *DeleteIstioGatewayDomainsResponseBody) *DeleteIstioGatewayDomainsResponse {
	s.Body = v
	return s
}

type DeleteServiceMeshRequest struct {
	// Specifies whether to forcibly delete the ASM instance. Valid values:
	//
	// *   `true`: forcibly deletes the ASM instance.
	// *   `false`: does not forcibly delete the ASM instance.
	//
	// Default value: false.
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// A JSON string that can be parsed into a string array. You can use this JSON string to specify the IDs of the resource instances that need to be retained when the ASM instance is deleted.
	RetainResources *string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceMeshRequest) SetForce(v bool) *DeleteServiceMeshRequest {
	s.Force = &v
	return s
}

func (s *DeleteServiceMeshRequest) SetRetainResources(v string) *DeleteServiceMeshRequest {
	s.RetainResources = &v
	return s
}

func (s *DeleteServiceMeshRequest) SetServiceMeshId(v string) *DeleteServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

type DeleteServiceMeshResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceMeshResponseBody) SetRequestId(v string) *DeleteServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceMeshResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceMeshResponse) SetHeaders(v map[string]*string) *DeleteServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceMeshResponse) SetStatusCode(v int32) *DeleteServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceMeshResponse) SetBody(v *DeleteServiceMeshResponseBody) *DeleteServiceMeshResponse {
	s.Body = v
	return s
}

type DeleteSwimLaneRequest struct {
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The name of the lane.
	SwimLaneName *string `json:"SwimLaneName,omitempty" xml:"SwimLaneName,omitempty"`
}

func (s DeleteSwimLaneRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSwimLaneRequest) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneRequest) SetGroupName(v string) *DeleteSwimLaneRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteSwimLaneRequest) SetServiceMeshId(v string) *DeleteSwimLaneRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DeleteSwimLaneRequest) SetSwimLaneName(v string) *DeleteSwimLaneRequest {
	s.SwimLaneName = &v
	return s
}

type DeleteSwimLaneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSwimLaneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSwimLaneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneResponseBody) SetRequestId(v string) *DeleteSwimLaneResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSwimLaneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSwimLaneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSwimLaneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSwimLaneResponse) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneResponse) SetHeaders(v map[string]*string) *DeleteSwimLaneResponse {
	s.Headers = v
	return s
}

func (s *DeleteSwimLaneResponse) SetStatusCode(v int32) *DeleteSwimLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSwimLaneResponse) SetBody(v *DeleteSwimLaneResponseBody) *DeleteSwimLaneResponse {
	s.Body = v
	return s
}

type DeleteSwimLaneGroupRequest struct {
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteSwimLaneGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSwimLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneGroupRequest) SetGroupName(v string) *DeleteSwimLaneGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteSwimLaneGroupRequest) SetServiceMeshId(v string) *DeleteSwimLaneGroupRequest {
	s.ServiceMeshId = &v
	return s
}

type DeleteSwimLaneGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSwimLaneGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSwimLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneGroupResponseBody) SetRequestId(v string) *DeleteSwimLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSwimLaneGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSwimLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSwimLaneGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSwimLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneGroupResponse) SetHeaders(v map[string]*string) *DeleteSwimLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSwimLaneGroupResponse) SetStatusCode(v int32) *DeleteSwimLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSwimLaneGroupResponse) SetBody(v *DeleteSwimLaneGroupResponseBody) *DeleteSwimLaneGroupResponse {
	s.Body = v
	return s
}

type DescribeASMGatewayImportedServicesRequest struct {
	// The name of the ASM gateway.
	ASMGatewayName *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The namespace in which the service resides.
	ServiceNamespace *string `json:"ServiceNamespace,omitempty" xml:"ServiceNamespace,omitempty"`
}

func (s DescribeASMGatewayImportedServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeASMGatewayImportedServicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeASMGatewayImportedServicesRequest) SetASMGatewayName(v string) *DescribeASMGatewayImportedServicesRequest {
	s.ASMGatewayName = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesRequest) SetServiceMeshId(v string) *DescribeASMGatewayImportedServicesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesRequest) SetServiceNamespace(v string) *DescribeASMGatewayImportedServicesRequest {
	s.ServiceNamespace = &v
	return s
}

type DescribeASMGatewayImportedServicesResponseBody struct {
	// The imported services.
	ImportedServices []*DescribeASMGatewayImportedServicesResponseBodyImportedServices `json:"ImportedServices,omitempty" xml:"ImportedServices,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeASMGatewayImportedServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeASMGatewayImportedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeASMGatewayImportedServicesResponseBody) SetImportedServices(v []*DescribeASMGatewayImportedServicesResponseBodyImportedServices) *DescribeASMGatewayImportedServicesResponseBody {
	s.ImportedServices = v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponseBody) SetRequestId(v string) *DescribeASMGatewayImportedServicesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeASMGatewayImportedServicesResponseBodyImportedServices struct {
	// The name of the service.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The namespace in which the service resides.
	ServiceNamespace *string `json:"ServiceNamespace,omitempty" xml:"ServiceNamespace,omitempty"`
}

func (s DescribeASMGatewayImportedServicesResponseBodyImportedServices) String() string {
	return tea.Prettify(s)
}

func (s DescribeASMGatewayImportedServicesResponseBodyImportedServices) GoString() string {
	return s.String()
}

func (s *DescribeASMGatewayImportedServicesResponseBodyImportedServices) SetServiceName(v string) *DescribeASMGatewayImportedServicesResponseBodyImportedServices {
	s.ServiceName = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponseBodyImportedServices) SetServiceNamespace(v string) *DescribeASMGatewayImportedServicesResponseBodyImportedServices {
	s.ServiceNamespace = &v
	return s
}

type DescribeASMGatewayImportedServicesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeASMGatewayImportedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeASMGatewayImportedServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeASMGatewayImportedServicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeASMGatewayImportedServicesResponse) SetHeaders(v map[string]*string) *DescribeASMGatewayImportedServicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponse) SetStatusCode(v int32) *DescribeASMGatewayImportedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeASMGatewayImportedServicesResponse) SetBody(v *DescribeASMGatewayImportedServicesResponseBody) *DescribeASMGatewayImportedServicesResponse {
	s.Body = v
	return s
}

type DescribeCCMVersionRequest struct {
	// The versions of the CCM component in all clusters on the data plane.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeCCMVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCCMVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeCCMVersionRequest) SetServiceMeshId(v string) *DescribeCCMVersionRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeCCMVersionResponseBody struct {
	// The ID of the request.
	CCMVersions map[string]*CCMVersionsValue `json:"CCMVersions,omitempty" xml:"CCMVersions,omitempty"`
	RequestId   *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCCMVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCCMVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCCMVersionResponseBody) SetCCMVersions(v map[string]*CCMVersionsValue) *DescribeCCMVersionResponseBody {
	s.CCMVersions = v
	return s
}

func (s *DescribeCCMVersionResponseBody) SetRequestId(v string) *DescribeCCMVersionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCCMVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCCMVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCCMVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCCMVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeCCMVersionResponse) SetHeaders(v map[string]*string) *DescribeCCMVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeCCMVersionResponse) SetStatusCode(v int32) *DescribeCCMVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCCMVersionResponse) SetBody(v *DescribeCCMVersionResponseBody) *DescribeCCMVersionResponse {
	s.Body = v
	return s
}

type DescribeCensRequest struct {
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeCensRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) SetServiceMeshId(v string) *DescribeCensRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeCensResponseBody struct {
	// The IDs of the queried Kubernetes clusters.
	Clusters []*string `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBody) SetClusters(v []*string) *DescribeCensResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeCensResponseBody) SetRequestId(v string) *DescribeCensResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCensResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCensResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCensResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponse) GoString() string {
	return s.String()
}

func (s *DescribeCensResponse) SetHeaders(v map[string]*string) *DescribeCensResponse {
	s.Headers = v
	return s
}

func (s *DescribeCensResponse) SetStatusCode(v int32) *DescribeCensResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCensResponse) SetBody(v *DescribeCensResponseBody) *DescribeCensResponse {
	s.Body = v
	return s
}

type DescribeClusterGrafanaRequest struct {
	// The ID of the cluster on the data plane.
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeClusterGrafanaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterGrafanaRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterGrafanaRequest) SetK8sClusterId(v string) *DescribeClusterGrafanaRequest {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeClusterGrafanaRequest) SetServiceMeshId(v string) *DescribeClusterGrafanaRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeClusterGrafanaResponseBody struct {
	// The information of Grafana dashboards.
	Dashboards []*DescribeClusterGrafanaResponseBodyDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterGrafanaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterGrafanaResponseBody) SetDashboards(v []*DescribeClusterGrafanaResponseBodyDashboards) *DescribeClusterGrafanaResponseBody {
	s.Dashboards = v
	return s
}

func (s *DescribeClusterGrafanaResponseBody) SetRequestId(v string) *DescribeClusterGrafanaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterGrafanaResponseBodyDashboards struct {
	// The title of the Grafana dashboard.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The endpoint of the Grafana dashboard.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeClusterGrafanaResponseBodyDashboards) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterGrafanaResponseBodyDashboards) GoString() string {
	return s.String()
}

func (s *DescribeClusterGrafanaResponseBodyDashboards) SetTitle(v string) *DescribeClusterGrafanaResponseBodyDashboards {
	s.Title = &v
	return s
}

func (s *DescribeClusterGrafanaResponseBodyDashboards) SetUrl(v string) *DescribeClusterGrafanaResponseBodyDashboards {
	s.Url = &v
	return s
}

type DescribeClusterGrafanaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClusterGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterGrafanaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterGrafanaResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterGrafanaResponse) SetHeaders(v map[string]*string) *DescribeClusterGrafanaResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterGrafanaResponse) SetStatusCode(v int32) *DescribeClusterGrafanaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterGrafanaResponse) SetBody(v *DescribeClusterGrafanaResponseBody) *DescribeClusterGrafanaResponse {
	s.Body = v
	return s
}

type DescribeClusterPrometheusRequest struct {
	// The ID of the cluster on the data plane.
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The ID of the region where the cluster on the data plane resides.
	K8sClusterRegionId *string `json:"K8sClusterRegionId,omitempty" xml:"K8sClusterRegionId,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeClusterPrometheusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterPrometheusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterPrometheusRequest) SetK8sClusterId(v string) *DescribeClusterPrometheusRequest {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeClusterPrometheusRequest) SetK8sClusterRegionId(v string) *DescribeClusterPrometheusRequest {
	s.K8sClusterRegionId = &v
	return s
}

func (s *DescribeClusterPrometheusRequest) SetServiceMeshId(v string) *DescribeClusterPrometheusRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeClusterPrometheusResponseBody struct {
	// The public endpoint of the Prometheus service that is used to monitor a cluster in the ASM instance.
	Prometheus *string `json:"Prometheus,omitempty" xml:"Prometheus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterPrometheusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterPrometheusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterPrometheusResponseBody) SetPrometheus(v string) *DescribeClusterPrometheusResponseBody {
	s.Prometheus = &v
	return s
}

func (s *DescribeClusterPrometheusResponseBody) SetRequestId(v string) *DescribeClusterPrometheusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterPrometheusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClusterPrometheusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterPrometheusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterPrometheusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterPrometheusResponse) SetHeaders(v map[string]*string) *DescribeClusterPrometheusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterPrometheusResponse) SetStatusCode(v int32) *DescribeClusterPrometheusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterPrometheusResponse) SetBody(v *DescribeClusterPrometheusResponseBody) *DescribeClusterPrometheusResponse {
	s.Body = v
	return s
}

type DescribeClustersInServiceMeshRequest struct {
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeClustersInServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersInServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshRequest) SetServiceMeshId(v string) *DescribeClustersInServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeClustersInServiceMeshResponseBody struct {
	// The clusters in the ASM instance.
	Clusters []*DescribeClustersInServiceMeshResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClustersInServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponseBody) SetClusters(v []*DescribeClustersInServiceMeshResponseBodyClusters) *DescribeClustersInServiceMeshResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBody) SetRequestId(v string) *DescribeClustersInServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClustersInServiceMeshResponseBodyClusters struct {
	// The configurations of access log collection.
	AccessLogDashboards []*DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards `json:"AccessLogDashboards,omitempty" xml:"AccessLogDashboards,omitempty" type:"Repeated"`
	// The domain of the cluster.
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The ID of the cluster.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the cluster.
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The point in time when the cluster was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message that is returned when the call failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the Logtail component is installed in the cluster. Valid values:
	//
	// *   `logtail_installed`: The Logtail component is installed.
	//
	// \-`logtail_uninstalled`: The Logtail component is not installed.
	//
	// *   `logtail_state_get_error`: The Logtail component fails to be installed.
	LogtailInstalledState *string `json:"LogtailInstalledState,omitempty" xml:"LogtailInstalledState,omitempty"`
	// The name of the cluster.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region in which the cluster resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	SgId *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	// The status of the cluster.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The point in time when the cluster was last modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version number of the cluster.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClustersInServiceMeshResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetAccessLogDashboards(v []*DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.AccessLogDashboards = v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetClusterDomain(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetClusterId(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetClusterType(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetCreationTime(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.CreationTime = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetErrorMessage(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetLogtailInstalledState(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.LogtailInstalledState = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetName(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetRegionId(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetSgId(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.SgId = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetState(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetUpdateTime(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.UpdateTime = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetVersion(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.Version = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetVpcId(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.VpcId = &v
	return s
}

type DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards struct {
	// The name of the dashboard for access logs.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the dashboard for access logs.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) SetTitle(v string) *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards {
	s.Title = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) SetUrl(v string) *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards {
	s.Url = &v
	return s
}

type DescribeClustersInServiceMeshResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClustersInServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClustersInServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponse) SetHeaders(v map[string]*string) *DescribeClustersInServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersInServiceMeshResponse) SetStatusCode(v int32) *DescribeClustersInServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponse) SetBody(v *DescribeClustersInServiceMeshResponseBody) *DescribeClustersInServiceMeshResponse {
	s.Body = v
	return s
}

type DescribeCrTemplatesRequest struct {
	// The Istio version used in ASM.
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	// The type of Istio resource whose common YAML templates you want to query. Valid values:
	//
	// *   AuthorizationPolicy
	// *   RequestAuthentication
	// *   PeerAuthentication
	// *   WorkloadGroup
	// *   WorkloadEntry
	// *   Sidecar
	// *   EnvoyFilter
	// *   ServiceEntry
	// *   Gateway
	// *   DestinationRule
	// *   VirtualService
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
}

func (s DescribeCrTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrTemplatesRequest) SetIstioVersion(v string) *DescribeCrTemplatesRequest {
	s.IstioVersion = &v
	return s
}

func (s *DescribeCrTemplatesRequest) SetKind(v string) *DescribeCrTemplatesRequest {
	s.Kind = &v
	return s
}

type DescribeCrTemplatesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The common YAML templates for the specified type of Istio resource.
	Templates []*DescribeCrTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DescribeCrTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrTemplatesResponseBody) SetRequestId(v string) *DescribeCrTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrTemplatesResponseBody) SetTemplates(v []*DescribeCrTemplatesResponseBodyTemplates) *DescribeCrTemplatesResponseBody {
	s.Templates = v
	return s
}

type DescribeCrTemplatesResponseBodyTemplates struct {
	// The Chinese name of the YAML template.
	ChineseName *string `json:"ChineseName,omitempty" xml:"ChineseName,omitempty"`
	// The English name of the YAML template.
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	// The content in the YAML template.
	Yaml *string `json:"Yaml,omitempty" xml:"Yaml,omitempty"`
}

func (s DescribeCrTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeCrTemplatesResponseBodyTemplates) SetChineseName(v string) *DescribeCrTemplatesResponseBodyTemplates {
	s.ChineseName = &v
	return s
}

func (s *DescribeCrTemplatesResponseBodyTemplates) SetEnglishName(v string) *DescribeCrTemplatesResponseBodyTemplates {
	s.EnglishName = &v
	return s
}

func (s *DescribeCrTemplatesResponseBodyTemplates) SetYaml(v string) *DescribeCrTemplatesResponseBodyTemplates {
	s.Yaml = &v
	return s
}

type DescribeCrTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCrTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCrTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrTemplatesResponse) SetHeaders(v map[string]*string) *DescribeCrTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrTemplatesResponse) SetStatusCode(v int32) *DescribeCrTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrTemplatesResponse) SetBody(v *DescribeCrTemplatesResponseBody) *DescribeCrTemplatesResponse {
	s.Body = v
	return s
}

type DescribeEipResourcesRequest struct {
	// The number of the page to return. Default value: 1.
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeEipResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesRequest) SetPageNum(v int32) *DescribeEipResourcesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeEipResourcesRequest) SetPageSize(v int32) *DescribeEipResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEipResourcesRequest) SetServiceMeshId(v string) *DescribeEipResourcesRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeEipResourcesResponseBody struct {
	// The EIPs.
	EipList []*DescribeEipResourcesResponseBodyEipList `json:"EipList,omitempty" xml:"EipList,omitempty" type:"Repeated"`
	// The pagination information.
	PageResult *DescribeEipResourcesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEipResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesResponseBody) SetEipList(v []*DescribeEipResourcesResponseBodyEipList) *DescribeEipResourcesResponseBody {
	s.EipList = v
	return s
}

func (s *DescribeEipResourcesResponseBody) SetPageResult(v *DescribeEipResourcesResponseBodyPageResult) *DescribeEipResourcesResponseBody {
	s.PageResult = v
	return s
}

func (s *DescribeEipResourcesResponseBody) SetRequestId(v string) *DescribeEipResourcesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEipResourcesResponseBodyEipList struct {
	// The ID of the EIP.
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The type of the resource that is associated with the EIP. Valid values:
	//
	// *   `EcsInstance`: an ECS instance in a VPC
	// *   `SlbInstance`: a Server Load Balancer (SLB) instance in a VPC
	// *   `Nat`: a NAT gateway
	// *   `HaVip`: a high-availability virtual IP address (HAVIP)
	// *   `NetworkInterface`: a secondary elastic network interface (ENI)
	//
	// Default value: `EcsInstance`.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the EIP.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The status of the EIP. Valid values:
	//
	// *   `Associating`: The EIP is being associated with a resource.
	// *   `Unassociating`: The EIP is being disassociated from a resource.
	// *   `InUse`: The EIP is associated with a resource.
	// *   `Available`: The EIP is available.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEipResourcesResponseBodyEipList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipResourcesResponseBodyEipList) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesResponseBodyEipList) SetAllocationId(v string) *DescribeEipResourcesResponseBodyEipList {
	s.AllocationId = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyEipList) SetInstanceType(v string) *DescribeEipResourcesResponseBodyEipList {
	s.InstanceType = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyEipList) SetIpAddress(v string) *DescribeEipResourcesResponseBodyEipList {
	s.IpAddress = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyEipList) SetStatus(v string) *DescribeEipResourcesResponseBodyEipList {
	s.Status = &v
	return s
}

type DescribeEipResourcesResponseBodyPageResult struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEipResourcesResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipResourcesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesResponseBodyPageResult) SetPageNumber(v int32) *DescribeEipResourcesResponseBodyPageResult {
	s.PageNumber = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyPageResult) SetPageSize(v int32) *DescribeEipResourcesResponseBodyPageResult {
	s.PageSize = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyPageResult) SetTotalCount(v int32) *DescribeEipResourcesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

type DescribeEipResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEipResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEipResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEipResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesResponse) SetHeaders(v map[string]*string) *DescribeEipResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEipResourcesResponse) SetStatusCode(v int32) *DescribeEipResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEipResourcesResponse) SetBody(v *DescribeEipResourcesResponseBody) *DescribeEipResourcesResponse {
	s.Body = v
	return s
}

type DescribeGatewaySecretDetailsRequest struct {
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeGatewaySecretDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySecretDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySecretDetailsRequest) SetIstioGatewayName(v string) *DescribeGatewaySecretDetailsRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DescribeGatewaySecretDetailsRequest) SetServiceMeshId(v string) *DescribeGatewaySecretDetailsRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeGatewaySecretDetailsResponseBody struct {
	// The detailed information about the secret of the ASM gateway.
	GatewaySecretDetails []*DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails `json:"GatewaySecretDetails,omitempty" xml:"GatewaySecretDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGatewaySecretDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySecretDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySecretDetailsResponseBody) SetGatewaySecretDetails(v []*DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) *DescribeGatewaySecretDetailsResponseBody {
	s.GatewaySecretDetails = v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBody) SetRequestId(v string) *DescribeGatewaySecretDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails struct {
	// The time when the certificate expires.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The name of the gateway.
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The time when the certificate was issued.
	IssueTime *string `json:"IssueTime,omitempty" xml:"IssueTime,omitempty"`
	// *   An error message is returned if the status of the gateway is abnormal. Examples: `tls.crt not exist`, `tls.key not exist`, and `secret type must be kubernetes.io/tls`.
	// *   An empty value is returned if the status of the gateway is normal.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The Server Name Indication (SNI) value that indicates the hostname of the service.
	SNI *string `json:"SNI,omitempty" xml:"SNI,omitempty"`
	// The name of the secret.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The status of the certificate. Valid values:
	//
	// *   `normal`
	// *   `abnormal`
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetExpiredTime(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetGatewayName(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.GatewayName = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetIssueTime(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.IssueTime = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetMessage(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.Message = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetSNI(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.SNI = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetSecretName(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.SecretName = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetState(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.State = &v
	return s
}

type DescribeGatewaySecretDetailsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGatewaySecretDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGatewaySecretDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewaySecretDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySecretDetailsResponse) SetHeaders(v map[string]*string) *DescribeGatewaySecretDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewaySecretDetailsResponse) SetStatusCode(v int32) *DescribeGatewaySecretDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponse) SetBody(v *DescribeGatewaySecretDetailsResponseBody) *DescribeGatewaySecretDetailsResponse {
	s.Body = v
	return s
}

type DescribeGuestClusterAccessLogDashboardsRequest struct {
	// The ID of the cluster on the data plane.
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
}

func (s DescribeGuestClusterAccessLogDashboardsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterAccessLogDashboardsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterAccessLogDashboardsRequest) SetK8sClusterId(v string) *DescribeGuestClusterAccessLogDashboardsRequest {
	s.K8sClusterId = &v
	return s
}

type DescribeGuestClusterAccessLogDashboardsResponseBody struct {
	// The access log reports of a cluster on the data plane.
	Dashboards []*DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The ID of the cluster on the data plane.
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGuestClusterAccessLogDashboardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterAccessLogDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) SetDashboards(v []*DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) *DescribeGuestClusterAccessLogDashboardsResponseBody {
	s.Dashboards = v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) SetK8sClusterId(v string) *DescribeGuestClusterAccessLogDashboardsResponseBody {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) SetRequestId(v string) *DescribeGuestClusterAccessLogDashboardsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards struct {
	// The title of the report.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the report.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) SetTitle(v string) *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards {
	s.Title = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) SetUrl(v string) *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards {
	s.Url = &v
	return s
}

type DescribeGuestClusterAccessLogDashboardsResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGuestClusterAccessLogDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGuestClusterAccessLogDashboardsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterAccessLogDashboardsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) SetHeaders(v map[string]*string) *DescribeGuestClusterAccessLogDashboardsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) SetStatusCode(v int32) *DescribeGuestClusterAccessLogDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) SetBody(v *DescribeGuestClusterAccessLogDashboardsResponseBody) *DescribeGuestClusterAccessLogDashboardsResponse {
	s.Body = v
	return s
}

type DescribeGuestClusterNamespacesRequest struct {
	// The ID of the Kubernetes cluster that is added to the ASM instance.
	GuestClusterID *string `json:"GuestClusterID,omitempty" xml:"GuestClusterID,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// Specifies whether to return the labels of the namespaces.
	ShowNsLabels *bool `json:"ShowNsLabels,omitempty" xml:"ShowNsLabels,omitempty"`
}

func (s DescribeGuestClusterNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterNamespacesRequest) SetGuestClusterID(v string) *DescribeGuestClusterNamespacesRequest {
	s.GuestClusterID = &v
	return s
}

func (s *DescribeGuestClusterNamespacesRequest) SetServiceMeshId(v string) *DescribeGuestClusterNamespacesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeGuestClusterNamespacesRequest) SetShowNsLabels(v bool) *DescribeGuestClusterNamespacesRequest {
	s.ShowNsLabels = &v
	return s
}

type DescribeGuestClusterNamespacesResponseBody struct {
	// The labels of the namespaces. Labels are returned only when the `ShowNsLabels` parameter is set to `true`.
	NsLabels map[string]interface{} `json:"NsLabels,omitempty" xml:"NsLabels,omitempty"`
	// The names of the namespaces.
	NsList []*string `json:"NsList,omitempty" xml:"NsList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGuestClusterNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterNamespacesResponseBody) SetNsLabels(v map[string]interface{}) *DescribeGuestClusterNamespacesResponseBody {
	s.NsLabels = v
	return s
}

func (s *DescribeGuestClusterNamespacesResponseBody) SetNsList(v []*string) *DescribeGuestClusterNamespacesResponseBody {
	s.NsList = v
	return s
}

func (s *DescribeGuestClusterNamespacesResponseBody) SetRequestId(v string) *DescribeGuestClusterNamespacesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGuestClusterNamespacesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGuestClusterNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGuestClusterNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterNamespacesResponse) SetHeaders(v map[string]*string) *DescribeGuestClusterNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGuestClusterNamespacesResponse) SetStatusCode(v int32) *DescribeGuestClusterNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGuestClusterNamespacesResponse) SetBody(v *DescribeGuestClusterNamespacesResponseBody) *DescribeGuestClusterNamespacesResponse {
	s.Body = v
	return s
}

type DescribeGuestClusterPodsRequest struct {
	// The ID of the Kubernetes cluster that is added to the ASM instance.
	GuestClusterID *string `json:"GuestClusterID,omitempty" xml:"GuestClusterID,omitempty"`
	// The name of the namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeGuestClusterPodsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterPodsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterPodsRequest) SetGuestClusterID(v string) *DescribeGuestClusterPodsRequest {
	s.GuestClusterID = &v
	return s
}

func (s *DescribeGuestClusterPodsRequest) SetNamespace(v string) *DescribeGuestClusterPodsRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeGuestClusterPodsRequest) SetServiceMeshId(v string) *DescribeGuestClusterPodsRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeGuestClusterPodsResponseBody struct {
	// The names of the queried pods.
	PodList []*string `json:"PodList,omitempty" xml:"PodList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGuestClusterPodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterPodsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterPodsResponseBody) SetPodList(v []*string) *DescribeGuestClusterPodsResponseBody {
	s.PodList = v
	return s
}

func (s *DescribeGuestClusterPodsResponseBody) SetRequestId(v string) *DescribeGuestClusterPodsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGuestClusterPodsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGuestClusterPodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGuestClusterPodsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterPodsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterPodsResponse) SetHeaders(v map[string]*string) *DescribeGuestClusterPodsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGuestClusterPodsResponse) SetStatusCode(v int32) *DescribeGuestClusterPodsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGuestClusterPodsResponse) SetBody(v *DescribeGuestClusterPodsResponseBody) *DescribeGuestClusterPodsResponse {
	s.Body = v
	return s
}

type DescribeImportedServicesDetailRequest struct {
	// The name of the service.
	ASMGatewayName *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	// The details of the services.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The namespace in which the service resides.
	ServiceNamespace *string `json:"ServiceNamespace,omitempty" xml:"ServiceNamespace,omitempty"`
}

func (s DescribeImportedServicesDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedServicesDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailRequest) SetASMGatewayName(v string) *DescribeImportedServicesDetailRequest {
	s.ASMGatewayName = &v
	return s
}

func (s *DescribeImportedServicesDetailRequest) SetServiceMeshId(v string) *DescribeImportedServicesDetailRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeImportedServicesDetailRequest) SetServiceNamespace(v string) *DescribeImportedServicesDetailRequest {
	s.ServiceNamespace = &v
	return s
}

type DescribeImportedServicesDetailResponseBody struct {
	// The IDs of the clusters to which the service belongs.
	Details []*DescribeImportedServicesDetailResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The labels of the service.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImportedServicesDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedServicesDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailResponseBody) SetDetails(v []*DescribeImportedServicesDetailResponseBodyDetails) *DescribeImportedServicesDetailResponseBody {
	s.Details = v
	return s
}

func (s *DescribeImportedServicesDetailResponseBody) SetRequestId(v string) *DescribeImportedServicesDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImportedServicesDetailResponseBodyDetails struct {
	// The name of the port.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	// The ports declared for the service.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Kubernetes
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The port number.
	Ports []*DescribeImportedServicesDetailResponseBodyDetailsPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The type of the service.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The protocol of the port.
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s DescribeImportedServicesDetailResponseBodyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedServicesDetailResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetClusterIds(v []*string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.ClusterIds = v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetLabels(v map[string]*string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.Labels = v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetNamespace(v string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.Namespace = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetPorts(v []*DescribeImportedServicesDetailResponseBodyDetailsPorts) *DescribeImportedServicesDetailResponseBodyDetails {
	s.Ports = v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetServiceName(v string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.ServiceName = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetails) SetServiceType(v string) *DescribeImportedServicesDetailResponseBodyDetails {
	s.ServiceType = &v
	return s
}

type DescribeImportedServicesDetailResponseBodyDetailsPorts struct {
	// The container port.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodePort *int32  `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	// The node port.
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TargetPort *int32  `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s DescribeImportedServicesDetailResponseBodyDetailsPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedServicesDetailResponseBodyDetailsPorts) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetName(v string) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.Name = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetNodePort(v int32) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.NodePort = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetPort(v int32) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.Port = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetProtocol(v string) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeImportedServicesDetailResponseBodyDetailsPorts) SetTargetPort(v int32) *DescribeImportedServicesDetailResponseBodyDetailsPorts {
	s.TargetPort = &v
	return s
}

type DescribeImportedServicesDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeImportedServicesDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImportedServicesDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedServicesDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailResponse) SetHeaders(v map[string]*string) *DescribeImportedServicesDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeImportedServicesDetailResponse) SetStatusCode(v int32) *DescribeImportedServicesDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImportedServicesDetailResponse) SetBody(v *DescribeImportedServicesDetailResponseBody) *DescribeImportedServicesDetailResponse {
	s.Body = v
	return s
}

type DescribeIstioGatewayDomainsRequest struct {
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The maximum number of ASM gateways to query.
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The namespace in which the ASM gateway resides.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeIstioGatewayDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayDomainsRequest) SetIstioGatewayName(v string) *DescribeIstioGatewayDomainsRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DescribeIstioGatewayDomainsRequest) SetLimit(v string) *DescribeIstioGatewayDomainsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeIstioGatewayDomainsRequest) SetNamespace(v string) *DescribeIstioGatewayDomainsRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayDomainsRequest) SetServiceMeshId(v string) *DescribeIstioGatewayDomainsRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeIstioGatewayDomainsResponseBody struct {
	// The domain names that are exposed by the ASM gateway.
	GatewaySecretDetails []*DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails `json:"GatewaySecretDetails,omitempty" xml:"GatewaySecretDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIstioGatewayDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayDomainsResponseBody) SetGatewaySecretDetails(v []*DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) *DescribeIstioGatewayDomainsResponseBody {
	s.GatewaySecretDetails = v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBody) SetRequestId(v string) *DescribeIstioGatewayDomainsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails struct {
	// The name of the secret that contains the Transport Layer Security (TLS) certificate and certificate authority (CA) certificate.
	CredentialName *string `json:"CredentialName,omitempty" xml:"CredentialName,omitempty"`
	// The details of the domain name in the JSON format.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The domain name.
	Domains       []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	GatewayCRName *string   `json:"GatewayCRName,omitempty" xml:"GatewayCRName,omitempty"`
	// The namespace in which the ASM gateway resides.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the port.
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The type of the protocol. Valid values: `HTTP`, `HTTPS`, `GRPC`, `HTTP2`, `MONGO`, `TCP`, and `TLS`.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetCredentialName(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.CredentialName = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetDetail(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.Detail = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetDomains(v []*string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.Domains = v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetGatewayCRName(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.GatewayCRName = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetNamespace(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetPortName(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.PortName = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetProtocol(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.Protocol = &v
	return s
}

type DescribeIstioGatewayDomainsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIstioGatewayDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIstioGatewayDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayDomainsResponse) SetHeaders(v map[string]*string) *DescribeIstioGatewayDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayDomainsResponse) SetStatusCode(v int32) *DescribeIstioGatewayDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponse) SetBody(v *DescribeIstioGatewayDomainsResponseBody) *DescribeIstioGatewayDomainsResponse {
	s.Body = v
	return s
}

type DescribeIstioGatewayRouteDetailRequest struct {
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The name of the routing rule.
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailRequest) SetIstioGatewayName(v string) *DescribeIstioGatewayRouteDetailRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailRequest) SetRouteName(v string) *DescribeIstioGatewayRouteDetailRequest {
	s.RouteName = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailRequest) SetServiceMeshId(v string) *DescribeIstioGatewayRouteDetailRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBody struct {
	// The description of the routing rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about the routing rule.
	RouteDetail *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail `json:"RouteDetail,omitempty" xml:"RouteDetail,omitempty" type:"Struct"`
	// The status of the routing rule. Valid values:
	//
	// *   `0`: The routing rule is valid.
	// *   `1`: The routing rule is invalid.
	// *   `2`: An error occurs during the creation or update of the routing rule.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetDescription(v string) *DescribeIstioGatewayRouteDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetNamespace(v string) *DescribeIstioGatewayRouteDetailResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetPriority(v int32) *DescribeIstioGatewayRouteDetailResponseBody {
	s.Priority = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetRequestId(v string) *DescribeIstioGatewayRouteDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetRouteDetail(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) *DescribeIstioGatewayRouteDetailResponseBody {
	s.RouteDetail = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBody) SetStatus(v int32) *DescribeIstioGatewayRouteDetailResponseBody {
	s.Status = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetail struct {
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The advanced settings for routing HTTP traffic.
	HTTPAdvancedOptions *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions `json:"HTTPAdvancedOptions,omitempty" xml:"HTTPAdvancedOptions,omitempty" type:"Struct"`
	HasUnsafeFeatures   *bool                                                                      `json:"HasUnsafeFeatures,omitempty" xml:"HasUnsafeFeatures,omitempty"`
	// The matching rules for traffic routing.
	MatchRequest *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest `json:"MatchRequest,omitempty" xml:"MatchRequest,omitempty" type:"Struct"`
	RawVSRoute   *string                                                             `json:"RawVSRoute,omitempty" xml:"RawVSRoute,omitempty"`
	// The endpoints of destination services for Layer 4 weighted routing.
	RouteDestinations []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations `json:"RouteDestinations,omitempty" xml:"RouteDestinations,omitempty" type:"Repeated"`
	// The name of the routing rule.
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The type of the traffic to be routed. Valid values: `HTTP`, `TLS`, and `TCP`.
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetDomains(v []*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.Domains = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetHTTPAdvancedOptions(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.HTTPAdvancedOptions = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetHasUnsafeFeatures(v bool) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.HasUnsafeFeatures = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetMatchRequest(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.MatchRequest = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetRawVSRoute(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.RawVSRoute = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetRouteDestinations(v []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.RouteDestinations = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetRouteName(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.RouteName = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail) SetRouteType(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetail {
	s.RouteType = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions struct {
	// The virtual service that defines traffic routing.
	Delegate *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate `json:"Delegate,omitempty" xml:"Delegate,omitempty" type:"Struct"`
	// The configurations of fault injection.
	Fault *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault `json:"Fault,omitempty" xml:"Fault,omitempty" type:"Struct"`
	// The HTTP redirection rule.
	HTTPRedirect *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect `json:"HTTPRedirect,omitempty" xml:"HTTPRedirect,omitempty" type:"Struct"`
	// The configurations for mirroring HTTP traffic to another destination in addition to forwarding requests to the specified destination.
	Mirror *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror `json:"Mirror,omitempty" xml:"Mirror,omitempty" type:"Struct"`
	// The percentage of requests that are aborted with the specified error code.
	MirrorPercentage *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage `json:"MirrorPercentage,omitempty" xml:"MirrorPercentage,omitempty" type:"Struct"`
	// The configurations of retries for failed requests.
	Retries *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries `json:"Retries,omitempty" xml:"Retries,omitempty" type:"Struct"`
	// The configurations for rewriting the virtual service.
	Rewrite *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Struct"`
	// The timeout period for requests.
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetDelegate(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Delegate = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetFault(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Fault = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetHTTPRedirect(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.HTTPRedirect = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetMirror(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Mirror = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetMirrorPercentage(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.MirrorPercentage = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetRetries(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Retries = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetRewrite(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Rewrite = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions) SetTimeout(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptions {
	s.Timeout = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate struct {
	// The name of the virtual service.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the virtual service belongs.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) SetName(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate {
	s.Name = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate) SetNamespace(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsDelegate {
	s.Namespace = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault struct {
	// The configurations for aborting requests with specified error codes.
	Abort *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort `json:"Abort,omitempty" xml:"Abort,omitempty" type:"Struct"`
	// The duration to delay a request.
	Delay *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay `json:"Delay,omitempty" xml:"Delay,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) SetAbort(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault {
	s.Abort = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault) SetDelay(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFault {
	s.Delay = v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort struct {
	// The HTTP status code.
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The percentage of requests that are aborted with the specified error code.
	Percentage *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) SetHttpStatus(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort {
	s.HttpStatus = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort) SetPercentage(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbort {
	s.Percentage = v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage struct {
	// The percentage of requests that are mirrored to another destination except for the original destination, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage) SetValue(v float32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultAbortPercentage {
	s.Value = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay struct {
	// The duration for request delay is expressed as 2 raised to the power of x. You must specify the value of x.
	ExponentialDelay *string `json:"ExponentialDelay,omitempty" xml:"ExponentialDelay,omitempty"`
	// The fixed duration for request delay.
	FixedDelay *string `json:"FixedDelay,omitempty" xml:"FixedDelay,omitempty"`
	// The percentage of requests to which the delay fault is injected.
	Percentage *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) SetExponentialDelay(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay {
	s.ExponentialDelay = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) SetFixedDelay(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay {
	s.FixedDelay = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay) SetPercentage(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelay {
	s.Percentage = v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage struct {
	// The percentage of requests that are aborted with the specified error code, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage) SetValue(v float32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsFaultDelayPercentage {
	s.Value = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect struct {
	// The value to be used to overwrite the value of the Authority or Host header during redirection.
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The HTTP status code to be used to indicate URL redirection. Default value: 301.
	RedirectCode *int32 `json:"RedirectCode,omitempty" xml:"RedirectCode,omitempty"`
	// The value to be used to overwrite the URL path during redirection.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) SetAuthority(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect {
	s.Authority = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) SetRedirectCode(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect {
	s.RedirectCode = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect) SetUri(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsHTTPRedirect {
	s.Uri = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror struct {
	// The name of the service defined in the service registry.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The name of the service subset.
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) SetHost(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror {
	s.Host = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror) SetSubset(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirror {
	s.Subset = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage struct {
	// The percentage of requests that are aborted with the specified error code, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage) SetValue(v float32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsMirrorPercentage {
	s.Value = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries struct {
	// The number of retries that are allowed for a request.
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// The timeout period for each retry.
	PerTryTimeout *string `json:"PerTryTimeout,omitempty" xml:"PerTryTimeout,omitempty"`
	// The condition for retries. Example: `connect-failure,refused-stream,503`.
	RetryOn *string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty"`
	// Specifies whether to allow retries to other localities.
	RetryRemoteLocalities *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities `json:"RetryRemoteLocalities,omitempty" xml:"RetryRemoteLocalities,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) SetAttempts(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	s.Attempts = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) SetPerTryTimeout(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	s.PerTryTimeout = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) SetRetryOn(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	s.RetryOn = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries) SetRetryRemoteLocalities(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetries {
	s.RetryRemoteLocalities = v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities struct {
	// Specifies whether to allow retries to other localities. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	Value *bool `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities) SetValue(v bool) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	s.Value = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite struct {
	// The value to be used to overwrite the value of the Authority or Host header.
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The value to be used to overwrite the path or prefix of the URI.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) SetAuthority(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite {
	s.Authority = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite) SetUri(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailHTTPAdvancedOptionsRewrite {
	s.Uri = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest struct {
	// The request headers to be matched.
	Headers []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The ports.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The matching rules for Transport Layer Security (TLS) traffic.
	TLSMatchAttributes []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes `json:"TLSMatchAttributes,omitempty" xml:"TLSMatchAttributes,omitempty" type:"Repeated"`
	// The matching rule for URIs.
	URI *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI `json:"URI,omitempty" xml:"URI,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) SetHeaders(v []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) SetPorts(v []*int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	s.Ports = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) SetTLSMatchAttributes(v []*DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	s.TLSMatchAttributes = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest) SetURI(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequest {
	s.URI = v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders struct {
	// The header value to be matched.
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode for the header value. Valid values:
	//
	// *   `exact`: exact match
	// *   `prefix`: match by prefix
	// *   `regex`: match by regular expression
	MatchingMode *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
	// The header key to be matched.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) SetMatchingContent(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders {
	s.MatchingContent = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) SetMatchingMode(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders {
	s.MatchingMode = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders) SetName(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestHeaders {
	s.Name = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes struct {
	// The Server Name Indication (SNI) values to be matched.
	SNIHosts []*string `json:"SNIHosts,omitempty" xml:"SNIHosts,omitempty" type:"Repeated"`
	// The TLS port.
	TLSPort *int32 `json:"TLSPort,omitempty" xml:"TLSPort,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) SetSNIHosts(v []*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes {
	s.SNIHosts = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes) SetTLSPort(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestTLSMatchAttributes {
	s.TLSPort = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI struct {
	// The content to be matched.
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode. Valid values:
	//
	// *   `exact`: exact match
	// *   `prefix`: match by prefix
	// *   `regex`: match by regular expression
	MatchingMode *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) SetMatchingContent(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI {
	s.MatchingContent = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI) SetMatchingMode(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailMatchRequestURI {
	s.MatchingMode = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations struct {
	// The unique endpoint of the destination service to which the specified requests are sent.
	Destination *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Struct"`
	// The request headers to be matched.
	Headers *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The traffic weight. Valid values: 1 to 100.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) SetDestination(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations {
	s.Destination = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) SetHeaders(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations) SetWeight(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinations {
	s.Weight = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination struct {
	// The name of the service defined in the service registry.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ports.
	Port *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Struct"`
	// The name of the service subset.
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) SetHost(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination {
	s.Host = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) SetPort(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination {
	s.Port = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination) SetSubset(v string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestination {
	s.Subset = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort struct {
	// The ports of the specified hosts to which the traffic is routed.
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort) SetNumber(v int32) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsDestinationPort {
	s.Number = &v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders struct {
	// The request header to be matched.
	Request *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// The processing of the headers of the response that is to be returned.
	Response *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) SetRequest(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders {
	s.Request = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders) SetResponse(v *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeaders {
	s.Response = v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest struct {
	// The values to be added to the header key.
	Add map[string]interface{} `json:"Add,omitempty" xml:"Add,omitempty"`
	// The header value to be deleted.
	Remove []*string `json:"Remove,omitempty" xml:"Remove,omitempty" type:"Repeated"`
	// The header key to be used to overwrite the original header key.
	Set map[string]*string `json:"Set,omitempty" xml:"Set,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) SetAdd(v map[string]interface{}) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest {
	s.Add = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) SetRemove(v []*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest {
	s.Remove = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest) SetSet(v map[string]*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersRequest {
	s.Set = v
	return s
}

type DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse struct {
	// The values to be added to the header key.
	Add map[string]interface{} `json:"Add,omitempty" xml:"Add,omitempty"`
	// The header value to be deleted.
	Remove []*string `json:"Remove,omitempty" xml:"Remove,omitempty" type:"Repeated"`
	// The header key to be used to overwrite the original header key.
	Set map[string]interface{} `json:"Set,omitempty" xml:"Set,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) SetAdd(v map[string]interface{}) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse {
	s.Add = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) SetRemove(v []*string) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse {
	s.Remove = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse) SetSet(v map[string]interface{}) *DescribeIstioGatewayRouteDetailResponseBodyRouteDetailRouteDestinationsHeadersResponse {
	s.Set = v
	return s
}

type DescribeIstioGatewayRouteDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIstioGatewayRouteDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIstioGatewayRouteDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponse) SetHeaders(v map[string]*string) *DescribeIstioGatewayRouteDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponse) SetStatusCode(v int32) *DescribeIstioGatewayRouteDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponse) SetBody(v *DescribeIstioGatewayRouteDetailResponseBody) *DescribeIstioGatewayRouteDetailResponse {
	s.Body = v
	return s
}

type DescribeIstioGatewayRoutesRequest struct {
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeIstioGatewayRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRoutesRequest) SetIstioGatewayName(v string) *DescribeIstioGatewayRoutesRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DescribeIstioGatewayRoutesRequest) SetServiceMeshId(v string) *DescribeIstioGatewayRoutesRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeIstioGatewayRoutesResponseBody struct {
	// The routing rules.
	ManagementRoutes []*DescribeIstioGatewayRoutesResponseBodyManagementRoutes `json:"ManagementRoutes,omitempty" xml:"ManagementRoutes,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIstioGatewayRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRoutesResponseBody) SetManagementRoutes(v []*DescribeIstioGatewayRoutesResponseBodyManagementRoutes) *DescribeIstioGatewayRoutesResponseBody {
	s.ManagementRoutes = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBody) SetRequestId(v string) *DescribeIstioGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIstioGatewayRoutesResponseBodyManagementRoutes struct {
	// The name of the ASM gateway.
	ASMGatewayName *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	// The description of the routing rule.
	Description       *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationHost   []*string `json:"DestinationHost,omitempty" xml:"DestinationHost,omitempty" type:"Repeated"`
	DestinationSubSet []*string `json:"DestinationSubSet,omitempty" xml:"DestinationSubSet,omitempty" type:"Repeated"`
	// The namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The name of the routing rule.
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The path that is used to match request URLs.
	RoutePath *string `json:"RoutePath,omitempty" xml:"RoutePath,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// *   `0`: The routing rule is valid.
	// *   `1`: The routing rule is invalid.
	// *   `2`: An error occurs during the creation or update of the routing rule.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIstioGatewayRoutesResponseBodyManagementRoutes) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRoutesResponseBodyManagementRoutes) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetASMGatewayName(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.ASMGatewayName = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetDescription(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.Description = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetDestinationHost(v []*string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.DestinationHost = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetDestinationSubSet(v []*string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.DestinationSubSet = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetNamespace(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetPriority(v int32) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.Priority = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetRouteName(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.RouteName = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetRoutePath(v string) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.RoutePath = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponseBodyManagementRoutes) SetStatus(v int32) *DescribeIstioGatewayRoutesResponseBodyManagementRoutes {
	s.Status = &v
	return s
}

type DescribeIstioGatewayRoutesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIstioGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIstioGatewayRoutesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIstioGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRoutesResponse) SetHeaders(v map[string]*string) *DescribeIstioGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponse) SetStatusCode(v int32) *DescribeIstioGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponse) SetBody(v *DescribeIstioGatewayRoutesResponseBody) *DescribeIstioGatewayRoutesResponse {
	s.Body = v
	return s
}

type DescribeMetadataResponseBody struct {
	// The metadata of ASM, which contains basic information about ASM.
	MetaData *DescribeMetadataResponseBodyMetaData `json:"MetaData,omitempty" xml:"MetaData,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponseBody) SetMetaData(v *DescribeMetadataResponseBodyMetaData) *DescribeMetadataResponseBody {
	s.MetaData = v
	return s
}

func (s *DescribeMetadataResponseBody) SetRequestId(v string) *DescribeMetadataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMetadataResponseBodyMetaData struct {
	// The current version.
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The version information about ASM of a commercial edition.
	ProEdition *DescribeMetadataResponseBodyMetaDataProEdition `json:"ProEdition,omitempty" xml:"ProEdition,omitempty" type:"Struct"`
	// The regions where ASM instances can be created.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The Custom Resource Definitions (CRDs) of the versions.
	VersionCrds []map[string]interface{} `json:"VersionCrds,omitempty" xml:"VersionCrds,omitempty" type:"Repeated"`
	// The ASM version and the corresponding Istio version.
	VersionRegistry []map[string]interface{} `json:"VersionRegistry,omitempty" xml:"VersionRegistry,omitempty" type:"Repeated"`
	// The supported versions.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeMetadataResponseBodyMetaData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetadataResponseBodyMetaData) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponseBodyMetaData) SetCurrentVersion(v string) *DescribeMetadataResponseBodyMetaData {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetProEdition(v *DescribeMetadataResponseBodyMetaDataProEdition) *DescribeMetadataResponseBodyMetaData {
	s.ProEdition = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetRegions(v []*string) *DescribeMetadataResponseBodyMetaData {
	s.Regions = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetVersionCrds(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaData {
	s.VersionCrds = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetVersionRegistry(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaData {
	s.VersionRegistry = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaData) SetVersions(v []*string) *DescribeMetadataResponseBodyMetaData {
	s.Versions = v
	return s
}

type DescribeMetadataResponseBodyMetaDataProEdition struct {
	// The current version.
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The CRDs of the versions.
	VersionCrds []map[string]interface{} `json:"VersionCrds,omitempty" xml:"VersionCrds,omitempty" type:"Repeated"`
	// The ASM version and the corresponding Istio version.
	VersionRegistry []map[string]interface{} `json:"VersionRegistry,omitempty" xml:"VersionRegistry,omitempty" type:"Repeated"`
	// The supported versions.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeMetadataResponseBodyMetaDataProEdition) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetadataResponseBodyMetaDataProEdition) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) SetCurrentVersion(v string) *DescribeMetadataResponseBodyMetaDataProEdition {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) SetVersionCrds(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaDataProEdition {
	s.VersionCrds = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) SetVersionRegistry(v []map[string]interface{}) *DescribeMetadataResponseBodyMetaDataProEdition {
	s.VersionRegistry = v
	return s
}

func (s *DescribeMetadataResponseBodyMetaDataProEdition) SetVersions(v []*string) *DescribeMetadataResponseBodyMetaDataProEdition {
	s.Versions = v
	return s
}

type DescribeMetadataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetadataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetadataResponse) SetHeaders(v map[string]*string) *DescribeMetadataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetadataResponse) SetStatusCode(v int32) *DescribeMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetadataResponse) SetBody(v *DescribeMetadataResponseBody) *DescribeMetadataResponse {
	s.Body = v
	return s
}

type DescribeNamespaceScopeSidecarConfigRequest struct {
	// The namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigRequest) SetNamespace(v string) *DescribeNamespaceScopeSidecarConfigRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigRequest) SetServiceMeshId(v string) *DescribeNamespaceScopeSidecarConfigRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBody struct {
	// The namespace-level sidecar proxy configurations.
	ConfigPatches *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches `json:"ConfigPatches,omitempty" xml:"ConfigPatches,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBody) SetConfigPatches(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) *DescribeNamespaceScopeSidecarConfigResponseBody {
	s.ConfigPatches = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBody) SetRequestId(v string) *DescribeNamespaceScopeSidecarConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches struct {
	// The number of worker threads to run in the istio-proxy container.
	Concurrency    *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	EnableCoreDump *bool  `json:"EnableCoreDump,omitempty" xml:"EnableCoreDump,omitempty"`
	// The inbound ports to be excluded from redirection to the sidecar proxy in the ASM instance.
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The outbound IP ranges in CIDR form to be excluded from redirection to the sidecar proxy in the ASM instance.
	ExcludeOutboundIPRanges *string `json:"ExcludeOutboundIPRanges,omitempty" xml:"ExcludeOutboundIPRanges,omitempty"`
	// The outbound ports to be excluded from redirection to the sidecar proxy in the ASM instance.
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// Indicates whether applications can be started only after the istio-proxy container starts. Valid values:
	//
	// *   `true`
	// *   false
	HoldApplicationUntilProxyStarts *bool `json:"HoldApplicationUntilProxyStarts,omitempty" xml:"HoldApplicationUntilProxyStarts,omitempty"`
	// The inbound ports for which traffic is to be redirected to the sidecar proxy in the ASM instance.
	IncludeInboundPorts *string `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	// The outbound IP ranges in CIDR form for which traffic is to be redirected to the sidecar proxy in the ASM instance.
	IncludeOutboundIPRanges *string `json:"IncludeOutboundIPRanges,omitempty" xml:"IncludeOutboundIPRanges,omitempty"`
	// The outbound ports for which traffic is to be redirected to the sidecar proxy in the ASM instance.
	IncludeOutboundPorts *string `json:"IncludeOutboundPorts,omitempty" xml:"IncludeOutboundPorts,omitempty"`
	// The mode in which the sidecar proxy intercepts inbound traffic. Valid values:
	//
	// *   `REDIRECT` (default): In this mode, source IP addresses are lost during traffic redirection to the sidecar proxy.
	// *   `TPROXY`: In this mode, both the source and destination IP addresses and ports are preserved.
	InterceptionMode *string `json:"InterceptionMode,omitempty" xml:"InterceptionMode,omitempty"`
	// Indicates whether the Domain Name System (DNS) proxy feature is enabled. Valid values:
	//
	// *   `true`
	// *   `false`
	IstioDNSProxyEnabled *bool `json:"IstioDNSProxyEnabled,omitempty" xml:"IstioDNSProxyEnabled,omitempty"`
	// The JSON string that describes the lifecycle of the sidecar proxy.
	LifecycleStr *string `json:"LifecycleStr,omitempty" xml:"LifecycleStr,omitempty"`
	// The log level. Valid values: `info`, `debug`, `trace`, and `error`.
	LogLevel      *string            `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	Privileged    *bool              `json:"Privileged,omitempty" xml:"Privileged,omitempty"`
	ProxyMetadata map[string]*string `json:"ProxyMetadata,omitempty" xml:"ProxyMetadata,omitempty"`
	// The custom Envoy statistics that are reported by the sidecar proxy.
	ProxyStatsMatcher              *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher              `json:"ProxyStatsMatcher,omitempty" xml:"ProxyStatsMatcher,omitempty" type:"Struct"`
	ReadinessFailureThreshold      *int32                                                                                      `json:"ReadinessFailureThreshold,omitempty" xml:"ReadinessFailureThreshold,omitempty"`
	ReadinessInitialDelaySeconds   *int32                                                                                      `json:"ReadinessInitialDelaySeconds,omitempty" xml:"ReadinessInitialDelaySeconds,omitempty"`
	ReadinessPeriodSeconds         *int32                                                                                      `json:"ReadinessPeriodSeconds,omitempty" xml:"ReadinessPeriodSeconds,omitempty"`
	SidecarProxyAckSloResource     *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource     `json:"SidecarProxyAckSloResource,omitempty" xml:"SidecarProxyAckSloResource,omitempty" type:"Struct"`
	SidecarProxyInitAckSloResource *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource `json:"SidecarProxyInitAckSloResource,omitempty" xml:"SidecarProxyInitAckSloResource,omitempty" type:"Struct"`
	// The maximum size of resources that are available to the istio-init container in the pod into which the sidecar proxy is injected. The istio-init container is used in this topic.
	SidecarProxyInitResourceLimit *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit `json:"SidecarProxyInitResourceLimit,omitempty" xml:"SidecarProxyInitResourceLimit,omitempty" type:"Struct"`
	// The minimum size of resources that are required by the istio-init container.
	SidecarProxyInitResourceRequest *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest `json:"SidecarProxyInitResourceRequest,omitempty" xml:"SidecarProxyInitResourceRequest,omitempty" type:"Struct"`
	// The maximum size of resources that are available to the sidecar proxy container.
	SidecarProxyResourceLimit *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit `json:"SidecarProxyResourceLimit,omitempty" xml:"SidecarProxyResourceLimit,omitempty" type:"Struct"`
	// The minimum size of resources that are required by the sidecar proxy container.
	SidecarProxyResourceRequest *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest `json:"SidecarProxyResourceRequest,omitempty" xml:"SidecarProxyResourceRequest,omitempty" type:"Struct"`
	// The maximum period of time allowed for connections to complete on sidecar proxy shutdown.
	TerminationDrainDuration *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
	// The custom configurations of Tracing Analysis.
	Tracing *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing `json:"Tracing,omitempty" xml:"Tracing,omitempty" type:"Struct"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetConcurrency(v int32) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.Concurrency = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetEnableCoreDump(v bool) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.EnableCoreDump = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetExcludeInboundPorts(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ExcludeInboundPorts = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetExcludeOutboundIPRanges(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ExcludeOutboundIPRanges = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetExcludeOutboundPorts(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ExcludeOutboundPorts = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetHoldApplicationUntilProxyStarts(v bool) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.HoldApplicationUntilProxyStarts = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetIncludeInboundPorts(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.IncludeInboundPorts = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetIncludeOutboundIPRanges(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.IncludeOutboundIPRanges = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetIncludeOutboundPorts(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.IncludeOutboundPorts = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetInterceptionMode(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.InterceptionMode = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetIstioDNSProxyEnabled(v bool) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.IstioDNSProxyEnabled = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetLifecycleStr(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.LifecycleStr = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetLogLevel(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.LogLevel = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetPrivileged(v bool) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.Privileged = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetProxyMetadata(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ProxyMetadata = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetProxyStatsMatcher(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ProxyStatsMatcher = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetReadinessFailureThreshold(v int32) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ReadinessFailureThreshold = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetReadinessInitialDelaySeconds(v int32) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ReadinessInitialDelaySeconds = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetReadinessPeriodSeconds(v int32) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ReadinessPeriodSeconds = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetSidecarProxyAckSloResource(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.SidecarProxyAckSloResource = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetSidecarProxyInitAckSloResource(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.SidecarProxyInitAckSloResource = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetSidecarProxyInitResourceLimit(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.SidecarProxyInitResourceLimit = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetSidecarProxyInitResourceRequest(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.SidecarProxyInitResourceRequest = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetSidecarProxyResourceLimit(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.SidecarProxyResourceLimit = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetSidecarProxyResourceRequest(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.SidecarProxyResourceRequest = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetTerminationDrainDuration(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.TerminationDrainDuration = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetTracing(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.Tracing = v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher struct {
	// The prefixes of the custom Envoy statistics that are reported by the sidecar proxy.
	InclusionPrefixes []*string `json:"InclusionPrefixes,omitempty" xml:"InclusionPrefixes,omitempty" type:"Repeated"`
	// The regular expressions for specifying the custom Envoy statistics that are reported by the sidecar proxy.
	InclusionRegexps []*string `json:"InclusionRegexps,omitempty" xml:"InclusionRegexps,omitempty" type:"Repeated"`
	// The suffixes of the custom Envoy statistics that are reported by the sidecar proxy.
	InclusionSuffixes []*string `json:"InclusionSuffixes,omitempty" xml:"InclusionSuffixes,omitempty" type:"Repeated"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) SetInclusionPrefixes(v []*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher {
	s.InclusionPrefixes = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) SetInclusionRegexps(v []*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher {
	s.InclusionRegexps = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) SetInclusionSuffixes(v []*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher {
	s.InclusionSuffixes = v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource struct {
	Limits   map[string]*string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	Requests map[string]*string `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) SetLimits(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource {
	s.Limits = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) SetRequests(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource {
	s.Requests = v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource struct {
	Limits   map[string]*string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	Requests map[string]*string `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) SetLimits(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource {
	s.Limits = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) SetRequests(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource {
	s.Requests = v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit struct {
	// The maximum number of CPU cores.
	ResourceCPULimit *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
	// The maximum size of the memory.
	ResourceMemoryLimit *string `json:"ResourceMemoryLimit,omitempty" xml:"ResourceMemoryLimit,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) SetResourceCPULimit(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit {
	s.ResourceCPULimit = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) SetResourceMemoryLimit(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit {
	s.ResourceMemoryLimit = &v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest struct {
	// The minimum number of CPU cores.
	ResourceCPURequest *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
	// The minimum size of the memory.
	ResourceMemoryRequest *string `json:"ResourceMemoryRequest,omitempty" xml:"ResourceMemoryRequest,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) SetResourceCPURequest(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest {
	s.ResourceCPURequest = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) SetResourceMemoryRequest(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest {
	s.ResourceMemoryRequest = &v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit struct {
	// The maximum number of CPU cores.
	ResourceCPULimit *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
	// The maximum size of the memory.
	ResourceMemoryLimit *string `json:"ResourceMemoryLimit,omitempty" xml:"ResourceMemoryLimit,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) SetResourceCPULimit(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit {
	s.ResourceCPULimit = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) SetResourceMemoryLimit(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit {
	s.ResourceMemoryLimit = &v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest struct {
	// The minimum number of CPU cores.
	ResourceCPURequest *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
	// The minimum size of the memory.
	ResourceMemoryRequest *string `json:"ResourceMemoryRequest,omitempty" xml:"ResourceMemoryRequest,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) SetResourceCPURequest(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest {
	s.ResourceCPURequest = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) SetResourceMemoryRequest(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest {
	s.ResourceMemoryRequest = &v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing struct {
	// The custom tags added to reported spans. The key of a tag is of the string type. The value of a tag is in the JSON format. A custom tag can belong to one of the following types:
	//
	// *   `literal`: The tag value is a fixed value in the JSON format. This tag must contain the `value` field that specifies a literal. Example: `{"value":"test"}`.
	// *   `header`: The tag value is a request header in the JSON format. This tag must contain the `name` field and the `defaultValue` field. The name field indicates the name of the request header. The defaultValue field indicates the default value that is used when no request header is available. Example: `{"name":"test","defaultValue":"test"}`.
	// *   `environment`: The tag value is an environment variable in the JSON format. This tag must contain the `name` field and the `defaultValue` field. The name field indicates the name of the environment variable. The defaultValue field indicates the environment variable that is used when no environment variable is available. Example: `{"name":"test","defaultValue":"test"}`.
	CustomTags map[string]interface{} `json:"CustomTags,omitempty" xml:"CustomTags,omitempty"`
	// The maximum tag length.
	MaxPathTagLength *int32 `json:"MaxPathTagLength,omitempty" xml:"MaxPathTagLength,omitempty"`
	// The sampling rate.
	Sampling *float64 `json:"Sampling,omitempty" xml:"Sampling,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) SetCustomTags(v map[string]interface{}) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing {
	s.CustomTags = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) SetMaxPathTagLength(v int32) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing {
	s.MaxPathTagLength = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) SetSampling(v float64) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing {
	s.Sampling = &v
	return s
}

type DescribeNamespaceScopeSidecarConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespaceScopeSidecarConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespaceScopeSidecarConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) SetHeaders(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) SetStatusCode(v int32) *DescribeNamespaceScopeSidecarConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponse) SetBody(v *DescribeNamespaceScopeSidecarConfigResponseBody) *DescribeNamespaceScopeSidecarConfigResponse {
	s.Body = v
	return s
}

type DescribeNodesInstanceTypeRequest struct {
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeNodesInstanceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodesInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeRequest) SetServiceMeshId(v string) *DescribeNodesInstanceTypeRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeNodesInstanceTypeResponseBody struct {
	// The instance types of the nodes.
	InstanceTypes []*DescribeNodesInstanceTypeResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodesInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodesInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeResponseBody) SetInstanceTypes(v []*DescribeNodesInstanceTypeResponseBodyInstanceTypes) *DescribeNodesInstanceTypeResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBody) SetRequestId(v string) *DescribeNodesInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNodesInstanceTypeResponseBodyInstanceTypes struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Indicates whether the instance type supports Multi-Buffer acceleration. Valid values:
	//
	// *   `true`
	// *   `false`
	MultiBufferEnabled *bool `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	// The instance type of the node.
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNodesInstanceTypeResponseBodyInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodesInstanceTypeResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetKey(v string) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.Key = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetMultiBufferEnabled(v bool) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.MultiBufferEnabled = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetNodeType(v string) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.NodeType = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetValue(v string) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.Value = &v
	return s
}

type DescribeNodesInstanceTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNodesInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodesInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodesInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeResponse) SetHeaders(v map[string]*string) *DescribeNodesInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodesInstanceTypeResponse) SetStatusCode(v int32) *DescribeNodesInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponse) SetBody(v *DescribeNodesInstanceTypeResponseBody) *DescribeNodesInstanceTypeResponse {
	s.Body = v
	return s
}

type DescribeReusableSlbRequest struct {
	// The ID of the Kubernetes cluster on the data plane.
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The network type of the SLB instance. Valid values:
	//
	// *   `intranet`
	// *   `internet`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s DescribeReusableSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeReusableSlbRequest) GoString() string {
	return s.String()
}

func (s *DescribeReusableSlbRequest) SetK8sClusterId(v string) *DescribeReusableSlbRequest {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeReusableSlbRequest) SetNetworkType(v string) *DescribeReusableSlbRequest {
	s.NetworkType = &v
	return s
}

type DescribeReusableSlbResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of SLB instances that can be reused.
	ReusableSlbList []*DescribeReusableSlbResponseBodyReusableSlbList `json:"ReusableSlbList,omitempty" xml:"ReusableSlbList,omitempty" type:"Repeated"`
}

func (s DescribeReusableSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeReusableSlbResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReusableSlbResponseBody) SetRequestId(v string) *DescribeReusableSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReusableSlbResponseBody) SetReusableSlbList(v []*DescribeReusableSlbResponseBodyReusableSlbList) *DescribeReusableSlbResponseBody {
	s.ReusableSlbList = v
	return s
}

type DescribeReusableSlbResponseBodyReusableSlbList struct {
	// The ID of the SLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the SLB instance.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
}

func (s DescribeReusableSlbResponseBodyReusableSlbList) String() string {
	return tea.Prettify(s)
}

func (s DescribeReusableSlbResponseBodyReusableSlbList) GoString() string {
	return s.String()
}

func (s *DescribeReusableSlbResponseBodyReusableSlbList) SetLoadBalancerId(v string) *DescribeReusableSlbResponseBodyReusableSlbList {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeReusableSlbResponseBodyReusableSlbList) SetLoadBalancerName(v string) *DescribeReusableSlbResponseBodyReusableSlbList {
	s.LoadBalancerName = &v
	return s
}

type DescribeReusableSlbResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeReusableSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeReusableSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeReusableSlbResponse) GoString() string {
	return s.String()
}

func (s *DescribeReusableSlbResponse) SetHeaders(v map[string]*string) *DescribeReusableSlbResponse {
	s.Headers = v
	return s
}

func (s *DescribeReusableSlbResponse) SetStatusCode(v int32) *DescribeReusableSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReusableSlbResponse) SetBody(v *DescribeReusableSlbResponseBody) *DescribeReusableSlbResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshAdditionalStatusRequest struct {
	// The check mode of the ASM instance. Valid values:
	//
	// *   `normal`: checks the Server Load Balancer (SLB) instances created for exposing the API server and Istio Pilot, audit logs, and installation of Logtail for clusters on the data plane.
	// *   `full`: checks control plane logs, access logs, security groups, and the elastic IP addresses (EIPs) of the API server in addition to the check items in normal mode.
	CheckMode *string `json:"CheckMode,omitempty" xml:"CheckMode,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusRequest) SetCheckMode(v string) *DescribeServiceMeshAdditionalStatusRequest {
	s.CheckMode = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusRequest) SetServiceMeshId(v string) *DescribeServiceMeshAdditionalStatusRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshAdditionalStatusResponseBody struct {
	// The status of the cluster.
	ClusterStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBody) SetClusterStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) *DescribeServiceMeshAdditionalStatusResponseBody {
	s.ClusterStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBody) SetRequestId(v string) *DescribeServiceMeshAdditionalStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus struct {
	// Indicates whether access logs exist. Valid values:
	//
	// *   `exist`: Access logs exist.
	// *   `not_exist`: Access logs do not exist.
	// *   `failed`: The check fails.
	// *   `time_out`: The check times out.
	AccessLogProjectStatus *string `json:"AccessLogProjectStatus,omitempty" xml:"AccessLogProjectStatus,omitempty"`
	// The check result of the EIP associated with the API server. Valid values:
	//
	// *   `exist`: The EIP exists.
	// *   `not_exist`: The EIP does not exist.
	// *   `failed`: The check fails.
	// *   `time_out`: The check times out.
	// *   `not_in_use`: The EIP is not associated with the API Server.
	// *   `locked`: The EIP is locked.
	ApiServerEIPStatus *string `json:"ApiServerEIPStatus,omitempty" xml:"ApiServerEIPStatus,omitempty"`
	// The check results of the SLB instance created for exposing the API server.
	ApiServerLoadBalancerStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus `json:"ApiServerLoadBalancerStatus,omitempty" xml:"ApiServerLoadBalancerStatus,omitempty" type:"Struct"`
	// Indicates whether audit logs exist. Valid values:
	//
	// *   `exist`
	// *   `not exist`
	AuditProjectStatus *string `json:"AuditProjectStatus,omitempty" xml:"AuditProjectStatus,omitempty"`
	// The check results of the SLB instance that is created for exposing Istio Pilot and used during canary release.
	CanaryPilotLoadBalancerStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus `json:"CanaryPilotLoadBalancerStatus,omitempty" xml:"CanaryPilotLoadBalancerStatus,omitempty" type:"Struct"`
	// Indicates whether control plane logs exist. Valid values:
	//
	// *   `exist`: Control plane logs exist.
	// *   `not_exist`: Control plane logs do not exist.
	// *   `failed`: The check fails.
	// *   `time_out`: The check times out.
	ControlPlaneProjectStatus *string `json:"ControlPlaneProjectStatus,omitempty" xml:"ControlPlaneProjectStatus,omitempty"`
	// Indicates whether Logtail is installed in clusters on the data plane.
	LogtailStatusRecord map[string]interface{} `json:"LogtailStatusRecord,omitempty" xml:"LogtailStatusRecord,omitempty"`
	// The check results of the SLB instance created for exposing Istio Pilot.
	PilotLoadBalancerStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus `json:"PilotLoadBalancerStatus,omitempty" xml:"PilotLoadBalancerStatus,omitempty" type:"Struct"`
	// The status of the RAM OAuth application that is integrated with Mesh Topology. Valid values:
	//
	// *   `exist`: The RAM OAuth application exists.
	// *   `reused`: The RAM OAuth application is reused.
	// *   `not_exist`: The RAM OAuth application does not exist.
	// *   `failed`: The check fails.
	// *   `time_out`: The check times out.
	RAMApplicationStatus *string `json:"RAMApplicationStatus,omitempty" xml:"RAMApplicationStatus,omitempty"`
	// Indicates whether the security group is reused. Valid values:
	//
	// *   `reused`: The security group is reused.
	// *   `not_reused`: The security group is not reused.
	// *   `failed`: The check fails.
	// *   `time_out`: The check times out.
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetAccessLogProjectStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.AccessLogProjectStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetApiServerEIPStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.ApiServerEIPStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetApiServerLoadBalancerStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.ApiServerLoadBalancerStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetAuditProjectStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.AuditProjectStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetCanaryPilotLoadBalancerStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.CanaryPilotLoadBalancerStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetControlPlaneProjectStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.ControlPlaneProjectStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetLogtailStatusRecord(v map[string]interface{}) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.LogtailStatusRecord = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetPilotLoadBalancerStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.PilotLoadBalancerStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetRAMApplicationStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.RAMApplicationStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetSgStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.SgStatus = &v
	return s
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus struct {
	// Indicates whether the SLB instance is locked. Valid values:
	//
	// *   `true`
	// *   `false`
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The billing method of the SLB instance. Valid values:
	//
	// *   `PrePay`: subscription
	// *   `PayOnDemand`: pay-as-you-go
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates whether the SLB instance is reused. Valid values:
	//
	// *   `true`
	// *   `false`
	Reused *bool `json:"Reused,omitempty" xml:"Reused,omitempty"`
	// The check result of the number of backend servers of the SLB instance created for exposing the API server. Valid values:
	//
	// *   `too_much`: An excessive number of backend servers are created.
	// *   `num_exact`: A proper number of backend servers are created.
	// *   `too_little`: The number of backend servers falls short.
	SLBBackEndServerNumStatus *string `json:"SLBBackEndServerNumStatus,omitempty" xml:"SLBBackEndServerNumStatus,omitempty"`
	// The check result of the SLB instance. Valid values:
	//
	// *   `exist`: The SLB instance exists.
	// *   `not_exist`: The SLB instance does not exist.
	// *   `conflict`: Conflicts are detected.
	// *   `failed`: The check fails.
	// *   `time_out`: The check times out.
	SLBExistStatus *string `json:"SLBExistStatus,omitempty" xml:"SLBExistStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetLocked(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.Locked = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetPayType(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.PayType = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetReused(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.Reused = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetSLBBackEndServerNumStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.SLBBackEndServerNumStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetSLBExistStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.SLBExistStatus = &v
	return s
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus struct {
	// Indicates whether the SLB instance is locked due to overdue payments. Valid values: `true` `false`
	Locked  *bool   `json:"Locked,omitempty" xml:"Locked,omitempty"`
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates whether the SLB instance is reused. Valid values:
	//
	// *   `true`: The SLB instance is reused. Non-ASM listener configuration is detected in the listener configurations of the SLB instance.
	// *   `false`: The SLB instance is not reused.
	Reused *bool `json:"Reused,omitempty" xml:"Reused,omitempty"`
	// The check result of the number of backend servers of the SLB instance created for exposing Istio Pilot. Valid values:
	//
	// *   `num_exact`: A proper number of backend servers are created.
	// *   `too_much`: An excessive number of backend servers are created.
	// *   `too_little`: The number of backend servers falls short.
	SLBBackEndServerNumStatus *string `json:"SLBBackEndServerNumStatus,omitempty" xml:"SLBBackEndServerNumStatus,omitempty"`
	// The check result of the SLB instance. Valid values:
	//
	// *   `exist`: The SLB instance exists.
	// *   `not_exist`: The SLB instance does not exist.
	// *   `time_out`: The check times out.
	// *   `failed`: The SLB instance has expired.
	SLBExistStatus *string `json:"SLBExistStatus,omitempty" xml:"SLBExistStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetLocked(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.Locked = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetPayType(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.PayType = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetReused(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.Reused = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetSLBBackEndServerNumStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.SLBBackEndServerNumStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetSLBExistStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.SLBExistStatus = &v
	return s
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus struct {
	// Indicates whether the SLB instance is locked. Valid values:
	//
	// *   `true`
	// *   `false`
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The billing method of the SLB instance. Valid values:
	//
	// *   `PrePay`: subscription
	// *   `PayOnDemand`: pay-as-you-go
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates whether the SLB instance is reused. Valid values:
	//
	// *   `true`
	// *   `false`
	Reused *bool `json:"Reused,omitempty" xml:"Reused,omitempty"`
	// The check result of the number of backend servers of the SLB instance created for exposing Istio Pilot. Valid values:
	//
	// *   `too_much`: An excessive number of backend servers are created.
	// *   `num_exact`: A proper number of backend servers are created.
	// *   `too_little`: The number of backend servers falls short.
	SLBBackEndServerNumStatus *string `json:"SLBBackEndServerNumStatus,omitempty" xml:"SLBBackEndServerNumStatus,omitempty"`
	// The check result of the SLB instance. Valid values:
	//
	// *   `exist`: The SLB instance exists.
	// *   `not_exist`: The SLB instance does not exist.
	// *   `conflict`: Conflicts are detected.
	// *   `failed`: The check fails.
	// *   `time_out`: The check times out.
	SLBExistStatus *string `json:"SLBExistStatus,omitempty" xml:"SLBExistStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetLocked(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.Locked = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetPayType(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.PayType = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetReused(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.Reused = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetSLBBackEndServerNumStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.SLBBackEndServerNumStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetSLBExistStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.SLBExistStatus = &v
	return s
}

type DescribeServiceMeshAdditionalStatusResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshAdditionalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshAdditionalStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshAdditionalStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponse) SetStatusCode(v int32) *DescribeServiceMeshAdditionalStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponse) SetBody(v *DescribeServiceMeshAdditionalStatusResponseBody) *DescribeServiceMeshAdditionalStatusResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshClustersRequest struct {
	// The maximum number of clusters in a cluster list.
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The position where the query starts.
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersRequest) SetLimit(v int64) *DescribeServiceMeshClustersRequest {
	s.Limit = &v
	return s
}

func (s *DescribeServiceMeshClustersRequest) SetOffset(v int64) *DescribeServiceMeshClustersRequest {
	s.Offset = &v
	return s
}

func (s *DescribeServiceMeshClustersRequest) SetServiceMeshId(v string) *DescribeServiceMeshClustersRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshClustersResponseBody struct {
	// The queried clusters.
	Clusters         []*DescribeServiceMeshClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	NumberOfClusters *int64                                             `json:"NumberOfClusters,omitempty" xml:"NumberOfClusters,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersResponseBody) SetClusters(v []*DescribeServiceMeshClustersResponseBodyClusters) *DescribeServiceMeshClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeServiceMeshClustersResponseBody) SetNumberOfClusters(v int64) *DescribeServiceMeshClustersResponseBody {
	s.NumberOfClusters = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBody) SetRequestId(v string) *DescribeServiceMeshClustersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceMeshClustersResponseBodyClusters struct {
	// The domain name of the cluster.
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster type.
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the cluster was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message about the cluster.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates that the cluster is available or the reason why the cluster cannot be added to the ASM instance. Valid values:
	//
	// *   `0`: The cluster can be added to the ASM instance.
	// *   `1`: The cluster cannot be added to the ASM instance because you do not have administrator permissions on the cluster.
	// *   `2`: The cluster cannot be added to the ASM instance because the cluster and the ASM instance reside in different VPCs between which no private connections are built.
	// *   `3`: The CIDR block of the cluster conflicts with that of the ASM instance.
	// *   `4`: The cluster has a namespace that is named istio system.
	ForbiddenFlag *int64  `json:"ForbiddenFlag,omitempty" xml:"ForbiddenFlag,omitempty"`
	ForbiddenInfo *string `json:"ForbiddenInfo,omitempty" xml:"ForbiddenInfo,omitempty"`
	// The name of the cluster.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region in which the cluster resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The ID of the security group.
	SgId *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	// The state of the cluster. Valid values:
	//
	// *   `running`: The cluster is running.
	// *   `starting`: The cluster is starting.
	// *   `stopping`: The cluster is being stopped.
	// *   `stopped`: The cluster is stopped.
	// *   `failed`: The cluster fails to be run.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the cluster was last modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version number of the cluster.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeServiceMeshClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetClusterDomain(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetClusterId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetClusterType(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetCreationTime(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.CreationTime = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetErrorMessage(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetForbiddenFlag(v int64) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ForbiddenFlag = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetForbiddenInfo(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ForbiddenInfo = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetName(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetRegionId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetServiceMeshId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetSgId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.SgId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetState(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetUpdateTime(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.UpdateTime = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetVersion(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.Version = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetVpcId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

type DescribeServiceMeshClustersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshClustersResponse) SetStatusCode(v int32) *DescribeServiceMeshClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshClustersResponse) SetBody(v *DescribeServiceMeshClustersResponseBody) *DescribeServiceMeshClustersResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshDetailRequest struct {
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailRequest) SetServiceMeshId(v string) *DescribeServiceMeshDetailRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshDetailResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the ASM instance.
	ServiceMesh *DescribeServiceMeshDetailResponseBodyServiceMesh `json:"ServiceMesh,omitempty" xml:"ServiceMesh,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBody) SetRequestId(v string) *DescribeServiceMeshDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBody) SetServiceMesh(v *DescribeServiceMeshDetailResponseBodyServiceMesh) *DescribeServiceMeshDetailResponseBody {
	s.ServiceMesh = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMesh struct {
	// The specification of the ASM instance. Valid values:
	//
	// - `standard`: Standard Edition
	// - `enterprise`: Enterprise Edition
	// - `ultimate`: Ultimate Edition
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The list of clusters.
	Clusters []*string `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The endpoints of the ASM instance.
	Endpoints *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud service instance for which the ASM instance is created.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud service for which the ASM instance is created. Valid values:
	//
	// - `ackone`: The ASM instance is created for Alibaba Cloud Distributed Cloud Container Platform (ACK One).
	// - An empty value indicates that the ASM instance is created by the user.
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The basic information about the ASM instance.
	ServiceMeshInfo *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo `json:"ServiceMeshInfo,omitempty" xml:"ServiceMeshInfo,omitempty" type:"Struct"`
	// The specifications of the ASM instance.
	Spec *DescribeServiceMeshDetailResponseBodyServiceMeshSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMesh) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMesh) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetClusterSpec(v string) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetClusters(v []*string) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.Clusters = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetEndpoints(v *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.Endpoints = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetOwnerId(v string) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.OwnerId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetOwnerType(v string) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.OwnerType = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetServiceMeshInfo(v *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.ServiceMeshInfo = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetSpec(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.Spec = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints struct {
	// The endpoint that is used to access the API server from the internal network.
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	// The endpoint that is used to access Istio Pilot from the internal network.
	IntranetPilotEndpoint *string `json:"IntranetPilotEndpoint,omitempty" xml:"IntranetPilotEndpoint,omitempty"`
	// The endpoint that is used to expose the API server to the Internet.
	PublicApiServerEndpoint *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
	// The endpoint that is used to expose Istio Pilot to the Internet.
	PublicPilotEndpoint *string `json:"PublicPilotEndpoint,omitempty" xml:"PublicPilotEndpoint,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) SetIntranetApiServerEndpoint(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	s.IntranetApiServerEndpoint = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) SetIntranetPilotEndpoint(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	s.IntranetPilotEndpoint = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) SetPublicApiServerEndpoint(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	s.PublicApiServerEndpoint = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) SetPublicPilotEndpoint(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	s.PublicPilotEndpoint = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo struct {
	// The time when the ASM instance was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message that is returned when the call failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the ASM instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The edition of the ASM instance. Valid values:
	//
	// *   `Default`: Standard Edition
	// *   `Pro`: Professional Edition
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region in which the ASM instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the ASM instance.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the ASM instance was last modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version of the ASM instance.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetCreationTime(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.CreationTime = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetErrorMessage(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetName(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetProfile(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.Profile = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetRegionId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetServiceMeshId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetState(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.State = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetUpdateTime(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) SetVersion(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	s.Version = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpec struct {
	// The information about load balancing.
	LoadBalancer *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Struct"`
	// The configurations of the ASM instance.
	MeshConfig *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	// The network configurations of the ASM instance.
	Network *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpec) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) SetLoadBalancer(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) *DescribeServiceMeshDetailResponseBodyServiceMeshSpec {
	s.LoadBalancer = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) SetMeshConfig(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) *DescribeServiceMeshDetailResponseBodyServiceMeshSpec {
	s.MeshConfig = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) SetNetwork(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) *DescribeServiceMeshDetailResponseBodyServiceMeshSpec {
	s.Network = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer struct {
	// The ID of the SLB instance that is used when the API server is exposed to the Internet.
	ApiServerLoadbalancerId *string `json:"ApiServerLoadbalancerId,omitempty" xml:"ApiServerLoadbalancerId,omitempty"`
	// Indicates whether the API server is exposed to the Internet. Valid values:
	//
	// *   `true`: The API server is exposed to the Internet.
	// *   `false`: The API server is not exposed to the Internet.
	ApiServerPublicEip *bool `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	// Indicates whether Istio Pilot is exposed to the Internet. Valid values:
	//
	// *   `true`: Istio Pilot is exposed to the Internet.
	// *   `false`: Istio Pilot is not exposed to the Internet.
	PilotPublicEip *bool `json:"PilotPublicEip,omitempty" xml:"PilotPublicEip,omitempty"`
	// The ID of the Server Load Balancer (SLB) instance that is used when Istio Pilot is exposed to the Internet.
	PilotPublicLoadbalancerId *string `json:"PilotPublicLoadbalancerId,omitempty" xml:"PilotPublicLoadbalancerId,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetApiServerLoadbalancerId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.ApiServerLoadbalancerId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetApiServerPublicEip(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.ApiServerPublicEip = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetPilotPublicEip(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.PilotPublicEip = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetPilotPublicLoadbalancerId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.PilotPublicLoadbalancerId = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig struct {
	// The configurations of access log collection.
	AccessLog *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog `json:"AccessLog,omitempty" xml:"AccessLog,omitempty" type:"Struct"`
	// The information about mesh audit.
	Audit *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit `json:"Audit,omitempty" xml:"Audit,omitempty" type:"Struct"`
	// The configurations of control-plane log collection.
	ControlPlaneLogInfo *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo `json:"ControlPlaneLogInfo,omitempty" xml:"ControlPlaneLogInfo,omitempty" type:"Struct"`
	// Indicates whether a custom Zipkin system is used. Valid values:
	//
	// *   `true`: A custom Zipkin system is used.
	// *   `false`: No custom Zipkin system is used.
	CustomizedZipkin *bool `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	// The information about the edition.
	Edition *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition `json:"Edition,omitempty" xml:"Edition,omitempty" type:"Struct"`
	// Indicates whether the feature that routes traffic to the nearest instance is enabled. Valid values:
	//
	// *   `true`: The feature is enabled.
	// *   `false`: The feature is disabled.
	EnableLocalityLB *bool `json:"EnableLocalityLB,omitempty" xml:"EnableLocalityLB,omitempty"`
	// The IP ranges in CIDR form to be excluded from redirection to sidecar proxies in the ASM instance.
	ExcludeIPRanges *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	// The inbound ports to be excluded from redirection to sidecar proxies in the ASM instance.
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The outbound ports to be excluded from redirection to sidecar proxies in the ASM instance.
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// The configurations of additional features for the ASM instance.
	ExtraConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration `json:"ExtraConfiguration,omitempty" xml:"ExtraConfiguration,omitempty" type:"Struct"`
	// The IP ranges in CIDR form to redirect to the sidecar proxies in the ASM instance.
	IncludeIPRanges *string `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	// The information about the Kubernetes API.
	K8sNewAPIsSupport *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport `json:"K8sNewAPIsSupport,omitempty" xml:"K8sNewAPIsSupport,omitempty" type:"Struct"`
	// The configurations of mesh topology.
	Kiali *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali `json:"Kiali,omitempty" xml:"Kiali,omitempty" type:"Struct"`
	// The configurations of cross-region load balancing.
	LocalityLB *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB `json:"LocalityLB,omitempty" xml:"LocalityLB,omitempty" type:"Struct"`
	// The configurations of Microservices Engine (MSE).
	MSE *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE `json:"MSE,omitempty" xml:"MSE,omitempty" type:"Struct"`
	// The information about the Open Policy Agent (OPA) plug-in.
	OPA *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA `json:"OPA,omitempty" xml:"OPA,omitempty" type:"Struct"`
	// The outbound traffic policy. Valid values:
	//
	// *   `ALLOW_ANY`: Outbound traffic to all external services is allowed.
	// *   `REGISTRY_ONLY`: Outbound traffic is allowed to only external services that are defined in the service registry of the ASM instance.
	OutboundTrafficPolicy *string `json:"OutboundTrafficPolicy,omitempty" xml:"OutboundTrafficPolicy,omitempty"`
	// The Pilot configurations.
	Pilot *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot `json:"Pilot,omitempty" xml:"Pilot,omitempty" type:"Struct"`
	// The configurations of Prometheus monitoring.
	Prometheus *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus `json:"Prometheus,omitempty" xml:"Prometheus,omitempty" type:"Struct"`
	// The configurations of protocol support.
	ProtocolSupport *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport `json:"ProtocolSupport,omitempty" xml:"ProtocolSupport,omitempty" type:"Struct"`
	// The proxy configurations.
	Proxy *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy `json:"Proxy,omitempty" xml:"Proxy,omitempty" type:"Struct"`
	// The configurations of the sidecar injector.
	SidecarInjector *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector `json:"SidecarInjector,omitempty" xml:"SidecarInjector,omitempty" type:"Struct"`
	// Indicates whether Prometheus monitoring is enabled. We recommend that you use [Prometheus Service of Application Real-Time Monitoring Service (ARMS)](https://arms.console.aliyun.com/). Valid values:
	//
	// *   `true`: Prometheus monitoring is enabled.
	// *   `false`: Prometheus monitoring is disabled.
	Telemetry *bool `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	// Indicates whether tracing analysis is enabled. This feature can be enabled only after [Tracing Analysis](https://tracing-analysis.console.aliyun.com/) is activated. Valid values:
	//
	// *   `true`: Tracing analysis is enabled.
	// *   `false`: Tracing analysis is disabled.
	Tracing *bool `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
	// The configurations of WebAssembly Filter.
	WebAssemblyFilterDeployment *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment `json:"WebAssemblyFilterDeployment,omitempty" xml:"WebAssemblyFilterDeployment,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetAccessLog(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.AccessLog = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetAudit(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.Audit = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetControlPlaneLogInfo(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.ControlPlaneLogInfo = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetCustomizedZipkin(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.CustomizedZipkin = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetEdition(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.Edition = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetEnableLocalityLB(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.EnableLocalityLB = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetExcludeIPRanges(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.ExcludeIPRanges = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetExcludeInboundPorts(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.ExcludeInboundPorts = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetExcludeOutboundPorts(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.ExcludeOutboundPorts = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetExtraConfiguration(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.ExtraConfiguration = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetIncludeIPRanges(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.IncludeIPRanges = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetK8sNewAPIsSupport(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.K8sNewAPIsSupport = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetKiali(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.Kiali = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetLocalityLB(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.LocalityLB = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetMSE(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.MSE = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetOPA(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.OPA = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetOutboundTrafficPolicy(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.OutboundTrafficPolicy = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetPilot(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.Pilot = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetPrometheus(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.Prometheus = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetProtocolSupport(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.ProtocolSupport = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetProxy(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.Proxy = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetSidecarInjector(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.SidecarInjector = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetTelemetry(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.Telemetry = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetTracing(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.Tracing = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) SetWebAssemblyFilterDeployment(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	s.WebAssemblyFilterDeployment = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog struct {
	// Indicates whether access log collection is enabled. Valid values:
	//
	// *   `true`: Access log collection is enabled.
	// *   `false`: Access log collection is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the Log Service project that stores access logs.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) SetProject(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog {
	s.Project = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit struct {
	// Indicates whether an audit project exists in the ASM instance. Valid values:
	//
	// *   `audit_project_exist`: An audit project exists.
	// *   `audit_project_not_exist`: No audit project exists.
	AuditProjectStatus *string `json:"AuditProjectStatus,omitempty" xml:"AuditProjectStatus,omitempty"`
	// Indicates whether mesh audit is enabled. Valid values:
	//
	// *   `true`: Mesh audit is enabled.
	// *   `false`: Mesh audit is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the Log Service project that is used for mesh audit.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) SetAuditProjectStatus(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit {
	s.AuditProjectStatus = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) SetProject(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit {
	s.Project = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo struct {
	// Indicates whether the collection of control-plane logs is enabled. Valid values:
	//
	// *   `true`: The collection of control-plane logs is enabled.
	// *   `false`: The collection of control-plane logs is disabled.
	Enabled *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	LogTTL  *int32 `json:"LogTTL,omitempty" xml:"LogTTL,omitempty"`
	// The name of the Log Service project that stores control-plane logs.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) SetLogTTL(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo {
	s.LogTTL = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) SetProject(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo {
	s.Project = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition struct {
	// The version of the Istiod image.
	IstiodImageTag *string `json:"IstiodImageTag,omitempty" xml:"IstiodImageTag,omitempty"`
	// The name of the edition.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the Istio Proxy image.
	ProxyImageTag *string `json:"ProxyImageTag,omitempty" xml:"ProxyImageTag,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) SetIstiodImageTag(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition {
	s.IstiodImageTag = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) SetName(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) SetProxyImageTag(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition {
	s.ProxyImageTag = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration struct {
	// The configurations of additional features for access log collection
	AccessLogExtraConf *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf `json:"AccessLogExtraConf,omitempty" xml:"AccessLogExtraConf,omitempty" type:"Struct"`
	// The configurations of adaptive xDS optimization.
	AdaptiveXdsConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration `json:"AdaptiveXdsConfiguration,omitempty" xml:"AdaptiveXdsConfiguration,omitempty" type:"Struct"`
	// The configurations of automatic diagnosis for the ASM instance.
	AutoDiagnosis *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis `json:"AutoDiagnosis,omitempty" xml:"AutoDiagnosis,omitempty" type:"Struct"`
	// Access to Istio resources by using the Kubernetes API on the data plane.
	CRAggregationConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration `json:"CRAggregationConfiguration,omitempty" xml:"CRAggregationConfiguration,omitempty" type:"Struct"`
	// Indicates whether the Kubernetes API of clusters on the data plane can be used to access Istio resources. Valid values:
	//
	// *   `true`: The Kubernetes API of clusters on the data plane can be used to access Istio resources.
	// *   `false`: The Kubernetes API of clusters on the data plane cannot be used to access Istio resources.
	CRAggregationEnabled *bool `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	// The label selectors used to specify namespaces on the data plane. The control plane discovers and process only application services in the specified namespaces.
	DiscoverySelectors []map[string]interface{} `json:"DiscoverySelectors,omitempty" xml:"DiscoverySelectors,omitempty" type:"Repeated"`
	// The configurations of the rollback feature for Istio resources.
	IstioCRHistory *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory `json:"IstioCRHistory,omitempty" xml:"IstioCRHistory,omitempty" type:"Struct"`
	// The lifecycle of Istio Proxy.
	Lifecycle *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty" type:"Struct"`
	// The information about Transport Layer Security (TLS) acceleration based on MulitiBuffer.
	MultiBuffer *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer `json:"MultiBuffer,omitempty" xml:"MultiBuffer,omitempty" type:"Struct"`
	// The configurations of Node Feature Discovery (NFD).
	NFDConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration `json:"NFDConfiguration,omitempty" xml:"NFDConfiguration,omitempty" type:"Struct"`
	// The configurations of the feature of controlling the OPA injection scope.
	OPAScopeInjection *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection `json:"OPAScopeInjection,omitempty" xml:"OPAScopeInjection,omitempty" type:"Struct"`
	// The resource limits on the istio-init container.
	SidecarProxyInitResourceLimit *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit `json:"SidecarProxyInitResourceLimit,omitempty" xml:"SidecarProxyInitResourceLimit,omitempty" type:"Struct"`
	// The resources that are requested by the istio-init container.
	SidecarProxyInitResourceRequest *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest `json:"SidecarProxyInitResourceRequest,omitempty" xml:"SidecarProxyInitResourceRequest,omitempty" type:"Struct"`
	// The maximum period of time that Istio Proxy waits for a request to end.
	TerminationDrainDuration *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetAccessLogExtraConf(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.AccessLogExtraConf = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetAdaptiveXdsConfiguration(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.AdaptiveXdsConfiguration = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetAutoDiagnosis(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.AutoDiagnosis = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetCRAggregationConfiguration(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.CRAggregationConfiguration = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetCRAggregationEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.CRAggregationEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetDiscoverySelectors(v []map[string]interface{}) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.DiscoverySelectors = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetIstioCRHistory(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.IstioCRHistory = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetLifecycle(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.Lifecycle = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetMultiBuffer(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.MultiBuffer = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetNFDConfiguration(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.NFDConfiguration = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetOPAScopeInjection(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.OPAScopeInjection = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetSidecarProxyInitResourceLimit(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.SidecarProxyInitResourceLimit = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetSidecarProxyInitResourceRequest(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.SidecarProxyInitResourceRequest = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetTerminationDrainDuration(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.TerminationDrainDuration = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf struct {
	GatewayEnabled *bool `json:"GatewayEnabled,omitempty" xml:"GatewayEnabled,omitempty"`
	// The retention period for the access logs of the ingress gateway. Unit: day. The logs are collected by using the Log Service. For example, a value of 30 indicates that the logs are retained for 30 days.
	GatewayLifecycle *int32 `json:"GatewayLifecycle,omitempty" xml:"GatewayLifecycle,omitempty"`
	SidecarEnabled   *bool  `json:"SidecarEnabled,omitempty" xml:"SidecarEnabled,omitempty"`
	// The retention period for the access logs of sidecar proxies. Unit: day. The logs are collected by using the Log Service. For example, a value of 30 indicates that the logs are retained for 30 days.
	SidecarLifecycle *int32 `json:"SidecarLifecycle,omitempty" xml:"SidecarLifecycle,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) SetGatewayEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf {
	s.GatewayEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) SetGatewayLifecycle(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf {
	s.GatewayLifecycle = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) SetSidecarEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf {
	s.SidecarEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) SetSidecarLifecycle(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf {
	s.SidecarLifecycle = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration struct {
	// Indicates whether Horizontal Pod Autoscaling (HPA) is enabled for the egress gateway.
	EgressAutoscaleEnabled *bool `json:"EgressAutoscaleEnabled,omitempty" xml:"EgressAutoscaleEnabled,omitempty"`
	// The CPU resource configurations of the egress gateway when HPA is enabled.
	EgressHpaCpu *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu `json:"EgressHpaCpu,omitempty" xml:"EgressHpaCpu,omitempty" type:"Struct"`
	// The memory resource configurations of the egress gateway when HPA is enabled.
	EgressHpaMemory *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory `json:"EgressHpaMemory,omitempty" xml:"EgressHpaMemory,omitempty" type:"Struct"`
	// The maximum number of egress gateway pod replicas when HPA is enabled.
	EgressMaxReplica *int32 `json:"EgressMaxReplica,omitempty" xml:"EgressMaxReplica,omitempty"`
	// The minimum number of egress gateway pod replicas when HPA is enabled.
	EgressMinReplica *int32 `json:"EgressMinReplica,omitempty" xml:"EgressMinReplica,omitempty"`
	// The number of the egress gateway pod replicas.
	EgressReplicaCount *int32 `json:"EgressReplicaCount,omitempty" xml:"EgressReplicaCount,omitempty"`
	// The resource configurations of the egress gateway that is used by adaptive xDS optimization.
	EgressResources *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources `json:"EgressResources,omitempty" xml:"EgressResources,omitempty" type:"Struct"`
	// Indicates whether adaptive xDS optimization is enabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) SetEgressAutoscaleEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	s.EgressAutoscaleEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) SetEgressHpaCpu(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	s.EgressHpaCpu = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) SetEgressHpaMemory(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	s.EgressHpaMemory = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) SetEgressMaxReplica(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	s.EgressMaxReplica = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) SetEgressMinReplica(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	s.EgressMinReplica = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) SetEgressReplicaCount(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	s.EgressReplicaCount = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) SetEgressResources(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	s.EgressResources = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	s.Enabled = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu struct {
	// The expected CPU utilization when HPA is enabled. Valid values: 1 to 100. If the CPU utilization exceeds this value, the number of pod replicas increases. If the CPU utilization is less than this value, the number of pod replicas decreases.
	TargetAverageUtilization *int32 `json:"TargetAverageUtilization,omitempty" xml:"TargetAverageUtilization,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) SetTargetAverageUtilization(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu {
	s.TargetAverageUtilization = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory struct {
	// The expected memory usage when HPA is enabled. Valid values: 1 to 100. If the memory usage exceeds this value, the number of pod replicas increases. If the memory usage is less than this value, the number of pod replicas decreases.
	TargetAverageUtilization *int32 `json:"TargetAverageUtilization,omitempty" xml:"TargetAverageUtilization,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) SetTargetAverageUtilization(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory {
	s.TargetAverageUtilization = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources struct {
	// The resources that are available to the egress gateway.
	Limits map[string]interface{} `json:"Limits,omitempty" xml:"Limits,omitempty"`
	// The resources that are requested by the egress gateway.
	Requests map[string]interface{} `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) SetLimits(v map[string]interface{}) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources {
	s.Limits = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) SetRequests(v map[string]interface{}) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources {
	s.Requests = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis struct {
	// Indicates whether automatic diagnosis is enabled for the ASM instance. If you enable this feature, the ASM instance is automatically diagnosed five minutes after you modify an Istio resource.
	AutoDiagnosisEnabled *bool `json:"AutoDiagnosisEnabled,omitempty" xml:"AutoDiagnosisEnabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) SetAutoDiagnosisEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis {
	s.AutoDiagnosisEnabled = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration struct {
	// Indicates whether Istio resources can be accessed by using the Kubernetes API on the data plane.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration {
	s.Enabled = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory struct {
	// Indicates whether the rollback feature for Istio resources is enabled. Valid values:
	//
	// *   `true`: The rollback feature for Istio resources is enabled.
	// *   `false`: The rollback feature for Istio resources is disabled.
	EnableHistory *bool `json:"EnableHistory,omitempty" xml:"EnableHistory,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) SetEnableHistory(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory {
	s.EnableHistory = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle struct {
	// The post-start parameters.
	PostStart *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart `json:"postStart,omitempty" xml:"postStart,omitempty" type:"Struct"`
	// The pre-close parameters.
	PreStop *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop `json:"preStop,omitempty" xml:"preStop,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) SetPostStart(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle {
	s.PostStart = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) SetPreStop(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle {
	s.PreStop = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart struct {
	// The post-start script.
	Exec *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec `json:"exec,omitempty" xml:"exec,omitempty" type:"Struct"`
	// The HTTP GET request that is sent before the instance stops.
	HttpGet *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet `json:"httpGet,omitempty" xml:"httpGet,omitempty" type:"Struct"`
	// The TCP socket request that is sent.
	TcpSocket *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket `json:"tcpSocket,omitempty" xml:"tcpSocket,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) SetExec(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart {
	s.Exec = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) SetHttpGet(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart {
	s.HttpGet = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) SetTcpSocket(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart {
	s.TcpSocket = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec struct {
	// The executed command.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) SetCommand(v []*string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec {
	s.Command = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet struct {
	// The URL of the request.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The HTTP request headers.
	HttpHeaders []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders `json:"httpHeaders,omitempty" xml:"httpHeaders,omitempty" type:"Repeated"`
	// The port number of the request.
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The request method. Valid values: `http` and `https`.
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) SetHttpHeaders(v []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet {
	s.HttpHeaders = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet {
	s.Port = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) SetScheme(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet {
	s.Scheme = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders struct {
	// The name of the HTTP request header.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the HTTP request header field.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) SetName(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) SetValue(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders {
	s.Value = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket struct {
	// The URL of the TCP socket request.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The port number of the TCP socket request.
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket {
	s.Port = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop struct {
	// The pre-close script.
	Exec *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec `json:"exec,omitempty" xml:"exec,omitempty" type:"Struct"`
	// The HTTP GET request that is sent before the instance stops.
	HttpGet *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet `json:"httpGet,omitempty" xml:"httpGet,omitempty" type:"Struct"`
	// The TCP socket request that is sent.
	TcpSocket *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket `json:"tcpSocket,omitempty" xml:"tcpSocket,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) SetExec(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop {
	s.Exec = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) SetHttpGet(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop {
	s.HttpGet = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) SetTcpSocket(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop {
	s.TcpSocket = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec struct {
	// The executed command.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) SetCommand(v []*string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec {
	s.Command = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet struct {
	// The URL of the request.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The HTTP request headers.
	HttpHeaders []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders `json:"httpHeaders,omitempty" xml:"httpHeaders,omitempty" type:"Repeated"`
	// The port number of the request.
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The request method. Valid values: `http` and `https`.
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) SetHttpHeaders(v []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet {
	s.HttpHeaders = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet {
	s.Port = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) SetScheme(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet {
	s.Scheme = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders struct {
	// The name of the HTTP request header.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the HTTP request header field.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) SetName(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) SetValue(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders {
	s.Value = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket struct {
	// The URL of the request.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The port number of the request.
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket {
	s.Port = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer struct {
	// Indicates whether MulitiBuffer-based TLS acceleration is enabled. Valid values:
	//
	// *   `true`: MulitiBuffer-based TLS acceleration is enabled.
	// *   `false`: MulitiBuffer-based TLS acceleration is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The pull-request latency.
	PollDelay *string `json:"PollDelay,omitempty" xml:"PollDelay,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) SetPollDelay(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer {
	s.PollDelay = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration struct {
	// Indicates whether NFD is enabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether feature labels on nodes are cleared when NFD is disabled.
	NFDLabelPruned *bool `json:"NFDLabelPruned,omitempty" xml:"NFDLabelPruned,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) SetNFDLabelPruned(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration {
	s.NFDLabelPruned = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection struct {
	// Indicates whether the feature of controlling the OPA injection scope is enabled. Valid values:
	//
	// *   `true`: The feature is enabled.
	// *   `false`: The feature is disabled.
	OPAScopeInjected *bool `json:"OPAScopeInjected,omitempty" xml:"OPAScopeInjected,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) SetOPAScopeInjected(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection {
	s.OPAScopeInjected = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit struct {
	// The maximum number of CPU cores that are available to the istio-init container.
	ResourceCPULimit *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
	// The maximum size of the memory that is available to the istio-init container.
	ResourceMemoryLimit *string `json:"ResourceMemoryLimit,omitempty" xml:"ResourceMemoryLimit,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) SetResourceCPULimit(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit {
	s.ResourceCPULimit = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) SetResourceMemoryLimit(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit {
	s.ResourceMemoryLimit = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest struct {
	// The number of CPU cores that are requested by the istio-init container.
	ResourceCPURequest *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
	// The size of the memory that is requested by the istio-init container.
	ResourceMemoryRequest *string `json:"ResourceMemoryRequest,omitempty" xml:"ResourceMemoryRequest,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) SetResourceCPURequest(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest {
	s.ResourceCPURequest = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) SetResourceMemoryRequest(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest {
	s.ResourceMemoryRequest = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport struct {
	// Indicates whether Gateway API is enabled. Valid values:
	//
	// *   `true`: Gateway API is enabled.
	// *   `false`: Gateway API is disabled.
	GatewayAPIEnabled *bool `json:"GatewayAPIEnabled,omitempty" xml:"GatewayAPIEnabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) SetGatewayAPIEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport {
	s.GatewayAPIEnabled = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali struct {
	// Indicates whether mesh topology is enabled. Mesh topology can be enabled only when Prometheus monitoring is enabled. If Prometheus monitoring is disabled, you must set this parameter to `false`. Valid values:
	//
	// *   `true`: Mesh topology is enabled.
	// *   `false`: Mesh topology is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The endpoint of the mesh topology service.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) SetUrl(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali {
	s.Url = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB struct {
	// The configurations of cross-region traffic distribution.
	//
	// >  Only one of `Failover` and Distribute parameters can be set. If you set the `Distribute` parameter, you cannot set the Failover parameter.
	Distribute map[string]interface{} `json:"Distribute,omitempty" xml:"Distribute,omitempty"`
	// Indicates whether cross-region load balancing is enabled. Valid values:
	//
	// *   `true`: Cross-region load balancing is enabled.
	// *   `false`: Cross-region load balancing is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The configurations of cross-region failover.
	//
	// >  Only one of Failover and `Distribute` parameters can be set. If you set the `Failover` parameter, you cannot set the `Distribute` parameter.
	Failover map[string]interface{} `json:"Failover,omitempty" xml:"Failover,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) SetDistribute(v map[string]interface{}) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB {
	s.Distribute = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) SetFailover(v map[string]interface{}) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB {
	s.Failover = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE struct {
	// Indicates whether MSE is enabled. Valid values:
	//
	// - `true`: MSE is enabled.
	// - `false`: MSE is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE {
	s.Enabled = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA struct {
	// Indicates whether the OPA plug-in is installed. Valid values:
	//
	// *   `true`: The OPA plug-in is installed.
	// *   `false`: The OPA plug-in is not installed.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of CPU cores that are available to the OPA proxy container.
	LimitCPU *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	// The maximum size of the memory that is available to the OPA proxy container.
	LimitMemory *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	// The level of the logs to be generated for OPA.
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The number of CPU cores that are requested by the OPA proxy container.
	RequestCPU *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The size of the memory that is requested by the OPA proxy container.
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) SetLimitCPU(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA {
	s.LimitCPU = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) SetLimitMemory(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA {
	s.LimitMemory = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) SetLogLevel(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA {
	s.LogLevel = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) SetRequestCPU(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA {
	s.RequestCPU = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) SetRequestMemory(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA {
	s.RequestMemory = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot struct {
	// The configurations of communication between external services and services in the mesh.
	ConfigSource *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource `json:"ConfigSource,omitempty" xml:"ConfigSource,omitempty" type:"Struct"`
	// The configurations of Pilot features.
	Feature *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature `json:"Feature,omitempty" xml:"Feature,omitempty" type:"Struct"`
	// Indicates whether HTTP/1.0 is supported. Valid values:
	//
	// *   `true`: HTTP/1.0 is supported.
	// *   `false`: HTTP/1.0 is not supported.
	Http10Enabled *bool `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
	// The sampling percentage of tracing analysis.
	TraceSampling *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) SetConfigSource(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot {
	s.ConfigSource = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) SetFeature(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot {
	s.Feature = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) SetHttp10Enabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot {
	s.Http10Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) SetTraceSampling(v float32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot {
	s.TraceSampling = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource struct {
	// Indicates whether communication is allowed between external services and services in the mesh. Valid values:
	//
	// *   `true`: The communication is allowed.
	// *   `false`: The communication is not allowed.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the Nacos instance that provides external service information.
	NacosID *string `json:"NacosID,omitempty" xml:"NacosID,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) SetNacosID(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource {
	s.NacosID = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature struct {
	// Indicates whether Secret Discovery Service (SDS) is enabled. Valid values:
	//
	// *   `true`: SDS is enabled.
	// *   `false`: SDS is disabled.
	EnableSDSServer *bool `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
	// Indicates whether gateway configuration filtering is enabled. Valid values:
	//
	// *   `true`: Gateway configuration filtering is enabled.
	// *   `false`: Gateway configuration filtering is disabled.
	FilterGatewayClusterConfig *bool `json:"FilterGatewayClusterConfig,omitempty" xml:"FilterGatewayClusterConfig,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) SetEnableSDSServer(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature {
	s.EnableSDSServer = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) SetFilterGatewayClusterConfig(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature {
	s.FilterGatewayClusterConfig = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus struct {
	// The endpoint of Prometheus monitoring. If you use a custom Prometheus instance, this parameter is populated by the system.
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// Indicates whether a custom Prometheus instance is used. Valid values:
	//
	// *   `true`: A custom Prometheus instance is used.
	// *   `false`: No custom Prometheus instance is used.
	UseExternal *bool `json:"UseExternal,omitempty" xml:"UseExternal,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) SetExternalUrl(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus {
	s.ExternalUrl = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) SetUseExternal(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus {
	s.UseExternal = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport struct {
	// Indicates whether Dubbo Filter is enabled. Valid values:
	//
	// *   `true`: Dubbo Filter is enabled.
	// *   `false`: Dubbo Filter is disabled.
	DubboFilterEnabled *bool `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	// Indicates whether MySQL Filter is enabled. Valid values:
	//
	// *   `true`: MySQL Filter is enabled.
	// *   `false`: MySQL Filter is disabled.
	MysqlFilterEnabled *bool `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	// Indicates whether Redis Filter is enabled. Valid values:
	//
	// *   `true`: Redis Filter is enabled.
	// *   `false`: Redis Filter is disabled.
	RedisFilterEnabled *bool `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
	// Indicates whether Thrift Filter is enabled. Valid values:
	//
	// *   `true`: Thrift Filter is enabled.
	// *   `false`: Thrift Filter is disabled.
	ThriftFilterEnabled *bool `json:"ThriftFilterEnabled,omitempty" xml:"ThriftFilterEnabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) SetDubboFilterEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport {
	s.DubboFilterEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) SetMysqlFilterEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport {
	s.MysqlFilterEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) SetRedisFilterEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport {
	s.RedisFilterEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) SetThriftFilterEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport {
	s.ThriftFilterEnabled = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy struct {
	// The path to the file that stores the access logs of sidecar proxies.
	AccessLogFile *string `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	// The format of the access logs of sidecar proxies.
	AccessLogFormat *string `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	// Indicates whether gRPC Access Log Service (ALS) for Envoy is enabled. Valid values:
	//
	// *   `true`: gRPC ALS for Envoy is enabled.
	// *   `false`: gRPC ALS for Envoy is disabled.
	AccessLogServiceEnabled *bool `json:"AccessLogServiceEnabled,omitempty" xml:"AccessLogServiceEnabled,omitempty"`
	// The endpoint of gRPC ALS for Envoy.
	AccessLogServiceHost *string `json:"AccessLogServiceHost,omitempty" xml:"AccessLogServiceHost,omitempty"`
	// The port of gRPC ALS for Envoy.
	AccessLogServicePort *int32 `json:"AccessLogServicePort,omitempty" xml:"AccessLogServicePort,omitempty"`
	// The trusted domain.
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// Indicates whether the Domain Name System (DNS) proxy feature is enabled. Valid values: Valid values:
	//
	// *   `true`: The DNS proxy feature is enabled.
	// *   `false`: The DNS proxy feature is disabled.
	EnableDNSProxying *bool `json:"EnableDNSProxying,omitempty" xml:"EnableDNSProxying,omitempty"`
	// The maximum number of CPU cores.
	LimitCPU *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	// The maximum size of the memory.
	LimitMemory *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	// The number of CPU cores that are requested.
	RequestCPU *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The size of the memory that is requested.
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetAccessLogFile(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.AccessLogFile = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetAccessLogFormat(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.AccessLogFormat = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetAccessLogServiceEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.AccessLogServiceEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetAccessLogServiceHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.AccessLogServiceHost = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetAccessLogServicePort(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.AccessLogServicePort = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetClusterDomain(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetEnableDNSProxying(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.EnableDNSProxying = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetLimitCPU(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.LimitCPU = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetLimitMemory(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.LimitMemory = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetRequestCPU(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.RequestCPU = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) SetRequestMemory(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	s.RequestMemory = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector struct {
	// Indicates whether automatic sidecar injection can be enabled by using pod annotations. Valid values:
	//
	// *   `true`: Automatic sidecar injection can be enabled by using pod annotations.
	// *   `false`: Automatic sidecar injection cannot be enabled by using pod annotations.
	AutoInjectionPolicyEnabled *bool `json:"AutoInjectionPolicyEnabled,omitempty" xml:"AutoInjectionPolicyEnabled,omitempty"`
	// Indicates whether automatic sidecar injection is enabled for all namespaces. Valid values:
	//
	// *   `true`: Automatic sidecar injection is enabled for all namespaces.
	// *   `false`: Automatic sidecar injection is not enabled for all namespaces.
	EnableNamespacesByDefault *bool `json:"EnableNamespacesByDefault,omitempty" xml:"EnableNamespacesByDefault,omitempty"`
	// The configurations of Container Network Interface (CNI).
	InitCNIConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration `json:"InitCNIConfiguration,omitempty" xml:"InitCNIConfiguration,omitempty" type:"Struct"`
	// The maximum number of CPU cores that are available to the sidecar injector pod.
	LimitCPU *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	// The maximum size of the memory that is available to the sidecar injector pod.
	LimitMemory *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	// The number of CPU cores that are requested by the sidecar injector pod.
	RequestCPU *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The size of the memory that is requested by the sidecar injector pod.
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// The number of component replicas that are used for sidecar injection. Default value: `1`.
	SidecarInjectorNum *int32 `json:"SidecarInjectorNum,omitempty" xml:"SidecarInjectorNum,omitempty"`
	// Other configurations of automatic sidecar injection, in the YAML format. For more information, see [Enable automatic sidecar injection by using multiple methods](~~186136~~).
	SidecarInjectorWebhookAsYaml *string `json:"SidecarInjectorWebhookAsYaml,omitempty" xml:"SidecarInjectorWebhookAsYaml,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetAutoInjectionPolicyEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.AutoInjectionPolicyEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetEnableNamespacesByDefault(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.EnableNamespacesByDefault = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetInitCNIConfiguration(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.InitCNIConfiguration = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetLimitCPU(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.LimitCPU = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetLimitMemory(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.LimitMemory = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetRequestCPU(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.RequestCPU = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetRequestMemory(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.RequestMemory = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetSidecarInjectorNum(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.SidecarInjectorNum = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetSidecarInjectorWebhookAsYaml(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.SidecarInjectorWebhookAsYaml = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration struct {
	// Indicates whether the CNI plug-in is enabled. Valid values:
	//
	// *   `true`: The CNI plug-in is enabled.
	// *   `false`: The CNI plug-in is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The namespaces to exclude. The CNI plug-in ignores pods in the excluded namespaces.
	ExcludeNamespaces *string `json:"ExcludeNamespaces,omitempty" xml:"ExcludeNamespaces,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) SetExcludeNamespaces(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	s.ExcludeNamespaces = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment struct {
	// Indicates whether WebAssembly Filter is enabled. Valid values:
	//
	// *   `true`:WebAssembly Filter is enabled.
	// *   `false`: WebAssembly Filter is disabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment {
	s.Enabled = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork struct {
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the vSwitch.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) SetSecurityGroupId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) SetVSwitches(v []*string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork {
	s.VSwitches = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) SetVpcId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork {
	s.VpcId = &v
	return s
}

type DescribeServiceMeshDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshDetailResponse) SetStatusCode(v int32) *DescribeServiceMeshDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshDetailResponse) SetBody(v *DescribeServiceMeshDetailResponseBody) *DescribeServiceMeshDetailResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshKubeconfigRequest struct {
	// Specifies whether to query the kubeconfig file that is used for Internet access or internal network access.
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshKubeconfigRequest) SetPrivateIpAddress(v bool) *DescribeServiceMeshKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigRequest) SetServiceMeshId(v string) *DescribeServiceMeshKubeconfigRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshKubeconfigResponseBody struct {
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The content of the kubeconfig file of the cluster.
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshKubeconfigResponseBody) SetExpireTime(v string) *DescribeServiceMeshKubeconfigResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponseBody) SetKubeconfig(v string) *DescribeServiceMeshKubeconfigResponseBody {
	s.Kubeconfig = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponseBody) SetRequestId(v string) *DescribeServiceMeshKubeconfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceMeshKubeconfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponse) SetStatusCode(v int32) *DescribeServiceMeshKubeconfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponse) SetBody(v *DescribeServiceMeshKubeconfigResponseBody) *DescribeServiceMeshKubeconfigResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshLogsRequest struct {
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshLogsRequest) SetServiceMeshId(v string) *DescribeServiceMeshLogsRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshLogsResponseBody struct {
	// The details of the logs.
	Logs []*DescribeServiceMeshLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshLogsResponseBody) SetLogs(v []*DescribeServiceMeshLogsResponseBodyLogs) *DescribeServiceMeshLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeServiceMeshLogsResponseBody) SetRequestId(v string) *DescribeServiceMeshLogsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceMeshLogsResponseBodyLogs struct {
	// The point in time when the logs were generated.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The content of the logs.
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) SetCreationTime(v string) *DescribeServiceMeshLogsResponseBodyLogs {
	s.CreationTime = &v
	return s
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) SetLog(v string) *DescribeServiceMeshLogsResponseBodyLogs {
	s.Log = &v
	return s
}

func (s *DescribeServiceMeshLogsResponseBodyLogs) SetServiceMeshId(v string) *DescribeServiceMeshLogsResponseBodyLogs {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshLogsResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshLogsResponse) SetStatusCode(v int32) *DescribeServiceMeshLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshLogsResponse) SetBody(v *DescribeServiceMeshLogsResponseBody) *DescribeServiceMeshLogsResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshProxyStatusRequest struct {
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshProxyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshProxyStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshProxyStatusRequest) SetServiceMeshId(v string) *DescribeServiceMeshProxyStatusRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshProxyStatusResponseBody struct {
	// The status code. Valid values:
	//
	// `200`: The status code returned because the operation is successful.
	//
	// *   `403`: The status code returned because you are not authorized to perform this operation.
	// *   `503`: The status code returned because a backend server error occurs.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about the status of the proxies on the data plane.
	ProxyStatus []*DescribeServiceMeshProxyStatusResponseBodyProxyStatus `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeServiceMeshProxyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshProxyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetCode(v string) *DescribeServiceMeshProxyStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetMessage(v string) *DescribeServiceMeshProxyStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetProxyStatus(v []*DescribeServiceMeshProxyStatusResponseBodyProxyStatus) *DescribeServiceMeshProxyStatusResponseBody {
	s.ProxyStatus = v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetRequestId(v string) *DescribeServiceMeshProxyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBody) SetSuccess(v string) *DescribeServiceMeshProxyStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeServiceMeshProxyStatusResponseBodyProxyStatus struct {
	// The update status of the proxy. Valid values:
	//
	// *   `SYNCED`: The status of the proxy is updated.
	// *   `NOT SENT`: The status of the proxy is not updated.
	// *   `STALE (Never Acknowledged)`: Istiod has sent multiple requests to the Envoy proxy to update the status of the proxy but receives no response.
	// *   `STALE`: Istiod has sent a request to the Envoy proxy to update the status of the proxy but receives no response.
	ClusterSynced *string `json:"ClusterSynced,omitempty" xml:"ClusterSynced,omitempty"`
	// The percentage of the updated endpoints.
	EndpointPercent *string `json:"EndpointPercent,omitempty" xml:"EndpointPercent,omitempty"`
	// The update status of the endpoint. Valid values:
	//
	// *   `SYNCED`: The status of the endpoint is updated.
	// *   `NOT SENT`: The status of the endpoint is not updated.
	// *   `STALE (Never Acknowledged)`: Istiod has sent multiple requests to the Envoy proxy to update the status of the endpoint but receives no response.
	// *   `STALE`: Istiod has sent a request to the Envoy proxy to update the status of the endpoint but receives no response.
	EndpointSynced *string `json:"EndpointSynced,omitempty" xml:"EndpointSynced,omitempty"`
	// The version of Istiod.
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	// The update status of the listener. Valid values:
	//
	// *   `SYNCED`: The status of the listener is updated.
	// *   `NOT SENT`: The status of the listener is not updated.
	// *   `STALE (Never Acknowledged)`: Istiod has sent multiple requests to the Envoy proxy to update the status of the listener but receives no response.
	// *   `STALE`: Istiod has sent a request to the Envoy proxy to update the status of the listener but receives no response.
	ListenerSynced *string `json:"ListenerSynced,omitempty" xml:"ListenerSynced,omitempty"`
	// The ID of the proxy on the data plane.
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The version number of the proxy on the data plane.
	ProxyVersion *string `json:"ProxyVersion,omitempty" xml:"ProxyVersion,omitempty"`
	// The update status of the route. Valid values:
	//
	// *   `SYNCED`: The status of the route is updated.
	// *   `NOT SENT`: The status of the route is not updated.
	// *   `STALE (Never Acknowledged)`: Istiod has sent multiple requests to the Envoy proxy to update the status of the route but receives no response.
	// *   `STALE`: Istiod has sent a request to the Envoy proxy to update the status of the route but receives no response.
	RouteSynced *string `json:"RouteSynced,omitempty" xml:"RouteSynced,omitempty"`
}

func (s DescribeServiceMeshProxyStatusResponseBodyProxyStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshProxyStatusResponseBodyProxyStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetClusterSynced(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.ClusterSynced = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetEndpointPercent(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.EndpointPercent = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetEndpointSynced(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.EndpointSynced = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetIstioVersion(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.IstioVersion = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetListenerSynced(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.ListenerSynced = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetProxyId(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.ProxyId = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetProxyVersion(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.ProxyVersion = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponseBodyProxyStatus) SetRouteSynced(v string) *DescribeServiceMeshProxyStatusResponseBodyProxyStatus {
	s.RouteSynced = &v
	return s
}

type DescribeServiceMeshProxyStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshProxyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshProxyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshProxyStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshProxyStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshProxyStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponse) SetStatusCode(v int32) *DescribeServiceMeshProxyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponse) SetBody(v *DescribeServiceMeshProxyStatusResponseBody) *DescribeServiceMeshProxyStatusResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshUpgradeStatusRequest struct {
	// The ID of the request.
	AllIstioGatewayFullNames *string `json:"AllIstioGatewayFullNames,omitempty" xml:"AllIstioGatewayFullNames,omitempty"`
	// The fully qualified names of ingress gateways in the ASM instance. Separate multiple names with commas (,).
	GuestClusterIds *string `json:"GuestClusterIds,omitempty" xml:"GuestClusterIds,omitempty"`
	// The IDs of the clusters on the data plane of the ASM instance. Separate multiple clusters with commas (,).
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshUpgradeStatusRequest) SetAllIstioGatewayFullNames(v string) *DescribeServiceMeshUpgradeStatusRequest {
	s.AllIstioGatewayFullNames = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusRequest) SetGuestClusterIds(v string) *DescribeServiceMeshUpgradeStatusRequest {
	s.GuestClusterIds = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusRequest) SetServiceMeshId(v string) *DescribeServiceMeshUpgradeStatusRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshUpgradeStatusResponseBody struct {
	// The upgrade results.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of ingress gateways that are upgraded.
	UpgradeDetail *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail `json:"UpgradeDetail,omitempty" xml:"UpgradeDetail,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshUpgradeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshUpgradeStatusResponseBody) SetRequestId(v string) *DescribeServiceMeshUpgradeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBody) SetUpgradeDetail(v *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) *DescribeServiceMeshUpgradeStatusResponseBody {
	s.UpgradeDetail = v
	return s
}

type DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail struct {
	// The status of the ASM instance. Valid values:
	//
	// *   running: The instance is running.
	// *   `upgrading`: The instance is being upgraded.
	// *   `upgrading_failed`: The upgrade of the instance fails.
	FinishedGatewaysNum *int64                                            `json:"FinishedGatewaysNum,omitempty" xml:"FinishedGatewaysNum,omitempty"`
	GatewayStatusRecord map[string]*UpgradeDetailGatewayStatusRecordValue `json:"GatewayStatusRecord,omitempty" xml:"GatewayStatusRecord,omitempty"`
	// The total number of ingress gateways in the ASM instance.
	MeshStatus *string `json:"MeshStatus,omitempty" xml:"MeshStatus,omitempty"`
	// The information about the status of the ingress gateways.
	TotalGatewaysNum *int64 `json:"TotalGatewaysNum,omitempty" xml:"TotalGatewaysNum,omitempty"`
}

func (s DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) SetFinishedGatewaysNum(v int64) *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	s.FinishedGatewaysNum = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) SetGatewayStatusRecord(v map[string]*UpgradeDetailGatewayStatusRecordValue) *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	s.GatewayStatusRecord = v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) SetMeshStatus(v string) *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	s.MeshStatus = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) SetTotalGatewaysNum(v int64) *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	s.TotalGatewaysNum = &v
	return s
}

type DescribeServiceMeshUpgradeStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshUpgradeStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponse) SetStatusCode(v int32) *DescribeServiceMeshUpgradeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponse) SetBody(v *DescribeServiceMeshUpgradeStatusResponseBody) *DescribeServiceMeshUpgradeStatusResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshVMsRequest struct {
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshVMsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshVMsRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshVMsRequest) SetServiceMeshId(v string) *DescribeServiceMeshVMsRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshVMsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ECS instances that reside in the same VPC as the ASM instance.
	VMs []*DescribeServiceMeshVMsResponseBodyVMs `json:"VMs,omitempty" xml:"VMs,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshVMsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshVMsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshVMsResponseBody) SetRequestId(v string) *DescribeServiceMeshVMsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBody) SetVMs(v []*DescribeServiceMeshVMsResponseBodyVMs) *DescribeServiceMeshVMsResponseBody {
	s.VMs = v
	return s
}

type DescribeServiceMeshVMsResponseBodyVMs struct {
	// Indicates whether the ECS instance has labels.
	HasTag *bool `json:"HasTag,omitempty" xml:"HasTag,omitempty"`
	// The host name.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the ECS instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the ECS instance.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The region ID.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The security group to which the ECS instance belongs.
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The state of the ECS instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeServiceMeshVMsResponseBodyVMs) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshVMsResponseBodyVMs) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetHasTag(v bool) *DescribeServiceMeshVMsResponseBodyVMs {
	s.HasTag = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetHostName(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.HostName = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetInstanceId(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.InstanceId = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetIpAddress(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.IpAddress = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetRegion(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.Region = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetSecurityGroupIds(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.SecurityGroupIds = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetServiceMeshId(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshVMsResponseBodyVMs) SetStatus(v string) *DescribeServiceMeshVMsResponseBodyVMs {
	s.Status = &v
	return s
}

type DescribeServiceMeshVMsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshVMsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshVMsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshVMsResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshVMsResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshVMsResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshVMsResponse) SetStatusCode(v int32) *DescribeServiceMeshVMsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshVMsResponse) SetBody(v *DescribeServiceMeshVMsResponseBody) *DescribeServiceMeshVMsResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshesRequest struct {
	Tag []*DescribeServiceMeshesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesRequest) SetTag(v []*DescribeServiceMeshesRequestTag) *DescribeServiceMeshesRequest {
	s.Tag = v
	return s
}

type DescribeServiceMeshesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeServiceMeshesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesRequestTag) SetKey(v string) *DescribeServiceMeshesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeServiceMeshesRequestTag) SetValue(v string) *DescribeServiceMeshesRequestTag {
	s.Value = &v
	return s
}

type DescribeServiceMeshesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the ASM instances.
	ServiceMeshes []*DescribeServiceMeshesResponseBodyServiceMeshes `json:"ServiceMeshes,omitempty" xml:"ServiceMeshes,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBody) SetRequestId(v string) *DescribeServiceMeshesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshesResponseBody) SetServiceMeshes(v []*DescribeServiceMeshesResponseBodyServiceMeshes) *DescribeServiceMeshesResponseBody {
	s.ServiceMeshes = v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshes struct {
	// The edition of the ASM instance. Valid values:
	//
	// - `standard`: Standard Edition
	// - `enterprise`: Enterprise Edition
	// - `ultimate`: Ultimate Edition
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The information about the clusters.
	Clusters []*string `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// All endpoints of the ASM instance.
	Endpoints *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud service instance for which the ASM instance is created.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud service for which the ASM instance is created. Valid values:
	//
	// - `ackone`: The ASM instance is created for Alibaba Cloud Distributed Cloud Container Platform (ACK One).
	// - An empty value indicates that the ASM instance is created by the user.
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The basic information about the ASM instance.
	ServiceMeshInfo *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo `json:"ServiceMeshInfo,omitempty" xml:"ServiceMeshInfo,omitempty" type:"Struct"`
	// The specifications of the ASM instance.
	Spec *DescribeServiceMeshesResponseBodyServiceMeshesSpec  `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	Tag  []*DescribeServiceMeshesResponseBodyServiceMeshesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshes) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshes) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetClusterSpec(v string) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetClusters(v []*string) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.Clusters = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetEndpoints(v *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.Endpoints = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetOwnerId(v string) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.OwnerId = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetOwnerType(v string) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.OwnerType = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetServiceMeshInfo(v *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.ServiceMeshInfo = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetSpec(v *DescribeServiceMeshesResponseBodyServiceMeshesSpec) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.Spec = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetTag(v []*DescribeServiceMeshesResponseBodyServiceMeshesTag) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.Tag = v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesEndpoints struct {
	// The endpoint that is used to access the API server from the internal network.
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	// The endpoint that is used to access Istio Pilot from the internal network.
	IntranetPilotEndpoint *string `json:"IntranetPilotEndpoint,omitempty" xml:"IntranetPilotEndpoint,omitempty"`
	// The endpoint that is used to expose the API server to the Internet.
	PublicApiServerEndpoint *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
	// The endpoint that is used to expose Istio Pilot to the Internet.
	PublicPilotEndpoint *string `json:"PublicPilotEndpoint,omitempty" xml:"PublicPilotEndpoint,omitempty"`
	// The endpoint of the reverse tunnel.
	ReverseTunnelEndpoint *string `json:"ReverseTunnelEndpoint,omitempty" xml:"ReverseTunnelEndpoint,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) SetIntranetApiServerEndpoint(v string) *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints {
	s.IntranetApiServerEndpoint = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) SetIntranetPilotEndpoint(v string) *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints {
	s.IntranetPilotEndpoint = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) SetPublicApiServerEndpoint(v string) *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints {
	s.PublicApiServerEndpoint = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) SetPublicPilotEndpoint(v string) *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints {
	s.PublicPilotEndpoint = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) SetReverseTunnelEndpoint(v string) *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints {
	s.ReverseTunnelEndpoint = &v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo struct {
	// The point in time when the ASM instance was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message that is returned when the call failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the ASM instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The edition of the ASM instance before ASM is available for commercial use. Valid values:
	//
	// *   `Pro`: Professional Edition
	// *   `Default`: Standard Edition
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region in which the ASM instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the ASM instance.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The point in time when the ASM instance was last modified.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version number of the ASM instance.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetCreationTime(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.CreationTime = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetErrorMessage(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetName(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetProfile(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.Profile = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetRegionId(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetServiceMeshId(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetState(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.State = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetUpdateTime(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) SetVersion(v string) *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	s.Version = &v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpec struct {
	// The information about load balancing.
	LoadBalancer *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Struct"`
	// The configurations of the ASM instance.
	MeshConfig *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	// The network configurations of the ASM instance.
	Network *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpec) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpec) SetLoadBalancer(v *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) *DescribeServiceMeshesResponseBodyServiceMeshesSpec {
	s.LoadBalancer = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpec) SetMeshConfig(v *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) *DescribeServiceMeshesResponseBodyServiceMeshesSpec {
	s.MeshConfig = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpec) SetNetwork(v *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) *DescribeServiceMeshesResponseBodyServiceMeshesSpec {
	s.Network = v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer struct {
	// The ID of the SLB instance that is used when the API server is exposed to the Internet.
	ApiServerLoadbalancerId *string `json:"ApiServerLoadbalancerId,omitempty" xml:"ApiServerLoadbalancerId,omitempty"`
	// Indicates whether the API Server is exposed to the Internet. Valid values:
	//
	// *   `true`: The API server is exposed to the Internet.
	// *   `false`: The API server is not exposed to the Internet.
	ApiServerPublicEip *bool `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	// Indicates whether Istio Pilot is exposed to the Internet. Valid values:
	//
	// *   `true`: Istio Pilot is exposed to the Internet.
	// *   `false`: Istio Pilot is not exposed to the Internet.
	PilotPublicEip *bool `json:"PilotPublicEip,omitempty" xml:"PilotPublicEip,omitempty"`
	// The ID of the Server Load Balancer (SLB) instance that is used when Istio Pilot is exposed to the Internet.
	PilotPublicLoadbalancerId *string `json:"PilotPublicLoadbalancerId,omitempty" xml:"PilotPublicLoadbalancerId,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) SetApiServerLoadbalancerId(v string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer {
	s.ApiServerLoadbalancerId = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) SetApiServerPublicEip(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer {
	s.ApiServerPublicEip = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) SetPilotPublicEip(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer {
	s.PilotPublicEip = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) SetPilotPublicLoadbalancerId(v string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer {
	s.PilotPublicLoadbalancerId = &v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig struct {
	// Indicates whether the feature of routing traffic to the nearest instance is enabled. Valid values:
	//
	// *   `true`: The feature is enabled.
	// *   `false`: The feature is disabled.
	Mtls *bool `json:"Mtls,omitempty" xml:"Mtls,omitempty"`
	// The outbound traffic policy. Valid values:
	//
	// *   `ALLOW_ANY`: Outbound traffic to an external service is allowed.
	// *   `REGISTRY_ONLY`: Outbound traffic is allowed to only external services that are defined in the service registry of the ASM instance.
	OutboundTrafficPolicy *string `json:"OutboundTrafficPolicy,omitempty" xml:"OutboundTrafficPolicy,omitempty"`
	// The configurations of the control plane.
	Pilot *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot `json:"Pilot,omitempty" xml:"Pilot,omitempty" type:"Struct"`
	// The configurations of sidecar injection.
	SidecarInjector *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector `json:"SidecarInjector,omitempty" xml:"SidecarInjector,omitempty" type:"Struct"`
	// Indicates whether mutual Transport Layer Security (mTLS) is strictly enforced.
	StrictMtls *bool `json:"StrictMtls,omitempty" xml:"StrictMtls,omitempty"`
	// Indicates whether Prometheus monitoring is enabled. We recommend that you use Prometheus Service of Application Real-Time Monitoring Service (ARMS). Valid values:
	//
	// *   `true`: Prometheus monitoring is enabled.
	// *   `false`: Prometheus monitoring is disabled.
	Telemetry *bool `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	// Indicates whether the tracing feature is enabled. This feature can be enabled only after Tracing Analysis is activated. Valid values:
	//
	// *   `true`: The tracing feature is enabled.
	// *   `false`: The tracing feature is disabled.
	Tracing *bool `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) SetMtls(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	s.Mtls = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) SetOutboundTrafficPolicy(v string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	s.OutboundTrafficPolicy = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) SetPilot(v *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	s.Pilot = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) SetSidecarInjector(v *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	s.SidecarInjector = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) SetStrictMtls(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	s.StrictMtls = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) SetTelemetry(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	s.Telemetry = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) SetTracing(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	s.Tracing = &v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot struct {
	// Indicates whether the support for HTTP 1.0 is enabled. Valid values:
	//
	// *   `true`: The support for HTTP 1.0 is enabled.
	// *   `false`: The support for HTTP 1.0 is disabled.
	Http10Enabled *bool `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
	// The sampling percentage of tracing.
	TraceSampling *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) SetHttp10Enabled(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot {
	s.Http10Enabled = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) SetTraceSampling(v float32) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot {
	s.TraceSampling = &v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector struct {
	// Indicates whether automatic sidecar injection is enabled by using annotations.
	AutoInjectionPolicyEnabled *bool `json:"AutoInjectionPolicyEnabled,omitempty" xml:"AutoInjectionPolicyEnabled,omitempty"`
	// Indicates whether automatic sidecar injection is enabled for all namespaces. Valid values:
	//
	// *   `true`: Automatic sidecar injection is enabled for all namespaces.
	// *   `false`: Automatic sidecar injection is disabled for all namespaces.
	EnableNamespacesByDefault *bool `json:"EnableNamespacesByDefault,omitempty" xml:"EnableNamespacesByDefault,omitempty"`
	// The initial configurations of Container Network Interface (CNI).
	InitCNIConfiguration *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration `json:"InitCNIConfiguration,omitempty" xml:"InitCNIConfiguration,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) SetAutoInjectionPolicyEnabled(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector {
	s.AutoInjectionPolicyEnabled = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) SetEnableNamespacesByDefault(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector {
	s.EnableNamespacesByDefault = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) SetInitCNIConfiguration(v *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector {
	s.InitCNIConfiguration = v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration struct {
	// Indicates whether elevated privileges are required for the istio-init container when you perform traffic redirection for the istio-proxy container. Valid values:
	//
	// *   `true`: Elevated privileges are required for the istio-init container.
	// *   `false`: Elevated privileges are not required for the istio-init container.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The namespace for which sidecar injection is disabled.
	ExcludeNamespaces *string `json:"ExcludeNamespaces,omitempty" xml:"ExcludeNamespaces,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) SetEnabled(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) SetExcludeNamespaces(v string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	s.ExcludeNamespaces = &v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork struct {
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of vSwitches.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) SetSecurityGroupId(v string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) SetVSwitches(v []*string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork {
	s.VSwitches = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) SetVpcId(v string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork {
	s.VpcId = &v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesTag) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesTag) SetKey(v string) *DescribeServiceMeshesResponseBodyServiceMeshesTag {
	s.Key = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesTag) SetValue(v string) *DescribeServiceMeshesResponseBodyServiceMeshesTag {
	s.Value = &v
	return s
}

type DescribeServiceMeshesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMeshesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMeshesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshesResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshesResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshesResponse) SetStatusCode(v int32) *DescribeServiceMeshesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshesResponse) SetBody(v *DescribeServiceMeshesResponseBody) *DescribeServiceMeshesResponse {
	s.Body = v
	return s
}

type DescribeUpgradeVersionRequest struct {
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeUpgradeVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpgradeVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeVersionRequest) SetServiceMeshId(v string) *DescribeUpgradeVersionRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeUpgradeVersionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version information.
	Version *DescribeUpgradeVersionResponseBodyVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
}

func (s DescribeUpgradeVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpgradeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeVersionResponseBody) SetRequestId(v string) *DescribeUpgradeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpgradeVersionResponseBody) SetVersion(v *DescribeUpgradeVersionResponseBodyVersion) *DescribeUpgradeVersionResponseBody {
	s.Version = v
	return s
}

type DescribeUpgradeVersionResponseBodyVersion struct {
	// The version of the ASM instance.
	IstioOperatorVersion *string `json:"IstioOperatorVersion,omitempty" xml:"IstioOperatorVersion,omitempty"`
	// The Istio version.
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	// The Kubernetes version.
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
}

func (s DescribeUpgradeVersionResponseBodyVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpgradeVersionResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeVersionResponseBodyVersion) SetIstioOperatorVersion(v string) *DescribeUpgradeVersionResponseBodyVersion {
	s.IstioOperatorVersion = &v
	return s
}

func (s *DescribeUpgradeVersionResponseBodyVersion) SetIstioVersion(v string) *DescribeUpgradeVersionResponseBodyVersion {
	s.IstioVersion = &v
	return s
}

func (s *DescribeUpgradeVersionResponseBodyVersion) SetKubernetesVersion(v string) *DescribeUpgradeVersionResponseBodyVersion {
	s.KubernetesVersion = &v
	return s
}

type DescribeUpgradeVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUpgradeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUpgradeVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpgradeVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeVersionResponse) SetHeaders(v map[string]*string) *DescribeUpgradeVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpgradeVersionResponse) SetStatusCode(v int32) *DescribeUpgradeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpgradeVersionResponse) SetBody(v *DescribeUpgradeVersionResponseBody) *DescribeUpgradeVersionResponse {
	s.Body = v
	return s
}

type DescribeUserPermissionsRequest struct {
	// The ID of the RAM user or RAM role.
	SubAccountUserId *string `json:"SubAccountUserId,omitempty" xml:"SubAccountUserId,omitempty"`
}

func (s DescribeUserPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsRequest) SetSubAccountUserId(v string) *DescribeUserPermissionsRequest {
	s.SubAccountUserId = &v
	return s
}

type DescribeUserPermissionsResponseBody struct {
	// The permissions that are granted to an entity.
	Permissions []*DescribeUserPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponseBody) SetPermissions(v []*DescribeUserPermissionsResponseBodyPermissions) *DescribeUserPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *DescribeUserPermissionsResponseBody) SetRequestId(v string) *DescribeUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserPermissionsResponseBodyPermissions struct {
	// The entity to which the permissions are granted. A value of `true` indicates that the permissions are granted to a RAM user. A value of `false` indicates that the permissions are granted to a RAM role.
	IsRamRole *string `json:"IsRamRole,omitempty" xml:"IsRamRole,omitempty"`
	// This parameter is required by the system. The return value is fixed to `0`.
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The ID of the ASM instance.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required by the system. The return value is fixed to `cluster`.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The name of the permissions. Valid values:
	//
	// *   `istio-admin`: the permissions of Alibaba Cloud Service Mesh (ASM) administrators.
	// *   `istio-ops`: the permissions of ASM restricted users.
	// *   `istio-readonly`: the read-only permissions.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// This parameter is required by the system. The return value is fixed to `custom`.
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeUserPermissionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetIsRamRole(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.IsRamRole = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetParentId(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.ParentId = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetResourceId(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.ResourceId = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetResourceType(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetRoleName(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.RoleName = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetRoleType(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.RoleType = &v
	return s
}

type DescribeUserPermissionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponse) SetHeaders(v map[string]*string) *DescribeUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserPermissionsResponse) SetStatusCode(v int32) *DescribeUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserPermissionsResponse) SetBody(v *DescribeUserPermissionsResponseBody) *DescribeUserPermissionsResponse {
	s.Body = v
	return s
}

type DescribeUsersWithPermissionsRequest struct {
	// Specifies whether to query the IDs of all RAM users or RAM roles to which an RBAC role is assigned. Valid values:
	//
	// *   `SubUser`: Query the IDs of all RAM users to which an RBAC role is assigned.
	// *   `RamRole`: Query the IDs of all RAM roles to which an RBAC role is assigned.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeUsersWithPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersWithPermissionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersWithPermissionsRequest) SetUserType(v string) *DescribeUsersWithPermissionsRequest {
	s.UserType = &v
	return s
}

type DescribeUsersWithPermissionsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the RAM users or RAM roles to which an RBAC role is assigned.
	UIDs []*string `json:"UIDs,omitempty" xml:"UIDs,omitempty" type:"Repeated"`
}

func (s DescribeUsersWithPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersWithPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersWithPermissionsResponseBody) SetRequestId(v string) *DescribeUsersWithPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersWithPermissionsResponseBody) SetUIDs(v []*string) *DescribeUsersWithPermissionsResponseBody {
	s.UIDs = v
	return s
}

type DescribeUsersWithPermissionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUsersWithPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUsersWithPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersWithPermissionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersWithPermissionsResponse) SetHeaders(v map[string]*string) *DescribeUsersWithPermissionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersWithPermissionsResponse) SetStatusCode(v int32) *DescribeUsersWithPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersWithPermissionsResponse) SetBody(v *DescribeUsersWithPermissionsResponseBody) *DescribeUsersWithPermissionsResponse {
	s.Body = v
	return s
}

type DescribeVMsInServiceMeshRequest struct {
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeVMsInServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVMsInServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *DescribeVMsInServiceMeshRequest) SetServiceMeshId(v string) *DescribeVMsInServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeVMsInServiceMeshResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VMs that are added to the ASM instance.
	VMs []*DescribeVMsInServiceMeshResponseBodyVMs `json:"VMs,omitempty" xml:"VMs,omitempty" type:"Repeated"`
}

func (s DescribeVMsInServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVMsInServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVMsInServiceMeshResponseBody) SetRequestId(v string) *DescribeVMsInServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBody) SetVMs(v []*DescribeVMsInServiceMeshResponseBodyVMs) *DescribeVMsInServiceMeshResponseBody {
	s.VMs = v
	return s
}

type DescribeVMsInServiceMeshResponseBodyVMs struct {
	// Indicates whether the ECS instance has labels.
	HasTag *bool `json:"HasTag,omitempty" xml:"HasTag,omitempty"`
	// The host name.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the ECS instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the ECS instance.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The region ID.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The security group to which the ECS instance belongs.
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	// The state of the ECS instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVMsInServiceMeshResponseBodyVMs) String() string {
	return tea.Prettify(s)
}

func (s DescribeVMsInServiceMeshResponseBodyVMs) GoString() string {
	return s.String()
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetHasTag(v bool) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.HasTag = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetHostName(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.HostName = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetInstanceId(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.InstanceId = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetIpAddress(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.IpAddress = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetRegion(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.Region = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetSecurityGroupIds(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.SecurityGroupIds = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponseBodyVMs) SetStatus(v string) *DescribeVMsInServiceMeshResponseBodyVMs {
	s.Status = &v
	return s
}

type DescribeVMsInServiceMeshResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVMsInServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVMsInServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVMsInServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *DescribeVMsInServiceMeshResponse) SetHeaders(v map[string]*string) *DescribeVMsInServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *DescribeVMsInServiceMeshResponse) SetStatusCode(v int32) *DescribeVMsInServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponse) SetBody(v *DescribeVMsInServiceMeshResponseBody) *DescribeVMsInServiceMeshResponse {
	s.Body = v
	return s
}

type DescribeVSwitchesRequest struct {
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) SetRegionId(v string) *DescribeVSwitchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVpcId(v string) *DescribeVSwitchesRequest {
	s.VpcId = &v
	return s
}

type DescribeVSwitchesResponseBody struct {
	// The maximum number of entries returned.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vSwitches that are deployed in the VPC in the region. This parameter is optional and is not returned by default.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The available vSwitches.
	VSwitches []*DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) SetMaxResults(v int32) *DescribeVSwitchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetNextToken(v string) *DescribeVSwitchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTotalCount(v int32) *DescribeVSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitches(v []*DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

type DescribeVSwitchesResponseBodyVSwitches struct {
	// Indicates whether the vSwitch is the default vSwitch. Valid values:
	//
	// *   `true`
	// *   `false`
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The state of the vSwitch. Valid values:
	//
	// *   `Pending`: The vSwitch is being configured.
	// *   `Available`: The vSwitch is available.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the VPC to which the vSwitch belongs.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 交换机所属的可用区。
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetIsDefault(v bool) *DescribeVSwitchesResponseBodyVSwitches {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVpcId(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetZoneId(v string) *DescribeVSwitchesResponseBodyVSwitches {
	s.ZoneId = &v
	return s
}

type DescribeVSwitchesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVSwitchesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponse) SetHeaders(v map[string]*string) *DescribeVSwitchesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVSwitchesResponse) SetStatusCode(v int32) *DescribeVSwitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVSwitchesResponse) SetBody(v *DescribeVSwitchesResponseBody) *DescribeVSwitchesResponse {
	s.Body = v
	return s
}

type DescribeVersionsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The available ASM versions.
	VersionInfo []*DescribeVersionsResponseBodyVersionInfo `json:"VersionInfo,omitempty" xml:"VersionInfo,omitempty" type:"Repeated"`
}

func (s DescribeVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionsResponseBody) SetRequestId(v string) *DescribeVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVersionsResponseBody) SetVersionInfo(v []*DescribeVersionsResponseBodyVersionInfo) *DescribeVersionsResponseBody {
	s.VersionInfo = v
	return s
}

type DescribeVersionsResponseBodyVersionInfo struct {
	// The edition of the ASM instance. Valid values:
	//
	// *   `Default`: Standard Edition
	// *   `Pro`: Professional Edition that is commercially released
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The ASM versions available for the ASM instance of the current edition.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeVersionsResponseBodyVersionInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionsResponseBodyVersionInfo) GoString() string {
	return s.String()
}

func (s *DescribeVersionsResponseBodyVersionInfo) SetEdition(v string) *DescribeVersionsResponseBodyVersionInfo {
	s.Edition = &v
	return s
}

func (s *DescribeVersionsResponseBodyVersionInfo) SetVersions(v []*string) *DescribeVersionsResponseBodyVersionInfo {
	s.Versions = v
	return s
}

type DescribeVersionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVersionsResponse) SetHeaders(v map[string]*string) *DescribeVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVersionsResponse) SetStatusCode(v int32) *DescribeVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVersionsResponse) SetBody(v *DescribeVersionsResponseBody) *DescribeVersionsResponse {
	s.Body = v
	return s
}

type DescribeVpcsRequest struct {
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeVpcsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcsRequest) SetRegionId(v string) *DescribeVpcsRequest {
	s.RegionId = &v
	return s
}

type DescribeVpcsResponseBody struct {
	// The maximum number of entries returned on a single page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, it indicates that you have retrieved all the data.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. By default, this parameter is not returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of VPCs that are available in the specified region.
	Vpcs []*DescribeVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBody) SetMaxResults(v int32) *DescribeVpcsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetNextToken(v string) *DescribeVpcsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetRequestId(v string) *DescribeVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetTotalCount(v int32) *DescribeVpcsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetVpcs(v []*DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody {
	s.Vpcs = v
	return s
}

type DescribeVpcsResponseBodyVpcs struct {
	// Indicates whether the VPC is the default VPC in the specified region. Valid values:
	//
	// *   `true`: yes
	// *   `false`: no
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The status of the VPC. Valid values:
	//
	// *   `Pending`: The VPC is being configured.
	// *   `Available`: The VPC is available for use.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcs) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetStatus(v string) *DescribeVpcsResponseBodyVpcs {
	s.Status = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpcId(v string) *DescribeVpcsResponseBodyVpcs {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpcName(v string) *DescribeVpcsResponseBodyVpcs {
	s.VpcName = &v
	return s
}

type DescribeVpcsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponse) SetHeaders(v map[string]*string) *DescribeVpcsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcsResponse) SetStatusCode(v int32) *DescribeVpcsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcsResponse) SetBody(v *DescribeVpcsResponseBody) *DescribeVpcsResponse {
	s.Body = v
	return s
}

type GetCaCertRequest struct {
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetCaCertRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCaCertRequest) GoString() string {
	return s.String()
}

func (s *GetCaCertRequest) SetServiceMeshId(v string) *GetCaCertRequest {
	s.ServiceMeshId = &v
	return s
}

type GetCaCertResponseBody struct {
	// The Base64-encoded content of the CA certificate.
	CaCert *string `json:"CaCert,omitempty" xml:"CaCert,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCaCertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCaCertResponseBody) GoString() string {
	return s.String()
}

func (s *GetCaCertResponseBody) SetCaCert(v string) *GetCaCertResponseBody {
	s.CaCert = &v
	return s
}

func (s *GetCaCertResponseBody) SetRequestId(v string) *GetCaCertResponseBody {
	s.RequestId = &v
	return s
}

type GetCaCertResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCaCertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCaCertResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCaCertResponse) GoString() string {
	return s.String()
}

func (s *GetCaCertResponse) SetHeaders(v map[string]*string) *GetCaCertResponse {
	s.Headers = v
	return s
}

func (s *GetCaCertResponse) SetStatusCode(v int32) *GetCaCertResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCaCertResponse) SetBody(v *GetCaCertResponseBody) *GetCaCertResponse {
	s.Body = v
	return s
}

type GetDeploymentBySelectorRequest struct {
	// The name of the cluster.
	GuestCluster *string `json:"GuestCluster,omitempty" xml:"GuestCluster,omitempty"`
	// The label selector information.
	LabelSelector map[string]*string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// The maximum number of returned data entries.
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The marker of data queried last time.
	Mark *string `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// The namespace.
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetDeploymentBySelectorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentBySelectorRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentBySelectorRequest) SetGuestCluster(v string) *GetDeploymentBySelectorRequest {
	s.GuestCluster = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetLabelSelector(v map[string]*string) *GetDeploymentBySelectorRequest {
	s.LabelSelector = v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetLimit(v int64) *GetDeploymentBySelectorRequest {
	s.Limit = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetMark(v string) *GetDeploymentBySelectorRequest {
	s.Mark = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetNameSpace(v string) *GetDeploymentBySelectorRequest {
	s.NameSpace = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetServiceMeshId(v string) *GetDeploymentBySelectorRequest {
	s.ServiceMeshId = &v
	return s
}

type GetDeploymentBySelectorShrinkRequest struct {
	// The name of the cluster.
	GuestCluster *string `json:"GuestCluster,omitempty" xml:"GuestCluster,omitempty"`
	// The label selector information.
	LabelSelectorShrink *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// The maximum number of returned data entries.
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The marker of data queried last time.
	Mark *string `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// The namespace.
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetDeploymentBySelectorShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentBySelectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentBySelectorShrinkRequest) SetGuestCluster(v string) *GetDeploymentBySelectorShrinkRequest {
	s.GuestCluster = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetLabelSelectorShrink(v string) *GetDeploymentBySelectorShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetLimit(v int64) *GetDeploymentBySelectorShrinkRequest {
	s.Limit = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetMark(v string) *GetDeploymentBySelectorShrinkRequest {
	s.Mark = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetNameSpace(v string) *GetDeploymentBySelectorShrinkRequest {
	s.NameSpace = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetServiceMeshId(v string) *GetDeploymentBySelectorShrinkRequest {
	s.ServiceMeshId = &v
	return s
}

type GetDeploymentBySelectorResponseBody struct {
	// The queried workloads.
	DeploymentNameList [][]byte `json:"DeploymentNameList,omitempty" xml:"DeploymentNameList,omitempty" type:"Repeated"`
	// The end-of-data marker.
	Mark *string `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeploymentBySelectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentBySelectorResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentBySelectorResponseBody) SetDeploymentNameList(v [][]byte) *GetDeploymentBySelectorResponseBody {
	s.DeploymentNameList = v
	return s
}

func (s *GetDeploymentBySelectorResponseBody) SetMark(v string) *GetDeploymentBySelectorResponseBody {
	s.Mark = &v
	return s
}

func (s *GetDeploymentBySelectorResponseBody) SetRequestId(v string) *GetDeploymentBySelectorResponseBody {
	s.RequestId = &v
	return s
}

type GetDeploymentBySelectorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeploymentBySelectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeploymentBySelectorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentBySelectorResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentBySelectorResponse) SetHeaders(v map[string]*string) *GetDeploymentBySelectorResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentBySelectorResponse) SetStatusCode(v int32) *GetDeploymentBySelectorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentBySelectorResponse) SetBody(v *GetDeploymentBySelectorResponseBody) *GetDeploymentBySelectorResponse {
	s.Body = v
	return s
}

type GetGrafanaDashboardUrlRequest struct {
	// The ID of the Container Service for Kubernetes (ACK) or serverless Kubernetes (ASK) cluster.
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The name of the dashboard.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetGrafanaDashboardUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGrafanaDashboardUrlRequest) GoString() string {
	return s.String()
}

func (s *GetGrafanaDashboardUrlRequest) SetK8sClusterId(v string) *GetGrafanaDashboardUrlRequest {
	s.K8sClusterId = &v
	return s
}

func (s *GetGrafanaDashboardUrlRequest) SetServiceMeshId(v string) *GetGrafanaDashboardUrlRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetGrafanaDashboardUrlRequest) SetTitle(v string) *GetGrafanaDashboardUrlRequest {
	s.Title = &v
	return s
}

type GetGrafanaDashboardUrlResponseBody struct {
	// The information about the dashboard.
	Dashboards []*GetGrafanaDashboardUrlResponseBodyDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGrafanaDashboardUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGrafanaDashboardUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetGrafanaDashboardUrlResponseBody) SetDashboards(v []*GetGrafanaDashboardUrlResponseBodyDashboards) *GetGrafanaDashboardUrlResponseBody {
	s.Dashboards = v
	return s
}

func (s *GetGrafanaDashboardUrlResponseBody) SetRequestId(v string) *GetGrafanaDashboardUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetGrafanaDashboardUrlResponseBodyDashboards struct {
	// The name of the dashboard.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the dashboard.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetGrafanaDashboardUrlResponseBodyDashboards) String() string {
	return tea.Prettify(s)
}

func (s GetGrafanaDashboardUrlResponseBodyDashboards) GoString() string {
	return s.String()
}

func (s *GetGrafanaDashboardUrlResponseBodyDashboards) SetTitle(v string) *GetGrafanaDashboardUrlResponseBodyDashboards {
	s.Title = &v
	return s
}

func (s *GetGrafanaDashboardUrlResponseBodyDashboards) SetUrl(v string) *GetGrafanaDashboardUrlResponseBodyDashboards {
	s.Url = &v
	return s
}

type GetGrafanaDashboardUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGrafanaDashboardUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGrafanaDashboardUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGrafanaDashboardUrlResponse) GoString() string {
	return s.String()
}

func (s *GetGrafanaDashboardUrlResponse) SetHeaders(v map[string]*string) *GetGrafanaDashboardUrlResponse {
	s.Headers = v
	return s
}

func (s *GetGrafanaDashboardUrlResponse) SetStatusCode(v int32) *GetGrafanaDashboardUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGrafanaDashboardUrlResponse) SetBody(v *GetGrafanaDashboardUrlResponseBody) *GetGrafanaDashboardUrlResponse {
	s.Body = v
	return s
}

type GetRegisteredServiceEndpointsRequest struct {
	// The name of the registered service.
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The type of the registered service. Valid values:
	//
	// *   `ServiceEntry`: indicates that the service is registered by creating a service entry.
	// *   `Kubernetes`: indicates that the service is registered on a Kubernetes cluster on the data plane.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of clusters in the ASM instance. Separate multiple cluster IDs with commas (,).
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the namespace.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The endpoints of the registered service.
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetRegisteredServiceEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceEndpointsRequest) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsRequest) SetClusterIds(v string) *GetRegisteredServiceEndpointsRequest {
	s.ClusterIds = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) SetName(v string) *GetRegisteredServiceEndpointsRequest {
	s.Name = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) SetNamespace(v string) *GetRegisteredServiceEndpointsRequest {
	s.Namespace = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) SetServiceMeshId(v string) *GetRegisteredServiceEndpointsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetRegisteredServiceEndpointsRequest) SetServiceType(v string) *GetRegisteredServiceEndpointsRequest {
	s.ServiceType = &v
	return s
}

type GetRegisteredServiceEndpointsResponseBody struct {
	// The name of the registered service.
	EndPointSlice *GetRegisteredServiceEndpointsResponseBodyEndPointSlice `json:"EndPointSlice,omitempty" xml:"EndPointSlice,omitempty" type:"Struct"`
	RequestId     *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IP address of the registered service.
	ServiceEndpoints []*GetRegisteredServiceEndpointsResponseBodyServiceEndpoints `json:"ServiceEndpoints,omitempty" xml:"ServiceEndpoints,omitempty" type:"Repeated"`
}

func (s GetRegisteredServiceEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponseBody) SetEndPointSlice(v *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) *GetRegisteredServiceEndpointsResponseBody {
	s.EndPointSlice = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBody) SetRequestId(v string) *GetRegisteredServiceEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBody) SetServiceEndpoints(v []*GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) *GetRegisteredServiceEndpointsResponseBody {
	s.ServiceEndpoints = v
	return s
}

type GetRegisteredServiceEndpointsResponseBodyEndPointSlice struct {
	// The name of the pod.
	EndpointsDetails []*GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails `json:"EndpointsDetails,omitempty" xml:"EndpointsDetails,omitempty" type:"Repeated"`
	// The details of the endpoint of the registered service.
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The location of the registered service. Valid values:
	//
	// *   `MESH_INTERNAL`: The service is deployed inside the ASM instance.
	// *   `MESH_EXTERNAL`: The service is deployed outside the ASM instance.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the namespace.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetRegisteredServiceEndpointsResponseBodyEndPointSlice) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponseBodyEndPointSlice) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) SetEndpointsDetails(v []*GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	s.EndpointsDetails = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) SetLocation(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	s.Location = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) SetNamespace(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	s.Namespace = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSlice) SetServiceName(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSlice {
	s.ServiceName = &v
	return s
}

type GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails struct {
	// The port of the registered service.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the region in which the registered service resides.
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The IP address of the registered service.
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// The host name of the registered service.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// Indicates whether sidecar proxies are injected. Valid values:
	//
	// *   `true`: yes
	// *   `false`: no
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the request.
	SidecarInjected *bool `json:"SidecarInjected,omitempty" xml:"SidecarInjected,omitempty"`
}

func (s GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetAddress(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.Address = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetHostname(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.Hostname = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetPodName(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.PodName = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetPorts(v []*int32) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.Ports = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetRegion(v string) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.Region = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails) SetSidecarInjected(v bool) *GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails {
	s.SidecarInjected = &v
	return s
}

type GetRegisteredServiceEndpointsResponseBodyServiceEndpoints struct {
	// The ID of the cluster on the data plane.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The details of the endpoints of the registered service.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) SetAddress(v string) *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints {
	s.Address = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints) SetClusterId(v string) *GetRegisteredServiceEndpointsResponseBodyServiceEndpoints {
	s.ClusterId = &v
	return s
}

type GetRegisteredServiceEndpointsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRegisteredServiceEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegisteredServiceEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponse) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponse) SetHeaders(v map[string]*string) *GetRegisteredServiceEndpointsResponse {
	s.Headers = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponse) SetStatusCode(v int32) *GetRegisteredServiceEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponse) SetBody(v *GetRegisteredServiceEndpointsResponseBody) *GetRegisteredServiceEndpointsResponse {
	s.Body = v
	return s
}

type GetRegisteredServiceNamespacesRequest struct {
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetRegisteredServiceNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceNamespacesRequest) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceNamespacesRequest) SetServiceMeshId(v string) *GetRegisteredServiceNamespacesRequest {
	s.ServiceMeshId = &v
	return s
}

type GetRegisteredServiceNamespacesResponseBody struct {
	// The names of the queried namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRegisteredServiceNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceNamespacesResponseBody) SetNamespaces(v []*string) *GetRegisteredServiceNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *GetRegisteredServiceNamespacesResponseBody) SetRequestId(v string) *GetRegisteredServiceNamespacesResponseBody {
	s.RequestId = &v
	return s
}

type GetRegisteredServiceNamespacesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRegisteredServiceNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegisteredServiceNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegisteredServiceNamespacesResponse) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceNamespacesResponse) SetHeaders(v map[string]*string) *GetRegisteredServiceNamespacesResponse {
	s.Headers = v
	return s
}

func (s *GetRegisteredServiceNamespacesResponse) SetStatusCode(v int32) *GetRegisteredServiceNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegisteredServiceNamespacesResponse) SetBody(v *GetRegisteredServiceNamespacesResponseBody) *GetRegisteredServiceNamespacesResponse {
	s.Body = v
	return s
}

type GetSwimLaneDetailRequest struct {
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The name of the lane.
	SwimLaneName *string `json:"SwimLaneName,omitempty" xml:"SwimLaneName,omitempty"`
}

func (s GetSwimLaneDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSwimLaneDetailRequest) SetGroupName(v string) *GetSwimLaneDetailRequest {
	s.GroupName = &v
	return s
}

func (s *GetSwimLaneDetailRequest) SetServiceMeshId(v string) *GetSwimLaneDetailRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetSwimLaneDetailRequest) SetSwimLaneName(v string) *GetSwimLaneDetailRequest {
	s.SwimLaneName = &v
	return s
}

type GetSwimLaneDetailResponseBody struct {
	// The traffic routing rule that routes traffic to the lane by using the ingress gateway. The traffic routing rule contains one or more custom routes.
	IngressRule *string `json:"IngressRule,omitempty" xml:"IngressRule,omitempty"`
	// This parameter is deprecated.
	IngressService *string `json:"IngressService,omitempty" xml:"IngressService,omitempty"`
	// Fixed value: **ASM_TRAFFIC_TAG**.
	LabelSelectorKey *string `json:"LabelSelectorKey,omitempty" xml:"LabelSelectorKey,omitempty"`
	// The value of ASM_TRAFFIC_TAG.
	LabelSelectorValue *string `json:"LabelSelectorValue,omitempty" xml:"LabelSelectorValue,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of services associated with the lane.
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
}

func (s GetSwimLaneDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwimLaneDetailResponseBody) SetIngressRule(v string) *GetSwimLaneDetailResponseBody {
	s.IngressRule = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetIngressService(v string) *GetSwimLaneDetailResponseBody {
	s.IngressService = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetLabelSelectorKey(v string) *GetSwimLaneDetailResponseBody {
	s.LabelSelectorKey = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetLabelSelectorValue(v string) *GetSwimLaneDetailResponseBody {
	s.LabelSelectorValue = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetRequestId(v string) *GetSwimLaneDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetServicesList(v string) *GetSwimLaneDetailResponseBody {
	s.ServicesList = &v
	return s
}

type GetSwimLaneDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSwimLaneDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSwimLaneDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSwimLaneDetailResponse) SetHeaders(v map[string]*string) *GetSwimLaneDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSwimLaneDetailResponse) SetStatusCode(v int32) *GetSwimLaneDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSwimLaneDetailResponse) SetBody(v *GetSwimLaneDetailResponseBody) *GetSwimLaneDetailResponse {
	s.Body = v
	return s
}

type GetSwimLaneGroupListRequest struct {
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetSwimLaneGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneGroupListRequest) GoString() string {
	return s.String()
}

func (s *GetSwimLaneGroupListRequest) SetServiceMeshId(v string) *GetSwimLaneGroupListRequest {
	s.ServiceMeshId = &v
	return s
}

type GetSwimLaneGroupListResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the lane groups.
	SwimLaneGroupList []*GetSwimLaneGroupListResponseBodySwimLaneGroupList `json:"SwimLaneGroupList,omitempty" xml:"SwimLaneGroupList,omitempty" type:"Repeated"`
}

func (s GetSwimLaneGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwimLaneGroupListResponseBody) SetRequestId(v string) *GetSwimLaneGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBody) SetSwimLaneGroupList(v []*GetSwimLaneGroupListResponseBodySwimLaneGroupList) *GetSwimLaneGroupListResponseBody {
	s.SwimLaneGroupList = v
	return s
}

type GetSwimLaneGroupListResponseBodySwimLaneGroupList struct {
	FallbackTarget *string `json:"FallbackTarget,omitempty" xml:"FallbackTarget,omitempty"`
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the ingress gateway.
	IngressGatewayName *string `json:"IngressGatewayName,omitempty" xml:"IngressGatewayName,omitempty"`
	// The ingress type. Traffic routing rules can be configured only in an ingress gateway.
	IngressType  *string `json:"IngressType,omitempty" xml:"IngressType,omitempty"`
	IsPermissive *bool   `json:"IsPermissive,omitempty" xml:"IsPermissive,omitempty"`
	RouteHeader  *string `json:"RouteHeader,omitempty" xml:"RouteHeader,omitempty"`
	// A list of services associated with the lane group.
	ServiceList    *string `json:"ServiceList,omitempty" xml:"ServiceList,omitempty"`
	SwimLaneLabels *string `json:"SwimLaneLabels,omitempty" xml:"SwimLaneLabels,omitempty"`
	TraceHeader    *string `json:"TraceHeader,omitempty" xml:"TraceHeader,omitempty"`
}

func (s GetSwimLaneGroupListResponseBodySwimLaneGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneGroupListResponseBodySwimLaneGroupList) GoString() string {
	return s.String()
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetFallbackTarget(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.FallbackTarget = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetGroupName(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.GroupName = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetIngressGatewayName(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.IngressGatewayName = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetIngressType(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.IngressType = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetIsPermissive(v bool) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.IsPermissive = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetRouteHeader(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.RouteHeader = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetServiceList(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.ServiceList = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetSwimLaneLabels(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.SwimLaneLabels = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetTraceHeader(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.TraceHeader = &v
	return s
}

type GetSwimLaneGroupListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSwimLaneGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSwimLaneGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneGroupListResponse) GoString() string {
	return s.String()
}

func (s *GetSwimLaneGroupListResponse) SetHeaders(v map[string]*string) *GetSwimLaneGroupListResponse {
	s.Headers = v
	return s
}

func (s *GetSwimLaneGroupListResponse) SetStatusCode(v int32) *GetSwimLaneGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSwimLaneGroupListResponse) SetBody(v *GetSwimLaneGroupListResponseBody) *GetSwimLaneGroupListResponse {
	s.Body = v
	return s
}

type GetSwimLaneListRequest struct {
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetSwimLaneListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneListRequest) GoString() string {
	return s.String()
}

func (s *GetSwimLaneListRequest) SetGroupName(v string) *GetSwimLaneListRequest {
	s.GroupName = &v
	return s
}

func (s *GetSwimLaneListRequest) SetServiceMeshId(v string) *GetSwimLaneListRequest {
	s.ServiceMeshId = &v
	return s
}

type GetSwimLaneListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of all the lanes in the lane group.
	SwimLaneList []*GetSwimLaneListResponseBodySwimLaneList `json:"SwimLaneList,omitempty" xml:"SwimLaneList,omitempty" type:"Repeated"`
}

func (s GetSwimLaneListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwimLaneListResponseBody) SetRequestId(v string) *GetSwimLaneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSwimLaneListResponseBody) SetSwimLaneList(v []*GetSwimLaneListResponseBodySwimLaneList) *GetSwimLaneListResponseBody {
	s.SwimLaneList = v
	return s
}

type GetSwimLaneListResponseBodySwimLaneList struct {
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The traffic routing rule associated with the lane.
	IngressRule *string `json:"IngressRule,omitempty" xml:"IngressRule,omitempty"`
	// This parameter is deprecated.
	IngressService *string `json:"IngressService,omitempty" xml:"IngressService,omitempty"`
	// The label key of the associated service workload. Fixed value: `ASM_TRAFFIC_TAG`.
	LabelSelectorKey *string `json:"LabelSelectorKey,omitempty" xml:"LabelSelectorKey,omitempty"`
	// The label value of the associated service workload.``
	LabelSelectorValue *string `json:"LabelSelectorValue,omitempty" xml:"LabelSelectorValue,omitempty"`
	// The name of the lane.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A list of services associated with the lane.
	ServiceList *string `json:"ServiceList,omitempty" xml:"ServiceList,omitempty"`
}

func (s GetSwimLaneListResponseBodySwimLaneList) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneListResponseBodySwimLaneList) GoString() string {
	return s.String()
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetGroupName(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.GroupName = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetIngressRule(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.IngressRule = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetIngressService(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.IngressService = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetLabelSelectorKey(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.LabelSelectorKey = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetLabelSelectorValue(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.LabelSelectorValue = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetName(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.Name = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetServiceList(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.ServiceList = &v
	return s
}

type GetSwimLaneListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSwimLaneListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSwimLaneListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSwimLaneListResponse) GoString() string {
	return s.String()
}

func (s *GetSwimLaneListResponse) SetHeaders(v map[string]*string) *GetSwimLaneListResponse {
	s.Headers = v
	return s
}

func (s *GetSwimLaneListResponse) SetStatusCode(v int32) *GetSwimLaneListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSwimLaneListResponse) SetBody(v *GetSwimLaneListResponseBody) *GetSwimLaneListResponse {
	s.Body = v
	return s
}

type GetVmAppMeshInfoRequest struct {
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetVmAppMeshInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVmAppMeshInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVmAppMeshInfoRequest) SetServiceMeshId(v string) *GetVmAppMeshInfoRequest {
	s.ServiceMeshId = &v
	return s
}

type GetVmAppMeshInfoResponseBody struct {
	// The returned information.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVmAppMeshInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVmAppMeshInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVmAppMeshInfoResponseBody) SetData(v string) *GetVmAppMeshInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetVmAppMeshInfoResponseBody) SetRequestId(v string) *GetVmAppMeshInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetVmAppMeshInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVmAppMeshInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVmAppMeshInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVmAppMeshInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVmAppMeshInfoResponse) SetHeaders(v map[string]*string) *GetVmAppMeshInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVmAppMeshInfoResponse) SetStatusCode(v int32) *GetVmAppMeshInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVmAppMeshInfoResponse) SetBody(v *GetVmAppMeshInfoResponseBody) *GetVmAppMeshInfoResponse {
	s.Body = v
	return s
}

type GetVmMetaRequest struct {
	// The namespace. This parameter is valid only after you set the Namespace and the ServiceAccount parameters.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The service account. This parameter is valid only after you set the Namespace and the ServiceAccount parameters.
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The trusted domain. Default value: cluster.local. This parameter is valid only after you set the Namespace and the ServiceAccount parameters.
	TrustDomain *string `json:"TrustDomain,omitempty" xml:"TrustDomain,omitempty"`
}

func (s GetVmMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVmMetaRequest) GoString() string {
	return s.String()
}

func (s *GetVmMetaRequest) SetNamespace(v string) *GetVmMetaRequest {
	s.Namespace = &v
	return s
}

func (s *GetVmMetaRequest) SetServiceAccount(v string) *GetVmMetaRequest {
	s.ServiceAccount = &v
	return s
}

func (s *GetVmMetaRequest) SetServiceMeshId(v string) *GetVmMetaRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetVmMetaRequest) SetTrustDomain(v string) *GetVmMetaRequest {
	s.TrustDomain = &v
	return s
}

type GetVmMetaResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata that is required to add a non-containerized application to the ASM instance.
	VmMetaInfo *GetVmMetaResponseBodyVmMetaInfo `json:"VmMetaInfo,omitempty" xml:"VmMetaInfo,omitempty" type:"Struct"`
}

func (s GetVmMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVmMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetVmMetaResponseBody) SetRequestId(v string) *GetVmMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVmMetaResponseBody) SetVmMetaInfo(v *GetVmMetaResponseBodyVmMetaInfo) *GetVmMetaResponseBody {
	s.VmMetaInfo = v
	return s
}

type GetVmMetaResponseBodyVmMetaInfo struct {
	// The content of the EnvoyEnv file.
	EnvoyEnvContent *string `json:"EnvoyEnvContent,omitempty" xml:"EnvoyEnvContent,omitempty"`
	// The content of the hosts file.
	HostsContent *string `json:"HostsContent,omitempty" xml:"HostsContent,omitempty"`
	// The content of the Token file.
	TokenContent *string `json:"TokenContent,omitempty" xml:"TokenContent,omitempty"`
}

func (s GetVmMetaResponseBodyVmMetaInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVmMetaResponseBodyVmMetaInfo) GoString() string {
	return s.String()
}

func (s *GetVmMetaResponseBodyVmMetaInfo) SetEnvoyEnvContent(v string) *GetVmMetaResponseBodyVmMetaInfo {
	s.EnvoyEnvContent = &v
	return s
}

func (s *GetVmMetaResponseBodyVmMetaInfo) SetHostsContent(v string) *GetVmMetaResponseBodyVmMetaInfo {
	s.HostsContent = &v
	return s
}

func (s *GetVmMetaResponseBodyVmMetaInfo) SetTokenContent(v string) *GetVmMetaResponseBodyVmMetaInfo {
	s.TokenContent = &v
	return s
}

type GetVmMetaResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVmMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVmMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVmMetaResponse) GoString() string {
	return s.String()
}

func (s *GetVmMetaResponse) SetHeaders(v map[string]*string) *GetVmMetaResponse {
	s.Headers = v
	return s
}

func (s *GetVmMetaResponse) SetStatusCode(v int32) *GetVmMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVmMetaResponse) SetBody(v *GetVmMetaResponseBody) *GetVmMetaResponse {
	s.Body = v
	return s
}

type GrantUserPermissionsRequest struct {
	// The permissions that are granted to an entity. The content is a string that consists of JSON arrays. You must specify all permissions that you want to grant to an entity. You can add or remove permissions by modifying the content. Field definition of the sample code:
	//
	// *   `IsCustom`: a parameter that is required by the system. Set the value to `true`.
	// *   `Cluster`: the ID of the Alibaba Cloud Service Mesh (ASM) instance.
	// *   `RoleName`: the name of the permissions. Valid values: `istio-admin`, `istio-ops`, and `istio-readonly`. A value of istio-admin indicates the permissions of ASM administrators. A value of istio-ops indicates the permissions of ASM restricted users. A value of istio-readonly indicates the read-only permissions.
	// *   `IsRamRole`: the entity to which you want to grant the permissions. To grant the permissions to a RAM role, set the value to `true`. Otherwise, set the value to `false`.
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The ID of the RAM user or RAM role.
	SubAccountUserId *string `json:"SubAccountUserId,omitempty" xml:"SubAccountUserId,omitempty"`
	// The ID list of the RAM user or RAM role.
	SubAccountUserIds []*string `json:"SubAccountUserIds,omitempty" xml:"SubAccountUserIds,omitempty" type:"Repeated"`
}

func (s GrantUserPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsRequest) SetPermissions(v string) *GrantUserPermissionsRequest {
	s.Permissions = &v
	return s
}

func (s *GrantUserPermissionsRequest) SetSubAccountUserId(v string) *GrantUserPermissionsRequest {
	s.SubAccountUserId = &v
	return s
}

func (s *GrantUserPermissionsRequest) SetSubAccountUserIds(v []*string) *GrantUserPermissionsRequest {
	s.SubAccountUserIds = v
	return s
}

type GrantUserPermissionsShrinkRequest struct {
	// The permissions that are granted to an entity. The content is a string that consists of JSON arrays. You must specify all permissions that you want to grant to an entity. You can add or remove permissions by modifying the content. Field definition of the sample code:
	//
	// *   `IsCustom`: a parameter that is required by the system. Set the value to `true`.
	// *   `Cluster`: the ID of the Alibaba Cloud Service Mesh (ASM) instance.
	// *   `RoleName`: the name of the permissions. Valid values: `istio-admin`, `istio-ops`, and `istio-readonly`. A value of istio-admin indicates the permissions of ASM administrators. A value of istio-ops indicates the permissions of ASM restricted users. A value of istio-readonly indicates the read-only permissions.
	// *   `IsRamRole`: the entity to which you want to grant the permissions. To grant the permissions to a RAM role, set the value to `true`. Otherwise, set the value to `false`.
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The ID of the RAM user or RAM role.
	SubAccountUserId *string `json:"SubAccountUserId,omitempty" xml:"SubAccountUserId,omitempty"`
	// The ID list of the RAM user or RAM role.
	SubAccountUserIdsShrink *string `json:"SubAccountUserIds,omitempty" xml:"SubAccountUserIds,omitempty"`
}

func (s GrantUserPermissionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsShrinkRequest) SetPermissions(v string) *GrantUserPermissionsShrinkRequest {
	s.Permissions = &v
	return s
}

func (s *GrantUserPermissionsShrinkRequest) SetSubAccountUserId(v string) *GrantUserPermissionsShrinkRequest {
	s.SubAccountUserId = &v
	return s
}

func (s *GrantUserPermissionsShrinkRequest) SetSubAccountUserIdsShrink(v string) *GrantUserPermissionsShrinkRequest {
	s.SubAccountUserIdsShrink = &v
	return s
}

type GrantUserPermissionsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantUserPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsResponseBody) SetRequestId(v string) *GrantUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

type GrantUserPermissionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantUserPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsResponse) SetHeaders(v map[string]*string) *GrantUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *GrantUserPermissionsResponse) SetStatusCode(v int32) *GrantUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantUserPermissionsResponse) SetBody(v *GrantUserPermissionsResponseBody) *GrantUserPermissionsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyApiServerEipResourceRequest struct {
	// The ID of the EIP.
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// The type of the operation. Valid values:
	//
	// *   `UnBindEip`: disassociates an EIP from the API server.
	// *   `BindEip`: associates an EIP with the API server.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ModifyApiServerEipResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiServerEipResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiServerEipResourceRequest) SetApiServerEipId(v string) *ModifyApiServerEipResourceRequest {
	s.ApiServerEipId = &v
	return s
}

func (s *ModifyApiServerEipResourceRequest) SetOperation(v string) *ModifyApiServerEipResourceRequest {
	s.Operation = &v
	return s
}

func (s *ModifyApiServerEipResourceRequest) SetServiceMeshId(v string) *ModifyApiServerEipResourceRequest {
	s.ServiceMeshId = &v
	return s
}

type ModifyApiServerEipResourceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiServerEipResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiServerEipResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiServerEipResourceResponseBody) SetRequestId(v string) *ModifyApiServerEipResourceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyApiServerEipResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyApiServerEipResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyApiServerEipResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiServerEipResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiServerEipResourceResponse) SetHeaders(v map[string]*string) *ModifyApiServerEipResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiServerEipResourceResponse) SetStatusCode(v int32) *ModifyApiServerEipResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiServerEipResourceResponse) SetBody(v *ModifyApiServerEipResourceResponseBody) *ModifyApiServerEipResourceResponse {
	s.Body = v
	return s
}

type ModifyServiceMeshNameRequest struct {
	// The new name of the ASM instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ModifyServiceMeshNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceMeshNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyServiceMeshNameRequest) SetName(v string) *ModifyServiceMeshNameRequest {
	s.Name = &v
	return s
}

func (s *ModifyServiceMeshNameRequest) SetServiceMeshId(v string) *ModifyServiceMeshNameRequest {
	s.ServiceMeshId = &v
	return s
}

type ModifyServiceMeshNameResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyServiceMeshNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceMeshNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServiceMeshNameResponseBody) SetRequestId(v string) *ModifyServiceMeshNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyServiceMeshNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyServiceMeshNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyServiceMeshNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceMeshNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyServiceMeshNameResponse) SetHeaders(v map[string]*string) *ModifyServiceMeshNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyServiceMeshNameResponse) SetStatusCode(v int32) *ModifyServiceMeshNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServiceMeshNameResponse) SetBody(v *ModifyServiceMeshNameResponseBody) *ModifyServiceMeshNameResponse {
	s.Body = v
	return s
}

type ReActivateAuditRequest struct {
	// Specifies whether to recreate a project that is used to store audit logs. Valid values:
	//
	// *   true: recreates a project.
	// *   false: does not recreate a project.
	EnableAudit *bool `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ReActivateAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s ReActivateAuditRequest) GoString() string {
	return s.String()
}

func (s *ReActivateAuditRequest) SetEnableAudit(v bool) *ReActivateAuditRequest {
	s.EnableAudit = &v
	return s
}

func (s *ReActivateAuditRequest) SetServiceMeshId(v string) *ReActivateAuditRequest {
	s.ServiceMeshId = &v
	return s
}

type ReActivateAuditResponseBody struct {
	// The name of the project that is used to store audit logs.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReActivateAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReActivateAuditResponseBody) GoString() string {
	return s.String()
}

func (s *ReActivateAuditResponseBody) SetData(v string) *ReActivateAuditResponseBody {
	s.Data = &v
	return s
}

func (s *ReActivateAuditResponseBody) SetRequestId(v string) *ReActivateAuditResponseBody {
	s.RequestId = &v
	return s
}

type ReActivateAuditResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReActivateAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReActivateAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s ReActivateAuditResponse) GoString() string {
	return s.String()
}

func (s *ReActivateAuditResponse) SetHeaders(v map[string]*string) *ReActivateAuditResponse {
	s.Headers = v
	return s
}

func (s *ReActivateAuditResponse) SetStatusCode(v int32) *ReActivateAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *ReActivateAuditResponse) SetBody(v *ReActivateAuditResponseBody) *ReActivateAuditResponse {
	s.Body = v
	return s
}

type RemoveClusterFromServiceMeshRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 移除集群时，保留istio-system 命名空间
	ReserveNamespace *bool   `json:"ReserveNamespace,omitempty" xml:"ReserveNamespace,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s RemoveClusterFromServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterFromServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *RemoveClusterFromServiceMeshRequest) SetClusterId(v string) *RemoveClusterFromServiceMeshRequest {
	s.ClusterId = &v
	return s
}

func (s *RemoveClusterFromServiceMeshRequest) SetReserveNamespace(v bool) *RemoveClusterFromServiceMeshRequest {
	s.ReserveNamespace = &v
	return s
}

func (s *RemoveClusterFromServiceMeshRequest) SetServiceMeshId(v string) *RemoveClusterFromServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

type RemoveClusterFromServiceMeshResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveClusterFromServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterFromServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClusterFromServiceMeshResponseBody) SetCode(v string) *RemoveClusterFromServiceMeshResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveClusterFromServiceMeshResponseBody) SetMessage(v string) *RemoveClusterFromServiceMeshResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveClusterFromServiceMeshResponseBody) SetRequestId(v string) *RemoveClusterFromServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

type RemoveClusterFromServiceMeshResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveClusterFromServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveClusterFromServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClusterFromServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterFromServiceMeshResponse) SetHeaders(v map[string]*string) *RemoveClusterFromServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *RemoveClusterFromServiceMeshResponse) SetStatusCode(v int32) *RemoveClusterFromServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClusterFromServiceMeshResponse) SetBody(v *RemoveClusterFromServiceMeshResponseBody) *RemoveClusterFromServiceMeshResponse {
	s.Body = v
	return s
}

type RemoveVMFromServiceMeshRequest struct {
	// The ID of the ECS instance.
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s RemoveVMFromServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveVMFromServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *RemoveVMFromServiceMeshRequest) SetEcsId(v string) *RemoveVMFromServiceMeshRequest {
	s.EcsId = &v
	return s
}

func (s *RemoveVMFromServiceMeshRequest) SetServiceMeshId(v string) *RemoveVMFromServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

type RemoveVMFromServiceMeshResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVMFromServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveVMFromServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVMFromServiceMeshResponseBody) SetRequestId(v string) *RemoveVMFromServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

type RemoveVMFromServiceMeshResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveVMFromServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveVMFromServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveVMFromServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *RemoveVMFromServiceMeshResponse) SetHeaders(v map[string]*string) *RemoveVMFromServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *RemoveVMFromServiceMeshResponse) SetStatusCode(v int32) *RemoveVMFromServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveVMFromServiceMeshResponse) SetBody(v *RemoveVMFromServiceMeshResponseBody) *RemoveVMFromServiceMeshResponse {
	s.Body = v
	return s
}

type RevokeKubeconfigRequest struct {
	// Specifies whether to return the kubeconfig file for private access.
	//
	// *   `true`: returns the kubeconfig file for private access.
	// *   `false`: returns the kubeconfig file for public access.
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the ASM instance for which you want to revoke a kubeconfig file.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s RevokeKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *RevokeKubeconfigRequest) SetPrivateIpAddress(v bool) *RevokeKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RevokeKubeconfigRequest) SetServiceMeshId(v string) *RevokeKubeconfigRequest {
	s.ServiceMeshId = &v
	return s
}

type RevokeKubeconfigResponseBody struct {
	// The new kubeconfig file generated.
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeKubeconfigResponseBody) SetKubeconfig(v string) *RevokeKubeconfigResponseBody {
	s.Kubeconfig = &v
	return s
}

func (s *RevokeKubeconfigResponseBody) SetRequestId(v string) *RevokeKubeconfigResponseBody {
	s.RequestId = &v
	return s
}

type RevokeKubeconfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *RevokeKubeconfigResponse) SetHeaders(v map[string]*string) *RevokeKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *RevokeKubeconfigResponse) SetStatusCode(v int32) *RevokeKubeconfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeKubeconfigResponse) SetBody(v *RevokeKubeconfigResponseBody) *RevokeKubeconfigResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateASMGatewayRequest struct {
	// The new YAML content of the ASM gateway.
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateASMGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayRequest) SetBody(v string) *UpdateASMGatewayRequest {
	s.Body = &v
	return s
}

func (s *UpdateASMGatewayRequest) SetIstioGatewayName(v string) *UpdateASMGatewayRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *UpdateASMGatewayRequest) SetServiceMeshId(v string) *UpdateASMGatewayRequest {
	s.ServiceMeshId = &v
	return s
}

type UpdateASMGatewayResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateASMGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayResponseBody) SetRequestId(v string) *UpdateASMGatewayResponseBody {
	s.RequestId = &v
	return s
}

type UpdateASMGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateASMGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateASMGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayResponse) SetHeaders(v map[string]*string) *UpdateASMGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpdateASMGatewayResponse) SetStatusCode(v int32) *UpdateASMGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateASMGatewayResponse) SetBody(v *UpdateASMGatewayResponseBody) *UpdateASMGatewayResponse {
	s.Body = v
	return s
}

type UpdateASMGatewayImportedServicesRequest struct {
	// The name of the ASM gateway.
	ASMGatewayName *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The names of the services. Separate multiple service names with commas (,). Example: reviews,sleep.
	ServiceNames *string `json:"ServiceNames,omitempty" xml:"ServiceNames,omitempty"`
	// The namespace in which the service resides.
	ServiceNamespace *string `json:"ServiceNamespace,omitempty" xml:"ServiceNamespace,omitempty"`
}

func (s UpdateASMGatewayImportedServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMGatewayImportedServicesRequest) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayImportedServicesRequest) SetASMGatewayName(v string) *UpdateASMGatewayImportedServicesRequest {
	s.ASMGatewayName = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesRequest) SetServiceMeshId(v string) *UpdateASMGatewayImportedServicesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesRequest) SetServiceNames(v string) *UpdateASMGatewayImportedServicesRequest {
	s.ServiceNames = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesRequest) SetServiceNamespace(v string) *UpdateASMGatewayImportedServicesRequest {
	s.ServiceNamespace = &v
	return s
}

type UpdateASMGatewayImportedServicesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateASMGatewayImportedServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMGatewayImportedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayImportedServicesResponseBody) SetRequestId(v string) *UpdateASMGatewayImportedServicesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateASMGatewayImportedServicesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateASMGatewayImportedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateASMGatewayImportedServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMGatewayImportedServicesResponse) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayImportedServicesResponse) SetHeaders(v map[string]*string) *UpdateASMGatewayImportedServicesResponse {
	s.Headers = v
	return s
}

func (s *UpdateASMGatewayImportedServicesResponse) SetStatusCode(v int32) *UpdateASMGatewayImportedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesResponse) SetBody(v *UpdateASMGatewayImportedServicesResponseBody) *UpdateASMGatewayImportedServicesResponse {
	s.Body = v
	return s
}

type UpdateASMNamespaceFromGuestClusterRequest struct {
	// The ID of the Kubernetes cluster whose namespace information you want to synchronize to ASM. The Kubernetes cluster is added to the ASM instance that is specified by the ServiceMeshId parameter.
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateASMNamespaceFromGuestClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMNamespaceFromGuestClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateASMNamespaceFromGuestClusterRequest) SetK8sClusterId(v string) *UpdateASMNamespaceFromGuestClusterRequest {
	s.K8sClusterId = &v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterRequest) SetServiceMeshId(v string) *UpdateASMNamespaceFromGuestClusterRequest {
	s.ServiceMeshId = &v
	return s
}

type UpdateASMNamespaceFromGuestClusterResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateASMNamespaceFromGuestClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMNamespaceFromGuestClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateASMNamespaceFromGuestClusterResponseBody) SetRequestId(v string) *UpdateASMNamespaceFromGuestClusterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateASMNamespaceFromGuestClusterResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateASMNamespaceFromGuestClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateASMNamespaceFromGuestClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateASMNamespaceFromGuestClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) SetHeaders(v map[string]*string) *UpdateASMNamespaceFromGuestClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) SetStatusCode(v int32) *UpdateASMNamespaceFromGuestClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterResponse) SetBody(v *UpdateASMNamespaceFromGuestClusterResponseBody) *UpdateASMNamespaceFromGuestClusterResponse {
	s.Body = v
	return s
}

type UpdateControlPlaneLogConfigRequest struct {
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	LogTTLInDay   *int32  `json:"LogTTLInDay,omitempty" xml:"LogTTLInDay,omitempty"`
	Project       *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateControlPlaneLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPlaneLogConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogConfigRequest) SetEnabled(v bool) *UpdateControlPlaneLogConfigRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateControlPlaneLogConfigRequest) SetLogTTLInDay(v int32) *UpdateControlPlaneLogConfigRequest {
	s.LogTTLInDay = &v
	return s
}

func (s *UpdateControlPlaneLogConfigRequest) SetProject(v string) *UpdateControlPlaneLogConfigRequest {
	s.Project = &v
	return s
}

func (s *UpdateControlPlaneLogConfigRequest) SetServiceMeshId(v string) *UpdateControlPlaneLogConfigRequest {
	s.ServiceMeshId = &v
	return s
}

type UpdateControlPlaneLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateControlPlaneLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPlaneLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogConfigResponseBody) SetRequestId(v string) *UpdateControlPlaneLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateControlPlaneLogConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateControlPlaneLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateControlPlaneLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPlaneLogConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogConfigResponse) SetHeaders(v map[string]*string) *UpdateControlPlaneLogConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateControlPlaneLogConfigResponse) SetStatusCode(v int32) *UpdateControlPlaneLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateControlPlaneLogConfigResponse) SetBody(v *UpdateControlPlaneLogConfigResponseBody) *UpdateControlPlaneLogConfigResponse {
	s.Body = v
	return s
}

type UpdateIstioGatewayRoutesRequest struct {
	// The description of the routing rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the routing rule to be updated for the ASM gateway.
	GatewayRoute *UpdateIstioGatewayRoutesRequestGatewayRoute `json:"GatewayRoute,omitempty" xml:"GatewayRoute,omitempty" type:"Struct"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// *   `0`: The routing rule is valid.
	// *   `1`: The routing rule is invalid.
	// *   `2`: An error occurs during the creation or update of the routing rule.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequest) SetDescription(v string) *UpdateIstioGatewayRoutesRequest {
	s.Description = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetGatewayRoute(v *UpdateIstioGatewayRoutesRequestGatewayRoute) *UpdateIstioGatewayRoutesRequest {
	s.GatewayRoute = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetIstioGatewayName(v string) *UpdateIstioGatewayRoutesRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetPriority(v int32) *UpdateIstioGatewayRoutesRequest {
	s.Priority = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetServiceMeshId(v string) *UpdateIstioGatewayRoutesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequest) SetStatus(v int32) *UpdateIstioGatewayRoutesRequest {
	s.Status = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRoute struct {
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The advanced settings for routing HTTP traffic.
	HTTPAdvancedOptions *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions `json:"HTTPAdvancedOptions,omitempty" xml:"HTTPAdvancedOptions,omitempty" type:"Struct"`
	// The matching rules for traffic routing.
	MatchRequest *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest `json:"MatchRequest,omitempty" xml:"MatchRequest,omitempty" type:"Struct"`
	// The namespace in which the destination service resides.
	Namespace  *string     `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RawVSRoute interface{} `json:"RawVSRoute,omitempty" xml:"RawVSRoute,omitempty"`
	// The endpoints of destination services for Layer 4 weighted routing.
	RouteDestinations []*UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations `json:"RouteDestinations,omitempty" xml:"RouteDestinations,omitempty" type:"Repeated"`
	// The name of the routing rule.
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The type of the traffic to be routed. Valid values: `HTTP`, `TLS`, and `TCP`.
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRoute) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRoute) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetDomains(v []*string) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.Domains = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetHTTPAdvancedOptions(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.HTTPAdvancedOptions = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetMatchRequest(v *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.MatchRequest = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetNamespace(v string) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.Namespace = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetRawVSRoute(v interface{}) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.RawVSRoute = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetRouteDestinations(v []*UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteDestinations = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetRouteName(v string) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteName = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRoute) SetRouteType(v string) *UpdateIstioGatewayRoutesRequestGatewayRoute {
	s.RouteType = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions struct {
	// The virtual service that defines traffic routing.
	Delegate *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate `json:"Delegate,omitempty" xml:"Delegate,omitempty" type:"Struct"`
	// The configurations of fault injection.
	Fault *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault `json:"Fault,omitempty" xml:"Fault,omitempty" type:"Struct"`
	// The HTTP redirection rule.
	HTTPRedirect *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect `json:"HTTPRedirect,omitempty" xml:"HTTPRedirect,omitempty" type:"Struct"`
	// The configurations for mirroring HTTP traffic to another destination in addition to forwarding requests to the specified destination.
	Mirror *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror `json:"Mirror,omitempty" xml:"Mirror,omitempty" type:"Struct"`
	// The percentage of requests that are mirrored to another destination except for the original destination.
	MirrorPercentage *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage `json:"MirrorPercentage,omitempty" xml:"MirrorPercentage,omitempty" type:"Struct"`
	// The configurations of retries for failed requests.
	Retries *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries `json:"Retries,omitempty" xml:"Retries,omitempty" type:"Struct"`
	// The configurations for rewriting the virtual service.
	Rewrite *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Struct"`
	// The timeout period for requests.
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetDelegate(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Delegate = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetFault(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Fault = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetHTTPRedirect(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.HTTPRedirect = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetMirror(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Mirror = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetMirrorPercentage(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.MirrorPercentage = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetRetries(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Retries = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetRewrite(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Rewrite = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions) SetTimeout(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptions {
	s.Timeout = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate struct {
	// The name of the virtual service.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the virtual service belongs.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) SetName(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	s.Name = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate) SetNamespace(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsDelegate {
	s.Namespace = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault struct {
	// The configurations for aborting requests with specified error codes.
	Abort *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort `json:"Abort,omitempty" xml:"Abort,omitempty" type:"Struct"`
	// The duration to delay a request.
	Delay *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay `json:"Delay,omitempty" xml:"Delay,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) SetAbort(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	s.Abort = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault) SetDelay(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFault {
	s.Delay = v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort struct {
	// The HTTP status code.
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The percentage of requests that are aborted with the specified error code.
	Percentage *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) SetHttpStatus(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	s.HttpStatus = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort) SetPercentage(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbort {
	s.Percentage = v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage struct {
	// The percentage of requests that are aborted with the specified error code, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage) SetValue(v float32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultAbortPercentage {
	s.Value = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay struct {
	// The fixed duration for request delay.
	FixedDelay *string `json:"FixedDelay,omitempty" xml:"FixedDelay,omitempty"`
	// The percentage of requests to which the delay fault is injected.
	Percentage *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage `json:"Percentage,omitempty" xml:"Percentage,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) SetFixedDelay(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	s.FixedDelay = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay) SetPercentage(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelay {
	s.Percentage = v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage struct {
	// The percentage of requests to which the delay fault is injected, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage) SetValue(v float32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsFaultDelayPercentage {
	s.Value = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect struct {
	// The value to be used to overwrite the value of the Authority or Host header during redirection.
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The HTTP status code to be used to indicate URL redirection. Default value: 301.
	RedirectCode *int32 `json:"RedirectCode,omitempty" xml:"RedirectCode,omitempty"`
	// The value to be used to overwrite the URL path during redirection.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetAuthority(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.Authority = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetRedirectCode(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.RedirectCode = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect) SetUri(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsHTTPRedirect {
	s.Uri = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror struct {
	// The name of the service defined in the service registry.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The name of the service subset.
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) SetHost(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	s.Host = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror) SetSubset(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirror {
	s.Subset = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage struct {
	// The percentage of requests that are mirrored to another destination except for the original destination, which is expressed as a decimal.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage) SetValue(v float32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsMirrorPercentage {
	s.Value = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries struct {
	// The number of retries allowed for a request.
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// The timeout period for each retry.
	PerTryTimeout *string `json:"PerTryTimeout,omitempty" xml:"PerTryTimeout,omitempty"`
	// The condition for retries. Example: `connect-failure,refused-stream,503`.
	RetryOn *string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty"`
	// Specifies whether to allow retries to other localities.
	RetryRemoteLocalities *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities `json:"RetryRemoteLocalities,omitempty" xml:"RetryRemoteLocalities,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetAttempts(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.Attempts = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetPerTryTimeout(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.PerTryTimeout = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetRetryOn(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.RetryOn = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries) SetRetryRemoteLocalities(v *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetries {
	s.RetryRemoteLocalities = v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities struct {
	// Specifies whether to allow retries to other localities. Valid values:
	//
	// *   `true`
	// *   `false`
	//
	// Default value: `false`.
	Value *bool `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities) SetValue(v bool) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRetriesRetryRemoteLocalities {
	s.Value = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite struct {
	// The value to be used to overwrite the value of the Authority or Host header.
	Authority *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	// The value to be used to overwrite the path or prefix of the URI.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) SetAuthority(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	s.Authority = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite) SetUri(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteHTTPAdvancedOptionsRewrite {
	s.Uri = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest struct {
	// The request headers to be matched.
	Headers []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The ports.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The matching rule for Transport Layer Security (TLS) traffic.
	TLSMatchAttributes []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes `json:"TLSMatchAttributes,omitempty" xml:"TLSMatchAttributes,omitempty" type:"Repeated"`
	// The matching rule for URIs.
	URI *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI `json:"URI,omitempty" xml:"URI,omitempty" type:"Struct"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetHeaders(v []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.Headers = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetPorts(v []*int32) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.Ports = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetTLSMatchAttributes(v []*UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.TLSMatchAttributes = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest) SetURI(v *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequest {
	s.URI = v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders struct {
	// The header value to be matched.
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode for the header value. Valid values:
	//
	// *   `exact`: exact match
	// *   `prefix`: match by prefix
	// *   `regex`: match by regular expression
	MatchingMode *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
	// The header key to be matched.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetMatchingContent(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.MatchingContent = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetMatchingMode(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.MatchingMode = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders) SetName(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestHeaders {
	s.Name = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes struct {
	// The Server Name Indication (SNI) values to be matched.
	SNIHosts []*string `json:"SNIHosts,omitempty" xml:"SNIHosts,omitempty" type:"Repeated"`
	// The TLS port.
	TLSPort *int32 `json:"TLSPort,omitempty" xml:"TLSPort,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) SetSNIHosts(v []*string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	s.SNIHosts = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes) SetTLSPort(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestTLSMatchAttributes {
	s.TLSPort = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI struct {
	// The content to be matched.
	MatchingContent *string `json:"MatchingContent,omitempty" xml:"MatchingContent,omitempty"`
	// The matching mode for the routing rule. Valid values:
	//
	// *   `exact`: exact match
	// *   `prefix`: match by prefix
	// *   `regex`: match by regular expression
	MatchingMode *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) SetMatchingContent(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	s.MatchingContent = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI) SetMatchingMode(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteMatchRequestURI {
	s.MatchingMode = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations struct {
	// The unique endpoint of the destination service to which the specified requests are sent.
	Destination *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination `json:"Destination,omitempty" xml:"Destination,omitempty" type:"Struct"`
	// The weight of the service subset.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) SetDestination(v *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	s.Destination = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations) SetWeight(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinations {
	s.Weight = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination struct {
	// The name of the service defined in the service registry.
	Host *string                                                                      `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort `json:"Port,omitempty" xml:"Port,omitempty" type:"Struct"`
	// The name of the service subset.
	Subset *string `json:"Subset,omitempty" xml:"Subset,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetHost(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Host = &v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetPort(v *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Port = v
	return s
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination) SetSubset(v string) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestination {
	s.Subset = &v
	return s
}

type UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort struct {
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort) SetNumber(v int32) *UpdateIstioGatewayRoutesRequestGatewayRouteRouteDestinationsDestinationPort {
	s.Number = &v
	return s
}

type UpdateIstioGatewayRoutesShrinkRequest struct {
	// The description of the routing rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the routing rule to be updated for the ASM gateway.
	GatewayRouteShrink *string `json:"GatewayRoute,omitempty" xml:"GatewayRoute,omitempty"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// *   `0`: The routing rule is valid.
	// *   `1`: The routing rule is invalid.
	// *   `2`: An error occurs during the creation or update of the routing rule.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIstioGatewayRoutesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetDescription(v string) *UpdateIstioGatewayRoutesShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetGatewayRouteShrink(v string) *UpdateIstioGatewayRoutesShrinkRequest {
	s.GatewayRouteShrink = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetIstioGatewayName(v string) *UpdateIstioGatewayRoutesShrinkRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetPriority(v int32) *UpdateIstioGatewayRoutesShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetServiceMeshId(v string) *UpdateIstioGatewayRoutesShrinkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetStatus(v int32) *UpdateIstioGatewayRoutesShrinkRequest {
	s.Status = &v
	return s
}

type UpdateIstioGatewayRoutesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIstioGatewayRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesResponseBody) SetRequestId(v string) *UpdateIstioGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIstioGatewayRoutesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIstioGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIstioGatewayRoutesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesResponse) SetHeaders(v map[string]*string) *UpdateIstioGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *UpdateIstioGatewayRoutesResponse) SetStatusCode(v int32) *UpdateIstioGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIstioGatewayRoutesResponse) SetBody(v *UpdateIstioGatewayRoutesResponseBody) *UpdateIstioGatewayRoutesResponse {
	s.Body = v
	return s
}

type UpdateIstioInjectionConfigRequest struct {
	DataPlaneMode *string `json:"DataPlaneMode,omitempty" xml:"DataPlaneMode,omitempty"`
	// Specifies whether to enable Istio automatic sidecar injection.
	EnableIstioInjection *bool `json:"EnableIstioInjection,omitempty" xml:"EnableIstioInjection,omitempty"`
	// Specifies whether to enable automatic sidecar injection by using SidecarSet.
	EnableSidecarSetInjection *bool `json:"EnableSidecarSetInjection,omitempty" xml:"EnableSidecarSetInjection,omitempty"`
	// Specifies the version to be injected into the namespace. This parameter is valid only when the ASM instance performs a canary release. When IstioRev is not empty, you must not specify EnableIstioInjection and EnableSidecarSetInjection.
	IstioRev *string `json:"IstioRev,omitempty" xml:"IstioRev,omitempty"`
	// The namespace for which you want to modify the sidecar injection setting.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateIstioInjectionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioInjectionConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioInjectionConfigRequest) SetDataPlaneMode(v string) *UpdateIstioInjectionConfigRequest {
	s.DataPlaneMode = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetEnableIstioInjection(v bool) *UpdateIstioInjectionConfigRequest {
	s.EnableIstioInjection = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetEnableSidecarSetInjection(v bool) *UpdateIstioInjectionConfigRequest {
	s.EnableSidecarSetInjection = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetIstioRev(v string) *UpdateIstioInjectionConfigRequest {
	s.IstioRev = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetNamespace(v string) *UpdateIstioInjectionConfigRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateIstioInjectionConfigRequest) SetServiceMeshId(v string) *UpdateIstioInjectionConfigRequest {
	s.ServiceMeshId = &v
	return s
}

type UpdateIstioInjectionConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIstioInjectionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioInjectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIstioInjectionConfigResponseBody) SetRequestId(v string) *UpdateIstioInjectionConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIstioInjectionConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIstioInjectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIstioInjectionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioInjectionConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateIstioInjectionConfigResponse) SetHeaders(v map[string]*string) *UpdateIstioInjectionConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateIstioInjectionConfigResponse) SetStatusCode(v int32) *UpdateIstioInjectionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIstioInjectionConfigResponse) SetBody(v *UpdateIstioInjectionConfigResponseBody) *UpdateIstioInjectionConfigResponse {
	s.Body = v
	return s
}

type UpdateIstioRouteAdditionalStatusRequest struct {
	// The description of the routing rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the ASM gateway.
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The name of the routing rule.
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// *   `0`: The routing rule is valid.
	// *   `1`: The routing rule is invalid.
	// *   `2`: An error occurs during the creation or update of the routing rule.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIstioRouteAdditionalStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioRouteAdditionalStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetDescription(v string) *UpdateIstioRouteAdditionalStatusRequest {
	s.Description = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetIstioGatewayName(v string) *UpdateIstioRouteAdditionalStatusRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetPriority(v int32) *UpdateIstioRouteAdditionalStatusRequest {
	s.Priority = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetRouteName(v string) *UpdateIstioRouteAdditionalStatusRequest {
	s.RouteName = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetServiceMeshId(v string) *UpdateIstioRouteAdditionalStatusRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetStatus(v int32) *UpdateIstioRouteAdditionalStatusRequest {
	s.Status = &v
	return s
}

type UpdateIstioRouteAdditionalStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIstioRouteAdditionalStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioRouteAdditionalStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIstioRouteAdditionalStatusResponseBody) SetRequestId(v string) *UpdateIstioRouteAdditionalStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIstioRouteAdditionalStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIstioRouteAdditionalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIstioRouteAdditionalStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIstioRouteAdditionalStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateIstioRouteAdditionalStatusResponse) SetHeaders(v map[string]*string) *UpdateIstioRouteAdditionalStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusResponse) SetStatusCode(v int32) *UpdateIstioRouteAdditionalStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusResponse) SetBody(v *UpdateIstioRouteAdditionalStatusResponseBody) *UpdateIstioRouteAdditionalStatusResponse {
	s.Body = v
	return s
}

type UpdateMeshCRAggregationRequest struct {
	// The maximum number of CPU cores that are available for the components installed in the ACK cluster on the data plane if you enable the Kubernetes API to access Istio resources in the ASM instance. You can specify the parameter value in the standard quantity representation used by Kubernetes.
	CPULimit *string `json:"CPULimit,omitempty" xml:"CPULimit,omitempty"`
	// The number of CPU cores that are requested by the components installed in the Container Service for Kubernetes (ACK) cluster on the data plane if you enable the Kubernetes API to access Istio resources in the ASM instance. You can specify the parameter value in the standard quantity representation used by Kubernetes.
	CPURequirement *string `json:"CPURequirement,omitempty" xml:"CPURequirement,omitempty"`
	// Specifies whether to enable the Kubernetes API on the data plane to access Istio resources in the ASM instance. Valid values:
	//
	// *   `true`: enables the Kubernetes API to access Istio resources in the ASM instance.
	// *   `false`: does not enable the Kubernetes API to access Istio resources in the ASM instance.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum size of the memory that is available for the components installed in the ACK cluster on the data plane if you enable the Kubernetes API to access Istio resources in the ASM instance. You can specify the parameter value in the standard quantity representation used by Kubernetes. 1 Mi equals 1,024 KB.
	MemoryLimit *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The size of the memory that is requested by the components installed in the ACK cluster on the data plane if you enable the Kubernetes API to access Istio resources in the ASM instance. You can specify the parameter value in the standard quantity representation used by Kubernetes. 1 Mi equals 1,024 KB.
	MemoryRequirement *string `json:"MemoryRequirement,omitempty" xml:"MemoryRequirement,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// Specifies whether the Kubernetes API on the data plane uses the public endpoint of the API server to access Istio resources in the ASM instance. Valid values:
	//
	// *   `true`: The Kubernetes API on the data plane uses the public endpoint of the API server to access Istio resources in the ASM instance.
	// *   `false`: The Kubernetes API on the data plane uses the private endpoint of the API server to access Istio resources in the ASM instance.
	//
	// Default value: `false`.
	UsePublicApiServer *bool `json:"UsePublicApiServer,omitempty" xml:"UsePublicApiServer,omitempty"`
}

func (s UpdateMeshCRAggregationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeshCRAggregationRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeshCRAggregationRequest) SetCPULimit(v string) *UpdateMeshCRAggregationRequest {
	s.CPULimit = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetCPURequirement(v string) *UpdateMeshCRAggregationRequest {
	s.CPURequirement = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetEnabled(v bool) *UpdateMeshCRAggregationRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetMemoryLimit(v string) *UpdateMeshCRAggregationRequest {
	s.MemoryLimit = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetMemoryRequirement(v string) *UpdateMeshCRAggregationRequest {
	s.MemoryRequirement = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetServiceMeshId(v string) *UpdateMeshCRAggregationRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetUsePublicApiServer(v bool) *UpdateMeshCRAggregationRequest {
	s.UsePublicApiServer = &v
	return s
}

type UpdateMeshCRAggregationResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMeshCRAggregationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeshCRAggregationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeshCRAggregationResponseBody) SetRequestId(v string) *UpdateMeshCRAggregationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMeshCRAggregationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMeshCRAggregationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMeshCRAggregationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeshCRAggregationResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeshCRAggregationResponse) SetHeaders(v map[string]*string) *UpdateMeshCRAggregationResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeshCRAggregationResponse) SetStatusCode(v int32) *UpdateMeshCRAggregationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeshCRAggregationResponse) SetBody(v *UpdateMeshCRAggregationResponseBody) *UpdateMeshCRAggregationResponse {
	s.Body = v
	return s
}

type UpdateMeshFeatureRequest struct {
	// Specifies whether to enable access log collection. Valid values:
	//
	// *   `true`: enables access log collection.
	// *   `false`: disables access log collection.
	//
	// Default value: `false`.
	AccessLogEnabled *bool `json:"AccessLogEnabled,omitempty" xml:"AccessLogEnabled,omitempty"`
	// Specifies whether to enable access logging. Valid values:
	//
	// *   `""`: disables access logging.
	// *   `/dev/stdout`: enables access logging. Access logs are written to /dev/stdout.
	AccessLogFile *string `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	// The custom format of access logs. To set this parameter, make sure that you have enabled access log collection. The value must be a JSON string. The following key names must be contained: authority_for, bytes_received, bytes_sent, downstream_local_address, downstream_remote_address, duration, istio_policy_status, method, path, protocol, requested_server_name, response_code, response_flags, route_name, start_time, trace_id, upstream_cluster, upstream_host, upstream_local_address, upstream_service_time, upstream_transport_failure_reason, user_agent, and x_forwarded_for.
	AccessLogFormat         *string `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	AccessLogGatewayEnabled *bool   `json:"AccessLogGatewayEnabled,omitempty" xml:"AccessLogGatewayEnabled,omitempty"`
	// The retention period for the access logs of the sidecar proxy. Unit: day. The logs are collected by using Log Service. For example, `30` indicates 30 days.
	AccessLogGatewayLifecycle *int32 `json:"AccessLogGatewayLifecycle,omitempty" xml:"AccessLogGatewayLifecycle,omitempty"`
	// The custom project on which the Log Service collects logs.
	AccessLogProject *string `json:"AccessLogProject,omitempty" xml:"AccessLogProject,omitempty"`
	// Specifies whether to enable gRPC Access Log Service (ALS) for Envoy. Valid values:
	//
	// *   `true`: enables gRPC ALS for Envoy.
	// *   `false`: disables gRPC ALS for Envoy.
	//
	// Default value: `false`.
	AccessLogServiceEnabled *bool `json:"AccessLogServiceEnabled,omitempty" xml:"AccessLogServiceEnabled,omitempty"`
	// The endpoint of gRPC ALS for Envoy.
	AccessLogServiceHost *string `json:"AccessLogServiceHost,omitempty" xml:"AccessLogServiceHost,omitempty"`
	// The port of gRPC ALS for Envoy.
	AccessLogServicePort    *int32 `json:"AccessLogServicePort,omitempty" xml:"AccessLogServicePort,omitempty"`
	AccessLogSidecarEnabled *bool  `json:"AccessLogSidecarEnabled,omitempty" xml:"AccessLogSidecarEnabled,omitempty"`
	// Specifies whether to enable automatic diagnostics for the ASM instance. If you enable this feature, the ASM instance is automatically diagnosed when you modify Istio resources in the ASM instance.
	AccessLogSidecarLifecycle *int32 `json:"AccessLogSidecarLifecycle,omitempty" xml:"AccessLogSidecarLifecycle,omitempty"`
	// The name of the Log Service project that is used for mesh audit.
	//
	// Default value: `mesh-log-{ASM instance ID}`.
	AuditProject *string `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	// Specifies whether to enable automatic sidecar proxy injection by using pod annotations. Valid values:
	//
	// *   `true`: enables automatic sidecar proxy injection by using pod annotations.
	// *   `false`: disables automatic sidecar proxy injection by using pod annotations.
	//
	// Default value: `false`.
	AutoInjectionPolicyEnabled *bool `json:"AutoInjectionPolicyEnabled,omitempty" xml:"AutoInjectionPolicyEnabled,omitempty"`
	// Specifies whether to use the Kubernetes API of clusters on the data plane to access Istio resources. To use this feature, the version of the ASM instance must be V1.9.7.93 or later.
	CRAggregationEnabled *bool `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	// Specifies whether to enable the feature of controlling the OPA injection scope. Valid values:
	//
	// *   `true`: enables the feature.
	// *   `false`: disables the feature.
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// Specifies whether to enable the Container Network Interface (CNI) plug-in. Valid values:
	//
	// *   `true`: enables the CNI plug-in.
	// *   `false`: disables the CNI plug-in.
	//
	// Default value: `false`.
	CniEnabled *bool `json:"CniEnabled,omitempty" xml:"CniEnabled,omitempty"`
	// The namespaces to be excluded for the CNI plug-in.
	CniExcludeNamespaces *string `json:"CniExcludeNamespaces,omitempty" xml:"CniExcludeNamespaces,omitempty"`
	// Specifies whether to delay application container startup until the sidecar proxy container is started in a pod.
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// Specifies whether to enable the external service registry. Valid values:
	//
	// *   `true`: enables the external service registry.
	// *   `false`: disables the external service registry.
	//
	// Default value: `false`.
	ConfigSourceEnabled *bool `json:"ConfigSourceEnabled,omitempty" xml:"ConfigSourceEnabled,omitempty"`
	// The instance ID of the Nacos registry.
	ConfigSourceNacosID *string `json:"ConfigSourceNacosID,omitempty" xml:"ConfigSourceNacosID,omitempty"`
	// Specifies whether to use a custom Prometheus instance. Valid values:
	//
	// *   `true`: uses a custom Prometheus instance.
	// *   `false`: does not use a custom Prometheus instance.
	//
	// Default value: `false`.
	CustomizedPrometheus *bool `json:"CustomizedPrometheus,omitempty" xml:"CustomizedPrometheus,omitempty"`
	// Specifies whether to use a self-managed Zipkin system to collect tracing data. Valid values:
	//
	// *   `true`: uses a self-managed Zipkin system.
	// *   `false`: does not use a self-managed Zipkin system.
	//
	// Default value: `false`.
	CustomizedZipkin *bool `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	// Specifies whether to enable DNS proxy. Valid values:
	//
	// *   `true`: enables the DNS proxy feature.
	// *   `false`: disables the DNS proxy feature.
	//
	// Default value: `false`.
	DNSProxyingEnabled *bool `json:"DNSProxyingEnabled,omitempty" xml:"DNSProxyingEnabled,omitempty"`
	// Specifies the default scheduling configurations that ASM delivers to components on the data plane. You can configure `nodeSelector` and `tolerations` in the JSON format.
	//
	// >
	//
	// *   Modifying the value of this parameter is a high-risk operation. The modification will cause all components on the data plane of ASM to restart. Exercise caution before modifying the value of this parameter.
	//
	// *   The configurations specified by this parameter do not apply to the ASM gateway. You can configure gateway-specific scheduling on the ASM gateway.
	DefaultComponentsScheduleConfig *string `json:"DefaultComponentsScheduleConfig,omitempty" xml:"DefaultComponentsScheduleConfig,omitempty"`
	// The label selectors used to specify the namespaces of the clusters on the data plane for selective service discovery.
	DiscoverySelectors *string `json:"DiscoverySelectors,omitempty" xml:"DiscoverySelectors,omitempty"`
	// Specifies whether to enable Dubbo Filter. Valid values:
	//
	// *   `true`: enables Dubbo Filter.
	// *   `false`: disables Dubbo Filter.
	//
	// Default value: `false`.
	DubboFilterEnabled *bool `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	// Specifies whether to enable the mesh audit feature. To enable this feature, make sure that you have activated [Log Service](https://sls.console.aliyun.com/). Valid values:
	//
	// *   `true`: enables the mesh audit feature.
	// *   `false`: disables the mesh audit feature.
	//
	// Default value: `false`.
	EnableAudit *bool `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	// The ports for which outbound traffic is redirected to the sidecar proxy.
	EnableAutoDiagnosis *bool `json:"EnableAutoDiagnosis,omitempty" xml:"EnableAutoDiagnosis,omitempty"`
	// Specifies the authentication token of an ARMS Prometheus instance when the Mesh Topology feature is enabled and ARMS Prometheus is used to collect monitoring metrics. The token is used to allow Mesh Topology to access the ARMS Prometheus instance. The token is in the JSON format. The key in the JSON object is the ID of the cluster on the data plane, and the value is the authentication token of the ARMS Prometheus instance deployed in the cluster.
	EnableBootstrapXdsAgent *bool `json:"EnableBootstrapXdsAgent,omitempty" xml:"EnableBootstrapXdsAgent,omitempty"`
	// Specifies whether to enable the rollback feature for Istio resources.
	EnableCRHistory *bool `json:"EnableCRHistory,omitempty" xml:"EnableCRHistory,omitempty"`
	// Specifies whether to enable automatic sidecar proxy injection for all namespaces. Valid values:
	//
	// *   `true`: enables automatic sidecar proxy injection for all namespaces.
	// *   `false`: disables automatic sidecar proxy injection for all namespaces.
	//
	// Default value: `false`.
	EnableNamespacesByDefault *bool `json:"EnableNamespacesByDefault,omitempty" xml:"EnableNamespacesByDefault,omitempty"`
	// Specifies whether to enable Secret Discovery Service (SDS). Valid values:
	//
	// *   `true`: enables SDS.
	// *   `false`: disables SDS.
	//
	// Default value: `false`.
	EnableSDSServer *bool `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
	// The IP addresses of external services to which traffic is not intercepted.
	ExcludeIPRanges *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	// The ports for which inbound traffic is not redirected to the sidecar proxy. Separate multiple ports with commas (,).
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The ports for which outbound traffic is not redirected to the sidecar proxy. Separate multiple ports with commas (,).
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// Specifies whether to enable gateway configuration filtering. Valid values:
	//
	// *   `true`: enables gateway configuration filtering.
	// *   `false`: disables gateway configuration filtering.
	//
	// Default value: `false`.
	FilterGatewayClusterConfig *bool `json:"FilterGatewayClusterConfig,omitempty" xml:"FilterGatewayClusterConfig,omitempty"`
	// Specifies whether to enable Gateway API. Valid values:
	//
	// *   `true`: enables Gateway API.
	// *   `false`: disables Gateway API.
	//
	// Default value: `false`.
	GatewayAPIEnabled *bool `json:"GatewayAPIEnabled,omitempty" xml:"GatewayAPIEnabled,omitempty"`
	// Other metrics of the sidecar proxy on the data plane.
	HoldApplicationUntilProxyStarts *bool `json:"HoldApplicationUntilProxyStarts,omitempty" xml:"HoldApplicationUntilProxyStarts,omitempty"`
	// Specifies whether to support HTTP 1.0. Valid values:
	//
	// *   `true`: supports HTTP 1.0.
	// *   `false`: does not support HTTP 1.0.
	//
	// Default value: `false`.
	Http10Enabled *bool `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
	// The IP addresses of external services to which traffic is intercepted.
	IncludeIPRanges *string `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	// The ports for which inbound traffic is redirected to the sidecar proxy.
	IncludeInboundPorts *string `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	// The log level of the sidecar proxy on the data plane. Log levels include `none`, `error`, `warn`, `info`, and `debug`. The levels correspond to different amounts of information reported by the logs. For example, none-level logs report the least information, while debug-level logs report the most information.
	IncludeOutboundPorts *string `json:"IncludeOutboundPorts,omitempty" xml:"IncludeOutboundPorts,omitempty"`
	// Specifies whether to enable Node Feature Discovery (NFD).
	IntegrateKiali *bool `json:"IntegrateKiali,omitempty" xml:"IntegrateKiali,omitempty"`
	// Specifies whether to load the bootstrap configuration before the sidecar proxy is started.
	InterceptionMode *string `json:"InterceptionMode,omitempty" xml:"InterceptionMode,omitempty"`
	// Specifies the default scheduling configurations that ASM delivers to components on the data plane. You can configure `nodeSelector` and tolerations in the JSON format.
	//
	// > *   Modifying the value of this parameter is a high-risk operation. The modification will cause all components on the data plane of ASM to restart. Exercise caution before modifying the value of this parameter.
	// >*   The configurations specified by this parameter do not apply to the ASM gateway. You can configure gateway-specific scheduling on the ASM gateway.
	KialiArmsAuthTokens *string `json:"KialiArmsAuthTokens,omitempty" xml:"KialiArmsAuthTokens,omitempty"`
	// Specifies whether to enable the Mesh Topology feature. To enable this feature, make sure that you have enabled Prometheus monitoring. If Prometheus monitoring is disabled, the Mesh Topology feature must be disabled. Valid values:````
	//
	// *   `true`: enables the Mesh Topology feature.
	// *   `false`: disables the Mesh Topology feature.
	//
	// Default value: `false`.
	KialiEnabled *bool `json:"KialiEnabled,omitempty" xml:"KialiEnabled,omitempty"`
	// Specifies Classic Load Balancer (CLB) instances by using annotations when the Mesh Topology feature is enabled. These CLB instances are used to access the Mesh Topology feature in different clusters.
	//
	// This parameter is a JSON-encoded string. The key in the JSON object is the ID of a cluster on the data plane, and the value is the annotation content of the Mesh Topology service in the cluster.
	//
	// For more information about how to configure CLB instances by using annotations, see [Add annotations to the YAML file of a Service to configure CLB instances](https://www.alibabacloud.com/help/container-service-for-kubernetes/latest/use-annotations-to-configure-load-balancing-1).
	KialiServiceAnnotations *string `json:"KialiServiceAnnotations,omitempty" xml:"KialiServiceAnnotations,omitempty"`
	// The lifecycle of the sidecar proxy.
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The configurations of cross-region load balancing. Valid values:
	//
	// *   `failover`: the configurations of cross-region failover. Example:
	//
	// <!---->
	//
	//     failover: [// Cross-region failover configurations of the struct type.
	//             {
	//                 // If a service fails in the China (Beijing) region, the traffic is redirected to the same service in the China (Hangzhou) region.
	//                 from: "cn-beijing",
	//                 to: "cn-hangzhou",
	//             }
	//         ]
	//
	// *   `distribute`: the configurations of cross-region traffic distribution. Example:
	//
	// <!---->
	//
	//     distribute: [// Cross-region traffic distribution configurations of the struct type.
	//             {
	//                 // For traffic that is routed to the China (Beijing) region, 70% of the traffic is allocated to the China (Beijing) region and the rest of the traffic is redirected to the China (Hangzhou) region.
	//                 "from": "cn-beijing",
	//                 "to": {
	//                     "cn-beijing": 70,
	//                     "cn-hangzhou": 30,
	//                 }
	//             }
	//         ]
	LocalityLBConf *string `json:"LocalityLBConf,omitempty" xml:"LocalityLBConf,omitempty"`
	// Specifies whether to enable cross-region load balancing. Valid values:
	//
	// *   `true`: enables cross-region load balancing.
	// *   `false`: disables cross-region load balancing.
	//
	// Default value: `false`.
	LocalityLoadBalancing *bool `json:"LocalityLoadBalancing,omitempty" xml:"LocalityLoadBalancing,omitempty"`
	// The number of worker threads used by the sidecar proxy on the data plane.
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// Specifies whether to enable Microservice Engine (MSE). Valid values:
	//
	// *   `true`: enables MSE.
	// *   `false`: disables MSE.
	//
	// Default value: `false`.
	MSEEnabled *bool `json:"MSEEnabled,omitempty" xml:"MSEEnabled,omitempty"`
	// Specifies whether to enable Transport Layer Security (TLS) acceleration based on MultiBuffer.
	MultiBufferEnabled *bool `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	// The pull-request latency. By default, this parameter is left empty.
	MultiBufferPollDelay *string `json:"MultiBufferPollDelay,omitempty" xml:"MultiBufferPollDelay,omitempty"`
	// Specifies whether to enable MySQL Filter. Valid values:
	//
	// *   `true`: enables MySQL Filter.
	// *   `false`: disables MySQL Filter.
	//
	// Default value: `false`.
	MysqlFilterEnabled *bool `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	// Specifies whether to clear feature labels on nodes when NFD is disabled.
	//
	// This parameter is valid only when the `NFDEnabled` parameter is set to `false`.
	NFDEnabled *bool `json:"NFDEnabled,omitempty" xml:"NFDEnabled,omitempty"`
	// The minimum number of CPU cores requested by the proxy service that exports Tracing Analysis data. For example, `1000m` indicates one CPU core.
	NFDLabelPruned *bool `json:"NFDLabelPruned,omitempty" xml:"NFDLabelPruned,omitempty"`
	// The maximum size of the memory that is available to the pod that injects OPA proxies into application pods. For example, `1024Mi` indicates 1024 MB.
	OPAInjectorCPULimit *string `json:"OPAInjectorCPULimit,omitempty" xml:"OPAInjectorCPULimit,omitempty"`
	// The minimum size of the memory requested by the pod that injects OPA proxies into application pods. For example, `50 Mi` indicates 50 MB.
	OPAInjectorCPURequirement *string `json:"OPAInjectorCPURequirement,omitempty" xml:"OPAInjectorCPURequirement,omitempty"`
	// Specifies whether to create an SLB instance for accessing the ASM mesh topology.
	OPAInjectorMemoryLimit *string `json:"OPAInjectorMemoryLimit,omitempty" xml:"OPAInjectorMemoryLimit,omitempty"`
	// The maximum number of CPU cores that are available to the pod that injects OPA proxies into application pods. For example, `1000m` indicates one CPU core.
	OPAInjectorMemoryRequirement *string `json:"OPAInjectorMemoryRequirement,omitempty" xml:"OPAInjectorMemoryRequirement,omitempty"`
	// The maximum number of CPU cores that are available to the OPA proxy container.
	OPALimitCPU *string `json:"OPALimitCPU,omitempty" xml:"OPALimitCPU,omitempty"`
	// The maximum size of the memory that is available to the OPA proxy container.
	OPALimitMemory *string `json:"OPALimitMemory,omitempty" xml:"OPALimitMemory,omitempty"`
	// The log level of the OPA proxy container.
	//
	// *   `info`: outputs all information.
	// *   `debug`: outputs debugging and error information.
	// *   `error`: outputs only error information.
	OPALogLevel *string `json:"OPALogLevel,omitempty" xml:"OPALogLevel,omitempty"`
	// The number of CPU cores that are requested by the OPA proxy container.
	OPARequestCPU *string `json:"OPARequestCPU,omitempty" xml:"OPARequestCPU,omitempty"`
	// The size of the memory that is requested by the OPA proxy container.
	OPARequestMemory *string `json:"OPARequestMemory,omitempty" xml:"OPARequestMemory,omitempty"`
	// The minimum number of CPU cores requested by the pod that injects OPA proxies into application pods. For example, `1000m` indicates one CPU core.
	OPAScopeInjected *bool `json:"OPAScopeInjected,omitempty" xml:"OPAScopeInjected,omitempty"`
	// Specifies whether to enable the OPA plug-in. Valid values:
	//
	// *   `true`: enables the OPA plug-in.
	// *   `false`: disables the OPA plug-in.
	//
	// Default value: `false`.
	OpaEnabled *bool `json:"OpaEnabled,omitempty" xml:"OpaEnabled,omitempty"`
	// Specifies whether to install the Open Policy Agent (OPA) plug-in. Valid values:
	//
	// *   `true`: installs the OPA plug-in.
	// *   `false`: does not install the OPA plug-in.
	//
	// Default value: `false`.
	OpenAgentPolicy *bool `json:"OpenAgentPolicy,omitempty" xml:"OpenAgentPolicy,omitempty"`
	// The policy for accessing external services. Valid values:
	//
	// *   `ALLOW_ANY`: allows access to all external services.
	// *   `REGISTRY_ONLY`: allows access to only the external services that are defined in the ServiceEntry of the ASM instance.
	OutboundTrafficPolicy *string `json:"OutboundTrafficPolicy,omitempty" xml:"OutboundTrafficPolicy,omitempty"`
	// The endpoint of Prometheus monitoring. If you use ARMS Prometheus, set this parameter to the endpoint of Prometheus provided by ARMS.
	PrometheusUrl *string `json:"PrometheusUrl,omitempty" xml:"PrometheusUrl,omitempty"`
	// The maximum number of CPU cores that are available to the istio-init container.
	ProxyInitCPUResourceLimit *string `json:"ProxyInitCPUResourceLimit,omitempty" xml:"ProxyInitCPUResourceLimit,omitempty"`
	// The number of CPU cores that are requested by the istio-init container.
	ProxyInitCPUResourceRequest *string `json:"ProxyInitCPUResourceRequest,omitempty" xml:"ProxyInitCPUResourceRequest,omitempty"`
	// The maximum size of the memory that is available to the istio-init container.
	ProxyInitMemoryResourceLimit *string `json:"ProxyInitMemoryResourceLimit,omitempty" xml:"ProxyInitMemoryResourceLimit,omitempty"`
	// The size of the memory that is requested by the istio-init container.
	ProxyInitMemoryResourceRequest *string `json:"ProxyInitMemoryResourceRequest,omitempty" xml:"ProxyInitMemoryResourceRequest,omitempty"`
	// The maximum number of CPU cores that are available to the sidecar proxy container.
	ProxyLimitCPU *string `json:"ProxyLimitCPU,omitempty" xml:"ProxyLimitCPU,omitempty"`
	// The maximum size of the memory that is available to the sidecar proxy container.
	ProxyLimitMemory *string `json:"ProxyLimitMemory,omitempty" xml:"ProxyLimitMemory,omitempty"`
	// The number of CPU cores that are requested by the sidecar proxy container.
	ProxyRequestCPU *string `json:"ProxyRequestCPU,omitempty" xml:"ProxyRequestCPU,omitempty"`
	// The size of the memory that is requested by the sidecar proxy container.
	ProxyRequestMemory *string `json:"ProxyRequestMemory,omitempty" xml:"ProxyRequestMemory,omitempty"`
	// The mode in which the sidecar proxy intercepts inbound traffic. Valid values:
	//
	// *   `REDIRECT`: The sidecar proxy intercepts inbound traffic in the REDIRECT mode.
	// *   `TPROXY`: The sidecar proxy intercepts inbound traffic in the TPROXY mode.
	ProxyStatsMatcher *string `json:"ProxyStatsMatcher,omitempty" xml:"ProxyStatsMatcher,omitempty"`
	// Specifies whether to enable Redis Filter. Valid values:
	//
	// *   `true`: enables Redis Filter.
	// *   `false`: disables Redis Filter.
	//
	// Default value: `false`.
	RedisFilterEnabled *bool `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The maximum number of CPU cores that are available to the pod where a sidecar proxy injector resides.
	SidecarInjectorLimitCPU *string `json:"SidecarInjectorLimitCPU,omitempty" xml:"SidecarInjectorLimitCPU,omitempty"`
	// The maximum size of the memory that is available to the pod where a sidecar proxy injector resides.
	SidecarInjectorLimitMemory *string `json:"SidecarInjectorLimitMemory,omitempty" xml:"SidecarInjectorLimitMemory,omitempty"`
	// The number of CPU cores that are requested by the pod where a sidecar proxy injector resides.
	SidecarInjectorRequestCPU *string `json:"SidecarInjectorRequestCPU,omitempty" xml:"SidecarInjectorRequestCPU,omitempty"`
	// The size of the memory that is requested by the pod where a sidecar proxy injector resides.
	SidecarInjectorRequestMemory *string `json:"SidecarInjectorRequestMemory,omitempty" xml:"SidecarInjectorRequestMemory,omitempty"`
	// Other configurations of automatic sidecar proxy injection, in the YAML format.
	SidecarInjectorWebhookAsYaml *string `json:"SidecarInjectorWebhookAsYaml,omitempty" xml:"SidecarInjectorWebhookAsYaml,omitempty"`
	// Specifies whether to enable Prometheus monitoring. We recommend that you enable [ARMS Prometheus](https://arms.console.aliyun.com/). Valid values:
	//
	// *   `true`: enables Prometheus monitoring.
	// *   `false`: disables Prometheus monitoring.
	//
	// Default value: `false`.
	Telemetry *bool `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	// The maximum period of time that the sidecar proxy waits for requests to be processed before the proxy is stopped. For example, if you want to specify a period of 5 seconds, set this parameter to 5s.
	TerminationDrainDuration *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
	// Specifies whether to enable Thrift Filter. Valid values:
	//
	// *   `true`: enables Thrift Filter.
	// *   `false`: disables Thrift Filter.
	//
	// Default value: `false`.
	ThriftFilterEnabled *bool `json:"ThriftFilterEnabled,omitempty" xml:"ThriftFilterEnabled,omitempty"`
	// The custom tag of Tracing Analysis. Specify this parameter in the JSON format.
	//
	//     {
	//         "name1": CustomTag,
	//         "name2": CustomTag
	//     }
	//
	// Tag key: literal, header, or environment.
	//
	//     {
	//         "literal": {
	//             "value": "Fixed value"
	//         }
	//         "header": {
	//             "name": "Header name"
	//             "defaultValue": "Default value that is used if the specified header does not exist"
	//         }
	//         "environment": {
	//             "name": "Environment variable name"
	//             "defaultValue": "Default value that is used if the specified environment variable does not exist"
	//         }
	//     }
	TraceCustomTags *string `json:"TraceCustomTags,omitempty" xml:"TraceCustomTags,omitempty"`
	// The maximum length of the request path contained in the HttpUrl span tag. Default value: `256`.
	TraceMaxPathTagLength *string `json:"TraceMaxPathTagLength,omitempty" xml:"TraceMaxPathTagLength,omitempty"`
	// The sampling percentage of Tracing Analysis.
	TraceSampling *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
	// Specifies whether to enable the Tracing Analysis feature. To enable this feature, make sure that you have activated [Tracing Analysis](https://tracing-analysis.console.aliyun.com/). Valid values:
	//
	// *   `true`: enables the Tracing Analysis feature.
	// *   `false`: disables the Tracing Analysis feature.
	//
	// Default value: `false`.
	Tracing *bool `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
	// The maximum size of the memory that is available to the proxy service that exports Tracing Analysis data. For example, `1Mi` indicates 1 MB.
	TracingOnExtZipkinLimitCPU *string `json:"TracingOnExtZipkinLimitCPU,omitempty" xml:"TracingOnExtZipkinLimitCPU,omitempty"`
	// The retention period for the access logs of the ingress gateway. Unit: day. The logs are collected by using Log Service. For example, `30` indicates 30 days.
	TracingOnExtZipkinLimitMemory *string `json:"TracingOnExtZipkinLimitMemory,omitempty" xml:"TracingOnExtZipkinLimitMemory,omitempty"`
	// The minimum size of the memory requested by the proxy service that exports Tracing Analysis data. For example, `1Mi` indicates 1 MB.
	TracingOnExtZipkinRequestCPU *string `json:"TracingOnExtZipkinRequestCPU,omitempty" xml:"TracingOnExtZipkinRequestCPU,omitempty"`
	// The maximum number of CPU cores that are available to the proxy service that exports Tracing Analysis data. For example, `1000m` indicates one CPU core.
	TracingOnExtZipkinRequestMemory *string `json:"TracingOnExtZipkinRequestMemory,omitempty" xml:"TracingOnExtZipkinRequestMemory,omitempty"`
	// Specifies whether to enable WebAssembly Filter. Valid values:
	//
	// *   `true`: enables WebAssembly Filter.
	// *   `false`: disables WebAssembly Filter.
	//
	// Default value: `false`.
	WebAssemblyFilterEnabled *bool `json:"WebAssemblyFilterEnabled,omitempty" xml:"WebAssemblyFilterEnabled,omitempty"`
}

func (s UpdateMeshFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeshFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeshFeatureRequest) SetAccessLogEnabled(v bool) *UpdateMeshFeatureRequest {
	s.AccessLogEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogFile(v string) *UpdateMeshFeatureRequest {
	s.AccessLogFile = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogFormat(v string) *UpdateMeshFeatureRequest {
	s.AccessLogFormat = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogGatewayEnabled(v bool) *UpdateMeshFeatureRequest {
	s.AccessLogGatewayEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogGatewayLifecycle(v int32) *UpdateMeshFeatureRequest {
	s.AccessLogGatewayLifecycle = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogProject(v string) *UpdateMeshFeatureRequest {
	s.AccessLogProject = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogServiceEnabled(v bool) *UpdateMeshFeatureRequest {
	s.AccessLogServiceEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogServiceHost(v string) *UpdateMeshFeatureRequest {
	s.AccessLogServiceHost = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogServicePort(v int32) *UpdateMeshFeatureRequest {
	s.AccessLogServicePort = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogSidecarEnabled(v bool) *UpdateMeshFeatureRequest {
	s.AccessLogSidecarEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAccessLogSidecarLifecycle(v int32) *UpdateMeshFeatureRequest {
	s.AccessLogSidecarLifecycle = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAuditProject(v string) *UpdateMeshFeatureRequest {
	s.AuditProject = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetAutoInjectionPolicyEnabled(v bool) *UpdateMeshFeatureRequest {
	s.AutoInjectionPolicyEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetCRAggregationEnabled(v bool) *UpdateMeshFeatureRequest {
	s.CRAggregationEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetClusterSpec(v string) *UpdateMeshFeatureRequest {
	s.ClusterSpec = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetCniEnabled(v bool) *UpdateMeshFeatureRequest {
	s.CniEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetCniExcludeNamespaces(v string) *UpdateMeshFeatureRequest {
	s.CniExcludeNamespaces = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetConcurrency(v int32) *UpdateMeshFeatureRequest {
	s.Concurrency = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetConfigSourceEnabled(v bool) *UpdateMeshFeatureRequest {
	s.ConfigSourceEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetConfigSourceNacosID(v string) *UpdateMeshFeatureRequest {
	s.ConfigSourceNacosID = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetCustomizedPrometheus(v bool) *UpdateMeshFeatureRequest {
	s.CustomizedPrometheus = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetCustomizedZipkin(v bool) *UpdateMeshFeatureRequest {
	s.CustomizedZipkin = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetDNSProxyingEnabled(v bool) *UpdateMeshFeatureRequest {
	s.DNSProxyingEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetDefaultComponentsScheduleConfig(v string) *UpdateMeshFeatureRequest {
	s.DefaultComponentsScheduleConfig = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetDiscoverySelectors(v string) *UpdateMeshFeatureRequest {
	s.DiscoverySelectors = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetDubboFilterEnabled(v bool) *UpdateMeshFeatureRequest {
	s.DubboFilterEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetEnableAudit(v bool) *UpdateMeshFeatureRequest {
	s.EnableAudit = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetEnableAutoDiagnosis(v bool) *UpdateMeshFeatureRequest {
	s.EnableAutoDiagnosis = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetEnableBootstrapXdsAgent(v bool) *UpdateMeshFeatureRequest {
	s.EnableBootstrapXdsAgent = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetEnableCRHistory(v bool) *UpdateMeshFeatureRequest {
	s.EnableCRHistory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetEnableNamespacesByDefault(v bool) *UpdateMeshFeatureRequest {
	s.EnableNamespacesByDefault = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetEnableSDSServer(v bool) *UpdateMeshFeatureRequest {
	s.EnableSDSServer = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetExcludeIPRanges(v string) *UpdateMeshFeatureRequest {
	s.ExcludeIPRanges = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetExcludeInboundPorts(v string) *UpdateMeshFeatureRequest {
	s.ExcludeInboundPorts = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetExcludeOutboundPorts(v string) *UpdateMeshFeatureRequest {
	s.ExcludeOutboundPorts = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetFilterGatewayClusterConfig(v bool) *UpdateMeshFeatureRequest {
	s.FilterGatewayClusterConfig = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetGatewayAPIEnabled(v bool) *UpdateMeshFeatureRequest {
	s.GatewayAPIEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetHoldApplicationUntilProxyStarts(v bool) *UpdateMeshFeatureRequest {
	s.HoldApplicationUntilProxyStarts = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetHttp10Enabled(v bool) *UpdateMeshFeatureRequest {
	s.Http10Enabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetIncludeIPRanges(v string) *UpdateMeshFeatureRequest {
	s.IncludeIPRanges = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetIncludeInboundPorts(v string) *UpdateMeshFeatureRequest {
	s.IncludeInboundPorts = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetIncludeOutboundPorts(v string) *UpdateMeshFeatureRequest {
	s.IncludeOutboundPorts = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetIntegrateKiali(v bool) *UpdateMeshFeatureRequest {
	s.IntegrateKiali = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetInterceptionMode(v string) *UpdateMeshFeatureRequest {
	s.InterceptionMode = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetKialiArmsAuthTokens(v string) *UpdateMeshFeatureRequest {
	s.KialiArmsAuthTokens = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetKialiEnabled(v bool) *UpdateMeshFeatureRequest {
	s.KialiEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetKialiServiceAnnotations(v string) *UpdateMeshFeatureRequest {
	s.KialiServiceAnnotations = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetLifecycle(v string) *UpdateMeshFeatureRequest {
	s.Lifecycle = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetLocalityLBConf(v string) *UpdateMeshFeatureRequest {
	s.LocalityLBConf = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetLocalityLoadBalancing(v bool) *UpdateMeshFeatureRequest {
	s.LocalityLoadBalancing = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetLogLevel(v string) *UpdateMeshFeatureRequest {
	s.LogLevel = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetMSEEnabled(v bool) *UpdateMeshFeatureRequest {
	s.MSEEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetMultiBufferEnabled(v bool) *UpdateMeshFeatureRequest {
	s.MultiBufferEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetMultiBufferPollDelay(v string) *UpdateMeshFeatureRequest {
	s.MultiBufferPollDelay = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetMysqlFilterEnabled(v bool) *UpdateMeshFeatureRequest {
	s.MysqlFilterEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetNFDEnabled(v bool) *UpdateMeshFeatureRequest {
	s.NFDEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetNFDLabelPruned(v bool) *UpdateMeshFeatureRequest {
	s.NFDLabelPruned = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPAInjectorCPULimit(v string) *UpdateMeshFeatureRequest {
	s.OPAInjectorCPULimit = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPAInjectorCPURequirement(v string) *UpdateMeshFeatureRequest {
	s.OPAInjectorCPURequirement = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPAInjectorMemoryLimit(v string) *UpdateMeshFeatureRequest {
	s.OPAInjectorMemoryLimit = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPAInjectorMemoryRequirement(v string) *UpdateMeshFeatureRequest {
	s.OPAInjectorMemoryRequirement = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPALimitCPU(v string) *UpdateMeshFeatureRequest {
	s.OPALimitCPU = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPALimitMemory(v string) *UpdateMeshFeatureRequest {
	s.OPALimitMemory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPALogLevel(v string) *UpdateMeshFeatureRequest {
	s.OPALogLevel = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPARequestCPU(v string) *UpdateMeshFeatureRequest {
	s.OPARequestCPU = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPARequestMemory(v string) *UpdateMeshFeatureRequest {
	s.OPARequestMemory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOPAScopeInjected(v bool) *UpdateMeshFeatureRequest {
	s.OPAScopeInjected = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOpaEnabled(v bool) *UpdateMeshFeatureRequest {
	s.OpaEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOpenAgentPolicy(v bool) *UpdateMeshFeatureRequest {
	s.OpenAgentPolicy = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetOutboundTrafficPolicy(v string) *UpdateMeshFeatureRequest {
	s.OutboundTrafficPolicy = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetPrometheusUrl(v string) *UpdateMeshFeatureRequest {
	s.PrometheusUrl = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyInitCPUResourceLimit(v string) *UpdateMeshFeatureRequest {
	s.ProxyInitCPUResourceLimit = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyInitCPUResourceRequest(v string) *UpdateMeshFeatureRequest {
	s.ProxyInitCPUResourceRequest = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyInitMemoryResourceLimit(v string) *UpdateMeshFeatureRequest {
	s.ProxyInitMemoryResourceLimit = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyInitMemoryResourceRequest(v string) *UpdateMeshFeatureRequest {
	s.ProxyInitMemoryResourceRequest = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyLimitCPU(v string) *UpdateMeshFeatureRequest {
	s.ProxyLimitCPU = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyLimitMemory(v string) *UpdateMeshFeatureRequest {
	s.ProxyLimitMemory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyRequestCPU(v string) *UpdateMeshFeatureRequest {
	s.ProxyRequestCPU = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyRequestMemory(v string) *UpdateMeshFeatureRequest {
	s.ProxyRequestMemory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetProxyStatsMatcher(v string) *UpdateMeshFeatureRequest {
	s.ProxyStatsMatcher = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetRedisFilterEnabled(v bool) *UpdateMeshFeatureRequest {
	s.RedisFilterEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetServiceMeshId(v string) *UpdateMeshFeatureRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetSidecarInjectorLimitCPU(v string) *UpdateMeshFeatureRequest {
	s.SidecarInjectorLimitCPU = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetSidecarInjectorLimitMemory(v string) *UpdateMeshFeatureRequest {
	s.SidecarInjectorLimitMemory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetSidecarInjectorRequestCPU(v string) *UpdateMeshFeatureRequest {
	s.SidecarInjectorRequestCPU = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetSidecarInjectorRequestMemory(v string) *UpdateMeshFeatureRequest {
	s.SidecarInjectorRequestMemory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetSidecarInjectorWebhookAsYaml(v string) *UpdateMeshFeatureRequest {
	s.SidecarInjectorWebhookAsYaml = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTelemetry(v bool) *UpdateMeshFeatureRequest {
	s.Telemetry = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTerminationDrainDuration(v string) *UpdateMeshFeatureRequest {
	s.TerminationDrainDuration = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetThriftFilterEnabled(v bool) *UpdateMeshFeatureRequest {
	s.ThriftFilterEnabled = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTraceCustomTags(v string) *UpdateMeshFeatureRequest {
	s.TraceCustomTags = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTraceMaxPathTagLength(v string) *UpdateMeshFeatureRequest {
	s.TraceMaxPathTagLength = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTraceSampling(v float32) *UpdateMeshFeatureRequest {
	s.TraceSampling = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTracing(v bool) *UpdateMeshFeatureRequest {
	s.Tracing = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTracingOnExtZipkinLimitCPU(v string) *UpdateMeshFeatureRequest {
	s.TracingOnExtZipkinLimitCPU = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTracingOnExtZipkinLimitMemory(v string) *UpdateMeshFeatureRequest {
	s.TracingOnExtZipkinLimitMemory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTracingOnExtZipkinRequestCPU(v string) *UpdateMeshFeatureRequest {
	s.TracingOnExtZipkinRequestCPU = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTracingOnExtZipkinRequestMemory(v string) *UpdateMeshFeatureRequest {
	s.TracingOnExtZipkinRequestMemory = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetWebAssemblyFilterEnabled(v bool) *UpdateMeshFeatureRequest {
	s.WebAssemblyFilterEnabled = &v
	return s
}

type UpdateMeshFeatureResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMeshFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeshFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeshFeatureResponseBody) SetRequestId(v string) *UpdateMeshFeatureResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMeshFeatureResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMeshFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMeshFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeshFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeshFeatureResponse) SetHeaders(v map[string]*string) *UpdateMeshFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeshFeatureResponse) SetStatusCode(v int32) *UpdateMeshFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeshFeatureResponse) SetBody(v *UpdateMeshFeatureResponseBody) *UpdateMeshFeatureResponse {
	s.Body = v
	return s
}

type UpdateNamespaceScopeSidecarConfigRequest struct {
	// The number of worker threads to run in Istio Proxy.
	Concurrency    *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	EnableCoreDump *bool  `json:"EnableCoreDump,omitempty" xml:"EnableCoreDump,omitempty"`
	// The range of IP addresses that are allowed to access external services. (`global.proxy.excludelPRanges`)
	ExcludeIPRanges *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	// The port that the inbound traffic of the sidecar proxy does not pass through.
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The port that the outbound traffic of the sidecar proxy does not pass through.
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// Specifies whether applications can be started only after Istio Proxy starts. Valid values:
	//
	// *   `true`: Applications can be started only after Istio Proxy starts.
	// *   `false`: Applications can be started before Istio Proxy starts.
	HoldApplicationUntilProxyStarts *bool `json:"HoldApplicationUntilProxyStarts,omitempty" xml:"HoldApplicationUntilProxyStarts,omitempty"`
	// The range of IP addresses that are denied to access external services. (`global.proxy.includelPRanges`)
	IncludeIPRanges *string `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	// The port that the inbound traffic of the sidecar proxy passes through.
	IncludeInboundPorts *string `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	// The port that the outbound traffic of the sidecar proxy passes through.
	IncludeOutboundPorts *string `json:"IncludeOutboundPorts,omitempty" xml:"IncludeOutboundPorts,omitempty"`
	// The mode in which the sidecar proxy intercepts inbound traffic. Valid values:
	//
	// *   `REDIRECT`: The sidecar proxy intercepts inbound traffic in the REDIRECT mode.
	// *   `TPROXY`: The sidecar proxy intercepts inbound traffic in the TPROXY mode.
	InterceptionMode *string `json:"InterceptionMode,omitempty" xml:"InterceptionMode,omitempty"`
	// Specifies whether to enable the Domain Name System (DNS) proxy feature. Valid values:
	//
	// *   `true`: The DNS proxy feature is enabled.
	// *   `false`: The DNS proxy feature is disabled.
	IstioDNSProxyEnabled *bool `json:"IstioDNSProxyEnabled,omitempty" xml:"IstioDNSProxyEnabled,omitempty"`
	// The lifecycle of the sidecar proxy.
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The log level. Valid values: `info`, `debug`, `tracing`, and `error`.
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The namespace for which you want to update the sidecar proxy configurations.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The post-start parameters of Istio Proxy.
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The pre-close parameters of Istio Proxy.
	PreStop                              *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	Privileged                           *bool   `json:"Privileged,omitempty" xml:"Privileged,omitempty"`
	ProxyInitAckSloCPUResourceLimit      *string `json:"ProxyInitAckSloCPUResourceLimit,omitempty" xml:"ProxyInitAckSloCPUResourceLimit,omitempty"`
	ProxyInitAckSloCPUResourceRequest    *string `json:"ProxyInitAckSloCPUResourceRequest,omitempty" xml:"ProxyInitAckSloCPUResourceRequest,omitempty"`
	ProxyInitAckSloMemoryResourceLimit   *string `json:"ProxyInitAckSloMemoryResourceLimit,omitempty" xml:"ProxyInitAckSloMemoryResourceLimit,omitempty"`
	ProxyInitAckSloMemoryResourceRequest *string `json:"ProxyInitAckSloMemoryResourceRequest,omitempty" xml:"ProxyInitAckSloMemoryResourceRequest,omitempty"`
	// The maximum number of CPU cores that are available to the sidecar proxy init container.
	ProxyInitCPUResourceLimit *string `json:"ProxyInitCPUResourceLimit,omitempty" xml:"ProxyInitCPUResourceLimit,omitempty"`
	// The minimum number of CPU cores that are requested by the sidecar proxy init container.
	ProxyInitCPUResourceRequest *string `json:"ProxyInitCPUResourceRequest,omitempty" xml:"ProxyInitCPUResourceRequest,omitempty"`
	// The maximum size of memory that is available to the sidecar proxy init container.
	ProxyInitMemoryResourceLimit *string `json:"ProxyInitMemoryResourceLimit,omitempty" xml:"ProxyInitMemoryResourceLimit,omitempty"`
	// The minimum size of memory that is requested by the sidecar proxy init container.
	ProxyInitMemoryResourceRequest *string `json:"ProxyInitMemoryResourceRequest,omitempty" xml:"ProxyInitMemoryResourceRequest,omitempty"`
	// The environment variables that are added to a sidecar proxy. The environment variables are represented as JSON objects. The keys and values in the JSON objects represent the keys and values added to the environment variables of the sidecar proxy.
	ProxyMetadata *string `json:"ProxyMetadata,omitempty" xml:"ProxyMetadata,omitempty"`
	// The monitoring metrics for data collected by Envoy proxies. The value is in the JSON format.
	ProxyStatsMatcher            *string `json:"ProxyStatsMatcher,omitempty" xml:"ProxyStatsMatcher,omitempty"`
	ReadinessFailureThreshold    *int32  `json:"ReadinessFailureThreshold,omitempty" xml:"ReadinessFailureThreshold,omitempty"`
	ReadinessInitialDelaySeconds *int32  `json:"ReadinessInitialDelaySeconds,omitempty" xml:"ReadinessInitialDelaySeconds,omitempty"`
	ReadinessPeriodSeconds       *int32  `json:"ReadinessPeriodSeconds,omitempty" xml:"ReadinessPeriodSeconds,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId                           *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	SidecarProxyAckSloCPUResourceLimit      *string `json:"SidecarProxyAckSloCPUResourceLimit,omitempty" xml:"SidecarProxyAckSloCPUResourceLimit,omitempty"`
	SidecarProxyAckSloCPUResourceRequest    *string `json:"SidecarProxyAckSloCPUResourceRequest,omitempty" xml:"SidecarProxyAckSloCPUResourceRequest,omitempty"`
	SidecarProxyAckSloMemoryResourceLimit   *string `json:"SidecarProxyAckSloMemoryResourceLimit,omitempty" xml:"SidecarProxyAckSloMemoryResourceLimit,omitempty"`
	SidecarProxyAckSloMemoryResourceRequest *string `json:"SidecarProxyAckSloMemoryResourceRequest,omitempty" xml:"SidecarProxyAckSloMemoryResourceRequest,omitempty"`
	// The maximum number of CPU cores that are available to the sidecar proxy container.
	SidecarProxyCPUResourceLimit *string `json:"SidecarProxyCPUResourceLimit,omitempty" xml:"SidecarProxyCPUResourceLimit,omitempty"`
	// The minimum number of CPU cores that are requested by the sidecar proxy container.
	SidecarProxyCPUResourceRequest *string `json:"SidecarProxyCPUResourceRequest,omitempty" xml:"SidecarProxyCPUResourceRequest,omitempty"`
	// The maximum size of memory that is available to the sidecar proxy container.
	SidecarProxyMemoryResourceLimit *string `json:"SidecarProxyMemoryResourceLimit,omitempty" xml:"SidecarProxyMemoryResourceLimit,omitempty"`
	// The minimum size of memory that is requested by the sidecar proxy container.
	SidecarProxyMemoryResourceRequest *string `json:"SidecarProxyMemoryResourceRequest,omitempty" xml:"SidecarProxyMemoryResourceRequest,omitempty"`
	// The maximum period of time that the sidecar proxy waits for a request to end.
	TerminationDrainDuration *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
	// The custom configurations of Tracing Analysis. The configurations must be serialized into JSON strings. The configurations contain the following parameters:
	//
	// *   `sampling`: The sampling rate, which is of the DOUBLE type.
	//
	// *   `custom_tags`: The custom tags added to reported spans, which are of the MAP type. The key of a tag is of the string type. The value of a tag is in the JSON format. A custom tag can belong to one of the following types:
	//
	//     *   `literal`: The tag value is a fixed value in the JSON format. This tag must contain the `value` field that specifies a literal. Example: `{"value":"test"}`.
	//     *   `header`: The tag value is a request header in the JSON format. This tag must contain the `name` field and `defaultValue` field.The name field indicates the name of the request header. The defaultValue field indicates the default value that is used when no request header is available. Example: `{"name":"test","defaultValue":"test"}`.
	//     *   `environment`: The tag value is an environment variable in the JSON format. This tag must contain the `name` field and `defaultValue` field. The name field indicates the name of the environment variable. The defaultValue field indicates the environment variable that is used when no environment variable is available. Example: `{"name":"test","defaultValue":"test"}`.
	Tracing *string `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
}

func (s UpdateNamespaceScopeSidecarConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetConcurrency(v int32) *UpdateNamespaceScopeSidecarConfigRequest {
	s.Concurrency = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetEnableCoreDump(v bool) *UpdateNamespaceScopeSidecarConfigRequest {
	s.EnableCoreDump = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetExcludeIPRanges(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ExcludeIPRanges = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetExcludeInboundPorts(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ExcludeInboundPorts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetExcludeOutboundPorts(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ExcludeOutboundPorts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetHoldApplicationUntilProxyStarts(v bool) *UpdateNamespaceScopeSidecarConfigRequest {
	s.HoldApplicationUntilProxyStarts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetIncludeIPRanges(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.IncludeIPRanges = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetIncludeInboundPorts(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.IncludeInboundPorts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetIncludeOutboundPorts(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.IncludeOutboundPorts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetInterceptionMode(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.InterceptionMode = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetIstioDNSProxyEnabled(v bool) *UpdateNamespaceScopeSidecarConfigRequest {
	s.IstioDNSProxyEnabled = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetLifecycle(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.Lifecycle = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetLogLevel(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.LogLevel = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetNamespace(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetPostStart(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.PostStart = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetPreStop(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.PreStop = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetPrivileged(v bool) *UpdateNamespaceScopeSidecarConfigRequest {
	s.Privileged = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyInitAckSloCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyInitAckSloCPUResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyInitAckSloCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyInitAckSloCPUResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyInitAckSloMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyInitAckSloMemoryResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyInitAckSloMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyInitAckSloMemoryResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyInitCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyInitCPUResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyInitCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyInitCPUResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyInitMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyInitMemoryResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyInitMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyInitMemoryResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyMetadata(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyMetadata = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetProxyStatsMatcher(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ProxyStatsMatcher = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetReadinessFailureThreshold(v int32) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ReadinessFailureThreshold = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetReadinessInitialDelaySeconds(v int32) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ReadinessInitialDelaySeconds = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetReadinessPeriodSeconds(v int32) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ReadinessPeriodSeconds = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetServiceMeshId(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSidecarProxyAckSloCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SidecarProxyAckSloCPUResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSidecarProxyAckSloCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SidecarProxyAckSloCPUResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSidecarProxyAckSloMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SidecarProxyAckSloMemoryResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSidecarProxyAckSloMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SidecarProxyAckSloMemoryResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSidecarProxyCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SidecarProxyCPUResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSidecarProxyCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SidecarProxyCPUResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSidecarProxyMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SidecarProxyMemoryResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSidecarProxyMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SidecarProxyMemoryResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetTerminationDrainDuration(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.TerminationDrainDuration = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetTracing(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.Tracing = &v
	return s
}

type UpdateNamespaceScopeSidecarConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNamespaceScopeSidecarConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceScopeSidecarConfigResponseBody) SetRequestId(v string) *UpdateNamespaceScopeSidecarConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNamespaceScopeSidecarConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNamespaceScopeSidecarConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNamespaceScopeSidecarConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) SetHeaders(v map[string]*string) *UpdateNamespaceScopeSidecarConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) SetStatusCode(v int32) *UpdateNamespaceScopeSidecarConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigResponse) SetBody(v *UpdateNamespaceScopeSidecarConfigResponseBody) *UpdateNamespaceScopeSidecarConfigResponse {
	s.Body = v
	return s
}

type UpdateSwimLaneRequest struct {
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The label key of the associated service workload. Set the value to `ASM_TRAFFIC_TAG`.
	LabelSelectorKey *string `json:"LabelSelectorKey,omitempty" xml:"LabelSelectorKey,omitempty"`
	// The label value of the associated service workload.``
	LabelSelectorValue *string `json:"LabelSelectorValue,omitempty" xml:"LabelSelectorValue,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// A list of services associated with the lane.
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
	// The name of the lane.
	SwimLaneName *string `json:"SwimLaneName,omitempty" xml:"SwimLaneName,omitempty"`
}

func (s UpdateSwimLaneRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSwimLaneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneRequest) SetGroupName(v string) *UpdateSwimLaneRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetLabelSelectorKey(v string) *UpdateSwimLaneRequest {
	s.LabelSelectorKey = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetLabelSelectorValue(v string) *UpdateSwimLaneRequest {
	s.LabelSelectorValue = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetServiceMeshId(v string) *UpdateSwimLaneRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetServicesList(v string) *UpdateSwimLaneRequest {
	s.ServicesList = &v
	return s
}

func (s *UpdateSwimLaneRequest) SetSwimLaneName(v string) *UpdateSwimLaneRequest {
	s.SwimLaneName = &v
	return s
}

type UpdateSwimLaneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSwimLaneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSwimLaneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneResponseBody) SetRequestId(v string) *UpdateSwimLaneResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSwimLaneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSwimLaneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSwimLaneResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSwimLaneResponse) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneResponse) SetHeaders(v map[string]*string) *UpdateSwimLaneResponse {
	s.Headers = v
	return s
}

func (s *UpdateSwimLaneResponse) SetStatusCode(v int32) *UpdateSwimLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSwimLaneResponse) SetBody(v *UpdateSwimLaneResponseBody) *UpdateSwimLaneResponse {
	s.Body = v
	return s
}

type UpdateSwimLaneGroupRequest struct {
	FallbackTarget *string `json:"FallbackTarget,omitempty" xml:"FallbackTarget,omitempty"`
	// The name of the lane group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// A list of services associated with the lane group.
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
}

func (s UpdateSwimLaneGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSwimLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneGroupRequest) SetFallbackTarget(v string) *UpdateSwimLaneGroupRequest {
	s.FallbackTarget = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetGroupName(v string) *UpdateSwimLaneGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetServiceMeshId(v string) *UpdateSwimLaneGroupRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetServicesList(v string) *UpdateSwimLaneGroupRequest {
	s.ServicesList = &v
	return s
}

type UpdateSwimLaneGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSwimLaneGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSwimLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneGroupResponseBody) SetRequestId(v string) *UpdateSwimLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSwimLaneGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSwimLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSwimLaneGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSwimLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneGroupResponse) SetHeaders(v map[string]*string) *UpdateSwimLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateSwimLaneGroupResponse) SetStatusCode(v int32) *UpdateSwimLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSwimLaneGroupResponse) SetBody(v *UpdateSwimLaneGroupResponseBody) *UpdateSwimLaneGroupResponse {
	s.Body = v
	return s
}

type UpgradeMeshEditionPartiallyRequest struct {
	// Specifies whether to upgrade the ASM gateways for the ASM instance. Valid values:
	//
	// *   `true`
	// *   `false`
	ASMGatewayContinue *bool   `json:"ASMGatewayContinue,omitempty" xml:"ASMGatewayContinue,omitempty"`
	ExpectedVersion    *string `json:"ExpectedVersion,omitempty" xml:"ExpectedVersion,omitempty"`
	// Specifies whether to perform an upgrade check. If the value of this parameter is set to true, only the upgrade check is performed and the ASM instance is not upgraded.
	PreCheck *bool `json:"PreCheck,omitempty" xml:"PreCheck,omitempty"`
	// The ASM instance ID.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// Deprecated
	// Specifies whether to upgrade the ASM instance to Professional Edition. Valid values:
	//
	// *   `true`
	// *   `false`
	SwitchToPro *bool `json:"SwitchToPro,omitempty" xml:"SwitchToPro,omitempty"`
	// Specifies the ASM gateways to be upgraded. Separate multiple ASM gateways with commas (,).
	UpgradeGatewayRecords *string `json:"UpgradeGatewayRecords,omitempty" xml:"UpgradeGatewayRecords,omitempty"`
}

func (s UpgradeMeshEditionPartiallyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMeshEditionPartiallyRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMeshEditionPartiallyRequest) SetASMGatewayContinue(v bool) *UpgradeMeshEditionPartiallyRequest {
	s.ASMGatewayContinue = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetExpectedVersion(v string) *UpgradeMeshEditionPartiallyRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetPreCheck(v bool) *UpgradeMeshEditionPartiallyRequest {
	s.PreCheck = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetServiceMeshId(v string) *UpgradeMeshEditionPartiallyRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetSwitchToPro(v bool) *UpgradeMeshEditionPartiallyRequest {
	s.SwitchToPro = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetUpgradeGatewayRecords(v string) *UpgradeMeshEditionPartiallyRequest {
	s.UpgradeGatewayRecords = &v
	return s
}

type UpgradeMeshEditionPartiallyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMeshEditionPartiallyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMeshEditionPartiallyResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMeshEditionPartiallyResponseBody) SetRequestId(v string) *UpgradeMeshEditionPartiallyResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeMeshEditionPartiallyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeMeshEditionPartiallyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeMeshEditionPartiallyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMeshEditionPartiallyResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMeshEditionPartiallyResponse) SetHeaders(v map[string]*string) *UpgradeMeshEditionPartiallyResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMeshEditionPartiallyResponse) SetStatusCode(v int32) *UpgradeMeshEditionPartiallyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyResponse) SetBody(v *UpgradeMeshEditionPartiallyResponseBody) *UpgradeMeshEditionPartiallyResponse {
	s.Body = v
	return s
}

type UpgradeMeshVersionRequest struct {
	PreCheck *bool `json:"PreCheck,omitempty" xml:"PreCheck,omitempty"`
	// The ID of the ASM instance.
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpgradeMeshVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMeshVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMeshVersionRequest) SetPreCheck(v bool) *UpgradeMeshVersionRequest {
	s.PreCheck = &v
	return s
}

func (s *UpgradeMeshVersionRequest) SetServiceMeshId(v string) *UpgradeMeshVersionRequest {
	s.ServiceMeshId = &v
	return s
}

type UpgradeMeshVersionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMeshVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMeshVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMeshVersionResponseBody) SetRequestId(v string) *UpgradeMeshVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeMeshVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeMeshVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeMeshVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMeshVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMeshVersionResponse) SetHeaders(v map[string]*string) *UpgradeMeshVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMeshVersionResponse) SetStatusCode(v int32) *UpgradeMeshVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMeshVersionResponse) SetBody(v *UpgradeMeshVersionResponseBody) *UpgradeMeshVersionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("servicemesh"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddClusterIntoServiceMeshWithOptions(request *AddClusterIntoServiceMeshRequest, runtime *util.RuntimeOptions) (_result *AddClusterIntoServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreNamespaceCheck)) {
		body["IgnoreNamespaceCheck"] = request.IgnoreNamespaceCheck
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddClusterIntoServiceMesh"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddClusterIntoServiceMeshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddClusterIntoServiceMesh(request *AddClusterIntoServiceMeshRequest) (_result *AddClusterIntoServiceMeshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddClusterIntoServiceMeshResponse{}
	_body, _err := client.AddClusterIntoServiceMeshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request AddVMIntoServiceMeshRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddVMIntoServiceMeshResponse
 */
// Deprecated
func (client *Client) AddVMIntoServiceMeshWithOptions(request *AddVMIntoServiceMeshRequest, runtime *util.RuntimeOptions) (_result *AddVMIntoServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcsId)) {
		query["EcsId"] = request.EcsId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVMIntoServiceMesh"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVMIntoServiceMeshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request AddVMIntoServiceMeshRequest
 * @return AddVMIntoServiceMeshResponse
 */
// Deprecated
func (client *Client) AddVMIntoServiceMesh(request *AddVMIntoServiceMeshRequest) (_result *AddVMIntoServiceMeshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVMIntoServiceMeshResponse{}
	_body, _err := client.AddVMIntoServiceMeshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateASMGatewayWithOptions(request *CreateASMGatewayRequest, runtime *util.RuntimeOptions) (_result *CreateASMGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateASMGateway"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateASMGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateASMGateway(request *CreateASMGatewayRequest) (_result *CreateASMGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateASMGatewayResponse{}
	_body, _err := client.CreateASMGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewaySecretWithOptions(request *CreateGatewaySecretRequest, runtime *util.RuntimeOptions) (_result *CreateGatewaySecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cert)) {
		body["Cert"] = request.Cert
	}

	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		body["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		body["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewaySecret"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewaySecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGatewaySecret(request *CreateGatewaySecretRequest) (_result *CreateGatewaySecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewaySecretResponse{}
	_body, _err := client.CreateGatewaySecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIstioGatewayDomainsWithOptions(request *CreateIstioGatewayDomainsRequest, runtime *util.RuntimeOptions) (_result *CreateIstioGatewayDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Credential)) {
		body["Credential"] = request.Credential
	}

	if !tea.BoolValue(util.IsUnset(request.ForceHttps)) {
		body["ForceHttps"] = request.ForceHttps
	}

	if !tea.BoolValue(util.IsUnset(request.Hosts)) {
		body["Hosts"] = request.Hosts
	}

	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		body["Number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.PortName)) {
		body["PortName"] = request.PortName
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIstioGatewayDomains"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIstioGatewayDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIstioGatewayDomains(request *CreateIstioGatewayDomainsRequest) (_result *CreateIstioGatewayDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIstioGatewayDomainsResponse{}
	_body, _err := client.CreateIstioGatewayDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIstioGatewayRoutesWithOptions(tmpReq *CreateIstioGatewayRoutesRequest, runtime *util.RuntimeOptions) (_result *CreateIstioGatewayRoutesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateIstioGatewayRoutesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GatewayRoute)) {
		request.GatewayRouteShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GatewayRoute, tea.String("GatewayRoute"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayRouteShrink)) {
		body["GatewayRoute"] = request.GatewayRouteShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIstioGatewayRoutes"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIstioGatewayRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIstioGatewayRoutes(request *CreateIstioGatewayRoutesRequest) (_result *CreateIstioGatewayRoutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIstioGatewayRoutesResponse{}
	_body, _err := client.CreateIstioGatewayRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceMeshWithOptions(request *CreateServiceMeshRequest, runtime *util.RuntimeOptions) (_result *CreateServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLogEnabled)) {
		body["AccessLogEnabled"] = request.AccessLogEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogFile)) {
		body["AccessLogFile"] = request.AccessLogFile
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogFormat)) {
		body["AccessLogFormat"] = request.AccessLogFormat
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogProject)) {
		body["AccessLogProject"] = request.AccessLogProject
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogServiceEnabled)) {
		body["AccessLogServiceEnabled"] = request.AccessLogServiceEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogServiceHost)) {
		body["AccessLogServiceHost"] = request.AccessLogServiceHost
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogServicePort)) {
		body["AccessLogServicePort"] = request.AccessLogServicePort
	}

	if !tea.BoolValue(util.IsUnset(request.ApiServerLoadBalancerSpec)) {
		body["ApiServerLoadBalancerSpec"] = request.ApiServerLoadBalancerSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ApiServerPublicEip)) {
		body["ApiServerPublicEip"] = request.ApiServerPublicEip
	}

	if !tea.BoolValue(util.IsUnset(request.AuditProject)) {
		body["AuditProject"] = request.AuditProject
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		body["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.CRAggregationEnabled)) {
		body["CRAggregationEnabled"] = request.CRAggregationEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterSpec)) {
		body["ClusterSpec"] = request.ClusterSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigSourceEnabled)) {
		body["ConfigSourceEnabled"] = request.ConfigSourceEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigSourceNacosID)) {
		body["ConfigSourceNacosID"] = request.ConfigSourceNacosID
	}

	if !tea.BoolValue(util.IsUnset(request.ControlPlaneLogEnabled)) {
		body["ControlPlaneLogEnabled"] = request.ControlPlaneLogEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ControlPlaneLogProject)) {
		body["ControlPlaneLogProject"] = request.ControlPlaneLogProject
	}

	if !tea.BoolValue(util.IsUnset(request.CustomizedPrometheus)) {
		body["CustomizedPrometheus"] = request.CustomizedPrometheus
	}

	if !tea.BoolValue(util.IsUnset(request.CustomizedZipkin)) {
		body["CustomizedZipkin"] = request.CustomizedZipkin
	}

	if !tea.BoolValue(util.IsUnset(request.DNSProxyingEnabled)) {
		body["DNSProxyingEnabled"] = request.DNSProxyingEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DubboFilterEnabled)) {
		body["DubboFilterEnabled"] = request.DubboFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Edition)) {
		body["Edition"] = request.Edition
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAmbient)) {
		body["EnableAmbient"] = request.EnableAmbient
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAudit)) {
		body["EnableAudit"] = request.EnableAudit
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCRHistory)) {
		body["EnableCRHistory"] = request.EnableCRHistory
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSDSServer)) {
		body["EnableSDSServer"] = request.EnableSDSServer
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeIPRanges)) {
		body["ExcludeIPRanges"] = request.ExcludeIPRanges
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeInboundPorts)) {
		body["ExcludeInboundPorts"] = request.ExcludeInboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeOutboundPorts)) {
		body["ExcludeOutboundPorts"] = request.ExcludeOutboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.ExistingCaCert)) {
		body["ExistingCaCert"] = request.ExistingCaCert
	}

	if !tea.BoolValue(util.IsUnset(request.ExistingCaKey)) {
		body["ExistingCaKey"] = request.ExistingCaKey
	}

	if !tea.BoolValue(util.IsUnset(request.ExistingCaType)) {
		body["ExistingCaType"] = request.ExistingCaType
	}

	if !tea.BoolValue(util.IsUnset(request.ExistingRootCaCert)) {
		body["ExistingRootCaCert"] = request.ExistingRootCaCert
	}

	if !tea.BoolValue(util.IsUnset(request.ExistingRootCaKey)) {
		body["ExistingRootCaKey"] = request.ExistingRootCaKey
	}

	if !tea.BoolValue(util.IsUnset(request.FilterGatewayClusterConfig)) {
		body["FilterGatewayClusterConfig"] = request.FilterGatewayClusterConfig
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayAPIEnabled)) {
		body["GatewayAPIEnabled"] = request.GatewayAPIEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.GuestCluster)) {
		body["GuestCluster"] = request.GuestCluster
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeIPRanges)) {
		body["IncludeIPRanges"] = request.IncludeIPRanges
	}

	if !tea.BoolValue(util.IsUnset(request.IstioVersion)) {
		body["IstioVersion"] = request.IstioVersion
	}

	if !tea.BoolValue(util.IsUnset(request.KialiEnabled)) {
		body["KialiEnabled"] = request.KialiEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.LocalityLBConf)) {
		body["LocalityLBConf"] = request.LocalityLBConf
	}

	if !tea.BoolValue(util.IsUnset(request.LocalityLoadBalancing)) {
		body["LocalityLoadBalancing"] = request.LocalityLoadBalancing
	}

	if !tea.BoolValue(util.IsUnset(request.MSEEnabled)) {
		body["MSEEnabled"] = request.MSEEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.MultiBufferEnabled)) {
		body["MultiBufferEnabled"] = request.MultiBufferEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.MultiBufferPollDelay)) {
		body["MultiBufferPollDelay"] = request.MultiBufferPollDelay
	}

	if !tea.BoolValue(util.IsUnset(request.MysqlFilterEnabled)) {
		body["MysqlFilterEnabled"] = request.MysqlFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OPALimitCPU)) {
		body["OPALimitCPU"] = request.OPALimitCPU
	}

	if !tea.BoolValue(util.IsUnset(request.OPALimitMemory)) {
		body["OPALimitMemory"] = request.OPALimitMemory
	}

	if !tea.BoolValue(util.IsUnset(request.OPALogLevel)) {
		body["OPALogLevel"] = request.OPALogLevel
	}

	if !tea.BoolValue(util.IsUnset(request.OPARequestCPU)) {
		body["OPARequestCPU"] = request.OPARequestCPU
	}

	if !tea.BoolValue(util.IsUnset(request.OPARequestMemory)) {
		body["OPARequestMemory"] = request.OPARequestMemory
	}

	if !tea.BoolValue(util.IsUnset(request.OpaEnabled)) {
		body["OpaEnabled"] = request.OpaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OpenAgentPolicy)) {
		body["OpenAgentPolicy"] = request.OpenAgentPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PilotLoadBalancerSpec)) {
		body["PilotLoadBalancerSpec"] = request.PilotLoadBalancerSpec
	}

	if !tea.BoolValue(util.IsUnset(request.PrometheusUrl)) {
		body["PrometheusUrl"] = request.PrometheusUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyLimitCPU)) {
		body["ProxyLimitCPU"] = request.ProxyLimitCPU
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyLimitMemory)) {
		body["ProxyLimitMemory"] = request.ProxyLimitMemory
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyRequestCPU)) {
		body["ProxyRequestCPU"] = request.ProxyRequestCPU
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyRequestMemory)) {
		body["ProxyRequestMemory"] = request.ProxyRequestMemory
	}

	if !tea.BoolValue(util.IsUnset(request.RedisFilterEnabled)) {
		body["RedisFilterEnabled"] = request.RedisFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Telemetry)) {
		body["Telemetry"] = request.Telemetry
	}

	if !tea.BoolValue(util.IsUnset(request.ThriftFilterEnabled)) {
		body["ThriftFilterEnabled"] = request.ThriftFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.TraceSampling)) {
		body["TraceSampling"] = request.TraceSampling
	}

	if !tea.BoolValue(util.IsUnset(request.Tracing)) {
		body["Tracing"] = request.Tracing
	}

	if !tea.BoolValue(util.IsUnset(request.UseExistingCA)) {
		body["UseExistingCA"] = request.UseExistingCA
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitches)) {
		body["VSwitches"] = request.VSwitches
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.WebAssemblyFilterEnabled)) {
		body["WebAssemblyFilterEnabled"] = request.WebAssemblyFilterEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceMesh"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceMeshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceMesh(request *CreateServiceMeshRequest) (_result *CreateServiceMeshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceMeshResponse{}
	_body, _err := client.CreateServiceMeshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSwimLaneWithOptions(request *CreateSwimLaneRequest, runtime *util.RuntimeOptions) (_result *CreateSwimLaneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorKey)) {
		body["LabelSelectorKey"] = request.LabelSelectorKey
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorValue)) {
		body["LabelSelectorValue"] = request.LabelSelectorValue
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicesList)) {
		body["ServicesList"] = request.ServicesList
	}

	if !tea.BoolValue(util.IsUnset(request.SwimLaneName)) {
		body["SwimLaneName"] = request.SwimLaneName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSwimLane"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSwimLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSwimLane(request *CreateSwimLaneRequest) (_result *CreateSwimLaneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSwimLaneResponse{}
	_body, _err := client.CreateSwimLaneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSwimLaneGroupWithOptions(request *CreateSwimLaneGroupRequest, runtime *util.RuntimeOptions) (_result *CreateSwimLaneGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IngressGatewayName)) {
		body["IngressGatewayName"] = request.IngressGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.IngressType)) {
		body["IngressType"] = request.IngressType
	}

	if !tea.BoolValue(util.IsUnset(request.IsPermissive)) {
		body["IsPermissive"] = request.IsPermissive
	}

	if !tea.BoolValue(util.IsUnset(request.RouteHeader)) {
		body["RouteHeader"] = request.RouteHeader
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicesList)) {
		body["ServicesList"] = request.ServicesList
	}

	if !tea.BoolValue(util.IsUnset(request.TraceHeader)) {
		body["TraceHeader"] = request.TraceHeader
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSwimLaneGroup"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSwimLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSwimLaneGroup(request *CreateSwimLaneGroupRequest) (_result *CreateSwimLaneGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSwimLaneGroupResponse{}
	_body, _err := client.CreateSwimLaneGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayRouteWithOptions(request *DeleteGatewayRouteRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.RouteName)) {
		body["RouteName"] = request.RouteName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayRoute"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGatewayRoute(request *DeleteGatewayRouteRequest) (_result *DeleteGatewayRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayRouteResponse{}
	_body, _err := client.DeleteGatewayRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewaySecretWithOptions(request *DeleteGatewaySecretRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewaySecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.SecretName)) {
		body["SecretName"] = request.SecretName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewaySecret"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewaySecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGatewaySecret(request *DeleteGatewaySecretRequest) (_result *DeleteGatewaySecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewaySecretResponse{}
	_body, _err := client.DeleteGatewaySecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIstioGatewayDomainsWithOptions(request *DeleteIstioGatewayDomainsRequest, runtime *util.RuntimeOptions) (_result *DeleteIstioGatewayDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Hosts)) {
		body["Hosts"] = request.Hosts
	}

	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PortName)) {
		body["PortName"] = request.PortName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIstioGatewayDomains"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIstioGatewayDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIstioGatewayDomains(request *DeleteIstioGatewayDomainsRequest) (_result *DeleteIstioGatewayDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIstioGatewayDomainsResponse{}
	_body, _err := client.DeleteIstioGatewayDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceMeshWithOptions(request *DeleteServiceMeshRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		body["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.RetainResources)) {
		body["RetainResources"] = request.RetainResources
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceMesh"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceMeshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceMesh(request *DeleteServiceMeshRequest) (_result *DeleteServiceMeshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceMeshResponse{}
	_body, _err := client.DeleteServiceMeshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSwimLaneWithOptions(request *DeleteSwimLaneRequest, runtime *util.RuntimeOptions) (_result *DeleteSwimLaneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.SwimLaneName)) {
		body["SwimLaneName"] = request.SwimLaneName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSwimLane"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSwimLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSwimLane(request *DeleteSwimLaneRequest) (_result *DeleteSwimLaneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSwimLaneResponse{}
	_body, _err := client.DeleteSwimLaneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSwimLaneGroupWithOptions(request *DeleteSwimLaneGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSwimLaneGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSwimLaneGroup"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSwimLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSwimLaneGroup(request *DeleteSwimLaneGroupRequest) (_result *DeleteSwimLaneGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSwimLaneGroupResponse{}
	_body, _err := client.DeleteSwimLaneGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeASMGatewayImportedServicesWithOptions(request *DescribeASMGatewayImportedServicesRequest, runtime *util.RuntimeOptions) (_result *DescribeASMGatewayImportedServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASMGatewayName)) {
		body["ASMGatewayName"] = request.ASMGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceNamespace)) {
		body["ServiceNamespace"] = request.ServiceNamespace
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeASMGatewayImportedServices"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeASMGatewayImportedServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeASMGatewayImportedServices(request *DescribeASMGatewayImportedServicesRequest) (_result *DescribeASMGatewayImportedServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeASMGatewayImportedServicesResponse{}
	_body, _err := client.DescribeASMGatewayImportedServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCCMVersionWithOptions(request *DescribeCCMVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeCCMVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCCMVersion"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCCMVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCCMVersion(request *DescribeCCMVersionRequest) (_result *DescribeCCMVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCCMVersionResponse{}
	_body, _err := client.DescribeCCMVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCensWithOptions(request *DescribeCensRequest, runtime *util.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCens"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCens(request *DescribeCensRequest) (_result *DescribeCensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCensResponse{}
	_body, _err := client.DescribeCensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterGrafanaWithOptions(request *DescribeClusterGrafanaRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterGrafanaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.K8sClusterId)) {
		query["K8sClusterId"] = request.K8sClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterGrafana"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterGrafanaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterGrafana(request *DescribeClusterGrafanaRequest) (_result *DescribeClusterGrafanaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterGrafanaResponse{}
	_body, _err := client.DescribeClusterGrafanaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterPrometheusWithOptions(request *DescribeClusterPrometheusRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterPrometheusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.K8sClusterId)) {
		query["K8sClusterId"] = request.K8sClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.K8sClusterRegionId)) {
		query["K8sClusterRegionId"] = request.K8sClusterRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterPrometheus"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterPrometheusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterPrometheus(request *DescribeClusterPrometheusRequest) (_result *DescribeClusterPrometheusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterPrometheusResponse{}
	_body, _err := client.DescribeClusterPrometheusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersInServiceMeshWithOptions(request *DescribeClustersInServiceMeshRequest, runtime *util.RuntimeOptions) (_result *DescribeClustersInServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClustersInServiceMesh"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClustersInServiceMeshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClustersInServiceMesh(request *DescribeClustersInServiceMeshRequest) (_result *DescribeClustersInServiceMeshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClustersInServiceMeshResponse{}
	_body, _err := client.DescribeClustersInServiceMeshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCrTemplatesWithOptions(request *DescribeCrTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeCrTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IstioVersion)) {
		body["IstioVersion"] = request.IstioVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Kind)) {
		body["Kind"] = request.Kind
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCrTemplates"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCrTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCrTemplates(request *DescribeCrTemplatesRequest) (_result *DescribeCrTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCrTemplatesResponse{}
	_body, _err := client.DescribeCrTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEipResourcesWithOptions(request *DescribeEipResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeEipResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEipResources"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEipResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEipResources(request *DescribeEipResourcesRequest) (_result *DescribeEipResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEipResourcesResponse{}
	_body, _err := client.DescribeEipResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGatewaySecretDetailsWithOptions(request *DescribeGatewaySecretDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeGatewaySecretDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGatewaySecretDetails"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewaySecretDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGatewaySecretDetails(request *DescribeGatewaySecretDetailsRequest) (_result *DescribeGatewaySecretDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGatewaySecretDetailsResponse{}
	_body, _err := client.DescribeGatewaySecretDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGuestClusterAccessLogDashboardsWithOptions(request *DescribeGuestClusterAccessLogDashboardsRequest, runtime *util.RuntimeOptions) (_result *DescribeGuestClusterAccessLogDashboardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.K8sClusterId)) {
		body["K8sClusterId"] = request.K8sClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGuestClusterAccessLogDashboards"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGuestClusterAccessLogDashboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGuestClusterAccessLogDashboards(request *DescribeGuestClusterAccessLogDashboardsRequest) (_result *DescribeGuestClusterAccessLogDashboardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGuestClusterAccessLogDashboardsResponse{}
	_body, _err := client.DescribeGuestClusterAccessLogDashboardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGuestClusterNamespacesWithOptions(request *DescribeGuestClusterNamespacesRequest, runtime *util.RuntimeOptions) (_result *DescribeGuestClusterNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GuestClusterID)) {
		body["GuestClusterID"] = request.GuestClusterID
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowNsLabels)) {
		body["ShowNsLabels"] = request.ShowNsLabels
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGuestClusterNamespaces"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGuestClusterNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGuestClusterNamespaces(request *DescribeGuestClusterNamespacesRequest) (_result *DescribeGuestClusterNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGuestClusterNamespacesResponse{}
	_body, _err := client.DescribeGuestClusterNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGuestClusterPodsWithOptions(request *DescribeGuestClusterPodsRequest, runtime *util.RuntimeOptions) (_result *DescribeGuestClusterPodsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GuestClusterID)) {
		body["GuestClusterID"] = request.GuestClusterID
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGuestClusterPods"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGuestClusterPodsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGuestClusterPods(request *DescribeGuestClusterPodsRequest) (_result *DescribeGuestClusterPodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGuestClusterPodsResponse{}
	_body, _err := client.DescribeGuestClusterPodsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImportedServicesDetailWithOptions(request *DescribeImportedServicesDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeImportedServicesDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASMGatewayName)) {
		body["ASMGatewayName"] = request.ASMGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceNamespace)) {
		body["ServiceNamespace"] = request.ServiceNamespace
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImportedServicesDetail"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImportedServicesDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImportedServicesDetail(request *DescribeImportedServicesDetailRequest) (_result *DescribeImportedServicesDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImportedServicesDetailResponse{}
	_body, _err := client.DescribeImportedServicesDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIstioGatewayDomainsWithOptions(request *DescribeIstioGatewayDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeIstioGatewayDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIstioGatewayDomains"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIstioGatewayDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIstioGatewayDomains(request *DescribeIstioGatewayDomainsRequest) (_result *DescribeIstioGatewayDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIstioGatewayDomainsResponse{}
	_body, _err := client.DescribeIstioGatewayDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIstioGatewayRouteDetailWithOptions(request *DescribeIstioGatewayRouteDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeIstioGatewayRouteDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.RouteName)) {
		body["RouteName"] = request.RouteName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIstioGatewayRouteDetail"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIstioGatewayRouteDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIstioGatewayRouteDetail(request *DescribeIstioGatewayRouteDetailRequest) (_result *DescribeIstioGatewayRouteDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIstioGatewayRouteDetailResponse{}
	_body, _err := client.DescribeIstioGatewayRouteDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIstioGatewayRoutesWithOptions(request *DescribeIstioGatewayRoutesRequest, runtime *util.RuntimeOptions) (_result *DescribeIstioGatewayRoutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIstioGatewayRoutes"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIstioGatewayRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIstioGatewayRoutes(request *DescribeIstioGatewayRoutesRequest) (_result *DescribeIstioGatewayRoutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIstioGatewayRoutesResponse{}
	_body, _err := client.DescribeIstioGatewayRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetadataWithOptions(runtime *util.RuntimeOptions) (_result *DescribeMetadataResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeMetadata"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetadata() (_result *DescribeMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetadataResponse{}
	_body, _err := client.DescribeMetadataWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespaceScopeSidecarConfigWithOptions(request *DescribeNamespaceScopeSidecarConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeNamespaceScopeSidecarConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespaceScopeSidecarConfig"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespaceScopeSidecarConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespaceScopeSidecarConfig(request *DescribeNamespaceScopeSidecarConfigRequest) (_result *DescribeNamespaceScopeSidecarConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNamespaceScopeSidecarConfigResponse{}
	_body, _err := client.DescribeNamespaceScopeSidecarConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodesInstanceTypeWithOptions(request *DescribeNodesInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeNodesInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNodesInstanceType"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodesInstanceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodesInstanceType(request *DescribeNodesInstanceTypeRequest) (_result *DescribeNodesInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodesInstanceTypeResponse{}
	_body, _err := client.DescribeNodesInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeReusableSlbWithOptions(request *DescribeReusableSlbRequest, runtime *util.RuntimeOptions) (_result *DescribeReusableSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.K8sClusterId)) {
		body["K8sClusterId"] = request.K8sClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["NetworkType"] = request.NetworkType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeReusableSlb"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeReusableSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeReusableSlb(request *DescribeReusableSlbRequest) (_result *DescribeReusableSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeReusableSlbResponse{}
	_body, _err := client.DescribeReusableSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMeshAdditionalStatusWithOptions(request *DescribeServiceMeshAdditionalStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshAdditionalStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckMode)) {
		body["CheckMode"] = request.CheckMode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshAdditionalStatus"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshAdditionalStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMeshAdditionalStatus(request *DescribeServiceMeshAdditionalStatusRequest) (_result *DescribeServiceMeshAdditionalStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshAdditionalStatusResponse{}
	_body, _err := client.DescribeServiceMeshAdditionalStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMeshClustersWithOptions(request *DescribeServiceMeshClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["Offset"] = request.Offset
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshClusters"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMeshClusters(request *DescribeServiceMeshClustersRequest) (_result *DescribeServiceMeshClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshClustersResponse{}
	_body, _err := client.DescribeServiceMeshClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMeshDetailWithOptions(request *DescribeServiceMeshDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshDetail"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMeshDetail(request *DescribeServiceMeshDetailRequest) (_result *DescribeServiceMeshDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshDetailResponse{}
	_body, _err := client.DescribeServiceMeshDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMeshKubeconfigWithOptions(request *DescribeServiceMeshKubeconfigRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshKubeconfig"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshKubeconfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMeshKubeconfig(request *DescribeServiceMeshKubeconfigRequest) (_result *DescribeServiceMeshKubeconfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshKubeconfigResponse{}
	_body, _err := client.DescribeServiceMeshKubeconfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMeshLogsWithOptions(request *DescribeServiceMeshLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshLogs"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMeshLogs(request *DescribeServiceMeshLogsRequest) (_result *DescribeServiceMeshLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshLogsResponse{}
	_body, _err := client.DescribeServiceMeshLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMeshProxyStatusWithOptions(request *DescribeServiceMeshProxyStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshProxyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshProxyStatus"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshProxyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMeshProxyStatus(request *DescribeServiceMeshProxyStatusRequest) (_result *DescribeServiceMeshProxyStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshProxyStatusResponse{}
	_body, _err := client.DescribeServiceMeshProxyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMeshUpgradeStatusWithOptions(request *DescribeServiceMeshUpgradeStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllIstioGatewayFullNames)) {
		body["AllIstioGatewayFullNames"] = request.AllIstioGatewayFullNames
	}

	if !tea.BoolValue(util.IsUnset(request.GuestClusterIds)) {
		body["GuestClusterIds"] = request.GuestClusterIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshUpgradeStatus"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshUpgradeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMeshUpgradeStatus(request *DescribeServiceMeshUpgradeStatusRequest) (_result *DescribeServiceMeshUpgradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshUpgradeStatusResponse{}
	_body, _err := client.DescribeServiceMeshUpgradeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DescribeServiceMeshVMsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeServiceMeshVMsResponse
 */
// Deprecated
func (client *Client) DescribeServiceMeshVMsWithOptions(request *DescribeServiceMeshVMsRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshVMsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshVMs"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshVMsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DescribeServiceMeshVMsRequest
 * @return DescribeServiceMeshVMsResponse
 */
// Deprecated
func (client *Client) DescribeServiceMeshVMs(request *DescribeServiceMeshVMsRequest) (_result *DescribeServiceMeshVMsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshVMsResponse{}
	_body, _err := client.DescribeServiceMeshVMsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMeshesWithOptions(request *DescribeServiceMeshesRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceMeshesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMeshes"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMeshesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMeshes(request *DescribeServiceMeshesRequest) (_result *DescribeServiceMeshesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshesResponse{}
	_body, _err := client.DescribeServiceMeshesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUpgradeVersionWithOptions(request *DescribeUpgradeVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeUpgradeVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUpgradeVersion"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUpgradeVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUpgradeVersion(request *DescribeUpgradeVersionRequest) (_result *DescribeUpgradeVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUpgradeVersionResponse{}
	_body, _err := client.DescribeUpgradeVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserPermissionsWithOptions(request *DescribeUserPermissionsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubAccountUserId)) {
		body["SubAccountUserId"] = request.SubAccountUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserPermissions"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserPermissions(request *DescribeUserPermissionsRequest) (_result *DescribeUserPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserPermissionsResponse{}
	_body, _err := client.DescribeUserPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsersWithPermissionsWithOptions(request *DescribeUsersWithPermissionsRequest, runtime *util.RuntimeOptions) (_result *DescribeUsersWithPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsersWithPermissions"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsersWithPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsersWithPermissions(request *DescribeUsersWithPermissionsRequest) (_result *DescribeUsersWithPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsersWithPermissionsResponse{}
	_body, _err := client.DescribeUsersWithPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DescribeVMsInServiceMeshRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVMsInServiceMeshResponse
 */
// Deprecated
func (client *Client) DescribeVMsInServiceMeshWithOptions(request *DescribeVMsInServiceMeshRequest, runtime *util.RuntimeOptions) (_result *DescribeVMsInServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVMsInServiceMesh"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVMsInServiceMeshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request DescribeVMsInServiceMeshRequest
 * @return DescribeVMsInServiceMeshResponse
 */
// Deprecated
func (client *Client) DescribeVMsInServiceMesh(request *DescribeVMsInServiceMeshRequest) (_result *DescribeVMsInServiceMeshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVMsInServiceMeshResponse{}
	_body, _err := client.DescribeVMsInServiceMeshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVSwitchesWithOptions(request *DescribeVSwitchesRequest, runtime *util.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVSwitches"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (_result *DescribeVSwitchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DescribeVSwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVersionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeVersionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeVersions"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVersions() (_result *DescribeVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVersionsResponse{}
	_body, _err := client.DescribeVersionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcsWithOptions(request *DescribeVpcsRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcs"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcs(request *DescribeVpcsRequest) (_result *DescribeVpcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcsResponse{}
	_body, _err := client.DescribeVpcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCaCertWithOptions(request *GetCaCertRequest, runtime *util.RuntimeOptions) (_result *GetCaCertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCaCert"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCaCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCaCert(request *GetCaCertRequest) (_result *GetCaCertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCaCertResponse{}
	_body, _err := client.GetCaCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeploymentBySelectorWithOptions(tmpReq *GetDeploymentBySelectorRequest, runtime *util.RuntimeOptions) (_result *GetDeploymentBySelectorResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDeploymentBySelectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LabelSelector)) {
		request.LabelSelectorShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSelector, tea.String("LabelSelector"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GuestCluster)) {
		body["GuestCluster"] = request.GuestCluster
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorShrink)) {
		body["LabelSelector"] = request.LabelSelectorShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Mark)) {
		body["Mark"] = request.Mark
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		body["NameSpace"] = request.NameSpace
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeploymentBySelector"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeploymentBySelectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeploymentBySelector(request *GetDeploymentBySelectorRequest) (_result *GetDeploymentBySelectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeploymentBySelectorResponse{}
	_body, _err := client.GetDeploymentBySelectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGrafanaDashboardUrlWithOptions(request *GetGrafanaDashboardUrlRequest, runtime *util.RuntimeOptions) (_result *GetGrafanaDashboardUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.K8sClusterId)) {
		body["K8sClusterId"] = request.K8sClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGrafanaDashboardUrl"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGrafanaDashboardUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGrafanaDashboardUrl(request *GetGrafanaDashboardUrlRequest) (_result *GetGrafanaDashboardUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGrafanaDashboardUrlResponse{}
	_body, _err := client.GetGrafanaDashboardUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRegisteredServiceEndpointsWithOptions(request *GetRegisteredServiceEndpointsRequest, runtime *util.RuntimeOptions) (_result *GetRegisteredServiceEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterIds)) {
		body["ClusterIds"] = request.ClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		body["ServiceType"] = request.ServiceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegisteredServiceEndpoints"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegisteredServiceEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegisteredServiceEndpoints(request *GetRegisteredServiceEndpointsRequest) (_result *GetRegisteredServiceEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRegisteredServiceEndpointsResponse{}
	_body, _err := client.GetRegisteredServiceEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRegisteredServiceNamespacesWithOptions(request *GetRegisteredServiceNamespacesRequest, runtime *util.RuntimeOptions) (_result *GetRegisteredServiceNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegisteredServiceNamespaces"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegisteredServiceNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegisteredServiceNamespaces(request *GetRegisteredServiceNamespacesRequest) (_result *GetRegisteredServiceNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRegisteredServiceNamespacesResponse{}
	_body, _err := client.GetRegisteredServiceNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSwimLaneDetailWithOptions(request *GetSwimLaneDetailRequest, runtime *util.RuntimeOptions) (_result *GetSwimLaneDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.SwimLaneName)) {
		body["SwimLaneName"] = request.SwimLaneName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSwimLaneDetail"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSwimLaneDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSwimLaneDetail(request *GetSwimLaneDetailRequest) (_result *GetSwimLaneDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSwimLaneDetailResponse{}
	_body, _err := client.GetSwimLaneDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSwimLaneGroupListWithOptions(request *GetSwimLaneGroupListRequest, runtime *util.RuntimeOptions) (_result *GetSwimLaneGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSwimLaneGroupList"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSwimLaneGroupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSwimLaneGroupList(request *GetSwimLaneGroupListRequest) (_result *GetSwimLaneGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSwimLaneGroupListResponse{}
	_body, _err := client.GetSwimLaneGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSwimLaneListWithOptions(request *GetSwimLaneListRequest, runtime *util.RuntimeOptions) (_result *GetSwimLaneListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSwimLaneList"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSwimLaneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSwimLaneList(request *GetSwimLaneListRequest) (_result *GetSwimLaneListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSwimLaneListResponse{}
	_body, _err := client.GetSwimLaneListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetVmAppMeshInfoRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetVmAppMeshInfoResponse
 */
// Deprecated
func (client *Client) GetVmAppMeshInfoWithOptions(request *GetVmAppMeshInfoRequest, runtime *util.RuntimeOptions) (_result *GetVmAppMeshInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVmAppMeshInfo"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVmAppMeshInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetVmAppMeshInfoRequest
 * @return GetVmAppMeshInfoResponse
 */
// Deprecated
func (client *Client) GetVmAppMeshInfo(request *GetVmAppMeshInfoRequest) (_result *GetVmAppMeshInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVmAppMeshInfoResponse{}
	_body, _err := client.GetVmAppMeshInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetVmMetaRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetVmMetaResponse
 */
// Deprecated
func (client *Client) GetVmMetaWithOptions(request *GetVmMetaRequest, runtime *util.RuntimeOptions) (_result *GetVmMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVmMeta"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVmMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetVmMetaRequest
 * @return GetVmMetaResponse
 */
// Deprecated
func (client *Client) GetVmMeta(request *GetVmMetaRequest) (_result *GetVmMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVmMetaResponse{}
	_body, _err := client.GetVmMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantUserPermissionsWithOptions(tmpReq *GrantUserPermissionsRequest, runtime *util.RuntimeOptions) (_result *GrantUserPermissionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GrantUserPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SubAccountUserIds)) {
		request.SubAccountUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubAccountUserIds, tea.String("SubAccountUserIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Permissions)) {
		body["Permissions"] = request.Permissions
	}

	if !tea.BoolValue(util.IsUnset(request.SubAccountUserId)) {
		body["SubAccountUserId"] = request.SubAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SubAccountUserIdsShrink)) {
		body["SubAccountUserIds"] = request.SubAccountUserIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantUserPermissions"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantUserPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantUserPermissions(request *GrantUserPermissionsRequest) (_result *GrantUserPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantUserPermissionsResponse{}
	_body, _err := client.GrantUserPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyApiServerEipResourceWithOptions(request *ModifyApiServerEipResourceRequest, runtime *util.RuntimeOptions) (_result *ModifyApiServerEipResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiServerEipId)) {
		body["ApiServerEipId"] = request.ApiServerEipId
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApiServerEipResource"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApiServerEipResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApiServerEipResource(request *ModifyApiServerEipResourceRequest) (_result *ModifyApiServerEipResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApiServerEipResourceResponse{}
	_body, _err := client.ModifyApiServerEipResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyServiceMeshNameWithOptions(request *ModifyServiceMeshNameRequest, runtime *util.RuntimeOptions) (_result *ModifyServiceMeshNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyServiceMeshName"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyServiceMeshNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyServiceMeshName(request *ModifyServiceMeshNameRequest) (_result *ModifyServiceMeshNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyServiceMeshNameResponse{}
	_body, _err := client.ModifyServiceMeshNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing methods of Log Service. For more information, visit the [pricing page](https://www.aliyun.com/price/product?spm=5176.10695662.1119587.4.194c6a67rcPWQH#/sls/detail).
 *
 * @param request ReActivateAuditRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReActivateAuditResponse
 */
func (client *Client) ReActivateAuditWithOptions(request *ReActivateAuditRequest, runtime *util.RuntimeOptions) (_result *ReActivateAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableAudit)) {
		body["EnableAudit"] = request.EnableAudit
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReActivateAudit"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReActivateAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing methods of Log Service. For more information, visit the [pricing page](https://www.aliyun.com/price/product?spm=5176.10695662.1119587.4.194c6a67rcPWQH#/sls/detail).
 *
 * @param request ReActivateAuditRequest
 * @return ReActivateAuditResponse
 */
func (client *Client) ReActivateAudit(request *ReActivateAuditRequest) (_result *ReActivateAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReActivateAuditResponse{}
	_body, _err := client.ReActivateAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveClusterFromServiceMeshWithOptions(request *RemoveClusterFromServiceMeshRequest, runtime *util.RuntimeOptions) (_result *RemoveClusterFromServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ReserveNamespace)) {
		body["ReserveNamespace"] = request.ReserveNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveClusterFromServiceMesh"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveClusterFromServiceMeshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveClusterFromServiceMesh(request *RemoveClusterFromServiceMeshRequest) (_result *RemoveClusterFromServiceMeshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveClusterFromServiceMeshResponse{}
	_body, _err := client.RemoveClusterFromServiceMeshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request RemoveVMFromServiceMeshRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveVMFromServiceMeshResponse
 */
// Deprecated
func (client *Client) RemoveVMFromServiceMeshWithOptions(request *RemoveVMFromServiceMeshRequest, runtime *util.RuntimeOptions) (_result *RemoveVMFromServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcsId)) {
		query["EcsId"] = request.EcsId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveVMFromServiceMesh"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveVMFromServiceMeshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request RemoveVMFromServiceMeshRequest
 * @return RemoveVMFromServiceMeshResponse
 */
// Deprecated
func (client *Client) RemoveVMFromServiceMesh(request *RemoveVMFromServiceMeshRequest) (_result *RemoveVMFromServiceMeshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveVMFromServiceMeshResponse{}
	_body, _err := client.RemoveVMFromServiceMeshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeKubeconfigWithOptions(request *RevokeKubeconfigRequest, runtime *util.RuntimeOptions) (_result *RevokeKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeKubeconfig"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeKubeconfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeKubeconfig(request *RevokeKubeconfigRequest) (_result *RevokeKubeconfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeKubeconfigResponse{}
	_body, _err := client.RevokeKubeconfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateASMGatewayWithOptions(request *UpdateASMGatewayRequest, runtime *util.RuntimeOptions) (_result *UpdateASMGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateASMGateway"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateASMGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateASMGateway(request *UpdateASMGatewayRequest) (_result *UpdateASMGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateASMGatewayResponse{}
	_body, _err := client.UpdateASMGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateASMGatewayImportedServicesWithOptions(request *UpdateASMGatewayImportedServicesRequest, runtime *util.RuntimeOptions) (_result *UpdateASMGatewayImportedServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASMGatewayName)) {
		body["ASMGatewayName"] = request.ASMGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceNames)) {
		body["ServiceNames"] = request.ServiceNames
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceNamespace)) {
		body["ServiceNamespace"] = request.ServiceNamespace
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateASMGatewayImportedServices"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateASMGatewayImportedServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateASMGatewayImportedServices(request *UpdateASMGatewayImportedServicesRequest) (_result *UpdateASMGatewayImportedServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateASMGatewayImportedServicesResponse{}
	_body, _err := client.UpdateASMGatewayImportedServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateASMNamespaceFromGuestClusterWithOptions(request *UpdateASMNamespaceFromGuestClusterRequest, runtime *util.RuntimeOptions) (_result *UpdateASMNamespaceFromGuestClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.K8sClusterId)) {
		body["K8sClusterId"] = request.K8sClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateASMNamespaceFromGuestCluster"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateASMNamespaceFromGuestClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateASMNamespaceFromGuestCluster(request *UpdateASMNamespaceFromGuestClusterRequest) (_result *UpdateASMNamespaceFromGuestClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateASMNamespaceFromGuestClusterResponse{}
	_body, _err := client.UpdateASMNamespaceFromGuestClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateControlPlaneLogConfigWithOptions(request *UpdateControlPlaneLogConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateControlPlaneLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		body["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.LogTTLInDay)) {
		body["LogTTLInDay"] = request.LogTTLInDay
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		body["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateControlPlaneLogConfig"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateControlPlaneLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateControlPlaneLogConfig(request *UpdateControlPlaneLogConfigRequest) (_result *UpdateControlPlaneLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateControlPlaneLogConfigResponse{}
	_body, _err := client.UpdateControlPlaneLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIstioGatewayRoutesWithOptions(tmpReq *UpdateIstioGatewayRoutesRequest, runtime *util.RuntimeOptions) (_result *UpdateIstioGatewayRoutesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateIstioGatewayRoutesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GatewayRoute)) {
		request.GatewayRouteShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GatewayRoute, tea.String("GatewayRoute"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayRouteShrink)) {
		body["GatewayRoute"] = request.GatewayRouteShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIstioGatewayRoutes"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIstioGatewayRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIstioGatewayRoutes(request *UpdateIstioGatewayRoutesRequest) (_result *UpdateIstioGatewayRoutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIstioGatewayRoutesResponse{}
	_body, _err := client.UpdateIstioGatewayRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIstioInjectionConfigWithOptions(request *UpdateIstioInjectionConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateIstioInjectionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataPlaneMode)) {
		body["DataPlaneMode"] = request.DataPlaneMode
	}

	if !tea.BoolValue(util.IsUnset(request.EnableIstioInjection)) {
		body["EnableIstioInjection"] = request.EnableIstioInjection
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSidecarSetInjection)) {
		body["EnableSidecarSetInjection"] = request.EnableSidecarSetInjection
	}

	if !tea.BoolValue(util.IsUnset(request.IstioRev)) {
		body["IstioRev"] = request.IstioRev
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIstioInjectionConfig"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIstioInjectionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIstioInjectionConfig(request *UpdateIstioInjectionConfigRequest) (_result *UpdateIstioInjectionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIstioInjectionConfigResponse{}
	_body, _err := client.UpdateIstioInjectionConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIstioRouteAdditionalStatusWithOptions(request *UpdateIstioRouteAdditionalStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateIstioRouteAdditionalStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RouteName)) {
		query["RouteName"] = request.RouteName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IstioGatewayName)) {
		body["IstioGatewayName"] = request.IstioGatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIstioRouteAdditionalStatus"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIstioRouteAdditionalStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIstioRouteAdditionalStatus(request *UpdateIstioRouteAdditionalStatusRequest) (_result *UpdateIstioRouteAdditionalStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIstioRouteAdditionalStatusResponse{}
	_body, _err := client.UpdateIstioRouteAdditionalStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMeshCRAggregationWithOptions(request *UpdateMeshCRAggregationRequest, runtime *util.RuntimeOptions) (_result *UpdateMeshCRAggregationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CPULimit)) {
		body["CPULimit"] = request.CPULimit
	}

	if !tea.BoolValue(util.IsUnset(request.CPURequirement)) {
		body["CPURequirement"] = request.CPURequirement
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		body["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.MemoryLimit)) {
		body["MemoryLimit"] = request.MemoryLimit
	}

	if !tea.BoolValue(util.IsUnset(request.MemoryRequirement)) {
		body["MemoryRequirement"] = request.MemoryRequirement
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.UsePublicApiServer)) {
		body["UsePublicApiServer"] = request.UsePublicApiServer
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMeshCRAggregation"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMeshCRAggregationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMeshCRAggregation(request *UpdateMeshCRAggregationRequest) (_result *UpdateMeshCRAggregationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMeshCRAggregationResponse{}
	_body, _err := client.UpdateMeshCRAggregationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMeshFeatureWithOptions(request *UpdateMeshFeatureRequest, runtime *util.RuntimeOptions) (_result *UpdateMeshFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLogGatewayEnabled)) {
		query["AccessLogGatewayEnabled"] = request.AccessLogGatewayEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogSidecarEnabled)) {
		query["AccessLogSidecarEnabled"] = request.AccessLogSidecarEnabled
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLogEnabled)) {
		body["AccessLogEnabled"] = request.AccessLogEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogFile)) {
		body["AccessLogFile"] = request.AccessLogFile
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogFormat)) {
		body["AccessLogFormat"] = request.AccessLogFormat
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogGatewayLifecycle)) {
		body["AccessLogGatewayLifecycle"] = request.AccessLogGatewayLifecycle
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogProject)) {
		body["AccessLogProject"] = request.AccessLogProject
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogServiceEnabled)) {
		body["AccessLogServiceEnabled"] = request.AccessLogServiceEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogServiceHost)) {
		body["AccessLogServiceHost"] = request.AccessLogServiceHost
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogServicePort)) {
		body["AccessLogServicePort"] = request.AccessLogServicePort
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogSidecarLifecycle)) {
		body["AccessLogSidecarLifecycle"] = request.AccessLogSidecarLifecycle
	}

	if !tea.BoolValue(util.IsUnset(request.AuditProject)) {
		body["AuditProject"] = request.AuditProject
	}

	if !tea.BoolValue(util.IsUnset(request.AutoInjectionPolicyEnabled)) {
		body["AutoInjectionPolicyEnabled"] = request.AutoInjectionPolicyEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.CRAggregationEnabled)) {
		body["CRAggregationEnabled"] = request.CRAggregationEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterSpec)) {
		body["ClusterSpec"] = request.ClusterSpec
	}

	if !tea.BoolValue(util.IsUnset(request.CniEnabled)) {
		body["CniEnabled"] = request.CniEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.CniExcludeNamespaces)) {
		body["CniExcludeNamespaces"] = request.CniExcludeNamespaces
	}

	if !tea.BoolValue(util.IsUnset(request.Concurrency)) {
		body["Concurrency"] = request.Concurrency
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigSourceEnabled)) {
		body["ConfigSourceEnabled"] = request.ConfigSourceEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigSourceNacosID)) {
		body["ConfigSourceNacosID"] = request.ConfigSourceNacosID
	}

	if !tea.BoolValue(util.IsUnset(request.CustomizedPrometheus)) {
		body["CustomizedPrometheus"] = request.CustomizedPrometheus
	}

	if !tea.BoolValue(util.IsUnset(request.CustomizedZipkin)) {
		body["CustomizedZipkin"] = request.CustomizedZipkin
	}

	if !tea.BoolValue(util.IsUnset(request.DNSProxyingEnabled)) {
		body["DNSProxyingEnabled"] = request.DNSProxyingEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultComponentsScheduleConfig)) {
		body["DefaultComponentsScheduleConfig"] = request.DefaultComponentsScheduleConfig
	}

	if !tea.BoolValue(util.IsUnset(request.DiscoverySelectors)) {
		body["DiscoverySelectors"] = request.DiscoverySelectors
	}

	if !tea.BoolValue(util.IsUnset(request.DubboFilterEnabled)) {
		body["DubboFilterEnabled"] = request.DubboFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAudit)) {
		body["EnableAudit"] = request.EnableAudit
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAutoDiagnosis)) {
		body["EnableAutoDiagnosis"] = request.EnableAutoDiagnosis
	}

	if !tea.BoolValue(util.IsUnset(request.EnableBootstrapXdsAgent)) {
		body["EnableBootstrapXdsAgent"] = request.EnableBootstrapXdsAgent
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCRHistory)) {
		body["EnableCRHistory"] = request.EnableCRHistory
	}

	if !tea.BoolValue(util.IsUnset(request.EnableNamespacesByDefault)) {
		body["EnableNamespacesByDefault"] = request.EnableNamespacesByDefault
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSDSServer)) {
		body["EnableSDSServer"] = request.EnableSDSServer
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeIPRanges)) {
		body["ExcludeIPRanges"] = request.ExcludeIPRanges
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeInboundPorts)) {
		body["ExcludeInboundPorts"] = request.ExcludeInboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeOutboundPorts)) {
		body["ExcludeOutboundPorts"] = request.ExcludeOutboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.FilterGatewayClusterConfig)) {
		body["FilterGatewayClusterConfig"] = request.FilterGatewayClusterConfig
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayAPIEnabled)) {
		body["GatewayAPIEnabled"] = request.GatewayAPIEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.HoldApplicationUntilProxyStarts)) {
		body["HoldApplicationUntilProxyStarts"] = request.HoldApplicationUntilProxyStarts
	}

	if !tea.BoolValue(util.IsUnset(request.Http10Enabled)) {
		body["Http10Enabled"] = request.Http10Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeIPRanges)) {
		body["IncludeIPRanges"] = request.IncludeIPRanges
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeInboundPorts)) {
		body["IncludeInboundPorts"] = request.IncludeInboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeOutboundPorts)) {
		body["IncludeOutboundPorts"] = request.IncludeOutboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrateKiali)) {
		body["IntegrateKiali"] = request.IntegrateKiali
	}

	if !tea.BoolValue(util.IsUnset(request.InterceptionMode)) {
		body["InterceptionMode"] = request.InterceptionMode
	}

	if !tea.BoolValue(util.IsUnset(request.KialiArmsAuthTokens)) {
		body["KialiArmsAuthTokens"] = request.KialiArmsAuthTokens
	}

	if !tea.BoolValue(util.IsUnset(request.KialiEnabled)) {
		body["KialiEnabled"] = request.KialiEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.KialiServiceAnnotations)) {
		body["KialiServiceAnnotations"] = request.KialiServiceAnnotations
	}

	if !tea.BoolValue(util.IsUnset(request.Lifecycle)) {
		body["Lifecycle"] = request.Lifecycle
	}

	if !tea.BoolValue(util.IsUnset(request.LocalityLBConf)) {
		body["LocalityLBConf"] = request.LocalityLBConf
	}

	if !tea.BoolValue(util.IsUnset(request.LocalityLoadBalancing)) {
		body["LocalityLoadBalancing"] = request.LocalityLoadBalancing
	}

	if !tea.BoolValue(util.IsUnset(request.LogLevel)) {
		body["LogLevel"] = request.LogLevel
	}

	if !tea.BoolValue(util.IsUnset(request.MSEEnabled)) {
		body["MSEEnabled"] = request.MSEEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.MultiBufferEnabled)) {
		body["MultiBufferEnabled"] = request.MultiBufferEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.MultiBufferPollDelay)) {
		body["MultiBufferPollDelay"] = request.MultiBufferPollDelay
	}

	if !tea.BoolValue(util.IsUnset(request.MysqlFilterEnabled)) {
		body["MysqlFilterEnabled"] = request.MysqlFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.NFDEnabled)) {
		body["NFDEnabled"] = request.NFDEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.NFDLabelPruned)) {
		body["NFDLabelPruned"] = request.NFDLabelPruned
	}

	if !tea.BoolValue(util.IsUnset(request.OPAInjectorCPULimit)) {
		body["OPAInjectorCPULimit"] = request.OPAInjectorCPULimit
	}

	if !tea.BoolValue(util.IsUnset(request.OPAInjectorCPURequirement)) {
		body["OPAInjectorCPURequirement"] = request.OPAInjectorCPURequirement
	}

	if !tea.BoolValue(util.IsUnset(request.OPAInjectorMemoryLimit)) {
		body["OPAInjectorMemoryLimit"] = request.OPAInjectorMemoryLimit
	}

	if !tea.BoolValue(util.IsUnset(request.OPAInjectorMemoryRequirement)) {
		body["OPAInjectorMemoryRequirement"] = request.OPAInjectorMemoryRequirement
	}

	if !tea.BoolValue(util.IsUnset(request.OPALimitCPU)) {
		body["OPALimitCPU"] = request.OPALimitCPU
	}

	if !tea.BoolValue(util.IsUnset(request.OPALimitMemory)) {
		body["OPALimitMemory"] = request.OPALimitMemory
	}

	if !tea.BoolValue(util.IsUnset(request.OPALogLevel)) {
		body["OPALogLevel"] = request.OPALogLevel
	}

	if !tea.BoolValue(util.IsUnset(request.OPARequestCPU)) {
		body["OPARequestCPU"] = request.OPARequestCPU
	}

	if !tea.BoolValue(util.IsUnset(request.OPARequestMemory)) {
		body["OPARequestMemory"] = request.OPARequestMemory
	}

	if !tea.BoolValue(util.IsUnset(request.OPAScopeInjected)) {
		body["OPAScopeInjected"] = request.OPAScopeInjected
	}

	if !tea.BoolValue(util.IsUnset(request.OpaEnabled)) {
		body["OpaEnabled"] = request.OpaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OpenAgentPolicy)) {
		body["OpenAgentPolicy"] = request.OpenAgentPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.OutboundTrafficPolicy)) {
		body["OutboundTrafficPolicy"] = request.OutboundTrafficPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.PrometheusUrl)) {
		body["PrometheusUrl"] = request.PrometheusUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitCPUResourceLimit)) {
		body["ProxyInitCPUResourceLimit"] = request.ProxyInitCPUResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitCPUResourceRequest)) {
		body["ProxyInitCPUResourceRequest"] = request.ProxyInitCPUResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitMemoryResourceLimit)) {
		body["ProxyInitMemoryResourceLimit"] = request.ProxyInitMemoryResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitMemoryResourceRequest)) {
		body["ProxyInitMemoryResourceRequest"] = request.ProxyInitMemoryResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyLimitCPU)) {
		body["ProxyLimitCPU"] = request.ProxyLimitCPU
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyLimitMemory)) {
		body["ProxyLimitMemory"] = request.ProxyLimitMemory
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyRequestCPU)) {
		body["ProxyRequestCPU"] = request.ProxyRequestCPU
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyRequestMemory)) {
		body["ProxyRequestMemory"] = request.ProxyRequestMemory
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyStatsMatcher)) {
		body["ProxyStatsMatcher"] = request.ProxyStatsMatcher
	}

	if !tea.BoolValue(util.IsUnset(request.RedisFilterEnabled)) {
		body["RedisFilterEnabled"] = request.RedisFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarInjectorLimitCPU)) {
		body["SidecarInjectorLimitCPU"] = request.SidecarInjectorLimitCPU
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarInjectorLimitMemory)) {
		body["SidecarInjectorLimitMemory"] = request.SidecarInjectorLimitMemory
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarInjectorRequestCPU)) {
		body["SidecarInjectorRequestCPU"] = request.SidecarInjectorRequestCPU
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarInjectorRequestMemory)) {
		body["SidecarInjectorRequestMemory"] = request.SidecarInjectorRequestMemory
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarInjectorWebhookAsYaml)) {
		body["SidecarInjectorWebhookAsYaml"] = request.SidecarInjectorWebhookAsYaml
	}

	if !tea.BoolValue(util.IsUnset(request.Telemetry)) {
		body["Telemetry"] = request.Telemetry
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationDrainDuration)) {
		body["TerminationDrainDuration"] = request.TerminationDrainDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ThriftFilterEnabled)) {
		body["ThriftFilterEnabled"] = request.ThriftFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.TraceCustomTags)) {
		body["TraceCustomTags"] = request.TraceCustomTags
	}

	if !tea.BoolValue(util.IsUnset(request.TraceMaxPathTagLength)) {
		body["TraceMaxPathTagLength"] = request.TraceMaxPathTagLength
	}

	if !tea.BoolValue(util.IsUnset(request.TraceSampling)) {
		body["TraceSampling"] = request.TraceSampling
	}

	if !tea.BoolValue(util.IsUnset(request.Tracing)) {
		body["Tracing"] = request.Tracing
	}

	if !tea.BoolValue(util.IsUnset(request.TracingOnExtZipkinLimitCPU)) {
		body["TracingOnExtZipkinLimitCPU"] = request.TracingOnExtZipkinLimitCPU
	}

	if !tea.BoolValue(util.IsUnset(request.TracingOnExtZipkinLimitMemory)) {
		body["TracingOnExtZipkinLimitMemory"] = request.TracingOnExtZipkinLimitMemory
	}

	if !tea.BoolValue(util.IsUnset(request.TracingOnExtZipkinRequestCPU)) {
		body["TracingOnExtZipkinRequestCPU"] = request.TracingOnExtZipkinRequestCPU
	}

	if !tea.BoolValue(util.IsUnset(request.TracingOnExtZipkinRequestMemory)) {
		body["TracingOnExtZipkinRequestMemory"] = request.TracingOnExtZipkinRequestMemory
	}

	if !tea.BoolValue(util.IsUnset(request.WebAssemblyFilterEnabled)) {
		body["WebAssemblyFilterEnabled"] = request.WebAssemblyFilterEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMeshFeature"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMeshFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMeshFeature(request *UpdateMeshFeatureRequest) (_result *UpdateMeshFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMeshFeatureResponse{}
	_body, _err := client.UpdateMeshFeatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNamespaceScopeSidecarConfigWithOptions(request *UpdateNamespaceScopeSidecarConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateNamespaceScopeSidecarConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Concurrency)) {
		body["Concurrency"] = request.Concurrency
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCoreDump)) {
		body["EnableCoreDump"] = request.EnableCoreDump
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeIPRanges)) {
		body["ExcludeIPRanges"] = request.ExcludeIPRanges
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeInboundPorts)) {
		body["ExcludeInboundPorts"] = request.ExcludeInboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeOutboundPorts)) {
		body["ExcludeOutboundPorts"] = request.ExcludeOutboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.HoldApplicationUntilProxyStarts)) {
		body["HoldApplicationUntilProxyStarts"] = request.HoldApplicationUntilProxyStarts
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeIPRanges)) {
		body["IncludeIPRanges"] = request.IncludeIPRanges
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeInboundPorts)) {
		body["IncludeInboundPorts"] = request.IncludeInboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeOutboundPorts)) {
		body["IncludeOutboundPorts"] = request.IncludeOutboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.InterceptionMode)) {
		body["InterceptionMode"] = request.InterceptionMode
	}

	if !tea.BoolValue(util.IsUnset(request.IstioDNSProxyEnabled)) {
		body["IstioDNSProxyEnabled"] = request.IstioDNSProxyEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Lifecycle)) {
		body["Lifecycle"] = request.Lifecycle
	}

	if !tea.BoolValue(util.IsUnset(request.LogLevel)) {
		body["LogLevel"] = request.LogLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PostStart)) {
		body["PostStart"] = request.PostStart
	}

	if !tea.BoolValue(util.IsUnset(request.PreStop)) {
		body["PreStop"] = request.PreStop
	}

	if !tea.BoolValue(util.IsUnset(request.Privileged)) {
		body["Privileged"] = request.Privileged
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitAckSloCPUResourceLimit)) {
		body["ProxyInitAckSloCPUResourceLimit"] = request.ProxyInitAckSloCPUResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitAckSloCPUResourceRequest)) {
		body["ProxyInitAckSloCPUResourceRequest"] = request.ProxyInitAckSloCPUResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitAckSloMemoryResourceLimit)) {
		body["ProxyInitAckSloMemoryResourceLimit"] = request.ProxyInitAckSloMemoryResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitAckSloMemoryResourceRequest)) {
		body["ProxyInitAckSloMemoryResourceRequest"] = request.ProxyInitAckSloMemoryResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitCPUResourceLimit)) {
		body["ProxyInitCPUResourceLimit"] = request.ProxyInitCPUResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitCPUResourceRequest)) {
		body["ProxyInitCPUResourceRequest"] = request.ProxyInitCPUResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitMemoryResourceLimit)) {
		body["ProxyInitMemoryResourceLimit"] = request.ProxyInitMemoryResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyInitMemoryResourceRequest)) {
		body["ProxyInitMemoryResourceRequest"] = request.ProxyInitMemoryResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyMetadata)) {
		body["ProxyMetadata"] = request.ProxyMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyStatsMatcher)) {
		body["ProxyStatsMatcher"] = request.ProxyStatsMatcher
	}

	if !tea.BoolValue(util.IsUnset(request.ReadinessFailureThreshold)) {
		body["ReadinessFailureThreshold"] = request.ReadinessFailureThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ReadinessInitialDelaySeconds)) {
		body["ReadinessInitialDelaySeconds"] = request.ReadinessInitialDelaySeconds
	}

	if !tea.BoolValue(util.IsUnset(request.ReadinessPeriodSeconds)) {
		body["ReadinessPeriodSeconds"] = request.ReadinessPeriodSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarProxyAckSloCPUResourceLimit)) {
		body["SidecarProxyAckSloCPUResourceLimit"] = request.SidecarProxyAckSloCPUResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarProxyAckSloCPUResourceRequest)) {
		body["SidecarProxyAckSloCPUResourceRequest"] = request.SidecarProxyAckSloCPUResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarProxyAckSloMemoryResourceLimit)) {
		body["SidecarProxyAckSloMemoryResourceLimit"] = request.SidecarProxyAckSloMemoryResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarProxyAckSloMemoryResourceRequest)) {
		body["SidecarProxyAckSloMemoryResourceRequest"] = request.SidecarProxyAckSloMemoryResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarProxyCPUResourceLimit)) {
		body["SidecarProxyCPUResourceLimit"] = request.SidecarProxyCPUResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarProxyCPUResourceRequest)) {
		body["SidecarProxyCPUResourceRequest"] = request.SidecarProxyCPUResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarProxyMemoryResourceLimit)) {
		body["SidecarProxyMemoryResourceLimit"] = request.SidecarProxyMemoryResourceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SidecarProxyMemoryResourceRequest)) {
		body["SidecarProxyMemoryResourceRequest"] = request.SidecarProxyMemoryResourceRequest
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationDrainDuration)) {
		body["TerminationDrainDuration"] = request.TerminationDrainDuration
	}

	if !tea.BoolValue(util.IsUnset(request.Tracing)) {
		body["Tracing"] = request.Tracing
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNamespaceScopeSidecarConfig"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNamespaceScopeSidecarConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNamespaceScopeSidecarConfig(request *UpdateNamespaceScopeSidecarConfigRequest) (_result *UpdateNamespaceScopeSidecarConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNamespaceScopeSidecarConfigResponse{}
	_body, _err := client.UpdateNamespaceScopeSidecarConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSwimLaneWithOptions(request *UpdateSwimLaneRequest, runtime *util.RuntimeOptions) (_result *UpdateSwimLaneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorKey)) {
		body["LabelSelectorKey"] = request.LabelSelectorKey
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelectorValue)) {
		body["LabelSelectorValue"] = request.LabelSelectorValue
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicesList)) {
		body["ServicesList"] = request.ServicesList
	}

	if !tea.BoolValue(util.IsUnset(request.SwimLaneName)) {
		body["SwimLaneName"] = request.SwimLaneName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSwimLane"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSwimLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSwimLane(request *UpdateSwimLaneRequest) (_result *UpdateSwimLaneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSwimLaneResponse{}
	_body, _err := client.UpdateSwimLaneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSwimLaneGroupWithOptions(request *UpdateSwimLaneGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateSwimLaneGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FallbackTarget)) {
		body["FallbackTarget"] = request.FallbackTarget
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicesList)) {
		body["ServicesList"] = request.ServicesList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSwimLaneGroup"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSwimLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSwimLaneGroup(request *UpdateSwimLaneGroupRequest) (_result *UpdateSwimLaneGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSwimLaneGroupResponse{}
	_body, _err := client.UpdateSwimLaneGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeMeshEditionPartiallyWithOptions(request *UpgradeMeshEditionPartiallyRequest, runtime *util.RuntimeOptions) (_result *UpgradeMeshEditionPartiallyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASMGatewayContinue)) {
		body["ASMGatewayContinue"] = request.ASMGatewayContinue
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectedVersion)) {
		body["ExpectedVersion"] = request.ExpectedVersion
	}

	if !tea.BoolValue(util.IsUnset(request.PreCheck)) {
		body["PreCheck"] = request.PreCheck
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchToPro)) {
		body["SwitchToPro"] = request.SwitchToPro
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeGatewayRecords)) {
		body["UpgradeGatewayRecords"] = request.UpgradeGatewayRecords
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeMeshEditionPartially"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeMeshEditionPartiallyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeMeshEditionPartially(request *UpgradeMeshEditionPartiallyRequest) (_result *UpgradeMeshEditionPartiallyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeMeshEditionPartiallyResponse{}
	_body, _err := client.UpgradeMeshEditionPartiallyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeMeshVersionWithOptions(request *UpgradeMeshVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeMeshVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PreCheck)) {
		query["PreCheck"] = request.PreCheck
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		query["ServiceMeshId"] = request.ServiceMeshId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeMeshVersion"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeMeshVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeMeshVersion(request *UpgradeMeshVersionRequest) (_result *UpgradeMeshVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeMeshVersionResponse{}
	_body, _err := client.UpgradeMeshVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
