// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ResizeDiskSizeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ResizeDiskSizeResponseBody
	GetRequestId() *string
}

type ResizeDiskSizeResponseBody struct {
	// example:
	//
	// 3C22622B-8555-42BF-AD8A-1B960743****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 493A762B-E4A6-44E9-B877-CA6D0CAF8B29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeDiskSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ResizeDiskSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeDiskSizeResponseBody) SetOrderId(v string) *ResizeDiskSizeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeDiskSizeResponseBody) SetRequestId(v string) *ResizeDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeDiskSizeResponseBody) Validate() error {
	return dara.Validate(s)
}
