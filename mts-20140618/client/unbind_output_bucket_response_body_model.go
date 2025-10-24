// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindOutputBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindOutputBucketResponseBody
	GetRequestId() *string
}

type UnbindOutputBucketResponseBody struct {
	// The operation that you want to perform. Set the value to **UnbindOutputBucket**.
	//
	// example:
	//
	// 4AEA0480-32F4-1656-92B3-F4D4CDE6BBB3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindOutputBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindOutputBucketResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindOutputBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindOutputBucketResponseBody) SetRequestId(v string) *UnbindOutputBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindOutputBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
