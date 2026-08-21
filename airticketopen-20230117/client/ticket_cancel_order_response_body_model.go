// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCancelOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *TicketCancelOrderResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v string) *TicketCancelOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketCancelOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketCancelOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketCancelOrderResponseBody
	GetSuccess() *bool
}

type TicketCancelOrderResponseBody struct {
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

func (s TicketCancelOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketCancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *TicketCancelOrderResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *TicketCancelOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketCancelOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketCancelOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketCancelOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketCancelOrderResponseBody) SetData(v map[string]interface{}) *TicketCancelOrderResponseBody {
	s.Data = v
	return s
}

func (s *TicketCancelOrderResponseBody) SetErrorCode(v string) *TicketCancelOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketCancelOrderResponseBody) SetErrorMsg(v string) *TicketCancelOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketCancelOrderResponseBody) SetRequestId(v string) *TicketCancelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketCancelOrderResponseBody) SetSuccess(v bool) *TicketCancelOrderResponseBody {
	s.Success = &v
	return s
}

func (s *TicketCancelOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
