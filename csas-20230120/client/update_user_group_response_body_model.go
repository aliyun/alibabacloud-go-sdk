// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserGroupResponseBody
	GetRequestId() *string
}

type UpdateUserGroupResponseBody struct {
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
