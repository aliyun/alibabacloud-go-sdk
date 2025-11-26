// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAckConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAckInstanceId(v string) *AckConfig
	GetAckInstanceId() *string
	SetCustomAnnotations(v []*Tag) *AckConfig
	GetCustomAnnotations() []*Tag
	SetCustomLabels(v []*Tag) *AckConfig
	GetCustomLabels() []*Tag
	SetDataDiskSize(v int32) *AckConfig
	GetDataDiskSize() *int32
	SetDataDiskStorageClass(v string) *AckConfig
	GetDataDiskStorageClass() *string
	SetLimitCpu(v float32) *AckConfig
	GetLimitCpu() *float32
	SetLimitMemory(v float32) *AckConfig
	GetLimitMemory() *float32
	SetMountHostCgroup(v bool) *AckConfig
	GetMountHostCgroup() *bool
	SetNamespace(v string) *AckConfig
	GetNamespace() *string
	SetNodeAffinity(v string) *AckConfig
	GetNodeAffinity() *string
	SetNodeSelectors(v []*Tag) *AckConfig
	GetNodeSelectors() []*Tag
	SetPodAffinity(v string) *AckConfig
	GetPodAffinity() *string
	SetPodAntiAffinity(v string) *AckConfig
	GetPodAntiAffinity() *string
	SetPreStartCommand(v []*string) *AckConfig
	GetPreStartCommand() []*string
	SetPvcs(v []*AckConfigPvcs) *AckConfig
	GetPvcs() []*AckConfigPvcs
	SetRequestCpu(v float32) *AckConfig
	GetRequestCpu() *float32
	SetRequestMemory(v float32) *AckConfig
	GetRequestMemory() *float32
	SetTolerations(v []*Toleration) *AckConfig
	GetTolerations() []*Toleration
	SetVolumeMounts(v []*AckConfigVolumeMounts) *AckConfig
	GetVolumeMounts() []*AckConfigVolumeMounts
	SetVolumes(v []*AckConfigVolumes) *AckConfig
	GetVolumes() []*AckConfigVolumes
}

type AckConfig struct {
	// ack集群id
	AckInstanceId        *string `json:"AckInstanceId,omitempty" xml:"AckInstanceId,omitempty"`
	CustomAnnotations    []*Tag  `json:"CustomAnnotations,omitempty" xml:"CustomAnnotations,omitempty" type:"Repeated"`
	CustomLabels         []*Tag  `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty" type:"Repeated"`
	DataDiskSize         *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DataDiskStorageClass *string `json:"DataDiskStorageClass,omitempty" xml:"DataDiskStorageClass,omitempty"`
	// Pod的CPU限制值。
	LimitCpu *float32 `json:"LimitCpu,omitempty" xml:"LimitCpu,omitempty"`
	// Pod的内存限制值。
	LimitMemory     *float32 `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	MountHostCgroup *bool    `json:"MountHostCgroup,omitempty" xml:"MountHostCgroup,omitempty"`
	// ack 命名空间
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NodeAffinity *string `json:"NodeAffinity,omitempty" xml:"NodeAffinity,omitempty"`
	// ack的节点标签限制
	NodeSelectors   []*Tag           `json:"NodeSelectors,omitempty" xml:"NodeSelectors,omitempty" type:"Repeated"`
	PodAffinity     *string          `json:"PodAffinity,omitempty" xml:"PodAffinity,omitempty"`
	PodAntiAffinity *string          `json:"PodAntiAffinity,omitempty" xml:"PodAntiAffinity,omitempty"`
	PreStartCommand []*string        `json:"PreStartCommand,omitempty" xml:"PreStartCommand,omitempty" type:"Repeated"`
	Pvcs            []*AckConfigPvcs `json:"Pvcs,omitempty" xml:"Pvcs,omitempty" type:"Repeated"`
	// Pod的CPU请求值
	RequestCpu *float32 `json:"RequestCpu,omitempty" xml:"RequestCpu,omitempty"`
	// Pod的内存请求值。
	RequestMemory *float32 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// ack的节点污点容忍
	Tolerations  []*Toleration            `json:"Tolerations,omitempty" xml:"Tolerations,omitempty" type:"Repeated"`
	VolumeMounts []*AckConfigVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	Volumes      []*AckConfigVolumes      `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s AckConfig) String() string {
	return dara.Prettify(s)
}

func (s AckConfig) GoString() string {
	return s.String()
}

func (s *AckConfig) GetAckInstanceId() *string {
	return s.AckInstanceId
}

func (s *AckConfig) GetCustomAnnotations() []*Tag {
	return s.CustomAnnotations
}

func (s *AckConfig) GetCustomLabels() []*Tag {
	return s.CustomLabels
}

func (s *AckConfig) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *AckConfig) GetDataDiskStorageClass() *string {
	return s.DataDiskStorageClass
}

func (s *AckConfig) GetLimitCpu() *float32 {
	return s.LimitCpu
}

func (s *AckConfig) GetLimitMemory() *float32 {
	return s.LimitMemory
}

func (s *AckConfig) GetMountHostCgroup() *bool {
	return s.MountHostCgroup
}

func (s *AckConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *AckConfig) GetNodeAffinity() *string {
	return s.NodeAffinity
}

func (s *AckConfig) GetNodeSelectors() []*Tag {
	return s.NodeSelectors
}

func (s *AckConfig) GetPodAffinity() *string {
	return s.PodAffinity
}

func (s *AckConfig) GetPodAntiAffinity() *string {
	return s.PodAntiAffinity
}

func (s *AckConfig) GetPreStartCommand() []*string {
	return s.PreStartCommand
}

func (s *AckConfig) GetPvcs() []*AckConfigPvcs {
	return s.Pvcs
}

func (s *AckConfig) GetRequestCpu() *float32 {
	return s.RequestCpu
}

func (s *AckConfig) GetRequestMemory() *float32 {
	return s.RequestMemory
}

func (s *AckConfig) GetTolerations() []*Toleration {
	return s.Tolerations
}

func (s *AckConfig) GetVolumeMounts() []*AckConfigVolumeMounts {
	return s.VolumeMounts
}

func (s *AckConfig) GetVolumes() []*AckConfigVolumes {
	return s.Volumes
}

func (s *AckConfig) SetAckInstanceId(v string) *AckConfig {
	s.AckInstanceId = &v
	return s
}

func (s *AckConfig) SetCustomAnnotations(v []*Tag) *AckConfig {
	s.CustomAnnotations = v
	return s
}

func (s *AckConfig) SetCustomLabels(v []*Tag) *AckConfig {
	s.CustomLabels = v
	return s
}

func (s *AckConfig) SetDataDiskSize(v int32) *AckConfig {
	s.DataDiskSize = &v
	return s
}

func (s *AckConfig) SetDataDiskStorageClass(v string) *AckConfig {
	s.DataDiskStorageClass = &v
	return s
}

func (s *AckConfig) SetLimitCpu(v float32) *AckConfig {
	s.LimitCpu = &v
	return s
}

func (s *AckConfig) SetLimitMemory(v float32) *AckConfig {
	s.LimitMemory = &v
	return s
}

func (s *AckConfig) SetMountHostCgroup(v bool) *AckConfig {
	s.MountHostCgroup = &v
	return s
}

func (s *AckConfig) SetNamespace(v string) *AckConfig {
	s.Namespace = &v
	return s
}

func (s *AckConfig) SetNodeAffinity(v string) *AckConfig {
	s.NodeAffinity = &v
	return s
}

func (s *AckConfig) SetNodeSelectors(v []*Tag) *AckConfig {
	s.NodeSelectors = v
	return s
}

func (s *AckConfig) SetPodAffinity(v string) *AckConfig {
	s.PodAffinity = &v
	return s
}

func (s *AckConfig) SetPodAntiAffinity(v string) *AckConfig {
	s.PodAntiAffinity = &v
	return s
}

func (s *AckConfig) SetPreStartCommand(v []*string) *AckConfig {
	s.PreStartCommand = v
	return s
}

func (s *AckConfig) SetPvcs(v []*AckConfigPvcs) *AckConfig {
	s.Pvcs = v
	return s
}

func (s *AckConfig) SetRequestCpu(v float32) *AckConfig {
	s.RequestCpu = &v
	return s
}

func (s *AckConfig) SetRequestMemory(v float32) *AckConfig {
	s.RequestMemory = &v
	return s
}

func (s *AckConfig) SetTolerations(v []*Toleration) *AckConfig {
	s.Tolerations = v
	return s
}

func (s *AckConfig) SetVolumeMounts(v []*AckConfigVolumeMounts) *AckConfig {
	s.VolumeMounts = v
	return s
}

func (s *AckConfig) SetVolumes(v []*AckConfigVolumes) *AckConfig {
	s.Volumes = v
	return s
}

func (s *AckConfig) Validate() error {
	if s.CustomAnnotations != nil {
		for _, item := range s.CustomAnnotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomLabels != nil {
		for _, item := range s.CustomLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeSelectors != nil {
		for _, item := range s.NodeSelectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pvcs != nil {
		for _, item := range s.Pvcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tolerations != nil {
		for _, item := range s.Tolerations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VolumeMounts != nil {
		for _, item := range s.VolumeMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Volumes != nil {
		for _, item := range s.Volumes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AckConfigPvcs struct {
	DataDiskSize         *int64  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DataDiskStorageClass *string `json:"DataDiskStorageClass,omitempty" xml:"DataDiskStorageClass,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path                 *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s AckConfigPvcs) String() string {
	return dara.Prettify(s)
}

func (s AckConfigPvcs) GoString() string {
	return s.String()
}

func (s *AckConfigPvcs) GetDataDiskSize() *int64 {
	return s.DataDiskSize
}

func (s *AckConfigPvcs) GetDataDiskStorageClass() *string {
	return s.DataDiskStorageClass
}

func (s *AckConfigPvcs) GetName() *string {
	return s.Name
}

func (s *AckConfigPvcs) GetPath() *string {
	return s.Path
}

func (s *AckConfigPvcs) SetDataDiskSize(v int64) *AckConfigPvcs {
	s.DataDiskSize = &v
	return s
}

func (s *AckConfigPvcs) SetDataDiskStorageClass(v string) *AckConfigPvcs {
	s.DataDiskStorageClass = &v
	return s
}

func (s *AckConfigPvcs) SetName(v string) *AckConfigPvcs {
	s.Name = &v
	return s
}

func (s *AckConfigPvcs) SetPath(v string) *AckConfigPvcs {
	s.Path = &v
	return s
}

func (s *AckConfigPvcs) Validate() error {
	return dara.Validate(s)
}

type AckConfigVolumeMounts struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s AckConfigVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s AckConfigVolumeMounts) GoString() string {
	return s.String()
}

func (s *AckConfigVolumeMounts) GetName() *string {
	return s.Name
}

func (s *AckConfigVolumeMounts) GetPath() *string {
	return s.Path
}

func (s *AckConfigVolumeMounts) SetName(v string) *AckConfigVolumeMounts {
	s.Name = &v
	return s
}

func (s *AckConfigVolumeMounts) SetPath(v string) *AckConfigVolumeMounts {
	s.Path = &v
	return s
}

func (s *AckConfigVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type AckConfigVolumes struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AckConfigVolumes) String() string {
	return dara.Prettify(s)
}

func (s AckConfigVolumes) GoString() string {
	return s.String()
}

func (s *AckConfigVolumes) GetName() *string {
	return s.Name
}

func (s *AckConfigVolumes) GetPath() *string {
	return s.Path
}

func (s *AckConfigVolumes) GetType() *string {
	return s.Type
}

func (s *AckConfigVolumes) SetName(v string) *AckConfigVolumes {
	s.Name = &v
	return s
}

func (s *AckConfigVolumes) SetPath(v string) *AckConfigVolumes {
	s.Path = &v
	return s
}

func (s *AckConfigVolumes) SetType(v string) *AckConfigVolumes {
	s.Type = &v
	return s
}

func (s *AckConfigVolumes) Validate() error {
	return dara.Validate(s)
}
