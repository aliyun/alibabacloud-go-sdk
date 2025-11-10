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
	SetJobSpecs(v []*JobSpec) *UpdateJobRequest
	GetJobSpecs() []*JobSpec
	SetPriority(v int32) *UpdateJobRequest
	GetPriority() *int32
}

type UpdateJobRequest struct {
	// The job visibility. Valid values:
	//
	// 	- PUBLIC: The job is visible to all members in the workspace.
	//
	// 	- PRIVATE: The job is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string    `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	JobSpecs      []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// The job priority. Valid values: 1 to 9.
	//
	// 	- 1: the lowest priority.
	//
	// 	- 9: the highest priority.
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
