// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelBucketRedundancyTransitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CancelBucketRedundancyTransitionRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *CancelBucketRedundancyTransitionRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *CancelBucketRedundancyTransitionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *CancelBucketRedundancyTransitionRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *CancelBucketRedundancyTransitionRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *CancelBucketRedundancyTransitionRequest
	GetStorageLocation() *string
	SetTaskId(v string) *CancelBucketRedundancyTransitionRequest
	GetTaskId() *string
}

type CancelBucketRedundancyTransitionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelBucketRedundancyTransitionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelBucketRedundancyTransitionRequest) GoString() string {
	return s.String()
}

func (s *CancelBucketRedundancyTransitionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelBucketRedundancyTransitionRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CancelBucketRedundancyTransitionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelBucketRedundancyTransitionRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *CancelBucketRedundancyTransitionRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *CancelBucketRedundancyTransitionRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *CancelBucketRedundancyTransitionRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelBucketRedundancyTransitionRequest) SetOwnerAccount(v string) *CancelBucketRedundancyTransitionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelBucketRedundancyTransitionRequest) SetOwnerId(v string) *CancelBucketRedundancyTransitionRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelBucketRedundancyTransitionRequest) SetResourceOwnerAccount(v string) *CancelBucketRedundancyTransitionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelBucketRedundancyTransitionRequest) SetResourceOwnerId(v string) *CancelBucketRedundancyTransitionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelBucketRedundancyTransitionRequest) SetResourceRealOwnerId(v int64) *CancelBucketRedundancyTransitionRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *CancelBucketRedundancyTransitionRequest) SetStorageLocation(v string) *CancelBucketRedundancyTransitionRequest {
	s.StorageLocation = &v
	return s
}

func (s *CancelBucketRedundancyTransitionRequest) SetTaskId(v string) *CancelBucketRedundancyTransitionRequest {
	s.TaskId = &v
	return s
}

func (s *CancelBucketRedundancyTransitionRequest) Validate() error {
	return dara.Validate(s)
}
