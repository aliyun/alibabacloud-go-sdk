// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInputBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *UnbindInputBucketRequest
	GetBucket() *string
	SetOwnerAccount(v string) *UnbindInputBucketRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnbindInputBucketRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UnbindInputBucketRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindInputBucketRequest
	GetResourceOwnerId() *int64
	SetRoleArn(v string) *UnbindInputBucketRequest
	GetRoleArn() *string
}

type UnbindInputBucketRequest struct {
	// The name of the input media bucket to be unbound. To obtain the media bucket name, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Media Buckets*	- in the left-side navigation pane. Alternatively, you can log on to the **Object Storage Service (OSS) console*	- and click **Historical Paths**.
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
	// The Alibaba Cloud Resource Name (ARN) of the role used for delegated authorization.
	//
	// example:
	//
	// acs:ram::174809843091****:role/exampleRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s UnbindInputBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindInputBucketRequest) GoString() string {
	return s.String()
}

func (s *UnbindInputBucketRequest) GetBucket() *string {
	return s.Bucket
}

func (s *UnbindInputBucketRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnbindInputBucketRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindInputBucketRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindInputBucketRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindInputBucketRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UnbindInputBucketRequest) SetBucket(v string) *UnbindInputBucketRequest {
	s.Bucket = &v
	return s
}

func (s *UnbindInputBucketRequest) SetOwnerAccount(v string) *UnbindInputBucketRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindInputBucketRequest) SetOwnerId(v int64) *UnbindInputBucketRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindInputBucketRequest) SetResourceOwnerAccount(v string) *UnbindInputBucketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindInputBucketRequest) SetResourceOwnerId(v int64) *UnbindInputBucketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindInputBucketRequest) SetRoleArn(v string) *UnbindInputBucketRequest {
	s.RoleArn = &v
	return s
}

func (s *UnbindInputBucketRequest) Validate() error {
	return dara.Validate(s)
}
