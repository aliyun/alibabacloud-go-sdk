// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketApplyRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *TicketApplyRefundResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v string) *TicketApplyRefundResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketApplyRefundResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketApplyRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketApplyRefundResponseBody
	GetSuccess() *bool
}

type TicketApplyRefundResponseBody struct {
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

func (s TicketApplyRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketApplyRefundResponseBody) GoString() string {
	return s.String()
}

func (s *TicketApplyRefundResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *TicketApplyRefundResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketApplyRefundResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketApplyRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketApplyRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketApplyRefundResponseBody) SetData(v map[string]interface{}) *TicketApplyRefundResponseBody {
	s.Data = v
	return s
}

func (s *TicketApplyRefundResponseBody) SetErrorCode(v string) *TicketApplyRefundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketApplyRefundResponseBody) SetErrorMsg(v string) *TicketApplyRefundResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketApplyRefundResponseBody) SetRequestId(v string) *TicketApplyRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketApplyRefundResponseBody) SetSuccess(v bool) *TicketApplyRefundResponseBody {
	s.Success = &v
	return s
}

func (s *TicketApplyRefundResponseBody) Validate() error {
	return dara.Validate(s)
}
