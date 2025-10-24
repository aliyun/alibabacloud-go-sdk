// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInputBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindInputBucketResponseBody
	GetRequestId() *string
}

type UnbindInputBucketResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4AEA0480-32F4-1656-92B3-F4D4CDE6BBB3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindInputBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindInputBucketResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindInputBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindInputBucketResponseBody) SetRequestId(v string) *UnbindInputBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindInputBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
