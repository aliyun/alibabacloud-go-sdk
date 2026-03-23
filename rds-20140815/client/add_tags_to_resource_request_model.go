// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsToResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*AddTagsToResourceRequestTag) *AddTagsToResourceRequest
	GetTag() []*AddTagsToResourceRequestTag
	SetClientToken(v string) *AddTagsToResourceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *AddTagsToResourceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *AddTagsToResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddTagsToResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddTagsToResourceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AddTagsToResourceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AddTagsToResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTagsToResourceRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *AddTagsToResourceRequest
	GetTags() *string
	SetProxyId(v string) *AddTagsToResourceRequest
	GetProxyId() *string
}

type AddTagsToResourceRequest struct {
	Tag         []*AddTagsToResourceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ClientToken *string                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s AddTagsToResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTagsToResourceRequest) GoString() string {
	return s.String()
}

func (s *AddTagsToResourceRequest) GetTag() []*AddTagsToResourceRequestTag {
	return s.Tag
}

func (s *AddTagsToResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddTagsToResourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AddTagsToResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddTagsToResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTagsToResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddTagsToResourceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddTagsToResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTagsToResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTagsToResourceRequest) GetTags() *string {
	return s.Tags
}

func (s *AddTagsToResourceRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *AddTagsToResourceRequest) SetTag(v []*AddTagsToResourceRequestTag) *AddTagsToResourceRequest {
	s.Tag = v
	return s
}

func (s *AddTagsToResourceRequest) SetClientToken(v string) *AddTagsToResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddTagsToResourceRequest) SetDBInstanceId(v string) *AddTagsToResourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetOwnerAccount(v string) *AddTagsToResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddTagsToResourceRequest) SetOwnerId(v int64) *AddTagsToResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetRegionId(v string) *AddTagsToResourceRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetResourceGroupId(v string) *AddTagsToResourceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetResourceOwnerAccount(v string) *AddTagsToResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTagsToResourceRequest) SetResourceOwnerId(v int64) *AddTagsToResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTagsToResourceRequest) SetTags(v string) *AddTagsToResourceRequest {
	s.Tags = &v
	return s
}

func (s *AddTagsToResourceRequest) SetProxyId(v string) *AddTagsToResourceRequest {
	s.ProxyId = &v
	return s
}

func (s *AddTagsToResourceRequest) Validate() error {
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

type AddTagsToResourceRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddTagsToResourceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddTagsToResourceRequestTag) GoString() string {
	return s.String()
}

func (s *AddTagsToResourceRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddTagsToResourceRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddTagsToResourceRequestTag) SetKey(v string) *AddTagsToResourceRequestTag {
	s.Key = &v
	return s
}

func (s *AddTagsToResourceRequestTag) SetValue(v string) *AddTagsToResourceRequestTag {
	s.Value = &v
	return s
}

func (s *AddTagsToResourceRequestTag) Validate() error {
	return dara.Validate(s)
}
