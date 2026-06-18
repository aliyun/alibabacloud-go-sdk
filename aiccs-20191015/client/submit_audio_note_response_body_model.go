// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAudioNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SubmitAudioNoteResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SubmitAudioNoteResponseBody
	GetCode() *string
	SetData(v string) *SubmitAudioNoteResponseBody
	GetData() *string
	SetMessage(v string) *SubmitAudioNoteResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitAudioNoteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAudioNoteResponseBody
	GetSuccess() *bool
}

type SubmitAudioNoteResponseBody struct {
	// The detailed reason why access is denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The recording notes task ID returned after successful submission. This value corresponds to the CallId in the subsequent asynchronous notes result callback.
	//
	// example:
	//
	// 100000000000000001_100000000000000002
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAudioNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAudioNoteResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAudioNoteResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SubmitAudioNoteResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitAudioNoteResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitAudioNoteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitAudioNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAudioNoteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAudioNoteResponseBody) SetAccessDeniedDetail(v string) *SubmitAudioNoteResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SubmitAudioNoteResponseBody) SetCode(v string) *SubmitAudioNoteResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAudioNoteResponseBody) SetData(v string) *SubmitAudioNoteResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitAudioNoteResponseBody) SetMessage(v string) *SubmitAudioNoteResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAudioNoteResponseBody) SetRequestId(v string) *SubmitAudioNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAudioNoteResponseBody) SetSuccess(v bool) *SubmitAudioNoteResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAudioNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
