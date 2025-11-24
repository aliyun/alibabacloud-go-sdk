// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceMeshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLogEnabled(v bool) *CreateServiceMeshRequest
	GetAccessLogEnabled() *bool
	SetAccessLogFile(v string) *CreateServiceMeshRequest
	GetAccessLogFile() *string
	SetAccessLogFormat(v string) *CreateServiceMeshRequest
	GetAccessLogFormat() *string
	SetAccessLogProject(v string) *CreateServiceMeshRequest
	GetAccessLogProject() *string
	SetAccessLogServiceEnabled(v bool) *CreateServiceMeshRequest
	GetAccessLogServiceEnabled() *bool
	SetAccessLogServiceHost(v string) *CreateServiceMeshRequest
	GetAccessLogServiceHost() *string
	SetAccessLogServicePort(v int32) *CreateServiceMeshRequest
	GetAccessLogServicePort() *int32
	SetApiServerLoadBalancerSpec(v string) *CreateServiceMeshRequest
	GetApiServerLoadBalancerSpec() *string
	SetApiServerPublicEip(v bool) *CreateServiceMeshRequest
	GetApiServerPublicEip() *bool
	SetAuditProject(v string) *CreateServiceMeshRequest
	GetAuditProject() *string
	SetAutoRenew(v bool) *CreateServiceMeshRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateServiceMeshRequest
	GetAutoRenewPeriod() *int32
	SetCRAggregationEnabled(v bool) *CreateServiceMeshRequest
	GetCRAggregationEnabled() *bool
	SetCertChain(v string) *CreateServiceMeshRequest
	GetCertChain() *string
	SetChargeType(v string) *CreateServiceMeshRequest
	GetChargeType() *string
	SetClusterDomain(v string) *CreateServiceMeshRequest
	GetClusterDomain() *string
	SetClusterSpec(v string) *CreateServiceMeshRequest
	GetClusterSpec() *string
	SetConfigSourceEnabled(v bool) *CreateServiceMeshRequest
	GetConfigSourceEnabled() *bool
	SetConfigSourceNacosID(v string) *CreateServiceMeshRequest
	GetConfigSourceNacosID() *string
	SetControlPlaneLogEnabled(v bool) *CreateServiceMeshRequest
	GetControlPlaneLogEnabled() *bool
	SetControlPlaneLogProject(v string) *CreateServiceMeshRequest
	GetControlPlaneLogProject() *string
	SetCustomizedPrometheus(v bool) *CreateServiceMeshRequest
	GetCustomizedPrometheus() *bool
	SetCustomizedZipkin(v bool) *CreateServiceMeshRequest
	GetCustomizedZipkin() *bool
	SetDNSProxyingEnabled(v bool) *CreateServiceMeshRequest
	GetDNSProxyingEnabled() *bool
	SetDubboFilterEnabled(v bool) *CreateServiceMeshRequest
	GetDubboFilterEnabled() *bool
	SetEdition(v string) *CreateServiceMeshRequest
	GetEdition() *string
	SetEnableACMG(v bool) *CreateServiceMeshRequest
	GetEnableACMG() *bool
	SetEnableAmbient(v bool) *CreateServiceMeshRequest
	GetEnableAmbient() *bool
	SetEnableAudit(v bool) *CreateServiceMeshRequest
	GetEnableAudit() *bool
	SetEnableCRHistory(v bool) *CreateServiceMeshRequest
	GetEnableCRHistory() *bool
	SetEnableSDSServer(v bool) *CreateServiceMeshRequest
	GetEnableSDSServer() *bool
	SetExcludeIPRanges(v string) *CreateServiceMeshRequest
	GetExcludeIPRanges() *string
	SetExcludeInboundPorts(v string) *CreateServiceMeshRequest
	GetExcludeInboundPorts() *string
	SetExcludeOutboundPorts(v string) *CreateServiceMeshRequest
	GetExcludeOutboundPorts() *string
	SetExistingCaCert(v string) *CreateServiceMeshRequest
	GetExistingCaCert() *string
	SetExistingCaKey(v string) *CreateServiceMeshRequest
	GetExistingCaKey() *string
	SetExistingCaType(v string) *CreateServiceMeshRequest
	GetExistingCaType() *string
	SetExistingRootCaCert(v string) *CreateServiceMeshRequest
	GetExistingRootCaCert() *string
	SetExistingRootCaKey(v string) *CreateServiceMeshRequest
	GetExistingRootCaKey() *string
	SetFilterGatewayClusterConfig(v bool) *CreateServiceMeshRequest
	GetFilterGatewayClusterConfig() *bool
	SetGatewayAPIEnabled(v bool) *CreateServiceMeshRequest
	GetGatewayAPIEnabled() *bool
	SetGuestCluster(v string) *CreateServiceMeshRequest
	GetGuestCluster() *string
	SetIncludeIPRanges(v string) *CreateServiceMeshRequest
	GetIncludeIPRanges() *string
	SetIstioVersion(v string) *CreateServiceMeshRequest
	GetIstioVersion() *string
	SetKialiEnabled(v bool) *CreateServiceMeshRequest
	GetKialiEnabled() *bool
	SetLocalityLBConf(v string) *CreateServiceMeshRequest
	GetLocalityLBConf() *string
	SetLocalityLoadBalancing(v bool) *CreateServiceMeshRequest
	GetLocalityLoadBalancing() *bool
	SetMSEEnabled(v bool) *CreateServiceMeshRequest
	GetMSEEnabled() *bool
	SetMultiBufferEnabled(v bool) *CreateServiceMeshRequest
	GetMultiBufferEnabled() *bool
	SetMultiBufferPollDelay(v string) *CreateServiceMeshRequest
	GetMultiBufferPollDelay() *string
	SetMysqlFilterEnabled(v bool) *CreateServiceMeshRequest
	GetMysqlFilterEnabled() *bool
	SetName(v string) *CreateServiceMeshRequest
	GetName() *string
	SetOPALimitCPU(v string) *CreateServiceMeshRequest
	GetOPALimitCPU() *string
	SetOPALimitMemory(v string) *CreateServiceMeshRequest
	GetOPALimitMemory() *string
	SetOPALogLevel(v string) *CreateServiceMeshRequest
	GetOPALogLevel() *string
	SetOPARequestCPU(v string) *CreateServiceMeshRequest
	GetOPARequestCPU() *string
	SetOPARequestMemory(v string) *CreateServiceMeshRequest
	GetOPARequestMemory() *string
	SetOpaEnabled(v bool) *CreateServiceMeshRequest
	GetOpaEnabled() *bool
	SetOpenAgentPolicy(v bool) *CreateServiceMeshRequest
	GetOpenAgentPolicy() *bool
	SetPeriod(v int32) *CreateServiceMeshRequest
	GetPeriod() *int32
	SetPilotLoadBalancerSpec(v string) *CreateServiceMeshRequest
	GetPilotLoadBalancerSpec() *string
	SetPlaygroundScene(v string) *CreateServiceMeshRequest
	GetPlaygroundScene() *string
	SetPrometheusUrl(v string) *CreateServiceMeshRequest
	GetPrometheusUrl() *string
	SetProxyLimitCPU(v string) *CreateServiceMeshRequest
	GetProxyLimitCPU() *string
	SetProxyLimitMemory(v string) *CreateServiceMeshRequest
	GetProxyLimitMemory() *string
	SetProxyRequestCPU(v string) *CreateServiceMeshRequest
	GetProxyRequestCPU() *string
	SetProxyRequestMemory(v string) *CreateServiceMeshRequest
	GetProxyRequestMemory() *string
	SetRedisFilterEnabled(v bool) *CreateServiceMeshRequest
	GetRedisFilterEnabled() *bool
	SetRegionId(v string) *CreateServiceMeshRequest
	GetRegionId() *string
	SetTag(v []*CreateServiceMeshRequestTag) *CreateServiceMeshRequest
	GetTag() []*CreateServiceMeshRequestTag
	SetTelemetry(v bool) *CreateServiceMeshRequest
	GetTelemetry() *bool
	SetThriftFilterEnabled(v bool) *CreateServiceMeshRequest
	GetThriftFilterEnabled() *bool
	SetTraceSampling(v float32) *CreateServiceMeshRequest
	GetTraceSampling() *float32
	SetTracing(v bool) *CreateServiceMeshRequest
	GetTracing() *bool
	SetUseExistingCA(v bool) *CreateServiceMeshRequest
	GetUseExistingCA() *bool
	SetVSwitches(v string) *CreateServiceMeshRequest
	GetVSwitches() *string
	SetVpcId(v string) *CreateServiceMeshRequest
	GetVpcId() *string
	SetWebAssemblyFilterEnabled(v bool) *CreateServiceMeshRequest
	GetWebAssemblyFilterEnabled() *bool
}

type CreateServiceMeshRequest struct {
	// Specifies whether to enable access log collection. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AccessLogEnabled *bool `json:"AccessLogEnabled,omitempty" xml:"AccessLogEnabled,omitempty"`
	// Specifies whether to enable access log collection. Valid values:
	//
	// 	- "": disables access log collection.
	//
	// 	- `/dev/stdout`: enables access log collection. Access logs are written to /dev/stdout.
	//
	// example:
	//
	// /dev/stdout
	AccessLogFile *string `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	// Custom fields of access logs. To set this parameter, you must enable access log collection. Otherwise, you cannot set this parameter. The value must be a JSON string. The following key values must be contained: authority_for, bytes_received, bytes_sent, downstream_local_address, downstream_remote_address, duration, istio_policy_status, method, path, protocol, requested_server_name, response_code, response_flags, route_name, start_time, trace_id, upstream_cluster, upstream_host, upstream_local_address, upstream_service_time, upstream_transport_failure_reason, user_agent, and x_forwarded_for.
	//
	// example:
	//
	// {"authority_for":"%REQ(:AUTHORITY)%","bytes_received":"%BYTES_RECEIVED%","bytes_sent":"%BYTES_SENT%","downstream_local_address":"%DOWNSTREAM_LOCAL_ADDRESS%","downstream_remote_address":"%DOWNSTREAM_REMOTE_ADDRESS%","duration":"%DURATION%","istio_policy_status":"%DYNAMIC_METADATA(istio.mixer:status)%","method":"%REQ(:METHOD)%","path":"%REQ(X-ENVOY-ORIGINAL-PATH?:PATH)%","protocol":"%PROTOCOL%","request_id":"%REQ(X-REQUEST-ID)%","requested_server_name":"%REQUESTED_SERVER_NAME%","response_code":"%RESPONSE_CODE%","response_flags":"%RESPONSE_FLAGS%","route_name":"%ROUTE_NAME%","start_time":"%START_TIME%","trace_id":"%REQ(X-B3-TRACEID)%","upstream_cluster":"%UPSTREAM_CLUSTER%","upstream_host":"%UPSTREAM_HOST%","upstream_local_address":"%UPSTREAM_LOCAL_ADDRESS%","upstream_service_time":"%RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)%","upstream_transport_failure_reason":"%UPSTREAM_TRANSPORT_FAILURE_REASON%","user_agent":"%REQ(USER-AGENT)%","x_forwarded_for":"%REQ(X-FORWARDED-FOR)%"}
	AccessLogFormat *string `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	// The SLS project from which access logs are collected.
	//
	// example:
	//
	// mesh-log-cf245a429b6ff4b6e97f20797758*****
	AccessLogProject *string `json:"AccessLogProject,omitempty" xml:"AccessLogProject,omitempty"`
	// Specifies whether to enable gRPC Access Log Service (ALS) of Envoy. gRPC is short for Google Remote Procedure Call. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AccessLogServiceEnabled *bool `json:"AccessLogServiceEnabled,omitempty" xml:"AccessLogServiceEnabled,omitempty"`
	// The endpoint of Envoy\\"s gRPC ALS.
	//
	// example:
	//
	// 0.0.0.0
	AccessLogServiceHost *string `json:"AccessLogServiceHost,omitempty" xml:"AccessLogServiceHost,omitempty"`
	// The port of Envoy\\"s gRPC ALS.
	//
	// example:
	//
	// 9999
	AccessLogServicePort *int32 `json:"AccessLogServicePort,omitempty" xml:"AccessLogServicePort,omitempty"`
	// The type of the Classic Load Balancer (CLB) instance that is bound to the API server. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`, `slb.s3.small`, `slb.s3.medium`, and `slb.s3.large`.
	//
	// example:
	//
	// slb.s1.small
	ApiServerLoadBalancerSpec *string `json:"ApiServerLoadBalancerSpec,omitempty" xml:"ApiServerLoadBalancerSpec,omitempty"`
	// Specifies whether to expose the API server to the Internet. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// > If you set this parameter to `false`, the API server cannot be accessed over the Internet.
	//
	// example:
	//
	// false
	ApiServerPublicEip *bool `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	// The name of the Log Service project that is used for mesh audit.
	//
	// Default value: mesh-log-{ASM instance ID}.
	//
	// example:
	//
	// mesh-log-xxxx
	AuditProject *string `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	// Specifies whether to enable auto-renewal for the CLB instance if the CLB instance uses the subscription billing method. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period of the subscription CLB instance. This parameter is valid only if `ChargeType` is set to `PrePay`. If the original subscription period of the CLB instance is less than one year, the value of this parameter indicates the number of months for auto-renewal. If the original subscription period of the CLB instance is more than one year, the value of this parameter indicates the number of years for auto-renewal.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to allow the Kubernetes API of clusters on the data plane to access Istio resources. The version of the ASM instance must be V1.9.7.93 or later. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CRAggregationEnabled *bool `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	// The certificate chain from the CA certificate to the root certificate. At least two certificates are included in the chain.
	//
	// example:
	//
	// Base64 encoded PEM certificate chain.
	CertChain *string `json:"CertChain,omitempty" xml:"CertChain,omitempty"`
	// The billing method of the CLB instance. Valid values:
	//
	// 	- `PayOnDemand`: pay-as-you-go
	//
	// 	- `PrePay`: subscription
	//
	// example:
	//
	// PrePay
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// ASM cluster domain.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The edition of the ASM instance. Valid values:
	//
	// - `standard`: Standard Edition
	//
	// - `enterprise`: Enterprise Edition
	//
	// - `ultimate`: Ultimate Edition
	//
	// example:
	//
	// standard
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// Specifies whether to enable the external service registry. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	ConfigSourceEnabled *bool `json:"ConfigSourceEnabled,omitempty" xml:"ConfigSourceEnabled,omitempty"`
	// The instance ID of the Nacos registry.
	//
	// example:
	//
	// mse-cn-tl326******
	ConfigSourceNacosID *string `json:"ConfigSourceNacosID,omitempty" xml:"ConfigSourceNacosID,omitempty"`
	// Specifies whether to enable the collection of control plane logs. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	ControlPlaneLogEnabled *bool `json:"ControlPlaneLogEnabled,omitempty" xml:"ControlPlaneLogEnabled,omitempty"`
	// The name of the Log Service project that is used to collect the logs of the control plane.
	//
	// example:
	//
	// mesh-log-cf245a429b6ff4b6e97f20797758*****
	ControlPlaneLogProject *string `json:"ControlPlaneLogProject,omitempty" xml:"ControlPlaneLogProject,omitempty"`
	// Specifies whether to use a custom Prometheus instance. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CustomizedPrometheus *bool `json:"CustomizedPrometheus,omitempty" xml:"CustomizedPrometheus,omitempty"`
	// Specifies whether to use a self-managed Zipkin system to collect tracing data. Valid values:
	//
	// 	- `true`: uses a self-managed Zipkin system to collect tracing data.
	//
	// 	- `false`: uses Alibaba Cloud Tracing Analysis to collect tracing data.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CustomizedZipkin *bool `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	// Specifies whether to enable the DNS proxy feature. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DNSProxyingEnabled *bool `json:"DNSProxyingEnabled,omitempty" xml:"DNSProxyingEnabled,omitempty"`
	// Specifies whether to enable Dubbo Filter. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DubboFilterEnabled *bool `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	// The edition of the ASM instance.
	//
	// example:
	//
	// Pro
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// Specifies whether to enable the ACMG mode for the ASM instance.
	//
	// example:
	//
	// false
	EnableACMG *bool `json:"EnableACMG,omitempty" xml:"EnableACMG,omitempty"`
	// Specifies whether to enable the Ambient Mesh mode for the ASM instance.
	//
	// example:
	//
	// false
	EnableAmbient *bool `json:"EnableAmbient,omitempty" xml:"EnableAmbient,omitempty"`
	// Specifies whether to enable the mesh audit feature. To enable this feature, make sure that you have activated [Log Service](https://sls.console.aliyun.com/). Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	EnableAudit *bool `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	// Specifies whether to enable the rollback feature for Istio resources. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	EnableCRHistory *bool `json:"EnableCRHistory,omitempty" xml:"EnableCRHistory,omitempty"`
	// Specifies whether to enable Secret Discovery Service (SDS). Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	EnableSDSServer *bool `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
	// The IP ranges in CIDR form to be excluded from redirection to the sidecar proxy in the ASM instance.
	//
	// example:
	//
	// 100.100.10*.***
	ExcludeIPRanges *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	// The inbound ports to be excluded from redirection to the sidecar proxy in the ASM instance. Separate multiple port numbers with commas (,).
	//
	// example:
	//
	// 80,81
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The outbound ports to be excluded from redirection to the sidecar proxy in the ASM instance. Separate multiple port numbers with commas (,).
	//
	// example:
	//
	// 80,81
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// The existing CA certificate, which is encoded in Base64. This parameter is used in scenarios where you migrate open source Istio to ASM. It specifies the content of the ca-cert.pem file in the istio-ca-secret secret. The secret is in the istio-system namespace of the Kubernetes cluster where the open source Istio is installed.
	//
	// example:
	//
	// CA cert content, base64 encoded format.
	ExistingCaCert *string `json:"ExistingCaCert,omitempty" xml:"ExistingCaCert,omitempty"`
	// The existing CA key, which is encoded in Base64. This parameter is used in scenarios where you migrate open source Istio to ASM. It specifies the content of the ca-key.pem file in the istio-ca-secret secret. The secret is in the istio-system namespace of the Kubernetes cluster where the open source Istio is installed.
	//
	// example:
	//
	// CA key content, base64 encoded format.
	ExistingCaKey *string `json:"ExistingCaKey,omitempty" xml:"ExistingCaKey,omitempty"`
	// Deprecated
	//
	// The type of the existing CA certificate. Valid values:
	//
	// 	- 1: Self-signed certificate generated by istiod. The certificate corresponds to the secret named istio-ca-secret in the istio-system namespace. If you use this type of certificate, you must set the `ExistingCaCert` and `ExistingCaKey` parameters.
	//
	// 	- 2: Administrator-specified certificate. For more information, see [plugin ca cert](https://istio.io/latest/docs/tasks/security/cert-management/plugin-ca-cert/). In most cases, the certificate corresponds to the secret named cacerts in the istio-system namespace. If you use this type of certificate, you must set the `ExisingRootCaCert` and `ExisingRootCaKey` parameters.
	//
	// example:
	//
	// 1
	ExistingCaType *string `json:"ExistingCaType,omitempty" xml:"ExistingCaType,omitempty"`
	// The existing root certificate, which is encoded in Base64.
	//
	// example:
	//
	// Existing CA cert content, base64 encoded format.
	ExistingRootCaCert *string `json:"ExistingRootCaCert,omitempty" xml:"ExistingRootCaCert,omitempty"`
	// Deprecated
	//
	// The private key that corresponds to the root certificate, which is encoded in Base64.
	//
	// example:
	//
	// Existing CA key content, base64 encoded format.
	ExistingRootCaKey *string `json:"ExistingRootCaKey,omitempty" xml:"ExistingRootCaKey,omitempty"`
	// Specifies whether to enable gateway configuration filtering. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	FilterGatewayClusterConfig *bool `json:"FilterGatewayClusterConfig,omitempty" xml:"FilterGatewayClusterConfig,omitempty"`
	// Specifies whether to enable Gateway API. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	GatewayAPIEnabled *bool `json:"GatewayAPIEnabled,omitempty" xml:"GatewayAPIEnabled,omitempty"`
	// When you create an ASM instance, you can add a cluster to the ASM instance. If you do not specify this parameter, no cluster is added to the ASM instance. The cluster and the ASM instance must be in the same vSwitch of the same VPC and have the same domain name.
	//
	// example:
	//
	// ACK cluster id
	GuestCluster *string `json:"GuestCluster,omitempty" xml:"GuestCluster,omitempty"`
	// The IP ranges in CIDR form for which traffic is to be redirected to the sidecar proxy in the ASM instance.
	//
	// example:
	//
	// *
	IncludeIPRanges *string `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	// The Istio version of the ASM instance.
	//
	// example:
	//
	// v1.5.4.1-g5960ec40-aliyun
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	// Specifies whether to enable the mesh topology feature. To enable this feature, make sure that you have enabled Prometheus monitoring. If Prometheus monitoring is disabled, you must set this parameter to `false`.`` Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	KialiEnabled *bool `json:"KialiEnabled,omitempty" xml:"KialiEnabled,omitempty"`
	// The configurations for the access to the nearest instance.
	//
	// example:
	//
	// {"failover":[{"from":"cn-hangzhou","to":"cn-shanghai"}]}
	LocalityLBConf *string `json:"LocalityLBConf,omitempty" xml:"LocalityLBConf,omitempty"`
	// Specifies whether to route traffic to the nearest instance. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	LocalityLoadBalancing *bool `json:"LocalityLoadBalancing,omitempty" xml:"LocalityLoadBalancing,omitempty"`
	// Specifies whether to enable Microservices Engine (MSE). Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	MSEEnabled *bool `json:"MSEEnabled,omitempty" xml:"MSEEnabled,omitempty"`
	// Specifies whether to enable MultiBuffer-based Transport Layer Security (TLS) acceleration. Valid values:
	//
	// - `true`
	//
	// - `false`
	//
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	MultiBufferEnabled *bool `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	// The pull-request latency. Default value: 30. Unit: seconds.
	//
	// example:
	//
	// 30s
	MultiBufferPollDelay *string `json:"MultiBufferPollDelay,omitempty" xml:"MultiBufferPollDelay,omitempty"`
	// Specifies whether to enable MySQL Filter. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	MysqlFilterEnabled *bool `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	// The name of the ASM instance.
	//
	// example:
	//
	// mesh1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of CPU cores that are available to the OPA container.
	//
	// example:
	//
	// 2
	OPALimitCPU *string `json:"OPALimitCPU,omitempty" xml:"OPALimitCPU,omitempty"`
	// The maximum size of the memory that is available to the OPA container. You can specify the parameter value in the standard quantity representation form used by Kubernetes. 1 Mi equals 1,024 KB.
	//
	// example:
	//
	// 1024Mi
	OPALimitMemory *string `json:"OPALimitMemory,omitempty" xml:"OPALimitMemory,omitempty"`
	// The log level of the OPA container.
	//
	// example:
	//
	// info
	OPALogLevel *string `json:"OPALogLevel,omitempty" xml:"OPALogLevel,omitempty"`
	// The minimum number of CPU cores that are required by the OPA container. You can specify the parameter value in the standard representation form of CPUs in Kubernetes. For example, if you set the value to 1, one CPU core is required.
	//
	// example:
	//
	// 1
	OPARequestCPU *string `json:"OPARequestCPU,omitempty" xml:"OPARequestCPU,omitempty"`
	// The minimum size of the memory that is required by the OPA container. You can specify the parameter value in the standard quantity representation form used by Kubernetes. 1 Mi equals 1,024 KB.
	//
	// example:
	//
	// 512Mi
	OPARequestMemory *string `json:"OPARequestMemory,omitempty" xml:"OPARequestMemory,omitempty"`
	// Specifies whether to enable the OPA plug-in. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	OpaEnabled *bool `json:"OpaEnabled,omitempty" xml:"OpaEnabled,omitempty"`
	// Specifies whether to install the Open Policy Agent (OPA) plug-in. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	OpenAgentPolicy *bool `json:"OpenAgentPolicy,omitempty" xml:"OpenAgentPolicy,omitempty"`
	// The subscription period of the CLB instance. This parameter is valid only if `ChargeType` is set to `PrePay`. The value of this parameter indicates the subscription period of the CLB instance. Unit: month. For example, if the subscription period is one year, set this parameter to 12.
	//
	// example:
	//
	// 3
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The type of the CLB instance that is bound to Istio Pilot. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`, `slb.s3.small`, `slb.s3.medium`, and `slb.s3.large`.
	//
	// example:
	//
	// slb.s1.small
	PilotLoadBalancerSpec *string `json:"PilotLoadBalancerSpec,omitempty" xml:"PilotLoadBalancerSpec,omitempty"`
	// The playground scenario. If you specify this parameter, an ASM playground instance is created. Valid values:
	//
	// 	- ewmaLb: the exponentially weighted moving average (EWMA) load balancing algorithm.
	//
	// example:
	//
	// ewmaLb
	PlaygroundScene *string `json:"PlaygroundScene,omitempty" xml:"PlaygroundScene,omitempty"`
	// The endpoint of the custom Prometheus instance.
	//
	// example:
	//
	// http://prometheus:9090
	PrometheusUrl *string `json:"PrometheusUrl,omitempty" xml:"PrometheusUrl,omitempty"`
	// The maximum number of CPU cores that are available to the proxy container.
	//
	// example:
	//
	// 2000m
	ProxyLimitCPU *string `json:"ProxyLimitCPU,omitempty" xml:"ProxyLimitCPU,omitempty"`
	// The maximum size of the memory that is available to the proxy container.
	//
	// example:
	//
	// 1024Mi
	ProxyLimitMemory *string `json:"ProxyLimitMemory,omitempty" xml:"ProxyLimitMemory,omitempty"`
	// The minimum number of CPU cores that are required by the proxy container.
	//
	// example:
	//
	// 100m
	ProxyRequestCPU *string `json:"ProxyRequestCPU,omitempty" xml:"ProxyRequestCPU,omitempty"`
	// The minimum size of the memory that is required by the proxy container.
	//
	// example:
	//
	// 128Mi
	ProxyRequestMemory *string `json:"ProxyRequestMemory,omitempty" xml:"ProxyRequestMemory,omitempty"`
	// Specifies whether to enable Redis Filter. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	RedisFilterEnabled *bool `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
	// The ID of the region in which the ASM instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Tag of the ASM instance. A tag contains the following information:
	//
	// 	- key: the name of the tag
	//
	// 	- value: the value of the tag
	Tag []*CreateServiceMeshRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to enable Prometheus monitoring. We recommend that you use Prometheus Service of [Application Real-Time Monitoring Service (ARMS)](https://arms.console.aliyun.com/). Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Telemetry *bool `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	// Specifies whether to enable Thrift Filter. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	ThriftFilterEnabled *bool `json:"ThriftFilterEnabled,omitempty" xml:"ThriftFilterEnabled,omitempty"`
	// The sampling percentage of Tracing Analysis.
	//
	// example:
	//
	// 100
	TraceSampling *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
	// Specifies whether to enable the Tracing Analysis feature. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Tracing *bool `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
	// Specifies whether to use an existing CA certificate and private key.
	//
	// example:
	//
	// false
	UseExistingCA *bool `json:"UseExistingCA,omitempty" xml:"UseExistingCA,omitempty"`
	// The ID of the vSwitch to which the ASM instance is connected.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["vsw-xzegf5dndkbf4m6eg****"]
	VSwitches *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the ASM instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xzelac2tw4ic7wz31****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Specifies whether to enable WebAssembly Filter. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	WebAssemblyFilterEnabled *bool `json:"WebAssemblyFilterEnabled,omitempty" xml:"WebAssemblyFilterEnabled,omitempty"`
}

func (s CreateServiceMeshRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshRequest) GetAccessLogEnabled() *bool {
	return s.AccessLogEnabled
}

func (s *CreateServiceMeshRequest) GetAccessLogFile() *string {
	return s.AccessLogFile
}

func (s *CreateServiceMeshRequest) GetAccessLogFormat() *string {
	return s.AccessLogFormat
}

func (s *CreateServiceMeshRequest) GetAccessLogProject() *string {
	return s.AccessLogProject
}

func (s *CreateServiceMeshRequest) GetAccessLogServiceEnabled() *bool {
	return s.AccessLogServiceEnabled
}

func (s *CreateServiceMeshRequest) GetAccessLogServiceHost() *string {
	return s.AccessLogServiceHost
}

func (s *CreateServiceMeshRequest) GetAccessLogServicePort() *int32 {
	return s.AccessLogServicePort
}

func (s *CreateServiceMeshRequest) GetApiServerLoadBalancerSpec() *string {
	return s.ApiServerLoadBalancerSpec
}

func (s *CreateServiceMeshRequest) GetApiServerPublicEip() *bool {
	return s.ApiServerPublicEip
}

func (s *CreateServiceMeshRequest) GetAuditProject() *string {
	return s.AuditProject
}

func (s *CreateServiceMeshRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateServiceMeshRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateServiceMeshRequest) GetCRAggregationEnabled() *bool {
	return s.CRAggregationEnabled
}

func (s *CreateServiceMeshRequest) GetCertChain() *string {
	return s.CertChain
}

func (s *CreateServiceMeshRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateServiceMeshRequest) GetClusterDomain() *string {
	return s.ClusterDomain
}

func (s *CreateServiceMeshRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *CreateServiceMeshRequest) GetConfigSourceEnabled() *bool {
	return s.ConfigSourceEnabled
}

func (s *CreateServiceMeshRequest) GetConfigSourceNacosID() *string {
	return s.ConfigSourceNacosID
}

func (s *CreateServiceMeshRequest) GetControlPlaneLogEnabled() *bool {
	return s.ControlPlaneLogEnabled
}

func (s *CreateServiceMeshRequest) GetControlPlaneLogProject() *string {
	return s.ControlPlaneLogProject
}

func (s *CreateServiceMeshRequest) GetCustomizedPrometheus() *bool {
	return s.CustomizedPrometheus
}

func (s *CreateServiceMeshRequest) GetCustomizedZipkin() *bool {
	return s.CustomizedZipkin
}

func (s *CreateServiceMeshRequest) GetDNSProxyingEnabled() *bool {
	return s.DNSProxyingEnabled
}

func (s *CreateServiceMeshRequest) GetDubboFilterEnabled() *bool {
	return s.DubboFilterEnabled
}

func (s *CreateServiceMeshRequest) GetEdition() *string {
	return s.Edition
}

func (s *CreateServiceMeshRequest) GetEnableACMG() *bool {
	return s.EnableACMG
}

func (s *CreateServiceMeshRequest) GetEnableAmbient() *bool {
	return s.EnableAmbient
}

func (s *CreateServiceMeshRequest) GetEnableAudit() *bool {
	return s.EnableAudit
}

func (s *CreateServiceMeshRequest) GetEnableCRHistory() *bool {
	return s.EnableCRHistory
}

func (s *CreateServiceMeshRequest) GetEnableSDSServer() *bool {
	return s.EnableSDSServer
}

func (s *CreateServiceMeshRequest) GetExcludeIPRanges() *string {
	return s.ExcludeIPRanges
}

func (s *CreateServiceMeshRequest) GetExcludeInboundPorts() *string {
	return s.ExcludeInboundPorts
}

func (s *CreateServiceMeshRequest) GetExcludeOutboundPorts() *string {
	return s.ExcludeOutboundPorts
}

func (s *CreateServiceMeshRequest) GetExistingCaCert() *string {
	return s.ExistingCaCert
}

func (s *CreateServiceMeshRequest) GetExistingCaKey() *string {
	return s.ExistingCaKey
}

func (s *CreateServiceMeshRequest) GetExistingCaType() *string {
	return s.ExistingCaType
}

func (s *CreateServiceMeshRequest) GetExistingRootCaCert() *string {
	return s.ExistingRootCaCert
}

func (s *CreateServiceMeshRequest) GetExistingRootCaKey() *string {
	return s.ExistingRootCaKey
}

func (s *CreateServiceMeshRequest) GetFilterGatewayClusterConfig() *bool {
	return s.FilterGatewayClusterConfig
}

func (s *CreateServiceMeshRequest) GetGatewayAPIEnabled() *bool {
	return s.GatewayAPIEnabled
}

func (s *CreateServiceMeshRequest) GetGuestCluster() *string {
	return s.GuestCluster
}

func (s *CreateServiceMeshRequest) GetIncludeIPRanges() *string {
	return s.IncludeIPRanges
}

func (s *CreateServiceMeshRequest) GetIstioVersion() *string {
	return s.IstioVersion
}

func (s *CreateServiceMeshRequest) GetKialiEnabled() *bool {
	return s.KialiEnabled
}

func (s *CreateServiceMeshRequest) GetLocalityLBConf() *string {
	return s.LocalityLBConf
}

func (s *CreateServiceMeshRequest) GetLocalityLoadBalancing() *bool {
	return s.LocalityLoadBalancing
}

func (s *CreateServiceMeshRequest) GetMSEEnabled() *bool {
	return s.MSEEnabled
}

func (s *CreateServiceMeshRequest) GetMultiBufferEnabled() *bool {
	return s.MultiBufferEnabled
}

func (s *CreateServiceMeshRequest) GetMultiBufferPollDelay() *string {
	return s.MultiBufferPollDelay
}

func (s *CreateServiceMeshRequest) GetMysqlFilterEnabled() *bool {
	return s.MysqlFilterEnabled
}

func (s *CreateServiceMeshRequest) GetName() *string {
	return s.Name
}

func (s *CreateServiceMeshRequest) GetOPALimitCPU() *string {
	return s.OPALimitCPU
}

func (s *CreateServiceMeshRequest) GetOPALimitMemory() *string {
	return s.OPALimitMemory
}

func (s *CreateServiceMeshRequest) GetOPALogLevel() *string {
	return s.OPALogLevel
}

func (s *CreateServiceMeshRequest) GetOPARequestCPU() *string {
	return s.OPARequestCPU
}

func (s *CreateServiceMeshRequest) GetOPARequestMemory() *string {
	return s.OPARequestMemory
}

func (s *CreateServiceMeshRequest) GetOpaEnabled() *bool {
	return s.OpaEnabled
}

func (s *CreateServiceMeshRequest) GetOpenAgentPolicy() *bool {
	return s.OpenAgentPolicy
}

func (s *CreateServiceMeshRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateServiceMeshRequest) GetPilotLoadBalancerSpec() *string {
	return s.PilotLoadBalancerSpec
}

func (s *CreateServiceMeshRequest) GetPlaygroundScene() *string {
	return s.PlaygroundScene
}

func (s *CreateServiceMeshRequest) GetPrometheusUrl() *string {
	return s.PrometheusUrl
}

func (s *CreateServiceMeshRequest) GetProxyLimitCPU() *string {
	return s.ProxyLimitCPU
}

func (s *CreateServiceMeshRequest) GetProxyLimitMemory() *string {
	return s.ProxyLimitMemory
}

func (s *CreateServiceMeshRequest) GetProxyRequestCPU() *string {
	return s.ProxyRequestCPU
}

func (s *CreateServiceMeshRequest) GetProxyRequestMemory() *string {
	return s.ProxyRequestMemory
}

func (s *CreateServiceMeshRequest) GetRedisFilterEnabled() *bool {
	return s.RedisFilterEnabled
}

func (s *CreateServiceMeshRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceMeshRequest) GetTag() []*CreateServiceMeshRequestTag {
	return s.Tag
}

func (s *CreateServiceMeshRequest) GetTelemetry() *bool {
	return s.Telemetry
}

func (s *CreateServiceMeshRequest) GetThriftFilterEnabled() *bool {
	return s.ThriftFilterEnabled
}

func (s *CreateServiceMeshRequest) GetTraceSampling() *float32 {
	return s.TraceSampling
}

func (s *CreateServiceMeshRequest) GetTracing() *bool {
	return s.Tracing
}

func (s *CreateServiceMeshRequest) GetUseExistingCA() *bool {
	return s.UseExistingCA
}

func (s *CreateServiceMeshRequest) GetVSwitches() *string {
	return s.VSwitches
}

func (s *CreateServiceMeshRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateServiceMeshRequest) GetWebAssemblyFilterEnabled() *bool {
	return s.WebAssemblyFilterEnabled
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

func (s *CreateServiceMeshRequest) SetCertChain(v string) *CreateServiceMeshRequest {
	s.CertChain = &v
	return s
}

func (s *CreateServiceMeshRequest) SetChargeType(v string) *CreateServiceMeshRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateServiceMeshRequest) SetClusterDomain(v string) *CreateServiceMeshRequest {
	s.ClusterDomain = &v
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

func (s *CreateServiceMeshRequest) SetEnableACMG(v bool) *CreateServiceMeshRequest {
	s.EnableACMG = &v
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

func (s *CreateServiceMeshRequest) SetPlaygroundScene(v string) *CreateServiceMeshRequest {
	s.PlaygroundScene = &v
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

func (s *CreateServiceMeshRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServiceMeshRequestTag struct {
	// The name of the tag.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceMeshRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceMeshRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateServiceMeshRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateServiceMeshRequestTag) SetKey(v string) *CreateServiceMeshRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceMeshRequestTag) SetValue(v string) *CreateServiceMeshRequestTag {
	s.Value = &v
	return s
}

func (s *CreateServiceMeshRequestTag) Validate() error {
	return dara.Validate(s)
}
