// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationChatMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *KopilotListConversationChatMessagesResponseBody
	GetCode() *int64
	SetData(v *KopilotListConversationChatMessagesResponseBodyData) *KopilotListConversationChatMessagesResponseBody
	GetData() *KopilotListConversationChatMessagesResponseBodyData
	SetRequestId(v string) *KopilotListConversationChatMessagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *KopilotListConversationChatMessagesResponseBody
	GetSuccess() *bool
}

type KopilotListConversationChatMessagesResponseBody struct {
	Code      *int64                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *KopilotListConversationChatMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *KopilotListConversationChatMessagesResponseBody) GetData() *KopilotListConversationChatMessagesResponseBodyData {
	return s.Data
}

func (s *KopilotListConversationChatMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KopilotListConversationChatMessagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *KopilotListConversationChatMessagesResponseBody) SetCode(v int64) *KopilotListConversationChatMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBody) SetData(v *KopilotListConversationChatMessagesResponseBodyData) *KopilotListConversationChatMessagesResponseBody {
	s.Data = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBody) SetRequestId(v string) *KopilotListConversationChatMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBody) SetSuccess(v bool) *KopilotListConversationChatMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationChatMessagesResponseBodyData struct {
	HasMore          *bool                                                          `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	Messages         []*KopilotListConversationChatMessagesResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	NextBeforeTurnId *int64                                                         `json:"NextBeforeTurnId,omitempty" xml:"NextBeforeTurnId,omitempty"`
	SessionId        *string                                                        `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TotalTurns       *int64                                                         `json:"TotalTurns,omitempty" xml:"TotalTurns,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetMessages() []*KopilotListConversationChatMessagesResponseBodyDataMessages {
	return s.Messages
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetNextBeforeTurnId() *int64 {
	return s.NextBeforeTurnId
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetTotalTurns() *int64 {
	return s.TotalTurns
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetHasMore(v bool) *KopilotListConversationChatMessagesResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetMessages(v []*KopilotListConversationChatMessagesResponseBodyDataMessages) *KopilotListConversationChatMessagesResponseBodyData {
	s.Messages = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetNextBeforeTurnId(v int64) *KopilotListConversationChatMessagesResponseBodyData {
	s.NextBeforeTurnId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetSessionId(v string) *KopilotListConversationChatMessagesResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetTotalTurns(v int64) *KopilotListConversationChatMessagesResponseBodyData {
	s.TotalTurns = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) Validate() error {
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

type KopilotListConversationChatMessagesResponseBodyDataMessages struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Feedback   *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	TurnId     *string `json:"TurnId,omitempty" xml:"TurnId,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetContent() *string {
	return s.Content
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetFeedback() *string {
	return s.Feedback
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetRole() *string {
	return s.Role
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetTurnId() *string {
	return s.TurnId
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetContent(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.Content = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetCreateTime(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.CreateTime = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetFeedback(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.Feedback = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetRole(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.Role = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetTurnId(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.TurnId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) Validate() error {
	return dara.Validate(s)
}
