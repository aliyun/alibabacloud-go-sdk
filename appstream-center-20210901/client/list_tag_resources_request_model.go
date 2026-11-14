// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *ListTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest
	GetTag() []*ListTagResourcesRequestTag
}

type ListTagResourcesRequest struct {
	// The pagination token for paged queries. You do not need to specify this parameter when you call this operation for the first time. If a non-empty NextToken is returned in the previous call, pass it as is to obtain subsequent results.
	//
	// example:
	//
	// AAAAAYRHtOLVQzCYj17y+OP7LZRrUJaF4rnBGQkWwMiVHlLZBB1w3Us37CVvhvyM0TXavA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs. You can specify up to 50 resource IDs.
	//
	// If the resource type is delivery group, specify the delivery group ID, which starts with `aig-`. You can obtain the ID by calling the ListAppInstanceGroup operation.
	//
	// > You must specify at least one of ResourceId.N and Tag.N. If only ResourceId.N is specified, all tags bound to these resources are returned.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// - APPINSTANCEGROUP: delivery group.
	//
	// The value is not case-sensitive. Only this value is supported. If you specify other values, the error code `InvalidResourceType.Invalid` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// APPINSTANCEGROUP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tag conditions. You can specify up to 20 tag conditions.
	//
	// > You must specify at least one of ResourceId.N and Tag.N. Multiple tags have an AND relationship. Only resources bound to all specified tags are returned.
	//
	// if can be null:
	// true
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetTag() []*ListTagResourcesRequestTag {
	return s.Tag
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesRequestTag struct {
	// The tag key. This parameter is required. The value is case-sensitive and must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Resolution
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value is case-sensitive and can be up to 256 characters in length. If this parameter is not specified, the value corresponding to the tag key is not limited, which means any tag value under the key is matched.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 720p
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *ListTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
