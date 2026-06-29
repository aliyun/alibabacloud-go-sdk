// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageTransformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAvif(v string) *UpdateImageTransformRequest
	GetAutoAvif() *string
	SetAutoWebp(v string) *UpdateImageTransformRequest
	GetAutoWebp() *string
	SetConfigId(v int64) *UpdateImageTransformRequest
	GetConfigId() *int64
	SetEnable(v string) *UpdateImageTransformRequest
	GetEnable() *string
	SetRule(v string) *UpdateImageTransformRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateImageTransformRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateImageTransformRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateImageTransformRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateImageTransformRequest
	GetSiteId() *int64
}

type UpdateImageTransformRequest struct {
	// The adaptive AVIF setting.
	//
	// example:
	//
	// on
	AutoAvif *string `json:"AutoAvif,omitempty" xml:"AutoAvif,omitempty"`
	// The adaptive WebP setting.
	//
	// example:
	//
	// on
	AutoWebp *string `json:"AutoWebp,omitempty" xml:"AutoWebp,omitempty"`
	// The configuration ID. You can call the [ListImageTransforms](https://help.aliyun.com/document_detail/2869056.html) operation to obtain the configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable image transformation. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateImageTransformRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageTransformRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageTransformRequest) GetAutoAvif() *string {
	return s.AutoAvif
}

func (s *UpdateImageTransformRequest) GetAutoWebp() *string {
	return s.AutoWebp
}

func (s *UpdateImageTransformRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateImageTransformRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateImageTransformRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateImageTransformRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateImageTransformRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateImageTransformRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateImageTransformRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateImageTransformRequest) SetAutoAvif(v string) *UpdateImageTransformRequest {
	s.AutoAvif = &v
	return s
}

func (s *UpdateImageTransformRequest) SetAutoWebp(v string) *UpdateImageTransformRequest {
	s.AutoWebp = &v
	return s
}

func (s *UpdateImageTransformRequest) SetConfigId(v int64) *UpdateImageTransformRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateImageTransformRequest) SetEnable(v string) *UpdateImageTransformRequest {
	s.Enable = &v
	return s
}

func (s *UpdateImageTransformRequest) SetRule(v string) *UpdateImageTransformRequest {
	s.Rule = &v
	return s
}

func (s *UpdateImageTransformRequest) SetRuleEnable(v string) *UpdateImageTransformRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateImageTransformRequest) SetRuleName(v string) *UpdateImageTransformRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateImageTransformRequest) SetSequence(v int32) *UpdateImageTransformRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateImageTransformRequest) SetSiteId(v int64) *UpdateImageTransformRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateImageTransformRequest) Validate() error {
	return dara.Validate(s)
}
