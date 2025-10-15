// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunChatResultGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChoices(v []*RunChatResultGenerationResponseBodyChoices) *RunChatResultGenerationResponseBody
	GetChoices() []*RunChatResultGenerationResponseBodyChoices
	SetCreated(v int64) *RunChatResultGenerationResponseBody
	GetCreated() *int64
	SetId(v string) *RunChatResultGenerationResponseBody
	GetId() *string
	SetModelId(v string) *RunChatResultGenerationResponseBody
	GetModelId() *string
	SetRequestId(v string) *RunChatResultGenerationResponseBody
	GetRequestId() *string
	SetTime(v string) *RunChatResultGenerationResponseBody
	GetTime() *string
	SetTotalTokens(v int32) *RunChatResultGenerationResponseBody
	GetTotalTokens() *int32
	SetUsage(v *RunChatResultGenerationResponseBodyUsage) *RunChatResultGenerationResponseBody
	GetUsage() *RunChatResultGenerationResponseBodyUsage
}

type RunChatResultGenerationResponseBody struct {
	Choices []*RunChatResultGenerationResponseBodyChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1720602203
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 500
	TotalTokens *int32                                    `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	Usage       *RunChatResultGenerationResponseBodyUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunChatResultGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBody) GetChoices() []*RunChatResultGenerationResponseBodyChoices {
	return s.Choices
}

func (s *RunChatResultGenerationResponseBody) GetCreated() *int64 {
	return s.Created
}

func (s *RunChatResultGenerationResponseBody) GetId() *string {
	return s.Id
}

func (s *RunChatResultGenerationResponseBody) GetModelId() *string {
	return s.ModelId
}

func (s *RunChatResultGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunChatResultGenerationResponseBody) GetTime() *string {
	return s.Time
}

func (s *RunChatResultGenerationResponseBody) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *RunChatResultGenerationResponseBody) GetUsage() *RunChatResultGenerationResponseBodyUsage {
	return s.Usage
}

func (s *RunChatResultGenerationResponseBody) SetChoices(v []*RunChatResultGenerationResponseBodyChoices) *RunChatResultGenerationResponseBody {
	s.Choices = v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetCreated(v int64) *RunChatResultGenerationResponseBody {
	s.Created = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetId(v string) *RunChatResultGenerationResponseBody {
	s.Id = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetModelId(v string) *RunChatResultGenerationResponseBody {
	s.ModelId = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetRequestId(v string) *RunChatResultGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetTime(v string) *RunChatResultGenerationResponseBody {
	s.Time = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetTotalTokens(v int32) *RunChatResultGenerationResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetUsage(v *RunChatResultGenerationResponseBodyUsage) *RunChatResultGenerationResponseBody {
	s.Usage = v
	return s
}

func (s *RunChatResultGenerationResponseBody) Validate() error {
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

type RunChatResultGenerationResponseBodyChoices struct {
	// example:
	//
	// null
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                             `json:"index,omitempty" xml:"index,omitempty"`
	Message *RunChatResultGenerationResponseBodyChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s RunChatResultGenerationResponseBodyChoices) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyChoices) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyChoices) GetFinishReason() *string {
	return s.FinishReason
}

func (s *RunChatResultGenerationResponseBodyChoices) GetIndex() *int32 {
	return s.Index
}

func (s *RunChatResultGenerationResponseBodyChoices) GetMessage() *RunChatResultGenerationResponseBodyChoicesMessage {
	return s.Message
}

func (s *RunChatResultGenerationResponseBodyChoices) SetFinishReason(v string) *RunChatResultGenerationResponseBodyChoices {
	s.FinishReason = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoices) SetIndex(v int32) *RunChatResultGenerationResponseBodyChoices {
	s.Index = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoices) SetMessage(v *RunChatResultGenerationResponseBodyChoicesMessage) *RunChatResultGenerationResponseBodyChoices {
	s.Message = v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoices) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunChatResultGenerationResponseBodyChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role      *string                  `json:"role,omitempty" xml:"role,omitempty"`
	ToolCalls []map[string]interface{} `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationResponseBodyChoicesMessage) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyChoicesMessage) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) GetContent() *string {
	return s.Content
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) GetRole() *string {
	return s.Role
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) GetToolCalls() []map[string]interface{} {
	return s.ToolCalls
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetContent(v string) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.Content = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetRole(v string) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.Role = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetToolCalls(v []map[string]interface{}) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.ToolCalls = v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) Validate() error {
	return dara.Validate(s)
}

type RunChatResultGenerationResponseBodyUsage struct {
	// example:
	//
	// 0
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 0
	ImageTokens *int32 `json:"imageTokens,omitempty" xml:"imageTokens,omitempty"`
	// example:
	//
	// 200
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 300
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 500
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunChatResultGenerationResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyUsage) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *RunChatResultGenerationResponseBodyUsage) GetImageTokens() *int32 {
	return s.ImageTokens
}

func (s *RunChatResultGenerationResponseBodyUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *RunChatResultGenerationResponseBodyUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *RunChatResultGenerationResponseBodyUsage) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *RunChatResultGenerationResponseBodyUsage) SetImageCount(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.ImageCount = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetImageTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.ImageTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetInputTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetOutputTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetTotalTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
