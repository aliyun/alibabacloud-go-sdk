// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindOutputBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *UnbindOutputBucketRequest
	GetBucket() *string
	SetOwnerAccount(v string) *UnbindOutputBucketRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnbindOutputBucketRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UnbindOutputBucketRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindOutputBucketRequest
	GetResourceOwnerId() *int64
}

type UnbindOutputBucketRequest struct {
	// The ID of the request.
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

func (s UnbindOutputBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindOutputBucketRequest) GoString() string {
	return s.String()
}

func (s *UnbindOutputBucketRequest) GetBucket() *string {
	return s.Bucket
}

func (s *UnbindOutputBucketRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnbindOutputBucketRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindOutputBucketRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindOutputBucketRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindOutputBucketRequest) SetBucket(v string) *UnbindOutputBucketRequest {
	s.Bucket = &v
	return s
}

func (s *UnbindOutputBucketRequest) SetOwnerAccount(v string) *UnbindOutputBucketRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindOutputBucketRequest) SetOwnerId(v int64) *UnbindOutputBucketRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindOutputBucketRequest) SetResourceOwnerAccount(v string) *UnbindOutputBucketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindOutputBucketRequest) SetResourceOwnerId(v int64) *UnbindOutputBucketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindOutputBucketRequest) Validate() error {
	return dara.Validate(s)
}
