// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceScopeSidecarConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *UpdateNamespaceScopeSidecarConfigRequest
	GetConcurrency() *int32
	SetEnableCoreDump(v bool) *UpdateNamespaceScopeSidecarConfigRequest
	GetEnableCoreDump() *bool
	SetExcludeIPRanges(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetExcludeIPRanges() *string
	SetExcludeInboundPorts(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetExcludeInboundPorts() *string
	SetExcludeOutboundPorts(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetExcludeOutboundPorts() *string
	SetHoldApplicationUntilProxyStarts(v bool) *UpdateNamespaceScopeSidecarConfigRequest
	GetHoldApplicationUntilProxyStarts() *bool
	SetIncludeIPRanges(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetIncludeIPRanges() *string
	SetIncludeInboundPorts(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetIncludeInboundPorts() *string
	SetIncludeOutboundPorts(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetIncludeOutboundPorts() *string
	SetInterceptionMode(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetInterceptionMode() *string
	SetIstioDNSProxyEnabled(v bool) *UpdateNamespaceScopeSidecarConfigRequest
	GetIstioDNSProxyEnabled() *bool
	SetLifecycle(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetLifecycle() *string
	SetLogLevel(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetLogLevel() *string
	SetNamespace(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetNamespace() *string
	SetPostStart(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetPostStart() *string
	SetPreStop(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetPreStop() *string
	SetPrivileged(v bool) *UpdateNamespaceScopeSidecarConfigRequest
	GetPrivileged() *bool
	SetProxyInitAckSloCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyInitAckSloCPUResourceLimit() *string
	SetProxyInitAckSloCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyInitAckSloCPUResourceRequest() *string
	SetProxyInitAckSloMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyInitAckSloMemoryResourceLimit() *string
	SetProxyInitAckSloMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyInitAckSloMemoryResourceRequest() *string
	SetProxyInitCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyInitCPUResourceLimit() *string
	SetProxyInitCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyInitCPUResourceRequest() *string
	SetProxyInitMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyInitMemoryResourceLimit() *string
	SetProxyInitMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyInitMemoryResourceRequest() *string
	SetProxyMetadata(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyMetadata() *string
	SetProxyStatsMatcher(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetProxyStatsMatcher() *string
	SetReadinessFailureThreshold(v int32) *UpdateNamespaceScopeSidecarConfigRequest
	GetReadinessFailureThreshold() *int32
	SetReadinessInitialDelaySeconds(v int32) *UpdateNamespaceScopeSidecarConfigRequest
	GetReadinessInitialDelaySeconds() *int32
	SetReadinessPeriodSeconds(v int32) *UpdateNamespaceScopeSidecarConfigRequest
	GetReadinessPeriodSeconds() *int32
	SetRuntimeValues(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetRuntimeValues() *string
	SetSMCEnabled(v bool) *UpdateNamespaceScopeSidecarConfigRequest
	GetSMCEnabled() *bool
	SetScaledSidecarResource(v *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) *UpdateNamespaceScopeSidecarConfigRequest
	GetScaledSidecarResource() *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource
	SetServiceMeshId(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetServiceMeshId() *string
	SetSidecarProxyAckSloCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetSidecarProxyAckSloCPUResourceLimit() *string
	SetSidecarProxyAckSloCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetSidecarProxyAckSloCPUResourceRequest() *string
	SetSidecarProxyAckSloMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetSidecarProxyAckSloMemoryResourceLimit() *string
	SetSidecarProxyAckSloMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetSidecarProxyAckSloMemoryResourceRequest() *string
	SetSidecarProxyCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetSidecarProxyCPUResourceLimit() *string
	SetSidecarProxyCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetSidecarProxyCPUResourceRequest() *string
	SetSidecarProxyMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetSidecarProxyMemoryResourceLimit() *string
	SetSidecarProxyMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetSidecarProxyMemoryResourceRequest() *string
	SetTerminationDrainDuration(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetTerminationDrainDuration() *string
	SetTracing(v string) *UpdateNamespaceScopeSidecarConfigRequest
	GetTracing() *string
}

type UpdateNamespaceScopeSidecarConfigRequest struct {
	// The number of worker threads to run in Istio Proxy.
	//
	// example:
	//
	// 2
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// Specifies whether to enable the core dump feature for the sidecar proxy containers. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	EnableCoreDump *bool `json:"EnableCoreDump,omitempty" xml:"EnableCoreDump,omitempty"`
	// The range of IP addresses that are allowed to access external services. (`global.proxy.excludelPRanges`)
	//
	// example:
	//
	// 172.16.0.0/12
	ExcludeIPRanges *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	// The port that the inbound traffic of the sidecar proxy does not pass through.
	//
	// example:
	//
	// 82
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The port that the outbound traffic of the sidecar proxy does not pass through.
	//
	// example:
	//
	// 81
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// Specifies whether applications can be started only after Istio Proxy starts. Valid values:
	//
	// 	- `true`: Applications can be started only after Istio Proxy starts.
	//
	// 	- `false`: Applications can be started before Istio Proxy starts.
	//
	// example:
	//
	// true
	HoldApplicationUntilProxyStarts *bool `json:"HoldApplicationUntilProxyStarts,omitempty" xml:"HoldApplicationUntilProxyStarts,omitempty"`
	// The range of IP addresses that are denied to access external services. (`global.proxy.includelPRanges`)
	//
	// example:
	//
	// *
	IncludeIPRanges *string `json:"IncludeIPRanges,omitempty" xml:"IncludeIPRanges,omitempty"`
	// The port that the inbound traffic of the sidecar proxy passes through.
	//
	// example:
	//
	// 83
	IncludeInboundPorts *string `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	// The port that the outbound traffic of the sidecar proxy passes through.
	//
	// example:
	//
	// 84
	IncludeOutboundPorts *string `json:"IncludeOutboundPorts,omitempty" xml:"IncludeOutboundPorts,omitempty"`
	// The mode in which the sidecar proxy intercepts inbound traffic. Valid values:
	//
	// 	- `REDIRECT`: The sidecar proxy intercepts inbound traffic in the REDIRECT mode.
	//
	// 	- `TPROXY`: The sidecar proxy intercepts inbound traffic in the TPROXY mode.
	//
	// example:
	//
	// TPROXY
	InterceptionMode *string `json:"InterceptionMode,omitempty" xml:"InterceptionMode,omitempty"`
	// Specifies whether to enable the Domain Name System (DNS) proxy feature. Valid values:
	//
	// 	- `true`: The DNS proxy feature is enabled.
	//
	// 	- `false`: The DNS proxy feature is disabled.
	//
	// example:
	//
	// true
	IstioDNSProxyEnabled *bool `json:"IstioDNSProxyEnabled,omitempty" xml:"IstioDNSProxyEnabled,omitempty"`
	// The lifecycle of the sidecar proxy.
	//
	// example:
	//
	// {"postStart":{"exec":{"command":["pilot-agent","wait"]}},"preStop":{"exec":{"command":["/bin/sh","-c","sleep 15"]}}}
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The log level. Valid values: `info`, `debug`, `tracing`, and `error`.
	//
	// example:
	//
	// info
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The namespace for which you want to update the sidecar proxy configurations.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The post-start parameters of Istio Proxy.
	//
	// example:
	//
	// {"exec":{"command":["pilot-agent","wait"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The pre-close parameters of Istio Proxy.
	//
	// example:
	//
	// {"exec":{"command":["/bin/sh","-c","sleep 15"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// Specifies whether to enable the privileged mode in the security context of the sidecar proxy containers. Valid values:
	//
	// 	- `true`: enables the privileged mode. This means that the sidecar proxy containers run in privileged mode.
	//
	// 	- `false`: does not enable the privileged mode.
	//
	// example:
	//
	// false
	Privileged *bool `json:"Privileged,omitempty" xml:"Privileged,omitempty"`
	// The maximum number of reclaimed CPU cores provided by Container Service for Kubernetes (ACK) that are available to the istio-init container. Reclaimed resources, including CPU cores and memory, are resources that can be dynamically overcommitted. This configuration item is used to set the maximum number of CPU cores that are available to the istio-init container in a pod labeled with `koordinator.sh/qosClass`. Unit: millicore.
	//
	// example:
	//
	// 2000
	ProxyInitAckSloCPUResourceLimit *string `json:"ProxyInitAckSloCPUResourceLimit,omitempty" xml:"ProxyInitAckSloCPUResourceLimit,omitempty"`
	// The minimum number of reclaimed CPU cores provided by ACK that the istio-init container requires at runtime. This configuration item is used to set the minimum number of reclaimed CPU cores for the istio-init container in a pod labeled with `koordinator.sh/qosClass`. Unit: millicore.
	//
	// example:
	//
	// 100
	ProxyInitAckSloCPUResourceRequest *string `json:"ProxyInitAckSloCPUResourceRequest,omitempty" xml:"ProxyInitAckSloCPUResourceRequest,omitempty"`
	// The maximum size of reclaimed memory resources provided by ACK that are available to the istio-init container. This configuration item is used to set the maximum size of memory that is available to the istio-init container in a pod labeled with `koordinator.sh/qosClass`.
	//
	// example:
	//
	// 2048Mi
	ProxyInitAckSloMemoryResourceLimit *string `json:"ProxyInitAckSloMemoryResourceLimit,omitempty" xml:"ProxyInitAckSloMemoryResourceLimit,omitempty"`
	// The minimum size of reclaimed memory provided by ACK that the istio-init container requires at runtime. This configuration item is used to set the minimum size of reclaimed memory for the istio-init container in a pod labeled with `koordinator.sh/qosClass`.
	//
	// example:
	//
	// 128Mi
	ProxyInitAckSloMemoryResourceRequest *string `json:"ProxyInitAckSloMemoryResourceRequest,omitempty" xml:"ProxyInitAckSloMemoryResourceRequest,omitempty"`
	// The maximum number of CPU cores that are available to the sidecar proxy init container.
	//
	// example:
	//
	// 2000 m
	ProxyInitCPUResourceLimit *string `json:"ProxyInitCPUResourceLimit,omitempty" xml:"ProxyInitCPUResourceLimit,omitempty"`
	// The minimum number of CPU cores that are requested by the sidecar proxy init container.
	//
	// example:
	//
	// 60 m
	ProxyInitCPUResourceRequest *string `json:"ProxyInitCPUResourceRequest,omitempty" xml:"ProxyInitCPUResourceRequest,omitempty"`
	// The maximum size of memory that is available to the sidecar proxy init container.
	//
	// example:
	//
	// 50 Mi
	ProxyInitMemoryResourceLimit *string `json:"ProxyInitMemoryResourceLimit,omitempty" xml:"ProxyInitMemoryResourceLimit,omitempty"`
	// The minimum size of memory that is requested by the sidecar proxy init container.
	//
	// example:
	//
	// 30 Mi
	ProxyInitMemoryResourceRequest *string `json:"ProxyInitMemoryResourceRequest,omitempty" xml:"ProxyInitMemoryResourceRequest,omitempty"`
	// The environment variables that are added to a sidecar proxy. The environment variables are represented as JSON objects. The keys and values in the JSON objects represent the keys and values added to the environment variables of the sidecar proxy.
	//
	// example:
	//
	// {"EXIT_ON_ZERO_ACTIVE_CONNECTIONS":"true"}
	ProxyMetadata *string `json:"ProxyMetadata,omitempty" xml:"ProxyMetadata,omitempty"`
	// The monitoring metrics for data collected by Envoy proxies. The value is in the JSON format.
	//
	// example:
	//
	// { "inclusionPrefixes": [ "cluster.outbound", "cluster_manager", "listener_manager", "server", "cluster.xds-grpc" ], "inclusionRegexps": [ "listener.*.downstream_cx_total", "listener.*.downstream_cx_active" ] }
	ProxyStatsMatcher *string `json:"ProxyStatsMatcher,omitempty" xml:"ProxyStatsMatcher,omitempty"`
	// The number of readiness check failures required before marking a sidecar proxy container as not ready.
	//
	// example:
	//
	// 5
	ReadinessFailureThreshold *int32 `json:"ReadinessFailureThreshold,omitempty" xml:"ReadinessFailureThreshold,omitempty"`
	// The amount of time to wait before the first readiness check of a sidecar proxy container is performed. Unit: seconds.
	//
	// example:
	//
	// 1
	ReadinessInitialDelaySeconds *int32 `json:"ReadinessInitialDelaySeconds,omitempty" xml:"ReadinessInitialDelaySeconds,omitempty"`
	// The interval between two readiness checks of a sidecar proxy container. Unit: seconds.
	//
	// example:
	//
	// 2
	ReadinessPeriodSeconds *int32 `json:"ReadinessPeriodSeconds,omitempty" xml:"ReadinessPeriodSeconds,omitempty"`
	// Indicates the runtime parameters of Envoy proxy processes in the sidecar proxy container. This parameter is a serialized JSON string. The keys and values of a JSON object are the keys and the values of a Envoy Runtime Parameter respectively.
	//
	// Valid values:
	//
	// 	- global_downstream_max_connections: indicates that the limits on the number of connections from downstream to Envoy.
	//
	// example:
	//
	// {"overload.global_downstream_max_connections":"65536"}
	RuntimeValues *string `json:"RuntimeValues,omitempty" xml:"RuntimeValues,omitempty"`
	// Specifies whether to enable Shared Memory Communications over Remote Direct Memory Access (SMC-R) optimization. The SMC-R optimization feature uses Alibaba Cloud Linux 3 and elastic remote direct memory access (eRDMA) network devices, which optimizes cross-node communication.
	//
	// example:
	//
	// false
	SMCEnabled            *bool                                                          `json:"SMCEnabled,omitempty" xml:"SMCEnabled,omitempty"`
	ScaledSidecarResource *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource `json:"ScaledSidecarResource,omitempty" xml:"ScaledSidecarResource,omitempty" type:"Struct"`
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca04bc38979214bf2882be79d39b4****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The maximum number of reclaimed CPU cores provided by ACK that are available to the sidecar proxy container. This configuration item is used to set the maximum number of CPU cores that are available to the sidecar proxy container in a pod labeled with `koordinator.sh/qosClass`. Unit: millicore.
	//
	// example:
	//
	// 2000
	SidecarProxyAckSloCPUResourceLimit *string `json:"SidecarProxyAckSloCPUResourceLimit,omitempty" xml:"SidecarProxyAckSloCPUResourceLimit,omitempty"`
	// The minimum number of reclaimed CPU cores provided by ACK that the sidecar proxy container requires at runtime. This configuration item is used to set the minimum number of reclaimed CPU cores for the sidecar proxy container in a pod labeled with `koordinator.sh/qosClass`. Unit: millicore.
	//
	// example:
	//
	// 100
	SidecarProxyAckSloCPUResourceRequest *string `json:"SidecarProxyAckSloCPUResourceRequest,omitempty" xml:"SidecarProxyAckSloCPUResourceRequest,omitempty"`
	// The maximum size of reclaimed memory resources provided by ACK that are available to the sidecar proxy container. This configuration item is used to set the maximum size of memory that is available to the sidecar proxy container in a pod labeled with `koordinator.sh/qosClass`.
	//
	// example:
	//
	// 2048Mi
	SidecarProxyAckSloMemoryResourceLimit *string `json:"SidecarProxyAckSloMemoryResourceLimit,omitempty" xml:"SidecarProxyAckSloMemoryResourceLimit,omitempty"`
	// The minimum size of reclaimed memory provided by ACK that the sidecar proxy container requires at runtime. This configuration item is used to set the minimum size of reclaimed memory for the sidecar proxy container in a pod labeled with `koordinator.sh/qosClass`.
	//
	// example:
	//
	// 128Mi
	SidecarProxyAckSloMemoryResourceRequest *string `json:"SidecarProxyAckSloMemoryResourceRequest,omitempty" xml:"SidecarProxyAckSloMemoryResourceRequest,omitempty"`
	// The maximum number of CPU cores that are available to the sidecar proxy container.
	//
	// example:
	//
	// 2000 m
	SidecarProxyCPUResourceLimit *string `json:"SidecarProxyCPUResourceLimit,omitempty" xml:"SidecarProxyCPUResourceLimit,omitempty"`
	// The minimum number of CPU cores that are requested by the sidecar proxy container.
	//
	// example:
	//
	// 60 m
	SidecarProxyCPUResourceRequest *string `json:"SidecarProxyCPUResourceRequest,omitempty" xml:"SidecarProxyCPUResourceRequest,omitempty"`
	// The maximum size of memory that is available to the sidecar proxy container.
	//
	// example:
	//
	// 50 Mi
	SidecarProxyMemoryResourceLimit *string `json:"SidecarProxyMemoryResourceLimit,omitempty" xml:"SidecarProxyMemoryResourceLimit,omitempty"`
	// The minimum size of memory that is requested by the sidecar proxy container.
	//
	// example:
	//
	// 30 Mi
	SidecarProxyMemoryResourceRequest *string `json:"SidecarProxyMemoryResourceRequest,omitempty" xml:"SidecarProxyMemoryResourceRequest,omitempty"`
	// The maximum period of time that the sidecar proxy waits for a request to end.
	//
	// example:
	//
	// 6s
	TerminationDrainDuration *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
	// The custom configurations of Tracing Analysis. The configurations must be serialized into JSON strings. The configurations contain the following parameters:
	//
	// 	- `sampling`: The sampling rate, which is of the DOUBLE type.
	//
	// 	- `custom_tags`: The custom tags added to reported spans, which are of the MAP type. The key of a tag is of the string type. The value of a tag is in the JSON format. A custom tag can belong to one of the following types:
	//
	//     	- `literal`: The tag value is a fixed value in the JSON format. This tag must contain the `value` field that specifies a literal. Example: `{"value":"test"}`.
	//
	//     	- `header`: The tag value is a request header in the JSON format. This tag must contain the `name` field and `defaultValue` field.The name field indicates the name of the request header. The defaultValue field indicates the default value that is used when no request header is available. Example: `{"name":"test","defaultValue":"test"}`.
	//
	//     	- `environment`: The tag value is an environment variable in the JSON format. This tag must contain the `name` field and `defaultValue` field. The name field indicates the name of the environment variable. The defaultValue field indicates the environment variable that is used when no environment variable is available. Example: `{"name":"test","defaultValue":"test"}`.
	//
	// example:
	//
	// {"sampling":99.8,"custom_tags":{"test":{"literal":{"value":"testnamespace"}}}}
	Tracing *string `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
}

func (s UpdateNamespaceScopeSidecarConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetEnableCoreDump() *bool {
	return s.EnableCoreDump
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetExcludeIPRanges() *string {
	return s.ExcludeIPRanges
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetExcludeInboundPorts() *string {
	return s.ExcludeInboundPorts
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetExcludeOutboundPorts() *string {
	return s.ExcludeOutboundPorts
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetHoldApplicationUntilProxyStarts() *bool {
	return s.HoldApplicationUntilProxyStarts
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetIncludeIPRanges() *string {
	return s.IncludeIPRanges
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetIncludeInboundPorts() *string {
	return s.IncludeInboundPorts
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetIncludeOutboundPorts() *string {
	return s.IncludeOutboundPorts
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetInterceptionMode() *string {
	return s.InterceptionMode
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetIstioDNSProxyEnabled() *bool {
	return s.IstioDNSProxyEnabled
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetLogLevel() *string {
	return s.LogLevel
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetPrivileged() *bool {
	return s.Privileged
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyInitAckSloCPUResourceLimit() *string {
	return s.ProxyInitAckSloCPUResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyInitAckSloCPUResourceRequest() *string {
	return s.ProxyInitAckSloCPUResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyInitAckSloMemoryResourceLimit() *string {
	return s.ProxyInitAckSloMemoryResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyInitAckSloMemoryResourceRequest() *string {
	return s.ProxyInitAckSloMemoryResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyInitCPUResourceLimit() *string {
	return s.ProxyInitCPUResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyInitCPUResourceRequest() *string {
	return s.ProxyInitCPUResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyInitMemoryResourceLimit() *string {
	return s.ProxyInitMemoryResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyInitMemoryResourceRequest() *string {
	return s.ProxyInitMemoryResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyMetadata() *string {
	return s.ProxyMetadata
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetProxyStatsMatcher() *string {
	return s.ProxyStatsMatcher
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetReadinessFailureThreshold() *int32 {
	return s.ReadinessFailureThreshold
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetReadinessInitialDelaySeconds() *int32 {
	return s.ReadinessInitialDelaySeconds
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetReadinessPeriodSeconds() *int32 {
	return s.ReadinessPeriodSeconds
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetRuntimeValues() *string {
	return s.RuntimeValues
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSMCEnabled() *bool {
	return s.SMCEnabled
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetScaledSidecarResource() *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource {
	return s.ScaledSidecarResource
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSidecarProxyAckSloCPUResourceLimit() *string {
	return s.SidecarProxyAckSloCPUResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSidecarProxyAckSloCPUResourceRequest() *string {
	return s.SidecarProxyAckSloCPUResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSidecarProxyAckSloMemoryResourceLimit() *string {
	return s.SidecarProxyAckSloMemoryResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSidecarProxyAckSloMemoryResourceRequest() *string {
	return s.SidecarProxyAckSloMemoryResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSidecarProxyCPUResourceLimit() *string {
	return s.SidecarProxyCPUResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSidecarProxyCPUResourceRequest() *string {
	return s.SidecarProxyCPUResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSidecarProxyMemoryResourceLimit() *string {
	return s.SidecarProxyMemoryResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetSidecarProxyMemoryResourceRequest() *string {
	return s.SidecarProxyMemoryResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetTerminationDrainDuration() *string {
	return s.TerminationDrainDuration
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) GetTracing() *string {
	return s.Tracing
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

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetRuntimeValues(v string) *UpdateNamespaceScopeSidecarConfigRequest {
	s.RuntimeValues = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetSMCEnabled(v bool) *UpdateNamespaceScopeSidecarConfigRequest {
	s.SMCEnabled = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequest) SetScaledSidecarResource(v *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) *UpdateNamespaceScopeSidecarConfigRequest {
	s.ScaledSidecarResource = v
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

func (s *UpdateNamespaceScopeSidecarConfigRequest) Validate() error {
	if s.ScaledSidecarResource != nil {
		if err := s.ScaledSidecarResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource struct {
	ContainerRef                *string `json:"ContainerRef,omitempty" xml:"ContainerRef,omitempty"`
	ResourceCalculationStrategy *string `json:"ResourceCalculationStrategy,omitempty" xml:"ResourceCalculationStrategy,omitempty"`
	ResourcePercentage          *int32  `json:"ResourcePercentage,omitempty" xml:"ResourcePercentage,omitempty"`
}

func (s UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) GetContainerRef() *string {
	return s.ContainerRef
}

func (s *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) GetResourceCalculationStrategy() *string {
	return s.ResourceCalculationStrategy
}

func (s *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) GetResourcePercentage() *int32 {
	return s.ResourcePercentage
}

func (s *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) SetContainerRef(v string) *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource {
	s.ContainerRef = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) SetResourceCalculationStrategy(v string) *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource {
	s.ResourceCalculationStrategy = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) SetResourcePercentage(v int32) *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource {
	s.ResourcePercentage = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigRequestScaledSidecarResource) Validate() error {
	return dara.Validate(s)
}
