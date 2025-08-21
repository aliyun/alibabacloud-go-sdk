// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *DeleteTagResourcesRequest
	GetAll() *bool
	SetRegionId(v string) *DeleteTagResourcesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteTagResourcesRequest
	GetResourceGroupId() *string
	SetResourceIds(v []*string) *DeleteTagResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *DeleteTagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *DeleteTagResourcesRequest
	GetTagKey() []*string
}

type DeleteTagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified resource. Valid values:
	//
	// 	- **true**: yes.
	//
	// 	- **false*	- no. This is the default value.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID of the instance. Set the value to **cn-hangzhou**, which indicates an Anti-DDoS Proxy (Chinese Mainland) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// An array consisting of the IDs of instances from which you want to remove tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the resource to which the tag belongs. Set the value to **INSTANCE**, which indicates an Anti-DDoS Pro instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// An array consisting of the keys of the tags that you want to remove.
	//
	// example:
	//
	// testkey
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s DeleteTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *DeleteTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTagResourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteTagResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DeleteTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteTagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *DeleteTagResourcesRequest) SetAll(v bool) *DeleteTagResourcesRequest {
	s.All = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetRegionId(v string) *DeleteTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetResourceGroupId(v string) *DeleteTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetResourceIds(v []*string) *DeleteTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *DeleteTagResourcesRequest) SetResourceType(v string) *DeleteTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetTagKey(v []*string) *DeleteTagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *DeleteTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
