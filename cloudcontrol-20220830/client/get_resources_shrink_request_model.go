// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *GetResourcesShrinkRequest
	GetFilterShrink() *string
	SetMaxResults(v int32) *GetResourcesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetResourcesShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *GetResourcesShrinkRequest
	GetRegionId() *string
}

type GetResourcesShrinkRequest struct {
	// The filter conditions for resources.
	//
	// Specify multiple key-value pairs in JSON format to filter resources. If a List or Get operation for a cloud product supports filtering by specific properties, you can use those properties as filter conditions for this parameter.
	//
	// > The supported filter fields may vary for different resource types. For more information about the supported fields, see the OpenAPI documentation for the specific resource.
	//
	// For example, DBInstance resources support filtering by the `EditionType` and `PaymentType` fields.
	//
	// example:
	//
	// {
	//
	//   "EditionType": "Community",
	//
	//   "PaymentType": "PostPaid"
	//
	// }
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The maximum number of records to return on each page for a paged query. Maximum value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// - You do not need to specify this parameter for the first query. The system returns data from the first page.
	//
	// - For subsequent queries, set this parameter to the nextToken value returned from the previous call.
	//
	// > If this parameter contains only digits, Cloud Control API treats it as the `PageNumber` for paging.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID. This parameter is required if the cloud product is region-specific.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetResourcesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourcesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetResourcesShrinkRequest) SetFilterShrink(v string) *GetResourcesShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetMaxResults(v int32) *GetResourcesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetNextToken(v string) *GetResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetRegionId(v string) *GetResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
