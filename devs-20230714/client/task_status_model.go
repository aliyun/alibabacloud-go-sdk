// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskStatus interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionDetails(v []*string) *TaskStatus
	GetExecutionDetails() []*string
	SetInvocations(v []*TaskInvocation) *TaskStatus
	GetInvocations() []*TaskInvocation
	SetLatestExecError(v *TaskExecError) *TaskStatus
	GetLatestExecError() *TaskExecError
	SetPhase(v string) *TaskStatus
	GetPhase() *string
	SetStatusGeneration(v int64) *TaskStatus
	GetStatusGeneration() *int64
}

type TaskStatus struct {
	ExecutionDetails []*string         `json:"executionDetails,omitempty" xml:"executionDetails,omitempty" type:"Repeated"`
	Invocations      []*TaskInvocation `json:"invocations,omitempty" xml:"invocations,omitempty" type:"Repeated"`
	LatestExecError  *TaskExecError    `json:"latestExecError,omitempty" xml:"latestExecError,omitempty"`
	// example:
	//
	// Success
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// 123
	StatusGeneration *int64 `json:"statusGeneration,omitempty" xml:"statusGeneration,omitempty"`
}

func (s TaskStatus) String() string {
	return dara.Prettify(s)
}

func (s TaskStatus) GoString() string {
	return s.String()
}

func (s *TaskStatus) GetExecutionDetails() []*string {
	return s.ExecutionDetails
}

func (s *TaskStatus) GetInvocations() []*TaskInvocation {
	return s.Invocations
}

func (s *TaskStatus) GetLatestExecError() *TaskExecError {
	return s.LatestExecError
}

func (s *TaskStatus) GetPhase() *string {
	return s.Phase
}

func (s *TaskStatus) GetStatusGeneration() *int64 {
	return s.StatusGeneration
}

func (s *TaskStatus) SetExecutionDetails(v []*string) *TaskStatus {
	s.ExecutionDetails = v
	return s
}

func (s *TaskStatus) SetInvocations(v []*TaskInvocation) *TaskStatus {
	s.Invocations = v
	return s
}

func (s *TaskStatus) SetLatestExecError(v *TaskExecError) *TaskStatus {
	s.LatestExecError = v
	return s
}

func (s *TaskStatus) SetPhase(v string) *TaskStatus {
	s.Phase = &v
	return s
}

func (s *TaskStatus) SetStatusGeneration(v int64) *TaskStatus {
	s.StatusGeneration = &v
	return s
}

func (s *TaskStatus) Validate() error {
	return dara.Validate(s)
}
