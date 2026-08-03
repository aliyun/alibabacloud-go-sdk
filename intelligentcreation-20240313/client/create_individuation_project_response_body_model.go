// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndividuationProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *CreateIndividuationProjectResponseBody
	GetProjectId() *string
	SetRequestId(v string) *CreateIndividuationProjectResponseBody
	GetRequestId() *string
}

type CreateIndividuationProjectResponseBody struct {
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateIndividuationProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIndividuationProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndividuationProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateIndividuationProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIndividuationProjectResponseBody) SetProjectId(v string) *CreateIndividuationProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateIndividuationProjectResponseBody) SetRequestId(v string) *CreateIndividuationProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndividuationProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
