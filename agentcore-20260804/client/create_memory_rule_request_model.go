// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateMemoryRuleRequestBody) *CreateMemoryRuleRequest
	GetBody() *CreateMemoryRuleRequestBody
}

type CreateMemoryRuleRequest struct {
	// The JSON request body.
	//
	// This parameter is required.
	Body *CreateMemoryRuleRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateMemoryRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryRuleRequest) GetBody() *CreateMemoryRuleRequestBody {
	return s.Body
}

func (s *CreateMemoryRuleRequest) SetBody(v *CreateMemoryRuleRequestBody) *CreateMemoryRuleRequest {
	s.Body = v
	return s
}

func (s *CreateMemoryRuleRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryRuleRequestBody struct {
	// description
	//
	// example:
	//
	// Extract long-term preferences explicitly expressed by users
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// enabled
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// name
	//
	// This parameter is required.
	//
	// example:
	//
	// Long-term preferences
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// strategy
	//
	// This parameter is required.
	Strategy *CreateMemoryRuleRequestBodyStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
}

func (s CreateMemoryRuleRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRuleRequestBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryRuleRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateMemoryRuleRequestBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMemoryRuleRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateMemoryRuleRequestBody) GetStrategy() *CreateMemoryRuleRequestBodyStrategy {
	return s.Strategy
}

func (s *CreateMemoryRuleRequestBody) SetDescription(v string) *CreateMemoryRuleRequestBody {
	s.Description = &v
	return s
}

func (s *CreateMemoryRuleRequestBody) SetEnabled(v bool) *CreateMemoryRuleRequestBody {
	s.Enabled = &v
	return s
}

func (s *CreateMemoryRuleRequestBody) SetName(v string) *CreateMemoryRuleRequestBody {
	s.Name = &v
	return s
}

func (s *CreateMemoryRuleRequestBody) SetStrategy(v *CreateMemoryRuleRequestBodyStrategy) *CreateMemoryRuleRequestBody {
	s.Strategy = v
	return s
}

func (s *CreateMemoryRuleRequestBody) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryRuleRequestBodyStrategy struct {
	// The strategy-level instruction. It takes effect together with the store-level instruction. If a conflict occurs, this field takes precedence.
	//
	// example:
	//
	// Retain only long-term preferences explicitly expressed by users
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// type
	//
	// This parameter is required.
	//
	// example:
	//
	// auto
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMemoryRuleRequestBodyStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRuleRequestBodyStrategy) GoString() string {
	return s.String()
}

func (s *CreateMemoryRuleRequestBodyStrategy) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateMemoryRuleRequestBodyStrategy) GetType() *string {
	return s.Type
}

func (s *CreateMemoryRuleRequestBodyStrategy) SetInstruction(v string) *CreateMemoryRuleRequestBodyStrategy {
	s.Instruction = &v
	return s
}

func (s *CreateMemoryRuleRequestBodyStrategy) SetType(v string) *CreateMemoryRuleRequestBodyStrategy {
	s.Type = &v
	return s
}

func (s *CreateMemoryRuleRequestBodyStrategy) Validate() error {
	return dara.Validate(s)
}
