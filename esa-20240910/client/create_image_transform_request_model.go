// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTransformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAvif(v string) *CreateImageTransformRequest
	GetAutoAvif() *string
	SetAutoWebp(v string) *CreateImageTransformRequest
	GetAutoWebp() *string
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
	// example:
	//
	// on
	AutoAvif *string `json:"AutoAvif,omitempty" xml:"AutoAvif,omitempty"`
	// example:
	//
	// on
	AutoWebp   *string `json:"AutoWebp,omitempty" xml:"AutoWebp,omitempty"`
	Enable     *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Rule       *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence   *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateImageTransformRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTransformRequest) GoString() string {
	return s.String()
}

func (s *CreateImageTransformRequest) GetAutoAvif() *string {
	return s.AutoAvif
}

func (s *CreateImageTransformRequest) GetAutoWebp() *string {
	return s.AutoWebp
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

func (s *CreateImageTransformRequest) SetAutoAvif(v string) *CreateImageTransformRequest {
	s.AutoAvif = &v
	return s
}

func (s *CreateImageTransformRequest) SetAutoWebp(v string) *CreateImageTransformRequest {
	s.AutoWebp = &v
	return s
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
