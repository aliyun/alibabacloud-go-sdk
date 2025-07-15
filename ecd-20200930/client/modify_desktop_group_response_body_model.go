// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDesktopGroupResponseBody
	GetRequestId() *string
}

type ModifyDesktopGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopGroupResponseBody) SetRequestId(v string) *ModifyDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
