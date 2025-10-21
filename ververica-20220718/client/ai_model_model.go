// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiModel interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *AiModel
	GetComment() *string
	SetInputSchema(v *Schema) *AiModel
	GetInputSchema() *Schema
	SetName(v string) *AiModel
	GetName() *string
	SetOptions(v map[string]*string) *AiModel
	GetOptions() map[string]*string
	SetOutputSchema(v *Schema) *AiModel
	GetOutputSchema() *Schema
}

type AiModel struct {
	Comment     *string `json:"comment,omitempty" xml:"comment,omitempty"`
	InputSchema *Schema `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Options      map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	OutputSchema *Schema            `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
}

func (s AiModel) String() string {
	return dara.Prettify(s)
}

func (s AiModel) GoString() string {
	return s.String()
}

func (s *AiModel) GetComment() *string {
	return s.Comment
}

func (s *AiModel) GetInputSchema() *Schema {
	return s.InputSchema
}

func (s *AiModel) GetName() *string {
	return s.Name
}

func (s *AiModel) GetOptions() map[string]*string {
	return s.Options
}

func (s *AiModel) GetOutputSchema() *Schema {
	return s.OutputSchema
}

func (s *AiModel) SetComment(v string) *AiModel {
	s.Comment = &v
	return s
}

func (s *AiModel) SetInputSchema(v *Schema) *AiModel {
	s.InputSchema = v
	return s
}

func (s *AiModel) SetName(v string) *AiModel {
	s.Name = &v
	return s
}

func (s *AiModel) SetOptions(v map[string]*string) *AiModel {
	s.Options = v
	return s
}

func (s *AiModel) SetOutputSchema(v *Schema) *AiModel {
	s.OutputSchema = v
	return s
}

func (s *AiModel) Validate() error {
	if s.InputSchema != nil {
		if err := s.InputSchema.Validate(); err != nil {
			return err
		}
	}
	if s.OutputSchema != nil {
		if err := s.OutputSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}
