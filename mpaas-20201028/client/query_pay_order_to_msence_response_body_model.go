// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPayOrderToMsenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMpaasUserGamecenterPaymentQuerystatusResponse(v *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse) *QueryPayOrderToMsenceResponseBody
	GetMpaasUserGamecenterPaymentQuerystatusResponse() *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse
	SetRequestId(v string) *QueryPayOrderToMsenceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryPayOrderToMsenceResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *QueryPayOrderToMsenceResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *QueryPayOrderToMsenceResponseBody
	GetSuccess() *bool
}

type QueryPayOrderToMsenceResponseBody struct {
	MpaasUserGamecenterPaymentQuerystatusResponse *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse `json:"MpaasUserGamecenterPaymentQuerystatusResponse,omitempty" xml:"MpaasUserGamecenterPaymentQuerystatusResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// SUCCESS
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPayOrderToMsenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPayOrderToMsenceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPayOrderToMsenceResponseBody) GetMpaasUserGamecenterPaymentQuerystatusResponse() *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse {
	return s.MpaasUserGamecenterPaymentQuerystatusResponse
}

func (s *QueryPayOrderToMsenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPayOrderToMsenceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryPayOrderToMsenceResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryPayOrderToMsenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPayOrderToMsenceResponseBody) SetMpaasUserGamecenterPaymentQuerystatusResponse(v *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse) *QueryPayOrderToMsenceResponseBody {
	s.MpaasUserGamecenterPaymentQuerystatusResponse = v
	return s
}

func (s *QueryPayOrderToMsenceResponseBody) SetRequestId(v string) *QueryPayOrderToMsenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPayOrderToMsenceResponseBody) SetResultCode(v string) *QueryPayOrderToMsenceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryPayOrderToMsenceResponseBody) SetResultMsg(v string) *QueryPayOrderToMsenceResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *QueryPayOrderToMsenceResponseBody) SetSuccess(v bool) *QueryPayOrderToMsenceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPayOrderToMsenceResponseBody) Validate() error {
	if s.MpaasUserGamecenterPaymentQuerystatusResponse != nil {
		if err := s.MpaasUserGamecenterPaymentQuerystatusResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse struct {
	// example:
	//
	// 1
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
}

func (s QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse) GoString() string {
	return s.String()
}

func (s *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse) SetOrderStatus(v string) *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse {
	s.OrderStatus = &v
	return s
}

func (s *QueryPayOrderToMsenceResponseBodyMpaasUserGamecenterPaymentQuerystatusResponse) Validate() error {
	return dara.Validate(s)
}
