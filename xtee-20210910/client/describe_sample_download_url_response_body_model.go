// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleDownloadUrlResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeSampleDownloadUrlResponseBody
	GetResultObject() *bool
}

type DescribeSampleDownloadUrlResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeSampleDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleDownloadUrlResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeSampleDownloadUrlResponseBody) SetRequestId(v string) *DescribeSampleDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleDownloadUrlResponseBody) SetResultObject(v bool) *DescribeSampleDownloadUrlResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeSampleDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
