// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientVersion(v string) *UpdateClusterShrinkRequest
	GetClientVersion() *string
	SetClusterCustomConfigurationShrink(v string) *UpdateClusterShrinkRequest
	GetClusterCustomConfigurationShrink() *string
	SetClusterDescription(v string) *UpdateClusterShrinkRequest
	GetClusterDescription() *string
	SetClusterId(v string) *UpdateClusterShrinkRequest
	GetClusterId() *string
	SetClusterName(v string) *UpdateClusterShrinkRequest
	GetClusterName() *string
	SetDeletionProtection(v bool) *UpdateClusterShrinkRequest
	GetDeletionProtection() *bool
	SetEnableScaleIn(v bool) *UpdateClusterShrinkRequest
	GetEnableScaleIn() *bool
	SetEnableScaleOut(v bool) *UpdateClusterShrinkRequest
	GetEnableScaleOut() *bool
	SetGrowInterval(v int32) *UpdateClusterShrinkRequest
	GetGrowInterval() *int32
	SetIdleInterval(v int32) *UpdateClusterShrinkRequest
	GetIdleInterval() *int32
	SetMaxCoreCount(v int32) *UpdateClusterShrinkRequest
	GetMaxCoreCount() *int32
	SetMaxCount(v int32) *UpdateClusterShrinkRequest
	GetMaxCount() *int32
	SetMonitorSpecShrink(v string) *UpdateClusterShrinkRequest
	GetMonitorSpecShrink() *string
	SetSchedulerSpecShrink(v string) *UpdateClusterShrinkRequest
	GetSchedulerSpecShrink() *string
}

type UpdateClusterShrinkRequest struct {
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
	ClusterCustomConfigurationShrink *string `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty"`
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
	MonitorSpecShrink *string `json:"MonitorSpec,omitempty" xml:"MonitorSpec,omitempty"`
	// The scheduler specifications of the cluster.
	SchedulerSpecShrink *string `json:"SchedulerSpec,omitempty" xml:"SchedulerSpec,omitempty"`
}

func (s UpdateClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterShrinkRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *UpdateClusterShrinkRequest) GetClusterCustomConfigurationShrink() *string {
	return s.ClusterCustomConfigurationShrink
}

func (s *UpdateClusterShrinkRequest) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *UpdateClusterShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateClusterShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *UpdateClusterShrinkRequest) GetEnableScaleIn() *bool {
	return s.EnableScaleIn
}

func (s *UpdateClusterShrinkRequest) GetEnableScaleOut() *bool {
	return s.EnableScaleOut
}

func (s *UpdateClusterShrinkRequest) GetGrowInterval() *int32 {
	return s.GrowInterval
}

func (s *UpdateClusterShrinkRequest) GetIdleInterval() *int32 {
	return s.IdleInterval
}

func (s *UpdateClusterShrinkRequest) GetMaxCoreCount() *int32 {
	return s.MaxCoreCount
}

func (s *UpdateClusterShrinkRequest) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *UpdateClusterShrinkRequest) GetMonitorSpecShrink() *string {
	return s.MonitorSpecShrink
}

func (s *UpdateClusterShrinkRequest) GetSchedulerSpecShrink() *string {
	return s.SchedulerSpecShrink
}

func (s *UpdateClusterShrinkRequest) SetClientVersion(v string) *UpdateClusterShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetClusterCustomConfigurationShrink(v string) *UpdateClusterShrinkRequest {
	s.ClusterCustomConfigurationShrink = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetClusterDescription(v string) *UpdateClusterShrinkRequest {
	s.ClusterDescription = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetClusterId(v string) *UpdateClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetClusterName(v string) *UpdateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetDeletionProtection(v bool) *UpdateClusterShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetEnableScaleIn(v bool) *UpdateClusterShrinkRequest {
	s.EnableScaleIn = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetEnableScaleOut(v bool) *UpdateClusterShrinkRequest {
	s.EnableScaleOut = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetGrowInterval(v int32) *UpdateClusterShrinkRequest {
	s.GrowInterval = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetIdleInterval(v int32) *UpdateClusterShrinkRequest {
	s.IdleInterval = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetMaxCoreCount(v int32) *UpdateClusterShrinkRequest {
	s.MaxCoreCount = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetMaxCount(v int32) *UpdateClusterShrinkRequest {
	s.MaxCount = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetMonitorSpecShrink(v string) *UpdateClusterShrinkRequest {
	s.MonitorSpecShrink = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetSchedulerSpecShrink(v string) *UpdateClusterShrinkRequest {
	s.SchedulerSpecShrink = &v
	return s
}

func (s *UpdateClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
