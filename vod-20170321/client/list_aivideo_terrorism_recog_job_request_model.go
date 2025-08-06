// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoTerrorismRecogJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoTerrorismRecogJobIds(v string) *ListAIVideoTerrorismRecogJobRequest
	GetAIVideoTerrorismRecogJobIds() *string
	SetOwnerAccount(v string) *ListAIVideoTerrorismRecogJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIVideoTerrorismRecogJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIVideoTerrorismRecogJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIVideoTerrorismRecogJobRequest
	GetResourceOwnerId() *string
}

type ListAIVideoTerrorismRecogJobRequest struct {
	// This parameter is required.
	AIVideoTerrorismRecogJobIds *string `json:"AIVideoTerrorismRecogJobIds,omitempty" xml:"AIVideoTerrorismRecogJobIds,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIVideoTerrorismRecogJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTerrorismRecogJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIVideoTerrorismRecogJobRequest) GetAIVideoTerrorismRecogJobIds() *string {
	return s.AIVideoTerrorismRecogJobIds
}

func (s *ListAIVideoTerrorismRecogJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIVideoTerrorismRecogJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIVideoTerrorismRecogJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIVideoTerrorismRecogJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIVideoTerrorismRecogJobRequest) SetAIVideoTerrorismRecogJobIds(v string) *ListAIVideoTerrorismRecogJobRequest {
	s.AIVideoTerrorismRecogJobIds = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobRequest) SetOwnerAccount(v string) *ListAIVideoTerrorismRecogJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobRequest) SetOwnerId(v string) *ListAIVideoTerrorismRecogJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobRequest) SetResourceOwnerAccount(v string) *ListAIVideoTerrorismRecogJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobRequest) SetResourceOwnerId(v string) *ListAIVideoTerrorismRecogJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobRequest) Validate() error {
	return dara.Validate(s)
}
