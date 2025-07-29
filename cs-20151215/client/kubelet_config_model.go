// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKubeletConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedUnsafeSysctls(v []*string) *KubeletConfig
	GetAllowedUnsafeSysctls() []*string
	SetClusterDNS(v []*string) *KubeletConfig
	GetClusterDNS() []*string
	SetContainerLogMaxFiles(v int64) *KubeletConfig
	GetContainerLogMaxFiles() *int64
	SetContainerLogMaxSize(v string) *KubeletConfig
	GetContainerLogMaxSize() *string
	SetContainerLogMaxWorkers(v int32) *KubeletConfig
	GetContainerLogMaxWorkers() *int32
	SetContainerLogMonitorInterval(v string) *KubeletConfig
	GetContainerLogMonitorInterval() *string
	SetCpuCFSQuota(v bool) *KubeletConfig
	GetCpuCFSQuota() *bool
	SetCpuCFSQuotaPeriod(v string) *KubeletConfig
	GetCpuCFSQuotaPeriod() *string
	SetCpuManagerPolicy(v string) *KubeletConfig
	GetCpuManagerPolicy() *string
	SetEventBurst(v int64) *KubeletConfig
	GetEventBurst() *int64
	SetEventRecordQPS(v int64) *KubeletConfig
	GetEventRecordQPS() *int64
	SetEvictionHard(v map[string]interface{}) *KubeletConfig
	GetEvictionHard() map[string]interface{}
	SetEvictionSoft(v map[string]interface{}) *KubeletConfig
	GetEvictionSoft() map[string]interface{}
	SetEvictionSoftGracePeriod(v map[string]interface{}) *KubeletConfig
	GetEvictionSoftGracePeriod() map[string]interface{}
	SetFeatureGates(v map[string]interface{}) *KubeletConfig
	GetFeatureGates() map[string]interface{}
	SetImageGCHighThresholdPercent(v int32) *KubeletConfig
	GetImageGCHighThresholdPercent() *int32
	SetImageGCLowThresholdPercent(v int32) *KubeletConfig
	GetImageGCLowThresholdPercent() *int32
	SetKubeAPIBurst(v int64) *KubeletConfig
	GetKubeAPIBurst() *int64
	SetKubeAPIQPS(v int64) *KubeletConfig
	GetKubeAPIQPS() *int64
	SetKubeReserved(v map[string]interface{}) *KubeletConfig
	GetKubeReserved() map[string]interface{}
	SetMaxPods(v int64) *KubeletConfig
	GetMaxPods() *int64
	SetMemoryManagerPolicy(v string) *KubeletConfig
	GetMemoryManagerPolicy() *string
	SetPodPidsLimit(v int64) *KubeletConfig
	GetPodPidsLimit() *int64
	SetReadOnlyPort(v int64) *KubeletConfig
	GetReadOnlyPort() *int64
	SetRegistryBurst(v int64) *KubeletConfig
	GetRegistryBurst() *int64
	SetRegistryPullQPS(v int64) *KubeletConfig
	GetRegistryPullQPS() *int64
	SetReservedMemory(v []*KubeletConfigReservedMemory) *KubeletConfig
	GetReservedMemory() []*KubeletConfigReservedMemory
	SetSerializeImagePulls(v bool) *KubeletConfig
	GetSerializeImagePulls() *bool
	SetSystemReserved(v map[string]interface{}) *KubeletConfig
	GetSystemReserved() map[string]interface{}
	SetTopologyManagerPolicy(v string) *KubeletConfig
	GetTopologyManagerPolicy() *string
	SetTracing(v *KubeletConfigTracing) *KubeletConfig
	GetTracing() *KubeletConfigTracing
}

type KubeletConfig struct {
	AllowedUnsafeSysctls []*string `json:"allowedUnsafeSysctls,omitempty" xml:"allowedUnsafeSysctls,omitempty" type:"Repeated"`
	ClusterDNS           []*string `json:"clusterDNS,omitempty" xml:"clusterDNS,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	ContainerLogMaxFiles *int64 `json:"containerLogMaxFiles,omitempty" xml:"containerLogMaxFiles,omitempty"`
	// example:
	//
	// 10Mi
	ContainerLogMaxSize *string `json:"containerLogMaxSize,omitempty" xml:"containerLogMaxSize,omitempty"`
	// example:
	//
	// 1
	ContainerLogMaxWorkers *int32 `json:"containerLogMaxWorkers,omitempty" xml:"containerLogMaxWorkers,omitempty"`
	// example:
	//
	// 10s
	ContainerLogMonitorInterval *string `json:"containerLogMonitorInterval,omitempty" xml:"containerLogMonitorInterval,omitempty"`
	// example:
	//
	// true
	CpuCFSQuota *bool `json:"cpuCFSQuota,omitempty" xml:"cpuCFSQuota,omitempty"`
	// example:
	//
	// 100ms
	CpuCFSQuotaPeriod *string `json:"cpuCFSQuotaPeriod,omitempty" xml:"cpuCFSQuotaPeriod,omitempty"`
	// example:
	//
	// none
	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty" xml:"cpuManagerPolicy,omitempty"`
	// example:
	//
	// 10
	EventBurst *int64 `json:"eventBurst,omitempty" xml:"eventBurst,omitempty"`
	// example:
	//
	// 5
	EventRecordQPS          *int64                 `json:"eventRecordQPS,omitempty" xml:"eventRecordQPS,omitempty"`
	EvictionHard            map[string]interface{} `json:"evictionHard,omitempty" xml:"evictionHard,omitempty"`
	EvictionSoft            map[string]interface{} `json:"evictionSoft,omitempty" xml:"evictionSoft,omitempty"`
	EvictionSoftGracePeriod map[string]interface{} `json:"evictionSoftGracePeriod,omitempty" xml:"evictionSoftGracePeriod,omitempty"`
	FeatureGates            map[string]interface{} `json:"featureGates,omitempty" xml:"featureGates,omitempty"`
	// example:
	//
	// 85
	ImageGCHighThresholdPercent *int32 `json:"imageGCHighThresholdPercent,omitempty" xml:"imageGCHighThresholdPercent,omitempty"`
	// example:
	//
	// 80
	ImageGCLowThresholdPercent *int32 `json:"imageGCLowThresholdPercent,omitempty" xml:"imageGCLowThresholdPercent,omitempty"`
	// example:
	//
	// 10
	KubeAPIBurst *int64 `json:"kubeAPIBurst,omitempty" xml:"kubeAPIBurst,omitempty"`
	// example:
	//
	// 5
	KubeAPIQPS   *int64                 `json:"kubeAPIQPS,omitempty" xml:"kubeAPIQPS,omitempty"`
	KubeReserved map[string]interface{} `json:"kubeReserved,omitempty" xml:"kubeReserved,omitempty"`
	// example:
	//
	// 110
	MaxPods *int64 `json:"maxPods,omitempty" xml:"maxPods,omitempty"`
	// example:
	//
	// none
	MemoryManagerPolicy *string `json:"memoryManagerPolicy,omitempty" xml:"memoryManagerPolicy,omitempty"`
	// example:
	//
	// -1
	PodPidsLimit *int64 `json:"podPidsLimit,omitempty" xml:"podPidsLimit,omitempty"`
	// example:
	//
	// 0
	ReadOnlyPort *int64 `json:"readOnlyPort,omitempty" xml:"readOnlyPort,omitempty"`
	// example:
	//
	// 10
	RegistryBurst *int64 `json:"registryBurst,omitempty" xml:"registryBurst,omitempty"`
	// example:
	//
	// 5
	RegistryPullQPS *int64                         `json:"registryPullQPS,omitempty" xml:"registryPullQPS,omitempty"`
	ReservedMemory  []*KubeletConfigReservedMemory `json:"reservedMemory,omitempty" xml:"reservedMemory,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SerializeImagePulls *bool                  `json:"serializeImagePulls,omitempty" xml:"serializeImagePulls,omitempty"`
	SystemReserved      map[string]interface{} `json:"systemReserved,omitempty" xml:"systemReserved,omitempty"`
	// example:
	//
	// restricted
	TopologyManagerPolicy *string               `json:"topologyManagerPolicy,omitempty" xml:"topologyManagerPolicy,omitempty"`
	Tracing               *KubeletConfigTracing `json:"tracing,omitempty" xml:"tracing,omitempty" type:"Struct"`
}

func (s KubeletConfig) String() string {
	return dara.Prettify(s)
}

func (s KubeletConfig) GoString() string {
	return s.String()
}

func (s *KubeletConfig) GetAllowedUnsafeSysctls() []*string {
	return s.AllowedUnsafeSysctls
}

func (s *KubeletConfig) GetClusterDNS() []*string {
	return s.ClusterDNS
}

func (s *KubeletConfig) GetContainerLogMaxFiles() *int64 {
	return s.ContainerLogMaxFiles
}

func (s *KubeletConfig) GetContainerLogMaxSize() *string {
	return s.ContainerLogMaxSize
}

func (s *KubeletConfig) GetContainerLogMaxWorkers() *int32 {
	return s.ContainerLogMaxWorkers
}

func (s *KubeletConfig) GetContainerLogMonitorInterval() *string {
	return s.ContainerLogMonitorInterval
}

func (s *KubeletConfig) GetCpuCFSQuota() *bool {
	return s.CpuCFSQuota
}

func (s *KubeletConfig) GetCpuCFSQuotaPeriod() *string {
	return s.CpuCFSQuotaPeriod
}

func (s *KubeletConfig) GetCpuManagerPolicy() *string {
	return s.CpuManagerPolicy
}

func (s *KubeletConfig) GetEventBurst() *int64 {
	return s.EventBurst
}

func (s *KubeletConfig) GetEventRecordQPS() *int64 {
	return s.EventRecordQPS
}

func (s *KubeletConfig) GetEvictionHard() map[string]interface{} {
	return s.EvictionHard
}

func (s *KubeletConfig) GetEvictionSoft() map[string]interface{} {
	return s.EvictionSoft
}

func (s *KubeletConfig) GetEvictionSoftGracePeriod() map[string]interface{} {
	return s.EvictionSoftGracePeriod
}

func (s *KubeletConfig) GetFeatureGates() map[string]interface{} {
	return s.FeatureGates
}

func (s *KubeletConfig) GetImageGCHighThresholdPercent() *int32 {
	return s.ImageGCHighThresholdPercent
}

func (s *KubeletConfig) GetImageGCLowThresholdPercent() *int32 {
	return s.ImageGCLowThresholdPercent
}

func (s *KubeletConfig) GetKubeAPIBurst() *int64 {
	return s.KubeAPIBurst
}

func (s *KubeletConfig) GetKubeAPIQPS() *int64 {
	return s.KubeAPIQPS
}

func (s *KubeletConfig) GetKubeReserved() map[string]interface{} {
	return s.KubeReserved
}

func (s *KubeletConfig) GetMaxPods() *int64 {
	return s.MaxPods
}

func (s *KubeletConfig) GetMemoryManagerPolicy() *string {
	return s.MemoryManagerPolicy
}

func (s *KubeletConfig) GetPodPidsLimit() *int64 {
	return s.PodPidsLimit
}

func (s *KubeletConfig) GetReadOnlyPort() *int64 {
	return s.ReadOnlyPort
}

func (s *KubeletConfig) GetRegistryBurst() *int64 {
	return s.RegistryBurst
}

func (s *KubeletConfig) GetRegistryPullQPS() *int64 {
	return s.RegistryPullQPS
}

func (s *KubeletConfig) GetReservedMemory() []*KubeletConfigReservedMemory {
	return s.ReservedMemory
}

func (s *KubeletConfig) GetSerializeImagePulls() *bool {
	return s.SerializeImagePulls
}

func (s *KubeletConfig) GetSystemReserved() map[string]interface{} {
	return s.SystemReserved
}

func (s *KubeletConfig) GetTopologyManagerPolicy() *string {
	return s.TopologyManagerPolicy
}

func (s *KubeletConfig) GetTracing() *KubeletConfigTracing {
	return s.Tracing
}

func (s *KubeletConfig) SetAllowedUnsafeSysctls(v []*string) *KubeletConfig {
	s.AllowedUnsafeSysctls = v
	return s
}

func (s *KubeletConfig) SetClusterDNS(v []*string) *KubeletConfig {
	s.ClusterDNS = v
	return s
}

func (s *KubeletConfig) SetContainerLogMaxFiles(v int64) *KubeletConfig {
	s.ContainerLogMaxFiles = &v
	return s
}

func (s *KubeletConfig) SetContainerLogMaxSize(v string) *KubeletConfig {
	s.ContainerLogMaxSize = &v
	return s
}

func (s *KubeletConfig) SetContainerLogMaxWorkers(v int32) *KubeletConfig {
	s.ContainerLogMaxWorkers = &v
	return s
}

func (s *KubeletConfig) SetContainerLogMonitorInterval(v string) *KubeletConfig {
	s.ContainerLogMonitorInterval = &v
	return s
}

func (s *KubeletConfig) SetCpuCFSQuota(v bool) *KubeletConfig {
	s.CpuCFSQuota = &v
	return s
}

func (s *KubeletConfig) SetCpuCFSQuotaPeriod(v string) *KubeletConfig {
	s.CpuCFSQuotaPeriod = &v
	return s
}

func (s *KubeletConfig) SetCpuManagerPolicy(v string) *KubeletConfig {
	s.CpuManagerPolicy = &v
	return s
}

func (s *KubeletConfig) SetEventBurst(v int64) *KubeletConfig {
	s.EventBurst = &v
	return s
}

func (s *KubeletConfig) SetEventRecordQPS(v int64) *KubeletConfig {
	s.EventRecordQPS = &v
	return s
}

func (s *KubeletConfig) SetEvictionHard(v map[string]interface{}) *KubeletConfig {
	s.EvictionHard = v
	return s
}

func (s *KubeletConfig) SetEvictionSoft(v map[string]interface{}) *KubeletConfig {
	s.EvictionSoft = v
	return s
}

func (s *KubeletConfig) SetEvictionSoftGracePeriod(v map[string]interface{}) *KubeletConfig {
	s.EvictionSoftGracePeriod = v
	return s
}

func (s *KubeletConfig) SetFeatureGates(v map[string]interface{}) *KubeletConfig {
	s.FeatureGates = v
	return s
}

func (s *KubeletConfig) SetImageGCHighThresholdPercent(v int32) *KubeletConfig {
	s.ImageGCHighThresholdPercent = &v
	return s
}

func (s *KubeletConfig) SetImageGCLowThresholdPercent(v int32) *KubeletConfig {
	s.ImageGCLowThresholdPercent = &v
	return s
}

func (s *KubeletConfig) SetKubeAPIBurst(v int64) *KubeletConfig {
	s.KubeAPIBurst = &v
	return s
}

func (s *KubeletConfig) SetKubeAPIQPS(v int64) *KubeletConfig {
	s.KubeAPIQPS = &v
	return s
}

func (s *KubeletConfig) SetKubeReserved(v map[string]interface{}) *KubeletConfig {
	s.KubeReserved = v
	return s
}

func (s *KubeletConfig) SetMaxPods(v int64) *KubeletConfig {
	s.MaxPods = &v
	return s
}

func (s *KubeletConfig) SetMemoryManagerPolicy(v string) *KubeletConfig {
	s.MemoryManagerPolicy = &v
	return s
}

func (s *KubeletConfig) SetPodPidsLimit(v int64) *KubeletConfig {
	s.PodPidsLimit = &v
	return s
}

func (s *KubeletConfig) SetReadOnlyPort(v int64) *KubeletConfig {
	s.ReadOnlyPort = &v
	return s
}

func (s *KubeletConfig) SetRegistryBurst(v int64) *KubeletConfig {
	s.RegistryBurst = &v
	return s
}

func (s *KubeletConfig) SetRegistryPullQPS(v int64) *KubeletConfig {
	s.RegistryPullQPS = &v
	return s
}

func (s *KubeletConfig) SetReservedMemory(v []*KubeletConfigReservedMemory) *KubeletConfig {
	s.ReservedMemory = v
	return s
}

func (s *KubeletConfig) SetSerializeImagePulls(v bool) *KubeletConfig {
	s.SerializeImagePulls = &v
	return s
}

func (s *KubeletConfig) SetSystemReserved(v map[string]interface{}) *KubeletConfig {
	s.SystemReserved = v
	return s
}

func (s *KubeletConfig) SetTopologyManagerPolicy(v string) *KubeletConfig {
	s.TopologyManagerPolicy = &v
	return s
}

func (s *KubeletConfig) SetTracing(v *KubeletConfigTracing) *KubeletConfig {
	s.Tracing = v
	return s
}

func (s *KubeletConfig) Validate() error {
	return dara.Validate(s)
}

type KubeletConfigReservedMemory struct {
	Limits   map[string]interface{} `json:"limits,omitempty" xml:"limits,omitempty"`
	NumaNode *int32                 `json:"numaNode,omitempty" xml:"numaNode,omitempty"`
}

func (s KubeletConfigReservedMemory) String() string {
	return dara.Prettify(s)
}

func (s KubeletConfigReservedMemory) GoString() string {
	return s.String()
}

func (s *KubeletConfigReservedMemory) GetLimits() map[string]interface{} {
	return s.Limits
}

func (s *KubeletConfigReservedMemory) GetNumaNode() *int32 {
	return s.NumaNode
}

func (s *KubeletConfigReservedMemory) SetLimits(v map[string]interface{}) *KubeletConfigReservedMemory {
	s.Limits = v
	return s
}

func (s *KubeletConfigReservedMemory) SetNumaNode(v int32) *KubeletConfigReservedMemory {
	s.NumaNode = &v
	return s
}

func (s *KubeletConfigReservedMemory) Validate() error {
	return dara.Validate(s)
}

type KubeletConfigTracing struct {
	// example:
	//
	// localhost:4317
	Endpoint               *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	SamplingRatePerMillion *int32  `json:"samplingRatePerMillion,omitempty" xml:"samplingRatePerMillion,omitempty"`
}

func (s KubeletConfigTracing) String() string {
	return dara.Prettify(s)
}

func (s KubeletConfigTracing) GoString() string {
	return s.String()
}

func (s *KubeletConfigTracing) GetEndpoint() *string {
	return s.Endpoint
}

func (s *KubeletConfigTracing) GetSamplingRatePerMillion() *int32 {
	return s.SamplingRatePerMillion
}

func (s *KubeletConfigTracing) SetEndpoint(v string) *KubeletConfigTracing {
	s.Endpoint = &v
	return s
}

func (s *KubeletConfigTracing) SetSamplingRatePerMillion(v int32) *KubeletConfigTracing {
	s.SamplingRatePerMillion = &v
	return s
}

func (s *KubeletConfigTracing) Validate() error {
	return dara.Validate(s)
}
