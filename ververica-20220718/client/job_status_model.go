// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobStatus interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentJobStatus(v string) *JobStatus
	GetCurrentJobStatus() *string
	SetFailure(v *JobFailure) *JobStatus
	GetFailure() *JobFailure
	SetHealthScore(v int32) *JobStatus
	GetHealthScore() *int32
	SetRiskLevel(v string) *JobStatus
	GetRiskLevel() *string
	SetRunning(v *JobStatusRunning) *JobStatus
	GetRunning() *JobStatusRunning
}

type JobStatus struct {
	// example:
	//
	// RUNNING
	CurrentJobStatus *string           `json:"currentJobStatus,omitempty" xml:"currentJobStatus,omitempty"`
	Failure          *JobFailure       `json:"failure,omitempty" xml:"failure,omitempty"`
	HealthScore      *int32            `json:"healthScore,omitempty" xml:"healthScore,omitempty"`
	RiskLevel        *string           `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	Running          *JobStatusRunning `json:"running,omitempty" xml:"running,omitempty"`
}

func (s JobStatus) String() string {
	return dara.Prettify(s)
}

func (s JobStatus) GoString() string {
	return s.String()
}

func (s *JobStatus) GetCurrentJobStatus() *string {
	return s.CurrentJobStatus
}

func (s *JobStatus) GetFailure() *JobFailure {
	return s.Failure
}

func (s *JobStatus) GetHealthScore() *int32 {
	return s.HealthScore
}

func (s *JobStatus) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *JobStatus) GetRunning() *JobStatusRunning {
	return s.Running
}

func (s *JobStatus) SetCurrentJobStatus(v string) *JobStatus {
	s.CurrentJobStatus = &v
	return s
}

func (s *JobStatus) SetFailure(v *JobFailure) *JobStatus {
	s.Failure = v
	return s
}

func (s *JobStatus) SetHealthScore(v int32) *JobStatus {
	s.HealthScore = &v
	return s
}

func (s *JobStatus) SetRiskLevel(v string) *JobStatus {
	s.RiskLevel = &v
	return s
}

func (s *JobStatus) SetRunning(v *JobStatusRunning) *JobStatus {
	s.Running = v
	return s
}

func (s *JobStatus) Validate() error {
	if s.Failure != nil {
		if err := s.Failure.Validate(); err != nil {
			return err
		}
	}
	if s.Running != nil {
		if err := s.Running.Validate(); err != nil {
			return err
		}
	}
	return nil
}
