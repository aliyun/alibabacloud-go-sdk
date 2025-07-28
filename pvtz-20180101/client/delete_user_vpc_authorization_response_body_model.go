// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserVpcAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserVpcAuthorizationResponseBody
	GetRequestId() *string
}

type DeleteUserVpcAuthorizationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserVpcAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserVpcAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserVpcAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserVpcAuthorizationResponseBody) SetRequestId(v string) *DeleteUserVpcAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserVpcAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
