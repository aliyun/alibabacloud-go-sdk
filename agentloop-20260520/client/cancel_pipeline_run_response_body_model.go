// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFinishTime(v string) *CancelPipelineRunResponseBody
	GetFinishTime() *string
	SetRequestId(v string) *CancelPipelineRunResponseBody
	GetRequestId() *string
	SetRunId(v string) *CancelPipelineRunResponseBody
	GetRunId() *string
	SetStatus(v string) *CancelPipelineRunResponseBody
	GetStatus() *string
}

type CancelPipelineRunResponseBody struct {
	// The time when the cancellation was completed, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:05.000Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the canceled run.
	//
	// example:
	//
	// run-20260101-0001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The status of the run after cancellation. The value is fixed to Cancelled.
	//
	// example:
	//
	// Cancelled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CancelPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPipelineRunResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *CancelPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelPipelineRunResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *CancelPipelineRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelPipelineRunResponseBody) SetFinishTime(v string) *CancelPipelineRunResponseBody {
	s.FinishTime = &v
	return s
}

func (s *CancelPipelineRunResponseBody) SetRequestId(v string) *CancelPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPipelineRunResponseBody) SetRunId(v string) *CancelPipelineRunResponseBody {
	s.RunId = &v
	return s
}

func (s *CancelPipelineRunResponseBody) SetStatus(v string) *CancelPipelineRunResponseBody {
	s.Status = &v
	return s
}

func (s *CancelPipelineRunResponseBody) Validate() error {
	return dara.Validate(s)
}
