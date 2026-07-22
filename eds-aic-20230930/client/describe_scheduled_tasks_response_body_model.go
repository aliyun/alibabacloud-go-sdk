// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeScheduledTasksResponseBody
	GetCode() *string
	SetMaxResults(v int32) *DescribeScheduledTasksResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeScheduledTasksResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeScheduledTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeScheduledTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribeScheduledTasksResponseBodyTasks) *DescribeScheduledTasksResponseBody
	GetTasks() []*DescribeScheduledTasksResponseBodyTasks
	SetTotalCount(v int32) *DescribeScheduledTasksResponseBody
	GetTotalCount() *int32
}

type DescribeScheduledTasksResponseBody struct {
	// The status code of the operation.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token that indicates the position from which to start reading. Leave this parameter empty to read from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A51B1DF-96FF-3BCC-B08C-783161D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of scheduled tasks.
	Tasks []*DescribeScheduledTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScheduledTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeScheduledTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeScheduledTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeScheduledTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeScheduledTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScheduledTasksResponseBody) GetTasks() []*DescribeScheduledTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeScheduledTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeScheduledTasksResponseBody) SetCode(v string) *DescribeScheduledTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetMaxResults(v int32) *DescribeScheduledTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetMessage(v string) *DescribeScheduledTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetNextToken(v string) *DescribeScheduledTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetRequestId(v string) *DescribeScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetTasks(v []*DescribeScheduledTasksResponseBodyTasks) *DescribeScheduledTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetTotalCount(v int32) *DescribeScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScheduledTasksResponseBodyTasks struct {
	// The cron expression.
	//
	// example:
	//
	// 0 0 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-01T00:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-06-12T10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The list of associated instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The last execution time.
	//
	// example:
	//
	// 2026-06-12T00:00:00
	LastExecutionAt *string `json:"LastExecutionAt,omitempty" xml:"LastExecutionAt,omitempty"`
	// The next execution time.
	//
	// example:
	//
	// 2026-06-13T00:00:00
	NextExecutionAt *string `json:"NextExecutionAt,omitempty" xml:"NextExecutionAt,omitempty"`
	// The run configuration.
	RunConfig *DescribeScheduledTasksResponseBodyTasksRunConfig `json:"RunConfig,omitempty" xml:"RunConfig,omitempty" type:"Struct"`
	// The scheduled task ID.
	//
	// example:
	//
	// sch-260705-agb*****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// The status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task configuration ID.
	//
	// example:
	//
	// tsk-260615-*****
	TaskConfigId *string `json:"TaskConfigId,omitempty" xml:"TaskConfigId,omitempty"`
	// The task name.
	//
	// example:
	//
	// Daily data synchronization task.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The total number of executions.
	//
	// example:
	//
	// 100
	TotalExecutions *int64 `json:"TotalExecutions,omitempty" xml:"TotalExecutions,omitempty"`
	// The total number of failures.
	//
	// example:
	//
	// 2
	TotalFailures *int64 `json:"TotalFailures,omitempty" xml:"TotalFailures,omitempty"`
	// The user prompt or task description.
	//
	// example:
	//
	// Execute daily data synchronization task.
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
	// The CAS version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeScheduledTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetLastExecutionAt() *string {
	return s.LastExecutionAt
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetNextExecutionAt() *string {
	return s.NextExecutionAt
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetRunConfig() *DescribeScheduledTasksResponseBodyTasksRunConfig {
	return s.RunConfig
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetTaskConfigId() *string {
	return s.TaskConfigId
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetTotalExecutions() *int64 {
	return s.TotalExecutions
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetTotalFailures() *int64 {
	return s.TotalFailures
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *DescribeScheduledTasksResponseBodyTasks) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetCronExpression(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.CronExpression = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetGmtCreate(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.GmtCreate = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetGmtModified(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.GmtModified = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetInstanceIds(v []*string) *DescribeScheduledTasksResponseBodyTasks {
	s.InstanceIds = v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetLastExecutionAt(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.LastExecutionAt = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetNextExecutionAt(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.NextExecutionAt = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetRunConfig(v *DescribeScheduledTasksResponseBodyTasksRunConfig) *DescribeScheduledTasksResponseBodyTasks {
	s.RunConfig = v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetScheduledId(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.ScheduledId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetStatus(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetTaskConfigId(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.TaskConfigId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetTaskName(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetTotalExecutions(v int64) *DescribeScheduledTasksResponseBodyTasks {
	s.TotalExecutions = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetTotalFailures(v int64) *DescribeScheduledTasksResponseBodyTasks {
	s.TotalFailures = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetUserPrompt(v string) *DescribeScheduledTasksResponseBodyTasks {
	s.UserPrompt = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) SetVersion(v int32) *DescribeScheduledTasksResponseBodyTasks {
	s.Version = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasks) Validate() error {
	if s.RunConfig != nil {
		if err := s.RunConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScheduledTasksResponseBodyTasksRunConfig struct {
	// The extra parameters.
	//
	// example:
	//
	// {"batchSize":"1000"}
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	// The maximum number of steps.
	//
	// example:
	//
	// 10
	MaxSteps *int32 `json:"MaxSteps,omitempty" xml:"MaxSteps,omitempty"`
	// The timeout period, in seconds.
	//
	// example:
	//
	// 3600
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s DescribeScheduledTasksResponseBodyTasksRunConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyTasksRunConfig) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyTasksRunConfig) GetExtraParams() *string {
	return s.ExtraParams
}

func (s *DescribeScheduledTasksResponseBodyTasksRunConfig) GetMaxSteps() *int32 {
	return s.MaxSteps
}

func (s *DescribeScheduledTasksResponseBodyTasksRunConfig) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *DescribeScheduledTasksResponseBodyTasksRunConfig) SetExtraParams(v string) *DescribeScheduledTasksResponseBodyTasksRunConfig {
	s.ExtraParams = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasksRunConfig) SetMaxSteps(v int32) *DescribeScheduledTasksResponseBodyTasksRunConfig {
	s.MaxSteps = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasksRunConfig) SetTimeoutSeconds(v int32) *DescribeScheduledTasksResponseBodyTasksRunConfig {
	s.TimeoutSeconds = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyTasksRunConfig) Validate() error {
	return dara.Validate(s)
}
