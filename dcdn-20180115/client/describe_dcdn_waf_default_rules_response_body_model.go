// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDefaultRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeDcdnWafDefaultRulesResponseBodyContent) *DescribeDcdnWafDefaultRulesResponseBody
	GetContent() []*DescribeDcdnWafDefaultRulesResponseBodyContent
	SetRequestId(v string) *DescribeDcdnWafDefaultRulesResponseBody
	GetRequestId() *string
}

type DescribeDcdnWafDefaultRulesResponseBody struct {
	// The rule configurations.
	Content []*DescribeDcdnWafDefaultRulesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnWafDefaultRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDefaultRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDefaultRulesResponseBody) GetContent() []*DescribeDcdnWafDefaultRulesResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnWafDefaultRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafDefaultRulesResponseBody) SetContent(v []*DescribeDcdnWafDefaultRulesResponseBodyContent) *DescribeDcdnWafDefaultRulesResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBody) SetRequestId(v string) *DescribeDcdnWafDefaultRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafDefaultRulesResponseBodyContent struct {
	// The protection scenario. Valid values:
	//
	// 	- **waf_group**: basic web protection.
	//
	// 	- **anti_scan**: scan protection.
	//
	// example:
	//
	// anti_scan
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The rules.
	Rules []*DescribeDcdnWafDefaultRulesResponseBodyContentRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeDcdnWafDefaultRulesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDefaultRulesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContent) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContent) GetRules() []*DescribeDcdnWafDefaultRulesResponseBodyContentRules {
	return s.Rules
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContent) SetDefenseScene(v string) *DescribeDcdnWafDefaultRulesResponseBodyContent {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContent) SetRules(v []*DescribeDcdnWafDefaultRulesResponseBodyContentRules) *DescribeDcdnWafDefaultRulesResponseBodyContent {
	s.Rules = v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafDefaultRulesResponseBodyContentRules struct {
	// The default action of the rule. Valid values:
	//
	// 	- **monitor**
	//
	// 	- **deny**
	//
	// 	- **block**
	//
	// example:
	//
	// block
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The default configuration of the rule.
	//
	// example:
	//
	// {\\"wafGroupIds\\":\\"1012\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The default name of the rule.
	//
	// example:
	//
	// Default_WafGroup_Rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The default status of the rule. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **waf_group**: basic web protection.
	//
	// 	- **high_frequency**: high-frequency scanning blocking.
	//
	// 	- **directory_traversal**: directory traversal blocking.
	//
	// 	- **scan_tools**: scanner blocking.
	//
	// example:
	//
	// waf_group
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnWafDefaultRulesResponseBodyContentRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDefaultRulesResponseBodyContentRules) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) GetAction() *string {
	return s.Action
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) GetConfig() *string {
	return s.Config
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) GetName() *string {
	return s.Name
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) SetAction(v string) *DescribeDcdnWafDefaultRulesResponseBodyContentRules {
	s.Action = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) SetConfig(v string) *DescribeDcdnWafDefaultRulesResponseBodyContentRules {
	s.Config = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) SetName(v string) *DescribeDcdnWafDefaultRulesResponseBodyContentRules {
	s.Name = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) SetStatus(v string) *DescribeDcdnWafDefaultRulesResponseBodyContentRules {
	s.Status = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) SetType(v string) *DescribeDcdnWafDefaultRulesResponseBodyContentRules {
	s.Type = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponseBodyContentRules) Validate() error {
	return dara.Validate(s)
}
