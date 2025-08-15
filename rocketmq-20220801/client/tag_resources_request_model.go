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
	SetResourceId(v string) *TagResourcesRequest
	GetResourceId() *string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTag(v string) *TagResourcesRequest
	GetTag() *string
}

type TagResourcesRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource IDs, in the JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq-cn-pe3355cs707
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of resource.
	//
	// Set this parameter to **instance**. The value of this parameter cannot be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// tag, in JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"key": "rmq-test", "value": "test"}]
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
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

func (s *TagResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTag() *string {
	return s.Tag
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v string) *TagResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v string) *TagResourcesRequest {
	s.Tag = &v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
