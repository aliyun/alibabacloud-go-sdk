// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMaskingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeMaskingRulesResponseBodyData) *DescribeMaskingRulesResponseBody
	GetData() *DescribeMaskingRulesResponseBodyData
	SetRequestId(v string) *DescribeMaskingRulesResponseBody
	GetRequestId() *string
}

type DescribeMaskingRulesResponseBody struct {
	Data *DescribeMaskingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 69779000-57A4-38F6-BF85-**********A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMaskingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesResponseBody) GetData() *DescribeMaskingRulesResponseBodyData {
	return s.Data
}

func (s *DescribeMaskingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMaskingRulesResponseBody) SetData(v *DescribeMaskingRulesResponseBodyData) *DescribeMaskingRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMaskingRulesResponseBody) SetRequestId(v string) *DescribeMaskingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMaskingRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMaskingRulesResponseBodyData struct {
	Rules []*DescribeMaskingRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeMaskingRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesResponseBodyData) GetRules() []*DescribeMaskingRulesResponseBodyDataRules {
	return s.Rules
}

func (s *DescribeMaskingRulesResponseBodyData) SetRules(v []*DescribeMaskingRulesResponseBodyDataRules) *DescribeMaskingRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *DescribeMaskingRulesResponseBodyData) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMaskingRulesResponseBodyDataRules struct {
	// example:
	//
	// aes-128-gcm
	DefaultAlgo *string `json:"DefaultAlgo,omitempty" xml:"DefaultAlgo,omitempty"`
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// [{"name": "aes-128-gcm"},
	//
	//         {"name":"sm4-128-gcm"}]
	MaskingAlgo *string                                              `json:"MaskingAlgo,omitempty" xml:"MaskingAlgo,omitempty"`
	RuleConfig  *DescribeMaskingRulesResponseBodyDataRulesRuleConfig `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty" type:"Struct"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeMaskingRulesResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesResponseBodyDataRules) GetDefaultAlgo() *string {
	return s.DefaultAlgo
}

func (s *DescribeMaskingRulesResponseBodyDataRules) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeMaskingRulesResponseBodyDataRules) GetMaskingAlgo() *string {
	return s.MaskingAlgo
}

func (s *DescribeMaskingRulesResponseBodyDataRules) GetRuleConfig() *DescribeMaskingRulesResponseBodyDataRulesRuleConfig {
	return s.RuleConfig
}

func (s *DescribeMaskingRulesResponseBodyDataRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeMaskingRulesResponseBodyDataRules) SetDefaultAlgo(v string) *DescribeMaskingRulesResponseBodyDataRules {
	s.DefaultAlgo = &v
	return s
}

func (s *DescribeMaskingRulesResponseBodyDataRules) SetEnabled(v string) *DescribeMaskingRulesResponseBodyDataRules {
	s.Enabled = &v
	return s
}

func (s *DescribeMaskingRulesResponseBodyDataRules) SetMaskingAlgo(v string) *DescribeMaskingRulesResponseBodyDataRules {
	s.MaskingAlgo = &v
	return s
}

func (s *DescribeMaskingRulesResponseBodyDataRules) SetRuleConfig(v *DescribeMaskingRulesResponseBodyDataRulesRuleConfig) *DescribeMaskingRulesResponseBodyDataRules {
	s.RuleConfig = v
	return s
}

func (s *DescribeMaskingRulesResponseBodyDataRules) SetRuleName(v string) *DescribeMaskingRulesResponseBodyDataRules {
	s.RuleName = &v
	return s
}

func (s *DescribeMaskingRulesResponseBodyDataRules) Validate() error {
	if s.RuleConfig != nil {
		if err := s.RuleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMaskingRulesResponseBodyDataRulesRuleConfig struct {
	Columns   []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeMaskingRulesResponseBodyDataRulesRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesResponseBodyDataRulesRuleConfig) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesResponseBodyDataRulesRuleConfig) GetColumns() []*string {
	return s.Columns
}

func (s *DescribeMaskingRulesResponseBodyDataRulesRuleConfig) GetDatabases() []*string {
	return s.Databases
}

func (s *DescribeMaskingRulesResponseBodyDataRulesRuleConfig) GetTables() []*string {
	return s.Tables
}

func (s *DescribeMaskingRulesResponseBodyDataRulesRuleConfig) SetColumns(v []*string) *DescribeMaskingRulesResponseBodyDataRulesRuleConfig {
	s.Columns = v
	return s
}

func (s *DescribeMaskingRulesResponseBodyDataRulesRuleConfig) SetDatabases(v []*string) *DescribeMaskingRulesResponseBodyDataRulesRuleConfig {
	s.Databases = v
	return s
}

func (s *DescribeMaskingRulesResponseBodyDataRulesRuleConfig) SetTables(v []*string) *DescribeMaskingRulesResponseBodyDataRulesRuleConfig {
	s.Tables = v
	return s
}

func (s *DescribeMaskingRulesResponseBodyDataRulesRuleConfig) Validate() error {
	return dara.Validate(s)
}
