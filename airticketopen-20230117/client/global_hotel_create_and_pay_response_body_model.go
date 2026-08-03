// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCreateAndPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelCreateAndPayResponseBodyData) *GlobalHotelCreateAndPayResponseBody
	GetData() *GlobalHotelCreateAndPayResponseBodyData
	SetErrorCode(v string) *GlobalHotelCreateAndPayResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelCreateAndPayResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelCreateAndPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelCreateAndPayResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelCreateAndPayResponseBody
	GetTracerId() *string
}

type GlobalHotelCreateAndPayResponseBody struct {
	Data *GlobalHotelCreateAndPayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GlobalHotelCreateAndPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateAndPayResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateAndPayResponseBody) GetData() *GlobalHotelCreateAndPayResponseBodyData {
	return s.Data
}

func (s *GlobalHotelCreateAndPayResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelCreateAndPayResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelCreateAndPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelCreateAndPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelCreateAndPayResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateAndPayResponseBody) SetData(v *GlobalHotelCreateAndPayResponseBodyData) *GlobalHotelCreateAndPayResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelCreateAndPayResponseBody) SetErrorCode(v string) *GlobalHotelCreateAndPayResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelCreateAndPayResponseBody) SetErrorMsg(v string) *GlobalHotelCreateAndPayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelCreateAndPayResponseBody) SetRequestId(v string) *GlobalHotelCreateAndPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelCreateAndPayResponseBody) SetSuccess(v bool) *GlobalHotelCreateAndPayResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelCreateAndPayResponseBody) SetTracerId(v string) *GlobalHotelCreateAndPayResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateAndPayResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelCreateAndPayResponseBodyData struct {
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCreateAndPayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateAndPayResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateAndPayResponseBodyData) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GlobalHotelCreateAndPayResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateAndPayResponseBodyData) SetOrderNo(v string) *GlobalHotelCreateAndPayResponseBodyData {
	s.OrderNo = &v
	return s
}

func (s *GlobalHotelCreateAndPayResponseBodyData) SetTracerId(v string) *GlobalHotelCreateAndPayResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateAndPayResponseBodyData) Validate() error {
	return dara.Validate(s)
}
