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
	SetSuccess(v bool) *ConvertInstanceResponseBody
	GetSuccess() *bool
}

type ConvertInstanceResponseBody struct {
	// example:
	//
	// 211473228320700
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *ConvertInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConvertInstanceResponseBody) SetOrderId(v int64) *ConvertInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertInstanceResponseBody) SetRequestId(v string) *ConvertInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertInstanceResponseBody) SetSuccess(v bool) *ConvertInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ConvertInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
