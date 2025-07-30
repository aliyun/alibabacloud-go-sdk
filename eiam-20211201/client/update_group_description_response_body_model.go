// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGroupDescriptionResponseBody
	GetRequestId() *string
}

type UpdateGroupDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupDescriptionResponseBody) SetRequestId(v string) *UpdateGroupDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
