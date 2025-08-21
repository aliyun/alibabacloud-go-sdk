// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketAcl(v string) *GetBucketAclResponseBody
	GetBucketAcl() *string
	SetRequestId(v string) *GetBucketAclResponseBody
	GetRequestId() *string
}

type GetBucketAclResponseBody struct {
	// The ACL of the bucket.
	//
	// example:
	//
	// private
	BucketAcl *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5C881388-2D4B-46F4-A96B-D4E6BD0886A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBucketAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketAclResponseBody) GetBucketAcl() *string {
	return s.BucketAcl
}

func (s *GetBucketAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBucketAclResponseBody) SetBucketAcl(v string) *GetBucketAclResponseBody {
	s.BucketAcl = &v
	return s
}

func (s *GetBucketAclResponseBody) SetRequestId(v string) *GetBucketAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBucketAclResponseBody) Validate() error {
	return dara.Validate(s)
}
