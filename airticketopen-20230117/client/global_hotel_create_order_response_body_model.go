// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelCreateOrderResponseBodyData) *GlobalHotelCreateOrderResponseBody
	GetData() *GlobalHotelCreateOrderResponseBodyData
	SetErrorCode(v string) *GlobalHotelCreateOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelCreateOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelCreateOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelCreateOrderResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelCreateOrderResponseBody
	GetTracerId() *string
}

type GlobalHotelCreateOrderResponseBody struct {
	Data *GlobalHotelCreateOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateOrderResponseBody) GetData() *GlobalHotelCreateOrderResponseBodyData {
	return s.Data
}

func (s *GlobalHotelCreateOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelCreateOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelCreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelCreateOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelCreateOrderResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateOrderResponseBody) SetData(v *GlobalHotelCreateOrderResponseBodyData) *GlobalHotelCreateOrderResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelCreateOrderResponseBody) SetErrorCode(v string) *GlobalHotelCreateOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelCreateOrderResponseBody) SetErrorMsg(v string) *GlobalHotelCreateOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelCreateOrderResponseBody) SetRequestId(v string) *GlobalHotelCreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelCreateOrderResponseBody) SetSuccess(v bool) *GlobalHotelCreateOrderResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelCreateOrderResponseBody) SetTracerId(v string) *GlobalHotelCreateOrderResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelCreateOrderResponseBodyData struct {
	// example:
	//
	// SO202606290001
	OrderNo  *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCreateOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateOrderResponseBodyData) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GlobalHotelCreateOrderResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateOrderResponseBodyData) SetOrderNo(v string) *GlobalHotelCreateOrderResponseBodyData {
	s.OrderNo = &v
	return s
}

func (s *GlobalHotelCreateOrderResponseBodyData) SetTracerId(v string) *GlobalHotelCreateOrderResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
