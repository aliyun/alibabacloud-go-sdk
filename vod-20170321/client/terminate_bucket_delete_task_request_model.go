// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateBucketDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *TerminateBucketDeleteTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *TerminateBucketDeleteTaskRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *TerminateBucketDeleteTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *TerminateBucketDeleteTaskRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *TerminateBucketDeleteTaskRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *TerminateBucketDeleteTaskRequest
	GetStorageLocation() *string
}

type TerminateBucketDeleteTaskRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s TerminateBucketDeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateBucketDeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *TerminateBucketDeleteTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TerminateBucketDeleteTaskRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *TerminateBucketDeleteTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TerminateBucketDeleteTaskRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *TerminateBucketDeleteTaskRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *TerminateBucketDeleteTaskRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *TerminateBucketDeleteTaskRequest) SetOwnerAccount(v string) *TerminateBucketDeleteTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TerminateBucketDeleteTaskRequest) SetOwnerId(v string) *TerminateBucketDeleteTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *TerminateBucketDeleteTaskRequest) SetResourceOwnerAccount(v string) *TerminateBucketDeleteTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TerminateBucketDeleteTaskRequest) SetResourceOwnerId(v string) *TerminateBucketDeleteTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TerminateBucketDeleteTaskRequest) SetResourceRealOwnerId(v int64) *TerminateBucketDeleteTaskRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *TerminateBucketDeleteTaskRequest) SetStorageLocation(v string) *TerminateBucketDeleteTaskRequest {
	s.StorageLocation = &v
	return s
}

func (s *TerminateBucketDeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
