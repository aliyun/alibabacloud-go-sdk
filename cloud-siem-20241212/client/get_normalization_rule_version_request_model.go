// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationRuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetNormalizationRuleVersionRequest
	GetLang() *string
	SetNormalizationRuleId(v string) *GetNormalizationRuleVersionRequest
	GetNormalizationRuleId() *string
	SetNormalizationRuleVersion(v int32) *GetNormalizationRuleVersionRequest
	GetNormalizationRuleVersion() *int32
	SetRegionId(v string) *GetNormalizationRuleVersionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetNormalizationRuleVersionRequest
	GetRoleFor() *int64
}

type GetNormalizationRuleVersionRequest struct {
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

func (s GetNormalizationRuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationRuleVersionRequest) GoString() string {
	return s.String()
}

func (s *GetNormalizationRuleVersionRequest) GetLang() *string {
	return s.Lang
}

func (s *GetNormalizationRuleVersionRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *GetNormalizationRuleVersionRequest) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *GetNormalizationRuleVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNormalizationRuleVersionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetNormalizationRuleVersionRequest) SetLang(v string) *GetNormalizationRuleVersionRequest {
	s.Lang = &v
	return s
}

func (s *GetNormalizationRuleVersionRequest) SetNormalizationRuleId(v string) *GetNormalizationRuleVersionRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *GetNormalizationRuleVersionRequest) SetNormalizationRuleVersion(v int32) *GetNormalizationRuleVersionRequest {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *GetNormalizationRuleVersionRequest) SetRegionId(v string) *GetNormalizationRuleVersionRequest {
	s.RegionId = &v
	return s
}

func (s *GetNormalizationRuleVersionRequest) SetRoleFor(v int64) *GetNormalizationRuleVersionRequest {
	s.RoleFor = &v
	return s
}

func (s *GetNormalizationRuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
