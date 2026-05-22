// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteProjectNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheck(v bool) *CheckSiteProjectNameResponseBody
	GetCheck() *bool
	SetDescription(v string) *CheckSiteProjectNameResponseBody
	GetDescription() *string
	SetProjectName(v string) *CheckSiteProjectNameResponseBody
	GetProjectName() *string
	SetRequestId(v string) *CheckSiteProjectNameResponseBody
	GetRequestId() *string
}

type CheckSiteProjectNameResponseBody struct {
	Check       *bool   `json:"Check,omitempty" xml:"Check,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckSiteProjectNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteProjectNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSiteProjectNameResponseBody) GetCheck() *bool {
	return s.Check
}

func (s *CheckSiteProjectNameResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CheckSiteProjectNameResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *CheckSiteProjectNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSiteProjectNameResponseBody) SetCheck(v bool) *CheckSiteProjectNameResponseBody {
	s.Check = &v
	return s
}

func (s *CheckSiteProjectNameResponseBody) SetDescription(v string) *CheckSiteProjectNameResponseBody {
	s.Description = &v
	return s
}

func (s *CheckSiteProjectNameResponseBody) SetProjectName(v string) *CheckSiteProjectNameResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CheckSiteProjectNameResponseBody) SetRequestId(v string) *CheckSiteProjectNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSiteProjectNameResponseBody) Validate() error {
	return dara.Validate(s)
}
