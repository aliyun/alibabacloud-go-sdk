// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *AddTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddTagsRequest
	GetRegionId() *string
	SetResourceId(v string) *AddTagsRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *AddTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTagsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *AddTagsRequest
	GetResourceType() *string
	SetTag(v []*AddTagsRequestTag) *AddTagsRequest
	GetTag() []*AddTagsRequestTag
}

type AddTagsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*AddTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AddTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTagsRequest) GoString() string {
	return s.String()
}

func (s *AddTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddTagsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddTagsRequest) GetTag() []*AddTagsRequestTag {
	return s.Tag
}

func (s *AddTagsRequest) SetOwnerId(v int64) *AddTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTagsRequest) SetRegionId(v string) *AddTagsRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsRequest) SetResourceId(v string) *AddTagsRequest {
	s.ResourceId = &v
	return s
}

func (s *AddTagsRequest) SetResourceOwnerAccount(v string) *AddTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTagsRequest) SetResourceOwnerId(v int64) *AddTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTagsRequest) SetResourceType(v string) *AddTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *AddTagsRequest) SetTag(v []*AddTagsRequestTag) *AddTagsRequest {
	s.Tag = v
	return s
}

func (s *AddTagsRequest) Validate() error {
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

type AddTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddTagsRequestTag) GoString() string {
	return s.String()
}

func (s *AddTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddTagsRequestTag) SetKey(v string) *AddTagsRequestTag {
	s.Key = &v
	return s
}

func (s *AddTagsRequestTag) SetValue(v string) *AddTagsRequestTag {
	s.Value = &v
	return s
}

func (s *AddTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
