// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentPolicyShrink(v string) *CreateJobShrinkRequest
	GetDeploymentPolicyShrink() *string
	SetJobDescription(v string) *CreateJobShrinkRequest
	GetJobDescription() *string
	SetJobName(v string) *CreateJobShrinkRequest
	GetJobName() *string
	SetJobScheduler(v string) *CreateJobShrinkRequest
	GetJobScheduler() *string
	SetSecurityPolicyShrink(v string) *CreateJobShrinkRequest
	GetSecurityPolicyShrink() *string
	SetTasksShrink(v string) *CreateJobShrinkRequest
	GetTasksShrink() *string
}

type CreateJobShrinkRequest struct {
	DeploymentPolicyShrink *string `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty"`
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testjob
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobScheduler         *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	// This parameter is required.
	TasksShrink *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
}

func (s CreateJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobShrinkRequest) GetDeploymentPolicyShrink() *string {
	return s.DeploymentPolicyShrink
}

func (s *CreateJobShrinkRequest) GetJobDescription() *string {
	return s.JobDescription
}

func (s *CreateJobShrinkRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateJobShrinkRequest) GetJobScheduler() *string {
	return s.JobScheduler
}

func (s *CreateJobShrinkRequest) GetSecurityPolicyShrink() *string {
	return s.SecurityPolicyShrink
}

func (s *CreateJobShrinkRequest) GetTasksShrink() *string {
	return s.TasksShrink
}

func (s *CreateJobShrinkRequest) SetDeploymentPolicyShrink(v string) *CreateJobShrinkRequest {
	s.DeploymentPolicyShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobDescription(v string) *CreateJobShrinkRequest {
	s.JobDescription = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobName(v string) *CreateJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobScheduler(v string) *CreateJobShrinkRequest {
	s.JobScheduler = &v
	return s
}

func (s *CreateJobShrinkRequest) SetSecurityPolicyShrink(v string) *CreateJobShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTasksShrink(v string) *CreateJobShrinkRequest {
	s.TasksShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
