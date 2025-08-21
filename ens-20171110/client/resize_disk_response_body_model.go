// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ResizeDiskResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ResizeDiskResponseBody
	GetRequestId() *string
}

type ResizeDiskResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 21522202681****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 689412F2-8402-181E-8C87-1EF62331DCC4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeDiskResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ResizeDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeDiskResponseBody) SetOrderId(v string) *ResizeDiskResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeDiskResponseBody) SetRequestId(v string) *ResizeDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
