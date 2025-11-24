// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServiceMeshDetailResponseBody
	GetRequestId() *string
	SetServiceMesh(v *DescribeServiceMeshDetailResponseBodyServiceMesh) *DescribeServiceMeshDetailResponseBody
	GetServiceMesh() *DescribeServiceMeshDetailResponseBodyServiceMesh
}

type DescribeServiceMeshDetailResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the ASM instance.
	ServiceMesh *DescribeServiceMeshDetailResponseBodyServiceMesh `json:"ServiceMesh,omitempty" xml:"ServiceMesh,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshDetailResponseBody) GetServiceMesh() *DescribeServiceMeshDetailResponseBodyServiceMesh {
	return s.ServiceMesh
}

func (s *DescribeServiceMeshDetailResponseBody) SetRequestId(v string) *DescribeServiceMeshDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBody) SetServiceMesh(v *DescribeServiceMeshDetailResponseBodyServiceMesh) *DescribeServiceMeshDetailResponseBody {
	s.ServiceMesh = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBody) Validate() error {
	if s.ServiceMesh != nil {
		if err := s.ServiceMesh.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMesh struct {
	// The specification of the ASM instance. Valid values:
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
	// The endpoints of the ASM instance.
	Endpoints *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
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
	// The basic information about the ASM instance.
	ServiceMeshInfo *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo `json:"ServiceMeshInfo,omitempty" xml:"ServiceMeshInfo,omitempty" type:"Struct"`
	// The specifications of the ASM instance.
	Spec *DescribeServiceMeshDetailResponseBodyServiceMeshSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMesh) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMesh) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) GetClusters() []*string {
	return s.Clusters
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) GetEndpoints() *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	return s.Endpoints
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) GetServiceMeshInfo() *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo {
	return s.ServiceMeshInfo
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) GetSpec() *DescribeServiceMeshDetailResponseBodyServiceMeshSpec {
	return s.Spec
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMesh) Validate() error {
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
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints struct {
	// The endpoint that is used to access the API server from the internal network.
	//
	// example:
	//
	// https://``192.168.**.**``:6443
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	// The endpoint that is used to access Istio Pilot from the internal network (Intranet) during canary release.
	//
	// example:
	//
	// ``192.168.**.**``:15011
	IntranetCanaryPilotEndpoint *string `json:"IntranetCanaryPilotEndpoint,omitempty" xml:"IntranetCanaryPilotEndpoint,omitempty"`
	// The endpoint that is used to access Istio Pilot from the internal network.
	//
	// example:
	//
	// ``192.168.**.**``:15011
	IntranetPilotEndpoint *string `json:"IntranetPilotEndpoint,omitempty" xml:"IntranetPilotEndpoint,omitempty"`
	// The endpoint that is used to access the API server over the Internet.
	//
	// example:
	//
	// https://``123.56.**.**``:6443
	PublicApiServerEndpoint *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
	// The endpoint that is used to expose Istio Pilot to the public network (Internet) during canary release.
	//
	// example:
	//
	// ``182.92.**.**``:15011
	PublicCanaryPilotEndpoint *string `json:"PublicCanaryPilotEndpoint,omitempty" xml:"PublicCanaryPilotEndpoint,omitempty"`
	// The endpoint that is used to expose Istio Pilot to the Internet.
	//
	// example:
	//
	// ``182.92.**.**``:15011
	PublicPilotEndpoint *string `json:"PublicPilotEndpoint,omitempty" xml:"PublicPilotEndpoint,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) GetIntranetApiServerEndpoint() *string {
	return s.IntranetApiServerEndpoint
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) GetIntranetCanaryPilotEndpoint() *string {
	return s.IntranetCanaryPilotEndpoint
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) GetIntranetPilotEndpoint() *string {
	return s.IntranetPilotEndpoint
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) GetPublicApiServerEndpoint() *string {
	return s.PublicApiServerEndpoint
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) GetPublicCanaryPilotEndpoint() *string {
	return s.PublicCanaryPilotEndpoint
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) GetPublicPilotEndpoint() *string {
	return s.PublicPilotEndpoint
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) SetIntranetApiServerEndpoint(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	s.IntranetApiServerEndpoint = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) SetIntranetCanaryPilotEndpoint(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	s.IntranetCanaryPilotEndpoint = &v
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) SetPublicCanaryPilotEndpoint(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	s.PublicCanaryPilotEndpoint = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) SetPublicPilotEndpoint(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints {
	s.PublicPilotEndpoint = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo struct {
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
	// mesh1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The edition of the ASM instance. Valid values:
	//
	// 	- `Default`: Standard Edition
	//
	// 	- `Pro`: Enterprise Edition and Ultimate Edition
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region in which the ASM instance resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ASM instance ID.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The state of the ASM instance.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the ASM instance was last modified.
	//
	// example:
	//
	// 2020-06-03T14:48:54+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version of the ASM instance.
	//
	// example:
	//
	// v1.7.4.0-gfb34ba99-aliyun
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetName() *string {
	return s.Name
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetProfile() *string {
	return s.Profile
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetState() *string {
	return s.State
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) GetVersion() *string {
	return s.Version
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshServiceMeshInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpec struct {
	// The information about the load balancer.
	LoadBalancer *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Struct"`
	// The configurations of the ASM instance.
	MeshConfig *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	// The network configurations of the ASM instance.
	Network *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpec) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) GetLoadBalancer() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	return s.LoadBalancer
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) GetMeshConfig() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig {
	return s.MeshConfig
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) GetNetwork() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork {
	return s.Network
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpec) Validate() error {
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

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer struct {
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
	// The ID of the endpoint that is used to expose API server to the Internet.
	//
	// example:
	//
	// eip-wz9gtwau6b2aklgjk****
	ApiServerPublicEipId *string `json:"ApiServerPublicEipId,omitempty" xml:"ApiServerPublicEipId,omitempty"`
	// The ID of the Classic Load Balancer (CLB) instance that is used during the canary release of Istio Pilot.
	//
	// example:
	//
	// lb-2zesa8qs8kbkj9jkl****
	CanaryPilotLoadBalancerId *string `json:"CanaryPilotLoadBalancerId,omitempty" xml:"CanaryPilotLoadBalancerId,omitempty"`
	// The ID of the endpoint that is used to expose Istio Pilot to the Internet during canary release.
	//
	// example:
	//
	// eip-wz9gtwau6b2aklgjk****
	CanaryPilotPublicEipId *string `json:"CanaryPilotPublicEipId,omitempty" xml:"CanaryPilotPublicEipId,omitempty"`
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
	// The ID of the endpoint that is used to expose Istio Pilot to the Internet.
	//
	// example:
	//
	// eip-wz9gtwau6b2aklgjk****
	PilotPublicEipId *string `json:"PilotPublicEipId,omitempty" xml:"PilotPublicEipId,omitempty"`
	// The ID of the Classic Load Balancer (CLB) instance that is used when Istio Pilot is exposed to the Internet.
	//
	// example:
	//
	// lb-2zesa8qs8kbkj9jkl****
	PilotPublicLoadbalancerId *string `json:"PilotPublicLoadbalancerId,omitempty" xml:"PilotPublicLoadbalancerId,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GetApiServerLoadbalancerId() *string {
	return s.ApiServerLoadbalancerId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GetApiServerPublicEip() *bool {
	return s.ApiServerPublicEip
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GetApiServerPublicEipId() *string {
	return s.ApiServerPublicEipId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GetCanaryPilotLoadBalancerId() *string {
	return s.CanaryPilotLoadBalancerId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GetCanaryPilotPublicEipId() *string {
	return s.CanaryPilotPublicEipId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GetPilotPublicEip() *bool {
	return s.PilotPublicEip
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GetPilotPublicEipId() *string {
	return s.PilotPublicEipId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) GetPilotPublicLoadbalancerId() *string {
	return s.PilotPublicLoadbalancerId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetApiServerLoadbalancerId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.ApiServerLoadbalancerId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetApiServerPublicEip(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.ApiServerPublicEip = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetApiServerPublicEipId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.ApiServerPublicEipId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetCanaryPilotLoadBalancerId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.CanaryPilotLoadBalancerId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetCanaryPilotPublicEipId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.CanaryPilotPublicEipId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetPilotPublicEip(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.PilotPublicEip = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetPilotPublicEipId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.PilotPublicEipId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) SetPilotPublicLoadbalancerId(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer {
	s.PilotPublicLoadbalancerId = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecLoadBalancer) Validate() error {
	return dara.Validate(s)
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
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	CustomizedZipkin *bool `json:"CustomizedZipkin,omitempty" xml:"CustomizedZipkin,omitempty"`
	// The information about the edition.
	Edition *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition `json:"Edition,omitempty" xml:"Edition,omitempty" type:"Struct"`
	// Indicates whether the feature that routes traffic to the nearest instance is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	EnableLocalityLB *bool `json:"EnableLocalityLB,omitempty" xml:"EnableLocalityLB,omitempty"`
	// The IP ranges in CIDR form to be excluded from redirection to sidecar proxies in the ASM instance.
	//
	// example:
	//
	// 172.16.0.0
	ExcludeIPRanges *string `json:"ExcludeIPRanges,omitempty" xml:"ExcludeIPRanges,omitempty"`
	// The inbound ports to be excluded from redirection to sidecar proxies in the ASM instance.
	//
	// example:
	//
	// 80,81
	ExcludeInboundPorts *string `json:"ExcludeInboundPorts,omitempty" xml:"ExcludeInboundPorts,omitempty"`
	// The outbound ports to be excluded from redirection to sidecar proxies in the ASM instance.
	//
	// example:
	//
	// 80,81
	ExcludeOutboundPorts *string `json:"ExcludeOutboundPorts,omitempty" xml:"ExcludeOutboundPorts,omitempty"`
	// The configurations of additional features for the ASM instance.
	ExtraConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration `json:"ExtraConfiguration,omitempty" xml:"ExtraConfiguration,omitempty" type:"Struct"`
	// The IP ranges in CIDR form to redirect to the sidecar proxies in the ASM instance.
	//
	// example:
	//
	// 192.168.0.0/16
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
	// 	- `ALLOW_ANY`: Outbound traffic to all external services is allowed.
	//
	// 	- `REGISTRY_ONLY`: Outbound traffic is allowed to only external services that are defined in the service registry of the ASM instance.
	//
	// example:
	//
	// ALLOW_ANY
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
	// Indicates whether Prometheus monitoring is enabled. We recommend that you use [Managed Service for Prometheus](https://arms.console.aliyun.com/). Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Telemetry *bool `json:"Telemetry,omitempty" xml:"Telemetry,omitempty"`
	// Indicates whether tracing analysis is enabled. This feature can be enabled only after [Managed Service for OpenTelemetry](https://tracing-analysis.console.aliyun.com/) is activated. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Tracing *bool `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
	// The configurations of WebAssembly Filter.
	WebAssemblyFilterDeployment *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment `json:"WebAssemblyFilterDeployment,omitempty" xml:"WebAssemblyFilterDeployment,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetAccessLog() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog {
	return s.AccessLog
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetAudit() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit {
	return s.Audit
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetControlPlaneLogInfo() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo {
	return s.ControlPlaneLogInfo
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetCustomizedZipkin() *bool {
	return s.CustomizedZipkin
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetEdition() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition {
	return s.Edition
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetEnableLocalityLB() *bool {
	return s.EnableLocalityLB
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetExcludeIPRanges() *string {
	return s.ExcludeIPRanges
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetExcludeInboundPorts() *string {
	return s.ExcludeInboundPorts
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetExcludeOutboundPorts() *string {
	return s.ExcludeOutboundPorts
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetExtraConfiguration() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	return s.ExtraConfiguration
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetIncludeIPRanges() *string {
	return s.IncludeIPRanges
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetK8sNewAPIsSupport() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport {
	return s.K8sNewAPIsSupport
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetKiali() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali {
	return s.Kiali
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetLocalityLB() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB {
	return s.LocalityLB
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetMSE() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE {
	return s.MSE
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetOPA() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA {
	return s.OPA
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetOutboundTrafficPolicy() *string {
	return s.OutboundTrafficPolicy
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetPilot() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot {
	return s.Pilot
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetPrometheus() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus {
	return s.Prometheus
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetProtocolSupport() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport {
	return s.ProtocolSupport
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetProxy() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy {
	return s.Proxy
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetSidecarInjector() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector {
	return s.SidecarInjector
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetTelemetry() *bool {
	return s.Telemetry
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetTracing() *bool {
	return s.Tracing
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) GetWebAssemblyFilterDeployment() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment {
	return s.WebAssemblyFilterDeployment
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfig) Validate() error {
	if s.AccessLog != nil {
		if err := s.AccessLog.Validate(); err != nil {
			return err
		}
	}
	if s.Audit != nil {
		if err := s.Audit.Validate(); err != nil {
			return err
		}
	}
	if s.ControlPlaneLogInfo != nil {
		if err := s.ControlPlaneLogInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Edition != nil {
		if err := s.Edition.Validate(); err != nil {
			return err
		}
	}
	if s.ExtraConfiguration != nil {
		if err := s.ExtraConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.K8sNewAPIsSupport != nil {
		if err := s.K8sNewAPIsSupport.Validate(); err != nil {
			return err
		}
	}
	if s.Kiali != nil {
		if err := s.Kiali.Validate(); err != nil {
			return err
		}
	}
	if s.LocalityLB != nil {
		if err := s.LocalityLB.Validate(); err != nil {
			return err
		}
	}
	if s.MSE != nil {
		if err := s.MSE.Validate(); err != nil {
			return err
		}
	}
	if s.OPA != nil {
		if err := s.OPA.Validate(); err != nil {
			return err
		}
	}
	if s.Pilot != nil {
		if err := s.Pilot.Validate(); err != nil {
			return err
		}
	}
	if s.Prometheus != nil {
		if err := s.Prometheus.Validate(); err != nil {
			return err
		}
	}
	if s.ProtocolSupport != nil {
		if err := s.ProtocolSupport.Validate(); err != nil {
			return err
		}
	}
	if s.Proxy != nil {
		if err := s.Proxy.Validate(); err != nil {
			return err
		}
	}
	if s.SidecarInjector != nil {
		if err := s.SidecarInjector.Validate(); err != nil {
			return err
		}
	}
	if s.WebAssemblyFilterDeployment != nil {
		if err := s.WebAssemblyFilterDeployment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog struct {
	// Indicates whether access log collection is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the Simple Log Service project that stores access logs.
	//
	// example:
	//
	// k8s-log-b7b05d08670e41ca8c8fc0b7718f*****
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) GetProject() *string {
	return s.Project
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) SetProject(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog {
	s.Project = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAccessLog) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit struct {
	// Indicates whether an audit project exists in the ASM instance. Valid values:
	//
	// 	- `audit_project_exist`: An audit project exists.
	//
	// 	- `audit_project_not_exist`: No audit project exists.
	//
	// example:
	//
	// audit_project_not_exist
	AuditProjectStatus *string `json:"AuditProjectStatus,omitempty" xml:"AuditProjectStatus,omitempty"`
	// Indicates whether mesh audit is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the Simple Log Service project that is used for mesh audit.
	//
	// example:
	//
	// audit-project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) GetAuditProjectStatus() *string {
	return s.AuditProjectStatus
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) GetProject() *string {
	return s.Project
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigAudit) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo struct {
	// Indicates whether the collection of control plane logs is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time to live (TTL) of the collected control-plane logs. Unit: day.
	//
	// example:
	//
	// 30
	LogTTL *int32 `json:"LogTTL,omitempty" xml:"LogTTL,omitempty"`
	// The name of the Simple Log Service project that stores control plane logs.
	//
	// example:
	//
	// mesh-log-cbeb85a09161b4a26ab73e0ac****
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) GetLogTTL() *int32 {
	return s.LogTTL
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) GetProject() *string {
	return s.Project
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigControlPlaneLogInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition struct {
	// The version of the Istiod image.
	//
	// example:
	//
	// v1.9.7.1-3-gb3f1ab3c9c-pro-aliyun
	IstiodImageTag *string `json:"IstiodImageTag,omitempty" xml:"IstiodImageTag,omitempty"`
	// The name of the edition.
	//
	// example:
	//
	// Pro
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the Istio Proxy image.
	//
	// example:
	//
	// v1.9.7.1-3-gb3f1ab3c9c-pro-aliyun
	ProxyImageTag *string `json:"ProxyImageTag,omitempty" xml:"ProxyImageTag,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) GetIstiodImageTag() *string {
	return s.IstiodImageTag
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) GetName() *string {
	return s.Name
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) GetProxyImageTag() *string {
	return s.ProxyImageTag
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigEdition) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration struct {
	// The configurations of additional features for access log collection.
	AccessLogExtraConf *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf `json:"AccessLogExtraConf,omitempty" xml:"AccessLogExtraConf,omitempty" type:"Struct"`
	// The configurations of adaptive xDS optimization.
	AdaptiveXdsConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration `json:"AdaptiveXdsConfiguration,omitempty" xml:"AdaptiveXdsConfiguration,omitempty" type:"Struct"`
	// The configurations of automatic diagnostics for the ASM instance.
	AutoDiagnosis *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis `json:"AutoDiagnosis,omitempty" xml:"AutoDiagnosis,omitempty" type:"Struct"`
	// Access to Istio resources by using the Kubernetes API on the data plane.
	CRAggregationConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration `json:"CRAggregationConfiguration,omitempty" xml:"CRAggregationConfiguration,omitempty" type:"Struct"`
	// Indicates whether the Kubernetes API of clusters on the data plane can be used to access Istio resources. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	CRAggregationEnabled *bool `json:"CRAggregationEnabled,omitempty" xml:"CRAggregationEnabled,omitempty"`
	// The label selectors used to specify the namespaces of the clusters on the data plane. The control plane discovers and processes only application services in the specified namespaces.
	DiscoverySelectors []map[string]interface{} `json:"DiscoverySelectors,omitempty" xml:"DiscoverySelectors,omitempty" type:"Repeated"`
	// The configurations of the rollback feature for Istio resources.
	IstioCRHistory *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory `json:"IstioCRHistory,omitempty" xml:"IstioCRHistory,omitempty" type:"Struct"`
	// Additional configurations for Istiod.
	IstiodExtraConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration `json:"IstiodExtraConfiguration,omitempty" xml:"IstiodExtraConfiguration,omitempty" type:"Struct"`
	// The lifecycle of Istio Proxy.
	Lifecycle *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty" type:"Struct"`
	// The information about Transport Layer Security (TLS) acceleration based on MultiBuffer.
	MultiBuffer *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer `json:"MultiBuffer,omitempty" xml:"MultiBuffer,omitempty" type:"Struct"`
	// The configurations of Node Feature Discovery (NFD).
	NFDConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration `json:"NFDConfiguration,omitempty" xml:"NFDConfiguration,omitempty" type:"Struct"`
	// The configurations of the feature of controlling the OPA injection scope.
	OPAScopeInjection *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection `json:"OPAScopeInjection,omitempty" xml:"OPAScopeInjection,omitempty" type:"Struct"`
	Playground        *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground        `json:"Playground,omitempty" xml:"Playground,omitempty" type:"Struct"`
	// The resource limits on the istio-init container.
	SidecarProxyInitResourceLimit *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit `json:"SidecarProxyInitResourceLimit,omitempty" xml:"SidecarProxyInitResourceLimit,omitempty" type:"Struct"`
	// The resources that are required by the istio-init container.
	SidecarProxyInitResourceRequest *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest `json:"SidecarProxyInitResourceRequest,omitempty" xml:"SidecarProxyInitResourceRequest,omitempty" type:"Struct"`
	// The maximum period of time that Istio Proxy waits for a request to end.
	//
	// example:
	//
	// 5s
	TerminationDrainDuration *string `json:"TerminationDrainDuration,omitempty" xml:"TerminationDrainDuration,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetAccessLogExtraConf() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf {
	return s.AccessLogExtraConf
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetAdaptiveXdsConfiguration() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration {
	return s.AdaptiveXdsConfiguration
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetAutoDiagnosis() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis {
	return s.AutoDiagnosis
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetCRAggregationConfiguration() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration {
	return s.CRAggregationConfiguration
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetCRAggregationEnabled() *bool {
	return s.CRAggregationEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetDiscoverySelectors() []map[string]interface{} {
	return s.DiscoverySelectors
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetIstioCRHistory() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory {
	return s.IstioCRHistory
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetIstiodExtraConfiguration() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration {
	return s.IstiodExtraConfiguration
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetLifecycle() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle {
	return s.Lifecycle
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetMultiBuffer() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer {
	return s.MultiBuffer
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetNFDConfiguration() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration {
	return s.NFDConfiguration
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetOPAScopeInjection() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection {
	return s.OPAScopeInjection
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetPlayground() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground {
	return s.Playground
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetSidecarProxyInitResourceLimit() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit {
	return s.SidecarProxyInitResourceLimit
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetSidecarProxyInitResourceRequest() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest {
	return s.SidecarProxyInitResourceRequest
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) GetTerminationDrainDuration() *string {
	return s.TerminationDrainDuration
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetIstiodExtraConfiguration(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.IstiodExtraConfiguration = v
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) SetPlayground(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration {
	s.Playground = v
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfiguration) Validate() error {
	if s.AccessLogExtraConf != nil {
		if err := s.AccessLogExtraConf.Validate(); err != nil {
			return err
		}
	}
	if s.AdaptiveXdsConfiguration != nil {
		if err := s.AdaptiveXdsConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.AutoDiagnosis != nil {
		if err := s.AutoDiagnosis.Validate(); err != nil {
			return err
		}
	}
	if s.CRAggregationConfiguration != nil {
		if err := s.CRAggregationConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.IstioCRHistory != nil {
		if err := s.IstioCRHistory.Validate(); err != nil {
			return err
		}
	}
	if s.IstiodExtraConfiguration != nil {
		if err := s.IstiodExtraConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Lifecycle != nil {
		if err := s.Lifecycle.Validate(); err != nil {
			return err
		}
	}
	if s.MultiBuffer != nil {
		if err := s.MultiBuffer.Validate(); err != nil {
			return err
		}
	}
	if s.NFDConfiguration != nil {
		if err := s.NFDConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.OPAScopeInjection != nil {
		if err := s.OPAScopeInjection.Validate(); err != nil {
			return err
		}
	}
	if s.Playground != nil {
		if err := s.Playground.Validate(); err != nil {
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
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf struct {
	// Indicates whether gateway log collection is enabled.
	//
	// example:
	//
	// true
	GatewayEnabled *bool `json:"GatewayEnabled,omitempty" xml:"GatewayEnabled,omitempty"`
	// The retention period for the access logs of the ingress gateway. Unit: day. The logs are collected by using Simple Log Service. For example, the value 30 indicates that the logs are retained for 30 days.
	//
	// example:
	//
	// 30
	GatewayLifecycle *int32 `json:"GatewayLifecycle,omitempty" xml:"GatewayLifecycle,omitempty"`
	// Indicates whether sidecar log collection is enabled.
	//
	// example:
	//
	// true
	SidecarEnabled *bool `json:"SidecarEnabled,omitempty" xml:"SidecarEnabled,omitempty"`
	// The retention period for the access logs of sidecar proxies. Unit: day. The logs are collected by using Simple Log Service. For example, the value 30 indicates that the logs are retained for 30 days.
	//
	// example:
	//
	// 30
	SidecarLifecycle *int32 `json:"SidecarLifecycle,omitempty" xml:"SidecarLifecycle,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) GetGatewayEnabled() *bool {
	return s.GatewayEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) GetGatewayLifecycle() *int32 {
	return s.GatewayLifecycle
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) GetSidecarEnabled() *bool {
	return s.SidecarEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) GetSidecarLifecycle() *int32 {
	return s.SidecarLifecycle
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAccessLogExtraConf) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration struct {
	// Indicates whether Horizontal Pod Autoscaling (HPA) is enabled for the egress gateway.
	//
	// example:
	//
	// true
	EgressAutoscaleEnabled *bool `json:"EgressAutoscaleEnabled,omitempty" xml:"EgressAutoscaleEnabled,omitempty"`
	// The CPU resource configurations of the egress gateway when HPA is enabled.
	EgressHpaCpu *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu `json:"EgressHpaCpu,omitempty" xml:"EgressHpaCpu,omitempty" type:"Struct"`
	// The memory resource configurations of the egress gateway when HPA is enabled.
	EgressHpaMemory *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory `json:"EgressHpaMemory,omitempty" xml:"EgressHpaMemory,omitempty" type:"Struct"`
	// The maximum number of egress gateway pod replicas when HPA is enabled.
	//
	// example:
	//
	// 2
	EgressMaxReplica *int32 `json:"EgressMaxReplica,omitempty" xml:"EgressMaxReplica,omitempty"`
	// The minimum number of egress gateway pod replicas when HPA is enabled.
	//
	// example:
	//
	// 1
	EgressMinReplica *int32 `json:"EgressMinReplica,omitempty" xml:"EgressMinReplica,omitempty"`
	// The number of the egress gateway pod replicas.
	//
	// example:
	//
	// 2
	EgressReplicaCount *int32 `json:"EgressReplicaCount,omitempty" xml:"EgressReplicaCount,omitempty"`
	// The resource configurations of the egress gateway that is used by adaptive xDS optimization.
	EgressResources *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources `json:"EgressResources,omitempty" xml:"EgressResources,omitempty" type:"Struct"`
	// Indicates whether adaptive xDS optimization is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GetEgressAutoscaleEnabled() *bool {
	return s.EgressAutoscaleEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GetEgressHpaCpu() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu {
	return s.EgressHpaCpu
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GetEgressHpaMemory() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory {
	return s.EgressHpaMemory
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GetEgressMaxReplica() *int32 {
	return s.EgressMaxReplica
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GetEgressMinReplica() *int32 {
	return s.EgressMinReplica
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GetEgressReplicaCount() *int32 {
	return s.EgressReplicaCount
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GetEgressResources() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources {
	return s.EgressResources
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) GetEnabled() *bool {
	return s.Enabled
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfiguration) Validate() error {
	if s.EgressHpaCpu != nil {
		if err := s.EgressHpaCpu.Validate(); err != nil {
			return err
		}
	}
	if s.EgressHpaMemory != nil {
		if err := s.EgressHpaMemory.Validate(); err != nil {
			return err
		}
	}
	if s.EgressResources != nil {
		if err := s.EgressResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu struct {
	// The expected CPU utilization when HPA is enabled. Valid values: 1 to 100. If the CPU utilization exceeds this value, the number of pod replicas increases. If the CPU utilization is less than this value, the number of pod replicas decreases.
	//
	// example:
	//
	// 80
	TargetAverageUtilization *int32 `json:"TargetAverageUtilization,omitempty" xml:"TargetAverageUtilization,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) GetTargetAverageUtilization() *int32 {
	return s.TargetAverageUtilization
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) SetTargetAverageUtilization(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu {
	s.TargetAverageUtilization = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaCpu) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory struct {
	// The expected memory usage when HPA is enabled. Valid values: 1 to 100. If the memory usage exceeds this value, the number of pod replicas increases. If the memory usage is less than this value, the number of pod replicas decreases.
	//
	// example:
	//
	// 80
	TargetAverageUtilization *int32 `json:"TargetAverageUtilization,omitempty" xml:"TargetAverageUtilization,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) GetTargetAverageUtilization() *int32 {
	return s.TargetAverageUtilization
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) SetTargetAverageUtilization(v int32) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory {
	s.TargetAverageUtilization = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressHpaMemory) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources struct {
	// The resources that are available to the egress gateway.
	//
	// example:
	//
	// {"cpu":"200m", "memory": "512Mi"}
	Limits map[string]interface{} `json:"Limits,omitempty" xml:"Limits,omitempty"`
	// The resources that are requested by the egress gateway.
	//
	// example:
	//
	// {"cpu":"100m", "memory": "256Mi"}
	Requests map[string]interface{} `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) GetLimits() map[string]interface{} {
	return s.Limits
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) GetRequests() map[string]interface{} {
	return s.Requests
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) SetLimits(v map[string]interface{}) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources {
	s.Limits = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) SetRequests(v map[string]interface{}) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources {
	s.Requests = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAdaptiveXdsConfigurationEgressResources) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis struct {
	// Indicates whether automatic diagnostics is enabled for the ASM instance. If you enable this feature, the ASM instance is automatically diagnosed 5 minutes after you modify an Istio resource.
	//
	// example:
	//
	// true
	AutoDiagnosisEnabled *bool `json:"AutoDiagnosisEnabled,omitempty" xml:"AutoDiagnosisEnabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) GetAutoDiagnosisEnabled() *bool {
	return s.AutoDiagnosisEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) SetAutoDiagnosisEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis {
	s.AutoDiagnosisEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationAutoDiagnosis) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration struct {
	// Indicates whether Istio resources can be accessed by using the Kubernetes API on the data plane.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationCRAggregationConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory struct {
	// Indicates whether the rollback feature for Istio resources is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	EnableHistory *bool `json:"EnableHistory,omitempty" xml:"EnableHistory,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) GetEnableHistory() *bool {
	return s.EnableHistory
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) SetEnableHistory(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory {
	s.EnableHistory = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstioCRHistory) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration struct {
	// The labels for isolated workloads.
	//
	// example:
	//
	// name=xx,region=xx
	LabelsForOffloadedWorkloads *string `json:"LabelsForOffloadedWorkloads,omitempty" xml:"LabelsForOffloadedWorkloads,omitempty"`
	// example:
	//
	// false
	PilotEnableQuicListeners *bool `json:"PilotEnableQuicListeners,omitempty" xml:"PilotEnableQuicListeners,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration) GetLabelsForOffloadedWorkloads() *string {
	return s.LabelsForOffloadedWorkloads
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration) GetPilotEnableQuicListeners() *bool {
	return s.PilotEnableQuicListeners
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration) SetLabelsForOffloadedWorkloads(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration {
	s.LabelsForOffloadedWorkloads = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration) SetPilotEnableQuicListeners(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration {
	s.PilotEnableQuicListeners = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationIstiodExtraConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle struct {
	// The post-start parameters.
	PostStart *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart `json:"postStart,omitempty" xml:"postStart,omitempty" type:"Struct"`
	// The pre-close parameters.
	PreStop *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop `json:"preStop,omitempty" xml:"preStop,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) GetPostStart() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart {
	return s.PostStart
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) GetPreStop() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop {
	return s.PreStop
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) SetPostStart(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle {
	s.PostStart = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) SetPreStop(v *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle {
	s.PreStop = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecycle) Validate() error {
	if s.PostStart != nil {
		if err := s.PostStart.Validate(); err != nil {
			return err
		}
	}
	if s.PreStop != nil {
		if err := s.PreStop.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) GetExec() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec {
	return s.Exec
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) GetHttpGet() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet {
	return s.HttpGet
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) GetTcpSocket() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket {
	return s.TcpSocket
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStart) Validate() error {
	if s.Exec != nil {
		if err := s.Exec.Validate(); err != nil {
			return err
		}
	}
	if s.HttpGet != nil {
		if err := s.HttpGet.Validate(); err != nil {
			return err
		}
	}
	if s.TcpSocket != nil {
		if err := s.TcpSocket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec struct {
	// The executed commands. The value is a string that consists of JSON arrays.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) GetCommand() []*string {
	return s.Command
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) SetCommand(v []*string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec {
	s.Command = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartExec) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet struct {
	// The URL of the request.
	//
	// example:
	//
	// 127.xx.xx.1
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The HTTP request headers.
	HttpHeaders []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders `json:"httpHeaders,omitempty" xml:"httpHeaders,omitempty" type:"Repeated"`
	// The port number of the request.
	//
	// example:
	//
	// 80
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The request method. Valid values: `http` and `https`.
	//
	// example:
	//
	// http
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) GetHost() *string {
	return s.Host
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) GetHttpHeaders() []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders {
	return s.HttpHeaders
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) GetPort() *string {
	return s.Port
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) GetScheme() *string {
	return s.Scheme
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGet) Validate() error {
	if s.HttpHeaders != nil {
		for _, item := range s.HttpHeaders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders struct {
	// The name of the HTTP request header.
	//
	// example:
	//
	// key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the HTTP request header.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) GetName() *string {
	return s.Name
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) GetValue() *string {
	return s.Value
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) SetName(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) SetValue(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders {
	s.Value = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartHttpGetHttpHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket struct {
	// The URL of the TCP socket request.
	//
	// example:
	//
	// 127.xx.xx.1
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The port number of the TCP socket request.
	//
	// example:
	//
	// 888
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) GetHost() *string {
	return s.Host
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) GetPort() *string {
	return s.Port
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket {
	s.Port = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePostStartTcpSocket) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) GetExec() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec {
	return s.Exec
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) GetHttpGet() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet {
	return s.HttpGet
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) GetTcpSocket() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket {
	return s.TcpSocket
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStop) Validate() error {
	if s.Exec != nil {
		if err := s.Exec.Validate(); err != nil {
			return err
		}
	}
	if s.HttpGet != nil {
		if err := s.HttpGet.Validate(); err != nil {
			return err
		}
	}
	if s.TcpSocket != nil {
		if err := s.TcpSocket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec struct {
	// The executed commands. The value is a string that consists of JSON arrays.
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) GetCommand() []*string {
	return s.Command
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) SetCommand(v []*string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec {
	s.Command = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopExec) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet struct {
	// The URL of the request.
	//
	// example:
	//
	// 127.xx.xx.1
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The HTTP request headers.
	HttpHeaders []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders `json:"httpHeaders,omitempty" xml:"httpHeaders,omitempty" type:"Repeated"`
	// The port number of the request.
	//
	// example:
	//
	// 80
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The request method. Valid values: `http` and `https`.
	//
	// example:
	//
	// http
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) GetHost() *string {
	return s.Host
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) GetHttpHeaders() []*DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders {
	return s.HttpHeaders
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) GetPort() *string {
	return s.Port
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) GetScheme() *string {
	return s.Scheme
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGet) Validate() error {
	if s.HttpHeaders != nil {
		for _, item := range s.HttpHeaders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders struct {
	// The name of the HTTP request header.
	//
	// example:
	//
	// key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the HTTP request header.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) GetName() *string {
	return s.Name
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) GetValue() *string {
	return s.Value
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) SetName(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) SetValue(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders {
	s.Value = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopHttpGetHttpHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket struct {
	// The URL of the request.
	//
	// example:
	//
	// 127.xx.xx.1
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The port number of the request.
	//
	// example:
	//
	// 888
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) GetHost() *string {
	return s.Host
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) GetPort() *string {
	return s.Port
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) SetHost(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) SetPort(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket {
	s.Port = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationLifecyclePreStopTcpSocket) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer struct {
	// Indicates whether MultiBuffer-based TLS acceleration is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The pull-request latency.
	//
	// example:
	//
	// 0.02s
	PollDelay *string `json:"PollDelay,omitempty" xml:"PollDelay,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) GetPollDelay() *string {
	return s.PollDelay
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) SetPollDelay(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer {
	s.PollDelay = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationMultiBuffer) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration struct {
	// Indicates whether NFD is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether feature labels on nodes are cleared when NFD is disabled.
	//
	// example:
	//
	// true
	NFDLabelPruned *bool `json:"NFDLabelPruned,omitempty" xml:"NFDLabelPruned,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) GetNFDLabelPruned() *bool {
	return s.NFDLabelPruned
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) SetNFDLabelPruned(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration {
	s.NFDLabelPruned = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationNFDConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection struct {
	// Indicates whether the feature of controlling the OPA injection scope is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	OPAScopeInjected *bool `json:"OPAScopeInjected,omitempty" xml:"OPAScopeInjected,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) GetOPAScopeInjected() *bool {
	return s.OPAScopeInjected
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) SetOPAScopeInjected(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection {
	s.OPAScopeInjected = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationOPAScopeInjection) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground struct {
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground) GetScene() *string {
	return s.Scene
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground) SetScene(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground {
	s.Scene = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationPlayground) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit struct {
	// The maximum number of CPU cores that are available to the istio-init container.
	//
	// example:
	//
	// 2000m
	ResourceCPULimit *string `json:"ResourceCPULimit,omitempty" xml:"ResourceCPULimit,omitempty"`
	// The maximum size of the memory that is available to the istio-init container.
	//
	// example:
	//
	// 1024Mi
	ResourceMemoryLimit *string `json:"ResourceMemoryLimit,omitempty" xml:"ResourceMemoryLimit,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) GetResourceCPULimit() *string {
	return s.ResourceCPULimit
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) GetResourceMemoryLimit() *string {
	return s.ResourceMemoryLimit
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) SetResourceCPULimit(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit {
	s.ResourceCPULimit = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) SetResourceMemoryLimit(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit {
	s.ResourceMemoryLimit = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceLimit) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest struct {
	// The number of CPU cores that are requested by the istio-init container.
	//
	// example:
	//
	// 10m
	ResourceCPURequest *string `json:"ResourceCPURequest,omitempty" xml:"ResourceCPURequest,omitempty"`
	// The size of the memory that is requested by the istio-init container.
	//
	// example:
	//
	// 10Mi
	ResourceMemoryRequest *string `json:"ResourceMemoryRequest,omitempty" xml:"ResourceMemoryRequest,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) GetResourceCPURequest() *string {
	return s.ResourceCPURequest
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) GetResourceMemoryRequest() *string {
	return s.ResourceMemoryRequest
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) SetResourceCPURequest(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest {
	s.ResourceCPURequest = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) SetResourceMemoryRequest(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest {
	s.ResourceMemoryRequest = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigExtraConfigurationSidecarProxyInitResourceRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport struct {
	// Indicates whether Gateway API is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	GatewayAPIEnabled *bool `json:"GatewayAPIEnabled,omitempty" xml:"GatewayAPIEnabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) GetGatewayAPIEnabled() *bool {
	return s.GatewayAPIEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) SetGatewayAPIEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport {
	s.GatewayAPIEnabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigK8sNewAPIsSupport) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali struct {
	// Indicates whether mesh topology is enabled. Mesh topology can be enabled only when Prometheus monitoring is enabled. If Prometheus monitoring is disabled, you must set this parameter to `false`.`` Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The endpoint of the mesh topology service.
	//
	// example:
	//
	// http://``1.2.**.**``:20001
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) GetUrl() *string {
	return s.Url
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) SetUrl(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali {
	s.Url = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigKiali) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB struct {
	// The configurations of cross-region traffic distribution.
	//
	// >  Either `Failover` or Distribute can be set. If you set `Distribute`, you cannot set Failover.
	//
	// example:
	//
	// [{"from":"cn-shanghai","to":{"cn-hangzhou/*":50,"cn-shanghai/*":25,"cn-zhangjiakou/*":25}},{"from":"cn-hangzhou","to":{"cn-hangzhou/*":50,"cn-shanghai/*":25,"cn-zhangjiakou/*":25}}]
	Distribute map[string]interface{} `json:"Distribute,omitempty" xml:"Distribute,omitempty"`
	// Indicates whether cross-region load balancing is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The configurations of cross-region failover.
	//
	// >  Either Failover or `Distribute` can be set. If you set `Failover`, you cannot set `Distribute`.
	//
	// example:
	//
	// {"failover":[{"from":"cn-hangzhou","to":"cn-shanghai"}]}
	Failover         map[string]interface{} `json:"Failover,omitempty" xml:"Failover,omitempty"`
	FailoverPriority map[string]interface{} `json:"FailoverPriority,omitempty" xml:"FailoverPriority,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) GetDistribute() map[string]interface{} {
	return s.Distribute
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) GetFailover() map[string]interface{} {
	return s.Failover
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) GetFailoverPriority() map[string]interface{} {
	return s.FailoverPriority
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) SetFailoverPriority(v map[string]interface{}) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB {
	s.FailoverPriority = v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigLocalityLB) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE struct {
	// Indicates whether MSE is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigMSE) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA struct {
	// Indicates whether the OPA plug-in is installed. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of CPU cores that are available to the OPA proxy container.
	//
	// example:
	//
	// 1
	LimitCPU *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	// The maximum size of the memory that is available to the OPA proxy container.
	//
	// example:
	//
	// 512Mi
	LimitMemory *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	// The level of the logs to be generated for OPA.
	//
	// example:
	//
	// info
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The number of CPU cores that are requested by the OPA proxy container.
	//
	// example:
	//
	// 2
	RequestCPU *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The size of the memory that is requested by OPA.
	//
	// example:
	//
	// 1024Mi
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) GetLimitCPU() *string {
	return s.LimitCPU
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) GetLimitMemory() *string {
	return s.LimitMemory
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) GetLogLevel() *string {
	return s.LogLevel
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) GetRequestCPU() *string {
	return s.RequestCPU
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) GetRequestMemory() *string {
	return s.RequestMemory
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigOPA) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot struct {
	// The configurations of communication between external services and services in the mesh.
	ConfigSource *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource `json:"ConfigSource,omitempty" xml:"ConfigSource,omitempty" type:"Struct"`
	// The configurations of Pilot features.
	Feature *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature `json:"Feature,omitempty" xml:"Feature,omitempty" type:"Struct"`
	// Indicates whether HTTP/1.0 is supported. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Http10Enabled *bool `json:"Http10Enabled,omitempty" xml:"Http10Enabled,omitempty"`
	// The sampling percentage of tracing analysis.
	//
	// example:
	//
	// 100
	TraceSampling *float32 `json:"TraceSampling,omitempty" xml:"TraceSampling,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) GetConfigSource() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource {
	return s.ConfigSource
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) GetFeature() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature {
	return s.Feature
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) GetHttp10Enabled() *bool {
	return s.Http10Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) GetTraceSampling() *float32 {
	return s.TraceSampling
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilot) Validate() error {
	if s.ConfigSource != nil {
		if err := s.ConfigSource.Validate(); err != nil {
			return err
		}
	}
	if s.Feature != nil {
		if err := s.Feature.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource struct {
	// Indicates whether communication is allowed between external services and services in the mesh. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the Nacos instance that provides external service information.
	//
	// example:
	//
	// mse-cn-tl326******
	NacosID *string `json:"NacosID,omitempty" xml:"NacosID,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) GetNacosID() *string {
	return s.NacosID
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) SetNacosID(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource {
	s.NacosID = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotConfigSource) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature struct {
	// Indicates whether Secret Discovery Service (SDS) is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	EnableSDSServer *bool `json:"EnableSDSServer,omitempty" xml:"EnableSDSServer,omitempty"`
	// Indicates whether gateway configuration filtering is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	FilterGatewayClusterConfig *bool `json:"FilterGatewayClusterConfig,omitempty" xml:"FilterGatewayClusterConfig,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) GetEnableSDSServer() *bool {
	return s.EnableSDSServer
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) GetFilterGatewayClusterConfig() *bool {
	return s.FilterGatewayClusterConfig
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) SetEnableSDSServer(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature {
	s.EnableSDSServer = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) SetFilterGatewayClusterConfig(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature {
	s.FilterGatewayClusterConfig = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPilotFeature) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus struct {
	// The endpoint of Prometheus monitoring. If you use a custom Prometheus instance, this parameter is populated by the system.
	//
	// example:
	//
	// http://prometheus:9090
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// Indicates whether a custom Prometheus instance is used. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	UseExternal *bool `json:"UseExternal,omitempty" xml:"UseExternal,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) GetUseExternal() *bool {
	return s.UseExternal
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) SetExternalUrl(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus {
	s.ExternalUrl = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) SetUseExternal(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus {
	s.UseExternal = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigPrometheus) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport struct {
	// Indicates whether Dubbo Filter is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	DubboFilterEnabled *bool `json:"DubboFilterEnabled,omitempty" xml:"DubboFilterEnabled,omitempty"`
	// Indicates whether MySQL Filter is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	MysqlFilterEnabled *bool `json:"MysqlFilterEnabled,omitempty" xml:"MysqlFilterEnabled,omitempty"`
	// Indicates whether Redis Filter is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	RedisFilterEnabled *bool `json:"RedisFilterEnabled,omitempty" xml:"RedisFilterEnabled,omitempty"`
	// Indicates whether Thrift Filter is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	ThriftFilterEnabled *bool `json:"ThriftFilterEnabled,omitempty" xml:"ThriftFilterEnabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) GetDubboFilterEnabled() *bool {
	return s.DubboFilterEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) GetMysqlFilterEnabled() *bool {
	return s.MysqlFilterEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) GetRedisFilterEnabled() *bool {
	return s.RedisFilterEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) GetThriftFilterEnabled() *bool {
	return s.ThriftFilterEnabled
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProtocolSupport) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy struct {
	// The path to the file that stores the access logs of sidecar proxies.
	//
	// example:
	//
	// /dev/stdout
	AccessLogFile *string `json:"AccessLogFile,omitempty" xml:"AccessLogFile,omitempty"`
	// The format of the access logs of sidecar proxies.
	//
	// example:
	//
	// {"authority_for":"%REQ(:AUTHORITY)%","bytes_received":"%BYTES_RECEIVED%","bytes_sent":"%BYTES_SENT%","downstream_local_address":"%DOWNSTREAM_LOCAL_ADDRESS%","downstream_remote_address":"%DOWNSTREAM_REMOTE_ADDRESS%","duration":"%DURATION%","istio_policy_status":"%DYNAMIC_METADATA(istio.mixer:status)%","method":"%REQ(:METHOD)%","path":"%REQ(X-ENVOY-ORIGINAL-PATH?:PATH)%","protocol":"%PROTOCOL%","request_id":"%REQ(X-REQUEST-ID)%","requested_server_name":"%REQUESTED_SERVER_NAME%","response_code":"%RESPONSE_CODE%","response_flags":"%RESPONSE_FLAGS%","route_name":"%ROUTE_NAME%","start_time":"%START_TIME%","trace_id":"%REQ(X-B3-TRACEID)%","upstream_cluster":"%UPSTREAM_CLUSTER%","upstream_host":"%UPSTREAM_HOST%","upstream_local_address":"%UPSTREAM_LOCAL_ADDRESS%","upstream_service_time":"%RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)%","upstream_transport_failure_reason":"%UPSTREAM_TRANSPORT_FAILURE_REASON%","user_agent":"%REQ(USER-AGENT)%","x_forwarded_for":"%REQ(X-FORWARDED-FOR)%"}
	AccessLogFormat *string `json:"AccessLogFormat,omitempty" xml:"AccessLogFormat,omitempty"`
	// Indicates whether gRPC Access Log Service (ALS) for Envoy is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
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
	// The trusted domain.
	//
	// example:
	//
	// cluster.domain
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// Indicates whether the Domain Name System (DNS) proxy feature is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	EnableDNSProxying *bool `json:"EnableDNSProxying,omitempty" xml:"EnableDNSProxying,omitempty"`
	// The maximum number of CPU cores.
	//
	// example:
	//
	// 2000m
	LimitCPU *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	// The maximum size of the memory.
	//
	// example:
	//
	// 1024Mi
	LimitMemory *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	// The number of CPU cores that are requested.
	//
	// example:
	//
	// 100m
	RequestCPU *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The size of the memory that is requested.
	//
	// example:
	//
	// 128Mi
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetAccessLogFile() *string {
	return s.AccessLogFile
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetAccessLogFormat() *string {
	return s.AccessLogFormat
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetAccessLogServiceEnabled() *bool {
	return s.AccessLogServiceEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetAccessLogServiceHost() *string {
	return s.AccessLogServiceHost
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetAccessLogServicePort() *int32 {
	return s.AccessLogServicePort
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetClusterDomain() *string {
	return s.ClusterDomain
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetEnableDNSProxying() *bool {
	return s.EnableDNSProxying
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetLimitCPU() *string {
	return s.LimitCPU
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetLimitMemory() *string {
	return s.LimitMemory
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetRequestCPU() *string {
	return s.RequestCPU
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) GetRequestMemory() *string {
	return s.RequestMemory
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigProxy) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector struct {
	// Indicates whether automatic sidecar proxy injection can be enabled by using pod annotations. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
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
	// The configurations of Container Network Interface (CNI).
	InitCNIConfiguration *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration `json:"InitCNIConfiguration,omitempty" xml:"InitCNIConfiguration,omitempty" type:"Struct"`
	// The maximum number of CPU cores that are available to the pod where the sidecar injector resides.
	//
	// example:
	//
	// 4000m
	LimitCPU *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	// The maximum size of the memory that is available to the pod where the sidecar injector resides.
	//
	// example:
	//
	// 2048Mi
	LimitMemory *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	// The number of CPU cores that are requested by the pod where the sidecar injector resides.
	//
	// example:
	//
	// 1000m
	RequestCPU *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The size of the memory that is requested by the pod where the sidecar injector resides.
	//
	// example:
	//
	// 512Mi
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// The number of component replicas that are used for sidecar proxy injection. Default value: `1`.
	//
	// example:
	//
	// 1
	SidecarInjectorNum *int32 `json:"SidecarInjectorNum,omitempty" xml:"SidecarInjectorNum,omitempty"`
	// Other configurations of automatic sidecar proxy injection, in the YAML format. For more information, see [Enable automatic sidecar proxy injection](https://help.aliyun.com/document_detail/186136.html).
	//
	// example:
	//
	// {"injectedAnnotations":{"test/istio-init":"runtime/default2","test/istio-proxy":"runtime/default"},"replicaCount":2,"nodeSelector":{"beta.kubernetes.io/os":"linux"}}
	SidecarInjectorWebhookAsYaml *string `json:"SidecarInjectorWebhookAsYaml,omitempty" xml:"SidecarInjectorWebhookAsYaml,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetAutoInjectionPolicyEnabled() *bool {
	return s.AutoInjectionPolicyEnabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetEnableNamespacesByDefault() *bool {
	return s.EnableNamespacesByDefault
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetInitCNIConfiguration() *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	return s.InitCNIConfiguration
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetLimitCPU() *string {
	return s.LimitCPU
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetLimitMemory() *string {
	return s.LimitMemory
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetRequestCPU() *string {
	return s.RequestCPU
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetRequestMemory() *string {
	return s.RequestMemory
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetSidecarInjectorNum() *int32 {
	return s.SidecarInjectorNum
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) GetSidecarInjectorWebhookAsYaml() *string {
	return s.SidecarInjectorWebhookAsYaml
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjector) Validate() error {
	if s.InitCNIConfiguration != nil {
		if err := s.InitCNIConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration struct {
	// Indicates whether the CNI plug-in is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The namespaces to exclude. The CNI plug-in ignores pods in the excluded namespaces.
	//
	// example:
	//
	// kube-system,istio-system
	ExcludeNamespaces *string `json:"ExcludeNamespaces,omitempty" xml:"ExcludeNamespaces,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) GetExcludeNamespaces() *string {
	return s.ExcludeNamespaces
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) SetExcludeNamespaces(v string) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration {
	s.ExcludeNamespaces = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigSidecarInjectorInitCNIConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment struct {
	// Indicates whether WebAssembly Filter is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) SetEnabled(v bool) *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment {
	s.Enabled = &v
	return s
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecMeshConfigWebAssemblyFilterDeployment) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork struct {
	// The security group ID.
	//
	// example:
	//
	// sg-2ze384sxttxbctnj****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The virtual switches (vSwitches).
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2zew0rajjkmxy2369****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) GetVSwitches() []*string {
	return s.VSwitches
}

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) GetVpcId() *string {
	return s.VpcId
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

func (s *DescribeServiceMeshDetailResponseBodyServiceMeshSpecNetwork) Validate() error {
	return dara.Validate(s)
}
