// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetStorageNotifyConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetStorageNotifyConfigRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetStorageNotifyConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetStorageNotifyConfigRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *GetStorageNotifyConfigRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *GetStorageNotifyConfigRequest
	GetStorageLocation() *string
}

type GetStorageNotifyConfigRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s GetStorageNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *GetStorageNotifyConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetStorageNotifyConfigRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetStorageNotifyConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetStorageNotifyConfigRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetStorageNotifyConfigRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetStorageNotifyConfigRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetStorageNotifyConfigRequest) SetOwnerAccount(v string) *GetStorageNotifyConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetStorageNotifyConfigRequest) SetOwnerId(v string) *GetStorageNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetStorageNotifyConfigRequest) SetResourceOwnerAccount(v string) *GetStorageNotifyConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetStorageNotifyConfigRequest) SetResourceOwnerId(v string) *GetStorageNotifyConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetStorageNotifyConfigRequest) SetResourceRealOwnerId(v int64) *GetStorageNotifyConfigRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetStorageNotifyConfigRequest) SetStorageLocation(v string) *GetStorageNotifyConfigRequest {
	s.StorageLocation = &v
	return s
}

func (s *GetStorageNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
