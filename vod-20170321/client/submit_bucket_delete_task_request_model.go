// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBucketDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteFiles(v bool) *SubmitBucketDeleteTaskRequest
	GetDeleteFiles() *bool
	SetOwnerAccount(v string) *SubmitBucketDeleteTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitBucketDeleteTaskRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitBucketDeleteTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitBucketDeleteTaskRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *SubmitBucketDeleteTaskRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *SubmitBucketDeleteTaskRequest
	GetStorageLocation() *string
}

type SubmitBucketDeleteTaskRequest struct {
	DeleteFiles          *bool   `json:"DeleteFiles,omitempty" xml:"DeleteFiles,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s SubmitBucketDeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBucketDeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitBucketDeleteTaskRequest) GetDeleteFiles() *bool {
	return s.DeleteFiles
}

func (s *SubmitBucketDeleteTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitBucketDeleteTaskRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitBucketDeleteTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitBucketDeleteTaskRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitBucketDeleteTaskRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *SubmitBucketDeleteTaskRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SubmitBucketDeleteTaskRequest) SetDeleteFiles(v bool) *SubmitBucketDeleteTaskRequest {
	s.DeleteFiles = &v
	return s
}

func (s *SubmitBucketDeleteTaskRequest) SetOwnerAccount(v string) *SubmitBucketDeleteTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitBucketDeleteTaskRequest) SetOwnerId(v string) *SubmitBucketDeleteTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitBucketDeleteTaskRequest) SetResourceOwnerAccount(v string) *SubmitBucketDeleteTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitBucketDeleteTaskRequest) SetResourceOwnerId(v string) *SubmitBucketDeleteTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitBucketDeleteTaskRequest) SetResourceRealOwnerId(v int64) *SubmitBucketDeleteTaskRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SubmitBucketDeleteTaskRequest) SetStorageLocation(v string) *SubmitBucketDeleteTaskRequest {
	s.StorageLocation = &v
	return s
}

func (s *SubmitBucketDeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
