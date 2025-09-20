// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachOSSBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachOSSBucketResponseBody
	GetRequestId() *string
}

type DetachOSSBucketResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5F74C5C9-5AC0-49F9-914D-E01589D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachOSSBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachOSSBucketResponseBody) GoString() string {
	return s.String()
}

func (s *DetachOSSBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachOSSBucketResponseBody) SetRequestId(v string) *DetachOSSBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachOSSBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
