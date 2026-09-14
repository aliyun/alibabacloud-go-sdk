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
	SetSuccess(v bool) *GetRecordingResponseBody
	GetSuccess() *bool
}

type GetRecordingResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetRecordingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *GetRecordingResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *GetRecordingResponseBody) SetSuccess(v bool) *GetRecordingResponseBody {
	s.Success = &v
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
	// The file name, including the file name extension.
	//
	// example:
	//
	// test_waveform.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The URL of the call recording file.
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The session ID.
	//
	// example:
	//
	// job-0b84bf6f-73dc-4462-bd8f-916e3a34c419
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
