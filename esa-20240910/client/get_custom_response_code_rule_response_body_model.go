// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomResponseCodeRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetCustomResponseCodeRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetCustomResponseCodeRuleResponseBody
	GetConfigType() *string
	SetPageId(v string) *GetCustomResponseCodeRuleResponseBody
	GetPageId() *string
	SetRequestId(v string) *GetCustomResponseCodeRuleResponseBody
	GetRequestId() *string
	SetReturnCode(v string) *GetCustomResponseCodeRuleResponseBody
	GetReturnCode() *string
	SetRule(v string) *GetCustomResponseCodeRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetCustomResponseCodeRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetCustomResponseCodeRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetCustomResponseCodeRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetCustomResponseCodeRuleResponseBody
	GetSiteVersion() *int32
}

type GetCustomResponseCodeRuleResponseBody struct {
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// example:
	//
	// 0
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetCustomResponseCodeRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomResponseCodeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomResponseCodeRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCustomResponseCodeRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetCustomResponseCodeRuleResponseBody) GetPageId() *string {
	return s.PageId
}

func (s *GetCustomResponseCodeRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomResponseCodeRuleResponseBody) GetReturnCode() *string {
	return s.ReturnCode
}

func (s *GetCustomResponseCodeRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetCustomResponseCodeRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetCustomResponseCodeRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetCustomResponseCodeRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetCustomResponseCodeRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetCustomResponseCodeRuleResponseBody) SetConfigId(v int64) *GetCustomResponseCodeRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetConfigType(v string) *GetCustomResponseCodeRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetPageId(v string) *GetCustomResponseCodeRuleResponseBody {
	s.PageId = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetRequestId(v string) *GetCustomResponseCodeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetReturnCode(v string) *GetCustomResponseCodeRuleResponseBody {
	s.ReturnCode = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetRule(v string) *GetCustomResponseCodeRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetRuleEnable(v string) *GetCustomResponseCodeRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetRuleName(v string) *GetCustomResponseCodeRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetSequence(v int32) *GetCustomResponseCodeRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) SetSiteVersion(v int32) *GetCustomResponseCodeRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
