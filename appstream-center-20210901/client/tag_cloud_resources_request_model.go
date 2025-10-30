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
	// The resource IDs. You can specify up to 50 resource IDs. You do not need to specify this parameter if you set ResourceType to AliUid.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the cloud resource.
	//
	// Valid values:
	//
	// 	- CenterImageId: center image ID.
	//
	// 	- AppId: app ID.
	//
	// 	- WyId: Alibaba Cloud Workspace user ID.
	//
	// 	- AppInstanceGroupId: delivery group ID.
	//
	// 	- AliUid: tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to remove from the cloud resources. System and custom tags are supported.
	//
	// 	- Valid values for system tags:
	//
	//     	- `System/Scheduler/GRAYSCALE`: canary tags.
	//
	//     	- `System/Scheduler/STOP_NEW_USER_CONNECTION`: tags used to stop new users bound to the delivery group from establishing a connection.
	//
	// 	- You can create up to 20 custom tags.
	//
	// > Each tag key on a resource can have only one tag value. If you create a tag that has the same key as an existing tag, the value of the existing tag is overwritten.
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
	// The key of a tag. The value must be 1 to 128 characters in length and is case-sensitive. The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Resolution
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag. The value must be 1 to 128 characters in length and is case-sensitive. The name must be 1 to 128 characters in length.
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
