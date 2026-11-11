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
	// The job details.
	Job *GetJobResponseBodyJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// The request ID.
	//
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
	// The full set of job parameters, including sensitive and non-sensitive parameters.
	AllParameters []*GetJobResponseBodyJobAllParameters `json:"allParameters,omitempty" xml:"allParameters,omitempty" type:"Repeated"`
	// The list of assertion checks, applicable to scenario-based testing tasks.
	AssertCheckDetail []*GetJobResponseBodyJobAssertCheckDetail `json:"assertCheckDetail,omitempty" xml:"assertCheckDetail,omitempty" type:"Repeated"`
	// The job configuration.
	Config *GetJobResponseBodyJobConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The job creation time.
	//
	// example:
	//
	// 2022-08-31T03:38:40Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The job description.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The download URL.
	//
	// example:
	//
	// url
	DownloadUrl map[string]interface{} `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 3s
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// The execution type. Valid values:
	//
	// - Manual: Manual execution (default).
	//
	// - Auto: Automatic execution.
	//
	// example:
	//
	// Manual
	ExecuteType *string `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// Indicates whether the assertion check is passed.
	//
	// example:
	//
	// true
	IsPassAssertCheck *bool `json:"isPassAssertCheck,omitempty" xml:"isPassAssertCheck,omitempty"`
	// The job ID.
	//
	// example:
	//
	// job-518855d9a058cfff0dc933e6b5767
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The job type.
	//
	// example:
	//
	// Default
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// The runtime logs. Currently, the following log content (that is, the values of the key) can be returned:
	//
	// - tf-init.run.error.log
	//
	// - tf-init.plan.log
	//
	// - tf-plan.run.log
	//
	// - tf-apply.run.log
	//
	// - tf-init.apply.log
	LogFile map[string]interface{} `json:"logFile,omitempty" xml:"logFile,omitempty"`
	// The job output.
	//
	// example:
	//
	// /
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// The change details of the Plan phase.
	//
	// example:
	//
	// {
	//
	//     "formatVersion": "1.2",
	//
	//     "terraformVersion": "1.5.7",
	//
	//     "providerVersion": "1.262.1",
	//
	//     "plannedValues": {
	//
	//         "root_module": {
	//
	//         }
	//
	//     },
	//
	//     "resourceChanges": [
	//
	//         {
	//
	//             "address": "alicloud_instance.uuid_ae98dda8_xxxxxxx",
	//
	//             "mode": "managed",
	//
	//             "type": "alicloud_instance",
	//
	//             "name": "uuid_ae98dda8_xxxxxx",
	//
	//             "providerName": "registry.terraform.io/aliyun/alicloud",
	//
	//             "change": {
	//
	//                 "actions": [
	//
	//                     "delete"
	//
	//                 ],
	//
	//                 "before": Object{...},
	//
	//                 "after_unknown": {
	//
	//                 },
	//
	//                 "before_sensitive": Object{...},
	//
	//                 "after_sensitive": false
	//
	//             },
	//
	//             "cloudSpecResourceCode": "ACS::ECS::Instance"
	//
	//         }
	//
	//     ],
	//
	//     "configuration": Object{...}
	//
	// }
	OutputJsonPlan interface{} `json:"outputJsonPlan,omitempty" xml:"outputJsonPlan,omitempty"`
	// The parameter collection.
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The job status. Valid values:
	//
	// - Pending: The initial state after the job is created.
	//
	// - PlanQueued: The state where the job is queued and waiting because no workflow is available after creation.
	//
	// - Planning: The state where the job is in the Plan phase.
	//
	// - ConfigProactiveInProgress: The state where the compliance pre-check is in progress. This requires the compliance pre-check feature to be enabled for the account.
	//
	// - ConfigProactiveSuccess: The state where the compliance pre-check is successful. This requires the compliance pre-check feature to be enabled for the account.
	//
	// - Planned: The state where the job has completed the Plan phase.
	//
	// - PlannedAndFinished: The final state where no differences are found after the Plan phase is completed.
	//
	// - Confirmed: The state where the job is waiting for confirmation after the Plan phase is completed.
	//
	// - ApplyQueued: The state where the job is queued and waiting because no workflow is available during execution.
	//
	// - Applying: The state where the job is in the Apply phase.
	//
	// - Applied: The final state where the job has completed the Apply phase.
	//
	// - Errored: The final state where an error occurs during job execution.
	//
	// - Canceled: The final state where the job execution is canceled.
	//
	// - Discarded: The final state where the plan of the job is discarded.
	//
	// - ConfigProactiveFailure: The state where the compliance pre-check fails. This requires the compliance pre-check feature to be enabled for the account.
	//
	// example:
	//
	// Errored
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The job status details.
	StatusDetail map[string]*JobStatusDetailValue `json:"statusDetail,omitempty" xml:"statusDetail,omitempty"`
	// The task ID.
	//
	// example:
	//
	// task-3b6cb9fa4751a1b9b5f22cbcf4e
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The task type. Valid values:
	//
	// - Task: Regular task (default).
	//
	// - SceneTestingTask: Scenario-based testing task.
	//
	// example:
	//
	// SceneTestingTask
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// The Terraform provider version.
	//
	// example:
	//
	// 1.230.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s GetJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJob) GetAllParameters() []*GetJobResponseBodyJobAllParameters {
	return s.AllParameters
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

func (s *GetJobResponseBodyJob) GetJobType() *string {
	return s.JobType
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

func (s *GetJobResponseBodyJob) SetAllParameters(v []*GetJobResponseBodyJobAllParameters) *GetJobResponseBodyJob {
	s.AllParameters = v
	return s
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

func (s *GetJobResponseBodyJob) SetJobType(v string) *GetJobResponseBodyJob {
	s.JobType = &v
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
	if s.AllParameters != nil {
		for _, item := range s.AllParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type GetJobResponseBodyJobAllParameters struct {
	// The parameter name.
	//
	// example:
	//
	// regionId
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether the parameter is confidential. Confidential parameters are hidden in API responses and console displays, and are stored in an encrypted manner.
	//
	// example:
	//
	// false
	Secret *bool `json:"secret,omitempty" xml:"secret,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// cn-zhangjiakou
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetJobResponseBodyJobAllParameters) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobAllParameters) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobAllParameters) GetName() *string {
	return s.Name
}

func (s *GetJobResponseBodyJobAllParameters) GetSecret() *bool {
	return s.Secret
}

func (s *GetJobResponseBodyJobAllParameters) GetValue() *string {
	return s.Value
}

func (s *GetJobResponseBodyJobAllParameters) SetName(v string) *GetJobResponseBodyJobAllParameters {
	s.Name = &v
	return s
}

func (s *GetJobResponseBodyJobAllParameters) SetSecret(v bool) *GetJobResponseBodyJobAllParameters {
	s.Secret = &v
	return s
}

func (s *GetJobResponseBodyJobAllParameters) SetValue(v string) *GetJobResponseBodyJobAllParameters {
	s.Value = &v
	return s
}

func (s *GetJobResponseBodyJobAllParameters) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobAssertCheckDetail struct {
	// The comparison operator. Valid values:
	//
	// - eq: Equal to.
	//
	// - n_eq: Not equal to.
	//
	// - ctn: Contains.
	//
	// - n_ctn: Does not contain.
	//
	// - regex: Regular expression match.
	//
	// example:
	//
	// eq
	Comparison *string `json:"comparison,omitempty" xml:"comparison,omitempty"`
	// The expected value.
	//
	// example:
	//
	// Expected value
	ExpectedValue *string `json:"expectedValue,omitempty" xml:"expectedValue,omitempty"`
	// Indicates whether the check is passed.
	//
	// example:
	//
	// true
	IsPass *bool `json:"isPass,omitempty" xml:"isPass,omitempty"`
	// The assertion type. Valid values:
	//
	// - state: Task status.
	//
	// - result: Execution result.
	//
	// - resourceChange: Resource change.
	//
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
	// Specifies whether to automatically execute the task.
	//
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// Specifies whether this job performs a compliance pre-check.
	//
	// example:
	//
	// true
	HasConfigProactive *string `json:"hasConfigProactive,omitempty" xml:"hasConfigProactive,omitempty"`
	// Specifies whether to destroy resources.
	//
	// example:
	//
	// fales
	IsDestroy *bool `json:"isDestroy,omitempty" xml:"isDestroy,omitempty"`
	// The template version.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// The resource changes.
	//
	// example:
	//
	// +0 ~0 -0
	ResourcesChanged *string `json:"resourcesChanged,omitempty" xml:"resourcesChanged,omitempty"`
	// The operation command.
	//
	// example:
	//
	// destroy
	SubCommand *string `json:"subCommand,omitempty" xml:"subCommand,omitempty"`
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
