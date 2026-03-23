// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationPromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationPromptRequest
	GetApplicationId() *string
	SetPromptName(v string) *CreateApplicationPromptRequest
	GetPromptName() *string
	SetPromptType(v string) *CreateApplicationPromptRequest
	GetPromptType() *string
	SetPromptValue(v string) *CreateApplicationPromptRequest
	GetPromptValue() *string
}

type CreateApplicationPromptRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my_prompt
	PromptName *string `json:"PromptName,omitempty" xml:"PromptName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CONVERSATION_SUMMERY
	//
	// SEMANTIC_MEMORY
	//
	// MEMORY_ANSWER_PROMPT
	//
	// USER_MEMORY_EXTRACTION_PROMPT
	//
	// AGENT_MEMORY_EXTRACTION_PROMPT
	//
	// PROCEDURAL_MEMORY_SYSTEM_PROMPT
	//
	// RETRIEVE_NODES_PROMPT
	//
	// EXTRACT_RELATIONS_PROMPT
	//
	// UPDATE_GRAPH_PROMPT
	//
	// DELETE_RELATIONS_SYSTEM_PROMPT
	PromptType *string `json:"PromptType,omitempty" xml:"PromptType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prompt content
	PromptValue *string `json:"PromptValue,omitempty" xml:"PromptValue,omitempty"`
}

func (s CreateApplicationPromptRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationPromptRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationPromptRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationPromptRequest) GetPromptName() *string {
	return s.PromptName
}

func (s *CreateApplicationPromptRequest) GetPromptType() *string {
	return s.PromptType
}

func (s *CreateApplicationPromptRequest) GetPromptValue() *string {
	return s.PromptValue
}

func (s *CreateApplicationPromptRequest) SetApplicationId(v string) *CreateApplicationPromptRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationPromptRequest) SetPromptName(v string) *CreateApplicationPromptRequest {
	s.PromptName = &v
	return s
}

func (s *CreateApplicationPromptRequest) SetPromptType(v string) *CreateApplicationPromptRequest {
	s.PromptType = &v
	return s
}

func (s *CreateApplicationPromptRequest) SetPromptValue(v string) *CreateApplicationPromptRequest {
	s.PromptValue = &v
	return s
}

func (s *CreateApplicationPromptRequest) Validate() error {
	return dara.Validate(s)
}
