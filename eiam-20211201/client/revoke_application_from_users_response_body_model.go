// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeApplicationFromUsersResponseBody
	GetRequestId() *string
}

type RevokeApplicationFromUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeApplicationFromUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeApplicationFromUsersResponseBody) SetRequestId(v string) *RevokeApplicationFromUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeApplicationFromUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
