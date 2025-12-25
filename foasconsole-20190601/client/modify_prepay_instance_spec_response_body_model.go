// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyPrepayInstanceSpecResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyPrepayInstanceSpecResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyPrepayInstanceSpecResponseBody
	GetSuccess() *bool
}

type ModifyPrepayInstanceSpecResponseBody struct {
	// example:
	//
	// 210406354690749
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

func (s ModifyPrepayInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyPrepayInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPrepayInstanceSpecResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetOrderId(v int64) *ModifyPrepayInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetRequestId(v string) *ModifyPrepayInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetSuccess(v bool) *ModifyPrepayInstanceSpecResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
