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
	// Specifies whether to enable the Alt-Svc feature, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	AltSvc *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	// Specifies whether to include the `clear` parameter in the Alt-Svc header, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	AltSvcClear *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	// The Alt-Svc max-age, in seconds. Default: `86400`.
	//
	// example:
	//
	// 86400
	AltSvcMa *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	// Specifies whether to include the `persist` parameter in the Alt-Svc header, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can query for a global or rule configuration based on this parameter. Valid values:
	//
	// - `global`: a global configuration.
	//
	// - `rule`: a rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Specifies whether to enable HSTS, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	Hsts *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	// Specifies whether to include subdomains in the HSTS policy, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	// The value of the `max-age` directive for HSTS, in seconds.
	//
	// example:
	//
	// 3600
	HstsMaxAge *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Specifies whether to enable HSTS preload, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	HstsPreload *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Specifies whether to enable force HTTPS, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	HttpsForce *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	// The status code for the force HTTPS redirect. Valid values:
	//
	// - `301`
	//
	// - `302`
	//
	// - `307`
	//
	// - `308`
	//
	// example:
	//
	// 301
	HttpsForceCode *string `json:"HttpsForceCode,omitempty" xml:"HttpsForceCode,omitempty"`
	// Specifies whether to deny TLS handshakes that lack an SNI, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	HttpsNoSniDeny *string `json:"HttpsNoSniDeny,omitempty" xml:"HttpsNoSniDeny,omitempty"`
	// Specifies whether to enable SNI verification, which is disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	HttpsSniVerify *string `json:"HttpsSniVerify,omitempty" xml:"HttpsSniVerify,omitempty"`
	// The SNI allowlist. Separate multiple values with a space.
	//
	// example:
	//
	// abc edf
	HttpsSniWhitelist *string `json:"HttpsSniWhitelist,omitempty" xml:"HttpsSniWhitelist,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3790430-3A06-535F-A424-0998BD9A6C9F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The conditional expression used to match user requests. This parameter is not required for a global configuration. There are two scenarios:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. This parameter is not required for a global configuration. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required for a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with configuration versioning enabled, this parameter specifies the applicable site version. The default is version `0`.
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
