// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateExperimentRequest
	GetAccessibility() *string
	SetName(v string) *UpdateExperimentRequest
	GetName() *string
}

type UpdateExperimentRequest struct {
	// The accessibility of the experiment in the workspace. Valid values:
	//
	// 	- PRIVATE: The experiment is accessible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: The experiment is accessible to all users in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The experiment name. The name must meet the following requirements:
	//
	// 	- The name must start with a letter.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be 1 to 63 characters in length.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateExperimentRequest) GetName() *string {
	return s.Name
}

func (s *UpdateExperimentRequest) SetAccessibility(v string) *UpdateExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateExperimentRequest) SetName(v string) *UpdateExperimentRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentRequest) Validate() error {
	return dara.Validate(s)
}
