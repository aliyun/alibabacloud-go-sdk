// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RenewInstanceResponseBody
	GetCode() *string
	SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody
	GetData() *RenewInstanceResponseBodyData
	SetMessage(v string) *RenewInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RenewInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewInstanceResponseBody
	GetSuccess() *bool
}

type RenewInstanceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *RenewInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s RenewInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenewInstanceResponseBody) GetData() *RenewInstanceResponseBodyData {
	return s.Data
}

func (s *RenewInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenewInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewInstanceResponseBody) SetCode(v string) *RenewInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RenewInstanceResponseBody) SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RenewInstanceResponseBody) SetMessage(v string) *RenewInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v bool) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RenewInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewInstanceResponseBodyData struct {
	// The ID of the order.
	//
	// example:
	//
	// 202657601410661
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewInstanceResponseBodyData) SetOrderId(v string) *RenewInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RenewInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
