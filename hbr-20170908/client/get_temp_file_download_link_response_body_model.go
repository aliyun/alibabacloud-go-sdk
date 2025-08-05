// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempFileDownloadLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTempFileDownloadLinkResponseBody
	GetCode() *string
	SetMessage(v string) *GetTempFileDownloadLinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTempFileDownloadLinkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTempFileDownloadLinkResponseBody
	GetSuccess() *bool
	SetUrl(v string) *GetTempFileDownloadLinkResponseBody
	GetUrl() *string
}

type GetTempFileDownloadLinkResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the request is successful, a value of successful is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The download URL of the file.
	//
	// example:
	//
	// https://a-hbr-temp-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/job-0007yg2i0m6705wdhgb6_0.csv?Expires=1649406469&OSSAccessKeyId=LTAI************&Signature=26%2BgjegCrRmMDCpS5jzyG4ivKU8%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTempFileDownloadLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTempFileDownloadLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetTempFileDownloadLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTempFileDownloadLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTempFileDownloadLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTempFileDownloadLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTempFileDownloadLinkResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetTempFileDownloadLinkResponseBody) SetCode(v string) *GetTempFileDownloadLinkResponseBody {
	s.Code = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) SetMessage(v string) *GetTempFileDownloadLinkResponseBody {
	s.Message = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) SetRequestId(v string) *GetTempFileDownloadLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) SetSuccess(v bool) *GetTempFileDownloadLinkResponseBody {
	s.Success = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) SetUrl(v string) *GetTempFileDownloadLinkResponseBody {
	s.Url = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
