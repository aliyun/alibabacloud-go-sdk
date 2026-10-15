// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationPlusAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageTranslationPlusAsyncResponseBody
	GetCode() *string
	SetData(v *ImageTranslationPlusAsyncResponseBodyData) *ImageTranslationPlusAsyncResponseBody
	GetData() *ImageTranslationPlusAsyncResponseBodyData
	SetMessage(v string) *ImageTranslationPlusAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageTranslationPlusAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageTranslationPlusAsyncResponseBody
	GetSuccess() *bool
}

type ImageTranslationPlusAsyncResponseBody struct {
	// The response code of the submission call. A value of success indicates a successful call.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result of the Image Translation Plus asynchronous task.
	Data *ImageTranslationPlusAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message, such as \\"The content contains sensitive data. Try another input.\\"
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which is used for troubleshooting and Tracing Analysis.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the submission call is successful. This does not mean that the translation is completed. A value of true indicates success, and a value of false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageTranslationPlusAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationPlusAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImageTranslationPlusAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageTranslationPlusAsyncResponseBody) GetData() *ImageTranslationPlusAsyncResponseBodyData {
	return s.Data
}

func (s *ImageTranslationPlusAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageTranslationPlusAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageTranslationPlusAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageTranslationPlusAsyncResponseBody) SetCode(v string) *ImageTranslationPlusAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImageTranslationPlusAsyncResponseBody) SetData(v *ImageTranslationPlusAsyncResponseBodyData) *ImageTranslationPlusAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImageTranslationPlusAsyncResponseBody) SetMessage(v string) *ImageTranslationPlusAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImageTranslationPlusAsyncResponseBody) SetRequestId(v string) *ImageTranslationPlusAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageTranslationPlusAsyncResponseBody) SetSuccess(v bool) *ImageTranslationPlusAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ImageTranslationPlusAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationPlusAsyncResponseBodyData struct {
	// The identifier of the asynchronous task, which is used for result query and notification association. Do not use the example prefix as a business validation rule.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImageTranslationPlusAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationPlusAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageTranslationPlusAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageTranslationPlusAsyncResponseBodyData) SetTaskId(v string) *ImageTranslationPlusAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageTranslationPlusAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
