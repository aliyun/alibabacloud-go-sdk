// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromDesktopOversoldUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserFromDesktopOversoldUserGroupResponseBody
	GetRequestId() *string
}

type RemoveUserFromDesktopOversoldUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromDesktopOversoldUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromDesktopOversoldUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponseBody) SetRequestId(v string) *RemoveUserFromDesktopOversoldUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
