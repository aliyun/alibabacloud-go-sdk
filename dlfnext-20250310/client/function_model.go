// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunction interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Function
	GetComment() *string
	SetCreatedAt(v int64) *Function
	GetCreatedAt() *int64
	SetCreatedBy(v string) *Function
	GetCreatedBy() *string
	SetDefinitions(v map[string]*FunctionDefinition) *Function
	GetDefinitions() map[string]*FunctionDefinition
	SetDeterministic(v bool) *Function
	GetDeterministic() *bool
	SetId(v string) *Function
	GetId() *string
	SetInputParams(v []*DataField) *Function
	GetInputParams() []*DataField
	SetName(v string) *Function
	GetName() *string
	SetOptions(v map[string]*string) *Function
	GetOptions() map[string]*string
	SetOwner(v string) *Function
	GetOwner() *string
	SetReturnParams(v []*DataField) *Function
	GetReturnParams() []*DataField
	SetUpdatedAt(v int64) *Function
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Function
	GetUpdatedBy() *string
}

type Function struct {
	Comment       *string                        `json:"comment,omitempty" xml:"comment,omitempty"`
	CreatedAt     *int64                         `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy     *string                        `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Definitions   map[string]*FunctionDefinition `json:"definitions,omitempty" xml:"definitions,omitempty"`
	Deterministic *bool                          `json:"deterministic,omitempty" xml:"deterministic,omitempty"`
	Id            *string                        `json:"id,omitempty" xml:"id,omitempty"`
	InputParams   []*DataField                   `json:"inputParams,omitempty" xml:"inputParams,omitempty" type:"Repeated"`
	Name          *string                        `json:"name,omitempty" xml:"name,omitempty"`
	Options       map[string]*string             `json:"options,omitempty" xml:"options,omitempty"`
	Owner         *string                        `json:"owner,omitempty" xml:"owner,omitempty"`
	ReturnParams  []*DataField                   `json:"returnParams,omitempty" xml:"returnParams,omitempty" type:"Repeated"`
	UpdatedAt     *int64                         `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy     *string                        `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Function) String() string {
	return dara.Prettify(s)
}

func (s Function) GoString() string {
	return s.String()
}

func (s *Function) GetComment() *string {
	return s.Comment
}

func (s *Function) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Function) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Function) GetDefinitions() map[string]*FunctionDefinition {
	return s.Definitions
}

func (s *Function) GetDeterministic() *bool {
	return s.Deterministic
}

func (s *Function) GetId() *string {
	return s.Id
}

func (s *Function) GetInputParams() []*DataField {
	return s.InputParams
}

func (s *Function) GetName() *string {
	return s.Name
}

func (s *Function) GetOptions() map[string]*string {
	return s.Options
}

func (s *Function) GetOwner() *string {
	return s.Owner
}

func (s *Function) GetReturnParams() []*DataField {
	return s.ReturnParams
}

func (s *Function) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Function) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Function) SetComment(v string) *Function {
	s.Comment = &v
	return s
}

func (s *Function) SetCreatedAt(v int64) *Function {
	s.CreatedAt = &v
	return s
}

func (s *Function) SetCreatedBy(v string) *Function {
	s.CreatedBy = &v
	return s
}

func (s *Function) SetDefinitions(v map[string]*FunctionDefinition) *Function {
	s.Definitions = v
	return s
}

func (s *Function) SetDeterministic(v bool) *Function {
	s.Deterministic = &v
	return s
}

func (s *Function) SetId(v string) *Function {
	s.Id = &v
	return s
}

func (s *Function) SetInputParams(v []*DataField) *Function {
	s.InputParams = v
	return s
}

func (s *Function) SetName(v string) *Function {
	s.Name = &v
	return s
}

func (s *Function) SetOptions(v map[string]*string) *Function {
	s.Options = v
	return s
}

func (s *Function) SetOwner(v string) *Function {
	s.Owner = &v
	return s
}

func (s *Function) SetReturnParams(v []*DataField) *Function {
	s.ReturnParams = v
	return s
}

func (s *Function) SetUpdatedAt(v int64) *Function {
	s.UpdatedAt = &v
	return s
}

func (s *Function) SetUpdatedBy(v string) *Function {
	s.UpdatedBy = &v
	return s
}

func (s *Function) Validate() error {
	if s.InputParams != nil {
		for _, item := range s.InputParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReturnParams != nil {
		for _, item := range s.ReturnParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
