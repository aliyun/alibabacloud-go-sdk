// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ChatWithDesensitizeResponseBodyData) *ChatWithDesensitizeResponseBody
	GetData() *ChatWithDesensitizeResponseBodyData
	SetErrorCode(v string) *ChatWithDesensitizeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ChatWithDesensitizeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ChatWithDesensitizeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatWithDesensitizeResponseBody
	GetSuccess() *bool
}

type ChatWithDesensitizeResponseBody struct {
	Data *ChatWithDesensitizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatWithDesensitizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeResponseBody) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeResponseBody) GetData() *ChatWithDesensitizeResponseBodyData {
	return s.Data
}

func (s *ChatWithDesensitizeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChatWithDesensitizeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ChatWithDesensitizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatWithDesensitizeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatWithDesensitizeResponseBody) SetData(v *ChatWithDesensitizeResponseBodyData) *ChatWithDesensitizeResponseBody {
	s.Data = v
	return s
}

func (s *ChatWithDesensitizeResponseBody) SetErrorCode(v string) *ChatWithDesensitizeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChatWithDesensitizeResponseBody) SetErrorMessage(v string) *ChatWithDesensitizeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ChatWithDesensitizeResponseBody) SetRequestId(v string) *ChatWithDesensitizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatWithDesensitizeResponseBody) SetSuccess(v bool) *ChatWithDesensitizeResponseBody {
	s.Success = &v
	return s
}

func (s *ChatWithDesensitizeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithDesensitizeResponseBodyData struct {
	Choices []*ChatWithDesensitizeResponseBodyDataChoices `json:"Choices,omitempty" xml:"Choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1763710100
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string                                   `json:"Model,omitempty" xml:"Model,omitempty"`
	Usage *ChatWithDesensitizeResponseBodyDataUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s ChatWithDesensitizeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeResponseBodyData) GetChoices() []*ChatWithDesensitizeResponseBodyDataChoices {
	return s.Choices
}

func (s *ChatWithDesensitizeResponseBodyData) GetCreated() *string {
	return s.Created
}

func (s *ChatWithDesensitizeResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ChatWithDesensitizeResponseBodyData) GetUsage() *ChatWithDesensitizeResponseBodyDataUsage {
	return s.Usage
}

func (s *ChatWithDesensitizeResponseBodyData) SetChoices(v []*ChatWithDesensitizeResponseBodyDataChoices) *ChatWithDesensitizeResponseBodyData {
	s.Choices = v
	return s
}

func (s *ChatWithDesensitizeResponseBodyData) SetCreated(v string) *ChatWithDesensitizeResponseBodyData {
	s.Created = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyData) SetModel(v string) *ChatWithDesensitizeResponseBodyData {
	s.Model = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyData) SetUsage(v *ChatWithDesensitizeResponseBodyDataUsage) *ChatWithDesensitizeResponseBodyData {
	s.Usage = v
	return s
}

func (s *ChatWithDesensitizeResponseBodyData) Validate() error {
	if s.Choices != nil {
		for _, item := range s.Choices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithDesensitizeResponseBodyDataChoices struct {
	// example:
	//
	// stop
	FinishReason *string                                            `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	Message      *ChatWithDesensitizeResponseBodyDataChoicesMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
}

func (s ChatWithDesensitizeResponseBodyDataChoices) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeResponseBodyDataChoices) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeResponseBodyDataChoices) GetFinishReason() *string {
	return s.FinishReason
}

func (s *ChatWithDesensitizeResponseBodyDataChoices) GetMessage() *ChatWithDesensitizeResponseBodyDataChoicesMessage {
	return s.Message
}

func (s *ChatWithDesensitizeResponseBodyDataChoices) SetFinishReason(v string) *ChatWithDesensitizeResponseBodyDataChoices {
	s.FinishReason = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataChoices) SetMessage(v *ChatWithDesensitizeResponseBodyDataChoicesMessage) *ChatWithDesensitizeResponseBodyDataChoices {
	s.Message = v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataChoices) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithDesensitizeResponseBodyDataChoicesMessage struct {
	// example:
	//
	// 你好呀！
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 嗯，用户发了个“你好”，看起来是想打招呼...
	ReasoningContent *string `json:"ReasoningContent,omitempty" xml:"ReasoningContent,omitempty"`
	// example:
	//
	// system
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ChatWithDesensitizeResponseBodyDataChoicesMessage) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeResponseBodyDataChoicesMessage) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeResponseBodyDataChoicesMessage) GetContent() *string {
	return s.Content
}

func (s *ChatWithDesensitizeResponseBodyDataChoicesMessage) GetReasoningContent() *string {
	return s.ReasoningContent
}

func (s *ChatWithDesensitizeResponseBodyDataChoicesMessage) GetRole() *string {
	return s.Role
}

func (s *ChatWithDesensitizeResponseBodyDataChoicesMessage) SetContent(v string) *ChatWithDesensitizeResponseBodyDataChoicesMessage {
	s.Content = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataChoicesMessage) SetReasoningContent(v string) *ChatWithDesensitizeResponseBodyDataChoicesMessage {
	s.ReasoningContent = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataChoicesMessage) SetRole(v string) *ChatWithDesensitizeResponseBodyDataChoicesMessage {
	s.Role = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataChoicesMessage) Validate() error {
	return dara.Validate(s)
}

type ChatWithDesensitizeResponseBodyDataUsage struct {
	// example:
	//
	// 10
	CompletionTokens *string `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// example:
	//
	// 9
	PromptTokens *string `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	// example:
	//
	// 19
	TotalTokens *string `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s ChatWithDesensitizeResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) GetCompletionTokens() *string {
	return s.CompletionTokens
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) GetPromptTokens() *string {
	return s.PromptTokens
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) SetCompletionTokens(v string) *ChatWithDesensitizeResponseBodyDataUsage {
	s.CompletionTokens = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) SetPromptTokens(v string) *ChatWithDesensitizeResponseBodyDataUsage {
	s.PromptTokens = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) SetTotalTokens(v string) *ChatWithDesensitizeResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
