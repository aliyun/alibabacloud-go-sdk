// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketCreateOrderResponseBodyData) *TicketCreateOrderResponseBody
	GetData() *TicketCreateOrderResponseBodyData
	SetErrorCode(v string) *TicketCreateOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketCreateOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketCreateOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketCreateOrderResponseBody
	GetSuccess() *bool
}

type TicketCreateOrderResponseBody struct {
	Data *TicketCreateOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DistributorOrderIdInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 分销商订单号不合法
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TicketCreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderResponseBody) GetData() *TicketCreateOrderResponseBodyData {
	return s.Data
}

func (s *TicketCreateOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketCreateOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketCreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketCreateOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketCreateOrderResponseBody) SetData(v *TicketCreateOrderResponseBodyData) *TicketCreateOrderResponseBody {
	s.Data = v
	return s
}

func (s *TicketCreateOrderResponseBody) SetErrorCode(v string) *TicketCreateOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketCreateOrderResponseBody) SetErrorMsg(v string) *TicketCreateOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketCreateOrderResponseBody) SetRequestId(v string) *TicketCreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketCreateOrderResponseBody) SetSuccess(v bool) *TicketCreateOrderResponseBody {
	s.Success = &v
	return s
}

func (s *TicketCreateOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketCreateOrderResponseBodyData struct {
	// example:
	//
	// 123456
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s TicketCreateOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *TicketCreateOrderResponseBodyData) SetOrderId(v string) *TicketCreateOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *TicketCreateOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
