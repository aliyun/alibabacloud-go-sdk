// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyInstanceResponseBody
	GetCode() *string
	SetData(v *ModifyInstanceResponseBodyData) *ModifyInstanceResponseBody
	GetData() *ModifyInstanceResponseBodyData
	SetMessage(v string) *ModifyInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyInstanceResponseBody
	GetSuccess() *bool
}

type ModifyInstanceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ModifyInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ModifyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyInstanceResponseBody) GetData() *ModifyInstanceResponseBodyData {
	return s.Data
}

func (s *ModifyInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyInstanceResponseBody) SetCode(v string) *ModifyInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetData(v *ModifyInstanceResponseBodyData) *ModifyInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ModifyInstanceResponseBody) SetMessage(v string) *ModifyInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetSuccess(v bool) *ModifyInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceResponseBodyData struct {
	// The ID of the host.
	//
	// example:
	//
	// testId
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the order that was created.
	//
	// example:
	//
	// 202653252354351
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *ModifyInstanceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyInstanceResponseBodyData) SetHostId(v string) *ModifyInstanceResponseBodyData {
	s.HostId = &v
	return s
}

func (s *ModifyInstanceResponseBodyData) SetOrderId(v string) *ModifyInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
