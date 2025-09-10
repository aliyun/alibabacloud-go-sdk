// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v []*ListProductQuotasRequestDimensions) *ListProductQuotasRequest
	GetDimensions() []*ListProductQuotasRequestDimensions
	SetGroupCode(v string) *ListProductQuotasRequest
	GetGroupCode() *string
	SetKeyWord(v string) *ListProductQuotasRequest
	GetKeyWord() *string
	SetMaxResults(v int32) *ListProductQuotasRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductQuotasRequest
	GetNextToken() *string
	SetProductCode(v string) *ListProductQuotasRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ListProductQuotasRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *ListProductQuotasRequest
	GetQuotaCategory() *string
	SetSortField(v string) *ListProductQuotasRequest
	GetSortField() *string
	SetSortOrder(v string) *ListProductQuotasRequest
	GetSortOrder() *string
}

type ListProductQuotasRequest struct {
	// The quota dimensions.
	Dimensions []*ListProductQuotasRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The code of the dimension group.
	//
	// example:
	//
	// entconsole_w1j3msbo2g
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The keyword that you want to use to search for the quotas.
	//
	// >  This parameter is available only for quotas that belong to ECS Quotas by Instance Type. The keyword is used to match the values of `QuotaName` and `QuotaActionCode`.
	//
	// example:
	//
	// ecs-spec
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// 	- You do not need to specify this parameter for the first and last requests.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of the `ProductCode` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The type of the quota.
	//
	// Default value: CommonQuota.
	//
	// Valid values:
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: privilege
	//
	// 	- CommonQuota: general quota
	//
	// example:
	//
	// FlowControl
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The field based on which you want to sort the returned records.
	//
	// >  This parameter is available only for quotas that belong to ECS Quotas by Instance Type. You can leave this parameter empty.
	//
	// Valid values:
	//
	// 	- TOTAL: sorts the returned records based on the total quota.
	//
	// 	- TIME: sorts the returned records based on the time when the quota was last modified.
	//
	// 	- RESERVED: sorts the returned records based on the reserved quota.
	//
	// example:
	//
	// TIME
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the returned records.
	//
	// >  This parameter is available only for quotas that belong to ECS Quotas by Instance Type. You can leave this parameter empty.
	//
	// Valid values:
	//
	// 	- Descending
	//
	// 	- Ascending
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProductQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListProductQuotasRequest) GetDimensions() []*ListProductQuotasRequestDimensions {
	return s.Dimensions
}

func (s *ListProductQuotasRequest) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ListProductQuotasRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListProductQuotasRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductQuotasRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductQuotasRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductQuotasRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListProductQuotasRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListProductQuotasRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListProductQuotasRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListProductQuotasRequest) SetDimensions(v []*ListProductQuotasRequestDimensions) *ListProductQuotasRequest {
	s.Dimensions = v
	return s
}

func (s *ListProductQuotasRequest) SetGroupCode(v string) *ListProductQuotasRequest {
	s.GroupCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetKeyWord(v string) *ListProductQuotasRequest {
	s.KeyWord = &v
	return s
}

func (s *ListProductQuotasRequest) SetMaxResults(v int32) *ListProductQuotasRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotasRequest) SetNextToken(v string) *ListProductQuotasRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotasRequest) SetProductCode(v string) *ListProductQuotasRequest {
	s.ProductCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetQuotaActionCode(v string) *ListProductQuotasRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListProductQuotasRequest) SetQuotaCategory(v string) *ListProductQuotasRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListProductQuotasRequest) SetSortField(v string) *ListProductQuotasRequest {
	s.SortField = &v
	return s
}

func (s *ListProductQuotasRequest) SetSortOrder(v string) *ListProductQuotasRequest {
	s.SortOrder = &v
	return s
}

func (s *ListProductQuotasRequest) Validate() error {
	return dara.Validate(s)
}

type ListProductQuotasRequestDimensions struct {
	// The key of the dimension.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >  The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductQuotasRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasRequestDimensions) GoString() string {
	return s.String()
}

func (s *ListProductQuotasRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *ListProductQuotasRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *ListProductQuotasRequestDimensions) SetKey(v string) *ListProductQuotasRequestDimensions {
	s.Key = &v
	return s
}

func (s *ListProductQuotasRequestDimensions) SetValue(v string) *ListProductQuotasRequestDimensions {
	s.Value = &v
	return s
}

func (s *ListProductQuotasRequestDimensions) Validate() error {
	return dara.Validate(s)
}
