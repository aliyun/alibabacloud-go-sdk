// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GlobalHotelPayResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelPayResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelPayResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelPayResponseBody
	GetTracerId() *string
}

type GlobalHotelPayResponseBody struct {
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

func (s GlobalHotelPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelPayResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelPayResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelPayResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelPayResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelPayResponseBody) SetErrorCode(v string) *GlobalHotelPayResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelPayResponseBody) SetErrorMsg(v string) *GlobalHotelPayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelPayResponseBody) SetRequestId(v string) *GlobalHotelPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelPayResponseBody) SetSuccess(v bool) *GlobalHotelPayResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelPayResponseBody) SetTracerId(v string) *GlobalHotelPayResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelPayResponseBody) Validate() error {
	return dara.Validate(s)
}
