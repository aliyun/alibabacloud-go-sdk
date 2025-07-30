// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStartEnabled(v bool) *CreateWorkflowInstancesRequest
	GetAutoStartEnabled() *bool
	SetComment(v string) *CreateWorkflowInstancesRequest
	GetComment() *string
	SetDefaultRunProperties(v *CreateWorkflowInstancesRequestDefaultRunProperties) *CreateWorkflowInstancesRequest
	GetDefaultRunProperties() *CreateWorkflowInstancesRequestDefaultRunProperties
	SetEnvType(v string) *CreateWorkflowInstancesRequest
	GetEnvType() *string
	SetName(v string) *CreateWorkflowInstancesRequest
	GetName() *string
	SetPeriods(v *CreateWorkflowInstancesRequestPeriods) *CreateWorkflowInstancesRequest
	GetPeriods() *CreateWorkflowInstancesRequestPeriods
	SetProjectId(v int64) *CreateWorkflowInstancesRequest
	GetProjectId() *int64
	SetTagCreationPolicy(v string) *CreateWorkflowInstancesRequest
	GetTagCreationPolicy() *string
	SetTags(v []*CreateWorkflowInstancesRequestTags) *CreateWorkflowInstancesRequest
	GetTags() []*CreateWorkflowInstancesRequestTags
	SetTaskParameters(v string) *CreateWorkflowInstancesRequest
	GetTaskParameters() *string
	SetType(v string) *CreateWorkflowInstancesRequest
	GetType() *string
	SetWorkflowId(v int64) *CreateWorkflowInstancesRequest
	GetWorkflowId() *int64
	SetWorkflowParameters(v string) *CreateWorkflowInstancesRequest
	GetWorkflowParameters() *string
}

type CreateWorkflowInstancesRequest struct {
	// The default value is true.
	//
	// example:
	//
	// true
	AutoStartEnabled *bool `json:"AutoStartEnabled,omitempty" xml:"AutoStartEnabled,omitempty"`
	// The reason for the creation.
	//
	// example:
	//
	// create for test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The runtime configuration.
	DefaultRunProperties *CreateWorkflowInstancesRequestDefaultRunProperties `json:"DefaultRunProperties,omitempty" xml:"DefaultRunProperties,omitempty" type:"Struct"`
	// The environment of the workspace. Valid values:
	//
	// 	- Prod
	//
	// 	- Dev
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// WorkflowInstance1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration of the data backfilling period.
	Periods *CreateWorkflowInstancesRequestPeriods `json:"Periods,omitempty" xml:"Periods,omitempty" type:"Struct"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId         *int64                                `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TagCreationPolicy *string                               `json:"TagCreationPolicy,omitempty" xml:"TagCreationPolicy,omitempty"`
	Tags              []*CreateWorkflowInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The task-specific parameters. The value is in the JSON format. The key specifies the task ID. You can call the GetTask operation to obtain the format of the value by querying the script parameters.
	//
	// example:
	//
	// {
	//
	//   "1001": "key1=val2 key2=val2",
	//
	//   "1002": "key1=val2 key2=val2"
	//
	// }
	TaskParameters *string `json:"TaskParameters,omitempty" xml:"TaskParameters,omitempty"`
	// The type of the workflow instance. Valid values:
	//
	// 	- SupplementData: The values of the RootTaskIds and IncludeTaskIds parameters vary based on the value of the Mode parameter. For more information, see the Mode parameter in this API operation.
	//
	// 	- ManualWorkflow: If you set the Type parameter to ManualWorkflow, you must set the WorkflowId parameter to the ID of the manually triggered workflow. The RootTaskIds parameter is optional. If you do not configure the RootTaskIds parameter, the IDs of the default root nodes of the manually triggered workflow are used.
	//
	// 	- Manual: You need to configure only the RootTaskIds parameter. The RootTaskIds parameter specifies the IDs of the manually triggered tasks that need to be run.
	//
	// 	- SmokeTest: You need to configure only the RootTaskIds parameter. The RootTaskIds parameter specifies the IDs of the test tasks that need to be run.
	//
	// This parameter is required.
	//
	// example:
	//
	// SupplementData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workflow to which the instance belongs. This parameter is set to 1 for auto triggered tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow parameters. This parameter takes effect only when you set the `WorkflowId` parameter to a value other than 1. If your workflow is an auto triggered workflow, configure this parameter in the key=value format. The parameters that you configure in this parameter have a lower priority than task parameters. If your workflow is a manually triggered workflow, configure this parameter in the JSON format. The parameters that you configure in this parameter have a higher priority than task parameters.
	//
	// example:
	//
	// {
	//
	//   "key1": "value1",
	//
	//   "key2": "value2"
	//
	// }
	WorkflowParameters *string `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty"`
}

func (s CreateWorkflowInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesRequest) GetAutoStartEnabled() *bool {
	return s.AutoStartEnabled
}

func (s *CreateWorkflowInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateWorkflowInstancesRequest) GetDefaultRunProperties() *CreateWorkflowInstancesRequestDefaultRunProperties {
	return s.DefaultRunProperties
}

func (s *CreateWorkflowInstancesRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateWorkflowInstancesRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowInstancesRequest) GetPeriods() *CreateWorkflowInstancesRequestPeriods {
	return s.Periods
}

func (s *CreateWorkflowInstancesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateWorkflowInstancesRequest) GetTagCreationPolicy() *string {
	return s.TagCreationPolicy
}

func (s *CreateWorkflowInstancesRequest) GetTags() []*CreateWorkflowInstancesRequestTags {
	return s.Tags
}

func (s *CreateWorkflowInstancesRequest) GetTaskParameters() *string {
	return s.TaskParameters
}

func (s *CreateWorkflowInstancesRequest) GetType() *string {
	return s.Type
}

func (s *CreateWorkflowInstancesRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *CreateWorkflowInstancesRequest) GetWorkflowParameters() *string {
	return s.WorkflowParameters
}

func (s *CreateWorkflowInstancesRequest) SetAutoStartEnabled(v bool) *CreateWorkflowInstancesRequest {
	s.AutoStartEnabled = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetComment(v string) *CreateWorkflowInstancesRequest {
	s.Comment = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetDefaultRunProperties(v *CreateWorkflowInstancesRequestDefaultRunProperties) *CreateWorkflowInstancesRequest {
	s.DefaultRunProperties = v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetEnvType(v string) *CreateWorkflowInstancesRequest {
	s.EnvType = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetName(v string) *CreateWorkflowInstancesRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetPeriods(v *CreateWorkflowInstancesRequestPeriods) *CreateWorkflowInstancesRequest {
	s.Periods = v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetProjectId(v int64) *CreateWorkflowInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetTagCreationPolicy(v string) *CreateWorkflowInstancesRequest {
	s.TagCreationPolicy = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetTags(v []*CreateWorkflowInstancesRequestTags) *CreateWorkflowInstancesRequest {
	s.Tags = v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetTaskParameters(v string) *CreateWorkflowInstancesRequest {
	s.TaskParameters = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetType(v string) *CreateWorkflowInstancesRequest {
	s.Type = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetWorkflowId(v int64) *CreateWorkflowInstancesRequest {
	s.WorkflowId = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) SetWorkflowParameters(v string) *CreateWorkflowInstancesRequest {
	s.WorkflowParameters = &v
	return s
}

func (s *CreateWorkflowInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstancesRequestDefaultRunProperties struct {
	// The alert settings.
	Alert *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert `json:"Alert,omitempty" xml:"Alert,omitempty" type:"Struct"`
	// The configurations for analysis. If you set the Type parameter to SupplementData, this parameter is required.
	Analysis *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The IDs of the projects that do not need to be run.
	ExcludeProjectIds []*int64 `json:"ExcludeProjectIds,omitempty" xml:"ExcludeProjectIds,omitempty" type:"Repeated"`
	// The IDs of the tasks that do not need to be run.
	ExcludeTaskIds []*int64 `json:"ExcludeTaskIds,omitempty" xml:"ExcludeTaskIds,omitempty" type:"Repeated"`
	// The IDs of the projects that need to be run.
	IncludeProjectIds []*int64 `json:"IncludeProjectIds,omitempty" xml:"IncludeProjectIds,omitempty" type:"Repeated"`
	// The IDs of the tasks that need to be run.
	IncludeTaskIds []*int64 `json:"IncludeTaskIds,omitempty" xml:"IncludeTaskIds,omitempty" type:"Repeated"`
	// The data backfill mode. Default value: ManualSelection. If you set the Type parameter to SupplementData, this parameter is required. Valid values:
	//
	// 	- General: You can specify only `one root task ID`. The `IncludeTaskIds` parameter is optional. If you do not configure the IncludeTaskIds parameter, the tasks that are specified by the `RootTaskIds` parameter are included by default.
	//
	// 	- ManualSelection: You can specify `multiple root task IDs`. The `IncludeTaskIds` parameter is optional. If you do not configure the IncludeTaskIds parameter, the tasks that are specified by the `RootTaskIds` parameter are included by default.
	//
	// 	- Chain: If you set the Mode parameter to Chain, you must leave the `RootTaskIds` parameter empty and set the `IncludeTaskIds` parameter to the start task ID and the end task ID.
	//
	// 	- AllDownstream: You can specify only one `root task ID`.
	//
	// example:
	//
	// ManualSelection
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The running order. Valid values:
	//
	// 	- Asc (default): The tasks are sorted by data timestamp in ascending order.
	//
	// 	- Desc: The tasks are sorted by data timestamp in descending order.
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of tasks that can be run in parallel. If you specify a value that ranges from 2 to 10, the value indicates the number of tasks that can be run in parallel. If you set the value to 1, the tasks are run one by one. If you set the Type parameter to SupplementData, this parameter is required.
	//
	// example:
	//
	// 2
	Parallelism *int32 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The running priority. Valid values: 1 to 11. A larger value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The weight policy for the priority. Valid values:
	//
	// 	- `Disable`: No weight policy for the priority is enabled. This is the default value.
	//
	// 	- `Upstream`: The sum of the weights for the ancestor nodes of the current node is calculated. The more ancestor nodes, the higher the weight.
	//
	// example:
	//
	// Upstream
	PriorityWeightStrategy *string `json:"PriorityWeightStrategy,omitempty" xml:"PriorityWeightStrategy,omitempty"`
	// The root task IDs.
	//
	// 	- If you set the Type parameter to SupplementData and the Mode parameter to a value other than Chain, the RootTaskIds parameter is required.
	//
	// 	- If you set the Type parameter to ManualWorkflow, the RootTaskIds parameter is optional. If you do not configure the RootTaskIds parameter, the IDs of the default root nodes of the manually triggered workflow are used.
	//
	// 	- If you set the Type parameter to Manual, the RootTaskIds parameter is required. The RootTaskIds parameter specifies the IDs of the manually triggered tasks that need to be run.
	//
	// 	- If you set the Type parameter to SmokeTest, the RootTaskIds parameter is required. The RootTaskIds parameter specifies the IDs of the test tasks that need to be run.
	RootTaskIds []*int64 `json:"RootTaskIds,omitempty" xml:"RootTaskIds,omitempty" type:"Repeated"`
	// The policy for running. If you leave this parameter empty, the task configuration is used.
	RunPolicy *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy `json:"RunPolicy,omitempty" xml:"RunPolicy,omitempty" type:"Struct"`
	// The identifier of the custom resource group for scheduling. If you leave this parameter empty, the runtime configuration is used.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	RuntimeResource *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
}

func (s CreateWorkflowInstancesRequestDefaultRunProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesRequestDefaultRunProperties) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetAlert() *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert {
	return s.Alert
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetAnalysis() *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis {
	return s.Analysis
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetExcludeProjectIds() []*int64 {
	return s.ExcludeProjectIds
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetExcludeTaskIds() []*int64 {
	return s.ExcludeTaskIds
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetIncludeProjectIds() []*int64 {
	return s.IncludeProjectIds
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetIncludeTaskIds() []*int64 {
	return s.IncludeTaskIds
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetMode() *string {
	return s.Mode
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetOrder() *string {
	return s.Order
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetParallelism() *int32 {
	return s.Parallelism
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetPriorityWeightStrategy() *string {
	return s.PriorityWeightStrategy
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetRootTaskIds() []*int64 {
	return s.RootTaskIds
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetRunPolicy() *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy {
	return s.RunPolicy
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) GetRuntimeResource() *string {
	return s.RuntimeResource
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetAlert(v *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.Alert = v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetAnalysis(v *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.Analysis = v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetExcludeProjectIds(v []*int64) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.ExcludeProjectIds = v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetExcludeTaskIds(v []*int64) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.ExcludeTaskIds = v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetIncludeProjectIds(v []*int64) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.IncludeProjectIds = v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetIncludeTaskIds(v []*int64) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.IncludeTaskIds = v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetMode(v string) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.Mode = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetOrder(v string) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.Order = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetParallelism(v int32) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.Parallelism = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetPriority(v int32) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.Priority = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetPriorityWeightStrategy(v string) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.PriorityWeightStrategy = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetRootTaskIds(v []*int64) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.RootTaskIds = v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetRunPolicy(v *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.RunPolicy = v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) SetRuntimeResource(v string) *CreateWorkflowInstancesRequestDefaultRunProperties {
	s.RuntimeResource = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunProperties) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstancesRequestDefaultRunPropertiesAlert struct {
	// The alert notification method. Valid values:
	//
	// 	- Sms
	//
	// 	- Mail
	//
	// 	- SmsMail
	//
	// example:
	//
	// Sms
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The alerting policy. Valid values:
	//
	// 	- Success: An alert is reported when data backfill succeeds.
	//
	// 	- Failure: An alert is reported when data backfill fails.
	//
	// 	- SuccessFailure: An alert is reported regardless of whether data backfill succeeds or fails.
	//
	// example:
	//
	// Succes
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateWorkflowInstancesRequestDefaultRunPropertiesAlert) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesRequestDefaultRunPropertiesAlert) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert) GetNoticeType() *string {
	return s.NoticeType
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert) GetType() *string {
	return s.Type
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert) SetNoticeType(v string) *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert {
	s.NoticeType = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert) SetType(v string) *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert {
	s.Type = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis struct {
	// Specifies whether to block running if the analysis fails. If you set the Type parameter to SupplementData, this parameter is required. Valid values: true and false.
	//
	// example:
	//
	// true
	Blocked *bool `json:"Blocked,omitempty" xml:"Blocked,omitempty"`
	// Specifies whether to enable the analysis feature. If you set the Type parameter to SupplementData, this parameter is required. Valid values: true and false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis) GetBlocked() *bool {
	return s.Blocked
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis) SetBlocked(v bool) *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis {
	s.Blocked = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis) SetEnabled(v bool) *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis {
	s.Enabled = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy struct {
	// The end time of running. Configure this parameter in the `hh:mm:ss` format. The time must be in the 24-hour clock. This parameter is required if you configure the RunPolicy parameter.
	//
	// example:
	//
	// 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether the instances can be run immediately during the time period in the future. Valid values: true and false.
	//
	// example:
	//
	// false
	Immediately *bool `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The start time of running. Configure this parameter in the `hh:mm:ss` format. The time must be in the 24-hour clock. This parameter is required if you configure the RunPolicy parameter.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the time period during which the data is backfilled. This parameter is required if you configure the RunPolicy parameter.
	//
	// 	- Daily
	//
	// 	- Weekend
	//
	// example:
	//
	// Daily
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) GetImmediately() *bool {
	return s.Immediately
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) GetType() *string {
	return s.Type
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) SetEndTime(v string) *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy {
	s.EndTime = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) SetImmediately(v bool) *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy {
	s.Immediately = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) SetStartTime(v string) *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy {
	s.StartTime = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) SetType(v string) *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy {
	s.Type = &v
	return s
}

func (s *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstancesRequestPeriods struct {
	// The data timestamps. You can specify up to seven data timestamps.
	//
	// This parameter is required.
	BizDates []*CreateWorkflowInstancesRequestPeriodsBizDates `json:"BizDates,omitempty" xml:"BizDates,omitempty" type:"Repeated"`
	// The end time of data backfill. Configure this parameter in the `hh:mm:ss` format. The time must be in the 24-hour clock. Default value: 23:59:59.
	//
	// If you configure this parameter, you must also configure the StartTime parameter.
	//
	// example:
	//
	// 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of data backfill. Configure this parameter in the `hh:mm:ss` format. The time must be in the 24-hour clock. Default value: 00:00:00.
	//
	// If you configure this parameter, you must also configure the EndTime parameter.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateWorkflowInstancesRequestPeriods) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesRequestPeriods) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesRequestPeriods) GetBizDates() []*CreateWorkflowInstancesRequestPeriodsBizDates {
	return s.BizDates
}

func (s *CreateWorkflowInstancesRequestPeriods) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateWorkflowInstancesRequestPeriods) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateWorkflowInstancesRequestPeriods) SetBizDates(v []*CreateWorkflowInstancesRequestPeriodsBizDates) *CreateWorkflowInstancesRequestPeriods {
	s.BizDates = v
	return s
}

func (s *CreateWorkflowInstancesRequestPeriods) SetEndTime(v string) *CreateWorkflowInstancesRequestPeriods {
	s.EndTime = &v
	return s
}

func (s *CreateWorkflowInstancesRequestPeriods) SetStartTime(v string) *CreateWorkflowInstancesRequestPeriods {
	s.StartTime = &v
	return s
}

func (s *CreateWorkflowInstancesRequestPeriods) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstancesRequestPeriodsBizDates struct {
	// The data timestamp at which data is no longer backfilled. Configure this parameter in the `yyyy-mm-dd` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-24
	EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	// The data timestamp at which the data starts to be backfilled. Configure this parameter in the `yyyy-mm-dd` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-20
	StartBizDate *string `json:"StartBizDate,omitempty" xml:"StartBizDate,omitempty"`
}

func (s CreateWorkflowInstancesRequestPeriodsBizDates) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesRequestPeriodsBizDates) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesRequestPeriodsBizDates) GetEndBizDate() *string {
	return s.EndBizDate
}

func (s *CreateWorkflowInstancesRequestPeriodsBizDates) GetStartBizDate() *string {
	return s.StartBizDate
}

func (s *CreateWorkflowInstancesRequestPeriodsBizDates) SetEndBizDate(v string) *CreateWorkflowInstancesRequestPeriodsBizDates {
	s.EndBizDate = &v
	return s
}

func (s *CreateWorkflowInstancesRequestPeriodsBizDates) SetStartBizDate(v string) *CreateWorkflowInstancesRequestPeriodsBizDates {
	s.StartBizDate = &v
	return s
}

func (s *CreateWorkflowInstancesRequestPeriodsBizDates) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstancesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateWorkflowInstancesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateWorkflowInstancesRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateWorkflowInstancesRequestTags) SetKey(v string) *CreateWorkflowInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *CreateWorkflowInstancesRequestTags) SetValue(v string) *CreateWorkflowInstancesRequestTags {
	s.Value = &v
	return s
}

func (s *CreateWorkflowInstancesRequestTags) Validate() error {
	return dara.Validate(s)
}
