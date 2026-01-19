// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessCustomIMCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *ProcessCustomIMCallbackRequest
	GetAccessChannelId() *string
	SetConversationId(v string) *ProcessCustomIMCallbackRequest
	GetConversationId() *string
	SetInstanceId(v string) *ProcessCustomIMCallbackRequest
	GetInstanceId() *string
	SetMessageContent(v string) *ProcessCustomIMCallbackRequest
	GetMessageContent() *string
	SetRequestId(v string) *ProcessCustomIMCallbackRequest
	GetRequestId() *string
	SetSenderAvatarMediaId(v string) *ProcessCustomIMCallbackRequest
	GetSenderAvatarMediaId() *string
	SetSenderId(v string) *ProcessCustomIMCallbackRequest
	GetSenderId() *string
	SetSenderName(v string) *ProcessCustomIMCallbackRequest
	GetSenderName() *string
}

type ProcessCustomIMCallbackRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 420F203D-87ED-599A-B137-D9D914EE9E70
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 27578380-d382-11ee-9cca-adec43112a87
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "ContentType": "Text",
	//
	//     "Text": "Hello"
	//
	//   }
	MessageContent *string `json:"MessageContent,omitempty" xml:"MessageContent,omitempty"`
	// example:
	//
	// FCEFE806-E67C-577E-865B-4ED398F2F648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// FCEFE806-E67C-577E-865B-4ED398F2F648
	SenderAvatarMediaId *string `json:"SenderAvatarMediaId,omitempty" xml:"SenderAvatarMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 654f1054dcda1b251282cbdf
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// example:
	//
	// tom
	SenderName *string `json:"SenderName,omitempty" xml:"SenderName,omitempty"`
}

func (s ProcessCustomIMCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s ProcessCustomIMCallbackRequest) GoString() string {
	return s.String()
}

func (s *ProcessCustomIMCallbackRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ProcessCustomIMCallbackRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ProcessCustomIMCallbackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ProcessCustomIMCallbackRequest) GetMessageContent() *string {
	return s.MessageContent
}

func (s *ProcessCustomIMCallbackRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ProcessCustomIMCallbackRequest) GetSenderAvatarMediaId() *string {
	return s.SenderAvatarMediaId
}

func (s *ProcessCustomIMCallbackRequest) GetSenderId() *string {
	return s.SenderId
}

func (s *ProcessCustomIMCallbackRequest) GetSenderName() *string {
	return s.SenderName
}

func (s *ProcessCustomIMCallbackRequest) SetAccessChannelId(v string) *ProcessCustomIMCallbackRequest {
	s.AccessChannelId = &v
	return s
}

func (s *ProcessCustomIMCallbackRequest) SetConversationId(v string) *ProcessCustomIMCallbackRequest {
	s.ConversationId = &v
	return s
}

func (s *ProcessCustomIMCallbackRequest) SetInstanceId(v string) *ProcessCustomIMCallbackRequest {
	s.InstanceId = &v
	return s
}

func (s *ProcessCustomIMCallbackRequest) SetMessageContent(v string) *ProcessCustomIMCallbackRequest {
	s.MessageContent = &v
	return s
}

func (s *ProcessCustomIMCallbackRequest) SetRequestId(v string) *ProcessCustomIMCallbackRequest {
	s.RequestId = &v
	return s
}

func (s *ProcessCustomIMCallbackRequest) SetSenderAvatarMediaId(v string) *ProcessCustomIMCallbackRequest {
	s.SenderAvatarMediaId = &v
	return s
}

func (s *ProcessCustomIMCallbackRequest) SetSenderId(v string) *ProcessCustomIMCallbackRequest {
	s.SenderId = &v
	return s
}

func (s *ProcessCustomIMCallbackRequest) SetSenderName(v string) *ProcessCustomIMCallbackRequest {
	s.SenderName = &v
	return s
}

func (s *ProcessCustomIMCallbackRequest) Validate() error {
	return dara.Validate(s)
}
