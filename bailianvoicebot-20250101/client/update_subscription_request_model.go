// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *UpdateSubscriptionRequest
	GetBusinessUnitId() *string
	SetEndpoint(v string) *UpdateSubscriptionRequest
	GetEndpoint() *string
	SetEventSubscriptions(v []*string) *UpdateSubscriptionRequest
	GetEventSubscriptions() []*string
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
	EventSubscriptions []*string `json:"EventSubscriptions,omitempty" xml:"EventSubscriptions,omitempty" type:"Repeated"`
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

func (s UpdateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *UpdateSubscriptionRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateSubscriptionRequest) GetEventSubscriptions() []*string {
	return s.EventSubscriptions
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

func (s *UpdateSubscriptionRequest) SetBusinessUnitId(v string) *UpdateSubscriptionRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetEndpoint(v string) *UpdateSubscriptionRequest {
	s.Endpoint = &v
	return s
}

func (s *UpdateSubscriptionRequest) SetEventSubscriptions(v []*string) *UpdateSubscriptionRequest {
	s.EventSubscriptions = v
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
