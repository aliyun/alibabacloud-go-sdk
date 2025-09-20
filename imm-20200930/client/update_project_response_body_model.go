// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProject(v *Project) *UpdateProjectResponseBody
	GetProject() *Project
	SetRequestId(v string) *UpdateProjectResponseBody
	GetRequestId() *string
}

type UpdateProjectResponseBody struct {
	// The project.
	Project *Project `json:"Project,omitempty" xml:"Project,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D33C3574-4093-448E-86E7-15BE2BD3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) GetProject() *Project {
	return s.Project
}

func (s *UpdateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectResponseBody) SetProject(v *Project) *UpdateProjectResponseBody {
	s.Project = v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
