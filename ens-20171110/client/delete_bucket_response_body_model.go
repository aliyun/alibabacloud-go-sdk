// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBucketResponseBody
	GetRequestId() *string
}

type DeleteBucketResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2180AB96-57CF-1C68-9FEE-D60E547FD79C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBucketResponseBody) SetRequestId(v string) *DeleteBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
