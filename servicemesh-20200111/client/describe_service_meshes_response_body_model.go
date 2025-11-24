// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServiceMeshesResponseBody
	GetRequestId() *string
	SetServiceMeshes(v []*DescribeServiceMeshesResponseBodyServiceMeshes) *DescribeServiceMeshesResponseBody
	GetServiceMeshes() []*DescribeServiceMeshesResponseBodyServiceMeshes
}

type DescribeServiceMeshesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the ASM instance.
	ServiceMeshes []*DescribeServiceMeshesResponseBodyServiceMeshes `json:"ServiceMeshes,omitempty" xml:"ServiceMeshes,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshesResponseBody) GetServiceMeshes() []*DescribeServiceMeshesResponseBodyServiceMeshes {
	return s.ServiceMeshes
}

func (s *DescribeServiceMeshesResponseBody) SetRequestId(v string) *DescribeServiceMeshesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshesResponseBody) SetServiceMeshes(v []*DescribeServiceMeshesResponseBodyServiceMeshes) *DescribeServiceMeshesResponseBody {
	s.ServiceMeshes = v
	return s
}

func (s *DescribeServiceMeshesResponseBody) Validate() error {
	if s.ServiceMeshes != nil {
		for _, item := range s.ServiceMeshes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServiceMeshesResponseBodyServiceMeshes struct {
	// The edition of the ASM instance. Valid values:
	//
	// 	- `standard`: Standard Edition
	//
	// 	- `enterprise`: Enterprise Edition
	//
	// 	- `ultimate`: Ultimate Edition
	//
	// example:
	//
	// standard
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The clusters.
	Clusters []*string `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The information about all endpoints of the ASM instance.
	Endpoints *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud service instance for which the ASM instance is created.
	//
	// example:
	//
	// cc3e96f249d124eb38b72718ec5*****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud service for which the ASM instance is created. Valid values:
	//
	// 	- `ackone`: The ASM instance is created for Alibaba Cloud Distributed Cloud Container Platform (ACK One).
	//
	// 	- An empty value indicates that the ASM instance is created by the user.
	//
	// example:
	//
	// ackone
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The basic information about the ASM instances.
	ServiceMeshInfo *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo `json:"ServiceMeshInfo,omitempty" xml:"ServiceMeshInfo,omitempty" type:"Struct"`
	// The specifications of the ASM instance.
	Spec *DescribeServiceMeshesResponseBodyServiceMeshesSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// The details about the tags.
	Tag []*DescribeServiceMeshesResponseBodyServiceMeshesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Indicates whether the ASM instance can be upgraded.
	//
	// example:
	//
	// false
	Upgradable *bool `json:"Upgradable,omitempty" xml:"Upgradable,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshes) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshes) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetClusters() []*string {
	return s.Clusters
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetEndpoints() *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints {
	return s.Endpoints
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetServiceMeshInfo() *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo {
	return s.ServiceMeshInfo
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetSpec() *DescribeServiceMeshesResponseBodyServiceMeshesSpec {
	return s.Spec
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetTag() []*DescribeServiceMeshesResponseBodyServiceMeshesTag {
	return s.Tag
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) GetUpgradable() *bool {
	return s.Upgradable
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) SetUpgradable(v bool) *DescribeServiceMeshesResponseBodyServiceMeshes {
	s.Upgradable = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshes) Validate() error {
	if s.Endpoints != nil {
		if err := s.Endpoints.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceMeshInfo != nil {
		if err := s.ServiceMeshInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			return err
		}
	}
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

type DescribeServiceMeshesResponseBodyServiceMeshesEndpoints struct {
	// The endpoint that is used to access the API server over the internal network.
	//
	// example:
	//
	// https://192.168.xx.xx:6443
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	// The endpoint that is used to access Istio Pilot from the internal network.
	//
	// example:
	//
	// 192.168.xx.xx:15011
	IntranetPilotEndpoint *string `json:"IntranetPilotEndpoint,omitempty" xml:"IntranetPilotEndpoint,omitempty"`
	// The endpoint that is used to access the API server over the Internet.
	//
	// example:
	//
	// https://123.56.xx.xx:6443
	PublicApiServerEndpoint *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
	// The public endpoint of the Pilot of the ASM instance.
	//
	// example:
	//
	// xx.xx.xx.xx
	PublicPilotEndpoint *string `json:"PublicPilotEndpoint,omitempty" xml:"PublicPilotEndpoint,omitempty"`
	// The endpoint of the reverse tunnel (Deprecated).
	//
	// example:
	//
	// none
	ReverseTunnelEndpoint *string `json:"ReverseTunnelEndpoint,omitempty" xml:"ReverseTunnelEndpoint,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) GetIntranetApiServerEndpoint() *string {
	return s.IntranetApiServerEndpoint
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) GetIntranetPilotEndpoint() *string {
	return s.IntranetPilotEndpoint
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) GetPublicApiServerEndpoint() *string {
	return s.PublicApiServerEndpoint
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) GetPublicPilotEndpoint() *string {
	return s.PublicPilotEndpoint
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) GetReverseTunnelEndpoint() *string {
	return s.ReverseTunnelEndpoint
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshesEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo struct {
	// The time when the ASM instance was created.
	//
	// example:
	//
	// 2020-04-21T09:42:20+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the ASM instance.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The edition of the ASM instance before ASM is available for commercial use. Valid values:
	//
	// 	- `Pro`: Professional Edition
	//
	// 	- `Default`: Standard Edition
	//
	// example:
	//
	// Pro
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The region ID of the ASM instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the ASM instance.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The state of the ASM instance.
	//
	// example:
	//
	// success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the ASM instance was last modified.
	//
	// example:
	//
	// 2020-04-21T09:42:20+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version number of the ASM instance.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetName() *string {
	return s.Name
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetProfile() *string {
	return s.Profile
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetState() *string {
	return s.State
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) GetVersion() *string {
	return s.Version
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshesServiceMeshInfo) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpec) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpec) GetLoadBalancer() *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer {
	return s.LoadBalancer
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpec) GetMeshConfig() *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	return s.MeshConfig
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpec) GetNetwork() *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork {
	return s.Network
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpec) Validate() error {
	if s.LoadBalancer != nil {
		if err := s.LoadBalancer.Validate(); err != nil {
			return err
		}
	}
	if s.MeshConfig != nil {
		if err := s.MeshConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer struct {
	// The ID of the CLB instance that is used when the API server is exposed to the Internet.
	//
	// example:
	//
	// lb-2zekaak10uxds44vx****
	ApiServerLoadbalancerId *string `json:"ApiServerLoadbalancerId,omitempty" xml:"ApiServerLoadbalancerId,omitempty"`
	// Indicates whether the API server is exposed to the Internet. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	ApiServerPublicEip *bool `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	// Indicates whether Istio Pilot is exposed to the Internet. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	PilotPublicEip *bool `json:"PilotPublicEip,omitempty" xml:"PilotPublicEip,omitempty"`
	// The ID of the Classic Load Balancer (CLB) instance that is used when Istio Pilot is exposed to the Internet.
	//
	// example:
	//
	// lb-2zesa8qs8kbkj9jkl****
	PilotPublicLoadbalancerId *string `json:"PilotPublicLoadbalancerId,omitempty" xml:"PilotPublicLoadbalancerId,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) GetApiServerLoadbalancerId() *string {
	return s.ApiServerLoadbalancerId
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) GetApiServerPublicEip() *bool {
	return s.ApiServerPublicEip
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) GetPilotPublicEip() *bool {
	return s.PilotPublicEip
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) GetPilotPublicLoadbalancerId() *string {
	return s.PilotPublicLoadbalancerId
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecLoadBalancer) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig struct {
	// The extended configurations of the ASM instance.
	ExtraConfiguration *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration `json:"ExtraConfiguration,omitempty" xml:"ExtraConfiguration,omitempty" type:"Struct"`
	// Indicates whether nearby access is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Mtls *bool `json:"Mtls,omitempty" xml:"Mtls,omitempty"`
	// The outbound traffic policy. Valid values:
	//
	// 	- `ALLOW_ANY`: Outbound traffic to an external service is allowed.
	//
	// 	- `REGISTRY_ONLY`: Outbound traffic is allowed to only external services that are defined in the service registry of the ASM instance.
	//
	// example:
	//
	// ALLOW_ANY
	OutboundTrafficPolicy *string `json:"OutboundTrafficPolicy,omitempty" xml:"OutboundTrafficPolicy,omitempty"`
	// The configurations of the control plane.
	Pilot *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot `json:"Pilot,omitempty" xml:"Pilot,omitempty" type:"Struct"`
	// The configurations of sidecar proxy injection.
	SidecarInjector *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector `json:"SidecarInjector,omitempty" xml:"SidecarInjector,omitempty" type:"Struct"`
	// Indicates whether mutual Transport Layer Security (mTLS) is strictly enforced.
	//
	// example:
	//
	// true
	StrictMtls *bool `json:"StrictMtls,omitempty" xml:"StrictMtls,omitempty"`
	// Indicates whether Prometheus monitoring is enabled. We recommend that you use Managed Service for Prometheus. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Telemetry *bool `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	// Indicates whether the tracing feature is enabled. This feature can be enabled only after Managed Service for OpenTelemetry is activated. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Tracing *bool `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GetExtraConfiguration() *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration {
	return s.ExtraConfiguration
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GetMtls() *bool {
	return s.Mtls
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GetOutboundTrafficPolicy() *string {
	return s.OutboundTrafficPolicy
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GetPilot() *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot {
	return s.Pilot
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GetSidecarInjector() *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector {
	return s.SidecarInjector
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GetStrictMtls() *bool {
	return s.StrictMtls
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GetTelemetry() *bool {
	return s.Telemetry
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) GetTracing() *bool {
	return s.Tracing
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) SetExtraConfiguration(v *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig {
	s.ExtraConfiguration = v
	return s
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfig) Validate() error {
	if s.ExtraConfiguration != nil {
		if err := s.ExtraConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Pilot != nil {
		if err := s.Pilot.Validate(); err != nil {
			return err
		}
	}
	if s.SidecarInjector != nil {
		if err := s.SidecarInjector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration struct {
	// The configurations of the ASM Playground (valid only for ASM Playground instances).
	Playground *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground `json:"Playground,omitempty" xml:"Playground,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration) GetPlayground() *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground {
	return s.Playground
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration) SetPlayground(v *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration {
	s.Playground = v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfiguration) Validate() error {
	if s.Playground != nil {
		if err := s.Playground.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground struct {
	// The ID of the ASM Playground scenario.
	//
	// example:
	//
	// ewmaLb
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground) GetScene() *string {
	return s.Scene
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground) SetScene(v string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground {
	s.Scene = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigExtraConfigurationPlayground) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot struct {
	// Indicates whether the support for HTTP 1.0 is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Http10Enabled *bool `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
	// The sampling rate when Managed Service for OpenTelemetry is enabled.
	//
	// example:
	//
	// 0.2
	TraceSampling *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) GetHttp10Enabled() *bool {
	return s.Http10Enabled
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) GetTraceSampling() *float32 {
	return s.TraceSampling
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) SetHttp10Enabled(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot {
	s.Http10Enabled = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) SetTraceSampling(v float32) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot {
	s.TraceSampling = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigPilot) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector struct {
	// Indicates whether automatic sidecar proxy injection is enabled by using annotations.
	//
	// example:
	//
	// true
	AutoInjectionPolicyEnabled *bool `json:"AutoInjectionPolicyEnabled,omitempty" xml:"AutoInjectionPolicyEnabled,omitempty"`
	// Indicates whether automatic sidecar proxy injection is enabled for all namespaces. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	EnableNamespacesByDefault *bool `json:"EnableNamespacesByDefault,omitempty" xml:"EnableNamespacesByDefault,omitempty"`
	// The initial configurations of Container Network Interface (CNI).
	InitCNIConfiguration *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration `json:"InitCNIConfiguration,omitempty" xml:"InitCNIConfiguration,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) GetAutoInjectionPolicyEnabled() *bool {
	return s.AutoInjectionPolicyEnabled
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) GetEnableNamespacesByDefault() *bool {
	return s.EnableNamespacesByDefault
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) GetInitCNIConfiguration() *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	return s.InitCNIConfiguration
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjector) Validate() error {
	if s.InitCNIConfiguration != nil {
		if err := s.InitCNIConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration struct {
	// Indicates whether elevated privileges are required for the istio-init container when you perform traffic redirection for the istio-proxy container. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The namespaces in which the CNI plug-in does not check the pods.
	//
	// example:
	//
	// default,foo
	ExcludeNamespaces *string `json:"ExcludeNamespaces,omitempty" xml:"ExcludeNamespaces,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) GetExcludeNamespaces() *string {
	return s.ExcludeNamespaces
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) SetEnabled(v bool) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) SetExcludeNamespaces(v string) *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	s.ExcludeNamespaces = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecMeshConfigSidecarInjectorInitCNIConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork struct {
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze384sxttxbctnj****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the vSwitches.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2zew0rajjkmxy2369****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) GetVSwitches() []*string {
	return s.VSwitches
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) GetVpcId() *string {
	return s.VpcId
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

func (s *DescribeServiceMeshesResponseBodyServiceMeshesSpecNetwork) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshesResponseBodyServiceMeshesTag struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// yahaha
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshesResponseBodyServiceMeshesTag) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesTag) GetKey() *string {
	return s.Key
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesTag) GetValue() *string {
	return s.Value
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesTag) SetKey(v string) *DescribeServiceMeshesResponseBodyServiceMeshesTag {
	s.Key = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesTag) SetValue(v string) *DescribeServiceMeshesResponseBodyServiceMeshesTag {
	s.Value = &v
	return s
}

func (s *DescribeServiceMeshesResponseBodyServiceMeshesTag) Validate() error {
	return dara.Validate(s)
}
