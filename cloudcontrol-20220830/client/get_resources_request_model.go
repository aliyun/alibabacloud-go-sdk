// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v map[string]interface{}) *GetResourcesRequest
	GetFilter() map[string]interface{}
	SetMaxResults(v int32) *GetResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *GetResourcesRequest
	GetRegionId() *string
}

type GetResourcesRequest struct {
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
	Filter map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
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

func (s GetResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesRequest) GetFilter() map[string]interface{} {
	return s.Filter
}

func (s *GetResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetResourcesRequest) SetFilter(v map[string]interface{}) *GetResourcesRequest {
	s.Filter = v
	return s
}

func (s *GetResourcesRequest) SetMaxResults(v int32) *GetResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResourcesRequest) SetNextToken(v string) *GetResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourcesRequest) SetRegionId(v string) *GetResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *GetResourcesRequest) Validate() error {
	return dara.Validate(s)
}
