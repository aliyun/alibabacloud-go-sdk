// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCensorJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCensorJobIds(v string) *ListAIVideoCensorJobRequest
	GetAIVideoCensorJobIds() *string
	SetOwnerAccount(v string) *ListAIVideoCensorJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIVideoCensorJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIVideoCensorJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIVideoCensorJobRequest
	GetResourceOwnerId() *string
}

type ListAIVideoCensorJobRequest struct {
	// This parameter is required.
	AIVideoCensorJobIds  *string `json:"AIVideoCensorJobIds,omitempty" xml:"AIVideoCensorJobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIVideoCensorJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCensorJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIVideoCensorJobRequest) GetAIVideoCensorJobIds() *string {
	return s.AIVideoCensorJobIds
}

func (s *ListAIVideoCensorJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIVideoCensorJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIVideoCensorJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIVideoCensorJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIVideoCensorJobRequest) SetAIVideoCensorJobIds(v string) *ListAIVideoCensorJobRequest {
	s.AIVideoCensorJobIds = &v
	return s
}

func (s *ListAIVideoCensorJobRequest) SetOwnerAccount(v string) *ListAIVideoCensorJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIVideoCensorJobRequest) SetOwnerId(v string) *ListAIVideoCensorJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIVideoCensorJobRequest) SetResourceOwnerAccount(v string) *ListAIVideoCensorJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIVideoCensorJobRequest) SetResourceOwnerId(v string) *ListAIVideoCensorJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIVideoCensorJobRequest) Validate() error {
	return dara.Validate(s)
}
