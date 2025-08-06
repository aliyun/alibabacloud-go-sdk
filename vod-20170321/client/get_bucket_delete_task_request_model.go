// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetBucketDeleteTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetBucketDeleteTaskRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetBucketDeleteTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetBucketDeleteTaskRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *GetBucketDeleteTaskRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *GetBucketDeleteTaskRequest
	GetStorageLocation() *string
}

type GetBucketDeleteTaskRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s GetBucketDeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *GetBucketDeleteTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetBucketDeleteTaskRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetBucketDeleteTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetBucketDeleteTaskRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetBucketDeleteTaskRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetBucketDeleteTaskRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetBucketDeleteTaskRequest) SetOwnerAccount(v string) *GetBucketDeleteTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetBucketDeleteTaskRequest) SetOwnerId(v string) *GetBucketDeleteTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *GetBucketDeleteTaskRequest) SetResourceOwnerAccount(v string) *GetBucketDeleteTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetBucketDeleteTaskRequest) SetResourceOwnerId(v string) *GetBucketDeleteTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetBucketDeleteTaskRequest) SetResourceRealOwnerId(v int64) *GetBucketDeleteTaskRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetBucketDeleteTaskRequest) SetStorageLocation(v string) *GetBucketDeleteTaskRequest {
	s.StorageLocation = &v
	return s
}

func (s *GetBucketDeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
