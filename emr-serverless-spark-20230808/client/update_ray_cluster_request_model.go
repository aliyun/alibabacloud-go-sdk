// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRayClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRayClusterRequest
	GetDescription() *string
	SetDisplayReleaseVersion(v string) *UpdateRayClusterRequest
	GetDisplayReleaseVersion() *string
	SetExtraParam(v string) *UpdateRayClusterRequest
	GetExtraParam() *string
	SetHeadSpec(v *UpdateRayClusterRequestHeadSpec) *UpdateRayClusterRequest
	GetHeadSpec() *UpdateRayClusterRequestHeadSpec
	SetName(v string) *UpdateRayClusterRequest
	GetName() *string
	SetNetworkServiceName(v string) *UpdateRayClusterRequest
	GetNetworkServiceName() *string
	SetVolumeIds(v []*string) *UpdateRayClusterRequest
	GetVolumeIds() []*string
	SetWorkerSpec(v []*UpdateRayClusterRequestWorkerSpec) *UpdateRayClusterRequest
	GetWorkerSpec() []*UpdateRayClusterRequestWorkerSpec
}

type UpdateRayClusterRequest struct {
	// The description.
	//
	// example:
	//
	// Ray Cluster for dev.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The Ray DPI engine version.
	//
	// example:
	//
	// ray-1.0.0 (Ray 2.47.1, Python 3.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The extra parameters. The value must be in JSON format.
	//
	// example:
	//
	// {}
	ExtraParam *string `json:"extraParam,omitempty" xml:"extraParam,omitempty"`
	// The Ray cluster head node information.
	HeadSpec *UpdateRayClusterRequestHeadSpec `json:"headSpec,omitempty" xml:"headSpec,omitempty" type:"Struct"`
	// The Ray cluster name. The name must be 1 to 64 characters in length.
	//
	// example:
	//
	// myRayCluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network connectivity name.
	//
	// example:
	//
	// vpc
	NetworkServiceName *string `json:"networkServiceName,omitempty" xml:"networkServiceName,omitempty"`
	// The list of managed folder IDs to mount.
	VolumeIds []*string `json:"volumeIds,omitempty" xml:"volumeIds,omitempty" type:"Repeated"`
	// The Ray cluster worker node information. A maximum of 50 groups are supported.
	WorkerSpec []*UpdateRayClusterRequestWorkerSpec `json:"workerSpec,omitempty" xml:"workerSpec,omitempty" type:"Repeated"`
}

func (s UpdateRayClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRayClusterRequest) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *UpdateRayClusterRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *UpdateRayClusterRequest) GetHeadSpec() *UpdateRayClusterRequestHeadSpec {
	return s.HeadSpec
}

func (s *UpdateRayClusterRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRayClusterRequest) GetNetworkServiceName() *string {
	return s.NetworkServiceName
}

func (s *UpdateRayClusterRequest) GetVolumeIds() []*string {
	return s.VolumeIds
}

func (s *UpdateRayClusterRequest) GetWorkerSpec() []*UpdateRayClusterRequestWorkerSpec {
	return s.WorkerSpec
}

func (s *UpdateRayClusterRequest) SetDescription(v string) *UpdateRayClusterRequest {
	s.Description = &v
	return s
}

func (s *UpdateRayClusterRequest) SetDisplayReleaseVersion(v string) *UpdateRayClusterRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *UpdateRayClusterRequest) SetExtraParam(v string) *UpdateRayClusterRequest {
	s.ExtraParam = &v
	return s
}

func (s *UpdateRayClusterRequest) SetHeadSpec(v *UpdateRayClusterRequestHeadSpec) *UpdateRayClusterRequest {
	s.HeadSpec = v
	return s
}

func (s *UpdateRayClusterRequest) SetName(v string) *UpdateRayClusterRequest {
	s.Name = &v
	return s
}

func (s *UpdateRayClusterRequest) SetNetworkServiceName(v string) *UpdateRayClusterRequest {
	s.NetworkServiceName = &v
	return s
}

func (s *UpdateRayClusterRequest) SetVolumeIds(v []*string) *UpdateRayClusterRequest {
	s.VolumeIds = v
	return s
}

func (s *UpdateRayClusterRequest) SetWorkerSpec(v []*UpdateRayClusterRequestWorkerSpec) *UpdateRayClusterRequest {
	s.WorkerSpec = v
	return s
}

func (s *UpdateRayClusterRequest) Validate() error {
	if s.HeadSpec != nil {
		if err := s.HeadSpec.Validate(); err != nil {
			return err
		}
	}
	if s.WorkerSpec != nil {
		for _, item := range s.WorkerSpec {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRayClusterRequestHeadSpec struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The Ray DPI engine version.
	//
	// example:
	//
	// ray-1.2.0 (Ray 2.55.1, Python 3.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// Specifies whether to enable automatic scaling.
	//
	// example:
	//
	// false
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty" xml:"enableAutoScaling,omitempty"`
	// The environment variables.
	//
	// example:
	//
	// MY_ENV=123456
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// The GCS Fault Tolerance configuration.
	GftConfig *UpdateRayClusterRequestHeadSpecGftConfig `json:"gftConfig,omitempty" xml:"gftConfig,omitempty" type:"Struct"`
	// Specifies whether to enable GCS Fault Tolerance.
	//
	// if can be null:
	// true
	GftEnabled *bool `json:"gftEnabled,omitempty" xml:"gftEnabled,omitempty"`
	// The GPU instance type.
	//
	// example:
	//
	// ecs.gn6i-c4g1.xlarge
	GpuSpec *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
	// The idle timeout period of workers after automatic scaling is enabled.
	//
	// example:
	//
	// 60
	IdleTimeoutSeconds *int32 `json:"idleTimeoutSeconds,omitempty" xml:"idleTimeoutSeconds,omitempty"`
	// The memory size. Unit: GiB.
	//
	// example:
	//
	// 8Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// The queue name.
	//
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The Ray startup parameters.
	//
	// example:
	//
	// --num-cpus=0
	RayStartParams *string `json:"rayStartParams,omitempty" xml:"rayStartParams,omitempty"`
}

func (s UpdateRayClusterRequestHeadSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterRequestHeadSpec) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterRequestHeadSpec) GetCpu() *string {
	return s.Cpu
}

func (s *UpdateRayClusterRequestHeadSpec) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *UpdateRayClusterRequestHeadSpec) GetEnableAutoScaling() *bool {
	return s.EnableAutoScaling
}

func (s *UpdateRayClusterRequestHeadSpec) GetEnv() *string {
	return s.Env
}

func (s *UpdateRayClusterRequestHeadSpec) GetGftConfig() *UpdateRayClusterRequestHeadSpecGftConfig {
	return s.GftConfig
}

func (s *UpdateRayClusterRequestHeadSpec) GetGftEnabled() *bool {
	return s.GftEnabled
}

func (s *UpdateRayClusterRequestHeadSpec) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *UpdateRayClusterRequestHeadSpec) GetIdleTimeoutSeconds() *int32 {
	return s.IdleTimeoutSeconds
}

func (s *UpdateRayClusterRequestHeadSpec) GetMemory() *string {
	return s.Memory
}

func (s *UpdateRayClusterRequestHeadSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateRayClusterRequestHeadSpec) GetRayStartParams() *string {
	return s.RayStartParams
}

func (s *UpdateRayClusterRequestHeadSpec) SetCpu(v string) *UpdateRayClusterRequestHeadSpec {
	s.Cpu = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetDisplayReleaseVersion(v string) *UpdateRayClusterRequestHeadSpec {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetEnableAutoScaling(v bool) *UpdateRayClusterRequestHeadSpec {
	s.EnableAutoScaling = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetEnv(v string) *UpdateRayClusterRequestHeadSpec {
	s.Env = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetGftConfig(v *UpdateRayClusterRequestHeadSpecGftConfig) *UpdateRayClusterRequestHeadSpec {
	s.GftConfig = v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetGftEnabled(v bool) *UpdateRayClusterRequestHeadSpec {
	s.GftEnabled = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetGpuSpec(v string) *UpdateRayClusterRequestHeadSpec {
	s.GpuSpec = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetIdleTimeoutSeconds(v int32) *UpdateRayClusterRequestHeadSpec {
	s.IdleTimeoutSeconds = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetMemory(v string) *UpdateRayClusterRequestHeadSpec {
	s.Memory = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetQueueName(v string) *UpdateRayClusterRequestHeadSpec {
	s.QueueName = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetRayStartParams(v string) *UpdateRayClusterRequestHeadSpec {
	s.RayStartParams = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) Validate() error {
	if s.GftConfig != nil {
		if err := s.GftConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRayClusterRequestHeadSpecGftConfig struct {
	// The Redis password.
	//
	// example:
	//
	// 123456
	RedisPassword *string `json:"redisPassword,omitempty" xml:"redisPassword,omitempty"`
	// The Redis address.
	//
	// example:
	//
	// 10.4.5.6:6379
	RedisUrl *string `json:"redisUrl,omitempty" xml:"redisUrl,omitempty"`
	// The Redis username.
	//
	// example:
	//
	// default
	RedisUsername *string `json:"redisUsername,omitempty" xml:"redisUsername,omitempty"`
}

func (s UpdateRayClusterRequestHeadSpecGftConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterRequestHeadSpecGftConfig) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterRequestHeadSpecGftConfig) GetRedisPassword() *string {
	return s.RedisPassword
}

func (s *UpdateRayClusterRequestHeadSpecGftConfig) GetRedisUrl() *string {
	return s.RedisUrl
}

func (s *UpdateRayClusterRequestHeadSpecGftConfig) GetRedisUsername() *string {
	return s.RedisUsername
}

func (s *UpdateRayClusterRequestHeadSpecGftConfig) SetRedisPassword(v string) *UpdateRayClusterRequestHeadSpecGftConfig {
	s.RedisPassword = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpecGftConfig) SetRedisUrl(v string) *UpdateRayClusterRequestHeadSpecGftConfig {
	s.RedisUrl = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpecGftConfig) SetRedisUsername(v string) *UpdateRayClusterRequestHeadSpecGftConfig {
	s.RedisUsername = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpecGftConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRayClusterRequestWorkerSpec struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The DPI engine version.
	//
	// example:
	//
	// ray-1.2.0 (Ray 2.55.1, Python 3.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The Ray environment variables.
	//
	// example:
	//
	// MY_ENV=123456
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// The GPU instance type.
	//
	// example:
	//
	// ecs.gn6i-c4g1.xlarge
	GpuSpec *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
	// The worker group name.
	//
	// example:
	//
	// WorkerGroup1
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The maximum number of workers. Minimum value: 1.
	//
	// example:
	//
	// 10
	MaxReplica *int32 `json:"maxReplica,omitempty" xml:"maxReplica,omitempty"`
	// The memory size. Unit: GiB.
	//
	// example:
	//
	// 8Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// The minimum number of workers. Minimum value: 1. The value must be less than or equal to maxReplica.
	//
	// example:
	//
	// 1
	MinReplica *int32 `json:"minReplica,omitempty" xml:"minReplica,omitempty"`
	// The queue name.
	//
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The Ray startup parameters.
	//
	// example:
	//
	// --num-cpus=0
	RayStartParams *string `json:"rayStartParams,omitempty" xml:"rayStartParams,omitempty"`
	// The number of workers. Minimum value: 1.
	//
	// example:
	//
	// 2
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// The worker type.
	//
	// example:
	//
	// CPU
	WorkerType *string `json:"workerType,omitempty" xml:"workerType,omitempty"`
}

func (s UpdateRayClusterRequestWorkerSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterRequestWorkerSpec) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterRequestWorkerSpec) GetCpu() *string {
	return s.Cpu
}

func (s *UpdateRayClusterRequestWorkerSpec) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *UpdateRayClusterRequestWorkerSpec) GetEnv() *string {
	return s.Env
}

func (s *UpdateRayClusterRequestWorkerSpec) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *UpdateRayClusterRequestWorkerSpec) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateRayClusterRequestWorkerSpec) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *UpdateRayClusterRequestWorkerSpec) GetMemory() *string {
	return s.Memory
}

func (s *UpdateRayClusterRequestWorkerSpec) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *UpdateRayClusterRequestWorkerSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateRayClusterRequestWorkerSpec) GetRayStartParams() *string {
	return s.RayStartParams
}

func (s *UpdateRayClusterRequestWorkerSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *UpdateRayClusterRequestWorkerSpec) GetWorkerType() *string {
	return s.WorkerType
}

func (s *UpdateRayClusterRequestWorkerSpec) SetCpu(v string) *UpdateRayClusterRequestWorkerSpec {
	s.Cpu = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetDisplayReleaseVersion(v string) *UpdateRayClusterRequestWorkerSpec {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetEnv(v string) *UpdateRayClusterRequestWorkerSpec {
	s.Env = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetGpuSpec(v string) *UpdateRayClusterRequestWorkerSpec {
	s.GpuSpec = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetGroupName(v string) *UpdateRayClusterRequestWorkerSpec {
	s.GroupName = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetMaxReplica(v int32) *UpdateRayClusterRequestWorkerSpec {
	s.MaxReplica = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetMemory(v string) *UpdateRayClusterRequestWorkerSpec {
	s.Memory = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetMinReplica(v int32) *UpdateRayClusterRequestWorkerSpec {
	s.MinReplica = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetQueueName(v string) *UpdateRayClusterRequestWorkerSpec {
	s.QueueName = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetRayStartParams(v string) *UpdateRayClusterRequestWorkerSpec {
	s.RayStartParams = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetReplica(v int32) *UpdateRayClusterRequestWorkerSpec {
	s.Replica = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetWorkerType(v string) *UpdateRayClusterRequestWorkerSpec {
	s.WorkerType = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) Validate() error {
	return dara.Validate(s)
}
