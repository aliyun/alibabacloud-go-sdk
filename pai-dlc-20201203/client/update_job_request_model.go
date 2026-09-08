// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateJobRequest
	GetAccessibility() *string
	SetDescription(v string) *UpdateJobRequest
	GetDescription() *string
	SetJobSpecs(v []*JobSpec) *UpdateJobRequest
	GetJobSpecs() []*JobSpec
	SetPriority(v int32) *UpdateJobRequest
	GetPriority() *int32
	SetUserCommand(v string) *UpdateJobRequest
	GetUserCommand() *string
}

type UpdateJobRequest struct {
	// The visibility of the job. The visibility can only be expanded, not reduced. Valid values:
	//
	// - PUBLIC: visible to all users in the workspace.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The job specification definition.
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// The priority of the job. Valid values: 1 to 9.
	//
	// - 1: the lowest priority.
	//
	// - 9: the highest priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The user command.
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateJobRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateJobRequest) GetJobSpecs() []*JobSpec {
	return s.JobSpecs
}

func (s *UpdateJobRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateJobRequest) GetUserCommand() *string {
	return s.UserCommand
}

func (s *UpdateJobRequest) SetAccessibility(v string) *UpdateJobRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateJobRequest) SetDescription(v string) *UpdateJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateJobRequest) SetJobSpecs(v []*JobSpec) *UpdateJobRequest {
	s.JobSpecs = v
	return s
}

func (s *UpdateJobRequest) SetPriority(v int32) *UpdateJobRequest {
	s.Priority = &v
	return s
}

func (s *UpdateJobRequest) SetUserCommand(v string) *UpdateJobRequest {
	s.UserCommand = &v
	return s
}

func (s *UpdateJobRequest) Validate() error {
	if s.JobSpecs != nil {
		for _, item := range s.JobSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
