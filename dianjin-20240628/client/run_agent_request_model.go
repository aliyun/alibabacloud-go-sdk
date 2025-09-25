// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotId(v string) *RunAgentRequest
	GetBotId() *string
	SetModelId(v string) *RunAgentRequest
	GetModelId() *string
	SetStream(v bool) *RunAgentRequest
	GetStream() *bool
	SetThreadId(v string) *RunAgentRequest
	GetThreadId() *string
	SetUseDraft(v bool) *RunAgentRequest
	GetUseDraft() *bool
	SetUserContent(v string) *RunAgentRequest
	GetUserContent() *string
	SetUserInputs(v map[string]interface{}) *RunAgentRequest
	GetUserInputs() map[string]interface{}
	SetVersionId(v string) *RunAgentRequest
	GetVersionId() *string
}

type RunAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d6zxykawk9
	BotId *string `json:"botId,omitempty" xml:"botId,omitempty"`
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// example:
	//
	// 4vlag5ken3
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// false
	UseDraft *bool `json:"useDraft,omitempty" xml:"useDraft,omitempty"`
	// This parameter is required.
	UserContent *string                `json:"userContent,omitempty" xml:"userContent,omitempty"`
	UserInputs  map[string]interface{} `json:"userInputs,omitempty" xml:"userInputs,omitempty"`
	// example:
	//
	// w4paqoezm2
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s RunAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s RunAgentRequest) GoString() string {
	return s.String()
}

func (s *RunAgentRequest) GetBotId() *string {
	return s.BotId
}

func (s *RunAgentRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunAgentRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunAgentRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *RunAgentRequest) GetUseDraft() *bool {
	return s.UseDraft
}

func (s *RunAgentRequest) GetUserContent() *string {
	return s.UserContent
}

func (s *RunAgentRequest) GetUserInputs() map[string]interface{} {
	return s.UserInputs
}

func (s *RunAgentRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *RunAgentRequest) SetBotId(v string) *RunAgentRequest {
	s.BotId = &v
	return s
}

func (s *RunAgentRequest) SetModelId(v string) *RunAgentRequest {
	s.ModelId = &v
	return s
}

func (s *RunAgentRequest) SetStream(v bool) *RunAgentRequest {
	s.Stream = &v
	return s
}

func (s *RunAgentRequest) SetThreadId(v string) *RunAgentRequest {
	s.ThreadId = &v
	return s
}

func (s *RunAgentRequest) SetUseDraft(v bool) *RunAgentRequest {
	s.UseDraft = &v
	return s
}

func (s *RunAgentRequest) SetUserContent(v string) *RunAgentRequest {
	s.UserContent = &v
	return s
}

func (s *RunAgentRequest) SetUserInputs(v map[string]interface{}) *RunAgentRequest {
	s.UserInputs = v
	return s
}

func (s *RunAgentRequest) SetVersionId(v string) *RunAgentRequest {
	s.VersionId = &v
	return s
}

func (s *RunAgentRequest) Validate() error {
	return dara.Validate(s)
}
