// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssBucketScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOssBucketScanTaskResponseBody
	GetRequestId() *string
}

type CreateOssBucketScanTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BBD75EC2-2F4F-5A7B-AA53-18724DC8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOssBucketScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOssBucketScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOssBucketScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOssBucketScanTaskResponseBody) SetRequestId(v string) *CreateOssBucketScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOssBucketScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
