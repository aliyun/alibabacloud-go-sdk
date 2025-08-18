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
	// Alt-Svc feature switch. Default is disabled. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	AltSvc *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	// Whether the Alt-Svc header includes the clear parameter. Default is disabled. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	AltSvcClear *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	// Alt-Svc validity period in seconds. The default is 86400 seconds.
	//
	// example:
	//
	// 86400
	AltSvcMa *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	// Whether the Alt-Svc header includes the persist parameter. Default is disabled. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// Configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Possible values:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Whether to enable HSTS. Default is disabled. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Hsts *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	// Whether to include subdomains in HSTS, default is off. Value range:
	//
	// - on: enabled.
	//
	// - off: disabled.
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
	// Whether to enable HSTS preload, default is off. Value range:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	HstsPreload *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Whether to enable forced HTTPS. Default is disabled. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	HttpsForce *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	// Status code for forced HTTPS redirection. Possible values:
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
	// Request ID.
	//
	// example:
	//
	// A3790430-3A06-535F-A424-0998BD9A6C9F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter does not need to be set when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq "video.example.com")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter does not need to be set when adding a global configuration. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter does not need to be set when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the version of the site for which the configuration takes effect. The default is version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
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
