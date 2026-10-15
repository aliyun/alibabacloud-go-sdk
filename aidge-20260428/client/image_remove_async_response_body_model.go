// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemoveAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageRemoveAsyncResponseBody
	GetCode() *string
	SetData(v *ImageRemoveAsyncResponseBodyData) *ImageRemoveAsyncResponseBody
	GetData() *ImageRemoveAsyncResponseBodyData
	SetMessage(v string) *ImageRemoveAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageRemoveAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageRemoveAsyncResponseBody
	GetSuccess() *bool
}

type ImageRemoveAsyncResponseBody struct {
	// The status code, such as "success".
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result struct.
	Data *ImageRemoveAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message, such as "The parameters contain sensitive information. Please try another input."
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to identify a unique request.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values: true (success) and false (failure).
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageRemoveAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImageRemoveAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageRemoveAsyncResponseBody) GetData() *ImageRemoveAsyncResponseBodyData {
	return s.Data
}

func (s *ImageRemoveAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageRemoveAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageRemoveAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageRemoveAsyncResponseBody) SetCode(v string) *ImageRemoveAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImageRemoveAsyncResponseBody) SetData(v *ImageRemoveAsyncResponseBodyData) *ImageRemoveAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImageRemoveAsyncResponseBody) SetMessage(v string) *ImageRemoveAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImageRemoveAsyncResponseBody) SetRequestId(v string) *ImageRemoveAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageRemoveAsyncResponseBody) SetSuccess(v bool) *ImageRemoveAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ImageRemoveAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageRemoveAsyncResponseBodyData struct {
	// The task ID used to query the result of the asynchronous task.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImageRemoveAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageRemoveAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageRemoveAsyncResponseBodyData) SetTaskId(v string) *ImageRemoveAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageRemoveAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
