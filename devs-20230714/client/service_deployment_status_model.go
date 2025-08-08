// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceDeploymentStatus interface {
	dara.Model
	String() string
	GoString() string
	SetFinishedTime(v string) *ServiceDeploymentStatus
	GetFinishedTime() *string
	SetPhase(v string) *ServiceDeploymentStatus
	GetPhase() *string
	SetPipelineName(v string) *ServiceDeploymentStatus
	GetPipelineName() *string
	SetStartTime(v string) *ServiceDeploymentStatus
	GetStartTime() *string
	SetTaskName(v string) *ServiceDeploymentStatus
	GetTaskName() *string
}

type ServiceDeploymentStatus struct {
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	// example:
	//
	// Running
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	StartTime    *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// my-task
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s ServiceDeploymentStatus) String() string {
	return dara.Prettify(s)
}

func (s ServiceDeploymentStatus) GoString() string {
	return s.String()
}

func (s *ServiceDeploymentStatus) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *ServiceDeploymentStatus) GetPhase() *string {
	return s.Phase
}

func (s *ServiceDeploymentStatus) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ServiceDeploymentStatus) GetStartTime() *string {
	return s.StartTime
}

func (s *ServiceDeploymentStatus) GetTaskName() *string {
	return s.TaskName
}

func (s *ServiceDeploymentStatus) SetFinishedTime(v string) *ServiceDeploymentStatus {
	s.FinishedTime = &v
	return s
}

func (s *ServiceDeploymentStatus) SetPhase(v string) *ServiceDeploymentStatus {
	s.Phase = &v
	return s
}

func (s *ServiceDeploymentStatus) SetPipelineName(v string) *ServiceDeploymentStatus {
	s.PipelineName = &v
	return s
}

func (s *ServiceDeploymentStatus) SetStartTime(v string) *ServiceDeploymentStatus {
	s.StartTime = &v
	return s
}

func (s *ServiceDeploymentStatus) SetTaskName(v string) *ServiceDeploymentStatus {
	s.TaskName = &v
	return s
}

func (s *ServiceDeploymentStatus) Validate() error {
	return dara.Validate(s)
}
