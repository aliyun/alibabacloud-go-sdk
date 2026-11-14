// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagCloudResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceIds(v []*string) *UntagCloudResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *UntagCloudResourcesRequest
	GetResourceType() *string
	SetTagKeys(v []*string) *UntagCloudResourcesRequest
	GetTagKeys() []*string
}

type UntagCloudResourcesRequest struct {
	// The list of resource IDs. A maximum of 50 resource IDs are supported. This parameter is not required when the resource type is tenant ID.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tag keys. Both system tags and custom tags are supported. A maximum of 10 tag keys can be passed.
	//
	// Valid values for system tags:
	//
	// - `System/Scheduler/GRAYSCALE`: Canary release tag.
	//
	// - `System/Scheduler/STOP_NEW_USER_CONNECTION`: Tag that prevents new users bound to the delivery group from establishing connections.
	//
	// This parameter is required.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s UntagCloudResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *UntagCloudResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagCloudResourcesRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *UntagCloudResourcesRequest) SetResourceIds(v []*string) *UntagCloudResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagCloudResourcesRequest) SetResourceType(v string) *UntagCloudResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagCloudResourcesRequest) SetTagKeys(v []*string) *UntagCloudResourcesRequest {
	s.TagKeys = v
	return s
}

func (s *UntagCloudResourcesRequest) Validate() error {
	return dara.Validate(s)
}
