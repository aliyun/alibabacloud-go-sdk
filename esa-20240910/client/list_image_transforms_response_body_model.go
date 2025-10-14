// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTransformsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListImageTransformsResponseBodyConfigs) *ListImageTransformsResponseBody
	GetConfigs() []*ListImageTransformsResponseBodyConfigs
	SetPageNumber(v int32) *ListImageTransformsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImageTransformsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListImageTransformsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListImageTransformsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListImageTransformsResponseBody
	GetTotalPage() *int32
}

type ListImageTransformsResponseBody struct {
	// Configuration list.
	Configs []*ListImageTransformsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size. Range: **1~500**, default is **500**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
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

func (s ListImageTransformsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageTransformsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageTransformsResponseBody) GetConfigs() []*ListImageTransformsResponseBodyConfigs {
	return s.Configs
}

func (s *ListImageTransformsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageTransformsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageTransformsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageTransformsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImageTransformsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListImageTransformsResponseBody) SetConfigs(v []*ListImageTransformsResponseBodyConfigs) *ListImageTransformsResponseBody {
	s.Configs = v
	return s
}

func (s *ListImageTransformsResponseBody) SetPageNumber(v int32) *ListImageTransformsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListImageTransformsResponseBody) SetPageSize(v int32) *ListImageTransformsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImageTransformsResponseBody) SetRequestId(v string) *ListImageTransformsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageTransformsResponseBody) SetTotalCount(v int32) *ListImageTransformsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImageTransformsResponseBody) SetTotalPage(v int32) *ListImageTransformsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListImageTransformsResponseBody) Validate() error {
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

type ListImageTransformsResponseBodyConfigs struct {
	// Configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type. Possible values:
	//
	// - global: Global configuration;
	//
	// - rule: Rule configuration;
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Switch. Possible values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
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
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site configuration version number. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. The default value is version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListImageTransformsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListImageTransformsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListImageTransformsResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListImageTransformsResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListImageTransformsResponseBodyConfigs) GetEnable() *string {
	return s.Enable
}

func (s *ListImageTransformsResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListImageTransformsResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListImageTransformsResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListImageTransformsResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListImageTransformsResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListImageTransformsResponseBodyConfigs) SetConfigId(v int64) *ListImageTransformsResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListImageTransformsResponseBodyConfigs) SetConfigType(v string) *ListImageTransformsResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListImageTransformsResponseBodyConfigs) SetEnable(v string) *ListImageTransformsResponseBodyConfigs {
	s.Enable = &v
	return s
}

func (s *ListImageTransformsResponseBodyConfigs) SetRule(v string) *ListImageTransformsResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListImageTransformsResponseBodyConfigs) SetRuleEnable(v string) *ListImageTransformsResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListImageTransformsResponseBodyConfigs) SetRuleName(v string) *ListImageTransformsResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListImageTransformsResponseBodyConfigs) SetSequence(v int32) *ListImageTransformsResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListImageTransformsResponseBodyConfigs) SetSiteVersion(v int32) *ListImageTransformsResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListImageTransformsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
