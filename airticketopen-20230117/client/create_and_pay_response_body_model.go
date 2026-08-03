// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAndPayResponseBodyData) *CreateAndPayResponseBody
	GetData() *CreateAndPayResponseBodyData
	SetErrorCode(v string) *CreateAndPayResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateAndPayResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateAndPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAndPayResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *CreateAndPayResponseBody
	GetTracerId() *string
}

type CreateAndPayResponseBody struct {
	Data *CreateAndPayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateAndPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndPayResponseBody) GetData() *CreateAndPayResponseBodyData {
	return s.Data
}

func (s *CreateAndPayResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAndPayResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateAndPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAndPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAndPayResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateAndPayResponseBody) SetData(v *CreateAndPayResponseBodyData) *CreateAndPayResponseBody {
	s.Data = v
	return s
}

func (s *CreateAndPayResponseBody) SetErrorCode(v string) *CreateAndPayResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAndPayResponseBody) SetErrorMsg(v string) *CreateAndPayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateAndPayResponseBody) SetRequestId(v string) *CreateAndPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndPayResponseBody) SetSuccess(v bool) *CreateAndPayResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAndPayResponseBody) SetTracerId(v string) *CreateAndPayResponseBody {
	s.TracerId = &v
	return s
}

func (s *CreateAndPayResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAndPayResponseBodyData struct {
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateAndPayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPayResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAndPayResponseBodyData) GetOrderNo() *string {
	return s.OrderNo
}

func (s *CreateAndPayResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateAndPayResponseBodyData) SetOrderNo(v string) *CreateAndPayResponseBodyData {
	s.OrderNo = &v
	return s
}

func (s *CreateAndPayResponseBodyData) SetTracerId(v string) *CreateAndPayResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *CreateAndPayResponseBodyData) Validate() error {
	return dara.Validate(s)
}
