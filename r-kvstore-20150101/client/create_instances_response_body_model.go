// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v *CreateInstancesResponseBodyInstanceIds) *CreateInstancesResponseBody
	GetInstanceIds() *CreateInstancesResponseBodyInstanceIds
	SetOrderId(v string) *CreateInstancesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateInstancesResponseBody
	GetRequestId() *string
}

type CreateInstancesResponseBody struct {
	// The IDs of instances that were created.
	InstanceIds *CreateInstancesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The ID of the order.
	//
	// example:
	//
	// 20905403119****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 561AFBF1-BE20-44DB-9BD1-6988B53E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstancesResponseBody) GetInstanceIds() *CreateInstancesResponseBodyInstanceIds {
	return s.InstanceIds
}

func (s *CreateInstancesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstancesResponseBody) SetInstanceIds(v *CreateInstancesResponseBodyInstanceIds) *CreateInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateInstancesResponseBody) SetOrderId(v string) *CreateInstancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateInstancesResponseBody) SetRequestId(v string) *CreateInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateInstancesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s CreateInstancesResponseBodyInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateInstancesResponseBodyInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *CreateInstancesResponseBodyInstanceIds) SetInstanceId(v []*string) *CreateInstancesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

func (s *CreateInstancesResponseBodyInstanceIds) Validate() error {
	return dara.Validate(s)
}
