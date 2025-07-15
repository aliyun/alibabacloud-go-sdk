// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindUserDesktopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindUserDesktopResponseBody
	GetRequestId() *string
}

type UnbindUserDesktopResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindUserDesktopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindUserDesktopResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindUserDesktopResponseBody) SetRequestId(v string) *UnbindUserDesktopResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindUserDesktopResponseBody) Validate() error {
	return dara.Validate(s)
}
