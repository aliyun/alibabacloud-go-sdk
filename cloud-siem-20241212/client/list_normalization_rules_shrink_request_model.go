// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationRulesShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationRulesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationRulesShrinkRequest
	GetNextToken() *string
	SetNormalizationCategoryId(v string) *ListNormalizationRulesShrinkRequest
	GetNormalizationCategoryId() *string
	SetNormalizationRuleIdsShrink(v string) *ListNormalizationRulesShrinkRequest
	GetNormalizationRuleIdsShrink() *string
	SetNormalizationRuleName(v string) *ListNormalizationRulesShrinkRequest
	GetNormalizationRuleName() *string
	SetNormalizationRuleType(v string) *ListNormalizationRulesShrinkRequest
	GetNormalizationRuleType() *string
	SetNormalizationSchemaId(v string) *ListNormalizationRulesShrinkRequest
	GetNormalizationSchemaId() *string
	SetOrderField(v string) *ListNormalizationRulesShrinkRequest
	GetOrderField() *string
	SetOrderType(v string) *ListNormalizationRulesShrinkRequest
	GetOrderType() *string
	SetPageNumber(v int32) *ListNormalizationRulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNormalizationRulesShrinkRequest
	GetPageSize() *int32
	SetProductId(v string) *ListNormalizationRulesShrinkRequest
	GetProductId() *string
	SetRegionId(v string) *ListNormalizationRulesShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationRulesShrinkRequest
	GetRoleFor() *int64
	SetVendorId(v string) *ListNormalizationRulesShrinkRequest
	GetVendorId() *string
}

type ListNormalizationRulesShrinkRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// NETWORK_CATEGORY。
	NormalizationCategoryId    *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	NormalizationRuleIdsShrink *string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty"`
	// example:
	//
	// normalization_rule_Z57np。
	NormalizationRuleName *string `json:"NormalizationRuleName,omitempty" xml:"NormalizationRuleName,omitempty"`
	// example:
	//
	// predefined。
	NormalizationRuleType *string `json:"NormalizationRuleType,omitempty" xml:"NormalizationRuleType,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// UpdateTime。
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// example:
	//
	// desc。
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 3。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s ListNormalizationRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationRulesShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationRulesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationRulesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationRulesShrinkRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ListNormalizationRulesShrinkRequest) GetNormalizationRuleIdsShrink() *string {
	return s.NormalizationRuleIdsShrink
}

func (s *ListNormalizationRulesShrinkRequest) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *ListNormalizationRulesShrinkRequest) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *ListNormalizationRulesShrinkRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *ListNormalizationRulesShrinkRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListNormalizationRulesShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListNormalizationRulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNormalizationRulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNormalizationRulesShrinkRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListNormalizationRulesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationRulesShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationRulesShrinkRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *ListNormalizationRulesShrinkRequest) SetLang(v string) *ListNormalizationRulesShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetMaxResults(v int32) *ListNormalizationRulesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetNextToken(v string) *ListNormalizationRulesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetNormalizationCategoryId(v string) *ListNormalizationRulesShrinkRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetNormalizationRuleIdsShrink(v string) *ListNormalizationRulesShrinkRequest {
	s.NormalizationRuleIdsShrink = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetNormalizationRuleName(v string) *ListNormalizationRulesShrinkRequest {
	s.NormalizationRuleName = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetNormalizationRuleType(v string) *ListNormalizationRulesShrinkRequest {
	s.NormalizationRuleType = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetNormalizationSchemaId(v string) *ListNormalizationRulesShrinkRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetOrderField(v string) *ListNormalizationRulesShrinkRequest {
	s.OrderField = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetOrderType(v string) *ListNormalizationRulesShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetPageNumber(v int32) *ListNormalizationRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetPageSize(v int32) *ListNormalizationRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetProductId(v string) *ListNormalizationRulesShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetRegionId(v string) *ListNormalizationRulesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetRoleFor(v int64) *ListNormalizationRulesShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) SetVendorId(v string) *ListNormalizationRulesShrinkRequest {
	s.VendorId = &v
	return s
}

func (s *ListNormalizationRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
