// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProcessDefinitionWithScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetAlertEmailAddress() *string
	SetDescription(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetDescription() *string
	SetExecutionType(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetExecutionType() *string
	SetGlobalParams(v []*UpdateProcessDefinitionWithScheduleRequestGlobalParams) *UpdateProcessDefinitionWithScheduleRequest
	GetGlobalParams() []*UpdateProcessDefinitionWithScheduleRequestGlobalParams
	SetName(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetName() *string
	SetProductNamespace(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetProductNamespace() *string
	SetPublish(v bool) *UpdateProcessDefinitionWithScheduleRequest
	GetPublish() *bool
	SetRegionId(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetRegionId() *string
	SetReleaseState(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetReleaseState() *string
	SetResourceQueue(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetResourceQueue() *string
	SetRetryTimes(v int32) *UpdateProcessDefinitionWithScheduleRequest
	GetRetryTimes() *int32
	SetRunAs(v string) *UpdateProcessDefinitionWithScheduleRequest
	GetRunAs() *string
	SetSchedule(v *UpdateProcessDefinitionWithScheduleRequestSchedule) *UpdateProcessDefinitionWithScheduleRequest
	GetSchedule() *UpdateProcessDefinitionWithScheduleRequestSchedule
	SetTags(v map[string]*string) *UpdateProcessDefinitionWithScheduleRequest
	GetTags() map[string]*string
	SetTaskDefinitionJson(v []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) *UpdateProcessDefinitionWithScheduleRequest
	GetTaskDefinitionJson() []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson
	SetTaskParallelism(v int32) *UpdateProcessDefinitionWithScheduleRequest
	GetTaskParallelism() *int32
	SetTaskRelationJson(v []*UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) *UpdateProcessDefinitionWithScheduleRequest
	GetTaskRelationJson() []*UpdateProcessDefinitionWithScheduleRequestTaskRelationJson
	SetTimeout(v int32) *UpdateProcessDefinitionWithScheduleRequest
	GetTimeout() *int32
}

type UpdateProcessDefinitionWithScheduleRequest struct {
	// The email address for alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The workflow description.
	//
	// example:
	//
	// ods batch workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// PARALLEL
	ExecutionType *string                                                   `json:"executionType,omitempty" xml:"executionType,omitempty"`
	GlobalParams  []*UpdateProcessDefinitionWithScheduleRequestGlobalParams `json:"globalParams,omitempty" xml:"globalParams,omitempty" type:"Repeated"`
	// The workflow name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// Specifies whether to publish the workflow.
	//
	// example:
	//
	// true
	Publish *bool `json:"publish,omitempty" xml:"publish,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The release state of the workflow.
	//
	// example:
	//
	// ONLINE
	ReleaseState *string `json:"releaseState,omitempty" xml:"releaseState,omitempty"`
	// The resource queue.
	//
	// example:
	//
	// root_queue
	ResourceQueue *string `json:"resourceQueue,omitempty" xml:"resourceQueue,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 1
	RetryTimes *int32 `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
	// The user to run the workflow.
	//
	// example:
	//
	// 113***************
	RunAs *string `json:"runAs,omitempty" xml:"runAs,omitempty"`
	// The scheduling configuration.
	Schedule *UpdateProcessDefinitionWithScheduleRequestSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	// The tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// A JSON array of task definitions. This array contains the descriptive information for all tasks in the workflow.
	//
	// This parameter is required.
	TaskDefinitionJson []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty" type:"Repeated"`
	// The degree of concurrent execution for workflow nodes.
	//
	// example:
	//
	// 1
	TaskParallelism *int32 `json:"taskParallelism,omitempty" xml:"taskParallelism,omitempty"`
	// A JSON array that defines the dependencies between tasks in the workflow. \\`preTaskCode\\` specifies the upstream task ID, and \\`postTaskCode\\` specifies the downstream task ID. Each task must have a unique ID. For a task node without an upstream task, add a dependency and set \\`preTaskCode\\` to 0.
	//
	// This parameter is required.
	TaskRelationJson []*UpdateProcessDefinitionWithScheduleRequestTaskRelationJson `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty" type:"Repeated"`
	// The default timeout period for the workflow execution.
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetAlertEmailAddress() *string {
	return s.AlertEmailAddress
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetGlobalParams() []*UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	return s.GlobalParams
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetProductNamespace() *string {
	return s.ProductNamespace
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetPublish() *bool {
	return s.Publish
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetReleaseState() *string {
	return s.ReleaseState
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetResourceQueue() *string {
	return s.ResourceQueue
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetRunAs() *string {
	return s.RunAs
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetSchedule() *UpdateProcessDefinitionWithScheduleRequestSchedule {
	return s.Schedule
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetTags() map[string]*string {
	return s.Tags
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetTaskDefinitionJson() []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	return s.TaskDefinitionJson
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetTaskParallelism() *int32 {
	return s.TaskParallelism
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetTaskRelationJson() []*UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	return s.TaskRelationJson
}

func (s *UpdateProcessDefinitionWithScheduleRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.AlertEmailAddress = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetDescription(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetExecutionType(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.ExecutionType = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetGlobalParams(v []*UpdateProcessDefinitionWithScheduleRequestGlobalParams) *UpdateProcessDefinitionWithScheduleRequest {
	s.GlobalParams = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetName(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetProductNamespace(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.ProductNamespace = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetPublish(v bool) *UpdateProcessDefinitionWithScheduleRequest {
	s.Publish = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetRegionId(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetReleaseState(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.ReleaseState = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetResourceQueue(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.ResourceQueue = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetRetryTimes(v int32) *UpdateProcessDefinitionWithScheduleRequest {
	s.RetryTimes = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetRunAs(v string) *UpdateProcessDefinitionWithScheduleRequest {
	s.RunAs = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetSchedule(v *UpdateProcessDefinitionWithScheduleRequestSchedule) *UpdateProcessDefinitionWithScheduleRequest {
	s.Schedule = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTags(v map[string]*string) *UpdateProcessDefinitionWithScheduleRequest {
	s.Tags = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTaskDefinitionJson(v []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) *UpdateProcessDefinitionWithScheduleRequest {
	s.TaskDefinitionJson = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTaskParallelism(v int32) *UpdateProcessDefinitionWithScheduleRequest {
	s.TaskParallelism = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTaskRelationJson(v []*UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) *UpdateProcessDefinitionWithScheduleRequest {
	s.TaskRelationJson = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) SetTimeout(v int32) *UpdateProcessDefinitionWithScheduleRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequest) Validate() error {
	if s.GlobalParams != nil {
		for _, item := range s.GlobalParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	if s.TaskDefinitionJson != nil {
		for _, item := range s.TaskDefinitionJson {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskRelationJson != nil {
		for _, item := range s.TaskRelationJson {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProcessDefinitionWithScheduleRequestGlobalParams struct {
	Direct *string `json:"direct,omitempty" xml:"direct,omitempty"`
	Prop   *string `json:"prop,omitempty" xml:"prop,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestGlobalParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestGlobalParams) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) GetDirect() *string {
	return s.Direct
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) GetProp() *string {
	return s.Prop
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) GetType() *string {
	return s.Type
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) GetValue() *string {
	return s.Value
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) SetDirect(v string) *UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Direct = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) SetProp(v string) *UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Prop = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) SetType(v string) *UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Type = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) SetValue(v string) *UpdateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Value = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestGlobalParams) Validate() error {
	return dara.Validate(s)
}

type UpdateProcessDefinitionWithScheduleRequestSchedule struct {
	// The cron expression for scheduling.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	Crontab *string `json:"crontab,omitempty" xml:"crontab,omitempty"`
	// The end time of the schedule.
	//
	// example:
	//
	// 2025-12-23 16:13:27
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2024-12-23 16:13:27
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The time zone ID.
	//
	// example:
	//
	// Asia/Shanghai
	TimezoneId *string `json:"timezoneId,omitempty" xml:"timezoneId,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestSchedule) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestSchedule) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) GetCrontab() *string {
	return s.Crontab
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) GetTimezoneId() *string {
	return s.TimezoneId
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) SetCrontab(v string) *UpdateProcessDefinitionWithScheduleRequestSchedule {
	s.Crontab = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) SetEndTime(v string) *UpdateProcessDefinitionWithScheduleRequestSchedule {
	s.EndTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) SetStartTime(v string) *UpdateProcessDefinitionWithScheduleRequestSchedule {
	s.StartTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) SetTimezoneId(v string) *UpdateProcessDefinitionWithScheduleRequestSchedule {
	s.TimezoneId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestSchedule) Validate() error {
	return dara.Validate(s)
}

type UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson struct {
	// The email address for alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The task definition ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23************
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the task definition.
	//
	// example:
	//
	// ods transform task
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable alerts when the task fails.
	//
	// example:
	//
	// true
	FailAlertEnable *bool `json:"failAlertEnable,omitempty" xml:"failAlertEnable,omitempty"`
	// The number of times to retry the task if it fails.
	//
	// example:
	//
	// 1
	FailRetryTimes *int32 `json:"failRetryTimes,omitempty" xml:"failRetryTimes,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_transform_task
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to enable alerts when the task starts.
	//
	// example:
	//
	// true
	StartAlertEnable *bool `json:"startAlertEnable,omitempty" xml:"startAlertEnable,omitempty"`
	// The tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The parameters for the task definition.
	//
	// This parameter is required.
	TaskParams *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams `json:"taskParams,omitempty" xml:"taskParams,omitempty" type:"Struct"`
	// The type of the workflow node.
	//
	// This parameter is required.
	//
	// example:
	//
	// EMR-SERVERLESS-SPARK
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// The default timeout period for the task execution.
	//
	// example:
	//
	// 30
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetAlertEmailAddress() *string {
	return s.AlertEmailAddress
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetCode() *int64 {
	return s.Code
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetDescription() *string {
	return s.Description
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetFailAlertEnable() *bool {
	return s.FailAlertEnable
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetFailRetryTimes() *int32 {
	return s.FailRetryTimes
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetName() *string {
	return s.Name
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetStartAlertEnable() *bool {
	return s.StartAlertEnable
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetTags() map[string]*string {
	return s.Tags
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetTaskParams() *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	return s.TaskParams
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetTaskType() *string {
	return s.TaskType
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.AlertEmailAddress = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetCode(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Code = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetDescription(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetFailAlertEnable(v bool) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.FailAlertEnable = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetFailRetryTimes(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.FailRetryTimes = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetName(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetStartAlertEnable(v bool) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.StartAlertEnable = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTags(v map[string]*string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Tags = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTaskParams(v *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.TaskParams = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTaskType(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.TaskType = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTimeout(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Timeout = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJson) Validate() error {
	if s.TaskParams != nil {
		if err := s.TaskParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams struct {
	// The display version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplaySparkVersion *string `json:"displaySparkVersion,omitempty" xml:"displaySparkVersion,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// ev-h*************
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Specifies whether to enable the Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion      *bool                                                                                `json:"fusion,omitempty" xml:"fusion,omitempty"`
	LocalParams []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams `json:"localParams,omitempty" xml:"localParams,omitempty" type:"Repeated"`
	// The name of the queue on which the task runs.
	//
	// This parameter is required.
	//
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The Spark task configurations.
	SparkConf []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf `json:"sparkConf,omitempty" xml:"sparkConf,omitempty" type:"Repeated"`
	// The number of cores for the Spark driver.
	//
	// example:
	//
	// 1
	SparkDriverCores *int32 `json:"sparkDriverCores,omitempty" xml:"sparkDriverCores,omitempty"`
	// The memory size of the Spark driver.
	//
	// example:
	//
	// 4g
	SparkDriverMemory *int64 `json:"sparkDriverMemory,omitempty" xml:"sparkDriverMemory,omitempty"`
	// The number of cores for each Spark executor.
	//
	// example:
	//
	// 1
	SparkExecutorCores *int32 `json:"sparkExecutorCores,omitempty" xml:"sparkExecutorCores,omitempty"`
	// The memory size of each Spark executor.
	//
	// example:
	//
	// 4g
	SparkExecutorMemory *int64 `json:"sparkExecutorMemory,omitempty" xml:"sparkExecutorMemory,omitempty"`
	// The Spark log level.
	//
	// example:
	//
	// INFO
	SparkLogLevel *string `json:"sparkLogLevel,omitempty" xml:"sparkLogLevel,omitempty"`
	// The path to store Spark task logs.
	SparkLogPath *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	// The Spark engine version.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	SparkVersion *string `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	// The ID of the Data Development task.
	//
	// This parameter is required.
	//
	// example:
	//
	// TSK-d87******************
	TaskBizId *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	// The Spark job type.
	//
	// example:
	//
	// SQL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-d8********
	WorkspaceBizId *string `json:"workspaceBizId,omitempty" xml:"workspaceBizId,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetDisplaySparkVersion() *string {
	return s.DisplaySparkVersion
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetFusion() *bool {
	return s.Fusion
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetLocalParams() []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	return s.LocalParams
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkConf() []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	return s.SparkConf
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkDriverCores() *int32 {
	return s.SparkDriverCores
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkDriverMemory() *int64 {
	return s.SparkDriverMemory
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkExecutorCores() *int32 {
	return s.SparkExecutorCores
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkExecutorMemory() *int64 {
	return s.SparkExecutorMemory
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkLogLevel() *string {
	return s.SparkLogLevel
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkLogPath() *string {
	return s.SparkLogPath
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkVersion() *string {
	return s.SparkVersion
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetTaskBizId() *string {
	return s.TaskBizId
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetType() *string {
	return s.Type
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetWorkspaceBizId() *string {
	return s.WorkspaceBizId
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetDisplaySparkVersion(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.DisplaySparkVersion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetEnvironmentId(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetFusion(v bool) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.Fusion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetLocalParams(v []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.LocalParams = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetResourceQueueId(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.ResourceQueueId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkConf(v []*UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkConf = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkDriverCores(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkDriverCores = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkDriverMemory(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkDriverMemory = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkExecutorCores(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkExecutorCores = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkExecutorMemory(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkExecutorMemory = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkLogLevel(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkLogLevel = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkLogPath(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkLogPath = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkVersion(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkVersion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetTaskBizId(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.TaskBizId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetType(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.Type = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetWorkspaceBizId(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.WorkspaceBizId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) Validate() error {
	if s.LocalParams != nil {
		for _, item := range s.LocalParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SparkConf != nil {
		for _, item := range s.SparkConf {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams struct {
	Direct *string `json:"direct,omitempty" xml:"direct,omitempty"`
	Prop   *string `json:"prop,omitempty" xml:"prop,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GetDirect() *string {
	return s.Direct
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GetProp() *string {
	return s.Prop
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GetType() *string {
	return s.Type
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GetValue() *string {
	return s.Value
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetDirect(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Direct = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetProp(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Prop = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetType(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Type = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetValue(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Value = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) Validate() error {
	return dara.Validate(s)
}

type UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf struct {
	// The key of the Spark configuration.
	//
	// example:
	//
	// spark.dynamicAllocation.enabled
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the Spark configuration.
	//
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) GetKey() *string {
	return s.Key
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) GetValue() *string {
	return s.Value
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) SetKey(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	s.Key = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) SetValue(v string) *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	s.Value = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) Validate() error {
	return dara.Validate(s)
}

type UpdateProcessDefinitionWithScheduleRequestTaskRelationJson struct {
	// The name of the task topology. You can use the workflow name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods batch workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The downstream task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19************
	PostTaskCode *int64 `json:"postTaskCode,omitempty" xml:"postTaskCode,omitempty"`
	// The downstream task version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PostTaskVersion *int32 `json:"postTaskVersion,omitempty" xml:"postTaskVersion,omitempty"`
	// The upstream task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16************
	PreTaskCode *int64 `json:"preTaskCode,omitempty" xml:"preTaskCode,omitempty"`
	// The upstream task version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PreTaskVersion *int32 `json:"preTaskVersion,omitempty" xml:"preTaskVersion,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) GetName() *string {
	return s.Name
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) GetPostTaskCode() *int64 {
	return s.PostTaskCode
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) GetPostTaskVersion() *int32 {
	return s.PostTaskVersion
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) GetPreTaskCode() *int64 {
	return s.PreTaskCode
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) GetPreTaskVersion() *int32 {
	return s.PreTaskVersion
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetName(v string) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPostTaskCode(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PostTaskCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPostTaskVersion(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PostTaskVersion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPreTaskCode(v int64) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PreTaskCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPreTaskVersion(v int32) *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PreTaskVersion = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleRequestTaskRelationJson) Validate() error {
	return dara.Validate(s)
}
