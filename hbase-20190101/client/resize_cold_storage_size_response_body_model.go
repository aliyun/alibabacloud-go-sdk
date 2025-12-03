// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeColdStorageSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ResizeColdStorageSizeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ResizeColdStorageSizeResponseBody
	GetRequestId() *string
}

type ResizeColdStorageSizeResponseBody struct {
	// example:
	//
	// 21711518427****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 5AA6F80E-535C-5611-BD13-3832D96A4D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeColdStorageSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeColdStorageSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ResizeColdStorageSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeColdStorageSizeResponseBody) SetOrderId(v string) *ResizeColdStorageSizeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeColdStorageSizeResponseBody) SetRequestId(v string) *ResizeColdStorageSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeColdStorageSizeResponseBody) Validate() error {
	return dara.Validate(s)
}
