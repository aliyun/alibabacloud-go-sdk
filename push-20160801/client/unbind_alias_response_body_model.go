// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindAliasResponseBody
	GetRequestId() *string
}

type UnbindAliasResponseBody struct {
	// example:
	//
	// 159E4422-6624-4750-8943-DFD98D34858C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindAliasResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindAliasResponseBody) SetRequestId(v string) *UnbindAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
