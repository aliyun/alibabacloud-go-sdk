// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *UpdateSubscriptionShrinkRequest
	GetEndpoint() *string
	SetEventSubscriptionsShrink(v string) *UpdateSubscriptionShrinkRequest
	GetEventSubscriptionsShrink() *string
	SetInstanceId(v string) *UpdateSubscriptionShrinkRequest
	GetInstanceId() *string
	SetMqInstanceId(v string) *UpdateSubscriptionShrinkRequest
	GetMqInstanceId() *string
	SetMqType(v string) *UpdateSubscriptionShrinkRequest
	GetMqType() *string
	SetPassword(v string) *UpdateSubscriptionShrinkRequest
	GetPassword() *string
	SetProducerId(v string) *UpdateSubscriptionShrinkRequest
	GetProducerId() *string
	SetTopic(v string) *UpdateSubscriptionShrinkRequest
	GetTopic() *string
	SetUserName(v string) *UpdateSubscriptionShrinkRequest
	GetUserName() *string
}

type UpdateSubscriptionShrinkRequest struct {
	// The endpoint.
	//
	// example:
	//
	// rmq-cn-h964u01wh12.cn-hangzhou.rmq.aliyuncs.com:8080
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The list of subscription items.
	EventSubscriptionsShrink *string `json:"EventSubscriptions,omitempty" xml:"EventSubscriptions,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance ID of the message queue.
	//
	// example:
	//
	// rmq-cn-3g84vpf3712
	MqInstanceId *string `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	// The MSMQ type.
	//
	// example:
	//
	// ROCKET_MQ_4
	MqType *string `json:"MqType,omitempty" xml:"MqType,omitempty"`
	// The password.
	//
	// example:
	//
	// pa44w0rd
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The producer ID.
	//
	// example:
	//
	// GID_123456
	ProducerId *string `json:"ProducerId,omitempty" xml:"ProducerId,omitempty"`
	// The topic.
	//
	// example:
	//
	// OUTBOUND_BOT_TOPIC
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The username.
	//
	// example:
	//
	// admin
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateSubscriptionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionShrinkRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateSubscriptionShrinkRequest) GetEventSubscriptionsShrink() *string {
	return s.EventSubscriptionsShrink
}

func (s *UpdateSubscriptionShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateSubscriptionShrinkRequest) GetMqInstanceId() *string {
	return s.MqInstanceId
}

func (s *UpdateSubscriptionShrinkRequest) GetMqType() *string {
	return s.MqType
}

func (s *UpdateSubscriptionShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateSubscriptionShrinkRequest) GetProducerId() *string {
	return s.ProducerId
}

func (s *UpdateSubscriptionShrinkRequest) GetTopic() *string {
	return s.Topic
}

func (s *UpdateSubscriptionShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateSubscriptionShrinkRequest) SetEndpoint(v string) *UpdateSubscriptionShrinkRequest {
	s.Endpoint = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetEventSubscriptionsShrink(v string) *UpdateSubscriptionShrinkRequest {
	s.EventSubscriptionsShrink = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetInstanceId(v string) *UpdateSubscriptionShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetMqInstanceId(v string) *UpdateSubscriptionShrinkRequest {
	s.MqInstanceId = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetMqType(v string) *UpdateSubscriptionShrinkRequest {
	s.MqType = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetPassword(v string) *UpdateSubscriptionShrinkRequest {
	s.Password = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetProducerId(v string) *UpdateSubscriptionShrinkRequest {
	s.ProducerId = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetTopic(v string) *UpdateSubscriptionShrinkRequest {
	s.Topic = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetUserName(v string) *UpdateSubscriptionShrinkRequest {
	s.UserName = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
