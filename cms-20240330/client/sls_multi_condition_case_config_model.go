// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSlsMultiConditionCaseConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *SlsMultiConditionCaseConfig
	GetCondition() *string
	SetCountCondition(v string) *SlsMultiConditionCaseConfig
	GetCountCondition() *string
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
	SetOperator(v string) *SlsMultiConditionCaseConfig
	GetOperator() *string
	SetRawCondition(v string) *SlsMultiConditionCaseConfig
	GetRawCondition() *string
	SetSeverity(v string) *SlsMultiConditionCaseConfig
	GetSeverity() *string
}

type SlsMultiConditionCaseConfig struct {
	// The match expression (corresponds to V1 condition, preserved as-is without structured parsing).
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The count match expression (corresponds to V1 countCondition, preserved as-is without structured parsing).
	CountCondition *string `json:"countCondition,omitempty" xml:"countCondition,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The write path is disabled. Use countCondition instead.
	CountOperator *string `json:"countOperator,omitempty" xml:"countOperator,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The write path is disabled. Use countCondition instead.
	CountThreshold *int64 `json:"countThreshold,omitempty" xml:"countThreshold,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The write path is disabled. Use condition instead.
	MatchField *string `json:"matchField,omitempty" xml:"matchField,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The write path is disabled. Use condition instead.
	MatchOperator *string `json:"matchOperator,omitempty" xml:"matchOperator,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The write path is disabled. Use condition instead.
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
	// The detection operator (aligned with V1 caseList.type): HAS_DATA / HAS_DATA_COUNT / HAS_DATA_MATCH / HAS_DATA_MATCH_COUNT.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The write path is disabled. Use condition instead.
	RawCondition *string `json:"rawCondition,omitempty" xml:"rawCondition,omitempty"`
	// The severity level (corresponds to V1 level).
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
}

func (s SlsMultiConditionCaseConfig) String() string {
	return dara.Prettify(s)
}

func (s SlsMultiConditionCaseConfig) GoString() string {
	return s.String()
}

func (s *SlsMultiConditionCaseConfig) GetCondition() *string {
	return s.Condition
}

func (s *SlsMultiConditionCaseConfig) GetCountCondition() *string {
	return s.CountCondition
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

func (s *SlsMultiConditionCaseConfig) GetOperator() *string {
	return s.Operator
}

func (s *SlsMultiConditionCaseConfig) GetRawCondition() *string {
	return s.RawCondition
}

func (s *SlsMultiConditionCaseConfig) GetSeverity() *string {
	return s.Severity
}

func (s *SlsMultiConditionCaseConfig) SetCondition(v string) *SlsMultiConditionCaseConfig {
	s.Condition = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) SetCountCondition(v string) *SlsMultiConditionCaseConfig {
	s.CountCondition = &v
	return s
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

func (s *SlsMultiConditionCaseConfig) SetOperator(v string) *SlsMultiConditionCaseConfig {
	s.Operator = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) SetRawCondition(v string) *SlsMultiConditionCaseConfig {
	s.RawCondition = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) SetSeverity(v string) *SlsMultiConditionCaseConfig {
	s.Severity = &v
	return s
}

func (s *SlsMultiConditionCaseConfig) Validate() error {
	return dara.Validate(s)
}
