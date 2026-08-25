// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyBindingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceIdsShrink(v string) *DescribePolicyBindingsShrinkRequest
	GetDataSourceIdsShrink() *string
	SetFilters(v []*DescribePolicyBindingsShrinkRequestFilters) *DescribePolicyBindingsShrinkRequest
	GetFilters() []*DescribePolicyBindingsShrinkRequestFilters
	SetMaxResults(v int32) *DescribePolicyBindingsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePolicyBindingsShrinkRequest
	GetNextToken() *string
	SetPolicyId(v string) *DescribePolicyBindingsShrinkRequest
	GetPolicyId() *string
	SetSourceType(v string) *DescribePolicyBindingsShrinkRequest
	GetSourceType() *string
}

type DescribePolicyBindingsShrinkRequest struct {
	// The list of data source IDs.
	DataSourceIdsShrink *string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty"`
	// The query filters.
	Filters []*DescribePolicyBindingsShrinkRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of results for each query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token required to obtain the next page of policy-data source associations.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// po-000************hky
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The data source type. Valid values:
	//
	// - **UDM_ECS**: ECS instance backup.
	//
	// - **OSS**: OSS backup.
	//
	// - **NAS**: Alibaba Cloud NAS backup.
	//
	// - **COMMON_NAS**: On-premises NAS backup.
	//
	// - **ECS_FILE**: ECS File Backup Essential Edition.
	//
	// - **File**: On-premises file backup.
	//
	// - **COMMON_FILE_SYSTEM**: CPFS backup.
	//
	// - **OTS**: Tablestore backup.
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribePolicyBindingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsShrinkRequest) GetDataSourceIdsShrink() *string {
	return s.DataSourceIdsShrink
}

func (s *DescribePolicyBindingsShrinkRequest) GetFilters() []*DescribePolicyBindingsShrinkRequestFilters {
	return s.Filters
}

func (s *DescribePolicyBindingsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePolicyBindingsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePolicyBindingsShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribePolicyBindingsShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribePolicyBindingsShrinkRequest) SetDataSourceIdsShrink(v string) *DescribePolicyBindingsShrinkRequest {
	s.DataSourceIdsShrink = &v
	return s
}

func (s *DescribePolicyBindingsShrinkRequest) SetFilters(v []*DescribePolicyBindingsShrinkRequestFilters) *DescribePolicyBindingsShrinkRequest {
	s.Filters = v
	return s
}

func (s *DescribePolicyBindingsShrinkRequest) SetMaxResults(v int32) *DescribePolicyBindingsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePolicyBindingsShrinkRequest) SetNextToken(v string) *DescribePolicyBindingsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyBindingsShrinkRequest) SetPolicyId(v string) *DescribePolicyBindingsShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribePolicyBindingsShrinkRequest) SetSourceType(v string) *DescribePolicyBindingsShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *DescribePolicyBindingsShrinkRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolicyBindingsShrinkRequestFilters struct {
	// The key in the query filter. Valid values:
	//
	// - **PolicyId**: backup policy ID
	//
	// - **DataSourceId**: ECS instance ID
	//
	// - **DataSourceType**: data source type
	//
	// example:
	//
	// DataSourceType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method. Default value: IN. The matching operation (Operator) supported by the Key and Value in the filter. Valid values:
	//
	// - **EQUAL**: equal to
	//
	// - **NOT_EQUAL**: not equal to
	//
	// - **GREATER_THAN**: greater than
	//
	// - **GREATER_THAN_OR_EQUAL**: greater than or equal to
	//
	// - **LESS_THAN**: less than
	//
	// - **LESS_THAN_OR_EQUAL**: less than or equal to
	//
	// - **BETWEEN**: range. The value is a JSON array `[lower bound, upper bound]`.
	//
	// - **IN**: in a set. The value is an array.
	//
	// example:
	//
	// IN
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The values to match in the query filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribePolicyBindingsShrinkRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsShrinkRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsShrinkRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribePolicyBindingsShrinkRequestFilters) GetOperator() *string {
	return s.Operator
}

func (s *DescribePolicyBindingsShrinkRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *DescribePolicyBindingsShrinkRequestFilters) SetKey(v string) *DescribePolicyBindingsShrinkRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribePolicyBindingsShrinkRequestFilters) SetOperator(v string) *DescribePolicyBindingsShrinkRequestFilters {
	s.Operator = &v
	return s
}

func (s *DescribePolicyBindingsShrinkRequestFilters) SetValues(v []*string) *DescribePolicyBindingsShrinkRequestFilters {
	s.Values = v
	return s
}

func (s *DescribePolicyBindingsShrinkRequestFilters) Validate() error {
	return dara.Validate(s)
}
