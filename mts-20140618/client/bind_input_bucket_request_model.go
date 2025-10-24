// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInputBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *BindInputBucketRequest
	GetBucket() *string
	SetOwnerAccount(v string) *BindInputBucketRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BindInputBucketRequest
	GetOwnerId() *int64
	SetReferer(v string) *BindInputBucketRequest
	GetReferer() *string
	SetResourceOwnerAccount(v string) *BindInputBucketRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindInputBucketRequest
	GetResourceOwnerId() *int64
}

type BindInputBucketRequest struct {
	// The name of the input media bucket to be bound. The name can be up to 64 bytes in size. To obtain the media bucket name, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Media Buckets*	- in the left-side navigation pane.
	//
	// > The bucket name can contain lowercase letters, digits, and hyphens (-), and cannot start or end with a hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// example-bucket-****
	Bucket       *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The settings of Object Storage Service (OSS) hotlink protection. For more information, see [Hotlink protection](https://help.aliyun.com/document_detail/31869.html).
	//
	// example:
	//
	// http://www.example.com
	Referer              *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindInputBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s BindInputBucketRequest) GoString() string {
	return s.String()
}

func (s *BindInputBucketRequest) GetBucket() *string {
	return s.Bucket
}

func (s *BindInputBucketRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BindInputBucketRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindInputBucketRequest) GetReferer() *string {
	return s.Referer
}

func (s *BindInputBucketRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindInputBucketRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindInputBucketRequest) SetBucket(v string) *BindInputBucketRequest {
	s.Bucket = &v
	return s
}

func (s *BindInputBucketRequest) SetOwnerAccount(v string) *BindInputBucketRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindInputBucketRequest) SetOwnerId(v int64) *BindInputBucketRequest {
	s.OwnerId = &v
	return s
}

func (s *BindInputBucketRequest) SetReferer(v string) *BindInputBucketRequest {
	s.Referer = &v
	return s
}

func (s *BindInputBucketRequest) SetResourceOwnerAccount(v string) *BindInputBucketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindInputBucketRequest) SetResourceOwnerId(v int64) *BindInputBucketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindInputBucketRequest) Validate() error {
	return dara.Validate(s)
}
