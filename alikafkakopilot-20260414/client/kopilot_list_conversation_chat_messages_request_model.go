// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationChatMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeforeTurnId(v int32) *KopilotListConversationChatMessagesRequest
	GetBeforeTurnId() *int32
	SetPageSize(v int32) *KopilotListConversationChatMessagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *KopilotListConversationChatMessagesRequest
	GetRegionId() *string
	SetSessionId(v string) *KopilotListConversationChatMessagesRequest
	GetSessionId() *string
}

type KopilotListConversationChatMessagesRequest struct {
	BeforeTurnId *int32 `json:"BeforeTurnId,omitempty" xml:"BeforeTurnId,omitempty"`
	PageSize     *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s KopilotListConversationChatMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesRequest) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesRequest) GetBeforeTurnId() *int32 {
	return s.BeforeTurnId
}

func (s *KopilotListConversationChatMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *KopilotListConversationChatMessagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotListConversationChatMessagesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *KopilotListConversationChatMessagesRequest) SetBeforeTurnId(v int32) *KopilotListConversationChatMessagesRequest {
	s.BeforeTurnId = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) SetPageSize(v int32) *KopilotListConversationChatMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) SetRegionId(v string) *KopilotListConversationChatMessagesRequest {
	s.RegionId = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) SetSessionId(v string) *KopilotListConversationChatMessagesRequest {
	s.SessionId = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) Validate() error {
	return dara.Validate(s)
}
