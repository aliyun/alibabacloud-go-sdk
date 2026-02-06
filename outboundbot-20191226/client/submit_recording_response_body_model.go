// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitRecordingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SubmitRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitRecordingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitRecordingResponseBody
	GetSuccess() *bool
}

type SubmitRecordingResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitRecordingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitRecordingResponseBody) SetCode(v string) *SubmitRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitRecordingResponseBody) SetHttpStatusCode(v int32) *SubmitRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitRecordingResponseBody) SetMessage(v string) *SubmitRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitRecordingResponseBody) SetRequestId(v string) *SubmitRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitRecordingResponseBody) SetSuccess(v bool) *SubmitRecordingResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}
