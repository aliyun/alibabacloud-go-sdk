// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *TagResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagResourcesRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *TagResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TagResourcesRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest
	GetTag() []*TagResourcesRequestTag
}

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Set the value to **cluster**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
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

func (s *TagResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TagResourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTag() []*TagResourcesRequestTag {
	return s.Tag
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
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
	return dara.Validate(s)
}

type TagResourcesRequestTag struct {
	// The key of the tag that you want to create for the cluster. To create multiple tags for a cluster at a time, click **Add*	- to add tag keys.
	//
	// >  You can create up to 20 tags for a cluster at a time. The value of `Tag.N.Key` is paired with the value of `Tag.N.Value`.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to create for the cluster. To create multiple tags for a cluster at a time, click **Add*	- to add tag values.
	//
	// >  You can create up to 20 tags for a cluster at a time. The value of `Tag.N.Key` is paired with the value of `Tag.N.Value`.
	//
	// example:
	//
	// test
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
