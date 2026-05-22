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
	// example:
	//
	// on
	AutoAvif *string `json:"AutoAvif,omitempty" xml:"AutoAvif,omitempty"`
	// example:
	//
	// on
	AutoWebp *string `json:"AutoWebp,omitempty" xml:"AutoWebp,omitempty"`
	// This parameter is required.
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enable     *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Rule       *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence   *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
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
