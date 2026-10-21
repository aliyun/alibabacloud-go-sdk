// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageMattingAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageMattingAsyncResponseBody
	GetCode() *string
	SetData(v *ImageMattingAsyncResponseBodyData) *ImageMattingAsyncResponseBody
	GetData() *ImageMattingAsyncResponseBodyData
	SetMessage(v string) *ImageMattingAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageMattingAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageMattingAsyncResponseBody
	GetSuccess() *bool
}

type ImageMattingAsyncResponseBody struct {
	// The status code, such as "success". This parameter is not returned for normal calls.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result of the asynchronous task for intelligent image matting.
	Data *ImageMattingAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned for normal calls.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. It is used to identify a unique request.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageMattingAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageMattingAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImageMattingAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageMattingAsyncResponseBody) GetData() *ImageMattingAsyncResponseBodyData {
	return s.Data
}

func (s *ImageMattingAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageMattingAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageMattingAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageMattingAsyncResponseBody) SetCode(v string) *ImageMattingAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImageMattingAsyncResponseBody) SetData(v *ImageMattingAsyncResponseBodyData) *ImageMattingAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImageMattingAsyncResponseBody) SetMessage(v string) *ImageMattingAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImageMattingAsyncResponseBody) SetRequestId(v string) *ImageMattingAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageMattingAsyncResponseBody) SetSuccess(v bool) *ImageMattingAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ImageMattingAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageMattingAsyncResponseBodyData struct {
	// The task ID used to query the result of the asynchronous task.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImageMattingAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageMattingAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageMattingAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageMattingAsyncResponseBodyData) SetTaskId(v string) *ImageMattingAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageMattingAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
