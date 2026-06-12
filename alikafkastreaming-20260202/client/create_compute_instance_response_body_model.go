// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateComputeInstanceResponseBody
	GetCode() *int64
	SetData(v *CreateComputeInstanceResponseBodyData) *CreateComputeInstanceResponseBody
	GetData() *CreateComputeInstanceResponseBodyData
	SetRequestId(v string) *CreateComputeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateComputeInstanceResponseBody
	GetSuccess() *bool
}

type CreateComputeInstanceResponseBody struct {
	Code      *int64                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateComputeInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateComputeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeInstanceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateComputeInstanceResponseBody) GetData() *CreateComputeInstanceResponseBodyData {
	return s.Data
}

func (s *CreateComputeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComputeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateComputeInstanceResponseBody) SetCode(v int64) *CreateComputeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeInstanceResponseBody) SetData(v *CreateComputeInstanceResponseBodyData) *CreateComputeInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateComputeInstanceResponseBody) SetRequestId(v string) *CreateComputeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeInstanceResponseBody) SetSuccess(v bool) *CreateComputeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateComputeInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateComputeInstanceResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateComputeInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateComputeInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateComputeInstanceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateComputeInstanceResponseBodyData) SetInstanceId(v string) *CreateComputeInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateComputeInstanceResponseBodyData) SetOrderId(v string) *CreateComputeInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateComputeInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
