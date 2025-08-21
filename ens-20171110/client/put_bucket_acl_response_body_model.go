// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutBucketAclResponseBody
	GetRequestId() *string
}

type PutBucketAclResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C0066F05-3116-4BAA-B588-52EB2E7F5D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutBucketAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutBucketAclResponseBody) GoString() string {
	return s.String()
}

func (s *PutBucketAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutBucketAclResponseBody) SetRequestId(v string) *PutBucketAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutBucketAclResponseBody) Validate() error {
	return dara.Validate(s)
}
