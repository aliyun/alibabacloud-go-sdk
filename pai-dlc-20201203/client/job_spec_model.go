// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAssignNodeSpec(v *AssignNodeSpec) *JobSpec
	GetAssignNodeSpec() *AssignNodeSpec
	SetAutoScalingSpec(v *AutoScalingSpec) *JobSpec
	GetAutoScalingSpec() *AutoScalingSpec
	SetEcsSpec(v string) *JobSpec
	GetEcsSpec() *string
	SetExtraPodSpec(v *ExtraPodSpec) *JobSpec
	GetExtraPodSpec() *ExtraPodSpec
	SetImage(v string) *JobSpec
	GetImage() *string
	SetImageConfig(v *ImageConfig) *JobSpec
	GetImageConfig() *ImageConfig
	SetIsCheif(v bool) *JobSpec
	GetIsCheif() *bool
	SetIsChief(v bool) *JobSpec
	GetIsChief() *bool
	SetLocalMountSpecs(v []*LocalMountSpec) *JobSpec
	GetLocalMountSpecs() []*LocalMountSpec
	SetPodCount(v int64) *JobSpec
	GetPodCount() *int64
	SetResourceConfig(v *ResourceConfig) *JobSpec
	GetResourceConfig() *ResourceConfig
	SetRestartPolicy(v string) *JobSpec
	GetRestartPolicy() *string
	SetServiceSpec(v *ServiceSpec) *JobSpec
	GetServiceSpec() *ServiceSpec
	SetSpotSpec(v *SpotSpec) *JobSpec
	GetSpotSpec() *SpotSpec
	SetSystemDisk(v *SystemDisk) *JobSpec
	GetSystemDisk() *SystemDisk
	SetType(v string) *JobSpec
	GetType() *string
	SetUseSpotInstance(v bool) *JobSpec
	GetUseSpotInstance() *bool
}

type JobSpec struct {
	AssignNodeSpec  *AssignNodeSpec  `json:"AssignNodeSpec,omitempty" xml:"AssignNodeSpec,omitempty"`
	AutoScalingSpec *AutoScalingSpec `json:"AutoScalingSpec,omitempty" xml:"AutoScalingSpec,omitempty"`
	// example:
	//
	// ecs.c6.large
	EcsSpec      *string       `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	ExtraPodSpec *ExtraPodSpec `json:"ExtraPodSpec,omitempty" xml:"ExtraPodSpec,omitempty"`
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/pai-dlc/tensorflow-training:1.12.2PAI-cpu-py27-ubuntu16.04
	Image       *string      `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageConfig *ImageConfig `json:"ImageConfig,omitempty" xml:"ImageConfig,omitempty"`
	// Deprecated
	IsCheif         *bool             `json:"IsCheif,omitempty" xml:"IsCheif,omitempty"`
	IsChief         *bool             `json:"IsChief,omitempty" xml:"IsChief,omitempty"`
	LocalMountSpecs []*LocalMountSpec `json:"LocalMountSpecs,omitempty" xml:"LocalMountSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PodCount       *int64          `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty"`
	RestartPolicy  *string         `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	ServiceSpec    *ServiceSpec    `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
	SpotSpec       *SpotSpec       `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty"`
	SystemDisk     *SystemDisk     `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Deprecated
	//
	// example:
	//
	// false
	UseSpotInstance *bool `json:"UseSpotInstance,omitempty" xml:"UseSpotInstance,omitempty"`
}

func (s JobSpec) String() string {
	return dara.Prettify(s)
}

func (s JobSpec) GoString() string {
	return s.String()
}

func (s *JobSpec) GetAssignNodeSpec() *AssignNodeSpec {
	return s.AssignNodeSpec
}

func (s *JobSpec) GetAutoScalingSpec() *AutoScalingSpec {
	return s.AutoScalingSpec
}

func (s *JobSpec) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *JobSpec) GetExtraPodSpec() *ExtraPodSpec {
	return s.ExtraPodSpec
}

func (s *JobSpec) GetImage() *string {
	return s.Image
}

func (s *JobSpec) GetImageConfig() *ImageConfig {
	return s.ImageConfig
}

func (s *JobSpec) GetIsCheif() *bool {
	return s.IsCheif
}

func (s *JobSpec) GetIsChief() *bool {
	return s.IsChief
}

func (s *JobSpec) GetLocalMountSpecs() []*LocalMountSpec {
	return s.LocalMountSpecs
}

func (s *JobSpec) GetPodCount() *int64 {
	return s.PodCount
}

func (s *JobSpec) GetResourceConfig() *ResourceConfig {
	return s.ResourceConfig
}

func (s *JobSpec) GetRestartPolicy() *string {
	return s.RestartPolicy
}

func (s *JobSpec) GetServiceSpec() *ServiceSpec {
	return s.ServiceSpec
}

func (s *JobSpec) GetSpotSpec() *SpotSpec {
	return s.SpotSpec
}

func (s *JobSpec) GetSystemDisk() *SystemDisk {
	return s.SystemDisk
}

func (s *JobSpec) GetType() *string {
	return s.Type
}

func (s *JobSpec) GetUseSpotInstance() *bool {
	return s.UseSpotInstance
}

func (s *JobSpec) SetAssignNodeSpec(v *AssignNodeSpec) *JobSpec {
	s.AssignNodeSpec = v
	return s
}

func (s *JobSpec) SetAutoScalingSpec(v *AutoScalingSpec) *JobSpec {
	s.AutoScalingSpec = v
	return s
}

func (s *JobSpec) SetEcsSpec(v string) *JobSpec {
	s.EcsSpec = &v
	return s
}

func (s *JobSpec) SetExtraPodSpec(v *ExtraPodSpec) *JobSpec {
	s.ExtraPodSpec = v
	return s
}

func (s *JobSpec) SetImage(v string) *JobSpec {
	s.Image = &v
	return s
}

func (s *JobSpec) SetImageConfig(v *ImageConfig) *JobSpec {
	s.ImageConfig = v
	return s
}

func (s *JobSpec) SetIsCheif(v bool) *JobSpec {
	s.IsCheif = &v
	return s
}

func (s *JobSpec) SetIsChief(v bool) *JobSpec {
	s.IsChief = &v
	return s
}

func (s *JobSpec) SetLocalMountSpecs(v []*LocalMountSpec) *JobSpec {
	s.LocalMountSpecs = v
	return s
}

func (s *JobSpec) SetPodCount(v int64) *JobSpec {
	s.PodCount = &v
	return s
}

func (s *JobSpec) SetResourceConfig(v *ResourceConfig) *JobSpec {
	s.ResourceConfig = v
	return s
}

func (s *JobSpec) SetRestartPolicy(v string) *JobSpec {
	s.RestartPolicy = &v
	return s
}

func (s *JobSpec) SetServiceSpec(v *ServiceSpec) *JobSpec {
	s.ServiceSpec = v
	return s
}

func (s *JobSpec) SetSpotSpec(v *SpotSpec) *JobSpec {
	s.SpotSpec = v
	return s
}

func (s *JobSpec) SetSystemDisk(v *SystemDisk) *JobSpec {
	s.SystemDisk = v
	return s
}

func (s *JobSpec) SetType(v string) *JobSpec {
	s.Type = &v
	return s
}

func (s *JobSpec) SetUseSpotInstance(v bool) *JobSpec {
	s.UseSpotInstance = &v
	return s
}

func (s *JobSpec) Validate() error {
	return dara.Validate(s)
}
