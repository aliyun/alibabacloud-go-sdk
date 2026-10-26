// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomScriptTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateCustomScriptTemplateRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCustomScriptTemplateRequest
	GetInstanceId() *string
	SetLabelsJson(v string) *CreateCustomScriptTemplateRequest
	GetLabelsJson() *string
	SetModel(v string) *CreateCustomScriptTemplateRequest
	GetModel() *string
	SetName(v string) *CreateCustomScriptTemplateRequest
	GetName() *string
	SetPrompt(v string) *CreateCustomScriptTemplateRequest
	GetPrompt() *string
	SetPromptJson(v string) *CreateCustomScriptTemplateRequest
	GetPromptJson() *string
	SetScriptId(v string) *CreateCustomScriptTemplateRequest
	GetScriptId() *string
}

type CreateCustomScriptTemplateRequest struct {
	// The description of the custom scenario.
	//
	// example:
	//
	// This is a large language model bot
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 0ec0c897-b92c-40e4-9ad7-e6e4f5ce13bb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The label descriptions in JSON array format.
	//
	// example:
	//
	// [{\\"name\\":\\"Delivered\\",\\"description\\":\\"Whether the purchased home appliance has been delivered\\",\\"valueList\\":\\"[\\\\\\"Yes\\\\\\",\\\\\\"No\\\\\\"]\\"},{\\"name\\":\\"On-site Appointment Time\\",\\"description\\":\\"Collect the customer\\"s expected on-site installation time\\",\\"valueList\\":\\"[]\\"}]
	LabelsJson *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// The ID of the model used to create the template.
	//
	// example:
	//
	// model_002
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name of the custom scenario.
	//
	// example:
	//
	// Test scenario
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The prompt in professional mode.
	//
	// example:
	//
	// \\nName: Annie \\nRole: Game Promoter\\nBackground: Promote the latest benefit activities to players over the phone, ensuring that every player stays informed and enjoys the game.\\nSkills: As a game customer service representative, recommend time-limited game benefit activities to users, highlight the benefits, and explain the participation process.\\nConstraints: You must remember that you are Annie,
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The agent configuration in JSON format.
	//
	// example:
	//
	// {"name":"Xiao X","gender":"Male","age":18,"role":"Game Promoter","communicationStyle":["Friendly"],"goals":"Hello","background":"Not very good","openingPrompt":"Hello, I am XXX"}
	PromptJson *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 43972417-0657-452a-81c2-c59d8245a9b2
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateCustomScriptTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomScriptTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomScriptTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomScriptTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomScriptTemplateRequest) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *CreateCustomScriptTemplateRequest) GetModel() *string {
	return s.Model
}

func (s *CreateCustomScriptTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomScriptTemplateRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateCustomScriptTemplateRequest) GetPromptJson() *string {
	return s.PromptJson
}

func (s *CreateCustomScriptTemplateRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateCustomScriptTemplateRequest) SetDescription(v string) *CreateCustomScriptTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomScriptTemplateRequest) SetInstanceId(v string) *CreateCustomScriptTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomScriptTemplateRequest) SetLabelsJson(v string) *CreateCustomScriptTemplateRequest {
	s.LabelsJson = &v
	return s
}

func (s *CreateCustomScriptTemplateRequest) SetModel(v string) *CreateCustomScriptTemplateRequest {
	s.Model = &v
	return s
}

func (s *CreateCustomScriptTemplateRequest) SetName(v string) *CreateCustomScriptTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomScriptTemplateRequest) SetPrompt(v string) *CreateCustomScriptTemplateRequest {
	s.Prompt = &v
	return s
}

func (s *CreateCustomScriptTemplateRequest) SetPromptJson(v string) *CreateCustomScriptTemplateRequest {
	s.PromptJson = &v
	return s
}

func (s *CreateCustomScriptTemplateRequest) SetScriptId(v string) *CreateCustomScriptTemplateRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateCustomScriptTemplateRequest) Validate() error {
	return dara.Validate(s)
}
