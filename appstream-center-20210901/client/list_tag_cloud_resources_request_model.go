// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagCloudResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagCloudResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagCloudResourcesRequest
	GetNextToken() *string
	SetResourceIds(v []*string) *ListTagCloudResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *ListTagCloudResourcesRequest
	GetResourceType() *string
	SetScope(v string) *ListTagCloudResourcesRequest
	GetScope() *string
}

type ListTagCloudResourcesRequest struct {
	// The number of entries per page.
	//
	// Maximum value: 1000. Default value: 50.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// ptnJAAAAAAAxNzE5OTEwNQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of resource IDs. A maximum of 50 resource IDs are supported. You do not need to specify this parameter when the resource type is tenant ID.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The cloud resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag type.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Custom
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListTagCloudResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagCloudResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagCloudResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListTagCloudResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagCloudResourcesRequest) GetScope() *string {
	return s.Scope
}

func (s *ListTagCloudResourcesRequest) SetMaxResults(v int32) *ListTagCloudResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagCloudResourcesRequest) SetNextToken(v string) *ListTagCloudResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagCloudResourcesRequest) SetResourceIds(v []*string) *ListTagCloudResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagCloudResourcesRequest) SetResourceType(v string) *ListTagCloudResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagCloudResourcesRequest) SetScope(v string) *ListTagCloudResourcesRequest {
	s.Scope = &v
	return s
}

func (s *ListTagCloudResourcesRequest) Validate() error {
	return dara.Validate(s)
}
