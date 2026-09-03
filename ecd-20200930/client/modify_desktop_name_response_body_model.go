// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDesktopNameResponseBody
	GetRequestId() *string
}

type ModifyDesktopNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopNameResponseBody) SetRequestId(v string) *ModifyDesktopNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopNameResponseBody) Validate() error {
	return dara.Validate(s)
}
