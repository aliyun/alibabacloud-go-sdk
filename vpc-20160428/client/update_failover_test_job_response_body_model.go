// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFailoverTestJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFailoverTestJobResponseBody
	GetRequestId() *string
}

type UpdateFailoverTestJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFailoverTestJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFailoverTestJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFailoverTestJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFailoverTestJobResponseBody) SetRequestId(v string) *UpdateFailoverTestJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFailoverTestJobResponseBody) Validate() error {
	return dara.Validate(s)
}
