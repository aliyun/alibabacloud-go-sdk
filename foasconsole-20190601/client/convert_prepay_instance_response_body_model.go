// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPrepayInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ConvertPrepayInstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ConvertPrepayInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConvertPrepayInstanceResponseBody
	GetSuccess() *bool
}

type ConvertPrepayInstanceResponseBody struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConvertPrepayInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertPrepayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ConvertPrepayInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertPrepayInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConvertPrepayInstanceResponseBody) SetOrderId(v int64) *ConvertPrepayInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertPrepayInstanceResponseBody) SetRequestId(v string) *ConvertPrepayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertPrepayInstanceResponseBody) SetSuccess(v bool) *ConvertPrepayInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ConvertPrepayInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
