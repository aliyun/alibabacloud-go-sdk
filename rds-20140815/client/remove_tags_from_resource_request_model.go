// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsFromResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*RemoveTagsFromResourceRequestTag) *RemoveTagsFromResourceRequest
	GetTag() []*RemoveTagsFromResourceRequestTag
	SetClientToken(v string) *RemoveTagsFromResourceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RemoveTagsFromResourceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *RemoveTagsFromResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveTagsFromResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveTagsFromResourceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RemoveTagsFromResourceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RemoveTagsFromResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveTagsFromResourceRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *RemoveTagsFromResourceRequest
	GetTags() *string
	SetProxyId(v string) *RemoveTagsFromResourceRequest
	GetProxyId() *string
}

type RemoveTagsFromResourceRequest struct {
	Tag []*RemoveTagsFromResourceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the ListResourceGroups operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A set of a TagKey and a TagValue that you use to unbind the tag. Format: {"key1":"value1"}.
	//
	// >  You cannot specify an empty string for TagKey. You can specify an empty string for TagValue.
	//
	// example:
	//
	// {"key1":"value1"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the proxy mode.
	//
	// example:
	//
	// API
	ProxyId *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s RemoveTagsFromResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsFromResourceRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsFromResourceRequest) GetTag() []*RemoveTagsFromResourceRequestTag {
	return s.Tag
}

func (s *RemoveTagsFromResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveTagsFromResourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RemoveTagsFromResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveTagsFromResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveTagsFromResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveTagsFromResourceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RemoveTagsFromResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveTagsFromResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveTagsFromResourceRequest) GetTags() *string {
	return s.Tags
}

func (s *RemoveTagsFromResourceRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *RemoveTagsFromResourceRequest) SetTag(v []*RemoveTagsFromResourceRequestTag) *RemoveTagsFromResourceRequest {
	s.Tag = v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetClientToken(v string) *RemoveTagsFromResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetDBInstanceId(v string) *RemoveTagsFromResourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetOwnerAccount(v string) *RemoveTagsFromResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetOwnerId(v int64) *RemoveTagsFromResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetRegionId(v string) *RemoveTagsFromResourceRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetResourceGroupId(v string) *RemoveTagsFromResourceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetResourceOwnerAccount(v string) *RemoveTagsFromResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetResourceOwnerId(v int64) *RemoveTagsFromResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetTags(v string) *RemoveTagsFromResourceRequest {
	s.Tags = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) SetProxyId(v string) *RemoveTagsFromResourceRequest {
	s.ProxyId = &v
	return s
}

func (s *RemoveTagsFromResourceRequest) Validate() error {
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

type RemoveTagsFromResourceRequestTag struct {
	// The TagKey of the first tag that you want to unbind. Each tag consists of a TagKey and a TagValue. You can specify up to five tags in a single request. You cannot specify an empty string as the tag key. You can specify an empty string as the tag value.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The TagValue of the first tag that you want to unbind. Each tag consists of a TagKey and a TagValue. You can specify up to five tags in a single request. You cannot specify an empty string as the tag key. You can specify an empty string as the tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RemoveTagsFromResourceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsFromResourceRequestTag) GoString() string {
	return s.String()
}

func (s *RemoveTagsFromResourceRequestTag) GetKey() *string {
	return s.Key
}

func (s *RemoveTagsFromResourceRequestTag) GetValue() *string {
	return s.Value
}

func (s *RemoveTagsFromResourceRequestTag) SetKey(v string) *RemoveTagsFromResourceRequestTag {
	s.Key = &v
	return s
}

func (s *RemoveTagsFromResourceRequestTag) SetValue(v string) *RemoveTagsFromResourceRequestTag {
	s.Value = &v
	return s
}

func (s *RemoveTagsFromResourceRequestTag) Validate() error {
	return dara.Validate(s)
}
