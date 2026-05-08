// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdatePipelineRequest
	GetDescription() *string
	SetExecutePolicy(v *UpdatePipelineRequestExecutePolicy) *UpdatePipelineRequest
	GetExecutePolicy() *UpdatePipelineRequestExecutePolicy
	SetPipeline(v *UpdatePipelineRequestPipeline) *UpdatePipelineRequest
	GetPipeline() *UpdatePipelineRequestPipeline
	SetSink(v *UpdatePipelineRequestSink) *UpdatePipelineRequest
	GetSink() *UpdatePipelineRequestSink
	SetSource(v *UpdatePipelineRequestSource) *UpdatePipelineRequest
	GetSource() *UpdatePipelineRequestSource
}

type UpdatePipelineRequest struct {
	// example:
	//
	// test pipeline
	Description   *string                             `json:"description,omitempty" xml:"description,omitempty"`
	ExecutePolicy *UpdatePipelineRequestExecutePolicy `json:"executePolicy,omitempty" xml:"executePolicy,omitempty" type:"Struct"`
	Pipeline      *UpdatePipelineRequestPipeline      `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	Sink          *UpdatePipelineRequestSink          `json:"sink,omitempty" xml:"sink,omitempty" type:"Struct"`
	Source        *UpdatePipelineRequestSource        `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s UpdatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePipelineRequest) GetExecutePolicy() *UpdatePipelineRequestExecutePolicy {
	return s.ExecutePolicy
}

func (s *UpdatePipelineRequest) GetPipeline() *UpdatePipelineRequestPipeline {
	return s.Pipeline
}

func (s *UpdatePipelineRequest) GetSink() *UpdatePipelineRequestSink {
	return s.Sink
}

func (s *UpdatePipelineRequest) GetSource() *UpdatePipelineRequestSource {
	return s.Source
}

func (s *UpdatePipelineRequest) SetDescription(v string) *UpdatePipelineRequest {
	s.Description = &v
	return s
}

func (s *UpdatePipelineRequest) SetExecutePolicy(v *UpdatePipelineRequestExecutePolicy) *UpdatePipelineRequest {
	s.ExecutePolicy = v
	return s
}

func (s *UpdatePipelineRequest) SetPipeline(v *UpdatePipelineRequestPipeline) *UpdatePipelineRequest {
	s.Pipeline = v
	return s
}

func (s *UpdatePipelineRequest) SetSink(v *UpdatePipelineRequestSink) *UpdatePipelineRequest {
	s.Sink = v
	return s
}

func (s *UpdatePipelineRequest) SetSource(v *UpdatePipelineRequestSource) *UpdatePipelineRequest {
	s.Source = v
	return s
}

func (s *UpdatePipelineRequest) Validate() error {
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

type UpdatePipelineRequestExecutePolicy struct {
	// example:
	//
	// runOnce
	Mode      *string                                      `json:"mode,omitempty" xml:"mode,omitempty"`
	RunOnce   *UpdatePipelineRequestExecutePolicyRunOnce   `json:"runOnce,omitempty" xml:"runOnce,omitempty" type:"Struct"`
	Scheduled *UpdatePipelineRequestExecutePolicyScheduled `json:"scheduled,omitempty" xml:"scheduled,omitempty" type:"Struct"`
}

func (s UpdatePipelineRequestExecutePolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestExecutePolicy) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestExecutePolicy) GetMode() *string {
	return s.Mode
}

func (s *UpdatePipelineRequestExecutePolicy) GetRunOnce() *UpdatePipelineRequestExecutePolicyRunOnce {
	return s.RunOnce
}

func (s *UpdatePipelineRequestExecutePolicy) GetScheduled() *UpdatePipelineRequestExecutePolicyScheduled {
	return s.Scheduled
}

func (s *UpdatePipelineRequestExecutePolicy) SetMode(v string) *UpdatePipelineRequestExecutePolicy {
	s.Mode = &v
	return s
}

func (s *UpdatePipelineRequestExecutePolicy) SetRunOnce(v *UpdatePipelineRequestExecutePolicyRunOnce) *UpdatePipelineRequestExecutePolicy {
	s.RunOnce = v
	return s
}

func (s *UpdatePipelineRequestExecutePolicy) SetScheduled(v *UpdatePipelineRequestExecutePolicyScheduled) *UpdatePipelineRequestExecutePolicy {
	s.Scheduled = v
	return s
}

func (s *UpdatePipelineRequestExecutePolicy) Validate() error {
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

type UpdatePipelineRequestExecutePolicyRunOnce struct {
	// example:
	//
	// 1772519013
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// example:
	//
	// 1772605413
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s UpdatePipelineRequestExecutePolicyRunOnce) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestExecutePolicyRunOnce) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestExecutePolicyRunOnce) GetFromTime() *int64 {
	return s.FromTime
}

func (s *UpdatePipelineRequestExecutePolicyRunOnce) GetToTime() *int64 {
	return s.ToTime
}

func (s *UpdatePipelineRequestExecutePolicyRunOnce) SetFromTime(v int64) *UpdatePipelineRequestExecutePolicyRunOnce {
	s.FromTime = &v
	return s
}

func (s *UpdatePipelineRequestExecutePolicyRunOnce) SetToTime(v int64) *UpdatePipelineRequestExecutePolicyRunOnce {
	s.ToTime = &v
	return s
}

func (s *UpdatePipelineRequestExecutePolicyRunOnce) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineRequestExecutePolicyScheduled struct {
	// example:
	//
	// 1772519013
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// example:
	//
	// 86400
	Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
}

func (s UpdatePipelineRequestExecutePolicyScheduled) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestExecutePolicyScheduled) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestExecutePolicyScheduled) GetFromTime() *int64 {
	return s.FromTime
}

func (s *UpdatePipelineRequestExecutePolicyScheduled) GetInterval() *string {
	return s.Interval
}

func (s *UpdatePipelineRequestExecutePolicyScheduled) SetFromTime(v int64) *UpdatePipelineRequestExecutePolicyScheduled {
	s.FromTime = &v
	return s
}

func (s *UpdatePipelineRequestExecutePolicyScheduled) SetInterval(v string) *UpdatePipelineRequestExecutePolicyScheduled {
	s.Interval = &v
	return s
}

func (s *UpdatePipelineRequestExecutePolicyScheduled) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineRequestPipeline struct {
	Nodes []*UpdatePipelineRequestPipelineNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s UpdatePipelineRequestPipeline) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestPipeline) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestPipeline) GetNodes() []*UpdatePipelineRequestPipelineNodes {
	return s.Nodes
}

func (s *UpdatePipelineRequestPipeline) SetNodes(v []*UpdatePipelineRequestPipelineNodes) *UpdatePipelineRequestPipeline {
	s.Nodes = v
	return s
}

func (s *UpdatePipelineRequestPipeline) Validate() error {
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

type UpdatePipelineRequestPipelineNodes struct {
	// example:
	//
	// node_1
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// dedup-fuzzy
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePipelineRequestPipelineNodes) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestPipelineNodes) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestPipelineNodes) GetId() *string {
	return s.Id
}

func (s *UpdatePipelineRequestPipelineNodes) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdatePipelineRequestPipelineNodes) GetType() *string {
	return s.Type
}

func (s *UpdatePipelineRequestPipelineNodes) SetId(v string) *UpdatePipelineRequestPipelineNodes {
	s.Id = &v
	return s
}

func (s *UpdatePipelineRequestPipelineNodes) SetParameters(v map[string]interface{}) *UpdatePipelineRequestPipelineNodes {
	s.Parameters = v
	return s
}

func (s *UpdatePipelineRequestPipelineNodes) SetType(v string) *UpdatePipelineRequestPipelineNodes {
	s.Type = &v
	return s
}

func (s *UpdatePipelineRequestPipelineNodes) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineRequestSink struct {
	Dataset *UpdatePipelineRequestSinkDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// example:
	//
	// dataset
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePipelineRequestSink) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestSink) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestSink) GetDataset() *UpdatePipelineRequestSinkDataset {
	return s.Dataset
}

func (s *UpdatePipelineRequestSink) GetType() *string {
	return s.Type
}

func (s *UpdatePipelineRequestSink) SetDataset(v *UpdatePipelineRequestSinkDataset) *UpdatePipelineRequestSink {
	s.Dataset = v
	return s
}

func (s *UpdatePipelineRequestSink) SetType(v string) *UpdatePipelineRequestSink {
	s.Type = &v
	return s
}

func (s *UpdatePipelineRequestSink) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePipelineRequestSinkDataset struct {
	// example:
	//
	// dataset_1
	Dataset *string `json:"dataset,omitempty" xml:"dataset,omitempty"`
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdatePipelineRequestSinkDataset) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestSinkDataset) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestSinkDataset) GetDataset() *string {
	return s.Dataset
}

func (s *UpdatePipelineRequestSinkDataset) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdatePipelineRequestSinkDataset) SetDataset(v string) *UpdatePipelineRequestSinkDataset {
	s.Dataset = &v
	return s
}

func (s *UpdatePipelineRequestSinkDataset) SetWorkspace(v string) *UpdatePipelineRequestSinkDataset {
	s.Workspace = &v
	return s
}

func (s *UpdatePipelineRequestSinkDataset) Validate() error {
	return dara.Validate(s)
}

type UpdatePipelineRequestSource struct {
	Logstore *UpdatePipelineRequestSourceLogstore `json:"logstore,omitempty" xml:"logstore,omitempty" type:"Struct"`
	// example:
	//
	// logstore
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePipelineRequestSource) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestSource) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestSource) GetLogstore() *UpdatePipelineRequestSourceLogstore {
	return s.Logstore
}

func (s *UpdatePipelineRequestSource) GetType() *string {
	return s.Type
}

func (s *UpdatePipelineRequestSource) SetLogstore(v *UpdatePipelineRequestSourceLogstore) *UpdatePipelineRequestSource {
	s.Logstore = v
	return s
}

func (s *UpdatePipelineRequestSource) SetType(v string) *UpdatePipelineRequestSource {
	s.Type = &v
	return s
}

func (s *UpdatePipelineRequestSource) Validate() error {
	if s.Logstore != nil {
		if err := s.Logstore.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePipelineRequestSourceLogstore struct {
	// example:
	//
	// test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// status:500 and method:GET
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s UpdatePipelineRequestSourceLogstore) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequestSourceLogstore) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequestSourceLogstore) GetLogstore() *string {
	return s.Logstore
}

func (s *UpdatePipelineRequestSourceLogstore) GetQuery() *string {
	return s.Query
}

func (s *UpdatePipelineRequestSourceLogstore) SetLogstore(v string) *UpdatePipelineRequestSourceLogstore {
	s.Logstore = &v
	return s
}

func (s *UpdatePipelineRequestSourceLogstore) SetQuery(v string) *UpdatePipelineRequestSourceLogstore {
	s.Query = &v
	return s
}

func (s *UpdatePipelineRequestSourceLogstore) Validate() error {
	return dara.Validate(s)
}
