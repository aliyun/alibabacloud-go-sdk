// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesRequest
	GetAll() *bool
	SetRegionId(v string) *UntagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UntagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all custom tags from the resources. **This parameter takes effect only when `TagKey.N` is not specified**. If `TagKey.N` is specified, this parameter is ignored. Valid values:
	//
	// - `true`: removes all custom tags from the resources, including Wuying system tags that start with `System/` and were bound by calling the [TagResources](~~TagResources~~) operation.
	//
	// - `false` (default): does not remove all tags. If `TagKey.N` is also not specified, the operation returns the error code `InvalidParameter.TagKeyListOrAll`.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID. This parameter is required. Specify the ID of the region where the app instance group resides, such as `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs from which you want to remove tags. This parameter is required. Specify the app instance group IDs. A maximum of 50 IDs can be specified in a single request. Duplicate IDs are automatically deduplicated.
	//
	// **All IDs must belong to existing app instance groups within the current Alibaba Cloud account.*	- If any ID does not exist or does not belong to the current account, the entire request fails and the operation returns the error code `InvalidAppInstanceGroup.NotFound`. No resources are untagged.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. This parameter is required. **Currently, only app instance groups are supported**. The value is case-insensitive. We recommend that you use uppercase letters:
	//
	// - `APPINSTANCEGROUP`: Wuying app instance group.
	//
	// If you specify other values, the operation returns the error code `InvalidResourceType.Invalid`.
	//
	// This parameter is required.
	//
	// example:
	//
	// APPINSTANCEGROUP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tag keys to be removed. A maximum of 20 tag keys can be specified in a single request. **You must specify at least one of `TagKey.N` and `All`**. If neither is specified, the operation returns the error code `InvalidParameter.TagKeyListOrAll`.
	//
	// - If `TagKey.N` is specified, only the tags corresponding to these tag keys are removed, and the `All` parameter is ignored.
	//
	// - If a specified tag key does not exist on a resource, the tag key is skipped and no error is returned.
	//
	// - If `TagKey.N` is not specified, you must set `All` to `true` to remove all custom tags from the resources.
	//
	// Tag keys that start with `System/` are Wuying system tags. Only the following values are supported:
	//
	// - `System/Scheduler/GRAYSCALE`: The grayscale tag for the app instance group.
	//
	// - `System/Scheduler/STOP_NEW_USER_CONNECTION`: Prevents newly bound users from establishing connections to the app instance group.
	//
	// If you specify other tag keys that start with `System/`, the operation returns the error code `InvalidTagPolicy.KeyInvalid` or `InvalidTag.SystemKeyNotAllow`.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
