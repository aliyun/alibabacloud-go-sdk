// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiStatisticsPathField interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *AiStatisticsPathField
	GetCategory() *string
	SetDescription(v string) *AiStatisticsPathField
	GetDescription() *string
	SetFieldKey(v string) *AiStatisticsPathField
	GetFieldKey() *string
	SetIo(v string) *AiStatisticsPathField
	GetIo() *string
	SetJsonPath(v string) *AiStatisticsPathField
	GetJsonPath() *string
	SetName(v string) *AiStatisticsPathField
	GetName() *string
	SetRecordEnabled(v bool) *AiStatisticsPathField
	GetRecordEnabled() *bool
	SetRule(v string) *AiStatisticsPathField
	GetRule() *string
	SetSensitive(v bool) *AiStatisticsPathField
	GetSensitive() *bool
	SetSource(v string) *AiStatisticsPathField
	GetSource() *string
}

type AiStatisticsPathField struct {
	Category      *string `json:"category,omitempty" xml:"category,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	FieldKey      *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	Io            *string `json:"io,omitempty" xml:"io,omitempty"`
	JsonPath      *string `json:"jsonPath,omitempty" xml:"jsonPath,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	RecordEnabled *bool   `json:"recordEnabled,omitempty" xml:"recordEnabled,omitempty"`
	Rule          *string `json:"rule,omitempty" xml:"rule,omitempty"`
	Sensitive     *bool   `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	Source        *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s AiStatisticsPathField) String() string {
	return dara.Prettify(s)
}

func (s AiStatisticsPathField) GoString() string {
	return s.String()
}

func (s *AiStatisticsPathField) GetCategory() *string {
	return s.Category
}

func (s *AiStatisticsPathField) GetDescription() *string {
	return s.Description
}

func (s *AiStatisticsPathField) GetFieldKey() *string {
	return s.FieldKey
}

func (s *AiStatisticsPathField) GetIo() *string {
	return s.Io
}

func (s *AiStatisticsPathField) GetJsonPath() *string {
	return s.JsonPath
}

func (s *AiStatisticsPathField) GetName() *string {
	return s.Name
}

func (s *AiStatisticsPathField) GetRecordEnabled() *bool {
	return s.RecordEnabled
}

func (s *AiStatisticsPathField) GetRule() *string {
	return s.Rule
}

func (s *AiStatisticsPathField) GetSensitive() *bool {
	return s.Sensitive
}

func (s *AiStatisticsPathField) GetSource() *string {
	return s.Source
}

func (s *AiStatisticsPathField) SetCategory(v string) *AiStatisticsPathField {
	s.Category = &v
	return s
}

func (s *AiStatisticsPathField) SetDescription(v string) *AiStatisticsPathField {
	s.Description = &v
	return s
}

func (s *AiStatisticsPathField) SetFieldKey(v string) *AiStatisticsPathField {
	s.FieldKey = &v
	return s
}

func (s *AiStatisticsPathField) SetIo(v string) *AiStatisticsPathField {
	s.Io = &v
	return s
}

func (s *AiStatisticsPathField) SetJsonPath(v string) *AiStatisticsPathField {
	s.JsonPath = &v
	return s
}

func (s *AiStatisticsPathField) SetName(v string) *AiStatisticsPathField {
	s.Name = &v
	return s
}

func (s *AiStatisticsPathField) SetRecordEnabled(v bool) *AiStatisticsPathField {
	s.RecordEnabled = &v
	return s
}

func (s *AiStatisticsPathField) SetRule(v string) *AiStatisticsPathField {
	s.Rule = &v
	return s
}

func (s *AiStatisticsPathField) SetSensitive(v bool) *AiStatisticsPathField {
	s.Sensitive = &v
	return s
}

func (s *AiStatisticsPathField) SetSource(v string) *AiStatisticsPathField {
	s.Source = &v
	return s
}

func (s *AiStatisticsPathField) Validate() error {
	return dara.Validate(s)
}
