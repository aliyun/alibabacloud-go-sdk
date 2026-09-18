// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPipelinesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPipelinesResponseBody
	GetNextToken() *string
	SetPipelines(v []*ListPipelinesResponseBodyPipelines) *ListPipelinesResponseBody
	GetPipelines() []*ListPipelinesResponseBodyPipelines
	SetRequestId(v string) *ListPipelinesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPipelinesResponseBody
	GetTotalCount() *int32
}

type ListPipelinesResponseBody struct {
	// The maximum number of entries per page that was specified in the request. This value is echoed back.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page. An empty string indicates that the current page is the last page.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of pipelines.
	Pipelines []*ListPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of pipelines that match the filter conditions.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPipelinesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelinesResponseBody) GetPipelines() []*ListPipelinesResponseBodyPipelines {
	return s.Pipelines
}

func (s *ListPipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelinesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPipelinesResponseBody) SetMaxResults(v int32) *ListPipelinesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPipelinesResponseBody) SetNextToken(v string) *ListPipelinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelinesResponseBody) SetPipelines(v []*ListPipelinesResponseBodyPipelines) *ListPipelinesResponseBody {
	s.Pipelines = v
	return s
}

func (s *ListPipelinesResponseBody) SetRequestId(v string) *ListPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelinesResponseBody) SetTotalCount(v int32) *ListPipelinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPipelinesResponseBody) Validate() error {
	if s.Pipelines != nil {
		for _, item := range s.Pipelines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelines struct {
	// The time when the pipeline was created, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the pipeline.
	//
	// example:
	//
	// My pipeline
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy.
	ExecutePolicy *ListPipelinesResponseBodyPipelinesExecutePolicy `json:"executePolicy,omitempty" xml:"executePolicy,omitempty" type:"Struct"`
	// The name of the pipeline.
	//
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The scheduling status. Valid values:
	//
	// - None: no scheduling. This value is returned for RunOnce pipelines.
	//
	// - Active: active.
	//
	// - Paused: paused.
	//
	// - Terminated: terminated.
	//
	// example:
	//
	// None
	ScheduleStatus *string `json:"scheduleStatus,omitempty" xml:"scheduleStatus,omitempty"`
	// The scheduling type. Valid values:
	//
	// - RunOnce: one-time execution.
	//
	// - Scheduled: periodic scheduling.
	//
	// example:
	//
	// RunOnce
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// The pipeline sink (data write destination).
	Sink *ListPipelinesResponseBodyPipelinesSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Struct"`
	// The pipeline data source.
	Source *ListPipelinesResponseBodyPipelinesSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The time when the pipeline was last updated, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The workspace associated with the pipeline.
	//
	// example:
	//
	// my-workspace
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListPipelinesResponseBodyPipelines) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelines) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelines) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPipelinesResponseBodyPipelines) GetDescription() *string {
	return s.Description
}

func (s *ListPipelinesResponseBodyPipelines) GetExecutePolicy() *ListPipelinesResponseBodyPipelinesExecutePolicy {
	return s.ExecutePolicy
}

func (s *ListPipelinesResponseBodyPipelines) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelinesResponseBodyPipelines) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPipelinesResponseBodyPipelines) GetScheduleStatus() *string {
	return s.ScheduleStatus
}

func (s *ListPipelinesResponseBodyPipelines) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListPipelinesResponseBodyPipelines) GetSink() *ListPipelinesResponseBodyPipelinesSink {
	return s.Sink
}

func (s *ListPipelinesResponseBodyPipelines) GetSource() *ListPipelinesResponseBodyPipelinesSource {
	return s.Source
}

func (s *ListPipelinesResponseBodyPipelines) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPipelinesResponseBodyPipelines) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListPipelinesResponseBodyPipelines) SetCreateTime(v string) *ListPipelinesResponseBodyPipelines {
	s.CreateTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetDescription(v string) *ListPipelinesResponseBodyPipelines {
	s.Description = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetExecutePolicy(v *ListPipelinesResponseBodyPipelinesExecutePolicy) *ListPipelinesResponseBodyPipelines {
	s.ExecutePolicy = v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetPipelineName(v string) *ListPipelinesResponseBodyPipelines {
	s.PipelineName = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetRegionId(v string) *ListPipelinesResponseBodyPipelines {
	s.RegionId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetScheduleStatus(v string) *ListPipelinesResponseBodyPipelines {
	s.ScheduleStatus = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetScheduleType(v string) *ListPipelinesResponseBodyPipelines {
	s.ScheduleType = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetSink(v *ListPipelinesResponseBodyPipelinesSink) *ListPipelinesResponseBodyPipelines {
	s.Sink = v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetSource(v *ListPipelinesResponseBodyPipelinesSource) *ListPipelinesResponseBodyPipelines {
	s.Source = v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetUpdateTime(v string) *ListPipelinesResponseBodyPipelines {
	s.UpdateTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetWorkspace(v string) *ListPipelinesResponseBodyPipelines {
	s.Workspace = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) Validate() error {
	if s.ExecutePolicy != nil {
		if err := s.ExecutePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.Sink != nil {
		if err := s.Sink.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelinesExecutePolicy struct {
	// The scheduling mode. Valid values:
	//
	// - RunOnce: one-time execution.
	//
	// - Scheduled: periodic scheduling.
	//
	// example:
	//
	// RunOnce
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The parameters for one-time execution. This parameter has a value only when mode is set to RunOnce.
	RunOnce *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce `json:"runOnce,omitempty" xml:"runOnce,omitempty" type:"Struct"`
	// The parameters for periodic scheduling. This parameter has a value only when mode is set to Scheduled.
	Scheduled *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled `json:"scheduled,omitempty" xml:"scheduled,omitempty" type:"Struct"`
}

func (s ListPipelinesResponseBodyPipelinesExecutePolicy) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesExecutePolicy) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicy) GetMode() *string {
	return s.Mode
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicy) GetRunOnce() *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce {
	return s.RunOnce
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicy) GetScheduled() *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled {
	return s.Scheduled
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicy) SetMode(v string) *ListPipelinesResponseBodyPipelinesExecutePolicy {
	s.Mode = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicy) SetRunOnce(v *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce) *ListPipelinesResponseBodyPipelinesExecutePolicy {
	s.RunOnce = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicy) SetScheduled(v *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled) *ListPipelinesResponseBodyPipelinesExecutePolicy {
	s.Scheduled = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicy) Validate() error {
	if s.RunOnce != nil {
		if err := s.RunOnce.Validate(); err != nil {
			return err
		}
	}
	if s.Scheduled != nil {
		if err := s.Scheduled.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce struct {
	// The start of the time slice, in UNIX millisecond timestamp format.
	//
	// example:
	//
	// 1735660800000
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The end of the time slice, in UNIX millisecond timestamp format.
	//
	// example:
	//
	// 1735747200000
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce) GetFromTime() *int64 {
	return s.FromTime
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce) GetToTime() *int64 {
	return s.ToTime
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce) SetFromTime(v int64) *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce {
	s.FromTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce) SetToTime(v int64) *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce {
	s.ToTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyRunOnce) Validate() error {
	return dara.Validate(s)
}

type ListPipelinesResponseBodyPipelinesExecutePolicyScheduled struct {
	// The scheduling start time, in UNIX millisecond timestamp format.
	//
	// example:
	//
	// 1735660800000
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The scheduling interval, such as 1h or 30m.
	//
	// example:
	//
	// 1h
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesExecutePolicyScheduled) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesExecutePolicyScheduled) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled) GetFromTime() *int64 {
	return s.FromTime
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled) GetInterval() *string {
	return s.Interval
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled) SetFromTime(v int64) *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled {
	s.FromTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled) SetInterval(v string) *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled {
	s.Interval = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesExecutePolicyScheduled) Validate() error {
	return dara.Validate(s)
}

type ListPipelinesResponseBodyPipelinesSink struct {
	// The conditional routing configuration. This parameter is used only when sink.type is set to condition.
	Condition *ListPipelinesResponseBodyPipelinesSinkCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	// The destination dataset configuration for the dataset sink. This parameter is used only when sink.type is set to dataset.
	Dataset *ListPipelinesResponseBodyPipelinesSinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The destination type. Valid values: dataset or condition.
	//
	// example:
	//
	// condition
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSink) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSink) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSink) GetCondition() *ListPipelinesResponseBodyPipelinesSinkCondition {
	return s.Condition
}

func (s *ListPipelinesResponseBodyPipelinesSink) GetDataset() *ListPipelinesResponseBodyPipelinesSinkDataset {
	return s.Dataset
}

func (s *ListPipelinesResponseBodyPipelinesSink) GetType() *string {
	return s.Type
}

func (s *ListPipelinesResponseBodyPipelinesSink) SetCondition(v *ListPipelinesResponseBodyPipelinesSinkCondition) *ListPipelinesResponseBodyPipelinesSink {
	s.Condition = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSink) SetDataset(v *ListPipelinesResponseBodyPipelinesSinkDataset) *ListPipelinesResponseBodyPipelinesSink {
	s.Dataset = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSink) SetType(v string) *ListPipelinesResponseBodyPipelinesSink {
	s.Type = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSink) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelinesSinkCondition struct {
	// The default write destination that is used when no condition route is matched.
	DefaultSink *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink `json:"defaultSink,omitempty" xml:"defaultSink,omitempty" type:"Struct"`
	// The route matching mode. Currently, only all is supported.
	//
	// example:
	//
	// all
	MatchMode *string `json:"matchMode,omitempty" xml:"matchMode,omitempty"`
	// The list of condition routes.
	Routes []*ListPipelinesResponseBodyPipelinesSinkConditionRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s ListPipelinesResponseBodyPipelinesSinkCondition) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSinkCondition) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSinkCondition) GetDefaultSink() *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink {
	return s.DefaultSink
}

func (s *ListPipelinesResponseBodyPipelinesSinkCondition) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListPipelinesResponseBodyPipelinesSinkCondition) GetRoutes() []*ListPipelinesResponseBodyPipelinesSinkConditionRoutes {
	return s.Routes
}

func (s *ListPipelinesResponseBodyPipelinesSinkCondition) SetDefaultSink(v *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink) *ListPipelinesResponseBodyPipelinesSinkCondition {
	s.DefaultSink = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkCondition) SetMatchMode(v string) *ListPipelinesResponseBodyPipelinesSinkCondition {
	s.MatchMode = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkCondition) SetRoutes(v []*ListPipelinesResponseBodyPipelinesSinkConditionRoutes) *ListPipelinesResponseBodyPipelinesSinkCondition {
	s.Routes = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkCondition) Validate() error {
	if s.DefaultSink != nil {
		if err := s.DefaultSink.Validate(); err != nil {
			return err
		}
	}
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink struct {
	// The default destination dataset.
	Dataset *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The type of the default destination. Currently, only dataset is supported.
	//
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink) GetDataset() *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset {
	return s.Dataset
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink) GetType() *string {
	return s.Type
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink) SetDataset(v *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset) *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink {
	s.Dataset = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink) SetType(v string) *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink {
	s.Type = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSink) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset struct {
	// The name of the AgentSpace to which the default destination dataset belongs.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The name of the default destination dataset.
	//
	// example:
	//
	// other-result
	Dataset *string `json:"dataset,omitempty" xml:"dataset,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset) GetDataset() *string {
	return s.Dataset
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset) SetAgentSpace(v string) *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset {
	s.AgentSpace = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset) SetDataset(v string) *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset {
	s.Dataset = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionDefaultSinkDataset) Validate() error {
	return dara.Validate(s)
}

type ListPipelinesResponseBodyPipelinesSinkConditionRoutes struct {
	// The route expression in SPL. Only where, project, and extend are supported.
	//
	// example:
	//
	// 	- | where intent = \\"refund\\"
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The route ID.
	//
	// example:
	//
	// refund
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The write destination of the route.
	Sink *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Struct"`
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionRoutes) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionRoutes) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutes) GetExpression() *string {
	return s.Expression
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutes) GetId() *string {
	return s.Id
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutes) GetSink() *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink {
	return s.Sink
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutes) SetExpression(v string) *ListPipelinesResponseBodyPipelinesSinkConditionRoutes {
	s.Expression = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutes) SetId(v string) *ListPipelinesResponseBodyPipelinesSinkConditionRoutes {
	s.Id = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutes) SetSink(v *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink) *ListPipelinesResponseBodyPipelinesSinkConditionRoutes {
	s.Sink = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutes) Validate() error {
	if s.Sink != nil {
		if err := s.Sink.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink struct {
	// The destination dataset of the route.
	Dataset *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The type of the route destination. Currently, only dataset is supported.
	//
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink) GetDataset() *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset {
	return s.Dataset
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink) GetType() *string {
	return s.Type
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink) SetDataset(v *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset) *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink {
	s.Dataset = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink) SetType(v string) *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink {
	s.Type = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSink) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset struct {
	// The name of the AgentSpace to which the destination dataset belongs.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The name of the destination dataset.
	//
	// example:
	//
	// refund-result
	Dataset *string `json:"dataset,omitempty" xml:"dataset,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset) GetDataset() *string {
	return s.Dataset
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset) SetAgentSpace(v string) *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset {
	s.AgentSpace = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset) SetDataset(v string) *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset {
	s.Dataset = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkConditionRoutesSinkDataset) Validate() error {
	return dara.Validate(s)
}

type ListPipelinesResponseBodyPipelinesSinkDataset struct {
	// The name of the AgentSpace to which the destination dataset belongs.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The name of the destination dataset.
	//
	// example:
	//
	// my-dataset
	Dataset *string `json:"dataset,omitempty" xml:"dataset,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSinkDataset) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSinkDataset) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSinkDataset) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListPipelinesResponseBodyPipelinesSinkDataset) GetDataset() *string {
	return s.Dataset
}

func (s *ListPipelinesResponseBodyPipelinesSinkDataset) SetAgentSpace(v string) *ListPipelinesResponseBodyPipelinesSinkDataset {
	s.AgentSpace = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkDataset) SetDataset(v string) *ListPipelinesResponseBodyPipelinesSinkDataset {
	s.Dataset = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSinkDataset) Validate() error {
	return dara.Validate(s)
}

type ListPipelinesResponseBodyPipelinesSource struct {
	// The dataset datasource config in the current AgentSpace.
	Dataset *ListPipelinesResponseBodyPipelinesSourceDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The Simple Log Service (SLS) Logstore datasource config.
	Logstore *ListPipelinesResponseBodyPipelinesSourceLogstore `json:"logstore,omitempty" xml:"logstore,omitempty" type:"Struct"`
	// The data source type. Valid values: logstore or dataset.
	//
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSource) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSource) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSource) GetDataset() *ListPipelinesResponseBodyPipelinesSourceDataset {
	return s.Dataset
}

func (s *ListPipelinesResponseBodyPipelinesSource) GetLogstore() *ListPipelinesResponseBodyPipelinesSourceLogstore {
	return s.Logstore
}

func (s *ListPipelinesResponseBodyPipelinesSource) GetType() *string {
	return s.Type
}

func (s *ListPipelinesResponseBodyPipelinesSource) SetDataset(v *ListPipelinesResponseBodyPipelinesSourceDataset) *ListPipelinesResponseBodyPipelinesSource {
	s.Dataset = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSource) SetLogstore(v *ListPipelinesResponseBodyPipelinesSourceLogstore) *ListPipelinesResponseBodyPipelinesSource {
	s.Logstore = v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSource) SetType(v string) *ListPipelinesResponseBodyPipelinesSource {
	s.Type = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSource) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	if s.Logstore != nil {
		if err := s.Logstore.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelinesSourceDataset struct {
	// The name of the source dataset.
	//
	// example:
	//
	// my-dataset
	Dataset *string `json:"dataset,omitempty" xml:"dataset,omitempty"`
	// The data filter condition for the dataset.
	//
	// example:
	//
	// status = \\"pending\\"
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSourceDataset) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSourceDataset) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSourceDataset) GetDataset() *string {
	return s.Dataset
}

func (s *ListPipelinesResponseBodyPipelinesSourceDataset) GetFilter() *string {
	return s.Filter
}

func (s *ListPipelinesResponseBodyPipelinesSourceDataset) SetDataset(v string) *ListPipelinesResponseBodyPipelinesSourceDataset {
	s.Dataset = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSourceDataset) SetFilter(v string) *ListPipelinesResponseBodyPipelinesSourceDataset {
	s.Filter = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSourceDataset) Validate() error {
	return dara.Validate(s)
}

type ListPipelinesResponseBodyPipelinesSourceLogstore struct {
	// The name of the SLS Logstore.
	//
	// example:
	//
	// my-sls-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the SLS project.
	//
	// example:
	//
	// my-sls-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The data filtered query statement in SLS query/analysis syntax.
	//
	// example:
	//
	// 	- | SELECT *
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListPipelinesResponseBodyPipelinesSourceLogstore) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelinesSourceLogstore) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelinesSourceLogstore) GetLogstore() *string {
	return s.Logstore
}

func (s *ListPipelinesResponseBodyPipelinesSourceLogstore) GetProject() *string {
	return s.Project
}

func (s *ListPipelinesResponseBodyPipelinesSourceLogstore) GetQuery() *string {
	return s.Query
}

func (s *ListPipelinesResponseBodyPipelinesSourceLogstore) SetLogstore(v string) *ListPipelinesResponseBodyPipelinesSourceLogstore {
	s.Logstore = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSourceLogstore) SetProject(v string) *ListPipelinesResponseBodyPipelinesSourceLogstore {
	s.Project = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSourceLogstore) SetQuery(v string) *ListPipelinesResponseBodyPipelinesSourceLogstore {
	s.Query = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelinesSourceLogstore) Validate() error {
	return dara.Validate(s)
}
