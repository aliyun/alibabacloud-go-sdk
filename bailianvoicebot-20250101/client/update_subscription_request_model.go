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
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// rmq-cn-l4p89zajz67.cn-hangzhou.rmq.aliyuncs.com:8080
	Endpoint           *string   `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EventSubscriptions []*string `json:"EventSubscriptions,omitempty" xml:"EventSubscriptions,omitempty" type:"Repeated"`
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
