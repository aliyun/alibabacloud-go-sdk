// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketAcl(v string) *PutBucketAclRequest
	GetBucketAcl() *string
	SetBucketName(v string) *PutBucketAclRequest
	GetBucketName() *string
}

type PutBucketAclRequest struct {
	// The access control list (ACL) of the bucket.
	//
	// 	- **public-read-write**
	//
	// 	- **public-read**
	//
	// 	- **private*	- (default)
	//
	// This parameter is required.
	//
	// example:
	//
	// private
	BucketAcl *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s PutBucketAclRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketAclRequest) GoString() string {
	return s.String()
}

func (s *PutBucketAclRequest) GetBucketAcl() *string {
	return s.BucketAcl
}

func (s *PutBucketAclRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *PutBucketAclRequest) SetBucketAcl(v string) *PutBucketAclRequest {
	s.BucketAcl = &v
	return s
}

func (s *PutBucketAclRequest) SetBucketName(v string) *PutBucketAclRequest {
	s.BucketName = &v
	return s
}

func (s *PutBucketAclRequest) Validate() error {
	return dara.Validate(s)
}
