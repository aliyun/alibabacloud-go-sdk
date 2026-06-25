// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteAll(v bool) *UntagResourcesRequest
	GetDeleteAll() *bool
	SetRegionId(v string) *UntagResourcesRequest
	GetRegionId() *string
	SetResourceIds(v string) *UntagResourcesRequest
	GetResourceIds() *string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKeys(v string) *UntagResourcesRequest
	GetTagKeys() *string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags. This parameter applies only when TagKeys is not specified. Valid values:
	//
	// - **true**: All tags are removed.
	//
	// - **false**: Only the tags specified in the TagKeys parameter are removed.
	//
	// example:
	//
	// false
	DeleteAll *bool `json:"DeleteAll,omitempty" xml:"DeleteAll,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. To specify multiple resource IDs, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["d42921c4-5433-4abd-8075-0e536f8b****"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The resource type. Only `application` is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys. To specify multiple tag keys, separate them with commas (,).
	//
	// example:
	//
	// ["k1","k2"]
	TagKeys *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetDeleteAll() *bool {
	return s.DeleteAll
}

func (s *UntagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKeys() *string {
	return s.TagKeys
}

func (s *UntagResourcesRequest) SetDeleteAll(v bool) *UntagResourcesRequest {
	s.DeleteAll = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v string) *UntagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v string) *UntagResourcesRequest {
	s.TagKeys = &v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
