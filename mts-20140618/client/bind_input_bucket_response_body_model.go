// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInputBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindInputBucketResponseBody
	GetRequestId() *string
}

type BindInputBucketResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4AEA0480-32F4-1656-92B3-F4D4CDE6BBB3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindInputBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindInputBucketResponseBody) GoString() string {
	return s.String()
}

func (s *BindInputBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindInputBucketResponseBody) SetRequestId(v string) *BindInputBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindInputBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
