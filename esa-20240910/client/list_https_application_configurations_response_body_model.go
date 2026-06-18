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
	// A list of HTTPS application configurations.
	Configs []*ListHttpsApplicationConfigurationsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
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
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHttpsApplicationConfigurationsResponseBodyConfigs struct {
	// Whether to enable the Alt-Svc feature. Default: `off`. Valid values:
	//
	// - `on`: The Alt-Svc feature is enabled.
	//
	// - `off`: The Alt-Svc feature is disabled.
	//
	// example:
	//
	// on
	AltSvc *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	// Whether the Alt-Svc header includes the `clear` parameter. Default: `off`. Valid values:
	//
	// - `on`: The `clear` parameter is included.
	//
	// - `off`: The `clear` parameter is not included.
	//
	// example:
	//
	// on
	AltSvcClear *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	// The Alt-Svc max-age in seconds. Default: `86400`.
	//
	// example:
	//
	// 86400
	AltSvcMa *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	// Whether the Alt-Svc header includes the `persist` parameter. Default: `off`. Valid values:
	//
	// - `on`: The `persist` parameter is included.
	//
	// - `off`: The `persist` parameter is not included.
	//
	// example:
	//
	// on
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The type of the configuration. Valid values:
	//
	// - `global`: A global configuration.
	//
	// - `rule`: A rule-based configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Whether to enable HSTS. Default: `off`. Valid values:
	//
	// - `on`: HSTS is enabled.
	//
	// - `off`: HSTS is disabled.
	//
	// example:
	//
	// on
	Hsts *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	// Whether the HSTS header includes the `includeSubDomains` directive. Default: `off`. Valid values:
	//
	// - `on`: The `includeSubDomains` directive is included.
	//
	// - `off`: The `includeSubDomains` directive is not included.
	//
	// example:
	//
	// on
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	// The HSTS `max-age`, in seconds.
	//
	// example:
	//
	// 3600
	HstsMaxAge *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Whether the HSTS header includes the `preload` directive. Default: `off`. Valid values:
	//
	// - `on`: The `preload` directive is included.
	//
	// - `off`: The `preload` directive is not included.
	//
	// example:
	//
	// on
	HstsPreload *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Whether to enable HTTPS redirection. Default: `off`. Valid values:
	//
	// - `on`: HTTPS redirection is enabled.
	//
	// - `off`: HTTPS redirection is disabled.
	//
	// example:
	//
	// on
	HttpsForce *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	// The status code for HTTPS redirection. Valid values:
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
	// Whether to reject TLS handshake requests that lack an SNI. Default: `off`. Valid values:
	//
	// - `on`: Requests that lack an SNI are rejected.
	//
	// - `off`: Requests that lack an SNI are not rejected.
	//
	// example:
	//
	// on
	HttpsNoSniDeny *string `json:"HttpsNoSniDeny,omitempty" xml:"HttpsNoSniDeny,omitempty"`
	// Whether to enable SNI verification. Default: `off`. Valid values:
	//
	// - `on`: SNI verification is enabled.
	//
	// - `off`: SNI verification is disabled.
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
	// The content of the rule, a conditional expression that matches user requests. This parameter is not required for a global configuration. The following use cases are supported:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Whether the rule is enabled. This parameter is not required for a global configuration. Valid values:
	//
	// - `on`: The rule is enabled.
	//
	// - `off`: The rule is disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter is not required for a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site configuration version. For sites with version management, this specifies the version to which the configuration applies. Default: `0`.
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
