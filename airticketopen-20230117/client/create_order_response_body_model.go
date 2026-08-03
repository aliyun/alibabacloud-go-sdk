// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateOrderResponseBodyData) *CreateOrderResponseBody
	GetData() *CreateOrderResponseBodyData
	SetErrorCode(v string) *CreateOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOrderResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *CreateOrderResponseBody
	GetTracerId() *string
}

type CreateOrderResponseBody struct {
	Data *CreateOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CreateOrderFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 创建订单失败
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) GetData() *CreateOrderResponseBodyData {
	return s.Data
}

func (s *CreateOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOrderResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateOrderResponseBody) SetData(v *CreateOrderResponseBodyData) *CreateOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateOrderResponseBody) SetErrorCode(v string) *CreateOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOrderResponseBody) SetErrorMsg(v string) *CreateOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetSuccess(v bool) *CreateOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrderResponseBody) SetTracerId(v string) *CreateOrderResponseBody {
	s.TracerId = &v
	return s
}

func (s *CreateOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrderResponseBodyData struct {
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyData) GetOrderNo() *string {
	return s.OrderNo
}

func (s *CreateOrderResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateOrderResponseBodyData) SetOrderNo(v string) *CreateOrderResponseBodyData {
	s.OrderNo = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetTracerId(v string) *CreateOrderResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *CreateOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
