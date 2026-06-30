// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAudioDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadAudioDataResponseBody
	GetCode() *string
	SetData(v string) *UploadAudioDataResponseBody
	GetData() *string
	SetMessage(v string) *UploadAudioDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadAudioDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadAudioDataResponseBody
	GetSuccess() *bool
}

type UploadAudioDataResponseBody struct {
	// The result code. A value of 200 indicates success. Other values indicate failure. You can use this field to determine the cause of a failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task ID. Use this ID to get the task result.
	//
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5***
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error details if the request fails. If the request is successful, the value is \\`successful\\`.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. This is the unique identifier of the request.
	//
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F8***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. You can use this field to determine if the request succeeded. A value of \\`true\\` indicates success. A value of \\`false\\` or \\`null\\` indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadAudioDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadAudioDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAudioDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadAudioDataResponseBody) GetData() *string {
	return s.Data
}

func (s *UploadAudioDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadAudioDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadAudioDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadAudioDataResponseBody) SetCode(v string) *UploadAudioDataResponseBody {
	s.Code = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetData(v string) *UploadAudioDataResponseBody {
	s.Data = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetMessage(v string) *UploadAudioDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetRequestId(v string) *UploadAudioDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetSuccess(v bool) *UploadAudioDataResponseBody {
	s.Success = &v
	return s
}

func (s *UploadAudioDataResponseBody) Validate() error {
	return dara.Validate(s)
}
