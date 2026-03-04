// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSubscriptionResponseBody
	GetCode() *string
	SetData(v *UpdateSubscriptionResponseBodyData) *UpdateSubscriptionResponseBody
	GetData() *UpdateSubscriptionResponseBodyData
	SetHttpStatusCode(v int32) *UpdateSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSubscriptionResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateSubscriptionResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateSubscriptionResponseBody
	GetRequestId() *string
}

type UpdateSubscriptionResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSubscriptionResponseBody) GetData() *UpdateSubscriptionResponseBodyData {
	return s.Data
}

func (s *UpdateSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSubscriptionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubscriptionResponseBody) SetCode(v string) *UpdateSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetData(v *UpdateSubscriptionResponseBodyData) *UpdateSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetHttpStatusCode(v int32) *UpdateSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetMessage(v string) *UpdateSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetParams(v []*string) *UpdateSubscriptionResponseBody {
	s.Params = v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetRequestId(v string) *UpdateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSubscriptionResponseBodyData struct {
	// example:
	//
	// rmq-cn-****.cn-shanghai.rmq.aliyuncs.com:8080
	AccessPoint *string                                        `json:"AccessPoint,omitempty" xml:"AccessPoint,omitempty"`
	EventList   []*UpdateSubscriptionResponseBodyDataEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rmq-cn-****
	MqInstanceId *string `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	// example:
	//
	// rocketmq5
	MqType *string `json:"MqType,omitempty" xml:"MqType,omitempty"`
	// example:
	//
	// GID_xxx
	ProducerId *string `json:"ProducerId,omitempty" xml:"ProducerId,omitempty"`
	// example:
	//
	// ccc-event
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateSubscriptionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBodyData) GetAccessPoint() *string {
	return s.AccessPoint
}

func (s *UpdateSubscriptionResponseBodyData) GetEventList() []*UpdateSubscriptionResponseBodyDataEventList {
	return s.EventList
}

func (s *UpdateSubscriptionResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateSubscriptionResponseBodyData) GetMqInstanceId() *string {
	return s.MqInstanceId
}

func (s *UpdateSubscriptionResponseBodyData) GetMqType() *string {
	return s.MqType
}

func (s *UpdateSubscriptionResponseBodyData) GetProducerId() *string {
	return s.ProducerId
}

func (s *UpdateSubscriptionResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *UpdateSubscriptionResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *UpdateSubscriptionResponseBodyData) SetAccessPoint(v string) *UpdateSubscriptionResponseBodyData {
	s.AccessPoint = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyData) SetEventList(v []*UpdateSubscriptionResponseBodyDataEventList) *UpdateSubscriptionResponseBodyData {
	s.EventList = v
	return s
}

func (s *UpdateSubscriptionResponseBodyData) SetInstanceId(v string) *UpdateSubscriptionResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyData) SetMqInstanceId(v string) *UpdateSubscriptionResponseBodyData {
	s.MqInstanceId = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyData) SetMqType(v string) *UpdateSubscriptionResponseBodyData {
	s.MqType = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyData) SetProducerId(v string) *UpdateSubscriptionResponseBodyData {
	s.ProducerId = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyData) SetTopic(v string) *UpdateSubscriptionResponseBodyData {
	s.Topic = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyData) SetUsername(v string) *UpdateSubscriptionResponseBodyData {
	s.Username = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyData) Validate() error {
	if s.EventList != nil {
		for _, item := range s.EventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateSubscriptionResponseBodyDataEventList struct {
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// example:
	//
	// Dialing
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateSubscriptionResponseBodyDataEventList) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionResponseBodyDataEventList) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBodyDataEventList) GetDisabled() *bool {
	return s.Disabled
}

func (s *UpdateSubscriptionResponseBodyDataEventList) GetName() *string {
	return s.Name
}

func (s *UpdateSubscriptionResponseBodyDataEventList) GetTopic() *string {
	return s.Topic
}

func (s *UpdateSubscriptionResponseBodyDataEventList) SetDisabled(v bool) *UpdateSubscriptionResponseBodyDataEventList {
	s.Disabled = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyDataEventList) SetName(v string) *UpdateSubscriptionResponseBodyDataEventList {
	s.Name = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyDataEventList) SetTopic(v string) *UpdateSubscriptionResponseBodyDataEventList {
	s.Topic = &v
	return s
}

func (s *UpdateSubscriptionResponseBodyDataEventList) Validate() error {
	return dara.Validate(s)
}
