// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNormalizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendContentPacked(v string) *CreateNormalizationRuleRequest
	GetExtendContentPacked() *string
	SetLang(v string) *CreateNormalizationRuleRequest
	GetLang() *string
	SetNormalizationCategoryId(v string) *CreateNormalizationRuleRequest
	GetNormalizationCategoryId() *string
	SetNormalizationRuleDescription(v string) *CreateNormalizationRuleRequest
	GetNormalizationRuleDescription() *string
	SetNormalizationRuleExpression(v string) *CreateNormalizationRuleRequest
	GetNormalizationRuleExpression() *string
	SetNormalizationRuleFormat(v string) *CreateNormalizationRuleRequest
	GetNormalizationRuleFormat() *string
	SetNormalizationRuleIds(v []*string) *CreateNormalizationRuleRequest
	GetNormalizationRuleIds() []*string
	SetNormalizationRuleMode(v string) *CreateNormalizationRuleRequest
	GetNormalizationRuleMode() *string
	SetNormalizationRuleName(v string) *CreateNormalizationRuleRequest
	GetNormalizationRuleName() *string
	SetNormalizationRuleType(v string) *CreateNormalizationRuleRequest
	GetNormalizationRuleType() *string
	SetNormalizationRuleVersion(v int32) *CreateNormalizationRuleRequest
	GetNormalizationRuleVersion() *int32
	SetNormalizationSchemaId(v string) *CreateNormalizationRuleRequest
	GetNormalizationSchemaId() *string
	SetOrderField(v string) *CreateNormalizationRuleRequest
	GetOrderField() *string
	SetProductId(v string) *CreateNormalizationRuleRequest
	GetProductId() *string
	SetRegionId(v string) *CreateNormalizationRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateNormalizationRuleRequest
	GetRoleFor() *int64
	SetVendorId(v string) *CreateNormalizationRuleRequest
	GetVendorId() *string
}

type CreateNormalizationRuleRequest struct {
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
	NormalizationRuleFormat *string   `json:"NormalizationRuleFormat,omitempty" xml:"NormalizationRuleFormat,omitempty"`
	NormalizationRuleIds    []*string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty" type:"Repeated"`
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

func (s CreateNormalizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateNormalizationRuleRequest) GetExtendContentPacked() *string {
	return s.ExtendContentPacked
}

func (s *CreateNormalizationRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateNormalizationRuleRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *CreateNormalizationRuleRequest) GetNormalizationRuleDescription() *string {
	return s.NormalizationRuleDescription
}

func (s *CreateNormalizationRuleRequest) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *CreateNormalizationRuleRequest) GetNormalizationRuleFormat() *string {
	return s.NormalizationRuleFormat
}

func (s *CreateNormalizationRuleRequest) GetNormalizationRuleIds() []*string {
	return s.NormalizationRuleIds
}

func (s *CreateNormalizationRuleRequest) GetNormalizationRuleMode() *string {
	return s.NormalizationRuleMode
}

func (s *CreateNormalizationRuleRequest) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *CreateNormalizationRuleRequest) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *CreateNormalizationRuleRequest) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *CreateNormalizationRuleRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *CreateNormalizationRuleRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *CreateNormalizationRuleRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateNormalizationRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNormalizationRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateNormalizationRuleRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *CreateNormalizationRuleRequest) SetExtendContentPacked(v string) *CreateNormalizationRuleRequest {
	s.ExtendContentPacked = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetLang(v string) *CreateNormalizationRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationCategoryId(v string) *CreateNormalizationRuleRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationRuleDescription(v string) *CreateNormalizationRuleRequest {
	s.NormalizationRuleDescription = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationRuleExpression(v string) *CreateNormalizationRuleRequest {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationRuleFormat(v string) *CreateNormalizationRuleRequest {
	s.NormalizationRuleFormat = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationRuleIds(v []*string) *CreateNormalizationRuleRequest {
	s.NormalizationRuleIds = v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationRuleMode(v string) *CreateNormalizationRuleRequest {
	s.NormalizationRuleMode = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationRuleName(v string) *CreateNormalizationRuleRequest {
	s.NormalizationRuleName = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationRuleType(v string) *CreateNormalizationRuleRequest {
	s.NormalizationRuleType = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationRuleVersion(v int32) *CreateNormalizationRuleRequest {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetNormalizationSchemaId(v string) *CreateNormalizationRuleRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetOrderField(v string) *CreateNormalizationRuleRequest {
	s.OrderField = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetProductId(v string) *CreateNormalizationRuleRequest {
	s.ProductId = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetRegionId(v string) *CreateNormalizationRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetRoleFor(v int64) *CreateNormalizationRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateNormalizationRuleRequest) SetVendorId(v string) *CreateNormalizationRuleRequest {
	s.VendorId = &v
	return s
}

func (s *CreateNormalizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
