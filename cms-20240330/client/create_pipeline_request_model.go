// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePipelineRequest
	GetDescription() *string
	SetExecutePolicy(v *CreatePipelineRequestExecutePolicy) *CreatePipelineRequest
	GetExecutePolicy() *CreatePipelineRequestExecutePolicy
	SetPipeline(v *CreatePipelineRequestPipeline) *CreatePipelineRequest
	GetPipeline() *CreatePipelineRequestPipeline
	SetPipelineName(v string) *CreatePipelineRequest
	GetPipelineName() *string
	SetSink(v *CreatePipelineRequestSink) *CreatePipelineRequest
	GetSink() *CreatePipelineRequestSink
	SetSource(v *CreatePipelineRequestSource) *CreatePipelineRequest
	GetSource() *CreatePipelineRequestSource
}

type CreatePipelineRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy.
	ExecutePolicy *CreatePipelineRequestExecutePolicy `json:"executePolicy,omitempty" xml:"executePolicy,omitempty" type:"Struct"`
	// The pipeline configuration.
	Pipeline *CreatePipelineRequestPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// The pipeline name.
	//
	// example:
	//
	// pipeline-name-1
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The data sink for the processed output.
	Sink *CreatePipelineRequestSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Struct"`
	// The data source.
	Source *CreatePipelineRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s CreatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePipelineRequest) GetExecutePolicy() *CreatePipelineRequestExecutePolicy {
	return s.ExecutePolicy
}

func (s *CreatePipelineRequest) GetPipeline() *CreatePipelineRequestPipeline {
	return s.Pipeline
}

func (s *CreatePipelineRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *CreatePipelineRequest) GetSink() *CreatePipelineRequestSink {
	return s.Sink
}

func (s *CreatePipelineRequest) GetSource() *CreatePipelineRequestSource {
	return s.Source
}

func (s *CreatePipelineRequest) SetDescription(v string) *CreatePipelineRequest {
	s.Description = &v
	return s
}

func (s *CreatePipelineRequest) SetExecutePolicy(v *CreatePipelineRequestExecutePolicy) *CreatePipelineRequest {
	s.ExecutePolicy = v
	return s
}

func (s *CreatePipelineRequest) SetPipeline(v *CreatePipelineRequestPipeline) *CreatePipelineRequest {
	s.Pipeline = v
	return s
}

func (s *CreatePipelineRequest) SetPipelineName(v string) *CreatePipelineRequest {
	s.PipelineName = &v
	return s
}

func (s *CreatePipelineRequest) SetSink(v *CreatePipelineRequestSink) *CreatePipelineRequest {
	s.Sink = v
	return s
}

func (s *CreatePipelineRequest) SetSource(v *CreatePipelineRequestSource) *CreatePipelineRequest {
	s.Source = v
	return s
}

func (s *CreatePipelineRequest) Validate() error {
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

type CreatePipelineRequestExecutePolicy struct {
	// The execution mode. Set to `runOnce` for a single execution, or `scheduled` for a recurring execution.
	//
	// example:
	//
	// runOnce
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The configuration for a one-time execution. This parameter is required when `executePolicy.mode` is set to `runOnce`.
	RunOnce *CreatePipelineRequestExecutePolicyRunOnce `json:"runOnce,omitempty" xml:"runOnce,omitempty" type:"Struct"`
	// The configuration for a scheduled execution. This parameter is required when `executePolicy.mode` is set to `scheduled`.
	Scheduled *CreatePipelineRequestExecutePolicyScheduled `json:"scheduled,omitempty" xml:"scheduled,omitempty" type:"Struct"`
}

func (s CreatePipelineRequestExecutePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestExecutePolicy) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestExecutePolicy) GetMode() *string {
	return s.Mode
}

func (s *CreatePipelineRequestExecutePolicy) GetRunOnce() *CreatePipelineRequestExecutePolicyRunOnce {
	return s.RunOnce
}

func (s *CreatePipelineRequestExecutePolicy) GetScheduled() *CreatePipelineRequestExecutePolicyScheduled {
	return s.Scheduled
}

func (s *CreatePipelineRequestExecutePolicy) SetMode(v string) *CreatePipelineRequestExecutePolicy {
	s.Mode = &v
	return s
}

func (s *CreatePipelineRequestExecutePolicy) SetRunOnce(v *CreatePipelineRequestExecutePolicyRunOnce) *CreatePipelineRequestExecutePolicy {
	s.RunOnce = v
	return s
}

func (s *CreatePipelineRequestExecutePolicy) SetScheduled(v *CreatePipelineRequestExecutePolicyScheduled) *CreatePipelineRequestExecutePolicy {
	s.Scheduled = v
	return s
}

func (s *CreatePipelineRequestExecutePolicy) Validate() error {
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

type CreatePipelineRequestExecutePolicyRunOnce struct {
	// The start timestamp.
	//
	// example:
	//
	// 1772519013
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The end timestamp.
	//
	// example:
	//
	// 1772519013
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s CreatePipelineRequestExecutePolicyRunOnce) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestExecutePolicyRunOnce) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestExecutePolicyRunOnce) GetFromTime() *int64 {
	return s.FromTime
}

func (s *CreatePipelineRequestExecutePolicyRunOnce) GetToTime() *int64 {
	return s.ToTime
}

func (s *CreatePipelineRequestExecutePolicyRunOnce) SetFromTime(v int64) *CreatePipelineRequestExecutePolicyRunOnce {
	s.FromTime = &v
	return s
}

func (s *CreatePipelineRequestExecutePolicyRunOnce) SetToTime(v int64) *CreatePipelineRequestExecutePolicyRunOnce {
	s.ToTime = &v
	return s
}

func (s *CreatePipelineRequestExecutePolicyRunOnce) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineRequestExecutePolicyScheduled struct {
	// The start timestamp.
	//
	// example:
	//
	// 1772519013
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The execution interval in seconds.
	//
	// example:
	//
	// 86400
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
}

func (s CreatePipelineRequestExecutePolicyScheduled) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestExecutePolicyScheduled) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestExecutePolicyScheduled) GetFromTime() *int64 {
	return s.FromTime
}

func (s *CreatePipelineRequestExecutePolicyScheduled) GetInterval() *string {
	return s.Interval
}

func (s *CreatePipelineRequestExecutePolicyScheduled) SetFromTime(v int64) *CreatePipelineRequestExecutePolicyScheduled {
	s.FromTime = &v
	return s
}

func (s *CreatePipelineRequestExecutePolicyScheduled) SetInterval(v string) *CreatePipelineRequestExecutePolicyScheduled {
	s.Interval = &v
	return s
}

func (s *CreatePipelineRequestExecutePolicyScheduled) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineRequestPipeline struct {
	// The pipeline nodes.
	Nodes []*CreatePipelineRequestPipelineNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s CreatePipelineRequestPipeline) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestPipeline) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestPipeline) GetNodes() []*CreatePipelineRequestPipelineNodes {
	return s.Nodes
}

func (s *CreatePipelineRequestPipeline) SetNodes(v []*CreatePipelineRequestPipelineNodes) *CreatePipelineRequestPipeline {
	s.Nodes = v
	return s
}

func (s *CreatePipelineRequestPipeline) Validate() error {
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

type CreatePipelineRequestPipelineNodes struct {
	// The node ID.
	//
	// example:
	//
	// node_1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The node parameters.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The node type.
	//
	// example:
	//
	// dedup-fuzzy
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreatePipelineRequestPipelineNodes) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestPipelineNodes) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestPipelineNodes) GetId() *string {
	return s.Id
}

func (s *CreatePipelineRequestPipelineNodes) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreatePipelineRequestPipelineNodes) GetType() *string {
	return s.Type
}

func (s *CreatePipelineRequestPipelineNodes) SetId(v string) *CreatePipelineRequestPipelineNodes {
	s.Id = &v
	return s
}

func (s *CreatePipelineRequestPipelineNodes) SetParameters(v map[string]interface{}) *CreatePipelineRequestPipelineNodes {
	s.Parameters = v
	return s
}

func (s *CreatePipelineRequestPipelineNodes) SetType(v string) *CreatePipelineRequestPipelineNodes {
	s.Type = &v
	return s
}

func (s *CreatePipelineRequestPipelineNodes) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineRequestSink struct {
	// The destination dataset configuration. This parameter is required when `sink.type` is set to `dataset`.
	Dataset *CreatePipelineRequestSinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The sink type.
	//
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreatePipelineRequestSink) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestSink) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestSink) GetDataset() *CreatePipelineRequestSinkDataset {
	return s.Dataset
}

func (s *CreatePipelineRequestSink) GetType() *string {
	return s.Type
}

func (s *CreatePipelineRequestSink) SetDataset(v *CreatePipelineRequestSinkDataset) *CreatePipelineRequestSink {
	s.Dataset = v
	return s
}

func (s *CreatePipelineRequestSink) SetType(v string) *CreatePipelineRequestSink {
	s.Type = &v
	return s
}

func (s *CreatePipelineRequestSink) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineRequestSinkDataset struct {
	// The dataset name.
	//
	// example:
	//
	// dataset_1
	Dataset *string `json:"dataset,omitempty" xml:"dataset,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreatePipelineRequestSinkDataset) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestSinkDataset) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestSinkDataset) GetDataset() *string {
	return s.Dataset
}

func (s *CreatePipelineRequestSinkDataset) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreatePipelineRequestSinkDataset) SetDataset(v string) *CreatePipelineRequestSinkDataset {
	s.Dataset = &v
	return s
}

func (s *CreatePipelineRequestSinkDataset) SetWorkspace(v string) *CreatePipelineRequestSinkDataset {
	s.Workspace = &v
	return s
}

func (s *CreatePipelineRequestSinkDataset) Validate() error {
	return dara.Validate(s)
}

type CreatePipelineRequestSource struct {
	// The Log Service Logstore configuration. This parameter is required when `source.type` is set to `logstore`.
	Logstore *CreatePipelineRequestSourceLogstore `json:"logstore,omitempty" xml:"logstore,omitempty" type:"Struct"`
	// The data source type.
	//
	// example:
	//
	// logstore
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreatePipelineRequestSource) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestSource) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestSource) GetLogstore() *CreatePipelineRequestSourceLogstore {
	return s.Logstore
}

func (s *CreatePipelineRequestSource) GetType() *string {
	return s.Type
}

func (s *CreatePipelineRequestSource) SetLogstore(v *CreatePipelineRequestSourceLogstore) *CreatePipelineRequestSource {
	s.Logstore = v
	return s
}

func (s *CreatePipelineRequestSource) SetType(v string) *CreatePipelineRequestSource {
	s.Type = &v
	return s
}

func (s *CreatePipelineRequestSource) Validate() error {
	if s.Logstore != nil {
		if err := s.Logstore.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineRequestSourceLogstore struct {
	// The Logstore name.
	//
	// example:
	//
	// test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The Log Service Project name.
	//
	// example:
	//
	// test-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The query statement to filter logs.
	//
	// example:
	//
	// status:500 and method:GET
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s CreatePipelineRequestSourceLogstore) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequestSourceLogstore) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequestSourceLogstore) GetLogstore() *string {
	return s.Logstore
}

func (s *CreatePipelineRequestSourceLogstore) GetProject() *string {
	return s.Project
}

func (s *CreatePipelineRequestSourceLogstore) GetQuery() *string {
	return s.Query
}

func (s *CreatePipelineRequestSourceLogstore) SetLogstore(v string) *CreatePipelineRequestSourceLogstore {
	s.Logstore = &v
	return s
}

func (s *CreatePipelineRequestSourceLogstore) SetProject(v string) *CreatePipelineRequestSourceLogstore {
	s.Project = &v
	return s
}

func (s *CreatePipelineRequestSourceLogstore) SetQuery(v string) *CreatePipelineRequestSourceLogstore {
	s.Query = &v
	return s
}

func (s *CreatePipelineRequestSourceLogstore) Validate() error {
	return dara.Validate(s)
}
