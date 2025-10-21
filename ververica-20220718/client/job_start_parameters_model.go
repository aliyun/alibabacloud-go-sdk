// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobStartParameters interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *JobStartParameters
	GetDeploymentId() *string
	SetJobId(v string) *JobStartParameters
	GetJobId() *string
	SetLocalVariables(v []*LocalVariable) *JobStartParameters
	GetLocalVariables() []*LocalVariable
	SetResourceQueueName(v string) *JobStartParameters
	GetResourceQueueName() *string
	SetRestoreStrategy(v *DeploymentRestoreStrategy) *JobStartParameters
	GetRestoreStrategy() *DeploymentRestoreStrategy
}

type JobStartParameters struct {
	DeploymentId   *string          `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	JobId          *string          `json:"jobId,omitempty" xml:"jobId,omitempty"`
	LocalVariables []*LocalVariable `json:"localVariables,omitempty" xml:"localVariables,omitempty" type:"Repeated"`
	// example:
	//
	// default-queue
	ResourceQueueName *string                    `json:"resourceQueueName,omitempty" xml:"resourceQueueName,omitempty"`
	RestoreStrategy   *DeploymentRestoreStrategy `json:"restoreStrategy,omitempty" xml:"restoreStrategy,omitempty"`
}

func (s JobStartParameters) String() string {
	return dara.Prettify(s)
}

func (s JobStartParameters) GoString() string {
	return s.String()
}

func (s *JobStartParameters) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *JobStartParameters) GetJobId() *string {
	return s.JobId
}

func (s *JobStartParameters) GetLocalVariables() []*LocalVariable {
	return s.LocalVariables
}

func (s *JobStartParameters) GetResourceQueueName() *string {
	return s.ResourceQueueName
}

func (s *JobStartParameters) GetRestoreStrategy() *DeploymentRestoreStrategy {
	return s.RestoreStrategy
}

func (s *JobStartParameters) SetDeploymentId(v string) *JobStartParameters {
	s.DeploymentId = &v
	return s
}

func (s *JobStartParameters) SetJobId(v string) *JobStartParameters {
	s.JobId = &v
	return s
}

func (s *JobStartParameters) SetLocalVariables(v []*LocalVariable) *JobStartParameters {
	s.LocalVariables = v
	return s
}

func (s *JobStartParameters) SetResourceQueueName(v string) *JobStartParameters {
	s.ResourceQueueName = &v
	return s
}

func (s *JobStartParameters) SetRestoreStrategy(v *DeploymentRestoreStrategy) *JobStartParameters {
	s.RestoreStrategy = v
	return s
}

func (s *JobStartParameters) Validate() error {
	if s.LocalVariables != nil {
		for _, item := range s.LocalVariables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RestoreStrategy != nil {
		if err := s.RestoreStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
