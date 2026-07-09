// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpaceName(v string) *PausePipelineResponseBody
	GetAgentSpaceName() *string
	SetPauseTime(v string) *PausePipelineResponseBody
	GetPauseTime() *string
	SetPausedReason(v string) *PausePipelineResponseBody
	GetPausedReason() *string
	SetPipelineName(v string) *PausePipelineResponseBody
	GetPipelineName() *string
	SetRequestId(v string) *PausePipelineResponseBody
	GetRequestId() *string
	SetScheduleStatus(v string) *PausePipelineResponseBody
	GetScheduleStatus() *string
}

type PausePipelineResponseBody struct {
	// The name of the AgentSpace where the pipeline is located.
	//
	// example:
	//
	// my-agent-space
	AgentSpaceName *string `json:"agentSpaceName,omitempty" xml:"agentSpaceName,omitempty"`
	// The time when the pipeline was paused, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:00.000Z
	PauseTime *string `json:"pauseTime,omitempty" xml:"pauseTime,omitempty"`
	// The reason for pausing the pipeline.
	//
	// example:
	//
	// manual maintenance
	PausedReason *string `json:"pausedReason,omitempty" xml:"pausedReason,omitempty"`
	// The name of the pipeline.
	//
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scheduling status. The value is fixed as Paused.
	//
	// example:
	//
	// Paused
	ScheduleStatus *string `json:"scheduleStatus,omitempty" xml:"scheduleStatus,omitempty"`
}

func (s PausePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PausePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *PausePipelineResponseBody) GetAgentSpaceName() *string {
	return s.AgentSpaceName
}

func (s *PausePipelineResponseBody) GetPauseTime() *string {
	return s.PauseTime
}

func (s *PausePipelineResponseBody) GetPausedReason() *string {
	return s.PausedReason
}

func (s *PausePipelineResponseBody) GetPipelineName() *string {
	return s.PipelineName
}

func (s *PausePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PausePipelineResponseBody) GetScheduleStatus() *string {
	return s.ScheduleStatus
}

func (s *PausePipelineResponseBody) SetAgentSpaceName(v string) *PausePipelineResponseBody {
	s.AgentSpaceName = &v
	return s
}

func (s *PausePipelineResponseBody) SetPauseTime(v string) *PausePipelineResponseBody {
	s.PauseTime = &v
	return s
}

func (s *PausePipelineResponseBody) SetPausedReason(v string) *PausePipelineResponseBody {
	s.PausedReason = &v
	return s
}

func (s *PausePipelineResponseBody) SetPipelineName(v string) *PausePipelineResponseBody {
	s.PipelineName = &v
	return s
}

func (s *PausePipelineResponseBody) SetRequestId(v string) *PausePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *PausePipelineResponseBody) SetScheduleStatus(v string) *PausePipelineResponseBody {
	s.ScheduleStatus = &v
	return s
}

func (s *PausePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
