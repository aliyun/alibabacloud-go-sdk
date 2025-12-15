// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUsersResponseBody
	GetRequestId() *string
}

type CreateUsersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUsersResponseBody) SetRequestId(v string) *CreateUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
