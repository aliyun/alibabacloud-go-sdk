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
	// C62888F6-254D-5589-BF05-0D9EE698C187
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
