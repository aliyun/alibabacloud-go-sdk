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
	SetExtendFieldStoreMode(v string) *UpdateNormalizationRuleShrinkRequest
	GetExtendFieldStoreMode() *string
	SetLang(v string) *UpdateNormalizationRuleShrinkRequest
	GetLang() *string
	SetNormalizationCategoryId(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationCategoryId() *string
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
	SetNormalizationSecurityDomainId(v string) *UpdateNormalizationRuleShrinkRequest
	GetNormalizationSecurityDomainId() *string
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
	// Specifies whether to package non-standard fields into the extend_content extension field. Valid values:
	//
	// - enabled: The feature is enabled.
	//
	// - disabled: The feature is disabled.
	//
	// example:
	//
	// enabled
	ExtendContentPacked *string `json:"ExtendContentPacked,omitempty" xml:"ExtendContentPacked,omitempty"`
	// The storage mode for extension fields. Valid values: flat, pack, and reject.
	//
	// example:
	//
	// flat
	ExtendFieldStoreMode *string `json:"ExtendFieldStoreMode,omitempty" xml:"ExtendFieldStoreMode,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The normalization category.
	//
	// example:
	//
	// HOST_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// The description of the normalization rule.
	//
	// example:
	//
	// normalization_rule_Z57np
	NormalizationRuleDescription *string `json:"NormalizationRuleDescription,omitempty" xml:"NormalizationRuleDescription,omitempty"`
	// The expression for the normalization rule.
	//
	// example:
	//
	// 	- | pack-fields -include=\\"[\\s\\S]+\\" as extend_content
	NormalizationRuleExpression *string `json:"NormalizationRuleExpression,omitempty" xml:"NormalizationRuleExpression,omitempty"`
	// The format of the normalization rule.
	//
	// example:
	//
	// SPL
	NormalizationRuleFormat *string `json:"NormalizationRuleFormat,omitempty" xml:"NormalizationRuleFormat,omitempty"`
	// The ID of the normalization rule.
	//
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// The list of normalization rule IDs.
	NormalizationRuleIdsShrink *string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty"`
	// The mode of the normalization rule. Valid values:
	//
	// - both
	//
	// - scan
	//
	// - realtime
	//
	// example:
	//
	// both
	NormalizationRuleMode *string `json:"NormalizationRuleMode,omitempty" xml:"NormalizationRuleMode,omitempty"`
	// The name of the normalization rule.
	//
	// example:
	//
	// normalization_rule_Z57np
	NormalizationRuleName *string `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	// The type of the normalization rule. Valid values:
	//
	// - predefined: predefined normalization rule.
	//
	// - custom: custom normalization rule.
	//
	// example:
	//
	// custom
	NormalizationRuleType *string `json:"NormalizationRuleType,omitempty" xml:"NormalizationRuleType,omitempty"`
	// The ID of the normalization structure.
	//
	// example:
	//
	// HTTP_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// NETWORK_AND_WEB_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The field to use for sorting the rule list. Valid values:
	//
	// - GmtModified: Sorts by modification time.
	//
	// - Id: Sorts by rule ID (default).
	//
	// example:
	//
	// Id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The product ID.
	//
	// example:
	//
	// alibaba_cloud_sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region where the Data Management center of threat analysis is located. Select a region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: an asset in the Chinese mainland.
	//
	// - ap-southeast-1: an asset outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. This parameter is used when an administrator switches to the perspective of the member.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The vendor ID that corresponds to the normalization rule.
	//
	// example:
	//
	// alibaba_cloud
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

func (s *UpdateNormalizationRuleShrinkRequest) GetExtendFieldStoreMode() *string {
	return s.ExtendFieldStoreMode
}

func (s *UpdateNormalizationRuleShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
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

func (s *UpdateNormalizationRuleShrinkRequest) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
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

func (s *UpdateNormalizationRuleShrinkRequest) SetExtendFieldStoreMode(v string) *UpdateNormalizationRuleShrinkRequest {
	s.ExtendFieldStoreMode = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetLang(v string) *UpdateNormalizationRuleShrinkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationCategoryId(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationCategoryId = &v
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

func (s *UpdateNormalizationRuleShrinkRequest) SetNormalizationSecurityDomainId(v string) *UpdateNormalizationRuleShrinkRequest {
	s.NormalizationSecurityDomainId = &v
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
