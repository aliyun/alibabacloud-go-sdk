// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecuteType(v string) *CreateJobRequest
	GetExecuteType() *string
	SetExperimentId(v string) *CreateJobRequest
	GetExperimentId() *string
	SetNodeId(v string) *CreateJobRequest
	GetNodeId() *string
	SetOptions(v string) *CreateJobRequest
	GetOptions() *string
	SetPipelineDraftId(v string) *CreateJobRequest
	GetPipelineDraftId() *string
}

type CreateJobRequest struct {
	// example:
	//
	// EXECUTE_ALL
	ExecuteType *string `json:"ExecuteType,omitempty" xml:"ExecuteType,omitempty"`
	// Deprecated
	//
	// example:
	//
	// draft-o1p0k444nlq3cd50zz
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// id-2d88-1608982098027-91558
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// {"mlflow":{"experimentId":"exp-1"}}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// draft-o1p0k444nlq3cd50zz
	PipelineDraftId *string `json:"PipelineDraftId,omitempty" xml:"PipelineDraftId,omitempty"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetExecuteType() *string {
	return s.ExecuteType
}

func (s *CreateJobRequest) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *CreateJobRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateJobRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateJobRequest) GetPipelineDraftId() *string {
	return s.PipelineDraftId
}

func (s *CreateJobRequest) SetExecuteType(v string) *CreateJobRequest {
	s.ExecuteType = &v
	return s
}

func (s *CreateJobRequest) SetExperimentId(v string) *CreateJobRequest {
	s.ExperimentId = &v
	return s
}

func (s *CreateJobRequest) SetNodeId(v string) *CreateJobRequest {
	s.NodeId = &v
	return s
}

func (s *CreateJobRequest) SetOptions(v string) *CreateJobRequest {
	s.Options = &v
	return s
}

func (s *CreateJobRequest) SetPipelineDraftId(v string) *CreateJobRequest {
	s.PipelineDraftId = &v
	return s
}

func (s *CreateJobRequest) Validate() error {
	return dara.Validate(s)
}
