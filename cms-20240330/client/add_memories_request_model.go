// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *AddMemoriesRequest
	GetAgentId() *string
	SetAppId(v string) *AddMemoriesRequest
	GetAppId() *string
	SetAsyncMode(v bool) *AddMemoriesRequest
	GetAsyncMode() *bool
	SetCustomInstructions(v string) *AddMemoriesRequest
	GetCustomInstructions() *string
	SetInfer(v bool) *AddMemoriesRequest
	GetInfer() *bool
	SetMessages(v []*AddMemoriesRequestMessages) *AddMemoriesRequest
	GetMessages() []*AddMemoriesRequestMessages
	SetMetadata(v map[string]interface{}) *AddMemoriesRequest
	GetMetadata() map[string]interface{}
	SetRunId(v string) *AddMemoriesRequest
	GetRunId() *string
	SetUserId(v string) *AddMemoriesRequest
	GetUserId() *string
}

type AddMemoriesRequest struct {
	// example:
	//
	// 952730733889060865
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// mm_480d961a1b5e4efe84603f4cbc0f
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// true
	AsyncMode *bool `json:"asyncMode,omitempty" xml:"asyncMode,omitempty"`
	// example:
	//
	// Your custom instructions here
	CustomInstructions *string `json:"customInstructions,omitempty" xml:"customInstructions,omitempty"`
	// example:
	//
	// true
	Infer    *bool                         `json:"infer,omitempty" xml:"infer,omitempty"`
	Messages []*AddMemoriesRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// {"sessionId":"test_session_001"}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// jr-80ded1d6953c64ea
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// test_user_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesRequest) GoString() string {
	return s.String()
}

func (s *AddMemoriesRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *AddMemoriesRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddMemoriesRequest) GetAsyncMode() *bool {
	return s.AsyncMode
}

func (s *AddMemoriesRequest) GetCustomInstructions() *string {
	return s.CustomInstructions
}

func (s *AddMemoriesRequest) GetInfer() *bool {
	return s.Infer
}

func (s *AddMemoriesRequest) GetMessages() []*AddMemoriesRequestMessages {
	return s.Messages
}

func (s *AddMemoriesRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *AddMemoriesRequest) GetRunId() *string {
	return s.RunId
}

func (s *AddMemoriesRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddMemoriesRequest) SetAgentId(v string) *AddMemoriesRequest {
	s.AgentId = &v
	return s
}

func (s *AddMemoriesRequest) SetAppId(v string) *AddMemoriesRequest {
	s.AppId = &v
	return s
}

func (s *AddMemoriesRequest) SetAsyncMode(v bool) *AddMemoriesRequest {
	s.AsyncMode = &v
	return s
}

func (s *AddMemoriesRequest) SetCustomInstructions(v string) *AddMemoriesRequest {
	s.CustomInstructions = &v
	return s
}

func (s *AddMemoriesRequest) SetInfer(v bool) *AddMemoriesRequest {
	s.Infer = &v
	return s
}

func (s *AddMemoriesRequest) SetMessages(v []*AddMemoriesRequestMessages) *AddMemoriesRequest {
	s.Messages = v
	return s
}

func (s *AddMemoriesRequest) SetMetadata(v map[string]interface{}) *AddMemoriesRequest {
	s.Metadata = v
	return s
}

func (s *AddMemoriesRequest) SetRunId(v string) *AddMemoriesRequest {
	s.RunId = &v
	return s
}

func (s *AddMemoriesRequest) SetUserId(v string) *AddMemoriesRequest {
	s.UserId = &v
	return s
}

func (s *AddMemoriesRequest) Validate() error {
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

type AddMemoriesRequestMessages struct {
	// example:
	//
	// My name is Zhang San and I live in Hangzhou.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s AddMemoriesRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesRequestMessages) GoString() string {
	return s.String()
}

func (s *AddMemoriesRequestMessages) GetContent() *string {
	return s.Content
}

func (s *AddMemoriesRequestMessages) GetRole() *string {
	return s.Role
}

func (s *AddMemoriesRequestMessages) SetContent(v string) *AddMemoriesRequestMessages {
	s.Content = &v
	return s
}

func (s *AddMemoriesRequestMessages) SetRole(v string) *AddMemoriesRequestMessages {
	s.Role = &v
	return s
}

func (s *AddMemoriesRequestMessages) Validate() error {
	return dara.Validate(s)
}
