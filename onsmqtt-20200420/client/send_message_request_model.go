// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SendMessageRequest
	GetInstanceId() *string
	SetMqttTopic(v string) *SendMessageRequest
	GetMqttTopic() *string
	SetPayload(v string) *SendMessageRequest
	GetPayload() *string
}

type SendMessageRequest struct {
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can view the instance ID in the **Basic Information*	- section on the **Instance Details*	- page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com).
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The topic to which you want to send a message on the ApsaraMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// TopicA
	MqttTopic *string `json:"MqttTopic,omitempty" xml:"MqttTopic,omitempty"`
	// The message content, which is the payload of the message. We recommend that you encode the content in Base64 to prevent non-printable characters from being transmitted.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s SendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendMessageRequest) GetMqttTopic() *string {
	return s.MqttTopic
}

func (s *SendMessageRequest) GetPayload() *string {
	return s.Payload
}

func (s *SendMessageRequest) SetInstanceId(v string) *SendMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *SendMessageRequest) SetMqttTopic(v string) *SendMessageRequest {
	s.MqttTopic = &v
	return s
}

func (s *SendMessageRequest) SetPayload(v string) *SendMessageRequest {
	s.Payload = &v
	return s
}

func (s *SendMessageRequest) Validate() error {
	return dara.Validate(s)
}
