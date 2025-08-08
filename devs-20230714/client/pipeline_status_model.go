// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPipelineStatus interface {
	dara.Model
	String() string
	GoString() string
	SetLatestExecError(v *TaskExecError) *PipelineStatus
	GetLatestExecError() *TaskExecError
	SetPhase(v string) *PipelineStatus
	GetPhase() *string
}

type PipelineStatus struct {
	LatestExecError *TaskExecError `json:"latestExecError,omitempty" xml:"latestExecError,omitempty"`
	// example:
	//
	// Success
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
}

func (s PipelineStatus) String() string {
	return dara.Prettify(s)
}

func (s PipelineStatus) GoString() string {
	return s.String()
}

func (s *PipelineStatus) GetLatestExecError() *TaskExecError {
	return s.LatestExecError
}

func (s *PipelineStatus) GetPhase() *string {
	return s.Phase
}

func (s *PipelineStatus) SetLatestExecError(v *TaskExecError) *PipelineStatus {
	s.LatestExecError = v
	return s
}

func (s *PipelineStatus) SetPhase(v string) *PipelineStatus {
	s.Phase = &v
	return s
}

func (s *PipelineStatus) Validate() error {
	return dara.Validate(s)
}
