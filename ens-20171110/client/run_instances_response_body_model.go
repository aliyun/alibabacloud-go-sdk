// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *RunInstancesResponseBody
	GetInstanceIds() []*string
	SetOrderId(v string) *RunInstancesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RunInstancesResponseBody
	GetRequestId() *string
}

type RunInstancesResponseBody struct {
	// The IDs of instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The order ID.
	//
	// example:
	//
	// 213177957850399
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91E4AFBE-4E35-5D2A-A886-BB477C9953D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RunInstancesResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RunInstancesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RunInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunInstancesResponseBody) SetInstanceIds(v []*string) *RunInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *RunInstancesResponseBody) SetOrderId(v string) *RunInstancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RunInstancesResponseBody) SetRequestId(v string) *RunInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
