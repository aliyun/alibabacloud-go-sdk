// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceDataDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDataSourceDataDownloadUrlResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeDataSourceDataDownloadUrlResponseBodyResultObject) *DescribeDataSourceDataDownloadUrlResponseBody
	GetResultObject() *DescribeDataSourceDataDownloadUrlResponseBodyResultObject
}

type DescribeDataSourceDataDownloadUrlResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeDataSourceDataDownloadUrlResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeDataSourceDataDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceDataDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceDataDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataSourceDataDownloadUrlResponseBody) GetResultObject() *DescribeDataSourceDataDownloadUrlResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeDataSourceDataDownloadUrlResponseBody) SetRequestId(v string) *DescribeDataSourceDataDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlResponseBody) SetResultObject(v *DescribeDataSourceDataDownloadUrlResponseBodyResultObject) *DescribeDataSourceDataDownloadUrlResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataSourceDataDownloadUrlResponseBodyResultObject struct {
	// Download URL.
	//
	// example:
	//
	// https:/xxxxx.oss-cn-xxxxxx.aliyuncs.com/xxxx/xxx/xxxxxx/xxx/xxxxxxxxxx?Expires=1753421064&OSSAccessKeyId=xxxxxxx&Signature=xxxxxxx%3D
	DownloadFileUrl *string `json:"downloadFileUrl,omitempty" xml:"downloadFileUrl,omitempty"`
}

func (s DescribeDataSourceDataDownloadUrlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceDataDownloadUrlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceDataDownloadUrlResponseBodyResultObject) GetDownloadFileUrl() *string {
	return s.DownloadFileUrl
}

func (s *DescribeDataSourceDataDownloadUrlResponseBodyResultObject) SetDownloadFileUrl(v string) *DescribeDataSourceDataDownloadUrlResponseBodyResultObject {
	s.DownloadFileUrl = &v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
