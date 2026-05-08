// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContextsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextType(v string) *AddContextsRequest
	GetContextType() *string
	SetItems(v []*AddContextsRequestItems) *AddContextsRequest
	GetItems() []*AddContextsRequestItems
	SetMemoryType(v string) *AddContextsRequest
	GetMemoryType() *string
}

type AddContextsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// memory
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// This parameter is required.
	Items []*AddContextsRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// short
	MemoryType *string `json:"memoryType,omitempty" xml:"memoryType,omitempty"`
}

func (s AddContextsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddContextsRequest) GoString() string {
	return s.String()
}

func (s *AddContextsRequest) GetContextType() *string {
	return s.ContextType
}

func (s *AddContextsRequest) GetItems() []*AddContextsRequestItems {
	return s.Items
}

func (s *AddContextsRequest) GetMemoryType() *string {
	return s.MemoryType
}

func (s *AddContextsRequest) SetContextType(v string) *AddContextsRequest {
	s.ContextType = &v
	return s
}

func (s *AddContextsRequest) SetItems(v []*AddContextsRequestItems) *AddContextsRequest {
	s.Items = v
	return s
}

func (s *AddContextsRequest) SetMemoryType(v string) *AddContextsRequest {
	s.MemoryType = &v
	return s
}

func (s *AddContextsRequest) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddContextsRequestItems struct {
	// example:
	//
	// 952730733889060865
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// mm_480d961a1b5e4efe84603f4cbc0f
	AppId      *string   `json:"appId,omitempty" xml:"appId,omitempty"`
	Categories []*string `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	// example:
	//
	// You are a conversation assistant.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// Your custom instructions here
	CustomInstructions *string `json:"customInstructions,omitempty" xml:"customInstructions,omitempty"`
	// example:
	//
	// {
	//
	// 	"taskType": "troubleshooting",
	//
	// 	"complexity": "medium",
	//
	// 	"confidence": 0.95
	//
	// }
	Experience map[string]interface{} `json:"experience,omitempty" xml:"experience,omitempty"`
	// example:
	//
	// 1731231212334396
	ExpirationDate *string `json:"expirationDate,omitempty" xml:"expirationDate,omitempty"`
	// example:
	//
	// true
	Immutable *bool `json:"immutable,omitempty" xml:"immutable,omitempty"`
	// example:
	//
	// true
	Infer    *bool                    `json:"infer,omitempty" xml:"infer,omitempty"`
	Labels   map[string]*string       `json:"labels,omitempty" xml:"labels,omitempty"`
	Messages []map[string]interface{} `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
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
	// 1774578167
	Timestamp        *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	TriggerCondition *string `json:"triggerCondition,omitempty" xml:"triggerCondition,omitempty"`
	// example:
	//
	// test_user_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddContextsRequestItems) String() string {
	return dara.Prettify(s)
}

func (s AddContextsRequestItems) GoString() string {
	return s.String()
}

func (s *AddContextsRequestItems) GetAgentId() *string {
	return s.AgentId
}

func (s *AddContextsRequestItems) GetAppId() *string {
	return s.AppId
}

func (s *AddContextsRequestItems) GetCategories() []*string {
	return s.Categories
}

func (s *AddContextsRequestItems) GetContent() *string {
	return s.Content
}

func (s *AddContextsRequestItems) GetCustomInstructions() *string {
	return s.CustomInstructions
}

func (s *AddContextsRequestItems) GetExperience() map[string]interface{} {
	return s.Experience
}

func (s *AddContextsRequestItems) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *AddContextsRequestItems) GetImmutable() *bool {
	return s.Immutable
}

func (s *AddContextsRequestItems) GetInfer() *bool {
	return s.Infer
}

func (s *AddContextsRequestItems) GetLabels() map[string]*string {
	return s.Labels
}

func (s *AddContextsRequestItems) GetMessages() []map[string]interface{} {
	return s.Messages
}

func (s *AddContextsRequestItems) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *AddContextsRequestItems) GetRunId() *string {
	return s.RunId
}

func (s *AddContextsRequestItems) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AddContextsRequestItems) GetTriggerCondition() *string {
	return s.TriggerCondition
}

func (s *AddContextsRequestItems) GetUserId() *string {
	return s.UserId
}

func (s *AddContextsRequestItems) SetAgentId(v string) *AddContextsRequestItems {
	s.AgentId = &v
	return s
}

func (s *AddContextsRequestItems) SetAppId(v string) *AddContextsRequestItems {
	s.AppId = &v
	return s
}

func (s *AddContextsRequestItems) SetCategories(v []*string) *AddContextsRequestItems {
	s.Categories = v
	return s
}

func (s *AddContextsRequestItems) SetContent(v string) *AddContextsRequestItems {
	s.Content = &v
	return s
}

func (s *AddContextsRequestItems) SetCustomInstructions(v string) *AddContextsRequestItems {
	s.CustomInstructions = &v
	return s
}

func (s *AddContextsRequestItems) SetExperience(v map[string]interface{}) *AddContextsRequestItems {
	s.Experience = v
	return s
}

func (s *AddContextsRequestItems) SetExpirationDate(v string) *AddContextsRequestItems {
	s.ExpirationDate = &v
	return s
}

func (s *AddContextsRequestItems) SetImmutable(v bool) *AddContextsRequestItems {
	s.Immutable = &v
	return s
}

func (s *AddContextsRequestItems) SetInfer(v bool) *AddContextsRequestItems {
	s.Infer = &v
	return s
}

func (s *AddContextsRequestItems) SetLabels(v map[string]*string) *AddContextsRequestItems {
	s.Labels = v
	return s
}

func (s *AddContextsRequestItems) SetMessages(v []map[string]interface{}) *AddContextsRequestItems {
	s.Messages = v
	return s
}

func (s *AddContextsRequestItems) SetMetadata(v map[string]interface{}) *AddContextsRequestItems {
	s.Metadata = v
	return s
}

func (s *AddContextsRequestItems) SetRunId(v string) *AddContextsRequestItems {
	s.RunId = &v
	return s
}

func (s *AddContextsRequestItems) SetTimestamp(v int64) *AddContextsRequestItems {
	s.Timestamp = &v
	return s
}

func (s *AddContextsRequestItems) SetTriggerCondition(v string) *AddContextsRequestItems {
	s.TriggerCondition = &v
	return s
}

func (s *AddContextsRequestItems) SetUserId(v string) *AddContextsRequestItems {
	s.UserId = &v
	return s
}

func (s *AddContextsRequestItems) Validate() error {
	return dara.Validate(s)
}
