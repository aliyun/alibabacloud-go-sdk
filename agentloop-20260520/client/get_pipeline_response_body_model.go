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
	// example:
	//
	// 1735660800
	CommittedWatermark *int64 `json:"committedWatermark,omitempty" xml:"committedWatermark,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 我的流水线
	Description   *string                               `json:"description,omitempty" xml:"description,omitempty"`
	ExecutePolicy *GetPipelineResponseBodyExecutePolicy `json:"executePolicy,omitempty" xml:"executePolicy,omitempty" type:"Struct"`
	// example:
	//
	// 1735661100
	NextTriggerTime *int64                           `json:"nextTriggerTime,omitempty" xml:"nextTriggerTime,omitempty"`
	Pipeline        *GetPipelineResponseBodyPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// Active
	ScheduleStatus *string                        `json:"scheduleStatus,omitempty" xml:"scheduleStatus,omitempty"`
	Sink           *GetPipelineResponseBodySink   `json:"sink,omitempty" xml:"sink,omitempty" type:"Struct"`
	Source         *GetPipelineResponseBodySource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
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
	// example:
	//
	// Scheduled
	Mode      *string                                        `json:"mode,omitempty" xml:"mode,omitempty"`
	RunOnce   *GetPipelineResponseBodyExecutePolicyRunOnce   `json:"runOnce,omitempty" xml:"runOnce,omitempty" type:"Struct"`
	Scheduled *GetPipelineResponseBodyExecutePolicyScheduled `json:"scheduled,omitempty" xml:"scheduled,omitempty" type:"Struct"`
}

func (s GetPipelineResponseBodyExecutePolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyExecutePolicy) GoString() string {
	return s.String()
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

type GetPipelineResponseBodyExecutePolicyRunOnce struct {
	// example:
	//
	// 1735660800000
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// example:
	//
	// 1735747200000
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
	// example:
	//
	// 1735660800000
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
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
	// example:
	//
	// node-1
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
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
	Dataset *GetPipelineResponseBodySinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// example:
	//
	// Dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodySink) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySink) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySink) GetDataset() *GetPipelineResponseBodySinkDataset {
	return s.Dataset
}

func (s *GetPipelineResponseBodySink) GetType() *string {
	return s.Type
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
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodySinkDataset struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
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
	Logstore *GetPipelineResponseBodySourceLogstore `json:"logstore,omitempty" xml:"logstore,omitempty" type:"Struct"`
	// example:
	//
	// SLS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodySource) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodySource) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodySource) GetLogstore() *GetPipelineResponseBodySourceLogstore {
	return s.Logstore
}

func (s *GetPipelineResponseBodySource) GetType() *string {
	return s.Type
}

func (s *GetPipelineResponseBodySource) SetLogstore(v *GetPipelineResponseBodySourceLogstore) *GetPipelineResponseBodySource {
	s.Logstore = v
	return s
}

func (s *GetPipelineResponseBodySource) SetType(v string) *GetPipelineResponseBodySource {
	s.Type = &v
	return s
}

func (s *GetPipelineResponseBodySource) Validate() error {
	if s.Logstore != nil {
		if err := s.Logstore.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineResponseBodySourceLogstore struct {
	// example:
	//
	// my-sls-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// my-sls-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
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
