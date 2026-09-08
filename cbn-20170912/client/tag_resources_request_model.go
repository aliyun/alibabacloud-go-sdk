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
	// The region ID of the resource instance.
	//
	// This parameter is not required when the resource type is Cen or BandwidthPackage. This parameter is required for all other resource types.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// **Cen**: CEN instance.
	//
	// **BandwidthPackage**: bandwidth package instance.
	//
	// **TransitRouter**: transit router instance.
	//
	// **TransitRouterVpcAttachment**: VPC connection instance.
	//
	// **TransitRouterVbrAttachment**: VBR connection instance.
	//
	// **TransitRouterPeerAttachment**: inter-region connection instance.
	//
	// **TransitRouterVpnAttachment**: VPN connection instance.
	//
	// **TransitRouterRouteTable**: route table instance.
	//
	// **Flowlog**: flow log instance.
	//
	// **TransitRouterMulticastDomain**: multicast domain instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag list.
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TagResourcesRequestTag struct {
	// The tag key.
	//
	// You can specify multiple tag keys. Valid values of **N**: **1*	- to **20**.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Each tag key corresponds to one tag value. Valid values of **N**: **1*	- to **20**.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TagValue
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
