// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessagePublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageTraceLists(v []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) *QueryMqttTraceMessagePublishResponseBody
	GetMessageTraceLists() []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists
	SetRequestId(v string) *QueryMqttTraceMessagePublishResponseBody
	GetRequestId() *string
}

type QueryMqttTraceMessagePublishResponseBody struct {
	// The message traces.
	MessageTraceLists []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists `json:"MessageTraceLists,omitempty" xml:"MessageTraceLists,omitempty" type:"Repeated"`
	// The request ID. You can use the ID to troubleshoot issues. This parameter is a common parameter.
	//
	// example:
	//
	// 69AD5550-BF22-438A-9202-A6E89185****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMqttTraceMessagePublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessagePublishResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishResponseBody) GetMessageTraceLists() []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	return s.MessageTraceLists
}

func (s *QueryMqttTraceMessagePublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMqttTraceMessagePublishResponseBody) SetMessageTraceLists(v []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) *QueryMqttTraceMessagePublishResponseBody {
	s.MessageTraceLists = v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBody) SetRequestId(v string) *QueryMqttTraceMessagePublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBody) Validate() error {
	if s.MessageTraceLists != nil {
		for _, item := range s.MessageTraceLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMqttTraceMessagePublishResponseBodyMessageTraceLists struct {
	// The action on the message. Valid values:
	//
	// 	- **pub_mqtt**: indicates that the message was sent by an ApsaraMQ for MQTT client.
	//
	// 	- **pub_mq**: indicates that the message was sent by an ApsaraMQ for RocketMQ client.
	//
	// example:
	//
	// pub_mqtt
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The returned code for the action on the message. Valid values:
	//
	// 	- **mqtt.trace.action.msg.pub.mqtt**: This value is returned if the value of Action is **pub_mqtt**.
	//
	// 	- **mqtt.trace.action.msg.pub.mq**: This value is returned if the value of Action is **pub_mq**.
	//
	// example:
	//
	// mqtt.trace.action.msg.pub.mqtt
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The returned information for the action on the message. Valid values:
	//
	// 	- **Pub From Mqtt Client**: This value is returned if the value of Action is **pub_mqtt**.
	//
	// 	- **Pub From MQ**: This value is returned if the value of Action is **pub_mq**.
	//
	// example:
	//
	// Pub From Mqtt Client
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The ID of the client that sends the message.
	//
	// example:
	//
	// GID_test@@@producer
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The message ID.
	//
	// example:
	//
	// AC1EC0030EAB78308DB16A3EC773BD95
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The time when the message was sent.
	//
	// example:
	//
	// 2021-05-21 15:08:19.210
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GetAction() *string {
	return s.Action
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GetActionCode() *string {
	return s.ActionCode
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GetActionInfo() *string {
	return s.ActionInfo
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GetClientId() *string {
	return s.ClientId
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GetMsgId() *string {
	return s.MsgId
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GetTime() *string {
	return s.Time
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetAction(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetActionCode(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetActionInfo(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.ActionInfo = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetClientId(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetMsgId(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetTime(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.Time = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) Validate() error {
	return dara.Validate(s)
}
