// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSubscriptionResponseBody
	GetCode() *string
	SetData(v *GetSubscriptionResponseBodyData) *GetSubscriptionResponseBody
	GetData() *GetSubscriptionResponseBodyData
	SetHttpStatusCode(v int32) *GetSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSubscriptionResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetSubscriptionResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetSubscriptionResponseBody
	GetRequestId() *string
}

type GetSubscriptionResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSubscriptionResponseBody) GetData() *GetSubscriptionResponseBodyData {
	return s.Data
}

func (s *GetSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubscriptionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubscriptionResponseBody) SetCode(v string) *GetSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetData(v *GetSubscriptionResponseBodyData) *GetSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionResponseBody) SetHttpStatusCode(v int32) *GetSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetMessage(v string) *GetSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetParams(v []*string) *GetSubscriptionResponseBody {
	s.Params = v
	return s
}

func (s *GetSubscriptionResponseBody) SetRequestId(v string) *GetSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSubscriptionResponseBodyData struct {
	// example:
	//
	// 0
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// example:
	//
	// rmq-cn-l4p89zajz67.cn-hangzhou.rmq.aliyuncs.com:8080
	Endpoint  *string                                     `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EventList []*GetSubscriptionResponseBodyDataEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rmq-cn-l4p89zajz67.cn
	MqInstanceId *string `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	// example:
	//
	// ROCKET_MQ_4
	MqType *string `json:"MqType,omitempty" xml:"MqType,omitempty"`
	// example:
	//
	// pwd
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// user1
	ProducerId *string `json:"ProducerId,omitempty" xml:"ProducerId,omitempty"`
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetSubscriptionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyData) GetDisabled() *bool {
	return s.Disabled
}

func (s *GetSubscriptionResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetSubscriptionResponseBodyData) GetEventList() []*GetSubscriptionResponseBodyDataEventList {
	return s.EventList
}

func (s *GetSubscriptionResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSubscriptionResponseBodyData) GetMqInstanceId() *string {
	return s.MqInstanceId
}

func (s *GetSubscriptionResponseBodyData) GetMqType() *string {
	return s.MqType
}

func (s *GetSubscriptionResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *GetSubscriptionResponseBodyData) GetProducerId() *string {
	return s.ProducerId
}

func (s *GetSubscriptionResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *GetSubscriptionResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetSubscriptionResponseBodyData) SetDisabled(v bool) *GetSubscriptionResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetEndpoint(v string) *GetSubscriptionResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetEventList(v []*GetSubscriptionResponseBodyDataEventList) *GetSubscriptionResponseBodyData {
	s.EventList = v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetInstanceId(v string) *GetSubscriptionResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetMqInstanceId(v string) *GetSubscriptionResponseBodyData {
	s.MqInstanceId = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetMqType(v string) *GetSubscriptionResponseBodyData {
	s.MqType = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetPassword(v string) *GetSubscriptionResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetProducerId(v string) *GetSubscriptionResponseBodyData {
	s.ProducerId = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetTopic(v string) *GetSubscriptionResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetUserName(v string) *GetSubscriptionResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) Validate() error {
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

type GetSubscriptionResponseBodyDataEventList struct {
	// example:
	//
	// false
	Disabled    *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// Released
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSubscriptionResponseBodyDataEventList) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionResponseBodyDataEventList) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyDataEventList) GetDisabled() *bool {
	return s.Disabled
}

func (s *GetSubscriptionResponseBodyDataEventList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetSubscriptionResponseBodyDataEventList) GetName() *string {
	return s.Name
}

func (s *GetSubscriptionResponseBodyDataEventList) SetDisabled(v bool) *GetSubscriptionResponseBodyDataEventList {
	s.Disabled = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataEventList) SetDisplayName(v string) *GetSubscriptionResponseBodyDataEventList {
	s.DisplayName = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataEventList) SetName(v string) *GetSubscriptionResponseBodyDataEventList {
	s.Name = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataEventList) Validate() error {
	return dara.Validate(s)
}
