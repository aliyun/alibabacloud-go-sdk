// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *UpdateSubscriptionShrinkRequest
	GetBusinessUnitId() *string
	SetEndpoint(v string) *UpdateSubscriptionShrinkRequest
	GetEndpoint() *string
	SetEventSubscriptionsShrink(v string) *UpdateSubscriptionShrinkRequest
	GetEventSubscriptionsShrink() *string
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
	// The ID of the Model Studio business space.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// The endpoint. This parameter is required if `MqType` is set to `ROCKET_MQ_4` or `ROCKET_MQ_5`.
	//
	// example:
	//
	// rmq-cn-l4p89zajz67.cn-hangzhou.rmq.aliyuncs.com:8080
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// A list of events to subscribe to.
	EventSubscriptionsShrink *string `json:"EventSubscriptions,omitempty" xml:"EventSubscriptions,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance. This parameter is required if `MqType` is set to `ROCKET_MQ_5`.
	//
	// example:
	//
	// rmq-cn-l4p89zajz67.cn
	MqInstanceId *string `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	// The type of the message queue service. Valid values are `ROCKET_MQ_4` and `ROCKET_MQ_5`, which correspond to ApsaraMQ for RocketMQ.
	//
	// example:
	//
	// ROCKET_MQ_4
	MqType *string `json:"MqType,omitempty" xml:"MqType,omitempty"`
	// The password for authentication. This parameter is required if `MqType` is set to `ROCKET_MQ_5`.
	//
	// example:
	//
	// pwd
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the producer. This parameter is required if `MqType` is set to `ROCKET_MQ_4`.
	//
	// example:
	//
	// user1
	ProducerId *string `json:"ProducerId,omitempty" xml:"ProducerId,omitempty"`
	// The topic of the queue. This parameter is required if `MqType` is set to `ROCKET_MQ_4` or `ROCKET_MQ_5`.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The username for authentication. This parameter is required if `MqType` is set to `ROCKET_MQ_5`.
	//
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

func (s *UpdateSubscriptionShrinkRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *UpdateSubscriptionShrinkRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateSubscriptionShrinkRequest) GetEventSubscriptionsShrink() *string {
	return s.EventSubscriptionsShrink
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

func (s *UpdateSubscriptionShrinkRequest) SetBusinessUnitId(v string) *UpdateSubscriptionShrinkRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetEndpoint(v string) *UpdateSubscriptionShrinkRequest {
	s.Endpoint = &v
	return s
}

func (s *UpdateSubscriptionShrinkRequest) SetEventSubscriptionsShrink(v string) *UpdateSubscriptionShrinkRequest {
	s.EventSubscriptionsShrink = &v
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
