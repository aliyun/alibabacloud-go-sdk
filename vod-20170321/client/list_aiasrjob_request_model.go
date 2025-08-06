// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIASRJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIASRJobIds(v string) *ListAIASRJobRequest
	GetAIASRJobIds() *string
	SetOwnerAccount(v string) *ListAIASRJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIASRJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIASRJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIASRJobRequest
	GetResourceOwnerId() *string
}

type ListAIASRJobRequest struct {
	// This parameter is required.
	AIASRJobIds          *string `json:"AIASRJobIds,omitempty" xml:"AIASRJobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIASRJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIASRJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIASRJobRequest) GetAIASRJobIds() *string {
	return s.AIASRJobIds
}

func (s *ListAIASRJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIASRJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIASRJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIASRJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIASRJobRequest) SetAIASRJobIds(v string) *ListAIASRJobRequest {
	s.AIASRJobIds = &v
	return s
}

func (s *ListAIASRJobRequest) SetOwnerAccount(v string) *ListAIASRJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIASRJobRequest) SetOwnerId(v string) *ListAIASRJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIASRJobRequest) SetResourceOwnerAccount(v string) *ListAIASRJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIASRJobRequest) SetResourceOwnerId(v string) *ListAIASRJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIASRJobRequest) Validate() error {
	return dara.Validate(s)
}
