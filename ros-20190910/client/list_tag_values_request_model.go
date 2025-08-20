// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ListTagValuesRequest
	GetKey() *string
	SetNextToken(v string) *ListTagValuesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagValuesRequest
	GetRegionId() *string
	SetResourceType(v string) *ListTagValuesRequest
	GetResourceType() *string
}

type ListTagValuesRequest struct {
	// The key of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// TagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the tag value. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- stack: stack
	//
	// 	- stackgroup: stack group
	//
	// 	- template: template
	//
	// 	- templatescratch: scenario
	//
	// This parameter is required.
	//
	// example:
	//
	// stack
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) GetKey() *string {
	return s.Key
}

func (s *ListTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagValuesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
