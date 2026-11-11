// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTransformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAvif(v string) *GetImageTransformResponseBody
	GetAutoAvif() *string
	SetAutoWebp(v string) *GetImageTransformResponseBody
	GetAutoWebp() *string
	SetConfigId(v int64) *GetImageTransformResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetImageTransformResponseBody
	GetConfigType() *string
	SetEnable(v string) *GetImageTransformResponseBody
	GetEnable() *string
	SetRequestId(v string) *GetImageTransformResponseBody
	GetRequestId() *string
	SetRule(v string) *GetImageTransformResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetImageTransformResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetImageTransformResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetImageTransformResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetImageTransformResponseBody
	GetSiteVersion() *int32
}

type GetImageTransformResponseBody struct {
	// Specifies whether adaptive AVIF is enabled. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// on
	AutoAvif *string `json:"AutoAvif,omitempty" xml:"AutoAvif,omitempty"`
	// Specifies whether adaptive WebP is enabled. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// on
	AutoWebp *string `json:"AutoWebp,omitempty" xml:"AutoWebp,omitempty"`
	// The ID of the configuration.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The type of the configuration. Valid values:
	//
	// - global: global configuration
	//
	// - rule: rule configuration
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Specifies whether image transform is enabled. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution sequence of the rule. A smaller value indicates a higher execution priority.
	//
	// example:
	//
	// 2
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetImageTransformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageTransformResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTransformResponseBody) GetAutoAvif() *string {
	return s.AutoAvif
}

func (s *GetImageTransformResponseBody) GetAutoWebp() *string {
	return s.AutoWebp
}

func (s *GetImageTransformResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetImageTransformResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetImageTransformResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetImageTransformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageTransformResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetImageTransformResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetImageTransformResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetImageTransformResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetImageTransformResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetImageTransformResponseBody) SetAutoAvif(v string) *GetImageTransformResponseBody {
	s.AutoAvif = &v
	return s
}

func (s *GetImageTransformResponseBody) SetAutoWebp(v string) *GetImageTransformResponseBody {
	s.AutoWebp = &v
	return s
}

func (s *GetImageTransformResponseBody) SetConfigId(v int64) *GetImageTransformResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetImageTransformResponseBody) SetConfigType(v string) *GetImageTransformResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetImageTransformResponseBody) SetEnable(v string) *GetImageTransformResponseBody {
	s.Enable = &v
	return s
}

func (s *GetImageTransformResponseBody) SetRequestId(v string) *GetImageTransformResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageTransformResponseBody) SetRule(v string) *GetImageTransformResponseBody {
	s.Rule = &v
	return s
}

func (s *GetImageTransformResponseBody) SetRuleEnable(v string) *GetImageTransformResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetImageTransformResponseBody) SetRuleName(v string) *GetImageTransformResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetImageTransformResponseBody) SetSequence(v int32) *GetImageTransformResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetImageTransformResponseBody) SetSiteVersion(v int32) *GetImageTransformResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetImageTransformResponseBody) Validate() error {
	return dara.Validate(s)
}
