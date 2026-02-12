// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsMessageDetailRequest
	GetInstanceId() *string
	SetMsgId(v string) *OnsMessageDetailRequest
	GetMsgId() *string
	SetTopic(v string) *OnsMessageDetailRequest
	GetTopic() *string
}

type OnsMessageDetailRequest struct {
	// The ID of the ApsaraMQ for RocketMQ Instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_184681981******_BXig0x6A
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C0******
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The name of the topic in which the message you want to query is stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageDetailRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessageDetailRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsMessageDetailRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessageDetailRequest) SetInstanceId(v string) *OnsMessageDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageDetailRequest) SetMsgId(v string) *OnsMessageDetailRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessageDetailRequest) SetTopic(v string) *OnsMessageDetailRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessageDetailRequest) Validate() error {
	return dara.Validate(s)
}
