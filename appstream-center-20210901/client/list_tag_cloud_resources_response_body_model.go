// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagCloudResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagCloudResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagCloudResourcesResponseBody
	GetRequestId() *string
	SetResourceTags(v []*ListTagCloudResourcesResponseBodyResourceTags) *ListTagCloudResourcesResponseBody
	GetResourceTags() []*ListTagCloudResourcesResponseBodyResourceTags
	SetTotalCount(v int32) *ListTagCloudResourcesResponseBody
	GetTotalCount() *int32
}

type ListTagCloudResourcesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// example:
	//
	// AAAAAYRHtOLVQzCYj17y+OP7LZRrUJaF4rnBGQkWwMiVHlLZBB1w3Us37CVvhvyM0TXavA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags added to the cloud resources.
	ResourceTags []*ListTagCloudResourcesResponseBodyResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagCloudResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagCloudResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagCloudResourcesResponseBody) GetResourceTags() []*ListTagCloudResourcesResponseBodyResourceTags {
	return s.ResourceTags
}

func (s *ListTagCloudResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTagCloudResourcesResponseBody) SetNextToken(v string) *ListTagCloudResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagCloudResourcesResponseBody) SetRequestId(v string) *ListTagCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagCloudResourcesResponseBody) SetResourceTags(v []*ListTagCloudResourcesResponseBodyResourceTags) *ListTagCloudResourcesResponseBody {
	s.ResourceTags = v
	return s
}

func (s *ListTagCloudResourcesResponseBody) SetTotalCount(v int32) *ListTagCloudResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagCloudResourcesResponseBody) Validate() error {
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagCloudResourcesResponseBodyResourceTags struct {
	// The resource ID.
	//
	// example:
	//
	// aig-0001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the cloud resource.
	//
	// Valid values:
	//
	// 	- AppId: app ID.
	//
	// 	- WyId: Alibaba Cloud Workspace user ID.
	//
	// 	- AppInstanceGroupId: delivery group ID.
	//
	// 	- AliUid: tenant ID.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*ListTagCloudResourcesResponseBodyResourceTagsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagCloudResourcesResponseBodyResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagCloudResourcesResponseBodyResourceTags) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) GetTags() []*ListTagCloudResourcesResponseBodyResourceTagsTags {
	return s.Tags
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) SetResourceId(v string) *ListTagCloudResourcesResponseBodyResourceTags {
	s.ResourceId = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) SetResourceType(v string) *ListTagCloudResourcesResponseBodyResourceTags {
	s.ResourceType = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) SetTags(v []*ListTagCloudResourcesResponseBodyResourceTagsTags) *ListTagCloudResourcesResponseBodyResourceTags {
	s.Tags = v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagCloudResourcesResponseBodyResourceTagsTags struct {
	// The tag key.
	//
	// example:
	//
	// Resolution
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag type.
	//
	// Valid values:
	//
	// 	- Custom: custom tag.
	//
	// 	- System: system tag.
	//
	// example:
	//
	// Custom
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1080p
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagCloudResourcesResponseBodyResourceTagsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagCloudResourcesResponseBodyResourceTagsTags) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) GetKey() *string {
	return s.Key
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) GetScope() *string {
	return s.Scope
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) GetValue() *string {
	return s.Value
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) SetKey(v string) *ListTagCloudResourcesResponseBodyResourceTagsTags {
	s.Key = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) SetScope(v string) *ListTagCloudResourcesResponseBodyResourceTagsTags {
	s.Scope = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) SetValue(v string) *ListTagCloudResourcesResponseBodyResourceTagsTags {
	s.Value = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) Validate() error {
	return dara.Validate(s)
}
