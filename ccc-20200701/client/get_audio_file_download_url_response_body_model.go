// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAudioFileDownloadUrlResponseBody
	GetCode() *string
	SetData(v string) *GetAudioFileDownloadUrlResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetAudioFileDownloadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAudioFileDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAudioFileDownloadUrlResponseBody
	GetRequestId() *string
}

type GetAudioFileDownloadUrlResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data. The OSS download link for the audio resource file. Note that the validity period of the download link is 24 hours. Do not store this link permanently.
	//
	// example:
	//
	// https://****.oss-cn-shanghai.aliyuncs.com/ccc-test/****.wav?Expires=1656472158&OSSAccessKeyId=****&Signature=****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 943D8EF3-3321-471F-A104-51C96FCA94D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAudioFileDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAudioFileDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAudioFileDownloadUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAudioFileDownloadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAudioFileDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAudioFileDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAudioFileDownloadUrlResponseBody) SetCode(v string) *GetAudioFileDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAudioFileDownloadUrlResponseBody) SetData(v string) *GetAudioFileDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetAudioFileDownloadUrlResponseBody) SetHttpStatusCode(v int32) *GetAudioFileDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAudioFileDownloadUrlResponseBody) SetMessage(v string) *GetAudioFileDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAudioFileDownloadUrlResponseBody) SetRequestId(v string) *GetAudioFileDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAudioFileDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
