// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCategoryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCategoryJobIds(v string) *ListAIVideoCategoryJobRequest
	GetAIVideoCategoryJobIds() *string
	SetOwnerAccount(v string) *ListAIVideoCategoryJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIVideoCategoryJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIVideoCategoryJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIVideoCategoryJobRequest
	GetResourceOwnerId() *string
}

type ListAIVideoCategoryJobRequest struct {
	// This parameter is required.
	AIVideoCategoryJobIds *string `json:"AIVideoCategoryJobIds,omitempty" xml:"AIVideoCategoryJobIds,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIVideoCategoryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCategoryJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIVideoCategoryJobRequest) GetAIVideoCategoryJobIds() *string {
	return s.AIVideoCategoryJobIds
}

func (s *ListAIVideoCategoryJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIVideoCategoryJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIVideoCategoryJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIVideoCategoryJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIVideoCategoryJobRequest) SetAIVideoCategoryJobIds(v string) *ListAIVideoCategoryJobRequest {
	s.AIVideoCategoryJobIds = &v
	return s
}

func (s *ListAIVideoCategoryJobRequest) SetOwnerAccount(v string) *ListAIVideoCategoryJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIVideoCategoryJobRequest) SetOwnerId(v string) *ListAIVideoCategoryJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIVideoCategoryJobRequest) SetResourceOwnerAccount(v string) *ListAIVideoCategoryJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIVideoCategoryJobRequest) SetResourceOwnerId(v string) *ListAIVideoCategoryJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIVideoCategoryJobRequest) Validate() error {
	return dara.Validate(s)
}
