// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetNormalizationRuleRequest
	GetLang() *string
	SetNormalizationRuleId(v string) *GetNormalizationRuleRequest
	GetNormalizationRuleId() *string
	SetRegionId(v string) *GetNormalizationRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetNormalizationRuleRequest
	GetRoleFor() *int64
}

type GetNormalizationRuleRequest struct {
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

func (s GetNormalizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationRuleRequest) GoString() string {
	return s.String()
}

func (s *GetNormalizationRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *GetNormalizationRuleRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *GetNormalizationRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNormalizationRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetNormalizationRuleRequest) SetLang(v string) *GetNormalizationRuleRequest {
	s.Lang = &v
	return s
}

func (s *GetNormalizationRuleRequest) SetNormalizationRuleId(v string) *GetNormalizationRuleRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *GetNormalizationRuleRequest) SetRegionId(v string) *GetNormalizationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *GetNormalizationRuleRequest) SetRoleFor(v int64) *GetNormalizationRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *GetNormalizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
