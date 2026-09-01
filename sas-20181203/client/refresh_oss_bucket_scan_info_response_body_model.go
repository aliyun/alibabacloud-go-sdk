// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshOssBucketScanInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshOssBucketScanInfoResponseBody
	GetRequestId() *string
}

type RefreshOssBucketScanInfoResponseBody struct {
	// The ID of the request. The China (Hangzhou) region generates a unique identifier for the request, which can be used for troubleshooting.
	//
	// example:
	//
	// CE290C1F-4B7D-5024-9D2F-E26D7B08****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshOssBucketScanInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshOssBucketScanInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshOssBucketScanInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshOssBucketScanInfoResponseBody) SetRequestId(v string) *RefreshOssBucketScanInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshOssBucketScanInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
