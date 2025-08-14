// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsUrlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSlsUrlConfigResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DescribeSlsUrlConfigResponseBody
	GetResultObject() *string
}

type DescribeSlsUrlConfigResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *string `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeSlsUrlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsUrlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsUrlConfigResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DescribeSlsUrlConfigResponseBody) SetRequestId(v string) *DescribeSlsUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsUrlConfigResponseBody) SetResultObject(v string) *DescribeSlsUrlConfigResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeSlsUrlConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
