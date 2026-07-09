// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpaceName(v string) *TerminatePipelineResponseBody
	GetAgentSpaceName() *string
	SetPipelineName(v string) *TerminatePipelineResponseBody
	GetPipelineName() *string
	SetRequestId(v string) *TerminatePipelineResponseBody
	GetRequestId() *string
	SetScheduleStatus(v string) *TerminatePipelineResponseBody
	GetScheduleStatus() *string
	SetTerminateTime(v string) *TerminatePipelineResponseBody
	GetTerminateTime() *string
	SetTerminatedReason(v string) *TerminatePipelineResponseBody
	GetTerminatedReason() *string
}

type TerminatePipelineResponseBody struct {
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
	// Terminated
	ScheduleStatus *string `json:"scheduleStatus,omitempty" xml:"scheduleStatus,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:00.000Z
	TerminateTime *string `json:"terminateTime,omitempty" xml:"terminateTime,omitempty"`
	// example:
	//
	// project deprecated
	TerminatedReason *string `json:"terminatedReason,omitempty" xml:"terminatedReason,omitempty"`
}

func (s TerminatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *TerminatePipelineResponseBody) GetAgentSpaceName() *string {
	return s.AgentSpaceName
}

func (s *TerminatePipelineResponseBody) GetPipelineName() *string {
	return s.PipelineName
}

func (s *TerminatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminatePipelineResponseBody) GetScheduleStatus() *string {
	return s.ScheduleStatus
}

func (s *TerminatePipelineResponseBody) GetTerminateTime() *string {
	return s.TerminateTime
}

func (s *TerminatePipelineResponseBody) GetTerminatedReason() *string {
	return s.TerminatedReason
}

func (s *TerminatePipelineResponseBody) SetAgentSpaceName(v string) *TerminatePipelineResponseBody {
	s.AgentSpaceName = &v
	return s
}

func (s *TerminatePipelineResponseBody) SetPipelineName(v string) *TerminatePipelineResponseBody {
	s.PipelineName = &v
	return s
}

func (s *TerminatePipelineResponseBody) SetRequestId(v string) *TerminatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminatePipelineResponseBody) SetScheduleStatus(v string) *TerminatePipelineResponseBody {
	s.ScheduleStatus = &v
	return s
}

func (s *TerminatePipelineResponseBody) SetTerminateTime(v string) *TerminatePipelineResponseBody {
	s.TerminateTime = &v
	return s
}

func (s *TerminatePipelineResponseBody) SetTerminatedReason(v string) *TerminatePipelineResponseBody {
	s.TerminatedReason = &v
	return s
}

func (s *TerminatePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
