// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationStandardAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageTranslationStandardAsyncResponseBody
	GetCode() *string
	SetData(v *ImageTranslationStandardAsyncResponseBodyData) *ImageTranslationStandardAsyncResponseBody
	GetData() *ImageTranslationStandardAsyncResponseBodyData
	SetMessage(v string) *ImageTranslationStandardAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageTranslationStandardAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageTranslationStandardAsyncResponseBody
	GetSuccess() *bool
}

type ImageTranslationStandardAsyncResponseBody struct {
	// The error code. This parameter is not returned for normal calls.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asynchronous task submission result of Image Translation Standard Edition.
	Data *ImageTranslationStandardAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned for normal calls.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageTranslationStandardAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageTranslationStandardAsyncResponseBody) GetData() *ImageTranslationStandardAsyncResponseBodyData {
	return s.Data
}

func (s *ImageTranslationStandardAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageTranslationStandardAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageTranslationStandardAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageTranslationStandardAsyncResponseBody) SetCode(v string) *ImageTranslationStandardAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImageTranslationStandardAsyncResponseBody) SetData(v *ImageTranslationStandardAsyncResponseBodyData) *ImageTranslationStandardAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImageTranslationStandardAsyncResponseBody) SetMessage(v string) *ImageTranslationStandardAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImageTranslationStandardAsyncResponseBody) SetRequestId(v string) *ImageTranslationStandardAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageTranslationStandardAsyncResponseBody) SetSuccess(v bool) *ImageTranslationStandardAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ImageTranslationStandardAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationStandardAsyncResponseBodyData struct {
	// The task ID used to query the asynchronous task result.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImageTranslationStandardAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageTranslationStandardAsyncResponseBodyData) SetTaskId(v string) *ImageTranslationStandardAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageTranslationStandardAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
