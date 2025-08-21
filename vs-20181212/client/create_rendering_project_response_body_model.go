// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *CreateRenderingProjectResponseBody
	GetProjectId() *string
	SetRequestId(v string) *CreateRenderingProjectResponseBody
	GetRequestId() *string
}

type CreateRenderingProjectResponseBody struct {
	// example:
	//
	// project-b93ea81de76f48609eed3cd420f0399f
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRenderingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRenderingProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateRenderingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRenderingProjectResponseBody) SetProjectId(v string) *CreateRenderingProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateRenderingProjectResponseBody) SetRequestId(v string) *CreateRenderingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRenderingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
