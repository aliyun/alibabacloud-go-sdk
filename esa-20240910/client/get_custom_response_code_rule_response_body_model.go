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
	ConfigId    *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType  *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageId      *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReturnCode  *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	Rule        *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable  *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence    *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
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
