// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDesktopGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserToDesktopGroupResponseBody
	GetRequestId() *string
}

type AddUserToDesktopGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToDesktopGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserToDesktopGroupResponseBody) SetRequestId(v string) *AddUserToDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToDesktopGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
