// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTransformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *CreateImageTransformRequest
	GetEnable() *string
	SetRule(v string) *CreateImageTransformRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateImageTransformRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateImageTransformRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateImageTransformRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateImageTransformRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateImageTransformRequest
	GetSiteVersion() *int32
}

type CreateImageTransformRequest struct {
	// Indicates whether image transformation is enabled. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - To match all incoming requests: Set the value to true
	//
	// - To match specific requests: Set the value to a custom expression, for example: (http.host eq "video.example.com")
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
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the site configuration, defaulting to version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateImageTransformRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTransformRequest) GoString() string {
	return s.String()
}

func (s *CreateImageTransformRequest) GetEnable() *string {
	return s.Enable
}

func (s *CreateImageTransformRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateImageTransformRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateImageTransformRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateImageTransformRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateImageTransformRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateImageTransformRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateImageTransformRequest) SetEnable(v string) *CreateImageTransformRequest {
	s.Enable = &v
	return s
}

func (s *CreateImageTransformRequest) SetRule(v string) *CreateImageTransformRequest {
	s.Rule = &v
	return s
}

func (s *CreateImageTransformRequest) SetRuleEnable(v string) *CreateImageTransformRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateImageTransformRequest) SetRuleName(v string) *CreateImageTransformRequest {
	s.RuleName = &v
	return s
}

func (s *CreateImageTransformRequest) SetSequence(v int32) *CreateImageTransformRequest {
	s.Sequence = &v
	return s
}

func (s *CreateImageTransformRequest) SetSiteId(v int64) *CreateImageTransformRequest {
	s.SiteId = &v
	return s
}

func (s *CreateImageTransformRequest) SetSiteVersion(v int32) *CreateImageTransformRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateImageTransformRequest) Validate() error {
	return dara.Validate(s)
}
