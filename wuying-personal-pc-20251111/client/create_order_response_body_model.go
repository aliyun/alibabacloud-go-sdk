// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateOrderResponseBody
	GetCode() *string
	SetData(v *CreateOrderResponseBodyData) *CreateOrderResponseBody
	GetData() *CreateOrderResponseBodyData
	SetMessage(v string) *CreateOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOrderResponseBody
	GetRequestId() *string
	SetTraceId(v string) *CreateOrderResponseBody
	GetTraceId() *string
}

type CreateOrderResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string                      `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateOrderResponseBody) GetData() *CreateOrderResponseBodyData {
	return s.Data
}

func (s *CreateOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateOrderResponseBody) SetCode(v string) *CreateOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOrderResponseBody) SetData(v *CreateOrderResponseBodyData) *CreateOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateOrderResponseBody) SetMessage(v string) *CreateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetTraceId(v string) *CreateOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrderResponseBodyData struct {
	OrderDetailList []*CreateOrderResponseBodyDataOrderDetailList `json:"OrderDetailList,omitempty" xml:"OrderDetailList,omitempty" type:"Repeated"`
	OrderId         *string                                       `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PayLink         *string                                       `json:"PayLink,omitempty" xml:"PayLink,omitempty"`
	ResourceId      *string                                       `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CreateOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyData) GetOrderDetailList() []*CreateOrderResponseBodyDataOrderDetailList {
	return s.OrderDetailList
}

func (s *CreateOrderResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateOrderResponseBodyData) GetPayLink() *string {
	return s.PayLink
}

func (s *CreateOrderResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateOrderResponseBodyData) SetOrderDetailList(v []*CreateOrderResponseBodyDataOrderDetailList) *CreateOrderResponseBodyData {
	s.OrderDetailList = v
	return s
}

func (s *CreateOrderResponseBodyData) SetOrderId(v string) *CreateOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetPayLink(v string) *CreateOrderResponseBodyData {
	s.PayLink = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetResourceId(v string) *CreateOrderResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *CreateOrderResponseBodyData) Validate() error {
	if s.OrderDetailList != nil {
		for _, item := range s.OrderDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrderResponseBodyDataOrderDetailList struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OrderId     *int64    `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateOrderResponseBodyDataOrderDetailList) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBodyDataOrderDetailList) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyDataOrderDetailList) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateOrderResponseBodyDataOrderDetailList) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateOrderResponseBodyDataOrderDetailList) SetInstanceIds(v []*string) *CreateOrderResponseBodyDataOrderDetailList {
	s.InstanceIds = v
	return s
}

func (s *CreateOrderResponseBodyDataOrderDetailList) SetOrderId(v int64) *CreateOrderResponseBodyDataOrderDetailList {
	s.OrderId = &v
	return s
}

func (s *CreateOrderResponseBodyDataOrderDetailList) Validate() error {
	return dara.Validate(s)
}
