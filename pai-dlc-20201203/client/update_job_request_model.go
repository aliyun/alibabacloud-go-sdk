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
}

type UpdateJobRequest struct {
	// The job\\"s visibility. You can only increase, not decrease, the visibility. Valid value:
	//
	// - `PUBLIC`: The job is visible to all users in the workspace.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The job specifications.
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// The priority of the job. Valid values range from 1 to 9:
	//
	// - 1 indicates the lowest priority.
	//
	// - 9 indicates the highest priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
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
