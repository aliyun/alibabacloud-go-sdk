// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKMSServiceKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKmsRegionId(v string) *GetKMSServiceKeyRequest
	GetKmsRegionId() *string
	SetOwnerAccount(v string) *GetKMSServiceKeyRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetKMSServiceKeyRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetKMSServiceKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetKMSServiceKeyRequest
	GetResourceOwnerId() *string
}

type GetKMSServiceKeyRequest struct {
	// This parameter is required.
	KmsRegionId          *string `json:"KmsRegionId,omitempty" xml:"KmsRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetKMSServiceKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKMSServiceKeyRequest) GoString() string {
	return s.String()
}

func (s *GetKMSServiceKeyRequest) GetKmsRegionId() *string {
	return s.KmsRegionId
}

func (s *GetKMSServiceKeyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetKMSServiceKeyRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetKMSServiceKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetKMSServiceKeyRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetKMSServiceKeyRequest) SetKmsRegionId(v string) *GetKMSServiceKeyRequest {
	s.KmsRegionId = &v
	return s
}

func (s *GetKMSServiceKeyRequest) SetOwnerAccount(v string) *GetKMSServiceKeyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetKMSServiceKeyRequest) SetOwnerId(v string) *GetKMSServiceKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetKMSServiceKeyRequest) SetResourceOwnerAccount(v string) *GetKMSServiceKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetKMSServiceKeyRequest) SetResourceOwnerId(v string) *GetKMSServiceKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetKMSServiceKeyRequest) Validate() error {
	return dara.Validate(s)
}
