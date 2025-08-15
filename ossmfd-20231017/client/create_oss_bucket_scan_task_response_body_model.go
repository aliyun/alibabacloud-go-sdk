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
	// example:
	//
	// E99D386F-5546-5BCD-BADD-F4EF4B******
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
