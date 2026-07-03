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
	SetNormalizationSecurityDomainId(v string) *ListNormalizationRulesRequest
	GetNormalizationSecurityDomainId() *string
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
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The category ID of the normalization rule.
	//
	// example:
	//
	// NETWORK_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// The list of normalization rule IDs.
	NormalizationRuleIds []*string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty" type:"Repeated"`
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
	// predefined
	NormalizationRuleType *string `json:"NormalizationRuleType,omitempty" xml:"NormalizationRuleType,omitempty"`
	// The normalization schema ID.
	//
	// example:
	//
	// HTTP_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// NETWORK_AND_WEB_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The field used for sorting.
	//
	// example:
	//
	// UpdateTime
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The sort order. Valid values:
	//
	// - desc
	//
	// - asc.
	//
	// example:
	//
	// desc
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number of the current page.
	//
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product ID.
	//
	// example:
	//
	// alibaba_cloud_sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region where the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: the Chinese mainland.
	//
	// - ap-southeast-1: outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
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

func (s *ListNormalizationRulesRequest) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
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

func (s *ListNormalizationRulesRequest) SetNormalizationSecurityDomainId(v string) *ListNormalizationRulesRequest {
	s.NormalizationSecurityDomainId = &v
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
