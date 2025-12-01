// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CompareRuleRequest
	GetLang() *string
	SetPreviousRuleVersionId(v int64) *CompareRuleRequest
	GetPreviousRuleVersionId() *int64
	SetRegId(v string) *CompareRuleRequest
	GetRegId() *string
	SetRuleVersionId(v int64) *CompareRuleRequest
	GetRuleVersionId() *int64
}

type CompareRuleRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary key ID of the previous policy version.
	//
	// example:
	//
	// 11518
	PreviousRuleVersionId *int64 `json:"previousRuleVersionId,omitempty" xml:"previousRuleVersionId,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Primary key ID of the policy version.
	//
	// example:
	//
	// 11519
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
}

func (s CompareRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareRuleRequest) GoString() string {
	return s.String()
}

func (s *CompareRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *CompareRuleRequest) GetPreviousRuleVersionId() *int64 {
	return s.PreviousRuleVersionId
}

func (s *CompareRuleRequest) GetRegId() *string {
	return s.RegId
}

func (s *CompareRuleRequest) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *CompareRuleRequest) SetLang(v string) *CompareRuleRequest {
	s.Lang = &v
	return s
}

func (s *CompareRuleRequest) SetPreviousRuleVersionId(v int64) *CompareRuleRequest {
	s.PreviousRuleVersionId = &v
	return s
}

func (s *CompareRuleRequest) SetRegId(v string) *CompareRuleRequest {
	s.RegId = &v
	return s
}

func (s *CompareRuleRequest) SetRuleVersionId(v int64) *CompareRuleRequest {
	s.RuleVersionId = &v
	return s
}

func (s *CompareRuleRequest) Validate() error {
	return dara.Validate(s)
}
