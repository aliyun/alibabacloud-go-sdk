// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetRedirectRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetRedirectRuleResponseBody
	GetConfigType() *string
	SetRequestId(v string) *GetRedirectRuleResponseBody
	GetRequestId() *string
	SetReserveQueryString(v string) *GetRedirectRuleResponseBody
	GetReserveQueryString() *string
	SetRule(v string) *GetRedirectRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetRedirectRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetRedirectRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetRedirectRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetRedirectRuleResponseBody
	GetSiteVersion() *int32
	SetStatusCode(v string) *GetRedirectRuleResponseBody
	GetStatusCode() *string
	SetTargetUrl(v string) *GetRedirectRuleResponseBody
	GetTargetUrl() *string
	SetType(v string) *GetRedirectRuleResponseBody
	GetType() *string
}

type GetRedirectRuleResponseBody struct {
	ConfigId           *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType         *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	Rule               *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable         *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence           *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion        *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	StatusCode         *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	TargetUrl          *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRedirectRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetRedirectRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRedirectRuleResponseBody) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *GetRedirectRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetRedirectRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetRedirectRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetRedirectRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetRedirectRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetRedirectRuleResponseBody) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetRedirectRuleResponseBody) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetRedirectRuleResponseBody) GetType() *string {
	return s.Type
}

func (s *GetRedirectRuleResponseBody) SetConfigId(v int64) *GetRedirectRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetConfigType(v string) *GetRedirectRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetRequestId(v string) *GetRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetReserveQueryString(v string) *GetRedirectRuleResponseBody {
	s.ReserveQueryString = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetRule(v string) *GetRedirectRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetRuleEnable(v string) *GetRedirectRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetRuleName(v string) *GetRedirectRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetSequence(v int32) *GetRedirectRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetSiteVersion(v int32) *GetRedirectRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetStatusCode(v string) *GetRedirectRuleResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetTargetUrl(v string) *GetRedirectRuleResponseBody {
	s.TargetUrl = &v
	return s
}

func (s *GetRedirectRuleResponseBody) SetType(v string) *GetRedirectRuleResponseBody {
	s.Type = &v
	return s
}

func (s *GetRedirectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
