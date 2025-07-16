// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateProjectResponseBody
	GetRequestId() *string
}

type UpdateProjectResponseBody struct {
	// example:
	//
	// 2150233F-A1F7-54D2-B5B5-8A70567549BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
