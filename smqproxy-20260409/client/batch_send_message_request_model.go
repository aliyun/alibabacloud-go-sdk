// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*BatchSendMessageRequestMessages) *BatchSendMessageRequest
	GetMessages() []*BatchSendMessageRequestMessages
}

type BatchSendMessageRequest struct {
	Messages []*BatchSendMessageRequestMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
}

func (s BatchSendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMessageRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMessageRequest) GetMessages() []*BatchSendMessageRequestMessages {
	return s.Messages
}

func (s *BatchSendMessageRequest) SetMessages(v []*BatchSendMessageRequestMessages) *BatchSendMessageRequest {
	s.Messages = v
	return s
}

func (s *BatchSendMessageRequest) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchSendMessageRequestMessages struct {
	// example:
	//
	// 60
	DelaySeconds *int32 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// example:
	//
	// This is test message 1.
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// example:
	//
	// group1
	MessageGroupId *string `json:"MessageGroupId,omitempty" xml:"MessageGroupId,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s BatchSendMessageRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMessageRequestMessages) GoString() string {
	return s.String()
}

func (s *BatchSendMessageRequestMessages) GetDelaySeconds() *int32 {
	return s.DelaySeconds
}

func (s *BatchSendMessageRequestMessages) GetMessageBody() *string {
	return s.MessageBody
}

func (s *BatchSendMessageRequestMessages) GetMessageGroupId() *string {
	return s.MessageGroupId
}

func (s *BatchSendMessageRequestMessages) GetPriority() *int32 {
	return s.Priority
}

func (s *BatchSendMessageRequestMessages) SetDelaySeconds(v int32) *BatchSendMessageRequestMessages {
	s.DelaySeconds = &v
	return s
}

func (s *BatchSendMessageRequestMessages) SetMessageBody(v string) *BatchSendMessageRequestMessages {
	s.MessageBody = &v
	return s
}

func (s *BatchSendMessageRequestMessages) SetMessageGroupId(v string) *BatchSendMessageRequestMessages {
	s.MessageGroupId = &v
	return s
}

func (s *BatchSendMessageRequestMessages) SetPriority(v int32) *BatchSendMessageRequestMessages {
	s.Priority = &v
	return s
}

func (s *BatchSendMessageRequestMessages) Validate() error {
	return dara.Validate(s)
}
