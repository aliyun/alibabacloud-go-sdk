// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagCloudResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceIds(v []*string) *TagCloudResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *TagCloudResourcesRequest
	GetResourceType() *string
	SetTags(v []*TagCloudResourcesRequestTags) *TagCloudResourcesRequest
	GetTags() []*TagCloudResourcesRequestTags
}

type TagCloudResourcesRequest struct {
	// The list of resource IDs. A maximum of 50 IDs are supported. You do not need to specify this parameter when the resource type is tenant ID.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The cloud resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags. System tags and custom tags are supported.
	//
	// - System tag enumeration values:
	//
	//    - `System/Scheduler/GRAYSCALE`: canary release tag
	//
	//    - `System/Scheduler/STOP_NEW_USER_CONNECTION`: tag that prevents newly bound users in a delivery group from establishing connections
	//
	// - Custom tags: A maximum of 20 custom tags can be created.
	//
	// > Each tag key on the same resource can have only one tag value. If you add a tag key that already exists, the corresponding tag value is updated to the new value.
	//
	// This parameter is required.
	Tags []*TagCloudResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagCloudResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *TagCloudResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagCloudResourcesRequest) GetTags() []*TagCloudResourcesRequestTags {
	return s.Tags
}

func (s *TagCloudResourcesRequest) SetResourceIds(v []*string) *TagCloudResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagCloudResourcesRequest) SetResourceType(v string) *TagCloudResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagCloudResourcesRequest) SetTags(v []*TagCloudResourcesRequestTags) *TagCloudResourcesRequest {
	s.Tags = v
	return s
}

func (s *TagCloudResourcesRequest) Validate() error {
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

type TagCloudResourcesRequestTags struct {
	// The tag key. Tags are case-sensitive. The key must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Resolution
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Tags are case-sensitive. The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 720p
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagCloudResourcesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s TagCloudResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesRequestTags) GetKey() *string {
	return s.Key
}

func (s *TagCloudResourcesRequestTags) GetValue() *string {
	return s.Value
}

func (s *TagCloudResourcesRequestTags) SetKey(v string) *TagCloudResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagCloudResourcesRequestTags) SetValue(v string) *TagCloudResourcesRequestTags {
	s.Value = &v
	return s
}

func (s *TagCloudResourcesRequestTags) Validate() error {
	return dara.Validate(s)
}
