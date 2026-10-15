// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateResourcePackageResponseBody
	GetCode() *string
	SetData(v *CreateResourcePackageResponseBodyData) *CreateResourcePackageResponseBody
	GetData() *CreateResourcePackageResponseBodyData
	SetMessage(v string) *CreateResourcePackageResponseBody
	GetMessage() *string
	SetOrderId(v int64) *CreateResourcePackageResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateResourcePackageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateResourcePackageResponseBody
	GetSuccess() *bool
}

type CreateResourcePackageResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateResourcePackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId   *int64                                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourcePackageResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateResourcePackageResponseBody) GetData() *CreateResourcePackageResponseBodyData {
	return s.Data
}

func (s *CreateResourcePackageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateResourcePackageResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourcePackageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateResourcePackageResponseBody) SetCode(v string) *CreateResourcePackageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetData(v *CreateResourcePackageResponseBodyData) *CreateResourcePackageResponseBody {
	s.Data = v
	return s
}

func (s *CreateResourcePackageResponseBody) SetMessage(v string) *CreateResourcePackageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetOrderId(v int64) *CreateResourcePackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetRequestId(v string) *CreateResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetSuccess(v bool) *CreateResourcePackageResponseBody {
	s.Success = &v
	return s
}

func (s *CreateResourcePackageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateResourcePackageResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateResourcePackageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateResourcePackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateResourcePackageResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateResourcePackageResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateResourcePackageResponseBodyData) SetInstanceId(v string) *CreateResourcePackageResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateResourcePackageResponseBodyData) SetOrderId(v int64) *CreateResourcePackageResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateResourcePackageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
