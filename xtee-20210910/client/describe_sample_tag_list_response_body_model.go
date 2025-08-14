// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleTagListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleTagListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeSampleTagListResponseBody
	GetResultObject() *bool
}

type DescribeSampleTagListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeSampleTagListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleTagListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleTagListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleTagListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeSampleTagListResponseBody) SetRequestId(v string) *DescribeSampleTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleTagListResponseBody) SetResultObject(v bool) *DescribeSampleTagListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeSampleTagListResponseBody) Validate() error {
	return dara.Validate(s)
}
