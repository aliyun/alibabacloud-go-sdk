// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest
	GetTag() []*TagResourcesRequestTag
}

type TagResourcesRequest struct {
	// The region ID. This parameter is required. Specify the ID of the region where the app instance group resides, such as `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs to which you want to bind tags. This parameter is required. Specify the app instance group IDs. You can specify up to 50 IDs in a single request. Duplicate IDs are automatically deduplicated.
	//
	// **All IDs must belong to existing app instance groups under the current Alibaba Cloud account.*	- If any ID does not exist or does not belong to the current account, the entire request fails and the error code `InvalidAppInstanceGroup.NotFound` is returned. No tags are bound to any resources.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. This parameter is required. **Currently, only app instance groups are supported**. The value is case-insensitive, but we recommend that you use uppercase letters:
	//
	// - `APPINSTANCEGROUP`: app instance group.
	//
	// If you specify other values, the error code `InvalidResourceType.Invalid` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// APPINSTANCEGROUP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags to bind. This parameter is required. You can specify up to 20 tags in a single request. You must specify both Key and Value for each tag.
	//
	// - Tag keys in the same request cannot be duplicate. Otherwise, the error code `InvalidTag.Duplicated` is returned.
	//
	// - If a tag with the same key already exists on the resource, the tag value is updated to the value specified in the current request.
	//
	// - A single resource can be bound to up to 20 custom tags. If this limit is exceeded, the error code `ResourceTag.CustomTagCountExceed` is returned.
	//
	// Tag keys starting with `System/` are Wuying system tags. Only the following values are supported, and the tag value can only be `true` or `false`:
	//
	// - `System/Scheduler/GRAYSCALE`: the canary release tag for the app instance group.
	//
	// - `System/Scheduler/STOP_NEW_USER_CONNECTION`: prevents newly bound users from establishing connections to the app instance group.
	//
	// If you specify other tag keys starting with `System/`, the error code `InvalidTag.SystemTagKeyInvalid` or `InvalidTag.SystemKeyNotAllow` is returned.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTag() []*TagResourcesRequestTag {
	return s.Tag
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
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

type TagResourcesRequestTag struct {
	// The tag key. This parameter is required. The tag key must be 1 to 128 characters in length and is case-sensitive. It cannot start with `aliyun` or `acs:` (case-insensitive) and cannot contain `http://` or `https://`. It supports letters, digits, spaces, and common Chinese and English punctuation marks. If the tag key does not meet the rules, the error code `InvalidTagPolicy.KeyInvalid` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// Resolution
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. This parameter is required. The tag value must be 0 to 256 characters in length and is case-sensitive. Empty strings are allowed. It cannot contain `http://` or `https://`. If the tag value does not meet the rules, the error code `InvalidTagPolicy.ValueInvalid` is returned.
	//
	// If the tag key is a system tag, the tag value can only be `true` or `false`.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 720p
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *TagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
