// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationProAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageTranslationProAsyncResponseBody
	GetCode() *string
	SetData(v *ImageTranslationProAsyncResponseBodyData) *ImageTranslationProAsyncResponseBody
	GetData() *ImageTranslationProAsyncResponseBodyData
	SetMessage(v string) *ImageTranslationProAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageTranslationProAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageTranslationProAsyncResponseBody
	GetSuccess() *bool
}

type ImageTranslationProAsyncResponseBody struct {
	// The response code. A value of success indicates that the task is submitted successfully.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result of the Image Translation Pro asynchronous task.
	Data *ImageTranslationProAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The call description. A successful call returns Success. If the call fails, read the actual error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the submission request. It is used for troubleshooting.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the submission call is successful. This does not indicate that the translation is complete.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageTranslationProAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImageTranslationProAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageTranslationProAsyncResponseBody) GetData() *ImageTranslationProAsyncResponseBodyData {
	return s.Data
}

func (s *ImageTranslationProAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageTranslationProAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageTranslationProAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageTranslationProAsyncResponseBody) SetCode(v string) *ImageTranslationProAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImageTranslationProAsyncResponseBody) SetData(v *ImageTranslationProAsyncResponseBodyData) *ImageTranslationProAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImageTranslationProAsyncResponseBody) SetMessage(v string) *ImageTranslationProAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImageTranslationProAsyncResponseBody) SetRequestId(v string) *ImageTranslationProAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageTranslationProAsyncResponseBody) SetSuccess(v bool) *ImageTranslationProAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ImageTranslationProAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationProAsyncResponseBodyData struct {
	// The ID of the asynchronous task. It is used to query results and associate notifications. Do not use the example prefix as a business validation rule.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImageTranslationProAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageTranslationProAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageTranslationProAsyncResponseBodyData) SetTaskId(v string) *ImageTranslationProAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageTranslationProAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
