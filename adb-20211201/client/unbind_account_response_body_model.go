// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindAccountResponseBody
	GetRequestId() *string
}

type UnbindAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 93E85E5C-C805-5837-8713-05B69A504EE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindAccountResponseBody) SetRequestId(v string) *UnbindAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
