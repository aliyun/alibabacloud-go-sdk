// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiChannelRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMultiChannelRecordingResponseBody
	GetCode() *string
	SetData(v *GetMultiChannelRecordingResponseBodyData) *GetMultiChannelRecordingResponseBody
	GetData() *GetMultiChannelRecordingResponseBodyData
	SetHttpStatusCode(v int32) *GetMultiChannelRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMultiChannelRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMultiChannelRecordingResponseBody
	GetRequestId() *string
}

type GetMultiChannelRecordingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetMultiChannelRecordingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultiChannelRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiChannelRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiChannelRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMultiChannelRecordingResponseBody) GetData() *GetMultiChannelRecordingResponseBodyData {
	return s.Data
}

func (s *GetMultiChannelRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMultiChannelRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMultiChannelRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiChannelRecordingResponseBody) SetCode(v string) *GetMultiChannelRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) SetData(v *GetMultiChannelRecordingResponseBodyData) *GetMultiChannelRecordingResponseBody {
	s.Data = v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) SetHttpStatusCode(v int32) *GetMultiChannelRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) SetMessage(v string) *GetMultiChannelRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) SetRequestId(v string) *GetMultiChannelRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMultiChannelRecordingResponseBodyData struct {
	// example:
	//
	// job-6538214103689****.mkv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com/ccc-record-mixed/ccc-test/2021/04/job-6538214103689****.mkv?Expires=1617435462&OSSAccessKeyId=****&Signature=****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetMultiChannelRecordingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMultiChannelRecordingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMultiChannelRecordingResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetMultiChannelRecordingResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetMultiChannelRecordingResponseBodyData) SetFileName(v string) *GetMultiChannelRecordingResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBodyData) SetFileUrl(v string) *GetMultiChannelRecordingResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
