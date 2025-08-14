// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSmapleBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DownloadSmapleBatchResponseBody
	GetRequestId() *string
	SetResultObject(v *DownloadSmapleBatchResponseBodyResultObject) *DownloadSmapleBatchResponseBody
	GetResultObject() *DownloadSmapleBatchResponseBodyResultObject
}

type DownloadSmapleBatchResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DownloadSmapleBatchResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DownloadSmapleBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadSmapleBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSmapleBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadSmapleBatchResponseBody) GetResultObject() *DownloadSmapleBatchResponseBodyResultObject {
	return s.ResultObject
}

func (s *DownloadSmapleBatchResponseBody) SetRequestId(v string) *DownloadSmapleBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadSmapleBatchResponseBody) SetResultObject(v *DownloadSmapleBatchResponseBodyResultObject) *DownloadSmapleBatchResponseBody {
	s.ResultObject = v
	return s
}

func (s *DownloadSmapleBatchResponseBody) Validate() error {
	return dara.Validate(s)
}

type DownloadSmapleBatchResponseBodyResultObject struct {
	// Download URL
	//
	// example:
	//
	// https://xxxxx-oss-xxxxx.xxxxxx.aliyuncs.com/xx/xx/xxx/xxxxxx.csv?Expires=1753433384&OSSAccessKeyId=xxxxxxxxx&Signature=%2F%xxxxxxxxxxxx%3D
	FileDownloadURL *string `json:"fileDownloadURL,omitempty" xml:"fileDownloadURL,omitempty"`
}

func (s DownloadSmapleBatchResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DownloadSmapleBatchResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DownloadSmapleBatchResponseBodyResultObject) GetFileDownloadURL() *string {
	return s.FileDownloadURL
}

func (s *DownloadSmapleBatchResponseBodyResultObject) SetFileDownloadURL(v string) *DownloadSmapleBatchResponseBodyResultObject {
	s.FileDownloadURL = &v
	return s
}

func (s *DownloadSmapleBatchResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
