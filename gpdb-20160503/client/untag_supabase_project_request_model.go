// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagSupabaseProjectRequest
	GetAll() *bool
	SetRegionId(v string) *UntagSupabaseProjectRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UntagSupabaseProjectRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagSupabaseProjectRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagSupabaseProjectRequest
	GetTagKey() []*string
}

type UntagSupabaseProjectRequest struct {
	// Specifies whether to remove all tags from the instance. This parameter takes effect only when `TagKey.N` is not specified. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance IDs. You can specify up to 50 instance IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// - `instance`: a reserved mode instance.
	//
	// - `ALIYUN::GPDB::INSTANCE`: an elastic mode instance.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags that you want to remove. You can specify up to 20 tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *UntagSupabaseProjectRequest) GetAll() *bool {
	return s.All
}

func (s *UntagSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagSupabaseProjectRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagSupabaseProjectRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagSupabaseProjectRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UntagSupabaseProjectRequest) SetAll(v bool) *UntagSupabaseProjectRequest {
	s.All = &v
	return s
}

func (s *UntagSupabaseProjectRequest) SetRegionId(v string) *UntagSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *UntagSupabaseProjectRequest) SetResourceId(v []*string) *UntagSupabaseProjectRequest {
	s.ResourceId = v
	return s
}

func (s *UntagSupabaseProjectRequest) SetResourceType(v string) *UntagSupabaseProjectRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagSupabaseProjectRequest) SetTagKey(v []*string) *UntagSupabaseProjectRequest {
	s.TagKey = v
	return s
}

func (s *UntagSupabaseProjectRequest) Validate() error {
	return dara.Validate(s)
}
