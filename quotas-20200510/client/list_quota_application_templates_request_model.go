// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v []*ListQuotaApplicationTemplatesRequestDimensions) *ListQuotaApplicationTemplatesRequest
	GetDimensions() []*ListQuotaApplicationTemplatesRequestDimensions
	SetId(v string) *ListQuotaApplicationTemplatesRequest
	GetId() *string
	SetMaxResults(v int32) *ListQuotaApplicationTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaApplicationTemplatesRequest
	GetNextToken() *string
	SetProductCode(v string) *ListQuotaApplicationTemplatesRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ListQuotaApplicationTemplatesRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *ListQuotaApplicationTemplatesRequest
	GetQuotaCategory() *string
}

type ListQuotaApplicationTemplatesRequest struct {
	// The quota dimensions.
	Dimensions []*ListQuotaApplicationTemplatesRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The ID of the quota item.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum number of records that can be returned for the query. Valid values: 1 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// > If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440555.html) operation and check the value of `ProductCode` in the response.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota type. Valid values:
	//
	// 	- CommonQuota: general quota.
	//
	// 	- WhiteListLabel: whitelist quota.
	//
	// 	- FlowControl: API rate limit.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ListQuotaApplicationTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesRequest) GetDimensions() []*ListQuotaApplicationTemplatesRequestDimensions {
	return s.Dimensions
}

func (s *ListQuotaApplicationTemplatesRequest) GetId() *string {
	return s.Id
}

func (s *ListQuotaApplicationTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaApplicationTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaApplicationTemplatesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaApplicationTemplatesRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaApplicationTemplatesRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListQuotaApplicationTemplatesRequest) SetDimensions(v []*ListQuotaApplicationTemplatesRequestDimensions) *ListQuotaApplicationTemplatesRequest {
	s.Dimensions = v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetId(v string) *ListQuotaApplicationTemplatesRequest {
	s.Id = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetMaxResults(v int32) *ListQuotaApplicationTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetNextToken(v string) *ListQuotaApplicationTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetProductCode(v string) *ListQuotaApplicationTemplatesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetQuotaActionCode(v string) *ListQuotaApplicationTemplatesRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) SetQuotaCategory(v string) *ListQuotaApplicationTemplatesRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

type ListQuotaApplicationTemplatesRequestDimensions struct {
	// The key of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- You must specify both Key and Value for each quota dimension.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- You must specify both Key and Value for each quota dimension.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotaApplicationTemplatesRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationTemplatesRequestDimensions) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *ListQuotaApplicationTemplatesRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *ListQuotaApplicationTemplatesRequestDimensions) SetKey(v string) *ListQuotaApplicationTemplatesRequestDimensions {
	s.Key = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequestDimensions) SetValue(v string) *ListQuotaApplicationTemplatesRequestDimensions {
	s.Value = &v
	return s
}

func (s *ListQuotaApplicationTemplatesRequestDimensions) Validate() error {
	return dara.Validate(s)
}
