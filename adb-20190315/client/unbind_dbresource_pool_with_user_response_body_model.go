// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourcePoolWithUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindDBResourcePoolWithUserResponseBody
	GetRequestId() *string
}

type UnbindDBResourcePoolWithUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDBResourcePoolWithUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindDBResourcePoolWithUserResponseBody) SetRequestId(v string) *UnbindDBResourcePoolWithUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserResponseBody) Validate() error {
	return dara.Validate(s)
}
