// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindOutputBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *BindOutputBucketRequest
	GetBucket() *string
	SetOwnerAccount(v string) *BindOutputBucketRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BindOutputBucketRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *BindOutputBucketRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindOutputBucketRequest
	GetResourceOwnerId() *int64
}

type BindOutputBucketRequest struct {
	// The name of the Object Storage Service (OSS) bucket that you want to bind. The name can be up to 64 bytes in size and can contain letters, digits, and hyphens (-). The name cannot start with a special character.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-bucket-****
	Bucket               *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindOutputBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s BindOutputBucketRequest) GoString() string {
	return s.String()
}

func (s *BindOutputBucketRequest) GetBucket() *string {
	return s.Bucket
}

func (s *BindOutputBucketRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BindOutputBucketRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindOutputBucketRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindOutputBucketRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindOutputBucketRequest) SetBucket(v string) *BindOutputBucketRequest {
	s.Bucket = &v
	return s
}

func (s *BindOutputBucketRequest) SetOwnerAccount(v string) *BindOutputBucketRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindOutputBucketRequest) SetOwnerId(v int64) *BindOutputBucketRequest {
	s.OwnerId = &v
	return s
}

func (s *BindOutputBucketRequest) SetResourceOwnerAccount(v string) *BindOutputBucketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindOutputBucketRequest) SetResourceOwnerId(v int64) *BindOutputBucketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindOutputBucketRequest) Validate() error {
	return dara.Validate(s)
}
