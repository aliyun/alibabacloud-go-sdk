// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateTagResourcesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTagResourcesRequest
	GetResourceGroupId() *string
	SetResourceIds(v []*string) *CreateTagResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *CreateTagResourcesRequest
	GetResourceType() *string
	SetTags(v []*CreateTagResourcesRequestTags) *CreateTagResourcesRequest
	GetTags() []*CreateTagResourcesRequestTags
}

type CreateTagResourcesRequest struct {
	// The region ID of the Anti-DDoS Proxy instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the Anti-DDoS Proxy instances to which you want to add tags.
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
	// An array that consists of the tags to add.
	Tags []*CreateTagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *CreateTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTagResourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTagResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *CreateTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateTagResourcesRequest) GetTags() []*CreateTagResourcesRequestTags {
	return s.Tags
}

func (s *CreateTagResourcesRequest) SetRegionId(v string) *CreateTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTagResourcesRequest) SetResourceGroupId(v string) *CreateTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTagResourcesRequest) SetResourceIds(v []*string) *CreateTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *CreateTagResourcesRequest) SetResourceType(v string) *CreateTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateTagResourcesRequest) SetTags(v []*CreateTagResourcesRequestTags) *CreateTagResourcesRequest {
	s.Tags = v
	return s
}

func (s *CreateTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTagResourcesRequestTags struct {
	// The key of the tag to add.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagResourcesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTagResourcesRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateTagResourcesRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateTagResourcesRequestTags) SetKey(v string) *CreateTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTagResourcesRequestTags) SetValue(v string) *CreateTagResourcesRequestTags {
	s.Value = &v
	return s
}

func (s *CreateTagResourcesRequestTags) Validate() error {
	return dara.Validate(s)
}
