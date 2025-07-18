// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserGroupResponseBody
	GetRequestId() *string
}

type DeleteUserGroupResponseBody struct {
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
