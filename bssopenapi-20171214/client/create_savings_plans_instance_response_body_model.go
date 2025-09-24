// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavingsPlansInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSavingsPlansInstanceResponseBody
	GetCode() *string
	SetData(v *CreateSavingsPlansInstanceResponseBodyData) *CreateSavingsPlansInstanceResponseBody
	GetData() *CreateSavingsPlansInstanceResponseBodyData
	SetMessage(v string) *CreateSavingsPlansInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSavingsPlansInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSavingsPlansInstanceResponseBody
	GetSuccess() *bool
}

type CreateSavingsPlansInstanceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// PARAM_ERROR
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CreateSavingsPlansInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// The parameter must be specified.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 26dabb0c-8ca0-4aa0-8143-30499f3fe304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSavingsPlansInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlansInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlansInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSavingsPlansInstanceResponseBody) GetData() *CreateSavingsPlansInstanceResponseBodyData {
	return s.Data
}

func (s *CreateSavingsPlansInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSavingsPlansInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSavingsPlansInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSavingsPlansInstanceResponseBody) SetCode(v string) *CreateSavingsPlansInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) SetData(v *CreateSavingsPlansInstanceResponseBodyData) *CreateSavingsPlansInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) SetMessage(v string) *CreateSavingsPlansInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) SetRequestId(v string) *CreateSavingsPlansInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) SetSuccess(v bool) *CreateSavingsPlansInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSavingsPlansInstanceResponseBodyData struct {
	// The ID of the order.
	//
	// example:
	//
	// 202110260001
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateSavingsPlansInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlansInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlansInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateSavingsPlansInstanceResponseBodyData) SetOrderId(v int64) *CreateSavingsPlansInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
