// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpaceName(v string) *RunPipelineResponseBody
	GetAgentSpaceName() *string
	SetPipelineName(v string) *RunPipelineResponseBody
	GetPipelineName() *string
	SetRequestId(v string) *RunPipelineResponseBody
	GetRequestId() *string
	SetRunId(v string) *RunPipelineResponseBody
	GetRunId() *string
	SetStatus(v string) *RunPipelineResponseBody
	GetStatus() *string
}

type RunPipelineResponseBody struct {
	// example:
	//
	// my-agent-space
	AgentSpaceName *string `json:"agentSpaceName,omitempty" xml:"agentSpaceName,omitempty"`
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// run-20260101-0001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RunPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *RunPipelineResponseBody) GetAgentSpaceName() *string {
	return s.AgentSpaceName
}

func (s *RunPipelineResponseBody) GetPipelineName() *string {
	return s.PipelineName
}

func (s *RunPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunPipelineResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *RunPipelineResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RunPipelineResponseBody) SetAgentSpaceName(v string) *RunPipelineResponseBody {
	s.AgentSpaceName = &v
	return s
}

func (s *RunPipelineResponseBody) SetPipelineName(v string) *RunPipelineResponseBody {
	s.PipelineName = &v
	return s
}

func (s *RunPipelineResponseBody) SetRequestId(v string) *RunPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPipelineResponseBody) SetRunId(v string) *RunPipelineResponseBody {
	s.RunId = &v
	return s
}

func (s *RunPipelineResponseBody) SetStatus(v string) *RunPipelineResponseBody {
	s.Status = &v
	return s
}

func (s *RunPipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
