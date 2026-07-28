// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesForExpressConnectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesForExpressConnectRequest
	GetAll() *bool
	SetOwnerAccount(v string) *UntagResourcesForExpressConnectRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UntagResourcesForExpressConnectRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UntagResourcesForExpressConnectRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UntagResourcesForExpressConnectRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *UntagResourcesForExpressConnectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UntagResourcesForExpressConnectRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *UntagResourcesForExpressConnectRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesForExpressConnectRequest
	GetTagKey() []*string
}

type UntagResourcesForExpressConnectRequest struct {
	// Specifies whether to unbind all tags from the resource. Valid values:
	//
	// - **true**: Unbinds all tags from the resource.
	//
	// - **false*	- (default): Does not unbind all tags from the resource.
	//
	// example:
	//
	// false
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - **PHYSICALCONNECTION**: Express Connect circuit instance.
	//
	// - **VIRTUALBORDERROUTER**: Virtual Border Router.
	//
	// - **ROUTERINTERFACE**: VBR uplink.
	//
	// - **TRAFFICQOS**: QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tag keys for the resource.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesForExpressConnectRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesForExpressConnectRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesForExpressConnectRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesForExpressConnectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UntagResourcesForExpressConnectRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UntagResourcesForExpressConnectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesForExpressConnectRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagResourcesForExpressConnectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UntagResourcesForExpressConnectRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UntagResourcesForExpressConnectRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesForExpressConnectRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UntagResourcesForExpressConnectRequest) SetAll(v bool) *UntagResourcesForExpressConnectRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) SetOwnerAccount(v string) *UntagResourcesForExpressConnectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) SetOwnerId(v int64) *UntagResourcesForExpressConnectRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) SetRegionId(v string) *UntagResourcesForExpressConnectRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) SetResourceId(v []*string) *UntagResourcesForExpressConnectRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) SetResourceOwnerAccount(v string) *UntagResourcesForExpressConnectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) SetResourceOwnerId(v int64) *UntagResourcesForExpressConnectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) SetResourceType(v string) *UntagResourcesForExpressConnectRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) SetTagKey(v []*string) *UntagResourcesForExpressConnectRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesForExpressConnectRequest) Validate() error {
	return dara.Validate(s)
}
