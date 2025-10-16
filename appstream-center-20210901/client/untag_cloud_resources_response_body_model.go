// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagCloudResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedResources(v []*UntagCloudResourcesResponseBodyFailedResources) *UntagCloudResourcesResponseBody
	GetFailedResources() []*UntagCloudResourcesResponseBodyFailedResources
	SetRequestId(v string) *UntagCloudResourcesResponseBody
	GetRequestId() *string
}

type UntagCloudResourcesResponseBody struct {
	// The cloud resources whose tags failed to be removed and the corresponding tags.
	FailedResources []*UntagCloudResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E25FC620-6B6F-12D2-A992-AD8727DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagCloudResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesResponseBody) GetFailedResources() []*UntagCloudResourcesResponseBodyFailedResources {
	return s.FailedResources
}

func (s *UntagCloudResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagCloudResourcesResponseBody) SetFailedResources(v []*UntagCloudResourcesResponseBodyFailedResources) *UntagCloudResourcesResponseBody {
	s.FailedResources = v
	return s
}

func (s *UntagCloudResourcesResponseBody) SetRequestId(v string) *UntagCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagCloudResourcesResponseBody) Validate() error {
	if s.FailedResources != nil {
		for _, item := range s.FailedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UntagCloudResourcesResponseBodyFailedResources struct {
	// The error code.
	//
	// example:
	//
	// UNTAG_RESOURCE_FAILED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to untag resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The resource IDs.
	//
	// example:
	//
	// aig-00000001
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
	// The tags that failed to be removed from the cloud resources.
	Tags []*UntagCloudResourcesResponseBodyFailedResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UntagCloudResourcesResponseBodyFailedResources) String() string {
	return dara.Prettify(s)
}

func (s UntagCloudResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesResponseBodyFailedResources) GetCode() *string {
	return s.Code
}

func (s *UntagCloudResourcesResponseBodyFailedResources) GetMessage() *string {
	return s.Message
}

func (s *UntagCloudResourcesResponseBodyFailedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *UntagCloudResourcesResponseBodyFailedResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagCloudResourcesResponseBodyFailedResources) GetTags() []*UntagCloudResourcesResponseBodyFailedResourcesTags {
	return s.Tags
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetCode(v string) *UntagCloudResourcesResponseBodyFailedResources {
	s.Code = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetMessage(v string) *UntagCloudResourcesResponseBodyFailedResources {
	s.Message = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetResourceId(v string) *UntagCloudResourcesResponseBodyFailedResources {
	s.ResourceId = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetResourceType(v string) *UntagCloudResourcesResponseBodyFailedResources {
	s.ResourceType = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetTags(v []*UntagCloudResourcesResponseBodyFailedResourcesTags) *UntagCloudResourcesResponseBodyFailedResources {
	s.Tags = v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) Validate() error {
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

type UntagCloudResourcesResponseBodyFailedResourcesTags struct {
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
}

func (s UntagCloudResourcesResponseBodyFailedResourcesTags) String() string {
	return dara.Prettify(s)
}

func (s UntagCloudResourcesResponseBodyFailedResourcesTags) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesResponseBodyFailedResourcesTags) GetKey() *string {
	return s.Key
}

func (s *UntagCloudResourcesResponseBodyFailedResourcesTags) GetScope() *string {
	return s.Scope
}

func (s *UntagCloudResourcesResponseBodyFailedResourcesTags) SetKey(v string) *UntagCloudResourcesResponseBodyFailedResourcesTags {
	s.Key = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResourcesTags) SetScope(v string) *UntagCloudResourcesResponseBodyFailedResourcesTags {
	s.Scope = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResourcesTags) Validate() error {
	return dara.Validate(s)
}
