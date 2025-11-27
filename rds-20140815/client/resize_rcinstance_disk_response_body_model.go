// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeRCInstanceDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ResizeRCInstanceDiskResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ResizeRCInstanceDiskResponseBody
	GetRequestId() *string
}

type ResizeRCInstanceDiskResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 230546833080102
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeRCInstanceDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeRCInstanceDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeRCInstanceDiskResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ResizeRCInstanceDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeRCInstanceDiskResponseBody) SetOrderId(v int64) *ResizeRCInstanceDiskResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeRCInstanceDiskResponseBody) SetRequestId(v string) *ResizeRCInstanceDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeRCInstanceDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
