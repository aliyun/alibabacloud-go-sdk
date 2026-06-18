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
	// Specifies whether to enable the `Alt-Svc` header. This feature is disabled by default. Valid values:
	//
	// - `on`: Enables the Alt-Svc header.
	//
	// - `off`: Disables the Alt-Svc header.
	//
	// example:
	//
	// on
	AltSvc *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	// Specifies whether to include the `clear` directive in the `Alt-Svc` header. The directive is not included by default. Valid values:
	//
	// - `on`: Includes the directive.
	//
	// - `off`: Does not include the directive.
	//
	// example:
	//
	// on
	AltSvcClear *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	// The max-age for the Alt-Svc header, in seconds. The default is 86400.
	//
	// example:
	//
	// 86400
	AltSvcMa *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	// Specifies whether to include the `persist` directive in the `Alt-Svc` header. The directive is not included by default. Valid values:
	//
	// - `on`: Includes the directive.
	//
	// - `off`: Does not include the directive.
	//
	// example:
	//
	// on
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable HTTP Strict Transport Security (HSTS). This feature is disabled by default. Valid values:
	//
	// - `on`: Enables HSTS.
	//
	// - `off`: Disables HSTS.
	//
	// example:
	//
	// on
	Hsts *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	// Specifies whether to include subdomains in the HSTS policy. Subdomains are not included by default. Valid values:
	//
	// - `on`: Includes subdomains.
	//
	// - `off`: Does not include subdomains.
	//
	// example:
	//
	// on
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	// The HSTS max-age in seconds.
	//
	// example:
	//
	// 3600
	HstsMaxAge *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Specifies whether to enable HSTS Preload. HSTS Preload is disabled by default. Valid values:
	//
	// - `on`: Enables Preload.
	//
	// - `off`: Disables Preload.
	//
	// example:
	//
	// on
	HstsPreload *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Specifies whether to enable Force HTTPS. This feature is disabled by default. Valid values:
	//
	// - `on`: Enables Force HTTPS.
	//
	// - `off`: Disables Force HTTPS.
	//
	// example:
	//
	// on
	HttpsForce *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	// The redirect status code for Force HTTPS. Valid values:
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
	// Specifies whether to deny TLS handshake requests that do not include an SNI. This feature is disabled by default. Valid values:
	//
	// - `on`: Denies requests without an SNI.
	//
	// - `off`: Allows requests without an SNI.
	//
	// example:
	//
	// on
	HttpsNoSniDeny *string `json:"HttpsNoSniDeny,omitempty" xml:"HttpsNoSniDeny,omitempty"`
	// Specifies whether to enable Server Name Indication (SNI) verification. This feature is disabled by default. Valid values:
	//
	// - `on`: Enables SNI verification.
	//
	// - `off`: Disables SNI verification.
	//
	// example:
	//
	// on
	HttpsSniVerify *string `json:"HttpsSniVerify,omitempty" xml:"HttpsSniVerify,omitempty"`
	// The SNI allowlist. Separate multiple hostnames with spaces.
	//
	// example:
	//
	// abc edf
	HttpsSniWhitelist *string `json:"HttpsSniWhitelist,omitempty" xml:"HttpsSniWhitelist,omitempty"`
	// The conditional expression that the rule uses to match requests. This parameter is not required for a global configuration.
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, provide a custom expression, such as `(http.host eq "video.example.com")`
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. This parameter is not required for a global configuration. Valid values:
	//
	// - `on`: Enables the rule.
	//
	// - `off`: Disables the rule.
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
	// The execution order of the rule. A lower value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. To obtain this ID, call the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
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
