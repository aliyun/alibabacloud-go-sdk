// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInquiryPriceRefundInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InquiryPriceRefundInstanceResponseBody
	GetCode() *string
	SetData(v *InquiryPriceRefundInstanceResponseBodyData) *InquiryPriceRefundInstanceResponseBody
	GetData() *InquiryPriceRefundInstanceResponseBodyData
	SetMessage(v string) *InquiryPriceRefundInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *InquiryPriceRefundInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InquiryPriceRefundInstanceResponseBody
	GetSuccess() *bool
}

type InquiryPriceRefundInstanceResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *InquiryPriceRefundInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the execution result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The Tracing Analysis character string for the execute operation. You can provide this string to the platform for troubleshooting.
	//
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is complete. A value of false indicates that an error occurred.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InquiryPriceRefundInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InquiryPriceRefundInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *InquiryPriceRefundInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *InquiryPriceRefundInstanceResponseBody) GetData() *InquiryPriceRefundInstanceResponseBodyData {
	return s.Data
}

func (s *InquiryPriceRefundInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InquiryPriceRefundInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InquiryPriceRefundInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InquiryPriceRefundInstanceResponseBody) SetCode(v string) *InquiryPriceRefundInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBody) SetData(v *InquiryPriceRefundInstanceResponseBodyData) *InquiryPriceRefundInstanceResponseBody {
	s.Data = v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBody) SetMessage(v string) *InquiryPriceRefundInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBody) SetRequestId(v string) *InquiryPriceRefundInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBody) SetSuccess(v bool) *InquiryPriceRefundInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InquiryPriceRefundInstanceResponseBodyData struct {
	// The currency.
	//
	// example:
	//
	// The currency of the refund amount: CNY: Chinese Yuan. USD: US Dollar. JPY: Japanese Yen
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The site where the execution host resides.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1etb69sqxgl4*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The refund amount.
	//
	// example:
	//
	// 12.34
	RefundAmount *float64 `json:"RefundAmount,omitempty" xml:"RefundAmount,omitempty"`
}

func (s InquiryPriceRefundInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InquiryPriceRefundInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *InquiryPriceRefundInstanceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *InquiryPriceRefundInstanceResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *InquiryPriceRefundInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InquiryPriceRefundInstanceResponseBodyData) GetRefundAmount() *float64 {
	return s.RefundAmount
}

func (s *InquiryPriceRefundInstanceResponseBodyData) SetCurrency(v string) *InquiryPriceRefundInstanceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBodyData) SetHostId(v string) *InquiryPriceRefundInstanceResponseBodyData {
	s.HostId = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBodyData) SetInstanceId(v string) *InquiryPriceRefundInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBodyData) SetRefundAmount(v float64) *InquiryPriceRefundInstanceResponseBodyData {
	s.RefundAmount = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
