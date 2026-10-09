// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromTime(v int64) *PreviewPipelineRequest
	GetFromTime() *int64
	SetPipeline(v *PreviewPipelineRequestPipeline) *PreviewPipelineRequest
	GetPipeline() *PreviewPipelineRequestPipeline
	SetSource(v *PreviewPipelineRequestSource) *PreviewPipelineRequest
	GetSource() *PreviewPipelineRequestSource
	SetToTime(v int64) *PreviewPipelineRequest
	GetToTime() *int64
}

type PreviewPipelineRequest struct {
	// The start time of the preview data window. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1735660800
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The pipeline configuration, including node orchestration.
	//
	// example:
	//
	// {"nodes":[{"id":"select-fields","type":"project","parameters":{"question":"user_query"}}]}
	Pipeline *PreviewPipelineRequestPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// The data source of the pipeline.
	//
	// example:
	//
	// {"type":"logstore","logstore":{"project":"my-sls-project","logstore":"agent-logs"},"inputFields":[{"name":"question","type":"text"}]}
	Source *PreviewPipelineRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The end time of the preview data window. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1735747200
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s PreviewPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequest) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequest) GetFromTime() *int64 {
	return s.FromTime
}

func (s *PreviewPipelineRequest) GetPipeline() *PreviewPipelineRequestPipeline {
	return s.Pipeline
}

func (s *PreviewPipelineRequest) GetSource() *PreviewPipelineRequestSource {
	return s.Source
}

func (s *PreviewPipelineRequest) GetToTime() *int64 {
	return s.ToTime
}

func (s *PreviewPipelineRequest) SetFromTime(v int64) *PreviewPipelineRequest {
	s.FromTime = &v
	return s
}

func (s *PreviewPipelineRequest) SetPipeline(v *PreviewPipelineRequestPipeline) *PreviewPipelineRequest {
	s.Pipeline = v
	return s
}

func (s *PreviewPipelineRequest) SetSource(v *PreviewPipelineRequestSource) *PreviewPipelineRequest {
	s.Source = v
	return s
}

func (s *PreviewPipelineRequest) SetToTime(v int64) *PreviewPipelineRequest {
	s.ToTime = &v
	return s
}

func (s *PreviewPipelineRequest) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
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

type PreviewPipelineRequestPipeline struct {
	// The list of nodes.
	//
	// example:
	//
	// [{"id":"select-fields","type":"project","parameters":{}}]
	Nodes []*PreviewPipelineRequestPipelineNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s PreviewPipelineRequestPipeline) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequestPipeline) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequestPipeline) GetNodes() []*PreviewPipelineRequestPipelineNodes {
	return s.Nodes
}

func (s *PreviewPipelineRequestPipeline) SetNodes(v []*PreviewPipelineRequestPipelineNodes) *PreviewPipelineRequestPipeline {
	s.Nodes = v
	return s
}

func (s *PreviewPipelineRequestPipeline) Validate() error {
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

type PreviewPipelineRequestPipelineNodes struct {
	// The ID of the node.
	//
	// example:
	//
	// node-1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The parameters of the node. The parameters are in key-value format and vary based on the node type.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// transform
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewPipelineRequestPipelineNodes) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequestPipelineNodes) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequestPipelineNodes) GetId() *string {
	return s.Id
}

func (s *PreviewPipelineRequestPipelineNodes) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *PreviewPipelineRequestPipelineNodes) GetType() *string {
	return s.Type
}

func (s *PreviewPipelineRequestPipelineNodes) SetId(v string) *PreviewPipelineRequestPipelineNodes {
	s.Id = &v
	return s
}

func (s *PreviewPipelineRequestPipelineNodes) SetParameters(v map[string]interface{}) *PreviewPipelineRequestPipelineNodes {
	s.Parameters = v
	return s
}

func (s *PreviewPipelineRequestPipelineNodes) SetType(v string) *PreviewPipelineRequestPipelineNodes {
	s.Type = &v
	return s
}

func (s *PreviewPipelineRequestPipelineNodes) Validate() error {
	return dara.Validate(s)
}

type PreviewPipelineRequestSource struct {
	// The dataset datasource config in the current AgentSpace.
	//
	// example:
	//
	// {"dataset":"my-dataset","filter":"status = \\"pending\\""}
	Dataset *PreviewPipelineRequestSourceDataset `json:"dataset,omitempty" xml:"dataset,omitempty" type:"Struct"`
	// The input fields and their data types. This applies to all data source types.
	//
	// example:
	//
	// [{"name":"question","type":"text"}]
	InputFields []*PreviewPipelineRequestSourceInputFields `json:"inputFields,omitempty" xml:"inputFields,omitempty" type:"Repeated"`
	// The Simple Log Service Logstore datasource config.
	//
	// example:
	//
	// {"project":"my-sls-project","logstore":"agent-logs"}
	Logstore *PreviewPipelineRequestSourceLogstore `json:"logstore,omitempty" xml:"logstore,omitempty" type:"Struct"`
	// The configuration of trajectory data. This parameter is optional and takes effect only when the type is set to trace. It retrieves ATIF standard trajectory data from the trajectory cleaning service and extends the data based on features.
	//
	// example:
	//
	// {"enrich":{"enabled":true,"columns":["input","output"]}}
	Trajectory *PreviewPipelineRequestSourceTrajectory `json:"trajectory,omitempty" xml:"trajectory,omitempty" type:"Struct"`
	// The type of the data source. Simple Log Service is currently supported.
	//
	// example:
	//
	// SLS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewPipelineRequestSource) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequestSource) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequestSource) GetDataset() *PreviewPipelineRequestSourceDataset {
	return s.Dataset
}

func (s *PreviewPipelineRequestSource) GetInputFields() []*PreviewPipelineRequestSourceInputFields {
	return s.InputFields
}

func (s *PreviewPipelineRequestSource) GetLogstore() *PreviewPipelineRequestSourceLogstore {
	return s.Logstore
}

func (s *PreviewPipelineRequestSource) GetTrajectory() *PreviewPipelineRequestSourceTrajectory {
	return s.Trajectory
}

func (s *PreviewPipelineRequestSource) GetType() *string {
	return s.Type
}

func (s *PreviewPipelineRequestSource) SetDataset(v *PreviewPipelineRequestSourceDataset) *PreviewPipelineRequestSource {
	s.Dataset = v
	return s
}

func (s *PreviewPipelineRequestSource) SetInputFields(v []*PreviewPipelineRequestSourceInputFields) *PreviewPipelineRequestSource {
	s.InputFields = v
	return s
}

func (s *PreviewPipelineRequestSource) SetLogstore(v *PreviewPipelineRequestSourceLogstore) *PreviewPipelineRequestSource {
	s.Logstore = v
	return s
}

func (s *PreviewPipelineRequestSource) SetTrajectory(v *PreviewPipelineRequestSourceTrajectory) *PreviewPipelineRequestSource {
	s.Trajectory = v
	return s
}

func (s *PreviewPipelineRequestSource) SetType(v string) *PreviewPipelineRequestSource {
	s.Type = &v
	return s
}

func (s *PreviewPipelineRequestSource) Validate() error {
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

type PreviewPipelineRequestSourceDataset struct {
	// The name of the source dataset.
	//
	// example:
	//
	// my-dataset
	Dataset *string `json:"dataset,omitempty" xml:"dataset,omitempty"`
	// The filter condition for the dataset data.
	//
	// example:
	//
	// status = \\"pending\\"
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s PreviewPipelineRequestSourceDataset) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequestSourceDataset) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequestSourceDataset) GetDataset() *string {
	return s.Dataset
}

func (s *PreviewPipelineRequestSourceDataset) GetFilter() *string {
	return s.Filter
}

func (s *PreviewPipelineRequestSourceDataset) SetDataset(v string) *PreviewPipelineRequestSourceDataset {
	s.Dataset = &v
	return s
}

func (s *PreviewPipelineRequestSourceDataset) SetFilter(v string) *PreviewPipelineRequestSourceDataset {
	s.Filter = &v
	return s
}

func (s *PreviewPipelineRequestSourceDataset) Validate() error {
	return dara.Validate(s)
}

type PreviewPipelineRequestSourceInputFields struct {
	// The name of the field.
	//
	// example:
	//
	// question
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the field. Valid values: text, long, double, and json.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreviewPipelineRequestSourceInputFields) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequestSourceInputFields) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequestSourceInputFields) GetName() *string {
	return s.Name
}

func (s *PreviewPipelineRequestSourceInputFields) GetType() *string {
	return s.Type
}

func (s *PreviewPipelineRequestSourceInputFields) SetName(v string) *PreviewPipelineRequestSourceInputFields {
	s.Name = &v
	return s
}

func (s *PreviewPipelineRequestSourceInputFields) SetType(v string) *PreviewPipelineRequestSourceInputFields {
	s.Type = &v
	return s
}

func (s *PreviewPipelineRequestSourceInputFields) Validate() error {
	return dara.Validate(s)
}

type PreviewPipelineRequestSourceLogstore struct {
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// my-sls-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// my-sls-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The filtered query statement (Simple Log Service query and analysis syntax).
	//
	// example:
	//
	// 	- | SELECT *
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s PreviewPipelineRequestSourceLogstore) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequestSourceLogstore) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequestSourceLogstore) GetLogstore() *string {
	return s.Logstore
}

func (s *PreviewPipelineRequestSourceLogstore) GetProject() *string {
	return s.Project
}

func (s *PreviewPipelineRequestSourceLogstore) GetQuery() *string {
	return s.Query
}

func (s *PreviewPipelineRequestSourceLogstore) SetLogstore(v string) *PreviewPipelineRequestSourceLogstore {
	s.Logstore = &v
	return s
}

func (s *PreviewPipelineRequestSourceLogstore) SetProject(v string) *PreviewPipelineRequestSourceLogstore {
	s.Project = &v
	return s
}

func (s *PreviewPipelineRequestSourceLogstore) SetQuery(v string) *PreviewPipelineRequestSourceLogstore {
	s.Query = &v
	return s
}

func (s *PreviewPipelineRequestSourceLogstore) Validate() error {
	return dara.Validate(s)
}

type PreviewPipelineRequestSourceTrajectory struct {
	// Trajectory enrichment: mounts trajectory data into the cleaning results based on the trace_id. When writing data to a dataset, the data is stored in the fixed agent_trajectory column, and the column value is the JSON content of the trajectory.
	//
	// example:
	//
	// {"enabled":true,"columns":["input","output"]}
	Enrich *PreviewPipelineRequestSourceTrajectoryEnrich `json:"enrich,omitempty" xml:"enrich,omitempty" type:"Struct"`
}

func (s PreviewPipelineRequestSourceTrajectory) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequestSourceTrajectory) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequestSourceTrajectory) GetEnrich() *PreviewPipelineRequestSourceTrajectoryEnrich {
	return s.Enrich
}

func (s *PreviewPipelineRequestSourceTrajectory) SetEnrich(v *PreviewPipelineRequestSourceTrajectoryEnrich) *PreviewPipelineRequestSourceTrajectory {
	s.Enrich = v
	return s
}

func (s *PreviewPipelineRequestSourceTrajectory) Validate() error {
	if s.Enrich != nil {
		if err := s.Enrich.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewPipelineRequestSourceTrajectoryEnrich struct {
	// The list of enrichment columns. This parameter is retained for compatibility. The current implementation outputs only the fixed agent_trajectory column, and this parameter no longer affects the output.
	//
	// example:
	//
	// ["input","output","session_id"]
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// Specifies whether to enable trajectory enrichment.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s PreviewPipelineRequestSourceTrajectoryEnrich) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineRequestSourceTrajectoryEnrich) GoString() string {
	return s.String()
}

func (s *PreviewPipelineRequestSourceTrajectoryEnrich) GetColumns() []*string {
	return s.Columns
}

func (s *PreviewPipelineRequestSourceTrajectoryEnrich) GetEnabled() *bool {
	return s.Enabled
}

func (s *PreviewPipelineRequestSourceTrajectoryEnrich) SetColumns(v []*string) *PreviewPipelineRequestSourceTrajectoryEnrich {
	s.Columns = v
	return s
}

func (s *PreviewPipelineRequestSourceTrajectoryEnrich) SetEnabled(v bool) *PreviewPipelineRequestSourceTrajectoryEnrich {
	s.Enabled = &v
	return s
}

func (s *PreviewPipelineRequestSourceTrajectoryEnrich) Validate() error {
	return dara.Validate(s)
}
