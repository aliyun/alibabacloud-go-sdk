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
	// The ID of the project.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
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
