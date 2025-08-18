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
	// The resource IDs. You can specify up to 50 resource IDs. You do not need to specify this parameter if you set ResourceType to AliUid.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the resource from which you want to remove tags.
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
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to remove from the cloud resources. System and custom tags are supported. You can specify up to 10 tags.
	//
	// Valid values for system tags:
	//
	// 	- `System/Scheduler/GRAYSCALE`: canary tags.
	//
	// 	- `System/Scheduler/STOP_NEW_USER_CONNECTION`: tags used to stop new users bound to the delivery group from establishing a connection.
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
