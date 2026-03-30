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
	// example:
	//
	// rmq-cn-l4p89zajz67.cn-hangzhou.rmq.aliyuncs.com:8080
	Endpoint                 *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EventSubscriptionsShrink *string `json:"EventSubscriptions,omitempty" xml:"EventSubscriptions,omitempty"`
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
