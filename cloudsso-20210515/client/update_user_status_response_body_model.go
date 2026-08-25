// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserStatusResponseBody
	GetRequestId() *string
}

type UpdateUserStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EE598602-AC67-56EF-B7CC-2927C30AA0A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserStatusResponseBody) SetRequestId(v string) *UpdateUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
