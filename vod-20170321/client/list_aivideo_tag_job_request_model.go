// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoTagJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoTagJobIds(v string) *ListAIVideoTagJobRequest
	GetAIVideoTagJobIds() *string
	SetOwnerAccount(v string) *ListAIVideoTagJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIVideoTagJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIVideoTagJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIVideoTagJobRequest
	GetResourceOwnerId() *string
}

type ListAIVideoTagJobRequest struct {
	// This parameter is required.
	AIVideoTagJobIds     *string `json:"AIVideoTagJobIds,omitempty" xml:"AIVideoTagJobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIVideoTagJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTagJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIVideoTagJobRequest) GetAIVideoTagJobIds() *string {
	return s.AIVideoTagJobIds
}

func (s *ListAIVideoTagJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIVideoTagJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIVideoTagJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIVideoTagJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIVideoTagJobRequest) SetAIVideoTagJobIds(v string) *ListAIVideoTagJobRequest {
	s.AIVideoTagJobIds = &v
	return s
}

func (s *ListAIVideoTagJobRequest) SetOwnerAccount(v string) *ListAIVideoTagJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIVideoTagJobRequest) SetOwnerId(v string) *ListAIVideoTagJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIVideoTagJobRequest) SetResourceOwnerAccount(v string) *ListAIVideoTagJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIVideoTagJobRequest) SetResourceOwnerId(v string) *ListAIVideoTagJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIVideoTagJobRequest) Validate() error {
	return dara.Validate(s)
}
