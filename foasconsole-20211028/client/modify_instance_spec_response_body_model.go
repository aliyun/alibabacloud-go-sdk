// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyInstanceSpecResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyInstanceSpecResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyInstanceSpecResponseBody
	GetSuccess() *bool
}

type ModifyInstanceSpecResponseBody struct {
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

func (s ModifyInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceSpecResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyInstanceSpecResponseBody) SetOrderId(v int64) *ModifyInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceSpecResponseBody) SetRequestId(v string) *ModifyInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceSpecResponseBody) SetSuccess(v bool) *ModifyInstanceSpecResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
