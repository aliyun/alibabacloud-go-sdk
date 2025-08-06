// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBucketRedundancyTransitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *SubmitBucketRedundancyTransitionRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitBucketRedundancyTransitionRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitBucketRedundancyTransitionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitBucketRedundancyTransitionRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *SubmitBucketRedundancyTransitionRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *SubmitBucketRedundancyTransitionRequest
	GetStorageLocation() *string
}

type SubmitBucketRedundancyTransitionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s SubmitBucketRedundancyTransitionRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBucketRedundancyTransitionRequest) GoString() string {
	return s.String()
}

func (s *SubmitBucketRedundancyTransitionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitBucketRedundancyTransitionRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitBucketRedundancyTransitionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitBucketRedundancyTransitionRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitBucketRedundancyTransitionRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *SubmitBucketRedundancyTransitionRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SubmitBucketRedundancyTransitionRequest) SetOwnerAccount(v string) *SubmitBucketRedundancyTransitionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionRequest) SetOwnerId(v string) *SubmitBucketRedundancyTransitionRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionRequest) SetResourceOwnerAccount(v string) *SubmitBucketRedundancyTransitionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionRequest) SetResourceOwnerId(v string) *SubmitBucketRedundancyTransitionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionRequest) SetResourceRealOwnerId(v int64) *SubmitBucketRedundancyTransitionRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionRequest) SetStorageLocation(v string) *SubmitBucketRedundancyTransitionRequest {
	s.StorageLocation = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionRequest) Validate() error {
	return dara.Validate(s)
}
