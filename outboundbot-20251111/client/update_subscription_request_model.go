// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *UpdateSubscriptionRequest
	GetEndpoint() *string
	SetEventSubscriptions(v []*string) *UpdateSubscriptionRequest
	GetEventSubscriptions() []*string
	SetInstanceId(v string) *UpdateSubscriptionRequest
	GetInstanceId() *string
	SetMqInstanceId(v string) *UpdateSubscriptionRequest
	GetMqInstanceId() *string
	SetMqType(v string) *UpdateSubscriptionRequest
	GetMqType() *string
	SetPassword(v string) *UpdateSubscriptionRequest
	GetPassword() *string
	SetProducerId(v string) *UpdateSubscriptionRequest
	GetProducerId() *string
	SetTopic(v string) *UpdateSubscriptionRequest
	GetTopic() *string
	SetUserName(v string) *UpdateSubscriptionRequest
	GetUserName() *string
}

type UpdateSubscriptionRequest struct {
	// The endpoint.
	//
	// example:
	//
	// rmq-cn-h964u01wh12.cn-hangzhou.rmq.aliyuncs.com:8080
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The list of subscription items.
	EventSubscriptions []*string `json:"EventSubscriptions,omitempty" xml:"EventSubscriptions,omitempty" type:"Repeated"`
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

func (s UpdateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateSubscriptionRequest) GetEventSubscriptions() []*string {
	return s.EventSubscriptions
}

func (s *UpdateSubscriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateSubscriptionRequest) GetMqInstanceId() *string {
	return s.MqInstanceId
}

func (s *UpdateSubscriptionRequest) GetMqType() *string {
	return s.MqType
}

func (s *UpdateSubscriptionRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateSubscriptionRequest) GetProducerId() *string {
	return s.ProducerId
}

func (s *UpdateSubscriptionRequest) GetTopic() *string {
	return s.Topic
}

func (s *UpdateSubscriptionRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateSubscriptionRequest) SetEndpoint(v string) *UpdateSubscriptionRequest {
	s.Endpoint = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetEventSubscriptions(v []*string) *UpdateSubscriptionRequest {
	s.EventSubscriptions = v
	return s
}

func (s *UpdateSubscriptionRequest) SetInstanceId(v string) *UpdateSubscriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetMqInstanceId(v string) *UpdateSubscriptionRequest {
	s.MqInstanceId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetMqType(v string) *UpdateSubscriptionRequest {
	s.MqType = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetPassword(v string) *UpdateSubscriptionRequest {
	s.Password = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetProducerId(v string) *UpdateSubscriptionRequest {
	s.ProducerId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetTopic(v string) *UpdateSubscriptionRequest {
	s.Topic = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetUserName(v string) *UpdateSubscriptionRequest {
	s.UserName = &v
	return s
}

func (s *UpdateSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
