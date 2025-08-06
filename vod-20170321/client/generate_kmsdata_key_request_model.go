// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateKMSDataKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GenerateKMSDataKeyRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GenerateKMSDataKeyRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GenerateKMSDataKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GenerateKMSDataKeyRequest
	GetResourceOwnerId() *string
}

type GenerateKMSDataKeyRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GenerateKMSDataKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateKMSDataKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateKMSDataKeyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GenerateKMSDataKeyRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GenerateKMSDataKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GenerateKMSDataKeyRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GenerateKMSDataKeyRequest) SetOwnerAccount(v string) *GenerateKMSDataKeyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GenerateKMSDataKeyRequest) SetOwnerId(v string) *GenerateKMSDataKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *GenerateKMSDataKeyRequest) SetResourceOwnerAccount(v string) *GenerateKMSDataKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GenerateKMSDataKeyRequest) SetResourceOwnerId(v string) *GenerateKMSDataKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GenerateKMSDataKeyRequest) Validate() error {
	return dara.Validate(s)
}
