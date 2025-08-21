// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaqResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFaqResponseBody
	GetRequestId() *string
}

type UpdateFaqResponseBody struct {
	// example:
	//
	// 736994BD-AA35-4742-88C9-E64BE4BAA14B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFaqResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaqResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaqResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFaqResponseBody) SetRequestId(v string) *UpdateFaqResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFaqResponseBody) Validate() error {
	return dara.Validate(s)
}
