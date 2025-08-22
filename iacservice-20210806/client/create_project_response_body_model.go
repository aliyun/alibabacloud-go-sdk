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
	// p-433aead7560572f8d95b25775c
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 3E49127A-BB65-5CCD-AB93-0EC0A43E5446
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
