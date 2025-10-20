// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunctionChange interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *FunctionChange
	GetAction() *string
	SetComment(v string) *FunctionChange
	GetComment() *string
	SetDefinition(v *FunctionDefinition) *FunctionChange
	GetDefinition() *FunctionDefinition
	SetKey(v string) *FunctionChange
	GetKey() *string
	SetName(v string) *FunctionChange
	GetName() *string
	SetValue(v string) *FunctionChange
	GetValue() *string
}

type FunctionChange struct {
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// required in UpdateFunctionComment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// required in AddDefinition/UpdateDefinition
	Definition *FunctionDefinition `json:"definition,omitempty" xml:"definition,omitempty"`
	// required in SetFunctionOption/RemoveFunctionOption
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// required in AddDefinition/UpdateDefinition/DropDefinition
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// required in SetFunctionOption
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FunctionChange) String() string {
	return dara.Prettify(s)
}

func (s FunctionChange) GoString() string {
	return s.String()
}

func (s *FunctionChange) GetAction() *string {
	return s.Action
}

func (s *FunctionChange) GetComment() *string {
	return s.Comment
}

func (s *FunctionChange) GetDefinition() *FunctionDefinition {
	return s.Definition
}

func (s *FunctionChange) GetKey() *string {
	return s.Key
}

func (s *FunctionChange) GetName() *string {
	return s.Name
}

func (s *FunctionChange) GetValue() *string {
	return s.Value
}

func (s *FunctionChange) SetAction(v string) *FunctionChange {
	s.Action = &v
	return s
}

func (s *FunctionChange) SetComment(v string) *FunctionChange {
	s.Comment = &v
	return s
}

func (s *FunctionChange) SetDefinition(v *FunctionDefinition) *FunctionChange {
	s.Definition = v
	return s
}

func (s *FunctionChange) SetKey(v string) *FunctionChange {
	s.Key = &v
	return s
}

func (s *FunctionChange) SetName(v string) *FunctionChange {
	s.Name = &v
	return s
}

func (s *FunctionChange) SetValue(v string) *FunctionChange {
	s.Value = &v
	return s
}

func (s *FunctionChange) Validate() error {
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
