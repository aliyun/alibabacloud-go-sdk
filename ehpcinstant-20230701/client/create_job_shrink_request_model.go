// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDependencyPolicyShrink(v string) *CreateJobShrinkRequest
	GetDependencyPolicyShrink() *string
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
	// Dependency policy.
	DependencyPolicyShrink *string `json:"DependencyPolicy,omitempty" xml:"DependencyPolicy,omitempty"`
	// The resource deployment policy.
	DeploymentPolicyShrink *string `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty"`
	// The description of the job.
	//
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// The job name. The name must be 2 to 64 characters in length and can contain letters, digits, and Chinese characters. It can contain hyphens (-) and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// testjob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The type of the job scheduler.
	//
	// 	- HPC
	//
	// 	- K8S
	//
	// Default value: HPC
	//
	// example:
	//
	// HPC
	JobScheduler *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	// The security policy.
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	// The list of tasks. Only one task is supported.
	//
	// This parameter is required.
	TasksShrink *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
}

func (s CreateJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobShrinkRequest) GetDependencyPolicyShrink() *string {
	return s.DependencyPolicyShrink
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

func (s *CreateJobShrinkRequest) SetDependencyPolicyShrink(v string) *CreateJobShrinkRequest {
	s.DependencyPolicyShrink = &v
	return s
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
