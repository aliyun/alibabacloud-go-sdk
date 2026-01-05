// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCapacityEachType(v map[string]interface{}) *QueryStatisticResponseBody
	GetInstanceCapacityEachType() map[string]interface{}
	SetInstanceNumEachType(v map[string]interface{}) *QueryStatisticResponseBody
	GetInstanceNumEachType() map[string]interface{}
	SetRequestId(v string) *QueryStatisticResponseBody
	GetRequestId() *string
	SetSlotNumEachType(v map[string]interface{}) *QueryStatisticResponseBody
	GetSlotNumEachType() map[string]interface{}
}

type QueryStatisticResponseBody struct {
	InstanceCapacityEachType map[string]interface{} `json:"InstanceCapacityEachType,omitempty" xml:"InstanceCapacityEachType,omitempty"`
	InstanceNumEachType      map[string]interface{} `json:"InstanceNumEachType,omitempty" xml:"InstanceNumEachType,omitempty"`
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId       *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlotNumEachType map[string]interface{} `json:"SlotNumEachType,omitempty" xml:"SlotNumEachType,omitempty"`
}

func (s QueryStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStatisticResponseBody) GetInstanceCapacityEachType() map[string]interface{} {
	return s.InstanceCapacityEachType
}

func (s *QueryStatisticResponseBody) GetInstanceNumEachType() map[string]interface{} {
	return s.InstanceNumEachType
}

func (s *QueryStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryStatisticResponseBody) GetSlotNumEachType() map[string]interface{} {
	return s.SlotNumEachType
}

func (s *QueryStatisticResponseBody) SetInstanceCapacityEachType(v map[string]interface{}) *QueryStatisticResponseBody {
	s.InstanceCapacityEachType = v
	return s
}

func (s *QueryStatisticResponseBody) SetInstanceNumEachType(v map[string]interface{}) *QueryStatisticResponseBody {
	s.InstanceNumEachType = v
	return s
}

func (s *QueryStatisticResponseBody) SetRequestId(v string) *QueryStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStatisticResponseBody) SetSlotNumEachType(v map[string]interface{}) *QueryStatisticResponseBody {
	s.SlotNumEachType = v
	return s
}

func (s *QueryStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
