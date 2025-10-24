// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindOutputBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindOutputBucketResponseBody
	GetRequestId() *string
}

type BindOutputBucketResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D0F80646-90D4-402F-9D56-CEFEAA6BCC9B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindOutputBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindOutputBucketResponseBody) GoString() string {
	return s.String()
}

func (s *BindOutputBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindOutputBucketResponseBody) SetRequestId(v string) *BindOutputBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindOutputBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
