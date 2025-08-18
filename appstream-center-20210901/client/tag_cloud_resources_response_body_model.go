// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagCloudResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedResources(v []*TagCloudResourcesResponseBodyFailedResources) *TagCloudResourcesResponseBody
	GetFailedResources() []*TagCloudResourcesResponseBodyFailedResources
	SetRequestId(v string) *TagCloudResourcesResponseBody
	GetRequestId() *string
}

type TagCloudResourcesResponseBody struct {
	FailedResources []*TagCloudResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagCloudResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesResponseBody) GetFailedResources() []*TagCloudResourcesResponseBodyFailedResources {
	return s.FailedResources
}

func (s *TagCloudResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagCloudResourcesResponseBody) SetFailedResources(v []*TagCloudResourcesResponseBodyFailedResources) *TagCloudResourcesResponseBody {
	s.FailedResources = v
	return s
}

func (s *TagCloudResourcesResponseBody) SetRequestId(v string) *TagCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagCloudResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type TagCloudResourcesResponseBodyFailedResources struct {
	// example:
	//
	// TAG_KEY_DUPLICATED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Duplicate tag keys exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// aig-001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string                                             `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags         []*TagCloudResourcesResponseBodyFailedResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagCloudResourcesResponseBodyFailedResources) String() string {
	return dara.Prettify(s)
}

func (s TagCloudResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesResponseBodyFailedResources) GetCode() *string {
	return s.Code
}

func (s *TagCloudResourcesResponseBodyFailedResources) GetMessage() *string {
	return s.Message
}

func (s *TagCloudResourcesResponseBodyFailedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *TagCloudResourcesResponseBodyFailedResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagCloudResourcesResponseBodyFailedResources) GetTags() []*TagCloudResourcesResponseBodyFailedResourcesTags {
	return s.Tags
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetCode(v string) *TagCloudResourcesResponseBodyFailedResources {
	s.Code = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetMessage(v string) *TagCloudResourcesResponseBodyFailedResources {
	s.Message = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetResourceId(v string) *TagCloudResourcesResponseBodyFailedResources {
	s.ResourceId = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetResourceType(v string) *TagCloudResourcesResponseBodyFailedResources {
	s.ResourceType = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetTags(v []*TagCloudResourcesResponseBodyFailedResourcesTags) *TagCloudResourcesResponseBodyFailedResources {
	s.Tags = v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) Validate() error {
	return dara.Validate(s)
}

type TagCloudResourcesResponseBodyFailedResourcesTags struct {
	// example:
	//
	// System/Scheduler/STOP_NEW_USER_CONNECTION
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// System
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagCloudResourcesResponseBodyFailedResourcesTags) String() string {
	return dara.Prettify(s)
}

func (s TagCloudResourcesResponseBodyFailedResourcesTags) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) GetKey() *string {
	return s.Key
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) GetScope() *string {
	return s.Scope
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) GetValue() *string {
	return s.Value
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) SetKey(v string) *TagCloudResourcesResponseBodyFailedResourcesTags {
	s.Key = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) SetScope(v string) *TagCloudResourcesResponseBodyFailedResourcesTags {
	s.Scope = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) SetValue(v string) *TagCloudResourcesResponseBodyFailedResourcesTags {
	s.Value = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) Validate() error {
	return dara.Validate(s)
}
