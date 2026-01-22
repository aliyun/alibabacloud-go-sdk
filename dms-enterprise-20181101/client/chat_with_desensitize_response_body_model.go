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
	// The data returned.
	Data *ChatWithDesensitizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request fails.
	//
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
	// The candidate array for model-generated content.
	Choices []*ChatWithDesensitizeResponseBodyDataChoices `json:"Choices,omitempty" xml:"Choices,omitempty" type:"Repeated"`
	// The Unix timestamp (in seconds) when the request was created.
	//
	// example:
	//
	// 1763710100
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// Error message, provided when StatusCode is not 200.
	//
	// example:
	//
	// InvalidParameter
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The model used for this request.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Error code, 200 for normal calls, others for exceptions.
	//
	// example:
	//
	// 200
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// Error type.
	//
	// example:
	//
	// invalid_request_error
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The token consumption information of this request.
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

func (s *ChatWithDesensitizeResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ChatWithDesensitizeResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ChatWithDesensitizeResponseBodyData) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ChatWithDesensitizeResponseBodyData) GetType() *string {
	return s.Type
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

func (s *ChatWithDesensitizeResponseBodyData) SetMessage(v string) *ChatWithDesensitizeResponseBodyData {
	s.Message = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyData) SetModel(v string) *ChatWithDesensitizeResponseBodyData {
	s.Model = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyData) SetStatusCode(v string) *ChatWithDesensitizeResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyData) SetType(v string) *ChatWithDesensitizeResponseBodyData {
	s.Type = &v
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
	// Finish reason: ● stop: The model reached a natural stop point or a specified stop sequence. ● length: Generation ended because the maximum number of tokens was reached. ● tool_calls: The model stopped because it needs to call a tool to proceed.
	//
	// example:
	//
	// stop
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// Token probability information of model output.
	//
	// example:
	//
	// {}
	Logprobs map[string]interface{} `json:"Logprobs,omitempty" xml:"Logprobs,omitempty"`
	// The message body output by the model.
	Message *ChatWithDesensitizeResponseBodyDataChoicesMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
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

func (s *ChatWithDesensitizeResponseBodyDataChoices) GetLogprobs() map[string]interface{} {
	return s.Logprobs
}

func (s *ChatWithDesensitizeResponseBodyDataChoices) GetMessage() *ChatWithDesensitizeResponseBodyDataChoicesMessage {
	return s.Message
}

func (s *ChatWithDesensitizeResponseBodyDataChoices) SetFinishReason(v string) *ChatWithDesensitizeResponseBodyDataChoices {
	s.FinishReason = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataChoices) SetLogprobs(v map[string]interface{}) *ChatWithDesensitizeResponseBodyDataChoices {
	s.Logprobs = v
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
	// The content of the model\\"s response.
	//
	// example:
	//
	// 你好呀！
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The internal reasoning content of the model.
	//
	// example:
	//
	// 嗯，用户发了个“你好”，看起来是想打招呼...
	ReasoningContent *string `json:"ReasoningContent,omitempty" xml:"ReasoningContent,omitempty"`
	// Message role.
	//
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
	// The number of output tokens.
	//
	// example:
	//
	// 10
	CompletionTokens *string `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// Fine-grained classification of output tokens when using the Qwen-VL model.
	//
	// example:
	//
	// {}
	CompletionTokensDetails map[string]*string `json:"CompletionTokensDetails,omitempty" xml:"CompletionTokensDetails,omitempty"`
	// The number of input tokens.
	//
	// example:
	//
	// 9
	PromptTokens *string `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	// Fine-grained classification of input tokens.
	//
	// example:
	//
	// {}
	PromptTokensDetails map[string]*string `json:"PromptTokensDetails,omitempty" xml:"PromptTokensDetails,omitempty"`
	// The total number of tokens consumed.
	//
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

func (s *ChatWithDesensitizeResponseBodyDataUsage) GetCompletionTokensDetails() map[string]*string {
	return s.CompletionTokensDetails
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) GetPromptTokens() *string {
	return s.PromptTokens
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) GetPromptTokensDetails() map[string]*string {
	return s.PromptTokensDetails
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) SetCompletionTokens(v string) *ChatWithDesensitizeResponseBodyDataUsage {
	s.CompletionTokens = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) SetCompletionTokensDetails(v map[string]*string) *ChatWithDesensitizeResponseBodyDataUsage {
	s.CompletionTokensDetails = v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) SetPromptTokens(v string) *ChatWithDesensitizeResponseBodyDataUsage {
	s.PromptTokens = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) SetPromptTokensDetails(v map[string]*string) *ChatWithDesensitizeResponseBodyDataUsage {
	s.PromptTokensDetails = v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) SetTotalTokens(v string) *ChatWithDesensitizeResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

func (s *ChatWithDesensitizeResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
