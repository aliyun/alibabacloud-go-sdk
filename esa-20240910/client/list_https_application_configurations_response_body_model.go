// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpsApplicationConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListHttpsApplicationConfigurationsResponseBodyConfigs) *ListHttpsApplicationConfigurationsResponseBody
	GetConfigs() []*ListHttpsApplicationConfigurationsResponseBodyConfigs
	SetPageNumber(v int32) *ListHttpsApplicationConfigurationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpsApplicationConfigurationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHttpsApplicationConfigurationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHttpsApplicationConfigurationsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListHttpsApplicationConfigurationsResponseBody
	GetTotalPage() *int32
}

type ListHttpsApplicationConfigurationsResponseBody struct {
	// Response body configurations.
	Configs []*ListHttpsApplicationConfigurationsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHttpsApplicationConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsApplicationConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpsApplicationConfigurationsResponseBody) GetConfigs() []*ListHttpsApplicationConfigurationsResponseBodyConfigs {
	return s.Configs
}

func (s *ListHttpsApplicationConfigurationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpsApplicationConfigurationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpsApplicationConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpsApplicationConfigurationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHttpsApplicationConfigurationsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHttpsApplicationConfigurationsResponseBody) SetConfigs(v []*ListHttpsApplicationConfigurationsResponseBodyConfigs) *ListHttpsApplicationConfigurationsResponseBody {
	s.Configs = v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBody) SetPageNumber(v int32) *ListHttpsApplicationConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBody) SetPageSize(v int32) *ListHttpsApplicationConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBody) SetRequestId(v string) *ListHttpsApplicationConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBody) SetTotalCount(v int32) *ListHttpsApplicationConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBody) SetTotalPage(v int32) *ListHttpsApplicationConfigurationsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHttpsApplicationConfigurationsResponseBodyConfigs struct {
	// Alt-Svc feature switch, default is off. Value range:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	AltSvc *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	// Whether the Alt-Svc header includes the clear parameter, default is off. Values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	AltSvcClear *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	// The validity period of Alt-Svc in seconds, default is 86400 seconds.
	//
	// example:
	//
	// 86400
	AltSvcMa *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	// Whether the Alt-Svc header includes the persist parameter, default is off. Values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// Configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule-based configurations. Possible values:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule-based configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Whether HSTS is enabled, default is off. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Hsts *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	// Whether to include subdomains in HSTS, default is off. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	// The expiration time of HSTS in seconds.
	//
	// example:
	//
	// 3600
	HstsMaxAge *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Whether HSTS preloading is enabled, default is off. Value range:
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
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
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
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site configuration version number. For sites with version management enabled, this parameter can specify the site version for which the configuration is effective, default is version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpsApplicationConfigurationsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsApplicationConfigurationsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetAltSvc() *string {
	return s.AltSvc
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetAltSvcClear() *string {
	return s.AltSvcClear
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetAltSvcMa() *string {
	return s.AltSvcMa
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetAltSvcPersist() *string {
	return s.AltSvcPersist
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHsts() *string {
	return s.Hsts
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHstsIncludeSubdomains() *string {
	return s.HstsIncludeSubdomains
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHstsMaxAge() *string {
	return s.HstsMaxAge
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHstsPreload() *string {
	return s.HstsPreload
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHttpsForce() *string {
	return s.HttpsForce
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHttpsForceCode() *string {
	return s.HttpsForceCode
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHttpsNoSniDeny() *string {
	return s.HttpsNoSniDeny
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHttpsSniVerify() *string {
	return s.HttpsSniVerify
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetHttpsSniWhitelist() *string {
	return s.HttpsSniWhitelist
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetAltSvc(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.AltSvc = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetAltSvcClear(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.AltSvcClear = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetAltSvcMa(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.AltSvcMa = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetAltSvcPersist(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.AltSvcPersist = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetConfigId(v int64) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetConfigType(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHsts(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.Hsts = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHstsIncludeSubdomains(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.HstsIncludeSubdomains = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHstsMaxAge(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.HstsMaxAge = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHstsPreload(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.HstsPreload = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHttpsForce(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.HttpsForce = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHttpsForceCode(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.HttpsForceCode = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHttpsNoSniDeny(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.HttpsNoSniDeny = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHttpsSniVerify(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.HttpsSniVerify = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetHttpsSniWhitelist(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.HttpsSniWhitelist = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetRule(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetRuleEnable(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetRuleName(v string) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetSequence(v int32) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) SetSiteVersion(v int32) *ListHttpsApplicationConfigurationsResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
