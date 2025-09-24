// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewChangeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RenewChangeInstanceResponseBody
	GetCode() *string
	SetData(v *RenewChangeInstanceResponseBodyData) *RenewChangeInstanceResponseBody
	GetData() *RenewChangeInstanceResponseBodyData
	SetMessage(v string) *RenewChangeInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RenewChangeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewChangeInstanceResponseBody
	GetSuccess() *bool
}

type RenewChangeInstanceResponseBody struct {
	// example:
	//
	// Success
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RenewChangeInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewChangeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewChangeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewChangeInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenewChangeInstanceResponseBody) GetData() *RenewChangeInstanceResponseBodyData {
	return s.Data
}

func (s *RenewChangeInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenewChangeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewChangeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewChangeInstanceResponseBody) SetCode(v string) *RenewChangeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RenewChangeInstanceResponseBody) SetData(v *RenewChangeInstanceResponseBodyData) *RenewChangeInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RenewChangeInstanceResponseBody) SetMessage(v string) *RenewChangeInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RenewChangeInstanceResponseBody) SetRequestId(v string) *RenewChangeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewChangeInstanceResponseBody) SetSuccess(v bool) *RenewChangeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RenewChangeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type RenewChangeInstanceResponseBodyData struct {
	// example:
	//
	// 100.100.64.1:8150
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 202407022550621
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewChangeInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RenewChangeInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewChangeInstanceResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *RenewChangeInstanceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewChangeInstanceResponseBodyData) SetHostId(v string) *RenewChangeInstanceResponseBodyData {
	s.HostId = &v
	return s
}

func (s *RenewChangeInstanceResponseBodyData) SetOrderId(v string) *RenewChangeInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RenewChangeInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
