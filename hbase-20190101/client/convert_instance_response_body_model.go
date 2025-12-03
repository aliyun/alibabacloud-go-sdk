// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ConvertInstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ConvertInstanceResponseBody
	GetRequestId() *string
}

type ConvertInstanceResponseBody struct {
	// example:
	//
	// 54124548879
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ConvertInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertInstanceResponseBody) SetOrderId(v int64) *ConvertInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertInstanceResponseBody) SetRequestId(v string) *ConvertInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
