// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonoRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMonoRecordingResponseBody
	GetCode() *string
	SetData(v *GetMonoRecordingResponseBodyData) *GetMonoRecordingResponseBody
	GetData() *GetMonoRecordingResponseBodyData
	SetHttpStatusCode(v int32) *GetMonoRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMonoRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMonoRecordingResponseBody
	GetRequestId() *string
}

type GetMonoRecordingResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetMonoRecordingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetMonoRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMonoRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonoRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMonoRecordingResponseBody) GetData() *GetMonoRecordingResponseBodyData {
	return s.Data
}

func (s *GetMonoRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMonoRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMonoRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMonoRecordingResponseBody) SetCode(v string) *GetMonoRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *GetMonoRecordingResponseBody) SetData(v *GetMonoRecordingResponseBodyData) *GetMonoRecordingResponseBody {
	s.Data = v
	return s
}

func (s *GetMonoRecordingResponseBody) SetHttpStatusCode(v int32) *GetMonoRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMonoRecordingResponseBody) SetMessage(v string) *GetMonoRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *GetMonoRecordingResponseBody) SetRequestId(v string) *GetMonoRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonoRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMonoRecordingResponseBodyData struct {
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// job-6538214103689****.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// http://ccc-v2-online.oss-cn-shanghai.aliyuncs.com/ccc-record/job-6538214103689****.wav?Expires=1610910578&OSSAccessKeyId=****&Signature=****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetMonoRecordingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMonoRecordingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMonoRecordingResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *GetMonoRecordingResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetMonoRecordingResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetMonoRecordingResponseBodyData) SetDuration(v int64) *GetMonoRecordingResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetMonoRecordingResponseBodyData) SetFileName(v string) *GetMonoRecordingResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetMonoRecordingResponseBodyData) SetFileUrl(v string) *GetMonoRecordingResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetMonoRecordingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
