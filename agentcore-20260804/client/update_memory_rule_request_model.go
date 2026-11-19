// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateMemoryRuleRequestBody) *UpdateMemoryRuleRequest
	GetBody() *UpdateMemoryRuleRequestBody
}

type UpdateMemoryRuleRequest struct {
	// The JSON request body.
	//
	// This parameter is required.
	Body *UpdateMemoryRuleRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateMemoryRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRuleRequest) GetBody() *UpdateMemoryRuleRequestBody {
	return s.Body
}

func (s *UpdateMemoryRuleRequest) SetBody(v *UpdateMemoryRuleRequestBody) *UpdateMemoryRuleRequest {
	s.Body = v
	return s
}

func (s *UpdateMemoryRuleRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryRuleRequestBody struct {
	// description
	//
	// example:
	//
	// Extract long-term preferences explicitly expressed by the user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// enabled
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// name
	//
	// example:
	//
	// Long-term preferences
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The extraction strategy. Uses partial merging, where omitted fields retain their original values.
	Strategy *UpdateMemoryRuleRequestBodyStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
}

func (s UpdateMemoryRuleRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRuleRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRuleRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateMemoryRuleRequestBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMemoryRuleRequestBody) GetName() *string {
	return s.Name
}

func (s *UpdateMemoryRuleRequestBody) GetStrategy() *UpdateMemoryRuleRequestBodyStrategy {
	return s.Strategy
}

func (s *UpdateMemoryRuleRequestBody) SetDescription(v string) *UpdateMemoryRuleRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateMemoryRuleRequestBody) SetEnabled(v bool) *UpdateMemoryRuleRequestBody {
	s.Enabled = &v
	return s
}

func (s *UpdateMemoryRuleRequestBody) SetName(v string) *UpdateMemoryRuleRequestBody {
	s.Name = &v
	return s
}

func (s *UpdateMemoryRuleRequestBody) SetStrategy(v *UpdateMemoryRuleRequestBodyStrategy) *UpdateMemoryRuleRequestBody {
	s.Strategy = v
	return s
}

func (s *UpdateMemoryRuleRequestBody) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryRuleRequestBodyStrategy struct {
	// The custom extraction instruction. If omitted, the original value is retained. An empty string clears the value.
	//
	// example:
	//
	// Retain only long-term preferences explicitly expressed by the user
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// type
	//
	// example:
	//
	// auto
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateMemoryRuleRequestBodyStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRuleRequestBodyStrategy) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRuleRequestBodyStrategy) GetInstruction() *string {
	return s.Instruction
}

func (s *UpdateMemoryRuleRequestBodyStrategy) GetType() *string {
	return s.Type
}

func (s *UpdateMemoryRuleRequestBodyStrategy) SetInstruction(v string) *UpdateMemoryRuleRequestBodyStrategy {
	s.Instruction = &v
	return s
}

func (s *UpdateMemoryRuleRequestBodyStrategy) SetType(v string) *UpdateMemoryRuleRequestBodyStrategy {
	s.Type = &v
	return s
}

func (s *UpdateMemoryRuleRequestBodyStrategy) Validate() error {
	return dara.Validate(s)
}
