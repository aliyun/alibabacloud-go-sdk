// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoPornRecogJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoPornRecogJobIds(v string) *ListAIVideoPornRecogJobRequest
	GetAIVideoPornRecogJobIds() *string
	SetOwnerAccount(v string) *ListAIVideoPornRecogJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIVideoPornRecogJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIVideoPornRecogJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIVideoPornRecogJobRequest
	GetResourceOwnerId() *string
}

type ListAIVideoPornRecogJobRequest struct {
	// This parameter is required.
	AIVideoPornRecogJobIds *string `json:"AIVideoPornRecogJobIds,omitempty" xml:"AIVideoPornRecogJobIds,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIVideoPornRecogJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoPornRecogJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIVideoPornRecogJobRequest) GetAIVideoPornRecogJobIds() *string {
	return s.AIVideoPornRecogJobIds
}

func (s *ListAIVideoPornRecogJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIVideoPornRecogJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIVideoPornRecogJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIVideoPornRecogJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIVideoPornRecogJobRequest) SetAIVideoPornRecogJobIds(v string) *ListAIVideoPornRecogJobRequest {
	s.AIVideoPornRecogJobIds = &v
	return s
}

func (s *ListAIVideoPornRecogJobRequest) SetOwnerAccount(v string) *ListAIVideoPornRecogJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIVideoPornRecogJobRequest) SetOwnerId(v string) *ListAIVideoPornRecogJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIVideoPornRecogJobRequest) SetResourceOwnerAccount(v string) *ListAIVideoPornRecogJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIVideoPornRecogJobRequest) SetResourceOwnerId(v string) *ListAIVideoPornRecogJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIVideoPornRecogJobRequest) Validate() error {
	return dara.Validate(s)
}
