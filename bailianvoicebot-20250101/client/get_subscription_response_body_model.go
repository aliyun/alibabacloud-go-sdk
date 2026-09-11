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
	// The API status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The subscription information.
	Data *GetSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance llm-zzu528i29ecnprcl does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of dynamic error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
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
	// Indicates whether the subscription is disabled. Valid values:
	//
	// - 0: enabled.
	//
	// - 1: disabled.
	//
	// example:
	//
	// 0
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The endpoint. This parameter is required when MqType is set to ROCKET_MQ_4 or ROCKET_MQ_5.
	//
	// example:
	//
	// rmq-cn-l4p89zajz67.cn-hangzhou.rmq.aliyuncs.com:8080
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The list of subscription events.
	EventList []*GetSubscriptionResponseBodyDataEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// The Bailian business workspace ID.
	//
	// example:
	//
	// llm-3pptowd2olrctsvc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The MQ instance ID. This parameter is required when MqType is set to ROCKET_MQ_5.
	//
	// example:
	//
	// rmq-cn-l4p89zajz67.cn
	MqInstanceId *string `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	// The message push type.
	//
	// example:
	//
	// ROCKET_MQ_4
	MqType *string `json:"MqType,omitempty" xml:"MqType,omitempty"`
	// The password. This parameter is required when MqType is set to ROCKET_MQ_5.
	//
	// example:
	//
	// pwd
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The producer ID. This parameter is required when MqType is set to ROCKET_MQ_4.
	//
	// example:
	//
	// PID_Webhook
	ProducerId *string `json:"ProducerId,omitempty" xml:"ProducerId,omitempty"`
	// The queue topic. This parameter is required when MqType is set to ROCKET_MQ_4 or ROCKET_MQ_5.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The username. This parameter is required when MqType is set to ROCKET_MQ_5.
	//
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
	// Indicates whether the event is pushed.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The display name of the subscription content.
	//
	// example:
	//
	// Call Ended
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The subscription content.
	//
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
