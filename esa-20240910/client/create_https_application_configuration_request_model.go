// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpsApplicationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAltSvc(v string) *CreateHttpsApplicationConfigurationRequest
	GetAltSvc() *string
	SetAltSvcClear(v string) *CreateHttpsApplicationConfigurationRequest
	GetAltSvcClear() *string
	SetAltSvcMa(v string) *CreateHttpsApplicationConfigurationRequest
	GetAltSvcMa() *string
	SetAltSvcPersist(v string) *CreateHttpsApplicationConfigurationRequest
	GetAltSvcPersist() *string
	SetHsts(v string) *CreateHttpsApplicationConfigurationRequest
	GetHsts() *string
	SetHstsIncludeSubdomains(v string) *CreateHttpsApplicationConfigurationRequest
	GetHstsIncludeSubdomains() *string
	SetHstsMaxAge(v string) *CreateHttpsApplicationConfigurationRequest
	GetHstsMaxAge() *string
	SetHstsPreload(v string) *CreateHttpsApplicationConfigurationRequest
	GetHstsPreload() *string
	SetHttpsForce(v string) *CreateHttpsApplicationConfigurationRequest
	GetHttpsForce() *string
	SetHttpsForceCode(v string) *CreateHttpsApplicationConfigurationRequest
	GetHttpsForceCode() *string
	SetHttpsNoSniDeny(v string) *CreateHttpsApplicationConfigurationRequest
	GetHttpsNoSniDeny() *string
	SetHttpsSniVerify(v string) *CreateHttpsApplicationConfigurationRequest
	GetHttpsSniVerify() *string
	SetHttpsSniWhitelist(v string) *CreateHttpsApplicationConfigurationRequest
	GetHttpsSniWhitelist() *string
	SetRule(v string) *CreateHttpsApplicationConfigurationRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpsApplicationConfigurationRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpsApplicationConfigurationRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpsApplicationConfigurationRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpsApplicationConfigurationRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpsApplicationConfigurationRequest
	GetSiteVersion() *int32
}

type CreateHttpsApplicationConfigurationRequest struct {
	// Alt-Svc feature switch, default is disabled. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	AltSvc *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	// Whether the Alt-Svc header includes the clear parameter, default is disabled. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	AltSvcClear *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	// Alt-Svc validity period in seconds, default is 86400 seconds.
	//
	// example:
	//
	// 86400
	AltSvcMa *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	// Whether the Alt-Svc header includes the persist parameter, default is disabled. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// Whether to enable HSTS, default is disabled. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Hsts *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	// Whether to include subdomains in HSTS, default is disabled. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	// HSTS expiration time in seconds.
	//
	// example:
	//
	// 3600
	HstsMaxAge *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Whether to enable HSTS preload, default is disabled. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	HstsPreload *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Whether to enable forced HTTPS, default is disabled. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	HttpsForce *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	// Forced HTTPS redirect status code. Possible values:
	//
	// - 301
	//
	// - 302
	//
	// - 307
	//
	// - 308
	//
	// example:
	//
	// 301
	HttpsForceCode    *string `json:"HttpsForceCode,omitempty" xml:"HttpsForceCode,omitempty"`
	HttpsNoSniDeny    *string `json:"HttpsNoSniDeny,omitempty" xml:"HttpsNoSniDeny,omitempty"`
	HttpsSniVerify    *string `json:"HttpsSniVerify,omitempty" xml:"HttpsSniVerify,omitempty"`
	HttpsSniWhitelist *string `json:"HttpsSniWhitelist,omitempty" xml:"HttpsSniWhitelist,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, this parameter can specify the version to which the configuration applies, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpsApplicationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpsApplicationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpsApplicationConfigurationRequest) GetAltSvc() *string {
	return s.AltSvc
}

func (s *CreateHttpsApplicationConfigurationRequest) GetAltSvcClear() *string {
	return s.AltSvcClear
}

func (s *CreateHttpsApplicationConfigurationRequest) GetAltSvcMa() *string {
	return s.AltSvcMa
}

func (s *CreateHttpsApplicationConfigurationRequest) GetAltSvcPersist() *string {
	return s.AltSvcPersist
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHsts() *string {
	return s.Hsts
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHstsIncludeSubdomains() *string {
	return s.HstsIncludeSubdomains
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHstsMaxAge() *string {
	return s.HstsMaxAge
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHstsPreload() *string {
	return s.HstsPreload
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHttpsForce() *string {
	return s.HttpsForce
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHttpsForceCode() *string {
	return s.HttpsForceCode
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHttpsNoSniDeny() *string {
	return s.HttpsNoSniDeny
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHttpsSniVerify() *string {
	return s.HttpsSniVerify
}

func (s *CreateHttpsApplicationConfigurationRequest) GetHttpsSniWhitelist() *string {
	return s.HttpsSniWhitelist
}

func (s *CreateHttpsApplicationConfigurationRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpsApplicationConfigurationRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpsApplicationConfigurationRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpsApplicationConfigurationRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpsApplicationConfigurationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpsApplicationConfigurationRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpsApplicationConfigurationRequest) SetAltSvc(v string) *CreateHttpsApplicationConfigurationRequest {
	s.AltSvc = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetAltSvcClear(v string) *CreateHttpsApplicationConfigurationRequest {
	s.AltSvcClear = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetAltSvcMa(v string) *CreateHttpsApplicationConfigurationRequest {
	s.AltSvcMa = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetAltSvcPersist(v string) *CreateHttpsApplicationConfigurationRequest {
	s.AltSvcPersist = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHsts(v string) *CreateHttpsApplicationConfigurationRequest {
	s.Hsts = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHstsIncludeSubdomains(v string) *CreateHttpsApplicationConfigurationRequest {
	s.HstsIncludeSubdomains = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHstsMaxAge(v string) *CreateHttpsApplicationConfigurationRequest {
	s.HstsMaxAge = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHstsPreload(v string) *CreateHttpsApplicationConfigurationRequest {
	s.HstsPreload = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHttpsForce(v string) *CreateHttpsApplicationConfigurationRequest {
	s.HttpsForce = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHttpsForceCode(v string) *CreateHttpsApplicationConfigurationRequest {
	s.HttpsForceCode = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHttpsNoSniDeny(v string) *CreateHttpsApplicationConfigurationRequest {
	s.HttpsNoSniDeny = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHttpsSniVerify(v string) *CreateHttpsApplicationConfigurationRequest {
	s.HttpsSniVerify = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetHttpsSniWhitelist(v string) *CreateHttpsApplicationConfigurationRequest {
	s.HttpsSniWhitelist = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetRule(v string) *CreateHttpsApplicationConfigurationRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetRuleEnable(v string) *CreateHttpsApplicationConfigurationRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetRuleName(v string) *CreateHttpsApplicationConfigurationRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetSequence(v int32) *CreateHttpsApplicationConfigurationRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetSiteId(v int64) *CreateHttpsApplicationConfigurationRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) SetSiteVersion(v int32) *CreateHttpsApplicationConfigurationRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
