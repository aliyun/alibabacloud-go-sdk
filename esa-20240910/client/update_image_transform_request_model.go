// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageTransformRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// Configuration ID. It can be obtained by calling the [ListImageTransforms](https://help.aliyun.com/document_detail/2869056.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Indicates whether to enable image transformation. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Rule content, used to match user requests with conditional expressions. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - To match all incoming requests: Set the value to true.
	//
	// - To match specific requests: Set the value to a custom expression, for example: (http.host eq "video.example.com")
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Possible values:
	//
	// - on: Enable.
	//
	// - off: Disable.
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
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
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
