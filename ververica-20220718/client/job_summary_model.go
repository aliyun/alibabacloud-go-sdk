// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobSummary interface {
	dara.Model
	String() string
	GoString() string
	SetCancelled(v int32) *JobSummary
	GetCancelled() *int32
	SetCancelling(v int32) *JobSummary
	GetCancelling() *int32
	SetFailed(v int32) *JobSummary
	GetFailed() *int32
	SetFinished(v int32) *JobSummary
	GetFinished() *int32
	SetRunning(v int32) *JobSummary
	GetRunning() *int32
	SetStarting(v int32) *JobSummary
	GetStarting() *int32
}

type JobSummary struct {
	// example:
	//
	// 1
	Cancelled *int32 `json:"cancelled,omitempty" xml:"cancelled,omitempty"`
	// example:
	//
	// 1
	Cancelling *int32 `json:"cancelling,omitempty" xml:"cancelling,omitempty"`
	// example:
	//
	// 1
	Failed *int32 `json:"failed,omitempty" xml:"failed,omitempty"`
	// example:
	//
	// 1
	Finished *int32 `json:"finished,omitempty" xml:"finished,omitempty"`
	// example:
	//
	// 1
	Running *int32 `json:"running,omitempty" xml:"running,omitempty"`
	// example:
	//
	// 1
	Starting *int32 `json:"starting,omitempty" xml:"starting,omitempty"`
}

func (s JobSummary) String() string {
	return dara.Prettify(s)
}

func (s JobSummary) GoString() string {
	return s.String()
}

func (s *JobSummary) GetCancelled() *int32 {
	return s.Cancelled
}

func (s *JobSummary) GetCancelling() *int32 {
	return s.Cancelling
}

func (s *JobSummary) GetFailed() *int32 {
	return s.Failed
}

func (s *JobSummary) GetFinished() *int32 {
	return s.Finished
}

func (s *JobSummary) GetRunning() *int32 {
	return s.Running
}

func (s *JobSummary) GetStarting() *int32 {
	return s.Starting
}

func (s *JobSummary) SetCancelled(v int32) *JobSummary {
	s.Cancelled = &v
	return s
}

func (s *JobSummary) SetCancelling(v int32) *JobSummary {
	s.Cancelling = &v
	return s
}

func (s *JobSummary) SetFailed(v int32) *JobSummary {
	s.Failed = &v
	return s
}

func (s *JobSummary) SetFinished(v int32) *JobSummary {
	s.Finished = &v
	return s
}

func (s *JobSummary) SetRunning(v int32) *JobSummary {
	s.Running = &v
	return s
}

func (s *JobSummary) SetStarting(v int32) *JobSummary {
	s.Starting = &v
	return s
}

func (s *JobSummary) Validate() error {
	return dara.Validate(s)
}
