// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadVerifyRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DownloadVerifyRecordsResponseBody
	GetRequestId() *string
	SetResultObject(v string) *DownloadVerifyRecordsResponseBody
	GetResultObject() *string
}

type DownloadVerifyRecordsResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// OSS link for file download.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth.oss-cn-shanghai.aliyuncs.com/invoke_download/0E30CD55-DCD4-5363-AA98-XXXXXXX
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s DownloadVerifyRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadVerifyRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadVerifyRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadVerifyRecordsResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *DownloadVerifyRecordsResponseBody) SetRequestId(v string) *DownloadVerifyRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadVerifyRecordsResponseBody) SetResultObject(v string) *DownloadVerifyRecordsResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DownloadVerifyRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
