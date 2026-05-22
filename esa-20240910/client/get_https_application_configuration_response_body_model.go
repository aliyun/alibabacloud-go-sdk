// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpsApplicationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAltSvc(v string) *GetHttpsApplicationConfigurationResponseBody
	GetAltSvc() *string
	SetAltSvcClear(v string) *GetHttpsApplicationConfigurationResponseBody
	GetAltSvcClear() *string
	SetAltSvcMa(v string) *GetHttpsApplicationConfigurationResponseBody
	GetAltSvcMa() *string
	SetAltSvcPersist(v string) *GetHttpsApplicationConfigurationResponseBody
	GetAltSvcPersist() *string
	SetConfigId(v int64) *GetHttpsApplicationConfigurationResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetHttpsApplicationConfigurationResponseBody
	GetConfigType() *string
	SetHsts(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHsts() *string
	SetHstsIncludeSubdomains(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHstsIncludeSubdomains() *string
	SetHstsMaxAge(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHstsMaxAge() *string
	SetHstsPreload(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHstsPreload() *string
	SetHttpsForce(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHttpsForce() *string
	SetHttpsForceCode(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHttpsForceCode() *string
	SetHttpsNoSniDeny(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHttpsNoSniDeny() *string
	SetHttpsSniVerify(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHttpsSniVerify() *string
	SetHttpsSniWhitelist(v string) *GetHttpsApplicationConfigurationResponseBody
	GetHttpsSniWhitelist() *string
	SetRequestId(v string) *GetHttpsApplicationConfigurationResponseBody
	GetRequestId() *string
	SetRule(v string) *GetHttpsApplicationConfigurationResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetHttpsApplicationConfigurationResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetHttpsApplicationConfigurationResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetHttpsApplicationConfigurationResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetHttpsApplicationConfigurationResponseBody
	GetSiteVersion() *int32
}

type GetHttpsApplicationConfigurationResponseBody struct {
	AltSvc                *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	AltSvcClear           *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	AltSvcMa              *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	AltSvcPersist         *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	ConfigId              *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType            *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Hsts                  *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	HstsMaxAge            *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	HstsPreload           *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	HttpsForce            *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	HttpsForceCode        *string `json:"HttpsForceCode,omitempty" xml:"HttpsForceCode,omitempty"`
	HttpsNoSniDeny        *string `json:"HttpsNoSniDeny,omitempty" xml:"HttpsNoSniDeny,omitempty"`
	HttpsSniVerify        *string `json:"HttpsSniVerify,omitempty" xml:"HttpsSniVerify,omitempty"`
	HttpsSniWhitelist     *string `json:"HttpsSniWhitelist,omitempty" xml:"HttpsSniWhitelist,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rule                  *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable            *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName              *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence              *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion           *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetHttpsApplicationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpsApplicationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetAltSvc() *string {
	return s.AltSvc
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetAltSvcClear() *string {
	return s.AltSvcClear
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetAltSvcMa() *string {
	return s.AltSvcMa
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetAltSvcPersist() *string {
	return s.AltSvcPersist
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHsts() *string {
	return s.Hsts
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHstsIncludeSubdomains() *string {
	return s.HstsIncludeSubdomains
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHstsMaxAge() *string {
	return s.HstsMaxAge
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHstsPreload() *string {
	return s.HstsPreload
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHttpsForce() *string {
	return s.HttpsForce
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHttpsForceCode() *string {
	return s.HttpsForceCode
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHttpsNoSniDeny() *string {
	return s.HttpsNoSniDeny
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHttpsSniVerify() *string {
	return s.HttpsSniVerify
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetHttpsSniWhitelist() *string {
	return s.HttpsSniWhitelist
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetHttpsApplicationConfigurationResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetAltSvc(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.AltSvc = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetAltSvcClear(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.AltSvcClear = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetAltSvcMa(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.AltSvcMa = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetAltSvcPersist(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.AltSvcPersist = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetConfigId(v int64) *GetHttpsApplicationConfigurationResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetConfigType(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHsts(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.Hsts = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHstsIncludeSubdomains(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.HstsIncludeSubdomains = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHstsMaxAge(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.HstsMaxAge = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHstsPreload(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.HstsPreload = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHttpsForce(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.HttpsForce = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHttpsForceCode(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.HttpsForceCode = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHttpsNoSniDeny(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.HttpsNoSniDeny = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHttpsSniVerify(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.HttpsSniVerify = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetHttpsSniWhitelist(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.HttpsSniWhitelist = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetRequestId(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetRule(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.Rule = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetRuleEnable(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetRuleName(v string) *GetHttpsApplicationConfigurationResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetSequence(v int32) *GetHttpsApplicationConfigurationResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) SetSiteVersion(v int32) *GetHttpsApplicationConfigurationResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
