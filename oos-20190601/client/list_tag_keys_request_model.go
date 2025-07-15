// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagKeysRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagKeysRequest
	GetRegionId() *string
	SetResourceType(v string) *ListTagKeysRequest
	GetResourceType() *string
}

type ListTagKeysRequest struct {
	// The maximum number of entries to return on each page. Valid value: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page.
	//
	// example:
	//
	// djsdlkasd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource to which the tag is added.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
