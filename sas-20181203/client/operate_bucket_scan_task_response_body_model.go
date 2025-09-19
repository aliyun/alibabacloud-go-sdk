// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBucketScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateBucketScanTaskResponseBody
	GetRequestId() *string
}

type OperateBucketScanTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0C8487EF-50C2-54BB-8634-10F8C35D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateBucketScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateBucketScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *OperateBucketScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateBucketScanTaskResponseBody) SetRequestId(v string) *OperateBucketScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateBucketScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
