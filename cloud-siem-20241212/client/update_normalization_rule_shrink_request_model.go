// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNormalizationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendContentPacked(v string) *UpdateNormalizationRuleShrinkRequest
	GetExtendContentPacked() *string
	SetLang(v string) *UpdateNormalizationRuleShrinkRequest
	GetLang() *string
	SetNormalizationRuleDescription(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationRuleDescription() *string
	SetNormalizationRuleExpression(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationRuleExpression() *string
	SetNormalizationRuleFormat(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationRuleFormat() *string
	SetNormalizationRuleId(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationRuleId() *string
	SetNormalizationRuleIdsShrink(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationRuleIdsShrink() *string
	SetNormalizationRuleMode(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationRuleMode() *string
	SetNormalizationRuleName(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationRuleName() *string
	SetNormalizationRuleType(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationRuleType() *string
	SetNormalizationSchemaId(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationSchemaId() *string
	SetOrderField(v string) *UpdateNormalizationRuleShrinkRequest
	GetOrderField() *string
	SetProductId(v string) *UpdateNormalizationRuleShrinkRequest
	GetProductId() *string
	SetRegionId(v string) *UpdateNormalizationRuleShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateNormalizationRuleShrinkRequest
	GetRoleFor() *int64
	SetVendorId(v string) *UpdateNormalizationRuleShrinkRequest
	GetVendorId() *string
}

type UpdateNormalizationRuleShrinkRequest struct {
	ExtendContentPacked *string `json:"ExtendContentPacked,omitempty" xml:"ExtendContentPacked,omitempty"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	NormalizationRuleFormat *string `json:"NormalizationRuleFormat,omitempty" xml:"NormalizationRuleFormat,omitempty"`
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp。
	NormalizationRuleId        *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	NormalizationRuleIdsShrink *string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty"`
	// example:
	//
	// both。
	NormalizationRuleMode *string `json:"NormalizationRuleMode,omitempty" xml:"NormalizationRuleMode,omitempty"`
	// example:
	//
	// normalization_rule_Z57np。
	NormalizationRuleName *string `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	NormalizationRuleType *string `json:"NormalizationRuleType,omitempty" xml:"NormalizationRuleType,omitempty"`
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

func (s UpdateNormalizationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationRuleShrinkRequest) GetExtendContentPacked() *string {
	return s.ExtendContentPacked
}

func (s *UpdateNormalizationRuleShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationRuleDescription() *string {
	return s.NormalizationRuleDescription
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationRuleFormat() *string {
	return s.NormalizationRuleFormat
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationRuleIdsShrink() *string {
	return s.NormalizationRuleIdsShrink
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationRuleMode() *string {
	return s.NormalizationRuleMode
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *UpdateNormalizationRuleShrinkRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *UpdateNormalizationRuleShrinkRequest) GetProductId() *string {
	return s.ProductId
}

func (s *UpdateNormalizationRuleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNormalizationRuleShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateNormalizationRuleShrinkRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *UpdateNormalizationRuleShrinkRequest) SetExtendContentPacked(v string) *UpdateNormalizationRuleShrinkRequest {
	s.ExtendContentPacked = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetLang(v string) *UpdateNormalizationRuleShrinkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationRuleDescription(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationRuleDescription = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationRuleExpression(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationRuleFormat(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationRuleFormat = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationRuleId(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationRuleIdsShrink(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationRuleIdsShrink = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationRuleMode(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationRuleMode = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationRuleName(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationRuleName = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationRuleType(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationRuleType = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationSchemaId(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetOrderField(v string) *UpdateNormalizationRuleShrinkRequest {
	s.OrderField = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetProductId(v string) *UpdateNormalizationRuleShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetRegionId(v string) *UpdateNormalizationRuleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetRoleFor(v int64) *UpdateNormalizationRuleShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetVendorId(v string) *UpdateNormalizationRuleShrinkRequest {
	s.VendorId = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
