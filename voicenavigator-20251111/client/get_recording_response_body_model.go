// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRecordingResponseBody
	GetCode() *string
	SetData(v *GetRecordingResponseBodyData) *GetRecordingResponseBody
	GetData() *GetRecordingResponseBodyData
	SetHttpStatusCode(v int32) *GetRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRecordingResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetRecordingResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetRecordingResponseBody
	GetRequestId() *string
}

type GetRecordingResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRecordingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// F132DDBA-66C4-5BD3-BF7E-9642FE877158
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRecordingResponseBody) GetData() *GetRecordingResponseBodyData {
	return s.Data
}

func (s *GetRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRecordingResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecordingResponseBody) SetCode(v string) *GetRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecordingResponseBody) SetData(v *GetRecordingResponseBodyData) *GetRecordingResponseBody {
	s.Data = v
	return s
}

func (s *GetRecordingResponseBody) SetHttpStatusCode(v int32) *GetRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRecordingResponseBody) SetMessage(v string) *GetRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecordingResponseBody) SetParams(v []*string) *GetRecordingResponseBody {
	s.Params = v
	return s
}

func (s *GetRecordingResponseBody) SetRequestId(v string) *GetRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRecordingResponseBodyData struct {
	// example:
	//
	// 61f7f225-3a2e-4b6e-8a1d-888f1862590f.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://bgc-us-east.oss-accelerate.aliyuncs.com/intellidocs/2025-08-27/bc84ce11-f110-436d-af70-382ce7c5690a.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// f814f3ae-b2a7-44fb-883c-771221274673
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetRecordingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRecordingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecordingResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetRecordingResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetRecordingResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetRecordingResponseBodyData) SetFileName(v string) *GetRecordingResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetRecordingResponseBodyData) SetFileUrl(v string) *GetRecordingResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetRecordingResponseBodyData) SetSessionId(v string) *GetRecordingResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetRecordingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
