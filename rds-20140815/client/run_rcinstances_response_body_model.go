// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdSets(v *RunRCInstancesResponseBodyInstanceIdSets) *RunRCInstancesResponseBody
	GetInstanceIdSets() *RunRCInstancesResponseBodyInstanceIdSets
	SetOrderId(v string) *RunRCInstancesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RunRCInstancesResponseBody
	GetRequestId() *string
}

type RunRCInstancesResponseBody struct {
	InstanceIdSets *RunRCInstancesResponseBodyInstanceIdSets `json:"InstanceIdSets,omitempty" xml:"InstanceIdSets,omitempty" type:"Struct"`
	// The order ID.
	//
	// example:
	//
	// 237850846720798
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 535BD857-E88F-5B4F-A18C-FAF59A74741F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunRCInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RunRCInstancesResponseBody) GetInstanceIdSets() *RunRCInstancesResponseBodyInstanceIdSets {
	return s.InstanceIdSets
}

func (s *RunRCInstancesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RunRCInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunRCInstancesResponseBody) SetInstanceIdSets(v *RunRCInstancesResponseBodyInstanceIdSets) *RunRCInstancesResponseBody {
	s.InstanceIdSets = v
	return s
}

func (s *RunRCInstancesResponseBody) SetOrderId(v string) *RunRCInstancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RunRCInstancesResponseBody) SetRequestId(v string) *RunRCInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunRCInstancesResponseBody) Validate() error {
	if s.InstanceIdSets != nil {
		if err := s.InstanceIdSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunRCInstancesResponseBodyInstanceIdSets struct {
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" xml:"InstanceIdSet,omitempty" type:"Repeated"`
}

func (s RunRCInstancesResponseBodyInstanceIdSets) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesResponseBodyInstanceIdSets) GoString() string {
	return s.String()
}

func (s *RunRCInstancesResponseBodyInstanceIdSets) GetInstanceIdSet() []*string {
	return s.InstanceIdSet
}

func (s *RunRCInstancesResponseBodyInstanceIdSets) SetInstanceIdSet(v []*string) *RunRCInstancesResponseBodyInstanceIdSets {
	s.InstanceIdSet = v
	return s
}

func (s *RunRCInstancesResponseBodyInstanceIdSets) Validate() error {
	return dara.Validate(s)
}
