// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafTemplateRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWafTemplateRulesResponseBody
	GetRequestId() *string
	SetRules(v []*ListWafTemplateRulesResponseBodyRules) *ListWafTemplateRulesResponseBody
	GetRules() []*ListWafTemplateRulesResponseBodyRules
}

type ListWafTemplateRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of template rules.
	Rules []*ListWafTemplateRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ListWafTemplateRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWafTemplateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWafTemplateRulesResponseBody) GetRules() []*ListWafTemplateRulesResponseBodyRules {
	return s.Rules
}

func (s *ListWafTemplateRulesResponseBody) SetRequestId(v string) *ListWafTemplateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafTemplateRulesResponseBody) SetRules(v []*ListWafTemplateRulesResponseBodyRules) *ListWafTemplateRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListWafTemplateRulesResponseBody) Validate() error {
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

type ListWafTemplateRulesResponseBodyRules struct {
	// The rule configuration.
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// The rule name.
	//
	// example:
	//
	// HTTP Directory Traversal Rule [Template]
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The WAF running phase.
	//
	// example:
	//
	// http_anti_scan
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The rule status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The rule type.
	//
	// example:
	//
	// http_directory_traversal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWafTemplateRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListWafTemplateRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesResponseBodyRules) GetConfig() *WafRuleConfig {
	return s.Config
}

func (s *ListWafTemplateRulesResponseBodyRules) GetName() *string {
	return s.Name
}

func (s *ListWafTemplateRulesResponseBodyRules) GetPhase() *string {
	return s.Phase
}

func (s *ListWafTemplateRulesResponseBodyRules) GetStatus() *string {
	return s.Status
}

func (s *ListWafTemplateRulesResponseBodyRules) GetType() *string {
	return s.Type
}

func (s *ListWafTemplateRulesResponseBodyRules) SetConfig(v *WafRuleConfig) *ListWafTemplateRulesResponseBodyRules {
	s.Config = v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) SetName(v string) *ListWafTemplateRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) SetPhase(v string) *ListWafTemplateRulesResponseBodyRules {
	s.Phase = &v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) SetStatus(v string) *ListWafTemplateRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) SetType(v string) *ListWafTemplateRulesResponseBodyRules {
	s.Type = &v
	return s
}

func (s *ListWafTemplateRulesResponseBodyRules) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}
