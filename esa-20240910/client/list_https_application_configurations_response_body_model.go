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
	Configs    []*ListHttpsApplicationConfigurationsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	PageNumber *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	Rule                  *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable            *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName              *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence              *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion           *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
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
