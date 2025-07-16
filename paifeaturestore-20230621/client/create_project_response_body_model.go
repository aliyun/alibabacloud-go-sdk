// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *CreateProjectResponseBody
	GetProjectId() *string
	SetRequestId(v string) *CreateProjectResponseBody
	GetRequestId() *string
}

type CreateProjectResponseBody struct {
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 977ADE3A-D69D-58AD-8323-96E2FB898E99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectResponseBody) SetProjectId(v string) *CreateProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
