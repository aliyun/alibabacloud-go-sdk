// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDcdnProjectExistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *CheckDcdnProjectExistRequest
	GetProjectName() *string
}

type CheckDcdnProjectExistRequest struct {
	// The name of a real-time log delivery project.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CheckDcdnProjectExistRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDcdnProjectExistRequest) GoString() string {
	return s.String()
}

func (s *CheckDcdnProjectExistRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CheckDcdnProjectExistRequest) SetProjectName(v string) *CheckDcdnProjectExistRequest {
	s.ProjectName = &v
	return s
}

func (s *CheckDcdnProjectExistRequest) Validate() error {
	return dara.Validate(s)
}
