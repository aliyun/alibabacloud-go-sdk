// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpaceName(v string) *ResumePipelineResponseBody
	GetAgentSpaceName() *string
	SetCommittedWatermark(v int64) *ResumePipelineResponseBody
	GetCommittedWatermark() *int64
	SetNextTriggerTime(v int64) *ResumePipelineResponseBody
	GetNextTriggerTime() *int64
	SetPipelineName(v string) *ResumePipelineResponseBody
	GetPipelineName() *string
	SetRequestId(v string) *ResumePipelineResponseBody
	GetRequestId() *string
	SetScheduleStatus(v string) *ResumePipelineResponseBody
	GetScheduleStatus() *string
}

type ResumePipelineResponseBody struct {
	// The name of the AgentSpace where the pipeline is located.
	//
	// example:
	//
	// my-agent-space
	AgentSpaceName *string `json:"agentSpaceName,omitempty" xml:"agentSpaceName,omitempty"`
	// The committed watermark, in UNIX seconds.
	//
	// example:
	//
	// 1735660800
	CommittedWatermark *int64 `json:"committedWatermark,omitempty" xml:"committedWatermark,omitempty"`
	// The next scheduling trigger time, in UNIX seconds.
	//
	// example:
	//
	// 1735664400
	NextTriggerTime *int64 `json:"nextTriggerTime,omitempty" xml:"nextTriggerTime,omitempty"`
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
	// The scheduling status. The value is fixed to Active.
	//
	// example:
	//
	// Active
	ScheduleStatus *string `json:"scheduleStatus,omitempty" xml:"scheduleStatus,omitempty"`
}

func (s ResumePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *ResumePipelineResponseBody) GetAgentSpaceName() *string {
	return s.AgentSpaceName
}

func (s *ResumePipelineResponseBody) GetCommittedWatermark() *int64 {
	return s.CommittedWatermark
}

func (s *ResumePipelineResponseBody) GetNextTriggerTime() *int64 {
	return s.NextTriggerTime
}

func (s *ResumePipelineResponseBody) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ResumePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumePipelineResponseBody) GetScheduleStatus() *string {
	return s.ScheduleStatus
}

func (s *ResumePipelineResponseBody) SetAgentSpaceName(v string) *ResumePipelineResponseBody {
	s.AgentSpaceName = &v
	return s
}

func (s *ResumePipelineResponseBody) SetCommittedWatermark(v int64) *ResumePipelineResponseBody {
	s.CommittedWatermark = &v
	return s
}

func (s *ResumePipelineResponseBody) SetNextTriggerTime(v int64) *ResumePipelineResponseBody {
	s.NextTriggerTime = &v
	return s
}

func (s *ResumePipelineResponseBody) SetPipelineName(v string) *ResumePipelineResponseBody {
	s.PipelineName = &v
	return s
}

func (s *ResumePipelineResponseBody) SetRequestId(v string) *ResumePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumePipelineResponseBody) SetScheduleStatus(v string) *ResumePipelineResponseBody {
	s.ScheduleStatus = &v
	return s
}

func (s *ResumePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
