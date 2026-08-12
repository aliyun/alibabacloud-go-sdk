// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSlsMultiConditionCaseConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCountOperator(v string) *SlsMultiConditionCaseConfig
	GetCountOperator() *string
	SetCountThreshold(v int64) *SlsMultiConditionCaseConfig
	GetCountThreshold() *int64
	SetMatchField(v string) *SlsMultiConditionCaseConfig
	GetMatchField() *string
	SetMatchOperator(v string) *SlsMultiConditionCaseConfig
	GetMatchOperator() *string
	SetMatchValue(v string) *SlsMultiConditionCaseConfig
	GetMatchValue() *string
	SetSeverity(v string) *SlsMultiConditionCaseConfig
	GetSeverity() *string
}

type SlsMultiConditionCaseConfig struct {
	// The count comparison operator. Valid values: GTE, GT, EQ, LTE, LT.
	CountOperator *string `json:"countOperator,omitempty" xml:"countOperator,omitempty"`
	// The count threshold. The alert is triggered when this threshold is met.
	CountThreshold *int64 `json:"countThreshold,omitempty" xml:"countThreshold,omitempty"`
	// The log field name. Required when matchOperator is set to CONTAINS, EQUALS, or REGEX. Specify the field name when matchOperator is set to PRESENT or NOT_PRESENT.
	MatchField *string `json:"matchField,omitempty" xml:"matchField,omitempty"`
	// The log matching operator. Valid values: PRESENT (field exists), NOT_PRESENT (field does not exist), CONTAINS (contains), EQUALS (equals), REGEX (regular expression). If left empty, any data matches.
	MatchOperator *string `json:"matchOperator,omitempty" xml:"matchOperator,omitempty"`
	// The log match value. Required when matchOperator is set to CONTAINS, EQUALS, or REGEX.
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
	// The severity level.
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
}

func (s SlsMultiConditionCaseConfig) String() string {
	return dara.Prettify(s)
}

func (s SlsMultiConditionCaseConfig) GoString() string {
	return s.String()
}

func (s *SlsMultiConditionCaseConfig) GetCountOperator() *string {
	return s.CountOperator
}

func (s *SlsMultiConditionCaseConfig) GetCountThreshold() *int64 {
	return s.CountThreshold
}

func (s *SlsMultiConditionCaseConfig) GetMatchField() *string {
	return s.MatchField
}

func (s *SlsMultiConditionCaseConfig) GetMatchOperator() *string {
	return s.MatchOperator
}

func (s *SlsMultiConditionCaseConfig) GetMatchValue() *string {
	return s.MatchValue
}

func (s *SlsMultiConditionCaseConfig) GetSeverity() *string {
	return s.Severity
}

func (s *SlsMultiConditionCaseConfig) SetCountOperator(v string) *SlsMultiConditionCaseConfig {
	s.CountOperator = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) SetCountThreshold(v int64) *SlsMultiConditionCaseConfig {
	s.CountThreshold = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) SetMatchField(v string) *SlsMultiConditionCaseConfig {
	s.MatchField = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) SetMatchOperator(v string) *SlsMultiConditionCaseConfig {
	s.MatchOperator = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) SetMatchValue(v string) *SlsMultiConditionCaseConfig {
	s.MatchValue = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) SetSeverity(v string) *SlsMultiConditionCaseConfig {
	s.Severity = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) Validate() error {
	return dara.Validate(s)
}
