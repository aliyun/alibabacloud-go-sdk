// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobSettings interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedSettings(v map[string]interface{}) *JobSettings
	GetAdvancedSettings() map[string]interface{}
	SetAllocateAllRDMADevices(v bool) *JobSettings
	GetAllocateAllRDMADevices() *bool
	SetAllowUnschedulableNodes(v bool) *JobSettings
	GetAllowUnschedulableNodes() *bool
	SetBusinessUserId(v string) *JobSettings
	GetBusinessUserId() *string
	SetCaller(v string) *JobSettings
	GetCaller() *string
	SetDataJuicerConfig(v *DataJuicerConfig) *JobSettings
	GetDataJuicerConfig() *DataJuicerConfig
	SetDisableEcsStockCheck(v bool) *JobSettings
	GetDisableEcsStockCheck() *bool
	SetDriver(v string) *JobSettings
	GetDriver() *string
	SetElasticSpotJobMaxRestartTimes(v int32) *JobSettings
	GetElasticSpotJobMaxRestartTimes() *int32
	SetEnableCPUAffinity(v bool) *JobSettings
	GetEnableCPUAffinity() *bool
	SetEnableDSWDev(v bool) *JobSettings
	GetEnableDSWDev() *bool
	SetEnableErrorMonitoringInAIMaster(v bool) *JobSettings
	GetEnableErrorMonitoringInAIMaster() *bool
	SetEnableOssAppend(v bool) *JobSettings
	GetEnableOssAppend() *bool
	SetEnableRDMA(v bool) *JobSettings
	GetEnableRDMA() *bool
	SetEnableSanityCheck(v bool) *JobSettings
	GetEnableSanityCheck() *bool
	SetEnableTideResource(v bool) *JobSettings
	GetEnableTideResource() *bool
	SetErrorMonitoringArgs(v string) *JobSettings
	GetErrorMonitoringArgs() *string
	SetJobReservedMinutes(v int32) *JobSettings
	GetJobReservedMinutes() *int32
	SetJobReservedPolicy(v string) *JobSettings
	GetJobReservedPolicy() *string
	SetModelConfig(v *ModelConfig) *JobSettings
	GetModelConfig() *ModelConfig
	SetOversoldType(v string) *JobSettings
	GetOversoldType() *string
	SetPipelineId(v string) *JobSettings
	GetPipelineId() *string
	SetSanityCheckArgs(v string) *JobSettings
	GetSanityCheckArgs() *string
	SetShell(v string) *JobSettings
	GetShell() *string
	SetTags(v map[string]*string) *JobSettings
	GetTags() map[string]*string
	SetTerminationGracePeriodSeconds(v int64) *JobSettings
	GetTerminationGracePeriodSeconds() *int64
}

type JobSettings struct {
	// The extra advanced parameter settings.
	AdvancedSettings map[string]interface{} `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// Specifies whether to mount all RDMA network interfaces.
	AllocateAllRDMADevices  *bool `json:"AllocateAllRDMADevices,omitempty" xml:"AllocateAllRDMADevices,omitempty"`
	AllowUnschedulableNodes *bool `json:"AllowUnschedulableNodes,omitempty" xml:"AllowUnschedulableNodes,omitempty"`
	// The user ID associated with the job.
	//
	// example:
	//
	// 16****
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// The caller.
	//
	// example:
	//
	// SilkFlow
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The DataJuicer task configuration.
	DataJuicerConfig *DataJuicerConfig `json:"DataJuicerConfig,omitempty" xml:"DataJuicerConfig,omitempty"`
	// Specifies whether to skip inventory check. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	DisableEcsStockCheck *bool `json:"DisableEcsStockCheck,omitempty" xml:"DisableEcsStockCheck,omitempty"`
	// The NVIDIA driver configuration.
	//
	// example:
	//
	// 535.54.03
	Driver                        *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	ElasticSpotJobMaxRestartTimes *int32  `json:"ElasticSpotJobMaxRestartTimes,omitempty" xml:"ElasticSpotJobMaxRestartTimes,omitempty"`
	// The CPU affinity setting. This setting is effective only when using general computing subscription resources.
	//
	// example:
	//
	// true
	EnableCPUAffinity *bool `json:"EnableCPUAffinity,omitempty" xml:"EnableCPUAffinity,omitempty"`
	EnableDSWDev      *bool `json:"EnableDSWDev,omitempty" xml:"EnableDSWDev,omitempty"`
	// Specifies whether to enable fault tolerance monitoring for the job. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	EnableErrorMonitoringInAIMaster *bool `json:"EnableErrorMonitoringInAIMaster,omitempty" xml:"EnableErrorMonitoringInAIMaster,omitempty"`
	// Specifies whether to allow OSS append write. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	EnableOssAppend *bool `json:"EnableOssAppend,omitempty" xml:"EnableOssAppend,omitempty"`
	// Specifies whether to allow the job to use RDMA. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	EnableRDMA *bool `json:"EnableRDMA,omitempty" xml:"EnableRDMA,omitempty"`
	// Specifies whether to enable computing power health check for the job. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	EnableSanityCheck *bool `json:"EnableSanityCheck,omitempty" xml:"EnableSanityCheck,omitempty"`
	// Specifies whether to allow the job to use tidal resources. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	EnableTideResource *bool `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	// The configuration parameters for fault tolerance monitoring after it is enabled, such as whether to enable log hang-based detection.
	//
	// example:
	//
	// --enable-log-hang-detection true
	ErrorMonitoringArgs *string `json:"ErrorMonitoringArgs,omitempty" xml:"ErrorMonitoringArgs,omitempty"`
	// The retention duration after job completion, in minutes.
	//
	// example:
	//
	// 30
	JobReservedMinutes *int32 `json:"JobReservedMinutes,omitempty" xml:"JobReservedMinutes,omitempty"`
	// The retention policy after job completion.
	//
	// example:
	//
	// Always
	JobReservedPolicy *string `json:"JobReservedPolicy,omitempty" xml:"JobReservedPolicy,omitempty"`
	// The output model configuration. This parameter is currently effective only in joint training scenarios.
	ModelConfig *ModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// The oversold resource usage mode for the job (reject/accept/only accept).
	//
	// example:
	//
	// AcceptQuotaOverSold
	OversoldType *string `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// pid-12****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The configuration parameters for computing power health check.
	//
	// example:
	//
	// --sanity-check-timing=AfterJobFaultTolerant --sanity-check-timeout-ops=MarkJobFail
	SanityCheckArgs *string `json:"SanityCheckArgs,omitempty" xml:"SanityCheckArgs,omitempty"`
	// example:
	//
	// /bin/bash
	Shell *string `json:"Shell,omitempty" xml:"Shell,omitempty"`
	// The custom tags.
	Tags                          map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TerminationGracePeriodSeconds *int64             `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
}

func (s JobSettings) String() string {
	return dara.Prettify(s)
}

func (s JobSettings) GoString() string {
	return s.String()
}

func (s *JobSettings) GetAdvancedSettings() map[string]interface{} {
	return s.AdvancedSettings
}

func (s *JobSettings) GetAllocateAllRDMADevices() *bool {
	return s.AllocateAllRDMADevices
}

func (s *JobSettings) GetAllowUnschedulableNodes() *bool {
	return s.AllowUnschedulableNodes
}

func (s *JobSettings) GetBusinessUserId() *string {
	return s.BusinessUserId
}

func (s *JobSettings) GetCaller() *string {
	return s.Caller
}

func (s *JobSettings) GetDataJuicerConfig() *DataJuicerConfig {
	return s.DataJuicerConfig
}

func (s *JobSettings) GetDisableEcsStockCheck() *bool {
	return s.DisableEcsStockCheck
}

func (s *JobSettings) GetDriver() *string {
	return s.Driver
}

func (s *JobSettings) GetElasticSpotJobMaxRestartTimes() *int32 {
	return s.ElasticSpotJobMaxRestartTimes
}

func (s *JobSettings) GetEnableCPUAffinity() *bool {
	return s.EnableCPUAffinity
}

func (s *JobSettings) GetEnableDSWDev() *bool {
	return s.EnableDSWDev
}

func (s *JobSettings) GetEnableErrorMonitoringInAIMaster() *bool {
	return s.EnableErrorMonitoringInAIMaster
}

func (s *JobSettings) GetEnableOssAppend() *bool {
	return s.EnableOssAppend
}

func (s *JobSettings) GetEnableRDMA() *bool {
	return s.EnableRDMA
}

func (s *JobSettings) GetEnableSanityCheck() *bool {
	return s.EnableSanityCheck
}

func (s *JobSettings) GetEnableTideResource() *bool {
	return s.EnableTideResource
}

func (s *JobSettings) GetErrorMonitoringArgs() *string {
	return s.ErrorMonitoringArgs
}

func (s *JobSettings) GetJobReservedMinutes() *int32 {
	return s.JobReservedMinutes
}

func (s *JobSettings) GetJobReservedPolicy() *string {
	return s.JobReservedPolicy
}

func (s *JobSettings) GetModelConfig() *ModelConfig {
	return s.ModelConfig
}

func (s *JobSettings) GetOversoldType() *string {
	return s.OversoldType
}

func (s *JobSettings) GetPipelineId() *string {
	return s.PipelineId
}

func (s *JobSettings) GetSanityCheckArgs() *string {
	return s.SanityCheckArgs
}

func (s *JobSettings) GetShell() *string {
	return s.Shell
}

func (s *JobSettings) GetTags() map[string]*string {
	return s.Tags
}

func (s *JobSettings) GetTerminationGracePeriodSeconds() *int64 {
	return s.TerminationGracePeriodSeconds
}

func (s *JobSettings) SetAdvancedSettings(v map[string]interface{}) *JobSettings {
	s.AdvancedSettings = v
	return s
}

func (s *JobSettings) SetAllocateAllRDMADevices(v bool) *JobSettings {
	s.AllocateAllRDMADevices = &v
	return s
}

func (s *JobSettings) SetAllowUnschedulableNodes(v bool) *JobSettings {
	s.AllowUnschedulableNodes = &v
	return s
}

func (s *JobSettings) SetBusinessUserId(v string) *JobSettings {
	s.BusinessUserId = &v
	return s
}

func (s *JobSettings) SetCaller(v string) *JobSettings {
	s.Caller = &v
	return s
}

func (s *JobSettings) SetDataJuicerConfig(v *DataJuicerConfig) *JobSettings {
	s.DataJuicerConfig = v
	return s
}

func (s *JobSettings) SetDisableEcsStockCheck(v bool) *JobSettings {
	s.DisableEcsStockCheck = &v
	return s
}

func (s *JobSettings) SetDriver(v string) *JobSettings {
	s.Driver = &v
	return s
}

func (s *JobSettings) SetElasticSpotJobMaxRestartTimes(v int32) *JobSettings {
	s.ElasticSpotJobMaxRestartTimes = &v
	return s
}

func (s *JobSettings) SetEnableCPUAffinity(v bool) *JobSettings {
	s.EnableCPUAffinity = &v
	return s
}

func (s *JobSettings) SetEnableDSWDev(v bool) *JobSettings {
	s.EnableDSWDev = &v
	return s
}

func (s *JobSettings) SetEnableErrorMonitoringInAIMaster(v bool) *JobSettings {
	s.EnableErrorMonitoringInAIMaster = &v
	return s
}

func (s *JobSettings) SetEnableOssAppend(v bool) *JobSettings {
	s.EnableOssAppend = &v
	return s
}

func (s *JobSettings) SetEnableRDMA(v bool) *JobSettings {
	s.EnableRDMA = &v
	return s
}

func (s *JobSettings) SetEnableSanityCheck(v bool) *JobSettings {
	s.EnableSanityCheck = &v
	return s
}

func (s *JobSettings) SetEnableTideResource(v bool) *JobSettings {
	s.EnableTideResource = &v
	return s
}

func (s *JobSettings) SetErrorMonitoringArgs(v string) *JobSettings {
	s.ErrorMonitoringArgs = &v
	return s
}

func (s *JobSettings) SetJobReservedMinutes(v int32) *JobSettings {
	s.JobReservedMinutes = &v
	return s
}

func (s *JobSettings) SetJobReservedPolicy(v string) *JobSettings {
	s.JobReservedPolicy = &v
	return s
}

func (s *JobSettings) SetModelConfig(v *ModelConfig) *JobSettings {
	s.ModelConfig = v
	return s
}

func (s *JobSettings) SetOversoldType(v string) *JobSettings {
	s.OversoldType = &v
	return s
}

func (s *JobSettings) SetPipelineId(v string) *JobSettings {
	s.PipelineId = &v
	return s
}

func (s *JobSettings) SetSanityCheckArgs(v string) *JobSettings {
	s.SanityCheckArgs = &v
	return s
}

func (s *JobSettings) SetShell(v string) *JobSettings {
	s.Shell = &v
	return s
}

func (s *JobSettings) SetTags(v map[string]*string) *JobSettings {
	s.Tags = v
	return s
}

func (s *JobSettings) SetTerminationGracePeriodSeconds(v int64) *JobSettings {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *JobSettings) Validate() error {
	if s.DataJuicerConfig != nil {
		if err := s.DataJuicerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ModelConfig != nil {
		if err := s.ModelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
