// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpsApplicationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAltSvc(v string) *UpdateHttpsApplicationConfigurationRequest
	GetAltSvc() *string
	SetAltSvcClear(v string) *UpdateHttpsApplicationConfigurationRequest
	GetAltSvcClear() *string
	SetAltSvcMa(v string) *UpdateHttpsApplicationConfigurationRequest
	GetAltSvcMa() *string
	SetAltSvcPersist(v string) *UpdateHttpsApplicationConfigurationRequest
	GetAltSvcPersist() *string
	SetConfigId(v int64) *UpdateHttpsApplicationConfigurationRequest
	GetConfigId() *int64
	SetHsts(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHsts() *string
	SetHstsIncludeSubdomains(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHstsIncludeSubdomains() *string
	SetHstsMaxAge(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHstsMaxAge() *string
	SetHstsPreload(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHstsPreload() *string
	SetHttpsForce(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHttpsForce() *string
	SetHttpsForceCode(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHttpsForceCode() *string
	SetHttpsNoSniDeny(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHttpsNoSniDeny() *string
	SetHttpsSniVerify(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHttpsSniVerify() *string
	SetHttpsSniWhitelist(v string) *UpdateHttpsApplicationConfigurationRequest
	GetHttpsSniWhitelist() *string
	SetRule(v string) *UpdateHttpsApplicationConfigurationRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpsApplicationConfigurationRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpsApplicationConfigurationRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpsApplicationConfigurationRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpsApplicationConfigurationRequest
	GetSiteId() *int64
}

type UpdateHttpsApplicationConfigurationRequest struct {
	AltSvc        *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	AltSvcClear   *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	AltSvcMa      *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// This parameter is required.
	ConfigId              *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Hsts                  *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	HstsMaxAge            *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	HstsPreload           *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	HttpsForce            *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	HttpsForceCode        *string `json:"HttpsForceCode,omitempty" xml:"HttpsForceCode,omitempty"`
	HttpsNoSniDeny        *string `json:"HttpsNoSniDeny,omitempty" xml:"HttpsNoSniDeny,omitempty"`
	HttpsSniVerify        *string `json:"HttpsSniVerify,omitempty" xml:"HttpsSniVerify,omitempty"`
	HttpsSniWhitelist     *string `json:"HttpsSniWhitelist,omitempty" xml:"HttpsSniWhitelist,omitempty"`
	Rule                  *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable            *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName              *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence              *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpsApplicationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpsApplicationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetAltSvc() *string {
	return s.AltSvc
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetAltSvcClear() *string {
	return s.AltSvcClear
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetAltSvcMa() *string {
	return s.AltSvcMa
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetAltSvcPersist() *string {
	return s.AltSvcPersist
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHsts() *string {
	return s.Hsts
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHstsIncludeSubdomains() *string {
	return s.HstsIncludeSubdomains
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHstsMaxAge() *string {
	return s.HstsMaxAge
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHstsPreload() *string {
	return s.HstsPreload
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHttpsForce() *string {
	return s.HttpsForce
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHttpsForceCode() *string {
	return s.HttpsForceCode
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHttpsNoSniDeny() *string {
	return s.HttpsNoSniDeny
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHttpsSniVerify() *string {
	return s.HttpsSniVerify
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetHttpsSniWhitelist() *string {
	return s.HttpsSniWhitelist
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpsApplicationConfigurationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetAltSvc(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.AltSvc = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetAltSvcClear(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.AltSvcClear = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetAltSvcMa(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.AltSvcMa = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetAltSvcPersist(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.AltSvcPersist = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetConfigId(v int64) *UpdateHttpsApplicationConfigurationRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHsts(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.Hsts = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHstsIncludeSubdomains(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.HstsIncludeSubdomains = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHstsMaxAge(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.HstsMaxAge = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHstsPreload(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.HstsPreload = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHttpsForce(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.HttpsForce = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHttpsForceCode(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.HttpsForceCode = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHttpsNoSniDeny(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.HttpsNoSniDeny = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHttpsSniVerify(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.HttpsSniVerify = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetHttpsSniWhitelist(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.HttpsSniWhitelist = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetRule(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetRuleEnable(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetRuleName(v string) *UpdateHttpsApplicationConfigurationRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetSequence(v int32) *UpdateHttpsApplicationConfigurationRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) SetSiteId(v int64) *UpdateHttpsApplicationConfigurationRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
