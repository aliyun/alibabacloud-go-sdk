// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUsersResponseBody
	GetRequestId() *string
}

type DeleteUsersResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUsersResponseBody) SetRequestId(v string) *DeleteUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
