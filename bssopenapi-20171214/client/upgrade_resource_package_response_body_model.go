// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpgradeResourcePackageResponseBody
	GetCode() *string
	SetData(v *UpgradeResourcePackageResponseBodyData) *UpgradeResourcePackageResponseBody
	GetData() *UpgradeResourcePackageResponseBodyData
	SetMessage(v string) *UpgradeResourcePackageResponseBody
	GetMessage() *string
	SetOrderId(v int64) *UpgradeResourcePackageResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *UpgradeResourcePackageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeResourcePackageResponseBody
	GetSuccess() *bool
}

type UpgradeResourcePackageResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *UpgradeResourcePackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 73387246238746
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

func (s UpgradeResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeResourcePackageResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeResourcePackageResponseBody) GetData() *UpgradeResourcePackageResponseBodyData {
	return s.Data
}

func (s *UpgradeResourcePackageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeResourcePackageResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *UpgradeResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeResourcePackageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeResourcePackageResponseBody) SetCode(v string) *UpgradeResourcePackageResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetData(v *UpgradeResourcePackageResponseBodyData) *UpgradeResourcePackageResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetMessage(v string) *UpgradeResourcePackageResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetOrderId(v int64) *UpgradeResourcePackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetRequestId(v string) *UpgradeResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetSuccess(v bool) *UpgradeResourcePackageResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeResourcePackageResponseBodyData struct {
	// The ID of the resource plan.
	//
	// example:
	//
	// OSSBAG-cn-0xl*****002
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 73387246238746
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s UpgradeResourcePackageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeResourcePackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeResourcePackageResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeResourcePackageResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *UpgradeResourcePackageResponseBodyData) SetInstanceId(v string) *UpgradeResourcePackageResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpgradeResourcePackageResponseBodyData) SetOrderId(v int64) *UpgradeResourcePackageResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *UpgradeResourcePackageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
