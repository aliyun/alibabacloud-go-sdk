// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiteTopicExpiration(v int64) *CreateTopicRequest
	GetLiteTopicExpiration() *int64
	SetMaxSendTps(v int64) *CreateTopicRequest
	GetMaxSendTps() *int64
	SetMessageType(v string) *CreateTopicRequest
	GetMessageType() *string
	SetRemark(v string) *CreateTopicRequest
	GetRemark() *string
}

type CreateTopicRequest struct {
	// The expiration time of the lite topic, in minutes.
	//
	// Valid values: 30 to 720.
	//
	// Default value: -1, which indicates the topic never expires.
	//
	// example:
	//
	// 60
	LiteTopicExpiration *int64 `json:"liteTopicExpiration,omitempty" xml:"liteTopicExpiration,omitempty"`
	// The maximum transactions per second (TPS) for sending messages.
	//
	// example:
	//
	// 1500
	MaxSendTps *int64 `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// The message type of the topic to be created.
	//
	// 	Notice: The type of the topic must be the same as the type of messages sent to the topic. For example, if you create a topic of the ordered message type, the topic can be used only to send and receive ordered messages, not other types of messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The remarks of the topic to be created.
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

func (s *CreateTopicRequest) GetLiteTopicExpiration() *int64 {
	return s.LiteTopicExpiration
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

func (s *CreateTopicRequest) SetLiteTopicExpiration(v int64) *CreateTopicRequest {
	s.LiteTopicExpiration = &v
	return s
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
