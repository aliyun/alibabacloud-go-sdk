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
	// Specifies whether to run the workflow instance immediately after creation. Default value: true.
	//
	// example:
	//
	// true
	AutoStartEnabled *bool `json:"AutoStartEnabled,omitempty" xml:"AutoStartEnabled,omitempty"`
	// The reason for creating the workflow instance.
	//
	// example:
	//
	// create for test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The runtime configurations.
	DefaultRunProperties *CreateWorkflowInstancesRequestDefaultRunProperties `json:"DefaultRunProperties,omitempty" xml:"DefaultRunProperties,omitempty" type:"Struct"`
	// The project environment. Valid values:
	//
	// - Prod: production
	//
	// - Dev: development
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
	// The data backfill period settings.
	Periods *CreateWorkflowInstancesRequestPeriods `json:"Periods,omitempty" xml:"Periods,omitempty" type:"Struct"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tag creation policy. Valid values:
	//
	// - Append: append mode. New tags are appended to the existing tags inherited from the manual workflow.
	//
	// - Overwrite: overwrite mode. Existing tags of the manual workflow are not inherited. Tags are created directly.
	//
	// example:
	//
	// Append
	TagCreationPolicy *string `json:"TagCreationPolicy,omitempty" xml:"TagCreationPolicy,omitempty"`
	// The list of node labels.
	Tags []*CreateWorkflowInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The node parameters used to set parameters for specific nodes. The value is in JSON format. The key is the node ID, and the value format refers to the node script parameter (the Task.Script.Parameter field in the GetTask response).
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
	// - SupplementData: data backfill. The method for specifying RootTaskIds and IncludeTaskIds varies based on the data backfill pattern. For more information, see the DefaultRunProperties.Mode parameter description.
	//
	// - ManualWorkflow: manual workflow. Set WorkflowId to the ID of the manual workflow. RootTaskIds is optional. If you do not specify RootTaskIds, the default root node list of the manual workflow is used.
	//
	// - Manual: manual node. Only RootTaskIds is required, which specifies the list of manual nodes to run.
	//
	// - SmokeTest: smoke test. Only RootTaskIds is required, which specifies the list of test nodes to run.
	//
	// - TriggerWorkflow: trigger-based workflow. Set WorkflowId to the ID of the trigger-based workflow. IncludeTaskIds is optional. If you do not specify IncludeTaskIds, the entire workflow is run.
	//
	// This parameter is required.
	//
	// example:
	//
	// SupplementData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workflow to which the instance belongs. The WorkflowId for periodic nodes is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow parameters. This parameter takes effect when a unique workflow is specified (`WorkflowId != 1`). For periodic workflows and trigger-based workflows, the format is key=value, and the priority is lower than node parameters. For manual workflows, the format is JSON, and the priority is higher than node parameters.
	//
	// example:
	//
	// "key=value" format:
	//
	// key1=value1 key2=value2
	//
	// JSON format:
	//
	// {"key1":"value1", "key2": "value2"}
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
	if s.DefaultRunProperties != nil {
		if err := s.DefaultRunProperties.Validate(); err != nil {
			return err
		}
	}
	if s.Periods != nil {
		if err := s.Periods.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkflowInstancesRequestDefaultRunProperties struct {
	// The alert configuration.
	Alert *CreateWorkflowInstancesRequestDefaultRunPropertiesAlert `json:"Alert,omitempty" xml:"Alert,omitempty" type:"Struct"`
	// The analysis configuration. This parameter is required when Type is set to SupplementData.
	Analysis *CreateWorkflowInstancesRequestDefaultRunPropertiesAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The list of project IDs to exclude.
	ExcludeProjectIds []*int64 `json:"ExcludeProjectIds,omitempty" xml:"ExcludeProjectIds,omitempty" type:"Repeated"`
	// The list of node IDs to exclude from running.
	ExcludeTaskIds []*int64 `json:"ExcludeTaskIds,omitempty" xml:"ExcludeTaskIds,omitempty" type:"Repeated"`
	// The list of project IDs to include.
	IncludeProjectIds []*int64 `json:"IncludeProjectIds,omitempty" xml:"IncludeProjectIds,omitempty" type:"Repeated"`
	// The list of node IDs to run.
	IncludeTaskIds []*int64 `json:"IncludeTaskIds,omitempty" xml:"IncludeTaskIds,omitempty" type:"Repeated"`
	// The data backfill mode. Default value: ManualSelection. This parameter is required when Type is set to SupplementData. Valid values:
	//
	// - General: general mode. Only one value can be specified for `RootTaskIds`. `IncludeTaskIds` is optional. If you do not specify IncludeTaskIds, the content in `RootTaskIds` is included by default.
	//
	// - ManualSelection: manual selection. Multiple values can be specified for `RootTaskIds`. `IncludeTaskIds` is optional. If you do not specify IncludeTaskIds, the content in `RootTaskIds` is included by default.
	//
	// - Chain: chain mode. `RootTaskIds` is empty. Specify two IDs in `IncludeTaskIds`, which are the start and end nodes.
	//
	// - AllDownstream: all downstream. Only one value can be specified for `RootTaskIds`.
	//
	// example:
	//
	// ManualSelection
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The run order. Default value: Asc. Valid values:
	//
	// - Asc: ascending order by business date.
	//
	// - Desc: descending order by business date.
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of parallel nodes. A value from 2 to 10 specifies the parallelism. A value of 1 specifies serial execution. This parameter is required when Type is set to SupplementData.
	//
	// example:
	//
	// 2
	Parallelism *int32 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The run priority. Valid values: 1 to 11. A larger value indicates a higher priority. This parameter settings only supports manual workflows and trigger-based workflows.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The priority weight policy. This parameter settings only supports manual workflows and trigger-based workflows. Valid values:
	//
	// - `Disable`: disabled (default)
	//
	// - `Upstream`: calculates the total weight of upstream nodes for the current node. The deeper the level, the higher the weight.
	//
	// example:
	//
	// Upstream
	PriorityWeightStrategy *string `json:"PriorityWeightStrategy,omitempty" xml:"PriorityWeightStrategy,omitempty"`
	// The list of root node IDs.
	//
	// - When Type is set to SupplementData, RootTaskIds is required except when Mode is set to Chain.
	//
	// - When Type is set to ManualWorkflow, RootTaskIds is optional. If you do not specify RootTaskIds, the default root node list of the manual workflow is used.
	//
	// - When Type is set to Manual, RootTaskIds is required, which specifies the list of manual nodes to run.
	//
	// - When Type is set to SmokeTest, RootTaskIds is required, which specifies the list of test nodes to run.
	RootTaskIds []*int64 `json:"RootTaskIds,omitempty" xml:"RootTaskIds,omitempty" type:"Repeated"`
	// The run policy. If this field is empty, the node configuration is used.
	RunPolicy *CreateWorkflowInstancesRequestDefaultRunPropertiesRunPolicy `json:"RunPolicy,omitempty" xml:"RunPolicy,omitempty" type:"Struct"`
	// The identifier of the custom schedule resource group. If this field is empty, the node configuration is used.
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
	if s.Alert != nil {
		if err := s.Alert.Validate(); err != nil {
			return err
		}
	}
	if s.Analysis != nil {
		if err := s.Analysis.Validate(); err != nil {
			return err
		}
	}
	if s.RunPolicy != nil {
		if err := s.RunPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkflowInstancesRequestDefaultRunPropertiesAlert struct {
	// The notification method. Valid values:
	//
	// - Sms: SMS only
	//
	// - Mail: email only
	//
	// - SmsMail: SMS and email
	//
	// example:
	//
	// Sms
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The alert policy. Valid values:
	//
	// - Success: alert on success
	//
	// - Failure: alert on failure
	//
	// - SuccessFailure: alert on both success and failure
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
	// Specifies whether to block running when the analysis does not pass. This parameter is required when Type is set to SupplementData.
	//
	// example:
	//
	// true
	Blocked *bool `json:"Blocked,omitempty" xml:"Blocked,omitempty"`
	// Specifies whether to enable analysis. This parameter is required when Type is set to SupplementData.
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
	// The end run time. Format: `hh:mm:ss` in 24-hour format. This field is required if you set the run policy.
	//
	// example:
	//
	// 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether the instance can start running immediately if the run time is in the future. Default value: false.
	//
	// example:
	//
	// false
	Immediately *bool `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The start run time. Format: `hh:mm:ss` in 24-hour format. This field is required if you set the run policy.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time period type. This field is required if you set the run policy. Valid values:
	//
	// - Daily: every day
	//
	// - Weekend: weekends only
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
	// The list of business dates. You can specify up to 7 business date ranges.
	//
	// This parameter is required.
	BizDates []*CreateWorkflowInstancesRequestPeriodsBizDates `json:"BizDates,omitempty" xml:"BizDates,omitempty" type:"Repeated"`
	// The end period time. Format: `hh:mm:ss` in 24-hour format. Default value: 23:59:59.
	//
	// If you specify this field, you must also specify StartTime.
	//
	// example:
	//
	// 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start period time. Format: `hh:mm:ss` in 24-hour format. Default value: 00:00:00.
	//
	// If you specify this field, you must also specify EndTime.
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
	if s.BizDates != nil {
		for _, item := range s.BizDates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkflowInstancesRequestPeriodsBizDates struct {
	// The end business date. Format: `yyyy-mm-dd`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-24
	EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	// The start business date. Format: `yyyy-mm-dd`.
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
	// The label key.
	//
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The label value.
	//
	// example:
	//
	// tagValue
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
