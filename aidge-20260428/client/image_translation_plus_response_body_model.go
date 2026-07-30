// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationPlusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageTranslationPlusResponseBody
	GetCode() *string
	SetData(v *ImageTranslationPlusResponseBodyData) *ImageTranslationPlusResponseBody
	GetData() *ImageTranslationPlusResponseBodyData
	SetMessage(v string) *ImageTranslationPlusResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageTranslationPlusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageTranslationPlusResponseBody
	GetSuccess() *bool
}

type ImageTranslationPlusResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submit status of the asynchronous task.
	Data *ImageTranslationPlusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
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
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageTranslationPlusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationPlusResponseBody) GoString() string {
	return s.String()
}

func (s *ImageTranslationPlusResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageTranslationPlusResponseBody) GetData() *ImageTranslationPlusResponseBodyData {
	return s.Data
}

func (s *ImageTranslationPlusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageTranslationPlusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageTranslationPlusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageTranslationPlusResponseBody) SetCode(v string) *ImageTranslationPlusResponseBody {
	s.Code = &v
	return s
}

func (s *ImageTranslationPlusResponseBody) SetData(v *ImageTranslationPlusResponseBodyData) *ImageTranslationPlusResponseBody {
	s.Data = v
	return s
}

func (s *ImageTranslationPlusResponseBody) SetMessage(v string) *ImageTranslationPlusResponseBody {
	s.Message = &v
	return s
}

func (s *ImageTranslationPlusResponseBody) SetRequestId(v string) *ImageTranslationPlusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageTranslationPlusResponseBody) SetSuccess(v bool) *ImageTranslationPlusResponseBody {
	s.Success = &v
	return s
}

func (s *ImageTranslationPlusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationPlusResponseBodyData struct {
	// The asynchronous task ID. Use the queryTaskResult API to poll for results.
	//
	// example:
	//
	// task-abc123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImageTranslationPlusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationPlusResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageTranslationPlusResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageTranslationPlusResponseBodyData) SetTaskId(v string) *ImageTranslationPlusResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageTranslationPlusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
