// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *TagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TagResourcesRequest
	GetRegionId() *string
	SetResourceIds(v []*string) *TagResourcesRequest
	GetResourceIds() []*string
	SetResourceOwnerAccount(v string) *TagResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest
	GetTags() []*TagResourcesRequestTags
}

type TagResourcesRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the Auto Scaling resources. You can specify 1 to 50 resource IDs.
	//
	// This parameter is required.
	ResourceIds          []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. Set the value to scalinggroup.
	//
	// This parameter is required.
	//
	// example:
	//
	// scalinggroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to add to the Auto Scaling resources.
	//
	// This parameter is required.
	Tags []*TagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *TagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTags() []*TagResourcesRequestTags {
	return s.Tags
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v []*string) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest {
	s.Tags = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type TagResourcesRequestTags struct {
	// The tag key.
	//
	// You cannot specify an empty string as a tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specifies whether to propagate the tag that you want to add. Valid values:
	//
	// 	- true: propagates the tag to new instances.
	//
	// 	- false: does not propagate the tag to any instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Propagate *bool `json:"Propagate,omitempty" xml:"Propagate,omitempty"`
	// The tag value.
	//
	// You can specify empty strings as tag values. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTags) GetKey() *string {
	return s.Key
}

func (s *TagResourcesRequestTags) GetPropagate() *bool {
	return s.Propagate
}

func (s *TagResourcesRequestTags) GetValue() *string {
	return s.Value
}

func (s *TagResourcesRequestTags) SetKey(v string) *TagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTags) SetPropagate(v bool) *TagResourcesRequestTags {
	s.Propagate = &v
	return s
}

func (s *TagResourcesRequestTags) SetValue(v string) *TagResourcesRequestTags {
	s.Value = &v
	return s
}

func (s *TagResourcesRequestTags) Validate() error {
	return dara.Validate(s)
}
