// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceScopeSidecarConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetConcurrency() *int32
	SetEnableCoreDump(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetEnableCoreDump() *bool
	SetExcludeIPRanges(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetExcludeIPRanges() *string
	SetExcludeInboundPorts(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetExcludeInboundPorts() *string
	SetExcludeOutboundPorts(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetExcludeOutboundPorts() *string
	SetHoldApplicationUntilProxyStarts(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetHoldApplicationUntilProxyStarts() *bool
	SetIncludeIPRanges(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetIncludeIPRanges() *string
	SetIncludeInboundPorts(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetIncludeInboundPorts() *string
	SetIncludeOutboundPorts(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetIncludeOutboundPorts() *string
	SetInterceptionMode(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetInterceptionMode() *string
	SetIstioDNSProxyEnabled(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetIstioDNSProxyEnabled() *bool
	SetLifecycle(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetLifecycle() *string
	SetLogLevel(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetLogLevel() *string
	SetNamespace(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetNamespace() *string
	SetPostStart(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetPostStart() *string
	SetPreStop(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetPreStop() *string
	SetPrivileged(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetPrivileged() *bool
	SetProxyInitAckSloCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyInitAckSloCPUResourceLimit() *string
	SetProxyInitAckSloCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyInitAckSloCPUResourceRequest() *string
	SetProxyInitAckSloMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyInitAckSloMemoryResourceLimit() *string
	SetProxyInitAckSloMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyInitAckSloMemoryResourceRequest() *string
	SetProxyInitCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyInitCPUResourceLimit() *string
	SetProxyInitCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyInitCPUResourceRequest() *string
	SetProxyInitMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyInitMemoryResourceLimit() *string
	SetProxyInitMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyInitMemoryResourceRequest() *string
	SetProxyMetadata(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyMetadata() *string
	SetProxyStatsMatcher(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetProxyStatsMatcher() *string
	SetReadinessFailureThreshold(v int32) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetReadinessFailureThreshold() *int32
	SetReadinessInitialDelaySeconds(v int32) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetReadinessInitialDelaySeconds() *int32
	SetReadinessPeriodSeconds(v int32) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetReadinessPeriodSeconds() *int32
	SetRuntimeValues(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetRuntimeValues() *string
	SetSMCEnabled(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSMCEnabled() *bool
	SetScaledSidecarResourceShrink(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetScaledSidecarResourceShrink() *string
	SetServiceMeshId(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetServiceMeshId() *string
	SetSidecarProxyAckSloCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSidecarProxyAckSloCPUResourceLimit() *string
	SetSidecarProxyAckSloCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSidecarProxyAckSloCPUResourceRequest() *string
	SetSidecarProxyAckSloMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSidecarProxyAckSloMemoryResourceLimit() *string
	SetSidecarProxyAckSloMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSidecarProxyAckSloMemoryResourceRequest() *string
	SetSidecarProxyCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSidecarProxyCPUResourceLimit() *string
	SetSidecarProxyCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSidecarProxyCPUResourceRequest() *string
	SetSidecarProxyMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSidecarProxyMemoryResourceLimit() *string
	SetSidecarProxyMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetSidecarProxyMemoryResourceRequest() *string
	SetTerminationDrainDuration(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetTerminationDrainDuration() *string
	SetTracing(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest
	GetTracing() *string
}

type UpdateNamespaceScopeSidecarConfigShrinkRequest struct {
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
	SMCEnabled                  *bool   `json:"SMCEnabled,omitempty" xml:"SMCEnabled,omitempty"`
	ScaledSidecarResourceShrink *string `json:"ScaledSidecarResource,omitempty" xml:"ScaledSidecarResource,omitempty"`
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

func (s UpdateNamespaceScopeSidecarConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetEnableCoreDump() *bool {
	return s.EnableCoreDump
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetExcludeIPRanges() *string {
	return s.ExcludeIPRanges
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetExcludeInboundPorts() *string {
	return s.ExcludeInboundPorts
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetExcludeOutboundPorts() *string {
	return s.ExcludeOutboundPorts
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetHoldApplicationUntilProxyStarts() *bool {
	return s.HoldApplicationUntilProxyStarts
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetIncludeIPRanges() *string {
	return s.IncludeIPRanges
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetIncludeInboundPorts() *string {
	return s.IncludeInboundPorts
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetIncludeOutboundPorts() *string {
	return s.IncludeOutboundPorts
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetInterceptionMode() *string {
	return s.InterceptionMode
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetIstioDNSProxyEnabled() *bool {
	return s.IstioDNSProxyEnabled
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetLogLevel() *string {
	return s.LogLevel
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetPrivileged() *bool {
	return s.Privileged
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyInitAckSloCPUResourceLimit() *string {
	return s.ProxyInitAckSloCPUResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyInitAckSloCPUResourceRequest() *string {
	return s.ProxyInitAckSloCPUResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyInitAckSloMemoryResourceLimit() *string {
	return s.ProxyInitAckSloMemoryResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyInitAckSloMemoryResourceRequest() *string {
	return s.ProxyInitAckSloMemoryResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyInitCPUResourceLimit() *string {
	return s.ProxyInitCPUResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyInitCPUResourceRequest() *string {
	return s.ProxyInitCPUResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyInitMemoryResourceLimit() *string {
	return s.ProxyInitMemoryResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyInitMemoryResourceRequest() *string {
	return s.ProxyInitMemoryResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyMetadata() *string {
	return s.ProxyMetadata
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetProxyStatsMatcher() *string {
	return s.ProxyStatsMatcher
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetReadinessFailureThreshold() *int32 {
	return s.ReadinessFailureThreshold
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetReadinessInitialDelaySeconds() *int32 {
	return s.ReadinessInitialDelaySeconds
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetReadinessPeriodSeconds() *int32 {
	return s.ReadinessPeriodSeconds
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetRuntimeValues() *string {
	return s.RuntimeValues
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSMCEnabled() *bool {
	return s.SMCEnabled
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetScaledSidecarResourceShrink() *string {
	return s.ScaledSidecarResourceShrink
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSidecarProxyAckSloCPUResourceLimit() *string {
	return s.SidecarProxyAckSloCPUResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSidecarProxyAckSloCPUResourceRequest() *string {
	return s.SidecarProxyAckSloCPUResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSidecarProxyAckSloMemoryResourceLimit() *string {
	return s.SidecarProxyAckSloMemoryResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSidecarProxyAckSloMemoryResourceRequest() *string {
	return s.SidecarProxyAckSloMemoryResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSidecarProxyCPUResourceLimit() *string {
	return s.SidecarProxyCPUResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSidecarProxyCPUResourceRequest() *string {
	return s.SidecarProxyCPUResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSidecarProxyMemoryResourceLimit() *string {
	return s.SidecarProxyMemoryResourceLimit
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetSidecarProxyMemoryResourceRequest() *string {
	return s.SidecarProxyMemoryResourceRequest
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetTerminationDrainDuration() *string {
	return s.TerminationDrainDuration
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) GetTracing() *string {
	return s.Tracing
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetConcurrency(v int32) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.Concurrency = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetEnableCoreDump(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.EnableCoreDump = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetExcludeIPRanges(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ExcludeIPRanges = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetExcludeInboundPorts(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ExcludeInboundPorts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetExcludeOutboundPorts(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ExcludeOutboundPorts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetHoldApplicationUntilProxyStarts(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.HoldApplicationUntilProxyStarts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetIncludeIPRanges(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.IncludeIPRanges = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetIncludeInboundPorts(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.IncludeInboundPorts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetIncludeOutboundPorts(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.IncludeOutboundPorts = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetInterceptionMode(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.InterceptionMode = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetIstioDNSProxyEnabled(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.IstioDNSProxyEnabled = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetLifecycle(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.Lifecycle = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetLogLevel(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.LogLevel = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetNamespace(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetPostStart(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.PostStart = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetPreStop(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.PreStop = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetPrivileged(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.Privileged = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyInitAckSloCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyInitAckSloCPUResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyInitAckSloCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyInitAckSloCPUResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyInitAckSloMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyInitAckSloMemoryResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyInitAckSloMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyInitAckSloMemoryResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyInitCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyInitCPUResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyInitCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyInitCPUResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyInitMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyInitMemoryResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyInitMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyInitMemoryResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyMetadata(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyMetadata = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetProxyStatsMatcher(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ProxyStatsMatcher = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetReadinessFailureThreshold(v int32) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ReadinessFailureThreshold = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetReadinessInitialDelaySeconds(v int32) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ReadinessInitialDelaySeconds = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetReadinessPeriodSeconds(v int32) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ReadinessPeriodSeconds = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetRuntimeValues(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.RuntimeValues = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSMCEnabled(v bool) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SMCEnabled = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetScaledSidecarResourceShrink(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ScaledSidecarResourceShrink = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetServiceMeshId(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSidecarProxyAckSloCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SidecarProxyAckSloCPUResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSidecarProxyAckSloCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SidecarProxyAckSloCPUResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSidecarProxyAckSloMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SidecarProxyAckSloMemoryResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSidecarProxyAckSloMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SidecarProxyAckSloMemoryResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSidecarProxyCPUResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SidecarProxyCPUResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSidecarProxyCPUResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SidecarProxyCPUResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSidecarProxyMemoryResourceLimit(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SidecarProxyMemoryResourceLimit = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetSidecarProxyMemoryResourceRequest(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.SidecarProxyMemoryResourceRequest = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetTerminationDrainDuration(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.TerminationDrainDuration = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) SetTracing(v string) *UpdateNamespaceScopeSidecarConfigShrinkRequest {
	s.Tracing = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
