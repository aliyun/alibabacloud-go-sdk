// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessagePushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *OnsMessagePushRequest
	GetClientId() *string
	SetGroupId(v string) *OnsMessagePushRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsMessagePushRequest
	GetInstanceId() *string
	SetMsgId(v string) *OnsMessagePushRequest
	GetMsgId() *string
	SetTopic(v string) *OnsMessagePushRequest
	GetTopic() *string
}

type OnsMessagePushRequest struct {
	// The ID of the consumer client. You can call the [OnsConsumerGetConnection](https://help.aliyun.com/document_detail/29598.html) operation to query client IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30.5.121.**@24813#-1999745829#-1737591554#453111174894656
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the consumer group. For information about what a consumer group is, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance to which the specified consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0BC1669963053CF68F733BB70396****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The topic to which the message is pushed.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessagePushRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePushRequest) GoString() string {
	return s.String()
}

func (s *OnsMessagePushRequest) GetClientId() *string {
	return s.ClientId
}

func (s *OnsMessagePushRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsMessagePushRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessagePushRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsMessagePushRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessagePushRequest) SetClientId(v string) *OnsMessagePushRequest {
	s.ClientId = &v
	return s
}

func (s *OnsMessagePushRequest) SetGroupId(v string) *OnsMessagePushRequest {
	s.GroupId = &v
	return s
}

func (s *OnsMessagePushRequest) SetInstanceId(v string) *OnsMessagePushRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessagePushRequest) SetMsgId(v string) *OnsMessagePushRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessagePushRequest) SetTopic(v string) *OnsMessagePushRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessagePushRequest) Validate() error {
	return dara.Validate(s)
}
