// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchemaIndexesSearchFieldsValue interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyzer(v string) *SchemaIndexesSearchFieldsValue
	GetAnalyzer() *string
	SetAnalyzerType(v string) *SchemaIndexesSearchFieldsValue
	GetAnalyzerType() *string
	SetAnalyzerGeneration(v string) *SchemaIndexesSearchFieldsValue
	GetAnalyzerGeneration() *string
	SetFields(v []*string) *SchemaIndexesSearchFieldsValue
	GetFields() []*string
	SetLabel(v string) *SchemaIndexesSearchFieldsValue
	GetLabel() *string
}

type SchemaIndexesSearchFieldsValue struct {
	Analyzer           *string   `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	AnalyzerType       *string   `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	AnalyzerGeneration *string   `json:"analyzerGeneration,omitempty" xml:"analyzerGeneration,omitempty"`
	Fields             []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Label              *string   `json:"label,omitempty" xml:"label,omitempty"`
}

func (s SchemaIndexesSearchFieldsValue) String() string {
	return dara.Prettify(s)
}

func (s SchemaIndexesSearchFieldsValue) GoString() string {
	return s.String()
}

func (s *SchemaIndexesSearchFieldsValue) GetAnalyzer() *string {
	return s.Analyzer
}

func (s *SchemaIndexesSearchFieldsValue) GetAnalyzerType() *string {
	return s.AnalyzerType
}

func (s *SchemaIndexesSearchFieldsValue) GetAnalyzerGeneration() *string {
	return s.AnalyzerGeneration
}

func (s *SchemaIndexesSearchFieldsValue) GetFields() []*string {
	return s.Fields
}

func (s *SchemaIndexesSearchFieldsValue) GetLabel() *string {
	return s.Label
}

func (s *SchemaIndexesSearchFieldsValue) SetAnalyzer(v string) *SchemaIndexesSearchFieldsValue {
	s.Analyzer = &v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) SetAnalyzerType(v string) *SchemaIndexesSearchFieldsValue {
	s.AnalyzerType = &v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) SetAnalyzerGeneration(v string) *SchemaIndexesSearchFieldsValue {
	s.AnalyzerGeneration = &v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) SetFields(v []*string) *SchemaIndexesSearchFieldsValue {
	s.Fields = v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) SetLabel(v string) *SchemaIndexesSearchFieldsValue {
	s.Label = &v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) Validate() error {
	return dara.Validate(s)
}
