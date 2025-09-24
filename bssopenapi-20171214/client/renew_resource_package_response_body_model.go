// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RenewResourcePackageResponseBody
	GetCode() *string
	SetData(v *RenewResourcePackageResponseBodyData) *RenewResourcePackageResponseBody
	GetData() *RenewResourcePackageResponseBodyData
	SetMessage(v string) *RenewResourcePackageResponseBody
	GetMessage() *string
	SetOrderId(v int64) *RenewResourcePackageResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *RenewResourcePackageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewResourcePackageResponseBody
	GetSuccess() *bool
}

type RenewResourcePackageResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *RenewResourcePackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 204322301110333
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *RenewResourcePackageResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenewResourcePackageResponseBody) GetData() *RenewResourcePackageResponseBodyData {
	return s.Data
}

func (s *RenewResourcePackageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenewResourcePackageResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewResourcePackageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewResourcePackageResponseBody) SetCode(v string) *RenewResourcePackageResponseBody {
	s.Code = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetData(v *RenewResourcePackageResponseBodyData) *RenewResourcePackageResponseBody {
	s.Data = v
	return s
}

func (s *RenewResourcePackageResponseBody) SetMessage(v string) *RenewResourcePackageResponseBody {
	s.Message = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetOrderId(v int64) *RenewResourcePackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetRequestId(v string) *RenewResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetSuccess(v bool) *RenewResourcePackageResponseBody {
	s.Success = &v
	return s
}

func (s *RenewResourcePackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type RenewResourcePackageResponseBodyData struct {
	// The ID of the resource plan.
	//
	// example:
	//
	// OSSBAG-cn-0xl0n****003
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 204322560333
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewResourcePackageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RenewResourcePackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewResourcePackageResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewResourcePackageResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewResourcePackageResponseBodyData) SetInstanceId(v string) *RenewResourcePackageResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *RenewResourcePackageResponseBodyData) SetOrderId(v int64) *RenewResourcePackageResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RenewResourcePackageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
