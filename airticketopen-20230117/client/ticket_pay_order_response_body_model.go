// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *TicketPayOrderResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v string) *TicketPayOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketPayOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketPayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketPayOrderResponseBody
	GetSuccess() *bool
}

type TicketPayOrderResponseBody struct {
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s TicketPayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *TicketPayOrderResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *TicketPayOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketPayOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketPayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketPayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketPayOrderResponseBody) SetData(v map[string]interface{}) *TicketPayOrderResponseBody {
	s.Data = v
	return s
}

func (s *TicketPayOrderResponseBody) SetErrorCode(v string) *TicketPayOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketPayOrderResponseBody) SetErrorMsg(v string) *TicketPayOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketPayOrderResponseBody) SetRequestId(v string) *TicketPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketPayOrderResponseBody) SetSuccess(v bool) *TicketPayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *TicketPayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
