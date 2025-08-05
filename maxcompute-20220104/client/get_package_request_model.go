// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceProject(v string) *GetPackageRequest
	GetSourceProject() *string
}

type GetPackageRequest struct {
	// The project to which the package belongs. This parameter is required if the package is installed in the MaxCompute project.
	//
	// example:
	//
	// projectB
	SourceProject *string `json:"sourceProject,omitempty" xml:"sourceProject,omitempty"`
}

func (s GetPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPackageRequest) GoString() string {
	return s.String()
}

func (s *GetPackageRequest) GetSourceProject() *string {
	return s.SourceProject
}

func (s *GetPackageRequest) SetSourceProject(v string) *GetPackageRequest {
	s.SourceProject = &v
	return s
}

func (s *GetPackageRequest) Validate() error {
	return dara.Validate(s)
}
