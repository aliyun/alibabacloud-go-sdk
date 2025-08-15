// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxSendTps(v int64) *CreateTopicRequest
	GetMaxSendTps() *int64
	SetMessageType(v string) *CreateTopicRequest
	GetMessageType() *string
	SetRemark(v string) *CreateTopicRequest
	GetRemark() *string
}

type CreateTopicRequest struct {
	// The maximum TPS for message sending.
	//
	// example:
	//
	// 1500
	MaxSendTps *int64 `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// The type of messages in the topic that you want to create.
	//
	// Valid values:
	//
	// 	- TRANSACTION: transactional messages
	//
	// 	- FIFO: ordered messages
	//
	// 	- DELAY: scheduled or delayed messages
	//
	// 	- NORMAL: normal messages
	//
	// >  The type of messages in the topic must be the same as the type of messages that you want to send. For example, if you create a topic whose message type is ordered messages, you can use the topic to send and receive only ordered messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The description of the topic that you want to create.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) GetMaxSendTps() *int64 {
	return s.MaxSendTps
}

func (s *CreateTopicRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *CreateTopicRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateTopicRequest) SetMaxSendTps(v int64) *CreateTopicRequest {
	s.MaxSendTps = &v
	return s
}

func (s *CreateTopicRequest) SetMessageType(v string) *CreateTopicRequest {
	s.MessageType = &v
	return s
}

func (s *CreateTopicRequest) SetRemark(v string) *CreateTopicRequest {
	s.Remark = &v
	return s
}

func (s *CreateTopicRequest) Validate() error {
	return dara.Validate(s)
}
