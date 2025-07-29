// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotTopicChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *RunHotTopicChatRequest
	GetCategory() *string
	SetGenerateOptions(v []*string) *RunHotTopicChatRequest
	GetGenerateOptions() []*string
	SetHotTopicVersion(v string) *RunHotTopicChatRequest
	GetHotTopicVersion() *string
	SetHotTopics(v []*string) *RunHotTopicChatRequest
	GetHotTopics() []*string
	SetImageCount(v int32) *RunHotTopicChatRequest
	GetImageCount() *int32
	SetMessages(v []*RunHotTopicChatRequestMessages) *RunHotTopicChatRequest
	GetMessages() []*RunHotTopicChatRequestMessages
	SetModelCustomPromptTemplate(v string) *RunHotTopicChatRequest
	GetModelCustomPromptTemplate() *string
	SetModelId(v string) *RunHotTopicChatRequest
	GetModelId() *string
	SetOriginalSessionId(v string) *RunHotTopicChatRequest
	GetOriginalSessionId() *string
	SetPrompt(v string) *RunHotTopicChatRequest
	GetPrompt() *string
	SetStepForBroadcastContentConfig(v *RunHotTopicChatRequestStepForBroadcastContentConfig) *RunHotTopicChatRequest
	GetStepForBroadcastContentConfig() *RunHotTopicChatRequestStepForBroadcastContentConfig
	SetTaskId(v string) *RunHotTopicChatRequest
	GetTaskId() *string
}

type RunHotTopicChatRequest struct {
	Category        *string   `json:"category,omitempty" xml:"category,omitempty"`
	GenerateOptions []*string `json:"generateOptions,omitempty" xml:"generateOptions,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-13_12
	HotTopicVersion *string   `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	HotTopics       []*string `json:"hotTopics,omitempty" xml:"hotTopics,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ImageCount *int32                            `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	Messages   []*RunHotTopicChatRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
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
	OriginalSessionId             *string                                              `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	Prompt                        *string                                              `json:"prompt,omitempty" xml:"prompt,omitempty"`
	StepForBroadcastContentConfig *RunHotTopicChatRequestStepForBroadcastContentConfig `json:"stepForBroadcastContentConfig,omitempty" xml:"stepForBroadcastContentConfig,omitempty" type:"Struct"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunHotTopicChatRequest) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatRequest) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatRequest) GetCategory() *string {
	return s.Category
}

func (s *RunHotTopicChatRequest) GetGenerateOptions() []*string {
	return s.GenerateOptions
}

func (s *RunHotTopicChatRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *RunHotTopicChatRequest) GetHotTopics() []*string {
	return s.HotTopics
}

func (s *RunHotTopicChatRequest) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *RunHotTopicChatRequest) GetMessages() []*RunHotTopicChatRequestMessages {
	return s.Messages
}

func (s *RunHotTopicChatRequest) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *RunHotTopicChatRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunHotTopicChatRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunHotTopicChatRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunHotTopicChatRequest) GetStepForBroadcastContentConfig() *RunHotTopicChatRequestStepForBroadcastContentConfig {
	return s.StepForBroadcastContentConfig
}

func (s *RunHotTopicChatRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunHotTopicChatRequest) SetCategory(v string) *RunHotTopicChatRequest {
	s.Category = &v
	return s
}

func (s *RunHotTopicChatRequest) SetGenerateOptions(v []*string) *RunHotTopicChatRequest {
	s.GenerateOptions = v
	return s
}

func (s *RunHotTopicChatRequest) SetHotTopicVersion(v string) *RunHotTopicChatRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicChatRequest) SetHotTopics(v []*string) *RunHotTopicChatRequest {
	s.HotTopics = v
	return s
}

func (s *RunHotTopicChatRequest) SetImageCount(v int32) *RunHotTopicChatRequest {
	s.ImageCount = &v
	return s
}

func (s *RunHotTopicChatRequest) SetMessages(v []*RunHotTopicChatRequestMessages) *RunHotTopicChatRequest {
	s.Messages = v
	return s
}

func (s *RunHotTopicChatRequest) SetModelCustomPromptTemplate(v string) *RunHotTopicChatRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunHotTopicChatRequest) SetModelId(v string) *RunHotTopicChatRequest {
	s.ModelId = &v
	return s
}

func (s *RunHotTopicChatRequest) SetOriginalSessionId(v string) *RunHotTopicChatRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunHotTopicChatRequest) SetPrompt(v string) *RunHotTopicChatRequest {
	s.Prompt = &v
	return s
}

func (s *RunHotTopicChatRequest) SetStepForBroadcastContentConfig(v *RunHotTopicChatRequestStepForBroadcastContentConfig) *RunHotTopicChatRequest {
	s.StepForBroadcastContentConfig = v
	return s
}

func (s *RunHotTopicChatRequest) SetTaskId(v string) *RunHotTopicChatRequest {
	s.TaskId = &v
	return s
}

func (s *RunHotTopicChatRequest) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatRequestMessages struct {
	// example:
	//
	// xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-12-10 18:51:29
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunHotTopicChatRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatRequestMessages) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatRequestMessages) GetContent() *string {
	return s.Content
}

func (s *RunHotTopicChatRequestMessages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *RunHotTopicChatRequestMessages) GetRole() *string {
	return s.Role
}

func (s *RunHotTopicChatRequestMessages) SetContent(v string) *RunHotTopicChatRequestMessages {
	s.Content = &v
	return s
}

func (s *RunHotTopicChatRequestMessages) SetCreateTime(v string) *RunHotTopicChatRequestMessages {
	s.CreateTime = &v
	return s
}

func (s *RunHotTopicChatRequestMessages) SetRole(v string) *RunHotTopicChatRequestMessages {
	s.Role = &v
	return s
}

func (s *RunHotTopicChatRequestMessages) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatRequestStepForBroadcastContentConfig struct {
	Categories            []*string                                                                   `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	CustomHotValueWeights []*RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights `json:"customHotValueWeights,omitempty" xml:"customHotValueWeights,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TopicCount *int32 `json:"topicCount,omitempty" xml:"topicCount,omitempty"`
}

func (s RunHotTopicChatRequestStepForBroadcastContentConfig) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatRequestStepForBroadcastContentConfig) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) GetCategories() []*string {
	return s.Categories
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) GetCustomHotValueWeights() []*RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights {
	return s.CustomHotValueWeights
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) GetTopicCount() *int32 {
	return s.TopicCount
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) SetCategories(v []*string) *RunHotTopicChatRequestStepForBroadcastContentConfig {
	s.Categories = v
	return s
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) SetCustomHotValueWeights(v []*RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) *RunHotTopicChatRequestStepForBroadcastContentConfig {
	s.CustomHotValueWeights = v
	return s
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) SetTopicCount(v int32) *RunHotTopicChatRequestStepForBroadcastContentConfig {
	s.TopicCount = &v
	return s
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfig) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights struct {
	// example:
	//
	// comments
	Dimension *string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) GetDimension() *string {
	return s.Dimension
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) SetDimension(v string) *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights {
	s.Dimension = &v
	return s
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) SetWeight(v int32) *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights {
	s.Weight = &v
	return s
}

func (s *RunHotTopicChatRequestStepForBroadcastContentConfigCustomHotValueWeights) Validate() error {
	return dara.Validate(s)
}
