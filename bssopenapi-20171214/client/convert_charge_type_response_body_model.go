// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConvertChargeTypeResponseBody
	GetCode() *string
	SetData(v *ConvertChargeTypeResponseBodyData) *ConvertChargeTypeResponseBody
	GetData() *ConvertChargeTypeResponseBodyData
	SetMessage(v string) *ConvertChargeTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConvertChargeTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConvertChargeTypeResponseBody
	GetSuccess() *bool
}

type ConvertChargeTypeResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ConvertChargeTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConvertChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertChargeTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConvertChargeTypeResponseBody) GetData() *ConvertChargeTypeResponseBodyData {
	return s.Data
}

func (s *ConvertChargeTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConvertChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertChargeTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConvertChargeTypeResponseBody) SetCode(v string) *ConvertChargeTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertChargeTypeResponseBody) SetData(v *ConvertChargeTypeResponseBodyData) *ConvertChargeTypeResponseBody {
	s.Data = v
	return s
}

func (s *ConvertChargeTypeResponseBody) SetMessage(v string) *ConvertChargeTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertChargeTypeResponseBody) SetRequestId(v string) *ConvertChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertChargeTypeResponseBody) SetSuccess(v bool) *ConvertChargeTypeResponseBody {
	s.Success = &v
	return s
}

func (s *ConvertChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ConvertChargeTypeResponseBodyData struct {
	// The ID of the order.
	//
	// example:
	//
	// 202657601410661
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ConvertChargeTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConvertChargeTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConvertChargeTypeResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *ConvertChargeTypeResponseBodyData) SetOrderId(v string) *ConvertChargeTypeResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ConvertChargeTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
