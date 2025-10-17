// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcessDefinitionWithScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEmailAddress(v string) *CreateProcessDefinitionWithScheduleRequest
	GetAlertEmailAddress() *string
	SetDescription(v string) *CreateProcessDefinitionWithScheduleRequest
	GetDescription() *string
	SetExecutionType(v string) *CreateProcessDefinitionWithScheduleRequest
	GetExecutionType() *string
	SetGlobalParams(v []*CreateProcessDefinitionWithScheduleRequestGlobalParams) *CreateProcessDefinitionWithScheduleRequest
	GetGlobalParams() []*CreateProcessDefinitionWithScheduleRequestGlobalParams
	SetName(v string) *CreateProcessDefinitionWithScheduleRequest
	GetName() *string
	SetProductNamespace(v string) *CreateProcessDefinitionWithScheduleRequest
	GetProductNamespace() *string
	SetPublish(v bool) *CreateProcessDefinitionWithScheduleRequest
	GetPublish() *bool
	SetRegionId(v string) *CreateProcessDefinitionWithScheduleRequest
	GetRegionId() *string
	SetResourceQueue(v string) *CreateProcessDefinitionWithScheduleRequest
	GetResourceQueue() *string
	SetRetryTimes(v int32) *CreateProcessDefinitionWithScheduleRequest
	GetRetryTimes() *int32
	SetRunAs(v string) *CreateProcessDefinitionWithScheduleRequest
	GetRunAs() *string
	SetSchedule(v *CreateProcessDefinitionWithScheduleRequestSchedule) *CreateProcessDefinitionWithScheduleRequest
	GetSchedule() *CreateProcessDefinitionWithScheduleRequestSchedule
	SetTags(v map[string]*string) *CreateProcessDefinitionWithScheduleRequest
	GetTags() map[string]*string
	SetTaskDefinitionJson(v []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) *CreateProcessDefinitionWithScheduleRequest
	GetTaskDefinitionJson() []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson
	SetTaskParallelism(v int32) *CreateProcessDefinitionWithScheduleRequest
	GetTaskParallelism() *int32
	SetTaskRelationJson(v []*CreateProcessDefinitionWithScheduleRequestTaskRelationJson) *CreateProcessDefinitionWithScheduleRequest
	GetTaskRelationJson() []*CreateProcessDefinitionWithScheduleRequestTaskRelationJson
	SetTimeout(v int32) *CreateProcessDefinitionWithScheduleRequest
	GetTimeout() *int32
}

type CreateProcessDefinitionWithScheduleRequest struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The description of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods batch workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy
	//
	// This parameter is required.
	//
	// example:
	//
	// PARALLEL
	ExecutionType *string                                                   `json:"executionType,omitempty" xml:"executionType,omitempty"`
	GlobalParams  []*CreateProcessDefinitionWithScheduleRequestGlobalParams `json:"globalParams,omitempty" xml:"globalParams,omitempty" type:"Repeated"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The code of the service.
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
	// The ID of the Alibaba Cloud account used by the user who creates the workflow.
	//
	// example:
	//
	// 113***************
	RunAs *string `json:"runAs,omitempty" xml:"runAs,omitempty"`
	// The scheduling settings.
	Schedule *CreateProcessDefinitionWithScheduleRequestSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	// The tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The descriptions of all nodes in the workflow.
	//
	// This parameter is required.
	TaskDefinitionJson []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty" type:"Repeated"`
	// The node parallelism.
	//
	// example:
	//
	// 1
	TaskParallelism *int32 `json:"taskParallelism,omitempty" xml:"taskParallelism,omitempty"`
	// The dependencies of all nodes in the workflow. preTaskCode specifies the ID of an upstream node, and postTaskCode specifies the ID of a downstream node. The ID of each node is unique. If a node does not have an upstream node, set preTaskCode to 0.
	//
	// This parameter is required.
	TaskRelationJson []*CreateProcessDefinitionWithScheduleRequestTaskRelationJson `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty" type:"Repeated"`
	// The default timeout period of the workflow.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetAlertEmailAddress() *string {
	return s.AlertEmailAddress
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetGlobalParams() []*CreateProcessDefinitionWithScheduleRequestGlobalParams {
	return s.GlobalParams
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetName() *string {
	return s.Name
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetProductNamespace() *string {
	return s.ProductNamespace
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetPublish() *bool {
	return s.Publish
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetResourceQueue() *string {
	return s.ResourceQueue
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetRunAs() *string {
	return s.RunAs
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetSchedule() *CreateProcessDefinitionWithScheduleRequestSchedule {
	return s.Schedule
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetTags() map[string]*string {
	return s.Tags
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetTaskDefinitionJson() []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	return s.TaskDefinitionJson
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetTaskParallelism() *int32 {
	return s.TaskParallelism
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetTaskRelationJson() []*CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	return s.TaskRelationJson
}

func (s *CreateProcessDefinitionWithScheduleRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetAlertEmailAddress(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.AlertEmailAddress = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetDescription(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.Description = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetExecutionType(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.ExecutionType = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetGlobalParams(v []*CreateProcessDefinitionWithScheduleRequestGlobalParams) *CreateProcessDefinitionWithScheduleRequest {
	s.GlobalParams = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetName(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetProductNamespace(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.ProductNamespace = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetPublish(v bool) *CreateProcessDefinitionWithScheduleRequest {
	s.Publish = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetRegionId(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetResourceQueue(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.ResourceQueue = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetRetryTimes(v int32) *CreateProcessDefinitionWithScheduleRequest {
	s.RetryTimes = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetRunAs(v string) *CreateProcessDefinitionWithScheduleRequest {
	s.RunAs = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetSchedule(v *CreateProcessDefinitionWithScheduleRequestSchedule) *CreateProcessDefinitionWithScheduleRequest {
	s.Schedule = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTags(v map[string]*string) *CreateProcessDefinitionWithScheduleRequest {
	s.Tags = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTaskDefinitionJson(v []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) *CreateProcessDefinitionWithScheduleRequest {
	s.TaskDefinitionJson = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTaskParallelism(v int32) *CreateProcessDefinitionWithScheduleRequest {
	s.TaskParallelism = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTaskRelationJson(v []*CreateProcessDefinitionWithScheduleRequestTaskRelationJson) *CreateProcessDefinitionWithScheduleRequest {
	s.TaskRelationJson = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) SetTimeout(v int32) *CreateProcessDefinitionWithScheduleRequest {
	s.Timeout = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequest) Validate() error {
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

type CreateProcessDefinitionWithScheduleRequestGlobalParams struct {
	Direct *string `json:"direct,omitempty" xml:"direct,omitempty"`
	Prop   *string `json:"prop,omitempty" xml:"prop,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestGlobalParams) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestGlobalParams) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) GetDirect() *string {
	return s.Direct
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) GetProp() *string {
	return s.Prop
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) GetType() *string {
	return s.Type
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) GetValue() *string {
	return s.Value
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) SetDirect(v string) *CreateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Direct = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) SetProp(v string) *CreateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Prop = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) SetType(v string) *CreateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) SetValue(v string) *CreateProcessDefinitionWithScheduleRequestGlobalParams {
	s.Value = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestGlobalParams) Validate() error {
	return dara.Validate(s)
}

type CreateProcessDefinitionWithScheduleRequestSchedule struct {
	// The CRON expression that is used for scheduling.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	Crontab *string `json:"crontab,omitempty" xml:"crontab,omitempty"`
	// The end time of the scheduling.
	//
	// example:
	//
	// 2025-12-23 16:13:27
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time of the scheduling.
	//
	// example:
	//
	// 2024-12-23 16:13:27
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimezoneId *string `json:"timezoneId,omitempty" xml:"timezoneId,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestSchedule) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestSchedule) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) GetCrontab() *string {
	return s.Crontab
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) GetTimezoneId() *string {
	return s.TimezoneId
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) SetCrontab(v string) *CreateProcessDefinitionWithScheduleRequestSchedule {
	s.Crontab = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) SetEndTime(v string) *CreateProcessDefinitionWithScheduleRequestSchedule {
	s.EndTime = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) SetStartTime(v string) *CreateProcessDefinitionWithScheduleRequestSchedule {
	s.StartTime = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) SetTimezoneId(v string) *CreateProcessDefinitionWithScheduleRequestSchedule {
	s.TimezoneId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestSchedule) Validate() error {
	return dara.Validate(s)
}

type CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36************
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The node description.
	//
	// example:
	//
	// ods transform task
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to send alerts when the node fails.
	//
	// example:
	//
	// false
	FailAlertEnable *bool `json:"failAlertEnable,omitempty" xml:"failAlertEnable,omitempty"`
	// The number of retries when the node fails.
	//
	// example:
	//
	// 1
	FailRetryTimes *int32 `json:"failRetryTimes,omitempty" xml:"failRetryTimes,omitempty"`
	// The name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_transform_task
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to send alerts when the node is started.
	//
	// example:
	//
	// false
	StartAlertEnable *bool `json:"startAlertEnable,omitempty" xml:"startAlertEnable,omitempty"`
	// The tags.
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The job parameters.
	//
	// This parameter is required.
	TaskParams *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams `json:"taskParams,omitempty" xml:"taskParams,omitempty" type:"Struct"`
	// The type of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// MigrateData
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// The timeout period of the callback. Unit: seconds.
	//
	// example:
	//
	// 1200
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetAlertEmailAddress() *string {
	return s.AlertEmailAddress
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetCode() *int64 {
	return s.Code
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetDescription() *string {
	return s.Description
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetFailAlertEnable() *bool {
	return s.FailAlertEnable
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetFailRetryTimes() *int32 {
	return s.FailRetryTimes
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetName() *string {
	return s.Name
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetStartAlertEnable() *bool {
	return s.StartAlertEnable
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetTags() map[string]*string {
	return s.Tags
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetTaskParams() *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	return s.TaskParams
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetAlertEmailAddress(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.AlertEmailAddress = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetCode(v int64) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Code = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetDescription(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Description = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetFailAlertEnable(v bool) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.FailAlertEnable = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetFailRetryTimes(v int32) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.FailRetryTimes = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetName(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetStartAlertEnable(v bool) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.StartAlertEnable = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTags(v map[string]*string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Tags = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTaskParams(v *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.TaskParams = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTaskType(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.TaskType = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) SetTimeout(v int32) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson {
	s.Timeout = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJson) Validate() error {
	if s.TaskParams != nil {
		if err := s.TaskParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams struct {
	// The displayed version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplaySparkVersion *string `json:"displaySparkVersion,omitempty" xml:"displaySparkVersion,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-crhq2h5lhtgju93buhkg
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Specifies whether to enable Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion      *bool                                                                                `json:"fusion,omitempty" xml:"fusion,omitempty"`
	LocalParams []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams `json:"localParams,omitempty" xml:"localParams,omitempty" type:"Repeated"`
	// The name of the resource queue on which the job runs.
	//
	// This parameter is required.
	//
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The configurations of the Spark job.
	SparkConf []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf `json:"sparkConf,omitempty" xml:"sparkConf,omitempty" type:"Repeated"`
	// The number of driver cores of the Spark job.
	//
	// example:
	//
	// 1
	SparkDriverCores *int32 `json:"sparkDriverCores,omitempty" xml:"sparkDriverCores,omitempty"`
	// The size of driver memory of the Spark job.
	//
	// example:
	//
	// 4g
	SparkDriverMemory *int64 `json:"sparkDriverMemory,omitempty" xml:"sparkDriverMemory,omitempty"`
	// The number of executor cores of the Spark job.
	//
	// example:
	//
	// 1
	SparkExecutorCores *int32 `json:"sparkExecutorCores,omitempty" xml:"sparkExecutorCores,omitempty"`
	// The size of executor memory of the Spark job.
	//
	// example:
	//
	// 4g
	SparkExecutorMemory *int64 `json:"sparkExecutorMemory,omitempty" xml:"sparkExecutorMemory,omitempty"`
	// The level of the Spark log.
	//
	// example:
	//
	// INFO
	SparkLogLevel *string `json:"sparkLogLevel,omitempty" xml:"sparkLogLevel,omitempty"`
	// The path where the operational logs of the Spark job are stored.
	SparkLogPath *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	SparkVersion *string `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	// The ID of the data development job.
	//
	// This parameter is required.
	//
	// example:
	//
	// TSK-d87******************
	TaskBizId *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	// The type of the Spark job.
	//
	// example:
	//
	// VPC
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

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetDisplaySparkVersion() *string {
	return s.DisplaySparkVersion
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetFusion() *bool {
	return s.Fusion
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetLocalParams() []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	return s.LocalParams
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkConf() []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	return s.SparkConf
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkDriverCores() *int32 {
	return s.SparkDriverCores
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkDriverMemory() *int64 {
	return s.SparkDriverMemory
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkExecutorCores() *int32 {
	return s.SparkExecutorCores
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkExecutorMemory() *int64 {
	return s.SparkExecutorMemory
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkLogLevel() *string {
	return s.SparkLogLevel
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkLogPath() *string {
	return s.SparkLogPath
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetSparkVersion() *string {
	return s.SparkVersion
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetTaskBizId() *string {
	return s.TaskBizId
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetType() *string {
	return s.Type
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) GetWorkspaceBizId() *string {
	return s.WorkspaceBizId
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetDisplaySparkVersion(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.DisplaySparkVersion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetEnvironmentId(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.EnvironmentId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetFusion(v bool) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.Fusion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetLocalParams(v []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.LocalParams = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetResourceQueueId(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.ResourceQueueId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkConf(v []*CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkConf = v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkDriverCores(v int32) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkDriverCores = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkDriverMemory(v int64) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkDriverMemory = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkExecutorCores(v int32) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkExecutorCores = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkExecutorMemory(v int64) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkExecutorMemory = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkLogLevel(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkLogLevel = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkLogPath(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkLogPath = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetSparkVersion(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.SparkVersion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetTaskBizId(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.TaskBizId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetType(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) SetWorkspaceBizId(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams {
	s.WorkspaceBizId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParams) Validate() error {
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

type CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams struct {
	Direct *string `json:"direct,omitempty" xml:"direct,omitempty"`
	Prop   *string `json:"prop,omitempty" xml:"prop,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GetDirect() *string {
	return s.Direct
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GetProp() *string {
	return s.Prop
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GetType() *string {
	return s.Type
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) GetValue() *string {
	return s.Value
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetDirect(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Direct = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetProp(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Prop = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetType(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) SetValue(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams {
	s.Value = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsLocalParams) Validate() error {
	return dara.Validate(s)
}

type CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf struct {
	// The key of the SparkConf object.
	//
	// example:
	//
	// spark.dynamicAllocation.enabled
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the SparkConf object.
	//
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) GetKey() *string {
	return s.Key
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) GetValue() *string {
	return s.Value
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) SetKey(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	s.Key = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) SetValue(v string) *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf {
	s.Value = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskDefinitionJsonTaskParamsSparkConf) Validate() error {
	return dara.Validate(s)
}

type CreateProcessDefinitionWithScheduleRequestTaskRelationJson struct {
	// The name of the node topology. You can enter a workflow name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods batch workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the downstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 28************
	PostTaskCode *int64 `json:"postTaskCode,omitempty" xml:"postTaskCode,omitempty"`
	// The version of the downstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PostTaskVersion *int32 `json:"postTaskVersion,omitempty" xml:"postTaskVersion,omitempty"`
	// The ID of the upstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16************
	PreTaskCode *int64 `json:"preTaskCode,omitempty" xml:"preTaskCode,omitempty"`
	// The version of the upstream node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PreTaskVersion *int32 `json:"preTaskVersion,omitempty" xml:"preTaskVersion,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleRequestTaskRelationJson) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleRequestTaskRelationJson) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) GetName() *string {
	return s.Name
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) GetPostTaskCode() *int64 {
	return s.PostTaskCode
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) GetPostTaskVersion() *int32 {
	return s.PostTaskVersion
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) GetPreTaskCode() *int64 {
	return s.PreTaskCode
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) GetPreTaskVersion() *int32 {
	return s.PreTaskVersion
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetName(v string) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPostTaskCode(v int64) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PostTaskCode = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPostTaskVersion(v int32) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PostTaskVersion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPreTaskCode(v int64) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PreTaskCode = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) SetPreTaskVersion(v int32) *CreateProcessDefinitionWithScheduleRequestTaskRelationJson {
	s.PreTaskVersion = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleRequestTaskRelationJson) Validate() error {
	return dara.Validate(s)
}
