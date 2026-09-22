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
	// Specifies whether to unbind all custom tags from the resource. **This parameter takes effect only when `TagKey.N` is not specified.*	- If `TagKey.N` is specified, this parameter is ignored. Valid values:
	//
	// - `true`: Unbinds all custom tags from the resource, including Wuying system tags that start with `System/` and were bound by calling [TagResources](~~TagResources~~).
	//
	// - `false` (default): Does not perform a full unbinding. If `TagKey.N` is also not specified, the error code `InvalidParameter.TagKeyListOrAll` is returned.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID. This parameter is required. Set this parameter to the ID of the region where the delivery group resides, such as `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs from which you want to unbind tags. This parameter is required. Specify delivery group IDs. You can specify up to 50 IDs at a time. Duplicate IDs are automatically deduplicated.
	//
	// **All IDs must correspond to existing delivery groups under the current Alibaba Cloud account.*	- If any ID does not exist or does not belong to the current account, the entire request fails with the error code `InvalidAppInstanceGroup.NotFound`, and no tags are unbound from any resource.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. This parameter is required. **Currently, only delivery groups are supported.*	- The value is case-insensitive. We recommend that you use uppercase.
	//
	// Valid values:
	//
	// - `APPINSTANCEGROUP`: Wuying delivery group.
	//
	// If you specify other values, the error code `InvalidResourceType.Invalid` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// APPINSTANCEGROUP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tag keys to unbind. You can specify up to 20 tag keys at a time. **Specify at least one of `TagKey.N` and `All`.*	- If neither is specified, the error code `InvalidParameter.TagKeyListOrAll` is returned.
	//
	// - If `TagKey.N` is specified, only the tags that correspond to the specified tag keys are unbound. The `All` parameter is ignored.
	//
	// - If a specified tag key does not exist on the resource, the tag key is skipped and no error is returned.
	//
	// - If `TagKey.N` is not specified, set `All` to `true` to unbind all custom tags from the resource.
	//
	// Tag keys that start with `System/` are Wuying system tags. Only the following values are supported:
	//
	// - `System/Scheduler/GRAYSCALE`: the canary release tag for the delivery group.
	//
	// - `System/Scheduler/STOP_NEW_USER_CONNECTION`: prevents newly bound users from establishing connections to the delivery group.
	//
	// If you specify other tag keys that start with `System/`, the error code `InvalidTagPolicy.KeyInvalid` or `InvalidTag.SystemKeyNotAllow` is returned.
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
