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
	SetOwnerAccount(v string) *UntagResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UntagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UntagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UntagResourcesRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *UntagResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UntagResourcesRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified resource. This parameter takes effect only if **TagKey.N*	- is empty.
	//
	// Valid values: **true*	- and **false**.
	//
	// example:
	//
	// true
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource. You can specify at most 20 resources.
	//
	// >  Set **ResourceId*	- of the **listener*	- to **LoadBalancerId_Listener protocol_Port**. Set LoadBalancerId to the ID of the CLB instance, and Port to the listener port.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp16qjewdsunr41m1****
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of resource. Valid values:
	//
	// 	- **instance**: CLB instance
	//
	// 	- **certificate**: certificate
	//
	// 	- **acl**: access control list (ACL)
	//
	// 	- **listener**: listener
	//
	// 	- **vservergroup**: vServer group
	//
	// 	- **masterslaveservergroup**: primary/secondary server group
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// >  If you do not specify **TagKey**, all tags are removed from the instance.
	//
	// example:
	//
	// FinanceDept
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

func (s *UntagResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UntagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UntagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UntagResourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
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

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
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
