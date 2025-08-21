// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutBucketResponseBody
	GetRequestId() *string
}

type PutBucketResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 85123E71-7710-4620-8AAB-133CCE49EC83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutBucketResponseBody) GoString() string {
	return s.String()
}

func (s *PutBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutBucketResponseBody) SetRequestId(v string) *PutBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
