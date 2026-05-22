// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserProjectNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheck(v bool) *CheckUserProjectNameResponseBody
	GetCheck() *bool
	SetDescription(v string) *CheckUserProjectNameResponseBody
	GetDescription() *string
	SetProjectName(v string) *CheckUserProjectNameResponseBody
	GetProjectName() *string
	SetRequestId(v string) *CheckUserProjectNameResponseBody
	GetRequestId() *string
}

type CheckUserProjectNameResponseBody struct {
	Check       *bool   `json:"Check,omitempty" xml:"Check,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckUserProjectNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUserProjectNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserProjectNameResponseBody) GetCheck() *bool {
	return s.Check
}

func (s *CheckUserProjectNameResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CheckUserProjectNameResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *CheckUserProjectNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUserProjectNameResponseBody) SetCheck(v bool) *CheckUserProjectNameResponseBody {
	s.Check = &v
	return s
}

func (s *CheckUserProjectNameResponseBody) SetDescription(v string) *CheckUserProjectNameResponseBody {
	s.Description = &v
	return s
}

func (s *CheckUserProjectNameResponseBody) SetProjectName(v string) *CheckUserProjectNameResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CheckUserProjectNameResponseBody) SetRequestId(v string) *CheckUserProjectNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserProjectNameResponseBody) Validate() error {
	return dara.Validate(s)
}
