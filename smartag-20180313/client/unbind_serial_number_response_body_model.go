// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSerialNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindSerialNumberResponseBody
	GetRequestId() *string
}

type UnbindSerialNumberResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 284045DE-4C2D-463D-9F27-B6898E67D120
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindSerialNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindSerialNumberResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSerialNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindSerialNumberResponseBody) SetRequestId(v string) *UnbindSerialNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSerialNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
