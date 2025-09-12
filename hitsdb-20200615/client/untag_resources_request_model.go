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
	SetResourceId(v []*string) *UntagResourcesRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *UntagResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UntagResourcesRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetSecurityToken(v string) *UntagResourcesRequest
	GetSecurityToken() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the instances. Valid values:
	//
	// 	- **true**: Remove all tags from the instances.
	//
	// 	- **false**: Do not remove all tags from the instances.
	//
	// >  The default value of this parameter is false.
	//
	//
	//
	// 	- If you specify this parameter together with the TagKey parameter, this parameter does not take effect.
	//
	// example:
	//
	// false
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of instances.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Set the value to **INSTANCE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The list of keys of the tags that you want to remove.
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

func (s *UntagResourcesRequest) GetSecurityToken() *string {
	return s.SecurityToken
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

func (s *UntagResourcesRequest) SetSecurityToken(v string) *UntagResourcesRequest {
	s.SecurityToken = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
