// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RefundInstanceResponseBody
	GetCode() *string
	SetData(v *RefundInstanceResponseBodyData) *RefundInstanceResponseBody
	GetData() *RefundInstanceResponseBodyData
	SetMessage(v string) *RefundInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RefundInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RefundInstanceResponseBody
	GetSuccess() *bool
}

type RefundInstanceResponseBody struct {
	// The response code.
	//
	// example:
	//
	// ResourceNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *RefundInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the execution result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The Tracing Analysis character string for the execute. Provide this string to the platform when troubleshooting issues.
	//
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is complete. A value of false indicates that an error occurred.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefundInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefundInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RefundInstanceResponseBody) GetData() *RefundInstanceResponseBodyData {
	return s.Data
}

func (s *RefundInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RefundInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefundInstanceResponseBody) SetCode(v string) *RefundInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RefundInstanceResponseBody) SetData(v *RefundInstanceResponseBodyData) *RefundInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RefundInstanceResponseBody) SetMessage(v string) *RefundInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RefundInstanceResponseBody) SetRequestId(v string) *RefundInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundInstanceResponseBody) SetSuccess(v bool) *RefundInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RefundInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefundInstanceResponseBodyData struct {
	// The site where the execution host resides.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The refund order ID returned when the unsubscription is successful.
	//
	// example:
	//
	// 2100000000***
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RefundInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefundInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *RefundInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RefundInstanceResponseBodyData) SetHostId(v string) *RefundInstanceResponseBodyData {
	s.HostId = &v
	return s
}

func (s *RefundInstanceResponseBodyData) SetOrderId(v int64) *RefundInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RefundInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
