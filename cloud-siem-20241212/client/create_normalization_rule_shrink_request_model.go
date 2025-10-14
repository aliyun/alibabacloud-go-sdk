// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNormalizationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendContentPacked(v string) *CreateNormalizationRuleShrinkRequest
	GetExtendContentPacked() *string
	SetLang(v string) *CreateNormalizationRuleShrinkRequest
	GetLang() *string
	SetNormalizationCategoryId(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationCategoryId() *string
	SetNormalizationRuleDescription(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationRuleDescription() *string
	SetNormalizationRuleExpression(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationRuleExpression() *string
	SetNormalizationRuleFormat(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationRuleFormat() *string
	SetNormalizationRuleIdsShrink(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationRuleIdsShrink() *string
	SetNormalizationRuleMode(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationRuleMode() *string
	SetNormalizationRuleName(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationRuleName() *string
	SetNormalizationRuleType(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationRuleType() *string
	SetNormalizationRuleVersion(v int32) *CreateNormalizationRuleShrinkRequest
	GetNormalizationRuleVersion() *int32
	SetNormalizationSchemaId(v string) *CreateNormalizationRuleShrinkRequest
	GetNormalizationSchemaId() *string
	SetOrderField(v string) *CreateNormalizationRuleShrinkRequest
	GetOrderField() *string
	SetProductId(v string) *CreateNormalizationRuleShrinkRequest
	GetProductId() *string
	SetRegionId(v string) *CreateNormalizationRuleShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateNormalizationRuleShrinkRequest
	GetRoleFor() *int64
	SetVendorId(v string) *CreateNormalizationRuleShrinkRequest
	GetVendorId() *string
}

type CreateNormalizationRuleShrinkRequest struct {
	ExtendContentPacked *string `json:"ExtendContentPacked,omitempty" xml:"ExtendContentPacked,omitempty"`
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
	// normalization_rule_Z57np。
	NormalizationRuleDescription *string `json:"NormalizationRuleDescription,omitempty" xml:"NormalizationRuleDescription,omitempty"`
	// example:
	//
	// 	- | pack-fields -include=\\"[\\s\\S]+\\" as extend_content。
	NormalizationRuleExpression *string `json:"NormalizationRuleExpression,omitempty" xml:"NormalizationRuleExpression,omitempty"`
	// example:
	//
	// SPL。
	NormalizationRuleFormat    *string `json:"NormalizationRuleFormat,omitempty" xml:"NormalizationRuleFormat,omitempty"`
	NormalizationRuleIdsShrink *string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty"`
	// example:
	//
	// both。
	NormalizationRuleMode *string `json:"NormalizationRuleMode,omitempty" xml:"NormalizationRuleMode,omitempty"`
	// example:
	//
	// normalization_rule_Z57np。
	NormalizationRuleName    *string `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	NormalizationRuleType    *string `json:"NormalizationRuleType,omitempty" xml:"NormalizationRuleType,omitempty"`
	NormalizationRuleVersion *int32  `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	OrderField            *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// alibaba_cloud。
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s CreateNormalizationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNormalizationRuleShrinkRequest) GetExtendContentPacked() *string {
	return s.ExtendContentPacked
}

func (s *CreateNormalizationRuleShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationRuleDescription() *string {
	return s.NormalizationRuleDescription
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationRuleFormat() *string {
	return s.NormalizationRuleFormat
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationRuleIdsShrink() *string {
	return s.NormalizationRuleIdsShrink
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationRuleMode() *string {
	return s.NormalizationRuleMode
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *CreateNormalizationRuleShrinkRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *CreateNormalizationRuleShrinkRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *CreateNormalizationRuleShrinkRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateNormalizationRuleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNormalizationRuleShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateNormalizationRuleShrinkRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *CreateNormalizationRuleShrinkRequest) SetExtendContentPacked(v string) *CreateNormalizationRuleShrinkRequest {
	s.ExtendContentPacked = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetLang(v string) *CreateNormalizationRuleShrinkRequest {
	s.Lang = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationCategoryId(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationRuleDescription(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationRuleDescription = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationRuleExpression(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationRuleFormat(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationRuleFormat = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationRuleIdsShrink(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationRuleIdsShrink = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationRuleMode(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationRuleMode = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationRuleName(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationRuleName = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationRuleType(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationRuleType = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationRuleVersion(v int32) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetNormalizationSchemaId(v string) *CreateNormalizationRuleShrinkRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetOrderField(v string) *CreateNormalizationRuleShrinkRequest {
	s.OrderField = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetProductId(v string) *CreateNormalizationRuleShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetRegionId(v string) *CreateNormalizationRuleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetRoleFor(v int64) *CreateNormalizationRuleShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) SetVendorId(v string) *CreateNormalizationRuleShrinkRequest {
	s.VendorId = &v
	return s
}

func (s *CreateNormalizationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
