// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceScopeSidecarConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigPatches(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) *DescribeNamespaceScopeSidecarConfigResponseBody
	GetConfigPatches() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches
	SetRequestId(v string) *DescribeNamespaceScopeSidecarConfigResponseBody
	GetRequestId() *string
}

type DescribeNamespaceScopeSidecarConfigResponseBody struct {
	// The namespace-level sidecar configurations.
	ConfigPatches *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches `json:"ConfigPatches,omitempty" xml:"ConfigPatches,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBody) GetConfigPatches() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	return s.ConfigPatches
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBody) SetConfigPatches(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) *DescribeNamespaceScopeSidecarConfigResponseBody {
	s.ConfigPatches = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBody) SetRequestId(v string) *DescribeNamespaceScopeSidecarConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBody) Validate() error {
	if s.ConfigPatches != nil {
		if err := s.ConfigPatches.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches struct {
	// The number of worker threads to run in the istio-proxy container.
	//
	// example:
	//
	// 2
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// Indicates whether the core dump feature is enabled for the sidecar proxy containers. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	EnableCoreDump *bool `json:"EnableCoreDump,omitempty" xml:"EnableCoreDump,omitempty"`
	// The inbound ports to be excluded from redirection to the sidecar proxy in the ASM instance.
	//
	// example:
	//
	// 82
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The outbound IP ranges in CIDR form to be excluded from redirection to the sidecar proxy in the ASM instance.
	//
	// example:
	//
	// 192.168.1.3/31
	ExcludeOutboundIPRanges *string `json:"ExcludeOutboundIPRanges,omitempty" xml:"ExcludeOutboundIPRanges,omitempty"`
	// The outbound ports to be excluded from redirection to the sidecar proxy in the ASM instance.
	//
	// example:
	//
	// 81
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// Indicates whether applications can be started only after the istio-proxy container starts. Valid values:
	//
	// 	- `true`: Applications can be started only after the istio-proxy container starts.
	//
	// 	- `false`: Applications can be started before the istio-proxy container starts.
	//
	// example:
	//
	// true
	HoldApplicationUntilProxyStarts *bool `json:"HoldApplicationUntilProxyStarts,omitempty" xml:"HoldApplicationUntilProxyStarts,omitempty"`
	// The inbound ports for which traffic is to be redirected to the sidecar proxy in the ASM instance.
	//
	// example:
	//
	// 83
	IncludeInboundPorts *string `json:"IncludeInboundPorts,omitempty" xml:"IncludeInboundPorts,omitempty"`
	// The outbound IP ranges in CIDR form for which traffic is to be redirected to the sidecar proxy in the ASM instance.
	//
	// example:
	//
	// 192.168.1.4/31
	IncludeOutboundIPRanges *string `json:"IncludeOutboundIPRanges,omitempty" xml:"IncludeOutboundIPRanges,omitempty"`
	// The outbound ports for which traffic is to be redirected to the sidecar proxy in the ASM instance.
	//
	// example:
	//
	// 84
	IncludeOutboundPorts *string `json:"IncludeOutboundPorts,omitempty" xml:"IncludeOutboundPorts,omitempty"`
	// The mode in which the sidecar proxy intercepts inbound traffic. Valid values:
	//
	// 	- `REDIRECT` (default): In this mode, source IP addresses are lost during traffic redirection to the sidecar proxy.
	//
	// 	- `TPROXY`: In this mode, both the source and destination IP addresses and ports are preserved.
	//
	// example:
	//
	// TPROXY
	InterceptionMode *string `json:"InterceptionMode,omitempty" xml:"InterceptionMode,omitempty"`
	// Indicates whether the Domain Name System (DNS) proxy feature is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	IstioDNSProxyEnabled *bool `json:"IstioDNSProxyEnabled,omitempty" xml:"IstioDNSProxyEnabled,omitempty"`
	// The JSON string that describes the lifecycle of the sidecar proxy.
	//
	// example:
	//
	// {"postStart":{"exec":{"command":["pilot-agent","wait"]}},"preStop":{"exec":{"command":["/bin/sh","-c","sleep 15"]}}}
	LifecycleStr *string `json:"LifecycleStr,omitempty" xml:"LifecycleStr,omitempty"`
	// The log level. Valid values: `info`, `debug`, `trace`, and `error`.
	//
	// example:
	//
	// info
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// Indicates whether the privileged mode is enabled in the security context of the sidecar proxy containers. Valid values:
	//
	// 	- `true`: The privileged mode is enabled, that is, the sidecar proxy containers run in privileged mode.
	//
	// 	- `false`: The privileged mode is not enabled.
	//
	// example:
	//
	// false
	Privileged    *bool              `json:"Privileged,omitempty" xml:"Privileged,omitempty"`
	ProxyMetadata map[string]*string `json:"ProxyMetadata,omitempty" xml:"ProxyMetadata,omitempty"`
	// The custom Envoy statistics that are reported by the sidecar proxy.
	ProxyStatsMatcher *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher `json:"ProxyStatsMatcher,omitempty" xml:"ProxyStatsMatcher,omitempty" type:"Struct"`
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
	// 2
	ReadinessInitialDelaySeconds *int32 `json:"ReadinessInitialDelaySeconds,omitempty" xml:"ReadinessInitialDelaySeconds,omitempty"`
	// The interval between two readiness checks of a sidecar proxy container. Unit: seconds.
	//
	// example:
	//
	// 3
	ReadinessPeriodSeconds *int32 `json:"ReadinessPeriodSeconds,omitempty" xml:"ReadinessPeriodSeconds,omitempty"`
	// Indicates the runtime parameters of Envoy proxy processes in the sidecar proxy container.
	RuntimeValues map[string]*string `json:"RuntimeValues,omitempty" xml:"RuntimeValues,omitempty"`
	// The configurations of the Shared Memory Communications over Remote Direct Memory Access (SMC-R) optimization feature.
	SMCConfiguration      *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration      `json:"SMCConfiguration,omitempty" xml:"SMCConfiguration,omitempty" type:"Struct"`
	ScaledSidecarResource *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource `json:"ScaledSidecarResource,omitempty" xml:"ScaledSidecarResource,omitempty" type:"Struct"`
	// The maximum size of reclaimed ACK resources that are available to the sidecar proxy container. This configuration item indicates the maximum size of resources that are available to the sidecar proxy container in a pod labeled with `koordinator.sh/qosClass`.
	SidecarProxyAckSloResource *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource `json:"SidecarProxyAckSloResource,omitempty" xml:"SidecarProxyAckSloResource,omitempty" type:"Struct"`
	// The reclaimed Container Service for Kubernetes (ACK) resources for the istio-init container. This configuration item indicates the resources of the istio-init container in a pod labeled with `koordinator.sh/qosClass`. Reclaimed ACK resources are resources that can be dynamically overcommitted.
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
	//
	// example:
	//
	// 6s
	TerminationDrainDuration *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
	// The custom configurations of Tracing Analysis.
	Tracing *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing `json:"Tracing,omitempty" xml:"Tracing,omitempty" type:"Struct"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetEnableCoreDump() *bool {
	return s.EnableCoreDump
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetExcludeInboundPorts() *string {
	return s.ExcludeInboundPorts
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetExcludeOutboundIPRanges() *string {
	return s.ExcludeOutboundIPRanges
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetExcludeOutboundPorts() *string {
	return s.ExcludeOutboundPorts
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetHoldApplicationUntilProxyStarts() *bool {
	return s.HoldApplicationUntilProxyStarts
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetIncludeInboundPorts() *string {
	return s.IncludeInboundPorts
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetIncludeOutboundIPRanges() *string {
	return s.IncludeOutboundIPRanges
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetIncludeOutboundPorts() *string {
	return s.IncludeOutboundPorts
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetInterceptionMode() *string {
	return s.InterceptionMode
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetIstioDNSProxyEnabled() *bool {
	return s.IstioDNSProxyEnabled
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetLifecycleStr() *string {
	return s.LifecycleStr
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetLogLevel() *string {
	return s.LogLevel
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetPrivileged() *bool {
	return s.Privileged
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetProxyMetadata() map[string]*string {
	return s.ProxyMetadata
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetProxyStatsMatcher() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher {
	return s.ProxyStatsMatcher
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetReadinessFailureThreshold() *int32 {
	return s.ReadinessFailureThreshold
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetReadinessInitialDelaySeconds() *int32 {
	return s.ReadinessInitialDelaySeconds
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetReadinessPeriodSeconds() *int32 {
	return s.ReadinessPeriodSeconds
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetRuntimeValues() map[string]*string {
	return s.RuntimeValues
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetSMCConfiguration() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration {
	return s.SMCConfiguration
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetScaledSidecarResource() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource {
	return s.ScaledSidecarResource
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetSidecarProxyAckSloResource() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource {
	return s.SidecarProxyAckSloResource
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetSidecarProxyInitAckSloResource() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource {
	return s.SidecarProxyInitAckSloResource
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetSidecarProxyInitResourceLimit() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit {
	return s.SidecarProxyInitResourceLimit
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetSidecarProxyInitResourceRequest() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest {
	return s.SidecarProxyInitResourceRequest
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetSidecarProxyResourceLimit() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit {
	return s.SidecarProxyResourceLimit
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetSidecarProxyResourceRequest() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest {
	return s.SidecarProxyResourceRequest
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetTerminationDrainDuration() *string {
	return s.TerminationDrainDuration
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) GetTracing() *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing {
	return s.Tracing
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

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetRuntimeValues(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.RuntimeValues = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetSMCConfiguration(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.SMCConfiguration = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) SetScaledSidecarResource(v *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches {
	s.ScaledSidecarResource = v
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

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatches) Validate() error {
	if s.ProxyStatsMatcher != nil {
		if err := s.ProxyStatsMatcher.Validate(); err != nil {
			return err
		}
	}
	if s.SMCConfiguration != nil {
		if err := s.SMCConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ScaledSidecarResource != nil {
		if err := s.ScaledSidecarResource.Validate(); err != nil {
			return err
		}
	}
	if s.SidecarProxyAckSloResource != nil {
		if err := s.SidecarProxyAckSloResource.Validate(); err != nil {
			return err
		}
	}
	if s.SidecarProxyInitAckSloResource != nil {
		if err := s.SidecarProxyInitAckSloResource.Validate(); err != nil {
			return err
		}
	}
	if s.SidecarProxyInitResourceLimit != nil {
		if err := s.SidecarProxyInitResourceLimit.Validate(); err != nil {
			return err
		}
	}
	if s.SidecarProxyInitResourceRequest != nil {
		if err := s.SidecarProxyInitResourceRequest.Validate(); err != nil {
			return err
		}
	}
	if s.SidecarProxyResourceLimit != nil {
		if err := s.SidecarProxyResourceLimit.Validate(); err != nil {
			return err
		}
	}
	if s.SidecarProxyResourceRequest != nil {
		if err := s.SidecarProxyResourceRequest.Validate(); err != nil {
			return err
		}
	}
	if s.Tracing != nil {
		if err := s.Tracing.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) GetInclusionPrefixes() []*string {
	return s.InclusionPrefixes
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) GetInclusionRegexps() []*string {
	return s.InclusionRegexps
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) GetInclusionSuffixes() []*string {
	return s.InclusionSuffixes
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

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesProxyStatsMatcher) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration struct {
	// Indicates whether the SMC-R optimization feature is enabled. The SMC-R optimization feature uses Alibaba Cloud Linux 3 and elastic remote direct memory access (eRDMA) network devices, which optimizes cross-node communication.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration) SetEnabled(v bool) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration {
	s.Enabled = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSMCConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource struct {
	ContainerRef                *string `json:"ContainerRef,omitempty" xml:"ContainerRef,omitempty"`
	ResourceCalculationStrategy *string `json:"ResourceCalculationStrategy,omitempty" xml:"ResourceCalculationStrategy,omitempty"`
	ResourcePercentage          *int32  `json:"ResourcePercentage,omitempty" xml:"ResourcePercentage,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) GetContainerRef() *string {
	return s.ContainerRef
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) GetResourceCalculationStrategy() *string {
	return s.ResourceCalculationStrategy
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) GetResourcePercentage() *int32 {
	return s.ResourcePercentage
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) SetContainerRef(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource {
	s.ContainerRef = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) SetResourceCalculationStrategy(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource {
	s.ResourceCalculationStrategy = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) SetResourcePercentage(v int32) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource {
	s.ResourcePercentage = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesScaledSidecarResource) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource struct {
	// The maximum size of reclaimed ACK resources that are available to the sidecar proxy container. This configuration item indicates the maximum size of resources that are available to the sidecar proxy container in a pod labeled with `koordinator.sh/qosClass`. The object can contain the following two types of keys, which indicate two types of resources:
	//
	// 	- `kubernetes.io/batch-cpu`: CPU resources that can be dynamically overcommitted. Unit: millicore.
	//
	// 	- `kubernetes.io/batch-memory`: memory resources that can be dynamically overcommitted.
	Limits map[string]*string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	// The minimum size of reclaimed ACK resources that the sidecar proxy container needs to use at runtime. This configuration item indicates the minimum size of reclaimed ACK resources for the sidecar proxy container in a pod labeled with `koordinator.sh/qosClass`. The object can contain the following two types of keys, which indicate two types of resources:
	//
	// 	- `kubernetes.io/batch-cpu`: CPU resources that can be dynamically overcommitted. Unit: millicore.
	//
	// 	- `kubernetes.io/batch-memory`: memory resources that can be dynamically overcommitted.
	Requests map[string]*string `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) GetLimits() map[string]*string {
	return s.Limits
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) GetRequests() map[string]*string {
	return s.Requests
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) SetLimits(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource {
	s.Limits = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) SetRequests(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource {
	s.Requests = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyAckSloResource) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource struct {
	// The maximum size of reclaimed ACK resources that are available to the istio-init container. This configuration item indicates the maximum size of resources that are available to the istio-init container in a pod labeled with `koordinator.sh/qosClass`. The object can contain the following two types of keys, which indicate two types of resources:
	//
	// 	- `kubernetes.io/batch-cpu`: CPU resources that can be dynamically overcommitted. Unit: millicore.
	//
	// 	- `kubernetes.io/batch-memory`: memory resources that can be dynamically overcommitted.
	Limits map[string]*string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	// The minimum size of reclaimed ACK resources that the istio-init container needs to use at runtime. This configuration item indicates the minimum size of reclaimed ACK resources for the istio-init container in a pod labeled with `koordinator.sh/qosClass`. The object can contain the following two types of keys, which indicate two types of resources:
	//
	// 	- `kubernetes.io/batch-cpu`: CPU resources that can be dynamically overcommitted. Unit: millicore.
	//
	// 	- `kubernetes.io/batch-memory`: memory resources that can be dynamically overcommitted.
	Requests map[string]*string `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) GetLimits() map[string]*string {
	return s.Limits
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) GetRequests() map[string]*string {
	return s.Requests
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) SetLimits(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource {
	s.Limits = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) SetRequests(v map[string]*string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource {
	s.Requests = v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitAckSloResource) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit struct {
	// The maximum number of CPU cores.
	//
	// example:
	//
	// 2000 m
	ResourceCPULimit *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
	// The maximum size of the memory.
	//
	// example:
	//
	// 50 Mi
	ResourceMemoryLimit *string `json:"ResourceMemoryLimit,omitempty" xml:"ResourceMemoryLimit,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) GetResourceCPULimit() *string {
	return s.ResourceCPULimit
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) GetResourceMemoryLimit() *string {
	return s.ResourceMemoryLimit
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) SetResourceCPULimit(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit {
	s.ResourceCPULimit = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) SetResourceMemoryLimit(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit {
	s.ResourceMemoryLimit = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceLimit) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest struct {
	// The minimum number of CPU cores.
	//
	// example:
	//
	// 60 m
	ResourceCPURequest *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
	// The minimum size of the memory.
	//
	// example:
	//
	// 30 Mi
	ResourceMemoryRequest *string `json:"ResourceMemoryRequest,omitempty" xml:"ResourceMemoryRequest,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) GetResourceCPURequest() *string {
	return s.ResourceCPURequest
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) GetResourceMemoryRequest() *string {
	return s.ResourceMemoryRequest
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) SetResourceCPURequest(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest {
	s.ResourceCPURequest = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) SetResourceMemoryRequest(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest {
	s.ResourceMemoryRequest = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyInitResourceRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit struct {
	// The maximum number of CPU cores.
	//
	// example:
	//
	// 2000 m
	ResourceCPULimit *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
	// The maximum size of the memory.
	//
	// example:
	//
	// 50 Mi
	ResourceMemoryLimit *string `json:"ResourceMemoryLimit,omitempty" xml:"ResourceMemoryLimit,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) GetResourceCPULimit() *string {
	return s.ResourceCPULimit
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) GetResourceMemoryLimit() *string {
	return s.ResourceMemoryLimit
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) SetResourceCPULimit(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit {
	s.ResourceCPULimit = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) SetResourceMemoryLimit(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit {
	s.ResourceMemoryLimit = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceLimit) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest struct {
	// The minimum number of CPU cores.
	//
	// example:
	//
	// 60 m
	ResourceCPURequest *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
	// The minimum size of the memory.
	//
	// example:
	//
	// 30 Mi
	ResourceMemoryRequest *string `json:"ResourceMemoryRequest,omitempty" xml:"ResourceMemoryRequest,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) GetResourceCPURequest() *string {
	return s.ResourceCPURequest
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) GetResourceMemoryRequest() *string {
	return s.ResourceMemoryRequest
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) SetResourceCPURequest(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest {
	s.ResourceCPURequest = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) SetResourceMemoryRequest(v string) *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest {
	s.ResourceMemoryRequest = &v
	return s
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesSidecarProxyResourceRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing struct {
	// The custom tags added to reported spans. The key of a tag is of the string type. The value of a tag is in the JSON format. A custom tag can belong to one of the following types:
	//
	// 	- `literal`: The tag value is a fixed value in the JSON format. This tag must contain the `value` field that specifies a literal. Example: `{"value":"test"}`.
	//
	// 	- `header`: The tag value is a request header in the JSON format. This tag must contain the `name` field and the `defaultValue` field. The name field indicates the name of the request header. The defaultValue field indicates the default value that is used when no request header is available. Example: `{"name":"test","defaultValue":"test"}`.
	//
	// 	- `environment`: The tag value is an environment variable in the JSON format. This tag must contain the `name` field and the `defaultValue` field. The name field indicates the name of the environment variable. The defaultValue field indicates the environment variable that is used when no environment variable is available. Example: `{"name":"test","defaultValue":"test"}`.
	//
	// example:
	//
	// {"test":{"literal":{"value":"test"}}}
	CustomTags map[string]interface{} `json:"CustomTags,omitempty" xml:"CustomTags,omitempty"`
	// The maximum tag length.
	//
	// example:
	//
	// 10
	MaxPathTagLength *int32 `json:"MaxPathTagLength,omitempty" xml:"MaxPathTagLength,omitempty"`
	// The sampling rate.
	//
	// example:
	//
	// 99.8
	Sampling *float64 `json:"Sampling,omitempty" xml:"Sampling,omitempty"`
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) GetCustomTags() map[string]interface{} {
	return s.CustomTags
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) GetMaxPathTagLength() *int32 {
	return s.MaxPathTagLength
}

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) GetSampling() *float64 {
	return s.Sampling
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

func (s *DescribeNamespaceScopeSidecarConfigResponseBodyConfigPatchesTracing) Validate() error {
	return dara.Validate(s)
}
