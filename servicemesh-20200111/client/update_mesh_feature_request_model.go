// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLogEnabled(v bool) *UpdateMeshFeatureRequest
	GetAccessLogEnabled() *bool
	SetAccessLogFile(v string) *UpdateMeshFeatureRequest
	GetAccessLogFile() *string
	SetAccessLogFormat(v string) *UpdateMeshFeatureRequest
	GetAccessLogFormat() *string
	SetAccessLogGatewayEnabled(v bool) *UpdateMeshFeatureRequest
	GetAccessLogGatewayEnabled() *bool
	SetAccessLogGatewayLifecycle(v int32) *UpdateMeshFeatureRequest
	GetAccessLogGatewayLifecycle() *int32
	SetAccessLogProject(v string) *UpdateMeshFeatureRequest
	GetAccessLogProject() *string
	SetAccessLogServiceEnabled(v bool) *UpdateMeshFeatureRequest
	GetAccessLogServiceEnabled() *bool
	SetAccessLogServiceHost(v string) *UpdateMeshFeatureRequest
	GetAccessLogServiceHost() *string
	SetAccessLogServicePort(v int32) *UpdateMeshFeatureRequest
	GetAccessLogServicePort() *int32
	SetAccessLogSidecarEnabled(v bool) *UpdateMeshFeatureRequest
	GetAccessLogSidecarEnabled() *bool
	SetAccessLogSidecarLifecycle(v int32) *UpdateMeshFeatureRequest
	GetAccessLogSidecarLifecycle() *int32
	SetAuditProject(v string) *UpdateMeshFeatureRequest
	GetAuditProject() *string
	SetAutoInjectionPolicyEnabled(v bool) *UpdateMeshFeatureRequest
	GetAutoInjectionPolicyEnabled() *bool
	SetCRAggregationEnabled(v bool) *UpdateMeshFeatureRequest
	GetCRAggregationEnabled() *bool
	SetCertChain(v string) *UpdateMeshFeatureRequest
	GetCertChain() *string
	SetClusterSpec(v string) *UpdateMeshFeatureRequest
	GetClusterSpec() *string
	SetCniEnabled(v bool) *UpdateMeshFeatureRequest
	GetCniEnabled() *bool
	SetCniExcludeNamespaces(v string) *UpdateMeshFeatureRequest
	GetCniExcludeNamespaces() *string
	SetConcurrency(v int32) *UpdateMeshFeatureRequest
	GetConcurrency() *int32
	SetConfigSourceEnabled(v bool) *UpdateMeshFeatureRequest
	GetConfigSourceEnabled() *bool
	SetConfigSourceNacosID(v string) *UpdateMeshFeatureRequest
	GetConfigSourceNacosID() *string
	SetCustomizedPrometheus(v bool) *UpdateMeshFeatureRequest
	GetCustomizedPrometheus() *bool
	SetCustomizedZipkin(v bool) *UpdateMeshFeatureRequest
	GetCustomizedZipkin() *bool
	SetDNSProxyingEnabled(v bool) *UpdateMeshFeatureRequest
	GetDNSProxyingEnabled() *bool
	SetDefaultComponentsScheduleConfig(v string) *UpdateMeshFeatureRequest
	GetDefaultComponentsScheduleConfig() *string
	SetDiscoverySelectors(v string) *UpdateMeshFeatureRequest
	GetDiscoverySelectors() *string
	SetDubboFilterEnabled(v bool) *UpdateMeshFeatureRequest
	GetDubboFilterEnabled() *bool
	SetEnableAudit(v bool) *UpdateMeshFeatureRequest
	GetEnableAudit() *bool
	SetEnableAutoDiagnosis(v bool) *UpdateMeshFeatureRequest
	GetEnableAutoDiagnosis() *bool
	SetEnableBootstrapXdsAgent(v bool) *UpdateMeshFeatureRequest
	GetEnableBootstrapXdsAgent() *bool
	SetEnableCRHistory(v bool) *UpdateMeshFeatureRequest
	GetEnableCRHistory() *bool
	SetEnableNamespacesByDefault(v bool) *UpdateMeshFeatureRequest
	GetEnableNamespacesByDefault() *bool
	SetEnableSDSServer(v bool) *UpdateMeshFeatureRequest
	GetEnableSDSServer() *bool
	SetExcludeIPRanges(v string) *UpdateMeshFeatureRequest
	GetExcludeIPRanges() *string
	SetExcludeInboundPorts(v string) *UpdateMeshFeatureRequest
	GetExcludeInboundPorts() *string
	SetExcludeOutboundPorts(v string) *UpdateMeshFeatureRequest
	GetExcludeOutboundPorts() *string
	SetExistingCaCert(v string) *UpdateMeshFeatureRequest
	GetExistingCaCert() *string
	SetExistingCaKey(v string) *UpdateMeshFeatureRequest
	GetExistingCaKey() *string
	SetExistingRootCaCert(v string) *UpdateMeshFeatureRequest
	GetExistingRootCaCert() *string
	SetFilterGatewayClusterConfig(v bool) *UpdateMeshFeatureRequest
	GetFilterGatewayClusterConfig() *bool
	SetGatewayAPIEnabled(v bool) *UpdateMeshFeatureRequest
	GetGatewayAPIEnabled() *bool
	SetHoldApplicationUntilProxyStarts(v bool) *UpdateMeshFeatureRequest
	GetHoldApplicationUntilProxyStarts() *bool
	SetHttp10Enabled(v bool) *UpdateMeshFeatureRequest
	GetHttp10Enabled() *bool
	SetIncludeIPRanges(v string) *UpdateMeshFeatureRequest
	GetIncludeIPRanges() *string
	SetIncludeInboundPorts(v string) *UpdateMeshFeatureRequest
	GetIncludeInboundPorts() *string
	SetIncludeOutboundPorts(v string) *UpdateMeshFeatureRequest
	GetIncludeOutboundPorts() *string
	SetIntegrateKiali(v bool) *UpdateMeshFeatureRequest
	GetIntegrateKiali() *bool
	SetInterceptionMode(v string) *UpdateMeshFeatureRequest
	GetInterceptionMode() *string
	SetKialiArmsAuthTokens(v string) *UpdateMeshFeatureRequest
	GetKialiArmsAuthTokens() *string
	SetKialiEnabled(v bool) *UpdateMeshFeatureRequest
	GetKialiEnabled() *bool
	SetKialiServiceAnnotations(v string) *UpdateMeshFeatureRequest
	GetKialiServiceAnnotations() *string
	SetLabelsForOffloadedWorkloads(v string) *UpdateMeshFeatureRequest
	GetLabelsForOffloadedWorkloads() *string
	SetLifecycle(v string) *UpdateMeshFeatureRequest
	GetLifecycle() *string
	SetLocalityLBConf(v string) *UpdateMeshFeatureRequest
	GetLocalityLBConf() *string
	SetLocalityLoadBalancing(v bool) *UpdateMeshFeatureRequest
	GetLocalityLoadBalancing() *bool
	SetLogLevel(v string) *UpdateMeshFeatureRequest
	GetLogLevel() *string
	SetMSEEnabled(v bool) *UpdateMeshFeatureRequest
	GetMSEEnabled() *bool
	SetMultiBufferEnabled(v bool) *UpdateMeshFeatureRequest
	GetMultiBufferEnabled() *bool
	SetMultiBufferPollDelay(v string) *UpdateMeshFeatureRequest
	GetMultiBufferPollDelay() *string
	SetMysqlFilterEnabled(v bool) *UpdateMeshFeatureRequest
	GetMysqlFilterEnabled() *bool
	SetNFDEnabled(v bool) *UpdateMeshFeatureRequest
	GetNFDEnabled() *bool
	SetNFDLabelPruned(v bool) *UpdateMeshFeatureRequest
	GetNFDLabelPruned() *bool
	SetOPAInjectorCPULimit(v string) *UpdateMeshFeatureRequest
	GetOPAInjectorCPULimit() *string
	SetOPAInjectorCPURequirement(v string) *UpdateMeshFeatureRequest
	GetOPAInjectorCPURequirement() *string
	SetOPAInjectorMemoryLimit(v string) *UpdateMeshFeatureRequest
	GetOPAInjectorMemoryLimit() *string
	SetOPAInjectorMemoryRequirement(v string) *UpdateMeshFeatureRequest
	GetOPAInjectorMemoryRequirement() *string
	SetOPALimitCPU(v string) *UpdateMeshFeatureRequest
	GetOPALimitCPU() *string
	SetOPALimitMemory(v string) *UpdateMeshFeatureRequest
	GetOPALimitMemory() *string
	SetOPALogLevel(v string) *UpdateMeshFeatureRequest
	GetOPALogLevel() *string
	SetOPARequestCPU(v string) *UpdateMeshFeatureRequest
	GetOPARequestCPU() *string
	SetOPARequestMemory(v string) *UpdateMeshFeatureRequest
	GetOPARequestMemory() *string
	SetOPAScopeInjected(v bool) *UpdateMeshFeatureRequest
	GetOPAScopeInjected() *bool
	SetOpaEnabled(v bool) *UpdateMeshFeatureRequest
	GetOpaEnabled() *bool
	SetOpenAgentPolicy(v bool) *UpdateMeshFeatureRequest
	GetOpenAgentPolicy() *bool
	SetOutboundTrafficPolicy(v string) *UpdateMeshFeatureRequest
	GetOutboundTrafficPolicy() *string
	SetPilotEnableQuicListeners(v bool) *UpdateMeshFeatureRequest
	GetPilotEnableQuicListeners() *bool
	SetPrometheusUrl(v string) *UpdateMeshFeatureRequest
	GetPrometheusUrl() *string
	SetProxyInitCPUResourceLimit(v string) *UpdateMeshFeatureRequest
	GetProxyInitCPUResourceLimit() *string
	SetProxyInitCPUResourceRequest(v string) *UpdateMeshFeatureRequest
	GetProxyInitCPUResourceRequest() *string
	SetProxyInitMemoryResourceLimit(v string) *UpdateMeshFeatureRequest
	GetProxyInitMemoryResourceLimit() *string
	SetProxyInitMemoryResourceRequest(v string) *UpdateMeshFeatureRequest
	GetProxyInitMemoryResourceRequest() *string
	SetProxyLimitCPU(v string) *UpdateMeshFeatureRequest
	GetProxyLimitCPU() *string
	SetProxyLimitMemory(v string) *UpdateMeshFeatureRequest
	GetProxyLimitMemory() *string
	SetProxyRequestCPU(v string) *UpdateMeshFeatureRequest
	GetProxyRequestCPU() *string
	SetProxyRequestMemory(v string) *UpdateMeshFeatureRequest
	GetProxyRequestMemory() *string
	SetProxyStatsMatcher(v string) *UpdateMeshFeatureRequest
	GetProxyStatsMatcher() *string
	SetRedisFilterEnabled(v bool) *UpdateMeshFeatureRequest
	GetRedisFilterEnabled() *bool
	SetSMCEnabled(v bool) *UpdateMeshFeatureRequest
	GetSMCEnabled() *bool
	SetServiceMeshId(v string) *UpdateMeshFeatureRequest
	GetServiceMeshId() *string
	SetSidecarInjectorLimitCPU(v string) *UpdateMeshFeatureRequest
	GetSidecarInjectorLimitCPU() *string
	SetSidecarInjectorLimitMemory(v string) *UpdateMeshFeatureRequest
	GetSidecarInjectorLimitMemory() *string
	SetSidecarInjectorRequestCPU(v string) *UpdateMeshFeatureRequest
	GetSidecarInjectorRequestCPU() *string
	SetSidecarInjectorRequestMemory(v string) *UpdateMeshFeatureRequest
	GetSidecarInjectorRequestMemory() *string
	SetSidecarInjectorWebhookAsYaml(v string) *UpdateMeshFeatureRequest
	GetSidecarInjectorWebhookAsYaml() *string
	SetTelemetry(v bool) *UpdateMeshFeatureRequest
	GetTelemetry() *bool
	SetTerminationDrainDuration(v string) *UpdateMeshFeatureRequest
	GetTerminationDrainDuration() *string
	SetThriftFilterEnabled(v bool) *UpdateMeshFeatureRequest
	GetThriftFilterEnabled() *bool
	SetTraceCustomTags(v string) *UpdateMeshFeatureRequest
	GetTraceCustomTags() *string
	SetTraceMaxPathTagLength(v string) *UpdateMeshFeatureRequest
	GetTraceMaxPathTagLength() *string
	SetTraceSampling(v float32) *UpdateMeshFeatureRequest
	GetTraceSampling() *float32
	SetTracing(v bool) *UpdateMeshFeatureRequest
	GetTracing() *bool
	SetTracingOnExtZipkinLimitCPU(v string) *UpdateMeshFeatureRequest
	GetTracingOnExtZipkinLimitCPU() *string
	SetTracingOnExtZipkinLimitMemory(v string) *UpdateMeshFeatureRequest
	GetTracingOnExtZipkinLimitMemory() *string
	SetTracingOnExtZipkinReplicaCount(v string) *UpdateMeshFeatureRequest
	GetTracingOnExtZipkinReplicaCount() *string
	SetTracingOnExtZipkinRequestCPU(v string) *UpdateMeshFeatureRequest
	GetTracingOnExtZipkinRequestCPU() *string
	SetTracingOnExtZipkinRequestMemory(v string) *UpdateMeshFeatureRequest
	GetTracingOnExtZipkinRequestMemory() *string
	SetWebAssemblyFilterEnabled(v bool) *UpdateMeshFeatureRequest
	GetWebAssemblyFilterEnabled() *bool
}

type UpdateMeshFeatureRequest struct {
	// Specifies whether to enable access log collection. Valid values:
	//
	// 	- `true`: enables access log collection.
	//
	// 	- `false`: disables access log collection.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AccessLogEnabled *bool `json:"AccessLogEnabled,omitempty" xml:"AccessLogEnabled,omitempty"`
	// Specifies whether to enable access logging. Valid values:
	//
	// 	- `""`: disables access logging.
	//
	// 	- `/dev/stdout`: enables access logging. Access logs are written to /dev/stdout.
	//
	// example:
	//
	// “”
	AccessLogFile *string `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	// The custom format of access logs. To set this parameter, make sure that you have enabled access log collection. The value must be a JSON string. The following key names must be contained: authority_for, bytes_received, bytes_sent, downstream_local_address, downstream_remote_address, duration, istio_policy_status, method, path, protocol, requested_server_name, response_code, response_flags, route_name, start_time, trace_id, upstream_cluster, upstream_host, upstream_local_address, upstream_service_time, upstream_transport_failure_reason, user_agent, and x_forwarded_for.
	//
	// example:
	//
	// {"authority_for":"%REQ(:AUTHORITY)%","bytes_received":"%BYTES_RECEIVED%","bytes_sent":"%BYTES_SENT%","downstream_local_address":"%DOWNSTREAM_LOCAL_ADDRESS%","downstream_remote_address":"%DOWNSTREAM_REMOTE_ADDRESS%","duration":"%DURATION%","istio_policy_status":"%DYNAMIC_METADATA(istio.mixer:status)%","method":"%REQ(:METHOD)%","path":"%REQ(X-ENVOY-ORIGINAL-PATH?:PATH)%","protocol":"%PROTOCOL%","request_id":"%REQ(X-REQUEST-ID)%","requested_server_name":"%REQUESTED_SERVER_NAME%","response_code":"%RESPONSE_CODE%","response_flags":"%RESPONSE_FLAGS%","route_name":"%ROUTE_NAME%","start_time":"%START_TIME%","trace_id":"%REQ(X-B3-TRACEID)%","upstream_cluster":"%UPSTREAM_CLUSTER%","upstream_host":"%UPSTREAM_HOST%","upstream_local_address":"%UPSTREAM_LOCAL_ADDRESS%","upstream_service_time":"%RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)%","upstream_transport_failure_reason":"%UPSTREAM_TRANSPORT_FAILURE_REASON%","user_agent":"%REQ(USER-AGENT)%","x_forwarded_for":"%REQ(X-FORWARDED-FOR)%"}
	AccessLogFormat *string `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	// Specifies whether to enable the collection of access logs of ASM gateways to Simple Log Service.
	//
	// example:
	//
	// false
	AccessLogGatewayEnabled *bool `json:"AccessLogGatewayEnabled,omitempty" xml:"AccessLogGatewayEnabled,omitempty"`
	// The retention period for the access logs of the sidecar proxy. Unit: day. The logs are collected by using Log Service. For example, `30` indicates 30 days.
	//
	// example:
	//
	// 30
	AccessLogGatewayLifecycle *int32 `json:"AccessLogGatewayLifecycle,omitempty" xml:"AccessLogGatewayLifecycle,omitempty"`
	// The custom project on which the Log Service collects logs.
	//
	// example:
	//
	// mesh-log-cf245a429b6ff4b6e97f20797758e****
	AccessLogProject *string `json:"AccessLogProject,omitempty" xml:"AccessLogProject,omitempty"`
	// Specifies whether to enable gRPC Access Log Service (ALS) for Envoy. Valid values:
	//
	// 	- `true`: enables gRPC ALS for Envoy.
	//
	// 	- `false`: disables gRPC ALS for Envoy.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AccessLogServiceEnabled *bool `json:"AccessLogServiceEnabled,omitempty" xml:"AccessLogServiceEnabled,omitempty"`
	// The endpoint of gRPC ALS for Envoy.
	//
	// example:
	//
	// 0.0.0.0
	AccessLogServiceHost *string `json:"AccessLogServiceHost,omitempty" xml:"AccessLogServiceHost,omitempty"`
	// The port of gRPC ALS for Envoy.
	//
	// example:
	//
	// 9999
	AccessLogServicePort *int32 `json:"AccessLogServicePort,omitempty" xml:"AccessLogServicePort,omitempty"`
	// Specifies whether to enable the collection of access logs of sidecar proxies to Simple Log Service.
	//
	// example:
	//
	// false
	AccessLogSidecarEnabled *bool `json:"AccessLogSidecarEnabled,omitempty" xml:"AccessLogSidecarEnabled,omitempty"`
	// Specifies whether to enable automatic diagnostics for the ASM instance. If you enable this feature, the ASM instance is automatically diagnosed when you modify Istio resources in the ASM instance.
	//
	// example:
	//
	// 30
	AccessLogSidecarLifecycle *int32 `json:"AccessLogSidecarLifecycle,omitempty" xml:"AccessLogSidecarLifecycle,omitempty"`
	// The name of the Log Service project that is used for mesh audit.
	//
	// Default value: `mesh-log-{ASM instance ID}`.
	//
	// example:
	//
	// mesh-log-c08ba3fd1e64xxb0f8cc1ad8****
	AuditProject *string `json:"AuditProject,omitempty" xml:"AuditProject,omitempty"`
	// Specifies whether to enable automatic sidecar proxy injection by using pod annotations. Valid values:
	//
	// 	- `true`: enables automatic sidecar proxy injection by using pod annotations.
	//
	// 	- `false`: disables automatic sidecar proxy injection by using pod annotations.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoInjectionPolicyEnabled *bool `json:"AutoInjectionPolicyEnabled,omitempty" xml:"AutoInjectionPolicyEnabled,omitempty"`
	// Specifies whether to use the Kubernetes API of clusters on the data plane to access Istio resources. To use this feature, the version of the ASM instance must be V1.9.7.93 or later.
	//
	// example:
	//
	// false
	CRAggregationEnabled *bool `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	// The certificate chain from the CA certificate to the root certificate. At least two certificates are included in the chain.
	//
	// example:
	//
	// Base64 encoded PEM cert chain.
	CertChain *string `json:"CertChain,omitempty" xml:"CertChain,omitempty"`
	// Specifies whether to enable the feature of controlling the OPA injection scope. Valid values:
	//
	// 	- `true`: enables the feature.
	//
	// 	- `false`: disables the feature.
	//
	// example:
	//
	// standard
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// Specifies whether to enable the Container Network Interface (CNI) plug-in. Valid values:
	//
	// 	- `true`: enables the CNI plug-in.
	//
	// 	- `false`: disables the CNI plug-in.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CniEnabled *bool `json:"CniEnabled,omitempty" xml:"CniEnabled,omitempty"`
	// The namespaces to be excluded for the CNI plug-in.
	//
	// example:
	//
	// kube-system
	CniExcludeNamespaces *string `json:"CniExcludeNamespaces,omitempty" xml:"CniExcludeNamespaces,omitempty"`
	// Specifies whether to delay application container startup until the sidecar proxy container is started in a pod.
	//
	// example:
	//
	// 2
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// Specifies whether to enable the external service registry. Valid values:
	//
	// 	- `true`: enables the external service registry.
	//
	// 	- `false`: disables the external service registry.
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
	// Specifies whether to use a custom Prometheus instance. Valid values:
	//
	// 	- `true`: uses a custom Prometheus instance.
	//
	// 	- `false`: does not use a custom Prometheus instance.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CustomizedPrometheus *bool `json:"CustomizedPrometheus,omitempty" xml:"CustomizedPrometheus,omitempty"`
	// Specifies whether to use a self-managed Zipkin system to collect tracing data. Valid values:
	//
	// 	- `true`: uses a self-managed Zipkin system.
	//
	// 	- `false`: does not use a self-managed Zipkin system.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CustomizedZipkin *bool `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	// Specifies whether to enable DNS proxy. Valid values:
	//
	// 	- `true`: enables the DNS proxy feature.
	//
	// 	- `false`: disables the DNS proxy feature.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DNSProxyingEnabled *bool `json:"DNSProxyingEnabled,omitempty" xml:"DNSProxyingEnabled,omitempty"`
	// Specifies the default scheduling configurations that ASM delivers to components on the data plane. You can configure `nodeSelector` and `tolerations` in the JSON format.
	//
	// >
	//
	// 	- Modifying the value of this parameter is a high-risk operation. The modification will cause all components on the data plane of ASM to restart. Exercise caution before modifying the value of this parameter.
	//
	// 	- The configurations specified by this parameter do not apply to ASM gateways. You can configure gateway-specific scheduling on ASM gateways.
	//
	// example:
	//
	// {"tolerations":[{"key":"test-taints", "operator":"Exists", "effect":"NoSchedule"}], "nodeSelector":{"kubernetes.io/hostname":"test-nodes"}}
	DefaultComponentsScheduleConfig *string `json:"DefaultComponentsScheduleConfig,omitempty" xml:"DefaultComponentsScheduleConfig,omitempty"`
	// The label selectors used to specify the namespaces of the clusters on the data plane for selective service discovery.
	//
	// example:
	//
	// [{"matchExpressions":[{"key":"asm-discovery","operator":"Exists"}]}]
	DiscoverySelectors *string `json:"DiscoverySelectors,omitempty" xml:"DiscoverySelectors,omitempty"`
	// Specifies whether to enable Dubbo Filter. Valid values:
	//
	// 	- `true`: enables Dubbo Filter.
	//
	// 	- `false`: disables Dubbo Filter.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DubboFilterEnabled *bool `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	// Specifies whether to enable the mesh audit feature. To enable this feature, make sure that you have activated [Log Service](https://sls.console.aliyun.com/). Valid values:
	//
	// 	- `true`: enables the mesh audit feature.
	//
	// 	- `false`: disables the mesh audit feature.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	EnableAudit *bool `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	// The ports for which outbound traffic is redirected to the sidecar proxy.
	//
	// example:
	//
	// true
	EnableAutoDiagnosis *bool `json:"EnableAutoDiagnosis,omitempty" xml:"EnableAutoDiagnosis,omitempty"`
	// Specifies the authentication token of an ARMS Prometheus instance when the Mesh Topology feature is enabled and ARMS Prometheus is used to collect monitoring metrics. The token is used to allow Mesh Topology to access the ARMS Prometheus instance. The token is in the JSON format. The key in the JSON object is the ID of the cluster on the data plane, and the value is the authentication token of the ARMS Prometheus instance deployed in the cluster.
	//
	// example:
	//
	// true
	EnableBootstrapXdsAgent *bool `json:"EnableBootstrapXdsAgent,omitempty" xml:"EnableBootstrapXdsAgent,omitempty"`
	// Specifies whether to enable the rollback feature for Istio resources.
	//
	// example:
	//
	// false
	EnableCRHistory *bool `json:"EnableCRHistory,omitempty" xml:"EnableCRHistory,omitempty"`
	// Specifies whether to enable automatic sidecar proxy injection for all namespaces. Valid values:
	//
	// 	- `true`: enables automatic sidecar proxy injection for all namespaces.
	//
	// 	- `false`: disables automatic sidecar proxy injection for all namespaces.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	EnableNamespacesByDefault *bool `json:"EnableNamespacesByDefault,omitempty" xml:"EnableNamespacesByDefault,omitempty"`
	// Specifies whether to enable Secret Discovery Service (SDS). Valid values:
	//
	// 	- `true`: enables SDS.
	//
	// 	- `false`: disables SDS.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	EnableSDSServer *bool `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
	// The IP addresses of external services to which traffic is not intercepted.
	//
	// example:
	//
	// 100.100.XXX.XXX
	ExcludeIPRanges *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	// The ports for which inbound traffic is not redirected to the sidecar proxy. Separate multiple ports with commas (,).
	//
	// example:
	//
	// 80,81
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The ports for which outbound traffic is not redirected to the sidecar proxy. Separate multiple ports with commas (,).
	//
	// example:
	//
	// 80,81
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// The updated certificate authority (CA) certificate. You can modify this parameter only if you use a custom certificate when you create an ASM instance.
	//
	// example:
	//
	// Base64 encoded PEM certificate.
	ExistingCaCert *string `json:"ExistingCaCert,omitempty" xml:"ExistingCaCert,omitempty"`
	// The updated CA certificate key. You can modify this parameter only if you use a custom certificate when you create an ASM instance.
	//
	// example:
	//
	// Base64 encoded PEM
	//
	//  private key.
	ExistingCaKey *string `json:"ExistingCaKey,omitempty" xml:"ExistingCaKey,omitempty"`
	// The updated root certificate. You can modify this parameter only if you use a custom certificate when you create a Service Mesh (ASM) instance.
	//
	// example:
	//
	// Base64 encoded PEM certificate.
	ExistingRootCaCert *string `json:"ExistingRootCaCert,omitempty" xml:"ExistingRootCaCert,omitempty"`
	// Specifies whether to enable gateway configuration filtering. Valid values:
	//
	// 	- `true`: enables gateway configuration filtering.
	//
	// 	- `false`: disables gateway configuration filtering.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	FilterGatewayClusterConfig *bool `json:"FilterGatewayClusterConfig,omitempty" xml:"FilterGatewayClusterConfig,omitempty"`
	// Specifies whether to enable Gateway API. Valid values:
	//
	// 	- `true`: enables Gateway API.
	//
	// 	- `false`: disables Gateway API.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	GatewayAPIEnabled *bool `json:"GatewayAPIEnabled,omitempty" xml:"GatewayAPIEnabled,omitempty"`
	// Other metrics of the sidecar proxy on the data plane.
	//
	// example:
	//
	// true
	HoldApplicationUntilProxyStarts *bool `json:"HoldApplicationUntilProxyStarts,omitempty" xml:"HoldApplicationUntilProxyStarts,omitempty"`
	// Specifies whether to support HTTP 1.0. Valid values:
	//
	// 	- `true`: supports HTTP 1.0.
	//
	// 	- `false`: does not support HTTP 1.0.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Http10Enabled *bool `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
	// The IP addresses of external services to which traffic is intercepted.
	//
	// example:
	//
	// *
	IncludeIPRanges *string `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	// The ports for which inbound traffic is redirected to the sidecar proxy.
	//
	// example:
	//
	// 80,81
	IncludeInboundPorts *string `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	// The log level of the sidecar proxy on the data plane. Log levels include `none`, `error`, `warn`, `info`, and `debug`. The levels correspond to different amounts of information reported by the logs. For example, none-level logs report the least information, while debug-level logs report the most information.
	//
	// example:
	//
	// 8000,8001
	IncludeOutboundPorts *string `json:"IncludeOutboundPorts,omitempty" xml:"IncludeOutboundPorts,omitempty"`
	// Specifies whether to create a Classic Load Balancer (CLB) instance for accessing Mesh Topology of Service Mesh (ASM).
	//
	// example:
	//
	// false
	IntegrateKiali *bool `json:"IntegrateKiali,omitempty" xml:"IntegrateKiali,omitempty"`
	// Specifies whether to load the bootstrap configuration before the sidecar proxy is started.
	//
	// example:
	//
	// TPROXY
	InterceptionMode *string `json:"InterceptionMode,omitempty" xml:"InterceptionMode,omitempty"`
	// Specifies the default scheduling configurations that ASM delivers to components on the data plane. You can configure `nodeSelector` and tolerations in the JSON format.
	//
	// > 	- Modifying the value of this parameter is a high-risk operation. The modification will cause all components on the data plane of ASM to restart. Exercise caution before modifying the value of this parameter.
	//
	// >	- The configurations specified by this parameter do not apply to the ASM gateway. You can configure gateway-specific scheduling on the ASM gateway.
	//
	// example:
	//
	// {"c31e3b******5634b":"token_example"}
	KialiArmsAuthTokens *string `json:"KialiArmsAuthTokens,omitempty" xml:"KialiArmsAuthTokens,omitempty"`
	// Specifies whether to enable the Mesh Topology feature. To enable this feature, make sure that you have enabled Prometheus monitoring. If Prometheus monitoring is disabled, the Mesh Topology feature must be disabled. Valid values:````
	//
	// 	- `true`: enables the Mesh Topology feature.
	//
	// 	- `false`: disables the Mesh Topology feature.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	KialiEnabled *bool `json:"KialiEnabled,omitempty" xml:"KialiEnabled,omitempty"`
	// Specifies Classic Load Balancer (CLB) instances by using annotations when the Mesh Topology feature is enabled. These CLB instances are used to access the Mesh Topology feature in different clusters.
	//
	// This parameter is a JSON-encoded string. The key in the JSON object is the ID of a cluster on the data plane, and the value is the annotation content of the Mesh Topology service in the cluster.
	//
	// For more information about how to configure CLB instances by using annotations, see [Add annotations to the YAML file of a Service to configure CLB instances](https://www.alibabacloud.com/help/container-service-for-kubernetes/latest/use-annotations-to-configure-load-balancing-1).
	//
	// example:
	//
	// {"c31e3b******5634b":{"service.beta.kubernetes.io/alibaba-cloud-loadbalancer-address-type":"intranet"}}
	KialiServiceAnnotations *string `json:"KialiServiceAnnotations,omitempty" xml:"KialiServiceAnnotations,omitempty"`
	// The labels for isolated workloads.
	//
	// example:
	//
	// name=xx,region=xx
	LabelsForOffloadedWorkloads *string `json:"LabelsForOffloadedWorkloads,omitempty" xml:"LabelsForOffloadedWorkloads,omitempty"`
	// The lifecycle of the sidecar proxy.
	//
	// example:
	//
	// {"postStart":{"exec":{"command":["pilot-agent","wait"]}},"preStop":{"exec":{"command":["/bin/sh","-c","sleep 15"]}}}
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The configurations of cross-region load balancing. Valid values:
	//
	// 	- `failover`: the configurations of cross-region failover. Example:
	//
	// <!---->
	//
	//     failover: [// Cross-region failover configurations of the struct type.
	//
	//             {
	//
	//                 // If a service fails in the China (Beijing) region, the traffic is redirected to the same service in the China (Hangzhou) region.
	//
	//                 from: "cn-beijing",
	//
	//                 to: "cn-hangzhou",
	//
	//             }
	//
	//         ]
	//
	// 	- `distribute`: the configurations of cross-region traffic distribution. Example:
	//
	// <!---->
	//
	//     distribute: [// Cross-region traffic distribution configurations of the struct type.
	//
	//             {
	//
	//                 // For traffic that is routed to the China (Beijing) region, 70% of the traffic is allocated to the China (Beijing) region and the rest of the traffic is redirected to the China (Hangzhou) region.
	//
	//                 "from": "cn-beijing",
	//
	//                 "to": {
	//
	//                     "cn-beijing": 70,
	//
	//                     "cn-hangzhou": 30,
	//
	//                 }
	//
	//             }
	//
	//         ]
	//
	// example:
	//
	// {"failover":[{"from":"cn-hangzhou","to":"cn-shanghai"}]}
	LocalityLBConf *string `json:"LocalityLBConf,omitempty" xml:"LocalityLBConf,omitempty"`
	// Specifies whether to enable cross-region load balancing. Valid values:
	//
	// 	- `true`: enables cross-region load balancing.
	//
	// 	- `false`: disables cross-region load balancing.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	LocalityLoadBalancing *bool `json:"LocalityLoadBalancing,omitempty" xml:"LocalityLoadBalancing,omitempty"`
	// The number of worker threads used by the sidecar proxy on the data plane.
	//
	// example:
	//
	// info
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// Deprecated
	//
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
	// Specifies whether to enable Transport Layer Security (TLS) acceleration based on MultiBuffer.
	//
	// example:
	//
	// false
	MultiBufferEnabled *bool `json:"MultiBufferEnabled,omitempty" xml:"MultiBufferEnabled,omitempty"`
	// The pull-request latency. By default, this parameter is left empty.
	//
	// example:
	//
	// 0.02s
	MultiBufferPollDelay *string `json:"MultiBufferPollDelay,omitempty" xml:"MultiBufferPollDelay,omitempty"`
	// Specifies whether to enable MySQL Filter. Valid values:
	//
	// 	- `true`: enables MySQL Filter.
	//
	// 	- `false`: disables MySQL Filter.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	MysqlFilterEnabled *bool `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	// Specifies whether to clear feature labels on nodes when NFD is disabled.
	//
	// This parameter is valid only when the `NFDEnabled` parameter is set to `false`.
	//
	// example:
	//
	// false
	NFDEnabled *bool `json:"NFDEnabled,omitempty" xml:"NFDEnabled,omitempty"`
	// The minimum number of CPU cores requested by the proxy service that exports Tracing Analysis data. For example, `1000m` indicates one CPU core.
	//
	// example:
	//
	// false
	NFDLabelPruned *bool `json:"NFDLabelPruned,omitempty" xml:"NFDLabelPruned,omitempty"`
	// The maximum size of the memory that is available to the pod that injects OPA proxies into application pods. For example, `1024Mi` indicates 1024 MB.
	//
	// example:
	//
	// 1000m
	OPAInjectorCPULimit *string `json:"OPAInjectorCPULimit,omitempty" xml:"OPAInjectorCPULimit,omitempty"`
	// The minimum size of the memory requested by the pod that injects OPA proxies into application pods. For example, `50 Mi` indicates 50 MB.
	//
	// example:
	//
	// 80m
	OPAInjectorCPURequirement *string `json:"OPAInjectorCPURequirement,omitempty" xml:"OPAInjectorCPURequirement,omitempty"`
	// Specifies whether to create a CLB instance for accessing the ASM mesh topology.
	//
	// example:
	//
	// 1024Mi
	OPAInjectorMemoryLimit *string `json:"OPAInjectorMemoryLimit,omitempty" xml:"OPAInjectorMemoryLimit,omitempty"`
	// The maximum number of CPU cores that are available to the pod that injects OPA proxies into application pods. For example, `1000m` indicates one CPU core.
	//
	// example:
	//
	// 50Mi
	OPAInjectorMemoryRequirement *string `json:"OPAInjectorMemoryRequirement,omitempty" xml:"OPAInjectorMemoryRequirement,omitempty"`
	// The maximum number of CPU cores that are available to the OPA proxy container.
	//
	// example:
	//
	// 2
	OPALimitCPU *string `json:"OPALimitCPU,omitempty" xml:"OPALimitCPU,omitempty"`
	// The maximum size of the memory that is available to the OPA proxy container.
	//
	// example:
	//
	// 1024Mi
	OPALimitMemory *string `json:"OPALimitMemory,omitempty" xml:"OPALimitMemory,omitempty"`
	// The log level of the OPA proxy container.
	//
	// 	- `info`: outputs all information.
	//
	// 	- `debug`: outputs debugging and error information.
	//
	// 	- `error`: outputs only error information.
	//
	// example:
	//
	// info
	OPALogLevel *string `json:"OPALogLevel,omitempty" xml:"OPALogLevel,omitempty"`
	// The number of CPU cores that are requested by the OPA proxy container.
	//
	// example:
	//
	// 1
	OPARequestCPU *string `json:"OPARequestCPU,omitempty" xml:"OPARequestCPU,omitempty"`
	// The size of the memory that is requested by the OPA proxy container.
	//
	// example:
	//
	// 512Mi
	OPARequestMemory *string `json:"OPARequestMemory,omitempty" xml:"OPARequestMemory,omitempty"`
	// The minimum number of CPU cores requested by the pod that injects OPA proxies into application pods. For example, `1000m` indicates one CPU core.
	//
	// example:
	//
	// false
	OPAScopeInjected *bool `json:"OPAScopeInjected,omitempty" xml:"OPAScopeInjected,omitempty"`
	// Specifies whether to enable the OPA plug-in. Valid values:
	//
	// 	- `true`: enables the OPA plug-in.
	//
	// 	- `false`: disables the OPA plug-in.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	OpaEnabled *bool `json:"OpaEnabled,omitempty" xml:"OpaEnabled,omitempty"`
	// Specifies whether to install the Open Policy Agent (OPA) plug-in. Valid values:
	//
	// 	- `true`: installs the OPA plug-in.
	//
	// 	- `false`: does not install the OPA plug-in.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	OpenAgentPolicy *bool `json:"OpenAgentPolicy,omitempty" xml:"OpenAgentPolicy,omitempty"`
	// The policy for accessing external services. Valid values:
	//
	// 	- `ALLOW_ANY`: allows access to all external services.
	//
	// 	- `REGISTRY_ONLY`: allows access to only the external services that are defined in the ServiceEntry of the ASM instance.
	//
	// example:
	//
	// ALLOW_ANY
	OutboundTrafficPolicy *string `json:"OutboundTrafficPolicy,omitempty" xml:"OutboundTrafficPolicy,omitempty"`
	// Specifies whether to support HTTP/3.
	//
	// example:
	//
	// false
	PilotEnableQuicListeners *bool `json:"PilotEnableQuicListeners,omitempty" xml:"PilotEnableQuicListeners,omitempty"`
	// The endpoint of Prometheus monitoring. If you use ARMS Prometheus, set this parameter to the endpoint of Prometheus provided by ARMS.
	//
	// example:
	//
	// http://prometheus:9090
	PrometheusUrl *string `json:"PrometheusUrl,omitempty" xml:"PrometheusUrl,omitempty"`
	// The maximum number of CPU cores that are available to the istio-init container.
	//
	// example:
	//
	// 2000m
	ProxyInitCPUResourceLimit *string `json:"ProxyInitCPUResourceLimit,omitempty" xml:"ProxyInitCPUResourceLimit,omitempty"`
	// The number of CPU cores that are requested by the istio-init container.
	//
	// example:
	//
	// 10m
	ProxyInitCPUResourceRequest *string `json:"ProxyInitCPUResourceRequest,omitempty" xml:"ProxyInitCPUResourceRequest,omitempty"`
	// The maximum size of the memory that is available to the istio-init container.
	//
	// example:
	//
	// 1024Mi
	ProxyInitMemoryResourceLimit *string `json:"ProxyInitMemoryResourceLimit,omitempty" xml:"ProxyInitMemoryResourceLimit,omitempty"`
	// The size of the memory that is requested by the istio-init container.
	//
	// example:
	//
	// 10Mi
	ProxyInitMemoryResourceRequest *string `json:"ProxyInitMemoryResourceRequest,omitempty" xml:"ProxyInitMemoryResourceRequest,omitempty"`
	// The maximum number of CPU cores that are available to the sidecar proxy container.
	//
	// example:
	//
	// 2000m
	ProxyLimitCPU *string `json:"ProxyLimitCPU,omitempty" xml:"ProxyLimitCPU,omitempty"`
	// The maximum size of the memory that is available to the sidecar proxy container.
	//
	// example:
	//
	// 1024Mi
	ProxyLimitMemory *string `json:"ProxyLimitMemory,omitempty" xml:"ProxyLimitMemory,omitempty"`
	// The number of CPU cores that are requested by the sidecar proxy container.
	//
	// example:
	//
	// 100m
	ProxyRequestCPU *string `json:"ProxyRequestCPU,omitempty" xml:"ProxyRequestCPU,omitempty"`
	// The size of the memory that is requested by the sidecar proxy container.
	//
	// example:
	//
	// 128Mi
	ProxyRequestMemory *string `json:"ProxyRequestMemory,omitempty" xml:"ProxyRequestMemory,omitempty"`
	// The mode in which the sidecar proxy intercepts inbound traffic. Valid values:
	//
	// 	- `REDIRECT`: The sidecar proxy intercepts inbound traffic in the REDIRECT mode.
	//
	// 	- `TPROXY`: The sidecar proxy intercepts inbound traffic in the TPROXY mode.
	//
	// example:
	//
	// {"inclusionRegexps":".*adaptive_concurrency.*"}
	ProxyStatsMatcher *string `json:"ProxyStatsMatcher,omitempty" xml:"ProxyStatsMatcher,omitempty"`
	// Specifies whether to enable Redis Filter. Valid values:
	//
	// 	- `true`: enables Redis Filter.
	//
	// 	- `false`: disables Redis Filter.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	RedisFilterEnabled *bool `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
	// Specifies whether to enable SMC optimization.
	//
	// example:
	//
	// false
	SMCEnabled *bool `json:"SMCEnabled,omitempty" xml:"SMCEnabled,omitempty"`
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The maximum number of CPU cores that are available to the pod where a sidecar proxy injector resides.
	//
	// example:
	//
	// 4000m
	SidecarInjectorLimitCPU *string `json:"SidecarInjectorLimitCPU,omitempty" xml:"SidecarInjectorLimitCPU,omitempty"`
	// The maximum size of the memory that is available to the pod where a sidecar proxy injector resides.
	//
	// example:
	//
	// 2048Mi
	SidecarInjectorLimitMemory *string `json:"SidecarInjectorLimitMemory,omitempty" xml:"SidecarInjectorLimitMemory,omitempty"`
	// The number of CPU cores that are requested by the pod where a sidecar proxy injector resides.
	//
	// example:
	//
	// 1000m
	SidecarInjectorRequestCPU *string `json:"SidecarInjectorRequestCPU,omitempty" xml:"SidecarInjectorRequestCPU,omitempty"`
	// The size of the memory that is requested by the pod where a sidecar proxy injector resides.
	//
	// example:
	//
	// 512Mi
	SidecarInjectorRequestMemory *string `json:"SidecarInjectorRequestMemory,omitempty" xml:"SidecarInjectorRequestMemory,omitempty"`
	// Other configurations of automatic sidecar proxy injection, in the YAML format.
	//
	// example:
	//
	// {"injectedAnnotations":{"test/istio-init":"runtime/default2","test/istio-proxy":"runtime/default"},"replicaCount":2,"nodeSelector":{"beta.kubernetes.io/os":"linux"}}
	SidecarInjectorWebhookAsYaml *string `json:"SidecarInjectorWebhookAsYaml,omitempty" xml:"SidecarInjectorWebhookAsYaml,omitempty"`
	// Specifies whether to enable Prometheus monitoring. We recommend that you enable [ARMS Prometheus](https://arms.console.aliyun.com/). Valid values:
	//
	// 	- `true`: enables Prometheus monitoring.
	//
	// 	- `false`: disables Prometheus monitoring.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Telemetry *bool `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	// The maximum period of time that the sidecar proxy waits for requests to be processed before the proxy is stopped. For example, if you want to specify a period of 5 seconds, set this parameter to 5s.
	//
	// example:
	//
	// 5s
	TerminationDrainDuration *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
	// Specifies whether to enable Thrift Filter. Valid values:
	//
	// 	- `true`: enables Thrift Filter.
	//
	// 	- `false`: disables Thrift Filter.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	ThriftFilterEnabled *bool `json:"ThriftFilterEnabled,omitempty" xml:"ThriftFilterEnabled,omitempty"`
	// The custom tag of Tracing Analysis. Specify this parameter in the JSON format.
	//
	//     {
	//
	//         "name1": CustomTag,
	//
	//         "name2": CustomTag
	//
	//     }
	//
	// Tag key: literal, header, or environment.
	//
	//     {
	//
	//         "literal": {
	//
	//             "value": "Fixed value"
	//
	//         }
	//
	//         "header": {
	//
	//             "name": "Header name"
	//
	//             "defaultValue": "Default value that is used if the specified header does not exist"
	//
	//         }
	//
	//         "environment": {
	//
	//             "name": "Environment variable name"
	//
	//             "defaultValue": "Default value that is used if the specified environment variable does not exist"
	//
	//         }
	//
	//     }
	//
	// example:
	//
	// {"mytag": {"literal":{"value":"test"}}}
	TraceCustomTags *string `json:"TraceCustomTags,omitempty" xml:"TraceCustomTags,omitempty"`
	// The maximum length of the request path contained in the HttpUrl span tag. Default value: `256`.
	//
	// example:
	//
	// 256
	TraceMaxPathTagLength *string `json:"TraceMaxPathTagLength,omitempty" xml:"TraceMaxPathTagLength,omitempty"`
	// The sampling percentage of Tracing Analysis.
	//
	// example:
	//
	// 100
	TraceSampling *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
	// Specifies whether to enable Managed Service for OpenTelemetry. (Before you enable [Managed Service for OpenTelemetry](https://tracing-analysis.console.aliyun.com/), make sure that you have activated it.) Valid values:
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
	// The maximum size of the memory that is available to the proxy service that exports Tracing Analysis data. For example, `1Mi` indicates 1 MB.
	//
	// example:
	//
	// 1000Mi
	TracingOnExtZipkinLimitCPU *string `json:"TracingOnExtZipkinLimitCPU,omitempty" xml:"TracingOnExtZipkinLimitCPU,omitempty"`
	// The retention period for the access logs of the ingress gateway. Unit: day. The logs are collected by using Log Service. For example, `30` indicates 30 days.
	//
	// example:
	//
	// 1024Mi
	TracingOnExtZipkinLimitMemory *string `json:"TracingOnExtZipkinLimitMemory,omitempty" xml:"TracingOnExtZipkinLimitMemory,omitempty"`
	// The number of replicas that are available to the proxy service that exports Managed Service for OpenTelemetry data.
	//
	// example:
	//
	// 2
	TracingOnExtZipkinReplicaCount *string `json:"TracingOnExtZipkinReplicaCount,omitempty" xml:"TracingOnExtZipkinReplicaCount,omitempty"`
	// The minimum size of the memory requested by the proxy service that exports Tracing Analysis data. For example, `1Mi` indicates 1 MB.
	//
	// example:
	//
	// 200m
	TracingOnExtZipkinRequestCPU *string `json:"TracingOnExtZipkinRequestCPU,omitempty" xml:"TracingOnExtZipkinRequestCPU,omitempty"`
	// The maximum number of CPU cores that are available to the proxy service that exports Tracing Analysis data. For example, `1000m` indicates one CPU core.
	//
	// example:
	//
	// 200Mi
	TracingOnExtZipkinRequestMemory *string `json:"TracingOnExtZipkinRequestMemory,omitempty" xml:"TracingOnExtZipkinRequestMemory,omitempty"`
	// Specifies whether to enable WebAssembly Filter. Valid values:
	//
	// 	- `true`: enables WebAssembly Filter.
	//
	// 	- `false`: disables WebAssembly Filter.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	WebAssemblyFilterEnabled *bool `json:"WebAssemblyFilterEnabled,omitempty" xml:"WebAssemblyFilterEnabled,omitempty"`
}

func (s UpdateMeshFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeshFeatureRequest) GetAccessLogEnabled() *bool {
	return s.AccessLogEnabled
}

func (s *UpdateMeshFeatureRequest) GetAccessLogFile() *string {
	return s.AccessLogFile
}

func (s *UpdateMeshFeatureRequest) GetAccessLogFormat() *string {
	return s.AccessLogFormat
}

func (s *UpdateMeshFeatureRequest) GetAccessLogGatewayEnabled() *bool {
	return s.AccessLogGatewayEnabled
}

func (s *UpdateMeshFeatureRequest) GetAccessLogGatewayLifecycle() *int32 {
	return s.AccessLogGatewayLifecycle
}

func (s *UpdateMeshFeatureRequest) GetAccessLogProject() *string {
	return s.AccessLogProject
}

func (s *UpdateMeshFeatureRequest) GetAccessLogServiceEnabled() *bool {
	return s.AccessLogServiceEnabled
}

func (s *UpdateMeshFeatureRequest) GetAccessLogServiceHost() *string {
	return s.AccessLogServiceHost
}

func (s *UpdateMeshFeatureRequest) GetAccessLogServicePort() *int32 {
	return s.AccessLogServicePort
}

func (s *UpdateMeshFeatureRequest) GetAccessLogSidecarEnabled() *bool {
	return s.AccessLogSidecarEnabled
}

func (s *UpdateMeshFeatureRequest) GetAccessLogSidecarLifecycle() *int32 {
	return s.AccessLogSidecarLifecycle
}

func (s *UpdateMeshFeatureRequest) GetAuditProject() *string {
	return s.AuditProject
}

func (s *UpdateMeshFeatureRequest) GetAutoInjectionPolicyEnabled() *bool {
	return s.AutoInjectionPolicyEnabled
}

func (s *UpdateMeshFeatureRequest) GetCRAggregationEnabled() *bool {
	return s.CRAggregationEnabled
}

func (s *UpdateMeshFeatureRequest) GetCertChain() *string {
	return s.CertChain
}

func (s *UpdateMeshFeatureRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *UpdateMeshFeatureRequest) GetCniEnabled() *bool {
	return s.CniEnabled
}

func (s *UpdateMeshFeatureRequest) GetCniExcludeNamespaces() *string {
	return s.CniExcludeNamespaces
}

func (s *UpdateMeshFeatureRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *UpdateMeshFeatureRequest) GetConfigSourceEnabled() *bool {
	return s.ConfigSourceEnabled
}

func (s *UpdateMeshFeatureRequest) GetConfigSourceNacosID() *string {
	return s.ConfigSourceNacosID
}

func (s *UpdateMeshFeatureRequest) GetCustomizedPrometheus() *bool {
	return s.CustomizedPrometheus
}

func (s *UpdateMeshFeatureRequest) GetCustomizedZipkin() *bool {
	return s.CustomizedZipkin
}

func (s *UpdateMeshFeatureRequest) GetDNSProxyingEnabled() *bool {
	return s.DNSProxyingEnabled
}

func (s *UpdateMeshFeatureRequest) GetDefaultComponentsScheduleConfig() *string {
	return s.DefaultComponentsScheduleConfig
}

func (s *UpdateMeshFeatureRequest) GetDiscoverySelectors() *string {
	return s.DiscoverySelectors
}

func (s *UpdateMeshFeatureRequest) GetDubboFilterEnabled() *bool {
	return s.DubboFilterEnabled
}

func (s *UpdateMeshFeatureRequest) GetEnableAudit() *bool {
	return s.EnableAudit
}

func (s *UpdateMeshFeatureRequest) GetEnableAutoDiagnosis() *bool {
	return s.EnableAutoDiagnosis
}

func (s *UpdateMeshFeatureRequest) GetEnableBootstrapXdsAgent() *bool {
	return s.EnableBootstrapXdsAgent
}

func (s *UpdateMeshFeatureRequest) GetEnableCRHistory() *bool {
	return s.EnableCRHistory
}

func (s *UpdateMeshFeatureRequest) GetEnableNamespacesByDefault() *bool {
	return s.EnableNamespacesByDefault
}

func (s *UpdateMeshFeatureRequest) GetEnableSDSServer() *bool {
	return s.EnableSDSServer
}

func (s *UpdateMeshFeatureRequest) GetExcludeIPRanges() *string {
	return s.ExcludeIPRanges
}

func (s *UpdateMeshFeatureRequest) GetExcludeInboundPorts() *string {
	return s.ExcludeInboundPorts
}

func (s *UpdateMeshFeatureRequest) GetExcludeOutboundPorts() *string {
	return s.ExcludeOutboundPorts
}

func (s *UpdateMeshFeatureRequest) GetExistingCaCert() *string {
	return s.ExistingCaCert
}

func (s *UpdateMeshFeatureRequest) GetExistingCaKey() *string {
	return s.ExistingCaKey
}

func (s *UpdateMeshFeatureRequest) GetExistingRootCaCert() *string {
	return s.ExistingRootCaCert
}

func (s *UpdateMeshFeatureRequest) GetFilterGatewayClusterConfig() *bool {
	return s.FilterGatewayClusterConfig
}

func (s *UpdateMeshFeatureRequest) GetGatewayAPIEnabled() *bool {
	return s.GatewayAPIEnabled
}

func (s *UpdateMeshFeatureRequest) GetHoldApplicationUntilProxyStarts() *bool {
	return s.HoldApplicationUntilProxyStarts
}

func (s *UpdateMeshFeatureRequest) GetHttp10Enabled() *bool {
	return s.Http10Enabled
}

func (s *UpdateMeshFeatureRequest) GetIncludeIPRanges() *string {
	return s.IncludeIPRanges
}

func (s *UpdateMeshFeatureRequest) GetIncludeInboundPorts() *string {
	return s.IncludeInboundPorts
}

func (s *UpdateMeshFeatureRequest) GetIncludeOutboundPorts() *string {
	return s.IncludeOutboundPorts
}

func (s *UpdateMeshFeatureRequest) GetIntegrateKiali() *bool {
	return s.IntegrateKiali
}

func (s *UpdateMeshFeatureRequest) GetInterceptionMode() *string {
	return s.InterceptionMode
}

func (s *UpdateMeshFeatureRequest) GetKialiArmsAuthTokens() *string {
	return s.KialiArmsAuthTokens
}

func (s *UpdateMeshFeatureRequest) GetKialiEnabled() *bool {
	return s.KialiEnabled
}

func (s *UpdateMeshFeatureRequest) GetKialiServiceAnnotations() *string {
	return s.KialiServiceAnnotations
}

func (s *UpdateMeshFeatureRequest) GetLabelsForOffloadedWorkloads() *string {
	return s.LabelsForOffloadedWorkloads
}

func (s *UpdateMeshFeatureRequest) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *UpdateMeshFeatureRequest) GetLocalityLBConf() *string {
	return s.LocalityLBConf
}

func (s *UpdateMeshFeatureRequest) GetLocalityLoadBalancing() *bool {
	return s.LocalityLoadBalancing
}

func (s *UpdateMeshFeatureRequest) GetLogLevel() *string {
	return s.LogLevel
}

func (s *UpdateMeshFeatureRequest) GetMSEEnabled() *bool {
	return s.MSEEnabled
}

func (s *UpdateMeshFeatureRequest) GetMultiBufferEnabled() *bool {
	return s.MultiBufferEnabled
}

func (s *UpdateMeshFeatureRequest) GetMultiBufferPollDelay() *string {
	return s.MultiBufferPollDelay
}

func (s *UpdateMeshFeatureRequest) GetMysqlFilterEnabled() *bool {
	return s.MysqlFilterEnabled
}

func (s *UpdateMeshFeatureRequest) GetNFDEnabled() *bool {
	return s.NFDEnabled
}

func (s *UpdateMeshFeatureRequest) GetNFDLabelPruned() *bool {
	return s.NFDLabelPruned
}

func (s *UpdateMeshFeatureRequest) GetOPAInjectorCPULimit() *string {
	return s.OPAInjectorCPULimit
}

func (s *UpdateMeshFeatureRequest) GetOPAInjectorCPURequirement() *string {
	return s.OPAInjectorCPURequirement
}

func (s *UpdateMeshFeatureRequest) GetOPAInjectorMemoryLimit() *string {
	return s.OPAInjectorMemoryLimit
}

func (s *UpdateMeshFeatureRequest) GetOPAInjectorMemoryRequirement() *string {
	return s.OPAInjectorMemoryRequirement
}

func (s *UpdateMeshFeatureRequest) GetOPALimitCPU() *string {
	return s.OPALimitCPU
}

func (s *UpdateMeshFeatureRequest) GetOPALimitMemory() *string {
	return s.OPALimitMemory
}

func (s *UpdateMeshFeatureRequest) GetOPALogLevel() *string {
	return s.OPALogLevel
}

func (s *UpdateMeshFeatureRequest) GetOPARequestCPU() *string {
	return s.OPARequestCPU
}

func (s *UpdateMeshFeatureRequest) GetOPARequestMemory() *string {
	return s.OPARequestMemory
}

func (s *UpdateMeshFeatureRequest) GetOPAScopeInjected() *bool {
	return s.OPAScopeInjected
}

func (s *UpdateMeshFeatureRequest) GetOpaEnabled() *bool {
	return s.OpaEnabled
}

func (s *UpdateMeshFeatureRequest) GetOpenAgentPolicy() *bool {
	return s.OpenAgentPolicy
}

func (s *UpdateMeshFeatureRequest) GetOutboundTrafficPolicy() *string {
	return s.OutboundTrafficPolicy
}

func (s *UpdateMeshFeatureRequest) GetPilotEnableQuicListeners() *bool {
	return s.PilotEnableQuicListeners
}

func (s *UpdateMeshFeatureRequest) GetPrometheusUrl() *string {
	return s.PrometheusUrl
}

func (s *UpdateMeshFeatureRequest) GetProxyInitCPUResourceLimit() *string {
	return s.ProxyInitCPUResourceLimit
}

func (s *UpdateMeshFeatureRequest) GetProxyInitCPUResourceRequest() *string {
	return s.ProxyInitCPUResourceRequest
}

func (s *UpdateMeshFeatureRequest) GetProxyInitMemoryResourceLimit() *string {
	return s.ProxyInitMemoryResourceLimit
}

func (s *UpdateMeshFeatureRequest) GetProxyInitMemoryResourceRequest() *string {
	return s.ProxyInitMemoryResourceRequest
}

func (s *UpdateMeshFeatureRequest) GetProxyLimitCPU() *string {
	return s.ProxyLimitCPU
}

func (s *UpdateMeshFeatureRequest) GetProxyLimitMemory() *string {
	return s.ProxyLimitMemory
}

func (s *UpdateMeshFeatureRequest) GetProxyRequestCPU() *string {
	return s.ProxyRequestCPU
}

func (s *UpdateMeshFeatureRequest) GetProxyRequestMemory() *string {
	return s.ProxyRequestMemory
}

func (s *UpdateMeshFeatureRequest) GetProxyStatsMatcher() *string {
	return s.ProxyStatsMatcher
}

func (s *UpdateMeshFeatureRequest) GetRedisFilterEnabled() *bool {
	return s.RedisFilterEnabled
}

func (s *UpdateMeshFeatureRequest) GetSMCEnabled() *bool {
	return s.SMCEnabled
}

func (s *UpdateMeshFeatureRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateMeshFeatureRequest) GetSidecarInjectorLimitCPU() *string {
	return s.SidecarInjectorLimitCPU
}

func (s *UpdateMeshFeatureRequest) GetSidecarInjectorLimitMemory() *string {
	return s.SidecarInjectorLimitMemory
}

func (s *UpdateMeshFeatureRequest) GetSidecarInjectorRequestCPU() *string {
	return s.SidecarInjectorRequestCPU
}

func (s *UpdateMeshFeatureRequest) GetSidecarInjectorRequestMemory() *string {
	return s.SidecarInjectorRequestMemory
}

func (s *UpdateMeshFeatureRequest) GetSidecarInjectorWebhookAsYaml() *string {
	return s.SidecarInjectorWebhookAsYaml
}

func (s *UpdateMeshFeatureRequest) GetTelemetry() *bool {
	return s.Telemetry
}

func (s *UpdateMeshFeatureRequest) GetTerminationDrainDuration() *string {
	return s.TerminationDrainDuration
}

func (s *UpdateMeshFeatureRequest) GetThriftFilterEnabled() *bool {
	return s.ThriftFilterEnabled
}

func (s *UpdateMeshFeatureRequest) GetTraceCustomTags() *string {
	return s.TraceCustomTags
}

func (s *UpdateMeshFeatureRequest) GetTraceMaxPathTagLength() *string {
	return s.TraceMaxPathTagLength
}

func (s *UpdateMeshFeatureRequest) GetTraceSampling() *float32 {
	return s.TraceSampling
}

func (s *UpdateMeshFeatureRequest) GetTracing() *bool {
	return s.Tracing
}

func (s *UpdateMeshFeatureRequest) GetTracingOnExtZipkinLimitCPU() *string {
	return s.TracingOnExtZipkinLimitCPU
}

func (s *UpdateMeshFeatureRequest) GetTracingOnExtZipkinLimitMemory() *string {
	return s.TracingOnExtZipkinLimitMemory
}

func (s *UpdateMeshFeatureRequest) GetTracingOnExtZipkinReplicaCount() *string {
	return s.TracingOnExtZipkinReplicaCount
}

func (s *UpdateMeshFeatureRequest) GetTracingOnExtZipkinRequestCPU() *string {
	return s.TracingOnExtZipkinRequestCPU
}

func (s *UpdateMeshFeatureRequest) GetTracingOnExtZipkinRequestMemory() *string {
	return s.TracingOnExtZipkinRequestMemory
}

func (s *UpdateMeshFeatureRequest) GetWebAssemblyFilterEnabled() *bool {
	return s.WebAssemblyFilterEnabled
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

func (s *UpdateMeshFeatureRequest) SetCertChain(v string) *UpdateMeshFeatureRequest {
	s.CertChain = &v
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

func (s *UpdateMeshFeatureRequest) SetExistingCaCert(v string) *UpdateMeshFeatureRequest {
	s.ExistingCaCert = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetExistingCaKey(v string) *UpdateMeshFeatureRequest {
	s.ExistingCaKey = &v
	return s
}

func (s *UpdateMeshFeatureRequest) SetExistingRootCaCert(v string) *UpdateMeshFeatureRequest {
	s.ExistingRootCaCert = &v
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

func (s *UpdateMeshFeatureRequest) SetLabelsForOffloadedWorkloads(v string) *UpdateMeshFeatureRequest {
	s.LabelsForOffloadedWorkloads = &v
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

func (s *UpdateMeshFeatureRequest) SetPilotEnableQuicListeners(v bool) *UpdateMeshFeatureRequest {
	s.PilotEnableQuicListeners = &v
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

func (s *UpdateMeshFeatureRequest) SetSMCEnabled(v bool) *UpdateMeshFeatureRequest {
	s.SMCEnabled = &v
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

func (s *UpdateMeshFeatureRequest) SetTracingOnExtZipkinReplicaCount(v string) *UpdateMeshFeatureRequest {
	s.TracingOnExtZipkinReplicaCount = &v
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

func (s *UpdateMeshFeatureRequest) Validate() error {
	return dara.Validate(s)
}
