// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDesktopManagerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchSetDesktopManagerResponseBody
	GetRequestId() *string
}

type BatchSetDesktopManagerResponseBody struct {
	// example:
	//
	// 868B8926-2E7A-5BE7-9897-E811E994****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetDesktopManagerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDesktopManagerResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetDesktopManagerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetDesktopManagerResponseBody) SetRequestId(v string) *BatchSetDesktopManagerResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetDesktopManagerResponseBody) Validate() error {
	return dara.Validate(s)
}
