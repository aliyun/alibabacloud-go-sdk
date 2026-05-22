// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserProjectNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *CheckUserProjectNameRequest
	GetProjectName() *string
}

type CheckUserProjectNameRequest struct {
	// The name of the real-time log delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ali-dcdn-log-56
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CheckUserProjectNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUserProjectNameRequest) GoString() string {
	return s.String()
}

func (s *CheckUserProjectNameRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CheckUserProjectNameRequest) SetProjectName(v string) *CheckUserProjectNameRequest {
	s.ProjectName = &v
	return s
}

func (s *CheckUserProjectNameRequest) Validate() error {
	return dara.Validate(s)
}
