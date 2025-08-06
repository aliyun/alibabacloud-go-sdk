// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketRedundancyTransitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ListBucketRedundancyTransitionRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListBucketRedundancyTransitionRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListBucketRedundancyTransitionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListBucketRedundancyTransitionRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *ListBucketRedundancyTransitionRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *ListBucketRedundancyTransitionRequest
	GetStorageLocation() *string
}

type ListBucketRedundancyTransitionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s ListBucketRedundancyTransitionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBucketRedundancyTransitionRequest) GoString() string {
	return s.String()
}

func (s *ListBucketRedundancyTransitionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListBucketRedundancyTransitionRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListBucketRedundancyTransitionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListBucketRedundancyTransitionRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListBucketRedundancyTransitionRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *ListBucketRedundancyTransitionRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *ListBucketRedundancyTransitionRequest) SetOwnerAccount(v string) *ListBucketRedundancyTransitionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListBucketRedundancyTransitionRequest) SetOwnerId(v string) *ListBucketRedundancyTransitionRequest {
	s.OwnerId = &v
	return s
}

func (s *ListBucketRedundancyTransitionRequest) SetResourceOwnerAccount(v string) *ListBucketRedundancyTransitionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListBucketRedundancyTransitionRequest) SetResourceOwnerId(v string) *ListBucketRedundancyTransitionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListBucketRedundancyTransitionRequest) SetResourceRealOwnerId(v int64) *ListBucketRedundancyTransitionRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *ListBucketRedundancyTransitionRequest) SetStorageLocation(v string) *ListBucketRedundancyTransitionRequest {
	s.StorageLocation = &v
	return s
}

func (s *ListBucketRedundancyTransitionRequest) Validate() error {
	return dara.Validate(s)
}
