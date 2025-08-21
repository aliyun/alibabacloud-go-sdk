// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateARMServerInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *CreateARMServerInstancesResponseBody
	GetInstanceIds() []*string
	SetOrderId(v string) *CreateARMServerInstancesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateARMServerInstancesResponseBody
	GetRequestId() *string
}

type CreateARMServerInstancesResponseBody struct {
	// The IDs of instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// 212630314490***
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateARMServerInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateARMServerInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateARMServerInstancesResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateARMServerInstancesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateARMServerInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateARMServerInstancesResponseBody) SetInstanceIds(v []*string) *CreateARMServerInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateARMServerInstancesResponseBody) SetOrderId(v string) *CreateARMServerInstancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateARMServerInstancesResponseBody) SetRequestId(v string) *CreateARMServerInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateARMServerInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
