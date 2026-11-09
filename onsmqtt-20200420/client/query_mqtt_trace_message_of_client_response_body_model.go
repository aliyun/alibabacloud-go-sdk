// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessageOfClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryMqttTraceMessageOfClientResponseBody
	GetCurrentPage() *int32
	SetMessageOfClientList(v []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) *QueryMqttTraceMessageOfClientResponseBody
	GetMessageOfClientList() []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList
	SetPageSize(v int32) *QueryMqttTraceMessageOfClientResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryMqttTraceMessageOfClientResponseBody
	GetRequestId() *string
	SetTotal(v int64) *QueryMqttTraceMessageOfClientResponseBody
	GetTotal() *int64
}

type QueryMqttTraceMessageOfClientResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned messages.
	MessageOfClientList []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList `json:"MessageOfClientList,omitempty" xml:"MessageOfClientList,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID. You can use the ID to troubleshoot issues. This parameter is a common parameter.
	//
	// example:
	//
	// B096B9D6-62F3-4567-BB59-58D1362E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of messages returned.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMqttTraceMessageOfClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMqttTraceMessageOfClientResponseBody) GetMessageOfClientList() []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	return s.MessageOfClientList
}

func (s *QueryMqttTraceMessageOfClientResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMqttTraceMessageOfClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMqttTraceMessageOfClientResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetCurrentPage(v int32) *QueryMqttTraceMessageOfClientResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetMessageOfClientList(v []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) *QueryMqttTraceMessageOfClientResponseBody {
	s.MessageOfClientList = v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetPageSize(v int32) *QueryMqttTraceMessageOfClientResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetRequestId(v string) *QueryMqttTraceMessageOfClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetTotal(v int64) *QueryMqttTraceMessageOfClientResponseBody {
	s.Total = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) Validate() error {
	if s.MessageOfClientList != nil {
		for _, item := range s.MessageOfClientList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList struct {
	// The action on the message. Valid values:
	//
	// 	- **pub_mqtt**: The ApsaraMQ for MQTT client sends the message.
	//
	// 	- **sub**: The ApsaraMQ for MQTT client subscribes to the message.
	//
	// 	- **push_offline**: The ApsaraMQ for MQTT broker pushes the offline message to the ApsaraMQ for MQTT client.
	//
	// example:
	//
	// pub_mqtt
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The returned code for the action on the message. Valid values:
	//
	// 	- **mqtt.trace.action.msg.pub.mqtt**: This value is returned if the value of Action is **pub_mqtt**.
	//
	// 	- **mqtt.trace.action.msg.sub**: This value is returned if the value of Action is **sub**.
	//
	// 	- **mqtt.trace.action.msg.push.offline**: This value is returned if the value of Action is **push_offline**.
	//
	// example:
	//
	// mqtt.trace.action.msg.pub.mqtt
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The information returned for the action on the message. Valid values:
	//
	// 	- **Pub From Mqtt Client**: This value is returned if the value of Action is **pub_mqtt**.
	//
	// 	- **Push To Mqtt Client**: This value is returned if the value of Action is **sub**.
	//
	// 	- **Push Offline Msg To Mqtt Client**: This value is returned if the value of Action is **push_offline**.
	//
	// example:
	//
	// Pub From Mqtt Client
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The client ID of the device.
	//
	// example:
	//
	// GID_test@@@producer
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The message ID.
	//
	// example:
	//
	// AC1EC0030EAB78308DB16A3EC773****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The time when the message was sent or received.
	//
	// example:
	//
	// 2021-05-21 15:08:19.234
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GetAction() *string {
	return s.Action
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GetActionCode() *string {
	return s.ActionCode
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GetActionInfo() *string {
	return s.ActionInfo
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GetClientId() *string {
	return s.ClientId
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GetMsgId() *string {
	return s.MsgId
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GetTime() *string {
	return s.Time
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetAction(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetActionCode(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetActionInfo(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.ActionInfo = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetClientId(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetMsgId(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetTime(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.Time = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) Validate() error {
	return dara.Validate(s)
}
