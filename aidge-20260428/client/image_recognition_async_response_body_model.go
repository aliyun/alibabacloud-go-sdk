// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRecognitionAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageRecognitionAsyncResponseBody
	GetCode() *string
	SetData(v *ImageRecognitionAsyncResponseBodyData) *ImageRecognitionAsyncResponseBody
	GetData() *ImageRecognitionAsyncResponseBodyData
	SetMessage(v string) *ImageRecognitionAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageRecognitionAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageRecognitionAsyncResponseBody
	GetSuccess() *bool
}

type ImageRecognitionAsyncResponseBody struct {
	// The error code. This parameter is not returned for normal calls.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct information of the returned recognition result.
	Data *ImageRecognitionAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message, such as "The parameter contains sensitive information. Try another input."
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to identify a unique request.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: successful
	//
	// - false: failed
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageRecognitionAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImageRecognitionAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageRecognitionAsyncResponseBody) GetData() *ImageRecognitionAsyncResponseBodyData {
	return s.Data
}

func (s *ImageRecognitionAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageRecognitionAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageRecognitionAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageRecognitionAsyncResponseBody) SetCode(v string) *ImageRecognitionAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImageRecognitionAsyncResponseBody) SetData(v *ImageRecognitionAsyncResponseBodyData) *ImageRecognitionAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImageRecognitionAsyncResponseBody) SetMessage(v string) *ImageRecognitionAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImageRecognitionAsyncResponseBody) SetRequestId(v string) *ImageRecognitionAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageRecognitionAsyncResponseBody) SetSuccess(v bool) *ImageRecognitionAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ImageRecognitionAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageRecognitionAsyncResponseBodyData struct {
	// The task ID used to query the result of the asynchronous task.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImageRecognitionAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageRecognitionAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageRecognitionAsyncResponseBodyData) SetTaskId(v string) *ImageRecognitionAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageRecognitionAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
