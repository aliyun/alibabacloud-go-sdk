// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsMessageTraceRequest
	GetInstanceId() *string
	SetMsgId(v string) *OnsMessageTraceRequest
	GetMsgId() *string
	SetTopic(v string) *OnsMessageTraceRequest
	GetTopic() *string
}

type OnsMessageTraceRequest struct {
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
	// 1E05791C117818B4AAC23B1BB0CE****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The topic to which the message belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageTraceRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessageTraceRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsMessageTraceRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessageTraceRequest) SetInstanceId(v string) *OnsMessageTraceRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageTraceRequest) SetMsgId(v string) *OnsMessageTraceRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessageTraceRequest) SetTopic(v string) *OnsMessageTraceRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessageTraceRequest) Validate() error {
	return dara.Validate(s)
}
