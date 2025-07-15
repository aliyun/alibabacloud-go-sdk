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
	// cf584733-***-***-9699-cb77aa3b7aa6
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d165de4f-9851-445e-9535-66ebfa72fa51
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MessageContent *string `json:"MessageContent,omitempty" xml:"MessageContent,omitempty"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// xxxx
	SenderAvatarMediaId *string `json:"SenderAvatarMediaId,omitempty" xml:"SenderAvatarMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 63061274befd6b545aab4c83
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
