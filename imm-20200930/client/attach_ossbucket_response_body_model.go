// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachOSSBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachOSSBucketResponseBody
	GetRequestId() *string
}

type AttachOSSBucketResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5F74C5C9-5AC0-49F9-914D-E01589D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachOSSBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachOSSBucketResponseBody) GoString() string {
	return s.String()
}

func (s *AttachOSSBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachOSSBucketResponseBody) SetRequestId(v string) *AttachOSSBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachOSSBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
