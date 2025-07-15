// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopHostNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDesktopHostNameResponseBody
	GetRequestId() *string
}

type ModifyDesktopHostNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopHostNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopHostNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopHostNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopHostNameResponseBody) SetRequestId(v string) *ModifyDesktopHostNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopHostNameResponseBody) Validate() error {
	return dara.Validate(s)
}
