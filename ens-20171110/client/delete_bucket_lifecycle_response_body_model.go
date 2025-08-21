// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketLifecycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBucketLifecycleResponseBody
	GetRequestId() *string
}

type DeleteBucketLifecycleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C6583E8B-B930-4F59-ADC0-0E209A45E860
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBucketLifecycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketLifecycleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBucketLifecycleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBucketLifecycleResponseBody) SetRequestId(v string) *DeleteBucketLifecycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBucketLifecycleResponseBody) Validate() error {
	return dara.Validate(s)
}
