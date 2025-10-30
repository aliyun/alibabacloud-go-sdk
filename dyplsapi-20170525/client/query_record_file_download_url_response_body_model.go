// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordFileDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRecordFileDownloadUrlResponseBody
	GetCode() *string
	SetDownloadUrl(v string) *QueryRecordFileDownloadUrlResponseBody
	GetDownloadUrl() *string
	SetMessage(v string) *QueryRecordFileDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRecordFileDownloadUrlResponseBody
	GetRequestId() *string
}

type QueryRecordFileDownloadUrlResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The download URL of the recording file. The download URL is valid for 2 hours.
	//
	// example:
	//
	// http://secret-axb-reco****cn-shanghai.aliyuncs.com/1000000820257625_66647243838006067251551752068865.mp3?Expires=155175****&OSSAccessKeyId=LTAIP00vvvv****v&Signature=tK6Yq9KusU4n%2BZ****7lg4/WmEA%3D
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AB3CEF7-DCBE-488C-9C33-D180982CE031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRecordFileDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordFileDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRecordFileDownloadUrlResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *QueryRecordFileDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRecordFileDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetCode(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetDownloadUrl(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetMessage(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) SetRequestId(v string) *QueryRecordFileDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
