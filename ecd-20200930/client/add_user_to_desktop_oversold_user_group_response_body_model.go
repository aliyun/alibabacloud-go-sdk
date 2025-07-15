// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDesktopOversoldUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserToDesktopOversoldUserGroupResponseBody
	GetRequestId() *string
}

type AddUserToDesktopOversoldUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToDesktopOversoldUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDesktopOversoldUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopOversoldUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserToDesktopOversoldUserGroupResponseBody) SetRequestId(v string) *AddUserToDesktopOversoldUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToDesktopOversoldUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
