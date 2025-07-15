// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEarlyMediaRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEarlyMediaRecordingResponseBody
	GetCode() *string
	SetData(v *GetEarlyMediaRecordingResponseBodyData) *GetEarlyMediaRecordingResponseBody
	GetData() *GetEarlyMediaRecordingResponseBodyData
	SetHttpStatusCode(v int32) *GetEarlyMediaRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetEarlyMediaRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEarlyMediaRecordingResponseBody
	GetRequestId() *string
}

type GetEarlyMediaRecordingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetEarlyMediaRecordingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEarlyMediaRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEarlyMediaRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *GetEarlyMediaRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEarlyMediaRecordingResponseBody) GetData() *GetEarlyMediaRecordingResponseBodyData {
	return s.Data
}

func (s *GetEarlyMediaRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetEarlyMediaRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEarlyMediaRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEarlyMediaRecordingResponseBody) SetCode(v string) *GetEarlyMediaRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *GetEarlyMediaRecordingResponseBody) SetData(v *GetEarlyMediaRecordingResponseBodyData) *GetEarlyMediaRecordingResponseBody {
	s.Data = v
	return s
}

func (s *GetEarlyMediaRecordingResponseBody) SetHttpStatusCode(v int32) *GetEarlyMediaRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetEarlyMediaRecordingResponseBody) SetMessage(v string) *GetEarlyMediaRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *GetEarlyMediaRecordingResponseBody) SetRequestId(v string) *GetEarlyMediaRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEarlyMediaRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEarlyMediaRecordingResponseBodyData struct {
	// example:
	//
	// job-6538214103689****-earlyMedia..wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com/ccc-record-mixed/ccc-test/2022/06/voicemail.job-054ded02****.wav?Expires=1656074923&OSSAccessKeyId=****&Signature=****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetEarlyMediaRecordingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEarlyMediaRecordingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEarlyMediaRecordingResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetEarlyMediaRecordingResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetEarlyMediaRecordingResponseBodyData) SetFileName(v string) *GetEarlyMediaRecordingResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetEarlyMediaRecordingResponseBodyData) SetFileUrl(v string) *GetEarlyMediaRecordingResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetEarlyMediaRecordingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
