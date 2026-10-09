// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommittedWatermark(v int64) *GetPipelineResponseBody
	GetCommittedWatermark() *int64
	SetCreateTime(v string) *GetPipelineResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetPipelineResponseBody
	GetDescription() *string
	SetExecutePolicy(v *GetPipelineResponseBodyExecutePolicy) *GetPipelineResponseBody
	GetExecutePolicy() *GetPipelineResponseBodyExecutePolicy
	SetNextTriggerTime(v int64) *GetPipelineResponseBody
	GetNextTriggerTime() *int64
	SetPipeline(v *GetPipelineResponseBodyPipeline) *GetPipelineResponseBody
	GetPipeline() *GetPipelineResponseBodyPipeline
	SetPipelineName(v string) *GetPipelineResponseBody
	GetPipelineName() *string
	SetRegionId(v string) *GetPipelineResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetPipelineResponseBody
	GetRequestId() *string
	SetScheduleStatus(v string) *GetPipelineResponseBody
	GetScheduleStatus() *string
	SetScheduleType(v string) *GetPipelineResponseBody
	GetScheduleType() *string
	SetSink(v *GetPipelineResponseBodySink) *GetPipelineResponseBody
	GetSink() *GetPipelineResponseBodySink
	SetSource(v *GetPipelineResponseBodySource) *GetPipelineResponseBody
	GetSource() *GetPipelineResponseBodySource
	SetUpdateTime(v string) *GetPipelineResponseBody
	GetUpdateTime() *string
	SetWorkspace(v string) *GetPipelineResponseBody
	GetWorkspace() *string
}

type GetPipelineResponseBody struct {
	// The committed watermark in UNIX seconds.
	//
	// example:
	//
	// 1735660800
	CommittedWatermark *int64 `json:"committedWatermark,omitempty" xml:"committedWatermark,omitempty"`
	// The pipeline creation time in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The pipeline description.
	//
	// example:
	//
	// My pipeline
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy.
	//
	// example:
	//
	// {"mode":"RunOnce","runOnce":{"fromTime":1735660800,"toTime":1735664400}}
	ExecutePolicy *GetPipelineResponseBodyExecutePolicy `json:"executePolicy,omitempty" xml:"executePolicy,omitempty" type:"Struct"`
	// The next scheduling trigger time in UNIX seconds.
	//
	// example:
	//
	// 1735661100
	NextTriggerTime *int64 `json:"nextTriggerTime,omitempty" xml:"nextTriggerTime,omitempty"`
	// The pipeline configuration for node orchestration.
	//
	// example:
	//
	// {"nodes":[{"id":"select-fields","type":"project","parameters":{"question":"user_query"}}]}
	Pipeline *GetPipelineResponseBodyPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// The pipeline name.
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
	// The request ID used to locate the request during troubleshooting.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scheduling status. Valid values: None (no scheduling), Active (active), Paused (paused), and Terminated (terminated).
	//
	// example:
	//
	// Active
	ScheduleStatus *string `json:"scheduleStatus,omitempty" xml:"scheduleStatus,omitempty"`
	// The scheduling type. Valid values: RunOnce (single execution), Scheduled (periodic scheduling), and Continuous (continuous execution driven by trace source signals).
	//
	// example:
	//
	// RunOnce
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// The pipeline sink, which is the destination for data writing.
	Sink *GetPipelineResponseBodySink `json:"sink,omitempty" xml:"sink,omitempty" type:"Struct"`
	// The pipeline data source.
	//
	// example:
	//
	// {"type":"logstore","logstore":{"project":"my-sls-project","logstore":"agent-logs"},"inputFields":[{"name":"question","type":"text"}]}
	Source *GetPipelineResponseBodySource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The last update time of the pipeline, in ISO 8601 UTC format.
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

func (s GetPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBody) GetCommittedWatermark() *int64 {
	return s.CommittedWatermark
}

func (s *GetPipelineResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPipelineResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetPipelineResponseBody) GetExecutePolicy() *GetPipelineResponseBodyExecutePolicy {
	return s.ExecutePolicy
}

func (s *GetPipelineResponseBody) GetNextTriggerTime() *int64 {
	return s.NextTriggerTime
}

func (s *GetPipelineResponseBody) GetPipeline() *GetPipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *GetPipelineResponseBody) GetPipelineName() *string {
	return s.PipelineName
}

func (s *GetPipelineResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineResponseBody) GetScheduleStatus() *string {
	return s.ScheduleStatus
}

func (s *GetPipelineResponseBody) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetPipelineResponseBody) GetSink() *GetPipelineResponseBodySink {
	return s.Sink
}

func (s *GetPipelineResponseBody) GetSource() *GetPipelineResponseBodySource {
	return s.Source
}

func (s *GetPipelineResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPipelineResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetPipelineResponseBody) SetCommittedWatermark(v int64) *GetPipelineResponseBody {
	s.CommittedWatermark = &v
	return s
}

func (s *GetPipelineResponseBody) SetCreateTime(v string) *GetPipelineResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineResponseBody) SetDescription(v string) *GetPipelineResponseBody {
	s.Description = &v
	return s
}

func (s *GetPipelineResponseBody) SetExecutePolicy(v *GetPipelineResponseBodyExecutePolicy) *GetPipelineResponseBody {
	s.ExecutePolicy = v
	return s
}

func (s *GetPipelineResponseBody) SetNextTriggerTime(v int64) *GetPipelineResponseBody {
	s.NextTriggerTime = &v
	return s
}

func (s *GetPipelineResponseBody) SetPipeline(v *GetPipelineResponseBodyPipeline) *GetPipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetPipelineResponseBody) SetPipelineName(v string) *GetPipelineResponseBody {
	s.PipelineName = &v
	return s
}

func (s *GetPipelineResponseBody) SetRegionId(v string) *GetPipelineResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetPipelineResponseBody) SetRequestId(v string) *GetPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineResponseBody) SetScheduleStatus(v string) *GetPipelineResponseBody {
	s.ScheduleStatus = &v
	return s
}

func (s *GetPipelineResponseBody) SetScheduleType(v string) *GetPipelineResponseBody {
	s.ScheduleType = &v
	return s
}

func (s *GetPipelineResponseBody) SetSink(v *GetPipelineResponseBodySink) *GetPipelineResponseBody {
	s.Sink = v
	return s
}

func (s *GetPipelineResponseBody) SetSource(v *GetPipelineResponseBodySource) *GetPipelineResponseBody {
	s.Source = v
	return s
}

func (s *GetPipelineResponseBody) SetUpdateTime(v string) *GetPipelineResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetPipelineResponseBody) SetWorkspace(v string) *GetPipelineResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetPipelineResponseBody) Validate() error {
	if s.ExecutePolicy != nil {
		if err := s.ExecutePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
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

type GetPipelineResponseBodyExecutePolicy struct {
	// The continuous execution configuration. It is used when the type is trace, and the processing frequency is a fixed value managed by the server.
	//
	// example:
	//
	// {"fromTime":1735660800}
	Continuous *GetPipelineResponseBodyExecutePolicyContinuous `json:"continuous,omitempty" xml:"continuous,omitempty" type:"Struct"`
	// The scheduling mode. Valid values: RunOnce (single execution), Scheduled (periodic execution), and Continuous (continuous execution, only for trace data sources; the processing frequency is a fixed value managed by the server, and automatic processing occurs at minute intervals after trace completion).
	//
	// example:
	//
	// scheduled
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The single execution configuration. This parameter is required only when the mode is RunOnce.
	//
	// example:
	//
	// {"fromTime":1735660800,"toTime":1735664400}
	RunOnce *GetPipelineResponseBodyExecutePolicyRunOnce `json:"runOnce,omitempty" xml:"runOnce,omitempty" type:"Struct"`
	// The periodic scheduling configuration. This parameter is required only when the mode is Scheduled.
	//
	// example:
	//
	// {"interval":"1h","fromTime":1735660800}
	Scheduled *GetPipelineResponseBodyExecutePolicyScheduled `json:"scheduled,omitempty" xml:"scheduled,omitempty" type:"Struct"`
}

func (s GetPipelineResponseBodyExecutePolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyExecutePolicy) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyExecutePolicy) GetContinuous() *GetPipelineResponseBodyExecutePolicyContinuous {
	return s.Continuous
}

func (s *GetPipelineResponseBodyExecutePolicy) GetMode() *string {
	return s.Mode
}

func (s *GetPipelineResponseBodyExecutePolicy) GetRunOnce() *GetPipelineResponseBodyExecutePolicyRunOnce {
	return s.RunOnce
}

func (s *GetPipelineResponseBodyExecutePolicy) GetScheduled() *GetPipelineResponseBodyExecutePolicyScheduled {
	return s.Scheduled
}

func (s *GetPipelineResponseBodyExecutePolicy) SetContinuous(v *GetPipelineResponseBodyExecutePolicyContinuous) *GetPipelineResponseBodyExecutePolicy {
	s.Continuous = v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicy) SetMode(v string) *GetPipelineResponseBodyExecutePolicy {
	s.Mode = &v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicy) SetRunOnce(v *GetPipelineResponseBodyExecutePolicyRunOnce) *GetPipelineResponseBodyExecutePolicy {
	s.RunOnce = v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicy) SetScheduled(v *GetPipelineResponseBodyExecutePolicyScheduled) *GetPipelineResponseBodyExecutePolicy {
	s.Scheduled = v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicy) Validate() error {
	if s.Continuous != nil {
		if err := s.Continuous.Validate(); err != nil {
			return err
		}
	}
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

type GetPipelineResponseBodyExecutePolicyContinuous struct {
	// The bootstrap start time in UNIX seconds. It has the same precision as runOnce.fromTime or scheduled.fromTime. Millisecond values greater than or equal to 1e12 are automatically converted. The cursor starts from this time aligned to the grid and catches up window by window. After catching up, it switches to minute intervals. By default, it starts from the current time and processes only incremental data.
	//
	// example:
	//
	// 1735660800
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
}

func (s GetPipelineResponseBodyExecutePolicyContinuous) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyExecutePolicyContinuous) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyExecutePolicyContinuous) GetFromTime() *int64 {
	return s.FromTime
}

func (s *GetPipelineResponseBodyExecutePolicyContinuous) SetFromTime(v int64) *GetPipelineResponseBodyExecutePolicyContinuous {
	s.FromTime = &v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicyContinuous) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodyExecutePolicyRunOnce struct {
	// The start time of the data processing window in UNIX seconds. The value must be less than the toTime value.
	//
	// example:
	//
	// 1735660800
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The end time of the data processing window in UNIX seconds. The value must be greater than the fromTime value.
	//
	// example:
	//
	// 1735747200
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s GetPipelineResponseBodyExecutePolicyRunOnce) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyExecutePolicyRunOnce) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyExecutePolicyRunOnce) GetFromTime() *int64 {
	return s.FromTime
}

func (s *GetPipelineResponseBodyExecutePolicyRunOnce) GetToTime() *int64 {
	return s.ToTime
}

func (s *GetPipelineResponseBodyExecutePolicyRunOnce) SetFromTime(v int64) *GetPipelineResponseBodyExecutePolicyRunOnce {
	s.FromTime = &v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicyRunOnce) SetToTime(v int64) *GetPipelineResponseBodyExecutePolicyRunOnce {
	s.ToTime = &v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicyRunOnce) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodyExecutePolicyScheduled struct {
	// The scheduling start time in UNIX seconds. It has the same precision as runOnce.fromTime. Millisecond values greater than or equal to 1e12 are automatically converted.
	//
	// example:
	//
	// 1735660800
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The scheduling interval. Valid values: 1h, 6h, 12h, and 1d.
	//
	// example:
	//
	// 1h
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
}

func (s GetPipelineResponseBodyExecutePolicyScheduled) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyExecutePolicyScheduled) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyExecutePolicyScheduled) GetFromTime() *int64 {
	return s.FromTime
}

func (s *GetPipelineResponseBodyExecutePolicyScheduled) GetInterval() *string {
	return s.Interval
}

func (s *GetPipelineResponseBodyExecutePolicyScheduled) SetFromTime(v int64) *GetPipelineResponseBodyExecutePolicyScheduled {
	s.FromTime = &v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicyScheduled) SetInterval(v string) *GetPipelineResponseBodyExecutePolicyScheduled {
	s.Interval = &v
	return s
}

func (s *GetPipelineResponseBodyExecutePolicyScheduled) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodyPipeline struct {
	// The list of nodes.
	//
	// example:
	//
	// [{"id":"select-fields","type":"project","parameters":{}}]
	Nodes []*GetPipelineResponseBodyPipelineNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s GetPipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipeline) GetNodes() []*GetPipelineResponseBodyPipelineNodes {
	return s.Nodes
}

func (s *GetPipelineResponseBodyPipeline) SetNodes(v []*GetPipelineResponseBodyPipelineNodes) *GetPipelineResponseBodyPipeline {
	s.Nodes = v
	return s
}

func (s *GetPipelineResponseBodyPipeline) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPipelineResponseBodyPipelineNodes struct {
	// The node ID.
	//
	// example:
	//
	// node-1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The node parameters in a key-value structure. The parameters vary based on the node type.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The node type.
	//
	// example:
	//
	// transform
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodyPipelineNodes) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyPipelineNodes) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelineNodes) GetId() *string {
	return s.Id
}

func (s *GetPipelineResponseBodyPipelineNodes) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetPipelineResponseBodyPipelineNodes) GetType() *string {
	return s.Type
}

func (s *GetPipelineResponseBodyPipelineNodes) SetId(v string) *GetPipelineResponseBodyPipelineNodes {
	s.Id = &v
	return s
}

func (s *GetPipelineResponseBodyPipelineNodes) SetParameters(v map[string]interface{}) *GetPipelineResponseBodyPipelineNodes {
	s.Parameters = v
	return s
}

func (s *GetPipelineResponseBodyPipelineNodes) SetType(v string) *GetPipelineResponseBodyPipelineNodes {
	s.Type = &v
	return s
}

func (s *GetPipelineResponseBodyPipelineNodes) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodySink struct {
	// The conditional routing configuration. This configuration is used only when the sink.type is condition.
	Condition *GetPipelineResponseBodySinkCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	// The destination dataset configuration for the dataset sink. This is used only when sink.type is set to dataset.
	Dataset *GetPipelineResponseBodySinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The destination type. Valid values: dataset and condition.
	//
	// example:
	//
	// condition
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodySink) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySink) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySink) GetCondition() *GetPipelineResponseBodySinkCondition {
	return s.Condition
}

func (s *GetPipelineResponseBodySink) GetDataset() *GetPipelineResponseBodySinkDataset {
	return s.Dataset
}

func (s *GetPipelineResponseBodySink) GetType() *string {
	return s.Type
}

func (s *GetPipelineResponseBodySink) SetCondition(v *GetPipelineResponseBodySinkCondition) *GetPipelineResponseBodySink {
	s.Condition = v
	return s
}

func (s *GetPipelineResponseBodySink) SetDataset(v *GetPipelineResponseBodySinkDataset) *GetPipelineResponseBodySink {
	s.Dataset = v
	return s
}

func (s *GetPipelineResponseBodySink) SetType(v string) *GetPipelineResponseBodySink {
	s.Type = &v
	return s
}

func (s *GetPipelineResponseBodySink) Validate() error {
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

type GetPipelineResponseBodySinkCondition struct {
	// The default write destination used when no conditional route is matched.
	DefaultSink *GetPipelineResponseBodySinkConditionDefaultSink `json:"defaultSink,omitempty" xml:"defaultSink,omitempty" type:"Struct"`
	// The route matching mode. Currently, only all is supported.
	//
	// example:
	//
	// all
	MatchMode *string `json:"matchMode,omitempty" xml:"matchMode,omitempty"`
	// The list of conditional routes.
	Routes []*GetPipelineResponseBodySinkConditionRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s GetPipelineResponseBodySinkCondition) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySinkCondition) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySinkCondition) GetDefaultSink() *GetPipelineResponseBodySinkConditionDefaultSink {
	return s.DefaultSink
}

func (s *GetPipelineResponseBodySinkCondition) GetMatchMode() *string {
	return s.MatchMode
}

func (s *GetPipelineResponseBodySinkCondition) GetRoutes() []*GetPipelineResponseBodySinkConditionRoutes {
	return s.Routes
}

func (s *GetPipelineResponseBodySinkCondition) SetDefaultSink(v *GetPipelineResponseBodySinkConditionDefaultSink) *GetPipelineResponseBodySinkCondition {
	s.DefaultSink = v
	return s
}

func (s *GetPipelineResponseBodySinkCondition) SetMatchMode(v string) *GetPipelineResponseBodySinkCondition {
	s.MatchMode = &v
	return s
}

func (s *GetPipelineResponseBodySinkCondition) SetRoutes(v []*GetPipelineResponseBodySinkConditionRoutes) *GetPipelineResponseBodySinkCondition {
	s.Routes = v
	return s
}

func (s *GetPipelineResponseBodySinkCondition) Validate() error {
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

type GetPipelineResponseBodySinkConditionDefaultSink struct {
	// The default destination dataset.
	Dataset *GetPipelineResponseBodySinkConditionDefaultSinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The default destination type. Currently, only dataset is supported.
	//
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodySinkConditionDefaultSink) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySinkConditionDefaultSink) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySinkConditionDefaultSink) GetDataset() *GetPipelineResponseBodySinkConditionDefaultSinkDataset {
	return s.Dataset
}

func (s *GetPipelineResponseBodySinkConditionDefaultSink) GetType() *string {
	return s.Type
}

func (s *GetPipelineResponseBodySinkConditionDefaultSink) SetDataset(v *GetPipelineResponseBodySinkConditionDefaultSinkDataset) *GetPipelineResponseBodySinkConditionDefaultSink {
	s.Dataset = v
	return s
}

func (s *GetPipelineResponseBodySinkConditionDefaultSink) SetType(v string) *GetPipelineResponseBodySinkConditionDefaultSink {
	s.Type = &v
	return s
}

func (s *GetPipelineResponseBodySinkConditionDefaultSink) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodySinkConditionDefaultSinkDataset struct {
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

func (s GetPipelineResponseBodySinkConditionDefaultSinkDataset) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySinkConditionDefaultSinkDataset) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySinkConditionDefaultSinkDataset) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetPipelineResponseBodySinkConditionDefaultSinkDataset) GetDataset() *string {
	return s.Dataset
}

func (s *GetPipelineResponseBodySinkConditionDefaultSinkDataset) SetAgentSpace(v string) *GetPipelineResponseBodySinkConditionDefaultSinkDataset {
	s.AgentSpace = &v
	return s
}

func (s *GetPipelineResponseBodySinkConditionDefaultSinkDataset) SetDataset(v string) *GetPipelineResponseBodySinkConditionDefaultSinkDataset {
	s.Dataset = &v
	return s
}

func (s *GetPipelineResponseBodySinkConditionDefaultSinkDataset) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodySinkConditionRoutes struct {
	// The route expression in Search Processing Language (SPL). Only where, project, and extend are supported.
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
	// The routing write destination.
	Sink *GetPipelineResponseBodySinkConditionRoutesSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Struct"`
}

func (s GetPipelineResponseBodySinkConditionRoutes) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySinkConditionRoutes) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySinkConditionRoutes) GetExpression() *string {
	return s.Expression
}

func (s *GetPipelineResponseBodySinkConditionRoutes) GetId() *string {
	return s.Id
}

func (s *GetPipelineResponseBodySinkConditionRoutes) GetSink() *GetPipelineResponseBodySinkConditionRoutesSink {
	return s.Sink
}

func (s *GetPipelineResponseBodySinkConditionRoutes) SetExpression(v string) *GetPipelineResponseBodySinkConditionRoutes {
	s.Expression = &v
	return s
}

func (s *GetPipelineResponseBodySinkConditionRoutes) SetId(v string) *GetPipelineResponseBodySinkConditionRoutes {
	s.Id = &v
	return s
}

func (s *GetPipelineResponseBodySinkConditionRoutes) SetSink(v *GetPipelineResponseBodySinkConditionRoutesSink) *GetPipelineResponseBodySinkConditionRoutes {
	s.Sink = v
	return s
}

func (s *GetPipelineResponseBodySinkConditionRoutes) Validate() error {
	if s.Sink != nil {
		if err := s.Sink.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodySinkConditionRoutesSink struct {
	// The routing destination dataset.
	Dataset *GetPipelineResponseBodySinkConditionRoutesSinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The routing destination type. Currently, only dataset is supported.
	//
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodySinkConditionRoutesSink) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySinkConditionRoutesSink) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySinkConditionRoutesSink) GetDataset() *GetPipelineResponseBodySinkConditionRoutesSinkDataset {
	return s.Dataset
}

func (s *GetPipelineResponseBodySinkConditionRoutesSink) GetType() *string {
	return s.Type
}

func (s *GetPipelineResponseBodySinkConditionRoutesSink) SetDataset(v *GetPipelineResponseBodySinkConditionRoutesSinkDataset) *GetPipelineResponseBodySinkConditionRoutesSink {
	s.Dataset = v
	return s
}

func (s *GetPipelineResponseBodySinkConditionRoutesSink) SetType(v string) *GetPipelineResponseBodySinkConditionRoutesSink {
	s.Type = &v
	return s
}

func (s *GetPipelineResponseBodySinkConditionRoutesSink) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodySinkConditionRoutesSinkDataset struct {
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

func (s GetPipelineResponseBodySinkConditionRoutesSinkDataset) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySinkConditionRoutesSinkDataset) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySinkConditionRoutesSinkDataset) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetPipelineResponseBodySinkConditionRoutesSinkDataset) GetDataset() *string {
	return s.Dataset
}

func (s *GetPipelineResponseBodySinkConditionRoutesSinkDataset) SetAgentSpace(v string) *GetPipelineResponseBodySinkConditionRoutesSinkDataset {
	s.AgentSpace = &v
	return s
}

func (s *GetPipelineResponseBodySinkConditionRoutesSinkDataset) SetDataset(v string) *GetPipelineResponseBodySinkConditionRoutesSinkDataset {
	s.Dataset = &v
	return s
}

func (s *GetPipelineResponseBodySinkConditionRoutesSinkDataset) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodySinkDataset struct {
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

func (s GetPipelineResponseBodySinkDataset) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySinkDataset) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySinkDataset) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetPipelineResponseBodySinkDataset) GetDataset() *string {
	return s.Dataset
}

func (s *GetPipelineResponseBodySinkDataset) SetAgentSpace(v string) *GetPipelineResponseBodySinkDataset {
	s.AgentSpace = &v
	return s
}

func (s *GetPipelineResponseBodySinkDataset) SetDataset(v string) *GetPipelineResponseBodySinkDataset {
	s.Dataset = &v
	return s
}

func (s *GetPipelineResponseBodySinkDataset) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodySource struct {
	// The dataset datasource config in the current AgentSpace.
	//
	// example:
	//
	// {"dataset":"my-dataset","filter":"status = \\"pending\\""}
	Dataset *GetPipelineResponseBodySourceDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The input fields and field types. This applies to all data source types.
	//
	// example:
	//
	// [{"name":"question","type":"text"}]
	InputFields []*GetPipelineResponseBodySourceInputFields `json:"inputFields,omitempty" xml:"inputFields,omitempty" type:"Repeated"`
	// The SLS Logstore datasource config.
	//
	// example:
	//
	// {"project":"my-sls-project","logstore":"agent-logs"}
	Logstore *GetPipelineResponseBodySourceLogstore `json:"logstore,omitempty" xml:"logstore,omitempty" type:"Struct"`
	// The trajectory data configuration. This is optional and takes effect only when the type is set to trace. It retrieves ATIF standard trajectory data from the trajectory scrubbing service and extends it by feature.
	//
	// example:
	//
	// {"enrich":{"enabled":true,"columns":["input","output"]}}
	Trajectory *GetPipelineResponseBodySourceTrajectory `json:"trajectory,omitempty" xml:"trajectory,omitempty" type:"Struct"`
	// The data source type. Valid values: logstore, dataset, and trace. The trace value indicates a trajectory signal-driven processing mode. The validity of the enum is verified by the server.
	//
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodySource) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySource) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySource) GetDataset() *GetPipelineResponseBodySourceDataset {
	return s.Dataset
}

func (s *GetPipelineResponseBodySource) GetInputFields() []*GetPipelineResponseBodySourceInputFields {
	return s.InputFields
}

func (s *GetPipelineResponseBodySource) GetLogstore() *GetPipelineResponseBodySourceLogstore {
	return s.Logstore
}

func (s *GetPipelineResponseBodySource) GetTrajectory() *GetPipelineResponseBodySourceTrajectory {
	return s.Trajectory
}

func (s *GetPipelineResponseBodySource) GetType() *string {
	return s.Type
}

func (s *GetPipelineResponseBodySource) SetDataset(v *GetPipelineResponseBodySourceDataset) *GetPipelineResponseBodySource {
	s.Dataset = v
	return s
}

func (s *GetPipelineResponseBodySource) SetInputFields(v []*GetPipelineResponseBodySourceInputFields) *GetPipelineResponseBodySource {
	s.InputFields = v
	return s
}

func (s *GetPipelineResponseBodySource) SetLogstore(v *GetPipelineResponseBodySourceLogstore) *GetPipelineResponseBodySource {
	s.Logstore = v
	return s
}

func (s *GetPipelineResponseBodySource) SetTrajectory(v *GetPipelineResponseBodySourceTrajectory) *GetPipelineResponseBodySource {
	s.Trajectory = v
	return s
}

func (s *GetPipelineResponseBodySource) SetType(v string) *GetPipelineResponseBodySource {
	s.Type = &v
	return s
}

func (s *GetPipelineResponseBodySource) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	if s.InputFields != nil {
		for _, item := range s.InputFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Logstore != nil {
		if err := s.Logstore.Validate(); err != nil {
			return err
		}
	}
	if s.Trajectory != nil {
		if err := s.Trajectory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodySourceDataset struct {
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

func (s GetPipelineResponseBodySourceDataset) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySourceDataset) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySourceDataset) GetDataset() *string {
	return s.Dataset
}

func (s *GetPipelineResponseBodySourceDataset) GetFilter() *string {
	return s.Filter
}

func (s *GetPipelineResponseBodySourceDataset) SetDataset(v string) *GetPipelineResponseBodySourceDataset {
	s.Dataset = &v
	return s
}

func (s *GetPipelineResponseBodySourceDataset) SetFilter(v string) *GetPipelineResponseBodySourceDataset {
	s.Filter = &v
	return s
}

func (s *GetPipelineResponseBodySourceDataset) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodySourceInputFields struct {
	// The name of the field.
	//
	// example:
	//
	// question
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The field type. Valid values: text, long, double, and json.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodySourceInputFields) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySourceInputFields) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySourceInputFields) GetName() *string {
	return s.Name
}

func (s *GetPipelineResponseBodySourceInputFields) GetType() *string {
	return s.Type
}

func (s *GetPipelineResponseBodySourceInputFields) SetName(v string) *GetPipelineResponseBodySourceInputFields {
	s.Name = &v
	return s
}

func (s *GetPipelineResponseBodySourceInputFields) SetType(v string) *GetPipelineResponseBodySourceInputFields {
	s.Type = &v
	return s
}

func (s *GetPipelineResponseBodySourceInputFields) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodySourceLogstore struct {
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
	// The data filtered query statement (SLS query and analysis syntax).
	//
	// example:
	//
	// 	- | SELECT *
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetPipelineResponseBodySourceLogstore) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySourceLogstore) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySourceLogstore) GetLogstore() *string {
	return s.Logstore
}

func (s *GetPipelineResponseBodySourceLogstore) GetProject() *string {
	return s.Project
}

func (s *GetPipelineResponseBodySourceLogstore) GetQuery() *string {
	return s.Query
}

func (s *GetPipelineResponseBodySourceLogstore) SetLogstore(v string) *GetPipelineResponseBodySourceLogstore {
	s.Logstore = &v
	return s
}

func (s *GetPipelineResponseBodySourceLogstore) SetProject(v string) *GetPipelineResponseBodySourceLogstore {
	s.Project = &v
	return s
}

func (s *GetPipelineResponseBodySourceLogstore) SetQuery(v string) *GetPipelineResponseBodySourceLogstore {
	s.Query = &v
	return s
}

func (s *GetPipelineResponseBodySourceLogstore) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodySourceTrajectory struct {
	// The trajectory enrichment. It mounts trajectory data into the scrubbing results by trace_id. When writing to a dataset, the data is stored in the fixed column agent_trajectory, where the column value is the trajectory JSON content.
	//
	// example:
	//
	// {"enabled":true,"columns":["input","output"]}
	Enrich *GetPipelineResponseBodySourceTrajectoryEnrich `json:"enrich,omitempty" xml:"enrich,omitempty" type:"Struct"`
}

func (s GetPipelineResponseBodySourceTrajectory) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySourceTrajectory) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySourceTrajectory) GetEnrich() *GetPipelineResponseBodySourceTrajectoryEnrich {
	return s.Enrich
}

func (s *GetPipelineResponseBodySourceTrajectory) SetEnrich(v *GetPipelineResponseBodySourceTrajectoryEnrich) *GetPipelineResponseBodySourceTrajectory {
	s.Enrich = v
	return s
}

func (s *GetPipelineResponseBodySourceTrajectory) Validate() error {
	if s.Enrich != nil {
		if err := s.Enrich.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodySourceTrajectoryEnrich struct {
	// The enrichment column list. This is retained for compatibility. The current implementation outputs a single fixed column agent_trajectory, and this parameter no longer affects the output.
	//
	// example:
	//
	// ["input","output","session_id"]
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// Specifies whether trajectory enrichment is enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetPipelineResponseBodySourceTrajectoryEnrich) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySourceTrajectoryEnrich) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySourceTrajectoryEnrich) GetColumns() []*string {
	return s.Columns
}

func (s *GetPipelineResponseBodySourceTrajectoryEnrich) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetPipelineResponseBodySourceTrajectoryEnrich) SetColumns(v []*string) *GetPipelineResponseBodySourceTrajectoryEnrich {
	s.Columns = v
	return s
}

func (s *GetPipelineResponseBodySourceTrajectoryEnrich) SetEnabled(v bool) *GetPipelineResponseBodySourceTrajectoryEnrich {
	s.Enabled = &v
	return s
}

func (s *GetPipelineResponseBodySourceTrajectoryEnrich) Validate() error {
	return dara.Validate(s)
}
