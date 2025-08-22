// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSecSpecInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDcdnSecSpecInfoResponseBody
	GetRequestId() *string
	SetSpecInfos(v []*DescribeDcdnSecSpecInfoResponseBodySpecInfos) *DescribeDcdnSecSpecInfoResponseBody
	GetSpecInfos() []*DescribeDcdnSecSpecInfoResponseBodySpecInfos
	SetVersion(v string) *DescribeDcdnSecSpecInfoResponseBody
	GetVersion() *string
}

type DescribeDcdnSecSpecInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 32f6cbb7-13e5-403a-9941-4d4e978dd227
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code and configurations of the security rules.
	SpecInfos []*DescribeDcdnSecSpecInfoResponseBodySpecInfos `json:"SpecInfos,omitempty" xml:"SpecInfos,omitempty" type:"Repeated"`
	// The version of secure DCDN.
	//
	// example:
	//
	// enterprise
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDcdnSecSpecInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnSecSpecInfoResponseBody) GetSpecInfos() []*DescribeDcdnSecSpecInfoResponseBodySpecInfos {
	return s.SpecInfos
}

func (s *DescribeDcdnSecSpecInfoResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeDcdnSecSpecInfoResponseBody) SetRequestId(v string) *DescribeDcdnSecSpecInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBody) SetSpecInfos(v []*DescribeDcdnSecSpecInfoResponseBodySpecInfos) *DescribeDcdnSecSpecInfoResponseBody {
	s.SpecInfos = v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBody) SetVersion(v string) *DescribeDcdnSecSpecInfoResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSecSpecInfoResponseBodySpecInfos struct {
	// The code of the security rule.
	//
	// example:
	//
	// accurate_***
	RuleCode *string `json:"RuleCode,omitempty" xml:"RuleCode,omitempty"`
	// The configurations of the security rule.
	RuleConfigs []*DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs `json:"RuleConfigs,omitempty" xml:"RuleConfigs,omitempty" type:"Repeated"`
}

func (s DescribeDcdnSecSpecInfoResponseBodySpecInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoResponseBodySpecInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfos) GetRuleCode() *string {
	return s.RuleCode
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfos) GetRuleConfigs() []*DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs {
	return s.RuleConfigs
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfos) SetRuleCode(v string) *DescribeDcdnSecSpecInfoResponseBodySpecInfos {
	s.RuleCode = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfos) SetRuleConfigs(v []*DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) *DescribeDcdnSecSpecInfoResponseBodySpecInfos {
	s.RuleConfigs = v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs struct {
	// The configuration code of the security rule.
	//
	// example:
	//
	// custom_****_number
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration expression of the security rule.
	//
	// example:
	//
	// equal
	Expr *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// The value of the configuration expression of the security rule.
	//
	// example:
	//
	// 20
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) GetCode() *string {
	return s.Code
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) GetExpr() *string {
	return s.Expr
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) SetCode(v string) *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs {
	s.Code = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) SetExpr(v string) *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs {
	s.Expr = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) SetValue(v string) *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs {
	s.Value = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) Validate() error {
	return dara.Validate(s)
}
