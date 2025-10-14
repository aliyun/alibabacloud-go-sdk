// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationRuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteNormalizationRuleVersionRequest
	GetLang() *string
	SetNormalizationRuleId(v string) *DeleteNormalizationRuleVersionRequest
	GetNormalizationRuleId() *string
	SetNormalizationRuleVersion(v int32) *DeleteNormalizationRuleVersionRequest
	GetNormalizationRuleVersion() *int32
	SetRegionId(v string) *DeleteNormalizationRuleVersionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteNormalizationRuleVersionRequest
	GetRoleFor() *int64
}

type DeleteNormalizationRuleVersionRequest struct {
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
	// 1。
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

func (s DeleteNormalizationRuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationRuleVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationRuleVersionRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteNormalizationRuleVersionRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *DeleteNormalizationRuleVersionRequest) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *DeleteNormalizationRuleVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNormalizationRuleVersionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteNormalizationRuleVersionRequest) SetLang(v string) *DeleteNormalizationRuleVersionRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNormalizationRuleVersionRequest) SetNormalizationRuleId(v string) *DeleteNormalizationRuleVersionRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *DeleteNormalizationRuleVersionRequest) SetNormalizationRuleVersion(v int32) *DeleteNormalizationRuleVersionRequest {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *DeleteNormalizationRuleVersionRequest) SetRegionId(v string) *DeleteNormalizationRuleVersionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNormalizationRuleVersionRequest) SetRoleFor(v int64) *DeleteNormalizationRuleVersionRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteNormalizationRuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
