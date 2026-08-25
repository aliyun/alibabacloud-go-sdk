// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody
	GetJobs() []*ListJobsResponseBodyJobs
	SetPageNumber(v int32) *ListJobsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListJobsResponseBody
	GetTotalCount() *int32
}

type ListJobsResponseBody struct {
	// The list of job information.
	Jobs []*ListJobsResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results returned per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 882304F9-6DB1-5593-A719-33473D082B9C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 11
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetJobs() []*ListJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobsResponseBodyJobs struct {
	// The job configuration.
	Config *ListJobsResponseBodyJobsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The time when the job was created, in UTC in the ISO 8601 format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-07-05T02:13:43Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The job description.
	//
	// example:
	//
	// plan
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 5m
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
	// Indicates whether the assertion check passed.
	//
	// example:
	//
	// true
	IsPassAssertCheck *bool `json:"isPassAssertCheck,omitempty" xml:"isPassAssertCheck,omitempty"`
	// The job ID.
	//
	// example:
	//
	// job-433aff9e4dca57b147c
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The job status. Valid values:
	//
	// - Pending: The initial status after the job is created.
	//
	// - PlanQueued: After the job is created, if no workflow is available, the job is queued.
	//
	// - Planning: The resource job is in the Plan execution phase.
	//
	// - ConfigProactiveInProgress: Compliance pre-check is in progress. The account must have the compliance pre-check feature enabled.
	//
	// - ConfigProactiveSuccess: Compliance pre-check succeeded. The account must have the compliance pre-check feature enabled.
	//
	// - Planned: The resource job has completed Plan execution.
	//
	// - PlannedAndFinished: After Plan execution is completed, no diff is found. This is a final status.
	//
	// - Confirmed: The resource job is waiting for confirmation after Plan execution is completed.
	//
	// - ApplyQueued: During job execution, if no workflow is available, the job is queued.
	//
	// - Applying: The resource job is in the Apply execution phase.
	//
	// - Applied: The resource job has completed Apply execution. This is a final status.
	//
	// - Errored: The job execution encountered an error. This is a final status.
	//
	// - Canceled: The job execution was canceled. This is a final status.
	//
	// - Discarded: The plan of the resource job was discarded. This is a final status.
	//
	// - ConfigProactiveFailure: Compliance pre-check failed. The account must have the compliance pre-check feature enabled.
	//
	// example:
	//
	// Errored
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status details.
	StatusDetail map[string]*JobsStatusDetailValue `json:"statusDetail,omitempty" xml:"statusDetail,omitempty"`
	// The task ID.
	//
	// example:
	//
	// task-518876866c2c3efb
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The Terraform provider version.
	//
	// example:
	//
	// 1.240.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) GetConfig() *ListJobsResponseBodyJobsConfig {
	return s.Config
}

func (s *ListJobsResponseBodyJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobsResponseBodyJobs) GetDescription() *string {
	return s.Description
}

func (s *ListJobsResponseBodyJobs) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *ListJobsResponseBodyJobs) GetExecuteType() *string {
	return s.ExecuteType
}

func (s *ListJobsResponseBodyJobs) GetIsPassAssertCheck() *bool {
	return s.IsPassAssertCheck
}

func (s *ListJobsResponseBodyJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListJobsResponseBodyJobs) GetStatusDetail() map[string]*JobsStatusDetailValue {
	return s.StatusDetail
}

func (s *ListJobsResponseBodyJobs) GetTaskId() *string {
	return s.TaskId
}

func (s *ListJobsResponseBodyJobs) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *ListJobsResponseBodyJobs) SetConfig(v *ListJobsResponseBodyJobsConfig) *ListJobsResponseBodyJobs {
	s.Config = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetCreateTime(v string) *ListJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetDescription(v string) *ListJobsResponseBodyJobs {
	s.Description = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetElapsedTime(v int64) *ListJobsResponseBodyJobs {
	s.ElapsedTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetExecuteType(v string) *ListJobsResponseBodyJobs {
	s.ExecuteType = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetIsPassAssertCheck(v bool) *ListJobsResponseBodyJobs {
	s.IsPassAssertCheck = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobId(v string) *ListJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatus(v string) *ListJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatusDetail(v map[string]*JobsStatusDetailValue) *ListJobsResponseBodyJobs {
	s.StatusDetail = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetTaskId(v string) *ListJobsResponseBodyJobs {
	s.TaskId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetTerraformProviderVersion(v string) *ListJobsResponseBodyJobs {
	s.TerraformProviderVersion = &v
	return s
}

func (s *ListJobsResponseBodyJobs) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyJobsConfig struct {
	// Indicates whether the job is a destroy job.
	//
	// example:
	//
	// true
	IsDestroy *bool `json:"isDestroy,omitempty" xml:"isDestroy,omitempty"`
	// The template description.
	//
	// example:
	//
	// moduleDescription
	ModuleDescription *string `json:"moduleDescription,omitempty" xml:"moduleDescription,omitempty"`
	// The template version.
	//
	// example:
	//
	// v4
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// The resource change content.
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

func (s ListJobsResponseBodyJobsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsConfig) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsConfig) GetIsDestroy() *bool {
	return s.IsDestroy
}

func (s *ListJobsResponseBodyJobsConfig) GetModuleDescription() *string {
	return s.ModuleDescription
}

func (s *ListJobsResponseBodyJobsConfig) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *ListJobsResponseBodyJobsConfig) GetResourcesChanged() *string {
	return s.ResourcesChanged
}

func (s *ListJobsResponseBodyJobsConfig) GetSubCommand() *string {
	return s.SubCommand
}

func (s *ListJobsResponseBodyJobsConfig) SetIsDestroy(v bool) *ListJobsResponseBodyJobsConfig {
	s.IsDestroy = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetModuleDescription(v string) *ListJobsResponseBodyJobsConfig {
	s.ModuleDescription = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetModuleVersion(v string) *ListJobsResponseBodyJobsConfig {
	s.ModuleVersion = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetResourcesChanged(v string) *ListJobsResponseBodyJobsConfig {
	s.ResourcesChanged = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetSubCommand(v string) *ListJobsResponseBodyJobsConfig {
	s.SubCommand = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) Validate() error {
	return dara.Validate(s)
}
