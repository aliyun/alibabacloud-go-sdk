// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageGetByMsgIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsMessageGetByMsgIdRequest
	GetInstanceId() *string
	SetMsgId(v string) *OnsMessageGetByMsgIdRequest
	GetMsgId() *string
	SetTopic(v string) *OnsMessageGetByMsgIdRequest
	GetTopic() *string
}

type OnsMessageGetByMsgIdRequest struct {
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C05F2*****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The topic that contains the message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageGetByMsgIdRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessageGetByMsgIdRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsMessageGetByMsgIdRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessageGetByMsgIdRequest) SetInstanceId(v string) *OnsMessageGetByMsgIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByMsgIdRequest) SetMsgId(v string) *OnsMessageGetByMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessageGetByMsgIdRequest) SetTopic(v string) *OnsMessageGetByMsgIdRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessageGetByMsgIdRequest) Validate() error {
	return dara.Validate(s)
}
