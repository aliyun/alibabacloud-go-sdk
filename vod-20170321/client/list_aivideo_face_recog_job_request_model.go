// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoFaceRecogJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoFaceRecogJobIds(v string) *ListAIVideoFaceRecogJobRequest
	GetAIVideoFaceRecogJobIds() *string
	SetOwnerAccount(v string) *ListAIVideoFaceRecogJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListAIVideoFaceRecogJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListAIVideoFaceRecogJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListAIVideoFaceRecogJobRequest
	GetResourceOwnerId() *string
}

type ListAIVideoFaceRecogJobRequest struct {
	// This parameter is required.
	AIVideoFaceRecogJobIds *string `json:"AIVideoFaceRecogJobIds,omitempty" xml:"AIVideoFaceRecogJobIds,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIVideoFaceRecogJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoFaceRecogJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIVideoFaceRecogJobRequest) GetAIVideoFaceRecogJobIds() *string {
	return s.AIVideoFaceRecogJobIds
}

func (s *ListAIVideoFaceRecogJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAIVideoFaceRecogJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListAIVideoFaceRecogJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAIVideoFaceRecogJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListAIVideoFaceRecogJobRequest) SetAIVideoFaceRecogJobIds(v string) *ListAIVideoFaceRecogJobRequest {
	s.AIVideoFaceRecogJobIds = &v
	return s
}

func (s *ListAIVideoFaceRecogJobRequest) SetOwnerAccount(v string) *ListAIVideoFaceRecogJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIVideoFaceRecogJobRequest) SetOwnerId(v string) *ListAIVideoFaceRecogJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIVideoFaceRecogJobRequest) SetResourceOwnerAccount(v string) *ListAIVideoFaceRecogJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIVideoFaceRecogJobRequest) SetResourceOwnerId(v string) *ListAIVideoFaceRecogJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAIVideoFaceRecogJobRequest) Validate() error {
	return dara.Validate(s)
}
