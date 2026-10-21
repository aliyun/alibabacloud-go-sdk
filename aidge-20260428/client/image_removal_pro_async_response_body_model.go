// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemovalProAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageRemovalProAsyncResponseBody
	GetCode() *string
	SetData(v *ImageRemovalProAsyncResponseBodyData) *ImageRemovalProAsyncResponseBody
	GetData() *ImageRemovalProAsyncResponseBodyData
	SetMessage(v string) *ImageRemovalProAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageRemovalProAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageRemovalProAsyncResponseBody
	GetSuccess() *bool
}

type ImageRemovalProAsyncResponseBody struct {
	// The error code. This parameter is not returned for successful calls.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result of the Intelligent Erasure Pro asynchronous task.
	Data *ImageRemovalProAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned for successful calls. Example: "The parameters contain sensitive information. Try other inputs."
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. This ID is used to uniquely identify the request.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: successful.
	//
	// - false: failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageRemovalProAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImageRemovalProAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageRemovalProAsyncResponseBody) GetData() *ImageRemovalProAsyncResponseBodyData {
	return s.Data
}

func (s *ImageRemovalProAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageRemovalProAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageRemovalProAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageRemovalProAsyncResponseBody) SetCode(v string) *ImageRemovalProAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImageRemovalProAsyncResponseBody) SetData(v *ImageRemovalProAsyncResponseBodyData) *ImageRemovalProAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImageRemovalProAsyncResponseBody) SetMessage(v string) *ImageRemovalProAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImageRemovalProAsyncResponseBody) SetRequestId(v string) *ImageRemovalProAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageRemovalProAsyncResponseBody) SetSuccess(v bool) *ImageRemovalProAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *ImageRemovalProAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageRemovalProAsyncResponseBodyData struct {
	// The ID of the task. You can use this ID to query the asynchronous task result.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImageRemovalProAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageRemovalProAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageRemovalProAsyncResponseBodyData) SetTaskId(v string) *ImageRemovalProAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageRemovalProAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
