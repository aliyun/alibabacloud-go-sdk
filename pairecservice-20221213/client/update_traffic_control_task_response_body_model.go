// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrafficControlTaskResponseBody
	GetRequestId() *string
}

type UpdateTrafficControlTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficControlTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrafficControlTaskResponseBody) SetRequestId(v string) *UpdateTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrafficControlTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
