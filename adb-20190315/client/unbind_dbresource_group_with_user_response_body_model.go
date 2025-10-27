// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourceGroupWithUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindDBResourceGroupWithUserResponseBody
	GetRequestId() *string
}

type UnbindDBResourceGroupWithUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDBResourceGroupWithUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindDBResourceGroupWithUserResponseBody) SetRequestId(v string) *UnbindDBResourceGroupWithUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserResponseBody) Validate() error {
	return dara.Validate(s)
}
