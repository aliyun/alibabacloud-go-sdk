// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultNormalizationRuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SetDefaultNormalizationRuleVersionRequest
	GetLang() *string
	SetNormalizationRuleId(v string) *SetDefaultNormalizationRuleVersionRequest
	GetNormalizationRuleId() *string
	SetNormalizationRuleVersion(v int32) *SetDefaultNormalizationRuleVersionRequest
	GetNormalizationRuleVersion() *int32
	SetRegionId(v string) *SetDefaultNormalizationRuleVersionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *SetDefaultNormalizationRuleVersionRequest
	GetRoleFor() *int64
}

type SetDefaultNormalizationRuleVersionRequest struct {
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
	// V1。
	NormalizationRuleVersion *int32 `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s SetDefaultNormalizationRuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultNormalizationRuleVersionRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultNormalizationRuleVersionRequest) GetLang() *string {
	return s.Lang
}

func (s *SetDefaultNormalizationRuleVersionRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *SetDefaultNormalizationRuleVersionRequest) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *SetDefaultNormalizationRuleVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDefaultNormalizationRuleVersionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *SetDefaultNormalizationRuleVersionRequest) SetLang(v string) *SetDefaultNormalizationRuleVersionRequest {
	s.Lang = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionRequest) SetNormalizationRuleId(v string) *SetDefaultNormalizationRuleVersionRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionRequest) SetNormalizationRuleVersion(v int32) *SetDefaultNormalizationRuleVersionRequest {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionRequest) SetRegionId(v string) *SetDefaultNormalizationRuleVersionRequest {
	s.RegionId = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionRequest) SetRoleFor(v int64) *SetDefaultNormalizationRuleVersionRequest {
	s.RoleFor = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
