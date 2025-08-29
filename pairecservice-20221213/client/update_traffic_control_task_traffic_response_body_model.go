// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTaskTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrafficControlTaskTrafficResponseBody
	GetRequestId() *string
}

type UpdateTrafficControlTaskTrafficResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficControlTaskTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTaskTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrafficControlTaskTrafficResponseBody) SetRequestId(v string) *UpdateTrafficControlTaskTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficResponseBody) Validate() error {
	return dara.Validate(s)
}
