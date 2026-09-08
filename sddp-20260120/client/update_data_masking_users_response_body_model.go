// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataMaskingUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataMaskingUsersResponseBody
	GetRequestId() *string
}

type UpdateDataMaskingUsersResponseBody struct {
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataMaskingUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingUsersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataMaskingUsersResponseBody) SetRequestId(v string) *UpdateDataMaskingUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataMaskingUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
