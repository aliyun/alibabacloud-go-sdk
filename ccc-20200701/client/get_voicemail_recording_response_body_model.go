// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVoicemailRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVoicemailRecordingResponseBody
	GetCode() *string
	SetData(v *GetVoicemailRecordingResponseBodyData) *GetVoicemailRecordingResponseBody
	GetData() *GetVoicemailRecordingResponseBodyData
	SetHttpStatusCode(v int32) *GetVoicemailRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVoicemailRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVoicemailRecordingResponseBody
	GetRequestId() *string
}

type GetVoicemailRecordingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetVoicemailRecordingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetVoicemailRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVoicemailRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *GetVoicemailRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVoicemailRecordingResponseBody) GetData() *GetVoicemailRecordingResponseBodyData {
	return s.Data
}

func (s *GetVoicemailRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVoicemailRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVoicemailRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVoicemailRecordingResponseBody) SetCode(v string) *GetVoicemailRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *GetVoicemailRecordingResponseBody) SetData(v *GetVoicemailRecordingResponseBodyData) *GetVoicemailRecordingResponseBody {
	s.Data = v
	return s
}

func (s *GetVoicemailRecordingResponseBody) SetHttpStatusCode(v int32) *GetVoicemailRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVoicemailRecordingResponseBody) SetMessage(v string) *GetVoicemailRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *GetVoicemailRecordingResponseBody) SetRequestId(v string) *GetVoicemailRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVoicemailRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVoicemailRecordingResponseBodyData struct {
	// example:
	//
	// voicemail.job-054ded02****.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com/ccc-record-mixed/ccc-test/2022/06/voicemail.job-054ded02****.wav?Expires=1656074923&OSSAccessKeyId=****&Signature=****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetVoicemailRecordingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVoicemailRecordingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVoicemailRecordingResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetVoicemailRecordingResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetVoicemailRecordingResponseBodyData) SetFileName(v string) *GetVoicemailRecordingResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetVoicemailRecordingResponseBodyData) SetFileUrl(v string) *GetVoicemailRecordingResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetVoicemailRecordingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
