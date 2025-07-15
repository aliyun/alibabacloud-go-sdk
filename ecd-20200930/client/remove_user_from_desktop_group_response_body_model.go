// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromDesktopGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserFromDesktopGroupResponseBody
	GetRequestId() *string
}

type RemoveUserFromDesktopGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromDesktopGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserFromDesktopGroupResponseBody) SetRequestId(v string) *RemoveUserFromDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
