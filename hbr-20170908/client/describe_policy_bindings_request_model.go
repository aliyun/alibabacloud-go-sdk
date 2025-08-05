// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceIds(v []*string) *DescribePolicyBindingsRequest
	GetDataSourceIds() []*string
	SetFilters(v []*DescribePolicyBindingsRequestFilters) *DescribePolicyBindingsRequest
	GetFilters() []*DescribePolicyBindingsRequestFilters
	SetMaxResults(v int32) *DescribePolicyBindingsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePolicyBindingsRequest
	GetNextToken() *string
	SetPolicyId(v string) *DescribePolicyBindingsRequest
	GetPolicyId() *string
	SetSourceType(v string) *DescribePolicyBindingsRequest
	GetSourceType() *string
}

type DescribePolicyBindingsRequest struct {
	// List of data source IDs.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// Query filters.
	Filters []*DescribePolicyBindingsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// Number of results per query.
	//
	// Range: 10~100. Default: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token required to fetch the next page of policy and data source associations.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// po-000************hky
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Data source type. Possible values:
	//
	// 	- **UDM_ECS**: Indicates ECS full machine backup.
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribePolicyBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsRequest) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *DescribePolicyBindingsRequest) GetFilters() []*DescribePolicyBindingsRequestFilters {
	return s.Filters
}

func (s *DescribePolicyBindingsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePolicyBindingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePolicyBindingsRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribePolicyBindingsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribePolicyBindingsRequest) SetDataSourceIds(v []*string) *DescribePolicyBindingsRequest {
	s.DataSourceIds = v
	return s
}

func (s *DescribePolicyBindingsRequest) SetFilters(v []*DescribePolicyBindingsRequestFilters) *DescribePolicyBindingsRequest {
	s.Filters = v
	return s
}

func (s *DescribePolicyBindingsRequest) SetMaxResults(v int32) *DescribePolicyBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePolicyBindingsRequest) SetNextToken(v string) *DescribePolicyBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyBindingsRequest) SetPolicyId(v string) *DescribePolicyBindingsRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribePolicyBindingsRequest) SetSourceType(v string) *DescribePolicyBindingsRequest {
	s.SourceType = &v
	return s
}

func (s *DescribePolicyBindingsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyBindingsRequestFilters struct {
	// Key in the query filter. Possible values include:
	//
	// - **PolicyId**: Backup policy ID
	//
	// - **DataSourceId**: ECS instance ID
	//
	// - **DataSourceType**: Data source type
	//
	// example:
	//
	// DataSourceType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Matching method. Default is IN. This refers to the matching operation (Operator) supported by the Key and Value in the filter. Possible values include:
	//
	// - **EQUAL**: Equal to
	//
	// - **NOT_EQUAL**: Not equal to
	//
	// - **GREATER_THAN**: Greater than
	//
	// - **GREATER_THAN_OR_EQUAL**: Greater than or equal to
	//
	// - **LESS_THAN**: Less than
	//
	// - **LESS_THAN_OR_EQUAL**: Less than or equal to
	//
	// - **BETWEEN**: Range, where value is a JSON array `[lower_bound, upper_bound]`.
	//
	// - **IN**: In the set, where value is an array.
	//
	// example:
	//
	// IN
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Values to be matched in the query filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribePolicyBindingsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribePolicyBindingsRequestFilters) GetOperator() *string {
	return s.Operator
}

func (s *DescribePolicyBindingsRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *DescribePolicyBindingsRequestFilters) SetKey(v string) *DescribePolicyBindingsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribePolicyBindingsRequestFilters) SetOperator(v string) *DescribePolicyBindingsRequestFilters {
	s.Operator = &v
	return s
}

func (s *DescribePolicyBindingsRequestFilters) SetValues(v []*string) *DescribePolicyBindingsRequestFilters {
	s.Values = v
	return s
}

func (s *DescribePolicyBindingsRequestFilters) Validate() error {
	return dara.Validate(s)
}
