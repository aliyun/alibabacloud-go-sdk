// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationRulesRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationRulesRequest
	GetNextToken() *string
	SetNormalizationCategoryId(v string) *ListNormalizationRulesRequest
	GetNormalizationCategoryId() *string
	SetNormalizationRuleIds(v []*string) *ListNormalizationRulesRequest
	GetNormalizationRuleIds() []*string
	SetNormalizationRuleName(v string) *ListNormalizationRulesRequest
	GetNormalizationRuleName() *string
	SetNormalizationRuleType(v string) *ListNormalizationRulesRequest
	GetNormalizationRuleType() *string
	SetNormalizationSchemaId(v string) *ListNormalizationRulesRequest
	GetNormalizationSchemaId() *string
	SetOrderField(v string) *ListNormalizationRulesRequest
	GetOrderField() *string
	SetOrderType(v string) *ListNormalizationRulesRequest
	GetOrderType() *string
	SetPageNumber(v int32) *ListNormalizationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNormalizationRulesRequest
	GetPageSize() *int32
	SetProductId(v string) *ListNormalizationRulesRequest
	GetProductId() *string
	SetRegionId(v string) *ListNormalizationRulesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationRulesRequest
	GetRoleFor() *int64
	SetVendorId(v string) *ListNormalizationRulesRequest
	GetVendorId() *string
}

type ListNormalizationRulesRequest struct {
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
	NormalizationCategoryId *string   `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	NormalizationRuleIds    []*string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty" type:"Repeated"`
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

func (s ListNormalizationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationRulesRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ListNormalizationRulesRequest) GetNormalizationRuleIds() []*string {
	return s.NormalizationRuleIds
}

func (s *ListNormalizationRulesRequest) GetNormalizationRuleName() *string {
	return s.NormalizationRuleName
}

func (s *ListNormalizationRulesRequest) GetNormalizationRuleType() *string {
	return s.NormalizationRuleType
}

func (s *ListNormalizationRulesRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *ListNormalizationRulesRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListNormalizationRulesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListNormalizationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNormalizationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNormalizationRulesRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListNormalizationRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationRulesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationRulesRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *ListNormalizationRulesRequest) SetLang(v string) *ListNormalizationRulesRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetMaxResults(v int32) *ListNormalizationRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetNextToken(v string) *ListNormalizationRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetNormalizationCategoryId(v string) *ListNormalizationRulesRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetNormalizationRuleIds(v []*string) *ListNormalizationRulesRequest {
	s.NormalizationRuleIds = v
	return s
}

func (s *ListNormalizationRulesRequest) SetNormalizationRuleName(v string) *ListNormalizationRulesRequest {
	s.NormalizationRuleName = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetNormalizationRuleType(v string) *ListNormalizationRulesRequest {
	s.NormalizationRuleType = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetNormalizationSchemaId(v string) *ListNormalizationRulesRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetOrderField(v string) *ListNormalizationRulesRequest {
	s.OrderField = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetOrderType(v string) *ListNormalizationRulesRequest {
	s.OrderType = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetPageNumber(v int32) *ListNormalizationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetPageSize(v int32) *ListNormalizationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetProductId(v string) *ListNormalizationRulesRequest {
	s.ProductId = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetRegionId(v string) *ListNormalizationRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetRoleFor(v int64) *ListNormalizationRulesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationRulesRequest) SetVendorId(v string) *ListNormalizationRulesRequest {
	s.VendorId = &v
	return s
}

func (s *ListNormalizationRulesRequest) Validate() error {
	return dara.Validate(s)
}
