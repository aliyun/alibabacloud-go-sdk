// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeNameListDownloadUrlResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeNameListDownloadUrlResponseBodyResultObject) *DescribeNameListDownloadUrlResponseBody
	GetResultObject() *DescribeNameListDownloadUrlResponseBodyResultObject
}

type DescribeNameListDownloadUrlResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeNameListDownloadUrlResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeNameListDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNameListDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNameListDownloadUrlResponseBody) GetResultObject() *DescribeNameListDownloadUrlResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeNameListDownloadUrlResponseBody) SetRequestId(v string) *DescribeNameListDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNameListDownloadUrlResponseBody) SetResultObject(v *DescribeNameListDownloadUrlResponseBodyResultObject) *DescribeNameListDownloadUrlResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeNameListDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNameListDownloadUrlResponseBodyResultObject struct {
	// Download URL.
	//
	// example:
	//
	// https://xxxxx-oss-xxxxx.xxxxxx.aliyuncs.com/xx/xx/xxx/xxxxxx.csv?Expires=1753433384&OSSAccessKeyId=xxxxxxxxx&Signature=%2F%xxxxxxxxxxxx%3D
	DownloadFileUrl *string `json:"downloadFileUrl,omitempty" xml:"downloadFileUrl,omitempty"`
}

func (s DescribeNameListDownloadUrlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListDownloadUrlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeNameListDownloadUrlResponseBodyResultObject) GetDownloadFileUrl() *string {
	return s.DownloadFileUrl
}

func (s *DescribeNameListDownloadUrlResponseBodyResultObject) SetDownloadFileUrl(v string) *DescribeNameListDownloadUrlResponseBodyResultObject {
	s.DownloadFileUrl = &v
	return s
}

func (s *DescribeNameListDownloadUrlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
