// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCoverJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCoverJobIds(v string) *ListAIVideoCoverJobRequest
	GetAIVideoCoverJobIds() *string
	SetOwnerAccount(v string) *ListAIVideoCoverJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIVideoCoverJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIVideoCoverJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIVideoCoverJobRequest
	GetResourceOwnerId() *string
}

type ListAIVideoCoverJobRequest struct {
	// This parameter is required.
	AIVideoCoverJobIds   *string `json:"AIVideoCoverJobIds,omitempty" xml:"AIVideoCoverJobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIVideoCoverJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCoverJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIVideoCoverJobRequest) GetAIVideoCoverJobIds() *string {
	return s.AIVideoCoverJobIds
}

func (s *ListAIVideoCoverJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIVideoCoverJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIVideoCoverJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIVideoCoverJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIVideoCoverJobRequest) SetAIVideoCoverJobIds(v string) *ListAIVideoCoverJobRequest {
	s.AIVideoCoverJobIds = &v
	return s
}

func (s *ListAIVideoCoverJobRequest) SetOwnerAccount(v string) *ListAIVideoCoverJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIVideoCoverJobRequest) SetOwnerId(v string) *ListAIVideoCoverJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIVideoCoverJobRequest) SetResourceOwnerAccount(v string) *ListAIVideoCoverJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIVideoCoverJobRequest) SetResourceOwnerId(v string) *ListAIVideoCoverJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIVideoCoverJobRequest) Validate() error {
	return dara.Validate(s)
}
