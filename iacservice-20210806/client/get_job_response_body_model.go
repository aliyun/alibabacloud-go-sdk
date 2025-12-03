// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetJobResponseBodyJob) *GetJobResponseBody
	GetJob() *GetJobResponseBodyJob
	SetRequestId(v string) *GetJobResponseBody
	GetRequestId() *string
}

type GetJobResponseBody struct {
	Job *GetJobResponseBodyJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// example:
	//
	// 1435C78A-AED9-53D6-B7A6-E2661D29B1FA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) GetJob() *GetJobResponseBodyJob {
	return s.Job
}

func (s *GetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobResponseBody) SetJob(v *GetJobResponseBodyJob) *GetJobResponseBody {
	s.Job = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobResponseBodyJob struct {
	AssertCheckDetail []*GetJobResponseBodyJobAssertCheckDetail `json:"assertCheckDetail,omitempty" xml:"assertCheckDetail,omitempty" type:"Repeated"`
	Config            *GetJobResponseBodyJobConfig              `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 2022-08-31T03:38:40Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test1
	Description *string                `json:"description,omitempty" xml:"description,omitempty"`
	DownloadUrl map[string]interface{} `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	ElapsedTime *int64                 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	ExecuteType *string                `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// example:
	//
	// true
	IsPassAssertCheck *bool `json:"isPassAssertCheck,omitempty" xml:"isPassAssertCheck,omitempty"`
	// example:
	//
	// job-518855d9a058cfff0dc933e6b5767
	JobId   *string                `json:"jobId,omitempty" xml:"jobId,omitempty"`
	LogFile map[string]interface{} `json:"logFile,omitempty" xml:"logFile,omitempty"`
	// example:
	//
	// /
	Output         *string            `json:"output,omitempty" xml:"output,omitempty"`
	OutputJsonPlan interface{}        `json:"outputJsonPlan,omitempty" xml:"outputJsonPlan,omitempty"`
	Parameters     map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// Errored
	Status       *string                          `json:"status,omitempty" xml:"status,omitempty"`
	StatusDetail map[string]*JobStatusDetailValue `json:"statusDetail,omitempty" xml:"statusDetail,omitempty"`
	// example:
	//
	// task-3b6cb9fa4751a1b9b5f22cbcf4e
	TaskId                   *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskType                 *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s GetJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJob) GetAssertCheckDetail() []*GetJobResponseBodyJobAssertCheckDetail {
	return s.AssertCheckDetail
}

func (s *GetJobResponseBodyJob) GetConfig() *GetJobResponseBodyJobConfig {
	return s.Config
}

func (s *GetJobResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobResponseBodyJob) GetDescription() *string {
	return s.Description
}

func (s *GetJobResponseBodyJob) GetDownloadUrl() map[string]interface{} {
	return s.DownloadUrl
}

func (s *GetJobResponseBodyJob) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *GetJobResponseBodyJob) GetExecuteType() *string {
	return s.ExecuteType
}

func (s *GetJobResponseBodyJob) GetIsPassAssertCheck() *bool {
	return s.IsPassAssertCheck
}

func (s *GetJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetJobResponseBodyJob) GetLogFile() map[string]interface{} {
	return s.LogFile
}

func (s *GetJobResponseBodyJob) GetOutput() *string {
	return s.Output
}

func (s *GetJobResponseBodyJob) GetOutputJsonPlan() interface{} {
	return s.OutputJsonPlan
}

func (s *GetJobResponseBodyJob) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *GetJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetJobResponseBodyJob) GetStatusDetail() map[string]*JobStatusDetailValue {
	return s.StatusDetail
}

func (s *GetJobResponseBodyJob) GetTaskId() *string {
	return s.TaskId
}

func (s *GetJobResponseBodyJob) GetTaskType() *string {
	return s.TaskType
}

func (s *GetJobResponseBodyJob) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *GetJobResponseBodyJob) SetAssertCheckDetail(v []*GetJobResponseBodyJobAssertCheckDetail) *GetJobResponseBodyJob {
	s.AssertCheckDetail = v
	return s
}

func (s *GetJobResponseBodyJob) SetConfig(v *GetJobResponseBodyJobConfig) *GetJobResponseBodyJob {
	s.Config = v
	return s
}

func (s *GetJobResponseBodyJob) SetCreateTime(v string) *GetJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyJob) SetDescription(v string) *GetJobResponseBodyJob {
	s.Description = &v
	return s
}

func (s *GetJobResponseBodyJob) SetDownloadUrl(v map[string]interface{}) *GetJobResponseBodyJob {
	s.DownloadUrl = v
	return s
}

func (s *GetJobResponseBodyJob) SetElapsedTime(v int64) *GetJobResponseBodyJob {
	s.ElapsedTime = &v
	return s
}

func (s *GetJobResponseBodyJob) SetExecuteType(v string) *GetJobResponseBodyJob {
	s.ExecuteType = &v
	return s
}

func (s *GetJobResponseBodyJob) SetIsPassAssertCheck(v bool) *GetJobResponseBodyJob {
	s.IsPassAssertCheck = &v
	return s
}

func (s *GetJobResponseBodyJob) SetJobId(v string) *GetJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBodyJob) SetLogFile(v map[string]interface{}) *GetJobResponseBodyJob {
	s.LogFile = v
	return s
}

func (s *GetJobResponseBodyJob) SetOutput(v string) *GetJobResponseBodyJob {
	s.Output = &v
	return s
}

func (s *GetJobResponseBodyJob) SetOutputJsonPlan(v interface{}) *GetJobResponseBodyJob {
	s.OutputJsonPlan = v
	return s
}

func (s *GetJobResponseBodyJob) SetParameters(v map[string]*string) *GetJobResponseBodyJob {
	s.Parameters = v
	return s
}

func (s *GetJobResponseBodyJob) SetStatus(v string) *GetJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyJob) SetStatusDetail(v map[string]*JobStatusDetailValue) *GetJobResponseBodyJob {
	s.StatusDetail = v
	return s
}

func (s *GetJobResponseBodyJob) SetTaskId(v string) *GetJobResponseBodyJob {
	s.TaskId = &v
	return s
}

func (s *GetJobResponseBodyJob) SetTaskType(v string) *GetJobResponseBodyJob {
	s.TaskType = &v
	return s
}

func (s *GetJobResponseBodyJob) SetTerraformProviderVersion(v string) *GetJobResponseBodyJob {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetJobResponseBodyJob) Validate() error {
	if s.AssertCheckDetail != nil {
		for _, item := range s.AssertCheckDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobResponseBodyJobAssertCheckDetail struct {
	// example:
	//
	// eq
	Comparison    *string `json:"comparison,omitempty" xml:"comparison,omitempty"`
	ExpectedValue *string `json:"expectedValue,omitempty" xml:"expectedValue,omitempty"`
	// example:
	//
	// true
	IsPass *bool `json:"isPass,omitempty" xml:"isPass,omitempty"`
	// example:
	//
	// result
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetJobResponseBodyJobAssertCheckDetail) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobAssertCheckDetail) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobAssertCheckDetail) GetComparison() *string {
	return s.Comparison
}

func (s *GetJobResponseBodyJobAssertCheckDetail) GetExpectedValue() *string {
	return s.ExpectedValue
}

func (s *GetJobResponseBodyJobAssertCheckDetail) GetIsPass() *bool {
	return s.IsPass
}

func (s *GetJobResponseBodyJobAssertCheckDetail) GetType() *string {
	return s.Type
}

func (s *GetJobResponseBodyJobAssertCheckDetail) SetComparison(v string) *GetJobResponseBodyJobAssertCheckDetail {
	s.Comparison = &v
	return s
}

func (s *GetJobResponseBodyJobAssertCheckDetail) SetExpectedValue(v string) *GetJobResponseBodyJobAssertCheckDetail {
	s.ExpectedValue = &v
	return s
}

func (s *GetJobResponseBodyJobAssertCheckDetail) SetIsPass(v bool) *GetJobResponseBodyJobAssertCheckDetail {
	s.IsPass = &v
	return s
}

func (s *GetJobResponseBodyJobAssertCheckDetail) SetType(v string) *GetJobResponseBodyJobAssertCheckDetail {
	s.Type = &v
	return s
}

func (s *GetJobResponseBodyJobAssertCheckDetail) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobConfig struct {
	// example:
	//
	// true
	AutoApply          *bool   `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	HasConfigProactive *string `json:"hasConfigProactive,omitempty" xml:"hasConfigProactive,omitempty"`
	// example:
	//
	// fales
	IsDestroy *bool `json:"isDestroy,omitempty" xml:"isDestroy,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// {}
	ResourcesChanged *string `json:"resourcesChanged,omitempty" xml:"resourcesChanged,omitempty"`
	SubCommand       *string `json:"subCommand,omitempty" xml:"subCommand,omitempty"`
}

func (s GetJobResponseBodyJobConfig) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobConfig) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobConfig) GetAutoApply() *bool {
	return s.AutoApply
}

func (s *GetJobResponseBodyJobConfig) GetHasConfigProactive() *string {
	return s.HasConfigProactive
}

func (s *GetJobResponseBodyJobConfig) GetIsDestroy() *bool {
	return s.IsDestroy
}

func (s *GetJobResponseBodyJobConfig) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *GetJobResponseBodyJobConfig) GetResourcesChanged() *string {
	return s.ResourcesChanged
}

func (s *GetJobResponseBodyJobConfig) GetSubCommand() *string {
	return s.SubCommand
}

func (s *GetJobResponseBodyJobConfig) SetAutoApply(v bool) *GetJobResponseBodyJobConfig {
	s.AutoApply = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetHasConfigProactive(v string) *GetJobResponseBodyJobConfig {
	s.HasConfigProactive = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetIsDestroy(v bool) *GetJobResponseBodyJobConfig {
	s.IsDestroy = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetModuleVersion(v string) *GetJobResponseBodyJobConfig {
	s.ModuleVersion = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetResourcesChanged(v string) *GetJobResponseBodyJobConfig {
	s.ResourcesChanged = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetSubCommand(v string) *GetJobResponseBodyJobConfig {
	s.SubCommand = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) Validate() error {
	return dara.Validate(s)
}
