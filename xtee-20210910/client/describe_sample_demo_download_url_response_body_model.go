// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDemoDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleDemoDownloadUrlResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeSampleDemoDownloadUrlResponseBody
	GetResultObject() *bool
}

type DescribeSampleDemoDownloadUrlResponseBody struct {
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

func (s DescribeSampleDemoDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDemoDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleDemoDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleDemoDownloadUrlResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeSampleDemoDownloadUrlResponseBody) SetRequestId(v string) *DescribeSampleDemoDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleDemoDownloadUrlResponseBody) SetResultObject(v bool) *DescribeSampleDemoDownloadUrlResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeSampleDemoDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
