// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantRoleToUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantRoleToUsersResponseBody
	GetRequestId() *string
}

type GrantRoleToUsersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GrantRoleToUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantRoleToUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantRoleToUsersResponseBody) SetRequestId(v string) *GrantRoleToUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantRoleToUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
