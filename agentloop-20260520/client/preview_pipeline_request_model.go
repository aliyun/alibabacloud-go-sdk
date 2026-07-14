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
	// The start time of the preview data window, in UNIX seconds.
	//
	// example:
	//
	// 1735660800
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The pipeline configuration, which defines the node orchestration.
	Pipeline *PreviewPipelineRequestPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// The pipeline data source.
	Source *PreviewPipelineRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The end time of the preview data window, in UNIX seconds.
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
	// The node ID.
	//
	// example:
	//
	// node-1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The node parameters in key-value structure. The parameters vary depending on the node type.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The node type.
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
	// The SLS Logstore datasource config.
	Logstore *PreviewPipelineRequestSourceLogstore `json:"logstore,omitempty" xml:"logstore,omitempty" type:"Struct"`
	// The data source type. Currently, only Simple Log Service (SLS) is supported.
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

func (s *PreviewPipelineRequestSource) GetLogstore() *PreviewPipelineRequestSourceLogstore {
	return s.Logstore
}

func (s *PreviewPipelineRequestSource) GetType() *string {
	return s.Type
}

func (s *PreviewPipelineRequestSource) SetLogstore(v *PreviewPipelineRequestSourceLogstore) *PreviewPipelineRequestSource {
	s.Logstore = v
	return s
}

func (s *PreviewPipelineRequestSource) SetType(v string) *PreviewPipelineRequestSource {
	s.Type = &v
	return s
}

func (s *PreviewPipelineRequestSource) Validate() error {
	if s.Logstore != nil {
		if err := s.Logstore.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewPipelineRequestSourceLogstore struct {
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
