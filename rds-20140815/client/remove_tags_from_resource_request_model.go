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
	Tag         []*RemoveTagsFromResourceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ClientToken *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ProxyId              *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
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
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
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
