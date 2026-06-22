// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartClusterRule interface {
	dara.Model
	String() string
	GoString() string
	SetBaseURIs(v []*string) *SmartClusterRule
	GetBaseURIs() []*string
	SetKeywords(v []*string) *SmartClusterRule
	GetKeywords() []*string
	SetRuleType(v string) *SmartClusterRule
	GetRuleType() *string
	SetSensitivity(v float32) *SmartClusterRule
	GetSensitivity() *float32
}

type SmartClusterRule struct {
	BaseURIs []*string `json:"BaseURIs,omitempty" xml:"BaseURIs,omitempty" type:"Repeated"`
	// An array of keywords for clustering.
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	RuleType *string   `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The sensitivity for clustering, ranging from 0.0 to 1.0.
	//
	// example:
	//
	// 0.5
	Sensitivity *float32 `json:"Sensitivity,omitempty" xml:"Sensitivity,omitempty"`
}

func (s SmartClusterRule) String() string {
	return dara.Prettify(s)
}

func (s SmartClusterRule) GoString() string {
	return s.String()
}

func (s *SmartClusterRule) GetBaseURIs() []*string {
	return s.BaseURIs
}

func (s *SmartClusterRule) GetKeywords() []*string {
	return s.Keywords
}

func (s *SmartClusterRule) GetRuleType() *string {
	return s.RuleType
}

func (s *SmartClusterRule) GetSensitivity() *float32 {
	return s.Sensitivity
}

func (s *SmartClusterRule) SetBaseURIs(v []*string) *SmartClusterRule {
	s.BaseURIs = v
	return s
}

func (s *SmartClusterRule) SetKeywords(v []*string) *SmartClusterRule {
	s.Keywords = v
	return s
}

func (s *SmartClusterRule) SetRuleType(v string) *SmartClusterRule {
	s.RuleType = &v
	return s
}

func (s *SmartClusterRule) SetSensitivity(v float32) *SmartClusterRule {
	s.Sensitivity = &v
	return s
}

func (s *SmartClusterRule) Validate() error {
	return dara.Validate(s)
}
