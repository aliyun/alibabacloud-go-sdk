// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEndUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachEndUserResponseBody
	GetRequestId() *string
}

type DetachEndUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachEndUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *DetachEndUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachEndUserResponseBody) SetRequestId(v string) *DetachEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachEndUserResponseBody) Validate() error {
	return dara.Validate(s)
}
