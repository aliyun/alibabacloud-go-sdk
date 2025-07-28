// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserVpcAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserVpcAuthorizationResponseBody
	GetRequestId() *string
}

type AddUserVpcAuthorizationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserVpcAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserVpcAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserVpcAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserVpcAuthorizationResponseBody) SetRequestId(v string) *AddUserVpcAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserVpcAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
