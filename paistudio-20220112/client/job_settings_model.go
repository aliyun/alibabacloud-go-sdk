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
	SetBusinessUserId(v string) *JobSettings
	GetBusinessUserId() *string
	SetCaller(v string) *JobSettings
	GetCaller() *string
	SetDriver(v string) *JobSettings
	GetDriver() *string
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
	SetOversoldType(v string) *JobSettings
	GetOversoldType() *string
	SetPipelineId(v string) *JobSettings
	GetPipelineId() *string
	SetSanityCheckArgs(v string) *JobSettings
	GetSanityCheckArgs() *string
	SetTags(v map[string]*string) *JobSettings
	GetTags() map[string]*string
}

type JobSettings struct {
	// Additional advanced parameter settings.
	AdvancedSettings map[string]interface{} `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// User ID associated with the job.
	//
	// example:
	//
	// 166924
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// Caller.
	//
	// example:
	//
	// SilkFlow
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// NVIDIA driver configuration.
	//
	// example:
	//
	// 535.54.03
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// Enable job fault tolerance monitoring.
	//
	// example:
	//
	// false
	EnableErrorMonitoringInAIMaster *bool `json:"EnableErrorMonitoringInAIMaster,omitempty" xml:"EnableErrorMonitoringInAIMaster,omitempty"`
	// Enable OSS append writes.
	//
	// example:
	//
	// true
	EnableOssAppend *bool `json:"EnableOssAppend,omitempty" xml:"EnableOssAppend,omitempty"`
	// Enable the job to use RDMA.
	//
	// example:
	//
	// true
	EnableRDMA *bool `json:"EnableRDMA,omitempty" xml:"EnableRDMA,omitempty"`
	// Enable computing power health check for the job.
	//
	// example:
	//
	// true
	EnableSanityCheck *bool `json:"EnableSanityCheck,omitempty" xml:"EnableSanityCheck,omitempty"`
	// Enable the job to use tide resources.
	//
	// example:
	//
	// true
	EnableTideResource *bool `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	// Specify configuration parameters for fault tolerance monitoring, such as whether to enable log hang-based detection.
	//
	// example:
	//
	// --enable-log-hang-detection true
	ErrorMonitoringArgs *string `json:"ErrorMonitoringArgs,omitempty" xml:"ErrorMonitoringArgs,omitempty"`
	// Duration in minutes to retain the job after completion.
	//
	// example:
	//
	// 30
	JobReservedMinutes *int32 `json:"JobReservedMinutes,omitempty" xml:"JobReservedMinutes,omitempty"`
	// Policy for retaining the job after completion.
	//
	// example:
	//
	// Always
	JobReservedPolicy *string `json:"JobReservedPolicy,omitempty" xml:"JobReservedPolicy,omitempty"`
	// Job\\"s oversold resource usage method (not accepted, accepted, or only accepted).
	//
	// example:
	//
	// AcceptQuotaOverSold
	OversoldType *string `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
	// Workflow ID.
	//
	// example:
	//
	// pid-123456
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// Configuration parameters for computing power health check.
	//
	// example:
	//
	// --sanity-check-timing=AfterJobFaultTolerant --sanity-check-timeout-ops=MarkJobFai
	SanityCheckArgs *string `json:"SanityCheckArgs,omitempty" xml:"SanityCheckArgs,omitempty"`
	// Custom tags.
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *JobSettings) GetBusinessUserId() *string {
	return s.BusinessUserId
}

func (s *JobSettings) GetCaller() *string {
	return s.Caller
}

func (s *JobSettings) GetDriver() *string {
	return s.Driver
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

func (s *JobSettings) GetOversoldType() *string {
	return s.OversoldType
}

func (s *JobSettings) GetPipelineId() *string {
	return s.PipelineId
}

func (s *JobSettings) GetSanityCheckArgs() *string {
	return s.SanityCheckArgs
}

func (s *JobSettings) GetTags() map[string]*string {
	return s.Tags
}

func (s *JobSettings) SetAdvancedSettings(v map[string]interface{}) *JobSettings {
	s.AdvancedSettings = v
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

func (s *JobSettings) SetDriver(v string) *JobSettings {
	s.Driver = &v
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

func (s *JobSettings) SetTags(v map[string]*string) *JobSettings {
	s.Tags = v
	return s
}

func (s *JobSettings) Validate() error {
	return dara.Validate(s)
}
