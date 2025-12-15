// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientVersion(v string) *UpdateClusterRequest
	GetClientVersion() *string
	SetClusterCustomConfiguration(v *UpdateClusterRequestClusterCustomConfiguration) *UpdateClusterRequest
	GetClusterCustomConfiguration() *UpdateClusterRequestClusterCustomConfiguration
	SetClusterDescription(v string) *UpdateClusterRequest
	GetClusterDescription() *string
	SetClusterId(v string) *UpdateClusterRequest
	GetClusterId() *string
	SetClusterName(v string) *UpdateClusterRequest
	GetClusterName() *string
	SetDeletionProtection(v bool) *UpdateClusterRequest
	GetDeletionProtection() *bool
	SetEnableScaleIn(v bool) *UpdateClusterRequest
	GetEnableScaleIn() *bool
	SetEnableScaleOut(v bool) *UpdateClusterRequest
	GetEnableScaleOut() *bool
	SetGrowInterval(v int32) *UpdateClusterRequest
	GetGrowInterval() *int32
	SetIdleInterval(v int32) *UpdateClusterRequest
	GetIdleInterval() *int32
	SetMaxCoreCount(v int32) *UpdateClusterRequest
	GetMaxCoreCount() *int32
	SetMaxCount(v int32) *UpdateClusterRequest
	GetMaxCount() *int32
	SetMonitorSpec(v *UpdateClusterRequestMonitorSpec) *UpdateClusterRequest
	GetMonitorSpec() *UpdateClusterRequestMonitorSpec
	SetSchedulerSpec(v *UpdateClusterRequestSchedulerSpec) *UpdateClusterRequest
	GetSchedulerSpec() *UpdateClusterRequestSchedulerSpec
}

type UpdateClusterRequest struct {
	// Specifies whether to enable auto scale-out for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 2.1.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// Specifies whether to enable auto scale-in for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	ClusterCustomConfiguration *UpdateClusterRequestClusterCustomConfiguration `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty" type:"Struct"`
	// The URL that is used to download the post-processing script.
	//
	// example:
	//
	// slurm22.05.8-serverless-cluster-20240805
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The client version. By default, the latest version is used.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The post-processing script of the cluster.
	//
	// example:
	//
	// slurm22.05.8-serverless-cluster-20240805
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The idle duration of the compute nodes allowed by the cluster.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The scheduler specifications of the cluster.
	//
	// example:
	//
	// 2
	GrowInterval *int32 `json:"GrowInterval,omitempty" xml:"GrowInterval,omitempty"`
	// Specifies whether to enable the topology awareness feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 4
	IdleInterval *int32 `json:"IdleInterval,omitempty" xml:"IdleInterval,omitempty"`
	// The interval at which the cluster is automatically scaled out.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The arguments that are used to run the post-processing script.
	//
	// example:
	//
	// 500
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The monitoring details of the cluster.
	MonitorSpec *UpdateClusterRequestMonitorSpec `json:"MonitorSpec,omitempty" xml:"MonitorSpec,omitempty" type:"Struct"`
	// The scheduler specifications of the cluster.
	SchedulerSpec *UpdateClusterRequestSchedulerSpec `json:"SchedulerSpec,omitempty" xml:"SchedulerSpec,omitempty" type:"Struct"`
}

func (s UpdateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *UpdateClusterRequest) GetClusterCustomConfiguration() *UpdateClusterRequestClusterCustomConfiguration {
	return s.ClusterCustomConfiguration
}

func (s *UpdateClusterRequest) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *UpdateClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateClusterRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *UpdateClusterRequest) GetEnableScaleIn() *bool {
	return s.EnableScaleIn
}

func (s *UpdateClusterRequest) GetEnableScaleOut() *bool {
	return s.EnableScaleOut
}

func (s *UpdateClusterRequest) GetGrowInterval() *int32 {
	return s.GrowInterval
}

func (s *UpdateClusterRequest) GetIdleInterval() *int32 {
	return s.IdleInterval
}

func (s *UpdateClusterRequest) GetMaxCoreCount() *int32 {
	return s.MaxCoreCount
}

func (s *UpdateClusterRequest) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *UpdateClusterRequest) GetMonitorSpec() *UpdateClusterRequestMonitorSpec {
	return s.MonitorSpec
}

func (s *UpdateClusterRequest) GetSchedulerSpec() *UpdateClusterRequestSchedulerSpec {
	return s.SchedulerSpec
}

func (s *UpdateClusterRequest) SetClientVersion(v string) *UpdateClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterCustomConfiguration(v *UpdateClusterRequestClusterCustomConfiguration) *UpdateClusterRequest {
	s.ClusterCustomConfiguration = v
	return s
}

func (s *UpdateClusterRequest) SetClusterDescription(v string) *UpdateClusterRequest {
	s.ClusterDescription = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterId(v string) *UpdateClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterName(v string) *UpdateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateClusterRequest) SetDeletionProtection(v bool) *UpdateClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *UpdateClusterRequest) SetEnableScaleIn(v bool) *UpdateClusterRequest {
	s.EnableScaleIn = &v
	return s
}

func (s *UpdateClusterRequest) SetEnableScaleOut(v bool) *UpdateClusterRequest {
	s.EnableScaleOut = &v
	return s
}

func (s *UpdateClusterRequest) SetGrowInterval(v int32) *UpdateClusterRequest {
	s.GrowInterval = &v
	return s
}

func (s *UpdateClusterRequest) SetIdleInterval(v int32) *UpdateClusterRequest {
	s.IdleInterval = &v
	return s
}

func (s *UpdateClusterRequest) SetMaxCoreCount(v int32) *UpdateClusterRequest {
	s.MaxCoreCount = &v
	return s
}

func (s *UpdateClusterRequest) SetMaxCount(v int32) *UpdateClusterRequest {
	s.MaxCount = &v
	return s
}

func (s *UpdateClusterRequest) SetMonitorSpec(v *UpdateClusterRequestMonitorSpec) *UpdateClusterRequest {
	s.MonitorSpec = v
	return s
}

func (s *UpdateClusterRequest) SetSchedulerSpec(v *UpdateClusterRequestSchedulerSpec) *UpdateClusterRequest {
	s.SchedulerSpec = v
	return s
}

func (s *UpdateClusterRequest) Validate() error {
	if s.ClusterCustomConfiguration != nil {
		if err := s.ClusterCustomConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.MonitorSpec != nil {
		if err := s.MonitorSpec.Validate(); err != nil {
			return err
		}
	}
	if s.SchedulerSpec != nil {
		if err := s.SchedulerSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateClusterRequestClusterCustomConfiguration struct {
	// Specifies whether to enable the monitoring component of compute nodes. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// E-HPC cn-hangzhou
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The monitoring details of the cluster.
	//
	// example:
	//
	// http://*****
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s UpdateClusterRequestClusterCustomConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterRequestClusterCustomConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequestClusterCustomConfiguration) GetArgs() *string {
	return s.Args
}

func (s *UpdateClusterRequestClusterCustomConfiguration) GetScript() *string {
	return s.Script
}

func (s *UpdateClusterRequestClusterCustomConfiguration) SetArgs(v string) *UpdateClusterRequestClusterCustomConfiguration {
	s.Args = &v
	return s
}

func (s *UpdateClusterRequestClusterCustomConfiguration) SetScript(v string) *UpdateClusterRequestClusterCustomConfiguration {
	s.Script = &v
	return s
}

func (s *UpdateClusterRequestClusterCustomConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateClusterRequestMonitorSpec struct {
	// Specifies whether to enable the monitoring component of compute nodes. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableComputeLoadMonitor *bool `json:"EnableComputeLoadMonitor,omitempty" xml:"EnableComputeLoadMonitor,omitempty"`
}

func (s UpdateClusterRequestMonitorSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterRequestMonitorSpec) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequestMonitorSpec) GetEnableComputeLoadMonitor() *bool {
	return s.EnableComputeLoadMonitor
}

func (s *UpdateClusterRequestMonitorSpec) SetEnableComputeLoadMonitor(v bool) *UpdateClusterRequestMonitorSpec {
	s.EnableComputeLoadMonitor = &v
	return s
}

func (s *UpdateClusterRequestMonitorSpec) Validate() error {
	return dara.Validate(s)
}

type UpdateClusterRequestSchedulerSpec struct {
	// Specifies whether to enable the topology awareness feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableTopologyAwareness *bool `json:"EnableTopologyAwareness,omitempty" xml:"EnableTopologyAwareness,omitempty"`
}

func (s UpdateClusterRequestSchedulerSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterRequestSchedulerSpec) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequestSchedulerSpec) GetEnableTopologyAwareness() *bool {
	return s.EnableTopologyAwareness
}

func (s *UpdateClusterRequestSchedulerSpec) SetEnableTopologyAwareness(v bool) *UpdateClusterRequestSchedulerSpec {
	s.EnableTopologyAwareness = &v
	return s
}

func (s *UpdateClusterRequestSchedulerSpec) Validate() error {
	return dara.Validate(s)
}
