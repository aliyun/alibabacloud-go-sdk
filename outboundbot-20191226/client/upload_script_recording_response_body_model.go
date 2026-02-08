// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadScriptRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadScriptRecordingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UploadScriptRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UploadScriptRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadScriptRecordingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadScriptRecordingResponseBody
	GetSuccess() *bool
	SetUuid(v string) *UploadScriptRecordingResponseBody
	GetUuid() *string
}

type UploadScriptRecordingResponseBody struct {
	// Return code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Unique ID of the recording
	//
	// example:
	//
	// 5feaab8a-97fd-4720-8108-79e017f2d3ac
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UploadScriptRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadScriptRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *UploadScriptRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadScriptRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UploadScriptRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadScriptRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadScriptRecordingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadScriptRecordingResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *UploadScriptRecordingResponseBody) SetCode(v string) *UploadScriptRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *UploadScriptRecordingResponseBody) SetHttpStatusCode(v int32) *UploadScriptRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadScriptRecordingResponseBody) SetMessage(v string) *UploadScriptRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *UploadScriptRecordingResponseBody) SetRequestId(v string) *UploadScriptRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadScriptRecordingResponseBody) SetSuccess(v bool) *UploadScriptRecordingResponseBody {
	s.Success = &v
	return s
}

func (s *UploadScriptRecordingResponseBody) SetUuid(v string) *UploadScriptRecordingResponseBody {
	s.Uuid = &v
	return s
}

func (s *UploadScriptRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}
