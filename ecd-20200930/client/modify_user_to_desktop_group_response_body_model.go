// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserToDesktopGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserToDesktopGroupResponseBody
	GetRequestId() *string
}

type ModifyUserToDesktopGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserToDesktopGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserToDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserToDesktopGroupResponseBody) SetRequestId(v string) *ModifyUserToDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserToDesktopGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
