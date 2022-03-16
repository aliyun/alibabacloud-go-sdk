// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddClusterIntoServiceMeshRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddClusterIntoServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddClusterIntoServiceMeshResponse) SetBody(v *AddClusterIntoServiceMeshResponseBody) *AddClusterIntoServiceMeshResponse {
	s.Body = v
	return s
}

type AddVMIntoServiceMeshRequest struct {
	EcsId         *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
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
	// Id of the request
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVMIntoServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddVMIntoServiceMeshResponse) SetBody(v *AddVMIntoServiceMeshResponseBody) *AddVMIntoServiceMeshResponse {
	s.Body = v
	return s
}

type CreateASMGatewayRequest struct {
	Body             *string `json:"Body,omitempty" xml:"Body,omitempty"`
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateASMGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateASMGatewayResponse) SetBody(v *CreateASMGatewayResponseBody) *CreateASMGatewayResponse {
	s.Body = v
	return s
}

type CreateServiceMeshRequest struct {
	AccessLogEnabled           *bool    `json:"AccessLogEnabled,omitempty" xml:"AccessLogEnabled,omitempty"`
	AccessLogFile              *string  `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	AccessLogFormat            *string  `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	AccessLogProject           *string  `json:"AccessLogProject,omitempty" xml:"AccessLogProject,omitempty"`
	AccessLogServiceEnabled    *bool    `json:"AccessLogServiceEnabled,omitempty" xml:"AccessLogServiceEnabled,omitempty"`
	AccessLogServiceHost       *string  `json:"AccessLogServiceHost,omitempty" xml:"AccessLogServiceHost,omitempty"`
	AccessLogServicePort       *int32   `json:"AccessLogServicePort,omitempty" xml:"AccessLogServicePort,omitempty"`
	ApiServerLoadBalancerSpec  *string  `json:"ApiServerLoadBalancerSpec,omitempty" xml:"ApiServerLoadBalancerSpec,omitempty"`
	ApiServerPublicEip         *bool    `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	AuditProject               *string  `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	AutoRenew                  *bool    `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod            *int32   `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	CRAggregationEnabled       *bool    `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	ChargeType                 *string  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterSpec                *string  `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	ConfigSourceEnabled        *bool    `json:"ConfigSourceEnabled,omitempty" xml:"ConfigSourceEnabled,omitempty"`
	ConfigSourceNacosID        *string  `json:"ConfigSourceNacosID,omitempty" xml:"ConfigSourceNacosID,omitempty"`
	ControlPlaneLogEnabled     *bool    `json:"ControlPlaneLogEnabled,omitempty" xml:"ControlPlaneLogEnabled,omitempty"`
	ControlPlaneLogProject     *string  `json:"ControlPlaneLogProject,omitempty" xml:"ControlPlaneLogProject,omitempty"`
	CustomizedPrometheus       *bool    `json:"CustomizedPrometheus,omitempty" xml:"CustomizedPrometheus,omitempty"`
	CustomizedZipkin           *bool    `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	DNSProxyingEnabled         *bool    `json:"DNSProxyingEnabled,omitempty" xml:"DNSProxyingEnabled,omitempty"`
	DubboFilterEnabled         *bool    `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	Edition                    *string  `json:"Edition,omitempty" xml:"Edition,omitempty"`
	EnableAudit                *bool    `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	EnableCRHistory            *bool    `json:"EnableCRHistory,omitempty" xml:"EnableCRHistory,omitempty"`
	EnableSDSServer            *bool    `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
	ExcludeIPRanges            *string  `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	ExcludeInboundPorts        *string  `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	ExcludeOutboundPorts       *string  `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	FilterGatewayClusterConfig *bool    `json:"FilterGatewayClusterConfig,omitempty" xml:"FilterGatewayClusterConfig,omitempty"`
	GatewayAPIEnabled          *bool    `json:"GatewayAPIEnabled,omitempty" xml:"GatewayAPIEnabled,omitempty"`
	GlobalRateLimitEnabled     *bool    `json:"GlobalRateLimitEnabled,omitempty" xml:"GlobalRateLimitEnabled,omitempty"`
	IncludeIPRanges            *string  `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	IstioVersion               *string  `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	KialiEnabled               *bool    `json:"KialiEnabled,omitempty" xml:"KialiEnabled,omitempty"`
	LocalityLBConf             *string  `json:"LocalityLBConf,omitempty" xml:"LocalityLBConf,omitempty"`
	LocalityLoadBalancing      *bool    `json:"LocalityLoadBalancing,omitempty" xml:"LocalityLoadBalancing,omitempty"`
	MSEEnabled                 *bool    `json:"MSEEnabled,omitempty" xml:"MSEEnabled,omitempty"`
	MultiBufferEnabled         *bool    `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	MultiBufferPollDelay       *string  `json:"MultiBufferPollDelay,omitempty" xml:"MultiBufferPollDelay,omitempty"`
	MysqlFilterEnabled         *bool    `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	Name                       *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OPALimitCPU                *string  `json:"OPALimitCPU,omitempty" xml:"OPALimitCPU,omitempty"`
	OPALimitMemory             *string  `json:"OPALimitMemory,omitempty" xml:"OPALimitMemory,omitempty"`
	OPALogLevel                *string  `json:"OPALogLevel,omitempty" xml:"OPALogLevel,omitempty"`
	OPARequestCPU              *string  `json:"OPARequestCPU,omitempty" xml:"OPARequestCPU,omitempty"`
	OPARequestMemory           *string  `json:"OPARequestMemory,omitempty" xml:"OPARequestMemory,omitempty"`
	OpaEnabled                 *bool    `json:"OpaEnabled,omitempty" xml:"OpaEnabled,omitempty"`
	OpenAgentPolicy            *bool    `json:"OpenAgentPolicy,omitempty" xml:"OpenAgentPolicy,omitempty"`
	Period                     *int32   `json:"Period,omitempty" xml:"Period,omitempty"`
	PilotLoadBalancerSpec      *string  `json:"PilotLoadBalancerSpec,omitempty" xml:"PilotLoadBalancerSpec,omitempty"`
	PrometheusUrl              *string  `json:"PrometheusUrl,omitempty" xml:"PrometheusUrl,omitempty"`
	ProxyLimitCPU              *string  `json:"ProxyLimitCPU,omitempty" xml:"ProxyLimitCPU,omitempty"`
	ProxyLimitMemory           *string  `json:"ProxyLimitMemory,omitempty" xml:"ProxyLimitMemory,omitempty"`
	ProxyRequestCPU            *string  `json:"ProxyRequestCPU,omitempty" xml:"ProxyRequestCPU,omitempty"`
	ProxyRequestMemory         *string  `json:"ProxyRequestMemory,omitempty" xml:"ProxyRequestMemory,omitempty"`
	RedisFilterEnabled         *bool    `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
	RegionId                   *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Telemetry                  *bool    `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	ThriftFilterEnabled        *bool    `json:"ThriftFilterEnabled,omitempty" xml:"ThriftFilterEnabled,omitempty"`
	TraceSampling              *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
	Tracing                    *bool    `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
	VSwitches                  *string  `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	VpcId                      *string  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WebAssemblyFilterEnabled   *bool    `json:"WebAssemblyFilterEnabled,omitempty" xml:"WebAssemblyFilterEnabled,omitempty"`
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

func (s *CreateServiceMeshRequest) SetFilterGatewayClusterConfig(v bool) *CreateServiceMeshRequest {
	s.FilterGatewayClusterConfig = &v
	return s
}

func (s *CreateServiceMeshRequest) SetGatewayAPIEnabled(v bool) *CreateServiceMeshRequest {
	s.GatewayAPIEnabled = &v
	return s
}

func (s *CreateServiceMeshRequest) SetGlobalRateLimitEnabled(v bool) *CreateServiceMeshRequest {
	s.GlobalRateLimitEnabled = &v
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

type CreateServiceMeshResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateServiceMeshResponse) SetBody(v *CreateServiceMeshResponseBody) *CreateServiceMeshResponse {
	s.Body = v
	return s
}

type DeleteServiceMeshRequest struct {
	Force           *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	RetainResources *string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty"`
	ServiceMeshId   *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteServiceMeshResponse) SetBody(v *DeleteServiceMeshResponseBody) *DeleteServiceMeshResponse {
	s.Body = v
	return s
}

type DescribeASMGatewayImportedServicesRequest struct {
	ASMGatewayName   *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	ImportedServices []*DescribeASMGatewayImportedServicesResponseBodyImportedServices `json:"ImportedServices,omitempty" xml:"ImportedServices,omitempty" type:"Repeated"`
	RequestId        *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ServiceName      *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeASMGatewayImportedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeASMGatewayImportedServicesResponse) SetBody(v *DescribeASMGatewayImportedServicesResponseBody) *DescribeASMGatewayImportedServicesResponse {
	s.Body = v
	return s
}

type DescribeAhasProRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAhasProRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAhasProRequest) GoString() string {
	return s.String()
}

func (s *DescribeAhasProRequest) SetRegionId(v string) *DescribeAhasProRequest {
	s.RegionId = &v
	return s
}

type DescribeAhasProResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAhasProResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAhasProResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAhasProResponseBody) SetRequestId(v string) *DescribeAhasProResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAhasProResponseBody) SetStatus(v bool) *DescribeAhasProResponseBody {
	s.Status = &v
	return s
}

type DescribeAhasProResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAhasProResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAhasProResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAhasProResponse) GoString() string {
	return s.String()
}

func (s *DescribeAhasProResponse) SetHeaders(v map[string]*string) *DescribeAhasProResponse {
	s.Headers = v
	return s
}

func (s *DescribeAhasProResponse) SetBody(v *DescribeAhasProResponseBody) *DescribeAhasProResponse {
	s.Body = v
	return s
}

type DescribeCCMVersionRequest struct {
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCCMVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCCMVersionResponse) SetBody(v *DescribeCCMVersionResponseBody) *DescribeCCMVersionResponse {
	s.Body = v
	return s
}

type DescribeCensRequest struct {
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
	Clusters  []*string `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCensResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCensResponse) SetBody(v *DescribeCensResponseBody) *DescribeCensResponse {
	s.Body = v
	return s
}

type DescribeClusterGrafanaRequest struct {
	K8sClusterId  *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
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
	Dashboards []*DescribeClusterGrafanaResponseBodyDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeClusterGrafanaResponse) SetBody(v *DescribeClusterGrafanaResponseBody) *DescribeClusterGrafanaResponse {
	s.Body = v
	return s
}

type DescribeClusterPrometheusRequest struct {
	K8sClusterId       *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	K8sClusterRegionId *string `json:"K8sClusterRegionId,omitempty" xml:"K8sClusterRegionId,omitempty"`
	ServiceMeshId      *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Prometheus *string `json:"Prometheus,omitempty" xml:"Prometheus,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterPrometheusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeClusterPrometheusResponse) SetBody(v *DescribeClusterPrometheusResponseBody) *DescribeClusterPrometheusResponse {
	s.Body = v
	return s
}

type DescribeClustersInServiceMeshRequest struct {
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
	Clusters  []*DescribeClustersInServiceMeshResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessLogDashboards   []*DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards `json:"AccessLogDashboards,omitempty" xml:"AccessLogDashboards,omitempty" type:"Repeated"`
	ClusterDomain         *string                                                                 `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	ClusterId             *string                                                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType           *string                                                                 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreationTime          *string                                                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorMessage          *string                                                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LogtailInstalledState *string                                                                 `json:"LogtailInstalledState,omitempty" xml:"LogtailInstalledState,omitempty"`
	Name                  *string                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId              *string                                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SgId                  *string                                                                 `json:"SgId,omitempty" xml:"SgId,omitempty"`
	State                 *string                                                                 `json:"State,omitempty" xml:"State,omitempty"`
	UpdateTime            *string                                                                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Version               *string                                                                 `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId                 *string                                                                 `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClustersInServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeClustersInServiceMeshResponse) SetBody(v *DescribeClustersInServiceMeshResponseBody) *DescribeClustersInServiceMeshResponse {
	s.Body = v
	return s
}

type DescribeCrTemplatesRequest struct {
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	Kind         *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
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
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ChineseName *string `json:"ChineseName,omitempty" xml:"ChineseName,omitempty"`
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	Yaml        *string `json:"Yaml,omitempty" xml:"Yaml,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCrTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCrTemplatesResponse) SetBody(v *DescribeCrTemplatesResponseBody) *DescribeCrTemplatesResponse {
	s.Body = v
	return s
}

type DescribeGuestClusterAccessLogDashboardsRequest struct {
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
	Dashboards   []*DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	K8sClusterId *string                                                          `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	RequestId    *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGuestClusterAccessLogDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeGuestClusterAccessLogDashboardsResponse) SetBody(v *DescribeGuestClusterAccessLogDashboardsResponseBody) *DescribeGuestClusterAccessLogDashboardsResponse {
	s.Body = v
	return s
}

type DescribeGuestClusterNamespacesRequest struct {
	GuestClusterID *string `json:"GuestClusterID,omitempty" xml:"GuestClusterID,omitempty"`
	ServiceMeshId  *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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

type DescribeGuestClusterNamespacesResponseBody struct {
	NsList    []*string `json:"NsList,omitempty" xml:"NsList,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGuestClusterNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGuestClusterNamespacesResponseBody) GoString() string {
	return s.String()
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGuestClusterNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeGuestClusterNamespacesResponse) SetBody(v *DescribeGuestClusterNamespacesResponseBody) *DescribeGuestClusterNamespacesResponse {
	s.Body = v
	return s
}

type DescribeGuestClusterPodsRequest struct {
	GuestClusterID *string `json:"GuestClusterID,omitempty" xml:"GuestClusterID,omitempty"`
	Namespace      *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ServiceMeshId  *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	PodList   []*string `json:"PodList,omitempty" xml:"PodList,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGuestClusterPodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeGuestClusterPodsResponse) SetBody(v *DescribeGuestClusterPodsResponseBody) *DescribeGuestClusterPodsResponse {
	s.Body = v
	return s
}

type DescribeImportedServicesDetailRequest struct {
	ASMGatewayName   *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Details   []*DescribeImportedServicesDetailResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ClusterIds  []*string                                                 `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	Labels      map[string]*string                                        `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Namespace   *string                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Ports       []*DescribeImportedServicesDetailResponseBodyDetailsPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	ServiceName *string                                                   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceType *string                                                   `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
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
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodePort   *int32  `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImportedServicesDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeImportedServicesDetailResponse) SetBody(v *DescribeImportedServicesDetailResponseBody) *DescribeImportedServicesDetailResponse {
	s.Body = v
	return s
}

type DescribeIngressGatewaysRequest struct {
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeIngressGatewaysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeIngressGatewaysRequest) SetServiceMeshId(v string) *DescribeIngressGatewaysRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeIngressGatewaysResponseBody struct {
	IngressGateways []map[string]interface{} `json:"IngressGateways,omitempty" xml:"IngressGateways,omitempty" type:"Repeated"`
	RequestId       *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIngressGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIngressGatewaysResponseBody) SetIngressGateways(v []map[string]interface{}) *DescribeIngressGatewaysResponseBody {
	s.IngressGateways = v
	return s
}

func (s *DescribeIngressGatewaysResponseBody) SetRequestId(v string) *DescribeIngressGatewaysResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIngressGatewaysResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIngressGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIngressGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressGatewaysResponse) GoString() string {
	return s.String()
}

func (s *DescribeIngressGatewaysResponse) SetHeaders(v map[string]*string) *DescribeIngressGatewaysResponse {
	s.Headers = v
	return s
}

func (s *DescribeIngressGatewaysResponse) SetBody(v *DescribeIngressGatewaysResponseBody) *DescribeIngressGatewaysResponse {
	s.Body = v
	return s
}

type DescribeNamespaceScopeSidecarConfigRequest struct {
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
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
	ConfigPatches *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches `json:"ConfigPatches,omitempty" xml:"ConfigPatches,omitempty" type:"Struct"`
	RequestId     *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ExcludeInboundPorts             *string                                                                                      `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	ExcludeOutboundIPRanges         *string                                                                                      `json:"ExcludeOutboundIPRanges,omitempty" xml:"ExcludeOutboundIPRanges,omitempty"`
	ExcludeOutboundPorts            *string                                                                                      `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	IncludeInboundPorts             *string                                                                                      `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	IncludeOutboundIPRanges         *string                                                                                      `json:"IncludeOutboundIPRanges,omitempty" xml:"IncludeOutboundIPRanges,omitempty"`
	IncludeOutboundPorts            *string                                                                                      `json:"IncludeOutboundPorts,omitempty" xml:"IncludeOutboundPorts,omitempty"`
	IstioDNSProxyEnabled            *bool                                                                                        `json:"IstioDNSProxyEnabled,omitempty" xml:"IstioDNSProxyEnabled,omitempty"`
	LifecycleStr                    *string                                                                                      `json:"LifecycleStr,omitempty" xml:"LifecycleStr,omitempty"`
	SidecarProxyInitResourceLimit   *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit   `json:"SidecarProxyInitResourceLimit,omitempty" xml:"SidecarProxyInitResourceLimit,omitempty" type:"Struct"`
	SidecarProxyInitResourceRequest *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest `json:"SidecarProxyInitResourceRequest,omitempty" xml:"SidecarProxyInitResourceRequest,omitempty" type:"Struct"`
	SidecarProxyResourceLimit       *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit       `json:"SidecarProxyResourceLimit,omitempty" xml:"SidecarProxyResourceLimit,omitempty" type:"Struct"`
	SidecarProxyResourceRequest     *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest     `json:"SidecarProxyResourceRequest,omitempty" xml:"SidecarProxyResourceRequest,omitempty" type:"Struct"`
	TerminationDrainDuration        *string                                                                                      `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GoString() string {
	return s.String()
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

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetIstioDNSProxyEnabled(v bool) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.IstioDNSProxyEnabled = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetLifecycleStr(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.LifecycleStr = &v
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

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit struct {
	ResourceCPULimit    *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
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
	ResourceCPURequest    *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
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
	ResourceCPULimit    *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
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
	ResourceCPURequest    *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
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

type DescribeNamespaceScopeSidecarConfigResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNamespaceScopeSidecarConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeNamespaceScopeSidecarConfigResponse) SetBody(v *DescribeNamespaceScopeSidecarConfigResponseBody) *DescribeNamespaceScopeSidecarConfigResponse {
	s.Body = v
	return s
}

type DescribeNodesInstanceTypeRequest struct {
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
	InstanceTypes []*DescribeNodesInstanceTypeResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	MultiBufferEnabled *bool   `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	NodeType           *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeNodesInstanceTypeResponseBodyInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodesInstanceTypeResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetMultiBufferEnabled(v bool) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.MultiBufferEnabled = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponseBodyInstanceTypes) SetNodeType(v string) *DescribeNodesInstanceTypeResponseBodyInstanceTypes {
	s.NodeType = &v
	return s
}

type DescribeNodesInstanceTypeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNodesInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeNodesInstanceTypeResponse) SetBody(v *DescribeNodesInstanceTypeResponseBody) *DescribeNodesInstanceTypeResponse {
	s.Body = v
	return s
}

type DescribeReusableSlbRequest struct {
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
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
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeReusableSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeReusableSlbResponse) SetBody(v *DescribeReusableSlbResponseBody) *DescribeReusableSlbResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshAdditionalStatusRequest struct {
	CheckMode     *string `json:"CheckMode,omitempty" xml:"CheckMode,omitempty"`
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
	ClusterStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty" type:"Struct"`
	RequestId     *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessLogProjectStatus      *string                                                                                  `json:"AccessLogProjectStatus,omitempty" xml:"AccessLogProjectStatus,omitempty"`
	ApiServerEIPStatus          *string                                                                                  `json:"ApiServerEIPStatus,omitempty" xml:"ApiServerEIPStatus,omitempty"`
	ApiServerLoadBalancerStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus `json:"ApiServerLoadBalancerStatus,omitempty" xml:"ApiServerLoadBalancerStatus,omitempty" type:"Struct"`
	AuditProjectStatus          *string                                                                                  `json:"AuditProjectStatus,omitempty" xml:"AuditProjectStatus,omitempty"`
	ControlPlaneProjectStatus   *string                                                                                  `json:"ControlPlaneProjectStatus,omitempty" xml:"ControlPlaneProjectStatus,omitempty"`
	LogtailStatusRecord         map[string]*ClusterStatusLogtailStatusRecordValue                                        `json:"LogtailStatusRecord,omitempty" xml:"LogtailStatusRecord,omitempty"`
	PilotLoadBalancerStatus     *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus     `json:"PilotLoadBalancerStatus,omitempty" xml:"PilotLoadBalancerStatus,omitempty" type:"Struct"`
	SgStatus                    *string                                                                                  `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
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

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetControlPlaneProjectStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.ControlPlaneProjectStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetLogtailStatusRecord(v map[string]*ClusterStatusLogtailStatusRecordValue) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.LogtailStatusRecord = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetPilotLoadBalancerStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.PilotLoadBalancerStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetSgStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.SgStatus = &v
	return s
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus struct {
	SLBBackEndServerNumStatus *string `json:"SLBBackEndServerNumStatus,omitempty" xml:"SLBBackEndServerNumStatus,omitempty"`
	SLBExistStatus            *string `json:"SLBExistStatus,omitempty" xml:"SLBExistStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetSLBBackEndServerNumStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.SLBBackEndServerNumStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetSLBExistStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.SLBExistStatus = &v
	return s
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus struct {
	SLBBackEndServerNumStatus *string `json:"SLBBackEndServerNumStatus,omitempty" xml:"SLBBackEndServerNumStatus,omitempty"`
	SLBExistStatus            *string `json:"SLBExistStatus,omitempty" xml:"SLBExistStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) GoString() string {
	return s.String()
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
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshAdditionalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshAdditionalStatusResponse) SetBody(v *DescribeServiceMeshAdditionalStatusResponseBody) *DescribeServiceMeshAdditionalStatusResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshClustersRequest struct {
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersRequest) SetServiceMeshId(v string) *DescribeServiceMeshClustersRequest {
	s.ServiceMeshId = &v
	return s
}

type DescribeServiceMeshClustersResponseBody struct {
	Clusters  []*DescribeServiceMeshClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeServiceMeshClustersResponseBody) SetRequestId(v string) *DescribeServiceMeshClustersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceMeshClustersResponseBodyClusters struct {
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType   *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreationTime  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ForbiddenFlag *int64  `json:"ForbiddenFlag,omitempty" xml:"ForbiddenFlag,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	SgId          *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshClustersResponse) SetBody(v *DescribeServiceMeshClustersResponseBody) *DescribeServiceMeshClustersResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshDetailRequest struct {
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
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ClusterSpec     *string                                                          `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	Clusters        []*string                                                        `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	Endpoints       *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints       `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	ServiceMeshInfo *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo `json:"ServiceMeshInfo,omitempty" xml:"ServiceMeshInfo,omitempty" type:"Struct"`
	Spec            *DescribeServiceMeshDetailResponseBodyServiceMeshSpec            `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetServiceMeshInfo(v *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.ServiceMeshInfo = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) SetSpec(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) *DescribeServiceMeshDetailResponseBodyServiceMesh {
	s.Spec = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints struct {
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	IntranetPilotEndpoint     *string `json:"IntranetPilotEndpoint,omitempty" xml:"IntranetPilotEndpoint,omitempty"`
	PublicApiServerEndpoint   *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
	PublicPilotEndpoint       *string `json:"PublicPilotEndpoint,omitempty" xml:"PublicPilotEndpoint,omitempty"`
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
	CreationTime  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile       *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	LoadBalancer *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Struct"`
	MeshConfig   *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig   `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	Network      *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork      `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
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
	ApiServerLoadbalancerId   *string `json:"ApiServerLoadbalancerId,omitempty" xml:"ApiServerLoadbalancerId,omitempty"`
	ApiServerPublicEip        *bool   `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	PilotPublicEip            *bool   `json:"PilotPublicEip,omitempty" xml:"PilotPublicEip,omitempty"`
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
	AccessLog                   *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog                   `json:"AccessLog,omitempty" xml:"AccessLog,omitempty" type:"Struct"`
	Audit                       *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit                       `json:"Audit,omitempty" xml:"Audit,omitempty" type:"Struct"`
	ControlPlaneLogInfo         *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo         `json:"ControlPlaneLogInfo,omitempty" xml:"ControlPlaneLogInfo,omitempty" type:"Struct"`
	CustomizedZipkin            *bool                                                                                      `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	Edition                     *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition                     `json:"Edition,omitempty" xml:"Edition,omitempty" type:"Struct"`
	EnableLocalityLB            *bool                                                                                      `json:"EnableLocalityLB,omitempty" xml:"EnableLocalityLB,omitempty"`
	ExcludeIPRanges             *string                                                                                    `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	ExcludeInboundPorts         *string                                                                                    `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	ExcludeOutboundPorts        *string                                                                                    `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	ExtraConfiguration          *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration          `json:"ExtraConfiguration,omitempty" xml:"ExtraConfiguration,omitempty" type:"Struct"`
	IncludeIPRanges             *string                                                                                    `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	K8sNewAPIsSupport           *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport           `json:"K8sNewAPIsSupport,omitempty" xml:"K8sNewAPIsSupport,omitempty" type:"Struct"`
	Kiali                       *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali                       `json:"Kiali,omitempty" xml:"Kiali,omitempty" type:"Struct"`
	LocalityLB                  *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB                  `json:"LocalityLB,omitempty" xml:"LocalityLB,omitempty" type:"Struct"`
	MSE                         *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE                         `json:"MSE,omitempty" xml:"MSE,omitempty" type:"Struct"`
	OPA                         *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA                         `json:"OPA,omitempty" xml:"OPA,omitempty" type:"Struct"`
	OutboundTrafficPolicy       *string                                                                                    `json:"OutboundTrafficPolicy,omitempty" xml:"OutboundTrafficPolicy,omitempty"`
	Pilot                       *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot                       `json:"Pilot,omitempty" xml:"Pilot,omitempty" type:"Struct"`
	Prometheus                  *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus                  `json:"Prometheus,omitempty" xml:"Prometheus,omitempty" type:"Struct"`
	ProtocolSupport             *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport             `json:"ProtocolSupport,omitempty" xml:"ProtocolSupport,omitempty" type:"Struct"`
	Proxy                       *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy                       `json:"Proxy,omitempty" xml:"Proxy,omitempty" type:"Struct"`
	SidecarInjector             *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector             `json:"SidecarInjector,omitempty" xml:"SidecarInjector,omitempty" type:"Struct"`
	Telemetry                   *bool                                                                                      `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	Tracing                     *bool                                                                                      `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
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
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
	AuditProjectStatus *string `json:"AuditProjectStatus,omitempty" xml:"AuditProjectStatus,omitempty"`
	Enabled            *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Project            *string `json:"Project,omitempty" xml:"Project,omitempty"`
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
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) SetProject(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo {
	s.Project = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition struct {
	IstiodImageTag *string `json:"IstiodImageTag,omitempty" xml:"IstiodImageTag,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProxyImageTag  *string `json:"ProxyImageTag,omitempty" xml:"ProxyImageTag,omitempty"`
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
	CRAggregationEnabled            *bool                                                                                                            `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	DiscoverySelectors              []map[string]interface{}                                                                                         `json:"DiscoverySelectors,omitempty" xml:"DiscoverySelectors,omitempty" type:"Repeated"`
	IstioCRHistory                  *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory                  `json:"IstioCRHistory,omitempty" xml:"IstioCRHistory,omitempty" type:"Struct"`
	Lifecycle                       *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle                       `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty" type:"Struct"`
	MultiBuffer                     *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer                     `json:"MultiBuffer,omitempty" xml:"MultiBuffer,omitempty" type:"Struct"`
	RateLimit                       *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationRateLimit                       `json:"RateLimit,omitempty" xml:"RateLimit,omitempty" type:"Struct"`
	SidecarProxyInitResourceLimit   *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit   `json:"SidecarProxyInitResourceLimit,omitempty" xml:"SidecarProxyInitResourceLimit,omitempty" type:"Struct"`
	SidecarProxyInitResourceRequest *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest `json:"SidecarProxyInitResourceRequest,omitempty" xml:"SidecarProxyInitResourceRequest,omitempty" type:"Struct"`
	TerminationDrainDuration        *string                                                                                                          `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GoString() string {
	return s.String()
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetRateLimit(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationRateLimit) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.RateLimit = v
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

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory struct {
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
	PostStart *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart `json:"PostStart,omitempty" xml:"PostStart,omitempty" type:"Struct"`
	PreStop   *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop   `json:"PreStop,omitempty" xml:"PreStop,omitempty" type:"Struct"`
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
	Exec      *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	HTTPGet   *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet   `json:"HTTPGet,omitempty" xml:"HTTPGet,omitempty" type:"Struct"`
	TCPSocket *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket `json:"TCPSocket,omitempty" xml:"TCPSocket,omitempty" type:"Struct"`
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) SetHTTPGet(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart {
	s.HTTPGet = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) SetTCPSocket(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart {
	s.TCPSocket = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
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

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet struct {
	HTTPHeaders []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders `json:"HTTPHeaders,omitempty" xml:"HTTPHeaders,omitempty" type:"Repeated"`
	Host        *string                                                                                                                 `json:"Host,omitempty" xml:"Host,omitempty"`
	Port        *string                                                                                                                 `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme      *string                                                                                                                 `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet) SetHTTPHeaders(v []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet {
	s.HTTPHeaders = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet {
	s.Port = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet) SetScheme(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGet {
	s.Scheme = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders) SetName(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders) SetValue(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHTTPGetHTTPHeaders {
	s.Value = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket struct {
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTCPSocket {
	s.Port = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop struct {
	Exec      *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	HTTPGet   *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet   `json:"HTTPGet,omitempty" xml:"HTTPGet,omitempty" type:"Struct"`
	TCPSocket *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket `json:"TCPSocket,omitempty" xml:"TCPSocket,omitempty" type:"Struct"`
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) SetHTTPGet(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop {
	s.HTTPGet = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) SetTCPSocket(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop {
	s.TCPSocket = v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
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

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet struct {
	HTTPHeaders []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders `json:"HTTPHeaders,omitempty" xml:"HTTPHeaders,omitempty" type:"Repeated"`
	Host        *string                                                                                                               `json:"Host,omitempty" xml:"Host,omitempty"`
	Port        *string                                                                                                               `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme      *string                                                                                                               `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet) SetHTTPHeaders(v []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet {
	s.HTTPHeaders = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet {
	s.Port = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet) SetScheme(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGet {
	s.Scheme = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders) SetName(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders) SetValue(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHTTPGetHTTPHeaders {
	s.Value = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket struct {
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTCPSocket {
	s.Port = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer struct {
	Enabled   *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationRateLimit struct {
	EnableGlobalRateLimit *bool `json:"EnableGlobalRateLimit,omitempty" xml:"EnableGlobalRateLimit,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationRateLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationRateLimit) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationRateLimit) SetEnableGlobalRateLimit(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationRateLimit {
	s.EnableGlobalRateLimit = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit struct {
	ResourceCPULimit    *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
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
	ResourceCPURequest    *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
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
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Distribute map[string]interface{} `json:"Distribute,omitempty" xml:"Distribute,omitempty"`
	Enabled    *bool                  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Failover   map[string]interface{} `json:"Failover,omitempty" xml:"Failover,omitempty"`
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
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	LimitCPU      *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	LimitMemory   *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	LogLevel      *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	RequestCPU    *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
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
	ConfigSource  *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource `json:"ConfigSource,omitempty" xml:"ConfigSource,omitempty" type:"Struct"`
	Feature       *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature      `json:"Feature,omitempty" xml:"Feature,omitempty" type:"Struct"`
	Http10Enabled *bool                                                                            `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
	TraceSampling *float32                                                                         `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
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
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
	EnableSDSServer            *bool `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
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
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	UseExternal *bool   `json:"UseExternal,omitempty" xml:"UseExternal,omitempty"`
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
	DubboFilterEnabled  *bool `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	MysqlFilterEnabled  *bool `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	RedisFilterEnabled  *bool `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
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
	AccessLogFile           *string `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	AccessLogFormat         *string `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	AccessLogServiceEnabled *bool   `json:"AccessLogServiceEnabled,omitempty" xml:"AccessLogServiceEnabled,omitempty"`
	AccessLogServiceHost    *string `json:"AccessLogServiceHost,omitempty" xml:"AccessLogServiceHost,omitempty"`
	AccessLogServicePort    *int32  `json:"AccessLogServicePort,omitempty" xml:"AccessLogServicePort,omitempty"`
	ClusterDomain           *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	EnableDNSProxying       *bool   `json:"EnableDNSProxying,omitempty" xml:"EnableDNSProxying,omitempty"`
	LimitCPU                *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	LimitMemory             *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	RequestCPU              *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestMemory           *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
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
	AutoInjectionPolicyEnabled   *bool                                                                                              `json:"AutoInjectionPolicyEnabled,omitempty" xml:"AutoInjectionPolicyEnabled,omitempty"`
	EnableNamespacesByDefault    *bool                                                                                              `json:"EnableNamespacesByDefault,omitempty" xml:"EnableNamespacesByDefault,omitempty"`
	InitCNIConfiguration         *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration `json:"InitCNIConfiguration,omitempty" xml:"InitCNIConfiguration,omitempty" type:"Struct"`
	LimitCPU                     *string                                                                                            `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	LimitMemory                  *string                                                                                            `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	RequestCPU                   *string                                                                                            `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestMemory                *string                                                                                            `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	SidecarInjectorWebhookAsYaml *string                                                                                            `json:"SidecarInjectorWebhookAsYaml,omitempty" xml:"SidecarInjectorWebhookAsYaml,omitempty"`
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) SetSidecarInjectorWebhookAsYaml(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	s.SidecarInjectorWebhookAsYaml = &v
	return s
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration struct {
	Enabled           *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitches       []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshDetailResponse) SetBody(v *DescribeServiceMeshDetailResponseBody) *DescribeServiceMeshDetailResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshKubeconfigRequest struct {
	PrivateIpAddress *bool   `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMeshKubeconfigResponseBody) GoString() string {
	return s.String()
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshKubeconfigResponse) SetBody(v *DescribeServiceMeshKubeconfigResponseBody) *DescribeServiceMeshKubeconfigResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshLogsRequest struct {
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
	Logs      []*DescribeServiceMeshLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreationTime  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Log           *string `json:"Log,omitempty" xml:"Log,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshLogsResponse) SetBody(v *DescribeServiceMeshLogsResponseBody) *DescribeServiceMeshLogsResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshProxyStatusRequest struct {
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
	Code        *string                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	ProxyStatus []*DescribeServiceMeshProxyStatusResponseBodyProxyStatus `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty" type:"Repeated"`
	RequestId   *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *string                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
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
	ClusterSynced   *string `json:"ClusterSynced,omitempty" xml:"ClusterSynced,omitempty"`
	EndpointPercent *string `json:"EndpointPercent,omitempty" xml:"EndpointPercent,omitempty"`
	EndpointSynced  *string `json:"EndpointSynced,omitempty" xml:"EndpointSynced,omitempty"`
	IstioVersion    *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	ListenerSynced  *string `json:"ListenerSynced,omitempty" xml:"ListenerSynced,omitempty"`
	ProxyId         *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	ProxyVersion    *string `json:"ProxyVersion,omitempty" xml:"ProxyVersion,omitempty"`
	RouteSynced     *string `json:"RouteSynced,omitempty" xml:"RouteSynced,omitempty"`
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshProxyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshProxyStatusResponse) SetBody(v *DescribeServiceMeshProxyStatusResponseBody) *DescribeServiceMeshProxyStatusResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshUpgradeStatusRequest struct {
	AllIstioGatewayFullNames *string `json:"AllIstioGatewayFullNames,omitempty" xml:"AllIstioGatewayFullNames,omitempty"`
	GuestClusterIds          *string `json:"GuestClusterIds,omitempty" xml:"GuestClusterIds,omitempty"`
	ServiceMeshId            *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	RequestId     *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	FinishedGatewaysNum *int64                                            `json:"FinishedGatewaysNum,omitempty" xml:"FinishedGatewaysNum,omitempty"`
	GatewayStatusRecord map[string]*UpgradeDetailGatewayStatusRecordValue `json:"GatewayStatusRecord,omitempty" xml:"GatewayStatusRecord,omitempty"`
	MeshStatus          *string                                           `json:"MeshStatus,omitempty" xml:"MeshStatus,omitempty"`
	TotalGatewaysNum    *int64                                            `json:"TotalGatewaysNum,omitempty" xml:"TotalGatewaysNum,omitempty"`
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshUpgradeStatusResponse) SetBody(v *DescribeServiceMeshUpgradeStatusResponseBody) *DescribeServiceMeshUpgradeStatusResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshVMsRequest struct {
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
	// Id of the request
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VMs       []*DescribeServiceMeshVMsResponseBodyVMs `json:"VMs,omitempty" xml:"VMs,omitempty" type:"Repeated"`
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
	HasTag           *bool   `json:"HasTag,omitempty" xml:"HasTag,omitempty"`
	HostName         *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpAddress        *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshVMsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshVMsResponse) SetBody(v *DescribeServiceMeshVMsResponseBody) *DescribeServiceMeshVMsResponse {
	s.Body = v
	return s
}

type DescribeServiceMeshesResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ClusterSpec     *string                                                        `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	Clusters        []*string                                                      `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	Endpoints       *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints       `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	ServiceMeshInfo *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo `json:"ServiceMeshInfo,omitempty" xml:"ServiceMeshInfo,omitempty" type:"Struct"`
	Spec            *DescribeServiceMeshesResponseBodyServiceMeshesSpec            `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetServiceMeshInfo(v *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.ServiceMeshInfo = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetSpec(v *DescribeServiceMeshesResponseBodyServiceMeshesSpec) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.Spec = v
	return s
}

type DescribeServiceMeshesResponseBodyServiceMeshesEndpoints struct {
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	IntranetPilotEndpoint     *string `json:"IntranetPilotEndpoint,omitempty" xml:"IntranetPilotEndpoint,omitempty"`
	PublicApiServerEndpoint   *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
	PublicPilotEndpoint       *string `json:"PublicPilotEndpoint,omitempty" xml:"PublicPilotEndpoint,omitempty"`
	ReverseTunnelEndpoint     *string `json:"ReverseTunnelEndpoint,omitempty" xml:"ReverseTunnelEndpoint,omitempty"`
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
	CreationTime  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile       *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	LoadBalancer *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Struct"`
	MeshConfig   *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig   `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	Network      *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork      `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
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
	ApiServerLoadbalancerId   *string `json:"ApiServerLoadbalancerId,omitempty" xml:"ApiServerLoadbalancerId,omitempty"`
	ApiServerPublicEip        *bool   `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	PilotPublicEip            *bool   `json:"PilotPublicEip,omitempty" xml:"PilotPublicEip,omitempty"`
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
	Mtls                  *bool                                                                        `json:"Mtls,omitempty" xml:"Mtls,omitempty"`
	OutboundTrafficPolicy *string                                                                      `json:"OutboundTrafficPolicy,omitempty" xml:"OutboundTrafficPolicy,omitempty"`
	Pilot                 *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot           `json:"Pilot,omitempty" xml:"Pilot,omitempty" type:"Struct"`
	SidecarInjector       *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector `json:"SidecarInjector,omitempty" xml:"SidecarInjector,omitempty" type:"Struct"`
	StrictMtls            *bool                                                                        `json:"StrictMtls,omitempty" xml:"StrictMtls,omitempty"`
	Telemetry             *bool                                                                        `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	Tracing               *bool                                                                        `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
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
	Http10Enabled *bool    `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
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
	AutoInjectionPolicyEnabled *bool                                                                                            `json:"AutoInjectionPolicyEnabled,omitempty" xml:"AutoInjectionPolicyEnabled,omitempty"`
	EnableNamespacesByDefault  *bool                                                                                            `json:"EnableNamespacesByDefault,omitempty" xml:"EnableNamespacesByDefault,omitempty"`
	InitCNIConfiguration       *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration `json:"InitCNIConfiguration,omitempty" xml:"InitCNIConfiguration,omitempty" type:"Struct"`
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
	Enabled           *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitches       []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

type DescribeServiceMeshesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServiceMeshesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeServiceMeshesResponse) SetBody(v *DescribeServiceMeshesResponseBody) *DescribeServiceMeshesResponse {
	s.Body = v
	return s
}

type DescribeUpgradeVersionRequest struct {
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
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version   *DescribeUpgradeVersionResponseBodyVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
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
	IstioOperatorVersion *string `json:"IstioOperatorVersion,omitempty" xml:"IstioOperatorVersion,omitempty"`
	IstioVersion         *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	KubernetesVersion    *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUpgradeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeUpgradeVersionResponse) SetBody(v *DescribeUpgradeVersionResponseBody) *DescribeUpgradeVersionResponse {
	s.Body = v
	return s
}

type DescribeUsersWithPermissionsRequest struct {
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
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UIDs      []*string `json:"UIDs,omitempty" xml:"UIDs,omitempty" type:"Repeated"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUsersWithPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeUsersWithPermissionsResponse) SetBody(v *DescribeUsersWithPermissionsResponseBody) *DescribeUsersWithPermissionsResponse {
	s.Body = v
	return s
}

type DescribeVMsInServiceMeshRequest struct {
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
	// Id of the request
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VMs       []*DescribeVMsInServiceMeshResponseBodyVMs `json:"VMs,omitempty" xml:"VMs,omitempty" type:"Repeated"`
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
	HasTag           *bool   `json:"HasTag,omitempty" xml:"HasTag,omitempty"`
	HostName         *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpAddress        *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVMsInServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVMsInServiceMeshResponse) SetBody(v *DescribeVMsInServiceMeshResponseBody) *DescribeVMsInServiceMeshResponse {
	s.Body = v
	return s
}

type DescribeVSwitchesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId    *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// VSwitches
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
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

type DescribeVSwitchesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVSwitchesResponse) SetBody(v *DescribeVSwitchesResponseBody) *DescribeVSwitchesResponse {
	s.Body = v
	return s
}

type DescribeVersionsResponseBody struct {
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Edition  *string   `json:"Edition,omitempty" xml:"Edition,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVersionsResponse) SetBody(v *DescribeVersionsResponseBody) *DescribeVersionsResponse {
	s.Body = v
	return s
}

type DescribeVpcsRequest struct {
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
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Vpcs
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
	IsDefault *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName   *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVpcsResponse) SetBody(v *DescribeVpcsResponseBody) *DescribeVpcsResponse {
	s.Body = v
	return s
}

type GetCaCertRequest struct {
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
	// base64 encode format
	CaCert    *string `json:"CaCert,omitempty" xml:"CaCert,omitempty"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCaCertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCaCertResponse) SetBody(v *GetCaCertResponseBody) *GetCaCertResponse {
	s.Body = v
	return s
}

type GetRegisteredServiceEndpointsRequest struct {
	ClusterIds    *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	ServiceType   *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
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
	EndPointSlice    *GetRegisteredServiceEndpointsResponseBodyEndPointSlice      `json:"EndPointSlice,omitempty" xml:"EndPointSlice,omitempty" type:"Struct"`
	RequestId        *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	EndpointsDetails []*GetRegisteredServiceEndpointsResponseBodyEndPointSliceEndpointsDetails `json:"EndpointsDetails,omitempty" xml:"EndpointsDetails,omitempty" type:"Repeated"`
	Location         *string                                                                   `json:"Location,omitempty" xml:"Location,omitempty"`
	Namespace        *string                                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ServiceName      *string                                                                   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
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
	Address         *string  `json:"Address,omitempty" xml:"Address,omitempty"`
	Hostname        *string  `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	PodName         *string  `json:"PodName,omitempty" xml:"PodName,omitempty"`
	Ports           []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	Region          *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	SidecarInjected *bool    `json:"SidecarInjected,omitempty" xml:"SidecarInjected,omitempty"`
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
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRegisteredServiceEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRegisteredServiceEndpointsResponse) SetBody(v *GetRegisteredServiceEndpointsResponseBody) *GetRegisteredServiceEndpointsResponse {
	s.Body = v
	return s
}

type GetRegisteredServiceNamespacesRequest struct {
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
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRegisteredServiceNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRegisteredServiceNamespacesResponse) SetBody(v *GetRegisteredServiceNamespacesResponseBody) *GetRegisteredServiceNamespacesResponse {
	s.Body = v
	return s
}

type GetVmAppMeshInfoRequest struct {
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
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVmAppMeshInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetVmAppMeshInfoResponse) SetBody(v *GetVmAppMeshInfoResponseBody) *GetVmAppMeshInfoResponse {
	s.Body = v
	return s
}

type GetVmMetaRequest struct {
	Namespace      *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	ServiceMeshId  *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	TrustDomain    *string `json:"TrustDomain,omitempty" xml:"TrustDomain,omitempty"`
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
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	EnvoyEnvContent *string `json:"EnvoyEnvContent,omitempty" xml:"EnvoyEnvContent,omitempty"`
	HostsContent    *string `json:"HostsContent,omitempty" xml:"HostsContent,omitempty"`
	TokenContent    *string `json:"TokenContent,omitempty" xml:"TokenContent,omitempty"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVmMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetVmMetaResponse) SetBody(v *GetVmMetaResponseBody) *GetVmMetaResponse {
	s.Body = v
	return s
}

type ModifyServiceMeshNameRequest struct {
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyServiceMeshNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyServiceMeshNameResponse) SetBody(v *ModifyServiceMeshNameResponseBody) *ModifyServiceMeshNameResponse {
	s.Body = v
	return s
}

type ReActivateAuditRequest struct {
	EnableAudit   *bool   `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
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
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReActivateAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReActivateAuditResponse) SetBody(v *ReActivateAuditResponseBody) *ReActivateAuditResponse {
	s.Body = v
	return s
}

type RemoveClusterFromServiceMeshRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveClusterFromServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveClusterFromServiceMeshResponse) SetBody(v *RemoveClusterFromServiceMeshResponseBody) *RemoveClusterFromServiceMeshResponse {
	s.Body = v
	return s
}

type RemoveVMFromServiceMeshRequest struct {
	EcsId         *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
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
	// Id of the request
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveVMFromServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveVMFromServiceMeshResponse) SetBody(v *RemoveVMFromServiceMeshResponseBody) *RemoveVMFromServiceMeshResponse {
	s.Body = v
	return s
}

type RevokeKubeconfigRequest struct {
	PrivateIpAddress *bool   `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokeKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RevokeKubeconfigResponse) SetBody(v *RevokeKubeconfigResponseBody) *RevokeKubeconfigResponse {
	s.Body = v
	return s
}

type RunDiagnosisRequest struct {
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s RunDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *RunDiagnosisRequest) SetServiceMeshId(v string) *RunDiagnosisRequest {
	s.ServiceMeshId = &v
	return s
}

type RunDiagnosisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RunDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *RunDiagnosisResponseBody) SetRequestId(v string) *RunDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDiagnosisResponseBody) SetResult(v string) *RunDiagnosisResponseBody {
	s.Result = &v
	return s
}

type RunDiagnosisResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *RunDiagnosisResponse) SetHeaders(v map[string]*string) *RunDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *RunDiagnosisResponse) SetBody(v *RunDiagnosisResponseBody) *RunDiagnosisResponse {
	s.Body = v
	return s
}

type UpdateASMGatewayRequest struct {
	Body             *string `json:"Body,omitempty" xml:"Body,omitempty"`
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateASMGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateASMGatewayResponse) SetBody(v *UpdateASMGatewayResponseBody) *UpdateASMGatewayResponse {
	s.Body = v
	return s
}

type UpdateASMGatewayImportedServicesRequest struct {
	ASMGatewayName   *string `json:"ASMGatewayName,omitempty" xml:"ASMGatewayName,omitempty"`
	ServiceMeshId    *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	ServiceNames     *string `json:"ServiceNames,omitempty" xml:"ServiceNames,omitempty"`
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateASMGatewayImportedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateASMGatewayImportedServicesResponse) SetBody(v *UpdateASMGatewayImportedServicesResponseBody) *UpdateASMGatewayImportedServicesResponse {
	s.Body = v
	return s
}

type UpdateMeshFeatureRequest struct {
	AccessLogEnabled               *bool    `json:"AccessLogEnabled,omitempty" xml:"AccessLogEnabled,omitempty"`
	AccessLogFile                  *string  `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	AccessLogFormat                *string  `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	AccessLogProject               *string  `json:"AccessLogProject,omitempty" xml:"AccessLogProject,omitempty"`
	AccessLogServiceEnabled        *bool    `json:"AccessLogServiceEnabled,omitempty" xml:"AccessLogServiceEnabled,omitempty"`
	AccessLogServiceHost           *string  `json:"AccessLogServiceHost,omitempty" xml:"AccessLogServiceHost,omitempty"`
	AccessLogServicePort           *int32   `json:"AccessLogServicePort,omitempty" xml:"AccessLogServicePort,omitempty"`
	AuditProject                   *string  `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	AutoInjectionPolicyEnabled     *bool    `json:"AutoInjectionPolicyEnabled,omitempty" xml:"AutoInjectionPolicyEnabled,omitempty"`
	CRAggregationEnabled           *bool    `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	ClusterSpec                    *string  `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	CniEnabled                     *bool    `json:"CniEnabled,omitempty" xml:"CniEnabled,omitempty"`
	CniExcludeNamespaces           *string  `json:"CniExcludeNamespaces,omitempty" xml:"CniExcludeNamespaces,omitempty"`
	ConfigSourceEnabled            *bool    `json:"ConfigSourceEnabled,omitempty" xml:"ConfigSourceEnabled,omitempty"`
	ConfigSourceNacosID            *string  `json:"ConfigSourceNacosID,omitempty" xml:"ConfigSourceNacosID,omitempty"`
	CustomizedPrometheus           *bool    `json:"CustomizedPrometheus,omitempty" xml:"CustomizedPrometheus,omitempty"`
	CustomizedZipkin               *bool    `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	DNSProxyingEnabled             *bool    `json:"DNSProxyingEnabled,omitempty" xml:"DNSProxyingEnabled,omitempty"`
	DiscoverySelectors             *string  `json:"DiscoverySelectors,omitempty" xml:"DiscoverySelectors,omitempty"`
	DubboFilterEnabled             *bool    `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	EnableAudit                    *bool    `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	EnableCRHistory                *bool    `json:"EnableCRHistory,omitempty" xml:"EnableCRHistory,omitempty"`
	EnableNamespacesByDefault      *bool    `json:"EnableNamespacesByDefault,omitempty" xml:"EnableNamespacesByDefault,omitempty"`
	EnableSDSServer                *bool    `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
	ExcludeIPRanges                *string  `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	ExcludeInboundPorts            *string  `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	ExcludeOutboundPorts           *string  `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	FilterGatewayClusterConfig     *bool    `json:"FilterGatewayClusterConfig,omitempty" xml:"FilterGatewayClusterConfig,omitempty"`
	GatewayAPIEnabled              *bool    `json:"GatewayAPIEnabled,omitempty" xml:"GatewayAPIEnabled,omitempty"`
	GlobalRateLimitEnabled         *bool    `json:"GlobalRateLimitEnabled,omitempty" xml:"GlobalRateLimitEnabled,omitempty"`
	Http10Enabled                  *bool    `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
	IncludeIPRanges                *string  `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	IncludeInboundPorts            *string  `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	KialiEnabled                   *bool    `json:"KialiEnabled,omitempty" xml:"KialiEnabled,omitempty"`
	Lifecycle                      *string  `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	LocalityLBConf                 *string  `json:"LocalityLBConf,omitempty" xml:"LocalityLBConf,omitempty"`
	LocalityLoadBalancing          *bool    `json:"LocalityLoadBalancing,omitempty" xml:"LocalityLoadBalancing,omitempty"`
	MSEEnabled                     *bool    `json:"MSEEnabled,omitempty" xml:"MSEEnabled,omitempty"`
	MultiBufferEnabled             *bool    `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	MultiBufferPollDelay           *string  `json:"MultiBufferPollDelay,omitempty" xml:"MultiBufferPollDelay,omitempty"`
	MysqlFilterEnabled             *bool    `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	OPALimitCPU                    *string  `json:"OPALimitCPU,omitempty" xml:"OPALimitCPU,omitempty"`
	OPALimitMemory                 *string  `json:"OPALimitMemory,omitempty" xml:"OPALimitMemory,omitempty"`
	OPALogLevel                    *string  `json:"OPALogLevel,omitempty" xml:"OPALogLevel,omitempty"`
	OPARequestCPU                  *string  `json:"OPARequestCPU,omitempty" xml:"OPARequestCPU,omitempty"`
	OPARequestMemory               *string  `json:"OPARequestMemory,omitempty" xml:"OPARequestMemory,omitempty"`
	OpaEnabled                     *bool    `json:"OpaEnabled,omitempty" xml:"OpaEnabled,omitempty"`
	OpenAgentPolicy                *bool    `json:"OpenAgentPolicy,omitempty" xml:"OpenAgentPolicy,omitempty"`
	OutboundTrafficPolicy          *string  `json:"OutboundTrafficPolicy,omitempty" xml:"OutboundTrafficPolicy,omitempty"`
	PrometheusUrl                  *string  `json:"PrometheusUrl,omitempty" xml:"PrometheusUrl,omitempty"`
	ProxyInitCPUResourceLimit      *string  `json:"ProxyInitCPUResourceLimit,omitempty" xml:"ProxyInitCPUResourceLimit,omitempty"`
	ProxyInitCPUResourceRequest    *string  `json:"ProxyInitCPUResourceRequest,omitempty" xml:"ProxyInitCPUResourceRequest,omitempty"`
	ProxyInitMemoryResourceLimit   *string  `json:"ProxyInitMemoryResourceLimit,omitempty" xml:"ProxyInitMemoryResourceLimit,omitempty"`
	ProxyInitMemoryResourceRequest *string  `json:"ProxyInitMemoryResourceRequest,omitempty" xml:"ProxyInitMemoryResourceRequest,omitempty"`
	ProxyLimitCPU                  *string  `json:"ProxyLimitCPU,omitempty" xml:"ProxyLimitCPU,omitempty"`
	ProxyLimitMemory               *string  `json:"ProxyLimitMemory,omitempty" xml:"ProxyLimitMemory,omitempty"`
	ProxyRequestCPU                *string  `json:"ProxyRequestCPU,omitempty" xml:"ProxyRequestCPU,omitempty"`
	ProxyRequestMemory             *string  `json:"ProxyRequestMemory,omitempty" xml:"ProxyRequestMemory,omitempty"`
	RedisFilterEnabled             *bool    `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
	ServiceMeshId                  *string  `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	SidecarInjectorLimitCPU        *string  `json:"SidecarInjectorLimitCPU,omitempty" xml:"SidecarInjectorLimitCPU,omitempty"`
	SidecarInjectorLimitMemory     *string  `json:"SidecarInjectorLimitMemory,omitempty" xml:"SidecarInjectorLimitMemory,omitempty"`
	SidecarInjectorRequestCPU      *string  `json:"SidecarInjectorRequestCPU,omitempty" xml:"SidecarInjectorRequestCPU,omitempty"`
	SidecarInjectorRequestMemory   *string  `json:"SidecarInjectorRequestMemory,omitempty" xml:"SidecarInjectorRequestMemory,omitempty"`
	SidecarInjectorWebhookAsYaml   *string  `json:"SidecarInjectorWebhookAsYaml,omitempty" xml:"SidecarInjectorWebhookAsYaml,omitempty"`
	Telemetry                      *bool    `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	TerminationDrainDuration       *string  `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
	ThriftFilterEnabled            *bool    `json:"ThriftFilterEnabled,omitempty" xml:"ThriftFilterEnabled,omitempty"`
	TraceSampling                  *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
	Tracing                        *bool    `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
	WebAssemblyFilterEnabled       *bool    `json:"WebAssemblyFilterEnabled,omitempty" xml:"WebAssemblyFilterEnabled,omitempty"`
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

func (s *UpdateMeshFeatureRequest) SetGlobalRateLimitEnabled(v bool) *UpdateMeshFeatureRequest {
	s.GlobalRateLimitEnabled = &v
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

func (s *UpdateMeshFeatureRequest) SetKialiEnabled(v bool) *UpdateMeshFeatureRequest {
	s.KialiEnabled = &v
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

func (s *UpdateMeshFeatureRequest) SetTraceSampling(v float32) *UpdateMeshFeatureRequest {
	s.TraceSampling = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetTracing(v bool) *UpdateMeshFeatureRequest {
	s.Tracing = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetWebAssemblyFilterEnabled(v bool) *UpdateMeshFeatureRequest {
	s.WebAssemblyFilterEnabled = &v
	return s
}

type UpdateMeshFeatureResponseBody struct {
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMeshFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateMeshFeatureResponse) SetBody(v *UpdateMeshFeatureResponseBody) *UpdateMeshFeatureResponse {
	s.Body = v
	return s
}

type UpdateNamespaceScopeSidecarConfigRequest struct {
	ExcludeIPRanges                   *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	ExcludeInboundPorts               *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	ExcludeOutboundPorts              *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	IncludeIPRanges                   *string `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	IncludeInboundPorts               *string `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	IncludeOutboundPorts              *string `json:"IncludeOutboundPorts,omitempty" xml:"IncludeOutboundPorts,omitempty"`
	IstioDNSProxyEnabled              *bool   `json:"IstioDNSProxyEnabled,omitempty" xml:"IstioDNSProxyEnabled,omitempty"`
	Lifecycle                         *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	Namespace                         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProxyInitCPUResourceLimit         *string `json:"ProxyInitCPUResourceLimit,omitempty" xml:"ProxyInitCPUResourceLimit,omitempty"`
	ProxyInitCPUResourceRequest       *string `json:"ProxyInitCPUResourceRequest,omitempty" xml:"ProxyInitCPUResourceRequest,omitempty"`
	ProxyInitMemoryResourceLimit      *string `json:"ProxyInitMemoryResourceLimit,omitempty" xml:"ProxyInitMemoryResourceLimit,omitempty"`
	ProxyInitMemoryResourceRequest    *string `json:"ProxyInitMemoryResourceRequest,omitempty" xml:"ProxyInitMemoryResourceRequest,omitempty"`
	ServiceMeshId                     *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	SidecarProxyCPUResourceLimit      *string `json:"SidecarProxyCPUResourceLimit,omitempty" xml:"SidecarProxyCPUResourceLimit,omitempty"`
	SidecarProxyCPUResourceRequest    *string `json:"SidecarProxyCPUResourceRequest,omitempty" xml:"SidecarProxyCPUResourceRequest,omitempty"`
	SidecarProxyMemoryResourceLimit   *string `json:"SidecarProxyMemoryResourceLimit,omitempty" xml:"SidecarProxyMemoryResourceLimit,omitempty"`
	SidecarProxyMemoryResourceRequest *string `json:"SidecarProxyMemoryResourceRequest,omitempty" xml:"SidecarProxyMemoryResourceRequest,omitempty"`
	TerminationDrainDuration          *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
}

func (s UpdateNamespaceScopeSidecarConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigRequest) GoString() string {
	return s.String()
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

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetIstioDNSProxyEnabled(v bool) *UpdateNamespaceScopeSidecarConfigRequest {
	s.IstioDNSProxyEnabled = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetLifecycle(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.Lifecycle = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetNamespace(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.Namespace = &v
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

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetServiceMeshId(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ServiceMeshId = &v
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

type UpdateNamespaceScopeSidecarConfigResponseBody struct {
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
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateNamespaceScopeSidecarConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateNamespaceScopeSidecarConfigResponse) SetBody(v *UpdateNamespaceScopeSidecarConfigResponseBody) *UpdateNamespaceScopeSidecarConfigResponse {
	s.Body = v
	return s
}

type UpgradeMeshEditionPartiallyRequest struct {
	ASMGatewayContinue    *bool   `json:"ASMGatewayContinue,omitempty" xml:"ASMGatewayContinue,omitempty"`
	ServiceMeshId         *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	SwitchToPro           *bool   `json:"SwitchToPro,omitempty" xml:"SwitchToPro,omitempty"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeMeshEditionPartiallyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeMeshEditionPartiallyResponse) SetBody(v *UpgradeMeshEditionPartiallyResponseBody) *UpgradeMeshEditionPartiallyResponse {
	s.Body = v
	return s
}

type UpgradeMeshVersionRequest struct {
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpgradeMeshVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMeshVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMeshVersionRequest) SetServiceMeshId(v string) *UpgradeMeshVersionRequest {
	s.ServiceMeshId = &v
	return s
}

type UpgradeMeshVersionResponseBody struct {
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeMeshVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeMeshVersionResponse) SetBody(v *UpgradeMeshVersionResponseBody) *UpgradeMeshVersionResponse {
	s.Body = v
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

type ClusterStatusLogtailStatusRecordValue struct {
	ClusterId           *string                                                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	LogtailDetailStatus *string                                                     `json:"LogtailDetailStatus,omitempty" xml:"LogtailDetailStatus,omitempty"`
	AccessLogDashboards []*ClusterStatusLogtailStatusRecordValueAccessLogDashboards `json:"AccessLogDashboards,omitempty" xml:"AccessLogDashboards,omitempty" type:"Repeated"`
}

func (s ClusterStatusLogtailStatusRecordValue) String() string {
	return tea.Prettify(s)
}

func (s ClusterStatusLogtailStatusRecordValue) GoString() string {
	return s.String()
}

func (s *ClusterStatusLogtailStatusRecordValue) SetClusterId(v string) *ClusterStatusLogtailStatusRecordValue {
	s.ClusterId = &v
	return s
}

func (s *ClusterStatusLogtailStatusRecordValue) SetLogtailDetailStatus(v string) *ClusterStatusLogtailStatusRecordValue {
	s.LogtailDetailStatus = &v
	return s
}

func (s *ClusterStatusLogtailStatusRecordValue) SetAccessLogDashboards(v []*ClusterStatusLogtailStatusRecordValueAccessLogDashboards) *ClusterStatusLogtailStatusRecordValue {
	s.AccessLogDashboards = v
	return s
}

type ClusterStatusLogtailStatusRecordValueAccessLogDashboards struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ClusterStatusLogtailStatusRecordValueAccessLogDashboards) String() string {
	return tea.Prettify(s)
}

func (s ClusterStatusLogtailStatusRecordValueAccessLogDashboards) GoString() string {
	return s.String()
}

func (s *ClusterStatusLogtailStatusRecordValueAccessLogDashboards) SetTitle(v string) *ClusterStatusLogtailStatusRecordValueAccessLogDashboards {
	s.Title = &v
	return s
}

func (s *ClusterStatusLogtailStatusRecordValueAccessLogDashboards) SetUrl(v string) *ClusterStatusLogtailStatusRecordValueAccessLogDashboards {
	s.Url = &v
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

type ClusterStatusLogtailStatusRecordValueValue struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ClusterStatusLogtailStatusRecordValueValue) String() string {
	return tea.Prettify(s)
}

func (s ClusterStatusLogtailStatusRecordValueValue) GoString() string {
	return s.String()
}

func (s *ClusterStatusLogtailStatusRecordValueValue) SetTitle(v string) *ClusterStatusLogtailStatusRecordValueValue {
	s.Title = &v
	return s
}

func (s *ClusterStatusLogtailStatusRecordValueValue) SetUrl(v string) *ClusterStatusLogtailStatusRecordValueValue {
	s.Url = &v
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

func (client *Client) CreateServiceMeshWithOptions(request *CreateServiceMeshRequest, runtime *util.RuntimeOptions) (_result *CreateServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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

	if !tea.BoolValue(util.IsUnset(request.FilterGatewayClusterConfig)) {
		body["FilterGatewayClusterConfig"] = request.FilterGatewayClusterConfig
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayAPIEnabled)) {
		body["GatewayAPIEnabled"] = request.GatewayAPIEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalRateLimitEnabled)) {
		body["GlobalRateLimitEnabled"] = request.GlobalRateLimitEnabled
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
		Body: openapiutil.ParseToMap(body),
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

func (client *Client) DescribeAhasProWithOptions(request *DescribeAhasProRequest, runtime *util.RuntimeOptions) (_result *DescribeAhasProResponse, _err error) {
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
		Action:      tea.String("DescribeAhasPro"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAhasProResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAhasPro(request *DescribeAhasProRequest) (_result *DescribeAhasProResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAhasProResponse{}
	_body, _err := client.DescribeAhasProWithOptions(request, runtime)
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

func (client *Client) DescribeIngressGatewaysWithOptions(request *DescribeIngressGatewaysRequest, runtime *util.RuntimeOptions) (_result *DescribeIngressGatewaysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIngressGateways"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIngressGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIngressGateways(request *DescribeIngressGatewaysRequest) (_result *DescribeIngressGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIngressGatewaysResponse{}
	_body, _err := client.DescribeIngressGatewaysWithOptions(request, runtime)
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
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

func (client *Client) DescribeServiceMeshesWithOptions(runtime *util.RuntimeOptions) (_result *DescribeServiceMeshesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
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

func (client *Client) DescribeServiceMeshes() (_result *DescribeServiceMeshesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceMeshesResponse{}
	_body, _err := client.DescribeServiceMeshesWithOptions(runtime)
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

func (client *Client) RunDiagnosisWithOptions(request *RunDiagnosisRequest, runtime *util.RuntimeOptions) (_result *RunDiagnosisResponse, _err error) {
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
		Action:      tea.String("RunDiagnosis"),
		Version:     tea.String("2020-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunDiagnosis(request *RunDiagnosisRequest) (_result *RunDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunDiagnosisResponse{}
	_body, _err := client.RunDiagnosisWithOptions(request, runtime)
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

func (client *Client) UpdateMeshFeatureWithOptions(request *UpdateMeshFeatureRequest, runtime *util.RuntimeOptions) (_result *UpdateMeshFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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

	if !tea.BoolValue(util.IsUnset(request.DiscoverySelectors)) {
		body["DiscoverySelectors"] = request.DiscoverySelectors
	}

	if !tea.BoolValue(util.IsUnset(request.DubboFilterEnabled)) {
		body["DubboFilterEnabled"] = request.DubboFilterEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAudit)) {
		body["EnableAudit"] = request.EnableAudit
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

	if !tea.BoolValue(util.IsUnset(request.GlobalRateLimitEnabled)) {
		body["GlobalRateLimitEnabled"] = request.GlobalRateLimitEnabled
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

	if !tea.BoolValue(util.IsUnset(request.KialiEnabled)) {
		body["KialiEnabled"] = request.KialiEnabled
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

	if !tea.BoolValue(util.IsUnset(request.TraceSampling)) {
		body["TraceSampling"] = request.TraceSampling
	}

	if !tea.BoolValue(util.IsUnset(request.Tracing)) {
		body["Tracing"] = request.Tracing
	}

	if !tea.BoolValue(util.IsUnset(request.WebAssemblyFilterEnabled)) {
		body["WebAssemblyFilterEnabled"] = request.WebAssemblyFilterEnabled
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
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
	if !tea.BoolValue(util.IsUnset(request.ExcludeIPRanges)) {
		body["ExcludeIPRanges"] = request.ExcludeIPRanges
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeInboundPorts)) {
		body["ExcludeInboundPorts"] = request.ExcludeInboundPorts
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeOutboundPorts)) {
		body["ExcludeOutboundPorts"] = request.ExcludeOutboundPorts
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

	if !tea.BoolValue(util.IsUnset(request.IstioDNSProxyEnabled)) {
		body["IstioDNSProxyEnabled"] = request.IstioDNSProxyEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Lifecycle)) {
		body["Lifecycle"] = request.Lifecycle
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
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

	if !tea.BoolValue(util.IsUnset(request.ServiceMeshId)) {
		body["ServiceMeshId"] = request.ServiceMeshId
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

func (client *Client) UpgradeMeshEditionPartiallyWithOptions(request *UpgradeMeshEditionPartiallyRequest, runtime *util.RuntimeOptions) (_result *UpgradeMeshEditionPartiallyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ASMGatewayContinue)) {
		body["ASMGatewayContinue"] = request.ASMGatewayContinue
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
