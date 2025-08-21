// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindTagResponseBody
	GetRequestId() *string
}

type UnbindTagResponseBody struct {
	// example:
	//
	// 159E4422-6624-4750-8943-DFD98D34858C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindTagResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindTagResponseBody) SetRequestId(v string) *UnbindTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindTagResponseBody) Validate() error {
	return dara.Validate(s)
}
