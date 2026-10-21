// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNormalizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendContentPacked(v string) *UpdateNormalizationRuleRequest
	GetExtendContentPacked() *string
	SetExtendFieldStoreMode(v string) *UpdateNormalizationRuleRequest
	GetExtendFieldStoreMode() *string
	SetLang(v string) *UpdateNormalizationRuleRequest
	GetLang() *string
	SetNormalizationCategoryId(v string) *UpdateNormalizationRuleRequest
	GetNormalizationCategoryId() *string
	SetNormalizationRuleDescription(v string) *UpdateNormalizationRuleRequest
	GetNormalizationRuleDescription() *string
	SetNormalizationRuleExpression(v string) *UpdateNormalizationRuleRequest
	GetNormalizationRuleExpression() *string
	SetNormalizationRuleFormat(v string) *UpdateNormalizationRuleRequest
	GetNormalizationRuleFormat() *string
	SetNormalizationRuleId(v string) *UpdateNormalizationRuleRequest
	GetNormalizationRuleId() *string
	SetNormalizationRuleIds(v []*string) *UpdateNormalizationRuleRequest
	GetNormalizationRuleIds() []*string
	SetNormalizationRuleMode(v string) *UpdateNormalizationRuleRequest
	GetNormalizationRuleMode() *string
	SetNormalizationRuleName(v string) *UpdateNormalizationRuleRequest
	GetNormalizationRuleName() *string
	SetNormalizationRuleType(v string) *UpdateNormalizationRuleRequest
	GetNormalizationRuleType() *string
	SetNormalizationSchemaId(v string) *UpdateNormalizationRuleRequest
	GetNormalizationSchemaId() *string
	SetNormalizationSecurityDomainId(v string) *UpdateNormalizationRuleRequest
	GetNormalizationSecurityDomainId() *string
	SetOrderField(v string) *UpdateNormalizationRuleRequest
	GetOrderField() *string
	SetProductId(v string) *UpdateNormalizationRuleRequest
	GetProductId() *string
	SetRegionId(v string) *UpdateNormalizationRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateNormalizationRuleRequest
	GetRoleFor() *int64
	SetVendorId(v string) *UpdateNormalizationRuleRequest
	GetVendorId() *string
}

type UpdateNormalizationRuleRequest struct {
	// Specifies whether to pack non-standard fields into the extended field extend_content. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	ExtendContentPacked *string `json:"ExtendContentPacked,omitempty" xml:"ExtendContentPacked,omitempty"`
	// The extended field storage mode. Valid values: flat, pack, reject.
	//
	// example:
	//
	// flat
	ExtendFieldStoreMode *string `json:"ExtendFieldStoreMode,omitempty" xml:"ExtendFieldStoreMode,omitempty"`
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The normalization category to which the rule belongs.
	//
	// example:
	//
	// HOST_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// The normalization rule description.
	//
	// example:
	//
	// normalization_rule_Z57np
	NormalizationRuleDescription *string `json:"NormalizationRuleDescription,omitempty" xml:"NormalizationRuleDescription,omitempty"`
	// The normalization rule expression.
	//
	// example:
	//
	// 	- | pack-fields -include=\\"[\\s\\S]+\\" as extend_content
	NormalizationRuleExpression *string `json:"NormalizationRuleExpression,omitempty" xml:"NormalizationRuleExpression,omitempty"`
	// The normalization rule format.
	//
	// example:
	//
	// SPL
	NormalizationRuleFormat *string `json:"NormalizationRuleFormat,omitempty" xml:"NormalizationRuleFormat,omitempty"`
	// The normalization rule ID.
	//
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// The list of normalization rule IDs.
	NormalizationRuleIds []*string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty" type:"Repeated"`
	// The normalization rule mode. Valid values:
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
	// The normalization rule name.
	//
	// example:
	//
	// normalization_rule_Z57np
	NormalizationRuleName *string `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	// The normalization rule type. Valid values:
	//
	// - predefined: predefined normalization rule.
	//
	// - custom: custom normalization rule.
	//
	// example:
	//
	// custom
	NormalizationRuleType *string `json:"NormalizationRuleType,omitempty" xml:"NormalizationRuleType,omitempty"`
	// The normalization schema ID.
	//
	// example:
	//
	// HTTP_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// The normalization security domain ID.
	//
	// example:
	//
	// NETWORK_AND_WEB_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The field used to sort the rule list. Valid values:
	//
	// - GmtModified: Sort by modification time.
	//
	// - Id: Sort by rule ID (default).
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
	// The region where the data management center of the threat detection service is located. Select the management center region based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member account that the administrator switches to.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The vendor ID associated with the normalization rule.
	//
	// example:
	//
	// alibaba_cloud
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s UpdateNormalizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationRuleRequest) GetExtendContentPacked() *string {
	return s.ExtendContentPacked
}

func (s *UpdateNormalizationRuleRequest) GetExtendFieldStoreMode() *string {
	return s.ExtendFieldStoreMode
}

func (s *UpdateNormalizationRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationRuleDescription() *string {
	return s.NormalizationRuleDescription
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationRuleFormat() *string {
	return s.NormalizationRuleFormat
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationRuleIds() []*string {
	return s.NormalizationRuleIds
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationRuleMode() *string {
	return s.NormalizationRuleMode
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *UpdateNormalizationRuleRequest) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
}

func (s *UpdateNormalizationRuleRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *UpdateNormalizationRuleRequest) GetProductId() *string {
	return s.ProductId
}

func (s *UpdateNormalizationRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNormalizationRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateNormalizationRuleRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *UpdateNormalizationRuleRequest) SetExtendContentPacked(v string) *UpdateNormalizationRuleRequest {
	s.ExtendContentPacked = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetExtendFieldStoreMode(v string) *UpdateNormalizationRuleRequest {
	s.ExtendFieldStoreMode = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetLang(v string) *UpdateNormalizationRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationCategoryId(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationRuleDescription(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationRuleDescription = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationRuleExpression(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationRuleFormat(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationRuleFormat = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationRuleId(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationRuleIds(v []*string) *UpdateNormalizationRuleRequest {
	s.NormalizationRuleIds = v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationRuleMode(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationRuleMode = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationRuleName(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationRuleName = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationRuleType(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationRuleType = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationSchemaId(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetNormalizationSecurityDomainId(v string) *UpdateNormalizationRuleRequest {
	s.NormalizationSecurityDomainId = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetOrderField(v string) *UpdateNormalizationRuleRequest {
	s.OrderField = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetProductId(v string) *UpdateNormalizationRuleRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetRegionId(v string) *UpdateNormalizationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetRoleFor(v int64) *UpdateNormalizationRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) SetVendorId(v string) *UpdateNormalizationRuleRequest {
	s.VendorId = &v
	return s
}

func (s *UpdateNormalizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
