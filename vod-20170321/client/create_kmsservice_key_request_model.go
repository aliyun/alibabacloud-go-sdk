// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKMSServiceKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKmsRegionId(v string) *CreateKMSServiceKeyRequest
	GetKmsRegionId() *string
	SetOwnerAccount(v string) *CreateKMSServiceKeyRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *CreateKMSServiceKeyRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *CreateKMSServiceKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *CreateKMSServiceKeyRequest
	GetResourceOwnerId() *string
}

type CreateKMSServiceKeyRequest struct {
	KmsRegionId          *string `json:"KmsRegionId,omitempty" xml:"KmsRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateKMSServiceKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKMSServiceKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateKMSServiceKeyRequest) GetKmsRegionId() *string {
	return s.KmsRegionId
}

func (s *CreateKMSServiceKeyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateKMSServiceKeyRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateKMSServiceKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateKMSServiceKeyRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *CreateKMSServiceKeyRequest) SetKmsRegionId(v string) *CreateKMSServiceKeyRequest {
	s.KmsRegionId = &v
	return s
}

func (s *CreateKMSServiceKeyRequest) SetOwnerAccount(v string) *CreateKMSServiceKeyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateKMSServiceKeyRequest) SetOwnerId(v string) *CreateKMSServiceKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateKMSServiceKeyRequest) SetResourceOwnerAccount(v string) *CreateKMSServiceKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateKMSServiceKeyRequest) SetResourceOwnerId(v string) *CreateKMSServiceKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateKMSServiceKeyRequest) Validate() error {
	return dara.Validate(s)
}
