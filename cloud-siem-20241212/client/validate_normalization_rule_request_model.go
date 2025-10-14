// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateNormalizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ValidateNormalizationRuleRequest
	GetData() *string
	SetLang(v string) *ValidateNormalizationRuleRequest
	GetLang() *string
	SetNormalizationCategoryId(v string) *ValidateNormalizationRuleRequest
	GetNormalizationCategoryId() *string
	SetNormalizationSchemaId(v string) *ValidateNormalizationRuleRequest
	GetNormalizationSchemaId() *string
	SetRegionId(v string) *ValidateNormalizationRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ValidateNormalizationRuleRequest
	GetRoleFor() *int64
}

type ValidateNormalizationRuleRequest struct {
	// example:
	//
	// 123456。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// NETWORK_CATEGORY。
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ValidateNormalizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateNormalizationRuleRequest) GoString() string {
	return s.String()
}

func (s *ValidateNormalizationRuleRequest) GetData() *string {
	return s.Data
}

func (s *ValidateNormalizationRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *ValidateNormalizationRuleRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ValidateNormalizationRuleRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *ValidateNormalizationRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ValidateNormalizationRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ValidateNormalizationRuleRequest) SetData(v string) *ValidateNormalizationRuleRequest {
	s.Data = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetLang(v string) *ValidateNormalizationRuleRequest {
	s.Lang = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetNormalizationCategoryId(v string) *ValidateNormalizationRuleRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetNormalizationSchemaId(v string) *ValidateNormalizationRuleRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetRegionId(v string) *ValidateNormalizationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) SetRoleFor(v int64) *ValidateNormalizationRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *ValidateNormalizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
