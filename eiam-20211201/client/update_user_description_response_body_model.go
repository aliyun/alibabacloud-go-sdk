// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserDescriptionResponseBody
	GetRequestId() *string
}

type UpdateUserDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserDescriptionResponseBody) SetRequestId(v string) *UpdateUserDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
