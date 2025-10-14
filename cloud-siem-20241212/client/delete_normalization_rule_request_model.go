// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteNormalizationRuleRequest
	GetLang() *string
	SetNormalizationRuleId(v string) *DeleteNormalizationRuleRequest
	GetNormalizationRuleId() *string
	SetRegionId(v string) *DeleteNormalizationRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteNormalizationRuleRequest
	GetRoleFor() *int64
}

type DeleteNormalizationRuleRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp。
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteNormalizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteNormalizationRuleRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *DeleteNormalizationRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNormalizationRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteNormalizationRuleRequest) SetLang(v string) *DeleteNormalizationRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNormalizationRuleRequest) SetNormalizationRuleId(v string) *DeleteNormalizationRuleRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *DeleteNormalizationRuleRequest) SetRegionId(v string) *DeleteNormalizationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNormalizationRuleRequest) SetRoleFor(v int64) *DeleteNormalizationRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteNormalizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
