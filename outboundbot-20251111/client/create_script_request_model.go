// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateScriptRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateScriptRequest
	GetInstanceId() *string
	SetName(v string) *CreateScriptRequest
	GetName() *string
	SetNluEngine(v string) *CreateScriptRequest
	GetNluEngine() *string
}

type CreateScriptRequest struct {
	// The description.
	//
	// example:
	//
	// For testing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// Test scenario
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The NLU engine type. Valid values:
	//
	// - BEEBOT: integrates with Chatbot.
	//
	// - PROMPTS: prompt mode.
	//
	// - FUNCTION: integrates with Function Compute.
	//
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
}

func (s CreateScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptRequest) GoString() string {
	return s.String()
}

func (s *CreateScriptRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScriptRequest) GetName() *string {
	return s.Name
}

func (s *CreateScriptRequest) GetNluEngine() *string {
	return s.NluEngine
}

func (s *CreateScriptRequest) SetDescription(v string) *CreateScriptRequest {
	s.Description = &v
	return s
}

func (s *CreateScriptRequest) SetInstanceId(v string) *CreateScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScriptRequest) SetName(v string) *CreateScriptRequest {
	s.Name = &v
	return s
}

func (s *CreateScriptRequest) SetNluEngine(v string) *CreateScriptRequest {
	s.NluEngine = &v
	return s
}

func (s *CreateScriptRequest) Validate() error {
	return dara.Validate(s)
}
