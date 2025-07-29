// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotTopicChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *RunHotTopicChatShrinkRequest
	GetCategory() *string
	SetGenerateOptionsShrink(v string) *RunHotTopicChatShrinkRequest
	GetGenerateOptionsShrink() *string
	SetHotTopicVersion(v string) *RunHotTopicChatShrinkRequest
	GetHotTopicVersion() *string
	SetHotTopicsShrink(v string) *RunHotTopicChatShrinkRequest
	GetHotTopicsShrink() *string
	SetImageCount(v int32) *RunHotTopicChatShrinkRequest
	GetImageCount() *int32
	SetMessagesShrink(v string) *RunHotTopicChatShrinkRequest
	GetMessagesShrink() *string
	SetModelCustomPromptTemplate(v string) *RunHotTopicChatShrinkRequest
	GetModelCustomPromptTemplate() *string
	SetModelId(v string) *RunHotTopicChatShrinkRequest
	GetModelId() *string
	SetOriginalSessionId(v string) *RunHotTopicChatShrinkRequest
	GetOriginalSessionId() *string
	SetPrompt(v string) *RunHotTopicChatShrinkRequest
	GetPrompt() *string
	SetStepForBroadcastContentConfigShrink(v string) *RunHotTopicChatShrinkRequest
	GetStepForBroadcastContentConfigShrink() *string
	SetTaskId(v string) *RunHotTopicChatShrinkRequest
	GetTaskId() *string
}

type RunHotTopicChatShrinkRequest struct {
	Category              *string `json:"category,omitempty" xml:"category,omitempty"`
	GenerateOptionsShrink *string `json:"generateOptions,omitempty" xml:"generateOptions,omitempty"`
	// example:
	//
	// 2024-09-13_12
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	HotTopicsShrink *string `json:"hotTopics,omitempty" xml:"hotTopics,omitempty"`
	// example:
	//
	// 1
	ImageCount     *int32  `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	MessagesShrink *string `json:"messages,omitempty" xml:"messages,omitempty"`
	// example:
	//
	// xx
	ModelCustomPromptTemplate *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5ax
	OriginalSessionId                   *string `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	Prompt                              *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	StepForBroadcastContentConfigShrink *string `json:"stepForBroadcastContentConfig,omitempty" xml:"stepForBroadcastContentConfig,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunHotTopicChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *RunHotTopicChatShrinkRequest) GetGenerateOptionsShrink() *string {
	return s.GenerateOptionsShrink
}

func (s *RunHotTopicChatShrinkRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *RunHotTopicChatShrinkRequest) GetHotTopicsShrink() *string {
	return s.HotTopicsShrink
}

func (s *RunHotTopicChatShrinkRequest) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *RunHotTopicChatShrinkRequest) GetMessagesShrink() *string {
	return s.MessagesShrink
}

func (s *RunHotTopicChatShrinkRequest) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *RunHotTopicChatShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunHotTopicChatShrinkRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunHotTopicChatShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunHotTopicChatShrinkRequest) GetStepForBroadcastContentConfigShrink() *string {
	return s.StepForBroadcastContentConfigShrink
}

func (s *RunHotTopicChatShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunHotTopicChatShrinkRequest) SetCategory(v string) *RunHotTopicChatShrinkRequest {
	s.Category = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetGenerateOptionsShrink(v string) *RunHotTopicChatShrinkRequest {
	s.GenerateOptionsShrink = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetHotTopicVersion(v string) *RunHotTopicChatShrinkRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetHotTopicsShrink(v string) *RunHotTopicChatShrinkRequest {
	s.HotTopicsShrink = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetImageCount(v int32) *RunHotTopicChatShrinkRequest {
	s.ImageCount = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetMessagesShrink(v string) *RunHotTopicChatShrinkRequest {
	s.MessagesShrink = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetModelCustomPromptTemplate(v string) *RunHotTopicChatShrinkRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetModelId(v string) *RunHotTopicChatShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetOriginalSessionId(v string) *RunHotTopicChatShrinkRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetPrompt(v string) *RunHotTopicChatShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetStepForBroadcastContentConfigShrink(v string) *RunHotTopicChatShrinkRequest {
	s.StepForBroadcastContentConfigShrink = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) SetTaskId(v string) *RunHotTopicChatShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunHotTopicChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
