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
	// Specifies whether to enable the Alt-Svc header. Disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	AltSvc *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	// Specifies whether to include the `clear` parameter in the Alt-Svc header. Disabled by default. Valid values:
	//
	// - `on`: The parameter is included.
	//
	// - `off`: The parameter is not included.
	//
	// example:
	//
	// on
	AltSvcClear *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	// The Max Age for the Alt-Svc header, in seconds. The default is 86400.
	//
	// example:
	//
	// 86400
	AltSvcMa *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	// Specifies whether to include the `persist` parameter in the Alt-Svc header. Disabled by default. Valid values:
	//
	// - `on`: The parameter is included.
	//
	// - `off`: The parameter is not included.
	//
	// example:
	//
	// on
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// Specifies whether to enable HTTP Strict Transport Security (HSTS). Disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	Hsts *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	// Specifies whether to include the `includeSubDomains` directive in the HSTS header. Disabled by default. Valid values:
	//
	// - `on`: The directive is included.
	//
	// - `off`: The directive is not included.
	//
	// example:
	//
	// on
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	// The value of the `max-age` directive for the HSTS header, in seconds.
	//
	// example:
	//
	// 3600
	HstsMaxAge *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Specifies whether to enable HSTS Preload by including the `preload` directive in the HSTS header. Disabled by default. Valid values:
	//
	// - `on`: The directive is included.
	//
	// - `off`: The directive is not included.
	//
	// example:
	//
	// on
	HstsPreload *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Specifies whether to enable Force HTTPS. Disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	HttpsForce *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	// The Redirection Status Code to use when Force HTTPS is enabled. Valid values:
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
	// Specifies whether to reject TLS Handshake Requests that do not include an SNI. Disabled by default. Valid values:
	//
	// - `on`: Rejects requests without an SNI.
	//
	// - `off`: Allows requests without an SNI.
	//
	// example:
	//
	// on
	HttpsNoSniDeny *string `json:"HttpsNoSniDeny,omitempty" xml:"HttpsNoSniDeny,omitempty"`
	// Specifies whether to enable Server Name Indication (SNI) verification. Disabled by default. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	HttpsSniVerify *string `json:"HttpsSniVerify,omitempty" xml:"HttpsSniVerify,omitempty"`
	// Specifies the allowlist of SNI values. Separate multiple values with a space.
	//
	// example:
	//
	// abc edf
	HttpsSniWhitelist *string `json:"HttpsSniWhitelist,omitempty" xml:"HttpsSniWhitelist,omitempty"`
	// The content of the Rule, which is a Conditional Expression that matches user Requests. This parameter is optional when adding a Global Configuration. Supported use cases include:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, use a custom expression. For example: `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. This parameter is optional when adding a Global Configuration. Valid values:
	//
	// - `on`: The rule is enabled.
	//
	// - `off`: The rule is disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the Rule. This parameter is optional when adding a Global Configuration.
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
	// The ID of the Site. You can get this ID by calling the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The Site\\"s configuration Version. For Sites with version management enabled, this parameter specifies the Version to which the configuration applies. The default is 0.
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
